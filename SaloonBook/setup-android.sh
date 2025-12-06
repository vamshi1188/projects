#!/bin/bash

# SaloonBook Android APK Setup Script
# This script automates the initial setup for building an Android APK

set -e

echo "=========================================="
echo "SaloonBook Android APK Setup"
echo "=========================================="
echo ""

# Navigate to frontend/web directory
cd "$(dirname "$0")/frontend/web"

echo "Step 1: Installing npm dependencies..."
npm install

echo ""
echo "Step 2: Building the web application..."
npm run build

echo ""
echo "Step 3: Checking for Android platform..."
if [ -d "android" ]; then
    echo "Android platform already exists. Skipping initialization."
else
    echo "Adding Android platform..."
    npx cap add android
fi

echo ""
echo "Step 4: Syncing web app with Android..."
npx cap sync

echo ""
echo "=========================================="
echo "Setup Complete! ✅"
echo "=========================================="
echo ""
echo "Next steps:"
echo "1. Open Android Studio:"
echo "   npx cap open android"
echo ""
echo "2. In Android Studio, go to:"
echo "   Build → Build Bundle(s) / APK(s) → Build APK(s)"
echo ""
echo "3. Or build from command line:"
echo "   cd frontend/web"
echo "   npm run android:build"
echo ""
echo "Your APK will be at:"
echo "   frontend/web/android/app/build/outputs/apk/debug/app-debug.apk"
echo ""
echo "See ANDROID_BUILD_GUIDE.md for detailed instructions."
echo "=========================================="

