package main


type Interpreter struct {
	Text string
	pos int
	current_char byte
	current_token Token
}
