// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package configurer

import (
	"bytes"
	"text/template"
)

// Source contains all the data required for a package source.
type PackageSource struct {
	Name string `yaml:"-"`
	Url  string `yaml:"source"`
	Key  string `yaml:"key, omitempty"`
}

// renderSourceFile renders the current source based on a template it recieves.
func (s *PackageSource) renderSourceFile(fileTemplate string) string {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(fileTemplate))

	if err := t.Execute(&buf, s); err != nil {
		panic(err)
	}

	return buf.String()
}
