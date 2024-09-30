# throttle

if we take:

~~~
[0].playerResponse.streamingData.adaptiveFormats[0].url
~~~

its throttled:

~~~
> curl -o o 'https://rr1---sn-q4fl6n66.googlevideo.com/videoplayback?expire=1...'
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                               Dload  Upload   Total   Spent    Left  Speed
0 1861M    0 1024k    0     0   203k      0  2:36:13  0:00:05  2:36:08  204k
~~~

this works:

~~~
> curl -o o 'https://rr2---sn-q4flrn7y.googlevideo.com/videoplayback?range=0-9999999'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 9765k  100 9765k    0     0  13.4M      0 --:--:-- --:--:-- --:--:-- 13.4M
~~~

this works:

~~~
> curl -o o -H range:bytes=0-9999999 https://rr2---sn-q4flrn7y.googlevideo.com/videoplayback
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 9765k  100 9765k    0     0  13.3M      0 --:--:-- --:--:-- --:--:-- 13.3M
~~~

## what does the web client do?

## what does the Android client do?

## other research

~~~
videoPlaybackUstreamerConfig
~~~

https://github.com/LuanRT/googlevideo/blob/main/examples/downloader/main.ts
