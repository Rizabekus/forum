package controllers

import "forum/internal/services"

type Controllers struct {
	Service *services.Services
}

func ControllersInstance(services *services.Services) *Controllers {
	return &Controllers{Service: services}
}
