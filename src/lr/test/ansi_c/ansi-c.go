// Auto-generated from source: ansi-c.lr

package ansi_c

import (
	fmt "fmt"
	io "io"
)

type CSymbolId int

const (
	CIdentifierToken    = CSymbolId(256)
	CConstantToken      = CSymbolId(257)
	CStringLiteralToken = CSymbolId(258)
	CSizeofToken        = CSymbolId(259)
	CPtrOpToken         = CSymbolId(260)
	CIncOpToken         = CSymbolId(261)
	CDecOpToken         = CSymbolId(262)
	CLeftOpToken        = CSymbolId(263)
	CRightOpToken       = CSymbolId(264)
	CLeOpToken          = CSymbolId(265)
	CGeOpToken          = CSymbolId(266)
	CEqOpToken          = CSymbolId(267)
	CNeOpToken          = CSymbolId(268)
	CAndOpToken         = CSymbolId(269)
	COrOpToken          = CSymbolId(270)
	CMulAssignToken     = CSymbolId(271)
	CDivAssignToken     = CSymbolId(272)
	CModAssignToken     = CSymbolId(273)
	CAddAssignToken     = CSymbolId(274)
	CSubAssignToken     = CSymbolId(275)
	CLeftAssignToken    = CSymbolId(276)
	CRightAssignToken   = CSymbolId(277)
	CAndAssignToken     = CSymbolId(278)
	CXorAssignToken     = CSymbolId(279)
	COrAssignToken      = CSymbolId(280)
	CTypeNameToken      = CSymbolId(281)
	CTypedefToken       = CSymbolId(282)
	CExternToken        = CSymbolId(283)
	CStaticToken        = CSymbolId(284)
	CAutoToken          = CSymbolId(285)
	CRegisterToken      = CSymbolId(286)
	CCharToken          = CSymbolId(287)
	CShortToken         = CSymbolId(288)
	CIntToken           = CSymbolId(289)
	CLongToken          = CSymbolId(290)
	CSignedToken        = CSymbolId(291)
	CUnsignedToken      = CSymbolId(292)
	CFloatToken         = CSymbolId(293)
	CDoubleToken        = CSymbolId(294)
	CConstToken         = CSymbolId(295)
	CVolatileToken      = CSymbolId(296)
	CVoidToken          = CSymbolId(297)
	CStructToken        = CSymbolId(298)
	CUnionToken         = CSymbolId(299)
	CEnumToken          = CSymbolId(300)
	CEllipsisToken      = CSymbolId(301)
	CCaseToken          = CSymbolId(302)
	CDefaultToken       = CSymbolId(303)
	CIfToken            = CSymbolId(304)
	CElseToken          = CSymbolId(305)
	CSwitchToken        = CSymbolId(306)
	CWhileToken         = CSymbolId(307)
	CDoToken            = CSymbolId(308)
	CForToken           = CSymbolId(309)
	CGotoToken          = CSymbolId(310)
	CContinueToken      = CSymbolId(311)
	CBreakToken         = CSymbolId(312)
	CReturnToken        = CSymbolId(313)
	CLparamToken        = CSymbolId(314)
	CRparamToken        = CSymbolId(315)
	CLcurlToken         = CSymbolId(316)
	CRcurlToken         = CSymbolId(317)
	CLbraceToken        = CSymbolId(318)
	CRbraceToken        = CSymbolId(319)
	CSemicolonToken     = CSymbolId(320)
	CColonToken         = CSymbolId(321)
	CCommaToken         = CSymbolId(322)
	CEqToken            = CSymbolId(323)
	CQuestionToken      = CSymbolId(324)
	CMulToken           = CSymbolId(325)
	CDivToken           = CSymbolId(326)
	CMinusToken         = CSymbolId(327)
	CPlusToken          = CSymbolId(328)
	CModToken           = CSymbolId(329)
	CAndToken           = CSymbolId(330)
	COrToken            = CSymbolId(331)
	CExclaimToken       = CSymbolId(332)
	CDotToken           = CSymbolId(333)
	CHatToken           = CSymbolId(334)
	CLtToken            = CSymbolId(335)
	CGtToken            = CSymbolId(336)
	CTildaToken         = CSymbolId(337)
)

type CLocation struct {
	FileName string
	Line     int
	Column   int
}

func (l CLocation) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l CLocation) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type CToken interface {
	Id() CSymbolId
	Loc() CLocation
}

type CGenericSymbol struct {
	CSymbolId
	CLocation
}

func (t *CGenericSymbol) Id() CSymbolId { return t.CSymbolId }

func (t *CGenericSymbol) Loc() CLocation { return t.CLocation }

type CLexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *CGenericSymbol
	Next() (CToken, error)

	CurrentLocation() CLocation
}

type CReducer interface {
	// 83:4: primary_expression -> a: ...
	AToPrimaryExpression(Identifier_ Symbol) (Symbol, error)

	// 84:4: primary_expression -> b: ...
	BToPrimaryExpression(Constant_ Symbol) (Symbol, error)

	// 85:4: primary_expression -> c: ...
	CToPrimaryExpression(StringLiteral_ Symbol) (Symbol, error)

	// 86:4: primary_expression -> d: ...
	DToPrimaryExpression(Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 89:4: postfix_expression -> a: ...
	AToPostfixExpression(PrimaryExpression_ Symbol) (Symbol, error)

	// 90:4: postfix_expression -> b: ...
	BToPostfixExpression(PostfixExpression_ Symbol, Lbrace_ Symbol, Expression_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 91:4: postfix_expression -> c: ...
	CToPostfixExpression(PostfixExpression_ Symbol, Lparam_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 92:4: postfix_expression -> d: ...
	DToPostfixExpression(PostfixExpression_ Symbol, Lparam_ Symbol, ArgumentExpressionList_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 93:4: postfix_expression -> e: ...
	EToPostfixExpression(PostfixExpression_ Symbol, Dot_ Symbol, Identifier_ Symbol) (Symbol, error)

	// 94:4: postfix_expression -> f: ...
	FToPostfixExpression(PostfixExpression_ Symbol, PtrOp_ Symbol, Identifier_ Symbol) (Symbol, error)

	// 95:4: postfix_expression -> g: ...
	GToPostfixExpression(PostfixExpression_ Symbol, IncOp_ Symbol) (Symbol, error)

	// 96:4: postfix_expression -> h: ...
	HToPostfixExpression(PostfixExpression_ Symbol, DecOp_ Symbol) (Symbol, error)

	// 99:4: argument_expression_list -> a: ...
	AToArgumentExpressionList(AssignmentExpression_ Symbol) (Symbol, error)

	// 100:4: argument_expression_list -> b: ...
	BToArgumentExpressionList(ArgumentExpressionList_ Symbol, Comma_ Symbol, AssignmentExpression_ Symbol) (Symbol, error)

	// 103:4: unary_expression -> a: ...
	AToUnaryExpression(PostfixExpression_ Symbol) (Symbol, error)

	// 104:4: unary_expression -> b: ...
	BToUnaryExpression(IncOp_ Symbol, UnaryExpression_ Symbol) (Symbol, error)

	// 105:4: unary_expression -> c: ...
	CToUnaryExpression(DecOp_ Symbol, UnaryExpression_ Symbol) (Symbol, error)

	// 106:4: unary_expression -> d: ...
	DToUnaryExpression(UnaryOperator_ Symbol, CastExpression_ Symbol) (Symbol, error)

	// 107:4: unary_expression -> e: ...
	EToUnaryExpression(Sizeof_ Symbol, UnaryExpression_ Symbol) (Symbol, error)

	// 108:4: unary_expression -> f: ...
	FToUnaryExpression(Sizeof_ Symbol, Lparam_ Symbol, TypeName_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 111:4: unary_operator -> a: ...
	AToUnaryOperator(And_ Symbol) (Symbol, error)

	// 112:4: unary_operator -> b: ...
	BToUnaryOperator(Mul_ Symbol) (Symbol, error)

	// 113:4: unary_operator -> c: ...
	CToUnaryOperator(Plus_ Symbol) (Symbol, error)

	// 114:4: unary_operator -> d: ...
	DToUnaryOperator(Minus_ Symbol) (Symbol, error)

	// 115:4: unary_operator -> e: ...
	EToUnaryOperator(Tilda_ Symbol) (Symbol, error)

	// 116:4: unary_operator -> f: ...
	FToUnaryOperator(Exclaim_ Symbol) (Symbol, error)

	// 119:4: cast_expression -> a: ...
	AToCastExpression(UnaryExpression_ Symbol) (Symbol, error)

	// 120:4: cast_expression -> b: ...
	BToCastExpression(Lparam_ Symbol, TypeName_ Symbol, Rparam_ Symbol, CastExpression_ Symbol) (Symbol, error)

	// 123:4: multiplicative_expression -> a: ...
	AToMultiplicativeExpression(CastExpression_ Symbol) (Symbol, error)

	// 124:4: multiplicative_expression -> b: ...
	BToMultiplicativeExpression(MultiplicativeExpression_ Symbol, Mul_ Symbol, CastExpression_ Symbol) (Symbol, error)

	// 125:4: multiplicative_expression -> c: ...
	CToMultiplicativeExpression(MultiplicativeExpression_ Symbol, Div_ Symbol, CastExpression_ Symbol) (Symbol, error)

	// 126:4: multiplicative_expression -> d: ...
	DToMultiplicativeExpression(MultiplicativeExpression_ Symbol, Mod_ Symbol, CastExpression_ Symbol) (Symbol, error)

	// 129:4: additive_expression -> a: ...
	AToAdditiveExpression(MultiplicativeExpression_ Symbol) (Symbol, error)

	// 130:4: additive_expression -> b: ...
	BToAdditiveExpression(AdditiveExpression_ Symbol, Plus_ Symbol, MultiplicativeExpression_ Symbol) (Symbol, error)

	// 131:4: additive_expression -> c: ...
	CToAdditiveExpression(AdditiveExpression_ Symbol, Minus_ Symbol, MultiplicativeExpression_ Symbol) (Symbol, error)

	// 134:4: shift_expression -> a: ...
	AToShiftExpression(AdditiveExpression_ Symbol) (Symbol, error)

	// 135:4: shift_expression -> b: ...
	BToShiftExpression(ShiftExpression_ Symbol, LeftOp_ Symbol, AdditiveExpression_ Symbol) (Symbol, error)

	// 136:4: shift_expression -> c: ...
	CToShiftExpression(ShiftExpression_ Symbol, RightOp_ Symbol, AdditiveExpression_ Symbol) (Symbol, error)

	// 139:4: relational_expression -> a: ...
	AToRelationalExpression(ShiftExpression_ Symbol) (Symbol, error)

	// 140:4: relational_expression -> b: ...
	BToRelationalExpression(RelationalExpression_ Symbol, Lt_ Symbol, ShiftExpression_ Symbol) (Symbol, error)

	// 141:4: relational_expression -> c: ...
	CToRelationalExpression(RelationalExpression_ Symbol, Gt_ Symbol, ShiftExpression_ Symbol) (Symbol, error)

	// 142:4: relational_expression -> d: ...
	DToRelationalExpression(RelationalExpression_ Symbol, LeOp_ Symbol, ShiftExpression_ Symbol) (Symbol, error)

	// 143:4: relational_expression -> e: ...
	EToRelationalExpression(RelationalExpression_ Symbol, GeOp_ Symbol, ShiftExpression_ Symbol) (Symbol, error)

	// 146:4: equality_expression -> a: ...
	AToEqualityExpression(RelationalExpression_ Symbol) (Symbol, error)

	// 147:4: equality_expression -> b: ...
	BToEqualityExpression(EqualityExpression_ Symbol, EqOp_ Symbol, RelationalExpression_ Symbol) (Symbol, error)

	// 148:4: equality_expression -> c: ...
	CToEqualityExpression(EqualityExpression_ Symbol, NeOp_ Symbol, RelationalExpression_ Symbol) (Symbol, error)

	// 151:4: and_expression -> a: ...
	AToAndExpression(EqualityExpression_ Symbol) (Symbol, error)

	// 152:4: and_expression -> b: ...
	BToAndExpression(AndExpression_ Symbol, And_ Symbol, EqualityExpression_ Symbol) (Symbol, error)

	// 155:4: exclusive_or_expression -> a: ...
	AToExclusiveOrExpression(AndExpression_ Symbol) (Symbol, error)

	// 156:4: exclusive_or_expression -> b: ...
	BToExclusiveOrExpression(ExclusiveOrExpression_ Symbol, Hat_ Symbol, AndExpression_ Symbol) (Symbol, error)

	// 159:4: inclusive_or_expression -> a: ...
	AToInclusiveOrExpression(ExclusiveOrExpression_ Symbol) (Symbol, error)

	// 160:4: inclusive_or_expression -> b: ...
	BToInclusiveOrExpression(InclusiveOrExpression_ Symbol, Or_ Symbol, ExclusiveOrExpression_ Symbol) (Symbol, error)

	// 163:4: logical_and_expression -> a: ...
	AToLogicalAndExpression(InclusiveOrExpression_ Symbol) (Symbol, error)

	// 164:4: logical_and_expression -> b: ...
	BToLogicalAndExpression(LogicalAndExpression_ Symbol, AndOp_ Symbol, InclusiveOrExpression_ Symbol) (Symbol, error)

	// 167:4: logical_or_expression -> a: ...
	AToLogicalOrExpression(LogicalAndExpression_ Symbol) (Symbol, error)

	// 168:4: logical_or_expression -> b: ...
	BToLogicalOrExpression(LogicalOrExpression_ Symbol, OrOp_ Symbol, LogicalAndExpression_ Symbol) (Symbol, error)

	// 171:4: conditional_expression -> a: ...
	AToConditionalExpression(LogicalOrExpression_ Symbol) (Symbol, error)

	// 172:4: conditional_expression -> b: ...
	BToConditionalExpression(LogicalOrExpression_ Symbol, Question_ Symbol, Expression_ Symbol, Colon_ Symbol, ConditionalExpression_ Symbol) (Symbol, error)

	// 175:4: assignment_expression -> a: ...
	AToAssignmentExpression(ConditionalExpression_ Symbol) (Symbol, error)

	// 176:4: assignment_expression -> b: ...
	BToAssignmentExpression(UnaryExpression_ Symbol, AssignmentOperator_ Symbol, AssignmentExpression_ Symbol) (Symbol, error)

	// 179:4: assignment_operator -> a: ...
	AToAssignmentOperator(Eq_ Symbol) (Symbol, error)

	// 180:4: assignment_operator -> b: ...
	BToAssignmentOperator(MulAssign_ Symbol) (Symbol, error)

	// 181:4: assignment_operator -> c: ...
	CToAssignmentOperator(DivAssign_ Symbol) (Symbol, error)

	// 182:4: assignment_operator -> d: ...
	DToAssignmentOperator(ModAssign_ Symbol) (Symbol, error)

	// 183:4: assignment_operator -> e: ...
	EToAssignmentOperator(AddAssign_ Symbol) (Symbol, error)

	// 184:4: assignment_operator -> f: ...
	FToAssignmentOperator(SubAssign_ Symbol) (Symbol, error)

	// 185:4: assignment_operator -> g: ...
	GToAssignmentOperator(LeftAssign_ Symbol) (Symbol, error)

	// 186:4: assignment_operator -> h: ...
	HToAssignmentOperator(RightAssign_ Symbol) (Symbol, error)

	// 187:4: assignment_operator -> i: ...
	IToAssignmentOperator(AndAssign_ Symbol) (Symbol, error)

	// 188:4: assignment_operator -> j: ...
	JToAssignmentOperator(XorAssign_ Symbol) (Symbol, error)

	// 189:4: assignment_operator -> k: ...
	KToAssignmentOperator(OrAssign_ Symbol) (Symbol, error)

	// 192:4: expression -> a: ...
	AToExpression(AssignmentExpression_ Symbol) (Symbol, error)

	// 193:4: expression -> b: ...
	BToExpression(Expression_ Symbol, Comma_ Symbol, AssignmentExpression_ Symbol) (Symbol, error)

	// 196:4: constant_expression -> a: ...
	AToConstantExpression(ConditionalExpression_ Symbol) (Symbol, error)

	// 199:4: declaration -> a: ...
	AToDeclaration(DeclarationSpecifiers_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 200:4: declaration -> b: ...
	BToDeclaration(DeclarationSpecifiers_ Symbol, InitDeclaratorList_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 203:4: declaration_specifiers -> a: ...
	AToDeclarationSpecifiers(StorageClassSpecifier_ Symbol) (Symbol, error)

	// 204:4: declaration_specifiers -> b: ...
	BToDeclarationSpecifiers(StorageClassSpecifier_ Symbol, DeclarationSpecifiers_ Symbol) (Symbol, error)

	// 205:4: declaration_specifiers -> c: ...
	CToDeclarationSpecifiers(TypeSpecifier_ Symbol) (Symbol, error)

	// 206:4: declaration_specifiers -> d: ...
	DToDeclarationSpecifiers(TypeSpecifier_ Symbol, DeclarationSpecifiers_ Symbol) (Symbol, error)

	// 207:4: declaration_specifiers -> e: ...
	EToDeclarationSpecifiers(TypeQualifier_ Symbol) (Symbol, error)

	// 208:4: declaration_specifiers -> f: ...
	FToDeclarationSpecifiers(TypeQualifier_ Symbol, DeclarationSpecifiers_ Symbol) (Symbol, error)

	// 211:4: init_declarator_list -> a: ...
	AToInitDeclaratorList(InitDeclarator_ Symbol) (Symbol, error)

	// 212:4: init_declarator_list -> b: ...
	BToInitDeclaratorList(InitDeclaratorList_ Symbol, Comma_ Symbol, InitDeclarator_ Symbol) (Symbol, error)

	// 215:4: init_declarator -> a: ...
	AToInitDeclarator(Declarator_ Symbol) (Symbol, error)

	// 216:4: init_declarator -> b: ...
	BToInitDeclarator(Declarator_ Symbol, Eq_ Symbol, Initializer_ Symbol) (Symbol, error)

	// 219:4: storage_class_specifier -> a: ...
	AToStorageClassSpecifier(Typedef_ Symbol) (Symbol, error)

	// 220:4: storage_class_specifier -> b: ...
	BToStorageClassSpecifier(Extern_ Symbol) (Symbol, error)

	// 221:4: storage_class_specifier -> c: ...
	CToStorageClassSpecifier(Static_ Symbol) (Symbol, error)

	// 222:4: storage_class_specifier -> d: ...
	DToStorageClassSpecifier(Auto_ Symbol) (Symbol, error)

	// 223:4: storage_class_specifier -> e: ...
	EToStorageClassSpecifier(Register_ Symbol) (Symbol, error)

	// 226:4: type_specifier -> a: ...
	AToTypeSpecifier(Void_ Symbol) (Symbol, error)

	// 227:4: type_specifier -> b: ...
	BToTypeSpecifier(Char_ Symbol) (Symbol, error)

	// 228:4: type_specifier -> c: ...
	CToTypeSpecifier(Short_ Symbol) (Symbol, error)

	// 229:4: type_specifier -> d: ...
	DToTypeSpecifier(Int_ Symbol) (Symbol, error)

	// 230:4: type_specifier -> e: ...
	EToTypeSpecifier(Long_ Symbol) (Symbol, error)

	// 231:4: type_specifier -> f: ...
	FToTypeSpecifier(Float_ Symbol) (Symbol, error)

	// 232:4: type_specifier -> g: ...
	GToTypeSpecifier(Double_ Symbol) (Symbol, error)

	// 233:4: type_specifier -> h: ...
	HToTypeSpecifier(Signed_ Symbol) (Symbol, error)

	// 234:4: type_specifier -> i: ...
	IToTypeSpecifier(Unsigned_ Symbol) (Symbol, error)

	// 235:4: type_specifier -> j: ...
	JToTypeSpecifier(StructOrUnionSpecifier_ Symbol) (Symbol, error)

	// 236:4: type_specifier -> k: ...
	KToTypeSpecifier(EnumSpecifier_ Symbol) (Symbol, error)

	// 237:4: type_specifier -> l: ...
	LToTypeSpecifier(TypeName_ Symbol) (Symbol, error)

	// 240:4: struct_or_union_specifier -> a: ...
	AToStructOrUnionSpecifier(StructOrUnion_ Symbol, Identifier_ Symbol, Lcurl_ Symbol, StructDeclarationList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 241:4: struct_or_union_specifier -> b: ...
	BToStructOrUnionSpecifier(StructOrUnion_ Symbol, Lcurl_ Symbol, StructDeclarationList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 242:4: struct_or_union_specifier -> c: ...
	CToStructOrUnionSpecifier(StructOrUnion_ Symbol, Identifier_ Symbol) (Symbol, error)

	// 245:4: struct_or_union -> a: ...
	AToStructOrUnion(Struct_ Symbol) (Symbol, error)

	// 246:4: struct_or_union -> b: ...
	BToStructOrUnion(Union_ Symbol) (Symbol, error)

	// 249:4: struct_declaration_list -> a: ...
	AToStructDeclarationList(StructDeclaration_ Symbol) (Symbol, error)

	// 250:4: struct_declaration_list -> b: ...
	BToStructDeclarationList(StructDeclarationList_ Symbol, StructDeclaration_ Symbol) (Symbol, error)

	// 253:4: struct_declaration -> a: ...
	AToStructDeclaration(SpecifierQualifierList_ Symbol, StructDeclaratorList_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 256:4: specifier_qualifier_list -> a: ...
	AToSpecifierQualifierList(TypeSpecifier_ Symbol, SpecifierQualifierList_ Symbol) (Symbol, error)

	// 257:4: specifier_qualifier_list -> b: ...
	BToSpecifierQualifierList(TypeSpecifier_ Symbol) (Symbol, error)

	// 258:4: specifier_qualifier_list -> c: ...
	CToSpecifierQualifierList(TypeQualifier_ Symbol, SpecifierQualifierList_ Symbol) (Symbol, error)

	// 259:4: specifier_qualifier_list -> d: ...
	DToSpecifierQualifierList(TypeQualifier_ Symbol) (Symbol, error)

	// 262:4: struct_declarator_list -> a: ...
	AToStructDeclaratorList(StructDeclarator_ Symbol) (Symbol, error)

	// 263:4: struct_declarator_list -> b: ...
	BToStructDeclaratorList(StructDeclaratorList_ Symbol, Comma_ Symbol, StructDeclarator_ Symbol) (Symbol, error)

	// 266:4: struct_declarator -> a: ...
	AToStructDeclarator(Declarator_ Symbol) (Symbol, error)

	// 267:4: struct_declarator -> b: ...
	BToStructDeclarator(Colon_ Symbol, ConstantExpression_ Symbol) (Symbol, error)

	// 268:4: struct_declarator -> c: ...
	CToStructDeclarator(Declarator_ Symbol, Colon_ Symbol, ConstantExpression_ Symbol) (Symbol, error)

	// 271:4: enum_specifier -> a: ...
	AToEnumSpecifier(Enum_ Symbol, Lcurl_ Symbol, EnumeratorList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 272:4: enum_specifier -> b: ...
	BToEnumSpecifier(Enum_ Symbol, Identifier_ Symbol, Lcurl_ Symbol, EnumeratorList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 273:4: enum_specifier -> c: ...
	CToEnumSpecifier(Enum_ Symbol, Identifier_ Symbol) (Symbol, error)

	// 276:4: enumerator_list -> a: ...
	AToEnumeratorList(Enumerator_ Symbol) (Symbol, error)

	// 277:4: enumerator_list -> b: ...
	BToEnumeratorList(EnumeratorList_ Symbol, Comma_ Symbol, Enumerator_ Symbol) (Symbol, error)

	// 280:4: enumerator -> a: ...
	AToEnumerator(Identifier_ Symbol) (Symbol, error)

	// 281:4: enumerator -> b: ...
	BToEnumerator(Identifier_ Symbol, Eq_ Symbol, ConstantExpression_ Symbol) (Symbol, error)

	// 284:4: type_qualifier -> a: ...
	AToTypeQualifier(Const_ Symbol) (Symbol, error)

	// 285:4: type_qualifier -> b: ...
	BToTypeQualifier(Volatile_ Symbol) (Symbol, error)

	// 288:4: declarator -> a: ...
	AToDeclarator(Pointer_ Symbol, DirectDeclarator_ Symbol) (Symbol, error)

	// 289:4: declarator -> b: ...
	BToDeclarator(DirectDeclarator_ Symbol) (Symbol, error)

	// 292:4: direct_declarator -> a: ...
	AToDirectDeclarator(Identifier_ Symbol) (Symbol, error)

	// 293:4: direct_declarator -> b: ...
	BToDirectDeclarator(Lparam_ Symbol, Declarator_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 294:4: direct_declarator -> c: ...
	CToDirectDeclarator(DirectDeclarator_ Symbol, Lbrace_ Symbol, ConstantExpression_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 295:4: direct_declarator -> d: ...
	DToDirectDeclarator(DirectDeclarator_ Symbol, Lbrace_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 296:4: direct_declarator -> e: ...
	EToDirectDeclarator(DirectDeclarator_ Symbol, Lparam_ Symbol, ParameterTypeList_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 297:4: direct_declarator -> f: ...
	FToDirectDeclarator(DirectDeclarator_ Symbol, Lparam_ Symbol, IdentifierList_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 298:4: direct_declarator -> g: ...
	GToDirectDeclarator(DirectDeclarator_ Symbol, Lparam_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 301:4: pointer -> a: ...
	AToPointer(Mul_ Symbol) (Symbol, error)

	// 302:4: pointer -> b: ...
	BToPointer(Mul_ Symbol, TypeQualifierList_ Symbol) (Symbol, error)

	// 303:4: pointer -> c: ...
	CToPointer(Mul_ Symbol, Pointer_ Symbol) (Symbol, error)

	// 304:4: pointer -> d: ...
	DToPointer(Mul_ Symbol, TypeQualifierList_ Symbol, Pointer_ Symbol) (Symbol, error)

	// 307:4: type_qualifier_list -> a: ...
	AToTypeQualifierList(TypeQualifier_ Symbol) (Symbol, error)

	// 308:4: type_qualifier_list -> b: ...
	BToTypeQualifierList(TypeQualifierList_ Symbol, TypeQualifier_ Symbol) (Symbol, error)

	// 312:4: parameter_type_list -> a: ...
	AToParameterTypeList(ParameterList_ Symbol) (Symbol, error)

	// 313:4: parameter_type_list -> b: ...
	BToParameterTypeList(ParameterList_ Symbol, Comma_ Symbol, Ellipsis_ Symbol) (Symbol, error)

	// 316:4: parameter_list -> a: ...
	AToParameterList(ParameterDeclaration_ Symbol) (Symbol, error)

	// 317:4: parameter_list -> b: ...
	BToParameterList(ParameterList_ Symbol, Comma_ Symbol, ParameterDeclaration_ Symbol) (Symbol, error)

	// 320:4: parameter_declaration -> a: ...
	AToParameterDeclaration(DeclarationSpecifiers_ Symbol, Declarator_ Symbol) (Symbol, error)

	// 321:4: parameter_declaration -> b: ...
	BToParameterDeclaration(DeclarationSpecifiers_ Symbol, AbstractDeclarator_ Symbol) (Symbol, error)

	// 322:4: parameter_declaration -> c: ...
	CToParameterDeclaration(DeclarationSpecifiers_ Symbol) (Symbol, error)

	// 325:4: identifier_list -> a: ...
	AToIdentifierList(Identifier_ Symbol) (Symbol, error)

	// 326:4: identifier_list -> b: ...
	BToIdentifierList(IdentifierList_ Symbol, Comma_ Symbol, Identifier_ Symbol) (Symbol, error)

	// 329:4: type_name -> a: ...
	AToTypeName(SpecifierQualifierList_ Symbol) (Symbol, error)

	// 330:4: type_name -> b: ...
	BToTypeName(SpecifierQualifierList_ Symbol, AbstractDeclarator_ Symbol) (Symbol, error)

	// 333:4: abstract_declarator -> a: ...
	AToAbstractDeclarator(Pointer_ Symbol) (Symbol, error)

	// 334:4: abstract_declarator -> b: ...
	BToAbstractDeclarator(DirectAbstractDeclarator_ Symbol) (Symbol, error)

	// 335:4: abstract_declarator -> c: ...
	CToAbstractDeclarator(Pointer_ Symbol, DirectAbstractDeclarator_ Symbol) (Symbol, error)

	// 338:4: direct_abstract_declarator -> a: ...
	AToDirectAbstractDeclarator(Lparam_ Symbol, AbstractDeclarator_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 339:4: direct_abstract_declarator -> b: ...
	BToDirectAbstractDeclarator(Lbrace_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 340:4: direct_abstract_declarator -> c: ...
	CToDirectAbstractDeclarator(Lbrace_ Symbol, ConstantExpression_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 341:4: direct_abstract_declarator -> d: ...
	DToDirectAbstractDeclarator(DirectAbstractDeclarator_ Symbol, Lbrace_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 342:4: direct_abstract_declarator -> e: ...
	EToDirectAbstractDeclarator(DirectAbstractDeclarator_ Symbol, Lbrace_ Symbol, ConstantExpression_ Symbol, Rbrace_ Symbol) (Symbol, error)

	// 343:4: direct_abstract_declarator -> f: ...
	FToDirectAbstractDeclarator(Lparam_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 344:4: direct_abstract_declarator -> g: ...
	GToDirectAbstractDeclarator(Lparam_ Symbol, ParameterTypeList_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 345:4: direct_abstract_declarator -> h: ...
	HToDirectAbstractDeclarator(DirectAbstractDeclarator_ Symbol, Lparam_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 346:4: direct_abstract_declarator -> i: ...
	IToDirectAbstractDeclarator(DirectAbstractDeclarator_ Symbol, Lparam_ Symbol, ParameterTypeList_ Symbol, Rparam_ Symbol) (Symbol, error)

	// 349:4: initializer -> a: ...
	AToInitializer(AssignmentExpression_ Symbol) (Symbol, error)

	// 350:4: initializer -> b: ...
	BToInitializer(Lcurl_ Symbol, InitializerList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 351:4: initializer -> c: ...
	CToInitializer(Lcurl_ Symbol, InitializerList_ Symbol, Comma_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 354:4: initializer_list -> a: ...
	AToInitializerList(Initializer_ Symbol) (Symbol, error)

	// 355:4: initializer_list -> b: ...
	BToInitializerList(InitializerList_ Symbol, Comma_ Symbol, Initializer_ Symbol) (Symbol, error)

	// 358:4: statement -> a: ...
	AToStatement(LabeledStatement_ Symbol) (Symbol, error)

	// 359:4: statement -> b: ...
	BToStatement(CompoundStatement_ Symbol) (Symbol, error)

	// 360:4: statement -> c: ...
	CToStatement(ExpressionStatement_ Symbol) (Symbol, error)

	// 361:4: statement -> d: ...
	DToStatement(SelectionStatement_ Symbol) (Symbol, error)

	// 362:4: statement -> e: ...
	EToStatement(IterationStatement_ Symbol) (Symbol, error)

	// 363:4: statement -> f: ...
	FToStatement(JumpStatement_ Symbol) (Symbol, error)

	// 366:4: labeled_statement -> a: ...
	AToLabeledStatement(Identifier_ Symbol, Colon_ Symbol, Statement_ Symbol) (Symbol, error)

	// 367:4: labeled_statement -> b: ...
	BToLabeledStatement(Case_ Symbol, ConstantExpression_ Symbol, Colon_ Symbol, Statement_ Symbol) (Symbol, error)

	// 368:4: labeled_statement -> c: ...
	CToLabeledStatement(Default_ Symbol, Colon_ Symbol, Statement_ Symbol) (Symbol, error)

	// 371:4: compound_statement -> a: ...
	AToCompoundStatement(Lcurl_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 372:4: compound_statement -> b: ...
	BToCompoundStatement(Lcurl_ Symbol, StatementList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 373:4: compound_statement -> c: ...
	CToCompoundStatement(Lcurl_ Symbol, DeclarationList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 374:4: compound_statement -> d: ...
	DToCompoundStatement(Lcurl_ Symbol, DeclarationList_ Symbol, StatementList_ Symbol, Rcurl_ Symbol) (Symbol, error)

	// 377:4: declaration_list -> a: ...
	AToDeclarationList(Declaration_ Symbol) (Symbol, error)

	// 378:4: declaration_list -> b: ...
	BToDeclarationList(DeclarationList_ Symbol, Declaration_ Symbol) (Symbol, error)

	// 381:4: statement_list -> a: ...
	AToStatementList(Statement_ Symbol) (Symbol, error)

	// 382:4: statement_list -> b: ...
	BToStatementList(StatementList_ Symbol, Statement_ Symbol) (Symbol, error)

	// 385:4: expression_statement -> a: ...
	AToExpressionStatement(Semicolon_ Symbol) (Symbol, error)

	// 386:4: expression_statement -> b: ...
	BToExpressionStatement(Expression_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 389:4: selection_statement -> a: ...
	AToSelectionStatement(If_ Symbol, Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol, Statement_ Symbol) (Symbol, error)

	// 390:4: selection_statement -> b: ...
	BToSelectionStatement(If_ Symbol, Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol, Statement_ Symbol, Else_ Symbol, Statement_2 Symbol) (Symbol, error)

	// 391:4: selection_statement -> c: ...
	CToSelectionStatement(Switch_ Symbol, Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol, Statement_ Symbol) (Symbol, error)

	// 394:4: iteration_statement -> a: ...
	AToIterationStatement(While_ Symbol, Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol, Statement_ Symbol) (Symbol, error)

	// 395:4: iteration_statement -> b: ...
	BToIterationStatement(Do_ Symbol, Statement_ Symbol, While_ Symbol, Lparam_ Symbol, Expression_ Symbol, Rparam_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 396:4: iteration_statement -> c: ...
	CToIterationStatement(For_ Symbol, Lparam_ Symbol, ExpressionStatement_ Symbol, ExpressionStatement_2 Symbol, Rparam_ Symbol, Statement_ Symbol) (Symbol, error)

	// 397:4: iteration_statement -> d: ...
	DToIterationStatement(For_ Symbol, Lparam_ Symbol, ExpressionStatement_ Symbol, ExpressionStatement_2 Symbol, Expression_ Symbol, Rparam_ Symbol, Statement_ Symbol) (Symbol, error)

	// 400:4: jump_statement -> a: ...
	AToJumpStatement(Goto_ Symbol, Identifier_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 401:4: jump_statement -> b: ...
	BToJumpStatement(Continue_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 402:4: jump_statement -> c: ...
	CToJumpStatement(Break_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 403:4: jump_statement -> d: ...
	DToJumpStatement(Return_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 404:4: jump_statement -> e: ...
	EToJumpStatement(Return_ Symbol, Expression_ Symbol, Semicolon_ Symbol) (Symbol, error)

	// 407:4: translation_unit -> a: ...
	AToTranslationUnit(ExternalDeclaration_ Symbol) (Symbol, error)

	// 408:4: translation_unit -> b: ...
	BToTranslationUnit(TranslationUnit_ Symbol, ExternalDeclaration_ Symbol) (Symbol, error)

	// 411:4: external_declaration -> a: ...
	AToExternalDeclaration(FunctionDefinition_ Symbol) (Symbol, error)

	// 412:4: external_declaration -> b: ...
	BToExternalDeclaration(Declaration_ Symbol) (Symbol, error)

	// 415:4: function_definition -> a: ...
	AToFunctionDefinition(DeclarationSpecifiers_ Symbol, Declarator_ Symbol, DeclarationList_ Symbol, CompoundStatement_ Symbol) (Symbol, error)

	// 416:4: function_definition -> b: ...
	BToFunctionDefinition(DeclarationSpecifiers_ Symbol, Declarator_ Symbol, CompoundStatement_ Symbol) (Symbol, error)

	// 417:4: function_definition -> c: ...
	CToFunctionDefinition(Declarator_ Symbol, DeclarationList_ Symbol, CompoundStatement_ Symbol) (Symbol, error)

	// 418:4: function_definition -> d: ...
	DToFunctionDefinition(Declarator_ Symbol, CompoundStatement_ Symbol) (Symbol, error)
}

type CParseErrorHandler interface {
	Error(nextToken CToken, parseStack _CStack) error
}

type CDefaultParseErrorHandler struct{}

func (CDefaultParseErrorHandler) Error(nextToken CToken, stack _CStack) error {
	return fmt.Errorf("Syntax error: unexpected symbol %v. Expecting: %v (%v)", nextToken.Id(), _CExpectedTerminals[stack[len(stack)-1].StateId], nextToken.Loc())
}

func CParse(lexer CLexer, reducer CReducer) (Symbol, error) {
	return CParseWithCustomErrorHandler(lexer, reducer, CDefaultParseErrorHandler{})
}

func CParseWithCustomErrorHandler(lexer CLexer, reducer CReducer, errHandler CParseErrorHandler) (Symbol, error) {
	item, err := _CParse(lexer, reducer, errHandler, _CState1)
	if err != nil {
		var errRetVal Symbol
		return errRetVal, err
	}
	return item.T, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _CParse(lexer CLexer, reducer CReducer, errHandler CParseErrorHandler, startState _CStateId) (*_CStackItem, error) {
	stateStack := _CStack{
		// Note: we don't have to populate the start symbol since its value is never accessed
		&_CStackItem{startState, nil},
	}
	symbolStack := &_CPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _CActionTable.Get(stateStack[len(stateStack)-1].StateId, nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}
		if action.ActionType == _CShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _CReduceAction {
			var reduceSymbol *CSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(reducer, stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _CAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil

		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i CSymbolId) String() string {
	switch i {
	case _CEndMarker:
		return "$"
	case _CWildcardMarker:
		return "*"
	case CIdentifierToken:
		return "IDENTIFIER"
	case CConstantToken:
		return "CONSTANT"
	case CStringLiteralToken:
		return "STRING_LITERAL"
	case CSizeofToken:
		return "SIZEOF"
	case CPtrOpToken:
		return "PTR_OP"
	case CIncOpToken:
		return "INC_OP"
	case CDecOpToken:
		return "DEC_OP"
	case CLeftOpToken:
		return "LEFT_OP"
	case CRightOpToken:
		return "RIGHT_OP"
	case CLeOpToken:
		return "LE_OP"
	case CGeOpToken:
		return "GE_OP"
	case CEqOpToken:
		return "EQ_OP"
	case CNeOpToken:
		return "NE_OP"
	case CAndOpToken:
		return "AND_OP"
	case COrOpToken:
		return "OR_OP"
	case CMulAssignToken:
		return "MUL_ASSIGN"
	case CDivAssignToken:
		return "DIV_ASSIGN"
	case CModAssignToken:
		return "MOD_ASSIGN"
	case CAddAssignToken:
		return "ADD_ASSIGN"
	case CSubAssignToken:
		return "SUB_ASSIGN"
	case CLeftAssignToken:
		return "LEFT_ASSIGN"
	case CRightAssignToken:
		return "RIGHT_ASSIGN"
	case CAndAssignToken:
		return "AND_ASSIGN"
	case CXorAssignToken:
		return "XOR_ASSIGN"
	case COrAssignToken:
		return "OR_ASSIGN"
	case CTypeNameToken:
		return "TYPE_NAME"
	case CTypedefToken:
		return "TYPEDEF"
	case CExternToken:
		return "EXTERN"
	case CStaticToken:
		return "STATIC"
	case CAutoToken:
		return "AUTO"
	case CRegisterToken:
		return "REGISTER"
	case CCharToken:
		return "CHAR"
	case CShortToken:
		return "SHORT"
	case CIntToken:
		return "INT"
	case CLongToken:
		return "LONG"
	case CSignedToken:
		return "SIGNED"
	case CUnsignedToken:
		return "UNSIGNED"
	case CFloatToken:
		return "FLOAT"
	case CDoubleToken:
		return "DOUBLE"
	case CConstToken:
		return "CONST"
	case CVolatileToken:
		return "VOLATILE"
	case CVoidToken:
		return "VOID"
	case CStructToken:
		return "STRUCT"
	case CUnionToken:
		return "UNION"
	case CEnumToken:
		return "ENUM"
	case CEllipsisToken:
		return "ELLIPSIS"
	case CCaseToken:
		return "CASE"
	case CDefaultToken:
		return "DEFAULT"
	case CIfToken:
		return "IF"
	case CElseToken:
		return "ELSE"
	case CSwitchToken:
		return "SWITCH"
	case CWhileToken:
		return "WHILE"
	case CDoToken:
		return "DO"
	case CForToken:
		return "FOR"
	case CGotoToken:
		return "GOTO"
	case CContinueToken:
		return "CONTINUE"
	case CBreakToken:
		return "BREAK"
	case CReturnToken:
		return "RETURN"
	case CLparamToken:
		return "LPARAM"
	case CRparamToken:
		return "RPARAM"
	case CLcurlToken:
		return "LCURL"
	case CRcurlToken:
		return "RCURL"
	case CLbraceToken:
		return "LBRACE"
	case CRbraceToken:
		return "RBRACE"
	case CSemicolonToken:
		return "SEMICOLON"
	case CColonToken:
		return "COLON"
	case CCommaToken:
		return "COMMA"
	case CEqToken:
		return "EQ"
	case CQuestionToken:
		return "QUESTION"
	case CMulToken:
		return "MUL"
	case CDivToken:
		return "DIV"
	case CMinusToken:
		return "MINUS"
	case CPlusToken:
		return "PLUS"
	case CModToken:
		return "MOD"
	case CAndToken:
		return "AND"
	case COrToken:
		return "OR"
	case CExclaimToken:
		return "EXCLAIM"
	case CDotToken:
		return "DOT"
	case CHatToken:
		return "HAT"
	case CLtToken:
		return "LT"
	case CGtToken:
		return "GT"
	case CTildaToken:
		return "TILDA"
	case CPrimaryExpressionType:
		return "primary_expression"
	case CPostfixExpressionType:
		return "postfix_expression"
	case CArgumentExpressionListType:
		return "argument_expression_list"
	case CUnaryExpressionType:
		return "unary_expression"
	case CUnaryOperatorType:
		return "unary_operator"
	case CCastExpressionType:
		return "cast_expression"
	case CMultiplicativeExpressionType:
		return "multiplicative_expression"
	case CAdditiveExpressionType:
		return "additive_expression"
	case CShiftExpressionType:
		return "shift_expression"
	case CRelationalExpressionType:
		return "relational_expression"
	case CEqualityExpressionType:
		return "equality_expression"
	case CAndExpressionType:
		return "and_expression"
	case CExclusiveOrExpressionType:
		return "exclusive_or_expression"
	case CInclusiveOrExpressionType:
		return "inclusive_or_expression"
	case CLogicalAndExpressionType:
		return "logical_and_expression"
	case CLogicalOrExpressionType:
		return "logical_or_expression"
	case CConditionalExpressionType:
		return "conditional_expression"
	case CAssignmentExpressionType:
		return "assignment_expression"
	case CAssignmentOperatorType:
		return "assignment_operator"
	case CExpressionType:
		return "expression"
	case CConstantExpressionType:
		return "constant_expression"
	case CDeclarationType:
		return "declaration"
	case CDeclarationSpecifiersType:
		return "declaration_specifiers"
	case CInitDeclaratorListType:
		return "init_declarator_list"
	case CInitDeclaratorType:
		return "init_declarator"
	case CStorageClassSpecifierType:
		return "storage_class_specifier"
	case CTypeSpecifierType:
		return "type_specifier"
	case CStructOrUnionSpecifierType:
		return "struct_or_union_specifier"
	case CStructOrUnionType:
		return "struct_or_union"
	case CStructDeclarationListType:
		return "struct_declaration_list"
	case CStructDeclarationType:
		return "struct_declaration"
	case CSpecifierQualifierListType:
		return "specifier_qualifier_list"
	case CStructDeclaratorListType:
		return "struct_declarator_list"
	case CStructDeclaratorType:
		return "struct_declarator"
	case CEnumSpecifierType:
		return "enum_specifier"
	case CEnumeratorListType:
		return "enumerator_list"
	case CEnumeratorType:
		return "enumerator"
	case CTypeQualifierType:
		return "type_qualifier"
	case CDeclaratorType:
		return "declarator"
	case CDirectDeclaratorType:
		return "direct_declarator"
	case CPointerType:
		return "pointer"
	case CTypeQualifierListType:
		return "type_qualifier_list"
	case CParameterTypeListType:
		return "parameter_type_list"
	case CParameterListType:
		return "parameter_list"
	case CParameterDeclarationType:
		return "parameter_declaration"
	case CIdentifierListType:
		return "identifier_list"
	case CTypeNameType:
		return "type_name"
	case CAbstractDeclaratorType:
		return "abstract_declarator"
	case CDirectAbstractDeclaratorType:
		return "direct_abstract_declarator"
	case CInitializerType:
		return "initializer"
	case CInitializerListType:
		return "initializer_list"
	case CStatementType:
		return "statement"
	case CLabeledStatementType:
		return "labeled_statement"
	case CCompoundStatementType:
		return "compound_statement"
	case CDeclarationListType:
		return "declaration_list"
	case CStatementListType:
		return "statement_list"
	case CExpressionStatementType:
		return "expression_statement"
	case CSelectionStatementType:
		return "selection_statement"
	case CIterationStatementType:
		return "iteration_statement"
	case CJumpStatementType:
		return "jump_statement"
	case CTranslationUnitType:
		return "translation_unit"
	case CExternalDeclarationType:
		return "external_declaration"
	case CFunctionDefinitionType:
		return "function_definition"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_CEndMarker      = CSymbolId(0)
	_CWildcardMarker = CSymbolId(-1)

	CPrimaryExpressionType        = CSymbolId(338)
	CPostfixExpressionType        = CSymbolId(339)
	CArgumentExpressionListType   = CSymbolId(340)
	CUnaryExpressionType          = CSymbolId(341)
	CUnaryOperatorType            = CSymbolId(342)
	CCastExpressionType           = CSymbolId(343)
	CMultiplicativeExpressionType = CSymbolId(344)
	CAdditiveExpressionType       = CSymbolId(345)
	CShiftExpressionType          = CSymbolId(346)
	CRelationalExpressionType     = CSymbolId(347)
	CEqualityExpressionType       = CSymbolId(348)
	CAndExpressionType            = CSymbolId(349)
	CExclusiveOrExpressionType    = CSymbolId(350)
	CInclusiveOrExpressionType    = CSymbolId(351)
	CLogicalAndExpressionType     = CSymbolId(352)
	CLogicalOrExpressionType      = CSymbolId(353)
	CConditionalExpressionType    = CSymbolId(354)
	CAssignmentExpressionType     = CSymbolId(355)
	CAssignmentOperatorType       = CSymbolId(356)
	CExpressionType               = CSymbolId(357)
	CConstantExpressionType       = CSymbolId(358)
	CDeclarationType              = CSymbolId(359)
	CDeclarationSpecifiersType    = CSymbolId(360)
	CInitDeclaratorListType       = CSymbolId(361)
	CInitDeclaratorType           = CSymbolId(362)
	CStorageClassSpecifierType    = CSymbolId(363)
	CTypeSpecifierType            = CSymbolId(364)
	CStructOrUnionSpecifierType   = CSymbolId(365)
	CStructOrUnionType            = CSymbolId(366)
	CStructDeclarationListType    = CSymbolId(367)
	CStructDeclarationType        = CSymbolId(368)
	CSpecifierQualifierListType   = CSymbolId(369)
	CStructDeclaratorListType     = CSymbolId(370)
	CStructDeclaratorType         = CSymbolId(371)
	CEnumSpecifierType            = CSymbolId(372)
	CEnumeratorListType           = CSymbolId(373)
	CEnumeratorType               = CSymbolId(374)
	CTypeQualifierType            = CSymbolId(375)
	CDeclaratorType               = CSymbolId(376)
	CDirectDeclaratorType         = CSymbolId(377)
	CPointerType                  = CSymbolId(378)
	CTypeQualifierListType        = CSymbolId(379)
	CParameterTypeListType        = CSymbolId(380)
	CParameterListType            = CSymbolId(381)
	CParameterDeclarationType     = CSymbolId(382)
	CIdentifierListType           = CSymbolId(383)
	CTypeNameType                 = CSymbolId(384)
	CAbstractDeclaratorType       = CSymbolId(385)
	CDirectAbstractDeclaratorType = CSymbolId(386)
	CInitializerType              = CSymbolId(387)
	CInitializerListType          = CSymbolId(388)
	CStatementType                = CSymbolId(389)
	CLabeledStatementType         = CSymbolId(390)
	CCompoundStatementType        = CSymbolId(391)
	CDeclarationListType          = CSymbolId(392)
	CStatementListType            = CSymbolId(393)
	CExpressionStatementType      = CSymbolId(394)
	CSelectionStatementType       = CSymbolId(395)
	CIterationStatementType       = CSymbolId(396)
	CJumpStatementType            = CSymbolId(397)
	CTranslationUnitType          = CSymbolId(398)
	CExternalDeclarationType      = CSymbolId(399)
	CFunctionDefinitionType       = CSymbolId(400)
)

type _CActionType int

const (
	// NOTE: error action is implicit
	_CShiftAction  = _CActionType(0)
	_CReduceAction = _CActionType(1)
	_CAcceptAction = _CActionType(2)
)

func (i _CActionType) String() string {
	switch i {
	case _CShiftAction:
		return "shift"
	case _CReduceAction:
		return "reduce"
	case _CAcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?unknown action %d", int(i))
	}
}

type _CReduceType int

const (
	_CReduceAToPrimaryExpression        = _CReduceType(1)
	_CReduceBToPrimaryExpression        = _CReduceType(2)
	_CReduceCToPrimaryExpression        = _CReduceType(3)
	_CReduceDToPrimaryExpression        = _CReduceType(4)
	_CReduceAToPostfixExpression        = _CReduceType(5)
	_CReduceBToPostfixExpression        = _CReduceType(6)
	_CReduceCToPostfixExpression        = _CReduceType(7)
	_CReduceDToPostfixExpression        = _CReduceType(8)
	_CReduceEToPostfixExpression        = _CReduceType(9)
	_CReduceFToPostfixExpression        = _CReduceType(10)
	_CReduceGToPostfixExpression        = _CReduceType(11)
	_CReduceHToPostfixExpression        = _CReduceType(12)
	_CReduceAToArgumentExpressionList   = _CReduceType(13)
	_CReduceBToArgumentExpressionList   = _CReduceType(14)
	_CReduceAToUnaryExpression          = _CReduceType(15)
	_CReduceBToUnaryExpression          = _CReduceType(16)
	_CReduceCToUnaryExpression          = _CReduceType(17)
	_CReduceDToUnaryExpression          = _CReduceType(18)
	_CReduceEToUnaryExpression          = _CReduceType(19)
	_CReduceFToUnaryExpression          = _CReduceType(20)
	_CReduceAToUnaryOperator            = _CReduceType(21)
	_CReduceBToUnaryOperator            = _CReduceType(22)
	_CReduceCToUnaryOperator            = _CReduceType(23)
	_CReduceDToUnaryOperator            = _CReduceType(24)
	_CReduceEToUnaryOperator            = _CReduceType(25)
	_CReduceFToUnaryOperator            = _CReduceType(26)
	_CReduceAToCastExpression           = _CReduceType(27)
	_CReduceBToCastExpression           = _CReduceType(28)
	_CReduceAToMultiplicativeExpression = _CReduceType(29)
	_CReduceBToMultiplicativeExpression = _CReduceType(30)
	_CReduceCToMultiplicativeExpression = _CReduceType(31)
	_CReduceDToMultiplicativeExpression = _CReduceType(32)
	_CReduceAToAdditiveExpression       = _CReduceType(33)
	_CReduceBToAdditiveExpression       = _CReduceType(34)
	_CReduceCToAdditiveExpression       = _CReduceType(35)
	_CReduceAToShiftExpression          = _CReduceType(36)
	_CReduceBToShiftExpression          = _CReduceType(37)
	_CReduceCToShiftExpression          = _CReduceType(38)
	_CReduceAToRelationalExpression     = _CReduceType(39)
	_CReduceBToRelationalExpression     = _CReduceType(40)
	_CReduceCToRelationalExpression     = _CReduceType(41)
	_CReduceDToRelationalExpression     = _CReduceType(42)
	_CReduceEToRelationalExpression     = _CReduceType(43)
	_CReduceAToEqualityExpression       = _CReduceType(44)
	_CReduceBToEqualityExpression       = _CReduceType(45)
	_CReduceCToEqualityExpression       = _CReduceType(46)
	_CReduceAToAndExpression            = _CReduceType(47)
	_CReduceBToAndExpression            = _CReduceType(48)
	_CReduceAToExclusiveOrExpression    = _CReduceType(49)
	_CReduceBToExclusiveOrExpression    = _CReduceType(50)
	_CReduceAToInclusiveOrExpression    = _CReduceType(51)
	_CReduceBToInclusiveOrExpression    = _CReduceType(52)
	_CReduceAToLogicalAndExpression     = _CReduceType(53)
	_CReduceBToLogicalAndExpression     = _CReduceType(54)
	_CReduceAToLogicalOrExpression      = _CReduceType(55)
	_CReduceBToLogicalOrExpression      = _CReduceType(56)
	_CReduceAToConditionalExpression    = _CReduceType(57)
	_CReduceBToConditionalExpression    = _CReduceType(58)
	_CReduceAToAssignmentExpression     = _CReduceType(59)
	_CReduceBToAssignmentExpression     = _CReduceType(60)
	_CReduceAToAssignmentOperator       = _CReduceType(61)
	_CReduceBToAssignmentOperator       = _CReduceType(62)
	_CReduceCToAssignmentOperator       = _CReduceType(63)
	_CReduceDToAssignmentOperator       = _CReduceType(64)
	_CReduceEToAssignmentOperator       = _CReduceType(65)
	_CReduceFToAssignmentOperator       = _CReduceType(66)
	_CReduceGToAssignmentOperator       = _CReduceType(67)
	_CReduceHToAssignmentOperator       = _CReduceType(68)
	_CReduceIToAssignmentOperator       = _CReduceType(69)
	_CReduceJToAssignmentOperator       = _CReduceType(70)
	_CReduceKToAssignmentOperator       = _CReduceType(71)
	_CReduceAToExpression               = _CReduceType(72)
	_CReduceBToExpression               = _CReduceType(73)
	_CReduceAToConstantExpression       = _CReduceType(74)
	_CReduceAToDeclaration              = _CReduceType(75)
	_CReduceBToDeclaration              = _CReduceType(76)
	_CReduceAToDeclarationSpecifiers    = _CReduceType(77)
	_CReduceBToDeclarationSpecifiers    = _CReduceType(78)
	_CReduceCToDeclarationSpecifiers    = _CReduceType(79)
	_CReduceDToDeclarationSpecifiers    = _CReduceType(80)
	_CReduceEToDeclarationSpecifiers    = _CReduceType(81)
	_CReduceFToDeclarationSpecifiers    = _CReduceType(82)
	_CReduceAToInitDeclaratorList       = _CReduceType(83)
	_CReduceBToInitDeclaratorList       = _CReduceType(84)
	_CReduceAToInitDeclarator           = _CReduceType(85)
	_CReduceBToInitDeclarator           = _CReduceType(86)
	_CReduceAToStorageClassSpecifier    = _CReduceType(87)
	_CReduceBToStorageClassSpecifier    = _CReduceType(88)
	_CReduceCToStorageClassSpecifier    = _CReduceType(89)
	_CReduceDToStorageClassSpecifier    = _CReduceType(90)
	_CReduceEToStorageClassSpecifier    = _CReduceType(91)
	_CReduceAToTypeSpecifier            = _CReduceType(92)
	_CReduceBToTypeSpecifier            = _CReduceType(93)
	_CReduceCToTypeSpecifier            = _CReduceType(94)
	_CReduceDToTypeSpecifier            = _CReduceType(95)
	_CReduceEToTypeSpecifier            = _CReduceType(96)
	_CReduceFToTypeSpecifier            = _CReduceType(97)
	_CReduceGToTypeSpecifier            = _CReduceType(98)
	_CReduceHToTypeSpecifier            = _CReduceType(99)
	_CReduceIToTypeSpecifier            = _CReduceType(100)
	_CReduceJToTypeSpecifier            = _CReduceType(101)
	_CReduceKToTypeSpecifier            = _CReduceType(102)
	_CReduceLToTypeSpecifier            = _CReduceType(103)
	_CReduceAToStructOrUnionSpecifier   = _CReduceType(104)
	_CReduceBToStructOrUnionSpecifier   = _CReduceType(105)
	_CReduceCToStructOrUnionSpecifier   = _CReduceType(106)
	_CReduceAToStructOrUnion            = _CReduceType(107)
	_CReduceBToStructOrUnion            = _CReduceType(108)
	_CReduceAToStructDeclarationList    = _CReduceType(109)
	_CReduceBToStructDeclarationList    = _CReduceType(110)
	_CReduceAToStructDeclaration        = _CReduceType(111)
	_CReduceAToSpecifierQualifierList   = _CReduceType(112)
	_CReduceBToSpecifierQualifierList   = _CReduceType(113)
	_CReduceCToSpecifierQualifierList   = _CReduceType(114)
	_CReduceDToSpecifierQualifierList   = _CReduceType(115)
	_CReduceAToStructDeclaratorList     = _CReduceType(116)
	_CReduceBToStructDeclaratorList     = _CReduceType(117)
	_CReduceAToStructDeclarator         = _CReduceType(118)
	_CReduceBToStructDeclarator         = _CReduceType(119)
	_CReduceCToStructDeclarator         = _CReduceType(120)
	_CReduceAToEnumSpecifier            = _CReduceType(121)
	_CReduceBToEnumSpecifier            = _CReduceType(122)
	_CReduceCToEnumSpecifier            = _CReduceType(123)
	_CReduceAToEnumeratorList           = _CReduceType(124)
	_CReduceBToEnumeratorList           = _CReduceType(125)
	_CReduceAToEnumerator               = _CReduceType(126)
	_CReduceBToEnumerator               = _CReduceType(127)
	_CReduceAToTypeQualifier            = _CReduceType(128)
	_CReduceBToTypeQualifier            = _CReduceType(129)
	_CReduceAToDeclarator               = _CReduceType(130)
	_CReduceBToDeclarator               = _CReduceType(131)
	_CReduceAToDirectDeclarator         = _CReduceType(132)
	_CReduceBToDirectDeclarator         = _CReduceType(133)
	_CReduceCToDirectDeclarator         = _CReduceType(134)
	_CReduceDToDirectDeclarator         = _CReduceType(135)
	_CReduceEToDirectDeclarator         = _CReduceType(136)
	_CReduceFToDirectDeclarator         = _CReduceType(137)
	_CReduceGToDirectDeclarator         = _CReduceType(138)
	_CReduceAToPointer                  = _CReduceType(139)
	_CReduceBToPointer                  = _CReduceType(140)
	_CReduceCToPointer                  = _CReduceType(141)
	_CReduceDToPointer                  = _CReduceType(142)
	_CReduceAToTypeQualifierList        = _CReduceType(143)
	_CReduceBToTypeQualifierList        = _CReduceType(144)
	_CReduceAToParameterTypeList        = _CReduceType(145)
	_CReduceBToParameterTypeList        = _CReduceType(146)
	_CReduceAToParameterList            = _CReduceType(147)
	_CReduceBToParameterList            = _CReduceType(148)
	_CReduceAToParameterDeclaration     = _CReduceType(149)
	_CReduceBToParameterDeclaration     = _CReduceType(150)
	_CReduceCToParameterDeclaration     = _CReduceType(151)
	_CReduceAToIdentifierList           = _CReduceType(152)
	_CReduceBToIdentifierList           = _CReduceType(153)
	_CReduceAToTypeName                 = _CReduceType(154)
	_CReduceBToTypeName                 = _CReduceType(155)
	_CReduceAToAbstractDeclarator       = _CReduceType(156)
	_CReduceBToAbstractDeclarator       = _CReduceType(157)
	_CReduceCToAbstractDeclarator       = _CReduceType(158)
	_CReduceAToDirectAbstractDeclarator = _CReduceType(159)
	_CReduceBToDirectAbstractDeclarator = _CReduceType(160)
	_CReduceCToDirectAbstractDeclarator = _CReduceType(161)
	_CReduceDToDirectAbstractDeclarator = _CReduceType(162)
	_CReduceEToDirectAbstractDeclarator = _CReduceType(163)
	_CReduceFToDirectAbstractDeclarator = _CReduceType(164)
	_CReduceGToDirectAbstractDeclarator = _CReduceType(165)
	_CReduceHToDirectAbstractDeclarator = _CReduceType(166)
	_CReduceIToDirectAbstractDeclarator = _CReduceType(167)
	_CReduceAToInitializer              = _CReduceType(168)
	_CReduceBToInitializer              = _CReduceType(169)
	_CReduceCToInitializer              = _CReduceType(170)
	_CReduceAToInitializerList          = _CReduceType(171)
	_CReduceBToInitializerList          = _CReduceType(172)
	_CReduceAToStatement                = _CReduceType(173)
	_CReduceBToStatement                = _CReduceType(174)
	_CReduceCToStatement                = _CReduceType(175)
	_CReduceDToStatement                = _CReduceType(176)
	_CReduceEToStatement                = _CReduceType(177)
	_CReduceFToStatement                = _CReduceType(178)
	_CReduceAToLabeledStatement         = _CReduceType(179)
	_CReduceBToLabeledStatement         = _CReduceType(180)
	_CReduceCToLabeledStatement         = _CReduceType(181)
	_CReduceAToCompoundStatement        = _CReduceType(182)
	_CReduceBToCompoundStatement        = _CReduceType(183)
	_CReduceCToCompoundStatement        = _CReduceType(184)
	_CReduceDToCompoundStatement        = _CReduceType(185)
	_CReduceAToDeclarationList          = _CReduceType(186)
	_CReduceBToDeclarationList          = _CReduceType(187)
	_CReduceAToStatementList            = _CReduceType(188)
	_CReduceBToStatementList            = _CReduceType(189)
	_CReduceAToExpressionStatement      = _CReduceType(190)
	_CReduceBToExpressionStatement      = _CReduceType(191)
	_CReduceAToSelectionStatement       = _CReduceType(192)
	_CReduceBToSelectionStatement       = _CReduceType(193)
	_CReduceCToSelectionStatement       = _CReduceType(194)
	_CReduceAToIterationStatement       = _CReduceType(195)
	_CReduceBToIterationStatement       = _CReduceType(196)
	_CReduceCToIterationStatement       = _CReduceType(197)
	_CReduceDToIterationStatement       = _CReduceType(198)
	_CReduceAToJumpStatement            = _CReduceType(199)
	_CReduceBToJumpStatement            = _CReduceType(200)
	_CReduceCToJumpStatement            = _CReduceType(201)
	_CReduceDToJumpStatement            = _CReduceType(202)
	_CReduceEToJumpStatement            = _CReduceType(203)
	_CReduceAToTranslationUnit          = _CReduceType(204)
	_CReduceBToTranslationUnit          = _CReduceType(205)
	_CReduceAToExternalDeclaration      = _CReduceType(206)
	_CReduceBToExternalDeclaration      = _CReduceType(207)
	_CReduceAToFunctionDefinition       = _CReduceType(208)
	_CReduceBToFunctionDefinition       = _CReduceType(209)
	_CReduceCToFunctionDefinition       = _CReduceType(210)
	_CReduceDToFunctionDefinition       = _CReduceType(211)
)

func (i _CReduceType) String() string {
	switch i {
	case _CReduceAToPrimaryExpression:
		return "AToPrimaryExpression"
	case _CReduceBToPrimaryExpression:
		return "BToPrimaryExpression"
	case _CReduceCToPrimaryExpression:
		return "CToPrimaryExpression"
	case _CReduceDToPrimaryExpression:
		return "DToPrimaryExpression"
	case _CReduceAToPostfixExpression:
		return "AToPostfixExpression"
	case _CReduceBToPostfixExpression:
		return "BToPostfixExpression"
	case _CReduceCToPostfixExpression:
		return "CToPostfixExpression"
	case _CReduceDToPostfixExpression:
		return "DToPostfixExpression"
	case _CReduceEToPostfixExpression:
		return "EToPostfixExpression"
	case _CReduceFToPostfixExpression:
		return "FToPostfixExpression"
	case _CReduceGToPostfixExpression:
		return "GToPostfixExpression"
	case _CReduceHToPostfixExpression:
		return "HToPostfixExpression"
	case _CReduceAToArgumentExpressionList:
		return "AToArgumentExpressionList"
	case _CReduceBToArgumentExpressionList:
		return "BToArgumentExpressionList"
	case _CReduceAToUnaryExpression:
		return "AToUnaryExpression"
	case _CReduceBToUnaryExpression:
		return "BToUnaryExpression"
	case _CReduceCToUnaryExpression:
		return "CToUnaryExpression"
	case _CReduceDToUnaryExpression:
		return "DToUnaryExpression"
	case _CReduceEToUnaryExpression:
		return "EToUnaryExpression"
	case _CReduceFToUnaryExpression:
		return "FToUnaryExpression"
	case _CReduceAToUnaryOperator:
		return "AToUnaryOperator"
	case _CReduceBToUnaryOperator:
		return "BToUnaryOperator"
	case _CReduceCToUnaryOperator:
		return "CToUnaryOperator"
	case _CReduceDToUnaryOperator:
		return "DToUnaryOperator"
	case _CReduceEToUnaryOperator:
		return "EToUnaryOperator"
	case _CReduceFToUnaryOperator:
		return "FToUnaryOperator"
	case _CReduceAToCastExpression:
		return "AToCastExpression"
	case _CReduceBToCastExpression:
		return "BToCastExpression"
	case _CReduceAToMultiplicativeExpression:
		return "AToMultiplicativeExpression"
	case _CReduceBToMultiplicativeExpression:
		return "BToMultiplicativeExpression"
	case _CReduceCToMultiplicativeExpression:
		return "CToMultiplicativeExpression"
	case _CReduceDToMultiplicativeExpression:
		return "DToMultiplicativeExpression"
	case _CReduceAToAdditiveExpression:
		return "AToAdditiveExpression"
	case _CReduceBToAdditiveExpression:
		return "BToAdditiveExpression"
	case _CReduceCToAdditiveExpression:
		return "CToAdditiveExpression"
	case _CReduceAToShiftExpression:
		return "AToShiftExpression"
	case _CReduceBToShiftExpression:
		return "BToShiftExpression"
	case _CReduceCToShiftExpression:
		return "CToShiftExpression"
	case _CReduceAToRelationalExpression:
		return "AToRelationalExpression"
	case _CReduceBToRelationalExpression:
		return "BToRelationalExpression"
	case _CReduceCToRelationalExpression:
		return "CToRelationalExpression"
	case _CReduceDToRelationalExpression:
		return "DToRelationalExpression"
	case _CReduceEToRelationalExpression:
		return "EToRelationalExpression"
	case _CReduceAToEqualityExpression:
		return "AToEqualityExpression"
	case _CReduceBToEqualityExpression:
		return "BToEqualityExpression"
	case _CReduceCToEqualityExpression:
		return "CToEqualityExpression"
	case _CReduceAToAndExpression:
		return "AToAndExpression"
	case _CReduceBToAndExpression:
		return "BToAndExpression"
	case _CReduceAToExclusiveOrExpression:
		return "AToExclusiveOrExpression"
	case _CReduceBToExclusiveOrExpression:
		return "BToExclusiveOrExpression"
	case _CReduceAToInclusiveOrExpression:
		return "AToInclusiveOrExpression"
	case _CReduceBToInclusiveOrExpression:
		return "BToInclusiveOrExpression"
	case _CReduceAToLogicalAndExpression:
		return "AToLogicalAndExpression"
	case _CReduceBToLogicalAndExpression:
		return "BToLogicalAndExpression"
	case _CReduceAToLogicalOrExpression:
		return "AToLogicalOrExpression"
	case _CReduceBToLogicalOrExpression:
		return "BToLogicalOrExpression"
	case _CReduceAToConditionalExpression:
		return "AToConditionalExpression"
	case _CReduceBToConditionalExpression:
		return "BToConditionalExpression"
	case _CReduceAToAssignmentExpression:
		return "AToAssignmentExpression"
	case _CReduceBToAssignmentExpression:
		return "BToAssignmentExpression"
	case _CReduceAToAssignmentOperator:
		return "AToAssignmentOperator"
	case _CReduceBToAssignmentOperator:
		return "BToAssignmentOperator"
	case _CReduceCToAssignmentOperator:
		return "CToAssignmentOperator"
	case _CReduceDToAssignmentOperator:
		return "DToAssignmentOperator"
	case _CReduceEToAssignmentOperator:
		return "EToAssignmentOperator"
	case _CReduceFToAssignmentOperator:
		return "FToAssignmentOperator"
	case _CReduceGToAssignmentOperator:
		return "GToAssignmentOperator"
	case _CReduceHToAssignmentOperator:
		return "HToAssignmentOperator"
	case _CReduceIToAssignmentOperator:
		return "IToAssignmentOperator"
	case _CReduceJToAssignmentOperator:
		return "JToAssignmentOperator"
	case _CReduceKToAssignmentOperator:
		return "KToAssignmentOperator"
	case _CReduceAToExpression:
		return "AToExpression"
	case _CReduceBToExpression:
		return "BToExpression"
	case _CReduceAToConstantExpression:
		return "AToConstantExpression"
	case _CReduceAToDeclaration:
		return "AToDeclaration"
	case _CReduceBToDeclaration:
		return "BToDeclaration"
	case _CReduceAToDeclarationSpecifiers:
		return "AToDeclarationSpecifiers"
	case _CReduceBToDeclarationSpecifiers:
		return "BToDeclarationSpecifiers"
	case _CReduceCToDeclarationSpecifiers:
		return "CToDeclarationSpecifiers"
	case _CReduceDToDeclarationSpecifiers:
		return "DToDeclarationSpecifiers"
	case _CReduceEToDeclarationSpecifiers:
		return "EToDeclarationSpecifiers"
	case _CReduceFToDeclarationSpecifiers:
		return "FToDeclarationSpecifiers"
	case _CReduceAToInitDeclaratorList:
		return "AToInitDeclaratorList"
	case _CReduceBToInitDeclaratorList:
		return "BToInitDeclaratorList"
	case _CReduceAToInitDeclarator:
		return "AToInitDeclarator"
	case _CReduceBToInitDeclarator:
		return "BToInitDeclarator"
	case _CReduceAToStorageClassSpecifier:
		return "AToStorageClassSpecifier"
	case _CReduceBToStorageClassSpecifier:
		return "BToStorageClassSpecifier"
	case _CReduceCToStorageClassSpecifier:
		return "CToStorageClassSpecifier"
	case _CReduceDToStorageClassSpecifier:
		return "DToStorageClassSpecifier"
	case _CReduceEToStorageClassSpecifier:
		return "EToStorageClassSpecifier"
	case _CReduceAToTypeSpecifier:
		return "AToTypeSpecifier"
	case _CReduceBToTypeSpecifier:
		return "BToTypeSpecifier"
	case _CReduceCToTypeSpecifier:
		return "CToTypeSpecifier"
	case _CReduceDToTypeSpecifier:
		return "DToTypeSpecifier"
	case _CReduceEToTypeSpecifier:
		return "EToTypeSpecifier"
	case _CReduceFToTypeSpecifier:
		return "FToTypeSpecifier"
	case _CReduceGToTypeSpecifier:
		return "GToTypeSpecifier"
	case _CReduceHToTypeSpecifier:
		return "HToTypeSpecifier"
	case _CReduceIToTypeSpecifier:
		return "IToTypeSpecifier"
	case _CReduceJToTypeSpecifier:
		return "JToTypeSpecifier"
	case _CReduceKToTypeSpecifier:
		return "KToTypeSpecifier"
	case _CReduceLToTypeSpecifier:
		return "LToTypeSpecifier"
	case _CReduceAToStructOrUnionSpecifier:
		return "AToStructOrUnionSpecifier"
	case _CReduceBToStructOrUnionSpecifier:
		return "BToStructOrUnionSpecifier"
	case _CReduceCToStructOrUnionSpecifier:
		return "CToStructOrUnionSpecifier"
	case _CReduceAToStructOrUnion:
		return "AToStructOrUnion"
	case _CReduceBToStructOrUnion:
		return "BToStructOrUnion"
	case _CReduceAToStructDeclarationList:
		return "AToStructDeclarationList"
	case _CReduceBToStructDeclarationList:
		return "BToStructDeclarationList"
	case _CReduceAToStructDeclaration:
		return "AToStructDeclaration"
	case _CReduceAToSpecifierQualifierList:
		return "AToSpecifierQualifierList"
	case _CReduceBToSpecifierQualifierList:
		return "BToSpecifierQualifierList"
	case _CReduceCToSpecifierQualifierList:
		return "CToSpecifierQualifierList"
	case _CReduceDToSpecifierQualifierList:
		return "DToSpecifierQualifierList"
	case _CReduceAToStructDeclaratorList:
		return "AToStructDeclaratorList"
	case _CReduceBToStructDeclaratorList:
		return "BToStructDeclaratorList"
	case _CReduceAToStructDeclarator:
		return "AToStructDeclarator"
	case _CReduceBToStructDeclarator:
		return "BToStructDeclarator"
	case _CReduceCToStructDeclarator:
		return "CToStructDeclarator"
	case _CReduceAToEnumSpecifier:
		return "AToEnumSpecifier"
	case _CReduceBToEnumSpecifier:
		return "BToEnumSpecifier"
	case _CReduceCToEnumSpecifier:
		return "CToEnumSpecifier"
	case _CReduceAToEnumeratorList:
		return "AToEnumeratorList"
	case _CReduceBToEnumeratorList:
		return "BToEnumeratorList"
	case _CReduceAToEnumerator:
		return "AToEnumerator"
	case _CReduceBToEnumerator:
		return "BToEnumerator"
	case _CReduceAToTypeQualifier:
		return "AToTypeQualifier"
	case _CReduceBToTypeQualifier:
		return "BToTypeQualifier"
	case _CReduceAToDeclarator:
		return "AToDeclarator"
	case _CReduceBToDeclarator:
		return "BToDeclarator"
	case _CReduceAToDirectDeclarator:
		return "AToDirectDeclarator"
	case _CReduceBToDirectDeclarator:
		return "BToDirectDeclarator"
	case _CReduceCToDirectDeclarator:
		return "CToDirectDeclarator"
	case _CReduceDToDirectDeclarator:
		return "DToDirectDeclarator"
	case _CReduceEToDirectDeclarator:
		return "EToDirectDeclarator"
	case _CReduceFToDirectDeclarator:
		return "FToDirectDeclarator"
	case _CReduceGToDirectDeclarator:
		return "GToDirectDeclarator"
	case _CReduceAToPointer:
		return "AToPointer"
	case _CReduceBToPointer:
		return "BToPointer"
	case _CReduceCToPointer:
		return "CToPointer"
	case _CReduceDToPointer:
		return "DToPointer"
	case _CReduceAToTypeQualifierList:
		return "AToTypeQualifierList"
	case _CReduceBToTypeQualifierList:
		return "BToTypeQualifierList"
	case _CReduceAToParameterTypeList:
		return "AToParameterTypeList"
	case _CReduceBToParameterTypeList:
		return "BToParameterTypeList"
	case _CReduceAToParameterList:
		return "AToParameterList"
	case _CReduceBToParameterList:
		return "BToParameterList"
	case _CReduceAToParameterDeclaration:
		return "AToParameterDeclaration"
	case _CReduceBToParameterDeclaration:
		return "BToParameterDeclaration"
	case _CReduceCToParameterDeclaration:
		return "CToParameterDeclaration"
	case _CReduceAToIdentifierList:
		return "AToIdentifierList"
	case _CReduceBToIdentifierList:
		return "BToIdentifierList"
	case _CReduceAToTypeName:
		return "AToTypeName"
	case _CReduceBToTypeName:
		return "BToTypeName"
	case _CReduceAToAbstractDeclarator:
		return "AToAbstractDeclarator"
	case _CReduceBToAbstractDeclarator:
		return "BToAbstractDeclarator"
	case _CReduceCToAbstractDeclarator:
		return "CToAbstractDeclarator"
	case _CReduceAToDirectAbstractDeclarator:
		return "AToDirectAbstractDeclarator"
	case _CReduceBToDirectAbstractDeclarator:
		return "BToDirectAbstractDeclarator"
	case _CReduceCToDirectAbstractDeclarator:
		return "CToDirectAbstractDeclarator"
	case _CReduceDToDirectAbstractDeclarator:
		return "DToDirectAbstractDeclarator"
	case _CReduceEToDirectAbstractDeclarator:
		return "EToDirectAbstractDeclarator"
	case _CReduceFToDirectAbstractDeclarator:
		return "FToDirectAbstractDeclarator"
	case _CReduceGToDirectAbstractDeclarator:
		return "GToDirectAbstractDeclarator"
	case _CReduceHToDirectAbstractDeclarator:
		return "HToDirectAbstractDeclarator"
	case _CReduceIToDirectAbstractDeclarator:
		return "IToDirectAbstractDeclarator"
	case _CReduceAToInitializer:
		return "AToInitializer"
	case _CReduceBToInitializer:
		return "BToInitializer"
	case _CReduceCToInitializer:
		return "CToInitializer"
	case _CReduceAToInitializerList:
		return "AToInitializerList"
	case _CReduceBToInitializerList:
		return "BToInitializerList"
	case _CReduceAToStatement:
		return "AToStatement"
	case _CReduceBToStatement:
		return "BToStatement"
	case _CReduceCToStatement:
		return "CToStatement"
	case _CReduceDToStatement:
		return "DToStatement"
	case _CReduceEToStatement:
		return "EToStatement"
	case _CReduceFToStatement:
		return "FToStatement"
	case _CReduceAToLabeledStatement:
		return "AToLabeledStatement"
	case _CReduceBToLabeledStatement:
		return "BToLabeledStatement"
	case _CReduceCToLabeledStatement:
		return "CToLabeledStatement"
	case _CReduceAToCompoundStatement:
		return "AToCompoundStatement"
	case _CReduceBToCompoundStatement:
		return "BToCompoundStatement"
	case _CReduceCToCompoundStatement:
		return "CToCompoundStatement"
	case _CReduceDToCompoundStatement:
		return "DToCompoundStatement"
	case _CReduceAToDeclarationList:
		return "AToDeclarationList"
	case _CReduceBToDeclarationList:
		return "BToDeclarationList"
	case _CReduceAToStatementList:
		return "AToStatementList"
	case _CReduceBToStatementList:
		return "BToStatementList"
	case _CReduceAToExpressionStatement:
		return "AToExpressionStatement"
	case _CReduceBToExpressionStatement:
		return "BToExpressionStatement"
	case _CReduceAToSelectionStatement:
		return "AToSelectionStatement"
	case _CReduceBToSelectionStatement:
		return "BToSelectionStatement"
	case _CReduceCToSelectionStatement:
		return "CToSelectionStatement"
	case _CReduceAToIterationStatement:
		return "AToIterationStatement"
	case _CReduceBToIterationStatement:
		return "BToIterationStatement"
	case _CReduceCToIterationStatement:
		return "CToIterationStatement"
	case _CReduceDToIterationStatement:
		return "DToIterationStatement"
	case _CReduceAToJumpStatement:
		return "AToJumpStatement"
	case _CReduceBToJumpStatement:
		return "BToJumpStatement"
	case _CReduceCToJumpStatement:
		return "CToJumpStatement"
	case _CReduceDToJumpStatement:
		return "DToJumpStatement"
	case _CReduceEToJumpStatement:
		return "EToJumpStatement"
	case _CReduceAToTranslationUnit:
		return "AToTranslationUnit"
	case _CReduceBToTranslationUnit:
		return "BToTranslationUnit"
	case _CReduceAToExternalDeclaration:
		return "AToExternalDeclaration"
	case _CReduceBToExternalDeclaration:
		return "BToExternalDeclaration"
	case _CReduceAToFunctionDefinition:
		return "AToFunctionDefinition"
	case _CReduceBToFunctionDefinition:
		return "BToFunctionDefinition"
	case _CReduceCToFunctionDefinition:
		return "CToFunctionDefinition"
	case _CReduceDToFunctionDefinition:
		return "DToFunctionDefinition"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _CStateId int

func (id _CStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_CState1   = _CStateId(1)
	_CState2   = _CStateId(2)
	_CState3   = _CStateId(3)
	_CState4   = _CStateId(4)
	_CState5   = _CStateId(5)
	_CState6   = _CStateId(6)
	_CState7   = _CStateId(7)
	_CState8   = _CStateId(8)
	_CState9   = _CStateId(9)
	_CState10  = _CStateId(10)
	_CState11  = _CStateId(11)
	_CState12  = _CStateId(12)
	_CState13  = _CStateId(13)
	_CState14  = _CStateId(14)
	_CState15  = _CStateId(15)
	_CState16  = _CStateId(16)
	_CState17  = _CStateId(17)
	_CState18  = _CStateId(18)
	_CState19  = _CStateId(19)
	_CState20  = _CStateId(20)
	_CState21  = _CStateId(21)
	_CState22  = _CStateId(22)
	_CState23  = _CStateId(23)
	_CState24  = _CStateId(24)
	_CState25  = _CStateId(25)
	_CState26  = _CStateId(26)
	_CState27  = _CStateId(27)
	_CState28  = _CStateId(28)
	_CState29  = _CStateId(29)
	_CState30  = _CStateId(30)
	_CState31  = _CStateId(31)
	_CState32  = _CStateId(32)
	_CState33  = _CStateId(33)
	_CState34  = _CStateId(34)
	_CState35  = _CStateId(35)
	_CState36  = _CStateId(36)
	_CState37  = _CStateId(37)
	_CState38  = _CStateId(38)
	_CState39  = _CStateId(39)
	_CState40  = _CStateId(40)
	_CState41  = _CStateId(41)
	_CState42  = _CStateId(42)
	_CState43  = _CStateId(43)
	_CState44  = _CStateId(44)
	_CState45  = _CStateId(45)
	_CState46  = _CStateId(46)
	_CState47  = _CStateId(47)
	_CState48  = _CStateId(48)
	_CState49  = _CStateId(49)
	_CState50  = _CStateId(50)
	_CState51  = _CStateId(51)
	_CState52  = _CStateId(52)
	_CState53  = _CStateId(53)
	_CState54  = _CStateId(54)
	_CState55  = _CStateId(55)
	_CState56  = _CStateId(56)
	_CState57  = _CStateId(57)
	_CState58  = _CStateId(58)
	_CState59  = _CStateId(59)
	_CState60  = _CStateId(60)
	_CState61  = _CStateId(61)
	_CState62  = _CStateId(62)
	_CState63  = _CStateId(63)
	_CState64  = _CStateId(64)
	_CState65  = _CStateId(65)
	_CState66  = _CStateId(66)
	_CState67  = _CStateId(67)
	_CState68  = _CStateId(68)
	_CState69  = _CStateId(69)
	_CState70  = _CStateId(70)
	_CState71  = _CStateId(71)
	_CState72  = _CStateId(72)
	_CState73  = _CStateId(73)
	_CState74  = _CStateId(74)
	_CState75  = _CStateId(75)
	_CState76  = _CStateId(76)
	_CState77  = _CStateId(77)
	_CState78  = _CStateId(78)
	_CState79  = _CStateId(79)
	_CState80  = _CStateId(80)
	_CState81  = _CStateId(81)
	_CState82  = _CStateId(82)
	_CState83  = _CStateId(83)
	_CState84  = _CStateId(84)
	_CState85  = _CStateId(85)
	_CState86  = _CStateId(86)
	_CState87  = _CStateId(87)
	_CState88  = _CStateId(88)
	_CState89  = _CStateId(89)
	_CState90  = _CStateId(90)
	_CState91  = _CStateId(91)
	_CState92  = _CStateId(92)
	_CState93  = _CStateId(93)
	_CState94  = _CStateId(94)
	_CState95  = _CStateId(95)
	_CState96  = _CStateId(96)
	_CState97  = _CStateId(97)
	_CState98  = _CStateId(98)
	_CState99  = _CStateId(99)
	_CState100 = _CStateId(100)
	_CState101 = _CStateId(101)
	_CState102 = _CStateId(102)
	_CState103 = _CStateId(103)
	_CState104 = _CStateId(104)
	_CState105 = _CStateId(105)
	_CState106 = _CStateId(106)
	_CState107 = _CStateId(107)
	_CState108 = _CStateId(108)
	_CState109 = _CStateId(109)
	_CState110 = _CStateId(110)
	_CState111 = _CStateId(111)
	_CState112 = _CStateId(112)
	_CState113 = _CStateId(113)
	_CState114 = _CStateId(114)
	_CState115 = _CStateId(115)
	_CState116 = _CStateId(116)
	_CState117 = _CStateId(117)
	_CState118 = _CStateId(118)
	_CState119 = _CStateId(119)
	_CState120 = _CStateId(120)
	_CState121 = _CStateId(121)
	_CState122 = _CStateId(122)
	_CState123 = _CStateId(123)
	_CState124 = _CStateId(124)
	_CState125 = _CStateId(125)
	_CState126 = _CStateId(126)
	_CState127 = _CStateId(127)
	_CState128 = _CStateId(128)
	_CState129 = _CStateId(129)
	_CState130 = _CStateId(130)
	_CState131 = _CStateId(131)
	_CState132 = _CStateId(132)
	_CState133 = _CStateId(133)
	_CState134 = _CStateId(134)
	_CState135 = _CStateId(135)
	_CState136 = _CStateId(136)
	_CState137 = _CStateId(137)
	_CState138 = _CStateId(138)
	_CState139 = _CStateId(139)
	_CState140 = _CStateId(140)
	_CState141 = _CStateId(141)
	_CState142 = _CStateId(142)
	_CState143 = _CStateId(143)
	_CState144 = _CStateId(144)
	_CState145 = _CStateId(145)
	_CState146 = _CStateId(146)
	_CState147 = _CStateId(147)
	_CState148 = _CStateId(148)
	_CState149 = _CStateId(149)
	_CState150 = _CStateId(150)
	_CState151 = _CStateId(151)
	_CState152 = _CStateId(152)
	_CState153 = _CStateId(153)
	_CState154 = _CStateId(154)
	_CState155 = _CStateId(155)
	_CState156 = _CStateId(156)
	_CState157 = _CStateId(157)
	_CState158 = _CStateId(158)
	_CState159 = _CStateId(159)
	_CState160 = _CStateId(160)
	_CState161 = _CStateId(161)
	_CState162 = _CStateId(162)
	_CState163 = _CStateId(163)
	_CState164 = _CStateId(164)
	_CState165 = _CStateId(165)
	_CState166 = _CStateId(166)
	_CState167 = _CStateId(167)
	_CState168 = _CStateId(168)
	_CState169 = _CStateId(169)
	_CState170 = _CStateId(170)
	_CState171 = _CStateId(171)
	_CState172 = _CStateId(172)
	_CState173 = _CStateId(173)
	_CState174 = _CStateId(174)
	_CState175 = _CStateId(175)
	_CState176 = _CStateId(176)
	_CState177 = _CStateId(177)
	_CState178 = _CStateId(178)
	_CState179 = _CStateId(179)
	_CState180 = _CStateId(180)
	_CState181 = _CStateId(181)
	_CState182 = _CStateId(182)
	_CState183 = _CStateId(183)
	_CState184 = _CStateId(184)
	_CState185 = _CStateId(185)
	_CState186 = _CStateId(186)
	_CState187 = _CStateId(187)
	_CState188 = _CStateId(188)
	_CState189 = _CStateId(189)
	_CState190 = _CStateId(190)
	_CState191 = _CStateId(191)
	_CState192 = _CStateId(192)
	_CState193 = _CStateId(193)
	_CState194 = _CStateId(194)
	_CState195 = _CStateId(195)
	_CState196 = _CStateId(196)
	_CState197 = _CStateId(197)
	_CState198 = _CStateId(198)
	_CState199 = _CStateId(199)
	_CState200 = _CStateId(200)
	_CState201 = _CStateId(201)
	_CState202 = _CStateId(202)
	_CState203 = _CStateId(203)
	_CState204 = _CStateId(204)
	_CState205 = _CStateId(205)
	_CState206 = _CStateId(206)
	_CState207 = _CStateId(207)
	_CState208 = _CStateId(208)
	_CState209 = _CStateId(209)
	_CState210 = _CStateId(210)
	_CState211 = _CStateId(211)
	_CState212 = _CStateId(212)
	_CState213 = _CStateId(213)
	_CState214 = _CStateId(214)
	_CState215 = _CStateId(215)
	_CState216 = _CStateId(216)
	_CState217 = _CStateId(217)
	_CState218 = _CStateId(218)
	_CState219 = _CStateId(219)
	_CState220 = _CStateId(220)
	_CState221 = _CStateId(221)
	_CState222 = _CStateId(222)
	_CState223 = _CStateId(223)
	_CState224 = _CStateId(224)
	_CState225 = _CStateId(225)
	_CState226 = _CStateId(226)
	_CState227 = _CStateId(227)
	_CState228 = _CStateId(228)
	_CState229 = _CStateId(229)
	_CState230 = _CStateId(230)
	_CState231 = _CStateId(231)
	_CState232 = _CStateId(232)
	_CState233 = _CStateId(233)
	_CState234 = _CStateId(234)
	_CState235 = _CStateId(235)
	_CState236 = _CStateId(236)
	_CState237 = _CStateId(237)
	_CState238 = _CStateId(238)
	_CState239 = _CStateId(239)
	_CState240 = _CStateId(240)
	_CState241 = _CStateId(241)
	_CState242 = _CStateId(242)
	_CState243 = _CStateId(243)
	_CState244 = _CStateId(244)
	_CState245 = _CStateId(245)
	_CState246 = _CStateId(246)
	_CState247 = _CStateId(247)
	_CState248 = _CStateId(248)
	_CState249 = _CStateId(249)
	_CState250 = _CStateId(250)
	_CState251 = _CStateId(251)
	_CState252 = _CStateId(252)
	_CState253 = _CStateId(253)
	_CState254 = _CStateId(254)
	_CState255 = _CStateId(255)
	_CState256 = _CStateId(256)
	_CState257 = _CStateId(257)
	_CState258 = _CStateId(258)
	_CState259 = _CStateId(259)
	_CState260 = _CStateId(260)
	_CState261 = _CStateId(261)
	_CState262 = _CStateId(262)
	_CState263 = _CStateId(263)
	_CState264 = _CStateId(264)
	_CState265 = _CStateId(265)
	_CState266 = _CStateId(266)
	_CState267 = _CStateId(267)
	_CState268 = _CStateId(268)
	_CState269 = _CStateId(269)
	_CState270 = _CStateId(270)
	_CState271 = _CStateId(271)
	_CState272 = _CStateId(272)
	_CState273 = _CStateId(273)
	_CState274 = _CStateId(274)
	_CState275 = _CStateId(275)
	_CState276 = _CStateId(276)
	_CState277 = _CStateId(277)
	_CState278 = _CStateId(278)
	_CState279 = _CStateId(279)
	_CState280 = _CStateId(280)
	_CState281 = _CStateId(281)
	_CState282 = _CStateId(282)
	_CState283 = _CStateId(283)
	_CState284 = _CStateId(284)
	_CState285 = _CStateId(285)
	_CState286 = _CStateId(286)
	_CState287 = _CStateId(287)
	_CState288 = _CStateId(288)
	_CState289 = _CStateId(289)
	_CState290 = _CStateId(290)
	_CState291 = _CStateId(291)
	_CState292 = _CStateId(292)
	_CState293 = _CStateId(293)
	_CState294 = _CStateId(294)
	_CState295 = _CStateId(295)
	_CState296 = _CStateId(296)
	_CState297 = _CStateId(297)
	_CState298 = _CStateId(298)
	_CState299 = _CStateId(299)
	_CState300 = _CStateId(300)
	_CState301 = _CStateId(301)
	_CState302 = _CStateId(302)
	_CState303 = _CStateId(303)
	_CState304 = _CStateId(304)
	_CState305 = _CStateId(305)
	_CState306 = _CStateId(306)
	_CState307 = _CStateId(307)
	_CState308 = _CStateId(308)
	_CState309 = _CStateId(309)
	_CState310 = _CStateId(310)
	_CState311 = _CStateId(311)
	_CState312 = _CStateId(312)
	_CState313 = _CStateId(313)
	_CState314 = _CStateId(314)
	_CState315 = _CStateId(315)
	_CState316 = _CStateId(316)
	_CState317 = _CStateId(317)
	_CState318 = _CStateId(318)
	_CState319 = _CStateId(319)
	_CState320 = _CStateId(320)
	_CState321 = _CStateId(321)
	_CState322 = _CStateId(322)
	_CState323 = _CStateId(323)
	_CState324 = _CStateId(324)
	_CState325 = _CStateId(325)
	_CState326 = _CStateId(326)
	_CState327 = _CStateId(327)
	_CState328 = _CStateId(328)
	_CState329 = _CStateId(329)
	_CState330 = _CStateId(330)
	_CState331 = _CStateId(331)
	_CState332 = _CStateId(332)
	_CState333 = _CStateId(333)
	_CState334 = _CStateId(334)
	_CState335 = _CStateId(335)
	_CState336 = _CStateId(336)
	_CState337 = _CStateId(337)
	_CState338 = _CStateId(338)
	_CState339 = _CStateId(339)
	_CState340 = _CStateId(340)
	_CState341 = _CStateId(341)
	_CState342 = _CStateId(342)
	_CState343 = _CStateId(343)
	_CState344 = _CStateId(344)
	_CState345 = _CStateId(345)
	_CState346 = _CStateId(346)
	_CState347 = _CStateId(347)
	_CState348 = _CStateId(348)
	_CState349 = _CStateId(349)
	_CState350 = _CStateId(350)
	_CState351 = _CStateId(351)
)

type CSymbol struct {
	SymbolId_ CSymbolId

	Generic_ *CGenericSymbol

	T Symbol
}

func NewSymbol(token CToken) (*CSymbol, error) {
	symbol, ok := token.(*CSymbol)
	if ok {
		return symbol, nil
	}

	symbol = &CSymbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _CEndMarker:
		val, ok := token.(*CGenericSymbol)
		if !ok {
			return nil, fmt.Errorf("Invalid value type for token %s.  Expecting *CGenericSymbol (%v)", token.Id(), token.Loc())
		}
		symbol.Generic_ = val
	case CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CPtrOpToken, CIncOpToken, CDecOpToken, CLeftOpToken, CRightOpToken, CLeOpToken, CGeOpToken, CEqOpToken, CNeOpToken, CAndOpToken, COrOpToken, CMulAssignToken, CDivAssignToken, CModAssignToken, CAddAssignToken, CSubAssignToken, CLeftAssignToken, CRightAssignToken, CAndAssignToken, CXorAssignToken, COrAssignToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CEllipsisToken, CCaseToken, CDefaultToken, CIfToken, CElseToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CRparamToken, CLcurlToken, CRcurlToken, CLbraceToken, CRbraceToken, CSemicolonToken, CColonToken, CCommaToken, CEqToken, CQuestionToken, CMulToken, CDivToken, CMinusToken, CPlusToken, CModToken, CAndToken, COrToken, CExclaimToken, CDotToken, CHatToken, CLtToken, CGtToken, CTildaToken:
		val, ok := token.(Symbol)
		if !ok {
			return nil, fmt.Errorf("Invalid value type for token %s.  Expecting Symbol (%v)", token.Id(), token.Loc())
		}
		symbol.T = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *CSymbol) Id() CSymbolId {
	return s.SymbolId_
}

func (s *CSymbol) Loc() CLocation {
	type locator interface{ Loc() CLocation }
	switch s.SymbolId_ {
	case CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CPtrOpToken, CIncOpToken, CDecOpToken, CLeftOpToken, CRightOpToken, CLeOpToken, CGeOpToken, CEqOpToken, CNeOpToken, CAndOpToken, COrOpToken, CMulAssignToken, CDivAssignToken, CModAssignToken, CAddAssignToken, CSubAssignToken, CLeftAssignToken, CRightAssignToken, CAndAssignToken, CXorAssignToken, COrAssignToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CEllipsisToken, CCaseToken, CDefaultToken, CIfToken, CElseToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CRparamToken, CLcurlToken, CRcurlToken, CLbraceToken, CRbraceToken, CSemicolonToken, CColonToken, CCommaToken, CEqToken, CQuestionToken, CMulToken, CDivToken, CMinusToken, CPlusToken, CModToken, CAndToken, COrToken, CExclaimToken, CDotToken, CHatToken, CLtToken, CGtToken, CTildaToken, CPrimaryExpressionType, CPostfixExpressionType, CArgumentExpressionListType, CUnaryExpressionType, CUnaryOperatorType, CCastExpressionType, CMultiplicativeExpressionType, CAdditiveExpressionType, CShiftExpressionType, CRelationalExpressionType, CEqualityExpressionType, CAndExpressionType, CExclusiveOrExpressionType, CInclusiveOrExpressionType, CLogicalAndExpressionType, CLogicalOrExpressionType, CConditionalExpressionType, CAssignmentExpressionType, CAssignmentOperatorType, CExpressionType, CConstantExpressionType, CDeclarationType, CDeclarationSpecifiersType, CInitDeclaratorListType, CInitDeclaratorType, CStorageClassSpecifierType, CTypeSpecifierType, CStructOrUnionSpecifierType, CStructOrUnionType, CStructDeclarationListType, CStructDeclarationType, CSpecifierQualifierListType, CStructDeclaratorListType, CStructDeclaratorType, CEnumSpecifierType, CEnumeratorListType, CEnumeratorType, CTypeQualifierType, CDeclaratorType, CDirectDeclaratorType, CPointerType, CTypeQualifierListType, CParameterTypeListType, CParameterListType, CParameterDeclarationType, CIdentifierListType, CTypeNameType, CAbstractDeclaratorType, CDirectAbstractDeclaratorType, CInitializerType, CInitializerListType, CStatementType, CLabeledStatementType, CCompoundStatementType, CDeclarationListType, CStatementListType, CExpressionStatementType, CSelectionStatementType, CIterationStatementType, CJumpStatementType, CTranslationUnitType, CExternalDeclarationType, CFunctionDefinitionType:
		loc, ok := interface{}(s.T).(locator)
		if ok {
			return loc.Loc()
		}
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return CLocation{}
}

type _CPseudoSymbolStack struct {
	lexer CLexer
	top   []*CSymbol
}

func (stack *_CPseudoSymbolStack) Top() (*CSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &CGenericSymbol{_CEndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_CPseudoSymbolStack) Push(symbol *CSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_CPseudoSymbolStack) Pop() (CToken, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _CStackItem struct {
	StateId _CStateId

	*CSymbol
}

type _CStack []*_CStackItem

type _CAction struct {
	ActionType _CActionType

	ShiftStateId _CStateId
	ReduceType   _CReduceType
}

func (act *_CAction) ShiftItem(symbol *CSymbol) *_CStackItem {
	return &_CStackItem{StateId: act.ShiftStateId, CSymbol: symbol}
}

func (act *_CAction) ReduceSymbol(reducer CReducer, stack _CStack) (_CStack, *CSymbol, error) {
	var err error
	symbol := &CSymbol{}
	switch act.ReduceType {
	case _CReduceAToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.T, err = reducer.AToPrimaryExpression(args[0].T)
	case _CReduceBToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.T, err = reducer.BToPrimaryExpression(args[0].T)
	case _CReduceCToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.T, err = reducer.CToPrimaryExpression(args[0].T)
	case _CReduceDToPrimaryExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPrimaryExpressionType
		symbol.T, err = reducer.DToPrimaryExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToPostfixExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.AToPostfixExpression(args[0].T)
	case _CReduceBToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.BToPostfixExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.CToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.DToPostfixExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceEToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.EToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceFToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.FToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceGToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.GToPostfixExpression(args[0].T, args[1].T)
	case _CReduceHToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPostfixExpressionType
		symbol.T, err = reducer.HToPostfixExpression(args[0].T, args[1].T)
	case _CReduceAToArgumentExpressionList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CArgumentExpressionListType
		symbol.T, err = reducer.AToArgumentExpressionList(args[0].T)
	case _CReduceBToArgumentExpressionList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CArgumentExpressionListType
		symbol.T, err = reducer.BToArgumentExpressionList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToUnaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.AToUnaryExpression(args[0].T)
	case _CReduceBToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.BToUnaryExpression(args[0].T, args[1].T)
	case _CReduceCToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.CToUnaryExpression(args[0].T, args[1].T)
	case _CReduceDToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.DToUnaryExpression(args[0].T, args[1].T)
	case _CReduceEToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.EToUnaryExpression(args[0].T, args[1].T)
	case _CReduceFToUnaryExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CUnaryExpressionType
		symbol.T, err = reducer.FToUnaryExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.AToUnaryOperator(args[0].T)
	case _CReduceBToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.BToUnaryOperator(args[0].T)
	case _CReduceCToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.CToUnaryOperator(args[0].T)
	case _CReduceDToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.DToUnaryOperator(args[0].T)
	case _CReduceEToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.EToUnaryOperator(args[0].T)
	case _CReduceFToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CUnaryOperatorType
		symbol.T, err = reducer.FToUnaryOperator(args[0].T)
	case _CReduceAToCastExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CCastExpressionType
		symbol.T, err = reducer.AToCastExpression(args[0].T)
	case _CReduceBToCastExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CCastExpressionType
		symbol.T, err = reducer.BToCastExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToMultiplicativeExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.T, err = reducer.AToMultiplicativeExpression(args[0].T)
	case _CReduceBToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.T, err = reducer.BToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.T, err = reducer.CToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CMultiplicativeExpressionType
		symbol.T, err = reducer.DToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAdditiveExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.T, err = reducer.AToAdditiveExpression(args[0].T)
	case _CReduceBToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.T, err = reducer.BToAdditiveExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAdditiveExpressionType
		symbol.T, err = reducer.CToAdditiveExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToShiftExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.T, err = reducer.AToShiftExpression(args[0].T)
	case _CReduceBToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.T, err = reducer.BToShiftExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CShiftExpressionType
		symbol.T, err = reducer.CToShiftExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToRelationalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.T, err = reducer.AToRelationalExpression(args[0].T)
	case _CReduceBToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.T, err = reducer.BToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.T, err = reducer.CToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.T, err = reducer.DToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceEToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CRelationalExpressionType
		symbol.T, err = reducer.EToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEqualityExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.T, err = reducer.AToEqualityExpression(args[0].T)
	case _CReduceBToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.T, err = reducer.BToEqualityExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEqualityExpressionType
		symbol.T, err = reducer.CToEqualityExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAndExpressionType
		symbol.T, err = reducer.AToAndExpression(args[0].T)
	case _CReduceBToAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAndExpressionType
		symbol.T, err = reducer.BToAndExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToExclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExclusiveOrExpressionType
		symbol.T, err = reducer.AToExclusiveOrExpression(args[0].T)
	case _CReduceBToExclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CExclusiveOrExpressionType
		symbol.T, err = reducer.BToExclusiveOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToInclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInclusiveOrExpressionType
		symbol.T, err = reducer.AToInclusiveOrExpression(args[0].T)
	case _CReduceBToInclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInclusiveOrExpressionType
		symbol.T, err = reducer.BToInclusiveOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToLogicalAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CLogicalAndExpressionType
		symbol.T, err = reducer.AToLogicalAndExpression(args[0].T)
	case _CReduceBToLogicalAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLogicalAndExpressionType
		symbol.T, err = reducer.BToLogicalAndExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToLogicalOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CLogicalOrExpressionType
		symbol.T, err = reducer.AToLogicalOrExpression(args[0].T)
	case _CReduceBToLogicalOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLogicalOrExpressionType
		symbol.T, err = reducer.BToLogicalOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToConditionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CConditionalExpressionType
		symbol.T, err = reducer.AToConditionalExpression(args[0].T)
	case _CReduceBToConditionalExpression:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CConditionalExpressionType
		symbol.T, err = reducer.BToConditionalExpression(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceAToAssignmentExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentExpressionType
		symbol.T, err = reducer.AToAssignmentExpression(args[0].T)
	case _CReduceBToAssignmentExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CAssignmentExpressionType
		symbol.T, err = reducer.BToAssignmentExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.AToAssignmentOperator(args[0].T)
	case _CReduceBToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.BToAssignmentOperator(args[0].T)
	case _CReduceCToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.CToAssignmentOperator(args[0].T)
	case _CReduceDToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.DToAssignmentOperator(args[0].T)
	case _CReduceEToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.EToAssignmentOperator(args[0].T)
	case _CReduceFToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.FToAssignmentOperator(args[0].T)
	case _CReduceGToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.GToAssignmentOperator(args[0].T)
	case _CReduceHToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.HToAssignmentOperator(args[0].T)
	case _CReduceIToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.IToAssignmentOperator(args[0].T)
	case _CReduceJToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.JToAssignmentOperator(args[0].T)
	case _CReduceKToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAssignmentOperatorType
		symbol.T, err = reducer.KToAssignmentOperator(args[0].T)
	case _CReduceAToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExpressionType
		symbol.T, err = reducer.AToExpression(args[0].T)
	case _CReduceBToExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CExpressionType
		symbol.T, err = reducer.BToExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToConstantExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CConstantExpressionType
		symbol.T, err = reducer.AToConstantExpression(args[0].T)
	case _CReduceAToDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationType
		symbol.T, err = reducer.AToDeclaration(args[0].T, args[1].T)
	case _CReduceBToDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDeclarationType
		symbol.T, err = reducer.BToDeclaration(args[0].T, args[1].T, args[2].T)
	case _CReduceAToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.AToDeclarationSpecifiers(args[0].T)
	case _CReduceBToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.BToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceCToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.CToDeclarationSpecifiers(args[0].T)
	case _CReduceDToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.DToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceEToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.EToDeclarationSpecifiers(args[0].T)
	case _CReduceFToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationSpecifiersType
		symbol.T, err = reducer.FToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceAToInitDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitDeclaratorListType
		symbol.T, err = reducer.AToInitDeclaratorList(args[0].T)
	case _CReduceBToInitDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitDeclaratorListType
		symbol.T, err = reducer.BToInitDeclaratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToInitDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitDeclaratorType
		symbol.T, err = reducer.AToInitDeclarator(args[0].T)
	case _CReduceBToInitDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitDeclaratorType
		symbol.T, err = reducer.BToInitDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.T, err = reducer.AToStorageClassSpecifier(args[0].T)
	case _CReduceBToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.T, err = reducer.BToStorageClassSpecifier(args[0].T)
	case _CReduceCToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.T, err = reducer.CToStorageClassSpecifier(args[0].T)
	case _CReduceDToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.T, err = reducer.DToStorageClassSpecifier(args[0].T)
	case _CReduceEToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStorageClassSpecifierType
		symbol.T, err = reducer.EToStorageClassSpecifier(args[0].T)
	case _CReduceAToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.AToTypeSpecifier(args[0].T)
	case _CReduceBToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.BToTypeSpecifier(args[0].T)
	case _CReduceCToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.CToTypeSpecifier(args[0].T)
	case _CReduceDToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.DToTypeSpecifier(args[0].T)
	case _CReduceEToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.EToTypeSpecifier(args[0].T)
	case _CReduceFToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.FToTypeSpecifier(args[0].T)
	case _CReduceGToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.GToTypeSpecifier(args[0].T)
	case _CReduceHToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.HToTypeSpecifier(args[0].T)
	case _CReduceIToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.IToTypeSpecifier(args[0].T)
	case _CReduceJToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.JToTypeSpecifier(args[0].T)
	case _CReduceKToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.KToTypeSpecifier(args[0].T)
	case _CReduceLToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeSpecifierType
		symbol.T, err = reducer.LToTypeSpecifier(args[0].T)
	case _CReduceAToStructOrUnionSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.T, err = reducer.AToStructOrUnionSpecifier(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToStructOrUnionSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.T, err = reducer.BToStructOrUnionSpecifier(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToStructOrUnionSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructOrUnionSpecifierType
		symbol.T, err = reducer.CToStructOrUnionSpecifier(args[0].T, args[1].T)
	case _CReduceAToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructOrUnionType
		symbol.T, err = reducer.AToStructOrUnion(args[0].T)
	case _CReduceBToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructOrUnionType
		symbol.T, err = reducer.BToStructOrUnion(args[0].T)
	case _CReduceAToStructDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclarationListType
		symbol.T, err = reducer.AToStructDeclarationList(args[0].T)
	case _CReduceBToStructDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructDeclarationListType
		symbol.T, err = reducer.BToStructDeclarationList(args[0].T, args[1].T)
	case _CReduceAToStructDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclarationType
		symbol.T, err = reducer.AToStructDeclaration(args[0].T, args[1].T, args[2].T)
	case _CReduceAToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.T, err = reducer.AToSpecifierQualifierList(args[0].T, args[1].T)
	case _CReduceBToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.T, err = reducer.BToSpecifierQualifierList(args[0].T)
	case _CReduceCToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.T, err = reducer.CToSpecifierQualifierList(args[0].T, args[1].T)
	case _CReduceDToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CSpecifierQualifierListType
		symbol.T, err = reducer.DToSpecifierQualifierList(args[0].T)
	case _CReduceAToStructDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclaratorListType
		symbol.T, err = reducer.AToStructDeclaratorList(args[0].T)
	case _CReduceBToStructDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclaratorListType
		symbol.T, err = reducer.BToStructDeclaratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStructDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.T, err = reducer.AToStructDeclarator(args[0].T)
	case _CReduceBToStructDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.T, err = reducer.BToStructDeclarator(args[0].T, args[1].T)
	case _CReduceCToStructDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CStructDeclaratorType
		symbol.T, err = reducer.CToStructDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEnumSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.T, err = reducer.AToEnumSpecifier(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceBToEnumSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.T, err = reducer.BToEnumSpecifier(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceCToEnumSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CEnumSpecifierType
		symbol.T, err = reducer.CToEnumSpecifier(args[0].T, args[1].T)
	case _CReduceAToEnumeratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEnumeratorListType
		symbol.T, err = reducer.AToEnumeratorList(args[0].T)
	case _CReduceBToEnumeratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEnumeratorListType
		symbol.T, err = reducer.BToEnumeratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEnumerator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CEnumeratorType
		symbol.T, err = reducer.AToEnumerator(args[0].T)
	case _CReduceBToEnumerator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CEnumeratorType
		symbol.T, err = reducer.BToEnumerator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierType
		symbol.T, err = reducer.AToTypeQualifier(args[0].T)
	case _CReduceBToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierType
		symbol.T, err = reducer.BToTypeQualifier(args[0].T)
	case _CReduceAToDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclaratorType
		symbol.T, err = reducer.AToDeclarator(args[0].T, args[1].T)
	case _CReduceBToDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclaratorType
		symbol.T, err = reducer.BToDeclarator(args[0].T)
	case _CReduceAToDirectDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.AToDirectDeclarator(args[0].T)
	case _CReduceBToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.BToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceCToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.CToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceDToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.DToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceEToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.EToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceFToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.FToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceGToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectDeclaratorType
		symbol.T, err = reducer.GToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToPointer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CPointerType
		symbol.T, err = reducer.AToPointer(args[0].T)
	case _CReduceBToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPointerType
		symbol.T, err = reducer.BToPointer(args[0].T, args[1].T)
	case _CReduceCToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CPointerType
		symbol.T, err = reducer.CToPointer(args[0].T, args[1].T)
	case _CReduceDToPointer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CPointerType
		symbol.T, err = reducer.DToPointer(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeQualifierListType
		symbol.T, err = reducer.AToTypeQualifierList(args[0].T)
	case _CReduceBToTypeQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTypeQualifierListType
		symbol.T, err = reducer.BToTypeQualifierList(args[0].T, args[1].T)
	case _CReduceAToParameterTypeList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterTypeListType
		symbol.T, err = reducer.AToParameterTypeList(args[0].T)
	case _CReduceBToParameterTypeList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CParameterTypeListType
		symbol.T, err = reducer.BToParameterTypeList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterListType
		symbol.T, err = reducer.AToParameterList(args[0].T)
	case _CReduceBToParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CParameterListType
		symbol.T, err = reducer.BToParameterList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.T, err = reducer.AToParameterDeclaration(args[0].T, args[1].T)
	case _CReduceBToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.T, err = reducer.BToParameterDeclaration(args[0].T, args[1].T)
	case _CReduceCToParameterDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CParameterDeclarationType
		symbol.T, err = reducer.CToParameterDeclaration(args[0].T)
	case _CReduceAToIdentifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CIdentifierListType
		symbol.T, err = reducer.AToIdentifierList(args[0].T)
	case _CReduceBToIdentifierList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CIdentifierListType
		symbol.T, err = reducer.BToIdentifierList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeName:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTypeNameType
		symbol.T, err = reducer.AToTypeName(args[0].T)
	case _CReduceBToTypeName:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTypeNameType
		symbol.T, err = reducer.BToTypeName(args[0].T, args[1].T)
	case _CReduceAToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.T, err = reducer.AToAbstractDeclarator(args[0].T)
	case _CReduceBToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.T, err = reducer.BToAbstractDeclarator(args[0].T)
	case _CReduceCToAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CAbstractDeclaratorType
		symbol.T, err = reducer.CToAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceAToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.AToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceBToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.BToDirectAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceCToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.CToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceDToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.DToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceEToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.EToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceFToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.FToDirectAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceGToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.GToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceHToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.HToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceIToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CDirectAbstractDeclaratorType
		symbol.T, err = reducer.IToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToInitializer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitializerType
		symbol.T, err = reducer.AToInitializer(args[0].T)
	case _CReduceBToInitializer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitializerType
		symbol.T, err = reducer.BToInitializer(args[0].T, args[1].T, args[2].T)
	case _CReduceCToInitializer:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CInitializerType
		symbol.T, err = reducer.CToInitializer(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToInitializerList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CInitializerListType
		symbol.T, err = reducer.AToInitializerList(args[0].T)
	case _CReduceBToInitializerList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CInitializerListType
		symbol.T, err = reducer.BToInitializerList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.AToStatement(args[0].T)
	case _CReduceBToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.BToStatement(args[0].T)
	case _CReduceCToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.CToStatement(args[0].T)
	case _CReduceDToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.DToStatement(args[0].T)
	case _CReduceEToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.EToStatement(args[0].T)
	case _CReduceFToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementType
		symbol.T, err = reducer.FToStatement(args[0].T)
	case _CReduceAToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.T, err = reducer.AToLabeledStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceBToLabeledStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.T, err = reducer.BToLabeledStatement(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CLabeledStatementType
		symbol.T, err = reducer.CToLabeledStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceAToCompoundStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.T, err = reducer.AToCompoundStatement(args[0].T, args[1].T)
	case _CReduceBToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.T, err = reducer.BToCompoundStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceCToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.T, err = reducer.CToCompoundStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceDToCompoundStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CCompoundStatementType
		symbol.T, err = reducer.DToCompoundStatement(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CDeclarationListType
		symbol.T, err = reducer.AToDeclarationList(args[0].T)
	case _CReduceBToDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CDeclarationListType
		symbol.T, err = reducer.BToDeclarationList(args[0].T, args[1].T)
	case _CReduceAToStatementList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CStatementListType
		symbol.T, err = reducer.AToStatementList(args[0].T)
	case _CReduceBToStatementList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CStatementListType
		symbol.T, err = reducer.BToStatementList(args[0].T, args[1].T)
	case _CReduceAToExpressionStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExpressionStatementType
		symbol.T, err = reducer.AToExpressionStatement(args[0].T)
	case _CReduceBToExpressionStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CExpressionStatementType
		symbol.T, err = reducer.BToExpressionStatement(args[0].T, args[1].T)
	case _CReduceAToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.T, err = reducer.AToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToSelectionStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.T, err = reducer.BToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceCToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CSelectionStatementType
		symbol.T, err = reducer.CToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceAToIterationStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CIterationStatementType
		symbol.T, err = reducer.AToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CIterationStatementType
		symbol.T, err = reducer.BToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceCToIterationStatement:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = CIterationStatementType
		symbol.T, err = reducer.CToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T)
	case _CReduceDToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CIterationStatementType
		symbol.T, err = reducer.DToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceAToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CJumpStatementType
		symbol.T, err = reducer.AToJumpStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceBToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.T, err = reducer.BToJumpStatement(args[0].T, args[1].T)
	case _CReduceCToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.T, err = reducer.CToJumpStatement(args[0].T, args[1].T)
	case _CReduceDToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CJumpStatementType
		symbol.T, err = reducer.DToJumpStatement(args[0].T, args[1].T)
	case _CReduceEToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CJumpStatementType
		symbol.T, err = reducer.EToJumpStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTranslationUnit:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CTranslationUnitType
		symbol.T, err = reducer.AToTranslationUnit(args[0].T)
	case _CReduceBToTranslationUnit:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CTranslationUnitType
		symbol.T, err = reducer.BToTranslationUnit(args[0].T, args[1].T)
	case _CReduceAToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExternalDeclarationType
		symbol.T, err = reducer.AToExternalDeclaration(args[0].T)
	case _CReduceBToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CExternalDeclarationType
		symbol.T, err = reducer.BToExternalDeclaration(args[0].T)
	case _CReduceAToFunctionDefinition:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.T, err = reducer.AToFunctionDefinition(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceBToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.T, err = reducer.BToFunctionDefinition(args[0].T, args[1].T, args[2].T)
	case _CReduceCToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.T, err = reducer.CToFunctionDefinition(args[0].T, args[1].T, args[2].T)
	case _CReduceDToFunctionDefinition:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CFunctionDefinitionType
		symbol.T, err = reducer.DToFunctionDefinition(args[0].T, args[1].T)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

var (
	_CGotoState1Action                        = &_CAction{_CShiftAction, _CState1, 0}
	_CGotoState2Action                        = &_CAction{_CShiftAction, _CState2, 0}
	_CGotoState3Action                        = &_CAction{_CShiftAction, _CState3, 0}
	_CGotoState4Action                        = &_CAction{_CShiftAction, _CState4, 0}
	_CGotoState5Action                        = &_CAction{_CShiftAction, _CState5, 0}
	_CGotoState6Action                        = &_CAction{_CShiftAction, _CState6, 0}
	_CGotoState7Action                        = &_CAction{_CShiftAction, _CState7, 0}
	_CGotoState8Action                        = &_CAction{_CShiftAction, _CState8, 0}
	_CGotoState9Action                        = &_CAction{_CShiftAction, _CState9, 0}
	_CGotoState10Action                       = &_CAction{_CShiftAction, _CState10, 0}
	_CGotoState11Action                       = &_CAction{_CShiftAction, _CState11, 0}
	_CGotoState12Action                       = &_CAction{_CShiftAction, _CState12, 0}
	_CGotoState13Action                       = &_CAction{_CShiftAction, _CState13, 0}
	_CGotoState14Action                       = &_CAction{_CShiftAction, _CState14, 0}
	_CGotoState15Action                       = &_CAction{_CShiftAction, _CState15, 0}
	_CGotoState16Action                       = &_CAction{_CShiftAction, _CState16, 0}
	_CGotoState17Action                       = &_CAction{_CShiftAction, _CState17, 0}
	_CGotoState18Action                       = &_CAction{_CShiftAction, _CState18, 0}
	_CGotoState19Action                       = &_CAction{_CShiftAction, _CState19, 0}
	_CGotoState20Action                       = &_CAction{_CShiftAction, _CState20, 0}
	_CGotoState21Action                       = &_CAction{_CShiftAction, _CState21, 0}
	_CGotoState22Action                       = &_CAction{_CShiftAction, _CState22, 0}
	_CGotoState23Action                       = &_CAction{_CShiftAction, _CState23, 0}
	_CGotoState24Action                       = &_CAction{_CShiftAction, _CState24, 0}
	_CGotoState25Action                       = &_CAction{_CShiftAction, _CState25, 0}
	_CGotoState26Action                       = &_CAction{_CShiftAction, _CState26, 0}
	_CGotoState27Action                       = &_CAction{_CShiftAction, _CState27, 0}
	_CGotoState28Action                       = &_CAction{_CShiftAction, _CState28, 0}
	_CGotoState29Action                       = &_CAction{_CShiftAction, _CState29, 0}
	_CGotoState30Action                       = &_CAction{_CShiftAction, _CState30, 0}
	_CGotoState31Action                       = &_CAction{_CShiftAction, _CState31, 0}
	_CGotoState32Action                       = &_CAction{_CShiftAction, _CState32, 0}
	_CGotoState33Action                       = &_CAction{_CShiftAction, _CState33, 0}
	_CGotoState34Action                       = &_CAction{_CShiftAction, _CState34, 0}
	_CGotoState35Action                       = &_CAction{_CShiftAction, _CState35, 0}
	_CGotoState36Action                       = &_CAction{_CShiftAction, _CState36, 0}
	_CGotoState37Action                       = &_CAction{_CShiftAction, _CState37, 0}
	_CGotoState38Action                       = &_CAction{_CShiftAction, _CState38, 0}
	_CGotoState39Action                       = &_CAction{_CShiftAction, _CState39, 0}
	_CGotoState40Action                       = &_CAction{_CShiftAction, _CState40, 0}
	_CGotoState41Action                       = &_CAction{_CShiftAction, _CState41, 0}
	_CGotoState42Action                       = &_CAction{_CShiftAction, _CState42, 0}
	_CGotoState43Action                       = &_CAction{_CShiftAction, _CState43, 0}
	_CGotoState44Action                       = &_CAction{_CShiftAction, _CState44, 0}
	_CGotoState45Action                       = &_CAction{_CShiftAction, _CState45, 0}
	_CGotoState46Action                       = &_CAction{_CShiftAction, _CState46, 0}
	_CGotoState47Action                       = &_CAction{_CShiftAction, _CState47, 0}
	_CGotoState48Action                       = &_CAction{_CShiftAction, _CState48, 0}
	_CGotoState49Action                       = &_CAction{_CShiftAction, _CState49, 0}
	_CGotoState50Action                       = &_CAction{_CShiftAction, _CState50, 0}
	_CGotoState51Action                       = &_CAction{_CShiftAction, _CState51, 0}
	_CGotoState52Action                       = &_CAction{_CShiftAction, _CState52, 0}
	_CGotoState53Action                       = &_CAction{_CShiftAction, _CState53, 0}
	_CGotoState54Action                       = &_CAction{_CShiftAction, _CState54, 0}
	_CGotoState55Action                       = &_CAction{_CShiftAction, _CState55, 0}
	_CGotoState56Action                       = &_CAction{_CShiftAction, _CState56, 0}
	_CGotoState57Action                       = &_CAction{_CShiftAction, _CState57, 0}
	_CGotoState58Action                       = &_CAction{_CShiftAction, _CState58, 0}
	_CGotoState59Action                       = &_CAction{_CShiftAction, _CState59, 0}
	_CGotoState60Action                       = &_CAction{_CShiftAction, _CState60, 0}
	_CGotoState61Action                       = &_CAction{_CShiftAction, _CState61, 0}
	_CGotoState62Action                       = &_CAction{_CShiftAction, _CState62, 0}
	_CGotoState63Action                       = &_CAction{_CShiftAction, _CState63, 0}
	_CGotoState64Action                       = &_CAction{_CShiftAction, _CState64, 0}
	_CGotoState65Action                       = &_CAction{_CShiftAction, _CState65, 0}
	_CGotoState66Action                       = &_CAction{_CShiftAction, _CState66, 0}
	_CGotoState67Action                       = &_CAction{_CShiftAction, _CState67, 0}
	_CGotoState68Action                       = &_CAction{_CShiftAction, _CState68, 0}
	_CGotoState69Action                       = &_CAction{_CShiftAction, _CState69, 0}
	_CGotoState70Action                       = &_CAction{_CShiftAction, _CState70, 0}
	_CGotoState71Action                       = &_CAction{_CShiftAction, _CState71, 0}
	_CGotoState72Action                       = &_CAction{_CShiftAction, _CState72, 0}
	_CGotoState73Action                       = &_CAction{_CShiftAction, _CState73, 0}
	_CGotoState74Action                       = &_CAction{_CShiftAction, _CState74, 0}
	_CGotoState75Action                       = &_CAction{_CShiftAction, _CState75, 0}
	_CGotoState76Action                       = &_CAction{_CShiftAction, _CState76, 0}
	_CGotoState77Action                       = &_CAction{_CShiftAction, _CState77, 0}
	_CGotoState78Action                       = &_CAction{_CShiftAction, _CState78, 0}
	_CGotoState79Action                       = &_CAction{_CShiftAction, _CState79, 0}
	_CGotoState80Action                       = &_CAction{_CShiftAction, _CState80, 0}
	_CGotoState81Action                       = &_CAction{_CShiftAction, _CState81, 0}
	_CGotoState82Action                       = &_CAction{_CShiftAction, _CState82, 0}
	_CGotoState83Action                       = &_CAction{_CShiftAction, _CState83, 0}
	_CGotoState84Action                       = &_CAction{_CShiftAction, _CState84, 0}
	_CGotoState85Action                       = &_CAction{_CShiftAction, _CState85, 0}
	_CGotoState86Action                       = &_CAction{_CShiftAction, _CState86, 0}
	_CGotoState87Action                       = &_CAction{_CShiftAction, _CState87, 0}
	_CGotoState88Action                       = &_CAction{_CShiftAction, _CState88, 0}
	_CGotoState89Action                       = &_CAction{_CShiftAction, _CState89, 0}
	_CGotoState90Action                       = &_CAction{_CShiftAction, _CState90, 0}
	_CGotoState91Action                       = &_CAction{_CShiftAction, _CState91, 0}
	_CGotoState92Action                       = &_CAction{_CShiftAction, _CState92, 0}
	_CGotoState93Action                       = &_CAction{_CShiftAction, _CState93, 0}
	_CGotoState94Action                       = &_CAction{_CShiftAction, _CState94, 0}
	_CGotoState95Action                       = &_CAction{_CShiftAction, _CState95, 0}
	_CGotoState96Action                       = &_CAction{_CShiftAction, _CState96, 0}
	_CGotoState97Action                       = &_CAction{_CShiftAction, _CState97, 0}
	_CGotoState98Action                       = &_CAction{_CShiftAction, _CState98, 0}
	_CGotoState99Action                       = &_CAction{_CShiftAction, _CState99, 0}
	_CGotoState100Action                      = &_CAction{_CShiftAction, _CState100, 0}
	_CGotoState101Action                      = &_CAction{_CShiftAction, _CState101, 0}
	_CGotoState102Action                      = &_CAction{_CShiftAction, _CState102, 0}
	_CGotoState103Action                      = &_CAction{_CShiftAction, _CState103, 0}
	_CGotoState104Action                      = &_CAction{_CShiftAction, _CState104, 0}
	_CGotoState105Action                      = &_CAction{_CShiftAction, _CState105, 0}
	_CGotoState106Action                      = &_CAction{_CShiftAction, _CState106, 0}
	_CGotoState107Action                      = &_CAction{_CShiftAction, _CState107, 0}
	_CGotoState108Action                      = &_CAction{_CShiftAction, _CState108, 0}
	_CGotoState109Action                      = &_CAction{_CShiftAction, _CState109, 0}
	_CGotoState110Action                      = &_CAction{_CShiftAction, _CState110, 0}
	_CGotoState111Action                      = &_CAction{_CShiftAction, _CState111, 0}
	_CGotoState112Action                      = &_CAction{_CShiftAction, _CState112, 0}
	_CGotoState113Action                      = &_CAction{_CShiftAction, _CState113, 0}
	_CGotoState114Action                      = &_CAction{_CShiftAction, _CState114, 0}
	_CGotoState115Action                      = &_CAction{_CShiftAction, _CState115, 0}
	_CGotoState116Action                      = &_CAction{_CShiftAction, _CState116, 0}
	_CGotoState117Action                      = &_CAction{_CShiftAction, _CState117, 0}
	_CGotoState118Action                      = &_CAction{_CShiftAction, _CState118, 0}
	_CGotoState119Action                      = &_CAction{_CShiftAction, _CState119, 0}
	_CGotoState120Action                      = &_CAction{_CShiftAction, _CState120, 0}
	_CGotoState121Action                      = &_CAction{_CShiftAction, _CState121, 0}
	_CGotoState122Action                      = &_CAction{_CShiftAction, _CState122, 0}
	_CGotoState123Action                      = &_CAction{_CShiftAction, _CState123, 0}
	_CGotoState124Action                      = &_CAction{_CShiftAction, _CState124, 0}
	_CGotoState125Action                      = &_CAction{_CShiftAction, _CState125, 0}
	_CGotoState126Action                      = &_CAction{_CShiftAction, _CState126, 0}
	_CGotoState127Action                      = &_CAction{_CShiftAction, _CState127, 0}
	_CGotoState128Action                      = &_CAction{_CShiftAction, _CState128, 0}
	_CGotoState129Action                      = &_CAction{_CShiftAction, _CState129, 0}
	_CGotoState130Action                      = &_CAction{_CShiftAction, _CState130, 0}
	_CGotoState131Action                      = &_CAction{_CShiftAction, _CState131, 0}
	_CGotoState132Action                      = &_CAction{_CShiftAction, _CState132, 0}
	_CGotoState133Action                      = &_CAction{_CShiftAction, _CState133, 0}
	_CGotoState134Action                      = &_CAction{_CShiftAction, _CState134, 0}
	_CGotoState135Action                      = &_CAction{_CShiftAction, _CState135, 0}
	_CGotoState136Action                      = &_CAction{_CShiftAction, _CState136, 0}
	_CGotoState137Action                      = &_CAction{_CShiftAction, _CState137, 0}
	_CGotoState138Action                      = &_CAction{_CShiftAction, _CState138, 0}
	_CGotoState139Action                      = &_CAction{_CShiftAction, _CState139, 0}
	_CGotoState140Action                      = &_CAction{_CShiftAction, _CState140, 0}
	_CGotoState141Action                      = &_CAction{_CShiftAction, _CState141, 0}
	_CGotoState142Action                      = &_CAction{_CShiftAction, _CState142, 0}
	_CGotoState143Action                      = &_CAction{_CShiftAction, _CState143, 0}
	_CGotoState144Action                      = &_CAction{_CShiftAction, _CState144, 0}
	_CGotoState145Action                      = &_CAction{_CShiftAction, _CState145, 0}
	_CGotoState146Action                      = &_CAction{_CShiftAction, _CState146, 0}
	_CGotoState147Action                      = &_CAction{_CShiftAction, _CState147, 0}
	_CGotoState148Action                      = &_CAction{_CShiftAction, _CState148, 0}
	_CGotoState149Action                      = &_CAction{_CShiftAction, _CState149, 0}
	_CGotoState150Action                      = &_CAction{_CShiftAction, _CState150, 0}
	_CGotoState151Action                      = &_CAction{_CShiftAction, _CState151, 0}
	_CGotoState152Action                      = &_CAction{_CShiftAction, _CState152, 0}
	_CGotoState153Action                      = &_CAction{_CShiftAction, _CState153, 0}
	_CGotoState154Action                      = &_CAction{_CShiftAction, _CState154, 0}
	_CGotoState155Action                      = &_CAction{_CShiftAction, _CState155, 0}
	_CGotoState156Action                      = &_CAction{_CShiftAction, _CState156, 0}
	_CGotoState157Action                      = &_CAction{_CShiftAction, _CState157, 0}
	_CGotoState158Action                      = &_CAction{_CShiftAction, _CState158, 0}
	_CGotoState159Action                      = &_CAction{_CShiftAction, _CState159, 0}
	_CGotoState160Action                      = &_CAction{_CShiftAction, _CState160, 0}
	_CGotoState161Action                      = &_CAction{_CShiftAction, _CState161, 0}
	_CGotoState162Action                      = &_CAction{_CShiftAction, _CState162, 0}
	_CGotoState163Action                      = &_CAction{_CShiftAction, _CState163, 0}
	_CGotoState164Action                      = &_CAction{_CShiftAction, _CState164, 0}
	_CGotoState165Action                      = &_CAction{_CShiftAction, _CState165, 0}
	_CGotoState166Action                      = &_CAction{_CShiftAction, _CState166, 0}
	_CGotoState167Action                      = &_CAction{_CShiftAction, _CState167, 0}
	_CGotoState168Action                      = &_CAction{_CShiftAction, _CState168, 0}
	_CGotoState169Action                      = &_CAction{_CShiftAction, _CState169, 0}
	_CGotoState170Action                      = &_CAction{_CShiftAction, _CState170, 0}
	_CGotoState171Action                      = &_CAction{_CShiftAction, _CState171, 0}
	_CGotoState172Action                      = &_CAction{_CShiftAction, _CState172, 0}
	_CGotoState173Action                      = &_CAction{_CShiftAction, _CState173, 0}
	_CGotoState174Action                      = &_CAction{_CShiftAction, _CState174, 0}
	_CGotoState175Action                      = &_CAction{_CShiftAction, _CState175, 0}
	_CGotoState176Action                      = &_CAction{_CShiftAction, _CState176, 0}
	_CGotoState177Action                      = &_CAction{_CShiftAction, _CState177, 0}
	_CGotoState178Action                      = &_CAction{_CShiftAction, _CState178, 0}
	_CGotoState179Action                      = &_CAction{_CShiftAction, _CState179, 0}
	_CGotoState180Action                      = &_CAction{_CShiftAction, _CState180, 0}
	_CGotoState181Action                      = &_CAction{_CShiftAction, _CState181, 0}
	_CGotoState182Action                      = &_CAction{_CShiftAction, _CState182, 0}
	_CGotoState183Action                      = &_CAction{_CShiftAction, _CState183, 0}
	_CGotoState184Action                      = &_CAction{_CShiftAction, _CState184, 0}
	_CGotoState185Action                      = &_CAction{_CShiftAction, _CState185, 0}
	_CGotoState186Action                      = &_CAction{_CShiftAction, _CState186, 0}
	_CGotoState187Action                      = &_CAction{_CShiftAction, _CState187, 0}
	_CGotoState188Action                      = &_CAction{_CShiftAction, _CState188, 0}
	_CGotoState189Action                      = &_CAction{_CShiftAction, _CState189, 0}
	_CGotoState190Action                      = &_CAction{_CShiftAction, _CState190, 0}
	_CGotoState191Action                      = &_CAction{_CShiftAction, _CState191, 0}
	_CGotoState192Action                      = &_CAction{_CShiftAction, _CState192, 0}
	_CGotoState193Action                      = &_CAction{_CShiftAction, _CState193, 0}
	_CGotoState194Action                      = &_CAction{_CShiftAction, _CState194, 0}
	_CGotoState195Action                      = &_CAction{_CShiftAction, _CState195, 0}
	_CGotoState196Action                      = &_CAction{_CShiftAction, _CState196, 0}
	_CGotoState197Action                      = &_CAction{_CShiftAction, _CState197, 0}
	_CGotoState198Action                      = &_CAction{_CShiftAction, _CState198, 0}
	_CGotoState199Action                      = &_CAction{_CShiftAction, _CState199, 0}
	_CGotoState200Action                      = &_CAction{_CShiftAction, _CState200, 0}
	_CGotoState201Action                      = &_CAction{_CShiftAction, _CState201, 0}
	_CGotoState202Action                      = &_CAction{_CShiftAction, _CState202, 0}
	_CGotoState203Action                      = &_CAction{_CShiftAction, _CState203, 0}
	_CGotoState204Action                      = &_CAction{_CShiftAction, _CState204, 0}
	_CGotoState205Action                      = &_CAction{_CShiftAction, _CState205, 0}
	_CGotoState206Action                      = &_CAction{_CShiftAction, _CState206, 0}
	_CGotoState207Action                      = &_CAction{_CShiftAction, _CState207, 0}
	_CGotoState208Action                      = &_CAction{_CShiftAction, _CState208, 0}
	_CGotoState209Action                      = &_CAction{_CShiftAction, _CState209, 0}
	_CGotoState210Action                      = &_CAction{_CShiftAction, _CState210, 0}
	_CGotoState211Action                      = &_CAction{_CShiftAction, _CState211, 0}
	_CGotoState212Action                      = &_CAction{_CShiftAction, _CState212, 0}
	_CGotoState213Action                      = &_CAction{_CShiftAction, _CState213, 0}
	_CGotoState214Action                      = &_CAction{_CShiftAction, _CState214, 0}
	_CGotoState215Action                      = &_CAction{_CShiftAction, _CState215, 0}
	_CGotoState216Action                      = &_CAction{_CShiftAction, _CState216, 0}
	_CGotoState217Action                      = &_CAction{_CShiftAction, _CState217, 0}
	_CGotoState218Action                      = &_CAction{_CShiftAction, _CState218, 0}
	_CGotoState219Action                      = &_CAction{_CShiftAction, _CState219, 0}
	_CGotoState220Action                      = &_CAction{_CShiftAction, _CState220, 0}
	_CGotoState221Action                      = &_CAction{_CShiftAction, _CState221, 0}
	_CGotoState222Action                      = &_CAction{_CShiftAction, _CState222, 0}
	_CGotoState223Action                      = &_CAction{_CShiftAction, _CState223, 0}
	_CGotoState224Action                      = &_CAction{_CShiftAction, _CState224, 0}
	_CGotoState225Action                      = &_CAction{_CShiftAction, _CState225, 0}
	_CGotoState226Action                      = &_CAction{_CShiftAction, _CState226, 0}
	_CGotoState227Action                      = &_CAction{_CShiftAction, _CState227, 0}
	_CGotoState228Action                      = &_CAction{_CShiftAction, _CState228, 0}
	_CGotoState229Action                      = &_CAction{_CShiftAction, _CState229, 0}
	_CGotoState230Action                      = &_CAction{_CShiftAction, _CState230, 0}
	_CGotoState231Action                      = &_CAction{_CShiftAction, _CState231, 0}
	_CGotoState232Action                      = &_CAction{_CShiftAction, _CState232, 0}
	_CGotoState233Action                      = &_CAction{_CShiftAction, _CState233, 0}
	_CGotoState234Action                      = &_CAction{_CShiftAction, _CState234, 0}
	_CGotoState235Action                      = &_CAction{_CShiftAction, _CState235, 0}
	_CGotoState236Action                      = &_CAction{_CShiftAction, _CState236, 0}
	_CGotoState237Action                      = &_CAction{_CShiftAction, _CState237, 0}
	_CGotoState238Action                      = &_CAction{_CShiftAction, _CState238, 0}
	_CGotoState239Action                      = &_CAction{_CShiftAction, _CState239, 0}
	_CGotoState240Action                      = &_CAction{_CShiftAction, _CState240, 0}
	_CGotoState241Action                      = &_CAction{_CShiftAction, _CState241, 0}
	_CGotoState242Action                      = &_CAction{_CShiftAction, _CState242, 0}
	_CGotoState243Action                      = &_CAction{_CShiftAction, _CState243, 0}
	_CGotoState244Action                      = &_CAction{_CShiftAction, _CState244, 0}
	_CGotoState245Action                      = &_CAction{_CShiftAction, _CState245, 0}
	_CGotoState246Action                      = &_CAction{_CShiftAction, _CState246, 0}
	_CGotoState247Action                      = &_CAction{_CShiftAction, _CState247, 0}
	_CGotoState248Action                      = &_CAction{_CShiftAction, _CState248, 0}
	_CGotoState249Action                      = &_CAction{_CShiftAction, _CState249, 0}
	_CGotoState250Action                      = &_CAction{_CShiftAction, _CState250, 0}
	_CGotoState251Action                      = &_CAction{_CShiftAction, _CState251, 0}
	_CGotoState252Action                      = &_CAction{_CShiftAction, _CState252, 0}
	_CGotoState253Action                      = &_CAction{_CShiftAction, _CState253, 0}
	_CGotoState254Action                      = &_CAction{_CShiftAction, _CState254, 0}
	_CGotoState255Action                      = &_CAction{_CShiftAction, _CState255, 0}
	_CGotoState256Action                      = &_CAction{_CShiftAction, _CState256, 0}
	_CGotoState257Action                      = &_CAction{_CShiftAction, _CState257, 0}
	_CGotoState258Action                      = &_CAction{_CShiftAction, _CState258, 0}
	_CGotoState259Action                      = &_CAction{_CShiftAction, _CState259, 0}
	_CGotoState260Action                      = &_CAction{_CShiftAction, _CState260, 0}
	_CGotoState261Action                      = &_CAction{_CShiftAction, _CState261, 0}
	_CGotoState262Action                      = &_CAction{_CShiftAction, _CState262, 0}
	_CGotoState263Action                      = &_CAction{_CShiftAction, _CState263, 0}
	_CGotoState264Action                      = &_CAction{_CShiftAction, _CState264, 0}
	_CGotoState265Action                      = &_CAction{_CShiftAction, _CState265, 0}
	_CGotoState266Action                      = &_CAction{_CShiftAction, _CState266, 0}
	_CGotoState267Action                      = &_CAction{_CShiftAction, _CState267, 0}
	_CGotoState268Action                      = &_CAction{_CShiftAction, _CState268, 0}
	_CGotoState269Action                      = &_CAction{_CShiftAction, _CState269, 0}
	_CGotoState270Action                      = &_CAction{_CShiftAction, _CState270, 0}
	_CGotoState271Action                      = &_CAction{_CShiftAction, _CState271, 0}
	_CGotoState272Action                      = &_CAction{_CShiftAction, _CState272, 0}
	_CGotoState273Action                      = &_CAction{_CShiftAction, _CState273, 0}
	_CGotoState274Action                      = &_CAction{_CShiftAction, _CState274, 0}
	_CGotoState275Action                      = &_CAction{_CShiftAction, _CState275, 0}
	_CGotoState276Action                      = &_CAction{_CShiftAction, _CState276, 0}
	_CGotoState277Action                      = &_CAction{_CShiftAction, _CState277, 0}
	_CGotoState278Action                      = &_CAction{_CShiftAction, _CState278, 0}
	_CGotoState279Action                      = &_CAction{_CShiftAction, _CState279, 0}
	_CGotoState280Action                      = &_CAction{_CShiftAction, _CState280, 0}
	_CGotoState281Action                      = &_CAction{_CShiftAction, _CState281, 0}
	_CGotoState282Action                      = &_CAction{_CShiftAction, _CState282, 0}
	_CGotoState283Action                      = &_CAction{_CShiftAction, _CState283, 0}
	_CGotoState284Action                      = &_CAction{_CShiftAction, _CState284, 0}
	_CGotoState285Action                      = &_CAction{_CShiftAction, _CState285, 0}
	_CGotoState286Action                      = &_CAction{_CShiftAction, _CState286, 0}
	_CGotoState287Action                      = &_CAction{_CShiftAction, _CState287, 0}
	_CGotoState288Action                      = &_CAction{_CShiftAction, _CState288, 0}
	_CGotoState289Action                      = &_CAction{_CShiftAction, _CState289, 0}
	_CGotoState290Action                      = &_CAction{_CShiftAction, _CState290, 0}
	_CGotoState291Action                      = &_CAction{_CShiftAction, _CState291, 0}
	_CGotoState292Action                      = &_CAction{_CShiftAction, _CState292, 0}
	_CGotoState293Action                      = &_CAction{_CShiftAction, _CState293, 0}
	_CGotoState294Action                      = &_CAction{_CShiftAction, _CState294, 0}
	_CGotoState295Action                      = &_CAction{_CShiftAction, _CState295, 0}
	_CGotoState296Action                      = &_CAction{_CShiftAction, _CState296, 0}
	_CGotoState297Action                      = &_CAction{_CShiftAction, _CState297, 0}
	_CGotoState298Action                      = &_CAction{_CShiftAction, _CState298, 0}
	_CGotoState299Action                      = &_CAction{_CShiftAction, _CState299, 0}
	_CGotoState300Action                      = &_CAction{_CShiftAction, _CState300, 0}
	_CGotoState301Action                      = &_CAction{_CShiftAction, _CState301, 0}
	_CGotoState302Action                      = &_CAction{_CShiftAction, _CState302, 0}
	_CGotoState303Action                      = &_CAction{_CShiftAction, _CState303, 0}
	_CGotoState304Action                      = &_CAction{_CShiftAction, _CState304, 0}
	_CGotoState305Action                      = &_CAction{_CShiftAction, _CState305, 0}
	_CGotoState306Action                      = &_CAction{_CShiftAction, _CState306, 0}
	_CGotoState307Action                      = &_CAction{_CShiftAction, _CState307, 0}
	_CGotoState308Action                      = &_CAction{_CShiftAction, _CState308, 0}
	_CGotoState309Action                      = &_CAction{_CShiftAction, _CState309, 0}
	_CGotoState310Action                      = &_CAction{_CShiftAction, _CState310, 0}
	_CGotoState311Action                      = &_CAction{_CShiftAction, _CState311, 0}
	_CGotoState312Action                      = &_CAction{_CShiftAction, _CState312, 0}
	_CGotoState313Action                      = &_CAction{_CShiftAction, _CState313, 0}
	_CGotoState314Action                      = &_CAction{_CShiftAction, _CState314, 0}
	_CGotoState315Action                      = &_CAction{_CShiftAction, _CState315, 0}
	_CGotoState316Action                      = &_CAction{_CShiftAction, _CState316, 0}
	_CGotoState317Action                      = &_CAction{_CShiftAction, _CState317, 0}
	_CGotoState318Action                      = &_CAction{_CShiftAction, _CState318, 0}
	_CGotoState319Action                      = &_CAction{_CShiftAction, _CState319, 0}
	_CGotoState320Action                      = &_CAction{_CShiftAction, _CState320, 0}
	_CGotoState321Action                      = &_CAction{_CShiftAction, _CState321, 0}
	_CGotoState322Action                      = &_CAction{_CShiftAction, _CState322, 0}
	_CGotoState323Action                      = &_CAction{_CShiftAction, _CState323, 0}
	_CGotoState324Action                      = &_CAction{_CShiftAction, _CState324, 0}
	_CGotoState325Action                      = &_CAction{_CShiftAction, _CState325, 0}
	_CGotoState326Action                      = &_CAction{_CShiftAction, _CState326, 0}
	_CGotoState327Action                      = &_CAction{_CShiftAction, _CState327, 0}
	_CGotoState328Action                      = &_CAction{_CShiftAction, _CState328, 0}
	_CGotoState329Action                      = &_CAction{_CShiftAction, _CState329, 0}
	_CGotoState330Action                      = &_CAction{_CShiftAction, _CState330, 0}
	_CGotoState331Action                      = &_CAction{_CShiftAction, _CState331, 0}
	_CGotoState332Action                      = &_CAction{_CShiftAction, _CState332, 0}
	_CGotoState333Action                      = &_CAction{_CShiftAction, _CState333, 0}
	_CGotoState334Action                      = &_CAction{_CShiftAction, _CState334, 0}
	_CGotoState335Action                      = &_CAction{_CShiftAction, _CState335, 0}
	_CGotoState336Action                      = &_CAction{_CShiftAction, _CState336, 0}
	_CGotoState337Action                      = &_CAction{_CShiftAction, _CState337, 0}
	_CGotoState338Action                      = &_CAction{_CShiftAction, _CState338, 0}
	_CGotoState339Action                      = &_CAction{_CShiftAction, _CState339, 0}
	_CGotoState340Action                      = &_CAction{_CShiftAction, _CState340, 0}
	_CGotoState341Action                      = &_CAction{_CShiftAction, _CState341, 0}
	_CGotoState342Action                      = &_CAction{_CShiftAction, _CState342, 0}
	_CGotoState343Action                      = &_CAction{_CShiftAction, _CState343, 0}
	_CGotoState344Action                      = &_CAction{_CShiftAction, _CState344, 0}
	_CGotoState345Action                      = &_CAction{_CShiftAction, _CState345, 0}
	_CGotoState346Action                      = &_CAction{_CShiftAction, _CState346, 0}
	_CGotoState347Action                      = &_CAction{_CShiftAction, _CState347, 0}
	_CGotoState348Action                      = &_CAction{_CShiftAction, _CState348, 0}
	_CGotoState349Action                      = &_CAction{_CShiftAction, _CState349, 0}
	_CGotoState350Action                      = &_CAction{_CShiftAction, _CState350, 0}
	_CGotoState351Action                      = &_CAction{_CShiftAction, _CState351, 0}
	_CReduceAToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceAToPrimaryExpression}
	_CReduceBToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceBToPrimaryExpression}
	_CReduceCToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceCToPrimaryExpression}
	_CReduceDToPrimaryExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceDToPrimaryExpression}
	_CReduceAToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceAToPostfixExpression}
	_CReduceBToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceBToPostfixExpression}
	_CReduceCToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceCToPostfixExpression}
	_CReduceDToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceDToPostfixExpression}
	_CReduceEToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceEToPostfixExpression}
	_CReduceFToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceFToPostfixExpression}
	_CReduceGToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceGToPostfixExpression}
	_CReduceHToPostfixExpressionAction        = &_CAction{_CReduceAction, 0, _CReduceHToPostfixExpression}
	_CReduceAToArgumentExpressionListAction   = &_CAction{_CReduceAction, 0, _CReduceAToArgumentExpressionList}
	_CReduceBToArgumentExpressionListAction   = &_CAction{_CReduceAction, 0, _CReduceBToArgumentExpressionList}
	_CReduceAToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceAToUnaryExpression}
	_CReduceBToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceBToUnaryExpression}
	_CReduceCToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceCToUnaryExpression}
	_CReduceDToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceDToUnaryExpression}
	_CReduceEToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceEToUnaryExpression}
	_CReduceFToUnaryExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceFToUnaryExpression}
	_CReduceAToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceAToUnaryOperator}
	_CReduceBToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceBToUnaryOperator}
	_CReduceCToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceCToUnaryOperator}
	_CReduceDToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceDToUnaryOperator}
	_CReduceEToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceEToUnaryOperator}
	_CReduceFToUnaryOperatorAction            = &_CAction{_CReduceAction, 0, _CReduceFToUnaryOperator}
	_CReduceAToCastExpressionAction           = &_CAction{_CReduceAction, 0, _CReduceAToCastExpression}
	_CReduceBToCastExpressionAction           = &_CAction{_CReduceAction, 0, _CReduceBToCastExpression}
	_CReduceAToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceAToMultiplicativeExpression}
	_CReduceBToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceBToMultiplicativeExpression}
	_CReduceCToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceCToMultiplicativeExpression}
	_CReduceDToMultiplicativeExpressionAction = &_CAction{_CReduceAction, 0, _CReduceDToMultiplicativeExpression}
	_CReduceAToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToAdditiveExpression}
	_CReduceBToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceBToAdditiveExpression}
	_CReduceCToAdditiveExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceCToAdditiveExpression}
	_CReduceAToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceAToShiftExpression}
	_CReduceBToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceBToShiftExpression}
	_CReduceCToShiftExpressionAction          = &_CAction{_CReduceAction, 0, _CReduceCToShiftExpression}
	_CReduceAToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToRelationalExpression}
	_CReduceBToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToRelationalExpression}
	_CReduceCToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceCToRelationalExpression}
	_CReduceDToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceDToRelationalExpression}
	_CReduceEToRelationalExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceEToRelationalExpression}
	_CReduceAToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToEqualityExpression}
	_CReduceBToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceBToEqualityExpression}
	_CReduceCToEqualityExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceCToEqualityExpression}
	_CReduceAToAndExpressionAction            = &_CAction{_CReduceAction, 0, _CReduceAToAndExpression}
	_CReduceBToAndExpressionAction            = &_CAction{_CReduceAction, 0, _CReduceBToAndExpression}
	_CReduceAToExclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToExclusiveOrExpression}
	_CReduceBToExclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToExclusiveOrExpression}
	_CReduceAToInclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToInclusiveOrExpression}
	_CReduceBToInclusiveOrExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToInclusiveOrExpression}
	_CReduceAToLogicalAndExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToLogicalAndExpression}
	_CReduceBToLogicalAndExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToLogicalAndExpression}
	_CReduceAToLogicalOrExpressionAction      = &_CAction{_CReduceAction, 0, _CReduceAToLogicalOrExpression}
	_CReduceBToLogicalOrExpressionAction      = &_CAction{_CReduceAction, 0, _CReduceBToLogicalOrExpression}
	_CReduceAToConditionalExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceAToConditionalExpression}
	_CReduceBToConditionalExpressionAction    = &_CAction{_CReduceAction, 0, _CReduceBToConditionalExpression}
	_CReduceAToAssignmentExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceAToAssignmentExpression}
	_CReduceBToAssignmentExpressionAction     = &_CAction{_CReduceAction, 0, _CReduceBToAssignmentExpression}
	_CReduceAToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceAToAssignmentOperator}
	_CReduceBToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceBToAssignmentOperator}
	_CReduceCToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceCToAssignmentOperator}
	_CReduceDToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceDToAssignmentOperator}
	_CReduceEToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceEToAssignmentOperator}
	_CReduceFToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceFToAssignmentOperator}
	_CReduceGToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceGToAssignmentOperator}
	_CReduceHToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceHToAssignmentOperator}
	_CReduceIToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceIToAssignmentOperator}
	_CReduceJToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceJToAssignmentOperator}
	_CReduceKToAssignmentOperatorAction       = &_CAction{_CReduceAction, 0, _CReduceKToAssignmentOperator}
	_CReduceAToExpressionAction               = &_CAction{_CReduceAction, 0, _CReduceAToExpression}
	_CReduceBToExpressionAction               = &_CAction{_CReduceAction, 0, _CReduceBToExpression}
	_CReduceAToConstantExpressionAction       = &_CAction{_CReduceAction, 0, _CReduceAToConstantExpression}
	_CReduceAToDeclarationAction              = &_CAction{_CReduceAction, 0, _CReduceAToDeclaration}
	_CReduceBToDeclarationAction              = &_CAction{_CReduceAction, 0, _CReduceBToDeclaration}
	_CReduceAToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceAToDeclarationSpecifiers}
	_CReduceBToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceBToDeclarationSpecifiers}
	_CReduceCToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceCToDeclarationSpecifiers}
	_CReduceDToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceDToDeclarationSpecifiers}
	_CReduceEToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceEToDeclarationSpecifiers}
	_CReduceFToDeclarationSpecifiersAction    = &_CAction{_CReduceAction, 0, _CReduceFToDeclarationSpecifiers}
	_CReduceAToInitDeclaratorListAction       = &_CAction{_CReduceAction, 0, _CReduceAToInitDeclaratorList}
	_CReduceBToInitDeclaratorListAction       = &_CAction{_CReduceAction, 0, _CReduceBToInitDeclaratorList}
	_CReduceAToInitDeclaratorAction           = &_CAction{_CReduceAction, 0, _CReduceAToInitDeclarator}
	_CReduceBToInitDeclaratorAction           = &_CAction{_CReduceAction, 0, _CReduceBToInitDeclarator}
	_CReduceAToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceAToStorageClassSpecifier}
	_CReduceBToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceBToStorageClassSpecifier}
	_CReduceCToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceCToStorageClassSpecifier}
	_CReduceDToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceDToStorageClassSpecifier}
	_CReduceEToStorageClassSpecifierAction    = &_CAction{_CReduceAction, 0, _CReduceEToStorageClassSpecifier}
	_CReduceAToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToTypeSpecifier}
	_CReduceBToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToTypeSpecifier}
	_CReduceCToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceCToTypeSpecifier}
	_CReduceDToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceDToTypeSpecifier}
	_CReduceEToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceEToTypeSpecifier}
	_CReduceFToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceFToTypeSpecifier}
	_CReduceGToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceGToTypeSpecifier}
	_CReduceHToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceHToTypeSpecifier}
	_CReduceIToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceIToTypeSpecifier}
	_CReduceJToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceJToTypeSpecifier}
	_CReduceKToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceKToTypeSpecifier}
	_CReduceLToTypeSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceLToTypeSpecifier}
	_CReduceAToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceAToStructOrUnionSpecifier}
	_CReduceBToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceBToStructOrUnionSpecifier}
	_CReduceCToStructOrUnionSpecifierAction   = &_CAction{_CReduceAction, 0, _CReduceCToStructOrUnionSpecifier}
	_CReduceAToStructOrUnionAction            = &_CAction{_CReduceAction, 0, _CReduceAToStructOrUnion}
	_CReduceBToStructOrUnionAction            = &_CAction{_CReduceAction, 0, _CReduceBToStructOrUnion}
	_CReduceAToStructDeclarationListAction    = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclarationList}
	_CReduceBToStructDeclarationListAction    = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclarationList}
	_CReduceAToStructDeclarationAction        = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclaration}
	_CReduceAToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceAToSpecifierQualifierList}
	_CReduceBToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceBToSpecifierQualifierList}
	_CReduceCToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceCToSpecifierQualifierList}
	_CReduceDToSpecifierQualifierListAction   = &_CAction{_CReduceAction, 0, _CReduceDToSpecifierQualifierList}
	_CReduceAToStructDeclaratorListAction     = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclaratorList}
	_CReduceBToStructDeclaratorListAction     = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclaratorList}
	_CReduceAToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceAToStructDeclarator}
	_CReduceBToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceBToStructDeclarator}
	_CReduceCToStructDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceCToStructDeclarator}
	_CReduceAToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToEnumSpecifier}
	_CReduceBToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToEnumSpecifier}
	_CReduceCToEnumSpecifierAction            = &_CAction{_CReduceAction, 0, _CReduceCToEnumSpecifier}
	_CReduceAToEnumeratorListAction           = &_CAction{_CReduceAction, 0, _CReduceAToEnumeratorList}
	_CReduceBToEnumeratorListAction           = &_CAction{_CReduceAction, 0, _CReduceBToEnumeratorList}
	_CReduceAToEnumeratorAction               = &_CAction{_CReduceAction, 0, _CReduceAToEnumerator}
	_CReduceBToEnumeratorAction               = &_CAction{_CReduceAction, 0, _CReduceBToEnumerator}
	_CReduceAToTypeQualifierAction            = &_CAction{_CReduceAction, 0, _CReduceAToTypeQualifier}
	_CReduceBToTypeQualifierAction            = &_CAction{_CReduceAction, 0, _CReduceBToTypeQualifier}
	_CReduceAToDeclaratorAction               = &_CAction{_CReduceAction, 0, _CReduceAToDeclarator}
	_CReduceBToDeclaratorAction               = &_CAction{_CReduceAction, 0, _CReduceBToDeclarator}
	_CReduceAToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceAToDirectDeclarator}
	_CReduceBToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceBToDirectDeclarator}
	_CReduceCToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceCToDirectDeclarator}
	_CReduceDToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceDToDirectDeclarator}
	_CReduceEToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceEToDirectDeclarator}
	_CReduceFToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceFToDirectDeclarator}
	_CReduceGToDirectDeclaratorAction         = &_CAction{_CReduceAction, 0, _CReduceGToDirectDeclarator}
	_CReduceAToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceAToPointer}
	_CReduceBToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceBToPointer}
	_CReduceCToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceCToPointer}
	_CReduceDToPointerAction                  = &_CAction{_CReduceAction, 0, _CReduceDToPointer}
	_CReduceAToTypeQualifierListAction        = &_CAction{_CReduceAction, 0, _CReduceAToTypeQualifierList}
	_CReduceBToTypeQualifierListAction        = &_CAction{_CReduceAction, 0, _CReduceBToTypeQualifierList}
	_CReduceAToParameterTypeListAction        = &_CAction{_CReduceAction, 0, _CReduceAToParameterTypeList}
	_CReduceBToParameterTypeListAction        = &_CAction{_CReduceAction, 0, _CReduceBToParameterTypeList}
	_CReduceAToParameterListAction            = &_CAction{_CReduceAction, 0, _CReduceAToParameterList}
	_CReduceBToParameterListAction            = &_CAction{_CReduceAction, 0, _CReduceBToParameterList}
	_CReduceAToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceAToParameterDeclaration}
	_CReduceBToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceBToParameterDeclaration}
	_CReduceCToParameterDeclarationAction     = &_CAction{_CReduceAction, 0, _CReduceCToParameterDeclaration}
	_CReduceAToIdentifierListAction           = &_CAction{_CReduceAction, 0, _CReduceAToIdentifierList}
	_CReduceBToIdentifierListAction           = &_CAction{_CReduceAction, 0, _CReduceBToIdentifierList}
	_CReduceAToTypeNameAction                 = &_CAction{_CReduceAction, 0, _CReduceAToTypeName}
	_CReduceBToTypeNameAction                 = &_CAction{_CReduceAction, 0, _CReduceBToTypeName}
	_CReduceAToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceAToAbstractDeclarator}
	_CReduceBToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceBToAbstractDeclarator}
	_CReduceCToAbstractDeclaratorAction       = &_CAction{_CReduceAction, 0, _CReduceCToAbstractDeclarator}
	_CReduceAToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceAToDirectAbstractDeclarator}
	_CReduceBToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceBToDirectAbstractDeclarator}
	_CReduceCToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceCToDirectAbstractDeclarator}
	_CReduceDToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceDToDirectAbstractDeclarator}
	_CReduceEToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceEToDirectAbstractDeclarator}
	_CReduceFToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceFToDirectAbstractDeclarator}
	_CReduceGToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceGToDirectAbstractDeclarator}
	_CReduceHToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceHToDirectAbstractDeclarator}
	_CReduceIToDirectAbstractDeclaratorAction = &_CAction{_CReduceAction, 0, _CReduceIToDirectAbstractDeclarator}
	_CReduceAToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceAToInitializer}
	_CReduceBToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceBToInitializer}
	_CReduceCToInitializerAction              = &_CAction{_CReduceAction, 0, _CReduceCToInitializer}
	_CReduceAToInitializerListAction          = &_CAction{_CReduceAction, 0, _CReduceAToInitializerList}
	_CReduceBToInitializerListAction          = &_CAction{_CReduceAction, 0, _CReduceBToInitializerList}
	_CReduceAToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceAToStatement}
	_CReduceBToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceBToStatement}
	_CReduceCToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceCToStatement}
	_CReduceDToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceDToStatement}
	_CReduceEToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceEToStatement}
	_CReduceFToStatementAction                = &_CAction{_CReduceAction, 0, _CReduceFToStatement}
	_CReduceAToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceAToLabeledStatement}
	_CReduceBToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceBToLabeledStatement}
	_CReduceCToLabeledStatementAction         = &_CAction{_CReduceAction, 0, _CReduceCToLabeledStatement}
	_CReduceAToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceAToCompoundStatement}
	_CReduceBToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceBToCompoundStatement}
	_CReduceCToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceCToCompoundStatement}
	_CReduceDToCompoundStatementAction        = &_CAction{_CReduceAction, 0, _CReduceDToCompoundStatement}
	_CReduceAToDeclarationListAction          = &_CAction{_CReduceAction, 0, _CReduceAToDeclarationList}
	_CReduceBToDeclarationListAction          = &_CAction{_CReduceAction, 0, _CReduceBToDeclarationList}
	_CReduceAToStatementListAction            = &_CAction{_CReduceAction, 0, _CReduceAToStatementList}
	_CReduceBToStatementListAction            = &_CAction{_CReduceAction, 0, _CReduceBToStatementList}
	_CReduceAToExpressionStatementAction      = &_CAction{_CReduceAction, 0, _CReduceAToExpressionStatement}
	_CReduceBToExpressionStatementAction      = &_CAction{_CReduceAction, 0, _CReduceBToExpressionStatement}
	_CReduceAToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceAToSelectionStatement}
	_CReduceBToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceBToSelectionStatement}
	_CReduceCToSelectionStatementAction       = &_CAction{_CReduceAction, 0, _CReduceCToSelectionStatement}
	_CReduceAToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceAToIterationStatement}
	_CReduceBToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceBToIterationStatement}
	_CReduceCToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceCToIterationStatement}
	_CReduceDToIterationStatementAction       = &_CAction{_CReduceAction, 0, _CReduceDToIterationStatement}
	_CReduceAToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceAToJumpStatement}
	_CReduceBToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceBToJumpStatement}
	_CReduceCToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceCToJumpStatement}
	_CReduceDToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceDToJumpStatement}
	_CReduceEToJumpStatementAction            = &_CAction{_CReduceAction, 0, _CReduceEToJumpStatement}
	_CReduceAToTranslationUnitAction          = &_CAction{_CReduceAction, 0, _CReduceAToTranslationUnit}
	_CReduceBToTranslationUnitAction          = &_CAction{_CReduceAction, 0, _CReduceBToTranslationUnit}
	_CReduceAToExternalDeclarationAction      = &_CAction{_CReduceAction, 0, _CReduceAToExternalDeclaration}
	_CReduceBToExternalDeclarationAction      = &_CAction{_CReduceAction, 0, _CReduceBToExternalDeclaration}
	_CReduceAToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceAToFunctionDefinition}
	_CReduceBToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceBToFunctionDefinition}
	_CReduceCToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceCToFunctionDefinition}
	_CReduceDToFunctionDefinitionAction       = &_CAction{_CReduceAction, 0, _CReduceDToFunctionDefinition}
)

type _CActionTableKey struct {
	_CStateId
	CSymbolId
}

type _CActionTableType map[_CActionTableKey]*_CAction

func (table _CActionTableType) Get(stateId _CStateId, symbol CSymbolId) (*_CAction, bool) {
	action, ok := table[_CActionTableKey{stateId, symbol}]
	if ok {
		return action, ok
	}
	action, ok = table[_CActionTableKey{stateId, _CWildcardMarker}]
	return action, ok
}

var _CActionTable = _CActionTableType{
	{_CState2, _CEndMarker}:                     &_CAction{_CAcceptAction, 0, 0},
	{_CState1, CIdentifierToken}:                _CGotoState10Action,
	{_CState1, CTypeNameToken}:                  _CGotoState21Action,
	{_CState1, CTypedefToken}:                   _CGotoState20Action,
	{_CState1, CExternToken}:                    _CGotoState8Action,
	{_CState1, CStaticToken}:                    _CGotoState18Action,
	{_CState1, CAutoToken}:                      _CGotoState3Action,
	{_CState1, CRegisterToken}:                  _CGotoState15Action,
	{_CState1, CCharToken}:                      _CGotoState4Action,
	{_CState1, CShortToken}:                     _CGotoState16Action,
	{_CState1, CIntToken}:                       _CGotoState11Action,
	{_CState1, CLongToken}:                      _CGotoState12Action,
	{_CState1, CSignedToken}:                    _CGotoState17Action,
	{_CState1, CUnsignedToken}:                  _CGotoState23Action,
	{_CState1, CFloatToken}:                     _CGotoState9Action,
	{_CState1, CDoubleToken}:                    _CGotoState6Action,
	{_CState1, CConstToken}:                     _CGotoState5Action,
	{_CState1, CVolatileToken}:                  _CGotoState25Action,
	{_CState1, CVoidToken}:                      _CGotoState24Action,
	{_CState1, CStructToken}:                    _CGotoState19Action,
	{_CState1, CUnionToken}:                     _CGotoState22Action,
	{_CState1, CEnumToken}:                      _CGotoState7Action,
	{_CState1, CLparamToken}:                    _CGotoState13Action,
	{_CState1, CMulToken}:                       _CGotoState14Action,
	{_CState1, CDeclarationType}:                _CGotoState26Action,
	{_CState1, CDeclarationSpecifiersType}:      _CGotoState27Action,
	{_CState1, CStorageClassSpecifierType}:      _CGotoState34Action,
	{_CState1, CTypeSpecifierType}:              _CGotoState38Action,
	{_CState1, CStructOrUnionSpecifierType}:     _CGotoState36Action,
	{_CState1, CStructOrUnionType}:              _CGotoState35Action,
	{_CState1, CEnumSpecifierType}:              _CGotoState30Action,
	{_CState1, CTypeQualifierType}:              _CGotoState37Action,
	{_CState1, CDeclaratorType}:                 _CGotoState28Action,
	{_CState1, CDirectDeclaratorType}:           _CGotoState29Action,
	{_CState1, CPointerType}:                    _CGotoState33Action,
	{_CState1, CTranslationUnitType}:            _CGotoState2Action,
	{_CState1, CExternalDeclarationType}:        _CGotoState31Action,
	{_CState1, CFunctionDefinitionType}:         _CGotoState32Action,
	{_CState2, CIdentifierToken}:                _CGotoState10Action,
	{_CState2, CTypeNameToken}:                  _CGotoState21Action,
	{_CState2, CTypedefToken}:                   _CGotoState20Action,
	{_CState2, CExternToken}:                    _CGotoState8Action,
	{_CState2, CStaticToken}:                    _CGotoState18Action,
	{_CState2, CAutoToken}:                      _CGotoState3Action,
	{_CState2, CRegisterToken}:                  _CGotoState15Action,
	{_CState2, CCharToken}:                      _CGotoState4Action,
	{_CState2, CShortToken}:                     _CGotoState16Action,
	{_CState2, CIntToken}:                       _CGotoState11Action,
	{_CState2, CLongToken}:                      _CGotoState12Action,
	{_CState2, CSignedToken}:                    _CGotoState17Action,
	{_CState2, CUnsignedToken}:                  _CGotoState23Action,
	{_CState2, CFloatToken}:                     _CGotoState9Action,
	{_CState2, CDoubleToken}:                    _CGotoState6Action,
	{_CState2, CConstToken}:                     _CGotoState5Action,
	{_CState2, CVolatileToken}:                  _CGotoState25Action,
	{_CState2, CVoidToken}:                      _CGotoState24Action,
	{_CState2, CStructToken}:                    _CGotoState19Action,
	{_CState2, CUnionToken}:                     _CGotoState22Action,
	{_CState2, CEnumToken}:                      _CGotoState7Action,
	{_CState2, CLparamToken}:                    _CGotoState13Action,
	{_CState2, CMulToken}:                       _CGotoState14Action,
	{_CState2, CDeclarationType}:                _CGotoState26Action,
	{_CState2, CDeclarationSpecifiersType}:      _CGotoState27Action,
	{_CState2, CStorageClassSpecifierType}:      _CGotoState34Action,
	{_CState2, CTypeSpecifierType}:              _CGotoState38Action,
	{_CState2, CStructOrUnionSpecifierType}:     _CGotoState36Action,
	{_CState2, CStructOrUnionType}:              _CGotoState35Action,
	{_CState2, CEnumSpecifierType}:              _CGotoState30Action,
	{_CState2, CTypeQualifierType}:              _CGotoState37Action,
	{_CState2, CDeclaratorType}:                 _CGotoState28Action,
	{_CState2, CDirectDeclaratorType}:           _CGotoState29Action,
	{_CState2, CPointerType}:                    _CGotoState33Action,
	{_CState2, CExternalDeclarationType}:        _CGotoState60Action,
	{_CState2, CFunctionDefinitionType}:         _CGotoState32Action,
	{_CState7, CIdentifierToken}:                _CGotoState39Action,
	{_CState7, CLcurlToken}:                     _CGotoState40Action,
	{_CState13, CIdentifierToken}:               _CGotoState10Action,
	{_CState13, CLparamToken}:                   _CGotoState13Action,
	{_CState13, CMulToken}:                      _CGotoState14Action,
	{_CState13, CDeclaratorType}:                _CGotoState41Action,
	{_CState13, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState13, CPointerType}:                   _CGotoState33Action,
	{_CState14, CConstToken}:                    _CGotoState5Action,
	{_CState14, CVolatileToken}:                 _CGotoState25Action,
	{_CState14, CMulToken}:                      _CGotoState14Action,
	{_CState14, CTypeQualifierType}:             _CGotoState43Action,
	{_CState14, CPointerType}:                   _CGotoState42Action,
	{_CState14, CTypeQualifierListType}:         _CGotoState44Action,
	{_CState27, CIdentifierToken}:               _CGotoState10Action,
	{_CState27, CLparamToken}:                   _CGotoState13Action,
	{_CState27, CSemicolonToken}:                _CGotoState45Action,
	{_CState27, CMulToken}:                      _CGotoState14Action,
	{_CState27, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState27, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState27, CDeclaratorType}:                _CGotoState46Action,
	{_CState27, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState27, CPointerType}:                   _CGotoState33Action,
	{_CState28, CTypeNameToken}:                 _CGotoState21Action,
	{_CState28, CTypedefToken}:                  _CGotoState20Action,
	{_CState28, CExternToken}:                   _CGotoState8Action,
	{_CState28, CStaticToken}:                   _CGotoState18Action,
	{_CState28, CAutoToken}:                     _CGotoState3Action,
	{_CState28, CRegisterToken}:                 _CGotoState15Action,
	{_CState28, CCharToken}:                     _CGotoState4Action,
	{_CState28, CShortToken}:                    _CGotoState16Action,
	{_CState28, CIntToken}:                      _CGotoState11Action,
	{_CState28, CLongToken}:                     _CGotoState12Action,
	{_CState28, CSignedToken}:                   _CGotoState17Action,
	{_CState28, CUnsignedToken}:                 _CGotoState23Action,
	{_CState28, CFloatToken}:                    _CGotoState9Action,
	{_CState28, CDoubleToken}:                   _CGotoState6Action,
	{_CState28, CConstToken}:                    _CGotoState5Action,
	{_CState28, CVolatileToken}:                 _CGotoState25Action,
	{_CState28, CVoidToken}:                     _CGotoState24Action,
	{_CState28, CStructToken}:                   _CGotoState19Action,
	{_CState28, CUnionToken}:                    _CGotoState22Action,
	{_CState28, CEnumToken}:                     _CGotoState7Action,
	{_CState28, CLcurlToken}:                    _CGotoState49Action,
	{_CState28, CDeclarationType}:               _CGotoState51Action,
	{_CState28, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState28, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState28, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState28, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState28, CStructOrUnionType}:             _CGotoState35Action,
	{_CState28, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState28, CTypeQualifierType}:             _CGotoState37Action,
	{_CState28, CCompoundStatementType}:         _CGotoState50Action,
	{_CState28, CDeclarationListType}:           _CGotoState52Action,
	{_CState29, CLparamToken}:                   _CGotoState55Action,
	{_CState29, CLbraceToken}:                   _CGotoState54Action,
	{_CState33, CIdentifierToken}:               _CGotoState10Action,
	{_CState33, CLparamToken}:                   _CGotoState13Action,
	{_CState33, CDirectDeclaratorType}:          _CGotoState56Action,
	{_CState34, CTypeNameToken}:                 _CGotoState21Action,
	{_CState34, CTypedefToken}:                  _CGotoState20Action,
	{_CState34, CExternToken}:                   _CGotoState8Action,
	{_CState34, CStaticToken}:                   _CGotoState18Action,
	{_CState34, CAutoToken}:                     _CGotoState3Action,
	{_CState34, CRegisterToken}:                 _CGotoState15Action,
	{_CState34, CCharToken}:                     _CGotoState4Action,
	{_CState34, CShortToken}:                    _CGotoState16Action,
	{_CState34, CIntToken}:                      _CGotoState11Action,
	{_CState34, CLongToken}:                     _CGotoState12Action,
	{_CState34, CSignedToken}:                   _CGotoState17Action,
	{_CState34, CUnsignedToken}:                 _CGotoState23Action,
	{_CState34, CFloatToken}:                    _CGotoState9Action,
	{_CState34, CDoubleToken}:                   _CGotoState6Action,
	{_CState34, CConstToken}:                    _CGotoState5Action,
	{_CState34, CVolatileToken}:                 _CGotoState25Action,
	{_CState34, CVoidToken}:                     _CGotoState24Action,
	{_CState34, CStructToken}:                   _CGotoState19Action,
	{_CState34, CUnionToken}:                    _CGotoState22Action,
	{_CState34, CEnumToken}:                     _CGotoState7Action,
	{_CState34, CDeclarationSpecifiersType}:     _CGotoState57Action,
	{_CState34, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState34, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState34, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState34, CStructOrUnionType}:             _CGotoState35Action,
	{_CState34, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState34, CTypeQualifierType}:             _CGotoState37Action,
	{_CState35, CIdentifierToken}:               _CGotoState58Action,
	{_CState35, CLcurlToken}:                    _CGotoState59Action,
	{_CState37, CTypeNameToken}:                 _CGotoState21Action,
	{_CState37, CTypedefToken}:                  _CGotoState20Action,
	{_CState37, CExternToken}:                   _CGotoState8Action,
	{_CState37, CStaticToken}:                   _CGotoState18Action,
	{_CState37, CAutoToken}:                     _CGotoState3Action,
	{_CState37, CRegisterToken}:                 _CGotoState15Action,
	{_CState37, CCharToken}:                     _CGotoState4Action,
	{_CState37, CShortToken}:                    _CGotoState16Action,
	{_CState37, CIntToken}:                      _CGotoState11Action,
	{_CState37, CLongToken}:                     _CGotoState12Action,
	{_CState37, CSignedToken}:                   _CGotoState17Action,
	{_CState37, CUnsignedToken}:                 _CGotoState23Action,
	{_CState37, CFloatToken}:                    _CGotoState9Action,
	{_CState37, CDoubleToken}:                   _CGotoState6Action,
	{_CState37, CConstToken}:                    _CGotoState5Action,
	{_CState37, CVolatileToken}:                 _CGotoState25Action,
	{_CState37, CVoidToken}:                     _CGotoState24Action,
	{_CState37, CStructToken}:                   _CGotoState19Action,
	{_CState37, CUnionToken}:                    _CGotoState22Action,
	{_CState37, CEnumToken}:                     _CGotoState7Action,
	{_CState37, CDeclarationSpecifiersType}:     _CGotoState61Action,
	{_CState37, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState37, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState37, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState37, CStructOrUnionType}:             _CGotoState35Action,
	{_CState37, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState37, CTypeQualifierType}:             _CGotoState37Action,
	{_CState38, CTypeNameToken}:                 _CGotoState21Action,
	{_CState38, CTypedefToken}:                  _CGotoState20Action,
	{_CState38, CExternToken}:                   _CGotoState8Action,
	{_CState38, CStaticToken}:                   _CGotoState18Action,
	{_CState38, CAutoToken}:                     _CGotoState3Action,
	{_CState38, CRegisterToken}:                 _CGotoState15Action,
	{_CState38, CCharToken}:                     _CGotoState4Action,
	{_CState38, CShortToken}:                    _CGotoState16Action,
	{_CState38, CIntToken}:                      _CGotoState11Action,
	{_CState38, CLongToken}:                     _CGotoState12Action,
	{_CState38, CSignedToken}:                   _CGotoState17Action,
	{_CState38, CUnsignedToken}:                 _CGotoState23Action,
	{_CState38, CFloatToken}:                    _CGotoState9Action,
	{_CState38, CDoubleToken}:                   _CGotoState6Action,
	{_CState38, CConstToken}:                    _CGotoState5Action,
	{_CState38, CVolatileToken}:                 _CGotoState25Action,
	{_CState38, CVoidToken}:                     _CGotoState24Action,
	{_CState38, CStructToken}:                   _CGotoState19Action,
	{_CState38, CUnionToken}:                    _CGotoState22Action,
	{_CState38, CEnumToken}:                     _CGotoState7Action,
	{_CState38, CDeclarationSpecifiersType}:     _CGotoState62Action,
	{_CState38, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState38, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState38, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState38, CStructOrUnionType}:             _CGotoState35Action,
	{_CState38, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState38, CTypeQualifierType}:             _CGotoState37Action,
	{_CState39, CLcurlToken}:                    _CGotoState63Action,
	{_CState40, CIdentifierToken}:               _CGotoState64Action,
	{_CState40, CEnumeratorListType}:            _CGotoState66Action,
	{_CState40, CEnumeratorType}:                _CGotoState65Action,
	{_CState41, CRparamToken}:                   _CGotoState67Action,
	{_CState44, CConstToken}:                    _CGotoState5Action,
	{_CState44, CVolatileToken}:                 _CGotoState25Action,
	{_CState44, CMulToken}:                      _CGotoState14Action,
	{_CState44, CTypeQualifierType}:             _CGotoState69Action,
	{_CState44, CPointerType}:                   _CGotoState68Action,
	{_CState46, CTypeNameToken}:                 _CGotoState21Action,
	{_CState46, CTypedefToken}:                  _CGotoState20Action,
	{_CState46, CExternToken}:                   _CGotoState8Action,
	{_CState46, CStaticToken}:                   _CGotoState18Action,
	{_CState46, CAutoToken}:                     _CGotoState3Action,
	{_CState46, CRegisterToken}:                 _CGotoState15Action,
	{_CState46, CCharToken}:                     _CGotoState4Action,
	{_CState46, CShortToken}:                    _CGotoState16Action,
	{_CState46, CIntToken}:                      _CGotoState11Action,
	{_CState46, CLongToken}:                     _CGotoState12Action,
	{_CState46, CSignedToken}:                   _CGotoState17Action,
	{_CState46, CUnsignedToken}:                 _CGotoState23Action,
	{_CState46, CFloatToken}:                    _CGotoState9Action,
	{_CState46, CDoubleToken}:                   _CGotoState6Action,
	{_CState46, CConstToken}:                    _CGotoState5Action,
	{_CState46, CVolatileToken}:                 _CGotoState25Action,
	{_CState46, CVoidToken}:                     _CGotoState24Action,
	{_CState46, CStructToken}:                   _CGotoState19Action,
	{_CState46, CUnionToken}:                    _CGotoState22Action,
	{_CState46, CEnumToken}:                     _CGotoState7Action,
	{_CState46, CLcurlToken}:                    _CGotoState49Action,
	{_CState46, CEqToken}:                       _CGotoState70Action,
	{_CState46, CDeclarationType}:               _CGotoState51Action,
	{_CState46, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState46, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState46, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState46, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState46, CStructOrUnionType}:             _CGotoState35Action,
	{_CState46, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState46, CTypeQualifierType}:             _CGotoState37Action,
	{_CState46, CCompoundStatementType}:         _CGotoState71Action,
	{_CState46, CDeclarationListType}:           _CGotoState72Action,
	{_CState48, CSemicolonToken}:                _CGotoState74Action,
	{_CState48, CCommaToken}:                    _CGotoState73Action,
	{_CState49, CIdentifierToken}:               _CGotoState86Action,
	{_CState49, CConstantToken}:                 _CGotoState78Action,
	{_CState49, CStringLiteralToken}:            _CGotoState97Action,
	{_CState49, CSizeofToken}:                   _CGotoState96Action,
	{_CState49, CIncOpToken}:                    _CGotoState88Action,
	{_CState49, CDecOpToken}:                    _CGotoState80Action,
	{_CState49, CTypeNameToken}:                 _CGotoState21Action,
	{_CState49, CTypedefToken}:                  _CGotoState20Action,
	{_CState49, CExternToken}:                   _CGotoState8Action,
	{_CState49, CStaticToken}:                   _CGotoState18Action,
	{_CState49, CAutoToken}:                     _CGotoState3Action,
	{_CState49, CRegisterToken}:                 _CGotoState15Action,
	{_CState49, CCharToken}:                     _CGotoState4Action,
	{_CState49, CShortToken}:                    _CGotoState16Action,
	{_CState49, CIntToken}:                      _CGotoState11Action,
	{_CState49, CLongToken}:                     _CGotoState12Action,
	{_CState49, CSignedToken}:                   _CGotoState17Action,
	{_CState49, CUnsignedToken}:                 _CGotoState23Action,
	{_CState49, CFloatToken}:                    _CGotoState9Action,
	{_CState49, CDoubleToken}:                   _CGotoState6Action,
	{_CState49, CConstToken}:                    _CGotoState5Action,
	{_CState49, CVolatileToken}:                 _CGotoState25Action,
	{_CState49, CVoidToken}:                     _CGotoState24Action,
	{_CState49, CStructToken}:                   _CGotoState19Action,
	{_CState49, CUnionToken}:                    _CGotoState22Action,
	{_CState49, CEnumToken}:                     _CGotoState7Action,
	{_CState49, CCaseToken}:                     _CGotoState77Action,
	{_CState49, CDefaultToken}:                  _CGotoState81Action,
	{_CState49, CIfToken}:                       _CGotoState87Action,
	{_CState49, CSwitchToken}:                   _CGotoState98Action,
	{_CState49, CWhileToken}:                    _CGotoState100Action,
	{_CState49, CDoToken}:                       _CGotoState82Action,
	{_CState49, CForToken}:                      _CGotoState84Action,
	{_CState49, CGotoToken}:                     _CGotoState85Action,
	{_CState49, CContinueToken}:                 _CGotoState79Action,
	{_CState49, CBreakToken}:                    _CGotoState76Action,
	{_CState49, CReturnToken}:                   _CGotoState94Action,
	{_CState49, CLparamToken}:                   _CGotoState89Action,
	{_CState49, CLcurlToken}:                    _CGotoState49Action,
	{_CState49, CRcurlToken}:                    _CGotoState93Action,
	{_CState49, CSemicolonToken}:                _CGotoState95Action,
	{_CState49, CMulToken}:                      _CGotoState91Action,
	{_CState49, CMinusToken}:                    _CGotoState90Action,
	{_CState49, CPlusToken}:                     _CGotoState92Action,
	{_CState49, CAndToken}:                      _CGotoState75Action,
	{_CState49, CExclaimToken}:                  _CGotoState83Action,
	{_CState49, CTildaToken}:                    _CGotoState99Action,
	{_CState49, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState49, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState49, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState49, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState49, CCastExpressionType}:            _CGotoState104Action,
	{_CState49, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState49, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState49, CShiftExpressionType}:           _CGotoState123Action,
	{_CState49, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState49, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState49, CAndExpressionType}:             _CGotoState102Action,
	{_CState49, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState49, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState49, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState49, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState49, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState49, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState49, CExpressionType}:                _CGotoState110Action,
	{_CState49, CDeclarationType}:               _CGotoState51Action,
	{_CState49, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState49, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState49, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState49, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState49, CStructOrUnionType}:             _CGotoState35Action,
	{_CState49, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState49, CTypeQualifierType}:             _CGotoState37Action,
	{_CState49, CStatementType}:                 _CGotoState124Action,
	{_CState49, CLabeledStatementType}:          _CGotoState115Action,
	{_CState49, CCompoundStatementType}:         _CGotoState105Action,
	{_CState49, CDeclarationListType}:           _CGotoState107Action,
	{_CState49, CStatementListType}:             _CGotoState125Action,
	{_CState49, CExpressionStatementType}:       _CGotoState111Action,
	{_CState49, CSelectionStatementType}:        _CGotoState122Action,
	{_CState49, CIterationStatementType}:        _CGotoState113Action,
	{_CState49, CJumpStatementType}:             _CGotoState114Action,
	{_CState52, CTypeNameToken}:                 _CGotoState21Action,
	{_CState52, CTypedefToken}:                  _CGotoState20Action,
	{_CState52, CExternToken}:                   _CGotoState8Action,
	{_CState52, CStaticToken}:                   _CGotoState18Action,
	{_CState52, CAutoToken}:                     _CGotoState3Action,
	{_CState52, CRegisterToken}:                 _CGotoState15Action,
	{_CState52, CCharToken}:                     _CGotoState4Action,
	{_CState52, CShortToken}:                    _CGotoState16Action,
	{_CState52, CIntToken}:                      _CGotoState11Action,
	{_CState52, CLongToken}:                     _CGotoState12Action,
	{_CState52, CSignedToken}:                   _CGotoState17Action,
	{_CState52, CUnsignedToken}:                 _CGotoState23Action,
	{_CState52, CFloatToken}:                    _CGotoState9Action,
	{_CState52, CDoubleToken}:                   _CGotoState6Action,
	{_CState52, CConstToken}:                    _CGotoState5Action,
	{_CState52, CVolatileToken}:                 _CGotoState25Action,
	{_CState52, CVoidToken}:                     _CGotoState24Action,
	{_CState52, CStructToken}:                   _CGotoState19Action,
	{_CState52, CUnionToken}:                    _CGotoState22Action,
	{_CState52, CEnumToken}:                     _CGotoState7Action,
	{_CState52, CLcurlToken}:                    _CGotoState49Action,
	{_CState52, CDeclarationType}:               _CGotoState129Action,
	{_CState52, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState52, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState52, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState52, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState52, CStructOrUnionType}:             _CGotoState35Action,
	{_CState52, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState52, CTypeQualifierType}:             _CGotoState37Action,
	{_CState52, CCompoundStatementType}:         _CGotoState128Action,
	{_CState53, CIdentifierToken}:               _CGotoState10Action,
	{_CState53, CLparamToken}:                   _CGotoState13Action,
	{_CState53, CSemicolonToken}:                _CGotoState45Action,
	{_CState53, CMulToken}:                      _CGotoState14Action,
	{_CState53, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState53, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState53, CDeclaratorType}:                _CGotoState130Action,
	{_CState53, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState53, CPointerType}:                   _CGotoState33Action,
	{_CState54, CIdentifierToken}:               _CGotoState131Action,
	{_CState54, CConstantToken}:                 _CGotoState78Action,
	{_CState54, CStringLiteralToken}:            _CGotoState97Action,
	{_CState54, CSizeofToken}:                   _CGotoState96Action,
	{_CState54, CIncOpToken}:                    _CGotoState88Action,
	{_CState54, CDecOpToken}:                    _CGotoState80Action,
	{_CState54, CLparamToken}:                   _CGotoState89Action,
	{_CState54, CRbraceToken}:                   _CGotoState132Action,
	{_CState54, CMulToken}:                      _CGotoState91Action,
	{_CState54, CMinusToken}:                    _CGotoState90Action,
	{_CState54, CPlusToken}:                     _CGotoState92Action,
	{_CState54, CAndToken}:                      _CGotoState75Action,
	{_CState54, CExclaimToken}:                  _CGotoState83Action,
	{_CState54, CTildaToken}:                    _CGotoState99Action,
	{_CState54, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState54, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState54, CUnaryExpressionType}:           _CGotoState135Action,
	{_CState54, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState54, CCastExpressionType}:            _CGotoState104Action,
	{_CState54, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState54, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState54, CShiftExpressionType}:           _CGotoState123Action,
	{_CState54, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState54, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState54, CAndExpressionType}:             _CGotoState102Action,
	{_CState54, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState54, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState54, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState54, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState54, CConditionalExpressionType}:     _CGotoState133Action,
	{_CState54, CConstantExpressionType}:        _CGotoState134Action,
	{_CState55, CIdentifierToken}:               _CGotoState136Action,
	{_CState55, CTypeNameToken}:                 _CGotoState21Action,
	{_CState55, CTypedefToken}:                  _CGotoState20Action,
	{_CState55, CExternToken}:                   _CGotoState8Action,
	{_CState55, CStaticToken}:                   _CGotoState18Action,
	{_CState55, CAutoToken}:                     _CGotoState3Action,
	{_CState55, CRegisterToken}:                 _CGotoState15Action,
	{_CState55, CCharToken}:                     _CGotoState4Action,
	{_CState55, CShortToken}:                    _CGotoState16Action,
	{_CState55, CIntToken}:                      _CGotoState11Action,
	{_CState55, CLongToken}:                     _CGotoState12Action,
	{_CState55, CSignedToken}:                   _CGotoState17Action,
	{_CState55, CUnsignedToken}:                 _CGotoState23Action,
	{_CState55, CFloatToken}:                    _CGotoState9Action,
	{_CState55, CDoubleToken}:                   _CGotoState6Action,
	{_CState55, CConstToken}:                    _CGotoState5Action,
	{_CState55, CVolatileToken}:                 _CGotoState25Action,
	{_CState55, CVoidToken}:                     _CGotoState24Action,
	{_CState55, CStructToken}:                   _CGotoState19Action,
	{_CState55, CUnionToken}:                    _CGotoState22Action,
	{_CState55, CEnumToken}:                     _CGotoState7Action,
	{_CState55, CRparamToken}:                   _CGotoState137Action,
	{_CState55, CDeclarationSpecifiersType}:     _CGotoState138Action,
	{_CState55, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState55, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState55, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState55, CStructOrUnionType}:             _CGotoState35Action,
	{_CState55, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState55, CTypeQualifierType}:             _CGotoState37Action,
	{_CState55, CParameterTypeListType}:         _CGotoState142Action,
	{_CState55, CParameterListType}:             _CGotoState141Action,
	{_CState55, CParameterDeclarationType}:      _CGotoState140Action,
	{_CState55, CIdentifierListType}:            _CGotoState139Action,
	{_CState56, CLparamToken}:                   _CGotoState55Action,
	{_CState56, CLbraceToken}:                   _CGotoState54Action,
	{_CState58, CLcurlToken}:                    _CGotoState143Action,
	{_CState59, CTypeNameToken}:                 _CGotoState21Action,
	{_CState59, CCharToken}:                     _CGotoState4Action,
	{_CState59, CShortToken}:                    _CGotoState16Action,
	{_CState59, CIntToken}:                      _CGotoState11Action,
	{_CState59, CLongToken}:                     _CGotoState12Action,
	{_CState59, CSignedToken}:                   _CGotoState17Action,
	{_CState59, CUnsignedToken}:                 _CGotoState23Action,
	{_CState59, CFloatToken}:                    _CGotoState9Action,
	{_CState59, CDoubleToken}:                   _CGotoState6Action,
	{_CState59, CConstToken}:                    _CGotoState5Action,
	{_CState59, CVolatileToken}:                 _CGotoState25Action,
	{_CState59, CVoidToken}:                     _CGotoState24Action,
	{_CState59, CStructToken}:                   _CGotoState19Action,
	{_CState59, CUnionToken}:                    _CGotoState22Action,
	{_CState59, CEnumToken}:                     _CGotoState7Action,
	{_CState59, CTypeSpecifierType}:             _CGotoState148Action,
	{_CState59, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState59, CStructOrUnionType}:             _CGotoState35Action,
	{_CState59, CStructDeclarationListType}:     _CGotoState146Action,
	{_CState59, CStructDeclarationType}:         _CGotoState145Action,
	{_CState59, CSpecifierQualifierListType}:    _CGotoState144Action,
	{_CState59, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState59, CTypeQualifierType}:             _CGotoState147Action,
	{_CState63, CIdentifierToken}:               _CGotoState64Action,
	{_CState63, CEnumeratorListType}:            _CGotoState149Action,
	{_CState63, CEnumeratorType}:                _CGotoState65Action,
	{_CState64, CEqToken}:                       _CGotoState150Action,
	{_CState66, CRcurlToken}:                    _CGotoState152Action,
	{_CState66, CCommaToken}:                    _CGotoState151Action,
	{_CState70, CIdentifierToken}:               _CGotoState131Action,
	{_CState70, CConstantToken}:                 _CGotoState78Action,
	{_CState70, CStringLiteralToken}:            _CGotoState97Action,
	{_CState70, CSizeofToken}:                   _CGotoState96Action,
	{_CState70, CIncOpToken}:                    _CGotoState88Action,
	{_CState70, CDecOpToken}:                    _CGotoState80Action,
	{_CState70, CLparamToken}:                   _CGotoState89Action,
	{_CState70, CLcurlToken}:                    _CGotoState153Action,
	{_CState70, CMulToken}:                      _CGotoState91Action,
	{_CState70, CMinusToken}:                    _CGotoState90Action,
	{_CState70, CPlusToken}:                     _CGotoState92Action,
	{_CState70, CAndToken}:                      _CGotoState75Action,
	{_CState70, CExclaimToken}:                  _CGotoState83Action,
	{_CState70, CTildaToken}:                    _CGotoState99Action,
	{_CState70, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState70, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState70, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState70, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState70, CCastExpressionType}:            _CGotoState104Action,
	{_CState70, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState70, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState70, CShiftExpressionType}:           _CGotoState123Action,
	{_CState70, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState70, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState70, CAndExpressionType}:             _CGotoState102Action,
	{_CState70, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState70, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState70, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState70, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState70, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState70, CAssignmentExpressionType}:      _CGotoState154Action,
	{_CState70, CInitializerType}:               _CGotoState155Action,
	{_CState72, CTypeNameToken}:                 _CGotoState21Action,
	{_CState72, CTypedefToken}:                  _CGotoState20Action,
	{_CState72, CExternToken}:                   _CGotoState8Action,
	{_CState72, CStaticToken}:                   _CGotoState18Action,
	{_CState72, CAutoToken}:                     _CGotoState3Action,
	{_CState72, CRegisterToken}:                 _CGotoState15Action,
	{_CState72, CCharToken}:                     _CGotoState4Action,
	{_CState72, CShortToken}:                    _CGotoState16Action,
	{_CState72, CIntToken}:                      _CGotoState11Action,
	{_CState72, CLongToken}:                     _CGotoState12Action,
	{_CState72, CSignedToken}:                   _CGotoState17Action,
	{_CState72, CUnsignedToken}:                 _CGotoState23Action,
	{_CState72, CFloatToken}:                    _CGotoState9Action,
	{_CState72, CDoubleToken}:                   _CGotoState6Action,
	{_CState72, CConstToken}:                    _CGotoState5Action,
	{_CState72, CVolatileToken}:                 _CGotoState25Action,
	{_CState72, CVoidToken}:                     _CGotoState24Action,
	{_CState72, CStructToken}:                   _CGotoState19Action,
	{_CState72, CUnionToken}:                    _CGotoState22Action,
	{_CState72, CEnumToken}:                     _CGotoState7Action,
	{_CState72, CLcurlToken}:                    _CGotoState49Action,
	{_CState72, CDeclarationType}:               _CGotoState129Action,
	{_CState72, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState72, CStorageClassSpecifierType}:     _CGotoState34Action,
	{_CState72, CTypeSpecifierType}:             _CGotoState38Action,
	{_CState72, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState72, CStructOrUnionType}:             _CGotoState35Action,
	{_CState72, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState72, CTypeQualifierType}:             _CGotoState37Action,
	{_CState72, CCompoundStatementType}:         _CGotoState156Action,
	{_CState73, CIdentifierToken}:               _CGotoState10Action,
	{_CState73, CLparamToken}:                   _CGotoState13Action,
	{_CState73, CMulToken}:                      _CGotoState14Action,
	{_CState73, CInitDeclaratorType}:            _CGotoState157Action,
	{_CState73, CDeclaratorType}:                _CGotoState130Action,
	{_CState73, CDirectDeclaratorType}:          _CGotoState29Action,
	{_CState73, CPointerType}:                   _CGotoState33Action,
	{_CState76, CSemicolonToken}:                _CGotoState158Action,
	{_CState77, CIdentifierToken}:               _CGotoState131Action,
	{_CState77, CConstantToken}:                 _CGotoState78Action,
	{_CState77, CStringLiteralToken}:            _CGotoState97Action,
	{_CState77, CSizeofToken}:                   _CGotoState96Action,
	{_CState77, CIncOpToken}:                    _CGotoState88Action,
	{_CState77, CDecOpToken}:                    _CGotoState80Action,
	{_CState77, CLparamToken}:                   _CGotoState89Action,
	{_CState77, CMulToken}:                      _CGotoState91Action,
	{_CState77, CMinusToken}:                    _CGotoState90Action,
	{_CState77, CPlusToken}:                     _CGotoState92Action,
	{_CState77, CAndToken}:                      _CGotoState75Action,
	{_CState77, CExclaimToken}:                  _CGotoState83Action,
	{_CState77, CTildaToken}:                    _CGotoState99Action,
	{_CState77, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState77, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState77, CUnaryExpressionType}:           _CGotoState135Action,
	{_CState77, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState77, CCastExpressionType}:            _CGotoState104Action,
	{_CState77, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState77, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState77, CShiftExpressionType}:           _CGotoState123Action,
	{_CState77, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState77, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState77, CAndExpressionType}:             _CGotoState102Action,
	{_CState77, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState77, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState77, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState77, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState77, CConditionalExpressionType}:     _CGotoState133Action,
	{_CState77, CConstantExpressionType}:        _CGotoState159Action,
	{_CState79, CSemicolonToken}:                _CGotoState160Action,
	{_CState80, CIdentifierToken}:               _CGotoState131Action,
	{_CState80, CConstantToken}:                 _CGotoState78Action,
	{_CState80, CStringLiteralToken}:            _CGotoState97Action,
	{_CState80, CSizeofToken}:                   _CGotoState96Action,
	{_CState80, CIncOpToken}:                    _CGotoState88Action,
	{_CState80, CDecOpToken}:                    _CGotoState80Action,
	{_CState80, CLparamToken}:                   _CGotoState161Action,
	{_CState80, CMulToken}:                      _CGotoState91Action,
	{_CState80, CMinusToken}:                    _CGotoState90Action,
	{_CState80, CPlusToken}:                     _CGotoState92Action,
	{_CState80, CAndToken}:                      _CGotoState75Action,
	{_CState80, CExclaimToken}:                  _CGotoState83Action,
	{_CState80, CTildaToken}:                    _CGotoState99Action,
	{_CState80, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState80, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState80, CUnaryExpressionType}:           _CGotoState162Action,
	{_CState80, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState81, CColonToken}:                    _CGotoState163Action,
	{_CState82, CIdentifierToken}:               _CGotoState86Action,
	{_CState82, CConstantToken}:                 _CGotoState78Action,
	{_CState82, CStringLiteralToken}:            _CGotoState97Action,
	{_CState82, CSizeofToken}:                   _CGotoState96Action,
	{_CState82, CIncOpToken}:                    _CGotoState88Action,
	{_CState82, CDecOpToken}:                    _CGotoState80Action,
	{_CState82, CCaseToken}:                     _CGotoState77Action,
	{_CState82, CDefaultToken}:                  _CGotoState81Action,
	{_CState82, CIfToken}:                       _CGotoState87Action,
	{_CState82, CSwitchToken}:                   _CGotoState98Action,
	{_CState82, CWhileToken}:                    _CGotoState100Action,
	{_CState82, CDoToken}:                       _CGotoState82Action,
	{_CState82, CForToken}:                      _CGotoState84Action,
	{_CState82, CGotoToken}:                     _CGotoState85Action,
	{_CState82, CContinueToken}:                 _CGotoState79Action,
	{_CState82, CBreakToken}:                    _CGotoState76Action,
	{_CState82, CReturnToken}:                   _CGotoState94Action,
	{_CState82, CLparamToken}:                   _CGotoState89Action,
	{_CState82, CLcurlToken}:                    _CGotoState49Action,
	{_CState82, CSemicolonToken}:                _CGotoState95Action,
	{_CState82, CMulToken}:                      _CGotoState91Action,
	{_CState82, CMinusToken}:                    _CGotoState90Action,
	{_CState82, CPlusToken}:                     _CGotoState92Action,
	{_CState82, CAndToken}:                      _CGotoState75Action,
	{_CState82, CExclaimToken}:                  _CGotoState83Action,
	{_CState82, CTildaToken}:                    _CGotoState99Action,
	{_CState82, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState82, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState82, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState82, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState82, CCastExpressionType}:            _CGotoState104Action,
	{_CState82, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState82, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState82, CShiftExpressionType}:           _CGotoState123Action,
	{_CState82, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState82, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState82, CAndExpressionType}:             _CGotoState102Action,
	{_CState82, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState82, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState82, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState82, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState82, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState82, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState82, CExpressionType}:                _CGotoState110Action,
	{_CState82, CStatementType}:                 _CGotoState164Action,
	{_CState82, CLabeledStatementType}:          _CGotoState115Action,
	{_CState82, CCompoundStatementType}:         _CGotoState105Action,
	{_CState82, CExpressionStatementType}:       _CGotoState111Action,
	{_CState82, CSelectionStatementType}:        _CGotoState122Action,
	{_CState82, CIterationStatementType}:        _CGotoState113Action,
	{_CState82, CJumpStatementType}:             _CGotoState114Action,
	{_CState84, CLparamToken}:                   _CGotoState165Action,
	{_CState85, CIdentifierToken}:               _CGotoState166Action,
	{_CState86, CColonToken}:                    _CGotoState167Action,
	{_CState87, CLparamToken}:                   _CGotoState168Action,
	{_CState88, CIdentifierToken}:               _CGotoState131Action,
	{_CState88, CConstantToken}:                 _CGotoState78Action,
	{_CState88, CStringLiteralToken}:            _CGotoState97Action,
	{_CState88, CSizeofToken}:                   _CGotoState96Action,
	{_CState88, CIncOpToken}:                    _CGotoState88Action,
	{_CState88, CDecOpToken}:                    _CGotoState80Action,
	{_CState88, CLparamToken}:                   _CGotoState161Action,
	{_CState88, CMulToken}:                      _CGotoState91Action,
	{_CState88, CMinusToken}:                    _CGotoState90Action,
	{_CState88, CPlusToken}:                     _CGotoState92Action,
	{_CState88, CAndToken}:                      _CGotoState75Action,
	{_CState88, CExclaimToken}:                  _CGotoState83Action,
	{_CState88, CTildaToken}:                    _CGotoState99Action,
	{_CState88, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState88, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState88, CUnaryExpressionType}:           _CGotoState169Action,
	{_CState88, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState89, CIdentifierToken}:               _CGotoState131Action,
	{_CState89, CConstantToken}:                 _CGotoState78Action,
	{_CState89, CStringLiteralToken}:            _CGotoState97Action,
	{_CState89, CSizeofToken}:                   _CGotoState96Action,
	{_CState89, CIncOpToken}:                    _CGotoState88Action,
	{_CState89, CDecOpToken}:                    _CGotoState80Action,
	{_CState89, CTypeNameToken}:                 _CGotoState21Action,
	{_CState89, CCharToken}:                     _CGotoState4Action,
	{_CState89, CShortToken}:                    _CGotoState16Action,
	{_CState89, CIntToken}:                      _CGotoState11Action,
	{_CState89, CLongToken}:                     _CGotoState12Action,
	{_CState89, CSignedToken}:                   _CGotoState17Action,
	{_CState89, CUnsignedToken}:                 _CGotoState23Action,
	{_CState89, CFloatToken}:                    _CGotoState9Action,
	{_CState89, CDoubleToken}:                   _CGotoState6Action,
	{_CState89, CConstToken}:                    _CGotoState5Action,
	{_CState89, CVolatileToken}:                 _CGotoState25Action,
	{_CState89, CVoidToken}:                     _CGotoState24Action,
	{_CState89, CStructToken}:                   _CGotoState19Action,
	{_CState89, CUnionToken}:                    _CGotoState22Action,
	{_CState89, CEnumToken}:                     _CGotoState7Action,
	{_CState89, CLparamToken}:                   _CGotoState89Action,
	{_CState89, CMulToken}:                      _CGotoState91Action,
	{_CState89, CMinusToken}:                    _CGotoState90Action,
	{_CState89, CPlusToken}:                     _CGotoState92Action,
	{_CState89, CAndToken}:                      _CGotoState75Action,
	{_CState89, CExclaimToken}:                  _CGotoState83Action,
	{_CState89, CTildaToken}:                    _CGotoState99Action,
	{_CState89, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState89, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState89, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState89, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState89, CCastExpressionType}:            _CGotoState104Action,
	{_CState89, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState89, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState89, CShiftExpressionType}:           _CGotoState123Action,
	{_CState89, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState89, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState89, CAndExpressionType}:             _CGotoState102Action,
	{_CState89, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState89, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState89, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState89, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState89, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState89, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState89, CExpressionType}:                _CGotoState170Action,
	{_CState89, CTypeSpecifierType}:             _CGotoState148Action,
	{_CState89, CStructOrUnionSpecifierType}:    _CGotoState36Action,
	{_CState89, CStructOrUnionType}:             _CGotoState35Action,
	{_CState89, CSpecifierQualifierListType}:    _CGotoState171Action,
	{_CState89, CEnumSpecifierType}:             _CGotoState30Action,
	{_CState89, CTypeQualifierType}:             _CGotoState147Action,
	{_CState89, CTypeNameType}:                  _CGotoState172Action,
	{_CState94, CIdentifierToken}:               _CGotoState131Action,
	{_CState94, CConstantToken}:                 _CGotoState78Action,
	{_CState94, CStringLiteralToken}:            _CGotoState97Action,
	{_CState94, CSizeofToken}:                   _CGotoState96Action,
	{_CState94, CIncOpToken}:                    _CGotoState88Action,
	{_CState94, CDecOpToken}:                    _CGotoState80Action,
	{_CState94, CLparamToken}:                   _CGotoState89Action,
	{_CState94, CSemicolonToken}:                _CGotoState173Action,
	{_CState94, CMulToken}:                      _CGotoState91Action,
	{_CState94, CMinusToken}:                    _CGotoState90Action,
	{_CState94, CPlusToken}:                     _CGotoState92Action,
	{_CState94, CAndToken}:                      _CGotoState75Action,
	{_CState94, CExclaimToken}:                  _CGotoState83Action,
	{_CState94, CTildaToken}:                    _CGotoState99Action,
	{_CState94, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState94, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState94, CUnaryExpressionType}:           _CGotoState126Action,
	{_CState94, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState94, CCastExpressionType}:            _CGotoState104Action,
	{_CState94, CMultiplicativeExpressionType}:  _CGotoState118Action,
	{_CState94, CAdditiveExpressionType}:        _CGotoState101Action,
	{_CState94, CShiftExpressionType}:           _CGotoState123Action,
	{_CState94, CRelationalExpressionType}:      _CGotoState121Action,
	{_CState94, CEqualityExpressionType}:        _CGotoState108Action,
	{_CState94, CAndExpressionType}:             _CGotoState102Action,
	{_CState94, CExclusiveOrExpressionType}:     _CGotoState109Action,
	{_CState94, CInclusiveOrExpressionType}:     _CGotoState112Action,
	{_CState94, CLogicalAndExpressionType}:      _CGotoState116Action,
	{_CState94, CLogicalOrExpressionType}:       _CGotoState117Action,
	{_CState94, CConditionalExpressionType}:     _CGotoState106Action,
	{_CState94, CAssignmentExpressionType}:      _CGotoState103Action,
	{_CState94, CExpressionType}:                _CGotoState174Action,
	{_CState96, CIdentifierToken}:               _CGotoState131Action,
	{_CState96, CConstantToken}:                 _CGotoState78Action,
	{_CState96, CStringLiteralToken}:            _CGotoState97Action,
	{_CState96, CSizeofToken}:                   _CGotoState96Action,
	{_CState96, CIncOpToken}:                    _CGotoState88Action,
	{_CState96, CDecOpToken}:                    _CGotoState80Action,
	{_CState96, CLparamToken}:                   _CGotoState175Action,
	{_CState96, CMulToken}:                      _CGotoState91Action,
	{_CState96, CMinusToken}:                    _CGotoState90Action,
	{_CState96, CPlusToken}:                     _CGotoState92Action,
	{_CState96, CAndToken}:                      _CGotoState75Action,
	{_CState96, CExclaimToken}:                  _CGotoState83Action,
	{_CState96, CTildaToken}:                    _CGotoState99Action,
	{_CState96, CPrimaryExpressionType}:         _CGotoState120Action,
	{_CState96, CPostfixExpressionType}:         _CGotoState119Action,
	{_CState96, CUnaryExpressionType}:           _CGotoState176Action,
	{_CState96, CUnaryOperatorType}:             _CGotoState127Action,
	{_CState98, CLparamToken}:                   _CGotoState177Action,
	{_CState100, CLparamToken}:                  _CGotoState178Action,
	{_CState101, CMinusToken}:                   _CGotoState179Action,
	{_CState101, CPlusToken}:                    _CGotoState180Action,
	{_CState102, CAndToken}:                     _CGotoState181Action,
	{_CState107, CIdentifierToken}:              _CGotoState86Action,
	{_CState107, CConstantToken}:                _CGotoState78Action,
	{_CState107, CStringLiteralToken}:           _CGotoState97Action,
	{_CState107, CSizeofToken}:                  _CGotoState96Action,
	{_CState107, CIncOpToken}:                   _CGotoState88Action,
	{_CState107, CDecOpToken}:                   _CGotoState80Action,
	{_CState107, CTypeNameToken}:                _CGotoState21Action,
	{_CState107, CTypedefToken}:                 _CGotoState20Action,
	{_CState107, CExternToken}:                  _CGotoState8Action,
	{_CState107, CStaticToken}:                  _CGotoState18Action,
	{_CState107, CAutoToken}:                    _CGotoState3Action,
	{_CState107, CRegisterToken}:                _CGotoState15Action,
	{_CState107, CCharToken}:                    _CGotoState4Action,
	{_CState107, CShortToken}:                   _CGotoState16Action,
	{_CState107, CIntToken}:                     _CGotoState11Action,
	{_CState107, CLongToken}:                    _CGotoState12Action,
	{_CState107, CSignedToken}:                  _CGotoState17Action,
	{_CState107, CUnsignedToken}:                _CGotoState23Action,
	{_CState107, CFloatToken}:                   _CGotoState9Action,
	{_CState107, CDoubleToken}:                  _CGotoState6Action,
	{_CState107, CConstToken}:                   _CGotoState5Action,
	{_CState107, CVolatileToken}:                _CGotoState25Action,
	{_CState107, CVoidToken}:                    _CGotoState24Action,
	{_CState107, CStructToken}:                  _CGotoState19Action,
	{_CState107, CUnionToken}:                   _CGotoState22Action,
	{_CState107, CEnumToken}:                    _CGotoState7Action,
	{_CState107, CCaseToken}:                    _CGotoState77Action,
	{_CState107, CDefaultToken}:                 _CGotoState81Action,
	{_CState107, CIfToken}:                      _CGotoState87Action,
	{_CState107, CSwitchToken}:                  _CGotoState98Action,
	{_CState107, CWhileToken}:                   _CGotoState100Action,
	{_CState107, CDoToken}:                      _CGotoState82Action,
	{_CState107, CForToken}:                     _CGotoState84Action,
	{_CState107, CGotoToken}:                    _CGotoState85Action,
	{_CState107, CContinueToken}:                _CGotoState79Action,
	{_CState107, CBreakToken}:                   _CGotoState76Action,
	{_CState107, CReturnToken}:                  _CGotoState94Action,
	{_CState107, CLparamToken}:                  _CGotoState89Action,
	{_CState107, CLcurlToken}:                   _CGotoState49Action,
	{_CState107, CRcurlToken}:                   _CGotoState182Action,
	{_CState107, CSemicolonToken}:               _CGotoState95Action,
	{_CState107, CMulToken}:                     _CGotoState91Action,
	{_CState107, CMinusToken}:                   _CGotoState90Action,
	{_CState107, CPlusToken}:                    _CGotoState92Action,
	{_CState107, CAndToken}:                     _CGotoState75Action,
	{_CState107, CExclaimToken}:                 _CGotoState83Action,
	{_CState107, CTildaToken}:                   _CGotoState99Action,
	{_CState107, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState107, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState107, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState107, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState107, CCastExpressionType}:           _CGotoState104Action,
	{_CState107, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState107, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState107, CShiftExpressionType}:          _CGotoState123Action,
	{_CState107, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState107, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState107, CAndExpressionType}:            _CGotoState102Action,
	{_CState107, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState107, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState107, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState107, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState107, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState107, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState107, CExpressionType}:               _CGotoState110Action,
	{_CState107, CDeclarationType}:              _CGotoState129Action,
	{_CState107, CDeclarationSpecifiersType}:    _CGotoState53Action,
	{_CState107, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState107, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState107, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState107, CStructOrUnionType}:            _CGotoState35Action,
	{_CState107, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState107, CTypeQualifierType}:            _CGotoState37Action,
	{_CState107, CStatementType}:                _CGotoState124Action,
	{_CState107, CLabeledStatementType}:         _CGotoState115Action,
	{_CState107, CCompoundStatementType}:        _CGotoState105Action,
	{_CState107, CStatementListType}:            _CGotoState183Action,
	{_CState107, CExpressionStatementType}:      _CGotoState111Action,
	{_CState107, CSelectionStatementType}:       _CGotoState122Action,
	{_CState107, CIterationStatementType}:       _CGotoState113Action,
	{_CState107, CJumpStatementType}:            _CGotoState114Action,
	{_CState108, CEqOpToken}:                    _CGotoState184Action,
	{_CState108, CNeOpToken}:                    _CGotoState185Action,
	{_CState109, CHatToken}:                     _CGotoState186Action,
	{_CState110, CSemicolonToken}:               _CGotoState188Action,
	{_CState110, CCommaToken}:                   _CGotoState187Action,
	{_CState112, COrToken}:                      _CGotoState189Action,
	{_CState116, CAndOpToken}:                   _CGotoState190Action,
	{_CState117, COrOpToken}:                    _CGotoState191Action,
	{_CState117, CQuestionToken}:                _CGotoState192Action,
	{_CState118, CMulToken}:                     _CGotoState195Action,
	{_CState118, CDivToken}:                     _CGotoState193Action,
	{_CState118, CModToken}:                     _CGotoState194Action,
	{_CState119, CPtrOpToken}:                   _CGotoState201Action,
	{_CState119, CIncOpToken}:                   _CGotoState198Action,
	{_CState119, CDecOpToken}:                   _CGotoState196Action,
	{_CState119, CLparamToken}:                  _CGotoState200Action,
	{_CState119, CLbraceToken}:                  _CGotoState199Action,
	{_CState119, CDotToken}:                     _CGotoState197Action,
	{_CState121, CLeOpToken}:                    _CGotoState204Action,
	{_CState121, CGeOpToken}:                    _CGotoState202Action,
	{_CState121, CLtToken}:                      _CGotoState205Action,
	{_CState121, CGtToken}:                      _CGotoState203Action,
	{_CState123, CLeftOpToken}:                  _CGotoState206Action,
	{_CState123, CRightOpToken}:                 _CGotoState207Action,
	{_CState125, CIdentifierToken}:              _CGotoState86Action,
	{_CState125, CConstantToken}:                _CGotoState78Action,
	{_CState125, CStringLiteralToken}:           _CGotoState97Action,
	{_CState125, CSizeofToken}:                  _CGotoState96Action,
	{_CState125, CIncOpToken}:                   _CGotoState88Action,
	{_CState125, CDecOpToken}:                   _CGotoState80Action,
	{_CState125, CCaseToken}:                    _CGotoState77Action,
	{_CState125, CDefaultToken}:                 _CGotoState81Action,
	{_CState125, CIfToken}:                      _CGotoState87Action,
	{_CState125, CSwitchToken}:                  _CGotoState98Action,
	{_CState125, CWhileToken}:                   _CGotoState100Action,
	{_CState125, CDoToken}:                      _CGotoState82Action,
	{_CState125, CForToken}:                     _CGotoState84Action,
	{_CState125, CGotoToken}:                    _CGotoState85Action,
	{_CState125, CContinueToken}:                _CGotoState79Action,
	{_CState125, CBreakToken}:                   _CGotoState76Action,
	{_CState125, CReturnToken}:                  _CGotoState94Action,
	{_CState125, CLparamToken}:                  _CGotoState89Action,
	{_CState125, CLcurlToken}:                   _CGotoState49Action,
	{_CState125, CRcurlToken}:                   _CGotoState208Action,
	{_CState125, CSemicolonToken}:               _CGotoState95Action,
	{_CState125, CMulToken}:                     _CGotoState91Action,
	{_CState125, CMinusToken}:                   _CGotoState90Action,
	{_CState125, CPlusToken}:                    _CGotoState92Action,
	{_CState125, CAndToken}:                     _CGotoState75Action,
	{_CState125, CExclaimToken}:                 _CGotoState83Action,
	{_CState125, CTildaToken}:                   _CGotoState99Action,
	{_CState125, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState125, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState125, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState125, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState125, CCastExpressionType}:           _CGotoState104Action,
	{_CState125, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState125, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState125, CShiftExpressionType}:          _CGotoState123Action,
	{_CState125, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState125, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState125, CAndExpressionType}:            _CGotoState102Action,
	{_CState125, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState125, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState125, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState125, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState125, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState125, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState125, CExpressionType}:               _CGotoState110Action,
	{_CState125, CStatementType}:                _CGotoState209Action,
	{_CState125, CLabeledStatementType}:         _CGotoState115Action,
	{_CState125, CCompoundStatementType}:        _CGotoState105Action,
	{_CState125, CExpressionStatementType}:      _CGotoState111Action,
	{_CState125, CSelectionStatementType}:       _CGotoState122Action,
	{_CState125, CIterationStatementType}:       _CGotoState113Action,
	{_CState125, CJumpStatementType}:            _CGotoState114Action,
	{_CState126, CMulAssignToken}:               _CGotoState216Action,
	{_CState126, CDivAssignToken}:               _CGotoState212Action,
	{_CState126, CModAssignToken}:               _CGotoState215Action,
	{_CState126, CAddAssignToken}:               _CGotoState210Action,
	{_CState126, CSubAssignToken}:               _CGotoState219Action,
	{_CState126, CLeftAssignToken}:              _CGotoState214Action,
	{_CState126, CRightAssignToken}:             _CGotoState218Action,
	{_CState126, CAndAssignToken}:               _CGotoState211Action,
	{_CState126, CXorAssignToken}:               _CGotoState220Action,
	{_CState126, COrAssignToken}:                _CGotoState217Action,
	{_CState126, CEqToken}:                      _CGotoState213Action,
	{_CState126, CAssignmentOperatorType}:       _CGotoState221Action,
	{_CState127, CIdentifierToken}:              _CGotoState131Action,
	{_CState127, CConstantToken}:                _CGotoState78Action,
	{_CState127, CStringLiteralToken}:           _CGotoState97Action,
	{_CState127, CSizeofToken}:                  _CGotoState96Action,
	{_CState127, CIncOpToken}:                   _CGotoState88Action,
	{_CState127, CDecOpToken}:                   _CGotoState80Action,
	{_CState127, CLparamToken}:                  _CGotoState89Action,
	{_CState127, CMulToken}:                     _CGotoState91Action,
	{_CState127, CMinusToken}:                   _CGotoState90Action,
	{_CState127, CPlusToken}:                    _CGotoState92Action,
	{_CState127, CAndToken}:                     _CGotoState75Action,
	{_CState127, CExclaimToken}:                 _CGotoState83Action,
	{_CState127, CTildaToken}:                   _CGotoState99Action,
	{_CState127, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState127, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState127, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState127, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState127, CCastExpressionType}:           _CGotoState222Action,
	{_CState130, CEqToken}:                      _CGotoState70Action,
	{_CState134, CRbraceToken}:                  _CGotoState223Action,
	{_CState138, CIdentifierToken}:              _CGotoState10Action,
	{_CState138, CLparamToken}:                  _CGotoState225Action,
	{_CState138, CLbraceToken}:                  _CGotoState224Action,
	{_CState138, CMulToken}:                     _CGotoState14Action,
	{_CState138, CDeclaratorType}:               _CGotoState227Action,
	{_CState138, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState138, CPointerType}:                  _CGotoState229Action,
	{_CState138, CAbstractDeclaratorType}:       _CGotoState226Action,
	{_CState138, CDirectAbstractDeclaratorType}: _CGotoState228Action,
	{_CState139, CRparamToken}:                  _CGotoState231Action,
	{_CState139, CCommaToken}:                   _CGotoState230Action,
	{_CState141, CCommaToken}:                   _CGotoState232Action,
	{_CState142, CRparamToken}:                  _CGotoState233Action,
	{_CState143, CTypeNameToken}:                _CGotoState21Action,
	{_CState143, CCharToken}:                    _CGotoState4Action,
	{_CState143, CShortToken}:                   _CGotoState16Action,
	{_CState143, CIntToken}:                     _CGotoState11Action,
	{_CState143, CLongToken}:                    _CGotoState12Action,
	{_CState143, CSignedToken}:                  _CGotoState17Action,
	{_CState143, CUnsignedToken}:                _CGotoState23Action,
	{_CState143, CFloatToken}:                   _CGotoState9Action,
	{_CState143, CDoubleToken}:                  _CGotoState6Action,
	{_CState143, CConstToken}:                   _CGotoState5Action,
	{_CState143, CVolatileToken}:                _CGotoState25Action,
	{_CState143, CVoidToken}:                    _CGotoState24Action,
	{_CState143, CStructToken}:                  _CGotoState19Action,
	{_CState143, CUnionToken}:                   _CGotoState22Action,
	{_CState143, CEnumToken}:                    _CGotoState7Action,
	{_CState143, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState143, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState143, CStructOrUnionType}:            _CGotoState35Action,
	{_CState143, CStructDeclarationListType}:    _CGotoState234Action,
	{_CState143, CStructDeclarationType}:        _CGotoState145Action,
	{_CState143, CSpecifierQualifierListType}:   _CGotoState144Action,
	{_CState143, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState143, CTypeQualifierType}:            _CGotoState147Action,
	{_CState144, CIdentifierToken}:              _CGotoState10Action,
	{_CState144, CLparamToken}:                  _CGotoState13Action,
	{_CState144, CColonToken}:                   _CGotoState235Action,
	{_CState144, CMulToken}:                     _CGotoState14Action,
	{_CState144, CStructDeclaratorListType}:     _CGotoState238Action,
	{_CState144, CStructDeclaratorType}:         _CGotoState237Action,
	{_CState144, CDeclaratorType}:               _CGotoState236Action,
	{_CState144, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState144, CPointerType}:                  _CGotoState33Action,
	{_CState146, CTypeNameToken}:                _CGotoState21Action,
	{_CState146, CCharToken}:                    _CGotoState4Action,
	{_CState146, CShortToken}:                   _CGotoState16Action,
	{_CState146, CIntToken}:                     _CGotoState11Action,
	{_CState146, CLongToken}:                    _CGotoState12Action,
	{_CState146, CSignedToken}:                  _CGotoState17Action,
	{_CState146, CUnsignedToken}:                _CGotoState23Action,
	{_CState146, CFloatToken}:                   _CGotoState9Action,
	{_CState146, CDoubleToken}:                  _CGotoState6Action,
	{_CState146, CConstToken}:                   _CGotoState5Action,
	{_CState146, CVolatileToken}:                _CGotoState25Action,
	{_CState146, CVoidToken}:                    _CGotoState24Action,
	{_CState146, CStructToken}:                  _CGotoState19Action,
	{_CState146, CUnionToken}:                   _CGotoState22Action,
	{_CState146, CEnumToken}:                    _CGotoState7Action,
	{_CState146, CRcurlToken}:                   _CGotoState239Action,
	{_CState146, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState146, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState146, CStructOrUnionType}:            _CGotoState35Action,
	{_CState146, CStructDeclarationType}:        _CGotoState240Action,
	{_CState146, CSpecifierQualifierListType}:   _CGotoState144Action,
	{_CState146, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState146, CTypeQualifierType}:            _CGotoState147Action,
	{_CState147, CTypeNameToken}:                _CGotoState21Action,
	{_CState147, CCharToken}:                    _CGotoState4Action,
	{_CState147, CShortToken}:                   _CGotoState16Action,
	{_CState147, CIntToken}:                     _CGotoState11Action,
	{_CState147, CLongToken}:                    _CGotoState12Action,
	{_CState147, CSignedToken}:                  _CGotoState17Action,
	{_CState147, CUnsignedToken}:                _CGotoState23Action,
	{_CState147, CFloatToken}:                   _CGotoState9Action,
	{_CState147, CDoubleToken}:                  _CGotoState6Action,
	{_CState147, CConstToken}:                   _CGotoState5Action,
	{_CState147, CVolatileToken}:                _CGotoState25Action,
	{_CState147, CVoidToken}:                    _CGotoState24Action,
	{_CState147, CStructToken}:                  _CGotoState19Action,
	{_CState147, CUnionToken}:                   _CGotoState22Action,
	{_CState147, CEnumToken}:                    _CGotoState7Action,
	{_CState147, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState147, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState147, CStructOrUnionType}:            _CGotoState35Action,
	{_CState147, CSpecifierQualifierListType}:   _CGotoState241Action,
	{_CState147, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState147, CTypeQualifierType}:            _CGotoState147Action,
	{_CState148, CTypeNameToken}:                _CGotoState21Action,
	{_CState148, CCharToken}:                    _CGotoState4Action,
	{_CState148, CShortToken}:                   _CGotoState16Action,
	{_CState148, CIntToken}:                     _CGotoState11Action,
	{_CState148, CLongToken}:                    _CGotoState12Action,
	{_CState148, CSignedToken}:                  _CGotoState17Action,
	{_CState148, CUnsignedToken}:                _CGotoState23Action,
	{_CState148, CFloatToken}:                   _CGotoState9Action,
	{_CState148, CDoubleToken}:                  _CGotoState6Action,
	{_CState148, CConstToken}:                   _CGotoState5Action,
	{_CState148, CVolatileToken}:                _CGotoState25Action,
	{_CState148, CVoidToken}:                    _CGotoState24Action,
	{_CState148, CStructToken}:                  _CGotoState19Action,
	{_CState148, CUnionToken}:                   _CGotoState22Action,
	{_CState148, CEnumToken}:                    _CGotoState7Action,
	{_CState148, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState148, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState148, CStructOrUnionType}:            _CGotoState35Action,
	{_CState148, CSpecifierQualifierListType}:   _CGotoState242Action,
	{_CState148, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState148, CTypeQualifierType}:            _CGotoState147Action,
	{_CState149, CRcurlToken}:                   _CGotoState243Action,
	{_CState149, CCommaToken}:                   _CGotoState151Action,
	{_CState150, CIdentifierToken}:              _CGotoState131Action,
	{_CState150, CConstantToken}:                _CGotoState78Action,
	{_CState150, CStringLiteralToken}:           _CGotoState97Action,
	{_CState150, CSizeofToken}:                  _CGotoState96Action,
	{_CState150, CIncOpToken}:                   _CGotoState88Action,
	{_CState150, CDecOpToken}:                   _CGotoState80Action,
	{_CState150, CLparamToken}:                  _CGotoState89Action,
	{_CState150, CMulToken}:                     _CGotoState91Action,
	{_CState150, CMinusToken}:                   _CGotoState90Action,
	{_CState150, CPlusToken}:                    _CGotoState92Action,
	{_CState150, CAndToken}:                     _CGotoState75Action,
	{_CState150, CExclaimToken}:                 _CGotoState83Action,
	{_CState150, CTildaToken}:                   _CGotoState99Action,
	{_CState150, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState150, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState150, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState150, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState150, CCastExpressionType}:           _CGotoState104Action,
	{_CState150, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState150, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState150, CShiftExpressionType}:          _CGotoState123Action,
	{_CState150, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState150, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState150, CAndExpressionType}:            _CGotoState102Action,
	{_CState150, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState150, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState150, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState150, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState150, CConditionalExpressionType}:    _CGotoState133Action,
	{_CState150, CConstantExpressionType}:       _CGotoState244Action,
	{_CState151, CIdentifierToken}:              _CGotoState64Action,
	{_CState151, CEnumeratorType}:               _CGotoState245Action,
	{_CState153, CIdentifierToken}:              _CGotoState131Action,
	{_CState153, CConstantToken}:                _CGotoState78Action,
	{_CState153, CStringLiteralToken}:           _CGotoState97Action,
	{_CState153, CSizeofToken}:                  _CGotoState96Action,
	{_CState153, CIncOpToken}:                   _CGotoState88Action,
	{_CState153, CDecOpToken}:                   _CGotoState80Action,
	{_CState153, CLparamToken}:                  _CGotoState89Action,
	{_CState153, CLcurlToken}:                   _CGotoState153Action,
	{_CState153, CMulToken}:                     _CGotoState91Action,
	{_CState153, CMinusToken}:                   _CGotoState90Action,
	{_CState153, CPlusToken}:                    _CGotoState92Action,
	{_CState153, CAndToken}:                     _CGotoState75Action,
	{_CState153, CExclaimToken}:                 _CGotoState83Action,
	{_CState153, CTildaToken}:                   _CGotoState99Action,
	{_CState153, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState153, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState153, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState153, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState153, CCastExpressionType}:           _CGotoState104Action,
	{_CState153, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState153, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState153, CShiftExpressionType}:          _CGotoState123Action,
	{_CState153, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState153, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState153, CAndExpressionType}:            _CGotoState102Action,
	{_CState153, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState153, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState153, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState153, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState153, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState153, CAssignmentExpressionType}:     _CGotoState154Action,
	{_CState153, CInitializerType}:              _CGotoState246Action,
	{_CState153, CInitializerListType}:          _CGotoState247Action,
	{_CState159, CColonToken}:                   _CGotoState248Action,
	{_CState161, CIdentifierToken}:              _CGotoState131Action,
	{_CState161, CConstantToken}:                _CGotoState78Action,
	{_CState161, CStringLiteralToken}:           _CGotoState97Action,
	{_CState161, CSizeofToken}:                  _CGotoState96Action,
	{_CState161, CIncOpToken}:                   _CGotoState88Action,
	{_CState161, CDecOpToken}:                   _CGotoState80Action,
	{_CState161, CLparamToken}:                  _CGotoState89Action,
	{_CState161, CMulToken}:                     _CGotoState91Action,
	{_CState161, CMinusToken}:                   _CGotoState90Action,
	{_CState161, CPlusToken}:                    _CGotoState92Action,
	{_CState161, CAndToken}:                     _CGotoState75Action,
	{_CState161, CExclaimToken}:                 _CGotoState83Action,
	{_CState161, CTildaToken}:                   _CGotoState99Action,
	{_CState161, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState161, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState161, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState161, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState161, CCastExpressionType}:           _CGotoState104Action,
	{_CState161, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState161, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState161, CShiftExpressionType}:          _CGotoState123Action,
	{_CState161, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState161, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState161, CAndExpressionType}:            _CGotoState102Action,
	{_CState161, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState161, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState161, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState161, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState161, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState161, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState161, CExpressionType}:               _CGotoState170Action,
	{_CState163, CIdentifierToken}:              _CGotoState86Action,
	{_CState163, CConstantToken}:                _CGotoState78Action,
	{_CState163, CStringLiteralToken}:           _CGotoState97Action,
	{_CState163, CSizeofToken}:                  _CGotoState96Action,
	{_CState163, CIncOpToken}:                   _CGotoState88Action,
	{_CState163, CDecOpToken}:                   _CGotoState80Action,
	{_CState163, CCaseToken}:                    _CGotoState77Action,
	{_CState163, CDefaultToken}:                 _CGotoState81Action,
	{_CState163, CIfToken}:                      _CGotoState87Action,
	{_CState163, CSwitchToken}:                  _CGotoState98Action,
	{_CState163, CWhileToken}:                   _CGotoState100Action,
	{_CState163, CDoToken}:                      _CGotoState82Action,
	{_CState163, CForToken}:                     _CGotoState84Action,
	{_CState163, CGotoToken}:                    _CGotoState85Action,
	{_CState163, CContinueToken}:                _CGotoState79Action,
	{_CState163, CBreakToken}:                   _CGotoState76Action,
	{_CState163, CReturnToken}:                  _CGotoState94Action,
	{_CState163, CLparamToken}:                  _CGotoState89Action,
	{_CState163, CLcurlToken}:                   _CGotoState49Action,
	{_CState163, CSemicolonToken}:               _CGotoState95Action,
	{_CState163, CMulToken}:                     _CGotoState91Action,
	{_CState163, CMinusToken}:                   _CGotoState90Action,
	{_CState163, CPlusToken}:                    _CGotoState92Action,
	{_CState163, CAndToken}:                     _CGotoState75Action,
	{_CState163, CExclaimToken}:                 _CGotoState83Action,
	{_CState163, CTildaToken}:                   _CGotoState99Action,
	{_CState163, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState163, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState163, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState163, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState163, CCastExpressionType}:           _CGotoState104Action,
	{_CState163, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState163, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState163, CShiftExpressionType}:          _CGotoState123Action,
	{_CState163, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState163, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState163, CAndExpressionType}:            _CGotoState102Action,
	{_CState163, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState163, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState163, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState163, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState163, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState163, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState163, CExpressionType}:               _CGotoState110Action,
	{_CState163, CStatementType}:                _CGotoState249Action,
	{_CState163, CLabeledStatementType}:         _CGotoState115Action,
	{_CState163, CCompoundStatementType}:        _CGotoState105Action,
	{_CState163, CExpressionStatementType}:      _CGotoState111Action,
	{_CState163, CSelectionStatementType}:       _CGotoState122Action,
	{_CState163, CIterationStatementType}:       _CGotoState113Action,
	{_CState163, CJumpStatementType}:            _CGotoState114Action,
	{_CState164, CWhileToken}:                   _CGotoState250Action,
	{_CState165, CIdentifierToken}:              _CGotoState131Action,
	{_CState165, CConstantToken}:                _CGotoState78Action,
	{_CState165, CStringLiteralToken}:           _CGotoState97Action,
	{_CState165, CSizeofToken}:                  _CGotoState96Action,
	{_CState165, CIncOpToken}:                   _CGotoState88Action,
	{_CState165, CDecOpToken}:                   _CGotoState80Action,
	{_CState165, CLparamToken}:                  _CGotoState89Action,
	{_CState165, CSemicolonToken}:               _CGotoState95Action,
	{_CState165, CMulToken}:                     _CGotoState91Action,
	{_CState165, CMinusToken}:                   _CGotoState90Action,
	{_CState165, CPlusToken}:                    _CGotoState92Action,
	{_CState165, CAndToken}:                     _CGotoState75Action,
	{_CState165, CExclaimToken}:                 _CGotoState83Action,
	{_CState165, CTildaToken}:                   _CGotoState99Action,
	{_CState165, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState165, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState165, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState165, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState165, CCastExpressionType}:           _CGotoState104Action,
	{_CState165, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState165, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState165, CShiftExpressionType}:          _CGotoState123Action,
	{_CState165, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState165, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState165, CAndExpressionType}:            _CGotoState102Action,
	{_CState165, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState165, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState165, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState165, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState165, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState165, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState165, CExpressionType}:               _CGotoState110Action,
	{_CState165, CExpressionStatementType}:      _CGotoState251Action,
	{_CState166, CSemicolonToken}:               _CGotoState252Action,
	{_CState167, CIdentifierToken}:              _CGotoState86Action,
	{_CState167, CConstantToken}:                _CGotoState78Action,
	{_CState167, CStringLiteralToken}:           _CGotoState97Action,
	{_CState167, CSizeofToken}:                  _CGotoState96Action,
	{_CState167, CIncOpToken}:                   _CGotoState88Action,
	{_CState167, CDecOpToken}:                   _CGotoState80Action,
	{_CState167, CCaseToken}:                    _CGotoState77Action,
	{_CState167, CDefaultToken}:                 _CGotoState81Action,
	{_CState167, CIfToken}:                      _CGotoState87Action,
	{_CState167, CSwitchToken}:                  _CGotoState98Action,
	{_CState167, CWhileToken}:                   _CGotoState100Action,
	{_CState167, CDoToken}:                      _CGotoState82Action,
	{_CState167, CForToken}:                     _CGotoState84Action,
	{_CState167, CGotoToken}:                    _CGotoState85Action,
	{_CState167, CContinueToken}:                _CGotoState79Action,
	{_CState167, CBreakToken}:                   _CGotoState76Action,
	{_CState167, CReturnToken}:                  _CGotoState94Action,
	{_CState167, CLparamToken}:                  _CGotoState89Action,
	{_CState167, CLcurlToken}:                   _CGotoState49Action,
	{_CState167, CSemicolonToken}:               _CGotoState95Action,
	{_CState167, CMulToken}:                     _CGotoState91Action,
	{_CState167, CMinusToken}:                   _CGotoState90Action,
	{_CState167, CPlusToken}:                    _CGotoState92Action,
	{_CState167, CAndToken}:                     _CGotoState75Action,
	{_CState167, CExclaimToken}:                 _CGotoState83Action,
	{_CState167, CTildaToken}:                   _CGotoState99Action,
	{_CState167, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState167, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState167, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState167, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState167, CCastExpressionType}:           _CGotoState104Action,
	{_CState167, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState167, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState167, CShiftExpressionType}:          _CGotoState123Action,
	{_CState167, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState167, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState167, CAndExpressionType}:            _CGotoState102Action,
	{_CState167, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState167, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState167, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState167, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState167, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState167, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState167, CExpressionType}:               _CGotoState110Action,
	{_CState167, CStatementType}:                _CGotoState253Action,
	{_CState167, CLabeledStatementType}:         _CGotoState115Action,
	{_CState167, CCompoundStatementType}:        _CGotoState105Action,
	{_CState167, CExpressionStatementType}:      _CGotoState111Action,
	{_CState167, CSelectionStatementType}:       _CGotoState122Action,
	{_CState167, CIterationStatementType}:       _CGotoState113Action,
	{_CState167, CJumpStatementType}:            _CGotoState114Action,
	{_CState168, CIdentifierToken}:              _CGotoState131Action,
	{_CState168, CConstantToken}:                _CGotoState78Action,
	{_CState168, CStringLiteralToken}:           _CGotoState97Action,
	{_CState168, CSizeofToken}:                  _CGotoState96Action,
	{_CState168, CIncOpToken}:                   _CGotoState88Action,
	{_CState168, CDecOpToken}:                   _CGotoState80Action,
	{_CState168, CLparamToken}:                  _CGotoState89Action,
	{_CState168, CMulToken}:                     _CGotoState91Action,
	{_CState168, CMinusToken}:                   _CGotoState90Action,
	{_CState168, CPlusToken}:                    _CGotoState92Action,
	{_CState168, CAndToken}:                     _CGotoState75Action,
	{_CState168, CExclaimToken}:                 _CGotoState83Action,
	{_CState168, CTildaToken}:                   _CGotoState99Action,
	{_CState168, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState168, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState168, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState168, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState168, CCastExpressionType}:           _CGotoState104Action,
	{_CState168, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState168, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState168, CShiftExpressionType}:          _CGotoState123Action,
	{_CState168, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState168, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState168, CAndExpressionType}:            _CGotoState102Action,
	{_CState168, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState168, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState168, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState168, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState168, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState168, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState168, CExpressionType}:               _CGotoState254Action,
	{_CState170, CRparamToken}:                  _CGotoState255Action,
	{_CState170, CCommaToken}:                   _CGotoState187Action,
	{_CState171, CLparamToken}:                  _CGotoState256Action,
	{_CState171, CLbraceToken}:                  _CGotoState224Action,
	{_CState171, CMulToken}:                     _CGotoState14Action,
	{_CState171, CPointerType}:                  _CGotoState258Action,
	{_CState171, CAbstractDeclaratorType}:       _CGotoState257Action,
	{_CState171, CDirectAbstractDeclaratorType}: _CGotoState228Action,
	{_CState172, CRparamToken}:                  _CGotoState259Action,
	{_CState174, CSemicolonToken}:               _CGotoState260Action,
	{_CState174, CCommaToken}:                   _CGotoState187Action,
	{_CState175, CIdentifierToken}:              _CGotoState131Action,
	{_CState175, CConstantToken}:                _CGotoState78Action,
	{_CState175, CStringLiteralToken}:           _CGotoState97Action,
	{_CState175, CSizeofToken}:                  _CGotoState96Action,
	{_CState175, CIncOpToken}:                   _CGotoState88Action,
	{_CState175, CDecOpToken}:                   _CGotoState80Action,
	{_CState175, CTypeNameToken}:                _CGotoState21Action,
	{_CState175, CCharToken}:                    _CGotoState4Action,
	{_CState175, CShortToken}:                   _CGotoState16Action,
	{_CState175, CIntToken}:                     _CGotoState11Action,
	{_CState175, CLongToken}:                    _CGotoState12Action,
	{_CState175, CSignedToken}:                  _CGotoState17Action,
	{_CState175, CUnsignedToken}:                _CGotoState23Action,
	{_CState175, CFloatToken}:                   _CGotoState9Action,
	{_CState175, CDoubleToken}:                  _CGotoState6Action,
	{_CState175, CConstToken}:                   _CGotoState5Action,
	{_CState175, CVolatileToken}:                _CGotoState25Action,
	{_CState175, CVoidToken}:                    _CGotoState24Action,
	{_CState175, CStructToken}:                  _CGotoState19Action,
	{_CState175, CUnionToken}:                   _CGotoState22Action,
	{_CState175, CEnumToken}:                    _CGotoState7Action,
	{_CState175, CLparamToken}:                  _CGotoState89Action,
	{_CState175, CMulToken}:                     _CGotoState91Action,
	{_CState175, CMinusToken}:                   _CGotoState90Action,
	{_CState175, CPlusToken}:                    _CGotoState92Action,
	{_CState175, CAndToken}:                     _CGotoState75Action,
	{_CState175, CExclaimToken}:                 _CGotoState83Action,
	{_CState175, CTildaToken}:                   _CGotoState99Action,
	{_CState175, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState175, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState175, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState175, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState175, CCastExpressionType}:           _CGotoState104Action,
	{_CState175, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState175, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState175, CShiftExpressionType}:          _CGotoState123Action,
	{_CState175, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState175, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState175, CAndExpressionType}:            _CGotoState102Action,
	{_CState175, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState175, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState175, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState175, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState175, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState175, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState175, CExpressionType}:               _CGotoState170Action,
	{_CState175, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState175, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState175, CStructOrUnionType}:            _CGotoState35Action,
	{_CState175, CSpecifierQualifierListType}:   _CGotoState171Action,
	{_CState175, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState175, CTypeQualifierType}:            _CGotoState147Action,
	{_CState175, CTypeNameType}:                 _CGotoState261Action,
	{_CState177, CIdentifierToken}:              _CGotoState131Action,
	{_CState177, CConstantToken}:                _CGotoState78Action,
	{_CState177, CStringLiteralToken}:           _CGotoState97Action,
	{_CState177, CSizeofToken}:                  _CGotoState96Action,
	{_CState177, CIncOpToken}:                   _CGotoState88Action,
	{_CState177, CDecOpToken}:                   _CGotoState80Action,
	{_CState177, CLparamToken}:                  _CGotoState89Action,
	{_CState177, CMulToken}:                     _CGotoState91Action,
	{_CState177, CMinusToken}:                   _CGotoState90Action,
	{_CState177, CPlusToken}:                    _CGotoState92Action,
	{_CState177, CAndToken}:                     _CGotoState75Action,
	{_CState177, CExclaimToken}:                 _CGotoState83Action,
	{_CState177, CTildaToken}:                   _CGotoState99Action,
	{_CState177, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState177, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState177, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState177, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState177, CCastExpressionType}:           _CGotoState104Action,
	{_CState177, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState177, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState177, CShiftExpressionType}:          _CGotoState123Action,
	{_CState177, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState177, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState177, CAndExpressionType}:            _CGotoState102Action,
	{_CState177, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState177, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState177, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState177, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState177, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState177, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState177, CExpressionType}:               _CGotoState262Action,
	{_CState178, CIdentifierToken}:              _CGotoState131Action,
	{_CState178, CConstantToken}:                _CGotoState78Action,
	{_CState178, CStringLiteralToken}:           _CGotoState97Action,
	{_CState178, CSizeofToken}:                  _CGotoState96Action,
	{_CState178, CIncOpToken}:                   _CGotoState88Action,
	{_CState178, CDecOpToken}:                   _CGotoState80Action,
	{_CState178, CLparamToken}:                  _CGotoState89Action,
	{_CState178, CMulToken}:                     _CGotoState91Action,
	{_CState178, CMinusToken}:                   _CGotoState90Action,
	{_CState178, CPlusToken}:                    _CGotoState92Action,
	{_CState178, CAndToken}:                     _CGotoState75Action,
	{_CState178, CExclaimToken}:                 _CGotoState83Action,
	{_CState178, CTildaToken}:                   _CGotoState99Action,
	{_CState178, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState178, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState178, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState178, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState178, CCastExpressionType}:           _CGotoState104Action,
	{_CState178, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState178, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState178, CShiftExpressionType}:          _CGotoState123Action,
	{_CState178, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState178, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState178, CAndExpressionType}:            _CGotoState102Action,
	{_CState178, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState178, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState178, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState178, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState178, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState178, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState178, CExpressionType}:               _CGotoState263Action,
	{_CState179, CIdentifierToken}:              _CGotoState131Action,
	{_CState179, CConstantToken}:                _CGotoState78Action,
	{_CState179, CStringLiteralToken}:           _CGotoState97Action,
	{_CState179, CSizeofToken}:                  _CGotoState96Action,
	{_CState179, CIncOpToken}:                   _CGotoState88Action,
	{_CState179, CDecOpToken}:                   _CGotoState80Action,
	{_CState179, CLparamToken}:                  _CGotoState89Action,
	{_CState179, CMulToken}:                     _CGotoState91Action,
	{_CState179, CMinusToken}:                   _CGotoState90Action,
	{_CState179, CPlusToken}:                    _CGotoState92Action,
	{_CState179, CAndToken}:                     _CGotoState75Action,
	{_CState179, CExclaimToken}:                 _CGotoState83Action,
	{_CState179, CTildaToken}:                   _CGotoState99Action,
	{_CState179, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState179, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState179, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState179, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState179, CCastExpressionType}:           _CGotoState104Action,
	{_CState179, CMultiplicativeExpressionType}: _CGotoState264Action,
	{_CState180, CIdentifierToken}:              _CGotoState131Action,
	{_CState180, CConstantToken}:                _CGotoState78Action,
	{_CState180, CStringLiteralToken}:           _CGotoState97Action,
	{_CState180, CSizeofToken}:                  _CGotoState96Action,
	{_CState180, CIncOpToken}:                   _CGotoState88Action,
	{_CState180, CDecOpToken}:                   _CGotoState80Action,
	{_CState180, CLparamToken}:                  _CGotoState89Action,
	{_CState180, CMulToken}:                     _CGotoState91Action,
	{_CState180, CMinusToken}:                   _CGotoState90Action,
	{_CState180, CPlusToken}:                    _CGotoState92Action,
	{_CState180, CAndToken}:                     _CGotoState75Action,
	{_CState180, CExclaimToken}:                 _CGotoState83Action,
	{_CState180, CTildaToken}:                   _CGotoState99Action,
	{_CState180, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState180, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState180, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState180, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState180, CCastExpressionType}:           _CGotoState104Action,
	{_CState180, CMultiplicativeExpressionType}: _CGotoState265Action,
	{_CState181, CIdentifierToken}:              _CGotoState131Action,
	{_CState181, CConstantToken}:                _CGotoState78Action,
	{_CState181, CStringLiteralToken}:           _CGotoState97Action,
	{_CState181, CSizeofToken}:                  _CGotoState96Action,
	{_CState181, CIncOpToken}:                   _CGotoState88Action,
	{_CState181, CDecOpToken}:                   _CGotoState80Action,
	{_CState181, CLparamToken}:                  _CGotoState89Action,
	{_CState181, CMulToken}:                     _CGotoState91Action,
	{_CState181, CMinusToken}:                   _CGotoState90Action,
	{_CState181, CPlusToken}:                    _CGotoState92Action,
	{_CState181, CAndToken}:                     _CGotoState75Action,
	{_CState181, CExclaimToken}:                 _CGotoState83Action,
	{_CState181, CTildaToken}:                   _CGotoState99Action,
	{_CState181, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState181, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState181, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState181, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState181, CCastExpressionType}:           _CGotoState104Action,
	{_CState181, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState181, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState181, CShiftExpressionType}:          _CGotoState123Action,
	{_CState181, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState181, CEqualityExpressionType}:       _CGotoState266Action,
	{_CState183, CIdentifierToken}:              _CGotoState86Action,
	{_CState183, CConstantToken}:                _CGotoState78Action,
	{_CState183, CStringLiteralToken}:           _CGotoState97Action,
	{_CState183, CSizeofToken}:                  _CGotoState96Action,
	{_CState183, CIncOpToken}:                   _CGotoState88Action,
	{_CState183, CDecOpToken}:                   _CGotoState80Action,
	{_CState183, CCaseToken}:                    _CGotoState77Action,
	{_CState183, CDefaultToken}:                 _CGotoState81Action,
	{_CState183, CIfToken}:                      _CGotoState87Action,
	{_CState183, CSwitchToken}:                  _CGotoState98Action,
	{_CState183, CWhileToken}:                   _CGotoState100Action,
	{_CState183, CDoToken}:                      _CGotoState82Action,
	{_CState183, CForToken}:                     _CGotoState84Action,
	{_CState183, CGotoToken}:                    _CGotoState85Action,
	{_CState183, CContinueToken}:                _CGotoState79Action,
	{_CState183, CBreakToken}:                   _CGotoState76Action,
	{_CState183, CReturnToken}:                  _CGotoState94Action,
	{_CState183, CLparamToken}:                  _CGotoState89Action,
	{_CState183, CLcurlToken}:                   _CGotoState49Action,
	{_CState183, CRcurlToken}:                   _CGotoState267Action,
	{_CState183, CSemicolonToken}:               _CGotoState95Action,
	{_CState183, CMulToken}:                     _CGotoState91Action,
	{_CState183, CMinusToken}:                   _CGotoState90Action,
	{_CState183, CPlusToken}:                    _CGotoState92Action,
	{_CState183, CAndToken}:                     _CGotoState75Action,
	{_CState183, CExclaimToken}:                 _CGotoState83Action,
	{_CState183, CTildaToken}:                   _CGotoState99Action,
	{_CState183, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState183, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState183, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState183, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState183, CCastExpressionType}:           _CGotoState104Action,
	{_CState183, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState183, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState183, CShiftExpressionType}:          _CGotoState123Action,
	{_CState183, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState183, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState183, CAndExpressionType}:            _CGotoState102Action,
	{_CState183, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState183, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState183, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState183, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState183, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState183, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState183, CExpressionType}:               _CGotoState110Action,
	{_CState183, CStatementType}:                _CGotoState209Action,
	{_CState183, CLabeledStatementType}:         _CGotoState115Action,
	{_CState183, CCompoundStatementType}:        _CGotoState105Action,
	{_CState183, CExpressionStatementType}:      _CGotoState111Action,
	{_CState183, CSelectionStatementType}:       _CGotoState122Action,
	{_CState183, CIterationStatementType}:       _CGotoState113Action,
	{_CState183, CJumpStatementType}:            _CGotoState114Action,
	{_CState184, CIdentifierToken}:              _CGotoState131Action,
	{_CState184, CConstantToken}:                _CGotoState78Action,
	{_CState184, CStringLiteralToken}:           _CGotoState97Action,
	{_CState184, CSizeofToken}:                  _CGotoState96Action,
	{_CState184, CIncOpToken}:                   _CGotoState88Action,
	{_CState184, CDecOpToken}:                   _CGotoState80Action,
	{_CState184, CLparamToken}:                  _CGotoState89Action,
	{_CState184, CMulToken}:                     _CGotoState91Action,
	{_CState184, CMinusToken}:                   _CGotoState90Action,
	{_CState184, CPlusToken}:                    _CGotoState92Action,
	{_CState184, CAndToken}:                     _CGotoState75Action,
	{_CState184, CExclaimToken}:                 _CGotoState83Action,
	{_CState184, CTildaToken}:                   _CGotoState99Action,
	{_CState184, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState184, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState184, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState184, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState184, CCastExpressionType}:           _CGotoState104Action,
	{_CState184, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState184, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState184, CShiftExpressionType}:          _CGotoState123Action,
	{_CState184, CRelationalExpressionType}:     _CGotoState268Action,
	{_CState185, CIdentifierToken}:              _CGotoState131Action,
	{_CState185, CConstantToken}:                _CGotoState78Action,
	{_CState185, CStringLiteralToken}:           _CGotoState97Action,
	{_CState185, CSizeofToken}:                  _CGotoState96Action,
	{_CState185, CIncOpToken}:                   _CGotoState88Action,
	{_CState185, CDecOpToken}:                   _CGotoState80Action,
	{_CState185, CLparamToken}:                  _CGotoState89Action,
	{_CState185, CMulToken}:                     _CGotoState91Action,
	{_CState185, CMinusToken}:                   _CGotoState90Action,
	{_CState185, CPlusToken}:                    _CGotoState92Action,
	{_CState185, CAndToken}:                     _CGotoState75Action,
	{_CState185, CExclaimToken}:                 _CGotoState83Action,
	{_CState185, CTildaToken}:                   _CGotoState99Action,
	{_CState185, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState185, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState185, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState185, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState185, CCastExpressionType}:           _CGotoState104Action,
	{_CState185, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState185, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState185, CShiftExpressionType}:          _CGotoState123Action,
	{_CState185, CRelationalExpressionType}:     _CGotoState269Action,
	{_CState186, CIdentifierToken}:              _CGotoState131Action,
	{_CState186, CConstantToken}:                _CGotoState78Action,
	{_CState186, CStringLiteralToken}:           _CGotoState97Action,
	{_CState186, CSizeofToken}:                  _CGotoState96Action,
	{_CState186, CIncOpToken}:                   _CGotoState88Action,
	{_CState186, CDecOpToken}:                   _CGotoState80Action,
	{_CState186, CLparamToken}:                  _CGotoState89Action,
	{_CState186, CMulToken}:                     _CGotoState91Action,
	{_CState186, CMinusToken}:                   _CGotoState90Action,
	{_CState186, CPlusToken}:                    _CGotoState92Action,
	{_CState186, CAndToken}:                     _CGotoState75Action,
	{_CState186, CExclaimToken}:                 _CGotoState83Action,
	{_CState186, CTildaToken}:                   _CGotoState99Action,
	{_CState186, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState186, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState186, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState186, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState186, CCastExpressionType}:           _CGotoState104Action,
	{_CState186, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState186, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState186, CShiftExpressionType}:          _CGotoState123Action,
	{_CState186, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState186, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState186, CAndExpressionType}:            _CGotoState270Action,
	{_CState187, CIdentifierToken}:              _CGotoState131Action,
	{_CState187, CConstantToken}:                _CGotoState78Action,
	{_CState187, CStringLiteralToken}:           _CGotoState97Action,
	{_CState187, CSizeofToken}:                  _CGotoState96Action,
	{_CState187, CIncOpToken}:                   _CGotoState88Action,
	{_CState187, CDecOpToken}:                   _CGotoState80Action,
	{_CState187, CLparamToken}:                  _CGotoState89Action,
	{_CState187, CMulToken}:                     _CGotoState91Action,
	{_CState187, CMinusToken}:                   _CGotoState90Action,
	{_CState187, CPlusToken}:                    _CGotoState92Action,
	{_CState187, CAndToken}:                     _CGotoState75Action,
	{_CState187, CExclaimToken}:                 _CGotoState83Action,
	{_CState187, CTildaToken}:                   _CGotoState99Action,
	{_CState187, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState187, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState187, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState187, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState187, CCastExpressionType}:           _CGotoState104Action,
	{_CState187, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState187, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState187, CShiftExpressionType}:          _CGotoState123Action,
	{_CState187, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState187, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState187, CAndExpressionType}:            _CGotoState102Action,
	{_CState187, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState187, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState187, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState187, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState187, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState187, CAssignmentExpressionType}:     _CGotoState271Action,
	{_CState189, CIdentifierToken}:              _CGotoState131Action,
	{_CState189, CConstantToken}:                _CGotoState78Action,
	{_CState189, CStringLiteralToken}:           _CGotoState97Action,
	{_CState189, CSizeofToken}:                  _CGotoState96Action,
	{_CState189, CIncOpToken}:                   _CGotoState88Action,
	{_CState189, CDecOpToken}:                   _CGotoState80Action,
	{_CState189, CLparamToken}:                  _CGotoState89Action,
	{_CState189, CMulToken}:                     _CGotoState91Action,
	{_CState189, CMinusToken}:                   _CGotoState90Action,
	{_CState189, CPlusToken}:                    _CGotoState92Action,
	{_CState189, CAndToken}:                     _CGotoState75Action,
	{_CState189, CExclaimToken}:                 _CGotoState83Action,
	{_CState189, CTildaToken}:                   _CGotoState99Action,
	{_CState189, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState189, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState189, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState189, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState189, CCastExpressionType}:           _CGotoState104Action,
	{_CState189, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState189, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState189, CShiftExpressionType}:          _CGotoState123Action,
	{_CState189, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState189, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState189, CAndExpressionType}:            _CGotoState102Action,
	{_CState189, CExclusiveOrExpressionType}:    _CGotoState272Action,
	{_CState190, CIdentifierToken}:              _CGotoState131Action,
	{_CState190, CConstantToken}:                _CGotoState78Action,
	{_CState190, CStringLiteralToken}:           _CGotoState97Action,
	{_CState190, CSizeofToken}:                  _CGotoState96Action,
	{_CState190, CIncOpToken}:                   _CGotoState88Action,
	{_CState190, CDecOpToken}:                   _CGotoState80Action,
	{_CState190, CLparamToken}:                  _CGotoState89Action,
	{_CState190, CMulToken}:                     _CGotoState91Action,
	{_CState190, CMinusToken}:                   _CGotoState90Action,
	{_CState190, CPlusToken}:                    _CGotoState92Action,
	{_CState190, CAndToken}:                     _CGotoState75Action,
	{_CState190, CExclaimToken}:                 _CGotoState83Action,
	{_CState190, CTildaToken}:                   _CGotoState99Action,
	{_CState190, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState190, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState190, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState190, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState190, CCastExpressionType}:           _CGotoState104Action,
	{_CState190, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState190, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState190, CShiftExpressionType}:          _CGotoState123Action,
	{_CState190, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState190, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState190, CAndExpressionType}:            _CGotoState102Action,
	{_CState190, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState190, CInclusiveOrExpressionType}:    _CGotoState273Action,
	{_CState191, CIdentifierToken}:              _CGotoState131Action,
	{_CState191, CConstantToken}:                _CGotoState78Action,
	{_CState191, CStringLiteralToken}:           _CGotoState97Action,
	{_CState191, CSizeofToken}:                  _CGotoState96Action,
	{_CState191, CIncOpToken}:                   _CGotoState88Action,
	{_CState191, CDecOpToken}:                   _CGotoState80Action,
	{_CState191, CLparamToken}:                  _CGotoState89Action,
	{_CState191, CMulToken}:                     _CGotoState91Action,
	{_CState191, CMinusToken}:                   _CGotoState90Action,
	{_CState191, CPlusToken}:                    _CGotoState92Action,
	{_CState191, CAndToken}:                     _CGotoState75Action,
	{_CState191, CExclaimToken}:                 _CGotoState83Action,
	{_CState191, CTildaToken}:                   _CGotoState99Action,
	{_CState191, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState191, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState191, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState191, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState191, CCastExpressionType}:           _CGotoState104Action,
	{_CState191, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState191, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState191, CShiftExpressionType}:          _CGotoState123Action,
	{_CState191, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState191, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState191, CAndExpressionType}:            _CGotoState102Action,
	{_CState191, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState191, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState191, CLogicalAndExpressionType}:     _CGotoState274Action,
	{_CState192, CIdentifierToken}:              _CGotoState131Action,
	{_CState192, CConstantToken}:                _CGotoState78Action,
	{_CState192, CStringLiteralToken}:           _CGotoState97Action,
	{_CState192, CSizeofToken}:                  _CGotoState96Action,
	{_CState192, CIncOpToken}:                   _CGotoState88Action,
	{_CState192, CDecOpToken}:                   _CGotoState80Action,
	{_CState192, CLparamToken}:                  _CGotoState89Action,
	{_CState192, CMulToken}:                     _CGotoState91Action,
	{_CState192, CMinusToken}:                   _CGotoState90Action,
	{_CState192, CPlusToken}:                    _CGotoState92Action,
	{_CState192, CAndToken}:                     _CGotoState75Action,
	{_CState192, CExclaimToken}:                 _CGotoState83Action,
	{_CState192, CTildaToken}:                   _CGotoState99Action,
	{_CState192, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState192, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState192, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState192, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState192, CCastExpressionType}:           _CGotoState104Action,
	{_CState192, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState192, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState192, CShiftExpressionType}:          _CGotoState123Action,
	{_CState192, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState192, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState192, CAndExpressionType}:            _CGotoState102Action,
	{_CState192, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState192, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState192, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState192, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState192, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState192, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState192, CExpressionType}:               _CGotoState275Action,
	{_CState193, CIdentifierToken}:              _CGotoState131Action,
	{_CState193, CConstantToken}:                _CGotoState78Action,
	{_CState193, CStringLiteralToken}:           _CGotoState97Action,
	{_CState193, CSizeofToken}:                  _CGotoState96Action,
	{_CState193, CIncOpToken}:                   _CGotoState88Action,
	{_CState193, CDecOpToken}:                   _CGotoState80Action,
	{_CState193, CLparamToken}:                  _CGotoState89Action,
	{_CState193, CMulToken}:                     _CGotoState91Action,
	{_CState193, CMinusToken}:                   _CGotoState90Action,
	{_CState193, CPlusToken}:                    _CGotoState92Action,
	{_CState193, CAndToken}:                     _CGotoState75Action,
	{_CState193, CExclaimToken}:                 _CGotoState83Action,
	{_CState193, CTildaToken}:                   _CGotoState99Action,
	{_CState193, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState193, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState193, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState193, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState193, CCastExpressionType}:           _CGotoState276Action,
	{_CState194, CIdentifierToken}:              _CGotoState131Action,
	{_CState194, CConstantToken}:                _CGotoState78Action,
	{_CState194, CStringLiteralToken}:           _CGotoState97Action,
	{_CState194, CSizeofToken}:                  _CGotoState96Action,
	{_CState194, CIncOpToken}:                   _CGotoState88Action,
	{_CState194, CDecOpToken}:                   _CGotoState80Action,
	{_CState194, CLparamToken}:                  _CGotoState89Action,
	{_CState194, CMulToken}:                     _CGotoState91Action,
	{_CState194, CMinusToken}:                   _CGotoState90Action,
	{_CState194, CPlusToken}:                    _CGotoState92Action,
	{_CState194, CAndToken}:                     _CGotoState75Action,
	{_CState194, CExclaimToken}:                 _CGotoState83Action,
	{_CState194, CTildaToken}:                   _CGotoState99Action,
	{_CState194, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState194, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState194, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState194, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState194, CCastExpressionType}:           _CGotoState277Action,
	{_CState195, CIdentifierToken}:              _CGotoState131Action,
	{_CState195, CConstantToken}:                _CGotoState78Action,
	{_CState195, CStringLiteralToken}:           _CGotoState97Action,
	{_CState195, CSizeofToken}:                  _CGotoState96Action,
	{_CState195, CIncOpToken}:                   _CGotoState88Action,
	{_CState195, CDecOpToken}:                   _CGotoState80Action,
	{_CState195, CLparamToken}:                  _CGotoState89Action,
	{_CState195, CMulToken}:                     _CGotoState91Action,
	{_CState195, CMinusToken}:                   _CGotoState90Action,
	{_CState195, CPlusToken}:                    _CGotoState92Action,
	{_CState195, CAndToken}:                     _CGotoState75Action,
	{_CState195, CExclaimToken}:                 _CGotoState83Action,
	{_CState195, CTildaToken}:                   _CGotoState99Action,
	{_CState195, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState195, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState195, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState195, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState195, CCastExpressionType}:           _CGotoState278Action,
	{_CState197, CIdentifierToken}:              _CGotoState279Action,
	{_CState199, CIdentifierToken}:              _CGotoState131Action,
	{_CState199, CConstantToken}:                _CGotoState78Action,
	{_CState199, CStringLiteralToken}:           _CGotoState97Action,
	{_CState199, CSizeofToken}:                  _CGotoState96Action,
	{_CState199, CIncOpToken}:                   _CGotoState88Action,
	{_CState199, CDecOpToken}:                   _CGotoState80Action,
	{_CState199, CLparamToken}:                  _CGotoState89Action,
	{_CState199, CMulToken}:                     _CGotoState91Action,
	{_CState199, CMinusToken}:                   _CGotoState90Action,
	{_CState199, CPlusToken}:                    _CGotoState92Action,
	{_CState199, CAndToken}:                     _CGotoState75Action,
	{_CState199, CExclaimToken}:                 _CGotoState83Action,
	{_CState199, CTildaToken}:                   _CGotoState99Action,
	{_CState199, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState199, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState199, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState199, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState199, CCastExpressionType}:           _CGotoState104Action,
	{_CState199, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState199, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState199, CShiftExpressionType}:          _CGotoState123Action,
	{_CState199, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState199, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState199, CAndExpressionType}:            _CGotoState102Action,
	{_CState199, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState199, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState199, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState199, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState199, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState199, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState199, CExpressionType}:               _CGotoState280Action,
	{_CState200, CIdentifierToken}:              _CGotoState131Action,
	{_CState200, CConstantToken}:                _CGotoState78Action,
	{_CState200, CStringLiteralToken}:           _CGotoState97Action,
	{_CState200, CSizeofToken}:                  _CGotoState96Action,
	{_CState200, CIncOpToken}:                   _CGotoState88Action,
	{_CState200, CDecOpToken}:                   _CGotoState80Action,
	{_CState200, CLparamToken}:                  _CGotoState89Action,
	{_CState200, CRparamToken}:                  _CGotoState281Action,
	{_CState200, CMulToken}:                     _CGotoState91Action,
	{_CState200, CMinusToken}:                   _CGotoState90Action,
	{_CState200, CPlusToken}:                    _CGotoState92Action,
	{_CState200, CAndToken}:                     _CGotoState75Action,
	{_CState200, CExclaimToken}:                 _CGotoState83Action,
	{_CState200, CTildaToken}:                   _CGotoState99Action,
	{_CState200, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState200, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState200, CArgumentExpressionListType}:   _CGotoState282Action,
	{_CState200, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState200, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState200, CCastExpressionType}:           _CGotoState104Action,
	{_CState200, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState200, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState200, CShiftExpressionType}:          _CGotoState123Action,
	{_CState200, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState200, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState200, CAndExpressionType}:            _CGotoState102Action,
	{_CState200, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState200, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState200, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState200, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState200, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState200, CAssignmentExpressionType}:     _CGotoState283Action,
	{_CState201, CIdentifierToken}:              _CGotoState284Action,
	{_CState202, CIdentifierToken}:              _CGotoState131Action,
	{_CState202, CConstantToken}:                _CGotoState78Action,
	{_CState202, CStringLiteralToken}:           _CGotoState97Action,
	{_CState202, CSizeofToken}:                  _CGotoState96Action,
	{_CState202, CIncOpToken}:                   _CGotoState88Action,
	{_CState202, CDecOpToken}:                   _CGotoState80Action,
	{_CState202, CLparamToken}:                  _CGotoState89Action,
	{_CState202, CMulToken}:                     _CGotoState91Action,
	{_CState202, CMinusToken}:                   _CGotoState90Action,
	{_CState202, CPlusToken}:                    _CGotoState92Action,
	{_CState202, CAndToken}:                     _CGotoState75Action,
	{_CState202, CExclaimToken}:                 _CGotoState83Action,
	{_CState202, CTildaToken}:                   _CGotoState99Action,
	{_CState202, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState202, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState202, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState202, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState202, CCastExpressionType}:           _CGotoState104Action,
	{_CState202, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState202, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState202, CShiftExpressionType}:          _CGotoState285Action,
	{_CState203, CIdentifierToken}:              _CGotoState131Action,
	{_CState203, CConstantToken}:                _CGotoState78Action,
	{_CState203, CStringLiteralToken}:           _CGotoState97Action,
	{_CState203, CSizeofToken}:                  _CGotoState96Action,
	{_CState203, CIncOpToken}:                   _CGotoState88Action,
	{_CState203, CDecOpToken}:                   _CGotoState80Action,
	{_CState203, CLparamToken}:                  _CGotoState89Action,
	{_CState203, CMulToken}:                     _CGotoState91Action,
	{_CState203, CMinusToken}:                   _CGotoState90Action,
	{_CState203, CPlusToken}:                    _CGotoState92Action,
	{_CState203, CAndToken}:                     _CGotoState75Action,
	{_CState203, CExclaimToken}:                 _CGotoState83Action,
	{_CState203, CTildaToken}:                   _CGotoState99Action,
	{_CState203, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState203, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState203, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState203, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState203, CCastExpressionType}:           _CGotoState104Action,
	{_CState203, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState203, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState203, CShiftExpressionType}:          _CGotoState286Action,
	{_CState204, CIdentifierToken}:              _CGotoState131Action,
	{_CState204, CConstantToken}:                _CGotoState78Action,
	{_CState204, CStringLiteralToken}:           _CGotoState97Action,
	{_CState204, CSizeofToken}:                  _CGotoState96Action,
	{_CState204, CIncOpToken}:                   _CGotoState88Action,
	{_CState204, CDecOpToken}:                   _CGotoState80Action,
	{_CState204, CLparamToken}:                  _CGotoState89Action,
	{_CState204, CMulToken}:                     _CGotoState91Action,
	{_CState204, CMinusToken}:                   _CGotoState90Action,
	{_CState204, CPlusToken}:                    _CGotoState92Action,
	{_CState204, CAndToken}:                     _CGotoState75Action,
	{_CState204, CExclaimToken}:                 _CGotoState83Action,
	{_CState204, CTildaToken}:                   _CGotoState99Action,
	{_CState204, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState204, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState204, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState204, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState204, CCastExpressionType}:           _CGotoState104Action,
	{_CState204, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState204, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState204, CShiftExpressionType}:          _CGotoState287Action,
	{_CState205, CIdentifierToken}:              _CGotoState131Action,
	{_CState205, CConstantToken}:                _CGotoState78Action,
	{_CState205, CStringLiteralToken}:           _CGotoState97Action,
	{_CState205, CSizeofToken}:                  _CGotoState96Action,
	{_CState205, CIncOpToken}:                   _CGotoState88Action,
	{_CState205, CDecOpToken}:                   _CGotoState80Action,
	{_CState205, CLparamToken}:                  _CGotoState89Action,
	{_CState205, CMulToken}:                     _CGotoState91Action,
	{_CState205, CMinusToken}:                   _CGotoState90Action,
	{_CState205, CPlusToken}:                    _CGotoState92Action,
	{_CState205, CAndToken}:                     _CGotoState75Action,
	{_CState205, CExclaimToken}:                 _CGotoState83Action,
	{_CState205, CTildaToken}:                   _CGotoState99Action,
	{_CState205, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState205, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState205, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState205, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState205, CCastExpressionType}:           _CGotoState104Action,
	{_CState205, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState205, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState205, CShiftExpressionType}:          _CGotoState288Action,
	{_CState206, CIdentifierToken}:              _CGotoState131Action,
	{_CState206, CConstantToken}:                _CGotoState78Action,
	{_CState206, CStringLiteralToken}:           _CGotoState97Action,
	{_CState206, CSizeofToken}:                  _CGotoState96Action,
	{_CState206, CIncOpToken}:                   _CGotoState88Action,
	{_CState206, CDecOpToken}:                   _CGotoState80Action,
	{_CState206, CLparamToken}:                  _CGotoState89Action,
	{_CState206, CMulToken}:                     _CGotoState91Action,
	{_CState206, CMinusToken}:                   _CGotoState90Action,
	{_CState206, CPlusToken}:                    _CGotoState92Action,
	{_CState206, CAndToken}:                     _CGotoState75Action,
	{_CState206, CExclaimToken}:                 _CGotoState83Action,
	{_CState206, CTildaToken}:                   _CGotoState99Action,
	{_CState206, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState206, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState206, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState206, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState206, CCastExpressionType}:           _CGotoState104Action,
	{_CState206, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState206, CAdditiveExpressionType}:       _CGotoState289Action,
	{_CState207, CIdentifierToken}:              _CGotoState131Action,
	{_CState207, CConstantToken}:                _CGotoState78Action,
	{_CState207, CStringLiteralToken}:           _CGotoState97Action,
	{_CState207, CSizeofToken}:                  _CGotoState96Action,
	{_CState207, CIncOpToken}:                   _CGotoState88Action,
	{_CState207, CDecOpToken}:                   _CGotoState80Action,
	{_CState207, CLparamToken}:                  _CGotoState89Action,
	{_CState207, CMulToken}:                     _CGotoState91Action,
	{_CState207, CMinusToken}:                   _CGotoState90Action,
	{_CState207, CPlusToken}:                    _CGotoState92Action,
	{_CState207, CAndToken}:                     _CGotoState75Action,
	{_CState207, CExclaimToken}:                 _CGotoState83Action,
	{_CState207, CTildaToken}:                   _CGotoState99Action,
	{_CState207, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState207, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState207, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState207, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState207, CCastExpressionType}:           _CGotoState104Action,
	{_CState207, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState207, CAdditiveExpressionType}:       _CGotoState290Action,
	{_CState221, CIdentifierToken}:              _CGotoState131Action,
	{_CState221, CConstantToken}:                _CGotoState78Action,
	{_CState221, CStringLiteralToken}:           _CGotoState97Action,
	{_CState221, CSizeofToken}:                  _CGotoState96Action,
	{_CState221, CIncOpToken}:                   _CGotoState88Action,
	{_CState221, CDecOpToken}:                   _CGotoState80Action,
	{_CState221, CLparamToken}:                  _CGotoState89Action,
	{_CState221, CMulToken}:                     _CGotoState91Action,
	{_CState221, CMinusToken}:                   _CGotoState90Action,
	{_CState221, CPlusToken}:                    _CGotoState92Action,
	{_CState221, CAndToken}:                     _CGotoState75Action,
	{_CState221, CExclaimToken}:                 _CGotoState83Action,
	{_CState221, CTildaToken}:                   _CGotoState99Action,
	{_CState221, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState221, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState221, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState221, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState221, CCastExpressionType}:           _CGotoState104Action,
	{_CState221, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState221, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState221, CShiftExpressionType}:          _CGotoState123Action,
	{_CState221, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState221, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState221, CAndExpressionType}:            _CGotoState102Action,
	{_CState221, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState221, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState221, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState221, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState221, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState221, CAssignmentExpressionType}:     _CGotoState291Action,
	{_CState224, CIdentifierToken}:              _CGotoState131Action,
	{_CState224, CConstantToken}:                _CGotoState78Action,
	{_CState224, CStringLiteralToken}:           _CGotoState97Action,
	{_CState224, CSizeofToken}:                  _CGotoState96Action,
	{_CState224, CIncOpToken}:                   _CGotoState88Action,
	{_CState224, CDecOpToken}:                   _CGotoState80Action,
	{_CState224, CLparamToken}:                  _CGotoState89Action,
	{_CState224, CRbraceToken}:                  _CGotoState292Action,
	{_CState224, CMulToken}:                     _CGotoState91Action,
	{_CState224, CMinusToken}:                   _CGotoState90Action,
	{_CState224, CPlusToken}:                    _CGotoState92Action,
	{_CState224, CAndToken}:                     _CGotoState75Action,
	{_CState224, CExclaimToken}:                 _CGotoState83Action,
	{_CState224, CTildaToken}:                   _CGotoState99Action,
	{_CState224, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState224, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState224, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState224, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState224, CCastExpressionType}:           _CGotoState104Action,
	{_CState224, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState224, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState224, CShiftExpressionType}:          _CGotoState123Action,
	{_CState224, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState224, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState224, CAndExpressionType}:            _CGotoState102Action,
	{_CState224, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState224, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState224, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState224, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState224, CConditionalExpressionType}:    _CGotoState133Action,
	{_CState224, CConstantExpressionType}:       _CGotoState293Action,
	{_CState225, CIdentifierToken}:              _CGotoState10Action,
	{_CState225, CTypeNameToken}:                _CGotoState21Action,
	{_CState225, CTypedefToken}:                 _CGotoState20Action,
	{_CState225, CExternToken}:                  _CGotoState8Action,
	{_CState225, CStaticToken}:                  _CGotoState18Action,
	{_CState225, CAutoToken}:                    _CGotoState3Action,
	{_CState225, CRegisterToken}:                _CGotoState15Action,
	{_CState225, CCharToken}:                    _CGotoState4Action,
	{_CState225, CShortToken}:                   _CGotoState16Action,
	{_CState225, CIntToken}:                     _CGotoState11Action,
	{_CState225, CLongToken}:                    _CGotoState12Action,
	{_CState225, CSignedToken}:                  _CGotoState17Action,
	{_CState225, CUnsignedToken}:                _CGotoState23Action,
	{_CState225, CFloatToken}:                   _CGotoState9Action,
	{_CState225, CDoubleToken}:                  _CGotoState6Action,
	{_CState225, CConstToken}:                   _CGotoState5Action,
	{_CState225, CVolatileToken}:                _CGotoState25Action,
	{_CState225, CVoidToken}:                    _CGotoState24Action,
	{_CState225, CStructToken}:                  _CGotoState19Action,
	{_CState225, CUnionToken}:                   _CGotoState22Action,
	{_CState225, CEnumToken}:                    _CGotoState7Action,
	{_CState225, CLparamToken}:                  _CGotoState225Action,
	{_CState225, CRparamToken}:                  _CGotoState294Action,
	{_CState225, CLbraceToken}:                  _CGotoState224Action,
	{_CState225, CMulToken}:                     _CGotoState14Action,
	{_CState225, CDeclarationSpecifiersType}:    _CGotoState138Action,
	{_CState225, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState225, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState225, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState225, CStructOrUnionType}:            _CGotoState35Action,
	{_CState225, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState225, CTypeQualifierType}:            _CGotoState37Action,
	{_CState225, CDeclaratorType}:               _CGotoState41Action,
	{_CState225, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState225, CPointerType}:                  _CGotoState229Action,
	{_CState225, CParameterTypeListType}:        _CGotoState296Action,
	{_CState225, CParameterListType}:            _CGotoState141Action,
	{_CState225, CParameterDeclarationType}:     _CGotoState140Action,
	{_CState225, CAbstractDeclaratorType}:       _CGotoState295Action,
	{_CState225, CDirectAbstractDeclaratorType}: _CGotoState228Action,
	{_CState228, CLparamToken}:                  _CGotoState298Action,
	{_CState228, CLbraceToken}:                  _CGotoState297Action,
	{_CState229, CIdentifierToken}:              _CGotoState10Action,
	{_CState229, CLparamToken}:                  _CGotoState225Action,
	{_CState229, CLbraceToken}:                  _CGotoState224Action,
	{_CState229, CDirectDeclaratorType}:         _CGotoState56Action,
	{_CState229, CDirectAbstractDeclaratorType}: _CGotoState299Action,
	{_CState230, CIdentifierToken}:              _CGotoState300Action,
	{_CState232, CTypeNameToken}:                _CGotoState21Action,
	{_CState232, CTypedefToken}:                 _CGotoState20Action,
	{_CState232, CExternToken}:                  _CGotoState8Action,
	{_CState232, CStaticToken}:                  _CGotoState18Action,
	{_CState232, CAutoToken}:                    _CGotoState3Action,
	{_CState232, CRegisterToken}:                _CGotoState15Action,
	{_CState232, CCharToken}:                    _CGotoState4Action,
	{_CState232, CShortToken}:                   _CGotoState16Action,
	{_CState232, CIntToken}:                     _CGotoState11Action,
	{_CState232, CLongToken}:                    _CGotoState12Action,
	{_CState232, CSignedToken}:                  _CGotoState17Action,
	{_CState232, CUnsignedToken}:                _CGotoState23Action,
	{_CState232, CFloatToken}:                   _CGotoState9Action,
	{_CState232, CDoubleToken}:                  _CGotoState6Action,
	{_CState232, CConstToken}:                   _CGotoState5Action,
	{_CState232, CVolatileToken}:                _CGotoState25Action,
	{_CState232, CVoidToken}:                    _CGotoState24Action,
	{_CState232, CStructToken}:                  _CGotoState19Action,
	{_CState232, CUnionToken}:                   _CGotoState22Action,
	{_CState232, CEnumToken}:                    _CGotoState7Action,
	{_CState232, CEllipsisToken}:                _CGotoState301Action,
	{_CState232, CDeclarationSpecifiersType}:    _CGotoState138Action,
	{_CState232, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState232, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState232, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState232, CStructOrUnionType}:            _CGotoState35Action,
	{_CState232, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState232, CTypeQualifierType}:            _CGotoState37Action,
	{_CState232, CParameterDeclarationType}:     _CGotoState302Action,
	{_CState234, CTypeNameToken}:                _CGotoState21Action,
	{_CState234, CCharToken}:                    _CGotoState4Action,
	{_CState234, CShortToken}:                   _CGotoState16Action,
	{_CState234, CIntToken}:                     _CGotoState11Action,
	{_CState234, CLongToken}:                    _CGotoState12Action,
	{_CState234, CSignedToken}:                  _CGotoState17Action,
	{_CState234, CUnsignedToken}:                _CGotoState23Action,
	{_CState234, CFloatToken}:                   _CGotoState9Action,
	{_CState234, CDoubleToken}:                  _CGotoState6Action,
	{_CState234, CConstToken}:                   _CGotoState5Action,
	{_CState234, CVolatileToken}:                _CGotoState25Action,
	{_CState234, CVoidToken}:                    _CGotoState24Action,
	{_CState234, CStructToken}:                  _CGotoState19Action,
	{_CState234, CUnionToken}:                   _CGotoState22Action,
	{_CState234, CEnumToken}:                    _CGotoState7Action,
	{_CState234, CRcurlToken}:                   _CGotoState303Action,
	{_CState234, CTypeSpecifierType}:            _CGotoState148Action,
	{_CState234, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState234, CStructOrUnionType}:            _CGotoState35Action,
	{_CState234, CStructDeclarationType}:        _CGotoState240Action,
	{_CState234, CSpecifierQualifierListType}:   _CGotoState144Action,
	{_CState234, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState234, CTypeQualifierType}:            _CGotoState147Action,
	{_CState235, CIdentifierToken}:              _CGotoState131Action,
	{_CState235, CConstantToken}:                _CGotoState78Action,
	{_CState235, CStringLiteralToken}:           _CGotoState97Action,
	{_CState235, CSizeofToken}:                  _CGotoState96Action,
	{_CState235, CIncOpToken}:                   _CGotoState88Action,
	{_CState235, CDecOpToken}:                   _CGotoState80Action,
	{_CState235, CLparamToken}:                  _CGotoState89Action,
	{_CState235, CMulToken}:                     _CGotoState91Action,
	{_CState235, CMinusToken}:                   _CGotoState90Action,
	{_CState235, CPlusToken}:                    _CGotoState92Action,
	{_CState235, CAndToken}:                     _CGotoState75Action,
	{_CState235, CExclaimToken}:                 _CGotoState83Action,
	{_CState235, CTildaToken}:                   _CGotoState99Action,
	{_CState235, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState235, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState235, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState235, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState235, CCastExpressionType}:           _CGotoState104Action,
	{_CState235, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState235, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState235, CShiftExpressionType}:          _CGotoState123Action,
	{_CState235, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState235, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState235, CAndExpressionType}:            _CGotoState102Action,
	{_CState235, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState235, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState235, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState235, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState235, CConditionalExpressionType}:    _CGotoState133Action,
	{_CState235, CConstantExpressionType}:       _CGotoState304Action,
	{_CState236, CColonToken}:                   _CGotoState305Action,
	{_CState238, CSemicolonToken}:               _CGotoState307Action,
	{_CState238, CCommaToken}:                   _CGotoState306Action,
	{_CState247, CRcurlToken}:                   _CGotoState309Action,
	{_CState247, CCommaToken}:                   _CGotoState308Action,
	{_CState248, CIdentifierToken}:              _CGotoState86Action,
	{_CState248, CConstantToken}:                _CGotoState78Action,
	{_CState248, CStringLiteralToken}:           _CGotoState97Action,
	{_CState248, CSizeofToken}:                  _CGotoState96Action,
	{_CState248, CIncOpToken}:                   _CGotoState88Action,
	{_CState248, CDecOpToken}:                   _CGotoState80Action,
	{_CState248, CCaseToken}:                    _CGotoState77Action,
	{_CState248, CDefaultToken}:                 _CGotoState81Action,
	{_CState248, CIfToken}:                      _CGotoState87Action,
	{_CState248, CSwitchToken}:                  _CGotoState98Action,
	{_CState248, CWhileToken}:                   _CGotoState100Action,
	{_CState248, CDoToken}:                      _CGotoState82Action,
	{_CState248, CForToken}:                     _CGotoState84Action,
	{_CState248, CGotoToken}:                    _CGotoState85Action,
	{_CState248, CContinueToken}:                _CGotoState79Action,
	{_CState248, CBreakToken}:                   _CGotoState76Action,
	{_CState248, CReturnToken}:                  _CGotoState94Action,
	{_CState248, CLparamToken}:                  _CGotoState89Action,
	{_CState248, CLcurlToken}:                   _CGotoState49Action,
	{_CState248, CSemicolonToken}:               _CGotoState95Action,
	{_CState248, CMulToken}:                     _CGotoState91Action,
	{_CState248, CMinusToken}:                   _CGotoState90Action,
	{_CState248, CPlusToken}:                    _CGotoState92Action,
	{_CState248, CAndToken}:                     _CGotoState75Action,
	{_CState248, CExclaimToken}:                 _CGotoState83Action,
	{_CState248, CTildaToken}:                   _CGotoState99Action,
	{_CState248, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState248, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState248, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState248, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState248, CCastExpressionType}:           _CGotoState104Action,
	{_CState248, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState248, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState248, CShiftExpressionType}:          _CGotoState123Action,
	{_CState248, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState248, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState248, CAndExpressionType}:            _CGotoState102Action,
	{_CState248, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState248, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState248, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState248, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState248, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState248, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState248, CExpressionType}:               _CGotoState110Action,
	{_CState248, CStatementType}:                _CGotoState310Action,
	{_CState248, CLabeledStatementType}:         _CGotoState115Action,
	{_CState248, CCompoundStatementType}:        _CGotoState105Action,
	{_CState248, CExpressionStatementType}:      _CGotoState111Action,
	{_CState248, CSelectionStatementType}:       _CGotoState122Action,
	{_CState248, CIterationStatementType}:       _CGotoState113Action,
	{_CState248, CJumpStatementType}:            _CGotoState114Action,
	{_CState250, CLparamToken}:                  _CGotoState311Action,
	{_CState251, CIdentifierToken}:              _CGotoState131Action,
	{_CState251, CConstantToken}:                _CGotoState78Action,
	{_CState251, CStringLiteralToken}:           _CGotoState97Action,
	{_CState251, CSizeofToken}:                  _CGotoState96Action,
	{_CState251, CIncOpToken}:                   _CGotoState88Action,
	{_CState251, CDecOpToken}:                   _CGotoState80Action,
	{_CState251, CLparamToken}:                  _CGotoState89Action,
	{_CState251, CSemicolonToken}:               _CGotoState95Action,
	{_CState251, CMulToken}:                     _CGotoState91Action,
	{_CState251, CMinusToken}:                   _CGotoState90Action,
	{_CState251, CPlusToken}:                    _CGotoState92Action,
	{_CState251, CAndToken}:                     _CGotoState75Action,
	{_CState251, CExclaimToken}:                 _CGotoState83Action,
	{_CState251, CTildaToken}:                   _CGotoState99Action,
	{_CState251, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState251, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState251, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState251, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState251, CCastExpressionType}:           _CGotoState104Action,
	{_CState251, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState251, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState251, CShiftExpressionType}:          _CGotoState123Action,
	{_CState251, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState251, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState251, CAndExpressionType}:            _CGotoState102Action,
	{_CState251, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState251, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState251, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState251, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState251, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState251, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState251, CExpressionType}:               _CGotoState110Action,
	{_CState251, CExpressionStatementType}:      _CGotoState312Action,
	{_CState254, CRparamToken}:                  _CGotoState313Action,
	{_CState254, CCommaToken}:                   _CGotoState187Action,
	{_CState256, CTypeNameToken}:                _CGotoState21Action,
	{_CState256, CTypedefToken}:                 _CGotoState20Action,
	{_CState256, CExternToken}:                  _CGotoState8Action,
	{_CState256, CStaticToken}:                  _CGotoState18Action,
	{_CState256, CAutoToken}:                    _CGotoState3Action,
	{_CState256, CRegisterToken}:                _CGotoState15Action,
	{_CState256, CCharToken}:                    _CGotoState4Action,
	{_CState256, CShortToken}:                   _CGotoState16Action,
	{_CState256, CIntToken}:                     _CGotoState11Action,
	{_CState256, CLongToken}:                    _CGotoState12Action,
	{_CState256, CSignedToken}:                  _CGotoState17Action,
	{_CState256, CUnsignedToken}:                _CGotoState23Action,
	{_CState256, CFloatToken}:                   _CGotoState9Action,
	{_CState256, CDoubleToken}:                  _CGotoState6Action,
	{_CState256, CConstToken}:                   _CGotoState5Action,
	{_CState256, CVolatileToken}:                _CGotoState25Action,
	{_CState256, CVoidToken}:                    _CGotoState24Action,
	{_CState256, CStructToken}:                  _CGotoState19Action,
	{_CState256, CUnionToken}:                   _CGotoState22Action,
	{_CState256, CEnumToken}:                    _CGotoState7Action,
	{_CState256, CLparamToken}:                  _CGotoState256Action,
	{_CState256, CRparamToken}:                  _CGotoState294Action,
	{_CState256, CLbraceToken}:                  _CGotoState224Action,
	{_CState256, CMulToken}:                     _CGotoState14Action,
	{_CState256, CDeclarationSpecifiersType}:    _CGotoState138Action,
	{_CState256, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState256, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState256, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState256, CStructOrUnionType}:            _CGotoState35Action,
	{_CState256, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState256, CTypeQualifierType}:            _CGotoState37Action,
	{_CState256, CPointerType}:                  _CGotoState258Action,
	{_CState256, CParameterTypeListType}:        _CGotoState296Action,
	{_CState256, CParameterListType}:            _CGotoState141Action,
	{_CState256, CParameterDeclarationType}:     _CGotoState140Action,
	{_CState256, CAbstractDeclaratorType}:       _CGotoState295Action,
	{_CState256, CDirectAbstractDeclaratorType}: _CGotoState228Action,
	{_CState258, CLparamToken}:                  _CGotoState256Action,
	{_CState258, CLbraceToken}:                  _CGotoState224Action,
	{_CState258, CDirectAbstractDeclaratorType}: _CGotoState299Action,
	{_CState259, CIdentifierToken}:              _CGotoState131Action,
	{_CState259, CConstantToken}:                _CGotoState78Action,
	{_CState259, CStringLiteralToken}:           _CGotoState97Action,
	{_CState259, CSizeofToken}:                  _CGotoState96Action,
	{_CState259, CIncOpToken}:                   _CGotoState88Action,
	{_CState259, CDecOpToken}:                   _CGotoState80Action,
	{_CState259, CLparamToken}:                  _CGotoState89Action,
	{_CState259, CMulToken}:                     _CGotoState91Action,
	{_CState259, CMinusToken}:                   _CGotoState90Action,
	{_CState259, CPlusToken}:                    _CGotoState92Action,
	{_CState259, CAndToken}:                     _CGotoState75Action,
	{_CState259, CExclaimToken}:                 _CGotoState83Action,
	{_CState259, CTildaToken}:                   _CGotoState99Action,
	{_CState259, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState259, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState259, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState259, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState259, CCastExpressionType}:           _CGotoState314Action,
	{_CState261, CRparamToken}:                  _CGotoState315Action,
	{_CState262, CRparamToken}:                  _CGotoState316Action,
	{_CState262, CCommaToken}:                   _CGotoState187Action,
	{_CState263, CRparamToken}:                  _CGotoState317Action,
	{_CState263, CCommaToken}:                   _CGotoState187Action,
	{_CState264, CMulToken}:                     _CGotoState195Action,
	{_CState264, CDivToken}:                     _CGotoState193Action,
	{_CState264, CModToken}:                     _CGotoState194Action,
	{_CState265, CMulToken}:                     _CGotoState195Action,
	{_CState265, CDivToken}:                     _CGotoState193Action,
	{_CState265, CModToken}:                     _CGotoState194Action,
	{_CState266, CEqOpToken}:                    _CGotoState184Action,
	{_CState266, CNeOpToken}:                    _CGotoState185Action,
	{_CState268, CLeOpToken}:                    _CGotoState204Action,
	{_CState268, CGeOpToken}:                    _CGotoState202Action,
	{_CState268, CLtToken}:                      _CGotoState205Action,
	{_CState268, CGtToken}:                      _CGotoState203Action,
	{_CState269, CLeOpToken}:                    _CGotoState204Action,
	{_CState269, CGeOpToken}:                    _CGotoState202Action,
	{_CState269, CLtToken}:                      _CGotoState205Action,
	{_CState269, CGtToken}:                      _CGotoState203Action,
	{_CState270, CAndToken}:                     _CGotoState181Action,
	{_CState272, CHatToken}:                     _CGotoState186Action,
	{_CState273, COrToken}:                      _CGotoState189Action,
	{_CState274, CAndOpToken}:                   _CGotoState190Action,
	{_CState275, CColonToken}:                   _CGotoState318Action,
	{_CState275, CCommaToken}:                   _CGotoState187Action,
	{_CState280, CRbraceToken}:                  _CGotoState319Action,
	{_CState280, CCommaToken}:                   _CGotoState187Action,
	{_CState282, CRparamToken}:                  _CGotoState321Action,
	{_CState282, CCommaToken}:                   _CGotoState320Action,
	{_CState285, CLeftOpToken}:                  _CGotoState206Action,
	{_CState285, CRightOpToken}:                 _CGotoState207Action,
	{_CState286, CLeftOpToken}:                  _CGotoState206Action,
	{_CState286, CRightOpToken}:                 _CGotoState207Action,
	{_CState287, CLeftOpToken}:                  _CGotoState206Action,
	{_CState287, CRightOpToken}:                 _CGotoState207Action,
	{_CState288, CLeftOpToken}:                  _CGotoState206Action,
	{_CState288, CRightOpToken}:                 _CGotoState207Action,
	{_CState289, CMinusToken}:                   _CGotoState179Action,
	{_CState289, CPlusToken}:                    _CGotoState180Action,
	{_CState290, CMinusToken}:                   _CGotoState179Action,
	{_CState290, CPlusToken}:                    _CGotoState180Action,
	{_CState293, CRbraceToken}:                  _CGotoState322Action,
	{_CState295, CRparamToken}:                  _CGotoState323Action,
	{_CState296, CRparamToken}:                  _CGotoState324Action,
	{_CState297, CIdentifierToken}:              _CGotoState131Action,
	{_CState297, CConstantToken}:                _CGotoState78Action,
	{_CState297, CStringLiteralToken}:           _CGotoState97Action,
	{_CState297, CSizeofToken}:                  _CGotoState96Action,
	{_CState297, CIncOpToken}:                   _CGotoState88Action,
	{_CState297, CDecOpToken}:                   _CGotoState80Action,
	{_CState297, CLparamToken}:                  _CGotoState89Action,
	{_CState297, CRbraceToken}:                  _CGotoState325Action,
	{_CState297, CMulToken}:                     _CGotoState91Action,
	{_CState297, CMinusToken}:                   _CGotoState90Action,
	{_CState297, CPlusToken}:                    _CGotoState92Action,
	{_CState297, CAndToken}:                     _CGotoState75Action,
	{_CState297, CExclaimToken}:                 _CGotoState83Action,
	{_CState297, CTildaToken}:                   _CGotoState99Action,
	{_CState297, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState297, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState297, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState297, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState297, CCastExpressionType}:           _CGotoState104Action,
	{_CState297, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState297, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState297, CShiftExpressionType}:          _CGotoState123Action,
	{_CState297, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState297, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState297, CAndExpressionType}:            _CGotoState102Action,
	{_CState297, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState297, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState297, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState297, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState297, CConditionalExpressionType}:    _CGotoState133Action,
	{_CState297, CConstantExpressionType}:       _CGotoState326Action,
	{_CState298, CTypeNameToken}:                _CGotoState21Action,
	{_CState298, CTypedefToken}:                 _CGotoState20Action,
	{_CState298, CExternToken}:                  _CGotoState8Action,
	{_CState298, CStaticToken}:                  _CGotoState18Action,
	{_CState298, CAutoToken}:                    _CGotoState3Action,
	{_CState298, CRegisterToken}:                _CGotoState15Action,
	{_CState298, CCharToken}:                    _CGotoState4Action,
	{_CState298, CShortToken}:                   _CGotoState16Action,
	{_CState298, CIntToken}:                     _CGotoState11Action,
	{_CState298, CLongToken}:                    _CGotoState12Action,
	{_CState298, CSignedToken}:                  _CGotoState17Action,
	{_CState298, CUnsignedToken}:                _CGotoState23Action,
	{_CState298, CFloatToken}:                   _CGotoState9Action,
	{_CState298, CDoubleToken}:                  _CGotoState6Action,
	{_CState298, CConstToken}:                   _CGotoState5Action,
	{_CState298, CVolatileToken}:                _CGotoState25Action,
	{_CState298, CVoidToken}:                    _CGotoState24Action,
	{_CState298, CStructToken}:                  _CGotoState19Action,
	{_CState298, CUnionToken}:                   _CGotoState22Action,
	{_CState298, CEnumToken}:                    _CGotoState7Action,
	{_CState298, CRparamToken}:                  _CGotoState327Action,
	{_CState298, CDeclarationSpecifiersType}:    _CGotoState138Action,
	{_CState298, CStorageClassSpecifierType}:    _CGotoState34Action,
	{_CState298, CTypeSpecifierType}:            _CGotoState38Action,
	{_CState298, CStructOrUnionSpecifierType}:   _CGotoState36Action,
	{_CState298, CStructOrUnionType}:            _CGotoState35Action,
	{_CState298, CEnumSpecifierType}:            _CGotoState30Action,
	{_CState298, CTypeQualifierType}:            _CGotoState37Action,
	{_CState298, CParameterTypeListType}:        _CGotoState328Action,
	{_CState298, CParameterListType}:            _CGotoState141Action,
	{_CState298, CParameterDeclarationType}:     _CGotoState140Action,
	{_CState299, CLparamToken}:                  _CGotoState298Action,
	{_CState299, CLbraceToken}:                  _CGotoState297Action,
	{_CState305, CIdentifierToken}:              _CGotoState131Action,
	{_CState305, CConstantToken}:                _CGotoState78Action,
	{_CState305, CStringLiteralToken}:           _CGotoState97Action,
	{_CState305, CSizeofToken}:                  _CGotoState96Action,
	{_CState305, CIncOpToken}:                   _CGotoState88Action,
	{_CState305, CDecOpToken}:                   _CGotoState80Action,
	{_CState305, CLparamToken}:                  _CGotoState89Action,
	{_CState305, CMulToken}:                     _CGotoState91Action,
	{_CState305, CMinusToken}:                   _CGotoState90Action,
	{_CState305, CPlusToken}:                    _CGotoState92Action,
	{_CState305, CAndToken}:                     _CGotoState75Action,
	{_CState305, CExclaimToken}:                 _CGotoState83Action,
	{_CState305, CTildaToken}:                   _CGotoState99Action,
	{_CState305, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState305, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState305, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState305, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState305, CCastExpressionType}:           _CGotoState104Action,
	{_CState305, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState305, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState305, CShiftExpressionType}:          _CGotoState123Action,
	{_CState305, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState305, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState305, CAndExpressionType}:            _CGotoState102Action,
	{_CState305, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState305, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState305, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState305, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState305, CConditionalExpressionType}:    _CGotoState133Action,
	{_CState305, CConstantExpressionType}:       _CGotoState329Action,
	{_CState306, CIdentifierToken}:              _CGotoState10Action,
	{_CState306, CLparamToken}:                  _CGotoState13Action,
	{_CState306, CColonToken}:                   _CGotoState235Action,
	{_CState306, CMulToken}:                     _CGotoState14Action,
	{_CState306, CStructDeclaratorType}:         _CGotoState330Action,
	{_CState306, CDeclaratorType}:               _CGotoState236Action,
	{_CState306, CDirectDeclaratorType}:         _CGotoState29Action,
	{_CState306, CPointerType}:                  _CGotoState33Action,
	{_CState308, CIdentifierToken}:              _CGotoState131Action,
	{_CState308, CConstantToken}:                _CGotoState78Action,
	{_CState308, CStringLiteralToken}:           _CGotoState97Action,
	{_CState308, CSizeofToken}:                  _CGotoState96Action,
	{_CState308, CIncOpToken}:                   _CGotoState88Action,
	{_CState308, CDecOpToken}:                   _CGotoState80Action,
	{_CState308, CLparamToken}:                  _CGotoState89Action,
	{_CState308, CLcurlToken}:                   _CGotoState153Action,
	{_CState308, CRcurlToken}:                   _CGotoState331Action,
	{_CState308, CMulToken}:                     _CGotoState91Action,
	{_CState308, CMinusToken}:                   _CGotoState90Action,
	{_CState308, CPlusToken}:                    _CGotoState92Action,
	{_CState308, CAndToken}:                     _CGotoState75Action,
	{_CState308, CExclaimToken}:                 _CGotoState83Action,
	{_CState308, CTildaToken}:                   _CGotoState99Action,
	{_CState308, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState308, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState308, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState308, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState308, CCastExpressionType}:           _CGotoState104Action,
	{_CState308, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState308, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState308, CShiftExpressionType}:          _CGotoState123Action,
	{_CState308, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState308, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState308, CAndExpressionType}:            _CGotoState102Action,
	{_CState308, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState308, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState308, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState308, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState308, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState308, CAssignmentExpressionType}:     _CGotoState154Action,
	{_CState308, CInitializerType}:              _CGotoState332Action,
	{_CState311, CIdentifierToken}:              _CGotoState131Action,
	{_CState311, CConstantToken}:                _CGotoState78Action,
	{_CState311, CStringLiteralToken}:           _CGotoState97Action,
	{_CState311, CSizeofToken}:                  _CGotoState96Action,
	{_CState311, CIncOpToken}:                   _CGotoState88Action,
	{_CState311, CDecOpToken}:                   _CGotoState80Action,
	{_CState311, CLparamToken}:                  _CGotoState89Action,
	{_CState311, CMulToken}:                     _CGotoState91Action,
	{_CState311, CMinusToken}:                   _CGotoState90Action,
	{_CState311, CPlusToken}:                    _CGotoState92Action,
	{_CState311, CAndToken}:                     _CGotoState75Action,
	{_CState311, CExclaimToken}:                 _CGotoState83Action,
	{_CState311, CTildaToken}:                   _CGotoState99Action,
	{_CState311, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState311, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState311, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState311, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState311, CCastExpressionType}:           _CGotoState104Action,
	{_CState311, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState311, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState311, CShiftExpressionType}:          _CGotoState123Action,
	{_CState311, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState311, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState311, CAndExpressionType}:            _CGotoState102Action,
	{_CState311, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState311, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState311, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState311, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState311, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState311, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState311, CExpressionType}:               _CGotoState333Action,
	{_CState312, CIdentifierToken}:              _CGotoState131Action,
	{_CState312, CConstantToken}:                _CGotoState78Action,
	{_CState312, CStringLiteralToken}:           _CGotoState97Action,
	{_CState312, CSizeofToken}:                  _CGotoState96Action,
	{_CState312, CIncOpToken}:                   _CGotoState88Action,
	{_CState312, CDecOpToken}:                   _CGotoState80Action,
	{_CState312, CLparamToken}:                  _CGotoState89Action,
	{_CState312, CRparamToken}:                  _CGotoState334Action,
	{_CState312, CMulToken}:                     _CGotoState91Action,
	{_CState312, CMinusToken}:                   _CGotoState90Action,
	{_CState312, CPlusToken}:                    _CGotoState92Action,
	{_CState312, CAndToken}:                     _CGotoState75Action,
	{_CState312, CExclaimToken}:                 _CGotoState83Action,
	{_CState312, CTildaToken}:                   _CGotoState99Action,
	{_CState312, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState312, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState312, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState312, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState312, CCastExpressionType}:           _CGotoState104Action,
	{_CState312, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState312, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState312, CShiftExpressionType}:          _CGotoState123Action,
	{_CState312, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState312, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState312, CAndExpressionType}:            _CGotoState102Action,
	{_CState312, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState312, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState312, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState312, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState312, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState312, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState312, CExpressionType}:               _CGotoState335Action,
	{_CState313, CIdentifierToken}:              _CGotoState86Action,
	{_CState313, CConstantToken}:                _CGotoState78Action,
	{_CState313, CStringLiteralToken}:           _CGotoState97Action,
	{_CState313, CSizeofToken}:                  _CGotoState96Action,
	{_CState313, CIncOpToken}:                   _CGotoState88Action,
	{_CState313, CDecOpToken}:                   _CGotoState80Action,
	{_CState313, CCaseToken}:                    _CGotoState77Action,
	{_CState313, CDefaultToken}:                 _CGotoState81Action,
	{_CState313, CIfToken}:                      _CGotoState87Action,
	{_CState313, CSwitchToken}:                  _CGotoState98Action,
	{_CState313, CWhileToken}:                   _CGotoState100Action,
	{_CState313, CDoToken}:                      _CGotoState82Action,
	{_CState313, CForToken}:                     _CGotoState84Action,
	{_CState313, CGotoToken}:                    _CGotoState85Action,
	{_CState313, CContinueToken}:                _CGotoState79Action,
	{_CState313, CBreakToken}:                   _CGotoState76Action,
	{_CState313, CReturnToken}:                  _CGotoState94Action,
	{_CState313, CLparamToken}:                  _CGotoState89Action,
	{_CState313, CLcurlToken}:                   _CGotoState49Action,
	{_CState313, CSemicolonToken}:               _CGotoState95Action,
	{_CState313, CMulToken}:                     _CGotoState91Action,
	{_CState313, CMinusToken}:                   _CGotoState90Action,
	{_CState313, CPlusToken}:                    _CGotoState92Action,
	{_CState313, CAndToken}:                     _CGotoState75Action,
	{_CState313, CExclaimToken}:                 _CGotoState83Action,
	{_CState313, CTildaToken}:                   _CGotoState99Action,
	{_CState313, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState313, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState313, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState313, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState313, CCastExpressionType}:           _CGotoState104Action,
	{_CState313, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState313, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState313, CShiftExpressionType}:          _CGotoState123Action,
	{_CState313, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState313, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState313, CAndExpressionType}:            _CGotoState102Action,
	{_CState313, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState313, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState313, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState313, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState313, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState313, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState313, CExpressionType}:               _CGotoState110Action,
	{_CState313, CStatementType}:                _CGotoState338Action,
	{_CState313, CLabeledStatementType}:         _CGotoState115Action,
	{_CState313, CCompoundStatementType}:        _CGotoState105Action,
	{_CState313, CExpressionStatementType}:      _CGotoState111Action,
	{_CState313, CSelectionStatementType}:       _CGotoState122Action,
	{_CState313, CIterationStatementType}:       _CGotoState113Action,
	{_CState313, CJumpStatementType}:            _CGotoState114Action,
	{_CState316, CIdentifierToken}:              _CGotoState86Action,
	{_CState316, CConstantToken}:                _CGotoState78Action,
	{_CState316, CStringLiteralToken}:           _CGotoState97Action,
	{_CState316, CSizeofToken}:                  _CGotoState96Action,
	{_CState316, CIncOpToken}:                   _CGotoState88Action,
	{_CState316, CDecOpToken}:                   _CGotoState80Action,
	{_CState316, CCaseToken}:                    _CGotoState77Action,
	{_CState316, CDefaultToken}:                 _CGotoState81Action,
	{_CState316, CIfToken}:                      _CGotoState87Action,
	{_CState316, CSwitchToken}:                  _CGotoState98Action,
	{_CState316, CWhileToken}:                   _CGotoState100Action,
	{_CState316, CDoToken}:                      _CGotoState82Action,
	{_CState316, CForToken}:                     _CGotoState84Action,
	{_CState316, CGotoToken}:                    _CGotoState85Action,
	{_CState316, CContinueToken}:                _CGotoState79Action,
	{_CState316, CBreakToken}:                   _CGotoState76Action,
	{_CState316, CReturnToken}:                  _CGotoState94Action,
	{_CState316, CLparamToken}:                  _CGotoState89Action,
	{_CState316, CLcurlToken}:                   _CGotoState49Action,
	{_CState316, CSemicolonToken}:               _CGotoState95Action,
	{_CState316, CMulToken}:                     _CGotoState91Action,
	{_CState316, CMinusToken}:                   _CGotoState90Action,
	{_CState316, CPlusToken}:                    _CGotoState92Action,
	{_CState316, CAndToken}:                     _CGotoState75Action,
	{_CState316, CExclaimToken}:                 _CGotoState83Action,
	{_CState316, CTildaToken}:                   _CGotoState99Action,
	{_CState316, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState316, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState316, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState316, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState316, CCastExpressionType}:           _CGotoState104Action,
	{_CState316, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState316, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState316, CShiftExpressionType}:          _CGotoState123Action,
	{_CState316, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState316, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState316, CAndExpressionType}:            _CGotoState102Action,
	{_CState316, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState316, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState316, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState316, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState316, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState316, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState316, CExpressionType}:               _CGotoState110Action,
	{_CState316, CStatementType}:                _CGotoState339Action,
	{_CState316, CLabeledStatementType}:         _CGotoState115Action,
	{_CState316, CCompoundStatementType}:        _CGotoState105Action,
	{_CState316, CExpressionStatementType}:      _CGotoState111Action,
	{_CState316, CSelectionStatementType}:       _CGotoState122Action,
	{_CState316, CIterationStatementType}:       _CGotoState113Action,
	{_CState316, CJumpStatementType}:            _CGotoState114Action,
	{_CState317, CIdentifierToken}:              _CGotoState86Action,
	{_CState317, CConstantToken}:                _CGotoState78Action,
	{_CState317, CStringLiteralToken}:           _CGotoState97Action,
	{_CState317, CSizeofToken}:                  _CGotoState96Action,
	{_CState317, CIncOpToken}:                   _CGotoState88Action,
	{_CState317, CDecOpToken}:                   _CGotoState80Action,
	{_CState317, CCaseToken}:                    _CGotoState77Action,
	{_CState317, CDefaultToken}:                 _CGotoState81Action,
	{_CState317, CIfToken}:                      _CGotoState87Action,
	{_CState317, CSwitchToken}:                  _CGotoState98Action,
	{_CState317, CWhileToken}:                   _CGotoState100Action,
	{_CState317, CDoToken}:                      _CGotoState82Action,
	{_CState317, CForToken}:                     _CGotoState84Action,
	{_CState317, CGotoToken}:                    _CGotoState85Action,
	{_CState317, CContinueToken}:                _CGotoState79Action,
	{_CState317, CBreakToken}:                   _CGotoState76Action,
	{_CState317, CReturnToken}:                  _CGotoState94Action,
	{_CState317, CLparamToken}:                  _CGotoState89Action,
	{_CState317, CLcurlToken}:                   _CGotoState49Action,
	{_CState317, CSemicolonToken}:               _CGotoState95Action,
	{_CState317, CMulToken}:                     _CGotoState91Action,
	{_CState317, CMinusToken}:                   _CGotoState90Action,
	{_CState317, CPlusToken}:                    _CGotoState92Action,
	{_CState317, CAndToken}:                     _CGotoState75Action,
	{_CState317, CExclaimToken}:                 _CGotoState83Action,
	{_CState317, CTildaToken}:                   _CGotoState99Action,
	{_CState317, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState317, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState317, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState317, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState317, CCastExpressionType}:           _CGotoState104Action,
	{_CState317, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState317, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState317, CShiftExpressionType}:          _CGotoState123Action,
	{_CState317, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState317, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState317, CAndExpressionType}:            _CGotoState102Action,
	{_CState317, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState317, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState317, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState317, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState317, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState317, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState317, CExpressionType}:               _CGotoState110Action,
	{_CState317, CStatementType}:                _CGotoState340Action,
	{_CState317, CLabeledStatementType}:         _CGotoState115Action,
	{_CState317, CCompoundStatementType}:        _CGotoState105Action,
	{_CState317, CExpressionStatementType}:      _CGotoState111Action,
	{_CState317, CSelectionStatementType}:       _CGotoState122Action,
	{_CState317, CIterationStatementType}:       _CGotoState113Action,
	{_CState317, CJumpStatementType}:            _CGotoState114Action,
	{_CState318, CIdentifierToken}:              _CGotoState131Action,
	{_CState318, CConstantToken}:                _CGotoState78Action,
	{_CState318, CStringLiteralToken}:           _CGotoState97Action,
	{_CState318, CSizeofToken}:                  _CGotoState96Action,
	{_CState318, CIncOpToken}:                   _CGotoState88Action,
	{_CState318, CDecOpToken}:                   _CGotoState80Action,
	{_CState318, CLparamToken}:                  _CGotoState89Action,
	{_CState318, CMulToken}:                     _CGotoState91Action,
	{_CState318, CMinusToken}:                   _CGotoState90Action,
	{_CState318, CPlusToken}:                    _CGotoState92Action,
	{_CState318, CAndToken}:                     _CGotoState75Action,
	{_CState318, CExclaimToken}:                 _CGotoState83Action,
	{_CState318, CTildaToken}:                   _CGotoState99Action,
	{_CState318, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState318, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState318, CUnaryExpressionType}:          _CGotoState135Action,
	{_CState318, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState318, CCastExpressionType}:           _CGotoState104Action,
	{_CState318, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState318, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState318, CShiftExpressionType}:          _CGotoState123Action,
	{_CState318, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState318, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState318, CAndExpressionType}:            _CGotoState102Action,
	{_CState318, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState318, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState318, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState318, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState318, CConditionalExpressionType}:    _CGotoState341Action,
	{_CState320, CIdentifierToken}:              _CGotoState131Action,
	{_CState320, CConstantToken}:                _CGotoState78Action,
	{_CState320, CStringLiteralToken}:           _CGotoState97Action,
	{_CState320, CSizeofToken}:                  _CGotoState96Action,
	{_CState320, CIncOpToken}:                   _CGotoState88Action,
	{_CState320, CDecOpToken}:                   _CGotoState80Action,
	{_CState320, CLparamToken}:                  _CGotoState89Action,
	{_CState320, CMulToken}:                     _CGotoState91Action,
	{_CState320, CMinusToken}:                   _CGotoState90Action,
	{_CState320, CPlusToken}:                    _CGotoState92Action,
	{_CState320, CAndToken}:                     _CGotoState75Action,
	{_CState320, CExclaimToken}:                 _CGotoState83Action,
	{_CState320, CTildaToken}:                   _CGotoState99Action,
	{_CState320, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState320, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState320, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState320, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState320, CCastExpressionType}:           _CGotoState104Action,
	{_CState320, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState320, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState320, CShiftExpressionType}:          _CGotoState123Action,
	{_CState320, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState320, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState320, CAndExpressionType}:            _CGotoState102Action,
	{_CState320, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState320, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState320, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState320, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState320, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState320, CAssignmentExpressionType}:     _CGotoState342Action,
	{_CState326, CRbraceToken}:                  _CGotoState343Action,
	{_CState328, CRparamToken}:                  _CGotoState344Action,
	{_CState333, CRparamToken}:                  _CGotoState345Action,
	{_CState333, CCommaToken}:                   _CGotoState187Action,
	{_CState334, CIdentifierToken}:              _CGotoState86Action,
	{_CState334, CConstantToken}:                _CGotoState78Action,
	{_CState334, CStringLiteralToken}:           _CGotoState97Action,
	{_CState334, CSizeofToken}:                  _CGotoState96Action,
	{_CState334, CIncOpToken}:                   _CGotoState88Action,
	{_CState334, CDecOpToken}:                   _CGotoState80Action,
	{_CState334, CCaseToken}:                    _CGotoState77Action,
	{_CState334, CDefaultToken}:                 _CGotoState81Action,
	{_CState334, CIfToken}:                      _CGotoState87Action,
	{_CState334, CSwitchToken}:                  _CGotoState98Action,
	{_CState334, CWhileToken}:                   _CGotoState100Action,
	{_CState334, CDoToken}:                      _CGotoState82Action,
	{_CState334, CForToken}:                     _CGotoState84Action,
	{_CState334, CGotoToken}:                    _CGotoState85Action,
	{_CState334, CContinueToken}:                _CGotoState79Action,
	{_CState334, CBreakToken}:                   _CGotoState76Action,
	{_CState334, CReturnToken}:                  _CGotoState94Action,
	{_CState334, CLparamToken}:                  _CGotoState89Action,
	{_CState334, CLcurlToken}:                   _CGotoState49Action,
	{_CState334, CSemicolonToken}:               _CGotoState95Action,
	{_CState334, CMulToken}:                     _CGotoState91Action,
	{_CState334, CMinusToken}:                   _CGotoState90Action,
	{_CState334, CPlusToken}:                    _CGotoState92Action,
	{_CState334, CAndToken}:                     _CGotoState75Action,
	{_CState334, CExclaimToken}:                 _CGotoState83Action,
	{_CState334, CTildaToken}:                   _CGotoState99Action,
	{_CState334, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState334, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState334, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState334, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState334, CCastExpressionType}:           _CGotoState104Action,
	{_CState334, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState334, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState334, CShiftExpressionType}:          _CGotoState123Action,
	{_CState334, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState334, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState334, CAndExpressionType}:            _CGotoState102Action,
	{_CState334, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState334, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState334, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState334, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState334, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState334, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState334, CExpressionType}:               _CGotoState110Action,
	{_CState334, CStatementType}:                _CGotoState346Action,
	{_CState334, CLabeledStatementType}:         _CGotoState115Action,
	{_CState334, CCompoundStatementType}:        _CGotoState105Action,
	{_CState334, CExpressionStatementType}:      _CGotoState111Action,
	{_CState334, CSelectionStatementType}:       _CGotoState122Action,
	{_CState334, CIterationStatementType}:       _CGotoState113Action,
	{_CState334, CJumpStatementType}:            _CGotoState114Action,
	{_CState335, CRparamToken}:                  _CGotoState347Action,
	{_CState335, CCommaToken}:                   _CGotoState187Action,
	{_CState336, CElseToken}:                    _CGotoState348Action,
	{_CState337, CElseToken}:                    _CGotoState348Action,
	{_CState338, CElseToken}:                    _CGotoState348Action,
	{_CState345, CSemicolonToken}:               _CGotoState349Action,
	{_CState347, CIdentifierToken}:              _CGotoState86Action,
	{_CState347, CConstantToken}:                _CGotoState78Action,
	{_CState347, CStringLiteralToken}:           _CGotoState97Action,
	{_CState347, CSizeofToken}:                  _CGotoState96Action,
	{_CState347, CIncOpToken}:                   _CGotoState88Action,
	{_CState347, CDecOpToken}:                   _CGotoState80Action,
	{_CState347, CCaseToken}:                    _CGotoState77Action,
	{_CState347, CDefaultToken}:                 _CGotoState81Action,
	{_CState347, CIfToken}:                      _CGotoState87Action,
	{_CState347, CSwitchToken}:                  _CGotoState98Action,
	{_CState347, CWhileToken}:                   _CGotoState100Action,
	{_CState347, CDoToken}:                      _CGotoState82Action,
	{_CState347, CForToken}:                     _CGotoState84Action,
	{_CState347, CGotoToken}:                    _CGotoState85Action,
	{_CState347, CContinueToken}:                _CGotoState79Action,
	{_CState347, CBreakToken}:                   _CGotoState76Action,
	{_CState347, CReturnToken}:                  _CGotoState94Action,
	{_CState347, CLparamToken}:                  _CGotoState89Action,
	{_CState347, CLcurlToken}:                   _CGotoState49Action,
	{_CState347, CSemicolonToken}:               _CGotoState95Action,
	{_CState347, CMulToken}:                     _CGotoState91Action,
	{_CState347, CMinusToken}:                   _CGotoState90Action,
	{_CState347, CPlusToken}:                    _CGotoState92Action,
	{_CState347, CAndToken}:                     _CGotoState75Action,
	{_CState347, CExclaimToken}:                 _CGotoState83Action,
	{_CState347, CTildaToken}:                   _CGotoState99Action,
	{_CState347, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState347, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState347, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState347, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState347, CCastExpressionType}:           _CGotoState104Action,
	{_CState347, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState347, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState347, CShiftExpressionType}:          _CGotoState123Action,
	{_CState347, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState347, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState347, CAndExpressionType}:            _CGotoState102Action,
	{_CState347, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState347, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState347, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState347, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState347, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState347, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState347, CExpressionType}:               _CGotoState110Action,
	{_CState347, CStatementType}:                _CGotoState350Action,
	{_CState347, CLabeledStatementType}:         _CGotoState115Action,
	{_CState347, CCompoundStatementType}:        _CGotoState105Action,
	{_CState347, CExpressionStatementType}:      _CGotoState111Action,
	{_CState347, CSelectionStatementType}:       _CGotoState122Action,
	{_CState347, CIterationStatementType}:       _CGotoState113Action,
	{_CState347, CJumpStatementType}:            _CGotoState114Action,
	{_CState348, CIdentifierToken}:              _CGotoState86Action,
	{_CState348, CConstantToken}:                _CGotoState78Action,
	{_CState348, CStringLiteralToken}:           _CGotoState97Action,
	{_CState348, CSizeofToken}:                  _CGotoState96Action,
	{_CState348, CIncOpToken}:                   _CGotoState88Action,
	{_CState348, CDecOpToken}:                   _CGotoState80Action,
	{_CState348, CCaseToken}:                    _CGotoState77Action,
	{_CState348, CDefaultToken}:                 _CGotoState81Action,
	{_CState348, CIfToken}:                      _CGotoState87Action,
	{_CState348, CSwitchToken}:                  _CGotoState98Action,
	{_CState348, CWhileToken}:                   _CGotoState100Action,
	{_CState348, CDoToken}:                      _CGotoState82Action,
	{_CState348, CForToken}:                     _CGotoState84Action,
	{_CState348, CGotoToken}:                    _CGotoState85Action,
	{_CState348, CContinueToken}:                _CGotoState79Action,
	{_CState348, CBreakToken}:                   _CGotoState76Action,
	{_CState348, CReturnToken}:                  _CGotoState94Action,
	{_CState348, CLparamToken}:                  _CGotoState89Action,
	{_CState348, CLcurlToken}:                   _CGotoState49Action,
	{_CState348, CSemicolonToken}:               _CGotoState95Action,
	{_CState348, CMulToken}:                     _CGotoState91Action,
	{_CState348, CMinusToken}:                   _CGotoState90Action,
	{_CState348, CPlusToken}:                    _CGotoState92Action,
	{_CState348, CAndToken}:                     _CGotoState75Action,
	{_CState348, CExclaimToken}:                 _CGotoState83Action,
	{_CState348, CTildaToken}:                   _CGotoState99Action,
	{_CState348, CPrimaryExpressionType}:        _CGotoState120Action,
	{_CState348, CPostfixExpressionType}:        _CGotoState119Action,
	{_CState348, CUnaryExpressionType}:          _CGotoState126Action,
	{_CState348, CUnaryOperatorType}:            _CGotoState127Action,
	{_CState348, CCastExpressionType}:           _CGotoState104Action,
	{_CState348, CMultiplicativeExpressionType}: _CGotoState118Action,
	{_CState348, CAdditiveExpressionType}:       _CGotoState101Action,
	{_CState348, CShiftExpressionType}:          _CGotoState123Action,
	{_CState348, CRelationalExpressionType}:     _CGotoState121Action,
	{_CState348, CEqualityExpressionType}:       _CGotoState108Action,
	{_CState348, CAndExpressionType}:            _CGotoState102Action,
	{_CState348, CExclusiveOrExpressionType}:    _CGotoState109Action,
	{_CState348, CInclusiveOrExpressionType}:    _CGotoState112Action,
	{_CState348, CLogicalAndExpressionType}:     _CGotoState116Action,
	{_CState348, CLogicalOrExpressionType}:      _CGotoState117Action,
	{_CState348, CConditionalExpressionType}:    _CGotoState106Action,
	{_CState348, CAssignmentExpressionType}:     _CGotoState103Action,
	{_CState348, CExpressionType}:               _CGotoState110Action,
	{_CState348, CStatementType}:                _CGotoState351Action,
	{_CState348, CLabeledStatementType}:         _CGotoState115Action,
	{_CState348, CCompoundStatementType}:        _CGotoState105Action,
	{_CState348, CExpressionStatementType}:      _CGotoState111Action,
	{_CState348, CSelectionStatementType}:       _CGotoState122Action,
	{_CState348, CIterationStatementType}:       _CGotoState113Action,
	{_CState348, CJumpStatementType}:            _CGotoState114Action,
	{_CState3, _CWildcardMarker}:                _CReduceDToStorageClassSpecifierAction,
	{_CState4, _CWildcardMarker}:                _CReduceBToTypeSpecifierAction,
	{_CState5, _CWildcardMarker}:                _CReduceAToTypeQualifierAction,
	{_CState6, _CWildcardMarker}:                _CReduceGToTypeSpecifierAction,
	{_CState8, _CWildcardMarker}:                _CReduceBToStorageClassSpecifierAction,
	{_CState9, _CWildcardMarker}:                _CReduceFToTypeSpecifierAction,
	{_CState10, _CWildcardMarker}:               _CReduceAToDirectDeclaratorAction,
	{_CState11, _CWildcardMarker}:               _CReduceDToTypeSpecifierAction,
	{_CState12, _CWildcardMarker}:               _CReduceEToTypeSpecifierAction,
	{_CState14, _CWildcardMarker}:               _CReduceAToPointerAction,
	{_CState15, _CWildcardMarker}:               _CReduceEToStorageClassSpecifierAction,
	{_CState16, _CWildcardMarker}:               _CReduceCToTypeSpecifierAction,
	{_CState17, _CWildcardMarker}:               _CReduceHToTypeSpecifierAction,
	{_CState18, _CWildcardMarker}:               _CReduceCToStorageClassSpecifierAction,
	{_CState19, _CWildcardMarker}:               _CReduceAToStructOrUnionAction,
	{_CState20, _CWildcardMarker}:               _CReduceAToStorageClassSpecifierAction,
	{_CState21, _CWildcardMarker}:               _CReduceLToTypeSpecifierAction,
	{_CState22, _CWildcardMarker}:               _CReduceBToStructOrUnionAction,
	{_CState23, _CWildcardMarker}:               _CReduceIToTypeSpecifierAction,
	{_CState24, _CWildcardMarker}:               _CReduceAToTypeSpecifierAction,
	{_CState25, _CWildcardMarker}:               _CReduceBToTypeQualifierAction,
	{_CState26, _CWildcardMarker}:               _CReduceBToExternalDeclarationAction,
	{_CState29, _CWildcardMarker}:               _CReduceBToDeclaratorAction,
	{_CState30, _CWildcardMarker}:               _CReduceKToTypeSpecifierAction,
	{_CState31, _CWildcardMarker}:               _CReduceAToTranslationUnitAction,
	{_CState32, _CWildcardMarker}:               _CReduceAToExternalDeclarationAction,
	{_CState34, _CWildcardMarker}:               _CReduceAToDeclarationSpecifiersAction,
	{_CState36, _CWildcardMarker}:               _CReduceJToTypeSpecifierAction,
	{_CState37, _CWildcardMarker}:               _CReduceEToDeclarationSpecifiersAction,
	{_CState38, _CWildcardMarker}:               _CReduceCToDeclarationSpecifiersAction,
	{_CState39, _CWildcardMarker}:               _CReduceCToEnumSpecifierAction,
	{_CState42, _CWildcardMarker}:               _CReduceCToPointerAction,
	{_CState43, _CWildcardMarker}:               _CReduceAToTypeQualifierListAction,
	{_CState44, _CWildcardMarker}:               _CReduceBToPointerAction,
	{_CState45, _CWildcardMarker}:               _CReduceAToDeclarationAction,
	{_CState46, _CWildcardMarker}:               _CReduceAToInitDeclaratorAction,
	{_CState47, _CWildcardMarker}:               _CReduceAToInitDeclaratorListAction,
	{_CState50, _CWildcardMarker}:               _CReduceDToFunctionDefinitionAction,
	{_CState51, _CWildcardMarker}:               _CReduceAToDeclarationListAction,
	{_CState56, _CWildcardMarker}:               _CReduceAToDeclaratorAction,
	{_CState57, _CWildcardMarker}:               _CReduceBToDeclarationSpecifiersAction,
	{_CState58, _CWildcardMarker}:               _CReduceCToStructOrUnionSpecifierAction,
	{_CState60, _CWildcardMarker}:               _CReduceBToTranslationUnitAction,
	{_CState61, _CWildcardMarker}:               _CReduceFToDeclarationSpecifiersAction,
	{_CState62, _CWildcardMarker}:               _CReduceDToDeclarationSpecifiersAction,
	{_CState64, _CWildcardMarker}:               _CReduceAToEnumeratorAction,
	{_CState65, _CWildcardMarker}:               _CReduceAToEnumeratorListAction,
	{_CState67, _CWildcardMarker}:               _CReduceBToDirectDeclaratorAction,
	{_CState68, _CWildcardMarker}:               _CReduceDToPointerAction,
	{_CState69, _CWildcardMarker}:               _CReduceBToTypeQualifierListAction,
	{_CState71, _CWildcardMarker}:               _CReduceBToFunctionDefinitionAction,
	{_CState74, _CWildcardMarker}:               _CReduceBToDeclarationAction,
	{_CState75, _CWildcardMarker}:               _CReduceAToUnaryOperatorAction,
	{_CState78, _CWildcardMarker}:               _CReduceBToPrimaryExpressionAction,
	{_CState83, _CWildcardMarker}:               _CReduceFToUnaryOperatorAction,
	{_CState86, _CWildcardMarker}:               _CReduceAToPrimaryExpressionAction,
	{_CState90, _CWildcardMarker}:               _CReduceDToUnaryOperatorAction,
	{_CState91, _CWildcardMarker}:               _CReduceBToUnaryOperatorAction,
	{_CState92, _CWildcardMarker}:               _CReduceCToUnaryOperatorAction,
	{_CState93, _CWildcardMarker}:               _CReduceAToCompoundStatementAction,
	{_CState95, _CWildcardMarker}:               _CReduceAToExpressionStatementAction,
	{_CState97, _CWildcardMarker}:               _CReduceCToPrimaryExpressionAction,
	{_CState99, _CWildcardMarker}:               _CReduceEToUnaryOperatorAction,
	{_CState101, _CWildcardMarker}:              _CReduceAToShiftExpressionAction,
	{_CState102, _CWildcardMarker}:              _CReduceAToExclusiveOrExpressionAction,
	{_CState103, _CWildcardMarker}:              _CReduceAToExpressionAction,
	{_CState104, _CWildcardMarker}:              _CReduceAToMultiplicativeExpressionAction,
	{_CState105, _CWildcardMarker}:              _CReduceBToStatementAction,
	{_CState106, _CWildcardMarker}:              _CReduceAToAssignmentExpressionAction,
	{_CState108, _CWildcardMarker}:              _CReduceAToAndExpressionAction,
	{_CState109, _CWildcardMarker}:              _CReduceAToInclusiveOrExpressionAction,
	{_CState111, _CWildcardMarker}:              _CReduceCToStatementAction,
	{_CState112, _CWildcardMarker}:              _CReduceAToLogicalAndExpressionAction,
	{_CState113, _CWildcardMarker}:              _CReduceEToStatementAction,
	{_CState114, _CWildcardMarker}:              _CReduceFToStatementAction,
	{_CState115, _CWildcardMarker}:              _CReduceAToStatementAction,
	{_CState116, _CWildcardMarker}:              _CReduceAToLogicalOrExpressionAction,
	{_CState117, _CWildcardMarker}:              _CReduceAToConditionalExpressionAction,
	{_CState118, _CWildcardMarker}:              _CReduceAToAdditiveExpressionAction,
	{_CState119, _CWildcardMarker}:              _CReduceAToUnaryExpressionAction,
	{_CState120, _CWildcardMarker}:              _CReduceAToPostfixExpressionAction,
	{_CState121, _CWildcardMarker}:              _CReduceAToEqualityExpressionAction,
	{_CState122, _CWildcardMarker}:              _CReduceDToStatementAction,
	{_CState123, _CWildcardMarker}:              _CReduceAToRelationalExpressionAction,
	{_CState124, _CWildcardMarker}:              _CReduceAToStatementListAction,
	{_CState126, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState128, _CWildcardMarker}:              _CReduceCToFunctionDefinitionAction,
	{_CState129, _CWildcardMarker}:              _CReduceBToDeclarationListAction,
	{_CState130, _CWildcardMarker}:              _CReduceAToInitDeclaratorAction,
	{_CState131, _CWildcardMarker}:              _CReduceAToPrimaryExpressionAction,
	{_CState132, _CWildcardMarker}:              _CReduceDToDirectDeclaratorAction,
	{_CState133, _CWildcardMarker}:              _CReduceAToConstantExpressionAction,
	{_CState135, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState136, _CWildcardMarker}:              _CReduceAToIdentifierListAction,
	{_CState137, _CWildcardMarker}:              _CReduceGToDirectDeclaratorAction,
	{_CState138, _CWildcardMarker}:              _CReduceCToParameterDeclarationAction,
	{_CState140, _CWildcardMarker}:              _CReduceAToParameterListAction,
	{_CState141, CRparamToken}:                  _CReduceAToParameterTypeListAction,
	{_CState145, _CWildcardMarker}:              _CReduceAToStructDeclarationListAction,
	{_CState147, _CWildcardMarker}:              _CReduceDToSpecifierQualifierListAction,
	{_CState148, _CWildcardMarker}:              _CReduceBToSpecifierQualifierListAction,
	{_CState152, _CWildcardMarker}:              _CReduceAToEnumSpecifierAction,
	{_CState154, _CWildcardMarker}:              _CReduceAToInitializerAction,
	{_CState155, _CWildcardMarker}:              _CReduceBToInitDeclaratorAction,
	{_CState156, _CWildcardMarker}:              _CReduceAToFunctionDefinitionAction,
	{_CState157, _CWildcardMarker}:              _CReduceBToInitDeclaratorListAction,
	{_CState158, _CWildcardMarker}:              _CReduceCToJumpStatementAction,
	{_CState160, _CWildcardMarker}:              _CReduceBToJumpStatementAction,
	{_CState162, _CWildcardMarker}:              _CReduceCToUnaryExpressionAction,
	{_CState169, _CWildcardMarker}:              _CReduceBToUnaryExpressionAction,
	{_CState171, CRparamToken}:                  _CReduceAToTypeNameAction,
	{_CState173, _CWildcardMarker}:              _CReduceDToJumpStatementAction,
	{_CState176, _CWildcardMarker}:              _CReduceEToUnaryExpressionAction,
	{_CState182, _CWildcardMarker}:              _CReduceCToCompoundStatementAction,
	{_CState188, _CWildcardMarker}:              _CReduceBToExpressionStatementAction,
	{_CState196, _CWildcardMarker}:              _CReduceHToPostfixExpressionAction,
	{_CState198, _CWildcardMarker}:              _CReduceGToPostfixExpressionAction,
	{_CState208, _CWildcardMarker}:              _CReduceBToCompoundStatementAction,
	{_CState209, _CWildcardMarker}:              _CReduceBToStatementListAction,
	{_CState210, _CWildcardMarker}:              _CReduceEToAssignmentOperatorAction,
	{_CState211, _CWildcardMarker}:              _CReduceIToAssignmentOperatorAction,
	{_CState212, _CWildcardMarker}:              _CReduceCToAssignmentOperatorAction,
	{_CState213, _CWildcardMarker}:              _CReduceAToAssignmentOperatorAction,
	{_CState214, _CWildcardMarker}:              _CReduceGToAssignmentOperatorAction,
	{_CState215, _CWildcardMarker}:              _CReduceDToAssignmentOperatorAction,
	{_CState216, _CWildcardMarker}:              _CReduceBToAssignmentOperatorAction,
	{_CState217, _CWildcardMarker}:              _CReduceKToAssignmentOperatorAction,
	{_CState218, _CWildcardMarker}:              _CReduceHToAssignmentOperatorAction,
	{_CState219, _CWildcardMarker}:              _CReduceFToAssignmentOperatorAction,
	{_CState220, _CWildcardMarker}:              _CReduceJToAssignmentOperatorAction,
	{_CState222, _CWildcardMarker}:              _CReduceDToUnaryExpressionAction,
	{_CState223, _CWildcardMarker}:              _CReduceCToDirectDeclaratorAction,
	{_CState226, _CWildcardMarker}:              _CReduceBToParameterDeclarationAction,
	{_CState227, _CWildcardMarker}:              _CReduceAToParameterDeclarationAction,
	{_CState228, _CWildcardMarker}:              _CReduceBToAbstractDeclaratorAction,
	{_CState229, _CWildcardMarker}:              _CReduceAToAbstractDeclaratorAction,
	{_CState231, _CWildcardMarker}:              _CReduceFToDirectDeclaratorAction,
	{_CState233, _CWildcardMarker}:              _CReduceEToDirectDeclaratorAction,
	{_CState236, _CWildcardMarker}:              _CReduceAToStructDeclaratorAction,
	{_CState237, _CWildcardMarker}:              _CReduceAToStructDeclaratorListAction,
	{_CState239, _CWildcardMarker}:              _CReduceBToStructOrUnionSpecifierAction,
	{_CState240, _CWildcardMarker}:              _CReduceBToStructDeclarationListAction,
	{_CState241, _CWildcardMarker}:              _CReduceCToSpecifierQualifierListAction,
	{_CState242, _CWildcardMarker}:              _CReduceAToSpecifierQualifierListAction,
	{_CState243, _CWildcardMarker}:              _CReduceBToEnumSpecifierAction,
	{_CState244, _CWildcardMarker}:              _CReduceBToEnumeratorAction,
	{_CState245, _CWildcardMarker}:              _CReduceBToEnumeratorListAction,
	{_CState246, _CWildcardMarker}:              _CReduceAToInitializerListAction,
	{_CState249, _CWildcardMarker}:              _CReduceCToLabeledStatementAction,
	{_CState252, _CWildcardMarker}:              _CReduceAToJumpStatementAction,
	{_CState253, _CWildcardMarker}:              _CReduceAToLabeledStatementAction,
	{_CState255, _CWildcardMarker}:              _CReduceDToPrimaryExpressionAction,
	{_CState257, CRparamToken}:                  _CReduceBToTypeNameAction,
	{_CState258, CRparamToken}:                  _CReduceAToAbstractDeclaratorAction,
	{_CState260, _CWildcardMarker}:              _CReduceEToJumpStatementAction,
	{_CState264, _CWildcardMarker}:              _CReduceCToAdditiveExpressionAction,
	{_CState265, _CWildcardMarker}:              _CReduceBToAdditiveExpressionAction,
	{_CState266, _CWildcardMarker}:              _CReduceBToAndExpressionAction,
	{_CState267, _CWildcardMarker}:              _CReduceDToCompoundStatementAction,
	{_CState268, _CWildcardMarker}:              _CReduceBToEqualityExpressionAction,
	{_CState269, _CWildcardMarker}:              _CReduceCToEqualityExpressionAction,
	{_CState270, _CWildcardMarker}:              _CReduceBToExclusiveOrExpressionAction,
	{_CState271, _CWildcardMarker}:              _CReduceBToExpressionAction,
	{_CState272, _CWildcardMarker}:              _CReduceBToInclusiveOrExpressionAction,
	{_CState273, _CWildcardMarker}:              _CReduceBToLogicalAndExpressionAction,
	{_CState274, _CWildcardMarker}:              _CReduceBToLogicalOrExpressionAction,
	{_CState276, _CWildcardMarker}:              _CReduceCToMultiplicativeExpressionAction,
	{_CState277, _CWildcardMarker}:              _CReduceDToMultiplicativeExpressionAction,
	{_CState278, _CWildcardMarker}:              _CReduceBToMultiplicativeExpressionAction,
	{_CState279, _CWildcardMarker}:              _CReduceEToPostfixExpressionAction,
	{_CState281, _CWildcardMarker}:              _CReduceCToPostfixExpressionAction,
	{_CState283, _CWildcardMarker}:              _CReduceAToArgumentExpressionListAction,
	{_CState284, _CWildcardMarker}:              _CReduceFToPostfixExpressionAction,
	{_CState285, _CWildcardMarker}:              _CReduceEToRelationalExpressionAction,
	{_CState286, _CWildcardMarker}:              _CReduceCToRelationalExpressionAction,
	{_CState287, _CWildcardMarker}:              _CReduceDToRelationalExpressionAction,
	{_CState288, _CWildcardMarker}:              _CReduceBToRelationalExpressionAction,
	{_CState289, _CWildcardMarker}:              _CReduceBToShiftExpressionAction,
	{_CState290, _CWildcardMarker}:              _CReduceCToShiftExpressionAction,
	{_CState291, _CWildcardMarker}:              _CReduceBToAssignmentExpressionAction,
	{_CState292, _CWildcardMarker}:              _CReduceBToDirectAbstractDeclaratorAction,
	{_CState294, _CWildcardMarker}:              _CReduceFToDirectAbstractDeclaratorAction,
	{_CState299, _CWildcardMarker}:              _CReduceCToAbstractDeclaratorAction,
	{_CState300, _CWildcardMarker}:              _CReduceBToIdentifierListAction,
	{_CState301, CRparamToken}:                  _CReduceBToParameterTypeListAction,
	{_CState302, _CWildcardMarker}:              _CReduceBToParameterListAction,
	{_CState303, _CWildcardMarker}:              _CReduceAToStructOrUnionSpecifierAction,
	{_CState304, _CWildcardMarker}:              _CReduceBToStructDeclaratorAction,
	{_CState307, _CWildcardMarker}:              _CReduceAToStructDeclarationAction,
	{_CState309, _CWildcardMarker}:              _CReduceBToInitializerAction,
	{_CState310, _CWildcardMarker}:              _CReduceBToLabeledStatementAction,
	{_CState314, _CWildcardMarker}:              _CReduceBToCastExpressionAction,
	{_CState315, _CWildcardMarker}:              _CReduceFToUnaryExpressionAction,
	{_CState319, _CWildcardMarker}:              _CReduceBToPostfixExpressionAction,
	{_CState321, _CWildcardMarker}:              _CReduceDToPostfixExpressionAction,
	{_CState322, _CWildcardMarker}:              _CReduceCToDirectAbstractDeclaratorAction,
	{_CState323, _CWildcardMarker}:              _CReduceAToDirectAbstractDeclaratorAction,
	{_CState324, _CWildcardMarker}:              _CReduceGToDirectAbstractDeclaratorAction,
	{_CState325, _CWildcardMarker}:              _CReduceDToDirectAbstractDeclaratorAction,
	{_CState327, _CWildcardMarker}:              _CReduceHToDirectAbstractDeclaratorAction,
	{_CState329, _CWildcardMarker}:              _CReduceCToStructDeclaratorAction,
	{_CState330, _CWildcardMarker}:              _CReduceBToStructDeclaratorListAction,
	{_CState331, _CWildcardMarker}:              _CReduceCToInitializerAction,
	{_CState332, _CWildcardMarker}:              _CReduceBToInitializerListAction,
	{_CState336, _CWildcardMarker}:              _CReduceAToSelectionStatementAction,
	{_CState337, CAndToken}:                     _CReduceAToSelectionStatementAction,
	{_CState337, CBreakToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CCaseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CConstantToken}:                _CReduceAToSelectionStatementAction,
	{_CState337, CContinueToken}:                _CReduceAToSelectionStatementAction,
	{_CState337, CDecOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CDefaultToken}:                 _CReduceAToSelectionStatementAction,
	{_CState337, CDoToken}:                      _CReduceAToSelectionStatementAction,
	{_CState337, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CExclaimToken}:                 _CReduceAToSelectionStatementAction,
	{_CState337, CForToken}:                     _CReduceAToSelectionStatementAction,
	{_CState337, CGotoToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CIdentifierToken}:              _CReduceAToSelectionStatementAction,
	{_CState337, CIfToken}:                      _CReduceAToSelectionStatementAction,
	{_CState337, CIncOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CLcurlToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CLparamToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CMinusToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CMulToken}:                     _CReduceAToSelectionStatementAction,
	{_CState337, CPlusToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CRcurlToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CReturnToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CSemicolonToken}:               _CReduceAToSelectionStatementAction,
	{_CState337, CSizeofToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CStringLiteralToken}:           _CReduceAToSelectionStatementAction,
	{_CState337, CSwitchToken}:                  _CReduceAToSelectionStatementAction,
	{_CState337, CTildaToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState338, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState338, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState339, _CWildcardMarker}:              _CReduceCToSelectionStatementAction,
	{_CState340, _CWildcardMarker}:              _CReduceAToIterationStatementAction,
	{_CState341, _CWildcardMarker}:              _CReduceBToConditionalExpressionAction,
	{_CState342, _CWildcardMarker}:              _CReduceBToArgumentExpressionListAction,
	{_CState343, _CWildcardMarker}:              _CReduceEToDirectAbstractDeclaratorAction,
	{_CState344, _CWildcardMarker}:              _CReduceIToDirectAbstractDeclaratorAction,
	{_CState346, _CWildcardMarker}:              _CReduceCToIterationStatementAction,
	{_CState349, _CWildcardMarker}:              _CReduceBToIterationStatementAction,
	{_CState350, _CWildcardMarker}:              _CReduceDToIterationStatementAction,
	{_CState351, _CWildcardMarker}:              _CReduceBToSelectionStatementAction,
}

var _CExpectedTerminals = map[_CStateId][]CSymbolId{
	_CState1:   []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken},
	_CState2:   []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, _CEndMarker},
	_CState7:   []CSymbolId{CIdentifierToken, CLcurlToken},
	_CState13:  []CSymbolId{CIdentifierToken, CLparamToken, CMulToken},
	_CState14:  []CSymbolId{CConstToken, CVolatileToken, CMulToken},
	_CState27:  []CSymbolId{CIdentifierToken, CLparamToken, CSemicolonToken, CMulToken},
	_CState28:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken},
	_CState29:  []CSymbolId{CLparamToken, CLbraceToken},
	_CState33:  []CSymbolId{CIdentifierToken, CLparamToken},
	_CState34:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState35:  []CSymbolId{CIdentifierToken, CLcurlToken},
	_CState37:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState38:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState39:  []CSymbolId{CLcurlToken},
	_CState40:  []CSymbolId{CIdentifierToken},
	_CState41:  []CSymbolId{CRparamToken},
	_CState44:  []CSymbolId{CConstToken, CVolatileToken, CMulToken},
	_CState46:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken, CEqToken},
	_CState48:  []CSymbolId{CSemicolonToken, CCommaToken},
	_CState49:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState52:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken},
	_CState53:  []CSymbolId{CIdentifierToken, CLparamToken, CSemicolonToken, CMulToken},
	_CState54:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRbraceToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState55:  []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRparamToken},
	_CState56:  []CSymbolId{CLparamToken, CLbraceToken},
	_CState58:  []CSymbolId{CLcurlToken},
	_CState59:  []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState63:  []CSymbolId{CIdentifierToken},
	_CState64:  []CSymbolId{CEqToken},
	_CState66:  []CSymbolId{CRcurlToken, CCommaToken},
	_CState70:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState72:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken},
	_CState73:  []CSymbolId{CIdentifierToken, CLparamToken, CMulToken},
	_CState76:  []CSymbolId{CSemicolonToken},
	_CState77:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState79:  []CSymbolId{CSemicolonToken},
	_CState80:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState81:  []CSymbolId{CColonToken},
	_CState82:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState84:  []CSymbolId{CLparamToken},
	_CState85:  []CSymbolId{CIdentifierToken},
	_CState86:  []CSymbolId{CColonToken},
	_CState87:  []CSymbolId{CLparamToken},
	_CState88:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState89:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState94:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState96:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState98:  []CSymbolId{CLparamToken},
	_CState100: []CSymbolId{CLparamToken},
	_CState101: []CSymbolId{CMinusToken, CPlusToken},
	_CState102: []CSymbolId{CAndToken},
	_CState107: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState108: []CSymbolId{CEqOpToken, CNeOpToken},
	_CState109: []CSymbolId{CHatToken},
	_CState110: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState112: []CSymbolId{COrToken},
	_CState116: []CSymbolId{CAndOpToken},
	_CState117: []CSymbolId{COrOpToken, CQuestionToken},
	_CState118: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState119: []CSymbolId{CPtrOpToken, CIncOpToken, CDecOpToken, CLparamToken, CLbraceToken, CDotToken},
	_CState121: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState123: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState125: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState126: []CSymbolId{CMulAssignToken, CDivAssignToken, CModAssignToken, CAddAssignToken, CSubAssignToken, CLeftAssignToken, CRightAssignToken, CAndAssignToken, CXorAssignToken, COrAssignToken, CEqToken},
	_CState127: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState130: []CSymbolId{CEqToken},
	_CState134: []CSymbolId{CRbraceToken},
	_CState138: []CSymbolId{CIdentifierToken, CLparamToken, CLbraceToken, CMulToken},
	_CState139: []CSymbolId{CRparamToken, CCommaToken},
	_CState141: []CSymbolId{CCommaToken, CRparamToken},
	_CState142: []CSymbolId{CRparamToken},
	_CState143: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState144: []CSymbolId{CIdentifierToken, CLparamToken, CColonToken, CMulToken},
	_CState146: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRcurlToken},
	_CState147: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState148: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState149: []CSymbolId{CRcurlToken, CCommaToken},
	_CState150: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState151: []CSymbolId{CIdentifierToken},
	_CState153: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState159: []CSymbolId{CColonToken},
	_CState161: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState163: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState164: []CSymbolId{CWhileToken},
	_CState165: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState166: []CSymbolId{CSemicolonToken},
	_CState167: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState168: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState170: []CSymbolId{CRparamToken, CCommaToken},
	_CState171: []CSymbolId{CLparamToken, CLbraceToken, CMulToken, CRparamToken},
	_CState172: []CSymbolId{CRparamToken},
	_CState174: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState175: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState177: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState178: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState179: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState180: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState181: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState183: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState184: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState185: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState186: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState187: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState189: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState190: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState191: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState192: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState193: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState194: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState195: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState197: []CSymbolId{CIdentifierToken},
	_CState199: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState200: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState201: []CSymbolId{CIdentifierToken},
	_CState202: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState203: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState204: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState205: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState206: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState207: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState221: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState224: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRbraceToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState225: []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CRparamToken, CLbraceToken, CMulToken},
	_CState228: []CSymbolId{CLparamToken, CLbraceToken},
	_CState229: []CSymbolId{CIdentifierToken, CLparamToken, CLbraceToken},
	_CState230: []CSymbolId{CIdentifierToken},
	_CState232: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CEllipsisToken},
	_CState234: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRcurlToken},
	_CState235: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState236: []CSymbolId{CColonToken},
	_CState238: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState247: []CSymbolId{CRcurlToken, CCommaToken},
	_CState248: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState250: []CSymbolId{CLparamToken},
	_CState251: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState254: []CSymbolId{CRparamToken, CCommaToken},
	_CState256: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CRparamToken, CLbraceToken, CMulToken},
	_CState257: []CSymbolId{CRparamToken},
	_CState258: []CSymbolId{CLparamToken, CLbraceToken, CRparamToken},
	_CState259: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState261: []CSymbolId{CRparamToken},
	_CState262: []CSymbolId{CRparamToken, CCommaToken},
	_CState263: []CSymbolId{CRparamToken, CCommaToken},
	_CState264: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState265: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState266: []CSymbolId{CEqOpToken, CNeOpToken},
	_CState268: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState269: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState270: []CSymbolId{CAndToken},
	_CState272: []CSymbolId{CHatToken},
	_CState273: []CSymbolId{COrToken},
	_CState274: []CSymbolId{CAndOpToken},
	_CState275: []CSymbolId{CColonToken, CCommaToken},
	_CState280: []CSymbolId{CRbraceToken, CCommaToken},
	_CState282: []CSymbolId{CRparamToken, CCommaToken},
	_CState285: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState286: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState287: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState288: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState289: []CSymbolId{CMinusToken, CPlusToken},
	_CState290: []CSymbolId{CMinusToken, CPlusToken},
	_CState293: []CSymbolId{CRbraceToken},
	_CState295: []CSymbolId{CRparamToken},
	_CState296: []CSymbolId{CRparamToken},
	_CState297: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRbraceToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState298: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRparamToken},
	_CState299: []CSymbolId{CLparamToken, CLbraceToken},
	_CState301: []CSymbolId{CRparamToken},
	_CState305: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState306: []CSymbolId{CIdentifierToken, CLparamToken, CColonToken, CMulToken},
	_CState308: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CRcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState311: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState312: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState313: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState316: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState317: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState318: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState320: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState326: []CSymbolId{CRbraceToken},
	_CState328: []CSymbolId{CRparamToken},
	_CState333: []CSymbolId{CRparamToken, CCommaToken},
	_CState334: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState335: []CSymbolId{CRparamToken, CCommaToken},
	_CState336: []CSymbolId{CElseToken},
	_CState337: []CSymbolId{CElseToken, CAndToken, CBreakToken, CCaseToken, CConstantToken, CContinueToken, CDecOpToken, CDefaultToken, CDoToken, CElseToken, CExclaimToken, CForToken, CGotoToken, CIdentifierToken, CIfToken, CIncOpToken, CLcurlToken, CLparamToken, CMinusToken, CMulToken, CPlusToken, CRcurlToken, CReturnToken, CSemicolonToken, CSizeofToken, CStringLiteralToken, CSwitchToken, CTildaToken, CWhileToken},
	_CState338: []CSymbolId{CElseToken, CElseToken, CWhileToken},
	_CState345: []CSymbolId{CSemicolonToken},
	_CState347: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState348: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.translation_unit
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 13
      MUL -> State 14
      declaration -> State 26
      declaration_specifiers -> State 27
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 28
      direct_declarator -> State 29
      pointer -> State 33
      translation_unit -> State 2
      external_declaration -> State 31
      function_definition -> State 32

  State 2:
    Kernel Items:
      #accept: ^ translation_unit., $
      translation_unit: translation_unit.external_declaration
    Reduce:
      $ -> [#accept]
    Goto:
      IDENTIFIER -> State 10
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 13
      MUL -> State 14
      declaration -> State 26
      declaration_specifiers -> State 27
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 28
      direct_declarator -> State 29
      pointer -> State 33
      external_declaration -> State 60
      function_definition -> State 32

  State 3:
    Kernel Items:
      storage_class_specifier: AUTO., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 4:
    Kernel Items:
      type_specifier: CHAR., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      type_qualifier: CONST., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      type_specifier: DOUBLE., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      enum_specifier: ENUM.IDENTIFIER
      enum_specifier: ENUM.IDENTIFIER LCURL enumerator_list RCURL
      enum_specifier: ENUM.LCURL enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 39
      LCURL -> State 40

  State 8:
    Kernel Items:
      storage_class_specifier: EXTERN., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      type_specifier: FLOAT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      direct_declarator: IDENTIFIER., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      type_specifier: INT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      type_specifier: LONG., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      direct_declarator: LPARAM.declarator RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      MUL -> State 14
      declarator -> State 41
      direct_declarator -> State 29
      pointer -> State 33

  State 14:
    Kernel Items:
      pointer: MUL., *
      pointer: MUL.pointer
      pointer: MUL.type_qualifier_list
      pointer: MUL.type_qualifier_list pointer
    Reduce:
      * -> [pointer]
    Goto:
      CONST -> State 5
      VOLATILE -> State 25
      MUL -> State 14
      type_qualifier -> State 43
      pointer -> State 42
      type_qualifier_list -> State 44

  State 15:
    Kernel Items:
      storage_class_specifier: REGISTER., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      type_specifier: SHORT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      type_specifier: SIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      storage_class_specifier: STATIC., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      struct_or_union: STRUCT., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      storage_class_specifier: TYPEDEF., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      type_specifier: TYPE_NAME., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      struct_or_union: UNION., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      type_specifier: UNSIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      type_specifier: VOID., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      type_qualifier: VOLATILE., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      external_declaration: declaration., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      declaration: declaration_specifiers.SEMICOLON
      declaration: declaration_specifiers.init_declarator_list SEMICOLON
      function_definition: declaration_specifiers.declarator compound_statement
      function_definition: declaration_specifiers.declarator declaration_list compound_statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      SEMICOLON -> State 45
      MUL -> State 14
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 46
      direct_declarator -> State 29
      pointer -> State 33

  State 28:
    Kernel Items:
      function_definition: declarator.compound_statement
      function_definition: declarator.declaration_list compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LCURL -> State 49
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 50
      declaration_list -> State 52

  State 29:
    Kernel Items:
      declarator: direct_declarator., *
      direct_declarator: direct_declarator.LBRACE RBRACE
      direct_declarator: direct_declarator.LBRACE constant_expression RBRACE
      direct_declarator: direct_declarator.LPARAM RPARAM
      direct_declarator: direct_declarator.LPARAM identifier_list RPARAM
      direct_declarator: direct_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [declarator]
    Goto:
      LPARAM -> State 55
      LBRACE -> State 54

  State 30:
    Kernel Items:
      type_specifier: enum_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      translation_unit: external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      external_declaration: function_definition., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      declarator: pointer.direct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      direct_declarator -> State 56

  State 34:
    Kernel Items:
      declaration_specifiers: storage_class_specifier., *
      declaration_specifiers: storage_class_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      declaration_specifiers -> State 57
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 35:
    Kernel Items:
      struct_or_union_specifier: struct_or_union.IDENTIFIER
      struct_or_union_specifier: struct_or_union.IDENTIFIER LCURL struct_declaration_list RCURL
      struct_or_union_specifier: struct_or_union.LCURL struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 58
      LCURL -> State 59

  State 36:
    Kernel Items:
      type_specifier: struct_or_union_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      declaration_specifiers: type_qualifier., *
      declaration_specifiers: type_qualifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      declaration_specifiers -> State 61
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 38:
    Kernel Items:
      declaration_specifiers: type_specifier., *
      declaration_specifiers: type_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      declaration_specifiers -> State 62
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37

  State 39:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER., *
      enum_specifier: ENUM IDENTIFIER.LCURL enumerator_list RCURL
    Reduce:
      * -> [enum_specifier]
    Goto:
      LCURL -> State 63

  State 40:
    Kernel Items:
      enum_specifier: ENUM LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 64
      enumerator_list -> State 66
      enumerator -> State 65

  State 41:
    Kernel Items:
      direct_declarator: LPARAM declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 67

  State 42:
    Kernel Items:
      pointer: MUL pointer., *
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      type_qualifier_list: type_qualifier., *
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      pointer: MUL type_qualifier_list., *
      pointer: MUL type_qualifier_list.pointer
      type_qualifier_list: type_qualifier_list.type_qualifier
    Reduce:
      * -> [pointer]
    Goto:
      CONST -> State 5
      VOLATILE -> State 25
      MUL -> State 14
      type_qualifier -> State 69
      pointer -> State 68

  State 45:
    Kernel Items:
      declaration: declaration_specifiers SEMICOLON., *
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      function_definition: declaration_specifiers declarator.compound_statement
      function_definition: declaration_specifiers declarator.declaration_list compound_statement
      init_declarator: declarator., *
      init_declarator: declarator.EQ initializer
    Reduce:
      * -> [init_declarator]
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LCURL -> State 49
      EQ -> State 70
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 71
      declaration_list -> State 72

  State 47:
    Kernel Items:
      init_declarator_list: init_declarator., *
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      declaration: declaration_specifiers init_declarator_list.SEMICOLON
      init_declarator_list: init_declarator_list.COMMA init_declarator
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 74
      COMMA -> State 73

  State 49:
    Kernel Items:
      compound_statement: LCURL.RCURL
      compound_statement: LCURL.declaration_list RCURL
      compound_statement: LCURL.declaration_list statement_list RCURL
      compound_statement: LCURL.statement_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      RCURL -> State 93
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      statement -> State 124
      labeled_statement -> State 115
      compound_statement -> State 105
      declaration_list -> State 107
      statement_list -> State 125
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 50:
    Kernel Items:
      function_definition: declarator compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      declaration_list: declaration., *
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      declaration_list: declaration_list.declaration
      function_definition: declarator declaration_list.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LCURL -> State 49
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 128

  State 53:
    Kernel Items:
      declaration: declaration_specifiers.SEMICOLON
      declaration: declaration_specifiers.init_declarator_list SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      SEMICOLON -> State 45
      MUL -> State 14
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 130
      direct_declarator -> State 29
      pointer -> State 33

  State 54:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE.RBRACE
      direct_declarator: direct_declarator LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      RBRACE -> State 132
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 134

  State 55:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM.RPARAM
      direct_declarator: direct_declarator LPARAM.identifier_list RPARAM
      direct_declarator: direct_declarator LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 136
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      RPARAM -> State 137
      declaration_specifiers -> State 138
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_type_list -> State 142
      parameter_list -> State 141
      parameter_declaration -> State 140
      identifier_list -> State 139

  State 56:
    Kernel Items:
      declarator: pointer direct_declarator., *
      direct_declarator: direct_declarator.LBRACE RBRACE
      direct_declarator: direct_declarator.LBRACE constant_expression RBRACE
      direct_declarator: direct_declarator.LPARAM RPARAM
      direct_declarator: direct_declarator.LPARAM identifier_list RPARAM
      direct_declarator: direct_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [declarator]
    Goto:
      LPARAM -> State 55
      LBRACE -> State 54

  State 57:
    Kernel Items:
      declaration_specifiers: storage_class_specifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER., *
      struct_or_union_specifier: struct_or_union IDENTIFIER.LCURL struct_declaration_list RCURL
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      LCURL -> State 143

  State 59:
    Kernel Items:
      struct_or_union_specifier: struct_or_union LCURL.struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration_list -> State 146
      struct_declaration -> State 145
      specifier_qualifier_list -> State 144
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 60:
    Kernel Items:
      translation_unit: translation_unit external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      declaration_specifiers: type_qualifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      declaration_specifiers: type_specifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 64
      enumerator_list -> State 149
      enumerator -> State 65

  State 64:
    Kernel Items:
      enumerator: IDENTIFIER., *
      enumerator: IDENTIFIER.EQ constant_expression
    Reduce:
      * -> [enumerator]
    Goto:
      EQ -> State 150

  State 65:
    Kernel Items:
      enumerator_list: enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      enum_specifier: ENUM LCURL enumerator_list.RCURL
      enumerator_list: enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 152
      COMMA -> State 151

  State 67:
    Kernel Items:
      direct_declarator: LPARAM declarator RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      pointer: MUL type_qualifier_list pointer., *
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      type_qualifier_list: type_qualifier_list type_qualifier., *
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      init_declarator: declarator EQ.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      LCURL -> State 153
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 155

  State 71:
    Kernel Items:
      function_definition: declaration_specifiers declarator compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      declaration_list: declaration_list.declaration
      function_definition: declaration_specifiers declarator declaration_list.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LCURL -> State 49
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      compound_statement -> State 156

  State 73:
    Kernel Items:
      init_declarator_list: init_declarator_list COMMA.init_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      MUL -> State 14
      init_declarator -> State 157
      declarator -> State 130
      direct_declarator -> State 29
      pointer -> State 33

  State 74:
    Kernel Items:
      declaration: declaration_specifiers init_declarator_list SEMICOLON., *
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      unary_operator: AND., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      jump_statement: BREAK.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 158

  State 77:
    Kernel Items:
      labeled_statement: CASE.constant_expression COLON statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 159

  State 78:
    Kernel Items:
      primary_expression: CONSTANT., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      jump_statement: CONTINUE.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 160

  State 80:
    Kernel Items:
      unary_expression: DEC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 161
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 162
      unary_operator -> State 127

  State 81:
    Kernel Items:
      labeled_statement: DEFAULT.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 163

  State 82:
    Kernel Items:
      iteration_statement: DO.statement WHILE LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 164
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 83:
    Kernel Items:
      unary_operator: EXCLAIM., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      iteration_statement: FOR.LPARAM expression_statement expression_statement RPARAM statement
      iteration_statement: FOR.LPARAM expression_statement expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 165

  State 85:
    Kernel Items:
      jump_statement: GOTO.IDENTIFIER SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166

  State 86:
    Kernel Items:
      labeled_statement: IDENTIFIER.COLON statement
      primary_expression: IDENTIFIER., *
    Reduce:
      * -> [primary_expression]
    Goto:
      COLON -> State 167

  State 87:
    Kernel Items:
      selection_statement: IF.LPARAM expression RPARAM statement
      selection_statement: IF.LPARAM expression RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 168

  State 88:
    Kernel Items:
      unary_expression: INC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 161
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 169
      unary_operator -> State 127

  State 89:
    Kernel Items:
      cast_expression: LPARAM.type_name RPARAM cast_expression
      primary_expression: LPARAM.expression RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 170
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 171
      enum_specifier -> State 30
      type_qualifier -> State 147
      type_name -> State 172

  State 90:
    Kernel Items:
      unary_operator: MINUS., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      unary_operator: MUL., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      unary_operator: PLUS., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      compound_statement: LCURL RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      jump_statement: RETURN.SEMICOLON
      jump_statement: RETURN.expression SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      SEMICOLON -> State 173
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 174

  State 95:
    Kernel Items:
      expression_statement: SEMICOLON., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      unary_expression: SIZEOF.LPARAM type_name RPARAM
      unary_expression: SIZEOF.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 175
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 176
      unary_operator -> State 127

  State 97:
    Kernel Items:
      primary_expression: STRING_LITERAL., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      selection_statement: SWITCH.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 177

  State 99:
    Kernel Items:
      unary_operator: TILDA., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      iteration_statement: WHILE.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 178

  State 101:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 179
      PLUS -> State 180

  State 102:
    Kernel Items:
      and_expression: and_expression.AND equality_expression
      exclusive_or_expression: and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 181

  State 103:
    Kernel Items:
      expression: assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      multiplicative_expression: cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      statement: compound_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      assignment_expression: conditional_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      compound_statement: LCURL declaration_list.RCURL
      compound_statement: LCURL declaration_list.statement_list RCURL
      declaration_list: declaration_list.declaration
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      RCURL -> State 182
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      declaration -> State 129
      declaration_specifiers -> State 53
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      statement -> State 124
      labeled_statement -> State 115
      compound_statement -> State 105
      statement_list -> State 183
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 108:
    Kernel Items:
      and_expression: equality_expression., *
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 184
      NE_OP -> State 185

  State 109:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.HAT and_expression
      inclusive_or_expression: exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 186

  State 110:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      expression_statement: expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 188
      COMMA -> State 187

  State 111:
    Kernel Items:
      statement: expression_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression: inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 189

  State 113:
    Kernel Items:
      statement: iteration_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      statement: jump_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      statement: labeled_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 190

  State 117:
    Kernel Items:
      conditional_expression: logical_or_expression., *
      conditional_expression: logical_or_expression.QUESTION expression COLON conditional_expression
      logical_or_expression: logical_or_expression.OR_OP logical_and_expression
    Reduce:
      * -> [conditional_expression]
    Goto:
      OR_OP -> State 191
      QUESTION -> State 192

  State 118:
    Kernel Items:
      additive_expression: multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 195
      DIV -> State 193
      MOD -> State 194

  State 119:
    Kernel Items:
      postfix_expression: postfix_expression.DEC_OP
      postfix_expression: postfix_expression.DOT IDENTIFIER
      postfix_expression: postfix_expression.INC_OP
      postfix_expression: postfix_expression.LBRACE expression RBRACE
      postfix_expression: postfix_expression.LPARAM RPARAM
      postfix_expression: postfix_expression.LPARAM argument_expression_list RPARAM
      postfix_expression: postfix_expression.PTR_OP IDENTIFIER
      unary_expression: postfix_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      PTR_OP -> State 201
      INC_OP -> State 198
      DEC_OP -> State 196
      LPARAM -> State 200
      LBRACE -> State 199
      DOT -> State 197

  State 120:
    Kernel Items:
      postfix_expression: primary_expression., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      equality_expression: relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 204
      GE_OP -> State 202
      LT -> State 205
      GT -> State 203

  State 122:
    Kernel Items:
      statement: selection_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      relational_expression: shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 124:
    Kernel Items:
      statement_list: statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      compound_statement: LCURL statement_list.RCURL
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      RCURL -> State 208
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 209
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 126:
    Kernel Items:
      assignment_expression: unary_expression.assignment_operator assignment_expression
      cast_expression: unary_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      MUL_ASSIGN -> State 216
      DIV_ASSIGN -> State 212
      MOD_ASSIGN -> State 215
      ADD_ASSIGN -> State 210
      SUB_ASSIGN -> State 219
      LEFT_ASSIGN -> State 214
      RIGHT_ASSIGN -> State 218
      AND_ASSIGN -> State 211
      XOR_ASSIGN -> State 220
      OR_ASSIGN -> State 217
      EQ -> State 213
      assignment_operator -> State 221

  State 127:
    Kernel Items:
      unary_expression: unary_operator.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 222

  State 128:
    Kernel Items:
      function_definition: declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      declaration_list: declaration_list declaration., *
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      init_declarator: declarator., *
      init_declarator: declarator.EQ initializer
    Reduce:
      * -> [init_declarator]
    Goto:
      EQ -> State 70

  State 131:
    Kernel Items:
      primary_expression: IDENTIFIER., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE RBRACE., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      constant_expression: conditional_expression., *
    Reduce:
      * -> [constant_expression]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 223

  State 135:
    Kernel Items:
      cast_expression: unary_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      identifier_list: IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      parameter_declaration: declaration_specifiers., *
      parameter_declaration: declaration_specifiers.abstract_declarator
      parameter_declaration: declaration_specifiers.declarator
    Reduce:
      * -> [parameter_declaration]
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 225
      LBRACE -> State 224
      MUL -> State 14
      declarator -> State 227
      direct_declarator -> State 29
      pointer -> State 229
      abstract_declarator -> State 226
      direct_abstract_declarator -> State 228

  State 139:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM identifier_list.RPARAM
      identifier_list: identifier_list.COMMA IDENTIFIER
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 231
      COMMA -> State 230

  State 140:
    Kernel Items:
      parameter_list: parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      parameter_list: parameter_list.COMMA parameter_declaration
      parameter_type_list: parameter_list., RPARAM
      parameter_type_list: parameter_list.COMMA ELLIPSIS
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      COMMA -> State 232

  State 142:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 233

  State 143:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL.struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration_list -> State 234
      struct_declaration -> State 145
      specifier_qualifier_list -> State 144
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 144:
    Kernel Items:
      struct_declaration: specifier_qualifier_list.struct_declarator_list SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      COLON -> State 235
      MUL -> State 14
      struct_declarator_list -> State 238
      struct_declarator -> State 237
      declarator -> State 236
      direct_declarator -> State 29
      pointer -> State 33

  State 145:
    Kernel Items:
      struct_declaration_list: struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      struct_declaration_list: struct_declaration_list.struct_declaration
      struct_or_union_specifier: struct_or_union LCURL struct_declaration_list.RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      RCURL -> State 239
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration -> State 240
      specifier_qualifier_list -> State 144
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 147:
    Kernel Items:
      specifier_qualifier_list: type_qualifier., *
      specifier_qualifier_list: type_qualifier.specifier_qualifier_list
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 241
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 148:
    Kernel Items:
      specifier_qualifier_list: type_specifier., *
      specifier_qualifier_list: type_specifier.specifier_qualifier_list
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 242
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 149:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL enumerator_list.RCURL
      enumerator_list: enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 243
      COMMA -> State 151

  State 150:
    Kernel Items:
      enumerator: IDENTIFIER EQ.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 244

  State 151:
    Kernel Items:
      enumerator_list: enumerator_list COMMA.enumerator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 64
      enumerator -> State 245

  State 152:
    Kernel Items:
      enum_specifier: ENUM LCURL enumerator_list RCURL., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      initializer: LCURL.initializer_list COMMA RCURL
      initializer: LCURL.initializer_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      LCURL -> State 153
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 246
      initializer_list -> State 247

  State 154:
    Kernel Items:
      initializer: assignment_expression., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      init_declarator: declarator EQ initializer., *
    Reduce:
      * -> [init_declarator]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      function_definition: declaration_specifiers declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      init_declarator_list: init_declarator_list COMMA init_declarator., *
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      jump_statement: BREAK SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      labeled_statement: CASE constant_expression.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 248

  State 160:
    Kernel Items:
      jump_statement: CONTINUE SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      primary_expression: LPARAM.expression RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 170

  State 162:
    Kernel Items:
      unary_expression: DEC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      labeled_statement: DEFAULT COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 249
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 164:
    Kernel Items:
      iteration_statement: DO statement.WHILE LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      WHILE -> State 250

  State 165:
    Kernel Items:
      iteration_statement: FOR LPARAM.expression_statement expression_statement RPARAM statement
      iteration_statement: FOR LPARAM.expression_statement expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      expression_statement -> State 251

  State 166:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 252

  State 167:
    Kernel Items:
      labeled_statement: IDENTIFIER COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 253
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 168:
    Kernel Items:
      selection_statement: IF LPARAM.expression RPARAM statement
      selection_statement: IF LPARAM.expression RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 254

  State 169:
    Kernel Items:
      unary_expression: INC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      primary_expression: LPARAM expression.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 255
      COMMA -> State 187

  State 171:
    Kernel Items:
      type_name: specifier_qualifier_list., RPARAM
      type_name: specifier_qualifier_list.abstract_declarator
    Reduce:
      RPARAM -> [type_name]
    Goto:
      LPARAM -> State 256
      LBRACE -> State 224
      MUL -> State 14
      pointer -> State 258
      abstract_declarator -> State 257
      direct_abstract_declarator -> State 228

  State 172:
    Kernel Items:
      cast_expression: LPARAM type_name.RPARAM cast_expression
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 259

  State 173:
    Kernel Items:
      jump_statement: RETURN SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      jump_statement: RETURN expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 260
      COMMA -> State 187

  State 175:
    Kernel Items:
      primary_expression: LPARAM.expression RPARAM
      unary_expression: SIZEOF LPARAM.type_name RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 170
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      specifier_qualifier_list -> State 171
      enum_specifier -> State 30
      type_qualifier -> State 147
      type_name -> State 261

  State 176:
    Kernel Items:
      unary_expression: SIZEOF unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      selection_statement: SWITCH LPARAM.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 262

  State 178:
    Kernel Items:
      iteration_statement: WHILE LPARAM.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 263

  State 179:
    Kernel Items:
      additive_expression: additive_expression MINUS.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 264

  State 180:
    Kernel Items:
      additive_expression: additive_expression PLUS.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 265

  State 181:
    Kernel Items:
      and_expression: and_expression AND.equality_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 266

  State 182:
    Kernel Items:
      compound_statement: LCURL declaration_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      compound_statement: LCURL declaration_list statement_list.RCURL
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      RCURL -> State 267
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 209
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 184:
    Kernel Items:
      equality_expression: equality_expression EQ_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 268

  State 185:
    Kernel Items:
      equality_expression: equality_expression NE_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 269

  State 186:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression HAT.and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 270

  State 187:
    Kernel Items:
      expression: expression COMMA.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 271

  State 188:
    Kernel Items:
      expression_statement: expression SEMICOLON., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression OR.exclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 272

  State 190:
    Kernel Items:
      logical_and_expression: logical_and_expression AND_OP.inclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 273

  State 191:
    Kernel Items:
      logical_or_expression: logical_or_expression OR_OP.logical_and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 274

  State 192:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION.expression COLON conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 275

  State 193:
    Kernel Items:
      multiplicative_expression: multiplicative_expression DIV.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 276

  State 194:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MOD.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 277

  State 195:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MUL.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 278

  State 196:
    Kernel Items:
      postfix_expression: postfix_expression DEC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      postfix_expression: postfix_expression DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 279

  State 198:
    Kernel Items:
      postfix_expression: postfix_expression INC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      postfix_expression: postfix_expression LBRACE.expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 280

  State 200:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM.RPARAM
      postfix_expression: postfix_expression LPARAM.argument_expression_list RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      RPARAM -> State 281
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      argument_expression_list -> State 282
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 283

  State 201:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 284

  State 202:
    Kernel Items:
      relational_expression: relational_expression GE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 285

  State 203:
    Kernel Items:
      relational_expression: relational_expression GT.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 286

  State 204:
    Kernel Items:
      relational_expression: relational_expression LE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 287

  State 205:
    Kernel Items:
      relational_expression: relational_expression LT.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 288

  State 206:
    Kernel Items:
      shift_expression: shift_expression LEFT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 289

  State 207:
    Kernel Items:
      shift_expression: shift_expression RIGHT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 290

  State 208:
    Kernel Items:
      compound_statement: LCURL statement_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      statement_list: statement_list statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      assignment_operator: ADD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      assignment_operator: AND_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      assignment_operator: DIV_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      assignment_operator: EQ., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      assignment_operator: LEFT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      assignment_operator: MOD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      assignment_operator: MUL_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      assignment_operator: OR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      assignment_operator: RIGHT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      assignment_operator: SUB_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      assignment_operator: XOR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 291

  State 222:
    Kernel Items:
      unary_expression: unary_operator cast_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      direct_abstract_declarator: LBRACE.RBRACE
      direct_abstract_declarator: LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      RBRACE -> State 292
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 293

  State 225:
    Kernel Items:
      direct_abstract_declarator: LPARAM.RPARAM
      direct_abstract_declarator: LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator: LPARAM.parameter_type_list RPARAM
      direct_declarator: LPARAM.declarator RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 225
      RPARAM -> State 294
      LBRACE -> State 224
      MUL -> State 14
      declaration_specifiers -> State 138
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      declarator -> State 41
      direct_declarator -> State 29
      pointer -> State 229
      parameter_type_list -> State 296
      parameter_list -> State 141
      parameter_declaration -> State 140
      abstract_declarator -> State 295
      direct_abstract_declarator -> State 228

  State 226:
    Kernel Items:
      parameter_declaration: declaration_specifiers abstract_declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      parameter_declaration: declaration_specifiers declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      abstract_declarator: direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator: direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 298
      LBRACE -> State 297

  State 229:
    Kernel Items:
      abstract_declarator: pointer., *
      abstract_declarator: pointer.direct_abstract_declarator
      declarator: pointer.direct_declarator
    Reduce:
      * -> [abstract_declarator]
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 225
      LBRACE -> State 224
      direct_declarator -> State 56
      direct_abstract_declarator -> State 299

  State 230:
    Kernel Items:
      identifier_list: identifier_list COMMA.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 300

  State 231:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM identifier_list RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      parameter_list: parameter_list COMMA.parameter_declaration
      parameter_type_list: parameter_list COMMA.ELLIPSIS
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      ELLIPSIS -> State 301
      declaration_specifiers -> State 138
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_declaration -> State 302

  State 233:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      struct_declaration_list: struct_declaration_list.struct_declaration
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL struct_declaration_list.RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      RCURL -> State 303
      type_specifier -> State 148
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      struct_declaration -> State 240
      specifier_qualifier_list -> State 144
      enum_specifier -> State 30
      type_qualifier -> State 147

  State 235:
    Kernel Items:
      struct_declarator: COLON.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 304

  State 236:
    Kernel Items:
      struct_declarator: declarator., *
      struct_declarator: declarator.COLON constant_expression
    Reduce:
      * -> [struct_declarator]
    Goto:
      COLON -> State 305

  State 237:
    Kernel Items:
      struct_declarator_list: struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list.SEMICOLON
      struct_declarator_list: struct_declarator_list.COMMA struct_declarator
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 307
      COMMA -> State 306

  State 239:
    Kernel Items:
      struct_or_union_specifier: struct_or_union LCURL struct_declaration_list RCURL., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      struct_declaration_list: struct_declaration_list struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      specifier_qualifier_list: type_qualifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      specifier_qualifier_list: type_specifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL enumerator_list RCURL., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      enumerator: IDENTIFIER EQ constant_expression., *
    Reduce:
      * -> [enumerator]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      enumerator_list: enumerator_list COMMA enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      initializer_list: initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      initializer: LCURL initializer_list.COMMA RCURL
      initializer: LCURL initializer_list.RCURL
      initializer_list: initializer_list.COMMA initializer
    Reduce:
      (nil)
    Goto:
      RCURL -> State 309
      COMMA -> State 308

  State 248:
    Kernel Items:
      labeled_statement: CASE constant_expression COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 310
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 249:
    Kernel Items:
      labeled_statement: DEFAULT COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      iteration_statement: DO statement WHILE.LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 311

  State 251:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement.expression_statement RPARAM statement
      iteration_statement: FOR LPARAM expression_statement.expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      expression_statement -> State 312

  State 252:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      labeled_statement: IDENTIFIER COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      selection_statement: IF LPARAM expression.RPARAM statement
      selection_statement: IF LPARAM expression.RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 313
      COMMA -> State 187

  State 255:
    Kernel Items:
      primary_expression: LPARAM expression RPARAM., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      direct_abstract_declarator: LPARAM.RPARAM
      direct_abstract_declarator: LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator: LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      LPARAM -> State 256
      RPARAM -> State 294
      LBRACE -> State 224
      MUL -> State 14
      declaration_specifiers -> State 138
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      pointer -> State 258
      parameter_type_list -> State 296
      parameter_list -> State 141
      parameter_declaration -> State 140
      abstract_declarator -> State 295
      direct_abstract_declarator -> State 228

  State 257:
    Kernel Items:
      type_name: specifier_qualifier_list abstract_declarator., RPARAM
    Reduce:
      RPARAM -> [type_name]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      abstract_declarator: pointer., RPARAM
      abstract_declarator: pointer.direct_abstract_declarator
    Reduce:
      RPARAM -> [abstract_declarator]
    Goto:
      LPARAM -> State 256
      LBRACE -> State 224
      direct_abstract_declarator -> State 299

  State 259:
    Kernel Items:
      cast_expression: LPARAM type_name RPARAM.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 314

  State 260:
    Kernel Items:
      jump_statement: RETURN expression SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      unary_expression: SIZEOF LPARAM type_name.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 315

  State 262:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      selection_statement: SWITCH LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 316
      COMMA -> State 187

  State 263:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: WHILE LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 317
      COMMA -> State 187

  State 264:
    Kernel Items:
      additive_expression: additive_expression MINUS multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 195
      DIV -> State 193
      MOD -> State 194

  State 265:
    Kernel Items:
      additive_expression: additive_expression PLUS multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 195
      DIV -> State 193
      MOD -> State 194

  State 266:
    Kernel Items:
      and_expression: and_expression AND equality_expression., *
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 184
      NE_OP -> State 185

  State 267:
    Kernel Items:
      compound_statement: LCURL declaration_list statement_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      equality_expression: equality_expression EQ_OP relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 204
      GE_OP -> State 202
      LT -> State 205
      GT -> State 203

  State 269:
    Kernel Items:
      equality_expression: equality_expression NE_OP relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 204
      GE_OP -> State 202
      LT -> State 205
      GT -> State 203

  State 270:
    Kernel Items:
      and_expression: and_expression.AND equality_expression
      exclusive_or_expression: exclusive_or_expression HAT and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 181

  State 271:
    Kernel Items:
      expression: expression COMMA assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.HAT and_expression
      inclusive_or_expression: inclusive_or_expression OR exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 186

  State 273:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression: logical_and_expression AND_OP inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 189

  State 274:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_or_expression OR_OP logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 190

  State 275:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression.COLON conditional_expression
      expression: expression.COMMA assignment_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 318
      COMMA -> State 187

  State 276:
    Kernel Items:
      multiplicative_expression: multiplicative_expression DIV cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MOD cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MUL cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      postfix_expression: postfix_expression DOT IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      postfix_expression: postfix_expression LBRACE expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 319
      COMMA -> State 187

  State 281:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM RPARAM., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      argument_expression_list: argument_expression_list.COMMA assignment_expression
      postfix_expression: postfix_expression LPARAM argument_expression_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 321
      COMMA -> State 320

  State 283:
    Kernel Items:
      argument_expression_list: assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      relational_expression: relational_expression GE_OP shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 286:
    Kernel Items:
      relational_expression: relational_expression GT shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 287:
    Kernel Items:
      relational_expression: relational_expression LE_OP shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 288:
    Kernel Items:
      relational_expression: relational_expression LT shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 206
      RIGHT_OP -> State 207

  State 289:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: shift_expression LEFT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 179
      PLUS -> State 180

  State 290:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: shift_expression RIGHT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 179
      PLUS -> State 180

  State 291:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator assignment_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      direct_abstract_declarator: LBRACE RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      direct_abstract_declarator: LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 322

  State 294:
    Kernel Items:
      direct_abstract_declarator: LPARAM RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      direct_abstract_declarator: LPARAM abstract_declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 323

  State 296:
    Kernel Items:
      direct_abstract_declarator: LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 324

  State 297:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE.RBRACE
      direct_abstract_declarator: direct_abstract_declarator LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      RBRACE -> State 325
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 326

  State 298:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM.RPARAM
      direct_abstract_declarator: direct_abstract_declarator LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 21
      TYPEDEF -> State 20
      EXTERN -> State 8
      STATIC -> State 18
      AUTO -> State 3
      REGISTER -> State 15
      CHAR -> State 4
      SHORT -> State 16
      INT -> State 11
      LONG -> State 12
      SIGNED -> State 17
      UNSIGNED -> State 23
      FLOAT -> State 9
      DOUBLE -> State 6
      CONST -> State 5
      VOLATILE -> State 25
      VOID -> State 24
      STRUCT -> State 19
      UNION -> State 22
      ENUM -> State 7
      RPARAM -> State 327
      declaration_specifiers -> State 138
      storage_class_specifier -> State 34
      type_specifier -> State 38
      struct_or_union_specifier -> State 36
      struct_or_union -> State 35
      enum_specifier -> State 30
      type_qualifier -> State 37
      parameter_type_list -> State 328
      parameter_list -> State 141
      parameter_declaration -> State 140

  State 299:
    Kernel Items:
      abstract_declarator: pointer direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator: direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 298
      LBRACE -> State 297

  State 300:
    Kernel Items:
      identifier_list: identifier_list COMMA IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      parameter_type_list: parameter_list COMMA ELLIPSIS., RPARAM
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      parameter_list: parameter_list COMMA parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL struct_declaration_list RCURL., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      struct_declarator: COLON constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      struct_declarator: declarator COLON.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 133
      constant_expression -> State 329

  State 306:
    Kernel Items:
      struct_declarator_list: struct_declarator_list COMMA.struct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 10
      LPARAM -> State 13
      COLON -> State 235
      MUL -> State 14
      struct_declarator -> State 330
      declarator -> State 236
      direct_declarator -> State 29
      pointer -> State 33

  State 307:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list SEMICOLON., *
    Reduce:
      * -> [struct_declaration]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      initializer: LCURL initializer_list COMMA.RCURL
      initializer_list: initializer_list COMMA.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      LCURL -> State 153
      RCURL -> State 331
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 154
      initializer -> State 332

  State 309:
    Kernel Items:
      initializer: LCURL initializer_list RCURL., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      labeled_statement: CASE constant_expression COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM.expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 333

  State 312:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement.RPARAM statement
      iteration_statement: FOR LPARAM expression_statement expression_statement.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      RPARAM -> State 334
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 335

  State 313:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM.statement
      selection_statement: IF LPARAM expression RPARAM.statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 338
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 314:
    Kernel Items:
      cast_expression: LPARAM type_name RPARAM cast_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      unary_expression: SIZEOF LPARAM type_name RPARAM., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      selection_statement: SWITCH LPARAM expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 339
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 317:
    Kernel Items:
      iteration_statement: WHILE LPARAM expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 340
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 318:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression COLON.conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 135
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 341

  State 319:
    Kernel Items:
      postfix_expression: postfix_expression LBRACE expression RBRACE., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      argument_expression_list: argument_expression_list COMMA.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      LPARAM -> State 89
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 342

  State 321:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM argument_expression_list RPARAM., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      direct_abstract_declarator: LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      direct_abstract_declarator: LPARAM abstract_declarator RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      direct_abstract_declarator: LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 343

  State 327:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 344

  State 329:
    Kernel Items:
      struct_declarator: declarator COLON constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      struct_declarator_list: struct_declarator_list COMMA struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      initializer: LCURL initializer_list COMMA RCURL., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      initializer_list: initializer_list COMMA initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: DO statement WHILE LPARAM expression.RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 345
      COMMA -> State 187

  State 334:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 346
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 335:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: FOR LPARAM expression_statement expression_statement expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 347
      COMMA -> State 187

  State 336:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement., *
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement
    Reduce:
      * -> [selection_statement]
    Goto:
      ELSE -> State 348

  State 337:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement., AND
      selection_statement: IF LPARAM expression RPARAM statement., BREAK
      selection_statement: IF LPARAM expression RPARAM statement., CASE
      selection_statement: IF LPARAM expression RPARAM statement., CONSTANT
      selection_statement: IF LPARAM expression RPARAM statement., CONTINUE
      selection_statement: IF LPARAM expression RPARAM statement., DEC_OP
      selection_statement: IF LPARAM expression RPARAM statement., DEFAULT
      selection_statement: IF LPARAM expression RPARAM statement., DO
      selection_statement: IF LPARAM expression RPARAM statement., ELSE
      selection_statement: IF LPARAM expression RPARAM statement., EXCLAIM
      selection_statement: IF LPARAM expression RPARAM statement., FOR
      selection_statement: IF LPARAM expression RPARAM statement., GOTO
      selection_statement: IF LPARAM expression RPARAM statement., IDENTIFIER
      selection_statement: IF LPARAM expression RPARAM statement., IF
      selection_statement: IF LPARAM expression RPARAM statement., INC_OP
      selection_statement: IF LPARAM expression RPARAM statement., LCURL
      selection_statement: IF LPARAM expression RPARAM statement., LPARAM
      selection_statement: IF LPARAM expression RPARAM statement., MINUS
      selection_statement: IF LPARAM expression RPARAM statement., MUL
      selection_statement: IF LPARAM expression RPARAM statement., PLUS
      selection_statement: IF LPARAM expression RPARAM statement., RCURL
      selection_statement: IF LPARAM expression RPARAM statement., RETURN
      selection_statement: IF LPARAM expression RPARAM statement., SEMICOLON
      selection_statement: IF LPARAM expression RPARAM statement., SIZEOF
      selection_statement: IF LPARAM expression RPARAM statement., STRING_LITERAL
      selection_statement: IF LPARAM expression RPARAM statement., SWITCH
      selection_statement: IF LPARAM expression RPARAM statement., TILDA
      selection_statement: IF LPARAM expression RPARAM statement., WHILE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , AND
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , BREAK
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , CASE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , CONSTANT
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , CONTINUE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , DEC_OP
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , DEFAULT
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , DO
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , ELSE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , EXCLAIM
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , FOR
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , GOTO
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , IDENTIFIER
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , IF
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , INC_OP
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , LCURL
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , LPARAM
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , MINUS
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , MUL
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , PLUS
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , RCURL
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , RETURN
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , SEMICOLON
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , SIZEOF
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , STRING_LITERAL
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , SWITCH
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , TILDA
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , WHILE
    Reduce:
      IDENTIFIER -> [selection_statement]
      CONSTANT -> [selection_statement]
      STRING_LITERAL -> [selection_statement]
      SIZEOF -> [selection_statement]
      INC_OP -> [selection_statement]
      DEC_OP -> [selection_statement]
      CASE -> [selection_statement]
      DEFAULT -> [selection_statement]
      IF -> [selection_statement]
      ELSE -> [selection_statement]
      SWITCH -> [selection_statement]
      WHILE -> [selection_statement]
      DO -> [selection_statement]
      FOR -> [selection_statement]
      GOTO -> [selection_statement]
      CONTINUE -> [selection_statement]
      BREAK -> [selection_statement]
      RETURN -> [selection_statement]
      LPARAM -> [selection_statement]
      LCURL -> [selection_statement]
      RCURL -> [selection_statement]
      SEMICOLON -> [selection_statement]
      MUL -> [selection_statement]
      MINUS -> [selection_statement]
      PLUS -> [selection_statement]
      AND -> [selection_statement]
      EXCLAIM -> [selection_statement]
      TILDA -> [selection_statement]
    Goto:
      ELSE -> State 348
    Shift/reduce conflict symbols:
      [ELSE]

  State 338:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement., ELSE
      selection_statement: IF LPARAM expression RPARAM statement., WHILE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , ELSE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , WHILE
    Reduce:
      ELSE -> [selection_statement]
      WHILE -> [selection_statement]
    Goto:
      ELSE -> State 348
    Shift/reduce conflict symbols:
      [ELSE]

  State 339:
    Kernel Items:
      selection_statement: SWITCH LPARAM expression RPARAM statement., *
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      iteration_statement: WHILE LPARAM expression RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression COLON conditional_expression., *
    Reduce:
      * -> [conditional_expression]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      argument_expression_list: argument_expression_list COMMA assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM expression RPARAM.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 349

  State 346:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 350
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 348:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement ELSE.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 86
      CONSTANT -> State 78
      STRING_LITERAL -> State 97
      SIZEOF -> State 96
      INC_OP -> State 88
      DEC_OP -> State 80
      CASE -> State 77
      DEFAULT -> State 81
      IF -> State 87
      SWITCH -> State 98
      WHILE -> State 100
      DO -> State 82
      FOR -> State 84
      GOTO -> State 85
      CONTINUE -> State 79
      BREAK -> State 76
      RETURN -> State 94
      LPARAM -> State 89
      LCURL -> State 49
      SEMICOLON -> State 95
      MUL -> State 91
      MINUS -> State 90
      PLUS -> State 92
      AND -> State 75
      EXCLAIM -> State 83
      TILDA -> State 99
      primary_expression -> State 120
      postfix_expression -> State 119
      unary_expression -> State 126
      unary_operator -> State 127
      cast_expression -> State 104
      multiplicative_expression -> State 118
      additive_expression -> State 101
      shift_expression -> State 123
      relational_expression -> State 121
      equality_expression -> State 108
      and_expression -> State 102
      exclusive_or_expression -> State 109
      inclusive_or_expression -> State 112
      logical_and_expression -> State 116
      logical_or_expression -> State 117
      conditional_expression -> State 106
      assignment_expression -> State 103
      expression -> State 110
      statement -> State 351
      labeled_statement -> State 115
      compound_statement -> State 105
      expression_statement -> State 111
      selection_statement -> State 122
      iteration_statement -> State 113
      jump_statement -> State 114

  State 349:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM expression RPARAM SEMICOLON., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement expression RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement ELSE statement., *
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

Number of states: 351
Number of shift actions: 2989
Number of reduce actions: 246
Number of shift/reduce conflicts: 2
Number of reduce/reduce conflicts: 0
*/
