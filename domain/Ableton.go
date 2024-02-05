package domain

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FileOrganize/utils"
)

const EXT_MP3 = ".mp3"		// mp3 extension.
const EXT_WAV = ".wav"		// wav extension.
var outputDirectory = ""

// Ableton organizes mp3, wav file from target project files.
func Ableton(target string, output string) error {
	outputDirectory = output
	err := filepath.Walk(target, visit)
	if err != nil {
		return fmt.Errorf("Ableton():" + err.Error())
	}
	return nil
}

// visit is WalkFunction.
func visit(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == EXT_MP3 || filepath.Ext(path) == EXT_WAV {
		fmt.Println(path + " : " + info.Name())
		err := utils.CopyFile(path, outputDirectory, info.Name())
		if err != nil {
			return fmt.Errorf("visit():" + err.Error())
		}
	}else {
		return fmt.Errorf("visit():Music file(mp3, wav) does not exist")
	}
	return nil
}