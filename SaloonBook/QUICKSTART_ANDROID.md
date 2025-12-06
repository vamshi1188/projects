# Quick Start: Building Android APK

## One-Time Setup

```bash
# Run the automated setup script
chmod +x setup-android.sh
./setup-android.sh
```

OR manually:

```bash
cd frontend/web
npm install
npm run build
npx cap add android
npx cap sync
```

## Build APK

### Option 1: Android Studio (Recommended for beginners)

```bash
cd frontend/web
npx cap open android
```

Then in Android Studio:
- **Build → Build Bundle(s) / APK(s) → Build APK(s)**
- APK location: `android/app/build/outputs/apk/debug/app-debug.apk`

### Option 2: Command Line (Quick)

```bash
cd frontend/web
npm run android:build
```

APK will be at: `frontend/web/android/app/build/outputs/apk/debug/app-debug.apk`

## After Making Changes

```bash
cd frontend/web
npm run cap:sync
```

Then rebuild in Android Studio or run `npm run android:build`

## Install APK on Device

```bash
# Via USB
adb install frontend/web/android/app/build/outputs/apk/debug/app-debug.apk

# Or just copy the APK to your phone and install it
```

## Prerequisites Checklist

- ✅ Node.js 18+ installed
- ✅ Android Studio installed
- ✅ Android SDK installed (via Android Studio)
- ✅ JDK 17+ installed
- ✅ ANDROID_HOME environment variable set

## Environment Variables

Add to `~/.bashrc` or `~/.zshrc`:

```bash
export ANDROID_HOME=$HOME/Android/Sdk
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/platform-tools
```

Then run: `source ~/.bashrc` (or restart terminal)

## Troubleshooting

**Build fails?**
- Run `npx cap sync` again
- Clean project in Android Studio: **Build → Clean Project**
- Check that ANDROID_HOME is set: `echo $ANDROID_HOME`

**App crashes?**
- Check Android Studio logcat for errors
- Verify `base: './'` is in vite.config.ts

**Can't find Android SDK?**
- Install via Android Studio SDK Manager
- Set ANDROID_HOME environment variable

## Full Documentation

See [ANDROID_BUILD_GUIDE.md](./ANDROID_BUILD_GUIDE.md) for complete instructions.

