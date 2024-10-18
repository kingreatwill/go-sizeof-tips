package pub

import "embed"

//go:embed images styles js
var StaticFiles embed.FS
