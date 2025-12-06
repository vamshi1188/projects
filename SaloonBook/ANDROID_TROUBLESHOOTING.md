# Android Build Troubleshooting Guide

Common issues and their solutions when building the SaloonBook Android APK.

## Environment Setup Issues

### Issue: "ANDROID_HOME not found"

**Symptoms:**
```
Error: ANDROID_HOME is not set
```

**Solution:**
```bash
# Find your Android SDK location (usually one of these):
# - $HOME/Android/Sdk
# - $HOME/Library/Android/sdk (macOS)
# - /usr/local/android-sdk

# Set environment variable
export ANDROID_HOME=$HOME/Android/Sdk
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/platform-tools

# Add to ~/.bashrc or ~/.zshrc to make permanent
echo 'export ANDROID_HOME=$HOME/Android/Sdk' >> ~/.bashrc
echo 'export PATH=$PATH:$ANDROID_HOME/tools' >> ~/.bashrc
echo 'export PATH=$PATH:$ANDROID_HOME/platform-tools' >> ~/.bashrc

# Reload
source ~/.bashrc
```

### Issue: "Java/JDK not found"

**Symptoms:**
```
Error: JAVA_HOME is not set
Could not find java executable
```

**Solution:**
```bash
# Install JDK 17 or higher
sudo apt install openjdk-17-jdk  # Ubuntu/Debian
brew install openjdk@17          # macOS

# Set JAVA_HOME
export JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64  # Ubuntu
export JAVA_HOME=/opt/homebrew/opt/openjdk@17         # macOS
```

### Issue: "Android SDK licenses not accepted"

**Symptoms:**
```
You have not accepted the license agreements
```

**Solution:**
```bash
cd $ANDROID_HOME/cmdline-tools/latest/bin
./sdkmanager --licenses
# Press 'y' to accept all licenses
```

## Build Issues

### Issue: "Gradle sync failed"

**Symptoms:**
- Android Studio shows "Gradle sync failed"
- Build won't start

**Solution:**
1. In Android Studio: **File → Invalidate Caches / Restart**
2. Delete Gradle cache:
   ```bash
   rm -rf ~/.gradle/caches/
   ```
3. In Android Studio: **Build → Clean Project**
4. **Build → Rebuild Project**

### Issue: "npm run build fails"

**Symptoms:**
```
Error: Cannot find module 'vite'
TypeScript compilation errors
```

**Solution:**
```bash
cd frontend/web

# Clean install
rm -rf node_modules package-lock.json
npm install

# Try build again
npm run build
```

### Issue: "android folder not created"

**Symptoms:**
```
Error: android platform not found
```

**Solution:**
```bash
cd frontend/web

# Make sure build succeeds first
npm run build

# Add android platform
npx cap add android

# If it says already exists but folder missing:
rm -rf android .capacitor
npx cap add android
```

### Issue: "gradlew: Permission denied"

**Symptoms:**
```
bash: ./gradlew: Permission denied
```

**Solution:**
```bash
cd frontend/web/android
chmod +x gradlew
./gradlew assembleDebug
```

### Issue: "Build tools version not found"

**Symptoms:**
```
Failed to find Build Tools revision XX.X.X
```

**Solution:**
1. Open Android Studio
2. **Tools → SDK Manager**
3. **SDK Tools** tab
4. Check **Android SDK Build-Tools**
5. Click **Apply** to install

## Runtime Issues

### Issue: "App crashes immediately on launch"

**Symptoms:**
- App opens briefly then closes
- White screen then crash

**Solution 1 - Check base path:**
```typescript
// frontend/web/vite.config.ts
export default defineConfig({
  // ... other config
  base: './',  // ← Must be './'
});
```

**Solution 2 - Rebuild and sync:**
```bash
cd frontend/web
rm -rf dist
npm run build
npx cap sync
# Rebuild in Android Studio
```

**Solution 3 - Check Android Studio logcat:**
1. Open Android Studio
2. **View → Tool Windows → Logcat**
3. Filter for your app package
4. Look for red error messages

### Issue: "Images/assets not loading"

**Symptoms:**
- App works but images are broken
- Console shows 404 errors for assets

**Solution:**
```bash
cd frontend/web

# Check vite config has correct alias
# frontend/web/vite.config.ts should have:
# alias: {
#   '@assets': path.resolve(__dirname, '../../attached_assets'),
# }

# Rebuild
npm run build

# Check dist folder has assets
ls -la dist/assets/

# Sync again
npx cap sync
```

### Issue: "White screen / blank app"

**Symptoms:**
- App launches but shows blank white screen
- No UI visible

**Solution:**
```bash
# Check browser console in Chrome DevTools
# 1. Connect phone via USB
# 2. Open Chrome on computer
# 3. Go to chrome://inspect
# 4. Find your app and click "inspect"

# Common fixes:
cd frontend/web

# Ensure proper build
npm run build

# Check capacitor config
cat capacitor.config.ts
# Should have: webDir: 'dist'

# Full rebuild
npx cap sync
npx cap copy
```

## Capacitor Issues

### Issue: "capacitor not found"

**Symptoms:**
```
'cap' is not recognized
npx: command not found: cap
```

**Solution:**
```bash
cd frontend/web

# Install Capacitor CLI
npm install @capacitor/cli --save-dev

# Or reinstall all dependencies
npm install
```

### Issue: "webDir does not exist"

**Symptoms:**
```
Error: webDir (dist) does not exist
```

**Solution:**
```bash
cd frontend/web

# Build first!
npm run build

# Then sync
npx cap sync
```

### Issue: "Multiple apps with same package name"

**Symptoms:**
```
Error: Package already exists
```

**Solution:**
Edit `capacitor.config.ts`:
```typescript
const config: CapacitorConfig = {
  appId: 'com.saloonbook.app.v2', // Change this
  // ...
};
```

Then:
```bash
npx cap sync
```

## Android Studio Issues

### Issue: "Android Studio won't open project"

**Solution:**
1. Don't open `frontend/web` folder
2. Open specifically: `frontend/web/android` folder
3. Or use: `npx cap open android` from `frontend/web` directory

### Issue: "Gradle build very slow"

**Solution:**
Edit `android/gradle.properties`:
```properties
org.gradle.daemon=true
org.gradle.parallel=true
org.gradle.jvmargs=-Xmx4096m -XX:MaxMetaspaceSize=512m
```

### Issue: "Emulator not starting"

**Solution:**
1. **Tools → Device Manager**
2. Create new virtual device
3. Choose a recent API level (33+)
4. Select x86_64 image for better performance

## APK Installation Issues

### Issue: "App not installed"

**Symptoms:**
- APK copied to phone but won't install
- "App not installed" error

**Solution:**
1. Enable **Unknown Sources** / **Install Unknown Apps**
2. Uninstall old version first if exists
3. Check APK is not corrupted (should be >10MB)

### Issue: "adb: device not found"

**Symptoms:**
```
error: no devices/emulators found
```

**Solution:**
```bash
# Enable USB debugging on phone
# Connect via USB

# Check device is detected
adb devices

# If not listed:
adb kill-server
adb start-server
adb devices

# On phone: Allow USB debugging when prompted
```

## Performance Issues

### Issue: "APK size very large (>100MB)"

**Solution:**
```bash
# Use release build instead of debug
npm run android:release

# Enable minification in android/app/build.gradle:
# buildTypes {
#     release {
#         minifyEnabled true
#         shrinkResources true
#     }
# }
```

### Issue: "App runs slow on device"

**Solution:**
1. Build release APK instead of debug
2. Check logcat for warnings
3. Profile app in Android Studio
4. Optimize images (compress, resize)

## Data Persistence Issues

### Issue: "Data doesn't persist between sessions"

**Symptoms:**
- Bookings disappear after closing app
- Settings reset

**Solution:**
- LocalStorage should work automatically
- Check browser console for errors
- Verify data is being written:
  ```javascript
  // In Chrome DevTools connected to device
  localStorage.getItem('your-key')
  ```

## Network/API Issues

### Issue: "API calls fail in app but work in browser"

**Solution:**
1. Check CORS settings if using external API
2. Verify network permissions in AndroidManifest.xml
3. Check if using HTTPS (required for Capacitor)

Add to `android/app/src/main/AndroidManifest.xml` if needed:
```xml
<uses-permission android:name="android.permission.INTERNET" />
```

## Getting More Help

### Check Logs

**Android Studio Logcat:**
1. **View → Tool Windows → Logcat**
2. Select your device
3. Filter by package name
4. Look for red error messages

**Chrome DevTools (for web content):**
1. Connect device via USB
2. Chrome → chrome://inspect
3. Click "inspect" on your app
4. Check Console tab

### Clean Everything and Start Fresh

```bash
cd frontend/web

# Clean node modules
rm -rf node_modules package-lock.json
npm install

# Clean build
rm -rf dist
npm run build

# Clean Android
rm -rf android .capacitor

# Start over
npx cap add android
npx cap sync

# Open in Android Studio
npx cap open android

# Clean and rebuild in Android Studio
# Build → Clean Project
# Build → Rebuild Project
```

## Still Having Issues?

1. Check the version compatibility:
   - Node.js 18+
   - Android SDK 33+
   - JDK 17+
   - Capacitor 6.x

2. Review the logs carefully
3. Search for error message in Capacitor docs
4. Check GitHub issues for similar problems

## Useful Commands for Debugging

```bash
# Check versions
node --version
npm --version
java -version
echo $ANDROID_HOME

# Check Android devices
adb devices

# View logs
adb logcat | grep -i "saloonbook"

# Clean everything
cd frontend/web
rm -rf node_modules dist android .capacitor
npm install
npm run build
npx cap add android
npx cap sync
```

---

Remember: Most issues are resolved by ensuring you run `npm run build` before `npx cap sync`!

