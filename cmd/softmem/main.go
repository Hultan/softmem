package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/hultan/softmem/internal/softmem"
	"os"
)

const ApplicationId string = "se.softteam.memo"
const ApplicationFlags = glib.APPLICATION_FLAGS_NONE

func main() {
	// Create a new application
	app, err := gtk.ApplicationNew(ApplicationId, ApplicationFlags)
	errorCheck(err)

	mainWindow := softmem.NewMainWindow()

	// Hook up the activate event handler
	_ = app.Connect("startup", mainWindow.OnStartup)
	_ = app.Connect("activate", mainWindow.OnActivate)
	_ = app.Connect("shutdown", mainWindow.OnShutdown)

	// Start the application
	status := app.Run(os.Args)

	// Exit
	os.Exit(status)
}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
