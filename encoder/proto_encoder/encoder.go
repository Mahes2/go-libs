package encoder

import (
	"encoding/json"
	"errors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Marshaller interface {
	Marshal(v interface{}) ([]byte, error)
}

type Encoder struct {
	marshaller Marshaller
	Options
}

type Options struct {
	SensitiveMessageOptions
}

type SensitiveMessageOptions struct {
	HideSensitiveMessage bool
	Extension            protoreflect.ExtensionType
}

type DefaultJSONMarshaller struct{}

func (DefaultJSONMarshaller) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func InitWithDefaultMarshaller(o Options) Encoder {
	return Encoder{
		marshaller: DefaultJSONMarshaller{},
		Options:    o,
	}
}

func Init(o Options, m Marshaller) Encoder {
	if m == nil {
		return InitWithDefaultMarshaller(o)
	}

	return Encoder{
		marshaller: m,
		Options:    o,
	}
}

func (e Encoder) Marshal(m proto.Message) ([]byte, error) {
	if e.marshaller == nil {
		return nil, errors.New("marshaller hasn't been initialized")
	}

	if e.SensitiveMessageOptions.HideSensitiveMessage {
		m = e.clearProtoFields(m)
	}

	return e.marshaller.Marshal(m)
}

func (e Encoder) clearProtoFields(msg proto.Message) proto.Message {
	clonedMsg := proto.Clone(msg)
	reflectMsg := clonedMsg.ProtoReflect()

	reflectMsg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		return e.visitMessage(reflectMsg, fd, val)
	})

	return clonedMsg
}

func (e Encoder) visitMessage(
	message protoreflect.Message,
	fd protoreflect.FieldDescriptor,
	val protoreflect.Value,
) bool {
	if e.clearField(message, fd) {
		return true
	}

	switch {
	case fd.Kind() != protoreflect.MessageKind:
	case fd.IsMap():
	case fd.IsList():
		listVal := val.List()
		for i := 0; i < listVal.Len(); i++ {
			elem := listVal.Get(i)
			elem.Message().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				return e.visitMessage(elem.Message(), fd, v)
			})
		}
	default:
		val.Message().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			return e.visitMessage(val.Message(), fd, v)
		})
	}

	return true
}

func (e Encoder) clearField(message protoreflect.Message, fd protoreflect.FieldDescriptor) bool {
	options := fd.Options()
	if options == nil || !proto.HasExtension(options, e.SensitiveMessageOptions.Extension) {
		return false
	}

	message.Clear(fd)
	return true
}
