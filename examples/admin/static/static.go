package static

import "embed"

//go:embed js/* css/* img/* fonts/* favicon.ico index.html
var Static embed.FS
