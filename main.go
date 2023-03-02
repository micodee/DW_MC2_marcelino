package main

import (
	"fmt"
	"heroes/controllers"
)

func main() {
	profile := controllers.MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := controllers.PowerUp(profile, 5)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}
