package main

import "errors"

// This example is simple cargo booking application

type Cargo struct {
	size float64
}

type Voyage struct {
	capacity float64 // The max capacity is sum of cargos sizes
	cargos   []Cargo
}

func (v *Voyage) bookedCargoSize() float64 {
	sum := 0.
	for _, cargo := range v.cargos {
		sum += cargo.size
	}
	return sum
}

func newVoyage(capacity float64) *Voyage {
	return &Voyage{
		capacity: capacity,
		cargos:   make([]Cargo, 0),
	}
}

func simpleMakeBooking(cargo *Cargo, voyage *Voyage) {
	voyage.cargos = append(voyage.cargos, *cargo)
}

// Add Validation sum of cargos size exceed voyage capacity * 1.1
func addSimpleValidationMakeBooking(cargo *Cargo, voyage *Voyage) error {
	maxBooking := voyage.capacity * 1.1
	if maxBooking < voyage.bookedCargoSize()+cargo.size {
		return errors.New("failed to add cargo: exceeded booked cargo size this voyage")
	}
	return nil
}

// consider the validation as a part of policies

type Policy interface {
	isAllowed(cargo *Cargo, voyage *Voyage) bool
}

type overBookingPolicy func()

func (o overBookingPolicy) isAllowed(cargo *Cargo, voyage *Voyage) bool {
	maxBooking := voyage.capacity * 1.1
	if maxBooking < voyage.bookedCargoSize()+cargo.size {
		return false
	}
	return true
}

func newOverBookingPolicy() Policy {
	return overBookingPolicy(func() {})
}

// finally make booking func is following
func makeBooking(cargo *Cargo, voyage *Voyage) error {
	overBookingPolicy := newOverBookingPolicy()
	if !overBookingPolicy.isAllowed(cargo, voyage) {
		return errors.New("failed to add cargo: exceeded booked cargo size this voyage")
	}
	return nil
}

func main() {

}
