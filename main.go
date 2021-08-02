package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/br3w0r/goitopdf-gui/view"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("goitopdf-gui")

	v := view.NewMainView(myApp, myWindow)

	myWindow.SetContent(v.Content())

	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}
