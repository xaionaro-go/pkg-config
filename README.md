# About

This is a minimalistic wrapper for `pkg-config` that allows to control which libraries should be linked dynamically or statically using two additional environment variables:
* `PKG_CONFIG_LIBS_FORCE_STATIC`
* `PKG_CONFIG_LIBS_FORCE_DYNAMIC`

Initially I had to implement this tool to make sure the CGo libraries I have as dependencies for my Android application are linked statically, because otherwise [it did not work](https://github.com/fyne-io/fyne/issues/5189).

# How to use

```sh
go install github.com/xaionaro-go/pkg-config@latest
PKG_CONFIG_LIBS_FORCE_STATIC="libav*,libvlc" PKG_CONFIG="$(go env GOPATH | awk -F ':' '{print $1}')/bin/pkg-config" go build PATH/TO/MY/PROJECT
```

It works not only with Go, but also with anything that understands variable `PKG_CONFIG`.

# Example of the output
```
$ go run ./ --libs-only-l libavcodec
-lavcodec
```
vs:
```
$ PKG_CONFIG_LIBS_FORCE_STATIC=libav* go run ./ --libs-only-l libavcodec
-Wl,-Bstatic -lavcodec -lvpx -lm -lvpx -lm -lvpx -lm -lvpx -lm -lwebpmux -lm -latomic -llzma -laribb24 -ldav1d -ldavs2 -lopencore-amrwb -lrsvg-2 -lm -lgio-2.0 -lgdk_pixbuf-2.0 -lgobject-2.0 -lglib-2.0 -lcairo -lzvbi -lpthread -lm -lpng -lz -lsnappy -lstdc++ -laom -lcodec2 -lfdk-aac -lgsm -lilbc -ljxl -ljxl_threads -lmp3lame -lm -lopencore-amrnb -lopenjp2 -lopus -lrav1e -lm -lshine -lspeex -lSvtAv1Enc -ltheoraenc -ltheoradec -logg -ltwolame -lvo-amrwbenc -lvorbis -lvorbisenc -lwebp -lx264 -lx265 -lxavs2 -lxvidcore -lopenh264 -lkvazaar -lz -lva -lvpl -ldl -lstdc++ -lswresample -lm -lsoxr -latomic -lva-drm -lva -lva-x11 -lva -lvdpau -lX11 -lgcrypt -lm -ldrm -lvpl -ldl -lstdc++ -lOpenCL -lssl -lcrypto -lva -latomic -lX11 -lavutil -lva-drm -lva -lva-x11 -lva -lvdpau -lX11 -lgcrypt -lm -ldrm -lvpl -ldl -lstdc++ -lOpenCL -lssl -lcrypto -lva -latomic -lX11
```
