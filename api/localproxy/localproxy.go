package main

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
)

func main() {
	module := app.DefaultModule(testfixtures.NewMockPersistence())
	errors.LogFatal(app.WithHandlers(module).Run("localhost:8080"))
}
