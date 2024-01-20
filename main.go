package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func main() {

	App := app.New()
	Window := App.NewWindow("Lampanel (golang)")

	Run := widget.NewLabel("Display")
	
	Window.SetContent(container.NewVBox(
		Run,
		widget.NewButton("clickme", func() {Run.SetText("yalpsiD")},
		),

	))

	Window.ShowAndRun()

	return
}
