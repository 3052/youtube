package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
   "net/url"
   "os"
   "time"
)

func main() {
   m := protobuf.Message{}
   m.Add(1, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.AddVarint(16, 3)
         m.AddBytes(17, []byte("19.33.35"))
      })
   })
   m := protobuf.Message{
      1: {protobuf.Message{
         1: {protobuf.Message{
            16: {protobuf.Varint(3)},
            17: {protobuf.Bytes("19.33.35")},
         }},
      }},
   }
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/visitor_id"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Body = io.NopCloser(bytes.NewReader(req_body.Marshal()))
   time.Sleep(time.Second)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
