# Android APK Build Guide for SaloonBook

This guide will help you build a workable APK for the SaloonBook application using Android Studio and Capacitor.

## Prerequisites

Before you start, make sure you have the following installed:

1. **Node.js** (18+ recommended) - [Download](https://nodejs.org/)
2. **Android Studio** - [Download](https://developer.android.com/studio)
3. **Java Development Kit (JDK)** 17 or higher
4. **Android SDK** (installed via Android Studio)

### Android Studio Setup

1. Install Android Studio
2. Open Android Studio → More Actions → SDK Manager
3. Install the following:
   - Android SDK Platform 33 (or latest)
   - Android SDK Build-Tools
   - Android SDK Command-line Tools
   - Android Emulator (optional, for testing)
4. Configure environment variables:
   ```bash
   export ANDROID_HOME=$HOME/Android/Sdk
   export PATH=$PATH:$ANDROID_HOME/tools
   export PATH=$PATH:$ANDROID_HOME/platform-tools
   ```
   Add these to your `~/.bashrc` or `~/.zshrc` file.

## Step-by-Step Build Process

### Step 1: Install Dependencies

Navigate to the frontend web directory and install all dependencies:

```bash
cd /home/vamshi/workspace/github/projects/SaloonBook/frontend/web
npm install
```

### Step 2: Build the Web Application

Build the production version of your React app:

```bash
npm run build
```

This will create a `dist` folder with your compiled web application.

### Step 3: Initialize Capacitor (First Time Only)

If this is your first time setting up Capacitor, run:

```bash
npx cap init
```

When prompted:
- **App name**: SaloonBook
- **App ID**: com.saloonbook.app (or your preferred package name)
- **Web asset directory**: dist

**Note:** The capacitor.config.ts file is already configured, so you can skip this if the file exists.

### Step 4: Add Android Platform

Add the Android platform to your project:

```bash
npx cap add android
```

This will create an `android` folder with a complete Android project.

### Step 5: Sync Your Web App with Android

Whenever you make changes to your web app, sync them:

```bash
npm run cap:sync
```

Or manually:

```bash
npm run build
npx cap sync
npx cap copy
```

### Step 6: Open in Android Studio

Open the Android project in Android Studio:

```bash
npx cap open android
```

Or manually open Android Studio and select "Open an Existing Project", then navigate to:
```
/home/vamshi/workspace/github/projects/SaloonBook/frontend/web/android
```

### Step 7: Build APK in Android Studio

#### Option A: Debug APK (for testing)

1. In Android Studio, go to **Build → Build Bundle(s) / APK(s) → Build APK(s)**
2. Wait for the build to complete
3. Click "locate" in the notification to find your APK
4. APK location: `android/app/build/outputs/apk/debug/app-debug.apk`

#### Option B: Using Command Line (Debug)

```bash
cd /home/vamshi/workspace/github/projects/SaloonBook/frontend/web
npm run android:build
```

The debug APK will be at: `android/app/build/outputs/apk/debug/app-debug.apk`

#### Option C: Release APK (for distribution)

For a release APK, you need to sign it:

1. **Generate a Keystore** (first time only):
   ```bash
   keytool -genkey -v -keystore saloonbook-release.keystore -alias saloonbook -keyalg RSA -keysize 2048 -validity 10000
   ```
   Follow the prompts and remember your passwords!

2. **Configure Signing** in `android/app/build.gradle`:
   
   Add this before the `android` block:
   ```gradle
   def keystorePropertiesFile = rootProject.file("keystore.properties")
   def keystoreProperties = new Properties()
   if (keystorePropertiesFile.exists()) {
       keystoreProperties.load(new FileInputStream(keystorePropertiesFile))
   }
   ```

   Inside the `android` block, add:
   ```gradle
   signingConfigs {
       release {
           keyAlias keystoreProperties['keyAlias']
           keyPassword keystoreProperties['keyPassword']
           storeFile file(keystoreProperties['storeFile'])
           storePassword keystoreProperties['storePassword']
       }
   }
   buildTypes {
       release {
           signingConfig signingConfigs.release
           minifyEnabled false
           proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
       }
   }
   ```

3. **Create keystore.properties** in the `android` folder:
   ```properties
   storeFile=/path/to/saloonbook-release.keystore
   storePassword=your_store_password
   keyAlias=saloonbook
   keyPassword=your_key_password
   ```

4. **Build Release APK**:
   ```bash
   npm run android:release
   ```
   
   Or in Android Studio: **Build → Build Bundle(s) / APK(s) → Build APK(s)** with "release" variant selected.

5. Release APK location: `android/app/build/outputs/apk/release/app-release.apk`

## Quick Commands Summary

```bash
# First time setup
cd /home/vamshi/workspace/github/projects/SaloonBook/frontend/web
npm install
npx cap add android

# Build debug APK
npm run android:build

# Open in Android Studio
npx cap open android

# After making changes to web app
npm run cap:sync
```

## Testing the APK

### Install on a Physical Device

1. Enable USB debugging on your Android device
2. Connect via USB
3. Run:
   ```bash
   adb install android/app/build/outputs/apk/debug/app-debug.apk
   ```

### Install on Emulator

1. Start an Android emulator from Android Studio
2. Drag and drop the APK file onto the emulator window
3. Or use adb:
   ```bash
   adb install android/app/build/outputs/apk/debug/app-debug.apk
   ```

## Troubleshooting

### Build Fails

- Make sure you have Android SDK installed
- Check that `ANDROID_HOME` is set correctly
- Run `npx cap sync` again
- Clean the project: In Android Studio, **Build → Clean Project**

### App Crashes on Launch

- Check that `base: './'` is set in `vite.config.ts`
- Ensure all assets are being bundled correctly
- Check browser console in Android Studio's logcat

### Cannot Find Android SDK

```bash
export ANDROID_HOME=$HOME/Android/Sdk
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/platform-tools
```

### Gradle Build Fails

- Update Gradle: In Android Studio, **File → Project Structure → Project**
- Accept licenses: `$ANDROID_HOME/cmdline-tools/latest/bin/sdkmanager --licenses`

## Customizing Your App

### Change App Icon

Replace these files in `android/app/src/main/res/`:
- `mipmap-hdpi/ic_launcher.png` (72x72)
- `mipmap-mdpi/ic_launcher.png` (48x48)
- `mipmap-xhdpi/ic_launcher.png` (96x96)
- `mipmap-xxhdpi/ic_launcher.png` (144x144)
- `mipmap-xxxhdpi/ic_launcher.png` (192x192)

### Change App Name

Edit `android/app/src/main/res/values/strings.xml`:
```xml
<string name="app_name">SaloonBook</string>
```

### Change Package Name

Edit `capacitor.config.ts`:
```typescript
appId: 'com.yourcompany.saloonbook'
```

Then run:
```bash
npx cap sync
```

## Distribution

### Google Play Store

1. Build a signed release APK (or AAB)
2. Create a Google Play Developer account
3. Create a new app in the Play Console
4. Upload your APK/AAB
5. Fill in store listing details
6. Submit for review

### Direct Distribution

1. Build a signed release APK
2. Upload to your website or file-sharing service
3. Users must enable "Install from Unknown Sources" on their devices

## Additional Resources

- [Capacitor Documentation](https://capacitorjs.com/docs)
- [Android Developer Guide](https://developer.android.com/guide)
- [Vite Build Guide](https://vitejs.dev/guide/build.html)

## Support

If you encounter issues:
1. Check the Android Studio logcat for errors
2. Verify all prerequisites are installed
3. Ensure environment variables are set correctly
4. Try cleaning and rebuilding the project

Good luck with your Android build! 🚀

