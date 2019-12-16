package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/alfred/alfredtoolkit/features"
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
	output             *widget.Label
	bitOutput          *widget.Label
	errOutput          *widget.Label
	stringSrcIutput    *widget.Entry
	stringResultOutput *widget.Label
	buttons            map[string]*widget.Button
	radios             map[string]*widget.Radio
	functions          map[string]func()
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

	c.stringSrcIutput = widget.NewEntry()
	c.stringSrcIutput.SetPlaceHolder("AABBCCDD(HEX)")
	c.stringResultOutput = widget.NewLabel("")

	c.buttons = make(map[string]*widget.Button)
	c.radios = make(map[string]*widget.Radio)
	return c
}

func makeCalcPageTab(win fyne.Window) fyne.CanvasObject {
	c := newCalc()
	win.Canvas().SetOnTypedRune(func(r rune) {
		c.output.SetText(c.output.Text + string(r))
	})

	regionWidget1 := widget.NewHBox(
		c.addRadio("digitShow", []string{"BIN", "DEC", "HEX"}, false),
		layout.NewSpacer(),
		c.output,
	)

	regionWidget2 := widget.NewVBox(
		c.bitOutput,
		c.addRadio("bitsShow", []string{"Byte", "short", "Int32", "Int 64"}, true),
	)

	functionGroup := widget.NewGroup("PEC Function",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			c.stringSrcIutput,
			c.addButtom("2's complement(1 byte)", func() {
				if len(c.stringSrcIutput.Text)%2 != 0 || len(c.stringSrcIutput.Text) == 0 {
					c.stringResultOutput.SetText("Wrong bytes.")
					return
				}
				calcByte := []byte(c.stringSrcIutput.Text)
				c.stringResultOutput.SetText(string(features.PEC1byte(calcByte)))
			}),
			c.addButtom("2's complement(2 byte)", func() {
				if len(c.stringSrcIutput.Text)%2 != 0 || len(c.stringSrcIutput.Text) == 0 {
					c.stringResultOutput.SetText("Wrong bytes.")
					return
				}
				calcByte := []byte(c.stringSrcIutput.Text)
				c.stringResultOutput.SetText(string(features.PEC2byte(calcByte)))
			}),
			c.stringResultOutput,
		),
	)

	regionWidget3 := widget.NewHBox(
		functionGroup,
	)

	return widget.NewVBox(
		regionWidget1,
		regionWidget2,
		widget.NewHBox(),
		regionWidget3,
	)
}
