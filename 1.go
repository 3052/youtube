package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/url"
)

func visitor_id() (string, error) {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/visitor_id"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body_1.Marshal()))
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return "", err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return "", err
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return "", err
   }
   if v, ok := message.Get(1)(); ok {
      if v, ok := v.GetBytes(2)(); ok {
         return string(v), nil
      }
   }
   return "", errors.New("visitor_id")
}

var body_1 = protobuf.Message{
   1: {
      protobuf.Message{
         1: {
            protobuf.Message{
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("19.33.35")},
            },
         },
      },
   },
}
