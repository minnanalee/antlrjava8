package vast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/minnanalee/antlrjava8/base"
	"reflect"
)

type TreeShapeListener struct {
	*base.BaseJava8ParserListener
	ClassName   []string
	Annotations []string
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// EnterTypeName is called when production typeName is entered.
func (s *TreeShapeListener) EnterTypeName(ctx *base.TypeNameContext) {
	parent := ctx.GetParent()
	for ; parent != nil; parent = parent.GetParent() {
		switch parentType := parent.GetPayload().(type) {
		case *antlr.BaseParserRuleContext:
			if RuleNames[parentType.GetRuleIndex()] == "normalClassDeclaration" {
				//s.ClassName=
			}
			fmt.Println("Ancestors Text:\t", parentType.GetText())
		}
	}

}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("Node RuleName\t", RuleNames[ctx.GetRuleContext().GetRuleIndex()])
	fmt.Println("Node Context: \t", ctx.GetText())
	fmt.Println("Node Type: \t", fmt.Sprint(reflect.TypeOf(ctx)))

	for i := 0; i < ctx.GetChildCount(); i++ {
		fmt.Println("Child Type", i, ": \t", fmt.Sprint(reflect.TypeOf(ctx.GetChild(i))))
		//fmt.Println("Child ",i,":\t", ctx.GetChild(i).)
	}
	//非叶子节点不打印全链路
	if !(ctx.GetChildCount() == 1 && fmt.Sprint(reflect.TypeOf(ctx.GetChild(0).GetPayload())) == "*antlr.CommonToken") {
		return
	}
	parent := ctx.GetParent()
	for ; parent != nil; parent = parent.GetParent() {
		switch parentType := parent.GetPayload().(type) {
		case *antlr.BaseParserRuleContext:
			fmt.Println("Ancestors Type:\t", RuleNames[parentType.GetRuleIndex()])
			fmt.Println("Ancestors Text:\t", parentType.GetText())
		}
	}
}

/*func main() {
	input, _ := antlr.NewFileStream("h.java")
	lexer := base.NewJava8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := base.NewJava8Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

}

*/
