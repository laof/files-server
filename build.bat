go build -ldflags="-s -w" -o filesserver.exe main.go 
upx filesserver.exe