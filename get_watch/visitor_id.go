package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

func (v visitor_id) id() (string, bool) {
   m, _ := v.message.Get(1)()
   b, ok := m.GetBytes(2)()
   return string(b), ok
}

func (v *visitor_id) New() error {
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
      return err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   v.message = protobuf.Message{}
   return v.message.Unmarshal(data)
}

type visitor_id struct {
   message protobuf.Message
}
