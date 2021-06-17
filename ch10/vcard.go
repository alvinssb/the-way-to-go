package main

import "time"

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddon string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtDate  time.Time
	Photo     string
	Addresses map[string]*Address
}
