package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFile copy directory.
func CopyFile(file string, target string, fileName string) error {
	err := os.MkdirAll(target, 0777)
	if err != nil {
		return fmt.Errorf("CopyFile():MkdirAll():" + target + ":" + err.Error())
	}

	newF, err := os.Create(filepath.Join(target, fileName))
	if err != nil {
		return fmt.Errorf("CopyFile():OpenFile():" + target + ":" + err.Error())
	}

	srcF, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("CopyFile():Open():" + file + ":" + err.Error())
	}

	_, err = io.Copy(newF, srcF)
	if err != nil {
		return fmt.Errorf("CopyFile():Copy():" + file + ":" + err.Error())
	}

	return nil
}