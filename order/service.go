package main

import "context"

type service struct {
	orderStore *store
}

func NewOrderService(orderStore *store) *service {
	return &service{
		orderStore,
	}
}

func (s *service) CreateService(context.Context) error {
	return nil
}
