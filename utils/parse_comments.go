package utils

import (
	"regexp"
)

func taskTitleFromComments(comments string) string {
	reqex := regexp.MustCompile(`^.*`)
	return reqex.FindString(comments)
}

func taskDescriptionFromComments(comments string) string {
	//reqex := regexp.MustCompile(`(?ms)(?:\n|\r\n|\r)(.*?)((?:\n|\r\n|\r)^Usage:|Examples:|\n$)`)
	reqex := regexp.MustCompile(`(?ms)(?:\n|\r\n|\r)(.*)`)
	match := reqex.FindStringSubmatch(comments)
	if len(match) > 1 {
		// return capture group `(.*?)`
		return match[1]
	}
	return ""
}
