package main

import "context"

func main() {
	store := NewOrderStore()
	svc := NewOrderService(store)

	svc.CreateService(context.Background())
}
