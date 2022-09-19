#!/bin/sh

go build -buildmode=c-archive -o internal.a ./internal
#TODO: Commented out for now but kept so I remember how
#clang -Wno-deprecated-declarations -framework Cocoa -framework WebKit -framework ImageIO -framework Foundation -framework QuickLook -framework CoreServices -framework Security -bundle -o QLArrow.qlgenerator/Contents/MacOS/QLArrow main.c GeneratePreviewForURL.m GenerateThumbnailForURL.m internal.a
