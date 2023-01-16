package utils

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

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

func processComments(comments string) string {
	reqex := regexp.MustCompile(`(?m)(^#(\s)?)`)
	return reqex.ReplaceAllString(comments, "")
}

func extractShortFromComments(comments string) string {
	reqex := regexp.MustCompile(`^.*`)
	return reqex.FindString(comments)
}
