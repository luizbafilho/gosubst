package gosubst

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/pelletier/go-toml"

	yaml "gopkg.in/yaml.v2"
)

type Values map[string]interface{}

func (v *Values) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

type Subst struct {
	values *Values
	input  io.Reader
	output io.Writer
}

// NewSubst substitues values
func NewSubst(valuesFile string, valuesType string, input io.Reader, output io.Writer) (*Subst, error) {
	values, err := loadValues(valuesFile, valuesType)
	if err != nil {
		return nil, err
	}

	return &Subst{
		values: &values,
		input:  input,
		output: output,
	}, nil

}

// Render render template with the given values
func (s Subst) Render() error {
	text, err := ioutil.ReadAll(s.input)
	if err != nil {
		return err
	}

	tmpl, err := template.New("base").Funcs(sprig.FuncMap()).Parse(string(text))
	if err != nil {
		return err
	}

	return tmpl.Execute(s.output, s.values)
}

func loadValues(path string, vType string) (Values, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var values Values

	switch strings.ToLower(vType) {
	case "json":
		if err := json.Unmarshal(file, &values); err != nil {
			return nil, err
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(file, &values); err != nil {
			return nil, err
		}
	case "toml":
		tree, err := toml.Load(string(file))
		if err != nil {
			return nil, err
		}

		values = tree.ToMap()
	}

	return values, nil
}
