# Notepage-V2
notepage started as nodejs project years ago and was a texteditor which used html to edit text this is notepage V2  a customizable editor implementing markdown with proper linting and modding support

## setting the dev environment up
in order to beeing able to work with fyne the following is needed to be installed:
https://www.msys2.org/
inside the msys window run the following commands:
```sh
pacman -Syu
pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-glew mingw-w64-x86_64-glfw
```

set these environment variables:
```
set GO111MODULE=on
set CGO_ENABLED=1
set CGO_CFLAGS=-IC:/msys64/mingw64/include
set CGO_LDFLAGS=-LC:/msys64/mingw64/lib
```
now run `go mod tidy` and then `go run main.go`
you get the error:
```
imports github.com/go-gl/gl/v2.1/gl: build constraints exclude all Go files in C:\Users\...\go\pkg\mod\github.com\go-gl\gl@v0.0.0-20211210172815-726fda9656d6\v2.1\gl
```
then run:
```
$env:CGO_ENABLED=1
$env:CGO_CFLAGS="-IC:\msys64\mingw64\include"
$env:CGO_LDFLAGS="-LC:\msys64\mingw64\lib"
$env:PATH += ";C:\msys64\mingw64\bin"
```
now start the main.go again
note: the first time starting the terminal will run the app without responding that is normal and can take a few minutes fyne is building 