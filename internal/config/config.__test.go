package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetVarSuccess(t *testing.T) {
	value, err := GetVar("API_URL")

	assert.NoError(t, err, "should not return an error for a valid value.")
	assert.Equal(t, "https://yahoo-finance15.p.rapidapi.com/api/", value, "value should be the expected value for the api url")
}

func Test_GetVarFail(t *testing.T) {
	value, err := GetVar("BAD_KEY")

	assert.Error(t, err, "should not have a value for the provided key")
	assert.Equal(t, "", value, "should be an empty string due to no actual key")
}
