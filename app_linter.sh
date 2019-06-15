#!/bin/sh

st=0
for pkg in $(go list ./... | grep -v /vendor/); do
    echo "==> $pkg"

    go vet "$pkg"
    [ $? -ne 0 ] && st=1

    golint "$pkg"
    [ $? -ne 0 ] && st=1

    # gofmt works on files, not packages
    gofmt -d "${f#arp242.net/trackwall}"*.go
    [ $? -ne 0 ] && st=1
done
exit $st