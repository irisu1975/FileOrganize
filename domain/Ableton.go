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
	filepath.Walk(target, visit)
}

// visit is WalkFunction.
func visit(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) == EXT_MP3 || filepath.Ext(path) == EXT_WAV {
		fmt.Println(path + " : " + info.Name())
		utils.CopyFile(path, outputDirectory, info.Name())
	}
	return nil
}