package gosubst

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createValuesFile(t *testing.T) {
	f, err := os.Create("/tmp/values.toml")
	assert.Nil(t, err)

	_, err = f.WriteString("foo=\"bar\"\n")
	assert.Nil(t, err)

}

func TestSubstRender(t *testing.T) {
	createValuesFile(t)

	template := bytes.NewBufferString(`Testing {{.foo}}`)
	output := bytes.NewBuffer(nil)

	subst, err := NewSubst("/tmp/values.toml", "toml", template, output)
	assert.Nil(t, err)

	subst.Render()

	assert.Equal(t, "Testing bar", output.String())
}

func TestSubstRender_NoValues(t *testing.T) {
	template := bytes.NewBufferString(`Testing {{.foo}}`)
	output := bytes.NewBuffer(nil)

	subst, err := NewSubst("", "toml", template, output)
	assert.Nil(t, err)

	subst.Render()

	assert.Equal(t, "Testing ", output.String())
}

func TestGetenv(t *testing.T) {
	values := &Values{}
	assert.Empty(t, values.Env()["BLA"])
	assert.Equal(t, os.Getenv("USER"), values.Env()["USER"])
}
