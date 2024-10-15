package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/egotch/go-timecard/utils"
)

type TimeDetailPane struct {
	*tview.Flex
	list               *tview.List
	newTimeItem        *tview.InputField
	timeDetailStarting int // the index in list where time entries start
}

// Initilializes the TimeDetailPane
func NewTimeDetailPane() *TimeDetailPane {

	// inite the time detail pane
	pane := TimeDetailPane{
		Flex:        tview.NewFlex().SetDirection(tview.FlexRow),
		list:        tview.NewList().ShowSecondaryText(false),
		newTimeItem: utils.MakeLightTextInput("+[New Time Entry]"),
	}

	// prompt to start with entering an initial time entry
	// add the item created
	pane.newTimeItem.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			fmt.Println("Add a new time log item here")
		case tcell.KeyEsc:
			app.SetFocus(timeDetailPane)
		}

	})

	pane.AddItem(pane.list, 0, 1, true).
		AddItem(pane.newTimeItem, 1, 0, false)

	// set up the pane (Boarder = true, give it a title)
	pane.SetBorder(true).SetTitle("[::u]T[::-]ime Entry")

	// reload the pane
	//pane.loadListItems(false)

	return &pane
}

// struct method for adding new time entry

// struct method for loading the items (see loadListItems on geek-life)
