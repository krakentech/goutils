package goutils

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestSliceIndexOf(t *testing.T) {
    m := []string{"a", "b", "C"}

    t.Run("not found", func(t *testing.T) {
        assert.Equal(t, -1, SliceIndexOf(m, "z"))
        assert.Equal(t, []string{"a", "b", "C"}, m)
    })

    t.Run("found", func(t *testing.T) {
        assert.Equal(t, 1, SliceIndexOf(m, "b"))
        assert.Equal(t, []string{"a", "b", "C"}, m)
    })

    t.Run("case is important", func(t *testing.T) {
        assert.Equal(t, 2, SliceIndexOf(m, "C"))
        assert.Equal(t, []string{"a", "b", "C"}, m)
        assert.Equal(t, -1, SliceIndexOf(m, "c"))
        assert.Equal(t, []string{"a", "b", "C"}, m)
    })

}

func TestSliceRemoveIdx(t *testing.T) {

    t.Run("index less than 0", func(t *testing.T) {
        data := []string{"a", "b", "C"}
        SliceRemoveIdx(data, -1)
        assert.Equal(t, []string{"a", "b", "C"}, data)
    })

    t.Run("index greater than slice len", func(t *testing.T) {
        data := []string{"a", "b", "C"}
        SliceRemoveIdx(data, 10)
        assert.Equal(t, []string{"a", "b", "C"}, data)
    })

    t.Run("remove item", func(t *testing.T) {
        data := []string{"a", "b", "C"}
        SliceRemoveIdx(data, 1)
        assert.Equal(t, []string{"a", "b", "C"}, data)
    })

}

func TestSliceRemoveItem(t *testing.T) {
    t.Run("not found", func(t *testing.T) {
        assert.Equal(t, []string{"a", "b", "C"}, SliceRemoveItem([]string{"a", "b", "C"}, "z"))
        assert.Equal(t, []string{"a", "b", "C"}, []string{"a", "b", "C"})
    })

    t.Run("found", func(t *testing.T) {
        assert.Equal(t, []string{"b", "C"}, SliceRemoveItem([]string{"a", "b", "C"}, "a"))
        assert.Equal(t, []string{"a", "b", "C"}, []string{"a", "b", "C"})
    })

    t.Run("case is important", func(t *testing.T) {
        assert.Equal(t, []string{"a", "b"}, SliceRemoveItem([]string{"a", "b", "C"}, "C"))
        assert.Equal(t, []string{"a", "b", "C"}, []string{"a", "b", "C"})
        assert.Equal(t, []string{"a", "b", "C"}, SliceRemoveItem([]string{"a", "b", "C"}, "c"))
        assert.Equal(t, []string{"a", "b", "C"}, []string{"a", "b", "C"})
    })

}
