set GO111MODULE=on
set CGO_ENABLED=1
set CGO_CFLAGS=-IC:/msys64/mingw64/include
set CGO_LDFLAGS=-LC:/msys64/mingw64/lib
$env:CGO_ENABLED=1
$env:CGO_CFLAGS="-IC:\msys64\mingw64\include"
$env:CGO_LDFLAGS="-LC:\msys64\mingw64\lib"
$env:PATH += ";C:\msys64\mingw64\bin"