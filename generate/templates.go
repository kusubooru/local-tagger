package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	inPath  = "web/"
	outFile = "templates.go"
)

var t = `
// generated by go generate; DO NOT EDIT

package main

import "html/template"

var (
	{{$lname:=.Layout.Name}}
	{{$lname}}Tmpl = template.Must(template.New("{{$lname}}").Funcs(fns).Parse({{$lname}}Template))
	{{range .TFiles}}
	{{.Name}}Tmpl = template.Must(template.Must({{$lname}}Tmpl.Clone()).Parse({{.Name}}Template))
	{{end}}

	{{.Layout}}

	{{range .TFiles}}{{.}}{{end}}
)
`

var templatesTmpl = template.Must(template.New("t").Parse(t))

func main() {
	if err := generateTemplates(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func generateTemplates() error {
	tfiles, err := loadFiles(inPath)
	if err != nil {
		return err
	}

	out, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); err == nil {
			err = cerr
		}
	}()

	err = templatesTmpl.Execute(out, tfiles)
	if err != nil {
		return err
	}

	cmd := exec.Command("gofmt", "-w", outFile)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return err
}

type templateFile struct {
	Name    string
	Content []byte
}

func (tf templateFile) String() string {
	return fmt.Sprintf("%vTemplate = `\n%v`\n", tf.Name, string(tf.Content))
}

type templateFiles struct {
	Layout templateFile
	TFiles []templateFile
}

func loadFiles(path string) (*templateFiles, error) {

	var tfiles []templateFile
	var layout *templateFile

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			if matches, _ := filepath.Match("*.tmpl", f.Name()); matches {
				if matches, _ := filepath.Match("layout.tmpl", f.Name()); matches {
					layout, err = readTemplateFile(filepath.Join(path, f.Name()))
					if err != nil {
						return nil, nil
					}
				} else {
					tf, err := readTemplateFile(filepath.Join(path, f.Name()))
					if err != nil {
						return nil, nil
					}
					tfiles = append(tfiles, *tf)
				}
			}
		}
	}
	return &templateFiles{Layout: *layout, TFiles: tfiles}, nil
}

func readTemplateFile(fpath string) (*templateFile, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return &templateFile{
		Name:    strings.TrimSuffix(filepath.Base(f.Name()), filepath.Ext(f.Name())),
		Content: c,
	}, nil
}
