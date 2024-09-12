package internal

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSONBody(t *testing.T) {
	t.Run("parse success", func(t *testing.T) {
		buf := bytes.Buffer{}
		buf.WriteString(`{"mockKey":"mockValue"}`)
		req, _ := http.NewRequest("", "", &buf)
		var responseLoader interface{}
		err := ParseJSONBody(req, &responseLoader)
		assert.Nil(t, err)
	})
	t.Run("parse fail", func(t *testing.T) {
		ch := make(chan int)
		buf := bytes.Buffer{}
		buf.WriteString(fmt.Sprint(ch))
		req, _ := http.NewRequest("", "", &buf)
		var responseLoader interface{}
		err := ParseJSONBody(req, responseLoader)
		assert.NotNil(t, err)
	})
}
