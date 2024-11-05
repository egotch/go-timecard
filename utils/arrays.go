package utils

import (
	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Checks if passed in value is in the array indicatd
// Returns true/false depending on result
func InArray(val interface{}, array interface{}) bool {
	return AtArrayPosition(val, array) != -1
}

// AtArrayPosition find the int position of val in a Slice
func AtArrayPosition(val interface{}, array interface{}) (index int) {
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				return
			}
		}
	}

	return
}

func MakeLightTextInput(placeholderTxt string) *tview.InputField {

	return tview.NewInputField().
		SetPlaceholder(placeholderTxt).
		SetPlaceholderTextColor(tcell.ColorDarkViolet).
		SetFieldTextColor(tcell.ColorMidnightBlue).
		SetFieldBackgroundColor(tcell.ColorLightSteelBlue)
}
