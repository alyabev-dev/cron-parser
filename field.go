package main

type CronField struct {
	Name   string
	Values []int
}

type FieldLimit struct {
	Title string
	Min   int
	Max   int
}

const CronFieldsAmount = 5

var fieldLimits = []FieldLimit{
	{"minute", 0, 59},
	{"hour", 0, 23},
	{"day of month", 1, 31},
	{"month", 1, 12},
	{"day of week", 0, 6},
}
