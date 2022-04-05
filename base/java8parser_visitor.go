// Code generated from ./Java8Parser.g4 by ANTLR 4.9. DO NOT EDIT.

package base // Java8Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Java8Parser.
type Java8ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Java8Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Java8Parser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#numericType.
	VisitNumericType(ctx *NumericTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#integralType.
	VisitIntegralType(ctx *IntegralTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#floatingPointType.
	VisitFloatingPointType(ctx *FloatingPointTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#referenceType.
	VisitReferenceType(ctx *ReferenceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#classType.
	VisitClassType(ctx *ClassTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#classType_lf_classOrInterfaceType.
	VisitClassType_lf_classOrInterfaceType(ctx *ClassType_lf_classOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#classType_lfno_classOrInterfaceType.
	VisitClassType_lfno_classOrInterfaceType(ctx *ClassType_lfno_classOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceType_lf_classOrInterfaceType.
	VisitInterfaceType_lf_classOrInterfaceType(ctx *InterfaceType_lf_classOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceType_lfno_classOrInterfaceType.
	VisitInterfaceType_lfno_classOrInterfaceType(ctx *InterfaceType_lfno_classOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeVariable.
	VisitTypeVariable(ctx *TypeVariableContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#dims.
	VisitDims(ctx *DimsContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeParameterModifier.
	VisitTypeParameterModifier(ctx *TypeParameterModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeBound.
	VisitTypeBound(ctx *TypeBoundContext) interface{}

	// Visit a parse tree produced by Java8Parser#additionalBound.
	VisitAdditionalBound(ctx *AdditionalBoundContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeArgumentList.
	VisitTypeArgumentList(ctx *TypeArgumentListContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by Java8Parser#wildcard.
	VisitWildcard(ctx *WildcardContext) interface{}

	// Visit a parse tree produced by Java8Parser#wildcardBounds.
	VisitWildcardBounds(ctx *WildcardBoundsContext) interface{}

	// Visit a parse tree produced by Java8Parser#packageName.
	VisitPackageName(ctx *PackageNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#packageOrTypeName.
	VisitPackageOrTypeName(ctx *PackageOrTypeNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#expressionName.
	VisitExpressionName(ctx *ExpressionNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodName.
	VisitMethodName(ctx *MethodNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#ambiguousName.
	VisitAmbiguousName(ctx *AmbiguousNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by Java8Parser#packageDeclaration.
	VisitPackageDeclaration(ctx *PackageDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#packageModifier.
	VisitPackageModifier(ctx *PackageModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#importDeclaration.
	VisitImportDeclaration(ctx *ImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#singleTypeImportDeclaration.
	VisitSingleTypeImportDeclaration(ctx *SingleTypeImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeImportOnDemandDeclaration.
	VisitTypeImportOnDemandDeclaration(ctx *TypeImportOnDemandDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#singleStaticImportDeclaration.
	VisitSingleStaticImportDeclaration(ctx *SingleStaticImportDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#staticImportOnDemandDeclaration.
	VisitStaticImportOnDemandDeclaration(ctx *StaticImportOnDemandDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#normalClassDeclaration.
	VisitNormalClassDeclaration(ctx *NormalClassDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#classModifier.
	VisitClassModifier(ctx *ClassModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeParameters.
	VisitTypeParameters(ctx *TypeParametersContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeParameterList.
	VisitTypeParameterList(ctx *TypeParameterListContext) interface{}

	// Visit a parse tree produced by Java8Parser#superclass.
	VisitSuperclass(ctx *SuperclassContext) interface{}

	// Visit a parse tree produced by Java8Parser#superinterfaces.
	VisitSuperinterfaces(ctx *SuperinterfacesContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceTypeList.
	VisitInterfaceTypeList(ctx *InterfaceTypeListContext) interface{}

	// Visit a parse tree produced by Java8Parser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#classMemberDeclaration.
	VisitClassMemberDeclaration(ctx *ClassMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#fieldModifier.
	VisitFieldModifier(ctx *FieldModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableDeclaratorList.
	VisitVariableDeclaratorList(ctx *VariableDeclaratorListContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannType.
	VisitUnannType(ctx *UnannTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannPrimitiveType.
	VisitUnannPrimitiveType(ctx *UnannPrimitiveTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannReferenceType.
	VisitUnannReferenceType(ctx *UnannReferenceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannClassOrInterfaceType.
	VisitUnannClassOrInterfaceType(ctx *UnannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannClassType.
	VisitUnannClassType(ctx *UnannClassTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannClassType_lf_unannClassOrInterfaceType.
	VisitUnannClassType_lf_unannClassOrInterfaceType(ctx *UnannClassType_lf_unannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannClassType_lfno_unannClassOrInterfaceType.
	VisitUnannClassType_lfno_unannClassOrInterfaceType(ctx *UnannClassType_lfno_unannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannInterfaceType.
	VisitUnannInterfaceType(ctx *UnannInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannInterfaceType_lf_unannClassOrInterfaceType.
	VisitUnannInterfaceType_lf_unannClassOrInterfaceType(ctx *UnannInterfaceType_lf_unannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannInterfaceType_lfno_unannClassOrInterfaceType.
	VisitUnannInterfaceType_lfno_unannClassOrInterfaceType(ctx *UnannInterfaceType_lfno_unannClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannTypeVariable.
	VisitUnannTypeVariable(ctx *UnannTypeVariableContext) interface{}

	// Visit a parse tree produced by Java8Parser#unannArrayType.
	VisitUnannArrayType(ctx *UnannArrayTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodModifier.
	VisitMethodModifier(ctx *MethodModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodHeader.
	VisitMethodHeader(ctx *MethodHeaderContext) interface{}

	// Visit a parse tree produced by Java8Parser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodDeclarator.
	VisitMethodDeclarator(ctx *MethodDeclaratorContext) interface{}

	// Visit a parse tree produced by Java8Parser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by Java8Parser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by Java8Parser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{}

	// Visit a parse tree produced by Java8Parser#receiverParameter.
	VisitReceiverParameter(ctx *ReceiverParameterContext) interface{}

	// Visit a parse tree produced by Java8Parser#throws_.
	VisitThrows_(ctx *Throws_Context) interface{}

	// Visit a parse tree produced by Java8Parser#exceptionTypeList.
	VisitExceptionTypeList(ctx *ExceptionTypeListContext) interface{}

	// Visit a parse tree produced by Java8Parser#exceptionType.
	VisitExceptionType(ctx *ExceptionTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#instanceInitializer.
	VisitInstanceInitializer(ctx *InstanceInitializerContext) interface{}

	// Visit a parse tree produced by Java8Parser#staticInitializer.
	VisitStaticInitializer(ctx *StaticInitializerContext) interface{}

	// Visit a parse tree produced by Java8Parser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#constructorModifier.
	VisitConstructorModifier(ctx *ConstructorModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#constructorDeclarator.
	VisitConstructorDeclarator(ctx *ConstructorDeclaratorContext) interface{}

	// Visit a parse tree produced by Java8Parser#simpleTypeName.
	VisitSimpleTypeName(ctx *SimpleTypeNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#constructorBody.
	VisitConstructorBody(ctx *ConstructorBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#explicitConstructorInvocation.
	VisitExplicitConstructorInvocation(ctx *ExplicitConstructorInvocationContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumConstantList.
	VisitEnumConstantList(ctx *EnumConstantListContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumConstant.
	VisitEnumConstant(ctx *EnumConstantContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumConstantModifier.
	VisitEnumConstantModifier(ctx *EnumConstantModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumBodyDeclarations.
	VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#normalInterfaceDeclaration.
	VisitNormalInterfaceDeclaration(ctx *NormalInterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceModifier.
	VisitInterfaceModifier(ctx *InterfaceModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#extendsInterfaces.
	VisitExtendsInterfaces(ctx *ExtendsInterfacesContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceMemberDeclaration.
	VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#constantDeclaration.
	VisitConstantDeclaration(ctx *ConstantDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#constantModifier.
	VisitConstantModifier(ctx *ConstantModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#interfaceMethodModifier.
	VisitInterfaceMethodModifier(ctx *InterfaceMethodModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotationTypeDeclaration.
	VisitAnnotationTypeDeclaration(ctx *AnnotationTypeDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotationTypeBody.
	VisitAnnotationTypeBody(ctx *AnnotationTypeBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotationTypeMemberDeclaration.
	VisitAnnotationTypeMemberDeclaration(ctx *AnnotationTypeMemberDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotationTypeElementDeclaration.
	VisitAnnotationTypeElementDeclaration(ctx *AnnotationTypeElementDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotationTypeElementModifier.
	VisitAnnotationTypeElementModifier(ctx *AnnotationTypeElementModifierContext) interface{}

	// Visit a parse tree produced by Java8Parser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by Java8Parser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by Java8Parser#normalAnnotation.
	VisitNormalAnnotation(ctx *NormalAnnotationContext) interface{}

	// Visit a parse tree produced by Java8Parser#elementValuePairList.
	VisitElementValuePairList(ctx *ElementValuePairListContext) interface{}

	// Visit a parse tree produced by Java8Parser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) interface{}

	// Visit a parse tree produced by Java8Parser#elementValue.
	VisitElementValue(ctx *ElementValueContext) interface{}

	// Visit a parse tree produced by Java8Parser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{}

	// Visit a parse tree produced by Java8Parser#elementValueList.
	VisitElementValueList(ctx *ElementValueListContext) interface{}

	// Visit a parse tree produced by Java8Parser#markerAnnotation.
	VisitMarkerAnnotation(ctx *MarkerAnnotationContext) interface{}

	// Visit a parse tree produced by Java8Parser#singleElementAnnotation.
	VisitSingleElementAnnotation(ctx *SingleElementAnnotationContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by Java8Parser#variableInitializerList.
	VisitVariableInitializerList(ctx *VariableInitializerListContext) interface{}

	// Visit a parse tree produced by Java8Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Java8Parser#blockStatements.
	VisitBlockStatements(ctx *BlockStatementsContext) interface{}

	// Visit a parse tree produced by Java8Parser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by Java8Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#statementNoShortIf.
	VisitStatementNoShortIf(ctx *StatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#statementWithoutTrailingSubstatement.
	VisitStatementWithoutTrailingSubstatement(ctx *StatementWithoutTrailingSubstatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by Java8Parser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#labeledStatementNoShortIf.
	VisitLabeledStatementNoShortIf(ctx *LabeledStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#ifThenStatement.
	VisitIfThenStatement(ctx *IfThenStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#ifThenElseStatement.
	VisitIfThenElseStatement(ctx *IfThenElseStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#ifThenElseStatementNoShortIf.
	VisitIfThenElseStatementNoShortIf(ctx *IfThenElseStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#assertStatement.
	VisitAssertStatement(ctx *AssertStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#switchBlock.
	VisitSwitchBlock(ctx *SwitchBlockContext) interface{}

	// Visit a parse tree produced by Java8Parser#switchBlockStatementGroup.
	VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{}

	// Visit a parse tree produced by Java8Parser#switchLabels.
	VisitSwitchLabels(ctx *SwitchLabelsContext) interface{}

	// Visit a parse tree produced by Java8Parser#switchLabel.
	VisitSwitchLabel(ctx *SwitchLabelContext) interface{}

	// Visit a parse tree produced by Java8Parser#enumConstantName.
	VisitEnumConstantName(ctx *EnumConstantNameContext) interface{}

	// Visit a parse tree produced by Java8Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#whileStatementNoShortIf.
	VisitWhileStatementNoShortIf(ctx *WhileStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#forStatementNoShortIf.
	VisitForStatementNoShortIf(ctx *ForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#basicForStatement.
	VisitBasicForStatement(ctx *BasicForStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#basicForStatementNoShortIf.
	VisitBasicForStatementNoShortIf(ctx *BasicForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by Java8Parser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by Java8Parser#statementExpressionList.
	VisitStatementExpressionList(ctx *StatementExpressionListContext) interface{}

	// Visit a parse tree produced by Java8Parser#enhancedForStatement.
	VisitEnhancedForStatement(ctx *EnhancedForStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#enhancedForStatementNoShortIf.
	VisitEnhancedForStatementNoShortIf(ctx *EnhancedForStatementNoShortIfContext) interface{}

	// Visit a parse tree produced by Java8Parser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#throwStatement.
	VisitThrowStatement(ctx *ThrowStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#synchronizedStatement.
	VisitSynchronizedStatement(ctx *SynchronizedStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#tryStatement.
	VisitTryStatement(ctx *TryStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#catches.
	VisitCatches(ctx *CatchesContext) interface{}

	// Visit a parse tree produced by Java8Parser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by Java8Parser#catchFormalParameter.
	VisitCatchFormalParameter(ctx *CatchFormalParameterContext) interface{}

	// Visit a parse tree produced by Java8Parser#catchType.
	VisitCatchType(ctx *CatchTypeContext) interface{}

	// Visit a parse tree produced by Java8Parser#finally_.
	VisitFinally_(ctx *Finally_Context) interface{}

	// Visit a parse tree produced by Java8Parser#tryWithResourcesStatement.
	VisitTryWithResourcesStatement(ctx *TryWithResourcesStatementContext) interface{}

	// Visit a parse tree produced by Java8Parser#resourceSpecification.
	VisitResourceSpecification(ctx *ResourceSpecificationContext) interface{}

	// Visit a parse tree produced by Java8Parser#resourceList.
	VisitResourceList(ctx *ResourceListContext) interface{}

	// Visit a parse tree produced by Java8Parser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by Java8Parser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray.
	VisitPrimaryNoNewArray(ctx *PrimaryNoNewArrayContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lf_arrayAccess.
	VisitPrimaryNoNewArray_lf_arrayAccess(ctx *PrimaryNoNewArray_lf_arrayAccessContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lfno_arrayAccess.
	VisitPrimaryNoNewArray_lfno_arrayAccess(ctx *PrimaryNoNewArray_lfno_arrayAccessContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lf_primary.
	VisitPrimaryNoNewArray_lf_primary(ctx *PrimaryNoNewArray_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary.
	VisitPrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lf_arrayAccess_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary.
	VisitPrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primary(ctx *PrimaryNoNewArray_lf_primary_lfno_arrayAccess_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lfno_primary.
	VisitPrimaryNoNewArray_lfno_primary(ctx *PrimaryNoNewArray_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary.
	VisitPrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lf_arrayAccess_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#primaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary.
	VisitPrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primary(ctx *PrimaryNoNewArray_lfno_primary_lfno_arrayAccess_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#classInstanceCreationExpression.
	VisitClassInstanceCreationExpression(ctx *ClassInstanceCreationExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#classInstanceCreationExpression_lf_primary.
	VisitClassInstanceCreationExpression_lf_primary(ctx *ClassInstanceCreationExpression_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#classInstanceCreationExpression_lfno_primary.
	VisitClassInstanceCreationExpression_lfno_primary(ctx *ClassInstanceCreationExpression_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#typeArgumentsOrDiamond.
	VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by Java8Parser#fieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by Java8Parser#fieldAccess_lf_primary.
	VisitFieldAccess_lf_primary(ctx *FieldAccess_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#fieldAccess_lfno_primary.
	VisitFieldAccess_lfno_primary(ctx *FieldAccess_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayAccess_lf_primary.
	VisitArrayAccess_lf_primary(ctx *ArrayAccess_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayAccess_lfno_primary.
	VisitArrayAccess_lfno_primary(ctx *ArrayAccess_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodInvocation.
	VisitMethodInvocation(ctx *MethodInvocationContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodInvocation_lf_primary.
	VisitMethodInvocation_lf_primary(ctx *MethodInvocation_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodInvocation_lfno_primary.
	VisitMethodInvocation_lfno_primary(ctx *MethodInvocation_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodReference.
	VisitMethodReference(ctx *MethodReferenceContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodReference_lf_primary.
	VisitMethodReference_lf_primary(ctx *MethodReference_lf_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#methodReference_lfno_primary.
	VisitMethodReference_lfno_primary(ctx *MethodReference_lfno_primaryContext) interface{}

	// Visit a parse tree produced by Java8Parser#arrayCreationExpression.
	VisitArrayCreationExpression(ctx *ArrayCreationExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#dimExprs.
	VisitDimExprs(ctx *DimExprsContext) interface{}

	// Visit a parse tree produced by Java8Parser#dimExpr.
	VisitDimExpr(ctx *DimExprContext) interface{}

	// Visit a parse tree produced by Java8Parser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#lambdaParameters.
	VisitLambdaParameters(ctx *LambdaParametersContext) interface{}

	// Visit a parse tree produced by Java8Parser#inferredFormalParameterList.
	VisitInferredFormalParameterList(ctx *InferredFormalParameterListContext) interface{}

	// Visit a parse tree produced by Java8Parser#lambdaBody.
	VisitLambdaBody(ctx *LambdaBodyContext) interface{}

	// Visit a parse tree produced by Java8Parser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by Java8Parser#leftHandSide.
	VisitLeftHandSide(ctx *LeftHandSideContext) interface{}

	// Visit a parse tree produced by Java8Parser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by Java8Parser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#conditionalOrExpression.
	VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#conditionalAndExpression.
	VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#preIncrementExpression.
	VisitPreIncrementExpression(ctx *PreIncrementExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#preDecrementExpression.
	VisitPreDecrementExpression(ctx *PreDecrementExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#unaryExpressionNotPlusMinus.
	VisitUnaryExpressionNotPlusMinus(ctx *UnaryExpressionNotPlusMinusContext) interface{}

	// Visit a parse tree produced by Java8Parser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#postIncrementExpression.
	VisitPostIncrementExpression(ctx *PostIncrementExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#postIncrementExpression_lf_postfixExpression.
	VisitPostIncrementExpression_lf_postfixExpression(ctx *PostIncrementExpression_lf_postfixExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#postDecrementExpression.
	VisitPostDecrementExpression(ctx *PostDecrementExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#postDecrementExpression_lf_postfixExpression.
	VisitPostDecrementExpression_lf_postfixExpression(ctx *PostDecrementExpression_lf_postfixExpressionContext) interface{}

	// Visit a parse tree produced by Java8Parser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}
}
