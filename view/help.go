package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Help struct {
	content *fyne.Container
}

func NewHelp() *Help {
	view := &Help{}

	titleText := widget.NewLabel("Goitopdf-gui")
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	titleText.Wrapping = fyne.TextWrapBreak

	mainText := widget.NewLabel("Made by br3w0r.\n App icon by Pixel Buddha.")
	mainText.Alignment = fyne.TextAlignCenter
	versionText := canvas.NewText("v0.7.0", color.White)
	versionText.Alignment = fyne.TextAlignCenter

	view.content = container.NewVBox(
		titleText,
		mainText,
		versionText,
	)

	return view
}

func (v *Help) Content() fyne.CanvasObject {
	return v.content
}
