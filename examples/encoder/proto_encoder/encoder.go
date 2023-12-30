package main

import (
	"encoding/json"
	"fmt"

	encoder "github.com/Mahes2/go-libs/encoder/proto_encoder"
)

type JsonMarshaller struct{}

func (j JsonMarshaller) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func main() {
	options := encoder.Options{
		SensitiveMessageOptions: encoder.SensitiveMessageOptions{
			HideSensitiveMessage: true,
			Extension:            encoder.E_SensitiveMessage,
		},
	}
	enc := encoder.Init(options, JsonMarshaller{})

	message := buildMessage()

	jsonBytes, _ := enc.Marshal(message)
	fmt.Println(string(jsonBytes))
}

func buildMessage() *encoder.GetResponse {
	return &encoder.GetResponse{
		Field1: 1,
		Field2: "Hello World",
		Field3: &encoder.Message1{
			Field1: 2,
			Field2: "Encoder",
		},
		Field4: &encoder.Message2{
			Field1: true,
			Field2: "Message",
		},
		Field5: []*encoder.Message3{
			{
				Field1: 3,
				Field2: []string{
					"A",
					"B",
					"C",
				},
			}, {
				Field1: 4,
				Field2: []string{
					"D",
					"E",
					"F",
					"G",
				},
			},
		},
		Field6: &encoder.Message4{
			Field1: []*encoder.Message2{
				{
					Field1: true,
					Field2: "true",
				},
				{
					Field1: false,
					Field2: "false",
				},
			},
		},
		Field7: map[string]bool{
			"K1": true,
			"K2": false,
		},
		Field8: true,
	}
}
