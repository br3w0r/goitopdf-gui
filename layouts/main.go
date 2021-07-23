package layouts

import (
	"fmt"

	"fyne.io/fyne/v2"
)

type FileChoose struct{}

func (l *FileChoose) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width

		if childSize.Height > h {
			h = childSize.Height
		}
	}
	return fyne.NewSize(w, h)
}

func (l *FileChoose) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	nObjects := len(objects)
	if nObjects != 2 {
		panic(fmt.Errorf("FileChoose layout requires exactly 2 CanvasObjects, got %d", nObjects))
	}

	btnSize := objects[1].MinSize()
	objects[1].Resize(btnSize)
	objects[1].Move(fyne.NewPos(containerSize.Width-btnSize.Width, (containerSize.Height-btnSize.Height)/2))

	textWidth := containerSize.Width - btnSize.Width - 5
	objects[0].Resize(fyne.NewSize(textWidth, containerSize.Height))
	objects[0].Move(fyne.NewPos(0, 0))
}
