@if exist "%~dp0Release\cefcapi.exe" (
    @del "%~dp0Release\cefcapi.exe"
)

@call gcc --version
@call gcc -c -Wall -Werror -I. -I.. -L../Release main_win.c -o number.o -lcef
@call ar rcs libnumber.a number.o