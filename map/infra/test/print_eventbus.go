package test

import "fmt"

type PrintEventbus struct {
}

func (p *PrintEventbus) Publish(event any) {
	fmt.Printf("Publish Event: %+v\n", event)
}

func (p *PrintEventbus) Subscribe(cb func(any)) {
	fmt.Printf("Subscribe Event: %p\n", cb)
}
