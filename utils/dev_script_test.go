package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDevScriptTasks_TasksSyntax(t *testing.T) {
	expected := []DevScriptTask{
		{Usage: "task-syntax-1a"},
		{Usage: "task-syntax-1b"},
		{Usage: "task-syntax-1c"},
		{Usage: "task-syntax-2a"},
		{Usage: "task-syntax-2b"},
		{Usage: "task-syntax-2c"},
		{Usage: "task-syntax-3a"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_syntax.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_IgnoreTasksStartingWithUnderscore(t *testing.T) {
	expected := []DevScriptTask{
		{Usage: "task-not-starting-with-underscore"},
		{Usage: "task-not-starting-with-underscore_"},
		{Usage: "task-not-starting-with_underscore"},
		// should not be parsed
		// {Usage: "_task-without-underscore"},
		// {Usage: "__task-without-underscore"},
		// {Usage: "___task-without-underscore"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_with_underscores.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_IgnoreCommentedOutTasks(t *testing.T) {
	expected := []DevScriptTask{
		{Usage: "task"},
		// should not be parsed
		// {Usage: "task-commented-out"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_commented_out.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_TaskCommentsShouldBeParsed(t *testing.T) {
	expected := []DevScriptTask{
		{
			Usage: "one-line-comment-block",
			Short: "one line comment block - 1",
			// TODO: fix trailing \n
			Long: "one line comment block - 1",
		},
		{
			Usage: "multiline-comment-block",
			Short: "multiline comment block - 1",
			// TODO: fix trailing \n
			Long: "multiline comment block - 1\nmultiline comment block - 2\nmultiline comment block - 3",
		},
		{
			Usage: "empty-lines-in-comment-block",
			Short: "multiline comment block with empty lines - 1",
			// TODO: fix trailing \n
			Long: "multiline comment block with empty lines - 1\n\nmultiline comment block with empty lines - 2\n\nmultiline comment block with empty lines - 3",
		},
		{
			Usage: "leading-spaces-or-none-in-comment-block",
			Short: "leading space missing",
			// TODO: fix trailing \n
			// remove # and one space always if present
			// keep \n but remove spaces if they are the only chars after #
			// also keep empty lines
			Long: "leading space missing\n  keep leading-spaces used for indents\n\n\n\ncomment block end",
		},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_with_comments.sh")
	assert.Equal(t, expected, actual)
}
