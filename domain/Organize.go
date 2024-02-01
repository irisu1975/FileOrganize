package domain

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// OrganizeFile organizes file in target directory.
func OrganizeFile(target string, output string, config string) error {

	files, err := os.ReadDir(target)
	if err != nil {
		return fmt.Errorf(target + ": " + err.Error())
	}

	configs, err := ReadConfig(config)
	if err != nil {
		return fmt.Errorf(target + ": " + err.Error())
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != "" &&  file.IsDir() == false{
			var	matchFlg = false

			for _, cfg := range configs.Config {
				for _, cfgExt := range cfg.Extension {
					if ext[1:] == cfgExt {
						err := copyFile(filepath.Join(target, file.Name()), filepath.Join(output, cfg.DirName), file.Name())
						if err != nil {
							fmt.Println("copyFile:" + err.Error())
						}
						matchFlg = true
					}
				}
			}
			if matchFlg == false {
				copyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
			}
		}else{
			copyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
		}
	}

	return nil
}

// makeDirectory make directory.
func copyFile(file string, target string, fileName string) error {
	err := os.MkdirAll(target, 0777)
	if err != nil {
		return fmt.Errorf("MkdirAll: " + target + ":" + err.Error())
	}

	newF, err := os.Create(filepath.Join(target, fileName))
	if err != nil {
		return fmt.Errorf("OpenFile: " + target + ":" + err.Error())
	}

	srcF, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Open: " + file + ":" + err.Error())
	}

	_, err = io.Copy(newF, srcF)
	if err != nil {
		return fmt.Errorf("Copy: " + file + ":" + err.Error())
	}

	return nil
}