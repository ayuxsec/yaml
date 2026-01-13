package yaml_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type ExampleStruct struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Place string `yaml:"place"`
}

func TestMarshalStruct(t *testing.T) {
	var exampleStruct = ExampleStruct{
		Name:  "spike spiegel",
		Age:   27,
		Place: "space",
	}
	data, err := yaml.Marshal(exampleStruct)
	require.NoError(t, err)
	t.Log(string(data))
}
