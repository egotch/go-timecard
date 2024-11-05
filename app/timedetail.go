package main

import (
	"fmt"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

  "github.com/egotch/go-timecard/model"
	"github.com/egotch/go-timecard/utils"
)

type TimeDetailPane struct {
	*tview.Flex
  timeEntries       []model.TimeEntry
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
			pane.addNewTimeEntry()
		case tcell.KeyEsc:
      pane.newTimeItem.SetText("")
			app.SetFocus(timeDetailPane)
		}

	})

	pane.AddItem(pane.list, 0, 1, true).
		AddItem(pane.newTimeItem, 1, 0, false)

	// set up the pane (Boarder = true, give it a title)
	pane.SetBorder(true).SetTitle(" [::u]T[::-]ime Entry ")

	// reload the pane
	//pane.loadListItems(false)

	return &pane
}

// struct method for adding new time entry
func (pane *TimeDetailPane) addNewTimeEntry() {

  var entry model.TimeEntry

	entry.Description = pane.newTimeItem.GetText()

	statusBar.showForSeconds(fmt.Sprintf("[yellow::]Time entry %s added!", entry.Description), 10)
  pane.timeEntries = append(pane.timeEntries, entry)
	pane.addTimeEntryToList(len(pane.timeEntries)-1, true)
	pane.newTimeItem.SetText("")

}

func (pane *TimeDetailPane) addTimeEntryToList(i int, selectItem bool) {

  pane.list.AddItem("- "+pane.timeEntries[i].Description, "", 0, nil)  

  if selectItem {
    fmt.Println("hello")
    // pane.list.SetCurrentItem(-1)
  }

}


// Keybindings for time detail pane
// 'n' -> go to newTimeItem pane
func (pane *TimeDetailPane) handlKeyBindings(event *tcell.EventKey) *tcell.EventKey {

	switch unicode.ToLower(event.Rune()) {

	case 'n':
		app.SetFocus(pane.newTimeItem)
		return nil

  case 'j':
    pane.list.SetCurrentItem(pane.list.GetCurrentItem() + 1)
    return nil

  case 'k':
    pane.list.SetCurrentItem(pane.list.GetCurrentItem() - 1)
    return nil

	}

	return event
}
