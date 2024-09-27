package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "net/http"
)

func (v *visitor_id) watch(
   video string, po_token protobuf.Message,
) (*http.Response, error) {
   message := protobuf.Message{}
   message.Add(1, func(m protobuf.Message) {
      m.Add(1, func(m protobuf.Message) {
         m.AddVarint(16, 3)
         m.AddBytes(17, []byte(youtube_version))
         m.AddBytes(19, []byte{'9'})
         m.AddVarint(64, 28)
      })
   })
   message.Add(2, func(m protobuf.Message) {
      m.AddBytes(2, []byte(video))
      if po_token != nil {
         m.Add(27, func(m protobuf.Message) {
            m.Add(1, func(m protobuf.Message) {
               m.AddMessage(1, po_token)
            })
         })
      }
   })
   req, err := http.NewRequest(
      "POST", "https://youtubei.googleapis.com/youtubei/v1/get_watch",
      bytes.NewReader(message.Marshal()),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type":      {"application/x-protobuf"},
      "user-agent":        {"com.google.android.youtube/" + youtube_version},
      "x-goog-visitor-id": {v.ResponseContext.VisitorData},
   }
   req.URL.RawQuery = "alt=json"
   return http.DefaultClient.Do(req)
}
