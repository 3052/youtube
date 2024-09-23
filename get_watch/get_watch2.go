package main

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/19.33.35(Linux; U; Android 9; en_US; Android SDK built for x86 Build/PSR1.180720.012) gzip"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/get_watch"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(req_body.Marshal()))
   time.Sleep(time.Second)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      panic(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      panic(err)
   }
   message := protobuf.Message{}
   err = message.Unmarshal(data)
   if err != nil {
      panic(err)
   }
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   message, _ = message.Get(3)()
   address, _ := message.GetBytes(2)()
   fmt.Println(string(address))
}

// github.com/ReVanced/revanced-patches/issues/1258
var req_body = protobuf.Message{
   1: { protobuf.Message{
         1: { protobuf.Message{
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("19.33.35")},
               19: {protobuf.Bytes("9")},
               64: {protobuf.Varint(28)},
               
               86: {protobuf.Unknown{
                  protobuf.Bytes("\x12\x04\b\x03\x12\x00\x18\x00"),
                  protobuf.Message{
                     2: {protobuf.Unknown{
                        protobuf.Bytes("\b\x03\x12\x00"),
                        protobuf.Message{
                           1: {protobuf.Varint(3)},
                           2: {protobuf.Bytes("")},
                        },
                     }},
                     3: {protobuf.Varint(0)},
                  },
               }},
               92: {protobuf.Bytes("ranchu;")},
               97: {protobuf.Unknown{
                  protobuf.Bytes("\b\x01\x10w"),
                  protobuf.Message{
                     1: {protobuf.Varint(1)},
                     2: {protobuf.Varint(119)},
                  },
               }},
               98: {protobuf.Bytes("Android")},
               100: {protobuf.Unknown{
                  protobuf.Bytes("\n\x05\b\x96\x0f\x18\x01"),
                  protobuf.Message{
                     1: {protobuf.Unknown{
                        protobuf.Bytes("\b\x96\x0f\x18\x01"),
                        protobuf.Message{
                           1: {protobuf.Varint(1942)},
                           3: {protobuf.Varint(1)},
                        },
                     }},
                  },
               }},
               102: {protobuf.Unknown{
                  protobuf.Bytes("\x10\x03\x18\x01"),
                  protobuf.Message{
                     2: {protobuf.Varint(3)},
                     3: {protobuf.Varint(1)},
                  },
               }},
            },
         },
         3: {protobuf.Unknown{
            protobuf.Bytes("8\x00x\x00"),
            protobuf.Message{
               7:  {protobuf.Varint(0)},
               15: {protobuf.Varint(0)},
            },
         }},
         5: {protobuf.Unknown{
            protobuf.Bytes("\x8a\x02\x00"),
            protobuf.Message{
               33: {protobuf.Bytes("")},
            },
         }},
         6: {protobuf.Unknown{
            protobuf.Bytes("\x12\"\"\x13\b\x87\xfe\xecϮ\x91\x88\x03\x15O\xab\xd1\x04\x1dN\xe1\a\x932\bexternal\x9a\x01\x00"),
            protobuf.Message{
               2: {protobuf.Unknown{
                  protobuf.Bytes("\"\x13\b\x87\xfe\xecϮ\x91\x88\x03\x15O\xab\xd1\x04\x1dN\xe1\a\x932\bexternal\x9a\x01\x00"),
                  protobuf.Message{
                     4: {protobuf.Unknown{
                        protobuf.Bytes("\b\x87\xfe\xecϮ\x91\x88\x03\x15O\xab\xd1\x04\x1dN\xe1\a\x93"),
                        protobuf.Message{
                           1: {protobuf.Varint(1724630863396615)},
                           2: {protobuf.Fixed32(80849743)},
                           3: {protobuf.Fixed32(2466767182)},
                        },
                     }},
                     6:  {protobuf.Bytes("external")},
                     19: {protobuf.Bytes("")},
                  },
               }},
            },
         }},
         9: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     1: {protobuf.Bytes("ms")},
                     2: {protobuf.Bytes("CoACZElenKPl7VFNkBvkfIXQiA6NylhnRZKMadRZrqH8ukZfD6YahUlkBLAse-JfKDH9kemp2sKZ15D6o0ardVj_djop7xIf8lJMaAyrvI3Zgs9-v1Q7LCRBCZ6DPeyQ-QzStNG_FzCtq0CPTosQWVEV6J-kse7G4UIyhdJn280swmlKAfKd4p0YFNPJUx2V3yA5t0hEDtvPE0vFfWP8IyfqODS1-VSTh94ozXT6D-B5Wag7_C1kz3xm2ArWfmHcpo3BOfLu-JC1P3ppqDCSA5GofI4tlT8aSg_EB9PqkD79NC-LYYDmaGbxSiS-jJ7hbQa4aGpLu287EvGblUpvRLzCWQqAAqkaYmpuDHlGUwG-eS23R6ObMM0xVxtqvwCFYksO-XGMDJ9kBQCng6kttO9wWWdctjEFCR6eObqmwGSaCdoxgz-3xtLh9-FOCRmdPTHh8pwMaDvJLXzUW7cpi1UINtV0bbTRoSPOLEFYF9AVED_ruAsmcAHMOb8cO5DR5ZrPlxbG--4nyrPdXEAdxJq_b8wlJyVukDnJrizw35Uf272evt9mjW3VdsphdID6JE-_PYHdX7AeIzE5ORXDeDF0F7xmSRr13sQwYY44auBEPQblSyzQ155M7Nq-IgEnIA93kI0WvJthnjd3Hqky1P87negnh3dZlBLQ7osQJZbIAIyRGJ0SEA2ffXou5JQBPmgVvpuMJhw")},
                  },
               },
            },
         },
         12: {protobuf.Unknown{
            protobuf.Bytes("\n\bCAESAggC"),
            protobuf.Message{
               1: {protobuf.Bytes("CAESAggC")},
            },
         }},
      },
   },
   2: {
      protobuf.Message{
         2: {protobuf.Bytes("40wkJJXfwQ0")},
         3: {protobuf.Varint(0)},
         4: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     3:  {protobuf.Bytes("android-google")},
                     4:  {protobuf.Varint(90)},
                     5:  {protobuf.Varint(185)},
                     6:  {protobuf.Varint(0)},
                     7:  {protobuf.Varint(3)},
                     8:  {protobuf.Varint(0)},
                     10: {protobuf.Varint(0)},
                     11: {protobuf.Varint(0)},
                     12: {protobuf.Bytes("sdkv=a.19.33.35&output=xml_vast2")},
                     29: {protobuf.Varint(0)},
                     31: {protobuf.Unknown{
                        protobuf.Bytes(":\x00"),
                        protobuf.Message{
                           7: {protobuf.Bytes("")},
                        },
                     }},
                     37: {protobuf.Varint(0)},
                     38: {protobuf.Varint(0)},
                     41: {protobuf.Varint(0)},
                  },
               },
            },
         },
         5:  {protobuf.Varint(0)},
         8:  {protobuf.Varint(0)},
         23: {protobuf.Bytes("Fb3SDK0qUXpDTg48")},
         26: {protobuf.Bytes("")},
         27: {
            protobuf.Message{
               1: {
                  protobuf.Message{
                     1: {
                        protobuf.Message{
                           1: {protobuf.Bytes("\x01/~e\x86\b\x11\xe8\xd4\xd1\x0fh\x7fm\xef\xfc\xb4n\t\xc6\xeaI\xa5,6\x99\xb1\xb6\xe8\xc4ܷ\x94\x90\x92\xa9K\x00\x1d\xfa?\xb6\x8e\x8d`\xfc7E\x90\x88\x9d\xfd\xcbˁ\xc5\xccnI\x16b\x03s\xb1om\xdd\xf6\xa9\xe4\xc8V\x19i\xab\x81\xd2\x19kZ\xd5\xfc\x11\xed\xaa f\x8e\xb5\xcc\x02\x89\xf8j\x11\xfe\x06m\a\x8dF\xe7/tX\xa6\xebȰ\xf8\xaa\x83\xa7ʜ\xd3m\xcbҡC\xda\xc3m\xf5\xad\xa4\xa0\xbc\x16bA\xd6\xc9\x05\xb5\xbe\xb7o\x99\xbb\v\x1e\xf8\xb3\xc8iP\xbcQ\xe7\xa9Qo\xa0O\xc6i\u17ef|?!\x89\x0e+\xff\x8fS\xce\x11\xbbߚ\xe6\xb4\xd75\xe2U\xdf\xf3\xb4\xbf\xd9;\xd0̝\xdd鎫K\xc3\xff\xee\xf3H\x1b\xd5*u\xa3;\x03\xe9\x13η\xc2\xeb\x8fG\xb7}\x16\x04J\xa6aR˜\x97.\xdf\\r\t\xbdIẻ\xd1\xf6\x9a{\n\x9e˴Ռ\x8d!\xecH")},
                           2: {protobuf.Bytes("\x01\x027\xfd\xb6\x1a\x91\xc5\xda\x19\xf2`k\x9f\xda\xe0\x18Bk\x00Y\xc4:\xcbٶI\xed7\xf9\xfd2\x1ae\x9c\x9a^\x91\xb8\xa2N\x0eԚڐ\x06\xbe\x81\x9a\xeff\x9a")},
                        },
                     },
                  },
               },
            },
         },
         28: {protobuf.Unknown{
            protobuf.Bytes("\b\x00\x10\x00\x18\x00"),
            protobuf.Message{
               1: {protobuf.Varint(0)},
               2: {protobuf.Varint(0)},
               3: {protobuf.Varint(0)},
            },
         }},
      },
   },
   3: {
      protobuf.Message{
         2:  {protobuf.Bytes("40wkJJXfwQ0")},
         6:  {protobuf.Bytes("")},
         9:  {protobuf.Varint(0)},
         10: {protobuf.Varint(0)},
         24: {protobuf.Varint(0)},
         25: {protobuf.Varint(0)},
         26: {protobuf.Varint(0)},
         28: {protobuf.Varint(1)},
         30: {protobuf.Varint(0)},
         36: {protobuf.Unknown{
            protobuf.Bytes("(\xbb\x01"),
            protobuf.Message{
               5: {protobuf.Varint(187)},
            },
         }},
         47: {protobuf.Varint(0)},
         48: {protobuf.Bytes("")},
         53: {protobuf.Unknown{
            protobuf.Bytes("\b\v\x10\xc0\ued0d\x02"),
            protobuf.Message{
               1: {protobuf.Varint(11)},
               2: {protobuf.Varint(565000000)},
            },
         }},
      },
   },
}
