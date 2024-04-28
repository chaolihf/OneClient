set GOOS=windows
set GOARCH=amd64
go build -ldflags -H=windowsgui -i -o Release/windowshelper.exe .

