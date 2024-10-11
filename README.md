# About

This is a minimalistic wrapper that adds `--static` to `--libs` requests, and adds `-Wl,-Bstatic` and `-Wl,-Bdynamic` around the response. This forces to link statically whatever pkg-config returns.

Initially I had to implement this tool to make sure the CGo libraries I have as dependencies for my Android application are linked statically, because otherwise [it did not work](https://github.com/fyne-io/fyne/issues/5189).

# How to use

```sh
go install github.com/xaionaro-go/pkg-config-static@latest
PKG_CONFIG="$HOME/go/bin/pkg-config-static" go build PATH/TO/MY/PROJECT
```

It works not only with Go, but also with anything that understands variable `PKG_CONFIG`.
