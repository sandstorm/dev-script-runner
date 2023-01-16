package utils

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func ParseDevScriptTasks(devScriptPath string) []DevScriptTask {
	// https://regex101.com/r/ZFB8ml/1
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
		comments := processComments(match[commentsIndex])

		if !strings.HasPrefix(task, "_") {
			results = append(results, DevScriptTask{
				Usage: task,
				Short: extractShortFromComments(comments),
				Long:  comments,
			})
		}
	}
	return results
}

// https://regex101.com/r/Jo4uSX/1
// `|\n$` matches new line and end of string
func processComments(comments string) string {
	reqex := regexp.MustCompile(`(?m)^# ?|\n$`)
	return reqex.ReplaceAllString(comments, "")
}

func extractShortFromComments(comments string) string {
	reqex := regexp.MustCompile(`^.*`)
	return reqex.FindString(comments)
}
