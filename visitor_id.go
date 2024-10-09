package youtube

import (
   "41.neocities.org/protobuf"
   "bytes"
   "encoding/json"
   "net/http"
)

const youtube_version = "19.38.37"

type visitor_id struct {
   ResponseContext struct {
      VisitorData string
   }
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
      "https://youtubei.googleapis.com/youtubei/v1/visitor_id?alt=json",
      "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(v)
}
