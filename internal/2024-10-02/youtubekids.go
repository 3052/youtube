package main

import (
   "net/http"
   "net/url"
   "os"
   "time"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "www.youtubekids.com"
   req.URL.Scheme = "http"
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version"}
   time.Sleep(time.Second)
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
