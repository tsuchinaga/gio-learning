package main

import (
	"image/color"
	"log"
	"os"
	"reflect"

	"gioui.org/text"

	"gioui.org/io/system"
	"gioui.org/layout"

	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/widget/material"

	"gioui.org/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	go func() {
		w := app.NewWindow()
		defer w.Close()
		if err := loop(w); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}()
	// app.Main()
	<-make(chan error)
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops
	for {
		e := <-w.Events()
		log.Printf("%+v, %+v\n", reflect.TypeOf(e), e)

		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			l := material.H1(th, "Hello, Gio!!!")
			l.Color = color.RGBA{R: 0x88, G: 0x00, B: 0x00, A: 0xFF}
			l.Alignment = text.Middle
			l.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
