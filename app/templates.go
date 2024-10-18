package app

import (
	"html/template"

	"github.com/kingreawill/go-sizeof-tips/templs"
)

var templates map[string]*template.Template

func prepareTemplates() error {
	templates = make(map[string]*template.Template)
	baseData, err := templs.TemplFiles.ReadFile("parts/base.tmpl")
	if err != nil {
		return err
	}
	var fns = template.FuncMap{
		"unvischunk": func(x int, len int) bool {
			return x > 2 && x < (len-1)
		},
	}
	for _, name := range []string{
		"index", "404", "500",
	} {
		assetData, err := templs.TemplFiles.ReadFile(name + ".tmpl")
		if err != nil {
			return err
		}
		templates[name], err = template.New(name).Funcs(fns).Parse(
			string(baseData) + string(assetData),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
