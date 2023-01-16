package utils

import (
	"embed"
	"github.com/gookit/color"
	"log"
	"os"
	"regexp"
)

// Assets represents the embedded files.
//
//go:embed templates/dev.sh templates/dev_setup.sh
var Assets embed.FS

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FileContains(filePath string, needle string) bool {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	markerIsPresent, err := regexp.Match(needle, b)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	return markerIsPresent
}

func CopyAssetToPath(embedPath string, targetPath string) {
	data, err := Assets.ReadFile(embedPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	// Shell scripts need to be executable -> 0755
	err = os.WriteFile(targetPath, []byte(data), 0755)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	color.Magenta.Println(targetPath + " was created.")
}
