# Android Build Checklist

Use this checklist to ensure your Android APK build will work correctly.

## ✅ Prerequisites Installed

- [ ] Node.js 18+ installed (`node --version`)
- [ ] npm installed (`npm --version`)
- [ ] Android Studio installed
- [ ] Android SDK installed (check SDK Manager in Android Studio)
- [ ] JDK 17+ installed (`java -version`)

## ✅ Environment Variables

- [ ] ANDROID_HOME is set (`echo $ANDROID_HOME` should show path)
- [ ] Android SDK tools in PATH
- [ ] Android platform-tools in PATH

Add to `~/.bashrc` or `~/.zshrc`:
```bash
export ANDROID_HOME=$HOME/Android/Sdk
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/platform-tools
```

## ✅ Project Setup

- [ ] Dependencies installed (`npm install` in frontend/web)
- [ ] Capacitor config exists (frontend/web/capacitor.config.ts)
- [ ] Package.json has Capacitor packages
- [ ] Vite config has `base: './'` setting

## ✅ Build Process

### First Time Setup
- [ ] Run `npm install` in frontend/web
- [ ] Run `npm run build` successfully
- [ ] Run `npx cap add android` (creates android folder)
- [ ] Run `npx cap sync` (syncs web assets)

### Subsequent Builds
- [ ] Make changes to your React app
- [ ] Run `npm run build`
- [ ] Run `npx cap sync` or `npm run cap:sync`
- [ ] Rebuild in Android Studio or run `npm run android:build`

## ✅ Verify Files Created

After `npx cap add android`, check these exist:
- [ ] frontend/web/android/ folder
- [ ] frontend/web/android/app/
- [ ] frontend/web/android/build.gradle
- [ ] frontend/web/android/app/build.gradle
- [ ] frontend/web/android/app/src/main/AndroidManifest.xml

## ✅ Build APK

Choose one method:

### Method 1: Android Studio (Recommended)
- [ ] Run `npx cap open android`
- [ ] Wait for Gradle sync to complete
- [ ] Go to Build → Build Bundle(s) / APK(s) → Build APK(s)
- [ ] Wait for build to complete
- [ ] Find APK at: `android/app/build/outputs/apk/debug/app-debug.apk`

### Method 2: Command Line
- [ ] Run `npm run android:build`
- [ ] Wait for build to complete
- [ ] Find APK at: `android/app/build/outputs/apk/debug/app-debug.apk`

## ✅ Test APK

- [ ] APK file exists and has reasonable size (> 10MB typically)
- [ ] Install on device: `adb install path/to/app-debug.apk`
- [ ] App launches without crashing
- [ ] App UI appears correctly
- [ ] All images/assets load properly
- [ ] Navigation works
- [ ] LocalStorage persists data

## ✅ Common Issues to Check

If build fails:
- [ ] Check Android Studio logcat for errors
- [ ] Run `npx cap sync` again
- [ ] Clean project in Android Studio (Build → Clean Project)
- [ ] Check Gradle version compatibility
- [ ] Verify all SDK packages are downloaded

If app crashes:
- [ ] Check `base: './'` in vite.config.ts
- [ ] Check that all assets are in dist folder after build
- [ ] Look for errors in Android Studio logcat
- [ ] Verify AndroidManifest.xml has proper permissions

If assets don't load:
- [ ] Check vite.config.ts has correct asset paths
- [ ] Verify images are copied to dist/assets folder
- [ ] Check that `@assets` alias is working
- [ ] Run `npm run build` before `npx cap sync`

## ✅ Release Build (Optional)

For production release APK:
- [ ] Generate keystore file
- [ ] Create keystore.properties file
- [ ] Update build.gradle with signing config
- [ ] Run `npm run android:release`
- [ ] Find release APK at: `android/app/build/outputs/apk/release/app-release.apk`

## 📝 Quick Commands Reference

```bash
# Setup (first time)
cd frontend/web
npm install
npm run build
npx cap add android
npx cap sync

# Build APK (command line)
npm run android:build

# Open in Android Studio
npx cap open android

# After making changes
npm run cap:sync

# Install on device
adb install android/app/build/outputs/apk/debug/app-debug.apk
```

## 🆘 Need Help?

See the full guides:
- [QUICKSTART_ANDROID.md](./QUICKSTART_ANDROID.md)
- [ANDROID_BUILD_GUIDE.md](./ANDROID_BUILD_GUIDE.md)

Or check common issues:
- Android SDK not found → Set ANDROID_HOME
- Build fails → Run `npx cap sync` and clean project
- App crashes → Check logcat and verify base path in vite.config.ts

