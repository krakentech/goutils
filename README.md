<picture>
    <source media="(prefers-color-scheme: dark)" srcset="./res/github-topper-dark.png" />
    <source media="(prefers-color-scheme: light)" srcset="./res/github-topper-light.png" />
    <img src="./res/github-topper-light.png" />
</picture>

![GitHub Release](https://img.shields.io/github/v/release/krakentech/goutils)
![Coverage](https://img.shields.io/badge/Coverage-0%25-red)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/krakentech/goutils)
![GitHub License](https://img.shields.io/github/license/krakentech/goutils)

## Installation

```bash
go get github.com/krakentech/goutils@latest
```

## Functions

### Files

* **FileExists:** checks if a file/directory exists at the path supplied may return error if fails to stat path
* **FileExistStrict:** checks if the file exists if an error occurs it returns false

### Slice

* **SliceIndexOf:** returns the index that holds the matching element otherwise a -1
* **SliceRemoveIdx:** returns a copy of the slice with the removed element at passed index
* **SliceRemoveItem:** is a shortcut for calling SliceIndexOf and them calling SliceRemoveIdx with the result

### String

* **StringToInt:** will convert the string to int or return 0
* **StringToInt64:** will convert the string to int64 or return 0
