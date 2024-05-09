package goutils

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestFileExists(t *testing.T) {
    tmpDir := t.TempDir()
    // Create a valid test file
    t.Run("Check For File Get Error", func(t *testing.T) {
        exists, err := FileExists(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.False(t, exists, "the file should NOT exist for this test")
        assert.NoError(t, err, "we should have no issues reading this file")
    })

    // Create a valid test file
    t.Run("Check For File Get Error", func(t *testing.T) {
        exists, err := FileExists("\000x")
        assert.False(t, exists, "the file should NOT exist for this test")
        assert.Equal(t, "stat \x00x: invalid argument", err.Error())
    })

    t.Run("Check For Existing File (no errors)", func(t *testing.T) {
        _, err := os.Create(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.NoError(t, err, "failed to create test file")
        exists, err := FileExists(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.True(t, exists, "the file should exist for this test")
        assert.NoError(t, err, "we should have no issues reading this file")
    })
}

func TestFileExistStrict(t *testing.T) {
    tmpDir := t.TempDir()
    // Create a valid test file
    t.Run("Check For File Get Error", func(t *testing.T) {
        exists := FileExistStrict(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.False(t, exists, "the file should NOT exist for this test")
    })

    // Create a valid test file
    t.Run("Check For File Get Error", func(t *testing.T) {
        exists := FileExistStrict("\000x")
        assert.False(t, exists, "the file should NOT exist for this test")
    })

    t.Run("Check For Existing File (no errors)", func(t *testing.T) {
        _, err := os.Create(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.NoError(t, err, "failed to create test file")
        exists := FileExistStrict(fmt.Sprintf("%s/file.txt", tmpDir))
        assert.True(t, exists, "the file should exist for this test")
    })
}
