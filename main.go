package main

import "goddb/cmd"

func main() {
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
