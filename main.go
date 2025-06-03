package main

import (
	"quin/state"
)

func main() {
	// Initialize the application state
	state.Setup()

	// Here you can add your application logic, for example:
	// fmt.Println("AI Data:", state.Config.AIData)
	// fmt.Println("Secrets:", state.Secrets.Secrets)
}
