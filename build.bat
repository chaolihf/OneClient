set GOOS=windows
set GOARCH=amd64
go build -ldflags -H=windowsgui -o Release/windowshelper.exe .

