package main

import "fmt"
type TokenType int

const (
	INTEGER=iota
	PLUS
	MINUS
	EOF
)

type Token struct {
	Type TokenType
	Value []byte
}


func main() {
	t := Token{Type:INTEGER, Value:[]byte("Test")}
	fmt.Println("vim-go")
	fmt.Printf("%+v\n", t)
	i := Interpreter{}
	fmt.Printf("%#v\n", i)
}
