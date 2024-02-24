package main

import (
	// "fmt"

	"image"
	"strings"

	//"image/draw"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	//"golang.org/x/exp/shiny/text"
	//"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"

	// "gioui.org/internal/ops"
	//"gioui.org/io/system"
	"gioui.org/unit"

	// //"gioui.org/layout"
	"gioui.org/op"
	// //"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)


func main() {
    go func() {

    w := app.NewWindow(
	app.Title("Calculator"),
	app.Size(unit.Dp(400), unit.Dp(400)),
    ) 
	err := run(w)
    if err != nil {
	log.Fatal(err)
    }
    os.Exit(0)

    
    }()
    app.Main()
}

func run(window *app.Window) error {

    th := material.NewTheme()
    var ops op.Ops
    var btn0 = new(widget.Clickable)
    var btn1 = new(widget.Clickable)
    var btn2 = new(widget.Clickable)
    var btn3 = new(widget.Clickable)
    var btn4 = new(widget.Clickable)
    var btn5 = new(widget.Clickable)
    var btn6 = new(widget.Clickable)
    var btn7 = new(widget.Clickable)
    var btn8 = new(widget.Clickable)
    var btn9 = new(widget.Clickable)
    var btndot = new(widget.Clickable)
    var btneq = new(widget.Clickable)
    var btnsum = new(widget.Clickable)
    var btnsub = new(widget.Clickable)
    var btnmul = new(widget.Clickable)
    var btndiv = new(widget.Clickable)
    var strabove string = " "
    var strbelow string = " "
    //inset := layout.Inset{Left: 10, Right: 10}
    
    for { 
	switch nEvent := window.NextEvent().(type) {
	
	case app.DestroyEvent:
	    return nEvent.Err
	
	case app.FrameEvent:
	    //log.Println("Hello World!") 
	    gtx := app.NewContext(&ops,  nEvent)
	    layout.Flex{ 
		Axis: layout.Vertical,
		Spacing: layout.SpaceEvenly, 
		}.Layout(gtx,
		layout.Flexed( 0.2, func(gtx layout.Context) layout.Dimensions { 
		    //drawTextBox(&gtx)
		    return layout.Inset{Top: 20, Bottom: 0, Left: 10, Right: 10}.Layout(gtx,  func(gtx layout.Context) layout.Dimensions { 
		    strabove = "Above"
		    textabove := material.Body1(th, strabove)
		    textabove.Alignment = text.End
		    textabove.TextSize = 24
		    return textabove.Layout(gtx)
		    },)
		    },
		),
		layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions { 
		    //strbelow = "Below"
		    return layout.Inset{Top: 0, Bottom: 0, Left: 10, Right: 10}.Layout(gtx,  func(gtx layout.Context) layout.Dimensions {
		    textbelow := material.Body1(th, strbelow)
		    textbelow.MaxLines = 1
		    textbelow.Alignment = text.End
		    textbelow.TextSize = 24
		    return textbelow.Layout(gtx)
		    },)
		    },
		),
		layout.Flexed(0.2, 
		    func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn7.Clicked(gtx){
					log.Println("7")
				    } 
				    tbtn := material.Button(th, btn7, "7")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn8.Clicked(gtx){
					log.Println("8")
				    } 
				    tbtn := material.Button(th, btn8, "8")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn9.Clicked(gtx){
					log.Println("9")
				    } 
				    tbtn := material.Button(th, btn9, "9")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btndiv.Clicked(gtx){
					log.Println("/")
				    } 
				    tbtn := material.Button(th, btndiv, "/")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)
				    })
			    }),
			)
		    },
		),
		layout.Flexed(0.2, 
		    func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
			    layout.Flexed(0.2, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn4.Clicked(gtx){
					log.Println("4")
				    } 
				    tbtn := material.Button(th, btn4, "4")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.2,func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn5.Clicked(gtx){
					log.Println("5")
				    } 
				    tbtn := material.Button(th, btn5, "5")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.2,  func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn6.Clicked(gtx){
					log.Println("6")
				    } 
				    tbtn := material.Button(th, btn6, "6")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.2,  func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btnmul.Clicked(gtx){
					log.Println("x")
				    } 
				    tbtn := material.Button(th, btnmul, "x")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)
				    })
			    }),
			)
		    },
		),
		layout.Flexed(0.2, 
		    func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn1.Clicked(gtx){
					log.Println("1")
				    } 
				    tbtn := material.Button(th, btn1, "1")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn2.Clicked(gtx){
					log.Println("2")
				    } 
				    tbtn := material.Button(th, btn2, "2")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn3.Clicked(gtx){
					log.Println("3")
				    } 
				    tbtn := material.Button(th, btn3, "3")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btnsub.Clicked(gtx){
					log.Println("-")
				    } 
				    tbtn := material.Button(th, btnsub, "-")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)
				    })
			    }),
			)
		    },
		),
		layout.Flexed(0.2, 
		    func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Spacing: layout.SpaceAround}.Layout(gtx,
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btndot.Clicked(gtx){
					log.Println(".")
					if !strings.Contains(strbelow, ".") {
					    strbelow = strbelow + "."
					}
				    } 
				    tbtn := material.Button(th, btndot, ".")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btn0.Clicked(gtx){
					log.Println("0")
					strbelow = strbelow + "0"
				    } 
				    tbtn := material.Button(th, btn0, "0")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btneq.Clicked(gtx){
					log.Println("=")
					strbelow = strbelow + "="
				    } 
				    tbtn := material.Button(th, btneq, "=")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)})
			    }),
			    layout.Flexed(0.25, func (gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(5)).Layout(gtx, func (gtx layout.Context) layout.Dimensions{	
				    for btnsum.Clicked(gtx){
					log.Println("+")
					strbelow = strbelow + "+"
				    } 
				    tbtn := material.Button(th, btnsum, "+")
				    tbtn.Background = color.NRGBA{R: 20, G: 20, B: 20, A: 220}
				    return tbtn.Layout(gtx)
				    })
			    }),
			)
		    },
		),	
	
	    )	
	    

	    nEvent.Frame(gtx.Ops)

	} 
			
    }
}


func drawTitle(th *material.Theme, gtx *layout.Context) {

    title := material.H1(th, "Hello Gio!")
    maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255} 
    title.Color = maroon 
    title.Alignment = text.Middle

    title.Layout(*gtx)

}
func drawRedRect(ops *op.Ops) {
    defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
    paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)

    paint.PaintOp{}.Add(ops)
}

func drawRedRect10PixelsRight(ops *op.Ops) {
    defer op.Offset(image.Pt(200, 0)).Push(ops).Pop()
    drawRedRect(ops)
}

func redButtonBackground(ops *op.Ops) {
    const r = 10 // roundness
    bounds := image.Rect(0, 0, 100, 100)
    clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Push(ops)
    drawRedRect(ops)
}

func drawTextBox(gtx *layout.Context) {

    const r = 10
    bounds := image.Rect(10, -10, 390,100)
    rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
    paint.FillShape(gtx.Ops,  color.NRGBA{A: 255},
	    clip.Stroke{
		    Path:  rrect.Path(gtx.Ops),
		    Width: 1,
	    }.Op(),
    )
}
func drawOutline(gtx *layout.Context) {

    const r = 10
    bounds := image.Rect(10, 10, 100, 100)
    rrect := clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}
    paint.FillShape(gtx.Ops,  color.NRGBA{A: 255},
	    clip.Stroke{
		    Path:  rrect.Path(gtx.Ops),
		    Width: 1,
	    }.Op(),
    )
}
