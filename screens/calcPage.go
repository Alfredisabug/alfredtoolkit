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

const (
	inType = iota
	hasOperator
	inFormula
)

type calc struct {
	output             *widget.Label
	formulaOutput      *widget.Label
	bitOutput          *widget.Label
	errOutput          *widget.Label
	stringSrcIutput    *widget.Entry
	stringResultOutput *widget.Label
	buttons            map[string]*widget.Button
	radios             map[string]*widget.Radio
	functions          map[string]func()
	typeStatus         int
}

func (c *calc) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button
	c.functions[text] = action
	return button
}

// +-*/專用
func (c *calc) operatorBtnFunc(text string) {
	if c.typeStatus == inFormula {
		temp := c.formulaOutput.Text
		if c.output.Text != "" {
			c.formulaOutput.SetText(temp + c.output.Text + text)
			c.output.SetText("")
			c.typeStatus = inFormula
			return
		}
		c.formulaOutput.SetText(temp[0:len(temp)-1] + text)
	}
	if c.typeStatus == inType {
		c.formulaOutput.SetText(c.formulaOutput.Text + c.output.Text + text)
		c.output.SetText("")
		c.typeStatus = inFormula
	}
}

//  數字專用
func (c *calc) numBtnFunc(text string) {
	if c.formulaOutput.Text != "" {
		c.typeStatus = inFormula
	}
	c.output.SetText(c.output.Text + text)
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
	c.typeStatus = inType

	c.formulaOutput = widget.NewLabel("")
	c.formulaOutput.TextStyle.Monospace = true
	c.formulaOutput.Alignment = fyne.TextAlignTrailing

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
	c.functions = make(map[string]func())
	return c
}

func makeCalcPageTab(win fyne.Window) fyne.CanvasObject {
	c := newCalc()
	win.Canvas().SetOnTypedRune(func(r rune) {
		action := c.functions[string(r)]
		if action != nil {
			action()
		}
	})
	win.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		if ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter {
			// TODO:
			// 	計算結果
			c.formulaOutput.SetText("")
			c.output.SetText("")
			c.typeStatus = inType
		}
		if ev.Name == fyne.KeyDelete || ev.Name == fyne.KeyBackspace {
			if c.output.Text != "" {
				c.output.SetText(c.output.Text[0 : len(c.output.Text)-1])
			}
		}
	})

	regionWidget1 := widget.NewHBox(
		c.addRadio("digitShow", []string{"BIN", "DEC", "HEX"}, false),
		layout.NewSpacer(),
		widget.NewVBox(
			c.formulaOutput,
			layout.NewSpacer(),
			c.output,
		),
	)

	regionWidget2 := widget.NewVBox(
		c.bitOutput,
		c.addRadio("bitsShow", []string{"Byte", "short", "Int32", "Int 64"}, true),
	)

	functionGroup := widget.NewGroup("PEC Function",
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			c.stringSrcIutput,
			c.addButton("2's complement(1 byte)", func() {
				if len(c.stringSrcIutput.Text)%2 != 0 || len(c.stringSrcIutput.Text) == 0 {
					c.stringResultOutput.SetText("Wrong bytes.")
					return
				}
				calcByte := []byte(c.stringSrcIutput.Text)
				c.stringResultOutput.SetText(string(features.PEC1byte(calcByte)))
			}),
			c.addButton("2's complement(2 byte)", func() {
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

	calcGroup := widget.NewGroup("calcator",
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(4),
				c.addButton("7", func() {
					c.numBtnFunc("7")
				}),
				c.addButton("8", func() {
					c.numBtnFunc("8")
				}),
				c.addButton("9", func() {
					c.numBtnFunc("9")
				}),
				c.addButton("+", func() {
					c.operatorBtnFunc("+")
				}),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(4),
				c.addButton("4", func() {
					c.numBtnFunc("4")
				}),
				c.addButton("5", func() {
					c.numBtnFunc("5")
				}),
				c.addButton("6", func() {
					c.numBtnFunc("6")
				}),
				c.addButton("-", func() {
					c.operatorBtnFunc("-")
				}),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(4),
				c.addButton("1", func() {
					c.numBtnFunc("1")
				}),
				c.addButton("2", func() {
					c.numBtnFunc("2")
				}),
				c.addButton("3", func() {
					c.numBtnFunc("3")
				}),
				c.addButton("*", func() {
					c.operatorBtnFunc("*")
				}),
			),
			fyne.NewContainerWithLayout(
				layout.NewGridLayout(4),
				c.addButton("", func() {}),
				c.addButton("0", func() {
					c.numBtnFunc("0")
				}),
				c.addButton("", func() {}),
				c.addButton("/", func() {
					c.operatorBtnFunc("/")
				}),
			),
		),
	)

	regionWidget3 := widget.NewHBox(
		functionGroup,
		calcGroup,
	)

	return widget.NewVBox(
		regionWidget1,
		regionWidget2,
		widget.NewHBox(),
		regionWidget3,
	)
}
