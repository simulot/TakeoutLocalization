package localizations

import (
	"embed"
)

//go:embed *.json
var Default embed.FS
