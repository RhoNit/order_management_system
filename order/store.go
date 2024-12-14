package main

import "context"

type store struct {
	// Todo
}

func NewOrderStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}
