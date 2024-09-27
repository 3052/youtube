package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

const youtube_version = "19.38.37"

type visitor_id struct {
   message protobuf.Message
}

func (v *visitor_id) New() error {
   data := protobuf.Message{
      1: {protobuf.Message{
         1: {protobuf.Message{
            16: {protobuf.Varint(3)},
            17: {protobuf.Bytes(youtube_version)},
         }},
      }},
   }.Marshal()
   resp, err := http.Post(
      "https://youtubei.googleapis.com/youtubei/v1/visitor_id",
      "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   v.message = protobuf.Message{}
   return v.message.Unmarshal(data)
}

func (v visitor_id) id() (string, error) {
   if v, ok := v.message.Get(1)(); ok {
      if v, ok := v.GetBytes(2)(); ok {
         return string(v), nil
      }
   }
   return "", errors.New("visitor_id.id")
}
