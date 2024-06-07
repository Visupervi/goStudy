package commands

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	flags "github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
)

type executable interface {
	Execute([]string) error
}

func testValidRefs(t *testing.T, v executable) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	specDoc := filepath.Join(fixtureBase, "expansion", "invalid-refs.json")
	result := v.Execute([]string{specDoc})
	assert.Error(t, result)
}

func testRequireParam(t *testing.T, v executable) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	result := v.Execute([]string{})
	assert.Error(t, result)

	result = v.Execute([]string{"nowhere.json"})
	assert.Error(t, result)
}

func getOutput(t *testing.T, specDoc, prefix, filename string) (string, string) {
	outDir, err := os.MkdirTemp(filepath.Dir(specDoc), "flatten")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return outDir, filepath.Join(outDir, filename)
}

func testProduceOutput(t *testing.T, v executable, specDoc, output string) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stdout)

	result := v.Execute([]string{specDoc})
	assert.NoError(t, result)
	_, exists := os.Stat(output)
	assert.True(t, !os.IsNotExist(exists))
}

// Commands requires at least one arg
func TestCmd_Flatten(t *testing.T) {
	v := &FlattenSpec{}
	testRequireParam(t, v)
}

func TestCmd_Flatten_Default(t *testing.T) {
	specDoc := filepath.Join(fixtureBase, "bugs", "1536", "fixture-1536.yaml")
	outDir, output := getOutput(t, specDoc, "flatten", "fixture-1536-flat-minimal.json")
	defer os.RemoveAll(outDir)
	v := &FlattenSpec{
		Format:  "json",
		Compact: true,
		Output:  flags.Filename(output),
	}
	testProduceOutput(t, v, specDoc, output)
}

func TestCmd_Flatten_Error(t *testing.T) {
	v := &FlattenSpec{}
	testValidRefs(t, v)
}

func TestCmd_Flatten_Issue2919(t *testing.T) {
	specDoc := filepath.Join(fixtureBase, "bugs", "2919", "edge-api", "client.yml")
	outDir, output := getOutput(t, specDoc, "flatten", "fixture-2919-flat-minimal.yml")
	defer os.RemoveAll(outDir)

	v := &FlattenSpec{
		Format:  "yaml",
		Compact: true,
		Output:  flags.Filename(output),
	}
	testProduceOutput(t, v, specDoc, output)
}

func TestCmd_FlattenKeepNames_Issue2334(t *testing.T) {
	specDoc := filepath.Join(fixtureBase, "bugs", "2334", "swagger.yaml")
	outDir, output := getOutput(t, specDoc, "flatten", "fixture-2334-flat-keep-names.yaml")
	defer os.RemoveAll(outDir)

	v := &FlattenSpec{
		Format:  "yaml",
		Compact: true,
		Output:  flags.Filename(output),
		FlattenCmdOptions: generate.FlattenCmdOptions{
			WithFlatten: []string{"keep-names"},
		},
	}
	testProduceOutput(t, v, specDoc, output)
	buf, err := os.ReadFile(output)
	require.NoError(t, err)
	spec := string(buf)

	require.Contains(t, spec, "$ref: '#/definitions/Bar'")
	require.Contains(t, spec, "Bar:")
	require.Contains(t, spec, "Baz:")
}
