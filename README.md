# QLArrow

[QuickLook](https://support.apple.com/guide/mac-help/view-and-edit-files-with-quick-look-mh14119/mac) plugin for various [Arrow](https://arrow.apache.org/) file formats.
[Arrow](https://arrow.apache.org/) is a high-performance, language-agnostic columnar memory format.

## Roadmap

[Parquet](https://parquet.apache.org/) is the main use case here but [Arrow IPC](https://arrow.apache.org/docs/format/Columnar.html) may be useful for some.

Currently, [QLArrow](https://github.com/amoeba/QLArrow) provides [QuickLook](https://support.apple.com/guide/mac-help/view-and-edit-files-with-quick-look-mh14119/mac) previews for:

- [x] [Parquet](https://parquet.apache.org/)
- [ ] Feather
- [ ] Arrow IPC


# Installation

- Download the [latest release](https://github.com/amoeba/QLArrow/releases/) and extract the ZIP
- Run `xattr -d com.apple.quarantine QLArrow.qlgenerator` on `QLArrow.qlgenerator`
- Move to `~/Library/QuickLook`

# Building

0. Pre-requisites

- XCode
- Go

1. Build the Go subproject

    ```sh
    go build -buildmode=c-archive -o internal.a ./internal
    ```

2. Build the XCode subproject
  * Open `QLArrow.xcodeproj`
  * Build the `QLArrow` target


## Credit

- [https://github.com/toland/qlmarkdown](https://github.com/toland/qlmarkdown)
- [https://github.com/remko/qlmka](https://github.com/remko/qlmka)
