package main

import "gosh/shell"
import "fmt"
import "os"
import "bufio"



func main() {
	
	GOSH_ENV, errs := shell.GetBasicShellInfo()
	if ContainsValues(errs) {
		fmt.Println("Could not get basic info")
		os.Exit(1)
	}
	
	showWelcome()
	exitChannel := make(chan int)
	go prompt(exitChannel, GOSH_ENV)
	exitVal := <-exitChannel
	os.Exit(exitVal)
}

func ContainsValues(nullable []error) (bool) {
	for _, item := range nullable {
		if item != nil {
			return true
		}
	}
	return false
}

func showWelcome() {
	fmt.Println("\n\nWelcome to gosh the Go shell.\nType help or ? for commands")
}
func prompt(exitChan chan int, GOSH_ENV map[string]string) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("gosh %s@%s %s> ", GOSH_ENV["user"], GOSH_ENV["host"], GOSH_ENV["pwd"])
		command, _, _ := stdin.ReadLine()
		switch string(command) {
			case "?":
				fmt.Println("HELP!")
			case "q":
				exitChan <- 0
				return
			default:
				fmt.Printf("Unknown command [%s]\n", command)
		}
	}
}