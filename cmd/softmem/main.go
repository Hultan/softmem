package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/hultan/softmem/internal/mainWindow"
	"os"
)

const ApplicationId string = "se.softteam.memo"
const ApplicationFlags glib.ApplicationFlags = glib.APPLICATION_FLAGS_NONE

func main() {
	// Create a new application
	app, err := gtk.ApplicationNew(ApplicationId, ApplicationFlags)
	errorCheck(err)

	mainWindow := mainWindow.NewMainWindow()

	// Hook up the activate event handler
	app.Connect("startup", mainWindow.OnStartup)
	app.Connect("activate", mainWindow.OnActivate)
	app.Connect("shutdown", mainWindow.OnShutdown)

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
