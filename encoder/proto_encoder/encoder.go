package encoder

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var marshaller Marshaller

type Marshaller interface {
	Marshal(v interface{}) ([]byte, error)
}

type DefaultJSONMarshaller struct{}

type Options struct {
	SensitiveMessageOptions
}

type SensitiveMessageOptions struct {
	HideSensitiveMessage bool
	Extension            protoreflect.ExtensionType
}

func (DefaultJSONMarshaller) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Init(m Marshaller) {
	marshaller = m
}

func (o Options) Marshal(m proto.Message) ([]byte, error) {
	if marshaller == nil {
		marshaller = DefaultJSONMarshaller{}
	}

	if o.SensitiveMessageOptions.HideSensitiveMessage {
		m = clearProtoFields(m, o.SensitiveMessageOptions.Extension)
	}

	return marshaller.Marshal(m)
}

func clearProtoFields(msg proto.Message, sensitiveFieldAnnotation protoreflect.ExtensionType) proto.Message {
	clonedMsg := proto.Clone(msg)
	reflectMsg := clonedMsg.ProtoReflect()

	reflectMsg.Range(func(fd protoreflect.FieldDescriptor, val protoreflect.Value) bool {
		return visitMessage(reflectMsg, fd, val, sensitiveFieldAnnotation)
	})

	return clonedMsg
}

func visitMessage(
	message protoreflect.Message,
	fd protoreflect.FieldDescriptor,
	val protoreflect.Value,
	sensitiveFieldAnnotation protoreflect.ExtensionType,
) bool {
	if clearField(sensitiveFieldAnnotation, message, fd) {
		return true
	}

	switch {
	case fd.Kind() != protoreflect.MessageKind:
		return true
	case fd.IsList():
		listVal := val.List()
		for i := 0; i < listVal.Len(); i++ {
			elem := listVal.Get(i)
			elem.Message().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				return visitMessage(elem.Message(), fd, v, sensitiveFieldAnnotation)
			})
		}
	case fd.IsMap():
		return true
	default:
		val.Message().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			return visitMessage(val.Message(), fd, v, sensitiveFieldAnnotation)
		})
	}

	return true
}

func clearField(sensitiveFieldAnnotation protoreflect.ExtensionType, message protoreflect.Message, fd protoreflect.FieldDescriptor) bool {
	options := fd.Options()
	if options != nil && proto.HasExtension(options, sensitiveFieldAnnotation) {
		message.Clear(fd)
		return true
	}

	return false
}
