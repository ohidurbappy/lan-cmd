
rm -rf build
mkdir build

# cd ..
# ./go-png2ico server/assets/server.png server/assets/server.ico

# cd server

go-winres simply --icon assets/server.ico --manifest gui
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-s -w -H=windowsgui" -o "build/LanCommanderClient_x64.exe"
# GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -ldflags="-s -w -H=windowsgui" -o "build/LanCommanderClient_x86.exe"