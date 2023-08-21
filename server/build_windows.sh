
rm -rf build
mkdir build
# go-winres simply --icon assets/server.png --manifest gui
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w -H=windowsgui" -o "build/LanCommanderClient_x64.exe"
GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -ldflags="-s -w -H=windowsgui" -o "build/LanCommanderClient_x86.exe"