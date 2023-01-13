package utils

import (
	"embed"
	"github.com/gookit/color"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

const DEV_SCRIPT_MARKER = "DEV_SCRIPT_MARKER"
const DEV_SCRIPT_NAME = "dev.sh"
const MAX_DEPTH = 100
const SECTION_SEPARATOR = "--------------------------------------------------------------------------"

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

func ParseDevScriptTasks(devScriptPath string) []DevScriptTask {
	// https://regex101.com/r/5LVRcP/1 -> Iteration 1 without comments before
	// https://regex101.com/r/XyB410/1 -> Final Iteration with comments before ;)
	devScriptBytes, err := os.ReadFile(devScriptPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}

	devScriptString := string(devScriptBytes)
	compiledRegex := regexp.MustCompile(`(?m)(?P<comments>(?:^#.*(?:\n|\r\n|\r))*)^(?:function )?(?P<name>[a-zA-Z0-9_-]+)\s?(?:\(\))?\s?{`)
	captureGroupNames := compiledRegex.SubexpNames()
	commentsIndex := sort.StringSlice(captureGroupNames).Search("comments")
	taskIndex := sort.StringSlice(captureGroupNames).Search("name")
	matches := compiledRegex.FindAllStringSubmatch(devScriptString, -1)

	var results = []DevScriptTask{}
	for _, match := range matches {
		task := match[taskIndex]
		comments := match[commentsIndex]
		if !strings.HasPrefix(task, "_") {
			results = append(results, DevScriptTask{Name: task, Comments: comments})
		}
	}
	return results
}
