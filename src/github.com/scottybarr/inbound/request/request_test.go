package request

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	assert.Panics(t, func() { check(errors.New("Test")) })
	assert.NotPanics(t, func() { check(nil) })
}

func TestSetAuth(t *testing.T) {
	fakeReq, _ := http.NewRequest("GET", "http://localhost", nil)

	setAuth(fakeReq)

	assert.NotEmpty(t, fakeReq)
}
