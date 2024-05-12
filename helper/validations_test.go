package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"


)

func Test_CheckEmpty(t *testing.T) {
	errs := []error{}
	CheckEmpty("", &errs, "name")

	assert.NotNil(t, errs)
	assert.Equal(t, "name can not be empty", errs[0].Error())

	errs = []error{}
	CheckEmpty("value", &errs, "name")
	assert.Empty(t, errs)
}
