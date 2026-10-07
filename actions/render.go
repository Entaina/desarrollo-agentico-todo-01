package actions

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"strings"
	"sync"

	"todo/public"
	"todo/templates"

	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine

var (
	assetVersions   = map[string]string{}
	assetVersionsMu sync.Mutex
)

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// fs.FS containing templates
		TemplatesFS: templates.FS(),

		// fs.FS containing assets
		AssetsFS: public.FS(),

		// Add template helpers here:
		Helpers: render.Helpers{
			"assetURL": assetURL,
		},
	})
}

// assetURL devuelve la ruta de un fichero de public/ con un parámetro de
// versión calculado a partir de su contenido, para que el navegador no
// use una copia antigua cuando el fichero cambia.
func assetURL(path string) string {
	path = "/" + strings.TrimPrefix(path, "/")

	assetVersionsMu.Lock()
	defer assetVersionsMu.Unlock()

	version, ok := assetVersions[path]
	if !ok {
		content, err := fs.ReadFile(public.FS(), strings.TrimPrefix(path, "/"))
		if err != nil {
			return path
		}
		sum := sha256.Sum256(content)
		version = hex.EncodeToString(sum[:])[:12]
		assetVersions[path] = version
	}

	return path + "?v=" + version
}
