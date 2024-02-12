package main
//
//import (
//	"fmt"
//	//"image/color"
//	//"os"
//	//"log"
//	"gioui.org/app"
//	//"gioui.org/io/system"
//	//"gioui.org/layout"
//	//"gioui.org/op"
//	//"gioui.org/text"
//	//"gioui.org/widget/material"
//
//)
//
//
//func main () {
//
//	fmt.Println("Hello World!")
//
//	go func() {
//		w := app.NewWindow()
//		
//		for range w.Events() {
//		}
//	}()
//	app.Main()
//}
import (
  "gioui.org/app"
)

func main() {
  go func() {
    // create new window
    w := app.NewWindow()

    // listen for events in the window.
    for { 
	w.NextEvent()
			
    }
  }()
  app.Main()
}

