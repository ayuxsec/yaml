package yaml_test

import (
	"testing"

	"github.com/ayuxsec/yaml"
	"github.com/stretchr/testify/require"
)

type exampleStruct struct {
	Name  string `yaml:"name" yaml_line_comment:"your name"`
	Age   int    `yaml:"age" yaml_line_comment:"your age"`
	Place string `yaml:"place" yaml_line_comment:"your location"`
}

func TestMarshalStruct(t *testing.T) {
	var exampleStruct = exampleStruct{
		Name:  "spike spiegel",
		Age:   27,
		Place: "space",
	}
	data, err := yaml.Marshal(exampleStruct)
	require.NoError(t, err)
	t.Log(string(data))
}
