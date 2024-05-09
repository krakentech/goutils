package goutils

// SliceIndexOf returns the index that holds the matching element otherwise a -1
func SliceIndexOf[T comparable](slice []T, element T) int {
    for i, x := range slice {
        if x == element {
            return i
        }
    }
    return -1
}

// SliceRemoveIdx returns a copy of the slice with the removed element at passed index
func SliceRemoveIdx[T comparable](slice []T, idx int) []T {
    out := make([]T, len(slice))
    copy(out, slice)
    if idx < 0 || idx > len(out)-1 {
        return out
    }
    return append(out[:idx], out[idx+1:]...)
}

// SliceRemoveItem is a shortcut for calling SliceIndexOf and them calling SliceRemoveIdx with the result
func SliceRemoveItem[T comparable](slice []T, element T) []T {
    idx := SliceIndexOf(slice, element)
    return SliceRemoveIdx(slice, idx)
}
