rsrc -manifest main.manifest -ico client.ico -o main.syso
set GOOS=windows
set GOARCH=amd64
go build -ldflags -H=windowsgui -o Release/windowshelper.exe .

