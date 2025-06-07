package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Brave File Manager")

	sidebar := widget.NewList(
		func() int { return 3 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i int, o fyne.CanvasObject) {
			labels := []string{"Local", "SFTP", "FTP"}
			o.(*widget.Label).SetText(labels[i])
		},
	)

	mainContent := widget.NewLabel("Select a location from the sidebar.")
	w.SetContent(container.NewHSplit(sidebar, mainContent))
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
