@if exist "%~dp0Release\cefcapi.exe" (
    @del "%~dp0Release\cefcapi.exe"
)

@call gcc --version
@call gcc -c -Wall -Werror -I. -I.. -L../Release main_win.c -o number.o -lcef
@call ar rcs libnumber.a number.o
cd /d/tools/golang/workspace/WindowsHelperClient
g++ -c -g cefapi/main_win.c -o cefapi/cefapi.o -LRelease -lcef -I. -I/c/msys64/mingw64/include/c++/13.2.0 -I/c/msys64/mingw64/include/c++/13.2.0/x86_64-w64-mingw32/bits