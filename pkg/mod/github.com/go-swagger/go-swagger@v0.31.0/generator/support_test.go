package generator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type relativePathTest struct {
	childpath  string
	parentpath string
	ok         bool
	path       string
}

func prefixAndFetchRelativePathFixtures() []relativePathTest {
	return []relativePathTest{
		// Positive
		{"/", "/", true, "."},
		{"/User/Gopher", "/", true, "User/Gopher"},
		{"/User/Gopher/Go", "/User/Gopher/Go", true, "."},
		{"/User/../User/Gopher", "/", true, "User/Gopher"},
		// Negative cases
		{"/", "/var", false, ""},
		{"/User/Gopher", "/User/SomethingElse", false, ""},
		{"/var", "/etc", false, ""},
		{"/mnt/dev3", "/mnt/dev3/dir", false, ""},
	}
}

type baseImportTest struct {
	title        string
	path         []string
	gopath       string
	targetpath   string
	symlinksrc   string
	symlinkdest  string // symlink is the last dir in targetpath
	expectedpath string
}

func baseImportTestFixtures(tempdir string) []baseImportTest {
	return []baseImportTest{
		{
			title:        "No sym link. Positive Test Case",
			path:         []string{tempdir + "/root/go/src/github.com/go-swagger"},
			gopath:       tempdir + "/root/go/",
			targetpath:   tempdir + "/root/go/src/github.com/go-swagger",
			symlinksrc:   "",
			symlinkdest:  "",
			expectedpath: "github.com/go-swagger",
		},
		{
			title:        "Symlink points inside GOPATH",
			path:         []string{tempdir + "/root/go/src/github.com/go-swagger"},
			gopath:       tempdir + "/root/go/",
			targetpath:   tempdir + "/root/symlink",
			symlinksrc:   tempdir + "/root/symlink",
			symlinkdest:  tempdir + "/root/go/src/",
			expectedpath: ".",
		},
		{
			title:        "Symlink points inside GOPATH (2)",
			path:         []string{tempdir + "/root/go/src/github.com/go-swagger"},
			gopath:       tempdir + "/root/go/",
			targetpath:   tempdir + "/root/symlink",
			symlinksrc:   tempdir + "/root/symlink",
			symlinkdest:  tempdir + "/root/go/src/github.com",
			expectedpath: "github.com",
		},
		{
			title:        "Symlink point outside GOPATH : Targets Case 1: in baseImport implementation",
			path:         []string{tempdir + "/root/go/src/github.com/go-swagger", tempdir + "/root/gopher/go/"},
			gopath:       tempdir + "/root/go/",
			targetpath:   tempdir + "/root/go/src/github.com/gopher",
			symlinksrc:   tempdir + "/root/go/src/github.com/gopher",
			symlinkdest:  tempdir + "/root/gopher/go",
			expectedpath: "github.com/gopher",
		},
	}
}

func TestCheckPrefixFetchRelPath(t *testing.T) {
	for _, item := range prefixAndFetchRelativePathFixtures() {
		actualok, actualpath := checkPrefixAndFetchRelativePath(item.childpath, item.parentpath)

		item.path = filepath.FromSlash(item.path)

		assert.Equalf(t, item.ok, actualok, "checkPrefixAndFetchRelativePath(%s, %s): expected %v, actual %v", item.childpath, item.parentpath, item.ok, actualok)
		assert.Equal(t, item.path, actualpath, "checkPrefixAndFetchRelativePath(%s, %s): expected %s, actual %s", item.childpath, item.parentpath, item.path, actualpath)
	}
}

func TestBaseImport(t *testing.T) {
	// 1. Create a root folder /tmp/root
	// 2. Simulate scenario
	//	2.a No Symlink
	//	2.b Symlink from outside of GOPATH to inside
	//  2.c Symlink from inside of GOPATH to outside.
	// 3. Check results.

	tempdir := t.TempDir()

	golang := GoLangOpts()

	for _, item := range baseImportTestFixtures(tempdir) {
		t.Logf("TestBaseImport(%q)", item.title)

		// Create Paths
		for _, paths := range item.path {
			require.NoError(t, os.MkdirAll(paths, 0o700))
		}

		if item.symlinksrc == "" {
			continue
		}

		// Create Symlink
		require.NoErrorf(t, os.Symlink(item.symlinkdest, item.symlinksrc),
			"WARNING:TestBaseImport with symlink could not be carried on. Symlink creation failed for %s -> %s\n%s",
			item.symlinksrc, item.symlinkdest,
			"NOTE:TestBaseImport with symlink on Windows requires extended privileges (admin or a user with SeCreateSymbolicLinkPrivilege)",
		)

		// Change GOPATH
		t.Setenv("GOPATH", item.gopath)

		// Test (baseImport always with /)
		actualpath := golang.baseImport(item.targetpath)

		require.Equalf(t, item.expectedpath, actualpath, "baseImport(%s): expected %s, actual %s", item.targetpath, item.expectedpath, actualpath)

		_ = os.RemoveAll(filepath.Join(tempdir, "root"))
	}
}

func TestGenerateMarkdown(t *testing.T) {
	defer discardOutput()()

	t.Run("should generate doc for demo fixture", func(t *testing.T) {
		opts := testGenOpts()
		opts.Spec = "../fixtures/enhancements/184/fixture-184.yaml"
		output := filepath.Join(t.TempDir(), "markdown.md")

		require.NoError(t, GenerateMarkdown(output, nil, nil, opts))
		expectedCode := []string{
			"# Markdown generator demo",
		}

		code, err := os.ReadFile(output)
		require.NoError(t, err)

		for line, codeLine := range expectedCode {
			if !assertInCode(t, strings.TrimSpace(codeLine), string(code)) {
				t.Logf("Code expected did not match in codegenfile %s for expected line %d: %q", output, line, expectedCode[line])
			}
		}
	})

	t.Run("should handle new lines in descriptions", func(t *testing.T) {
		opts := testGenOpts()
		opts.Spec = "../fixtures/bugs/2700/2700.yaml"
		output := filepath.Join(t.TempDir(), "markdown.md")

		require.NoError(t, GenerateMarkdown(output, nil, nil, opts))
		expectedCode := []string{
			`| Filesystem type of the volume that you want to mount.</br>Tip: Ensure that the filesystem type is supported by the host operating system.</br>Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.</br>More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore</br></br>TODO: how do we prevent errors in the filesystem from compromising the machine |`,
		}

		code, err := os.ReadFile(output)
		require.NoError(t, err)

		for line, codeLine := range expectedCode {
			if !assertInCode(t, strings.TrimSpace(codeLine), string(code)) {
				t.Logf("Code expected did not match in codegenfile %s for expected line %d: %q", output, line, expectedCode[line])
			}
		}
	})
}
