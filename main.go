package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	icon, err := getIcon()
	log.Printf("loaded icon %dx%d", icon.GetWidth(), icon.GetHeight())
	if err != nil {
		log.Fatal("Unable to load an icon:", err)
	}
	win.SetIcon(icon)
	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add the label to the window.
	win.Add(l)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}

func getIcon() (*gdk.Pixbuf, error) {
	data, err := Asset("data/cc-icon.png")
	if err != nil {
		return nil, err
	}
	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		return nil, err
	}
	icon, err := loader.WriteAndReturnPixbuf(data)
	if err != nil {
		return nil, err
	}
	return icon, nil
}
