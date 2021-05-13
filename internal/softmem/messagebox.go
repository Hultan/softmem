package softmem

import (
	"github.com/gotk3/gotk3/gtk"
)

//buttons := []messagebox.Button{{ "Absolutely!", gtk.RESPONSE_OK}, {"Hell no!", gtk.RESPONSE_CANCEL}}
//box := messagebox.NewMessageBoxWithButtons("Update statistics?", "Do you want to update statistics?", m.window, buttons)
//
//if box.Question() == gtk.RESPONSE_OK {
//	do stuff
//}

type MessageBox struct {
	Header  string
	Message string
	Parent  gtk.IWindow
	Buttons []Button
}

type Button struct {
	Text         string
	ResponseType gtk.ResponseType
}

func NewMessageBox(header, message string, parent gtk.IWindow) *MessageBox {
	box := new(MessageBox)
	box.Header = header
	box.Message = message
	box.Parent = parent
	return box
}

func NewMessageBoxWithButtons(header, message string, parent gtk.IWindow, buttons []Button) *MessageBox {
	box := new(MessageBox)
	box.Header = header
	box.Message = message
	box.Parent = parent
	box.Buttons = buttons
	return box
}

func (m *MessageBox) Information() {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_NONE, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	for _, button := range m.Buttons {
		dialog.AddButton(button.Text, button.ResponseType)
	}
	dialog.Run()
}

func (m *MessageBox) InformationOK() {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	dialog.Run()
}

func (m *MessageBox) Question() gtk.ResponseType {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_QUESTION, gtk.BUTTONS_NONE, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	for _, button := range m.Buttons {
		dialog.AddButton(button.Text, button.ResponseType)
	}
	return dialog.Run()
}

func (m *MessageBox) QuestionYesNo() bool {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_QUESTION, gtk.BUTTONS_YES_NO, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	return dialog.Run() == gtk.RESPONSE_YES
}

func (m *MessageBox) QuestionOkCancel() bool {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_QUESTION, gtk.BUTTONS_OK_CANCEL, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	return dialog.Run() == gtk.RESPONSE_OK
}

func (m *MessageBox) Warning() gtk.ResponseType {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_NONE, m.Message)
	defer dialog.Destroy()
	for _, button := range m.Buttons {
		dialog.AddButton(button.Text, button.ResponseType)
	}
	return dialog.Run()
}

func (m *MessageBox) WarningOKCancel() bool {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_OK_CANCEL, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	return dialog.Run() == gtk.RESPONSE_OK
}

func (m *MessageBox) WarningYesNo() bool {
	dialog := gtk.MessageDialogNewWithMarkup(m.Parent, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_YES_NO, m.Message)
	defer dialog.Destroy()
	dialog.SetTitle(m.Header)
	return dialog.Run() == gtk.RESPONSE_YES
}
