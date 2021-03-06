package cluster

import (
	"github.com/hashicorp/go-msgpack/codec"
	"unsafe"
)

type messageType int

var (
	mh codec.MsgpackHandle
	h  = &mh // or mh to use msgpack
)

const (
	nextIdResponse messageType = iota

	nextIdRequest

	// Guble protocol.Message
	gubleMessage

	stringMessage
)

type message struct {
	NodeID int
	Type   messageType
	Body   []byte
}

func (cmsg *message) encode() ([]byte, error) {
	logger.WithField("clusterMessage", cmsg).Debug("encode")
	encodedBytes := make([]byte, cmsg.length())
	encoder := codec.NewEncoderBytes(&encodedBytes, h)
	err := encoder.Encode(cmsg)
	if err != nil {
		logger.WithField("err", err).Error("Encoding failed")
		return nil, err
	}
	return encodedBytes, nil
}

func (cmsg *message) length() int {
	return int(unsafe.Sizeof(cmsg.Type)) + int(unsafe.Sizeof(cmsg.NodeID)) + len(cmsg.Body)
}

func decode(cmsgBytes []byte) (*message, error) {
	var cmsg message
	logger.WithField("clusterMessageBytes", string(cmsgBytes)).Debug("decode")

	decoder := codec.NewDecoderBytes(cmsgBytes, h)
	err := decoder.Decode(&cmsg)
	if err != nil {
		logger.WithField("err", err).Error("Decoding failed")
		return nil, err
	}
	return &cmsg, nil
}
