package serializer_test

import (
	"mygrpc/pb"
	"mygrpc/sample"
	"mygrpc/serializer"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(jsonFile, laptop1)
	require.NoError(t, err)

	laptop3 := &pb.Laptop{}
	err = serializer.ReadProtobufFromJSONFile(jsonFile, laptop3)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop3))
}
