package models

type Profile struct {
	ID     int
	Name   string
	Health int
	Power  int
	Exp    int
}

var Heroes = []Profile{
	{
		Name:   "Goku",
		Health: 200,
		Power:  100,
		Exp:    0,
	},
	{
		Name:   "Sasuke",
		Health: 400,
		Power:  300,
		Exp:    100,
	},
	{
		Name:   "Naruto",
		Health: 150,
		Power:  200,
		Exp:    50,
	},
}