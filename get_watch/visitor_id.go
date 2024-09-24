package youtube

import (
	"154.pages.dev/protobuf"
	"bytes"
	"io"
	"net/http"
)

func visitor_id() (protobuf.Message, error) {
	message := protobuf.Message{
		1: {protobuf.Message{
			1: {protobuf.Message{
				16: {protobuf.Varint(3)},
				17: {protobuf.Bytes("19.33.35")},
			}},
		}},
	}
	resp, err := http.Post(
		"https://youtubei.googleapis.com/youtubei/v1/visitor_id",
		"application/x-protobuf", bytes.NewReader(message.Marshal()),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	message = protobuf.Message{}
	err = message.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return message, nil
}
