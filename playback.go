package youtube

import (
   "154.pages.dev/protobuf"
   "strconv"
)

func (p playback) String() string {
   b := []byte("itag = ")
   b = strconv.AppendUint(b, p.itag(), 10)
   b = append(b, "\nlabel = "...)
   b = append(b, p.label()...)
   return string(b)
}

func (p playback) label() string {
   v, _ := p.message.GetBytes(26)()
   return string(v)
}

func (p playback) itag() uint64 {
   v, _ := p.message.GetVarint(1)()
   return uint64(v)
}

type playback struct {
   message protobuf.Message
}
