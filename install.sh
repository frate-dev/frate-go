OS=$(uname |tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
 
if [[ $ARCH == "x86_64" ]]; then
    ARCH="amd64"
elif [[ $ARCH == "aarch64" ]]; then
    ARCH="arm64"
fi

URL="https://github.com/frate-dev/frate-go/releases/download/main/frate-$OS-$ARCH"
wget -O frate "$URL" > /dev/null 2>&1 
chmod +x frate 
sudo mv frate /usr/local/bin
