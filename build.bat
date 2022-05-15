go build -ldflags="-s -w" -o ./build/fs.exe main.go
upx ./build/fs.exe