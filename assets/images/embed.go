package images

import (
	_ "embed"
)

var (
	//go:embed gopher-l.png
	GopherL_png []byte
	//go:embed gopher-r.png
	GopherR_png []byte
)
