package frames

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type WindowUpdateFrame struct {
	StreamId            uint32
	WindowSizeIncrement uint32
}

func NewWindowUpdateFrame(streamId uint32, windowSizeIncrement uint32) *WindowUpdateFrame {
	return &WindowUpdateFrame{
		StreamId:            streamId,
		WindowSizeIncrement: windowSizeIncrement,
	}
}

func DecodeWindowUpdateFrame(flags byte, streamId uint32, payload []byte, context *DecodingContext) (Frame, error) {
	if len(payload) < 4 {
		return nil, fmt.Errorf("FRAME_SIZE_ERROR: Received WINDOW_UPDATE frame of length %v", len(payload))
	}
	return NewWindowUpdateFrame(streamId, uint32_ignoreFirstBit(payload[0:4])), nil
}

func (f *WindowUpdateFrame) Type() Type {
	return WINDOW_UPDATE_TYPE
}

func (f *WindowUpdateFrame) Encode(context *EncodingContext) ([]byte, error) {
	payload := make([]byte, 4)
	binary.BigEndian.PutUint32(payload, uint32(f.WindowSizeIncrement))

	var result bytes.Buffer
	result.Write(encodeHeader(f.Type(), f.StreamId, uint32(len(payload)), []Flag{}))
	result.Write(payload)
	return result.Bytes(), nil
}

func (f *WindowUpdateFrame) GetStreamId() uint32 {
	return f.StreamId
}
