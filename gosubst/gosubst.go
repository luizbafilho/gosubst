package gosubst

import (
	"encoding/json"
	"text/template"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"

	yaml "gopkg.in/yaml.v2"
)

// Values holds data present in the values file
type Values map[string]interface{}

// Env exports all enviroments variable
func (v *Values) Env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

// Subst receives a values file, values type, input and renders the template replacing
// all variables present in the values file sending it to the output
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

	tmpl, err := template.New("base").Option("missingkey=error").Funcs(sprig.TxtFuncMap()).Parse(string(text))
	if err != nil {
		return err
	}

	return tmpl.Execute(s.output, s.values)
}

func loadValues(path string, vType string) (Values, error) {
	var values Values

	if path == "" {
		return values, nil
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Reading values file failed")
	}

	switch strings.ToLower(vType) {
	case "json":
		if err := json.Unmarshal(file, &values); err != nil {
			return nil, errors.Wrapf(err, "Parsing %s as json failed", path)
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(file, &values); err != nil {
			return nil, errors.Wrapf(err, "Parsing %s as yaml failed", path)
		}
	case "toml":
		tree, err := toml.Load(string(file))
		if err != nil {
			return nil, errors.Wrapf(err, "Parsing %s as toml failed", path)
		}

		values = tree.ToMap()
	}

	return values, nil
}
