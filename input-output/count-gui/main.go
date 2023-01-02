package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Count chars")

	i := binding.NewInt()
	str := binding.IntToString(i)
	entry := widget.NewEntry()
	entry.OnChanged = func(s string) { i.Set(len(s)) }

	w.SetContent(container.NewVBox(
		entry,
		widget.NewLabelWithData(str),
	))

	w.ShowAndRun()
}
