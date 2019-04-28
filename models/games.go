package models

type Games struct {
	IdGame int
	Type string
	Participants []int
	Win bool
	Date string
}
