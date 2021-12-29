package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type cdata struct {
	mem int
	cal string
	flg bool
}

// createNumButtons create number buttons.
func createNumButtons(f func(int)) *fyne.Container {
	numLayout := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		widget.NewButton(strconv.Itoa(7), func() { f(7) }),
		widget.NewButton(strconv.Itoa(8), func() { f(8) }),
		widget.NewButton(strconv.Itoa(9), func() { f(9) }),
		widget.NewButton(strconv.Itoa(4), func() { f(4) }),
		widget.NewButton(strconv.Itoa(5), func() { f(5) }),
		widget.NewButton(strconv.Itoa(6), func() { f(6) }),
		widget.NewButton(strconv.Itoa(1), func() { f(1) }),
		widget.NewButton(strconv.Itoa(2), func() { f(2) }),
		widget.NewButton(strconv.Itoa(3), func() { f(3) }),
		widget.NewButton(strconv.Itoa(0), func() { f(0) }),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() { f(10) }),
	)
	return numLayout
}

// createNSignButtons create all numerical operation buttons.
func createOperatorButtons(f func(string)) *fyne.Container {
	signLayout := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewButton("clear", func() { f("clear") }),
		widget.NewButton("/", func() { f("/") }),
		widget.NewButton("*", func() { f("*") }),
		widget.NewButton("+", func() { f("+") }),
		widget.NewButton("-", func() { f("-") }),
	)
	return signLayout
}

// main function.
func main() {
	a := app.New()
	w := a.NewWindow("Calclator")
	w.SetFixedSize(false)
	topNum := widget.NewLabel("0")
	DisplayContainer := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		topNum,
	)

	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}

	calculate := func(num int) {
		switch data.cal {
		case "/":
			defer func() {
				if num == 0 {
					p := recover()
					if p != nil {
						panic("not divisible by 0")
					}
				}
			}()
			data.mem /= num
		case "*":
			data.mem *= num
		case "+":
			data.mem += num
		case "-":
			data.mem -= num
		}
	}

	InputNum := func(num int) {
		stringNum := topNum.Text

		if stringNum == "0" {
			stringNum = strconv.Itoa(num)
		} else if num == 10 && stringNum != "0" {
			stringNum = stringNum[:len(stringNum)-1]
		} else {
			stringNum += strconv.Itoa(num)
		}
		displayNum, err := strconv.Atoi(stringNum)

		if err == nil {
			if data.flg {
				calculate(num)
				topNum.SetText(strconv.Itoa(data.mem))
			} else {
				data.mem = displayNum
				topNum.SetText(strconv.Itoa(displayNum))
			}
		}
	}

	inputOperator := func(operator string) {
		if operator == "clear" {
			topNum.SetText("0")
			data.mem = 0
		} else {
			data.cal = operator
			data.flg = true
		}
	}

	numButton := createNumButtons(InputNum)
	signButton := createOperatorButtons(inputOperator)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				DisplayContainer, nil, nil, signButton,
			),
			DisplayContainer, numButton, signButton,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
