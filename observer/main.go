package main

func main() {
	shirtItem := newItem("Nike Shirt")
	observerFirst := &customer{
		id: "santos@gmail.com",
	}
	observerSecond := &customer{
		id: "mendax@gmail.com",
	}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability()
}
