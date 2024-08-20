package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

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
	
}