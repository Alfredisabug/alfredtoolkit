package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/alfred/alfredtoolkit/icon"
	"net/url"
)

func makeMainPageTab() fyne.Widget {
	logo := canvas.NewImageFromResource(icon.WindowIcon)
	logo.SetMinSize(fyne.NewSize(200, 200))

	link, err := url.Parse("https://jicommand.com")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	name := widget.NewLabelWithStyle("Author: Alfred Wu", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	return widget.NewVBox(
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		layout.NewSpacer(),
		name,
		widget.NewHyperlinkWithStyle("https://jicommand.com", link, fyne.TextAlignCenter, fyne.TextStyle{}),
	)
}

// WidgetScreen 最高層父視窗
func WidgetScreen(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
		widget.NewTabContainer(
			widget.NewTabItem("Home", makeMainPageTab()),
			widget.NewTabItem("Calc", makeCalcPageTab(win)),
		),
	)
}
