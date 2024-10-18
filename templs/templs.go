package templs

import "embed"

//go:embed parts/*.tmpl *.tmpl
var TemplFiles embed.FS
