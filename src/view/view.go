package view

import (
	"reflect"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

type View struct {
	MainWindow *gtk.ApplicationWindow `gtk:"main_window"`

	MainHeaderBar *gtk.HeaderBar `gtk:"main_header_bar"`

	ButtonStack *gtk.Stack `gtk:"button_stack"`

	VerifyButton *gtk.Button `gtk:"verify_button"`
	CancelButton *gtk.Button `gtk:"cancel_button"`

	FileChooserButton *gtk.FileChooserButton `gtk:"file_chooser_button"`
	HashValueEntry    *gtk.Entry             `gtk:"hash_value_entry"`

	ErrorDialog      *gtk.MessageDialog `gtk:"error_dialog"`
	ResultOkDialog   *gtk.MessageDialog `gtk:"result_ok_dialog"`
	ResultFailDialog *gtk.MessageDialog `gtk:"result_fail_dialog"`
}

func New() *View {
	return &View{}
}

func (view *View) Activate(app *gtk.Application, uiResource string) {
	// Load UI from resources
	builder, err := gtk.BuilderNewFromResource(uiResource)
	if err != nil {
		panic(err)
	}

	// Get widgets from UI definition
	vStruct := reflect.ValueOf(view).Elem()
	for i := 0; i < vStruct.NumField(); i++ {
		field := vStruct.Field(i)
		structField := vStruct.Type().Field(i)
		widget := structField.Tag.Get("gtk")

		if widget == "" {
			continue
		}

		obj, err := builder.GetObject(widget)
		if err != nil {
			panic(err)
		}

		field.Set(reflect.ValueOf(obj).Convert(field.Type()))
	}

	// Show main window
	view.MainWindow.Present()

	// Add main window
	app.AddWindow(view.MainWindow)
}

func (view *View) notify(title, body string) {
	notification := glib.NotificationNew(title)
	notification.SetBody(body)
	app, _ := view.MainWindow.GetApplication()
	app.SendNotification(app.GetApplicationID(), notification)
}
