package encoder

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type JsonMarshaller struct{}

func (j JsonMarshaller) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		name               string
		marshaller         Marshaller
		options            Options
		message            protoreflect.ProtoMessage
		err                error
		expectedJsonString string
	}{
		{
			name: "HideSensitiveDataWithDefaultMarshaller",
			options: Options{
				SensitiveMessageOptions: SensitiveMessageOptions{
					HideSensitiveMessage: true,
					Extension:            E_SensitiveMessage,
				},
			},
			message: &GetResponse{
				Field1: 1,
				Field2: "Hello World",
				Field3: &Message1{
					Field1: 2,
					Field2: "Encoder",
				},
				Field4: &Message2{
					Field1: true,
					Field2: "Message",
				},
				Field5: []*Message3{
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
				Field6: &Message4{
					Field1: []*Message2{
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
			},
			expectedJsonString: `{"field1":1,"field2":"Hello World","field3":{"field2":"Encoder"},"field5":[{"field1":3,"field2":["A","B","C"]},{"field1":4,"field2":["D","E","F","G"]}],"field6":{},"field8":true}`,
		},
		{
			name:       "HideSensitiveDataWithOtherMarshaller",
			marshaller: JsonMarshaller{},
			options: Options{
				SensitiveMessageOptions: SensitiveMessageOptions{
					HideSensitiveMessage: true,
					Extension:            E_SensitiveMessage,
				},
			},
			message: &GetResponse{
				Field1: 1,
				Field2: "Hello World",
				Field3: &Message1{
					Field1: 2,
					Field2: "Encoder",
				},
				Field4: &Message2{
					Field1: true,
					Field2: "Message",
				},
				Field5: []*Message3{
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
				Field6: &Message4{
					Field1: []*Message2{
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
			},
			expectedJsonString: `{"field1":1,"field2":"Hello World","field3":{"field2":"Encoder"},"field5":[{"field1":3,"field2":["A","B","C"]},{"field1":4,"field2":["D","E","F","G"]}],"field6":{},"field8":true}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encoder := Init(test.options, test.marshaller)
			jsonBytes, err := encoder.Marshal(test.message)
			if test.err != nil {
				if err != test.err {
					t.Errorf("got error %v, want %v", err, test.err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error %q", err)
			}
			if string(jsonBytes) != test.expectedJsonString {
				t.Errorf("got json string %s, want %s", string(jsonBytes), test.expectedJsonString)
			}
		})
	}
}

func BenchmarkMarshal_HidingSensitiveData(b *testing.B) {
	encoder := InitWithDefaultMarshaller(Options{
		SensitiveMessageOptions: SensitiveMessageOptions{
			HideSensitiveMessage: true,
			Extension:            E_SensitiveMessage,
		},
	})
	message := &GetResponse{
		Field1: 1,
		Field2: "Hello World",
		Field3: &Message1{
			Field1: 2,
			Field2: "Encoder",
		},
		Field4: &Message2{
			Field1: true,
			Field2: "Message",
		},
		Field5: []*Message3{
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
		Field6: &Message4{
			Field1: []*Message2{
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

	for i := 0; i < b.N; i++ {
		encoder.Marshal(message)
	}
}

func BenchmarkMarshal_NoHidingSensitiveData(b *testing.B) {
	encoder := InitWithDefaultMarshaller(Options{
		SensitiveMessageOptions: SensitiveMessageOptions{
			HideSensitiveMessage: false,
		},
	})
	message := &GetResponse{
		Field1: 1,
		Field2: "Hello World",
		Field3: &Message1{
			Field1: 2,
			Field2: "Encoder",
		},
		Field4: &Message2{
			Field1: true,
			Field2: "Message",
		},
		Field5: []*Message3{
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
		Field6: &Message4{
			Field1: []*Message2{
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

	for i := 0; i < b.N; i++ {
		encoder.Marshal(message)
	}
}
