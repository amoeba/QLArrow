#!/bin/sh

go build -buildmode=c-archive -o internal.a ./internal
clang -Wno-deprecated-declarations -framework Cocoa -framework WebKit -framework ImageIO -framework Foundation -framework QuickLook -framework CoreServices -bundle -o QLArrow.qlgenerator/Contents/MacOS/QLArrow main.c GeneratePreviewForURL.m GenerateThumbnailForURL.m internal.a
