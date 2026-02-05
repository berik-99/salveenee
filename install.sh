#!/bin/bash

# Repository Configuration
REPO="berik-99/salveenee"
BINARY_NAME="salveenee"
INSTALL_PATH="/usr/local/bin"

# Colors
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}Starting $BINARY_NAME installation...${NC}"

# 1. Architecture Detection
OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH_TYPE=$(uname -m)

case $ARCH_TYPE in
    x86_64)
        ARCH_SUFFIX="amd64"
        ;;
    aarch64|arm64)
        ARCH_SUFFIX="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture: $ARCH_TYPE${NC}"
        exit 1
        ;;
esac

# OS Check (Currently supporting Linux/WSL)
if [ "$OS_TYPE" != "linux" ]; then
    echo -e "${RED}This installer currently supports Linux only. For other OS, please build from source.${NC}"
    exit 1
fi

# 2. Fetch Latest Release URL via GitHub API
echo -e "Checking for the latest $ARCH_SUFFIX release..."
RELEASE_JSON=$(curl -s https://api.github.com/repos/$REPO/releases/latest)

# Extracting the download URL for the specific architecture
DOWNLOAD_URL=$(echo "$RELEASE_JSON" | grep "browser_download_url" | grep "_$ARCH_SUFFIX" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo -e "${RED}Error: Could not find binary 'salveenee_$ARCH_SUFFIX' in the latest GitHub release.${NC}"
    echo "Please ensure you have uploaded assets named exactly 'salveenee_amd64' and 'salveenee_arm64'."
    exit 1
fi

# 3. Downloading Binary
echo -e "Downloading binary..."
curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"

# 4. Installation
chmod +x "$BINARY_NAME"
echo -e "Moving binary to $INSTALL_PATH (sudo password may be required)..."

if [ -w "$INSTALL_PATH" ]; then
    mv "$BINARY_NAME" "$INSTALL_PATH/$BINARY_NAME"
else
    sudo mv "$BINARY_NAME" "$INSTALL_PATH/$BINARY_NAME"
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}------------------------------------------${NC}"
    echo -e "${GREEN}Installation successful!${NC}"
    echo -e "You can now run it by typing: ${BLUE}$BINARY_NAME${NC}"
    echo -e "${GREEN}------------------------------------------${NC}"
else
    echo -e "${RED}Error: Failed to move the binary to $INSTALL_PATH.${NC}"
    exit 1
fi