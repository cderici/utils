// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package configurer

const (
	// AptSourcesFile is the default file which list all core
	// sources for apt packages on an apt-based system.
	AptSourcesFile = "/etc/apt/sources.list"

	// AptListsDirectory is the location of the APT sources list.
	AptListsDirectory = "/var/lib/apt/lists"

	// AptConfigDirectory is the default directory in which
	// apt configuration files are stored.
	UbuntuAptConfigFile = "/etc/apt/apt.conf.d"

	// ExtractAptSource is a shell command that will extract the
	// currently configured APT source location. We assume that
	// the first source for "main" in the file is the one that
	// should be replaced throughout the file.
	ExtractAptSource = `awk "/^deb .* $(lsb_release -sc) .*main.*\$/{print \$2;exit}" ` + AptSourcesFile

	// AptSourceListPrefix is a shell program that translates an
	// APT source (piped from stdin) to a file prefix. The algorithm
	// involves stripping up to one trailing slash, stripping the
	// URL scheme prefix, and finally translating slashes to
	// underscores.
	AptSourceListPrefix = `sed 's,.*://,,' | sed 's,/$,,' | tr / _`

	// AptSourceTemplate is the template specific to an apt source file.
	AptSourceTemplate = `
# {{.Name}} (added by Juju)
deb {{.Url}} %s main
# deb-src {{.Url}} %s main
`
)
