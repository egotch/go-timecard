package main

import (
	"reflect"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/egotch/go-timecard/utils"
)

var (
	app    *tview.Application
	layout *tview.Flex

	timeDetailPane *TimeDetailPane
	statusBar      *StatusBar
)

func main() {

	app = tview.NewApplication()

	// Set up the app layout
	// [             Title Bar            ]
	// [Granular Details] [ [Summary]     ]
	// [                ] [ [Punch in/out ]
	// [            Key Bindings          ]
	layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(makeTitleBar(), 1, 1, false).
		AddItem(makeContentPane(), 0, 10, false).
		AddItem(prepareStatusBar(app), 1, 1, false)

	setKeyboardShortcuts()

	if err := app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func ignoreKeyEvt() bool {
	textInputs := []string{"*tview.InputField", "*femto.View"}
	return utils.InArray(reflect.TypeOf(app.GetFocus()).String(), textInputs)
}

// Set custom keybindings for the application
// h - display help modal popup
// q - quit - will display a "Do you want to Exit?" questionaire
// a - display a "hello world" modal
func setKeyboardShortcuts() *tview.Application {

	e := app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if ignoreKeyEvt() {
			return event
		}

		// Global shortcuts
		switch unicode.ToLower(event.Rune()) {
		case 'q':
			AskYesNo("Exit Application?", app.Stop)
			return nil
		}

		return event
	})
	return e
}

// renders titlebar TextView
func makeTitleBar() *tview.TextView {
	titleText := tview.NewTextView().SetText("[lime::b]Go TimeCard [::-]- Terminal Time Card Manager!").SetDynamicColors(true)

	return titleText
}

// renders the main content pane
// Flex view that's 2 columns
// Col 1 = granular/detail time view
// Col 2 = 2 x Rows:
//
//	Row1 = summary view
//	Row2 = punch in/out view
func makeContentPane() *tview.Flex {

	// detailView := tview.NewBox().SetTitle("Time Card").SetBorder(true)
	timeDetailPane = NewTimeDetailPane()
	summaryView := tview.NewBox().SetTitle("Summary").SetBorder(true)
	punchInOut := tview.NewBox().SetTitle("Punch In/Out").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(timeDetailPane, 0, 4, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(summaryView, 0, 1, false).
			AddItem(punchInOut, 10, 1, false), 0, 1, false)

	return flex
}

// modal pop up
// Asks the question (string) passed in
// and prompts for a yes / no answer
// if "yes" is selected, then the passed in function is executed and focus
// returns to previous panel
// if "no" is selected the modal goes away
func AskYesNo(text string, f func()) {

	activePane := app.GetFocus()
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				f()
			}
			app.SetRoot(layout, true).EnableMouse(true)
			app.SetFocus(activePane)
		})

	pages := tview.NewPages().
		AddPage("background", layout, true, true).
		AddPage("modal", modal, true, true)

	_ = app.SetRoot(pages, true).EnableMouse(true)
}
