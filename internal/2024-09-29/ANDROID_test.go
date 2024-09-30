package youtube

import (
   "testing"
)

func TestAndroidCreator(t *testing.T) {
   const name = "ANDROID_CREATOR"
   version, err := appVersion("com.google.android.apps.youtube.creator", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidKids(t *testing.T) {
   const name = "ANDROID_KIDS"
   version, err := appVersion("com.google.android.apps.youtube.kids", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidLite(t *testing.T) {
   const name = "ANDROID_LITE"
   version, err := appVersion("com.google.android.apps.youtube.mango", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidMusic(t *testing.T) {
   const name = "ANDROID_MUSIC"
   version, err := appVersion("com.google.android.apps.youtube.music", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidCast(t *testing.T) {
   const name = "TVHTML5_CAST"
   version, err := appVersion(
      "com.google.android.apps.youtube.music.pwa", tablet,
   )
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidUnplugged(t *testing.T) {
   const name = "ANDROID_UNPLUGGED"
   version, err := appVersion("com.google.android.apps.youtube.unplugged", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidVr(t *testing.T) {
   const name = "ANDROID_VR"
   version, err := appVersion("com.google.android.apps.youtube.vr", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroid(t *testing.T) {
   const name = "ANDROID"
   version, err := appVersion("com.google.android.youtube", phone)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidTv(t *testing.T) {
   const name = "ANDROID_TV"
   version, err := appVersion("com.google.android.youtube.tv", tv)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidTvKids(t *testing.T) {
   const name = "ANDROID_TV_KIDS"
   version, err := appVersion("com.google.android.youtube.tvkids", tv)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}

func TestAndroidTvunplug(t *testing.T) {
   const name = "TV_UNPLUGGED_ANDROID"
   version, err := appVersion("com.google.android.youtube.tvunplugged", tv)
   if err != nil {
      t.Fatal(err)
   }
   if version != names[name] {
      t.Fatal(name, version)
   }
}
