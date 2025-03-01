package providers

import (
	"fmt"
	"strings"
)

const (
	maxDisplayNameLen int = 16
	maxDisplayURLLen  int = 19
)

func (app *App) fixURL() {
	if !(strings.HasPrefix(app.URL, "https://") || strings.HasPrefix(app.URL, "http://")) {
		app.URL = fmt.Sprintf("//%s", app.URL)
	}
}

func (app *App) formatDisplayURL(maxLength int) {
	app.fixURL()
	app.DisplayURL = app.URL
	if strings.HasPrefix(app.URL, "https://") || strings.HasPrefix(app.URL, "http://") {
		app.DisplayURL = strings.Split(app.URL, "://")[1]
	}
	if len(app.DisplayURL) > maxLength {
		app.DisplayURL = fmt.Sprintf("%s...", app.DisplayURL[:(maxLength-3)])
	}
}

func (app *App) formatDisplayName(name string, maxLength int) {
	if len(name) > maxLength {
		app.DisplayName = fmt.Sprintf("%s...", name[:(maxLength-3)])
	} else {
		app.DisplayName = name
	}
}
