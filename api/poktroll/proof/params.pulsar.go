// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package proof

import (
	_ "cosmossdk.io/api/amino"
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	binary "encoding/binary"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	math "math"
	reflect "reflect"
	sync "sync"
)

var (
	md_Params                              protoreflect.MessageDescriptor
	fd_Params_relay_difficulty_target_hash protoreflect.FieldDescriptor
	fd_Params_proof_request_probability    protoreflect.FieldDescriptor
	fd_Params_proof_requirement_threshold  protoreflect.FieldDescriptor
	fd_Params_proof_missing_penalty        protoreflect.FieldDescriptor
)

func init() {
	file_poktroll_proof_params_proto_init()
	md_Params = File_poktroll_proof_params_proto.Messages().ByName("Params")
	fd_Params_relay_difficulty_target_hash = md_Params.Fields().ByName("relay_difficulty_target_hash")
	fd_Params_proof_request_probability = md_Params.Fields().ByName("proof_request_probability")
	fd_Params_proof_requirement_threshold = md_Params.Fields().ByName("proof_requirement_threshold")
	fd_Params_proof_missing_penalty = md_Params.Fields().ByName("proof_missing_penalty")
}

var _ protoreflect.Message = (*fastReflection_Params)(nil)

type fastReflection_Params Params

func (x *Params) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Params)(x)
}

func (x *Params) slowProtoReflect() protoreflect.Message {
	mi := &file_poktroll_proof_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Params_messageType fastReflection_Params_messageType
var _ protoreflect.MessageType = fastReflection_Params_messageType{}

type fastReflection_Params_messageType struct{}

func (x fastReflection_Params_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Params)(nil)
}
func (x fastReflection_Params_messageType) New() protoreflect.Message {
	return new(fastReflection_Params)
}
func (x fastReflection_Params_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Params) Descriptor() protoreflect.MessageDescriptor {
	return md_Params
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Params) Type() protoreflect.MessageType {
	return _fastReflection_Params_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Params) New() protoreflect.Message {
	return new(fastReflection_Params)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Params) Interface() protoreflect.ProtoMessage {
	return (*Params)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Params) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.RelayDifficultyTargetHash) != 0 {
		value := protoreflect.ValueOfBytes(x.RelayDifficultyTargetHash)
		if !f(fd_Params_relay_difficulty_target_hash, value) {
			return
		}
	}
	if x.ProofRequestProbability != float32(0) || math.Signbit(float64(x.ProofRequestProbability)) {
		value := protoreflect.ValueOfFloat32(x.ProofRequestProbability)
		if !f(fd_Params_proof_request_probability, value) {
			return
		}
	}
	if x.ProofRequirementThreshold != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProofRequirementThreshold)
		if !f(fd_Params_proof_requirement_threshold, value) {
			return
		}
	}
	if x.ProofMissingPenalty != nil {
		value := protoreflect.ValueOfMessage(x.ProofMissingPenalty.ProtoReflect())
		if !f(fd_Params_proof_missing_penalty, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Params) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		return len(x.RelayDifficultyTargetHash) != 0
	case "poktroll.proof.Params.proof_request_probability":
		return x.ProofRequestProbability != float32(0) || math.Signbit(float64(x.ProofRequestProbability))
	case "poktroll.proof.Params.proof_requirement_threshold":
		return x.ProofRequirementThreshold != uint64(0)
	case "poktroll.proof.Params.proof_missing_penalty":
		return x.ProofMissingPenalty != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		x.RelayDifficultyTargetHash = nil
	case "poktroll.proof.Params.proof_request_probability":
		x.ProofRequestProbability = float32(0)
	case "poktroll.proof.Params.proof_requirement_threshold":
		x.ProofRequirementThreshold = uint64(0)
	case "poktroll.proof.Params.proof_missing_penalty":
		x.ProofMissingPenalty = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Params) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		value := x.RelayDifficultyTargetHash
		return protoreflect.ValueOfBytes(value)
	case "poktroll.proof.Params.proof_request_probability":
		value := x.ProofRequestProbability
		return protoreflect.ValueOfFloat32(value)
	case "poktroll.proof.Params.proof_requirement_threshold":
		value := x.ProofRequirementThreshold
		return protoreflect.ValueOfUint64(value)
	case "poktroll.proof.Params.proof_missing_penalty":
		value := x.ProofMissingPenalty
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		x.RelayDifficultyTargetHash = value.Bytes()
	case "poktroll.proof.Params.proof_request_probability":
		x.ProofRequestProbability = float32(value.Float())
	case "poktroll.proof.Params.proof_requirement_threshold":
		x.ProofRequirementThreshold = value.Uint()
	case "poktroll.proof.Params.proof_missing_penalty":
		x.ProofMissingPenalty = value.Message().Interface().(*v1beta1.Coin)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.proof.Params.proof_missing_penalty":
		if x.ProofMissingPenalty == nil {
			x.ProofMissingPenalty = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.ProofMissingPenalty.ProtoReflect())
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		panic(fmt.Errorf("field relay_difficulty_target_hash of message poktroll.proof.Params is not mutable"))
	case "poktroll.proof.Params.proof_request_probability":
		panic(fmt.Errorf("field proof_request_probability of message poktroll.proof.Params is not mutable"))
	case "poktroll.proof.Params.proof_requirement_threshold":
		panic(fmt.Errorf("field proof_requirement_threshold of message poktroll.proof.Params is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Params) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.proof.Params.relay_difficulty_target_hash":
		return protoreflect.ValueOfBytes(nil)
	case "poktroll.proof.Params.proof_request_probability":
		return protoreflect.ValueOfFloat32(float32(0))
	case "poktroll.proof.Params.proof_requirement_threshold":
		return protoreflect.ValueOfUint64(uint64(0))
	case "poktroll.proof.Params.proof_missing_penalty":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.proof.Params"))
		}
		panic(fmt.Errorf("message poktroll.proof.Params does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Params) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in poktroll.proof.Params", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Params) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Params) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Params) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Params) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.RelayDifficultyTargetHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.ProofRequestProbability != 0 || math.Signbit(float64(x.ProofRequestProbability)) {
			n += 5
		}
		if x.ProofRequirementThreshold != 0 {
			n += 1 + runtime.Sov(uint64(x.ProofRequirementThreshold))
		}
		if x.ProofMissingPenalty != nil {
			l = options.Size(x.ProofMissingPenalty)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.ProofMissingPenalty != nil {
			encoded, err := options.Marshal(x.ProofMissingPenalty)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.ProofRequirementThreshold != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProofRequirementThreshold))
			i--
			dAtA[i] = 0x18
		}
		if x.ProofRequestProbability != 0 || math.Signbit(float64(x.ProofRequestProbability)) {
			i -= 4
			binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(x.ProofRequestProbability))))
			i--
			dAtA[i] = 0x15
		}
		if len(x.RelayDifficultyTargetHash) > 0 {
			i -= len(x.RelayDifficultyTargetHash)
			copy(dAtA[i:], x.RelayDifficultyTargetHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RelayDifficultyTargetHash)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Params)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RelayDifficultyTargetHash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RelayDifficultyTargetHash = append(x.RelayDifficultyTargetHash[:0], dAtA[iNdEx:postIndex]...)
				if x.RelayDifficultyTargetHash == nil {
					x.RelayDifficultyTargetHash = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 5 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProofRequestProbability", wireType)
				}
				var v uint32
				if (iNdEx + 4) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
				iNdEx += 4
				x.ProofRequestProbability = float32(math.Float32frombits(v))
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProofRequirementThreshold", wireType)
				}
				x.ProofRequirementThreshold = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProofRequirementThreshold |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProofMissingPenalty", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.ProofMissingPenalty == nil {
					x.ProofMissingPenalty = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.ProofMissingPenalty); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: poktroll/proof/params.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Params defines the parameters for the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO_FOLLOWUP(@olshansk, #690): Either delete this or change it to be named "minimum"
	// relay_difficulty_target_hash is the maximum value a relay hash must be less than to be volume/reward applicable.
	RelayDifficultyTargetHash []byte `protobuf:"bytes,1,opt,name=relay_difficulty_target_hash,json=relayDifficultyTargetHash,proto3" json:"relay_difficulty_target_hash,omitempty"`
	// proof_request_probability is the probability of a session requiring a proof
	// if it's cost (i.e. compute unit consumption) is below the ProofRequirementThreshold.
	ProofRequestProbability float32 `protobuf:"fixed32,2,opt,name=proof_request_probability,json=proofRequestProbability,proto3" json:"proof_request_probability,omitempty"`
	// proof_requirement_threshold is the session cost (i.e. compute unit consumption)
	// threshold which asserts that a session MUST have a corresponding proof when its cost
	// is equal to or above the threshold. This is in contrast to the this requirement
	// being determined probabilistically via ProofRequestProbability.
	//
	// TODO_MAINNET: Consider renaming this to `proof_requirement_threshold_compute_units`.
	ProofRequirementThreshold uint64 `protobuf:"varint,3,opt,name=proof_requirement_threshold,json=proofRequirementThreshold,proto3" json:"proof_requirement_threshold,omitempty"`
	// proof_missing_penalty is the number of tokens (uPOKT) which should be slashed from a supplier
	// when a proof is required (either via proof_requirement_threshold or proof_missing_penalty)
	// but is not provided.
	ProofMissingPenalty *v1beta1.Coin `protobuf:"bytes,4,opt,name=proof_missing_penalty,json=proofMissingPenalty,proto3" json:"proof_missing_penalty,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poktroll_proof_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_poktroll_proof_params_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetRelayDifficultyTargetHash() []byte {
	if x != nil {
		return x.RelayDifficultyTargetHash
	}
	return nil
}

func (x *Params) GetProofRequestProbability() float32 {
	if x != nil {
		return x.ProofRequestProbability
	}
	return 0
}

func (x *Params) GetProofRequirementThreshold() uint64 {
	if x != nil {
		return x.ProofRequirementThreshold
	}
	return 0
}

func (x *Params) GetProofMissingPenalty() *v1beta1.Coin {
	if x != nil {
		return x.ProofMissingPenalty
	}
	return nil
}

var File_poktroll_proof_params_proto protoreflect.FileDescriptor

var file_poktroll_proof_params_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70,
	0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x1a, 0x11, 0x61,
	0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x03, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x61, 0x0a, 0x1c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x20, 0xea, 0xde, 0x1f, 0x1c, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x52, 0x19, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x59, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42, 0x1d, 0xea, 0xde, 0x1f, 0x19, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x5f, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x1f, 0xea, 0xde, 0x1f, 0x1b, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x68, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x19, 0xea, 0xde, 0x1f, 0x15,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x50, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x3a, 0x20, 0xe8, 0xa0, 0x1f, 0x01,
	0x8a, 0xe7, 0xb0, 0x2a, 0x17, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x9f, 0x01, 0xd8,
	0xe2, 0x1e, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64,
	0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa, 0x02, 0x0e,
	0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0xca, 0x02,
	0x0e, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0xe2,
	0x02, 0x1a, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x50,
	0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poktroll_proof_params_proto_rawDescOnce sync.Once
	file_poktroll_proof_params_proto_rawDescData = file_poktroll_proof_params_proto_rawDesc
)

func file_poktroll_proof_params_proto_rawDescGZIP() []byte {
	file_poktroll_proof_params_proto_rawDescOnce.Do(func() {
		file_poktroll_proof_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_poktroll_proof_params_proto_rawDescData)
	})
	return file_poktroll_proof_params_proto_rawDescData
}

var file_poktroll_proof_params_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_poktroll_proof_params_proto_goTypes = []interface{}{
	(*Params)(nil),       // 0: poktroll.proof.Params
	(*v1beta1.Coin)(nil), // 1: cosmos.base.v1beta1.Coin
}
var file_poktroll_proof_params_proto_depIdxs = []int32{
	1, // 0: poktroll.proof.Params.proof_missing_penalty:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_poktroll_proof_params_proto_init() }
func file_poktroll_proof_params_proto_init() {
	if File_poktroll_proof_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_poktroll_proof_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poktroll_proof_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poktroll_proof_params_proto_goTypes,
		DependencyIndexes: file_poktroll_proof_params_proto_depIdxs,
		MessageInfos:      file_poktroll_proof_params_proto_msgTypes,
	}.Build()
	File_poktroll_proof_params_proto = out.File
	file_poktroll_proof_params_proto_rawDesc = nil
	file_poktroll_proof_params_proto_goTypes = nil
	file_poktroll_proof_params_proto_depIdxs = nil
}
