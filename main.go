package main

import (
	"antlrjava8/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseJava8ParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(RuleNames[ctx.GetRuleContext().GetRuleIndex()])
	fmt.Println(ctx.GetParent().GetChildCount())
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream("h.java")
	lexer := parser.NewJava8Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJava8Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

var RuleNames = []string{
	"literal", "primitiveType", "numericType", "integralType", "floatingPointType",
	"referenceType", "classOrInterfaceType", "classType", "classType_lf_classOrInterfaceType",
	"classType_lfno_classOrInterfaceType", "interfaceType", "interfaceType_lf_classOrInterfaceType",
	"interfaceType_lfno_classOrInterfaceType", "typeVariable", "arrayType",
	"dims", "typeParameter", "typeParameterModifier", "typeBound", "additionalBound",
	"typeArguments", "typeArgumentList", "typeArgument", "wildcard", "wildcardBounds",
	"packageName", "typeName", "packageOrTypeName", "expressionName", "methodName",
	"ambiguousName", "compilationUnit", "packageDeclaration", "packageModifier",
	"importDeclaration", "singleTypeImportDeclaration", "typeImportOnDemandDeclaration",
	"singleStaticImportDeclaration", "staticImportOnDemandDeclaration", "typeDeclaration",
	"classDeclaration", "normalClassDeclaration", "classModifier", "typeParameters",
	"typeParameterList", "superclass", "superinterfaces", "interfaceTypeList",
	"classBody", "classBodyDeclaration", "classMemberDeclaration", "fieldDeclaration",
	"fieldModifier", "variableDeclaratorList", "variableDeclarator", "variableDeclaratorId",
	"variableInitializer", "unannType", "unannPrimitiveType", "unannReferenceType",
	"unannClassOrInterfaceType", "unannClassType", "unannClassType_lf_unannClassOrInterfaceType",
	"unannClassType_lfno_unannClassOrInterfaceType", "unannInterfaceType",
	"unannInterfaceType_lf_unannClassOrInterfaceType", "unannInterfaceType_lfno_unannClassOrInterfaceType",
	"unannTypeVariable", "unannArrayType", "methodDeclaration", "methodModifier",
	"methodHeader", "result", "methodDeclarator", "formalParameterList", "formalParameters",
	"formalParameter", "variableModifier", "lastFormalParameter", "receiverParameter",
	"throws_", "exceptionTypeList", "exceptionType", "methodBody", "instanceInitializer",
	"staticInitializer", "constructorDeclaration", "constructorModifier", "constructorDeclarator",
	"simpleTypeName", "constructorBody", "explicitConstructorInvocation", "enumDeclaration",
	"enumBody", "enumConstantList", "enumConstant", "enumConstantModifier",
	"enumBodyDeclarations", "interfaceDeclaration", "normalInterfaceDeclaration",
	"interfaceModifier", "extendsInterfaces", "interfaceBody", "interfaceMemberDeclaration",
	"constantDeclaration", "constantModifier", "interfaceMethodDeclaration",
	"interfaceMethodModifier", "annotationTypeDeclaration", "annotationTypeBody",
	"annotationTypeMemberDeclaration", "annotationTypeElementDeclaration",
	"annotationTypeElementModifier", "defaultValue", "annotation", "normalAnnotation",
	"elementValuePairList", "elementValuePair", "elementValue", "elementValueArrayInitializer",
	"elementValueList", "markerAnnotation", "singleElementAnnotation", "arrayInitializer",
	"variableInitializerList", "block", "blockStatements", "blockStatement",
	"localVariableDeclarationStatement", "localVariableDeclaration", "statement",
	"statementNoShortIf", "statementWithoutTrailingSubstatement", "emptyStatement_",
	"labeledStatement", "labeledStatementNoShortIf", "expressionStatement",
	"statementExpression", "ifThenStatement", "ifThenElseStatement", "ifThenElseStatementNoShortIf",
	"assertStatement", "switchStatement", "switchBlock", "switchBlockStatementGroup",
	"switchLabels", "switchLabel", "enumConstantName", "whileStatement", "whileStatementNoShortIf",
	"doStatement", "forStatement", "forStatementNoShortIf", "basicForStatement",
	"basicForStatementNoShortIf", "forInit", "forUpdate", "statementExpressionList",
	"enhancedForStatement", "enhancedForStatementNoShortIf", "breakStatement",
	"continueStatement", "returnStatement", "throwStatement", "synchronizedStatement",
	"tryStatement", "catches", "catchClause", "catchFormalParameter", "catchType",
	"finally_", "tryWithResourcesStatement", "resourceSpecification", "resourceList",
	"resource", "primary", "primaryNoNewArray", "primaryNoNewArray_lf_arrayAccess",
	"primaryNoNewArray_lfno_arrayAccess", "primaryNoNewArray_lf_primary", "primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary",
	"primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary", "primaryNoNewArray_lfno_primary",
	"primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary", "primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary",
	"classInstanceCreationExpression", "classInstanceCreationExpression_lf_primary",
	"classInstanceCreationExpression_lfno_primary", "typeArgumentsOrDiamond",
	"fieldAccess", "fieldAccess_lf_primary", "fieldAccess_lfno_primary", "arrayAccess",
	"arrayAccess_lf_primary", "arrayAccess_lfno_primary", "methodInvocation",
	"methodInvocation_lf_primary", "methodInvocation_lfno_primary", "argumentList",
	"methodReference", "methodReference_lf_primary", "methodReference_lfno_primary",
	"arrayCreationExpression", "dimExprs", "dimExpr", "constantExpression",
	"expression", "lambdaExpression", "lambdaParameters", "inferredFormalParameterList",
	"lambdaBody", "assignmentExpression", "assignment", "leftHandSide", "assignmentOperator",
	"conditionalExpression", "conditionalOrExpression", "conditionalAndExpression",
	"inclusiveOrExpression", "exclusiveOrExpression", "andExpression", "equalityExpression",
	"relationalExpression", "shiftExpression", "additiveExpression", "multiplicativeExpression",
	"unaryExpression", "preIncrementExpression", "preDecrementExpression",
	"unaryExpressionNotPlusMinus", "postfixExpression", "postIncrementExpression",
	"postIncrementExpression_lf_postfixExpression", "postDecrementExpression",
	"postDecrementExpression_lf_postfixExpression", "castExpression",
}
