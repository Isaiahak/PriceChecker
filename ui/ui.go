package ui

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context

type D = layout.Dimensions

var tag = new(bool)

var altPressed = false
var ctrlPressed = false

type Stage int

const (
	Start Stage = iota
	Click
	Copy
	Extract
	Parse
	Display
)

var state Stage = Start

func StartUI() {
	go func() {

		w := new(app.Window)
		w.Option(app.Title("PriceChecker"))
		w.Option(app.Size(unit.Dp(400), unit.Dp(600)))

		// ops are the operations from the ui
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)

	}()
	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops

	var startButton widget.Clickable

	theme := material.NewTheme()

	for {

		// detect the type
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			switch state {
			case Start:
			}
			layout.Flex{
				// vertical alignment, from top to bottom
				Axis: layout.Vertical,

				// Empty space is left to the start, i.e at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}
						return margins.Layout(gtx,
							func(gtx C) D {
								var text string
								btn := material.Button(theme, &startButton, text)
								return btn.Layout(gtx)
							},
						)
					},
				),
			)

			checkStart(gtx)

			e.Frame(gtx.Ops)
		case Click:
		case Copy:
		case Extract:
		case Parse:
		case Display:

		case app.DestroyEvent:
			return e.Err
		}
	}
}

// checks if the user input ctrl and alt
func checkStart(gtx C) {
	// area of interest screen
	defer clip.Rect{Max: image.Pt(1920, 1080)}.Push(ops).Pop()

	// processing events that arrive between the last frame and current frame
	for {
		ev, ok := gtx.Event(
			key.Filter{Name: key.NameCtrl},
			key.Filter{Name: key.NameAlt},
		)
		if !ok {
			break
		}

		if x, ok := ev.(key.Event); ok {
			switch x.Name {
			case key.NameCtrl:
				ctrlPressed = (x.State == key.Press)
			case key.NameAlt:
				altPressed = (x.State == key.Press)
			}
		}

		var c color.NRGBA
		if altPressed && ctrlPressed {
			c = color.NRGBA{R: 0xff, A: 0xff}
			log.Print("pressed")
		} else {
			c = color.NRGBA{G: 0xff, A: 0xff}
			log.Print("released")
		}

		paint.ColorOp{Color: c}.Add(ops)
		paint.PaintOp{}.Add(ops)
	}

}
