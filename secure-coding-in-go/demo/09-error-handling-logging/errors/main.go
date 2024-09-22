package main

import "fmt"

func main() {
	start()
	fmt.Println("Returned normally from start().")
}

func start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in start()")
		}
	}()
	fmt.Println("Called start()")
	part2(0)
	// Note this line is not called because of panic.
	fmt.Println("Returned normally from part2().")
}

func part2(i int) {
	// note when i == 1, the defer function is not registered, so
	// "Defer in part2()" is only called once.
	if i > 0 {
		fmt.Println("Panicking in par2()!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in part2()")
	fmt.Println("Executing part2()")
	part2(i + 1)
}


// For log.Fatal, defer function is not executed, buffer will not be flushed, and temporary files and directories will not be removed. Differ from panic.