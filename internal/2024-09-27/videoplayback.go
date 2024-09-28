package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Cache-Control"] = []string{"no-cache"}
   req.Header["Connection"] = []string{"keep-alive"}
   req.Header["Dnt"] = []string{"1"}
   req.Header["Host"] = []string{"rr2---sn-q4flrn7y.googlevideo.com"}
   req.Header["Origin"] = []string{"https://www.youtube.com"}
   req.Header["Pragma"] = []string{"no-cache"}
   req.Header["Referer"] = []string{"https://www.youtube.com/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"cross-site"}
   req.Header["Sec-Gpc"] = []string{"1"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "rr2---sn-q4flrn7y.googlevideo.com"
   req.URL.Path = "/videoplayback"
   value := url.Values{}
   value["c"] = []string{"WEB"}
   value["cpn"] = []string{"VbortoAYAvj_rLBY"}
   value["cver"] = []string{"2.20240926.00.00"}
   value["ei"] = []string{"xjr3ZqHfBKGN2_gPgpTYqQg"}
   value["expire"] = []string{"1727500070"}
   value["fexp"] = []string{"51299152"}
   value["fvip"] = []string{"1"}
   value["id"] = []string{"o-APfR73aaofu1T94nXV0JsIcoCglGu8xBxjODtYtfqv_w"}
   value["initcwndbps"] = []string{"1276250"}
   value["ip"] = []string{"72.181.16.91"}
   value["keepalive"] = []string{"yes"}
   value["lsig"] = []string{"ABPmVW0wRQIhAOptJIs1SZTuGYZhbPYsVazs7kv-AlfJMgzBpBuTHUWXAiAYFs4bHiacdVEC6GbI8s3voO3vRCbA7KmiKKxtyS9P1Q=="}
   value["lsparams"] = []string{"mh,mm,mn,ms,mv,mvi,pl,initcwndbps"}
   value["mh"] = []string{"am"}
   value["mm"] = []string{"31,29"}
   value["mn"] = []string{"sn-q4flrn7y,sn-q4fl6n66"}
   value["ms"] = []string{"au,rdu"}
   value["mt"] = []string{"1727478056"}
   value["mv"] = []string{"m"}
   value["mvi"] = []string{"2"}
   value["n"] = []string{"GXINieRLk3xLFQ"}
   value["ns"] = []string{"kHNtSl4Tzx-R5qzDSTxnkG8Q"}
   value["pl"] = []string{"19"}
   value["requiressl"] = []string{"yes"}
   value["rn"] = []string{"19"}
   value["rqh"] = []string{"1"}
   value["sabr"] = []string{"1"}
   value["sig"] = []string{"AJfQdSswRgIhANp-z6q5u9n_bYG2TX1TtHcm8r0IwLhH0x8TGZlrbKgwAiEAxP93ElyKsvKe6vZFH90j8orfxHGJHoNFi7mIPNmp2E4="}
   value["source"] = []string{"youtube"}
   value["sparams"] = []string{"expire,ei,ip,id,source,requiressl,xpc,spc,svpuc,ns,sabr,rqh"}
   value["spc"] = []string{"54MbxXih-0O8ob7r1DNOJxfkPaIxRZVY0qS0YOZ153j8gHg7nSpm0P7W5mOA"}
   value["svpuc"] = []string{"1"}
   value["xpc"] = []string{"EgVo2aDSNQ=="}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   protobuf.Deterministic = true
   req.Body = io.NopCloser(bytes.NewReader(req_body.Marshal()))
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   var buf bytes.Buffer
   resp.Write(&buf)
   fmt.Printf("%q\n", buf.Bytes()[:1599])
}

var req_body = protobuf.Message{
   5: {
      protobuf.Message{
         1: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     1:  {protobuf.Varint(0)},
                     4:  {protobuf.Fixed32(1065353216)},
                     5:  {protobuf.Fixed32(1064514355)},
                     6:  {protobuf.Fixed32(1066863165)},
                     11: {protobuf.Varint(1)},
                     13: {protobuf.Varint(1)},
                     14: {
                        protobuf.Message{
                           1: {protobuf.Bytes("mfs2_cmfs_web_v3_2_003")},
                           3: {protobuf.Varint(0)},
                        },
                     },
                     15: {protobuf.Varint(9999)},
                     20: {protobuf.Varint(1)},
                     21: {protobuf.Varint(0)},
                     34: {protobuf.Varint(1)},
                     39: {protobuf.Varint(0)},
                     41: {protobuf.Varint(1)},
                     43: {
                        protobuf.Message{
                           2:         {protobuf.Varint(30000)},
                           3:         {protobuf.Varint(3600000)},
                           4:         {protobuf.Varint(20000)},
                           5:         {protobuf.Varint(20000)},
                           6:         {protobuf.Varint(15000)},
                           14:        {protobuf.Varint(5000)},
                           16:        {protobuf.Varint(500)},
                           23:        {protobuf.Varint(1)},
                           28:        {protobuf.Varint(3)},
                           34:        {protobuf.Varint(1)},
                           35:        {protobuf.Varint(12)},
                           36:        {protobuf.Varint(1)},
                           40:        {protobuf.Varint(1)},
                           42:        {protobuf.Varint(2)},
                           44:        {protobuf.Varint(1)},
                           45:        {protobuf.Varint(4)},
                           48:        {protobuf.Varint(2)},
                           49:        {protobuf.Varint(5000)},
                           51:        {protobuf.Varint(1)},
                           53:        {protobuf.Varint(3)},
                           56:        {protobuf.Varint(1)},
                           57:        {protobuf.Varint(1)},
                           58:        {protobuf.Varint(1)},
                           63:        {protobuf.Varint(1)},
                           64:        {protobuf.Varint(1)},
                           65:        {protobuf.Varint(1)},
                           66:        {protobuf.Varint(1)},
                           67:        {protobuf.Varint(1)},
                           68:        {protobuf.Varint(1)},
                           69:        {protobuf.Varint(1)},
                           73:        {protobuf.Varint(1)},
                           74:        {protobuf.Varint(1)},
                           75:        {protobuf.Varint(1)},
                           76:        {protobuf.Varint(0)},
                           77:        {protobuf.Varint(1)},
                           79:        {protobuf.Varint(7)},
                           80:        {protobuf.Varint(125)},
                           81:        {protobuf.Varint(1)},
                           86:        {protobuf.Varint(1)},
                           87:        {protobuf.Varint(1)},
                           88:        {protobuf.Varint(1)},
                           89:        {protobuf.Varint(1)},
                           90:        {protobuf.Varint(1)},
                           91:        {protobuf.Varint(1)},
                           92:        {protobuf.Varint(2000)},
                           93:        {protobuf.Varint(1)},
                           95:        {protobuf.Varint(2000)},
                           96:        {protobuf.Varint(1)},
                           98:        {protobuf.Varint(1)},
                           103:       {protobuf.Varint(1)},
                           104:       {protobuf.Varint(1)},
                           106:       {protobuf.Varint(1)},
                           107:       {protobuf.Varint(1)},
                           109:       {protobuf.Varint(1)},
                           110:       {protobuf.Varint(1)},
                           111:       {protobuf.Varint(1)},
                           112:       {protobuf.Varint(2000)},
                           114:       {protobuf.Varint(1)},
                           117:       {protobuf.Varint(1)},
                           121:       {protobuf.Varint(1)},
                           122:       {protobuf.Varint(1)},
                           123:       {protobuf.Varint(1)},
                           126:       {protobuf.Varint(1)},
                           131:       {protobuf.Fixed32(3212836864)},
                           132:       {protobuf.Varint(1000)},
                           429165407: {protobuf.Varint(1)},
                        },
                     },
                     47: {
                        protobuf.Message{
                           5:   {protobuf.Fixed32(1115815936)},
                           6:   {protobuf.Fixed32(1117126656)},
                           9:   {protobuf.Varint(1)},
                           12:  {protobuf.Fixed32(1082130432)},
                           13:  {protobuf.Varint(14400)},
                           21:  {protobuf.Varint(50000)},
                           22:  {protobuf.Varint(480)},
                           23:  {protobuf.Varint(1)},
                           25:  {protobuf.Fixed32(1065353216)},
                           30:  {protobuf.Varint(1)},
                           31:  {protobuf.Fixed32(1065353216)},
                           32:  {protobuf.Fixed32(1058642330)},
                           33:  {protobuf.Fixed32(1065353216)},
                           34:  {protobuf.Fixed32(1107427328)},
                           35:  {protobuf.Varint(1)},
                           38:  {protobuf.Fixed32(1065353216)},
                           40:  {protobuf.Varint(480)},
                           42:  {protobuf.Bytes("\xb0\xff\xff\xff\xff\xff\xff\xff\xff\x01\x1e<FZ\\]^")},
                           43:  {protobuf.Bytes("20:00")},
                           44:  {protobuf.Varint(120)},
                           45:  {protobuf.Varint(360)},
                           46:  {protobuf.Fixed32(1000593162)},
                           47:  {protobuf.Fixed32(1036831949)},
                           48:  {protobuf.Varint(1)},
                           50:  {protobuf.Varint(1)},
                           51:  {protobuf.Fixed32(1025758986)},
                           52:  {protobuf.Varint(1)},
                           55:  {protobuf.Varint(1)},
                           57:  {protobuf.Varint(1)},
                           59:  {protobuf.Varint(1)},
                           60:  {protobuf.Fixed32(1078217314)},
                           61:  {protobuf.Fixed32(1056164402)},
                           62:  {protobuf.Varint(1)},
                           63:  {protobuf.Fixed32(1065772646)},
                           64:  {protobuf.Fixed32(1082130432)},
                           67:  {protobuf.Varint(1)},
                           74:  {protobuf.Fixed32(1092616192)},
                           77:  {protobuf.Varint(2160)},
                           78:  {protobuf.Varint(1)},
                           89:  {protobuf.Varint(1)},
                           92:  {protobuf.Varint(1)},
                           99:  {protobuf.Varint(1)},
                           101: {protobuf.Varint(1)},
                           102: {protobuf.Fixed32(897988541)},
                           103: {protobuf.Fixed32(1082340147)},
                           114: {protobuf.Varint(1)},
                           118: {protobuf.Varint(1)},
                           120: {protobuf.Varint(1)},
                           121: {protobuf.Varint(1)},
                           122: {protobuf.Fixed32(1134395392)},
                           124: {protobuf.Fixed32(1141473280)},
                           132: {protobuf.Fixed64(13830554455654793216)},
                           133: {protobuf.Fixed64(13830554455654793216)},
                           134: {protobuf.Varint(240)},
                           135: {protobuf.Varint(1)},
                           139: {protobuf.Varint(240)},
                           141: {protobuf.Varint(1)},
                           145: {protobuf.Varint(1)},
                           146: {protobuf.Fixed32(1134395392)},
                        },
                     },
                     48: {protobuf.Bytes{}},
                     50: {protobuf.Varint(1)},
                     53: {protobuf.Varint(1)},
                     54: {protobuf.Varint(3)},
                     58: {protobuf.Varint(1)},
                     59: {protobuf.Varint(1)},
                     60: {protobuf.Varint(10000)},
                     71: {protobuf.Varint(1)},
                     73: {
                        protobuf.Message{
                           1: {
                              protobuf.Message{
                                 1: {protobuf.Varint(120000)},
                                 2: {protobuf.Varint(15000)},
                                 3: {protobuf.Varint(1000)},
                                 4: {protobuf.Fixed32(0)},
                                 5: {protobuf.Varint(0)},
                                 6: {protobuf.Varint(0)},
                              },
                           },
                           2: {protobuf.Varint(60000)},
                           3: {protobuf.Varint(2000)},
                        },
                     },
                     74: {
                        protobuf.Message{
                           1: {
                              protobuf.Message{
                                 1: {protobuf.Varint(1200)},
                                 2: {protobuf.Varint(1200)},
                                 4: {protobuf.Varint(1)},
                              },
                           },
                           4: {protobuf.Varint(5000)},
                        },
                     },
                     75: {
                        protobuf.Message{
                           1: {
                              protobuf.Message{
                                 1: {protobuf.Varint(6000)},
                                 2: {protobuf.Varint(6000)},
                              },
                           },
                           4: {protobuf.Varint(6000)},
                        },
                     },
                     79:  {protobuf.Varint(1)},
                     82:  {protobuf.Varint(1)},
                     85:  {protobuf.Varint(1)},
                     90:  {protobuf.Varint(1)},
                     91:  {protobuf.Varint(1)},
                     93:  {protobuf.Varint(1)},
                     94:  {protobuf.Varint(1)},
                     97:  {protobuf.Varint(1)},
                     99:  {protobuf.Varint(1)},
                     101: {protobuf.Varint(32768)},
                     104: {protobuf.Varint(1)},
                     105: {protobuf.Varint(1)},
                     106: {
                        protobuf.Message{
                           1: {protobuf.Varint(1000)},
                           2: {protobuf.Varint(100)},
                           3: {
                              protobuf.Message{
                                 1: {protobuf.Varint(5000)},
                                 2: {protobuf.Fixed32(1056964608)},
                                 3: {protobuf.Fixed32(1061997773)},
                              },
                           },
                        },
                     },
                     108: {protobuf.Varint(1)},
                     112: {
                        protobuf.Message{
                           2: {protobuf.Fixed32(1065353216)},
                           3: {protobuf.Varint(100)},
                           4: {protobuf.Varint(10000)},
                        },
                     },
                     113: {protobuf.Varint(1)},
                     116: {protobuf.Varint(1)},
                     119: {protobuf.Varint(1)},
                     120: {protobuf.Varint(1)},
                     128: {protobuf.Varint(1)},
                     132: {protobuf.Varint(1)},
                     134: {protobuf.Varint(1)},
                     135: {protobuf.Varint(1)},
                     138: {
                        protobuf.Message{
                           1: {protobuf.Varint(1)},
                           2: {protobuf.Varint(1)},
                           3: {protobuf.Varint(1)},
                        },
                     },
                     147:       {protobuf.Varint(1)},
                     149:       {protobuf.Fixed64(13830554455654793216)},
                     150:       {protobuf.Fixed64(13830554455654793216)},
                     153:       {protobuf.Varint(1)},
                     155:       {protobuf.Bytes("SMEJUlmS1oOXDt1ju4gFJeM0dvgO8GHoe5Ul")},
                     156:       {protobuf.Varint(1)},
                     166:       {protobuf.Varint(1)},
                     171:       {protobuf.Varint(1)},
                     177:       {protobuf.Varint(1)},
                     179:       {protobuf.Varint(1)},
                     184:       {protobuf.Varint(1)},
                     185:       {protobuf.Varint(1)},
                     186:       {protobuf.Varint(1)},
                     187:       {protobuf.Varint(1)},
                     189:       {protobuf.Bytes("\x8b\x06\x8c\x06")},
                     191:       {protobuf.Varint(1)},
                     194:       {protobuf.Varint(1)},
                     197:       {protobuf.Varint(144)},
                     198:       {protobuf.Varint(1)},
                     199:       {protobuf.Varint(1)},
                     200:       {protobuf.Varint(1)},
                     202:       {protobuf.Varint(1)},
                     204:       {protobuf.Varint(1)},
                     208:       {protobuf.Varint(1)},
                     212:       {protobuf.Varint(1)},
                     218:       {protobuf.Varint(1)},
                     220:       {protobuf.Varint(1)},
                     225:       {protobuf.Varint(1)},
                     226:       {protobuf.Varint(1)},
                     230:       {protobuf.Varint(1)},
                     233:       {protobuf.Varint(1)},
                     250:       {protobuf.Varint(1)},
                     388565617: {protobuf.Varint(1)},
                  },
               },
               3: {protobuf.Varint(1)},
               4: {protobuf.Varint(1)},
               6: {
                  protobuf.Message{
                     1: {protobuf.Varint(137)},
                     2: {protobuf.Varint(1699655337114949)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(248)},
                     2: {protobuf.Varint(1699664105050004)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(399)},
                     2: {protobuf.Varint(1699655077389253)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(136)},
                     2: {protobuf.Varint(1699655258474940)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(247)},
                     2: {protobuf.Varint(1699694363397823)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(398)},
                     2: {protobuf.Varint(1699658205612393)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(135)},
                     2: {protobuf.Varint(1699655480045133)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(244)},
                     2: {protobuf.Varint(1699694941187198)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(397)},
                     2: {protobuf.Varint(1699656423406986)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(134)},
                     2: {protobuf.Varint(1699656208055818)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(243)},
                     2: {protobuf.Varint(1699691229355595)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(396)},
                     2: {protobuf.Varint(1699654960551981)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(133)},
                     2: {protobuf.Varint(1699654821561017)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(242)},
                     2: {protobuf.Varint(1699694312030064)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(395)},
                     2: {protobuf.Varint(1699654771170051)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(160)},
                     2: {protobuf.Varint(1699655509406462)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(278)},
                     2: {protobuf.Varint(1699691101075592)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(394)},
                     2: {protobuf.Varint(1699654694538392)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(140)},
                     2: {protobuf.Varint(1699614126944198)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(140)},
                     2: {protobuf.Varint(1699617903657306)},
                     3: {protobuf.Bytes("CggKA2RyYxIBMQ")},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(249)},
                     2: {protobuf.Varint(1699617755965408)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(249)},
                     2: {protobuf.Varint(1699619606942568)},
                     3: {protobuf.Bytes("CggKA2RyYxIBMQ")},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(250)},
                     2: {protobuf.Varint(1699617773001086)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(250)},
                     2: {protobuf.Varint(1699619589835173)},
                     3: {protobuf.Bytes("CggKA2RyYxIBMQ")},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(251)},
                     2: {protobuf.Varint(1699617745862517)},
                  },
                  protobuf.Message{
                     1: {protobuf.Varint(251)},
                     2: {protobuf.Varint(1699619606771124)},
                     3: {protobuf.Bytes("CggKA2RyYxIBMQ")},
                  },
               },
               7: {protobuf.Bytes{}},
               9: {protobuf.Varint(0)},
               10: {
                  protobuf.Message{
                     3:  {protobuf.Bytes("en")},
                     5:  {protobuf.Varint(0)},
                     6:  {protobuf.Bytes("UCqeAJqujkbzivXidDfN8kug")},
                     7:  {protobuf.Varint(0)},
                     8:  {protobuf.Varint(0)},
                     11: {protobuf.Varint(0)},
                  },
               },
               473865394: {protobuf.Varint(1)},
            },
         },
         2: {protobuf.Bytes("\x00\r\xe9F\n0E\x02!\x00\xab8t\x0e~P\xbcr\xd5[-\"\xc2\xf7\x8f\xbd=Vc\xa0\x13\x89\x126E\xf37m\xa8\xf6\x14\xa1\x02 ?\x9a\xc9:\x81\"+\xab\v\xa9\xec\x13B\xfa\x1c\xd6E\xf9ZpGVh\xef6\xd8;C\xc6b\x99\xd1")},
         3: {protobuf.Bytes("ei")},
      },
   },
}
