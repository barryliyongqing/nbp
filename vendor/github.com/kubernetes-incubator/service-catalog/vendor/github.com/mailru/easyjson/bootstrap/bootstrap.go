// Package bootstrap implements the bootstrapping logic: generation of a .go file to
// launch the actual generator and launching the generator itself.
//
// The package may be preferred to a command-line utility if generating the serializers
// from golang code is required.
package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const genPackage = "github.com/mailru/easyjson/gen"
const pkgWriter = "github.com/mailru/easyjson/jwriter"
const pkgLexer = "github.com/mailru/easyjson/jlexer"

type Generator struct {
	PkgPath, PkgName string
	Types            []string

	NoStdMarshalers bool
	SnakeCase       bool
	OmitEmpty       bool

	OutName   string
	BuildTags string

	StubsOnly  bool
	LeaveTemps bool
	NoFormat   bool
}

// writeStub outputs an initial stubs for marshalers/unmarshalers so that the package
// using marshalers/unmarshales compiles correctly for boostrapping code.
func (g *Generator) writeStub() error {
	f, err := os.Create(g.OutName)
	if err != nil {
		return err
	}
	defer f.Close()

	if g.BuildTags != "" {
		fmt.Fprintln(f, "// +build ", g.BuildTags)
		fmt.Fprintln(f)
	}
	fmt.Fprintln(f, "// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package")
	fmt.Fprintln(f, "// compilable during generation.")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "package ", g.PkgName)

	if len(g.Types) > 0 {
		fmt.Fprintln(f)
		fmt.Fprintln(f, "import (")
		fmt.Fprintln(f, `  "`+pkgWriter+`"`)
		fmt.Fprintln(f, `  "`+pkgLexer+`"`)
		fmt.Fprintln(f, ")")
	}

	for _, t := range g.Types {
		fmt.Fprintln(f)
		if !g.NoStdMarshalers {
			fmt.Fprintln(f, "func (", t, ") MarshalJSON() ([]byte, error) { return nil, nil }")
			fmt.Fprintln(f, "func (*", t, ") UnmarshalJSON([]byte) error { return nil }")
		}

		fmt.Fprintln(f, "func (", t, ") MarshalEasyJSON(w *jwriter.Writer) {}")
		fmt.Fprintln(f, "func (*", t, ") UnmarshalEasyJSON(l *jlexer.Lexer) {}")
		fmt.Fprintln(f)
		fmt.Fprintln(f, "type EasyJSON_exporter_"+t+" *"+t)
	}
	return nil
}

// writeMain creates a .go file that launches the generator if 'go run'.
func (g *Generator) writeMain() (path string, err error) {
	f, err := ioutil.TempFile(filepath.Dir(g.OutName), "easyjson-bootstrap")
	if err != nil {
		return "", err
	}

	fmt.Fprintln(f, "// +build ignore")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// TEMPORARY AUTOGENERATED FILE: easyjson bootstapping code to launch")
	fmt.Fprintln(f, "// the actual generator.")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import (")
	fmt.Fprintln(f, `  "fmt"`)
	fmt.Fprintln(f, `  "os"`)
	fmt.Fprintln(f)
	fmt.Fprintf(f, "  %q\n", genPackage)
	if len(g.Types) > 0 {
		fmt.Fprintln(f)
		fmt.Fprintf(f, "  pkg %q\n", g.PkgPath)
	}
	fmt.Fprintln(f, ")")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "func main() {")
	fmt.Fprintf(f, "  g := gen.NewGenerator(%q)\n", filepath.Base(g.OutName))
	fmt.Fprintf(f, "  g.SetPkg(%q, %q)\n", g.PkgName, g.PkgPath)
	if g.BuildTags != "" {
		fmt.Fprintf(f, "  g.SetBuildTags(%q)\n", g.BuildTags)
	}
	if g.SnakeCase {
		fmt.Fprintln(f, "  g.UseSnakeCase()")
	}
	if g.OmitEmpty {
		fmt.Fprintln(f, "  g.OmitEmpty()")
	}
	if g.NoStdMarshalers {
		fmt.Fprintln(f, "  g.NoStdMarshalers()")
	}
	for _, v := range g.Types {
		fmt.Fprintln(f, "  g.Add(pkg.EasyJSON_exporter_"+v+"(nil))")
	}

	fmt.Fprintln(f, "  if err := g.Run(os.Stdout); err != nil {")
	fmt.Fprintln(f, "    fmt.Fprintln(os.Stderr, err)")
	fmt.Fprintln(f, "    os.Exit(1)")
	fmt.Fprintln(f, "  }")
	fmt.Fprintln(f, "}")

	src := f.Name()
	if err := f.Close(); err != nil {
		return src, err
	}

	dest := src + ".go"
	return dest, os.Rename(src, dest)
}

func (g *Generator) Run() error {
	if err := g.writeStub(); err != nil {
		return err
	}
	if g.StubsOnly {
		return nil
	}

	path, err := g.writeMain()
	if err != nil {
		return err
	}
	if !g.LeaveTemps {
		defer os.Remove(path)
	}

	f, err := os.Create(g.OutName + ".tmp")
	if err != nil {
		return err
	}
	if !g.LeaveTemps {
		defer os.Remove(f.Name()) // will not remove after rename
	}

	cmd := exec.Command("go", "run", "-tags", g.BuildTags, path)
	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		return err
	}

	f.Close()

	if !g.NoFormat {
		cmd = exec.Command("gofmt", "-w", f.Name())
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			return err
		}
	}

	return os.Rename(f.Name(), g.OutName)
}
