#!/usr/bin/env bash

# download dependencies
go mod download

# clean up build directory
if [ -d "build" ]; then
    rm -rf build
fi

# create build directory
mkdir build

# introduce os and arch martix
os_arch_matrix=(
    "linux/amd64"
    "linux/arm64"
    "linux/386"
    "linux/arm"
    "darwin/amd64"
    "windows/amd64"
)

# generate code
go generate .

# build for each os and arch combination
for os_arch in "${os_arch_matrix[@]}"
do
    echo "Building for $os_arch"
    GOOS=${os_arch%/*}
    GOARCH=${os_arch#*/}
    OUTPUT_NAME="build/sfcrypt-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi
    env GOOS="$GOOS" GOARCH="$GOARCH" go build -v=0 -a -trimpath -ldflags "-s -w -extldflags '-static'" -o "$OUTPUT_NAME" github.com/kmou424/sfcrypt/cmd/v2/sfcrypt
done