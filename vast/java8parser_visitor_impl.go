package vast

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/minnanalee/antlrjava8/base"
	"reflect"
)

type ClassType struct {
	Annotation []string
	ClassName  string
	Methods    []string
}

type CallGraphLog struct {
	Package string
	Imports []string
	Classes []struct {
		ClassName       string
		Class_SQN       string
		ClassAnnotation string
		Methods         []struct {
			MethodName string
			Method_SQN string
		}
	}
}

type Method struct {
}

type SourceInfo struct {
	Start, End, Line, Column int
	Source                   string
}

// 解析的源代码的起止位置，行、列号，内容
func getSourceInfo(ctx antlr.BaseParserRuleContext) *SourceInfo {
	start, end := ctx.GetStart().GetStart(), ctx.GetStop().GetStop()
	return &SourceInfo{Line: ctx.GetStart().GetLine(), Start: start, End: end,
		Column: ctx.GetStart().GetColumn(),
		Source: ctx.GetStart().GetInputStream().GetTextFromInterval(&antlr.Interval{
			Start: start, Stop: end})}

}

type Visitor struct {
	antlr.BaseParseTreeVisitor
	ClassMap map[string]ClassType
}

func NewVisitor() *Visitor {
	classMap := make(map[string]ClassType)
	return &Visitor{
		BaseParseTreeVisitor: antlr.BaseParseTreeVisitor{},
		ClassMap:             classMap,
	}
}

func (c *Visitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Println("Visit->Tree type: ", reflect.TypeOf(tree))

	switch val := tree.(type) {
	default:
		return val.Accept(c)
		/*	case *CalculateContext:
				return val.Accept(c)
			case *ToSetVarContext:
				return val.Accept(c)
		*/
	}
	return nil
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, ch := range node.GetChildren() {
		switch rr := ch.(type) {
		case *antlr.TerminalNodeImpl:
			fmt.Println("Terminal Node Value: \t" + rr.GetText())
			/*		default:
					fmt.Println("Middle Node Value: \t"+ch.(antlr.ParseTree).ToStringTree(nil,antlr.Parser)
			*/
		}
		ch.(antlr.ParseTree).Accept(v)
	}
	return nil
}
func (v *Visitor) VisitNormalClassDeclaration(ctx *base.NormalClassDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	jSourceInfo, _ := json.Marshal(getSourceInfo(*ctx.BaseParserRuleContext))
	fmt.Println("RuleSourceInfo: ", string(jSourceInfo))
	for i := 0; i < ctx.GetChildCount(); i++ {
		fmt.Println("normalClassDeclaration child ", i)
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), ctx.GetChild(i))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLiteral(ctx *base.LiteralContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimitiveType(ctx *base.PrimitiveTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNumericType(ctx *base.NumericTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIntegralType(ctx *base.IntegralTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFloatingPointType(ctx *base.FloatingPointTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReferenceType(ctx *base.ReferenceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassOrInterfaceType(ctx *base.ClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassType(ctx *base.ClassTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassType_lf_classOrInterfaceType(ctx *base.ClassType_lf_classOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassType_lfno_classOrInterfaceType(ctx *base.ClassType_lfno_classOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceType(ctx *base.InterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceType_lf_classOrInterfaceType(ctx *base.InterfaceType_lf_classOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceType_lfno_classOrInterfaceType(ctx *base.InterfaceType_lfno_classOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeVariable(ctx *base.TypeVariableContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayType(ctx *base.ArrayTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDims(ctx *base.DimsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeParameter(ctx *base.TypeParameterContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeParameterModifier(ctx *base.TypeParameterModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeBound(ctx *base.TypeBoundContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAdditionalBound(ctx *base.AdditionalBoundContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeArguments(ctx *base.TypeArgumentsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeArgumentList(ctx *base.TypeArgumentListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeArgument(ctx *base.TypeArgumentContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWildcard(ctx *base.WildcardContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWildcardBounds(ctx *base.WildcardBoundsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPackageName(ctx *base.PackageNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeName(ctx *base.TypeNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPackageOrTypeName(ctx *base.PackageOrTypeNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpressionName(ctx *base.ExpressionNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodName(ctx *base.MethodNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAmbiguousName(ctx *base.AmbiguousNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCompilationUnit(ctx *base.CompilationUnitContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPackageDeclaration(ctx *base.PackageDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPackageModifier(ctx *base.PackageModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitImportDeclaration(ctx *base.ImportDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSingleTypeImportDeclaration(ctx *base.SingleTypeImportDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeImportOnDemandDeclaration(ctx *base.TypeImportOnDemandDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSingleStaticImportDeclaration(ctx *base.SingleStaticImportDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStaticImportOnDemandDeclaration(ctx *base.StaticImportOnDemandDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeDeclaration(ctx *base.TypeDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassDeclaration(ctx *base.ClassDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassModifier(ctx *base.ClassModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeParameters(ctx *base.TypeParametersContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeParameterList(ctx *base.TypeParameterListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSuperclass(ctx *base.SuperclassContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSuperinterfaces(ctx *base.SuperinterfacesContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceTypeList(ctx *base.InterfaceTypeListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassBody(ctx *base.ClassBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassBodyDeclaration(ctx *base.ClassBodyDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassMemberDeclaration(ctx *base.ClassMemberDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldDeclaration(ctx *base.FieldDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldModifier(ctx *base.FieldModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableDeclaratorList(ctx *base.VariableDeclaratorListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableDeclarator(ctx *base.VariableDeclaratorContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableDeclaratorId(ctx *base.VariableDeclaratorIdContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableInitializer(ctx *base.VariableInitializerContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannType(ctx *base.UnannTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannPrimitiveType(ctx *base.UnannPrimitiveTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannReferenceType(ctx *base.UnannReferenceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannClassOrInterfaceType(ctx *base.UnannClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannClassType(ctx *base.UnannClassTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannClassType_lf_unannClassOrInterfaceType(ctx *base.UnannClassType_lf_unannClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannClassType_lfno_unannClassOrInterfaceType(ctx *base.UnannClassType_lfno_unannClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannInterfaceType(ctx *base.UnannInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *base.UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *base.UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannTypeVariable(ctx *base.UnannTypeVariableContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnannArrayType(ctx *base.UnannArrayTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodDeclaration(ctx *base.MethodDeclarationContext) interface{} {

	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))

	jSourceInfo, _ := json.Marshal(getSourceInfo(*ctx.BaseParserRuleContext))
	fmt.Println("RuleSourceInfo: ", string(jSourceInfo))
	for i := 0; i < ctx.GetChildCount(); i++ {
		fmt.Println("normalClassDeclaration child ", i)
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), ctx.GetChild(i))
	}

	return v.VisitChildren(ctx)

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodModifier(ctx *base.MethodModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodHeader(ctx *base.MethodHeaderContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResult(ctx *base.ResultContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodDeclarator(ctx *base.MethodDeclaratorContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFormalParameterList(ctx *base.FormalParameterListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFormalParameters(ctx *base.FormalParametersContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFormalParameter(ctx *base.FormalParameterContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableModifier(ctx *base.VariableModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLastFormalParameter(ctx *base.LastFormalParameterContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReceiverParameter(ctx *base.ReceiverParameterContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitThrows_(ctx *base.Throws_Context) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExceptionTypeList(ctx *base.ExceptionTypeListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExceptionType(ctx *base.ExceptionTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodBody(ctx *base.MethodBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInstanceInitializer(ctx *base.InstanceInitializerContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStaticInitializer(ctx *base.StaticInitializerContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstructorDeclaration(ctx *base.ConstructorDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstructorModifier(ctx *base.ConstructorModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstructorDeclarator(ctx *base.ConstructorDeclaratorContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSimpleTypeName(ctx *base.SimpleTypeNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstructorBody(ctx *base.ConstructorBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExplicitConstructorInvocation(ctx *base.ExplicitConstructorInvocationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumDeclaration(ctx *base.EnumDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumBody(ctx *base.EnumBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumConstantList(ctx *base.EnumConstantListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumConstant(ctx *base.EnumConstantContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumConstantModifier(ctx *base.EnumConstantModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumBodyDeclarations(ctx *base.EnumBodyDeclarationsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceDeclaration(ctx *base.InterfaceDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNormalInterfaceDeclaration(ctx *base.NormalInterfaceDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceModifier(ctx *base.InterfaceModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExtendsInterfaces(ctx *base.ExtendsInterfacesContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceBody(ctx *base.InterfaceBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceMemberDeclaration(ctx *base.InterfaceMemberDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstantDeclaration(ctx *base.ConstantDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstantModifier(ctx *base.ConstantModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceMethodDeclaration(ctx *base.InterfaceMethodDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInterfaceMethodModifier(ctx *base.InterfaceMethodModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotationTypeDeclaration(ctx *base.AnnotationTypeDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotationTypeBody(ctx *base.AnnotationTypeBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotationTypeMemberDeclaration(ctx *base.AnnotationTypeMemberDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotationTypeElementDeclaration(ctx *base.AnnotationTypeElementDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotationTypeElementModifier(ctx *base.AnnotationTypeElementModifierContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDefaultValue(ctx *base.DefaultValueContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAnnotation(ctx *base.AnnotationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitNormalAnnotation(ctx *base.NormalAnnotationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementValuePairList(ctx *base.ElementValuePairListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementValuePair(ctx *base.ElementValuePairContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementValue(ctx *base.ElementValueContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementValueArrayInitializer(ctx *base.ElementValueArrayInitializerContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitElementValueList(ctx *base.ElementValueListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMarkerAnnotation(ctx *base.MarkerAnnotationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSingleElementAnnotation(ctx *base.SingleElementAnnotationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayInitializer(ctx *base.ArrayInitializerContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitVariableInitializerList(ctx *base.VariableInitializerListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBlock(ctx *base.BlockContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBlockStatements(ctx *base.BlockStatementsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBlockStatement(ctx *base.BlockStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLocalVariableDeclarationStatement(ctx *base.LocalVariableDeclarationStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLocalVariableDeclaration(ctx *base.LocalVariableDeclarationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatement(ctx *base.StatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatementNoShortIf(ctx *base.StatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatementWithoutTrailingSubstatement(ctx *base.StatementWithoutTrailingSubstatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEmptyStatement_(ctx *base.EmptyStatement_Context) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLabeledStatement(ctx *base.LabeledStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLabeledStatementNoShortIf(ctx *base.LabeledStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpressionStatement(ctx *base.ExpressionStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatementExpression(ctx *base.StatementExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfThenStatement(ctx *base.IfThenStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfThenElseStatement(ctx *base.IfThenElseStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitIfThenElseStatementNoShortIf(ctx *base.IfThenElseStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssertStatement(ctx *base.AssertStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchStatement(ctx *base.SwitchStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchBlock(ctx *base.SwitchBlockContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchBlockStatementGroup(ctx *base.SwitchBlockStatementGroupContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchLabels(ctx *base.SwitchLabelsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSwitchLabel(ctx *base.SwitchLabelContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnumConstantName(ctx *base.EnumConstantNameContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWhileStatement(ctx *base.WhileStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWhileStatementNoShortIf(ctx *base.WhileStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDoStatement(ctx *base.DoStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForStatement(ctx *base.ForStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForStatementNoShortIf(ctx *base.ForStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBasicForStatement(ctx *base.BasicForStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBasicForStatementNoShortIf(ctx *base.BasicForStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForInit(ctx *base.ForInitContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitForUpdate(ctx *base.ForUpdateContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitStatementExpressionList(ctx *base.StatementExpressionListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnhancedForStatement(ctx *base.EnhancedForStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEnhancedForStatementNoShortIf(ctx *base.EnhancedForStatementNoShortIfContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitBreakStatement(ctx *base.BreakStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitContinueStatement(ctx *base.ContinueStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitReturnStatement(ctx *base.ReturnStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitThrowStatement(ctx *base.ThrowStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitSynchronizedStatement(ctx *base.SynchronizedStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTryStatement(ctx *base.TryStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCatches(ctx *base.CatchesContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCatchClause(ctx *base.CatchClauseContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCatchFormalParameter(ctx *base.CatchFormalParameterContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCatchType(ctx *base.CatchTypeContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFinally_(ctx *base.Finally_Context) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTryWithResourcesStatement(ctx *base.TryWithResourcesStatementContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResourceSpecification(ctx *base.ResourceSpecificationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResourceList(ctx *base.ResourceListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitResource(ctx *base.ResourceContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimary(ctx *base.PrimaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray(ctx *base.PrimaryNoNewArrayContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lf_arrayAccess(ctx *base.PrimaryNoNewArray_lf_arrayAccessContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lfno_arrayAccess(ctx *base.PrimaryNoNewArray_lfno_arrayAccessContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lf_primary(ctx *base.PrimaryNoNewArray_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *base.PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *base.PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lfno_primary(ctx *base.PrimaryNoNewArray_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *base.PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *base.PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassInstanceCreationExpression(ctx *base.ClassInstanceCreationExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassInstanceCreationExpression_lf_primary(ctx *base.ClassInstanceCreationExpression_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitClassInstanceCreationExpression_lfno_primary(ctx *base.ClassInstanceCreationExpression_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitTypeArgumentsOrDiamond(ctx *base.TypeArgumentsOrDiamondContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldAccess(ctx *base.FieldAccessContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldAccess_lf_primary(ctx *base.FieldAccess_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitFieldAccess_lfno_primary(ctx *base.FieldAccess_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayAccess(ctx *base.ArrayAccessContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayAccess_lf_primary(ctx *base.ArrayAccess_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayAccess_lfno_primary(ctx *base.ArrayAccess_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodInvocation(ctx *base.MethodInvocationContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodInvocation_lf_primary(ctx *base.MethodInvocation_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodInvocation_lfno_primary(ctx *base.MethodInvocation_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArgumentList(ctx *base.ArgumentListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodReference(ctx *base.MethodReferenceContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodReference_lf_primary(ctx *base.MethodReference_lf_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMethodReference_lfno_primary(ctx *base.MethodReference_lfno_primaryContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitArrayCreationExpression(ctx *base.ArrayCreationExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDimExprs(ctx *base.DimExprsContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitDimExpr(ctx *base.DimExprContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConstantExpression(ctx *base.ConstantExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExpression(ctx *base.ExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLambdaExpression(ctx *base.LambdaExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLambdaParameters(ctx *base.LambdaParametersContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInferredFormalParameterList(ctx *base.InferredFormalParameterListContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLambdaBody(ctx *base.LambdaBodyContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssignmentExpression(ctx *base.AssignmentExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssignment(ctx *base.AssignmentContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitLeftHandSide(ctx *base.LeftHandSideContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAssignmentOperator(ctx *base.AssignmentOperatorContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConditionalExpression(ctx *base.ConditionalExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConditionalOrExpression(ctx *base.ConditionalOrExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitConditionalAndExpression(ctx *base.ConditionalAndExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitInclusiveOrExpression(ctx *base.InclusiveOrExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitExclusiveOrExpression(ctx *base.ExclusiveOrExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAndExpression(ctx *base.AndExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitEqualityExpression(ctx *base.EqualityExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitRelationalExpression(ctx *base.RelationalExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitShiftExpression(ctx *base.ShiftExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitAdditiveExpression(ctx *base.AdditiveExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *base.MultiplicativeExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnaryExpression(ctx *base.UnaryExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPreIncrementExpression(ctx *base.PreIncrementExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPreDecrementExpression(ctx *base.PreDecrementExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitUnaryExpressionNotPlusMinus(ctx *base.UnaryExpressionNotPlusMinusContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostfixExpression(ctx *base.PostfixExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostIncrementExpression(ctx *base.PostIncrementExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostIncrementExpression_lf_postfixExpression(ctx *base.PostIncrementExpression_lf_postfixExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostDecrementExpression(ctx *base.PostDecrementExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitPostDecrementExpression_lf_postfixExpression(ctx *base.PostDecrementExpression_lf_postfixExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitCastExpression(ctx *base.CastExpressionContext) interface{} {
	fmt.Println("RuleName: " + RuleNames[ctx.GetRuleIndex()] + "\t->\t Text: " + ctx.GetText() + "\t->\tType:" + fmt.Sprint(reflect.TypeOf(ctx)))
	return v.VisitChildren(ctx)
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
