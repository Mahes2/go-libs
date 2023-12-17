package encoder

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Marshaller interface {
	Marshal(v interface{}) ([]byte, error)
}

var marshaller Marshaller

func Init(m Marshaller) {
	marshaller = m
}

type DefaultJSONMarshaller struct{}

func (DefaultJSONMarshaller) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

type SensitiveMessageOptions struct {
	HideSensitiveMessage bool
	Extension            protoreflect.ExtensionType
}

type Options struct {
	SensitiveMessageOptions
}

func (o Options) Marshal(m proto.Message) ([]byte, error) {
	if marshaller == nil {
		marshaller = DefaultJSONMarshaller{}
	}

	if o.SensitiveMessageOptions.HideSensitiveMessage {
		m = o.clearProtoFields(m)
	}

	return marshaller.Marshal(m)
}

func (o Options) clearProtoFields(msg proto.Message) proto.Message {
	clonedMsg := proto.Clone(msg)
	reflectMsg := clonedMsg.ProtoReflect()

	reflectMsg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		return o.visitMessage(reflectMsg, fd, val)
	})

	return clonedMsg
}

func (o Options) visitMessage(
	message protoreflect.Message,
	fd protoreflect.FieldDescriptor,
	val protoreflect.Value,
) bool {
	if o.clearField(message, fd) {
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
				return o.visitMessage(elem.Message(), fd, v)
			})
		}
	default:
		val.Message().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			return o.visitMessage(val.Message(), fd, v)
		})
	}

	return true
}

func (o Options) clearField(message protoreflect.Message, fd protoreflect.FieldDescriptor) bool {
	options := fd.Options()
	if options != nil && proto.HasExtension(options, o.SensitiveMessageOptions.Extension) {
		message.Clear(fd)
		return true
	}

	return false
}
