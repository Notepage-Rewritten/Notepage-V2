package components

import "fyne.io/fyne/v2/widget"

type button struct {
	button *widget.Button
}

func NewButton(buttonText string) button {
	buttonObj := button{}

	buttonObj.button = widget.NewButton(buttonText, func() {})
	return buttonObj
}

// GetButton returns the button widget
func (b button) GetButton() *widget.Button {
	return b.button
}

// SetEventHandler overrides whatever function the tapped event fires
func (b button) SetEventHandler(event func()) {
	b.button.OnTapped = event
}
