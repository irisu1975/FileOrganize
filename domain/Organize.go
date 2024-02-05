package domain

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FileOrganize/utils"
)

// OrganizeFile organizes file in target directory.
func OrganizeFile(target string, output string, config string) error {

	files, err := os.ReadDir(target)
	if err != nil {
		return fmt.Errorf("OrganizeFile():" + err.Error())
	}

	configs, err := ReadConfig(config)
	if err != nil {
		return fmt.Errorf("OrganizeFile():" + err.Error())
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext != "" &&  file.IsDir() == false{
			var	matchFlg = false

			for _, cfg := range configs.Config {
				for _, cfgExt := range cfg.Extension {
					if ext[1:] == cfgExt {
						err := utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, cfg.DirName), file.Name())
						if err != nil {
							fmt.Println("OrganizeFile():" + err.Error())
						}
						matchFlg = true
					}
				}
			}
			if matchFlg == false {
				err := utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
				if err != nil {
					fmt.Println("OrganizeFile():" + err.Error())
				}

			}
		}else{
			err := utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
			if err != nil {
				fmt.Println("OrganizeFile():" + err.Error())
			}
		}
	}

	return nil
}