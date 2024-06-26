package librarytutorial

type Student struct {
	Name  string
	grade int // ! if it's called to the main.go, it will error because use normal letters
}

// ! Public (Capital Letters), private (normal letters)
