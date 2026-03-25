package main

import "fmt"

type Device interface {
	Activate() string
	GetStatus() string
}

type Lamp struct {
	Brightnes int
	IsOn bool
}

func (l *Lamp) Activate() string {
	var light string
	if !l.IsOn {
		l.IsOn = true
		light = fmt.Sprintf("Лампа включена, яркость: %d", l.Brightnes + 50)
	} else if l.IsOn {
		light = fmt.Sprintf("Яркость лампы увеличена до %d", l.Brightnes + 10)
	}
	return light
}

func (l Lamp) GetStatus() string {
	var status string
	if l.IsOn {
		status = "[вкл]"
	} else if !l.IsOn {
		status = "[выкл]"
	}
	return fmt.Sprintf("Лампа %s, яркость: %d", status, l.Brightnes)
}

