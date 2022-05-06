go build -ldflags="-s -w" -o ./build/filesserver.exe main.go
upx ./build/filesserver.exe