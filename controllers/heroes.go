package controllers

import "heroes/models"

func MakeProfile(name string) *models.Profile {
	for _, hero := range models.Heroes {
		if hero.Name == name {
			return &models.Profile{
				Name:   hero.Name,
				Health: hero.Health,
				Power:  hero.Power,
				Exp:    hero.Exp,
			}
		}
	}
	return nil
}

func PowerUp(pwr *models.Profile, multiplier int) *models.Profile {
	pwr.Health += pwr.Health * multiplier
	pwr.Power += pwr.Power * multiplier
	pwr.Exp += pwr.Exp * multiplier
	return pwr
}