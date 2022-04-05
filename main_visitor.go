package main

import (
	"antlrjava8/base"
	"antlrjava8/vast"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
)

func main() {
	fs, err := antlr.NewFileStream("h.java")
	if err != nil {
		log.Fatal(err)
	}
	lex := base.NewJava8Lexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := base.NewJava8Parser(tokens)
	visitor := vast.NewVisitor()
	tree := p.CompilationUnit()
	fmt.Println("test")
	fmt.Println(visitor.Visit(tree))

}
