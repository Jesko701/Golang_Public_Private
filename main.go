package main

// * 1 folder can handle lot's of function from all file inside library-tutorial
import library "tutorial-go-access-public-private/library-tutorial"

func main() {
	library.SayHello()     // file from library.go
	library.Testing("Joy") // file from testing.go (ada fungsi untuk memanggil private)
	library.ABC()          // file from testing.go
}
