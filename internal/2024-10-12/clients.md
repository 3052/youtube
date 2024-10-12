# clients

## `/videoplayback` requires `pot`

56 `WEB_EMBEDDED_PLAYER`:

https://www.youtube.com/embed/x4qKS3UQXFk

76 `WEB_KIDS`:

https://www.youtubekids.com

## only returns one video option

62 `WEB_CREATOR`:

https://studio.youtube.com

## signatureCipher

2 MWEB:

~~~
general.useragent.override
~~~

https://whatismybrowser.com/guides/how-to-change-your-user-agent/firefox

then:

https://m.youtube.com/watch?v=KwRxeZ9Ro24

~~~
"signatureCipher": "s=%3D%3D%3DgMNosw63xQJg-J41zzo9qJJ_v9vrkEnIkzDVidlxV7bCQI..."
~~~

67 `WEB_REMIX`:

https://music.youtube.com

or:

https://play.google.com/store/apps/details?id=com.google.android.apps.youtube.music.pwa

fail:

~~~
"signatureCipher": "s=%3D%3D%3DgMNosw63xQJg-J41zzo9qJJ_v9vrkEnIkzDVidlxV7bCQI..."
~~~
