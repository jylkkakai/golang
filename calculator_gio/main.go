
package main 
import ( 
    // "fmt" 
    
    //"image/color" 
    //"image/draw" 
    "log"
    "os"

    "gioui.org/app"
    // "gioui.org/internal/ops"
    "gioui.org/io/system"

    // //"gioui.org/layout"
    "gioui.org/op"
    // //"gioui.org/text"
    //"gioui.org/widget/material"
)


func main() {
    go func() {

	w := app.NewWindow()
	err := run(w)
	if err != nil {
	    log.Fatal(err)
	}
	os.Exit(0)

    
    }()
    app.Main()
}

func run(window *app.Window) error {

    //////th := material.NewTheme()
    var ops op.Ops
    for { 
	switch nEvent := window.NextEvent().(type) {

	case system.DestroyEvent:
	    return nEvent.Err

	case system.FrameEvent:
	    log.Println("Hello World!")

	    //title := material.H1(th, "Hello Gio!")

	    ops.Reset()

	    //draw(&ops)

	    nEvent.Frame(&ops)

	} 
			
    }
}

