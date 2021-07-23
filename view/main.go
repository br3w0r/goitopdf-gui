package view

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/br3w0r/goitopdf-gui/layouts"
	"github.com/br3w0r/goitopdf/itopdf"
	"github.com/skratchdot/open-golang/open"
)

type data struct {
	saveFile       string
	saveFileDirUri fyne.ListableURI
}

type MainView struct {
	data          *data
	window        *fyne.Window
	inDirLabel    *canvas.Text
	inDirText     *widget.Entry
	inDirBtn      *widget.Button
	saveDirLabel  *canvas.Text
	saveDirText   *widget.Entry
	saveDirBtn    *widget.Button
	outNameEntry  *widget.Entry
	saveBtn       *widget.Button
	saveLabel     *widget.Label
	saveFile      *widget.Label
	openFolderBtn *widget.Button
	openFileBtn   *widget.Button
	content       *fyne.Container
}

func NewMainView(window *fyne.Window) *MainView {
	view := &MainView{window: window, data: &data{}}

	view.inDirLabel = canvas.NewText("Input folder:", color.White)
	view.inDirText = widget.NewEntry()
	view.inDirText.SetPlaceHolder("(not chosen)")
	view.inDirText.OnChanged = view.enableSaveBtn
	view.inDirBtn = widget.NewButton("Choose", view.chooseFolder(0))

	view.saveDirLabel = canvas.NewText("Output folder:", color.White)
	view.saveDirText = widget.NewEntry()
	view.saveDirText.SetPlaceHolder("(same as input folder)")
	view.saveDirBtn = widget.NewButton("Choose", view.chooseFolder(1))

	view.outNameEntry = widget.NewEntry()
	view.outNameEntry.SetPlaceHolder("Output file name")
	view.outNameEntry.OnChanged = view.enableSaveBtn

	view.saveBtn = widget.NewButton("Save", view.save)
	view.saveBtn.Disable()

	view.saveLabel = widget.NewLabel("")
	view.saveFile = widget.NewLabel("")
	view.saveFile.Wrapping = fyne.TextWrapBreak

	view.openFileBtn = widget.NewButton("Open File", func() {
		open.Start(view.data.saveFile)
	})
	view.openFileBtn.Hide()

	view.openFolderBtn = widget.NewButton("Open Folder", view.openFolder)
	view.openFolderBtn.Hide()

	view.content = container.NewVBox(
		view.inDirLabel,
		container.New(
			&layouts.FileChoose{},
			view.inDirText,
			view.inDirBtn,
		),
		view.saveDirLabel,
		container.New(
			&layouts.FileChoose{},
			view.saveDirText,
			view.saveDirBtn,
		),
		view.outNameEntry,
		view.saveBtn,
		view.saveLabel,
		view.saveFile,
		container.New(
			layout.NewGridLayoutWithColumns(2),
			view.openFileBtn,
			view.openFolderBtn,
		),
	)

	return view
}

func (v *MainView) Content() fyne.CanvasObject {
	return v.content
}

func (v *MainView) enableSaveBtn(s string) {
	if v.outNameEntry.Text != "" && v.inDirText.Text != "" {
		v.saveBtn.Enable()
	} else {
		v.saveBtn.Disable()
	}
}

func (v *MainView) openFolder() {
	if v.saveDirText.Text == "" {
		open.Start(v.inDirText.Text)
	} else {
		open.Start(v.saveDirText.Text)
	}
}

func (v *MainView) chooseFolder(t int) func() {
	return func() {
		d := dialog.NewFolderOpen(v.chooseDir(t), *v.window)
		d.Resize(v.content.Size())
		if v.data.saveFileDirUri != nil {
			d.SetLocation(v.data.saveFileDirUri)
		}
		d.Show()
	}
}

func (v *MainView) chooseDir(t int) func(uri fyne.ListableURI, err error) {
	return func(uri fyne.ListableURI, err error) {
		if err != nil {
			panic(err)
		}

		if uri == nil {
			return
		}

		switch t {
		case 0:
			v.inDirText.SetText(uri.Path())
			v.data.saveFileDirUri = uri
		case 1:
			v.saveDirText.SetText(uri.Path())
		default:
			panic(fmt.Errorf("Unsupported dir choose type: %d", t))
		}
	}
}

func (v *MainView) save() {
	pdf := itopdf.NewInstance()

	v.saveLabel.SetText("Saving...")
	err := pdf.WalkDir(v.inDirText.Text, v.logSave)
	if err != nil {
		err = pdf.Save(v.data.saveFile)
		v.errorMessage("Save error:", err.Error())
		return
	}

	if v.saveDirText.Text == "" {
		v.data.saveFile = v.inDirText.Text
	} else {
		v.data.saveFile = v.saveDirText.Text
	}
	v.data.saveFile += "/" + v.outNameEntry.Text

	err = pdf.Save(v.data.saveFile)
	if err != nil {
		v.errorMessage("Save error:", err.Error())
		return
	}

	v.saveLabel.SetText("Successfully saved to:")
	v.saveFile.SetText(v.data.saveFile)

	v.openFileBtn.Show()
	v.openFolderBtn.Show()
}

func (v *MainView) errorMessage(label string, message string) {
	v.saveLabel.SetText(label)
	v.saveFile.SetText(message)
}

func (v *MainView) logSave(s string) {
	v.saveFile.SetText(s)
}
