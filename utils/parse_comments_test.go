package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseComments_TitleAndDescription(t *testing.T) {
	// title only
	comment := "some title"
	titleExpected := "some title"
	titleActual := taskTitleFromComments(comment)
	assert.Equal(t, titleExpected, titleActual)
	descriptionExpected := ""
	descriptionActual := taskDescriptionFromComments(comment)
	assert.Equal(t, descriptionExpected, descriptionActual)

	// title only with linebreak
	comment = "some title\n"
	titleExpected = "some title"
	titleActual = taskTitleFromComments(comment)
	assert.Equal(t, titleExpected, titleActual)
	descriptionExpected = ""
	descriptionActual = taskDescriptionFromComments(comment)
	assert.Equal(t, descriptionExpected, descriptionActual)

	// title with comment
	comment = ""
	comment += "some title\n"
	comment += "Description Line 1\n"
	comment += "Description Line 2\n"
	titleExpected = "some title"
	titleActual = taskTitleFromComments(comment)
	assert.Equal(t, titleExpected, titleActual)
	descriptionExpected = "Description Line 1\n"
	descriptionExpected += "Description Line 2\n"
	descriptionActual = taskDescriptionFromComments(comment)
	assert.Equal(t, descriptionExpected, descriptionActual)

	// title with comment and Usage or Examples
	//
	/*
		comment = ""
		comment += "some title\n"
		comment += "Description Line 1\n"
		comment += "Description Line 2\n"
		comment += "Usage:\n"
		titleExpected = "some title"
		titleActual = taskTitleFromComments(comment)
		assert.Equal(t, titleExpected, titleActual)
		descriptionExpected = "Description Line 1\n"
		descriptionExpected += "Description Line 2"
		descriptionActual = taskDescriptionFromComments(comment)
		assert.Equal(t, descriptionExpected, descriptionActual)

		// title with comment and "Examples:"
		comment = ""
		comment += "some title\n"
		comment += "Description Line 1\n"
		comment += "Description Line 2\n"
		comment += "Examples:\n"
		titleExpected = "some title"
		titleActual = taskTitleFromComments(comment)
		assert.Equal(t, titleExpected, titleActual)
		descriptionExpected = "Description Line 1\n"
		// TODO: fix reqex to exclude \n of last line
		descriptionExpected += "Description Line 2\n"
		descriptionActual = taskDescriptionFromComments(comment)
		assert.Equal(t, descriptionExpected, descriptionActual)
	*/
}
