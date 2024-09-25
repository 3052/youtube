package youtube

import (
   "154.pages.dev/protobuf"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
   "bytes"
)

func visitor_id() (string, error) {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Priority"] = []string{"u=2, i"}
   req.Header["User-Agent"] = []string{"com.google.android.youtube/19.33.35(Linux; U; Android 9; en_US; Android SDK built for x86 Build/PSR1.180720.012) gzip"}
   req.Header["X-Goog-Api-Format-Version"] = []string{"2"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "youtubei.googleapis.com"
   req.URL.Path = "/youtubei/v1/visitor_id"
   value := url.Values{}
   value["key"] = []string{"AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(bytes.NewReader(body_1.Marshal()))
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
   1: {protobuf.Unknown{
      protobuf.Bytes("\n\xdf\x01b\aunknownj\x19Android SDK built for x86\x80\x01\x03\x8a\x01\b19.33.35\x92\x01\aAndroid\x9a\x01\x019\xaa\x01\x05en-US\xb2\x01\x02US\xa8\x02\x9b\x03\xb0\x02\xab\x05\xbd\x02I\x92$@\xc5\x02\x8b\xaf\x88@\xc8\x02\x03\xf0\x02\x01\x8a\x03\t\xaa\x03\x06310260\x90\x03\xd6\xf1\xabi\xa0\x03\x04\xb8\x03\x9b\x03\xc0\x03\xab\x05\xe8\x03\x03\xf2\x03\x00\x80\x04\x1c\x8d\x04\x00\x00(@\x98\x04\xd4\xfd\xff\xff\xff\xff\xff\xff\xff\x01\xed\x04\x00\x00\x80?\xf0\x04\x01\x82\x05\x0fAmerica/Chicago\xe2\x05\aranchu;\x92\x06\aAndroid\xa2\x06\a\n\x05\b\x96\x0f\x18\x01\xb2\x06\x04\x10\x03\x18\x01\x1a\x048\x00x\x00*\x03\x8a\x02\x002\x02\x12\x00J\xd5\x05\n\xd2\x05\n\x02ms\x12\xcb\x05CoACZElenKPl7VFNkBvkfIXQiA6NylhnRZKMadRZrqH8ukZfD6YahUlkBLAse-JfKDH9kemp2sKZ15D6o0ardVj_djop7xIf8lJMaAyrvI3Zgs9-v1Q7LCRBCZ6DPeyQ-QzStNG_FzCtq0CPTosQWVEV6J-kse7G4UIyhdJn280swmlKAfKd4p0YFNPJUx2V3yA5t0hEDtvPE0vFfWP8IyfqODS1-VSTh94ozXT6D-B5Wag7_C1kz3xm2ArWfmHcpo3BOfLu-JC1P3ppqDCSA5GofI4tlT8aSg_EB9PqkD79NC-LYYDmaGbxSiS-jJ7hbQa4aGpLu287EvGblUpvRLzCWQqAAqkaYmpuDHlGUwG-eS23R6ObMM0xVxtqvwCFYksO-XGMDJ9kBQCng6kttO9wWWdctjEFCR6eObqmwGSaCdoxgz-3xtLh9-FOCRmdPTHh8pwMaDvJLXzUW7cpi1UINtV0bbTRoSPOLEFYF9AVED_ruAsmcAHMOb8cO5DR5ZrPlxbG--4nyrPdXEAdxJq_b8wlJyVukDnJrizw35Uf272evt9mjW3VdsphdID6JE-_PYHdX7AeIzE5ORXDeDF0F7xmSRr13sQwYY44auBEPQblSyzQ155M7Nq-IgEnIA93kI0WvJthnjd3Hqky1P87negnh3dZlBLQ7osQJZbIAIyRGJ0SEA2ffXou5JQBPmgVvpuMJhw"),
      protobuf.Message{
         1: {protobuf.Unknown{
            protobuf.Bytes("b\aunknownj\x19Android SDK built for x86\x80\x01\x03\x8a\x01\b19.33.35\x92\x01\aAndroid\x9a\x01\x019\xaa\x01\x05en-US\xb2\x01\x02US\xa8\x02\x9b\x03\xb0\x02\xab\x05\xbd\x02I\x92$@\xc5\x02\x8b\xaf\x88@\xc8\x02\x03\xf0\x02\x01\x8a\x03\t\xaa\x03\x06310260\x90\x03\xd6\xf1\xabi\xa0\x03\x04\xb8\x03\x9b\x03\xc0\x03\xab\x05\xe8\x03\x03\xf2\x03\x00\x80\x04\x1c\x8d\x04\x00\x00(@\x98\x04\xd4\xfd\xff\xff\xff\xff\xff\xff\xff\x01\xed\x04\x00\x00\x80?\xf0\x04\x01\x82\x05\x0fAmerica/Chicago\xe2\x05\aranchu;\x92\x06\aAndroid\xa2\x06\a\n\x05\b\x96\x0f\x18\x01\xb2\x06\x04\x10\x03\x18\x01"),
            protobuf.Message{
               12: {protobuf.Bytes("unknown")},
               13: {protobuf.Bytes("Android SDK built for x86")},
               16: {protobuf.Varint(3)},
               17: {protobuf.Bytes("19.33.35")},
               18: {protobuf.Bytes("Android")},
               19: {protobuf.Bytes("9")},
               21: {protobuf.Unknown{
                  protobuf.Bytes("en-US"),
                  protobuf.Message{
                     12: {protobuf.Fixed32(1398091118)},
                  },
               }},
               22: {protobuf.Bytes("US")},
               37: {protobuf.Varint(411)},
               38: {protobuf.Varint(683)},
               39: {protobuf.Fixed32(1076138569)},
               40: {protobuf.Fixed32(1082699659)},
               41: {protobuf.Varint(3)},
               46: {protobuf.Varint(1)},
               49: {protobuf.Unknown{
                  protobuf.Bytes("\xaa\x03\x06310260"),
                  protobuf.Message{
                     53: {protobuf.Bytes("310260")},
                  },
               }},
               50: {protobuf.Varint(220920022)},
               52: {protobuf.Varint(4)},
               55: {protobuf.Varint(411)},
               56: {protobuf.Varint(683)},
               61: {protobuf.Varint(3)},
               62: {protobuf.Bytes("")},
               64: {protobuf.Varint(28)},
               65: {protobuf.Fixed32(1076363264)},
               67: {protobuf.Varint(18446744073709551316)},
               77: {protobuf.Fixed32(1065353216)},
               78: {protobuf.Varint(1)},
               80: {protobuf.Bytes("America/Chicago")},
               92: {protobuf.Bytes("ranchu;")},
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
         }},
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
            protobuf.Bytes("\x12\x00"),
            protobuf.Message{
               2: {protobuf.Bytes("")},
            },
         }},
         9: {protobuf.Unknown{
            protobuf.Bytes("\n\xd2\x05\n\x02ms\x12\xcb\x05CoACZElenKPl7VFNkBvkfIXQiA6NylhnRZKMadRZrqH8ukZfD6YahUlkBLAse-JfKDH9kemp2sKZ15D6o0ardVj_djop7xIf8lJMaAyrvI3Zgs9-v1Q7LCRBCZ6DPeyQ-QzStNG_FzCtq0CPTosQWVEV6J-kse7G4UIyhdJn280swmlKAfKd4p0YFNPJUx2V3yA5t0hEDtvPE0vFfWP8IyfqODS1-VSTh94ozXT6D-B5Wag7_C1kz3xm2ArWfmHcpo3BOfLu-JC1P3ppqDCSA5GofI4tlT8aSg_EB9PqkD79NC-LYYDmaGbxSiS-jJ7hbQa4aGpLu287EvGblUpvRLzCWQqAAqkaYmpuDHlGUwG-eS23R6ObMM0xVxtqvwCFYksO-XGMDJ9kBQCng6kttO9wWWdctjEFCR6eObqmwGSaCdoxgz-3xtLh9-FOCRmdPTHh8pwMaDvJLXzUW7cpi1UINtV0bbTRoSPOLEFYF9AVED_ruAsmcAHMOb8cO5DR5ZrPlxbG--4nyrPdXEAdxJq_b8wlJyVukDnJrizw35Uf272evt9mjW3VdsphdID6JE-_PYHdX7AeIzE5ORXDeDF0F7xmSRr13sQwYY44auBEPQblSyzQ155M7Nq-IgEnIA93kI0WvJthnjd3Hqky1P87negnh3dZlBLQ7osQJZbIAIyRGJ0SEA2ffXou5JQBPmgVvpuMJhw"),
            protobuf.Message{
               1: {protobuf.Unknown{
                  protobuf.Bytes("\n\x02ms\x12\xcb\x05CoACZElenKPl7VFNkBvkfIXQiA6NylhnRZKMadRZrqH8ukZfD6YahUlkBLAse-JfKDH9kemp2sKZ15D6o0ardVj_djop7xIf8lJMaAyrvI3Zgs9-v1Q7LCRBCZ6DPeyQ-QzStNG_FzCtq0CPTosQWVEV6J-kse7G4UIyhdJn280swmlKAfKd4p0YFNPJUx2V3yA5t0hEDtvPE0vFfWP8IyfqODS1-VSTh94ozXT6D-B5Wag7_C1kz3xm2ArWfmHcpo3BOfLu-JC1P3ppqDCSA5GofI4tlT8aSg_EB9PqkD79NC-LYYDmaGbxSiS-jJ7hbQa4aGpLu287EvGblUpvRLzCWQqAAqkaYmpuDHlGUwG-eS23R6ObMM0xVxtqvwCFYksO-XGMDJ9kBQCng6kttO9wWWdctjEFCR6eObqmwGSaCdoxgz-3xtLh9-FOCRmdPTHh8pwMaDvJLXzUW7cpi1UINtV0bbTRoSPOLEFYF9AVED_ruAsmcAHMOb8cO5DR5ZrPlxbG--4nyrPdXEAdxJq_b8wlJyVukDnJrizw35Uf272evt9mjW3VdsphdID6JE-_PYHdX7AeIzE5ORXDeDF0F7xmSRr13sQwYY44auBEPQblSyzQ155M7Nq-IgEnIA93kI0WvJthnjd3Hqky1P87negnh3dZlBLQ7osQJZbIAIyRGJ0SEA2ffXou5JQBPmgVvpuMJhw"),
                  protobuf.Message{
                     1: {protobuf.Bytes("ms")},
                     2: {protobuf.Bytes("CoACZElenKPl7VFNkBvkfIXQiA6NylhnRZKMadRZrqH8ukZfD6YahUlkBLAse-JfKDH9kemp2sKZ15D6o0ardVj_djop7xIf8lJMaAyrvI3Zgs9-v1Q7LCRBCZ6DPeyQ-QzStNG_FzCtq0CPTosQWVEV6J-kse7G4UIyhdJn280swmlKAfKd4p0YFNPJUx2V3yA5t0hEDtvPE0vFfWP8IyfqODS1-VSTh94ozXT6D-B5Wag7_C1kz3xm2ArWfmHcpo3BOfLu-JC1P3ppqDCSA5GofI4tlT8aSg_EB9PqkD79NC-LYYDmaGbxSiS-jJ7hbQa4aGpLu287EvGblUpvRLzCWQqAAqkaYmpuDHlGUwG-eS23R6ObMM0xVxtqvwCFYksO-XGMDJ9kBQCng6kttO9wWWdctjEFCR6eObqmwGSaCdoxgz-3xtLh9-FOCRmdPTHh8pwMaDvJLXzUW7cpi1UINtV0bbTRoSPOLEFYF9AVED_ruAsmcAHMOb8cO5DR5ZrPlxbG--4nyrPdXEAdxJq_b8wlJyVukDnJrizw35Uf272evt9mjW3VdsphdID6JE-_PYHdX7AeIzE5ORXDeDF0F7xmSRr13sQwYY44auBEPQblSyzQ155M7Nq-IgEnIA93kI0WvJthnjd3Hqky1P87negnh3dZlBLQ7osQJZbIAIyRGJ0SEA2ffXou5JQBPmgVvpuMJhw")},
                  },
               }},
            },
         }},
      },
   }},
}
