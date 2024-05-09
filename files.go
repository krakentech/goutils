package goutils

import (
    "errors"
    "os"
)

// FileExists checks if a file/directory exists at the path supplied may return error if fails to stat path
func FileExists(path string) (bool, error) {
    if _, err := os.Stat(path); err == nil {
        return true, nil
    } else if errors.Is(err, os.ErrNotExist) {
        return false, nil
    } else {
        return false, err
    }
}

// FileExistStrict checks if the file exists or if it has an error with the file.
func FileExistStrict(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true
    }
    return false
}
