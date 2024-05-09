set GOOS=windows
set GOARCH=amd64
go clean --cache
go build -ldflags -H=windowsgui -o Release/windowshelper.exe .

