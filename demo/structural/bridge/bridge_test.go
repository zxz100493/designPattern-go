package bridge

func ExampleCommonSMS() {
	m := NewCommonMesage(ViaSMS())
	m.SendMessage("hava a drink?", "bob")
}

func ExampleCommonEmail() {
	m := NewCommonMesage(ViaEmail())
	m.SendMessage("hava a drink?", "bob")
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("hava a drink?", "bob")
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("hava a drink?", "bob")
}
