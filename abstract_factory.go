package main

import "fmt"

var DisplayStandard string = "Motif"

type ScrollerFactoryer interface {
	ScrollerTo(position int)
}

type MotifScroller struct{}

type PMScroller struct{}

func NewScrollerFactory() (s ScrollerFactoryer) {
	switch DisplayStandard {
	case "Motif":
		s = &MotifScroller{}
	case "PM":
		s = &PMScroller{}
	}
	return s
}

func (b *MotifScroller) ScrollerTo(position int) {
	fmt.Println("Create MotifScrollerBar")
	fmt.Println("MotifScrollerBar move to position ", position)
}

func (b *PMScroller) ScrollerTo(position int) {
	fmt.Println("Create PMScrollerBar")
	fmt.Println("PMScrollerBar move to position ", position)
}

func main() {
	DisplayStandard = "PM"
	scroFactory := NewScrollerFactory()
	scroFactory.ScrollerTo(12)
}
