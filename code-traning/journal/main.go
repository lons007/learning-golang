package main

import "fmt"

type Device interface {
	Activate() string
	GetStatus() string
}

type Lamp struct {
	Brightness int
	IsOn       bool
}

func (l *Lamp) Activate() string {
	var light string
	if !l.IsOn {
		l.IsOn = true
		l.Brightness += 50
		light = fmt.Sprintf("Лампа включена, яркость: %d", l.Brightness)
	} else if l.IsOn {

		if l.Brightness+10 <= 100 {
			l.Brightness += 10
		} else {
			l.Brightness = 100
		}
		light = fmt.Sprintf("Яркость лампы увеличена до %d", l.Brightness)

	}
	return light
}

func (l *Lamp) GetStatus() string {
	var status string
	if l.IsOn {
		status = "[вкл]"
	} else if !l.IsOn {
		status = "[выкл]"
	}
	return fmt.Sprintf("Лампа %s, яркость: %d", status, l.Brightness)
}

type Thermostat struct {
	Temperature float64
	TargetTemp  float64
}

func (t *Thermostat) Activate() string {
	if t.Temperature < t.TargetTemp {
		t.Temperature += 1
		return fmt.Sprintf("Температура повышена до %0.0f °C", t.Temperature)
	} else if t.Temperature == t.TargetTemp {
		return fmt.Sprintf("Температура достигла цели %0.0f °C", t.Temperature)
	} else {
		return fmt.Sprintf("Температура выше целевой %0.0f °C", t.Temperature)
	}
}

func (t *Thermostat) GetStatus() string {
	return fmt.Sprintf("Текущая температура: %0.0f, цель: %0.0f", t.Temperature, t.TargetTemp)
}

func ControlDevice(d Device, times int) {
	for range times {
		fmt.Println(d.Activate())
	}
}

func ShowStatus(d Device) {
	fmt.Println(d.GetStatus())
}

func BatchControl(devices []Device, times int) {
	for _, device := range devices {
		ControlDevice(device, times)
	}
}

func main () {
	lamp := &Lamp{Brightness: 0, IsOn: false}
	term := &Thermostat{Temperature: 18, TargetTemp: 22}
	devices := []Device{lamp, term}
	BatchControl(devices, 3)
	ShowStatus(lamp)
	ShowStatus(term)
}