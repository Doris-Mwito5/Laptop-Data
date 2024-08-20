package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
)

// WriteProtobufToJSON converts a Protobuf message to JSON and writes it to a file.
func WriteProtobufToJSONFile(message protoiface.MessageV1, filename string) error {
	// Convert the Protobuf message to JSON
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal the protobuf message to JSON: %w", err)
	}

	// Write the JSON data to a file
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON to file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile serializes the Protobuf message and writes it to a binary file.
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// Serialize the message to binary format
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	// Write the binary data to the file with the appropriate file permissions
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

//function to read the protobuf message from binary form
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	//os reads the data from the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from the file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal the binaey file to a protobuf message: %w", err)
	}
	return nil
}

