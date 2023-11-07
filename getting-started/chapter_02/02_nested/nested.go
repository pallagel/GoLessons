package main

import "fmt"

// simple struct
type messageToSend struct {
	message string
	name    string
	number  int
}

// function to check if name or number is empty
func canSendMessage(mToSend messageToSend) bool {
	if mToSend.name == "" || mToSend.number == 0 {
		return false
	}

	return true
}

// Test function
func test(mToSend messageToSend) {
	fmt.Printf("Sending '%s' from %s (%v) ....",
		mToSend.message,
		mToSend.name,
		mToSend.number,
	)

	fmt.Println()

	if canSendMessage((mToSend)) {
		fmt.Println("...... message sent")
	} else {
		fmt.Println(".... can't send message")
	}
	fmt.Println("**********************")
}

// main function for the application
func main() {

	test(messageToSend{
		message: "you have a doc app tomorrow",
		name:    "Charlie Brow",
		number:  6699966969,
	})

	test(messageToSend{
		message: "you have an event tommorow",
		name:    "Suzie Sall",
		number:  0,
	})

	test(messageToSend{
		message: "you have an party tommorow",
		name:    "Njorn Halafax",
		number:  16545550987,
	})

	test(messageToSend{
		message: "you have a birthday tommorow",
		name:    "Eli Halafax",
		number:  0,
	})
}
