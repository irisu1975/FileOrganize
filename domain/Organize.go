package domain

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// OrganizeFile organizes file in target directory.
func OrganizeFile(target string, output string) error {

	files, err := os.ReadDir(target)
	if err != nil {
		return fmt.Errorf(target + ": " + err.Error())
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != ""{
			fmt.Println(file.Name() + " : " + ext[1:])
			copyFile(file.Name(), filepath.Join(output, ext[1:]))
		}else {
			fmt.Println(file.Name())
		}
	}

	return nil
}

// makeDirectory make directory.
func copyFile(file string, target string) error {
	err := os.MkdirAll(target, 0777)
	if err != nil {
		return err
	}

	newF, err := os.Create(target)
	if err != nil {
		return err
	}

	srcF, err := os.Open(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(newF, srcF)
	if err != nil {
		return err
	}

	return nil
}