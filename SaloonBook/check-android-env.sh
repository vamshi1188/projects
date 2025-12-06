#!/bin/bash

# Post-Installation Verification Script for SaloonBook Android Build
# Run this script to verify your environment is ready to build Android APK

echo "=========================================="
echo "SaloonBook Android Build - Environment Check"
echo "=========================================="
echo ""

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

ISSUES=0

# Check Node.js
echo -n "Checking Node.js... "
if command -v node &> /dev/null; then
    NODE_VERSION=$(node --version)
    echo -e "${GREEN}✓${NC} Found: $NODE_VERSION"

    # Check if version is 18+
    MAJOR_VERSION=$(echo $NODE_VERSION | cut -d'v' -f2 | cut -d'.' -f1)
    if [ "$MAJOR_VERSION" -lt 18 ]; then
        echo -e "${YELLOW}  ⚠ Warning: Node.js 18+ recommended (you have v$MAJOR_VERSION)${NC}"
        ISSUES=$((ISSUES+1))
    fi
else
    echo -e "${RED}✗${NC} Not found"
    echo "  Install from: https://nodejs.org/"
    ISSUES=$((ISSUES+1))
fi

# Check npm
echo -n "Checking npm... "
if command -v npm &> /dev/null; then
    NPM_VERSION=$(npm --version)
    echo -e "${GREEN}✓${NC} Found: v$NPM_VERSION"
else
    echo -e "${RED}✗${NC} Not found"
    ISSUES=$((ISSUES+1))
fi

# Check Java
echo -n "Checking Java/JDK... "
if command -v java &> /dev/null; then
    JAVA_VERSION=$(java -version 2>&1 | head -n 1)
    echo -e "${GREEN}✓${NC} Found: $JAVA_VERSION"

    # Check if version is 17+
    if ! echo "$JAVA_VERSION" | grep -q "17\|18\|19\|20\|21"; then
        echo -e "${YELLOW}  ⚠ Warning: JDK 17+ recommended${NC}"
        ISSUES=$((ISSUES+1))
    fi
else
    echo -e "${RED}✗${NC} Not found"
    echo "  Install: sudo apt install openjdk-17-jdk (Ubuntu/Debian)"
    ISSUES=$((ISSUES+1))
fi

# Check ANDROID_HOME
echo -n "Checking ANDROID_HOME... "
if [ -n "$ANDROID_HOME" ]; then
    echo -e "${GREEN}✓${NC} Set: $ANDROID_HOME"

    # Check if directory exists
    if [ ! -d "$ANDROID_HOME" ]; then
        echo -e "${RED}  ✗ Directory does not exist!${NC}"
        ISSUES=$((ISSUES+1))
    fi
else
    echo -e "${RED}✗${NC} Not set"
    echo "  Add to ~/.bashrc:"
    echo "  export ANDROID_HOME=\$HOME/Android/Sdk"
    echo "  export PATH=\$PATH:\$ANDROID_HOME/tools"
    echo "  export PATH=\$PATH:\$ANDROID_HOME/platform-tools"
    ISSUES=$((ISSUES+1))
fi

# Check Android Studio
echo -n "Checking Android Studio... "
if [ -d "$HOME/android-studio" ] || [ -d "/opt/android-studio" ] || [ -d "/usr/local/android-studio" ]; then
    echo -e "${GREEN}✓${NC} Found"
elif command -v studio.sh &> /dev/null; then
    echo -e "${GREEN}✓${NC} Found"
else
    echo -e "${YELLOW}⚠${NC} Not found in common locations"
    echo "  Install from: https://developer.android.com/studio"
    ISSUES=$((ISSUES+1))
fi

# Check adb
echo -n "Checking adb (Android Debug Bridge)... "
if command -v adb &> /dev/null; then
    ADB_VERSION=$(adb version | head -n 1)
    echo -e "${GREEN}✓${NC} Found: $ADB_VERSION"
else
    echo -e "${YELLOW}⚠${NC} Not found in PATH"
    echo "  Add to PATH or install Android SDK Platform-Tools"
fi

# Check project dependencies
echo ""
echo "Checking project setup..."

if [ -d "frontend/web/node_modules" ]; then
    echo -e "${GREEN}✓${NC} node_modules installed"
else
    echo -e "${YELLOW}⚠${NC} node_modules not found"
    echo "  Run: cd frontend/web && npm install"
    ISSUES=$((ISSUES+1))
fi

if [ -f "frontend/web/capacitor.config.ts" ]; then
    echo -e "${GREEN}✓${NC} Capacitor config exists"
else
    echo -e "${RED}✗${NC} Capacitor config missing"
    ISSUES=$((ISSUES+1))
fi

if [ -f "frontend/web/package.json" ]; then
    if grep -q "@capacitor/android" "frontend/web/package.json"; then
        echo -e "${GREEN}✓${NC} Capacitor dependencies in package.json"
    else
        echo -e "${YELLOW}⚠${NC} Capacitor dependencies not found"
        echo "  Run: cd frontend/web && npm install"
        ISSUES=$((ISSUES+1))
    fi
fi

if [ -d "frontend/web/android" ]; then
    echo -e "${GREEN}✓${NC} Android platform added"
else
    echo -e "${YELLOW}⚠${NC} Android platform not added yet"
    echo "  Run: cd frontend/web && npx cap add android"
fi

if [ -d "frontend/web/dist" ]; then
    echo -e "${GREEN}✓${NC} Build output exists"
else
    echo -e "${YELLOW}⚠${NC} Build not done yet"
    echo "  Run: cd frontend/web && npm run build"
fi

# Summary
echo ""
echo "=========================================="
if [ $ISSUES -eq 0 ]; then
    echo -e "${GREEN}✓ All checks passed! You're ready to build.${NC}"
    echo ""
    echo "Next steps:"
    echo "1. Run: ./setup-android.sh"
    echo "2. Or manually: cd frontend/web && npx cap open android"
else
    echo -e "${YELLOW}⚠ Found $ISSUES issue(s) that need attention.${NC}"
    echo ""
    echo "Please resolve the issues above before building."
    echo "See ANDROID_BUILD_GUIDE.md for detailed instructions."
fi
echo "=========================================="

