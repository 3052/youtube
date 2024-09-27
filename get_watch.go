package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "errors"
   "io"
   "net/http"
)

type get_watch struct {
   message protobuf.Message
}

func (g get_watch) watch() (string, bool) {
   m, _ := g.message.Get(1)()
   m, _ = m.Get(2)()
   m, _ = m.Get(4)()
   m, _ = m.Get(3)()
   watch, ok := m.GetBytes(2)()
   return string(watch), ok
}

func (v visitor_id) watch(
   video string, po_token protobuf.Message,
) (*get_watch, error) {
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
   id, err := v.id()
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type":      {"application/x-protobuf"},
      "user-agent":        {"com.google.android.youtube/" + youtube_version},
      "x-goog-visitor-id": {id},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message = protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return &get_watch{message}, nil
}
