package main

import (
	"github.com/fergusstrange/roundofbeer/api/app"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
)

func main() {
	module := app.NewApplicationModule(testfixtures.NewMockPersistence())
	errors.LogFatal(app.Handlers(module).Run("localhost:8080"))
}
