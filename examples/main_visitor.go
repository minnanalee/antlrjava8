package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/minnanalee/antlrjava8/base"
	"github.com/minnanalee/antlrjava8/vast"
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
	fmt.Println(tree.ToStringTree(nil, p))
	fmt.Println(visitor.Visit(tree))

}
