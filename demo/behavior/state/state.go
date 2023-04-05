package state

import "fmt"

type Week interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{today: &Sunday{}}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct {
}

func (s *Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (s *Sunday) Next(d *DayContext) {
	d.today = &Monday{}
}

type Monday struct {
}

func (s *Monday) Today() {
	fmt.Printf("Monday\n")
}

func (s *Monday) Next(d *DayContext) {
	d.today = &Tuesday{}
}

type Tuesday struct {
}

func (s *Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (s *Tuesday) Next(d *DayContext) {
	d.today = &Monday{}
}

type Wednesday struct {
}

func (s *Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (s *Wednesday) Next(d *DayContext) {
	d.today = &Thursday{}
}

type Thursday struct {
}

func (s *Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (s *Thursday) Next(d *DayContext) {
	d.today = &Friday{}
}

type Friday struct {
}

func (s *Friday) Today() {
	fmt.Printf("Friday\n")
}

func (s *Friday) Next(d *DayContext) {
	d.today = &Monday{}
}

type Saturday struct {
}

func (s *Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (s *Saturday) Next(d *DayContext) {
	d.today = &Sunday{}
}
