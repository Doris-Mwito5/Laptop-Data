package serializer_test

import (
	"github/Doris-Mwito5/pcbook/sample"
	"github/Doris-Mwito5/pcbook/serializer"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryfile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryfile)
	require.NoError(t, err)
}