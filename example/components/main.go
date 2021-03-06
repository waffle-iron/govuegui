package main

import (
	"bytes"

	"log"

	"github.com/as27/govuegui"
	"github.com/as27/govuegui/gui/bulma"
)

var gui = govuegui.NewGui(bulma.Template)

func main() {
	inputElements := gui.Form("Elements").Box("Input")
	inText := inputElements.Input("Input Text")
	inText.Set("Some Text")
	inTextArea := inputElements.Textarea("Input Textarea")
	inTextArea.Set(`A longer text
	with more 
	lines...`)
	dropdown := inputElements.Dropdown("My Options")
	dropdown.Option("key1", "First Option")
	dropdown.Option("key2", "Second Option")
	dropdown.Set("")
	preview := inputElements.Text("Preview")
	preview.Set("")
	inTextArea.Action(func(gui *govuegui.Gui) {
		setText(preview, inText, inTextArea)
	})
	dropdown.Action(func(gui *govuegui.Gui) {
		err := preview.Set("Dropdown changed...")
		if err != nil {
			log.Println("Error: ", err)
		}
		preview.Update()
	})
	inText.Action(func(gui *govuegui.Gui) {
		setText(preview, inText, inTextArea)
	})
	outputElements := gui.Form("Elements").Box("Output")
	outputElements.Text("HTML").Set("Some <b>HTML</b> comes <br>here.")
	myList := []string{"Fist entry", "Second list item", "more", "another item"}
	outputElements.List("A simple list").Set(myList)
	myTable := [][]string{
		{"Header1", "Header2", "Header3"},
		{"1-1", "1-2", "1-3"},
		{"2-1", "2-2", "2-3"},
	}
	outputElements.Table("A simple Table").Set(myTable)
	govuegui.Serve(gui)
}

func setText(out *govuegui.Element, ins ...*govuegui.Element) {
	b := &bytes.Buffer{}
	for _, in := range ins {
		b.WriteString(in.Get().(string))
		b.WriteString("<br>----------------------<br>")
	}
	out.Set(b.String())
	out.Update()
}
