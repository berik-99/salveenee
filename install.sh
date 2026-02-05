#!/bin/bash

# Configurazione repository
REPO="berik-99/salveenee"
BINARY_NAME="salveenee"
INSTALL_PATH="/usr/local/bin"

# Colori
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${BLUE}Inizio installazione di $BINARY_NAME...${NC}"

# 1. Rilevamento Architettura
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
        echo -e "${RED}Architettura non supportata: $ARCH_TYPE${NC}"
        exit 1
        ;;
esac

# Al momento supportiamo solo Linux (WSL incluso)
if [ "$OS_TYPE" != "linux" ]; then
    echo -e "${RED}Questo installer supporta solo Linux. Per altri OS, compila dai sorgenti.${NC}"
    exit 1
fi

# 2. Recupero URL ultima release tramite API GitHub
echo -e "Ricerca dell'ultima versione per $ARCH_SUFFIX..."
RELEASE_JSON=$(curl -s https://api.github.com/repos/$REPO/releases/latest)
# Estraiamo l'URL che contiene il suffisso corretto (es: salveenee_amd64)
DOWNLOAD_URL=$(echo "$RELEASE_JSON" | grep "browser_download_url" | grep "_$ARCH_SUFFIX" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo -e "${RED}Errore: Impossibile trovare il binario 'salveenee_$ARCH_SUFFIX' nelle release di GitHub.${NC}"
    exit 1
fi

# 3. Download
echo -e "Download in corso..."
curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"

# 4. Installazione
chmod +x "$BINARY_NAME"
echo -e "Spostamento del binario in $INSTALL_PATH (potrebbe servire la password di sudo)..."

if [ -w "$INSTALL_PATH" ]; then
    mv "$BINARY_NAME" "$INSTALL_PATH/$BINARY_NAME"
else
    sudo mv "$BINARY_NAME" "$INSTALL_PATH/$BINARY_NAME"
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}------------------------------------------${NC}"
    echo -e "${GREEN}Installazione riuscita!${NC}"
    echo -e "Digita '${BLUE}$BINARY_NAME${NC}' per avviare lo script."
    echo -e "${GREEN}------------------------------------------${NC}"
else
    echo -e "${RED}Errore durante lo spostamento del file.${NC}"
    exit 1
fi