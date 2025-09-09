package ui

import (
	"image"
	"image/color"
	"log"
	"os"

	"golang.design/x/clipboard"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
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

type Stage int

const (
	Start Stage = iota
	Click
	Copy
	Extract
	Parse
	Modify
	Search
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

	theme := material.NewTheme()

	for {

		// detect the type
		switch e := w.Event().(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			switch state {
			
			case Start:
				// checks if alt and ctrl/ cmd are clicked
				checkStart(ops,gtx)
				e.Frame(gtx.Ops)
			case Click:
				// waits for the user to copy the item to the clipboard

				// layout for copy prompt
				layout.Flex{
					Axis: layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(
						func(gtx C) D{
							rect := clip.Rect{Max: image.Pt(100,100)}.Push(ops).Pop()
							rectColor := color.NRGBA{R:98, G:98, B:98, A:255}
							paint.FillShape(gtx.Ops, reactColor, )
						},
					),
				)

				checkClick(ops,gtx)
				e.Frame(gtx.Ops)
			case Copy:
				copyItem()
				e.Frame(gtx.Ops)
			case Extract:
				checkItem()
				e.Frame(gtx.Ops)
			case Parse:
				parseItem()
				e.Frame(gtx.Ops)
			case Modify:
				// we need to pass the attributes of the item to the modify layout for dynamically 
				// creating the different selectors and modifiers
				/*
				func modifyLayout(gtx C) D {
					return layout.Flex{}.Layout{gtx,


					}
				}
				*/
				modifyItem()
				e.Frame(gtx.Ops)
			case Search:
				searchItem()
				e.Frame(gtx.Ops)
			case Display:
				displayResult()
				e.Frame(gtx.Ops)
		}

		case app.DestroyEvent:
			return e.Err
		}
	}
}

// checks if the user input ctrl and alt
func checkStart(ops *op.Ops, gtx C) {
	defer clip.Rect{Max: image.Pt(1920, 1080)}.Push(ops).Pop()

	var altPressed = false
	var ctrlPressed = false
	var cmdPressed = false

	for {
		ev, ok := gtx.Event(
			key.Filter{Name: key.NameCtrl},
			key.Filter{Name: key.NameAlt},
			key.Filter{Name: key.NameCommand},
		)
		if !ok {
			break
		}
		// sets the pressed bools
		if x, ok := ev.(key.Event); ok {
			switch x.Name {
			case key.NameCtrl:
				ctrlPressed = (x.State == key.Press)
			case key.NameAlt:
				altPressed = (x.State == key.Press)
			case key.NameCommand:
				cmdPressed = (x.State == key.Press)
			}

		}

		// checks if ctrl or cmd and alt are pressed
		if altPressed && (ctrlPressed || cmdPressed) {
			state = Click
			log.Print("both are pressed")
		} else {
			state = Start
			log.Print("both are not pressed")
		}
	}
}

// checks if the user has clicked on the screen
func checkClick(ops *op.Ops, gtx C) {
	defer clip.Rect{Max: image.Pt(1920,1080)}.Push(ops).Pop()
	
	var ctrl = false
	var c = false
	var click = false

	for {
		ev, ok := gtx.Event(
			key.Filter{Name: key.NameCtrl},
			key.Filter{Name: "c"},
			pointer.Filter{
				Target: tag,
				Kinds: pointer.Press,
			},
		)

		if !ok {
			break
		}

		switch x := ev.(type){

		case pointer.Event: 
			if x.Kind == pointer.Press {
				click = true
			}

		case key.Event:
			if x.State == key.Press {
				switch x.Name {
				case key.NameCtrl:
					ctrl = true
				case "c":
					c = true
				}
			}
		}
		 
		if c && ctrl && click {
			state = Copy
			log.Print(" user tried to copy item")
		} else {
			state = Click
			log.Print(" no copy was detected")
		}
	}
}

// copies the object at the mouse location to the clip board
func copyItem() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Read(clipboard.FmtText)
}
// checks if the item copied exists if not prompts user to click again
func checkItem(){
}
// formats the copied item to be used for the item search
func parseItem(){
}
// allows user to modify the properities of the item they clicked
func modifyItem(){
}
// looks for the item on the trade site, taking note of a number of prices associated with the states on the item
func searchItem(){
}
// display the results of the search
func  displayResult()



