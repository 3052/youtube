package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "io"
   "net/http"
)

var po_token = protobuf.Message{
   1: {protobuf.Bytes("\x01/~e\x86\b\x11\xe8\xd4\xd1\x0fh\x7fm\xef\xfc\xb4n\t\xc6\xeaI\xa5,6\x99\xb1\xb6\xe8\xc4ܷ\x94\x90\x92\xa9K\x00\x1d\xfa?\xb6\x8e\x8d`\xfc7E\x90\x88\x9d\xfd\xcbˁ\xc5\xccnI\x16b\x03s\xb1om\xdd\xf6\xa9\xe4\xc8V\x19i\xab\x81\xd2\x19kZ\xd5\xfc\x11\xed\xaa f\x8e\xb5\xcc\x02\x89\xf8j\x11\xfe\x06m\a\x8dF\xe7/tX\xa6\xebȰ\xf8\xaa\x83\xa7ʜ\xd3m\xcbҡC\xda\xc3m\xf5\xad\xa4\xa0\xbc\x16bA\xd6\xc9\x05\xb5\xbe\xb7o\x99\xbb\v\x1e\xf8\xb3\xc8iP\xbcQ\xe7\xa9Qo\xa0O\xc6i\u17ef|?!\x89\x0e+\xff\x8fS\xce\x11\xbbߚ\xe6\xb4\xd75\xe2U\xdf\xf3\xb4\xbf\xd9;\xd0̝\xdd鎫K\xc3\xff\xee\xf3H\x1b\xd5*u\xa3;\x03\xe9\x13η\xc2\xeb\x8fG\xb7}\x16\x04J\xa6aR˜\x97.\xdf\\r\t\xbdIẻ\xd1\xf6\x9a{\n\x9e˴Ռ\x8d!\xecH")},
   2: {protobuf.Bytes("\x01\x027\xfd\xb6\x1a\x91\xc5\xda\x19\xf2`k\x9f\xda\xe0\x18Bk\x00Y\xc4:\xcbٶI\xed7\xf9\xfd2\x1ae\x9c\x9a^\x91\xb8\xa2N\x0eԚڐ\x06\xbe\x81\x9a\xeff\x9a")},
}

func (v visitor_id) watch(token protobuf.Message) (*get_watch, error) {
   message := protobuf.Message{}
   message.AddFunc(1, func(m protobuf.Message) {
      m.AddFunc(1, func(m protobuf.Message) {
         m.AddVarint(16, 3)
         m.AddBytes(17, []byte("19.33.35"))
         m.AddBytes(19, []byte("9"))
         m.AddVarint(64, 28)
      })
   })
   message.AddFunc(2, func(m protobuf.Message) {
      m.AddBytes(2, []byte("40wkJJXfwQ0"))
      if token != nil {
         m.AddFunc(27, func(m protobuf.Message) {
            m.AddFunc(1, func(m protobuf.Message) {
               m.Add(1, po_token)
            })
         })
      }
   })
   req, err := http.NewRequest(
      "POST", "https://youtubei.googleapis.com/youtubei/v1/get_watch",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   id, err := v.id()
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/x-protobuf"},
      "user-agent": {"com.google.android.youtube/19.33.35(Linux; U; Android 9; en_US; Android SDK built for x86 Build/PSR1.180720.012) gzip"},
      "x-goog-visitor-id": {id},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return &get_watch{message}, nil
}
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
