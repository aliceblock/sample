package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsd413dffbac67c655ac111807d5f7cc8b9d4c69c9 = "<body>\r\n  <h1>API</h1>\r\n  <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.</p>\r\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"web"}, "/web": []string{"templates"}, "/web/templates": []string{"api.tmpl"}}, map[string]*assets.File{
	"/web/templates": &assets.File{
		Path:     "/web/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1528227690, 1528227690293079800),
		Data:     nil,
	}, "/web/templates/api.tmpl": &assets.File{
		Path:     "/web/templates/api.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1528227690, 1528227690294000000),
		Data:     []byte(_Assetsd413dffbac67c655ac111807d5f7cc8b9d4c69c9),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1528227973, 1528227973613674900),
		Data:     nil,
	}, "/web": &assets.File{
		Path:     "/web",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1528226330, 1528226330191866100),
		Data:     nil,
	}}, "")
