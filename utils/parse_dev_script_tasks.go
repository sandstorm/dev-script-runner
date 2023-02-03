package utils

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func getSourcedScriptContent(sourcePath string, devScriptPath string) string {
	absolutePath := ""
	if filepath.IsAbs(sourcePath) {
		// sourcePath is absolute, we can ignore the devScriptPath
		absolutePath = sourcePath
	} else {
		// sourcePath is relative
		absolutePath = path.Join(path.Dir(devScriptPath), sourcePath)
	}
	sourceBytes, err := os.ReadFile(absolutePath)
	if err != nil {
		return ""
	}
	return string(sourceBytes)
}

func collectSourcedScripts(devScriptPath string) string {
	devScriptBytes, err := os.ReadFile(devScriptPath)
	if err != nil {
		log.Fatalf("Failed to execute: '%s'", err.Error())
	}
	result := string(devScriptBytes)
	// https://regex101.com/r/POMvPQ/1
	compiledRegex := regexp.MustCompile(`source (.*\.sh)`)
	matches := compiledRegex.FindAllStringSubmatch(result, -1)

	for _, match := range matches {
		if len(match) > 1 {
			// capture group `(.*\.sh)`
			result += "\n"
			result += getSourcedScriptContent(match[1], devScriptPath)
		}
	}

	return result
}

func ParseDevScriptTasks(devScriptPath string) []DevScriptTask {
	// https://regex101.com/r/ZFB8ml/1
	compiledRegex := regexp.MustCompile(`(?m)(?P<comments>(?:^#.*(?:\n|\r\n|\r))*)^(?:function )?(?P<name>[a-zA-Z0-9_-]+)\s?(?:\(\))?\s?{`)
	captureGroupNames := compiledRegex.SubexpNames()
	commentsIndex := sort.StringSlice(captureGroupNames).Search("comments")
	taskIndex := sort.StringSlice(captureGroupNames).Search("name")
	matches := compiledRegex.FindAllStringSubmatch(collectSourcedScripts(devScriptPath), -1)

	var results = []DevScriptTask{}
	for _, match := range matches {
		task := match[taskIndex]
		comments := prepareComments(match[commentsIndex])

		if !strings.HasPrefix(task, "_") {
			results = append(results, DevScriptTask{
				Name:        task,
				Title:       taskTitleFromComments(comments),
				Description: taskDescriptionFromComments(comments),
			})
		}
	}
	return results
}

// https://regex101.com/r/Jo4uSX/1
// `|\n$` matches new line and end of string
func prepareComments(comments string) string {
	reqex := regexp.MustCompile(`(?m)^# ?|\n$`)
	return reqex.ReplaceAllString(comments, "")
}
