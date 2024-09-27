package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "encoding/json"
   "net/http"
   "strconv"
)

func (v *visitor_id) watch(
   video string, po_token protobuf.Message,
) (adaptive_formats, error) {
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var watch []struct {
      PlayerResponse struct {
         StreamingData struct {
            AdaptiveFormats adaptive_formats
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&watch)
   if err != nil {
      return nil, err
   }
   return watch[0].PlayerResponse.StreamingData.AdaptiveFormats, nil
}

type adaptive_formats []adaptive_format

type adaptive_format struct {
   AudioQuality string
   Bitrate int64
   IsDrc bool // wikipedia.org/wiki/Dynamic_range_compression
   MimeType string
   QualityLabel string
   Url string
}

func (a *adaptive_format) String() string {
   b := []byte("bitrate = ")
   b = strconv.AppendInt(b, a.Bitrate, 10)
   if a.IsDrc {
      b = append(b, "\ndrc = true"...)
   }
   if a.AudioQuality != "" {
      b = append(b, "\naudio = "...)
      b = append(b, a.AudioQuality...)
   }
   b = append(b, "\ntype = "...)
   b = append(b, a.MimeType...)
   if a.QualityLabel != "" {
      b = append(b, "\nlabel = "...)
      b = append(b, a.QualityLabel...)
   }
   return string(b)
}

func (a adaptive_formats) String() string {
   var b []byte
   for i, format := range a {
      if i >= 1 {
         b = append(b, "\n\n"...)
      }
      b = append(b, "index = "...)
      b = strconv.AppendInt(b, int64(i), 10)
      b = append(b, '\n')
      b = append(b, format.String()...)
   }
   return string(b)
}
