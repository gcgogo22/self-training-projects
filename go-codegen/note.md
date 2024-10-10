# Code Gen

[Code Gen Link](https://www.youtube.com/watch?v=ClW_g1iDGi4)

    main.go

    package main
    //go:generate echo xxx  // -> Run with <code>go generate</code> command 
    //go:generate bash -c "env | grep PWD"
    //go:generate -command foo echo xxx  // alias
    //go:generate foo

- Need to be valid executable
- Can't not do the full shell scripting here. 
- 