# clients

## 0 `UNKNOWN_INTERFACE`

I dont know what client this is

## 1 WEB

ump:

~~~
"POST /videoplayback?expire=1727646217&ei=qXX5ZtvPEv2Z2_gPjerBuQg&ip=72.181.16.91&id=o-AOjx3hN1pSaRRd3_PIAj40DuwR1fxvZ5WFVqdLSpc0R8&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=am&mm=31%2C29&mn=sn-q4fl6n66%2Csn-q4flrn7y&ms=au%2Crdu&mv=m&mvi=1&pl=19&initcwndbps=1088750&spc=54MbxcPragR4oT0P2DMgTIiq-YEXZD2NztrO5K-ISHw6gs-5S0Bh58rgKJCH&svpuc=1&ns=1t-Z2QnTKpGR6CZCglhNWvMQ&sabr=1&rqh=1&mt=1727624226&fvip=2&keepalive=yes&fexp=51299152&c=WEB&n=78_onccMq-WDBw&sparams=expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Cxpc%2Cspc%2Csvpuc%2Cns%2Csabr%2Crqh&sig=AJfQdSswRAIgJyZWdWRfpdJmRVhZduPvhaxi5DYr0Gp427j1_wwMDQMCICXlvg1Cxnbf4ydySG0I6N9AWuuKQyWbLvfdC0NZmCWI&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=ABPmVW0wRQIhAP2cVPxzbOFWNxXuJiDTq1kTPhTMCq8vuaUGvc5A1w73AiBYGMOlenrh--07zMjL5N9Zn9yN-aEmg1G_XB_Qgz5ErA%3D%3D&cpn=0I5Kau1O22UZFNbW&cver=2.20240926.00.00&rn=1 HTTP/1.1"
"Host: rr1---sn-q4fl6n66.googlevideo.com"
"User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0"
"Accept: */*"
"Accept-Language: en-US,en;q=0.5"
"Accept-Encoding: gzip, deflate, br"
"Content-Length: 1903"
"Origin: https://www.youtube.com"
"DNT: 1"
"Sec-GPC: 1"
"Connection: keep-alive"
"Referer: https://www.youtube.com/"
"Sec-Fetch-Dest: empty"
"Sec-Fetch-Mode: cors"
"Sec-Fetch-Site: cross-site"
"Pragma: no-cache"
"Cache-Control: no-cache"
""
"\n\xa1\x80\x00\x90\x01\xf5\x06\x98\x01\xf1\x03\xa8\x01\x00\xe0\x01\x00\x90..."
"\b\b\xb0\t\x10\xb0\t \x01 \x88'\xda\x04\v\n\x06\b\xf0.\x10\xf0. \xf0.\xf8..."
~~~

## 2 MWEB

~~~
general.useragent.override
~~~

https://whatismybrowser.com/guides/how-to-change-your-user-agent/firefox

then:

https://m.youtube.com/watch?v=KwRxeZ9Ro24

fail:

~~~
"signatureCipher": "s=%3D%3D%3DgMNosw63xQJg-J41zzo9qJJ_v9vrkEnIkzDVidlxV7bCQI..."
~~~

## 3 ANDROID

ump:

~~~
"POST /videoplayback?expire=1724652395&ei=CsfLZoO2O4rkir4P0taY8Q0&ip=72.181.16.91&id=o-ACCjhjANFZlVfaNe1JZwKleNFPNIbmQuWMovDELsK8X6&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=IN&mm=31%2C26&mn=sn-q4fl6n6y%2Csn-a5mlrnlz&ms=au%2Conr&mv=m&mvi=1&pl=23&initcwndbps=1342500&spc=Mv1m9kC_OK5E_6M6kNGwgFszAnm4Jv2jE_OMHDkzv40Y1AC2U3IrojBc8A&svpuc=1&sabr=1&rqh=1&mt=1724630451&fvip=3&keepalive=yes&c=ANDROID&sparams=expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Cxpc%2Cspc%2Csvpuc%2Csabr%2Crqh&sig=AJfQdSswRQIgJ4OQfR9D-q0hAxvOPzUeyGZsH6byIPIob-O_71y7GAYCIQClzAXUSGu0JU2-0LuEvB9CJwfrdoVXDivGzgxG2ATARA%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AGtxev0wRAIgPl3jlEueiwtDrep2WhsrCA_MlA6uhSGCpbdaid95p7ICIB4Ta9RzhuoievlP9_rnwcu1N_VosN8qnxzr_H_rYyoA&cpn=Wlf_0-AoX2uquEvg&rn=6 HTTP/1.1"
"Host: rr1---sn-q4fl6n6y.googlevideo.com"
"Connection: keep-alive"
"Content-Length: 2325"
"Content-Type: application/x-protobuf"
"User-Agent: com.google.android.youtube/1547951552 (Linux; U; Android 9; en_US; Android SDK built for x86; Build/PSR1.180720.012; Cronet/129.0.6614.4)"
"Accept-Encoding: gzip, deflate, br"
""
"\n9\x88\x01t\x90\x01\xb7\b\x98\x01\xdf\x04\xb0\x01\x01\xb8\x01\xbb\xf3c\xd0..."
~~~

https://play.google.com/store/apps/details?id=com.google.android.youtube

## 5 IOS

I dont know how to capture this client

## 7 TVHTML5

~~~
general.useragent.override
Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version
~~~

https://www.youtube.com/tv

<https://github.com/youtube/cobalt/blob/25.lts.10/cobalt/browser/user_agent_string.cc>

ump:

~~~
"POST /videoplayback?expire=1727768530&ei=clP7ZpWkN5G2ir4PismO8Qc&ip=72.181.16.91&id=o-AIaY6w_YnVWWgJC83fAateS4FhWoClal5hEFgx1hazxI&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=vE&mm=31%2C26&mn=sn-q4fl6nss%2Csn-a5mekndl&ms=au%2Conr&mv=m&mvi=4&pl=19&pcm2=no&gcr=us&initcwndbps=1048750&svpuc=1&ns=VWS3EmKFn86onDqTKK_8K2EQ&sabr=1&rqh=1&mt=1727746616&fvip=3&keepalive=yes&fexp=51299152%2C51300761&c=TVHTML5&n=Z_evkyL-2Vjd0w&sparams=expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Cxpc%2Cpcm2%2Cgcr%2Csvpuc%2Cns%2Csabr%2Crqh&sig=AJfQdSswRQIgcCDKJ8lEzv18FDjCPrsN-QMJY8JIuaRI0SYwA7WSIqQCIQCnG-B6Ma1BREzaHxGRDqfE4yjUS-Nv5xgjLKibiOc0WQ%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=ABPmVW0wRgIhAMOfKRjcpdSx-s4IPGdxxx5sD5anG851b1xqvhJLKeMqAiEA8MM7gBWCLPHUKYWWQDJovARdR3UpK_tRY6JXIvrelFk%3D&cpn=BHPZePgvGfCR30N0&cver=7.20240929.15.00&rn=1 HTTP/1.1"
"Host: rr4---sn-q4fl6nss.googlevideo.com"
"User-Agent: Mozilla/5.0 (ChromiumStylePlatform) Cobalt/Version"
"Accept: */*"
"Accept-Language: en-US,en;q=0.5"
"Accept-Encoding: gzip, deflate, br"
"Referer: https://www.youtube.com/"
"Content-Length: 1741"
"Origin: https://www.youtube.com"
"DNT: 1"
"Sec-GPC: 1"
"Connection: keep-alive"
"Sec-Fetch-Dest: empty"
"Sec-Fetch-Mode: cors"
"Sec-Fetch-Site: cross-site"
"Pragma: no-cache"
"Cache-Control: no-cache"
""
"\n\xa3\x80\x00\x90\x01\xc5\v\x98\x01\x93\x06\xa8\x01\x00\xe0\x01\x00\x90\x02..."
"\b\b\xb0\t\x10\xb0\t \x01 \x88'\xda\x04\v\n\x06\b\xf0.\x10\xf0. \xf0.\xf8\x04..."
~~~

## 8 TVLITE

I dont know what client this is

## 10 TVANDROID

I dont know what client this is

## 11 XBOX

I dont know how to capture this client

## 12 CLIENTX

I dont know what client this is

## 13 XBOXONEGUIDE

I dont know how to capture this client

## 14 `ANDROID_CREATOR`

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.creator

if you try to play a video it just opens another app

## 15 `IOS_CREATOR`

I dont know how to capture this client

## 16 TVAPPLE

I dont know how to capture this client

## 17 IOS\_INSTANT

I dont know how to capture this client

## 18 `ANDROID_KIDS`

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.kids

ump:

~~~
"POST /videoplayback?expire=1727930103&ei=l8r9ZsyPOe3IzN0PkM6Y8AM&ip=72.181.16.91&id=o-AEMRc2lRRCIZLPL7Z7Ie8AVo-EJPYum9veez0urzNGxb&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=rj&mm=31%2C26&mn=sn-q4flrnsk%2Csn-a5meknde&ms=au%2Conr&mv=m&mvi=4&pl=19&pcm2=yes&initcwndbps=1233750&svpuc=1&sabr=1&rqh=1&mt=1727908136&fvip=5&keepalive=yes&fexp=51300760&c=ANDROID_KIDS&sparams=expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Cxpc%2Cpcm2%2Csvpuc%2Csabr%2Crqh&sig=AJfQdSswRQIhAIUd41mWRwIfm3QuETiyxZTxpxw5cYaPp1ZtHp31n2oWAiBCoTpA0IJkBesHwb9sbUfYh57_ImHFAY0gDzFfgtrRGg%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=ABPmVW0wRQIgRl8U-S70jy-IOS7l4q6a_RUNvdtYT9Is7iJA7gZN74ECIQCSg5SY3bA9V9m8zfCQIpAjiqQGoL2sHbN3-b1Vgm7EpA%3D%3D&cpn=ipy1CXXvnPS9vgJv&rn=2 HTTP/1.1"
"Host: rr4---sn-q4flrnsk.googlevideo.com"
"Connection: keep-alive"
"Content-Length: 1858"
"Content-Type: application/x-protobuf"
"User-Agent: com.google.android.apps.youtube.kids/153801030 (Linux; U; Android 9; en_US; AOSP on IA Emulator; Build/PSR1.180720.122; Cronet/129.0.6614.4)"
"Accept-Encoding: gzip, deflate, br"
""
"\n6\x88\x01t\x90\x01\xff\x0e\x98\x01\xb8\b\xb0\x01\x00\xb8\x01\xb4\xa0\x81\x..."
~~~

## 19 IOS\_KIDS

I dont know how to capture this client

## 20 `ANDROID_INSTANT`

I dont know what client this is

## 21 `ANDROID_MUSIC`

https://play.google.com/store/apps/details?id=com.google.android.youtube.tvmusic

wont run on virtual device, even with proxy disabled. tried Android 9

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.music

ump:

~~~
"POST /videoplayback?expire=1727931210&ei=6s79ZpvJFqK0ir4PyryJmA0&ip=72.181.16.91&id=o-AE-D7jiwQnqG9sPkd2KltnaNQGlKvYSwPGQGwGGJwBeg&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=xj&mm=31%2C26&mn=sn-q4flrnez%2Csn-a5mlrnll&ms=au%2Conr&mv=m&mvi=4&pl=19&ctier=A&pfa=5&gcr=us&initcwndbps=1240000&hightc=yes&siu=1&spc=54MbxQu13pamnn0c5dd1KqWEh1ljwnnODzH8clzS8HZW2LTzMM9IA_pvKWvh&svpuc=1&sabr=1&rqh=1&mt=1727909099&fvip=5&keepalive=yes&fexp=51300760&c=ANDROID_MUSIC&sparams=expire%2Cei%2Cip%2Cid%2Csource%2Crequiressl%2Cxpc%2Cctier%2Cpfa%2Cgcr%2Chightc%2Csiu%2Cspc%2Csvpuc%2Csabr%2Crqh&sig=AJfQdSswRQIhAL-RAWRbUt3wXuh9nWCQZJh4m0h1xTHbwZxqz-Q2NcFaAiBDnOkVsM508tRUsvTGD_BjRKpVGwvmkcKsjWlYIQn2eQ%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=ABPmVW0wRQIhAIfAnsyd2FcfZ4nSkWELEzQrY6t8PEX0gmFdQBXE39u6AiAjSnVzcVZR98q--aAVYG3auWe10Ggp4GeVA_WtxQyZWw%3D%3D&cpn=UBoejYYSjpWH11j-&rn=12 HTTP/1.1"
"Host: rr4---sn-q4flrnez.googlevideo.com"
"Connection: keep-alive"
"Content-Length: 1754"
"Content-Type: application/x-protobuf"
"User-Agent: com.google.android.apps.youtube.music/72150270 (Linux; U; Android 9; en_US; AOSP on IA Emulator; Build/PSR1.180720.122; Cronet/129.0.6614.4)"
"Accept-Encoding: gzip, deflate, br"
""
"\n6\x88\x01t\x90\x01\xa0\x06\x98\x01\xa0\x06\xb0\x01\x01\xb8\x01\x96\xfe\x97..."
~~~

## 22 IOS\_TABLOID

I dont know how to capture this client

## 23 `ANDROID_TV`

https://play.google.com/store/apps/details?id=com.google.android.youtube.tv

error after capturing two requests

## 24 `ANDROID_GAMING`

https://apkmirror.com/apk/google-inc/youtube-gaming

discontinued:

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.gaming

## 25 IOS\_GAMING

I dont know how to capture this client

## 26 IOS\_MUSIC

I dont know how to capture this client

## 27 MWEB\_TIER\_2

I dont know what client this is

## 28 `ANDROID_VR`

https://apkmirror.com/apk/google-inc/youtube-vr-daydream

discontinued:

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.vr

## 29 `ANDROID_UNPLUGGED`

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged

then:

https://apkmirror.com/apk/google-inc/youtube-tv-watch-record-tv

cost money

## 30 `ANDROID_TESTSUITE`

I dont know what client this is

## 31 `WEB_MUSIC_ANALYTICS`

I dont know what client this is

## 32 `WEB_GAMING`

I dont know what client this is

## 33 `IOS_UNPLUGGED`

I dont know how to capture this client

## 34 `ANDROID_WITNESS`

I dont know what client this is

## 35 `IOS_WITNESS`

I dont know how to capture this client

## 36 `ANDROID_SPORTS`

I dont know what client this is

## 37 `IOS_SPORTS`

I dont know how to capture this client

## 38 `ANDROID_LITE`

https://apkmirror.com/apk/google-inc/youtube-go

discontinued:

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.mango

## 39 `IOS_EMBEDDED_PLAYER`

I dont know how to capture this client

## 40 `IOS_DIRECTOR`

I dont know how to capture this client

## 41 `WEB_UNPLUGGED`

cost money

## 42 `WEB_EXPERIMENTS`

I dont know what client this is

## 43 `TVHTML5_CAST`

I dont know what client this is

## 53 `IOS_PILOT_STUDIO`

I dont know how to capture this client

## 54 `ANDROID_CASUAL`

I dont know what client this is

## 55 `ANDROID_EMBEDDED_PLAYER`

I dont know what client this is

## 56 `WEB_EMBEDDED_PLAYER`

`/videoplayback` requires `pot`

## 57 `TVHTML5_AUDIO`

I dont know what client this is

## 58 `TV_UNPLUGGED_CAST`

cost money

## 59 `TVHTML5_KIDS`

I dont know what client this is

## 60 `WEB_HEROES`

I dont know what client this is

## 61 `WEB_MUSIC`

I dont know what client this is

## 62 `WEB_CREATOR`

only returns one video option

https://studio.youtube.com

## 63 `TV_UNPLUGGED_ANDROID`

cost money

https://play.google.com/store/apps/details?id=com.google.android.youtube.tvunplugged

## 64 `IOS_LIVE_CREATION_EXTENSION`

I dont know how to capture this client

## 65 `TVHTML5_UNPLUGGED`

cost money

## 66 `IOS_MESSAGES_EXTENSION`

I dont know how to capture this client

## 67 `WEB_REMIX`

https://music.youtube.com

or:

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.music.pwa

fail:

~~~
"signatureCipher": "s=%3D%3D%3DgMNosw63xQJg-J41zzo9qJJ_v9vrkEnIkzDVidlxV7bCQI..."
~~~

## 68 `IOS_UPTIME`

I dont know how to capture this client

## 69 `WEB_UNPLUGGED_ONBOARDING`

cost money

## 70 `WEB_UNPLUGGED_OPS`

cost money

## 71 `WEB_UNPLUGGED_PUBLIC`

cost money

## 72 `TVHTML5_VR`

I dont know what client this is

## 73 `WEB_LIVE_STREAMING`

I dont know what client this is

## 74 `ANDROID_TV_KIDS`

https://play.google.com/store/apps/details?id=com.google.android.youtube.tvkids

error after capturing two requests

## 75 `TVHTML5_SIMPLY`

I dont know what client this is

## 76 `WEB_KIDS`

`/videoplayback` requires `pot`

## 77 `MUSIC_INTEGRATIONS`

I dont know what client this is

## 80 `TVHTML5_YONGLE`

I dont know what client this is

## 84 `GOOGLE_ASSISTANT`

I dont know what client this is

## 85 `TVHTML5_SIMPLY_EMBEDDED_PLAYER`

I dont know what client this is

## 86 `WEB_MUSIC_EMBEDDED_PLAYER`

I dont know what client this is

## 87 `WEB_INTERNAL_ANALYTICS`

I dont know what client this is

## 88 `WEB_PARENT_TOOLS`

I dont know what client this is

## 89 `GOOGLE_MEDIA_ACTIONS`

I dont know what client this is

## 90 `WEB_PHONE_VERIFICATION`

I dont know what client this is

## 91 `ANDROID_PRODUCER`

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.producer

unable to sign in

## 92 `IOS_PRODUCER`

I dont know how to capture this client

## 93 `TVHTML5_FOR_KIDS`

I dont know what client this is

## 94 `GOOGLE_LIST_RECS`

I dont know what client this is

## 95 `MEDIA_CONNECT_FRONTEND`

I dont know what client this is

## 98 `WEB_EFFECT_MAKER`

I dont know what client this is

## 99 `WEB_SHOPPING_EXTENSION`

I dont know what client this is

## 100 `WEB_PLAYABLES_PORTAL`

I dont know what client this is
