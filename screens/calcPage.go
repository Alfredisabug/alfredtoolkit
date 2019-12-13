package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	log "github.com/sirupsen/logrus"
)

/*
	TODO List:
		UI Layout
		位元顯示轉換
		計算功能
		pec計算功能
*/

type calc struct {
	functions map[string]func()
	output    *widget.Label
	bitOutput *widget.Label
	errOutput *widget.Label
	buttons   map[string]*widget.Button
	radios    map[string]*widget.Radio
}

func (c *calc) addButtom(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button
	return button
}

func (c *calc) addRadio(named string, text []string, isHorizontal bool) *widget.Radio {
	radio := widget.NewRadio(
		text,
		func(s string) {
			log.Printf(s)
		})
	radio.Horizontal = isHorizontal
	c.radios[named] = radio
	return radio
}

func newCalc() *calc {
	c := &calc{}
	c.output = widget.NewLabel("")
	c.output.TextStyle.Monospace = true
	c.output.Alignment = fyne.TextAlignTrailing

	c.bitOutput = widget.NewLabel("0000 0000")
	c.bitOutput.TextStyle.Monospace = true
	c.bitOutput.Alignment = fyne.TextAlignTrailing

	c.buttons = make(map[string]*widget.Button)
	c.radios = make(map[string]*widget.Radio)
	return c
}

func makeCalcPageTab(win fyne.Window) fyne.CanvasObject {
	c := newCalc()
	win.Canvas().SetOnTypedRune(func(r rune) {
		c.output.SetText(c.output.Text + string(r))
	})

	return widget.NewVBox(
		widget.NewHBox(
			c.addRadio("digitShow", []string{"BIN", "DEC", "HEX"}, false),
			layout.NewSpacer(),
			c.output,
		),
		widget.NewVBox(
			c.bitOutput,
			c.addRadio("bitsShow", []string{"Byte", "short", "Int32", "Int 64"}, true),
		),
		widget.NewHBox(),
		widget.NewHBox(
			widget.NewGroup("PEC Function",
				c.addRadio("calcType", []string{"2's complement(2 byte)", "2's complement(1 byte)"}, false),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				fyne.NewContainerWithLayout(
					layout.NewGridLayout(4),
					c.addButtom("(", func() {}),
					c.addButtom(")", func() {}),
					c.addButtom("%", func() {}),
					c.addButtom("AC", func() {
						c.output.SetText("")
					}),
				),
				fyne.NewContainerWithLayout(
					layout.NewGridLayout(4),
					c.addButtom("7", func() {
						c.output.SetText(c.output.Text + "7")
					}),
					c.addButtom("8", func() {
						c.output.SetText(c.output.Text + "8")
					}),
					c.addButtom("9", func() {
						c.output.SetText(c.output.Text + "9")
					}),
					c.addButtom("*", func() {
						c.output.SetText(c.output.Text + "*")
					}),
				),
				fyne.NewContainerWithLayout(
					layout.NewGridLayout(4),
					c.addButtom("4", func() {
						c.output.SetText(c.output.Text + "4")
					}),
					c.addButtom("5", func() {
						c.output.SetText(c.output.Text + "5")
					}),
					c.addButtom("6", func() {
						c.output.SetText(c.output.Text + "6")
					}),
					c.addButtom("-", func() {
						c.output.SetText(c.output.Text + "-")
					}),
				),
				fyne.NewContainerWithLayout(
					layout.NewGridLayout(4),
					c.addButtom("1", func() {
						c.output.SetText(c.output.Text + "1")
					}),
					c.addButtom("2", func() {
						c.output.SetText(c.output.Text + "2")
					}),
					c.addButtom("3", func() {
						c.output.SetText(c.output.Text + "3")
					}),
					c.addButtom("+", func() {
						c.output.SetText(c.output.Text + "+")
					}),
				),
				fyne.NewContainerWithLayout(
					layout.NewGridLayout(4),
					c.addButtom(".", func() {
						c.output.SetText(c.output.Text + ".")
					}),
					c.addButtom("0", func() {
						c.output.SetText(c.output.Text + "0")
					}),
					c.addButtom("=", func() {}),
					c.addButtom("/", func() {}),
				),
			),
		),
	)
}
