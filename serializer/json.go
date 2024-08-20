package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToJSON converts a Protobuf message to a JSON string using jsonpb.
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false, // Enums are represented as strings
		EmitDefaults: true,  // Include fields with default values
		Indent:       " ",   // Pretty-print with indentation
		OrigName:     true,  // Use original field names as defined in the .proto file
	}
	return marshaler.MarshalToString(message)
}
