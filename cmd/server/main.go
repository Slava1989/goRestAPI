package main

import "fmt"

// Run - is going to be responsible for
// the instantiation and startup of
// go app
func Run() error {
	fmt.Println("starting up")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println((err))
	}
}