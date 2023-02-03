package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDevScriptTasks_TasksSyntax(t *testing.T) {
	expected := []DevScriptTask{
		{Name: "task-syntax-1a"},
		{Name: "task-syntax-1b"},
		{Name: "task-syntax-1c"},
		{Name: "task-syntax-2a"},
		{Name: "task-syntax-2b"},
		{Name: "task-syntax-2c"},
		{Name: "task-syntax-3a"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_syntax.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_IgnoreTasksStartingWithUnderscore(t *testing.T) {
	expected := []DevScriptTask{
		{Name: "task-not-starting-with-underscore"},
		{Name: "task-not-starting-with-underscore_"},
		{Name: "task-not-starting-with_underscore"},
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
		{Name: "task"},
		// should not be parsed
		// {Usage: "task-commented-out"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_commented_out.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_TaskCommentsShouldBeParsed(t *testing.T) {
	expected := []DevScriptTask{
		{
			Name:        "one-line-comment-block",
			Title:       "one line comment block - 1",
			Description: "",
		},
		{
			Name:        "multiline-comment-block",
			Title:       "multiline comment block - 1",
			Description: "multiline comment block - 2\nmultiline comment block - 3",
		},
		{
			Name:  "empty-lines-in-comment-block",
			Title: "multiline comment block with empty lines - 1",
			// TODO: remove trailing \n
			Description: "\nmultiline comment block with empty lines - 2\n\nmultiline comment block with empty lines - 3",
		},
		{
			Name:  "leading-spaces-or-none-in-comment-block",
			Title: "leading space missing",
			// remove # and one space always if present
			// keep \n but remove spaces if they are the only chars after #
			// also keep empty lines
			Description: "  keep leading-spaces used for indents\n\n\n\ncomment block end",
		},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_with_comments.sh")
	assert.Equal(t, expected, actual)
}

func TestParseDevScriptTasks_TasksFromImports(t *testing.T) {
	expected := []DevScriptTask{
		{Name: "task"},
		{Name: "import1-task1"},
		{Name: "import1-task2"},
		{Name: "import2-task1"},
		{Name: "import2-task2"},
		{Name: "nested-import-task1"},
		{Name: "nested-import-task2"},
	}
	actual := ParseDevScriptTasks("./fixtures/tasks_from_imports.sh")
	assert.Equal(t, expected, actual)
}
