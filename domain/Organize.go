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
						err := utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, cfg.DirName), file.Name())
						if err != nil {
							fmt.Println("copyFile:" + err.Error())
						}
						matchFlg = true
					}
				}
			}
			if matchFlg == false {
				utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
			}
		}else{
			utils.CopyFile(filepath.Join(target, file.Name()), filepath.Join(output, "Other"), file.Name())
		}
	}

	return nil
}