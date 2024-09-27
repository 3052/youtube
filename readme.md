# YouTube

setup device with Android API 32. earlier versions likely fail to load. YouTube
should already be installed. Then install system certificate

~~~
adb shell am start -a android.intent.action.VIEW `
-d https://www.youtube.com/watch?v=40wkJJXfwQ0
~~~

https://play.google.com/store/apps/details?id=com.google.android.youtube

windows defender firewall:

1. private profile
2. on
3. outbound rules
4. new rule
5. port
6. UDP
7. 443
8. block the connection
9. QUIC

- https://community.zscaler.com/s/question/0D54u00009evmkVCAQ/recommended-method-to-block-quic-and-http3
- https://superuser.com/questions/1182658/how-to-block-everything-all-incoming
