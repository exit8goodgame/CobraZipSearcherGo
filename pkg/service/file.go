package service

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/yeka/zip"
)

func CrateDemoFile(fileName string) error {
	contents := []byte("Hello World")
	fZip, err := os.Create(fileName)
	if err != nil {
		return err
	}
	wZip := zip.NewWriter(fZip)
	defer wZip.Close()
	w, err := wZip.Encrypt(`sample.txt`, `123`, zip.AES256Encryption)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, bytes.NewReader(contents))
	if err != nil {
		return err
	}
	wZip.Flush()
	return nil
}

func CheckFileOpen(fullPath string) error {
	_, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	return nil
}

func GetFullPath(path string) (string, error) {
	basePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(basePath[:(len(basePath)-len(filepath.Base(basePath)))], path), nil
}
