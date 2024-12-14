package main

import "context"

type OrderService interface {
	CreateService(context.Context) error
}

type OrderStore interface {
	Create(context.Context) error
}
