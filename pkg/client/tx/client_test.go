package tx_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"cosmossdk.io/depinject"
	cometbytes "github.com/cometbft/cometbft/libs/bytes"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/internal/mocks/mockclient"
	"github.com/pokt-network/poktroll/internal/testclient"
	"github.com/pokt-network/poktroll/internal/testclient/testblock"
	"github.com/pokt-network/poktroll/internal/testclient/testeventsquery"
	"github.com/pokt-network/poktroll/internal/testclient/testtx"
	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/tx"
	"github.com/pokt-network/poktroll/pkg/either"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
)

const (
	testSigningKeyName = "test_signer"
	// NB: testServiceIdPrefix must not be longer than 7 characters due to
	// maxServiceIdLen.
	testServiceIdPrefix = "testsvc"
	txCommitTimeout     = 10 * time.Millisecond
)

// TODO_TECHDEBT: add coverage for the transactions client handling an events bytes error either.

func TestTxClient_SignAndBroadcast_Succeeds(t *testing.T) {
	var (
		// expectedTx is the expected transactions bytes that will be signed and broadcast
		// by the transaction client. It is computed and assigned in the
		// testtx.NewOneTimeTxTxContext helper function. The same reference needs
		// to be used across the expectations that are set on the transactions context mock.
		expectedTx cometbytes.HexBytes
		// eventsBzPublishCh is the channel that the mock events query client
		// will use to publish the transactions event bytes. It is used near the end of
		// the test to mock the network signaling that the transactions was committed.
		eventsBzPublishCh chan<- either.Bytes
		blocksPublishCh   chan client.Block
		ctx               = context.Background()
	)

	keyring, signingKey := newTestKeyringWithKey(t)

	eventsQueryClient := testeventsquery.NewOneTimeTxEventsQueryClient(
		ctx, t, signingKey, &eventsBzPublishCh,
	)

	txCtxMock := testtx.NewOneTimeTxTxContext(
		t, keyring,
		testSigningKeyName,
		&expectedTx,
	)

	// Construct a new mock block client because it is a required dependency.
	// Since we're not exercising transactions timeouts in this test, we don't need to
	// set any particular expectations on it, nor do we care about the contents
	// of the latest block.
	blockClientMock := testblock.NewOneTimeCommittedBlocksSequenceBlockClient(
		t, blocksPublishCh,
	)

	// Construct a new depinject config with the mocks we created above.
	txClientDeps := depinject.Supply(
		eventsQueryClient,
		txCtxMock,
		blockClientMock,
	)

	// Construct the transaction client.
	txClient, err := tx.NewTxClient(
		ctx, txClientDeps, tx.WithSigningKeyName(testSigningKeyName),
	)
	require.NoError(t, err)

	signingKeyAddr, err := signingKey.GetAddress()
	require.NoError(t, err)

	// Construct a valid (arbitrary) message to sign, encode, and broadcast.
	appStake := types.NewCoin("upokt", types.NewInt(1000000))
	appStakeMsg := &apptypes.MsgStakeApplication{
		Address:  signingKeyAddr.String(),
		Stake:    &appStake,
		Services: client.NewTestApplicationServiceConfig(testServiceIdPrefix, 2),
	}

	// Sign and broadcast the message.
	eitherErr := txClient.SignAndBroadcast(ctx, appStakeMsg)
	err, errCh := eitherErr.SyncOrAsyncError()
	require.NoError(t, err)

	// Construct the expected transaction event bytes from the expected transaction bytes.
	txEventBz, err := json.Marshal(&tx.TxEvent{Tx: expectedTx})
	require.NoError(t, err)

	// Publish the transaction event bytes to the events query client so that the transaction client
	// registers the transactions as committed (i.e. removes it from the timeout pool).
	eventsBzPublishCh <- either.Success[[]byte](txEventBz)

	// Assert that the error channel was closed without receiving.
	select {
	case err, ok := <-errCh:
		require.NoError(t, err)
		require.Falsef(t, ok, "expected errCh to be closed")
	case <-time.After(txCommitTimeout):
		t.Fatal("test timed out waiting for errCh to receive")
	}
}

func TestTxClient_NewTxClient_Error(t *testing.T) {
	// Construct an empty in-memory keyring.
	keyring := cosmoskeyring.NewInMemory(testclient.EncodingConfig.Marshaler)

	tests := []struct {
		name           string
		signingKeyName string
		expectedErr    error
	}{
		{
			name:           "empty signing key name",
			signingKeyName: "",
			expectedErr:    tx.ErrEmptySigningKeyName,
		},
		{
			name:           "signing key does not exist",
			signingKeyName: "nonexistent",
			expectedErr:    tx.ErrNoSuchSigningKey,
		},
		// TODO_TECHDEBT: add coverage for this error case
		// {
		// 	name:        "failed to get address",
		//   testSigningKeyName: "incompatible",
		// 	expectedErr: tx.ErrSigningKeyAddr,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				ctrl = gomock.NewController(t)
				ctx  = context.Background()
			)

			// Construct a new mock events query client. Since we expect the
			// NewTxClient call to fail, we don't need to set any expectations
			// on this mock.
			eventsQueryClient := mockclient.NewMockEventsQueryClient(ctrl)

			// Construct a new mock transactions context.
			txCtxMock, _ := testtx.NewAnyTimesTxTxContext(t, keyring)

			// Construct a new mock block client. Since we expect the NewTxClient
			// call to fail, we don't need to set any expectations on this mock.
			blockClientMock := mockclient.NewMockBlockClient(ctrl)

			// Construct a new depinject config with the mocks we created above.
			txClientDeps := depinject.Supply(
				eventsQueryClient,
				txCtxMock,
				blockClientMock,
			)

			// Construct a signing key option using the test signing key name.
			signingKeyOpt := tx.WithSigningKeyName(tt.signingKeyName)

			// Attempt to create the transactions client.
			txClient, err := tx.NewTxClient(ctx, txClientDeps, signingKeyOpt)
			require.ErrorIs(t, err, tt.expectedErr)
			require.Nil(t, txClient)
		})
	}
}

func TestTxClient_SignAndBroadcast_SyncError(t *testing.T) {
	var (
		// eventsBzPublishCh is the channel that the mock events query client
		// will use to publish the transactions event bytes. It is not used in
		// this test but is required to use the NewOneTimeTxEventsQueryClient
		// helper.
		eventsBzPublishCh chan<- either.Bytes
		// blocksPublishCh is the channel that the mock block client will use
		// to publish the latest block. It is not used in this test but is
		// required to use the NewOneTimeCommittedBlocksSequenceBlockClient
		// helper.
		blocksPublishCh chan client.Block
		ctx             = context.Background()
	)

	keyring, signingKey := newTestKeyringWithKey(t)

	// Construct a new mock events query client. Since we expect the
	// NewTxClient call to fail, we don't need to set any expectations
	// on this mock.
	eventsQueryClient := testeventsquery.NewOneTimeTxEventsQueryClient(
		ctx, t, signingKey, &eventsBzPublishCh,
	)

	// Construct a new mock transaction context.
	txCtxMock, _ := testtx.NewAnyTimesTxTxContext(t, keyring)

	// Construct a new mock block client because it is a required dependency.
	// Since we're not exercising transactions timeouts in this test, we don't need to
	// set any particular expectations on it, nor do we care about the contents
	// of the latest block.
	blockClientMock := testblock.NewOneTimeCommittedBlocksSequenceBlockClient(
		t, blocksPublishCh,
	)

	// Construct a new depinject config with the mocks we created above.
	txClientDeps := depinject.Supply(
		eventsQueryClient,
		txCtxMock,
		blockClientMock,
	)

	// Construct the transaction client.
	txClient, err := tx.NewTxClient(
		ctx, txClientDeps, tx.WithSigningKeyName(testSigningKeyName),
	)
	require.NoError(t, err)

	// Construct an invalid (arbitrary) message to sign, encode, and broadcast.
	signingAddr, err := signingKey.GetAddress()
	require.NoError(t, err)
	appStakeMsg := &apptypes.MsgStakeApplication{
		// Providing address to avoid panic from #GetSigners().
		Address: signingAddr.String(),
		Stake:   nil,
		// NB: explicitly omitting required fields
	}

	eitherErr := txClient.SignAndBroadcast(ctx, appStakeMsg)
	err, _ = eitherErr.SyncOrAsyncError()
	require.ErrorIs(t, err, tx.ErrInvalidMsg)

	time.Sleep(10 * time.Millisecond)
}

// TODO_INCOMPLETE: add coverage for async error; i.e. insufficient gas or on-chain error
func TestTxClient_SignAndBroadcast_CheckTxError(t *testing.T) {
	var (
		// expectedErrMsg is the expected error message that will be returned
		// by the transaction client. It is computed and assigned in the
		// testtx.NewOneTimeErrCheckTxTxContext helper function.
		expectedErrMsg string
		// eventsBzPublishCh is the channel that the mock events query client
		// will use to publish the transactions event bytes. It is used near the end of
		// the test to mock the network signaling that the transactions was committed.
		eventsBzPublishCh chan<- either.Bytes
		blocksPublishCh   chan client.Block
		ctx               = context.Background()
	)

	keyring, signingKey := newTestKeyringWithKey(t)

	eventsQueryClient := testeventsquery.NewOneTimeTxEventsQueryClient(
		ctx, t, signingKey, &eventsBzPublishCh,
	)

	txCtxMock := testtx.NewOneTimeErrCheckTxTxContext(
		t, keyring,
		testSigningKeyName,
		&expectedErrMsg,
	)

	// Construct a new mock block client because it is a required dependency.
	// Since we're not exercising transactions timeouts in this test, we don't need to
	// set any particular expectations on it, nor do we care about the contents
	// of the latest block.
	blockClientMock := testblock.NewOneTimeCommittedBlocksSequenceBlockClient(
		t, blocksPublishCh,
	)

	// Construct a new depinject config with the mocks we created above.
	txClientDeps := depinject.Supply(
		eventsQueryClient,
		txCtxMock,
		blockClientMock,
	)

	// Construct the transaction client.
	txClient, err := tx.NewTxClient(ctx, txClientDeps, tx.WithSigningKeyName(testSigningKeyName))
	require.NoError(t, err)

	signingKeyAddr, err := signingKey.GetAddress()
	require.NoError(t, err)

	// Construct a valid (arbitrary) message to sign, encode, and broadcast.
	appStake := types.NewCoin("upokt", types.NewInt(1000000))
	appStakeMsg := &apptypes.MsgStakeApplication{
		Address:  signingKeyAddr.String(),
		Stake:    &appStake,
		Services: client.NewTestApplicationServiceConfig(testServiceIdPrefix, 2),
	}

	// Sign and broadcast the message.
	eitherErr := txClient.SignAndBroadcast(ctx, appStakeMsg)
	err, _ = eitherErr.SyncOrAsyncError()
	require.ErrorIs(t, err, tx.ErrCheckTx)
	require.ErrorContains(t, err, expectedErrMsg)
}

func TestTxClient_SignAndBroadcast_Timeout(t *testing.T) {
	var (
		// expectedErrMsg is the expected error message that will be returned
		// by the transaction client. It is computed and assigned in the
		// testtx.NewOneTimeErrCheckTxTxContext helper function.
		expectedErrMsg string
		// eventsBzPublishCh is the channel that the mock events query client
		// will use to publish the transaction event bytes. It is used near the end of
		// the test to mock the network signaling that the transaction was committed.
		eventsBzPublishCh chan<- either.Bytes
		blocksPublishCh   = make(chan client.Block, tx.DefaultCommitTimeoutHeightOffset)
		ctx               = context.Background()
	)

	keyring, signingKey := newTestKeyringWithKey(t)

	eventsQueryClient := testeventsquery.NewOneTimeTxEventsQueryClient(
		ctx, t, signingKey, &eventsBzPublishCh,
	)

	txCtxMock := testtx.NewOneTimeErrTxTimeoutTxContext(
		t, keyring,
		testSigningKeyName,
		&expectedErrMsg,
	)

	// Construct a new mock block client because it is a required dependency.
	// Since we're not exercising transaction timeouts in this test, we don't need to
	// set any particular expectations on it, nor do we care about the contents
	// of the latest block.
	blockClientMock := testblock.NewOneTimeCommittedBlocksSequenceBlockClient(
		t, blocksPublishCh,
	)

	// Construct a new depinject config with the mocks we created above.
	txClientDeps := depinject.Supply(
		eventsQueryClient,
		txCtxMock,
		blockClientMock,
	)

	// Construct the transaction client.
	txClient, err := tx.NewTxClient(
		ctx, txClientDeps, tx.WithSigningKeyName(testSigningKeyName),
	)
	require.NoError(t, err)

	signingKeyAddr, err := signingKey.GetAddress()
	require.NoError(t, err)

	// Construct a valid (arbitrary) message to sign, encode, and broadcast.
	appStake := types.NewCoin("upokt", types.NewInt(1000000))
	appStakeMsg := &apptypes.MsgStakeApplication{
		Address:  signingKeyAddr.String(),
		Stake:    &appStake,
		Services: client.NewTestApplicationServiceConfig(testServiceIdPrefix, 2),
	}

	// Sign and broadcast the message in a transaction.
	eitherErr := txClient.SignAndBroadcast(ctx, appStakeMsg)
	err, errCh := eitherErr.SyncOrAsyncError()
	require.NoError(t, err)

	for i := 0; i < tx.DefaultCommitTimeoutHeightOffset; i++ {
		blocksPublishCh <- testblock.NewAnyTimesBlock(t, []byte{}, int64(i+1))
	}

	// Assert that we receive the expected error type & message.
	select {
	case err := <-errCh:
		require.ErrorIs(t, err, tx.ErrTxTimeout)
		require.ErrorContains(t, err, expectedErrMsg)
	// NB: wait 110% of txCommitTimeout; a bit longer than strictly necessary in
	// order to mitigate flakiness.
	case <-time.After(txCommitTimeout * 110 / 100):
		t.Fatal("test timed out waiting for errCh to receive")
	}

	// Assert that the error channel was closed.
	select {
	case err, ok := <-errCh:
		require.Falsef(t, ok, "expected errCh to be closed")
		require.NoError(t, err)
	// NB: Give the error channel some time to be ready to receive in order to
	// mitigate flakiness.
	case <-time.After(50 * time.Millisecond):
		t.Fatal("expected errCh to be closed")
	}
}

// TODO_TECHDEBT: add coverage for sending multiple messages simultaneously
func TestTxClient_SignAndBroadcast_MultipleMsgs(t *testing.T) {
	t.SkipNow()
}

// newTestKeyringWithKey creates a new in-memory keyring with a test key
// with testSigningKeyName as its name.
func newTestKeyringWithKey(t *testing.T) (cosmoskeyring.Keyring, *cosmoskeyring.Record) {
	keyring := cosmoskeyring.NewInMemory(testclient.EncodingConfig.Marshaler)
	key, _ := testclient.NewKey(t, testSigningKeyName, keyring)
	return keyring, key
}