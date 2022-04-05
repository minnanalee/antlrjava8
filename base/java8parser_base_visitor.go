// Code generated from ./Java8Parser.g4 by ANTLR 4.9. DO NOT EDIT.

package base // Java8Parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type BaseJava8ParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJava8ParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitNumericType(ctx *NumericTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitIntegralType(ctx *IntegralTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFloatingPointType(ctx *FloatingPointTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitReferenceType(ctx *ReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassType(ctx *ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeVariable(ctx *TypeVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitDims(ctx *DimsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeBound(ctx *TypeBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAdditionalBound(ctx *AdditionalBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeArguments(ctx *TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitWildcard(ctx *WildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitWildcardBounds(ctx *WildcardBoundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPackageName(ctx *PackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPackageOrTypeName(ctx *PackageOrTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExpressionName(ctx *ExpressionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodName(ctx *MethodNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAmbiguousName(ctx *AmbiguousNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPackageModifier(ctx *PackageModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitImportDeclaration(ctx *ImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitNormalClassDeclaration(ctx *NormalClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassModifier(ctx *ClassModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeParameters(ctx *TypeParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeParameterList(ctx *TypeParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSuperclass(ctx *SuperclassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSuperinterfaces(ctx *SuperinterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceTypeList(ctx *InterfaceTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFieldModifier(ctx *FieldModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannType(ctx *UnannTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannReferenceType(ctx *UnannReferenceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannClassType(ctx *UnannClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannInterfaceType(ctx *UnannInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannTypeVariable(ctx *UnannTypeVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnannArrayType(ctx *UnannArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodModifier(ctx *MethodModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodHeader(ctx *MethodHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodDeclarator(ctx *MethodDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitReceiverParameter(ctx *ReceiverParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitThrows_(ctx *Throws_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExceptionTypeList(ctx *ExceptionTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExceptionType(ctx *ExceptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodBody(ctx *MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInstanceInitializer(ctx *InstanceInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStaticInitializer(ctx *StaticInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstructorModifier(ctx *ConstructorModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstructorDeclarator(ctx *ConstructorDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSimpleTypeName(ctx *SimpleTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstructorBody(ctx *ConstructorBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumConstantList(ctx *EnumConstantListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumConstant(ctx *EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumConstantModifier(ctx *EnumConstantModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceModifier(ctx *InterfaceModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExtendsInterfaces(ctx *ExtendsInterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceBody(ctx *InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstantModifier(ctx *ConstantModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitNormalAnnotation(ctx *NormalAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitElementValuePairList(ctx *ElementValuePairListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitElementValuePair(ctx *ElementValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitElementValue(ctx *ElementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitElementValueList(ctx *ElementValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMarkerAnnotation(ctx *MarkerAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSingleElementAnnotation(ctx *SingleElementAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLabeledStatement(ctx *LabeledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStatementExpression(ctx *StatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitIfThenStatement(ctx *IfThenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAssertStatement(ctx *AssertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSwitchBlock(ctx *SwitchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSwitchLabels(ctx *SwitchLabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSwitchLabel(ctx *SwitchLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnumConstantName(ctx *EnumConstantNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBasicForStatement(ctx *BasicForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitForUpdate(ctx *ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitThrowStatement(ctx *ThrowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitSynchronizedStatement(ctx *SynchronizedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTryStatement(ctx *TryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCatches(ctx *CatchesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCatchClause(ctx *CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCatchFormalParameter(ctx *CatchFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCatchType(ctx *CatchTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFinally_(ctx *Finally_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitResourceList(ctx *ResourceListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFieldAccess(ctx *FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayAccess(ctx *ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodInvocation(ctx *MethodInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodReference(ctx *MethodReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitArrayCreationExpression(ctx *ArrayCreationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitDimExprs(ctx *DimExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitDimExpr(ctx *DimExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConstantExpression(ctx *ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLambdaParameters(ctx *LambdaParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInferredFormalParameterList(ctx *InferredFormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLambdaBody(ctx *LambdaBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitLeftHandSide(ctx *LeftHandSideContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitShiftExpression(ctx *ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPostDecrementExpression(ctx *PostDecrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJava8ParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
