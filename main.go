// alfred tool kit
/*
	自己用的工具
	author： Alfred Wu

*/
package main

import (
	"github.com/alfred/alfredtoolkit/icon"
	"github.com/alfred/alfredtoolkit/screens"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("Alfred Toolkit")

	mainLayout := screens.WidgetScreen(w)
	w.SetContent(mainLayout)
	w.SetIcon(icon.WindowIcon)

	// w.Resize(fyne.NewSize(1024, 768))
	w.CenterOnScreen()
	w.ShowAndRun()
}
