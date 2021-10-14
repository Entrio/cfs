// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue

	p2 := widgets.NewParagraph()
	p2.Text = "Press q to quit\nPress h or l to switch tabs\n"
	p2.Title = "Keys"
	p2.SetRect(5, 5, 40, 15)
	p2.BorderStyle.Fg = ui.ColorYellow

	bc := widgets.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []float64{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(5, 5, 35, 10)
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}

	tabpane := widgets.NewTabPane("pierwszy", "drugi", "trzeci", "żółw", "four", "five")
	tabpane.SetRect(0, 1, 60, 4)
	tabpane.Border = true

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] foo",
		"[8] bar",
		"[9] baz",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(5, 5, 40, 13)

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(p2)
		case 1:
			ui.Render(bc)
		case 2:
			ui.Render(l)
		}
	}

	ui.Render(header, tabpane, p2)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Left>":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "<Right>":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "j", "<Down>":
			ui.Clear()
			ui.Render(header, tabpane)
			l.ScrollDown()
			renderTab()
		case "k", "<Up>":
			ui.Clear()
			ui.Render(header, tabpane)
			l.ScrollUp()
			renderTab()
		}
	}
}
