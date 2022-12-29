package main

import o "observer"

func main() {
	shirtItem := o.NewItem("Nike shirt")
	observerFirst := &o.Customer{ID:"abc@gmail.com"}
	observerSecond := &o.Customer{ID:"avc@gmail.com"}
	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)
	shirtItem.UpdateAvailability()
	
}