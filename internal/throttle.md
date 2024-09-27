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
