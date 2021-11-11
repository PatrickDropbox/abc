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
	var errRetVal Symbol
	stateStack := _CStack{
		// Note: we don't have to populate the start symbol since its value is never accessed
		&_CStackItem{_CState0, nil},
	}
	symbolStack := &_CPseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return errRetVal, err
		}

		action, ok := _CActionTable.Get(stateStack[len(stateStack)-1].StateId, nextSymbol.Id())
		if !ok {
			return errRetVal, errHandler.Error(nextSymbol, stateStack)
		}
		if action.ActionType == _CShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return errRetVal, err
			}
		} else if action.ActionType == _CReduceAction {
			var reduceSymbol *CSymbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(reducer, stateStack)
			if err != nil {
				return errRetVal, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _CAcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1].T, nil

		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

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

type _CReduceType string

const (
	_CReduceAToPrimaryExpression        = _CReduceType("AToPrimaryExpression")
	_CReduceBToPrimaryExpression        = _CReduceType("BToPrimaryExpression")
	_CReduceCToPrimaryExpression        = _CReduceType("CToPrimaryExpression")
	_CReduceDToPrimaryExpression        = _CReduceType("DToPrimaryExpression")
	_CReduceAToPostfixExpression        = _CReduceType("AToPostfixExpression")
	_CReduceBToPostfixExpression        = _CReduceType("BToPostfixExpression")
	_CReduceCToPostfixExpression        = _CReduceType("CToPostfixExpression")
	_CReduceDToPostfixExpression        = _CReduceType("DToPostfixExpression")
	_CReduceEToPostfixExpression        = _CReduceType("EToPostfixExpression")
	_CReduceFToPostfixExpression        = _CReduceType("FToPostfixExpression")
	_CReduceGToPostfixExpression        = _CReduceType("GToPostfixExpression")
	_CReduceHToPostfixExpression        = _CReduceType("HToPostfixExpression")
	_CReduceAToArgumentExpressionList   = _CReduceType("AToArgumentExpressionList")
	_CReduceBToArgumentExpressionList   = _CReduceType("BToArgumentExpressionList")
	_CReduceAToUnaryExpression          = _CReduceType("AToUnaryExpression")
	_CReduceBToUnaryExpression          = _CReduceType("BToUnaryExpression")
	_CReduceCToUnaryExpression          = _CReduceType("CToUnaryExpression")
	_CReduceDToUnaryExpression          = _CReduceType("DToUnaryExpression")
	_CReduceEToUnaryExpression          = _CReduceType("EToUnaryExpression")
	_CReduceFToUnaryExpression          = _CReduceType("FToUnaryExpression")
	_CReduceAToUnaryOperator            = _CReduceType("AToUnaryOperator")
	_CReduceBToUnaryOperator            = _CReduceType("BToUnaryOperator")
	_CReduceCToUnaryOperator            = _CReduceType("CToUnaryOperator")
	_CReduceDToUnaryOperator            = _CReduceType("DToUnaryOperator")
	_CReduceEToUnaryOperator            = _CReduceType("EToUnaryOperator")
	_CReduceFToUnaryOperator            = _CReduceType("FToUnaryOperator")
	_CReduceAToCastExpression           = _CReduceType("AToCastExpression")
	_CReduceBToCastExpression           = _CReduceType("BToCastExpression")
	_CReduceAToMultiplicativeExpression = _CReduceType("AToMultiplicativeExpression")
	_CReduceBToMultiplicativeExpression = _CReduceType("BToMultiplicativeExpression")
	_CReduceCToMultiplicativeExpression = _CReduceType("CToMultiplicativeExpression")
	_CReduceDToMultiplicativeExpression = _CReduceType("DToMultiplicativeExpression")
	_CReduceAToAdditiveExpression       = _CReduceType("AToAdditiveExpression")
	_CReduceBToAdditiveExpression       = _CReduceType("BToAdditiveExpression")
	_CReduceCToAdditiveExpression       = _CReduceType("CToAdditiveExpression")
	_CReduceAToShiftExpression          = _CReduceType("AToShiftExpression")
	_CReduceBToShiftExpression          = _CReduceType("BToShiftExpression")
	_CReduceCToShiftExpression          = _CReduceType("CToShiftExpression")
	_CReduceAToRelationalExpression     = _CReduceType("AToRelationalExpression")
	_CReduceBToRelationalExpression     = _CReduceType("BToRelationalExpression")
	_CReduceCToRelationalExpression     = _CReduceType("CToRelationalExpression")
	_CReduceDToRelationalExpression     = _CReduceType("DToRelationalExpression")
	_CReduceEToRelationalExpression     = _CReduceType("EToRelationalExpression")
	_CReduceAToEqualityExpression       = _CReduceType("AToEqualityExpression")
	_CReduceBToEqualityExpression       = _CReduceType("BToEqualityExpression")
	_CReduceCToEqualityExpression       = _CReduceType("CToEqualityExpression")
	_CReduceAToAndExpression            = _CReduceType("AToAndExpression")
	_CReduceBToAndExpression            = _CReduceType("BToAndExpression")
	_CReduceAToExclusiveOrExpression    = _CReduceType("AToExclusiveOrExpression")
	_CReduceBToExclusiveOrExpression    = _CReduceType("BToExclusiveOrExpression")
	_CReduceAToInclusiveOrExpression    = _CReduceType("AToInclusiveOrExpression")
	_CReduceBToInclusiveOrExpression    = _CReduceType("BToInclusiveOrExpression")
	_CReduceAToLogicalAndExpression     = _CReduceType("AToLogicalAndExpression")
	_CReduceBToLogicalAndExpression     = _CReduceType("BToLogicalAndExpression")
	_CReduceAToLogicalOrExpression      = _CReduceType("AToLogicalOrExpression")
	_CReduceBToLogicalOrExpression      = _CReduceType("BToLogicalOrExpression")
	_CReduceAToConditionalExpression    = _CReduceType("AToConditionalExpression")
	_CReduceBToConditionalExpression    = _CReduceType("BToConditionalExpression")
	_CReduceAToAssignmentExpression     = _CReduceType("AToAssignmentExpression")
	_CReduceBToAssignmentExpression     = _CReduceType("BToAssignmentExpression")
	_CReduceAToAssignmentOperator       = _CReduceType("AToAssignmentOperator")
	_CReduceBToAssignmentOperator       = _CReduceType("BToAssignmentOperator")
	_CReduceCToAssignmentOperator       = _CReduceType("CToAssignmentOperator")
	_CReduceDToAssignmentOperator       = _CReduceType("DToAssignmentOperator")
	_CReduceEToAssignmentOperator       = _CReduceType("EToAssignmentOperator")
	_CReduceFToAssignmentOperator       = _CReduceType("FToAssignmentOperator")
	_CReduceGToAssignmentOperator       = _CReduceType("GToAssignmentOperator")
	_CReduceHToAssignmentOperator       = _CReduceType("HToAssignmentOperator")
	_CReduceIToAssignmentOperator       = _CReduceType("IToAssignmentOperator")
	_CReduceJToAssignmentOperator       = _CReduceType("JToAssignmentOperator")
	_CReduceKToAssignmentOperator       = _CReduceType("KToAssignmentOperator")
	_CReduceAToExpression               = _CReduceType("AToExpression")
	_CReduceBToExpression               = _CReduceType("BToExpression")
	_CReduceAToConstantExpression       = _CReduceType("AToConstantExpression")
	_CReduceAToDeclaration              = _CReduceType("AToDeclaration")
	_CReduceBToDeclaration              = _CReduceType("BToDeclaration")
	_CReduceAToDeclarationSpecifiers    = _CReduceType("AToDeclarationSpecifiers")
	_CReduceBToDeclarationSpecifiers    = _CReduceType("BToDeclarationSpecifiers")
	_CReduceCToDeclarationSpecifiers    = _CReduceType("CToDeclarationSpecifiers")
	_CReduceDToDeclarationSpecifiers    = _CReduceType("DToDeclarationSpecifiers")
	_CReduceEToDeclarationSpecifiers    = _CReduceType("EToDeclarationSpecifiers")
	_CReduceFToDeclarationSpecifiers    = _CReduceType("FToDeclarationSpecifiers")
	_CReduceAToInitDeclaratorList       = _CReduceType("AToInitDeclaratorList")
	_CReduceBToInitDeclaratorList       = _CReduceType("BToInitDeclaratorList")
	_CReduceAToInitDeclarator           = _CReduceType("AToInitDeclarator")
	_CReduceBToInitDeclarator           = _CReduceType("BToInitDeclarator")
	_CReduceAToStorageClassSpecifier    = _CReduceType("AToStorageClassSpecifier")
	_CReduceBToStorageClassSpecifier    = _CReduceType("BToStorageClassSpecifier")
	_CReduceCToStorageClassSpecifier    = _CReduceType("CToStorageClassSpecifier")
	_CReduceDToStorageClassSpecifier    = _CReduceType("DToStorageClassSpecifier")
	_CReduceEToStorageClassSpecifier    = _CReduceType("EToStorageClassSpecifier")
	_CReduceAToTypeSpecifier            = _CReduceType("AToTypeSpecifier")
	_CReduceBToTypeSpecifier            = _CReduceType("BToTypeSpecifier")
	_CReduceCToTypeSpecifier            = _CReduceType("CToTypeSpecifier")
	_CReduceDToTypeSpecifier            = _CReduceType("DToTypeSpecifier")
	_CReduceEToTypeSpecifier            = _CReduceType("EToTypeSpecifier")
	_CReduceFToTypeSpecifier            = _CReduceType("FToTypeSpecifier")
	_CReduceGToTypeSpecifier            = _CReduceType("GToTypeSpecifier")
	_CReduceHToTypeSpecifier            = _CReduceType("HToTypeSpecifier")
	_CReduceIToTypeSpecifier            = _CReduceType("IToTypeSpecifier")
	_CReduceJToTypeSpecifier            = _CReduceType("JToTypeSpecifier")
	_CReduceKToTypeSpecifier            = _CReduceType("KToTypeSpecifier")
	_CReduceLToTypeSpecifier            = _CReduceType("LToTypeSpecifier")
	_CReduceAToStructOrUnionSpecifier   = _CReduceType("AToStructOrUnionSpecifier")
	_CReduceBToStructOrUnionSpecifier   = _CReduceType("BToStructOrUnionSpecifier")
	_CReduceCToStructOrUnionSpecifier   = _CReduceType("CToStructOrUnionSpecifier")
	_CReduceAToStructOrUnion            = _CReduceType("AToStructOrUnion")
	_CReduceBToStructOrUnion            = _CReduceType("BToStructOrUnion")
	_CReduceAToStructDeclarationList    = _CReduceType("AToStructDeclarationList")
	_CReduceBToStructDeclarationList    = _CReduceType("BToStructDeclarationList")
	_CReduceAToStructDeclaration        = _CReduceType("AToStructDeclaration")
	_CReduceAToSpecifierQualifierList   = _CReduceType("AToSpecifierQualifierList")
	_CReduceBToSpecifierQualifierList   = _CReduceType("BToSpecifierQualifierList")
	_CReduceCToSpecifierQualifierList   = _CReduceType("CToSpecifierQualifierList")
	_CReduceDToSpecifierQualifierList   = _CReduceType("DToSpecifierQualifierList")
	_CReduceAToStructDeclaratorList     = _CReduceType("AToStructDeclaratorList")
	_CReduceBToStructDeclaratorList     = _CReduceType("BToStructDeclaratorList")
	_CReduceAToStructDeclarator         = _CReduceType("AToStructDeclarator")
	_CReduceBToStructDeclarator         = _CReduceType("BToStructDeclarator")
	_CReduceCToStructDeclarator         = _CReduceType("CToStructDeclarator")
	_CReduceAToEnumSpecifier            = _CReduceType("AToEnumSpecifier")
	_CReduceBToEnumSpecifier            = _CReduceType("BToEnumSpecifier")
	_CReduceCToEnumSpecifier            = _CReduceType("CToEnumSpecifier")
	_CReduceAToEnumeratorList           = _CReduceType("AToEnumeratorList")
	_CReduceBToEnumeratorList           = _CReduceType("BToEnumeratorList")
	_CReduceAToEnumerator               = _CReduceType("AToEnumerator")
	_CReduceBToEnumerator               = _CReduceType("BToEnumerator")
	_CReduceAToTypeQualifier            = _CReduceType("AToTypeQualifier")
	_CReduceBToTypeQualifier            = _CReduceType("BToTypeQualifier")
	_CReduceAToDeclarator               = _CReduceType("AToDeclarator")
	_CReduceBToDeclarator               = _CReduceType("BToDeclarator")
	_CReduceAToDirectDeclarator         = _CReduceType("AToDirectDeclarator")
	_CReduceBToDirectDeclarator         = _CReduceType("BToDirectDeclarator")
	_CReduceCToDirectDeclarator         = _CReduceType("CToDirectDeclarator")
	_CReduceDToDirectDeclarator         = _CReduceType("DToDirectDeclarator")
	_CReduceEToDirectDeclarator         = _CReduceType("EToDirectDeclarator")
	_CReduceFToDirectDeclarator         = _CReduceType("FToDirectDeclarator")
	_CReduceGToDirectDeclarator         = _CReduceType("GToDirectDeclarator")
	_CReduceAToPointer                  = _CReduceType("AToPointer")
	_CReduceBToPointer                  = _CReduceType("BToPointer")
	_CReduceCToPointer                  = _CReduceType("CToPointer")
	_CReduceDToPointer                  = _CReduceType("DToPointer")
	_CReduceAToTypeQualifierList        = _CReduceType("AToTypeQualifierList")
	_CReduceBToTypeQualifierList        = _CReduceType("BToTypeQualifierList")
	_CReduceAToParameterTypeList        = _CReduceType("AToParameterTypeList")
	_CReduceBToParameterTypeList        = _CReduceType("BToParameterTypeList")
	_CReduceAToParameterList            = _CReduceType("AToParameterList")
	_CReduceBToParameterList            = _CReduceType("BToParameterList")
	_CReduceAToParameterDeclaration     = _CReduceType("AToParameterDeclaration")
	_CReduceBToParameterDeclaration     = _CReduceType("BToParameterDeclaration")
	_CReduceCToParameterDeclaration     = _CReduceType("CToParameterDeclaration")
	_CReduceAToIdentifierList           = _CReduceType("AToIdentifierList")
	_CReduceBToIdentifierList           = _CReduceType("BToIdentifierList")
	_CReduceAToTypeName                 = _CReduceType("AToTypeName")
	_CReduceBToTypeName                 = _CReduceType("BToTypeName")
	_CReduceAToAbstractDeclarator       = _CReduceType("AToAbstractDeclarator")
	_CReduceBToAbstractDeclarator       = _CReduceType("BToAbstractDeclarator")
	_CReduceCToAbstractDeclarator       = _CReduceType("CToAbstractDeclarator")
	_CReduceAToDirectAbstractDeclarator = _CReduceType("AToDirectAbstractDeclarator")
	_CReduceBToDirectAbstractDeclarator = _CReduceType("BToDirectAbstractDeclarator")
	_CReduceCToDirectAbstractDeclarator = _CReduceType("CToDirectAbstractDeclarator")
	_CReduceDToDirectAbstractDeclarator = _CReduceType("DToDirectAbstractDeclarator")
	_CReduceEToDirectAbstractDeclarator = _CReduceType("EToDirectAbstractDeclarator")
	_CReduceFToDirectAbstractDeclarator = _CReduceType("FToDirectAbstractDeclarator")
	_CReduceGToDirectAbstractDeclarator = _CReduceType("GToDirectAbstractDeclarator")
	_CReduceHToDirectAbstractDeclarator = _CReduceType("HToDirectAbstractDeclarator")
	_CReduceIToDirectAbstractDeclarator = _CReduceType("IToDirectAbstractDeclarator")
	_CReduceAToInitializer              = _CReduceType("AToInitializer")
	_CReduceBToInitializer              = _CReduceType("BToInitializer")
	_CReduceCToInitializer              = _CReduceType("CToInitializer")
	_CReduceAToInitializerList          = _CReduceType("AToInitializerList")
	_CReduceBToInitializerList          = _CReduceType("BToInitializerList")
	_CReduceAToStatement                = _CReduceType("AToStatement")
	_CReduceBToStatement                = _CReduceType("BToStatement")
	_CReduceCToStatement                = _CReduceType("CToStatement")
	_CReduceDToStatement                = _CReduceType("DToStatement")
	_CReduceEToStatement                = _CReduceType("EToStatement")
	_CReduceFToStatement                = _CReduceType("FToStatement")
	_CReduceAToLabeledStatement         = _CReduceType("AToLabeledStatement")
	_CReduceBToLabeledStatement         = _CReduceType("BToLabeledStatement")
	_CReduceCToLabeledStatement         = _CReduceType("CToLabeledStatement")
	_CReduceAToCompoundStatement        = _CReduceType("AToCompoundStatement")
	_CReduceBToCompoundStatement        = _CReduceType("BToCompoundStatement")
	_CReduceCToCompoundStatement        = _CReduceType("CToCompoundStatement")
	_CReduceDToCompoundStatement        = _CReduceType("DToCompoundStatement")
	_CReduceAToDeclarationList          = _CReduceType("AToDeclarationList")
	_CReduceBToDeclarationList          = _CReduceType("BToDeclarationList")
	_CReduceAToStatementList            = _CReduceType("AToStatementList")
	_CReduceBToStatementList            = _CReduceType("BToStatementList")
	_CReduceAToExpressionStatement      = _CReduceType("AToExpressionStatement")
	_CReduceBToExpressionStatement      = _CReduceType("BToExpressionStatement")
	_CReduceAToSelectionStatement       = _CReduceType("AToSelectionStatement")
	_CReduceBToSelectionStatement       = _CReduceType("BToSelectionStatement")
	_CReduceCToSelectionStatement       = _CReduceType("CToSelectionStatement")
	_CReduceAToIterationStatement       = _CReduceType("AToIterationStatement")
	_CReduceBToIterationStatement       = _CReduceType("BToIterationStatement")
	_CReduceCToIterationStatement       = _CReduceType("CToIterationStatement")
	_CReduceDToIterationStatement       = _CReduceType("DToIterationStatement")
	_CReduceAToJumpStatement            = _CReduceType("AToJumpStatement")
	_CReduceBToJumpStatement            = _CReduceType("BToJumpStatement")
	_CReduceCToJumpStatement            = _CReduceType("CToJumpStatement")
	_CReduceDToJumpStatement            = _CReduceType("DToJumpStatement")
	_CReduceEToJumpStatement            = _CReduceType("EToJumpStatement")
	_CReduceAToTranslationUnit          = _CReduceType("AToTranslationUnit")
	_CReduceBToTranslationUnit          = _CReduceType("BToTranslationUnit")
	_CReduceAToExternalDeclaration      = _CReduceType("AToExternalDeclaration")
	_CReduceBToExternalDeclaration      = _CReduceType("BToExternalDeclaration")
	_CReduceAToFunctionDefinition       = _CReduceType("AToFunctionDefinition")
	_CReduceBToFunctionDefinition       = _CReduceType("BToFunctionDefinition")
	_CReduceCToFunctionDefinition       = _CReduceType("CToFunctionDefinition")
	_CReduceDToFunctionDefinition       = _CReduceType("DToFunctionDefinition")
)

type _CStateId int

func (id _CStateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_CState0   = _CStateId(0)
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
		panic("Unknown reduce type: " + act.ReduceType)
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

var (
	_CGotoState0Action                        = &_CAction{_CShiftAction, _CState0, ""}
	_CGotoState1Action                        = &_CAction{_CShiftAction, _CState1, ""}
	_CGotoState2Action                        = &_CAction{_CShiftAction, _CState2, ""}
	_CGotoState3Action                        = &_CAction{_CShiftAction, _CState3, ""}
	_CGotoState4Action                        = &_CAction{_CShiftAction, _CState4, ""}
	_CGotoState5Action                        = &_CAction{_CShiftAction, _CState5, ""}
	_CGotoState6Action                        = &_CAction{_CShiftAction, _CState6, ""}
	_CGotoState7Action                        = &_CAction{_CShiftAction, _CState7, ""}
	_CGotoState8Action                        = &_CAction{_CShiftAction, _CState8, ""}
	_CGotoState9Action                        = &_CAction{_CShiftAction, _CState9, ""}
	_CGotoState10Action                       = &_CAction{_CShiftAction, _CState10, ""}
	_CGotoState11Action                       = &_CAction{_CShiftAction, _CState11, ""}
	_CGotoState12Action                       = &_CAction{_CShiftAction, _CState12, ""}
	_CGotoState13Action                       = &_CAction{_CShiftAction, _CState13, ""}
	_CGotoState14Action                       = &_CAction{_CShiftAction, _CState14, ""}
	_CGotoState15Action                       = &_CAction{_CShiftAction, _CState15, ""}
	_CGotoState16Action                       = &_CAction{_CShiftAction, _CState16, ""}
	_CGotoState17Action                       = &_CAction{_CShiftAction, _CState17, ""}
	_CGotoState18Action                       = &_CAction{_CShiftAction, _CState18, ""}
	_CGotoState19Action                       = &_CAction{_CShiftAction, _CState19, ""}
	_CGotoState20Action                       = &_CAction{_CShiftAction, _CState20, ""}
	_CGotoState21Action                       = &_CAction{_CShiftAction, _CState21, ""}
	_CGotoState22Action                       = &_CAction{_CShiftAction, _CState22, ""}
	_CGotoState23Action                       = &_CAction{_CShiftAction, _CState23, ""}
	_CGotoState24Action                       = &_CAction{_CShiftAction, _CState24, ""}
	_CGotoState25Action                       = &_CAction{_CShiftAction, _CState25, ""}
	_CGotoState26Action                       = &_CAction{_CShiftAction, _CState26, ""}
	_CGotoState27Action                       = &_CAction{_CShiftAction, _CState27, ""}
	_CGotoState28Action                       = &_CAction{_CShiftAction, _CState28, ""}
	_CGotoState29Action                       = &_CAction{_CShiftAction, _CState29, ""}
	_CGotoState30Action                       = &_CAction{_CShiftAction, _CState30, ""}
	_CGotoState31Action                       = &_CAction{_CShiftAction, _CState31, ""}
	_CGotoState32Action                       = &_CAction{_CShiftAction, _CState32, ""}
	_CGotoState33Action                       = &_CAction{_CShiftAction, _CState33, ""}
	_CGotoState34Action                       = &_CAction{_CShiftAction, _CState34, ""}
	_CGotoState35Action                       = &_CAction{_CShiftAction, _CState35, ""}
	_CGotoState36Action                       = &_CAction{_CShiftAction, _CState36, ""}
	_CGotoState37Action                       = &_CAction{_CShiftAction, _CState37, ""}
	_CGotoState38Action                       = &_CAction{_CShiftAction, _CState38, ""}
	_CGotoState39Action                       = &_CAction{_CShiftAction, _CState39, ""}
	_CGotoState40Action                       = &_CAction{_CShiftAction, _CState40, ""}
	_CGotoState41Action                       = &_CAction{_CShiftAction, _CState41, ""}
	_CGotoState42Action                       = &_CAction{_CShiftAction, _CState42, ""}
	_CGotoState43Action                       = &_CAction{_CShiftAction, _CState43, ""}
	_CGotoState44Action                       = &_CAction{_CShiftAction, _CState44, ""}
	_CGotoState45Action                       = &_CAction{_CShiftAction, _CState45, ""}
	_CGotoState46Action                       = &_CAction{_CShiftAction, _CState46, ""}
	_CGotoState47Action                       = &_CAction{_CShiftAction, _CState47, ""}
	_CGotoState48Action                       = &_CAction{_CShiftAction, _CState48, ""}
	_CGotoState49Action                       = &_CAction{_CShiftAction, _CState49, ""}
	_CGotoState50Action                       = &_CAction{_CShiftAction, _CState50, ""}
	_CGotoState51Action                       = &_CAction{_CShiftAction, _CState51, ""}
	_CGotoState52Action                       = &_CAction{_CShiftAction, _CState52, ""}
	_CGotoState53Action                       = &_CAction{_CShiftAction, _CState53, ""}
	_CGotoState54Action                       = &_CAction{_CShiftAction, _CState54, ""}
	_CGotoState55Action                       = &_CAction{_CShiftAction, _CState55, ""}
	_CGotoState56Action                       = &_CAction{_CShiftAction, _CState56, ""}
	_CGotoState57Action                       = &_CAction{_CShiftAction, _CState57, ""}
	_CGotoState58Action                       = &_CAction{_CShiftAction, _CState58, ""}
	_CGotoState59Action                       = &_CAction{_CShiftAction, _CState59, ""}
	_CGotoState60Action                       = &_CAction{_CShiftAction, _CState60, ""}
	_CGotoState61Action                       = &_CAction{_CShiftAction, _CState61, ""}
	_CGotoState62Action                       = &_CAction{_CShiftAction, _CState62, ""}
	_CGotoState63Action                       = &_CAction{_CShiftAction, _CState63, ""}
	_CGotoState64Action                       = &_CAction{_CShiftAction, _CState64, ""}
	_CGotoState65Action                       = &_CAction{_CShiftAction, _CState65, ""}
	_CGotoState66Action                       = &_CAction{_CShiftAction, _CState66, ""}
	_CGotoState67Action                       = &_CAction{_CShiftAction, _CState67, ""}
	_CGotoState68Action                       = &_CAction{_CShiftAction, _CState68, ""}
	_CGotoState69Action                       = &_CAction{_CShiftAction, _CState69, ""}
	_CGotoState70Action                       = &_CAction{_CShiftAction, _CState70, ""}
	_CGotoState71Action                       = &_CAction{_CShiftAction, _CState71, ""}
	_CGotoState72Action                       = &_CAction{_CShiftAction, _CState72, ""}
	_CGotoState73Action                       = &_CAction{_CShiftAction, _CState73, ""}
	_CGotoState74Action                       = &_CAction{_CShiftAction, _CState74, ""}
	_CGotoState75Action                       = &_CAction{_CShiftAction, _CState75, ""}
	_CGotoState76Action                       = &_CAction{_CShiftAction, _CState76, ""}
	_CGotoState77Action                       = &_CAction{_CShiftAction, _CState77, ""}
	_CGotoState78Action                       = &_CAction{_CShiftAction, _CState78, ""}
	_CGotoState79Action                       = &_CAction{_CShiftAction, _CState79, ""}
	_CGotoState80Action                       = &_CAction{_CShiftAction, _CState80, ""}
	_CGotoState81Action                       = &_CAction{_CShiftAction, _CState81, ""}
	_CGotoState82Action                       = &_CAction{_CShiftAction, _CState82, ""}
	_CGotoState83Action                       = &_CAction{_CShiftAction, _CState83, ""}
	_CGotoState84Action                       = &_CAction{_CShiftAction, _CState84, ""}
	_CGotoState85Action                       = &_CAction{_CShiftAction, _CState85, ""}
	_CGotoState86Action                       = &_CAction{_CShiftAction, _CState86, ""}
	_CGotoState87Action                       = &_CAction{_CShiftAction, _CState87, ""}
	_CGotoState88Action                       = &_CAction{_CShiftAction, _CState88, ""}
	_CGotoState89Action                       = &_CAction{_CShiftAction, _CState89, ""}
	_CGotoState90Action                       = &_CAction{_CShiftAction, _CState90, ""}
	_CGotoState91Action                       = &_CAction{_CShiftAction, _CState91, ""}
	_CGotoState92Action                       = &_CAction{_CShiftAction, _CState92, ""}
	_CGotoState93Action                       = &_CAction{_CShiftAction, _CState93, ""}
	_CGotoState94Action                       = &_CAction{_CShiftAction, _CState94, ""}
	_CGotoState95Action                       = &_CAction{_CShiftAction, _CState95, ""}
	_CGotoState96Action                       = &_CAction{_CShiftAction, _CState96, ""}
	_CGotoState97Action                       = &_CAction{_CShiftAction, _CState97, ""}
	_CGotoState98Action                       = &_CAction{_CShiftAction, _CState98, ""}
	_CGotoState99Action                       = &_CAction{_CShiftAction, _CState99, ""}
	_CGotoState100Action                      = &_CAction{_CShiftAction, _CState100, ""}
	_CGotoState101Action                      = &_CAction{_CShiftAction, _CState101, ""}
	_CGotoState102Action                      = &_CAction{_CShiftAction, _CState102, ""}
	_CGotoState103Action                      = &_CAction{_CShiftAction, _CState103, ""}
	_CGotoState104Action                      = &_CAction{_CShiftAction, _CState104, ""}
	_CGotoState105Action                      = &_CAction{_CShiftAction, _CState105, ""}
	_CGotoState106Action                      = &_CAction{_CShiftAction, _CState106, ""}
	_CGotoState107Action                      = &_CAction{_CShiftAction, _CState107, ""}
	_CGotoState108Action                      = &_CAction{_CShiftAction, _CState108, ""}
	_CGotoState109Action                      = &_CAction{_CShiftAction, _CState109, ""}
	_CGotoState110Action                      = &_CAction{_CShiftAction, _CState110, ""}
	_CGotoState111Action                      = &_CAction{_CShiftAction, _CState111, ""}
	_CGotoState112Action                      = &_CAction{_CShiftAction, _CState112, ""}
	_CGotoState113Action                      = &_CAction{_CShiftAction, _CState113, ""}
	_CGotoState114Action                      = &_CAction{_CShiftAction, _CState114, ""}
	_CGotoState115Action                      = &_CAction{_CShiftAction, _CState115, ""}
	_CGotoState116Action                      = &_CAction{_CShiftAction, _CState116, ""}
	_CGotoState117Action                      = &_CAction{_CShiftAction, _CState117, ""}
	_CGotoState118Action                      = &_CAction{_CShiftAction, _CState118, ""}
	_CGotoState119Action                      = &_CAction{_CShiftAction, _CState119, ""}
	_CGotoState120Action                      = &_CAction{_CShiftAction, _CState120, ""}
	_CGotoState121Action                      = &_CAction{_CShiftAction, _CState121, ""}
	_CGotoState122Action                      = &_CAction{_CShiftAction, _CState122, ""}
	_CGotoState123Action                      = &_CAction{_CShiftAction, _CState123, ""}
	_CGotoState124Action                      = &_CAction{_CShiftAction, _CState124, ""}
	_CGotoState125Action                      = &_CAction{_CShiftAction, _CState125, ""}
	_CGotoState126Action                      = &_CAction{_CShiftAction, _CState126, ""}
	_CGotoState127Action                      = &_CAction{_CShiftAction, _CState127, ""}
	_CGotoState128Action                      = &_CAction{_CShiftAction, _CState128, ""}
	_CGotoState129Action                      = &_CAction{_CShiftAction, _CState129, ""}
	_CGotoState130Action                      = &_CAction{_CShiftAction, _CState130, ""}
	_CGotoState131Action                      = &_CAction{_CShiftAction, _CState131, ""}
	_CGotoState132Action                      = &_CAction{_CShiftAction, _CState132, ""}
	_CGotoState133Action                      = &_CAction{_CShiftAction, _CState133, ""}
	_CGotoState134Action                      = &_CAction{_CShiftAction, _CState134, ""}
	_CGotoState135Action                      = &_CAction{_CShiftAction, _CState135, ""}
	_CGotoState136Action                      = &_CAction{_CShiftAction, _CState136, ""}
	_CGotoState137Action                      = &_CAction{_CShiftAction, _CState137, ""}
	_CGotoState138Action                      = &_CAction{_CShiftAction, _CState138, ""}
	_CGotoState139Action                      = &_CAction{_CShiftAction, _CState139, ""}
	_CGotoState140Action                      = &_CAction{_CShiftAction, _CState140, ""}
	_CGotoState141Action                      = &_CAction{_CShiftAction, _CState141, ""}
	_CGotoState142Action                      = &_CAction{_CShiftAction, _CState142, ""}
	_CGotoState143Action                      = &_CAction{_CShiftAction, _CState143, ""}
	_CGotoState144Action                      = &_CAction{_CShiftAction, _CState144, ""}
	_CGotoState145Action                      = &_CAction{_CShiftAction, _CState145, ""}
	_CGotoState146Action                      = &_CAction{_CShiftAction, _CState146, ""}
	_CGotoState147Action                      = &_CAction{_CShiftAction, _CState147, ""}
	_CGotoState148Action                      = &_CAction{_CShiftAction, _CState148, ""}
	_CGotoState149Action                      = &_CAction{_CShiftAction, _CState149, ""}
	_CGotoState150Action                      = &_CAction{_CShiftAction, _CState150, ""}
	_CGotoState151Action                      = &_CAction{_CShiftAction, _CState151, ""}
	_CGotoState152Action                      = &_CAction{_CShiftAction, _CState152, ""}
	_CGotoState153Action                      = &_CAction{_CShiftAction, _CState153, ""}
	_CGotoState154Action                      = &_CAction{_CShiftAction, _CState154, ""}
	_CGotoState155Action                      = &_CAction{_CShiftAction, _CState155, ""}
	_CGotoState156Action                      = &_CAction{_CShiftAction, _CState156, ""}
	_CGotoState157Action                      = &_CAction{_CShiftAction, _CState157, ""}
	_CGotoState158Action                      = &_CAction{_CShiftAction, _CState158, ""}
	_CGotoState159Action                      = &_CAction{_CShiftAction, _CState159, ""}
	_CGotoState160Action                      = &_CAction{_CShiftAction, _CState160, ""}
	_CGotoState161Action                      = &_CAction{_CShiftAction, _CState161, ""}
	_CGotoState162Action                      = &_CAction{_CShiftAction, _CState162, ""}
	_CGotoState163Action                      = &_CAction{_CShiftAction, _CState163, ""}
	_CGotoState164Action                      = &_CAction{_CShiftAction, _CState164, ""}
	_CGotoState165Action                      = &_CAction{_CShiftAction, _CState165, ""}
	_CGotoState166Action                      = &_CAction{_CShiftAction, _CState166, ""}
	_CGotoState167Action                      = &_CAction{_CShiftAction, _CState167, ""}
	_CGotoState168Action                      = &_CAction{_CShiftAction, _CState168, ""}
	_CGotoState169Action                      = &_CAction{_CShiftAction, _CState169, ""}
	_CGotoState170Action                      = &_CAction{_CShiftAction, _CState170, ""}
	_CGotoState171Action                      = &_CAction{_CShiftAction, _CState171, ""}
	_CGotoState172Action                      = &_CAction{_CShiftAction, _CState172, ""}
	_CGotoState173Action                      = &_CAction{_CShiftAction, _CState173, ""}
	_CGotoState174Action                      = &_CAction{_CShiftAction, _CState174, ""}
	_CGotoState175Action                      = &_CAction{_CShiftAction, _CState175, ""}
	_CGotoState176Action                      = &_CAction{_CShiftAction, _CState176, ""}
	_CGotoState177Action                      = &_CAction{_CShiftAction, _CState177, ""}
	_CGotoState178Action                      = &_CAction{_CShiftAction, _CState178, ""}
	_CGotoState179Action                      = &_CAction{_CShiftAction, _CState179, ""}
	_CGotoState180Action                      = &_CAction{_CShiftAction, _CState180, ""}
	_CGotoState181Action                      = &_CAction{_CShiftAction, _CState181, ""}
	_CGotoState182Action                      = &_CAction{_CShiftAction, _CState182, ""}
	_CGotoState183Action                      = &_CAction{_CShiftAction, _CState183, ""}
	_CGotoState184Action                      = &_CAction{_CShiftAction, _CState184, ""}
	_CGotoState185Action                      = &_CAction{_CShiftAction, _CState185, ""}
	_CGotoState186Action                      = &_CAction{_CShiftAction, _CState186, ""}
	_CGotoState187Action                      = &_CAction{_CShiftAction, _CState187, ""}
	_CGotoState188Action                      = &_CAction{_CShiftAction, _CState188, ""}
	_CGotoState189Action                      = &_CAction{_CShiftAction, _CState189, ""}
	_CGotoState190Action                      = &_CAction{_CShiftAction, _CState190, ""}
	_CGotoState191Action                      = &_CAction{_CShiftAction, _CState191, ""}
	_CGotoState192Action                      = &_CAction{_CShiftAction, _CState192, ""}
	_CGotoState193Action                      = &_CAction{_CShiftAction, _CState193, ""}
	_CGotoState194Action                      = &_CAction{_CShiftAction, _CState194, ""}
	_CGotoState195Action                      = &_CAction{_CShiftAction, _CState195, ""}
	_CGotoState196Action                      = &_CAction{_CShiftAction, _CState196, ""}
	_CGotoState197Action                      = &_CAction{_CShiftAction, _CState197, ""}
	_CGotoState198Action                      = &_CAction{_CShiftAction, _CState198, ""}
	_CGotoState199Action                      = &_CAction{_CShiftAction, _CState199, ""}
	_CGotoState200Action                      = &_CAction{_CShiftAction, _CState200, ""}
	_CGotoState201Action                      = &_CAction{_CShiftAction, _CState201, ""}
	_CGotoState202Action                      = &_CAction{_CShiftAction, _CState202, ""}
	_CGotoState203Action                      = &_CAction{_CShiftAction, _CState203, ""}
	_CGotoState204Action                      = &_CAction{_CShiftAction, _CState204, ""}
	_CGotoState205Action                      = &_CAction{_CShiftAction, _CState205, ""}
	_CGotoState206Action                      = &_CAction{_CShiftAction, _CState206, ""}
	_CGotoState207Action                      = &_CAction{_CShiftAction, _CState207, ""}
	_CGotoState208Action                      = &_CAction{_CShiftAction, _CState208, ""}
	_CGotoState209Action                      = &_CAction{_CShiftAction, _CState209, ""}
	_CGotoState210Action                      = &_CAction{_CShiftAction, _CState210, ""}
	_CGotoState211Action                      = &_CAction{_CShiftAction, _CState211, ""}
	_CGotoState212Action                      = &_CAction{_CShiftAction, _CState212, ""}
	_CGotoState213Action                      = &_CAction{_CShiftAction, _CState213, ""}
	_CGotoState214Action                      = &_CAction{_CShiftAction, _CState214, ""}
	_CGotoState215Action                      = &_CAction{_CShiftAction, _CState215, ""}
	_CGotoState216Action                      = &_CAction{_CShiftAction, _CState216, ""}
	_CGotoState217Action                      = &_CAction{_CShiftAction, _CState217, ""}
	_CGotoState218Action                      = &_CAction{_CShiftAction, _CState218, ""}
	_CGotoState219Action                      = &_CAction{_CShiftAction, _CState219, ""}
	_CGotoState220Action                      = &_CAction{_CShiftAction, _CState220, ""}
	_CGotoState221Action                      = &_CAction{_CShiftAction, _CState221, ""}
	_CGotoState222Action                      = &_CAction{_CShiftAction, _CState222, ""}
	_CGotoState223Action                      = &_CAction{_CShiftAction, _CState223, ""}
	_CGotoState224Action                      = &_CAction{_CShiftAction, _CState224, ""}
	_CGotoState225Action                      = &_CAction{_CShiftAction, _CState225, ""}
	_CGotoState226Action                      = &_CAction{_CShiftAction, _CState226, ""}
	_CGotoState227Action                      = &_CAction{_CShiftAction, _CState227, ""}
	_CGotoState228Action                      = &_CAction{_CShiftAction, _CState228, ""}
	_CGotoState229Action                      = &_CAction{_CShiftAction, _CState229, ""}
	_CGotoState230Action                      = &_CAction{_CShiftAction, _CState230, ""}
	_CGotoState231Action                      = &_CAction{_CShiftAction, _CState231, ""}
	_CGotoState232Action                      = &_CAction{_CShiftAction, _CState232, ""}
	_CGotoState233Action                      = &_CAction{_CShiftAction, _CState233, ""}
	_CGotoState234Action                      = &_CAction{_CShiftAction, _CState234, ""}
	_CGotoState235Action                      = &_CAction{_CShiftAction, _CState235, ""}
	_CGotoState236Action                      = &_CAction{_CShiftAction, _CState236, ""}
	_CGotoState237Action                      = &_CAction{_CShiftAction, _CState237, ""}
	_CGotoState238Action                      = &_CAction{_CShiftAction, _CState238, ""}
	_CGotoState239Action                      = &_CAction{_CShiftAction, _CState239, ""}
	_CGotoState240Action                      = &_CAction{_CShiftAction, _CState240, ""}
	_CGotoState241Action                      = &_CAction{_CShiftAction, _CState241, ""}
	_CGotoState242Action                      = &_CAction{_CShiftAction, _CState242, ""}
	_CGotoState243Action                      = &_CAction{_CShiftAction, _CState243, ""}
	_CGotoState244Action                      = &_CAction{_CShiftAction, _CState244, ""}
	_CGotoState245Action                      = &_CAction{_CShiftAction, _CState245, ""}
	_CGotoState246Action                      = &_CAction{_CShiftAction, _CState246, ""}
	_CGotoState247Action                      = &_CAction{_CShiftAction, _CState247, ""}
	_CGotoState248Action                      = &_CAction{_CShiftAction, _CState248, ""}
	_CGotoState249Action                      = &_CAction{_CShiftAction, _CState249, ""}
	_CGotoState250Action                      = &_CAction{_CShiftAction, _CState250, ""}
	_CGotoState251Action                      = &_CAction{_CShiftAction, _CState251, ""}
	_CGotoState252Action                      = &_CAction{_CShiftAction, _CState252, ""}
	_CGotoState253Action                      = &_CAction{_CShiftAction, _CState253, ""}
	_CGotoState254Action                      = &_CAction{_CShiftAction, _CState254, ""}
	_CGotoState255Action                      = &_CAction{_CShiftAction, _CState255, ""}
	_CGotoState256Action                      = &_CAction{_CShiftAction, _CState256, ""}
	_CGotoState257Action                      = &_CAction{_CShiftAction, _CState257, ""}
	_CGotoState258Action                      = &_CAction{_CShiftAction, _CState258, ""}
	_CGotoState259Action                      = &_CAction{_CShiftAction, _CState259, ""}
	_CGotoState260Action                      = &_CAction{_CShiftAction, _CState260, ""}
	_CGotoState261Action                      = &_CAction{_CShiftAction, _CState261, ""}
	_CGotoState262Action                      = &_CAction{_CShiftAction, _CState262, ""}
	_CGotoState263Action                      = &_CAction{_CShiftAction, _CState263, ""}
	_CGotoState264Action                      = &_CAction{_CShiftAction, _CState264, ""}
	_CGotoState265Action                      = &_CAction{_CShiftAction, _CState265, ""}
	_CGotoState266Action                      = &_CAction{_CShiftAction, _CState266, ""}
	_CGotoState267Action                      = &_CAction{_CShiftAction, _CState267, ""}
	_CGotoState268Action                      = &_CAction{_CShiftAction, _CState268, ""}
	_CGotoState269Action                      = &_CAction{_CShiftAction, _CState269, ""}
	_CGotoState270Action                      = &_CAction{_CShiftAction, _CState270, ""}
	_CGotoState271Action                      = &_CAction{_CShiftAction, _CState271, ""}
	_CGotoState272Action                      = &_CAction{_CShiftAction, _CState272, ""}
	_CGotoState273Action                      = &_CAction{_CShiftAction, _CState273, ""}
	_CGotoState274Action                      = &_CAction{_CShiftAction, _CState274, ""}
	_CGotoState275Action                      = &_CAction{_CShiftAction, _CState275, ""}
	_CGotoState276Action                      = &_CAction{_CShiftAction, _CState276, ""}
	_CGotoState277Action                      = &_CAction{_CShiftAction, _CState277, ""}
	_CGotoState278Action                      = &_CAction{_CShiftAction, _CState278, ""}
	_CGotoState279Action                      = &_CAction{_CShiftAction, _CState279, ""}
	_CGotoState280Action                      = &_CAction{_CShiftAction, _CState280, ""}
	_CGotoState281Action                      = &_CAction{_CShiftAction, _CState281, ""}
	_CGotoState282Action                      = &_CAction{_CShiftAction, _CState282, ""}
	_CGotoState283Action                      = &_CAction{_CShiftAction, _CState283, ""}
	_CGotoState284Action                      = &_CAction{_CShiftAction, _CState284, ""}
	_CGotoState285Action                      = &_CAction{_CShiftAction, _CState285, ""}
	_CGotoState286Action                      = &_CAction{_CShiftAction, _CState286, ""}
	_CGotoState287Action                      = &_CAction{_CShiftAction, _CState287, ""}
	_CGotoState288Action                      = &_CAction{_CShiftAction, _CState288, ""}
	_CGotoState289Action                      = &_CAction{_CShiftAction, _CState289, ""}
	_CGotoState290Action                      = &_CAction{_CShiftAction, _CState290, ""}
	_CGotoState291Action                      = &_CAction{_CShiftAction, _CState291, ""}
	_CGotoState292Action                      = &_CAction{_CShiftAction, _CState292, ""}
	_CGotoState293Action                      = &_CAction{_CShiftAction, _CState293, ""}
	_CGotoState294Action                      = &_CAction{_CShiftAction, _CState294, ""}
	_CGotoState295Action                      = &_CAction{_CShiftAction, _CState295, ""}
	_CGotoState296Action                      = &_CAction{_CShiftAction, _CState296, ""}
	_CGotoState297Action                      = &_CAction{_CShiftAction, _CState297, ""}
	_CGotoState298Action                      = &_CAction{_CShiftAction, _CState298, ""}
	_CGotoState299Action                      = &_CAction{_CShiftAction, _CState299, ""}
	_CGotoState300Action                      = &_CAction{_CShiftAction, _CState300, ""}
	_CGotoState301Action                      = &_CAction{_CShiftAction, _CState301, ""}
	_CGotoState302Action                      = &_CAction{_CShiftAction, _CState302, ""}
	_CGotoState303Action                      = &_CAction{_CShiftAction, _CState303, ""}
	_CGotoState304Action                      = &_CAction{_CShiftAction, _CState304, ""}
	_CGotoState305Action                      = &_CAction{_CShiftAction, _CState305, ""}
	_CGotoState306Action                      = &_CAction{_CShiftAction, _CState306, ""}
	_CGotoState307Action                      = &_CAction{_CShiftAction, _CState307, ""}
	_CGotoState308Action                      = &_CAction{_CShiftAction, _CState308, ""}
	_CGotoState309Action                      = &_CAction{_CShiftAction, _CState309, ""}
	_CGotoState310Action                      = &_CAction{_CShiftAction, _CState310, ""}
	_CGotoState311Action                      = &_CAction{_CShiftAction, _CState311, ""}
	_CGotoState312Action                      = &_CAction{_CShiftAction, _CState312, ""}
	_CGotoState313Action                      = &_CAction{_CShiftAction, _CState313, ""}
	_CGotoState314Action                      = &_CAction{_CShiftAction, _CState314, ""}
	_CGotoState315Action                      = &_CAction{_CShiftAction, _CState315, ""}
	_CGotoState316Action                      = &_CAction{_CShiftAction, _CState316, ""}
	_CGotoState317Action                      = &_CAction{_CShiftAction, _CState317, ""}
	_CGotoState318Action                      = &_CAction{_CShiftAction, _CState318, ""}
	_CGotoState319Action                      = &_CAction{_CShiftAction, _CState319, ""}
	_CGotoState320Action                      = &_CAction{_CShiftAction, _CState320, ""}
	_CGotoState321Action                      = &_CAction{_CShiftAction, _CState321, ""}
	_CGotoState322Action                      = &_CAction{_CShiftAction, _CState322, ""}
	_CGotoState323Action                      = &_CAction{_CShiftAction, _CState323, ""}
	_CGotoState324Action                      = &_CAction{_CShiftAction, _CState324, ""}
	_CGotoState325Action                      = &_CAction{_CShiftAction, _CState325, ""}
	_CGotoState326Action                      = &_CAction{_CShiftAction, _CState326, ""}
	_CGotoState327Action                      = &_CAction{_CShiftAction, _CState327, ""}
	_CGotoState328Action                      = &_CAction{_CShiftAction, _CState328, ""}
	_CGotoState329Action                      = &_CAction{_CShiftAction, _CState329, ""}
	_CGotoState330Action                      = &_CAction{_CShiftAction, _CState330, ""}
	_CGotoState331Action                      = &_CAction{_CShiftAction, _CState331, ""}
	_CGotoState332Action                      = &_CAction{_CShiftAction, _CState332, ""}
	_CGotoState333Action                      = &_CAction{_CShiftAction, _CState333, ""}
	_CGotoState334Action                      = &_CAction{_CShiftAction, _CState334, ""}
	_CGotoState335Action                      = &_CAction{_CShiftAction, _CState335, ""}
	_CGotoState336Action                      = &_CAction{_CShiftAction, _CState336, ""}
	_CGotoState337Action                      = &_CAction{_CShiftAction, _CState337, ""}
	_CGotoState338Action                      = &_CAction{_CShiftAction, _CState338, ""}
	_CGotoState339Action                      = &_CAction{_CShiftAction, _CState339, ""}
	_CGotoState340Action                      = &_CAction{_CShiftAction, _CState340, ""}
	_CGotoState341Action                      = &_CAction{_CShiftAction, _CState341, ""}
	_CGotoState342Action                      = &_CAction{_CShiftAction, _CState342, ""}
	_CGotoState343Action                      = &_CAction{_CShiftAction, _CState343, ""}
	_CGotoState344Action                      = &_CAction{_CShiftAction, _CState344, ""}
	_CGotoState345Action                      = &_CAction{_CShiftAction, _CState345, ""}
	_CGotoState346Action                      = &_CAction{_CShiftAction, _CState346, ""}
	_CGotoState347Action                      = &_CAction{_CShiftAction, _CState347, ""}
	_CGotoState348Action                      = &_CAction{_CShiftAction, _CState348, ""}
	_CGotoState349Action                      = &_CAction{_CShiftAction, _CState349, ""}
	_CGotoState350Action                      = &_CAction{_CShiftAction, _CState350, ""}
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
	{_CState1, _CEndMarker}:                     &_CAction{_CAcceptAction, 0, ""},
	{_CState0, CIdentifierToken}:                _CGotoState9Action,
	{_CState0, CTypeNameToken}:                  _CGotoState20Action,
	{_CState0, CTypedefToken}:                   _CGotoState19Action,
	{_CState0, CExternToken}:                    _CGotoState7Action,
	{_CState0, CStaticToken}:                    _CGotoState17Action,
	{_CState0, CAutoToken}:                      _CGotoState2Action,
	{_CState0, CRegisterToken}:                  _CGotoState14Action,
	{_CState0, CCharToken}:                      _CGotoState3Action,
	{_CState0, CShortToken}:                     _CGotoState15Action,
	{_CState0, CIntToken}:                       _CGotoState10Action,
	{_CState0, CLongToken}:                      _CGotoState11Action,
	{_CState0, CSignedToken}:                    _CGotoState16Action,
	{_CState0, CUnsignedToken}:                  _CGotoState22Action,
	{_CState0, CFloatToken}:                     _CGotoState8Action,
	{_CState0, CDoubleToken}:                    _CGotoState5Action,
	{_CState0, CConstToken}:                     _CGotoState4Action,
	{_CState0, CVolatileToken}:                  _CGotoState24Action,
	{_CState0, CVoidToken}:                      _CGotoState23Action,
	{_CState0, CStructToken}:                    _CGotoState18Action,
	{_CState0, CUnionToken}:                     _CGotoState21Action,
	{_CState0, CEnumToken}:                      _CGotoState6Action,
	{_CState0, CLparamToken}:                    _CGotoState12Action,
	{_CState0, CMulToken}:                       _CGotoState13Action,
	{_CState0, CDeclarationType}:                _CGotoState25Action,
	{_CState0, CDeclarationSpecifiersType}:      _CGotoState26Action,
	{_CState0, CStorageClassSpecifierType}:      _CGotoState33Action,
	{_CState0, CTypeSpecifierType}:              _CGotoState37Action,
	{_CState0, CStructOrUnionSpecifierType}:     _CGotoState35Action,
	{_CState0, CStructOrUnionType}:              _CGotoState34Action,
	{_CState0, CEnumSpecifierType}:              _CGotoState29Action,
	{_CState0, CTypeQualifierType}:              _CGotoState36Action,
	{_CState0, CDeclaratorType}:                 _CGotoState27Action,
	{_CState0, CDirectDeclaratorType}:           _CGotoState28Action,
	{_CState0, CPointerType}:                    _CGotoState32Action,
	{_CState0, CTranslationUnitType}:            _CGotoState1Action,
	{_CState0, CExternalDeclarationType}:        _CGotoState30Action,
	{_CState0, CFunctionDefinitionType}:         _CGotoState31Action,
	{_CState1, CIdentifierToken}:                _CGotoState9Action,
	{_CState1, CTypeNameToken}:                  _CGotoState20Action,
	{_CState1, CTypedefToken}:                   _CGotoState19Action,
	{_CState1, CExternToken}:                    _CGotoState7Action,
	{_CState1, CStaticToken}:                    _CGotoState17Action,
	{_CState1, CAutoToken}:                      _CGotoState2Action,
	{_CState1, CRegisterToken}:                  _CGotoState14Action,
	{_CState1, CCharToken}:                      _CGotoState3Action,
	{_CState1, CShortToken}:                     _CGotoState15Action,
	{_CState1, CIntToken}:                       _CGotoState10Action,
	{_CState1, CLongToken}:                      _CGotoState11Action,
	{_CState1, CSignedToken}:                    _CGotoState16Action,
	{_CState1, CUnsignedToken}:                  _CGotoState22Action,
	{_CState1, CFloatToken}:                     _CGotoState8Action,
	{_CState1, CDoubleToken}:                    _CGotoState5Action,
	{_CState1, CConstToken}:                     _CGotoState4Action,
	{_CState1, CVolatileToken}:                  _CGotoState24Action,
	{_CState1, CVoidToken}:                      _CGotoState23Action,
	{_CState1, CStructToken}:                    _CGotoState18Action,
	{_CState1, CUnionToken}:                     _CGotoState21Action,
	{_CState1, CEnumToken}:                      _CGotoState6Action,
	{_CState1, CLparamToken}:                    _CGotoState12Action,
	{_CState1, CMulToken}:                       _CGotoState13Action,
	{_CState1, CDeclarationType}:                _CGotoState25Action,
	{_CState1, CDeclarationSpecifiersType}:      _CGotoState26Action,
	{_CState1, CStorageClassSpecifierType}:      _CGotoState33Action,
	{_CState1, CTypeSpecifierType}:              _CGotoState37Action,
	{_CState1, CStructOrUnionSpecifierType}:     _CGotoState35Action,
	{_CState1, CStructOrUnionType}:              _CGotoState34Action,
	{_CState1, CEnumSpecifierType}:              _CGotoState29Action,
	{_CState1, CTypeQualifierType}:              _CGotoState36Action,
	{_CState1, CDeclaratorType}:                 _CGotoState27Action,
	{_CState1, CDirectDeclaratorType}:           _CGotoState28Action,
	{_CState1, CPointerType}:                    _CGotoState32Action,
	{_CState1, CExternalDeclarationType}:        _CGotoState38Action,
	{_CState1, CFunctionDefinitionType}:         _CGotoState31Action,
	{_CState6, CIdentifierToken}:                _CGotoState39Action,
	{_CState6, CLcurlToken}:                     _CGotoState40Action,
	{_CState12, CIdentifierToken}:               _CGotoState9Action,
	{_CState12, CLparamToken}:                   _CGotoState12Action,
	{_CState12, CMulToken}:                      _CGotoState13Action,
	{_CState12, CDeclaratorType}:                _CGotoState41Action,
	{_CState12, CDirectDeclaratorType}:          _CGotoState28Action,
	{_CState12, CPointerType}:                   _CGotoState32Action,
	{_CState13, CConstToken}:                    _CGotoState4Action,
	{_CState13, CVolatileToken}:                 _CGotoState24Action,
	{_CState13, CMulToken}:                      _CGotoState13Action,
	{_CState13, CTypeQualifierType}:             _CGotoState43Action,
	{_CState13, CPointerType}:                   _CGotoState42Action,
	{_CState13, CTypeQualifierListType}:         _CGotoState44Action,
	{_CState26, CIdentifierToken}:               _CGotoState9Action,
	{_CState26, CLparamToken}:                   _CGotoState12Action,
	{_CState26, CSemicolonToken}:                _CGotoState45Action,
	{_CState26, CMulToken}:                      _CGotoState13Action,
	{_CState26, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState26, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState26, CDeclaratorType}:                _CGotoState46Action,
	{_CState26, CDirectDeclaratorType}:          _CGotoState28Action,
	{_CState26, CPointerType}:                   _CGotoState32Action,
	{_CState27, CTypeNameToken}:                 _CGotoState20Action,
	{_CState27, CTypedefToken}:                  _CGotoState19Action,
	{_CState27, CExternToken}:                   _CGotoState7Action,
	{_CState27, CStaticToken}:                   _CGotoState17Action,
	{_CState27, CAutoToken}:                     _CGotoState2Action,
	{_CState27, CRegisterToken}:                 _CGotoState14Action,
	{_CState27, CCharToken}:                     _CGotoState3Action,
	{_CState27, CShortToken}:                    _CGotoState15Action,
	{_CState27, CIntToken}:                      _CGotoState10Action,
	{_CState27, CLongToken}:                     _CGotoState11Action,
	{_CState27, CSignedToken}:                   _CGotoState16Action,
	{_CState27, CUnsignedToken}:                 _CGotoState22Action,
	{_CState27, CFloatToken}:                    _CGotoState8Action,
	{_CState27, CDoubleToken}:                   _CGotoState5Action,
	{_CState27, CConstToken}:                    _CGotoState4Action,
	{_CState27, CVolatileToken}:                 _CGotoState24Action,
	{_CState27, CVoidToken}:                     _CGotoState23Action,
	{_CState27, CStructToken}:                   _CGotoState18Action,
	{_CState27, CUnionToken}:                    _CGotoState21Action,
	{_CState27, CEnumToken}:                     _CGotoState6Action,
	{_CState27, CLcurlToken}:                    _CGotoState49Action,
	{_CState27, CDeclarationType}:               _CGotoState51Action,
	{_CState27, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState27, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState27, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState27, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState27, CStructOrUnionType}:             _CGotoState34Action,
	{_CState27, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState27, CTypeQualifierType}:             _CGotoState36Action,
	{_CState27, CCompoundStatementType}:         _CGotoState50Action,
	{_CState27, CDeclarationListType}:           _CGotoState52Action,
	{_CState28, CLparamToken}:                   _CGotoState55Action,
	{_CState28, CLbraceToken}:                   _CGotoState54Action,
	{_CState32, CIdentifierToken}:               _CGotoState9Action,
	{_CState32, CLparamToken}:                   _CGotoState12Action,
	{_CState32, CDirectDeclaratorType}:          _CGotoState56Action,
	{_CState33, CTypeNameToken}:                 _CGotoState20Action,
	{_CState33, CTypedefToken}:                  _CGotoState19Action,
	{_CState33, CExternToken}:                   _CGotoState7Action,
	{_CState33, CStaticToken}:                   _CGotoState17Action,
	{_CState33, CAutoToken}:                     _CGotoState2Action,
	{_CState33, CRegisterToken}:                 _CGotoState14Action,
	{_CState33, CCharToken}:                     _CGotoState3Action,
	{_CState33, CShortToken}:                    _CGotoState15Action,
	{_CState33, CIntToken}:                      _CGotoState10Action,
	{_CState33, CLongToken}:                     _CGotoState11Action,
	{_CState33, CSignedToken}:                   _CGotoState16Action,
	{_CState33, CUnsignedToken}:                 _CGotoState22Action,
	{_CState33, CFloatToken}:                    _CGotoState8Action,
	{_CState33, CDoubleToken}:                   _CGotoState5Action,
	{_CState33, CConstToken}:                    _CGotoState4Action,
	{_CState33, CVolatileToken}:                 _CGotoState24Action,
	{_CState33, CVoidToken}:                     _CGotoState23Action,
	{_CState33, CStructToken}:                   _CGotoState18Action,
	{_CState33, CUnionToken}:                    _CGotoState21Action,
	{_CState33, CEnumToken}:                     _CGotoState6Action,
	{_CState33, CDeclarationSpecifiersType}:     _CGotoState57Action,
	{_CState33, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState33, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState33, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState33, CStructOrUnionType}:             _CGotoState34Action,
	{_CState33, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState33, CTypeQualifierType}:             _CGotoState36Action,
	{_CState34, CIdentifierToken}:               _CGotoState58Action,
	{_CState34, CLcurlToken}:                    _CGotoState59Action,
	{_CState36, CTypeNameToken}:                 _CGotoState20Action,
	{_CState36, CTypedefToken}:                  _CGotoState19Action,
	{_CState36, CExternToken}:                   _CGotoState7Action,
	{_CState36, CStaticToken}:                   _CGotoState17Action,
	{_CState36, CAutoToken}:                     _CGotoState2Action,
	{_CState36, CRegisterToken}:                 _CGotoState14Action,
	{_CState36, CCharToken}:                     _CGotoState3Action,
	{_CState36, CShortToken}:                    _CGotoState15Action,
	{_CState36, CIntToken}:                      _CGotoState10Action,
	{_CState36, CLongToken}:                     _CGotoState11Action,
	{_CState36, CSignedToken}:                   _CGotoState16Action,
	{_CState36, CUnsignedToken}:                 _CGotoState22Action,
	{_CState36, CFloatToken}:                    _CGotoState8Action,
	{_CState36, CDoubleToken}:                   _CGotoState5Action,
	{_CState36, CConstToken}:                    _CGotoState4Action,
	{_CState36, CVolatileToken}:                 _CGotoState24Action,
	{_CState36, CVoidToken}:                     _CGotoState23Action,
	{_CState36, CStructToken}:                   _CGotoState18Action,
	{_CState36, CUnionToken}:                    _CGotoState21Action,
	{_CState36, CEnumToken}:                     _CGotoState6Action,
	{_CState36, CDeclarationSpecifiersType}:     _CGotoState60Action,
	{_CState36, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState36, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState36, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState36, CStructOrUnionType}:             _CGotoState34Action,
	{_CState36, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState36, CTypeQualifierType}:             _CGotoState36Action,
	{_CState37, CTypeNameToken}:                 _CGotoState20Action,
	{_CState37, CTypedefToken}:                  _CGotoState19Action,
	{_CState37, CExternToken}:                   _CGotoState7Action,
	{_CState37, CStaticToken}:                   _CGotoState17Action,
	{_CState37, CAutoToken}:                     _CGotoState2Action,
	{_CState37, CRegisterToken}:                 _CGotoState14Action,
	{_CState37, CCharToken}:                     _CGotoState3Action,
	{_CState37, CShortToken}:                    _CGotoState15Action,
	{_CState37, CIntToken}:                      _CGotoState10Action,
	{_CState37, CLongToken}:                     _CGotoState11Action,
	{_CState37, CSignedToken}:                   _CGotoState16Action,
	{_CState37, CUnsignedToken}:                 _CGotoState22Action,
	{_CState37, CFloatToken}:                    _CGotoState8Action,
	{_CState37, CDoubleToken}:                   _CGotoState5Action,
	{_CState37, CConstToken}:                    _CGotoState4Action,
	{_CState37, CVolatileToken}:                 _CGotoState24Action,
	{_CState37, CVoidToken}:                     _CGotoState23Action,
	{_CState37, CStructToken}:                   _CGotoState18Action,
	{_CState37, CUnionToken}:                    _CGotoState21Action,
	{_CState37, CEnumToken}:                     _CGotoState6Action,
	{_CState37, CDeclarationSpecifiersType}:     _CGotoState61Action,
	{_CState37, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState37, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState37, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState37, CStructOrUnionType}:             _CGotoState34Action,
	{_CState37, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState37, CTypeQualifierType}:             _CGotoState36Action,
	{_CState39, CLcurlToken}:                    _CGotoState62Action,
	{_CState40, CIdentifierToken}:               _CGotoState63Action,
	{_CState40, CEnumeratorListType}:            _CGotoState65Action,
	{_CState40, CEnumeratorType}:                _CGotoState64Action,
	{_CState41, CRparamToken}:                   _CGotoState66Action,
	{_CState44, CConstToken}:                    _CGotoState4Action,
	{_CState44, CVolatileToken}:                 _CGotoState24Action,
	{_CState44, CMulToken}:                      _CGotoState13Action,
	{_CState44, CTypeQualifierType}:             _CGotoState68Action,
	{_CState44, CPointerType}:                   _CGotoState67Action,
	{_CState46, CTypeNameToken}:                 _CGotoState20Action,
	{_CState46, CTypedefToken}:                  _CGotoState19Action,
	{_CState46, CExternToken}:                   _CGotoState7Action,
	{_CState46, CStaticToken}:                   _CGotoState17Action,
	{_CState46, CAutoToken}:                     _CGotoState2Action,
	{_CState46, CRegisterToken}:                 _CGotoState14Action,
	{_CState46, CCharToken}:                     _CGotoState3Action,
	{_CState46, CShortToken}:                    _CGotoState15Action,
	{_CState46, CIntToken}:                      _CGotoState10Action,
	{_CState46, CLongToken}:                     _CGotoState11Action,
	{_CState46, CSignedToken}:                   _CGotoState16Action,
	{_CState46, CUnsignedToken}:                 _CGotoState22Action,
	{_CState46, CFloatToken}:                    _CGotoState8Action,
	{_CState46, CDoubleToken}:                   _CGotoState5Action,
	{_CState46, CConstToken}:                    _CGotoState4Action,
	{_CState46, CVolatileToken}:                 _CGotoState24Action,
	{_CState46, CVoidToken}:                     _CGotoState23Action,
	{_CState46, CStructToken}:                   _CGotoState18Action,
	{_CState46, CUnionToken}:                    _CGotoState21Action,
	{_CState46, CEnumToken}:                     _CGotoState6Action,
	{_CState46, CLcurlToken}:                    _CGotoState49Action,
	{_CState46, CEqToken}:                       _CGotoState69Action,
	{_CState46, CDeclarationType}:               _CGotoState51Action,
	{_CState46, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState46, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState46, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState46, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState46, CStructOrUnionType}:             _CGotoState34Action,
	{_CState46, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState46, CTypeQualifierType}:             _CGotoState36Action,
	{_CState46, CCompoundStatementType}:         _CGotoState70Action,
	{_CState46, CDeclarationListType}:           _CGotoState71Action,
	{_CState48, CSemicolonToken}:                _CGotoState73Action,
	{_CState48, CCommaToken}:                    _CGotoState72Action,
	{_CState49, CIdentifierToken}:               _CGotoState85Action,
	{_CState49, CConstantToken}:                 _CGotoState77Action,
	{_CState49, CStringLiteralToken}:            _CGotoState96Action,
	{_CState49, CSizeofToken}:                   _CGotoState95Action,
	{_CState49, CIncOpToken}:                    _CGotoState87Action,
	{_CState49, CDecOpToken}:                    _CGotoState79Action,
	{_CState49, CTypeNameToken}:                 _CGotoState20Action,
	{_CState49, CTypedefToken}:                  _CGotoState19Action,
	{_CState49, CExternToken}:                   _CGotoState7Action,
	{_CState49, CStaticToken}:                   _CGotoState17Action,
	{_CState49, CAutoToken}:                     _CGotoState2Action,
	{_CState49, CRegisterToken}:                 _CGotoState14Action,
	{_CState49, CCharToken}:                     _CGotoState3Action,
	{_CState49, CShortToken}:                    _CGotoState15Action,
	{_CState49, CIntToken}:                      _CGotoState10Action,
	{_CState49, CLongToken}:                     _CGotoState11Action,
	{_CState49, CSignedToken}:                   _CGotoState16Action,
	{_CState49, CUnsignedToken}:                 _CGotoState22Action,
	{_CState49, CFloatToken}:                    _CGotoState8Action,
	{_CState49, CDoubleToken}:                   _CGotoState5Action,
	{_CState49, CConstToken}:                    _CGotoState4Action,
	{_CState49, CVolatileToken}:                 _CGotoState24Action,
	{_CState49, CVoidToken}:                     _CGotoState23Action,
	{_CState49, CStructToken}:                   _CGotoState18Action,
	{_CState49, CUnionToken}:                    _CGotoState21Action,
	{_CState49, CEnumToken}:                     _CGotoState6Action,
	{_CState49, CCaseToken}:                     _CGotoState76Action,
	{_CState49, CDefaultToken}:                  _CGotoState80Action,
	{_CState49, CIfToken}:                       _CGotoState86Action,
	{_CState49, CSwitchToken}:                   _CGotoState97Action,
	{_CState49, CWhileToken}:                    _CGotoState99Action,
	{_CState49, CDoToken}:                       _CGotoState81Action,
	{_CState49, CForToken}:                      _CGotoState83Action,
	{_CState49, CGotoToken}:                     _CGotoState84Action,
	{_CState49, CContinueToken}:                 _CGotoState78Action,
	{_CState49, CBreakToken}:                    _CGotoState75Action,
	{_CState49, CReturnToken}:                   _CGotoState93Action,
	{_CState49, CLparamToken}:                   _CGotoState88Action,
	{_CState49, CLcurlToken}:                    _CGotoState49Action,
	{_CState49, CRcurlToken}:                    _CGotoState92Action,
	{_CState49, CSemicolonToken}:                _CGotoState94Action,
	{_CState49, CMulToken}:                      _CGotoState90Action,
	{_CState49, CMinusToken}:                    _CGotoState89Action,
	{_CState49, CPlusToken}:                     _CGotoState91Action,
	{_CState49, CAndToken}:                      _CGotoState74Action,
	{_CState49, CExclaimToken}:                  _CGotoState82Action,
	{_CState49, CTildaToken}:                    _CGotoState98Action,
	{_CState49, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState49, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState49, CUnaryExpressionType}:           _CGotoState125Action,
	{_CState49, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState49, CCastExpressionType}:            _CGotoState103Action,
	{_CState49, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState49, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState49, CShiftExpressionType}:           _CGotoState122Action,
	{_CState49, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState49, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState49, CAndExpressionType}:             _CGotoState101Action,
	{_CState49, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState49, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState49, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState49, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState49, CConditionalExpressionType}:     _CGotoState105Action,
	{_CState49, CAssignmentExpressionType}:      _CGotoState102Action,
	{_CState49, CExpressionType}:                _CGotoState109Action,
	{_CState49, CDeclarationType}:               _CGotoState51Action,
	{_CState49, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState49, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState49, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState49, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState49, CStructOrUnionType}:             _CGotoState34Action,
	{_CState49, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState49, CTypeQualifierType}:             _CGotoState36Action,
	{_CState49, CStatementType}:                 _CGotoState123Action,
	{_CState49, CLabeledStatementType}:          _CGotoState114Action,
	{_CState49, CCompoundStatementType}:         _CGotoState104Action,
	{_CState49, CDeclarationListType}:           _CGotoState106Action,
	{_CState49, CStatementListType}:             _CGotoState124Action,
	{_CState49, CExpressionStatementType}:       _CGotoState110Action,
	{_CState49, CSelectionStatementType}:        _CGotoState121Action,
	{_CState49, CIterationStatementType}:        _CGotoState112Action,
	{_CState49, CJumpStatementType}:             _CGotoState113Action,
	{_CState52, CTypeNameToken}:                 _CGotoState20Action,
	{_CState52, CTypedefToken}:                  _CGotoState19Action,
	{_CState52, CExternToken}:                   _CGotoState7Action,
	{_CState52, CStaticToken}:                   _CGotoState17Action,
	{_CState52, CAutoToken}:                     _CGotoState2Action,
	{_CState52, CRegisterToken}:                 _CGotoState14Action,
	{_CState52, CCharToken}:                     _CGotoState3Action,
	{_CState52, CShortToken}:                    _CGotoState15Action,
	{_CState52, CIntToken}:                      _CGotoState10Action,
	{_CState52, CLongToken}:                     _CGotoState11Action,
	{_CState52, CSignedToken}:                   _CGotoState16Action,
	{_CState52, CUnsignedToken}:                 _CGotoState22Action,
	{_CState52, CFloatToken}:                    _CGotoState8Action,
	{_CState52, CDoubleToken}:                   _CGotoState5Action,
	{_CState52, CConstToken}:                    _CGotoState4Action,
	{_CState52, CVolatileToken}:                 _CGotoState24Action,
	{_CState52, CVoidToken}:                     _CGotoState23Action,
	{_CState52, CStructToken}:                   _CGotoState18Action,
	{_CState52, CUnionToken}:                    _CGotoState21Action,
	{_CState52, CEnumToken}:                     _CGotoState6Action,
	{_CState52, CLcurlToken}:                    _CGotoState49Action,
	{_CState52, CDeclarationType}:               _CGotoState128Action,
	{_CState52, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState52, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState52, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState52, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState52, CStructOrUnionType}:             _CGotoState34Action,
	{_CState52, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState52, CTypeQualifierType}:             _CGotoState36Action,
	{_CState52, CCompoundStatementType}:         _CGotoState127Action,
	{_CState53, CIdentifierToken}:               _CGotoState9Action,
	{_CState53, CLparamToken}:                   _CGotoState12Action,
	{_CState53, CSemicolonToken}:                _CGotoState45Action,
	{_CState53, CMulToken}:                      _CGotoState13Action,
	{_CState53, CInitDeclaratorListType}:        _CGotoState48Action,
	{_CState53, CInitDeclaratorType}:            _CGotoState47Action,
	{_CState53, CDeclaratorType}:                _CGotoState129Action,
	{_CState53, CDirectDeclaratorType}:          _CGotoState28Action,
	{_CState53, CPointerType}:                   _CGotoState32Action,
	{_CState54, CIdentifierToken}:               _CGotoState130Action,
	{_CState54, CConstantToken}:                 _CGotoState77Action,
	{_CState54, CStringLiteralToken}:            _CGotoState96Action,
	{_CState54, CSizeofToken}:                   _CGotoState95Action,
	{_CState54, CIncOpToken}:                    _CGotoState87Action,
	{_CState54, CDecOpToken}:                    _CGotoState79Action,
	{_CState54, CLparamToken}:                   _CGotoState88Action,
	{_CState54, CRbraceToken}:                   _CGotoState131Action,
	{_CState54, CMulToken}:                      _CGotoState90Action,
	{_CState54, CMinusToken}:                    _CGotoState89Action,
	{_CState54, CPlusToken}:                     _CGotoState91Action,
	{_CState54, CAndToken}:                      _CGotoState74Action,
	{_CState54, CExclaimToken}:                  _CGotoState82Action,
	{_CState54, CTildaToken}:                    _CGotoState98Action,
	{_CState54, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState54, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState54, CUnaryExpressionType}:           _CGotoState134Action,
	{_CState54, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState54, CCastExpressionType}:            _CGotoState103Action,
	{_CState54, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState54, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState54, CShiftExpressionType}:           _CGotoState122Action,
	{_CState54, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState54, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState54, CAndExpressionType}:             _CGotoState101Action,
	{_CState54, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState54, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState54, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState54, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState54, CConditionalExpressionType}:     _CGotoState132Action,
	{_CState54, CConstantExpressionType}:        _CGotoState133Action,
	{_CState55, CIdentifierToken}:               _CGotoState135Action,
	{_CState55, CTypeNameToken}:                 _CGotoState20Action,
	{_CState55, CTypedefToken}:                  _CGotoState19Action,
	{_CState55, CExternToken}:                   _CGotoState7Action,
	{_CState55, CStaticToken}:                   _CGotoState17Action,
	{_CState55, CAutoToken}:                     _CGotoState2Action,
	{_CState55, CRegisterToken}:                 _CGotoState14Action,
	{_CState55, CCharToken}:                     _CGotoState3Action,
	{_CState55, CShortToken}:                    _CGotoState15Action,
	{_CState55, CIntToken}:                      _CGotoState10Action,
	{_CState55, CLongToken}:                     _CGotoState11Action,
	{_CState55, CSignedToken}:                   _CGotoState16Action,
	{_CState55, CUnsignedToken}:                 _CGotoState22Action,
	{_CState55, CFloatToken}:                    _CGotoState8Action,
	{_CState55, CDoubleToken}:                   _CGotoState5Action,
	{_CState55, CConstToken}:                    _CGotoState4Action,
	{_CState55, CVolatileToken}:                 _CGotoState24Action,
	{_CState55, CVoidToken}:                     _CGotoState23Action,
	{_CState55, CStructToken}:                   _CGotoState18Action,
	{_CState55, CUnionToken}:                    _CGotoState21Action,
	{_CState55, CEnumToken}:                     _CGotoState6Action,
	{_CState55, CRparamToken}:                   _CGotoState136Action,
	{_CState55, CDeclarationSpecifiersType}:     _CGotoState137Action,
	{_CState55, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState55, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState55, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState55, CStructOrUnionType}:             _CGotoState34Action,
	{_CState55, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState55, CTypeQualifierType}:             _CGotoState36Action,
	{_CState55, CParameterTypeListType}:         _CGotoState141Action,
	{_CState55, CParameterListType}:             _CGotoState140Action,
	{_CState55, CParameterDeclarationType}:      _CGotoState139Action,
	{_CState55, CIdentifierListType}:            _CGotoState138Action,
	{_CState56, CLparamToken}:                   _CGotoState55Action,
	{_CState56, CLbraceToken}:                   _CGotoState54Action,
	{_CState58, CLcurlToken}:                    _CGotoState142Action,
	{_CState59, CTypeNameToken}:                 _CGotoState20Action,
	{_CState59, CCharToken}:                     _CGotoState3Action,
	{_CState59, CShortToken}:                    _CGotoState15Action,
	{_CState59, CIntToken}:                      _CGotoState10Action,
	{_CState59, CLongToken}:                     _CGotoState11Action,
	{_CState59, CSignedToken}:                   _CGotoState16Action,
	{_CState59, CUnsignedToken}:                 _CGotoState22Action,
	{_CState59, CFloatToken}:                    _CGotoState8Action,
	{_CState59, CDoubleToken}:                   _CGotoState5Action,
	{_CState59, CConstToken}:                    _CGotoState4Action,
	{_CState59, CVolatileToken}:                 _CGotoState24Action,
	{_CState59, CVoidToken}:                     _CGotoState23Action,
	{_CState59, CStructToken}:                   _CGotoState18Action,
	{_CState59, CUnionToken}:                    _CGotoState21Action,
	{_CState59, CEnumToken}:                     _CGotoState6Action,
	{_CState59, CTypeSpecifierType}:             _CGotoState147Action,
	{_CState59, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState59, CStructOrUnionType}:             _CGotoState34Action,
	{_CState59, CStructDeclarationListType}:     _CGotoState145Action,
	{_CState59, CStructDeclarationType}:         _CGotoState144Action,
	{_CState59, CSpecifierQualifierListType}:    _CGotoState143Action,
	{_CState59, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState59, CTypeQualifierType}:             _CGotoState146Action,
	{_CState62, CIdentifierToken}:               _CGotoState63Action,
	{_CState62, CEnumeratorListType}:            _CGotoState148Action,
	{_CState62, CEnumeratorType}:                _CGotoState64Action,
	{_CState63, CEqToken}:                       _CGotoState149Action,
	{_CState65, CRcurlToken}:                    _CGotoState151Action,
	{_CState65, CCommaToken}:                    _CGotoState150Action,
	{_CState69, CIdentifierToken}:               _CGotoState130Action,
	{_CState69, CConstantToken}:                 _CGotoState77Action,
	{_CState69, CStringLiteralToken}:            _CGotoState96Action,
	{_CState69, CSizeofToken}:                   _CGotoState95Action,
	{_CState69, CIncOpToken}:                    _CGotoState87Action,
	{_CState69, CDecOpToken}:                    _CGotoState79Action,
	{_CState69, CLparamToken}:                   _CGotoState88Action,
	{_CState69, CLcurlToken}:                    _CGotoState152Action,
	{_CState69, CMulToken}:                      _CGotoState90Action,
	{_CState69, CMinusToken}:                    _CGotoState89Action,
	{_CState69, CPlusToken}:                     _CGotoState91Action,
	{_CState69, CAndToken}:                      _CGotoState74Action,
	{_CState69, CExclaimToken}:                  _CGotoState82Action,
	{_CState69, CTildaToken}:                    _CGotoState98Action,
	{_CState69, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState69, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState69, CUnaryExpressionType}:           _CGotoState125Action,
	{_CState69, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState69, CCastExpressionType}:            _CGotoState103Action,
	{_CState69, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState69, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState69, CShiftExpressionType}:           _CGotoState122Action,
	{_CState69, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState69, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState69, CAndExpressionType}:             _CGotoState101Action,
	{_CState69, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState69, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState69, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState69, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState69, CConditionalExpressionType}:     _CGotoState105Action,
	{_CState69, CAssignmentExpressionType}:      _CGotoState153Action,
	{_CState69, CInitializerType}:               _CGotoState154Action,
	{_CState71, CTypeNameToken}:                 _CGotoState20Action,
	{_CState71, CTypedefToken}:                  _CGotoState19Action,
	{_CState71, CExternToken}:                   _CGotoState7Action,
	{_CState71, CStaticToken}:                   _CGotoState17Action,
	{_CState71, CAutoToken}:                     _CGotoState2Action,
	{_CState71, CRegisterToken}:                 _CGotoState14Action,
	{_CState71, CCharToken}:                     _CGotoState3Action,
	{_CState71, CShortToken}:                    _CGotoState15Action,
	{_CState71, CIntToken}:                      _CGotoState10Action,
	{_CState71, CLongToken}:                     _CGotoState11Action,
	{_CState71, CSignedToken}:                   _CGotoState16Action,
	{_CState71, CUnsignedToken}:                 _CGotoState22Action,
	{_CState71, CFloatToken}:                    _CGotoState8Action,
	{_CState71, CDoubleToken}:                   _CGotoState5Action,
	{_CState71, CConstToken}:                    _CGotoState4Action,
	{_CState71, CVolatileToken}:                 _CGotoState24Action,
	{_CState71, CVoidToken}:                     _CGotoState23Action,
	{_CState71, CStructToken}:                   _CGotoState18Action,
	{_CState71, CUnionToken}:                    _CGotoState21Action,
	{_CState71, CEnumToken}:                     _CGotoState6Action,
	{_CState71, CLcurlToken}:                    _CGotoState49Action,
	{_CState71, CDeclarationType}:               _CGotoState128Action,
	{_CState71, CDeclarationSpecifiersType}:     _CGotoState53Action,
	{_CState71, CStorageClassSpecifierType}:     _CGotoState33Action,
	{_CState71, CTypeSpecifierType}:             _CGotoState37Action,
	{_CState71, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState71, CStructOrUnionType}:             _CGotoState34Action,
	{_CState71, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState71, CTypeQualifierType}:             _CGotoState36Action,
	{_CState71, CCompoundStatementType}:         _CGotoState155Action,
	{_CState72, CIdentifierToken}:               _CGotoState9Action,
	{_CState72, CLparamToken}:                   _CGotoState12Action,
	{_CState72, CMulToken}:                      _CGotoState13Action,
	{_CState72, CInitDeclaratorType}:            _CGotoState156Action,
	{_CState72, CDeclaratorType}:                _CGotoState129Action,
	{_CState72, CDirectDeclaratorType}:          _CGotoState28Action,
	{_CState72, CPointerType}:                   _CGotoState32Action,
	{_CState75, CSemicolonToken}:                _CGotoState157Action,
	{_CState76, CIdentifierToken}:               _CGotoState130Action,
	{_CState76, CConstantToken}:                 _CGotoState77Action,
	{_CState76, CStringLiteralToken}:            _CGotoState96Action,
	{_CState76, CSizeofToken}:                   _CGotoState95Action,
	{_CState76, CIncOpToken}:                    _CGotoState87Action,
	{_CState76, CDecOpToken}:                    _CGotoState79Action,
	{_CState76, CLparamToken}:                   _CGotoState88Action,
	{_CState76, CMulToken}:                      _CGotoState90Action,
	{_CState76, CMinusToken}:                    _CGotoState89Action,
	{_CState76, CPlusToken}:                     _CGotoState91Action,
	{_CState76, CAndToken}:                      _CGotoState74Action,
	{_CState76, CExclaimToken}:                  _CGotoState82Action,
	{_CState76, CTildaToken}:                    _CGotoState98Action,
	{_CState76, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState76, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState76, CUnaryExpressionType}:           _CGotoState134Action,
	{_CState76, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState76, CCastExpressionType}:            _CGotoState103Action,
	{_CState76, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState76, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState76, CShiftExpressionType}:           _CGotoState122Action,
	{_CState76, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState76, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState76, CAndExpressionType}:             _CGotoState101Action,
	{_CState76, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState76, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState76, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState76, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState76, CConditionalExpressionType}:     _CGotoState132Action,
	{_CState76, CConstantExpressionType}:        _CGotoState158Action,
	{_CState78, CSemicolonToken}:                _CGotoState159Action,
	{_CState79, CIdentifierToken}:               _CGotoState130Action,
	{_CState79, CConstantToken}:                 _CGotoState77Action,
	{_CState79, CStringLiteralToken}:            _CGotoState96Action,
	{_CState79, CSizeofToken}:                   _CGotoState95Action,
	{_CState79, CIncOpToken}:                    _CGotoState87Action,
	{_CState79, CDecOpToken}:                    _CGotoState79Action,
	{_CState79, CLparamToken}:                   _CGotoState160Action,
	{_CState79, CMulToken}:                      _CGotoState90Action,
	{_CState79, CMinusToken}:                    _CGotoState89Action,
	{_CState79, CPlusToken}:                     _CGotoState91Action,
	{_CState79, CAndToken}:                      _CGotoState74Action,
	{_CState79, CExclaimToken}:                  _CGotoState82Action,
	{_CState79, CTildaToken}:                    _CGotoState98Action,
	{_CState79, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState79, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState79, CUnaryExpressionType}:           _CGotoState161Action,
	{_CState79, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState80, CColonToken}:                    _CGotoState162Action,
	{_CState81, CIdentifierToken}:               _CGotoState85Action,
	{_CState81, CConstantToken}:                 _CGotoState77Action,
	{_CState81, CStringLiteralToken}:            _CGotoState96Action,
	{_CState81, CSizeofToken}:                   _CGotoState95Action,
	{_CState81, CIncOpToken}:                    _CGotoState87Action,
	{_CState81, CDecOpToken}:                    _CGotoState79Action,
	{_CState81, CCaseToken}:                     _CGotoState76Action,
	{_CState81, CDefaultToken}:                  _CGotoState80Action,
	{_CState81, CIfToken}:                       _CGotoState86Action,
	{_CState81, CSwitchToken}:                   _CGotoState97Action,
	{_CState81, CWhileToken}:                    _CGotoState99Action,
	{_CState81, CDoToken}:                       _CGotoState81Action,
	{_CState81, CForToken}:                      _CGotoState83Action,
	{_CState81, CGotoToken}:                     _CGotoState84Action,
	{_CState81, CContinueToken}:                 _CGotoState78Action,
	{_CState81, CBreakToken}:                    _CGotoState75Action,
	{_CState81, CReturnToken}:                   _CGotoState93Action,
	{_CState81, CLparamToken}:                   _CGotoState88Action,
	{_CState81, CLcurlToken}:                    _CGotoState49Action,
	{_CState81, CSemicolonToken}:                _CGotoState94Action,
	{_CState81, CMulToken}:                      _CGotoState90Action,
	{_CState81, CMinusToken}:                    _CGotoState89Action,
	{_CState81, CPlusToken}:                     _CGotoState91Action,
	{_CState81, CAndToken}:                      _CGotoState74Action,
	{_CState81, CExclaimToken}:                  _CGotoState82Action,
	{_CState81, CTildaToken}:                    _CGotoState98Action,
	{_CState81, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState81, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState81, CUnaryExpressionType}:           _CGotoState125Action,
	{_CState81, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState81, CCastExpressionType}:            _CGotoState103Action,
	{_CState81, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState81, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState81, CShiftExpressionType}:           _CGotoState122Action,
	{_CState81, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState81, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState81, CAndExpressionType}:             _CGotoState101Action,
	{_CState81, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState81, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState81, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState81, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState81, CConditionalExpressionType}:     _CGotoState105Action,
	{_CState81, CAssignmentExpressionType}:      _CGotoState102Action,
	{_CState81, CExpressionType}:                _CGotoState109Action,
	{_CState81, CStatementType}:                 _CGotoState163Action,
	{_CState81, CLabeledStatementType}:          _CGotoState114Action,
	{_CState81, CCompoundStatementType}:         _CGotoState104Action,
	{_CState81, CExpressionStatementType}:       _CGotoState110Action,
	{_CState81, CSelectionStatementType}:        _CGotoState121Action,
	{_CState81, CIterationStatementType}:        _CGotoState112Action,
	{_CState81, CJumpStatementType}:             _CGotoState113Action,
	{_CState83, CLparamToken}:                   _CGotoState164Action,
	{_CState84, CIdentifierToken}:               _CGotoState165Action,
	{_CState85, CColonToken}:                    _CGotoState166Action,
	{_CState86, CLparamToken}:                   _CGotoState167Action,
	{_CState87, CIdentifierToken}:               _CGotoState130Action,
	{_CState87, CConstantToken}:                 _CGotoState77Action,
	{_CState87, CStringLiteralToken}:            _CGotoState96Action,
	{_CState87, CSizeofToken}:                   _CGotoState95Action,
	{_CState87, CIncOpToken}:                    _CGotoState87Action,
	{_CState87, CDecOpToken}:                    _CGotoState79Action,
	{_CState87, CLparamToken}:                   _CGotoState160Action,
	{_CState87, CMulToken}:                      _CGotoState90Action,
	{_CState87, CMinusToken}:                    _CGotoState89Action,
	{_CState87, CPlusToken}:                     _CGotoState91Action,
	{_CState87, CAndToken}:                      _CGotoState74Action,
	{_CState87, CExclaimToken}:                  _CGotoState82Action,
	{_CState87, CTildaToken}:                    _CGotoState98Action,
	{_CState87, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState87, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState87, CUnaryExpressionType}:           _CGotoState168Action,
	{_CState87, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState88, CIdentifierToken}:               _CGotoState130Action,
	{_CState88, CConstantToken}:                 _CGotoState77Action,
	{_CState88, CStringLiteralToken}:            _CGotoState96Action,
	{_CState88, CSizeofToken}:                   _CGotoState95Action,
	{_CState88, CIncOpToken}:                    _CGotoState87Action,
	{_CState88, CDecOpToken}:                    _CGotoState79Action,
	{_CState88, CTypeNameToken}:                 _CGotoState20Action,
	{_CState88, CCharToken}:                     _CGotoState3Action,
	{_CState88, CShortToken}:                    _CGotoState15Action,
	{_CState88, CIntToken}:                      _CGotoState10Action,
	{_CState88, CLongToken}:                     _CGotoState11Action,
	{_CState88, CSignedToken}:                   _CGotoState16Action,
	{_CState88, CUnsignedToken}:                 _CGotoState22Action,
	{_CState88, CFloatToken}:                    _CGotoState8Action,
	{_CState88, CDoubleToken}:                   _CGotoState5Action,
	{_CState88, CConstToken}:                    _CGotoState4Action,
	{_CState88, CVolatileToken}:                 _CGotoState24Action,
	{_CState88, CVoidToken}:                     _CGotoState23Action,
	{_CState88, CStructToken}:                   _CGotoState18Action,
	{_CState88, CUnionToken}:                    _CGotoState21Action,
	{_CState88, CEnumToken}:                     _CGotoState6Action,
	{_CState88, CLparamToken}:                   _CGotoState88Action,
	{_CState88, CMulToken}:                      _CGotoState90Action,
	{_CState88, CMinusToken}:                    _CGotoState89Action,
	{_CState88, CPlusToken}:                     _CGotoState91Action,
	{_CState88, CAndToken}:                      _CGotoState74Action,
	{_CState88, CExclaimToken}:                  _CGotoState82Action,
	{_CState88, CTildaToken}:                    _CGotoState98Action,
	{_CState88, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState88, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState88, CUnaryExpressionType}:           _CGotoState125Action,
	{_CState88, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState88, CCastExpressionType}:            _CGotoState103Action,
	{_CState88, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState88, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState88, CShiftExpressionType}:           _CGotoState122Action,
	{_CState88, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState88, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState88, CAndExpressionType}:             _CGotoState101Action,
	{_CState88, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState88, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState88, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState88, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState88, CConditionalExpressionType}:     _CGotoState105Action,
	{_CState88, CAssignmentExpressionType}:      _CGotoState102Action,
	{_CState88, CExpressionType}:                _CGotoState169Action,
	{_CState88, CTypeSpecifierType}:             _CGotoState147Action,
	{_CState88, CStructOrUnionSpecifierType}:    _CGotoState35Action,
	{_CState88, CStructOrUnionType}:             _CGotoState34Action,
	{_CState88, CSpecifierQualifierListType}:    _CGotoState170Action,
	{_CState88, CEnumSpecifierType}:             _CGotoState29Action,
	{_CState88, CTypeQualifierType}:             _CGotoState146Action,
	{_CState88, CTypeNameType}:                  _CGotoState171Action,
	{_CState93, CIdentifierToken}:               _CGotoState130Action,
	{_CState93, CConstantToken}:                 _CGotoState77Action,
	{_CState93, CStringLiteralToken}:            _CGotoState96Action,
	{_CState93, CSizeofToken}:                   _CGotoState95Action,
	{_CState93, CIncOpToken}:                    _CGotoState87Action,
	{_CState93, CDecOpToken}:                    _CGotoState79Action,
	{_CState93, CLparamToken}:                   _CGotoState88Action,
	{_CState93, CSemicolonToken}:                _CGotoState172Action,
	{_CState93, CMulToken}:                      _CGotoState90Action,
	{_CState93, CMinusToken}:                    _CGotoState89Action,
	{_CState93, CPlusToken}:                     _CGotoState91Action,
	{_CState93, CAndToken}:                      _CGotoState74Action,
	{_CState93, CExclaimToken}:                  _CGotoState82Action,
	{_CState93, CTildaToken}:                    _CGotoState98Action,
	{_CState93, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState93, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState93, CUnaryExpressionType}:           _CGotoState125Action,
	{_CState93, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState93, CCastExpressionType}:            _CGotoState103Action,
	{_CState93, CMultiplicativeExpressionType}:  _CGotoState117Action,
	{_CState93, CAdditiveExpressionType}:        _CGotoState100Action,
	{_CState93, CShiftExpressionType}:           _CGotoState122Action,
	{_CState93, CRelationalExpressionType}:      _CGotoState120Action,
	{_CState93, CEqualityExpressionType}:        _CGotoState107Action,
	{_CState93, CAndExpressionType}:             _CGotoState101Action,
	{_CState93, CExclusiveOrExpressionType}:     _CGotoState108Action,
	{_CState93, CInclusiveOrExpressionType}:     _CGotoState111Action,
	{_CState93, CLogicalAndExpressionType}:      _CGotoState115Action,
	{_CState93, CLogicalOrExpressionType}:       _CGotoState116Action,
	{_CState93, CConditionalExpressionType}:     _CGotoState105Action,
	{_CState93, CAssignmentExpressionType}:      _CGotoState102Action,
	{_CState93, CExpressionType}:                _CGotoState173Action,
	{_CState95, CIdentifierToken}:               _CGotoState130Action,
	{_CState95, CConstantToken}:                 _CGotoState77Action,
	{_CState95, CStringLiteralToken}:            _CGotoState96Action,
	{_CState95, CSizeofToken}:                   _CGotoState95Action,
	{_CState95, CIncOpToken}:                    _CGotoState87Action,
	{_CState95, CDecOpToken}:                    _CGotoState79Action,
	{_CState95, CLparamToken}:                   _CGotoState174Action,
	{_CState95, CMulToken}:                      _CGotoState90Action,
	{_CState95, CMinusToken}:                    _CGotoState89Action,
	{_CState95, CPlusToken}:                     _CGotoState91Action,
	{_CState95, CAndToken}:                      _CGotoState74Action,
	{_CState95, CExclaimToken}:                  _CGotoState82Action,
	{_CState95, CTildaToken}:                    _CGotoState98Action,
	{_CState95, CPrimaryExpressionType}:         _CGotoState119Action,
	{_CState95, CPostfixExpressionType}:         _CGotoState118Action,
	{_CState95, CUnaryExpressionType}:           _CGotoState175Action,
	{_CState95, CUnaryOperatorType}:             _CGotoState126Action,
	{_CState97, CLparamToken}:                   _CGotoState176Action,
	{_CState99, CLparamToken}:                   _CGotoState177Action,
	{_CState100, CMinusToken}:                   _CGotoState178Action,
	{_CState100, CPlusToken}:                    _CGotoState179Action,
	{_CState101, CAndToken}:                     _CGotoState180Action,
	{_CState106, CIdentifierToken}:              _CGotoState85Action,
	{_CState106, CConstantToken}:                _CGotoState77Action,
	{_CState106, CStringLiteralToken}:           _CGotoState96Action,
	{_CState106, CSizeofToken}:                  _CGotoState95Action,
	{_CState106, CIncOpToken}:                   _CGotoState87Action,
	{_CState106, CDecOpToken}:                   _CGotoState79Action,
	{_CState106, CTypeNameToken}:                _CGotoState20Action,
	{_CState106, CTypedefToken}:                 _CGotoState19Action,
	{_CState106, CExternToken}:                  _CGotoState7Action,
	{_CState106, CStaticToken}:                  _CGotoState17Action,
	{_CState106, CAutoToken}:                    _CGotoState2Action,
	{_CState106, CRegisterToken}:                _CGotoState14Action,
	{_CState106, CCharToken}:                    _CGotoState3Action,
	{_CState106, CShortToken}:                   _CGotoState15Action,
	{_CState106, CIntToken}:                     _CGotoState10Action,
	{_CState106, CLongToken}:                    _CGotoState11Action,
	{_CState106, CSignedToken}:                  _CGotoState16Action,
	{_CState106, CUnsignedToken}:                _CGotoState22Action,
	{_CState106, CFloatToken}:                   _CGotoState8Action,
	{_CState106, CDoubleToken}:                  _CGotoState5Action,
	{_CState106, CConstToken}:                   _CGotoState4Action,
	{_CState106, CVolatileToken}:                _CGotoState24Action,
	{_CState106, CVoidToken}:                    _CGotoState23Action,
	{_CState106, CStructToken}:                  _CGotoState18Action,
	{_CState106, CUnionToken}:                   _CGotoState21Action,
	{_CState106, CEnumToken}:                    _CGotoState6Action,
	{_CState106, CCaseToken}:                    _CGotoState76Action,
	{_CState106, CDefaultToken}:                 _CGotoState80Action,
	{_CState106, CIfToken}:                      _CGotoState86Action,
	{_CState106, CSwitchToken}:                  _CGotoState97Action,
	{_CState106, CWhileToken}:                   _CGotoState99Action,
	{_CState106, CDoToken}:                      _CGotoState81Action,
	{_CState106, CForToken}:                     _CGotoState83Action,
	{_CState106, CGotoToken}:                    _CGotoState84Action,
	{_CState106, CContinueToken}:                _CGotoState78Action,
	{_CState106, CBreakToken}:                   _CGotoState75Action,
	{_CState106, CReturnToken}:                  _CGotoState93Action,
	{_CState106, CLparamToken}:                  _CGotoState88Action,
	{_CState106, CLcurlToken}:                   _CGotoState49Action,
	{_CState106, CRcurlToken}:                   _CGotoState181Action,
	{_CState106, CSemicolonToken}:               _CGotoState94Action,
	{_CState106, CMulToken}:                     _CGotoState90Action,
	{_CState106, CMinusToken}:                   _CGotoState89Action,
	{_CState106, CPlusToken}:                    _CGotoState91Action,
	{_CState106, CAndToken}:                     _CGotoState74Action,
	{_CState106, CExclaimToken}:                 _CGotoState82Action,
	{_CState106, CTildaToken}:                   _CGotoState98Action,
	{_CState106, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState106, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState106, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState106, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState106, CCastExpressionType}:           _CGotoState103Action,
	{_CState106, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState106, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState106, CShiftExpressionType}:          _CGotoState122Action,
	{_CState106, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState106, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState106, CAndExpressionType}:            _CGotoState101Action,
	{_CState106, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState106, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState106, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState106, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState106, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState106, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState106, CExpressionType}:               _CGotoState109Action,
	{_CState106, CDeclarationType}:              _CGotoState128Action,
	{_CState106, CDeclarationSpecifiersType}:    _CGotoState53Action,
	{_CState106, CStorageClassSpecifierType}:    _CGotoState33Action,
	{_CState106, CTypeSpecifierType}:            _CGotoState37Action,
	{_CState106, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState106, CStructOrUnionType}:            _CGotoState34Action,
	{_CState106, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState106, CTypeQualifierType}:            _CGotoState36Action,
	{_CState106, CStatementType}:                _CGotoState123Action,
	{_CState106, CLabeledStatementType}:         _CGotoState114Action,
	{_CState106, CCompoundStatementType}:        _CGotoState104Action,
	{_CState106, CStatementListType}:            _CGotoState182Action,
	{_CState106, CExpressionStatementType}:      _CGotoState110Action,
	{_CState106, CSelectionStatementType}:       _CGotoState121Action,
	{_CState106, CIterationStatementType}:       _CGotoState112Action,
	{_CState106, CJumpStatementType}:            _CGotoState113Action,
	{_CState107, CEqOpToken}:                    _CGotoState183Action,
	{_CState107, CNeOpToken}:                    _CGotoState184Action,
	{_CState108, CHatToken}:                     _CGotoState185Action,
	{_CState109, CSemicolonToken}:               _CGotoState187Action,
	{_CState109, CCommaToken}:                   _CGotoState186Action,
	{_CState111, COrToken}:                      _CGotoState188Action,
	{_CState115, CAndOpToken}:                   _CGotoState189Action,
	{_CState116, COrOpToken}:                    _CGotoState190Action,
	{_CState116, CQuestionToken}:                _CGotoState191Action,
	{_CState117, CMulToken}:                     _CGotoState194Action,
	{_CState117, CDivToken}:                     _CGotoState192Action,
	{_CState117, CModToken}:                     _CGotoState193Action,
	{_CState118, CPtrOpToken}:                   _CGotoState200Action,
	{_CState118, CIncOpToken}:                   _CGotoState197Action,
	{_CState118, CDecOpToken}:                   _CGotoState195Action,
	{_CState118, CLparamToken}:                  _CGotoState199Action,
	{_CState118, CLbraceToken}:                  _CGotoState198Action,
	{_CState118, CDotToken}:                     _CGotoState196Action,
	{_CState120, CLeOpToken}:                    _CGotoState203Action,
	{_CState120, CGeOpToken}:                    _CGotoState201Action,
	{_CState120, CLtToken}:                      _CGotoState204Action,
	{_CState120, CGtToken}:                      _CGotoState202Action,
	{_CState122, CLeftOpToken}:                  _CGotoState205Action,
	{_CState122, CRightOpToken}:                 _CGotoState206Action,
	{_CState124, CIdentifierToken}:              _CGotoState85Action,
	{_CState124, CConstantToken}:                _CGotoState77Action,
	{_CState124, CStringLiteralToken}:           _CGotoState96Action,
	{_CState124, CSizeofToken}:                  _CGotoState95Action,
	{_CState124, CIncOpToken}:                   _CGotoState87Action,
	{_CState124, CDecOpToken}:                   _CGotoState79Action,
	{_CState124, CCaseToken}:                    _CGotoState76Action,
	{_CState124, CDefaultToken}:                 _CGotoState80Action,
	{_CState124, CIfToken}:                      _CGotoState86Action,
	{_CState124, CSwitchToken}:                  _CGotoState97Action,
	{_CState124, CWhileToken}:                   _CGotoState99Action,
	{_CState124, CDoToken}:                      _CGotoState81Action,
	{_CState124, CForToken}:                     _CGotoState83Action,
	{_CState124, CGotoToken}:                    _CGotoState84Action,
	{_CState124, CContinueToken}:                _CGotoState78Action,
	{_CState124, CBreakToken}:                   _CGotoState75Action,
	{_CState124, CReturnToken}:                  _CGotoState93Action,
	{_CState124, CLparamToken}:                  _CGotoState88Action,
	{_CState124, CLcurlToken}:                   _CGotoState49Action,
	{_CState124, CRcurlToken}:                   _CGotoState207Action,
	{_CState124, CSemicolonToken}:               _CGotoState94Action,
	{_CState124, CMulToken}:                     _CGotoState90Action,
	{_CState124, CMinusToken}:                   _CGotoState89Action,
	{_CState124, CPlusToken}:                    _CGotoState91Action,
	{_CState124, CAndToken}:                     _CGotoState74Action,
	{_CState124, CExclaimToken}:                 _CGotoState82Action,
	{_CState124, CTildaToken}:                   _CGotoState98Action,
	{_CState124, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState124, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState124, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState124, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState124, CCastExpressionType}:           _CGotoState103Action,
	{_CState124, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState124, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState124, CShiftExpressionType}:          _CGotoState122Action,
	{_CState124, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState124, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState124, CAndExpressionType}:            _CGotoState101Action,
	{_CState124, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState124, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState124, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState124, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState124, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState124, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState124, CExpressionType}:               _CGotoState109Action,
	{_CState124, CStatementType}:                _CGotoState208Action,
	{_CState124, CLabeledStatementType}:         _CGotoState114Action,
	{_CState124, CCompoundStatementType}:        _CGotoState104Action,
	{_CState124, CExpressionStatementType}:      _CGotoState110Action,
	{_CState124, CSelectionStatementType}:       _CGotoState121Action,
	{_CState124, CIterationStatementType}:       _CGotoState112Action,
	{_CState124, CJumpStatementType}:            _CGotoState113Action,
	{_CState125, CMulAssignToken}:               _CGotoState215Action,
	{_CState125, CDivAssignToken}:               _CGotoState211Action,
	{_CState125, CModAssignToken}:               _CGotoState214Action,
	{_CState125, CAddAssignToken}:               _CGotoState209Action,
	{_CState125, CSubAssignToken}:               _CGotoState218Action,
	{_CState125, CLeftAssignToken}:              _CGotoState213Action,
	{_CState125, CRightAssignToken}:             _CGotoState217Action,
	{_CState125, CAndAssignToken}:               _CGotoState210Action,
	{_CState125, CXorAssignToken}:               _CGotoState219Action,
	{_CState125, COrAssignToken}:                _CGotoState216Action,
	{_CState125, CEqToken}:                      _CGotoState212Action,
	{_CState125, CAssignmentOperatorType}:       _CGotoState220Action,
	{_CState126, CIdentifierToken}:              _CGotoState130Action,
	{_CState126, CConstantToken}:                _CGotoState77Action,
	{_CState126, CStringLiteralToken}:           _CGotoState96Action,
	{_CState126, CSizeofToken}:                  _CGotoState95Action,
	{_CState126, CIncOpToken}:                   _CGotoState87Action,
	{_CState126, CDecOpToken}:                   _CGotoState79Action,
	{_CState126, CLparamToken}:                  _CGotoState88Action,
	{_CState126, CMulToken}:                     _CGotoState90Action,
	{_CState126, CMinusToken}:                   _CGotoState89Action,
	{_CState126, CPlusToken}:                    _CGotoState91Action,
	{_CState126, CAndToken}:                     _CGotoState74Action,
	{_CState126, CExclaimToken}:                 _CGotoState82Action,
	{_CState126, CTildaToken}:                   _CGotoState98Action,
	{_CState126, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState126, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState126, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState126, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState126, CCastExpressionType}:           _CGotoState221Action,
	{_CState129, CEqToken}:                      _CGotoState69Action,
	{_CState133, CRbraceToken}:                  _CGotoState222Action,
	{_CState137, CIdentifierToken}:              _CGotoState9Action,
	{_CState137, CLparamToken}:                  _CGotoState224Action,
	{_CState137, CLbraceToken}:                  _CGotoState223Action,
	{_CState137, CMulToken}:                     _CGotoState13Action,
	{_CState137, CDeclaratorType}:               _CGotoState226Action,
	{_CState137, CDirectDeclaratorType}:         _CGotoState28Action,
	{_CState137, CPointerType}:                  _CGotoState228Action,
	{_CState137, CAbstractDeclaratorType}:       _CGotoState225Action,
	{_CState137, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState138, CRparamToken}:                  _CGotoState230Action,
	{_CState138, CCommaToken}:                   _CGotoState229Action,
	{_CState140, CCommaToken}:                   _CGotoState231Action,
	{_CState141, CRparamToken}:                  _CGotoState232Action,
	{_CState142, CTypeNameToken}:                _CGotoState20Action,
	{_CState142, CCharToken}:                    _CGotoState3Action,
	{_CState142, CShortToken}:                   _CGotoState15Action,
	{_CState142, CIntToken}:                     _CGotoState10Action,
	{_CState142, CLongToken}:                    _CGotoState11Action,
	{_CState142, CSignedToken}:                  _CGotoState16Action,
	{_CState142, CUnsignedToken}:                _CGotoState22Action,
	{_CState142, CFloatToken}:                   _CGotoState8Action,
	{_CState142, CDoubleToken}:                  _CGotoState5Action,
	{_CState142, CConstToken}:                   _CGotoState4Action,
	{_CState142, CVolatileToken}:                _CGotoState24Action,
	{_CState142, CVoidToken}:                    _CGotoState23Action,
	{_CState142, CStructToken}:                  _CGotoState18Action,
	{_CState142, CUnionToken}:                   _CGotoState21Action,
	{_CState142, CEnumToken}:                    _CGotoState6Action,
	{_CState142, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState142, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState142, CStructOrUnionType}:            _CGotoState34Action,
	{_CState142, CStructDeclarationListType}:    _CGotoState233Action,
	{_CState142, CStructDeclarationType}:        _CGotoState144Action,
	{_CState142, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState142, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState142, CTypeQualifierType}:            _CGotoState146Action,
	{_CState143, CIdentifierToken}:              _CGotoState9Action,
	{_CState143, CLparamToken}:                  _CGotoState12Action,
	{_CState143, CColonToken}:                   _CGotoState234Action,
	{_CState143, CMulToken}:                     _CGotoState13Action,
	{_CState143, CStructDeclaratorListType}:     _CGotoState237Action,
	{_CState143, CStructDeclaratorType}:         _CGotoState236Action,
	{_CState143, CDeclaratorType}:               _CGotoState235Action,
	{_CState143, CDirectDeclaratorType}:         _CGotoState28Action,
	{_CState143, CPointerType}:                  _CGotoState32Action,
	{_CState145, CTypeNameToken}:                _CGotoState20Action,
	{_CState145, CCharToken}:                    _CGotoState3Action,
	{_CState145, CShortToken}:                   _CGotoState15Action,
	{_CState145, CIntToken}:                     _CGotoState10Action,
	{_CState145, CLongToken}:                    _CGotoState11Action,
	{_CState145, CSignedToken}:                  _CGotoState16Action,
	{_CState145, CUnsignedToken}:                _CGotoState22Action,
	{_CState145, CFloatToken}:                   _CGotoState8Action,
	{_CState145, CDoubleToken}:                  _CGotoState5Action,
	{_CState145, CConstToken}:                   _CGotoState4Action,
	{_CState145, CVolatileToken}:                _CGotoState24Action,
	{_CState145, CVoidToken}:                    _CGotoState23Action,
	{_CState145, CStructToken}:                  _CGotoState18Action,
	{_CState145, CUnionToken}:                   _CGotoState21Action,
	{_CState145, CEnumToken}:                    _CGotoState6Action,
	{_CState145, CRcurlToken}:                   _CGotoState238Action,
	{_CState145, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState145, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState145, CStructOrUnionType}:            _CGotoState34Action,
	{_CState145, CStructDeclarationType}:        _CGotoState239Action,
	{_CState145, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState145, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState145, CTypeQualifierType}:            _CGotoState146Action,
	{_CState146, CTypeNameToken}:                _CGotoState20Action,
	{_CState146, CCharToken}:                    _CGotoState3Action,
	{_CState146, CShortToken}:                   _CGotoState15Action,
	{_CState146, CIntToken}:                     _CGotoState10Action,
	{_CState146, CLongToken}:                    _CGotoState11Action,
	{_CState146, CSignedToken}:                  _CGotoState16Action,
	{_CState146, CUnsignedToken}:                _CGotoState22Action,
	{_CState146, CFloatToken}:                   _CGotoState8Action,
	{_CState146, CDoubleToken}:                  _CGotoState5Action,
	{_CState146, CConstToken}:                   _CGotoState4Action,
	{_CState146, CVolatileToken}:                _CGotoState24Action,
	{_CState146, CVoidToken}:                    _CGotoState23Action,
	{_CState146, CStructToken}:                  _CGotoState18Action,
	{_CState146, CUnionToken}:                   _CGotoState21Action,
	{_CState146, CEnumToken}:                    _CGotoState6Action,
	{_CState146, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState146, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState146, CStructOrUnionType}:            _CGotoState34Action,
	{_CState146, CSpecifierQualifierListType}:   _CGotoState240Action,
	{_CState146, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState146, CTypeQualifierType}:            _CGotoState146Action,
	{_CState147, CTypeNameToken}:                _CGotoState20Action,
	{_CState147, CCharToken}:                    _CGotoState3Action,
	{_CState147, CShortToken}:                   _CGotoState15Action,
	{_CState147, CIntToken}:                     _CGotoState10Action,
	{_CState147, CLongToken}:                    _CGotoState11Action,
	{_CState147, CSignedToken}:                  _CGotoState16Action,
	{_CState147, CUnsignedToken}:                _CGotoState22Action,
	{_CState147, CFloatToken}:                   _CGotoState8Action,
	{_CState147, CDoubleToken}:                  _CGotoState5Action,
	{_CState147, CConstToken}:                   _CGotoState4Action,
	{_CState147, CVolatileToken}:                _CGotoState24Action,
	{_CState147, CVoidToken}:                    _CGotoState23Action,
	{_CState147, CStructToken}:                  _CGotoState18Action,
	{_CState147, CUnionToken}:                   _CGotoState21Action,
	{_CState147, CEnumToken}:                    _CGotoState6Action,
	{_CState147, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState147, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState147, CStructOrUnionType}:            _CGotoState34Action,
	{_CState147, CSpecifierQualifierListType}:   _CGotoState241Action,
	{_CState147, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState147, CTypeQualifierType}:            _CGotoState146Action,
	{_CState148, CRcurlToken}:                   _CGotoState242Action,
	{_CState148, CCommaToken}:                   _CGotoState150Action,
	{_CState149, CIdentifierToken}:              _CGotoState130Action,
	{_CState149, CConstantToken}:                _CGotoState77Action,
	{_CState149, CStringLiteralToken}:           _CGotoState96Action,
	{_CState149, CSizeofToken}:                  _CGotoState95Action,
	{_CState149, CIncOpToken}:                   _CGotoState87Action,
	{_CState149, CDecOpToken}:                   _CGotoState79Action,
	{_CState149, CLparamToken}:                  _CGotoState88Action,
	{_CState149, CMulToken}:                     _CGotoState90Action,
	{_CState149, CMinusToken}:                   _CGotoState89Action,
	{_CState149, CPlusToken}:                    _CGotoState91Action,
	{_CState149, CAndToken}:                     _CGotoState74Action,
	{_CState149, CExclaimToken}:                 _CGotoState82Action,
	{_CState149, CTildaToken}:                   _CGotoState98Action,
	{_CState149, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState149, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState149, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState149, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState149, CCastExpressionType}:           _CGotoState103Action,
	{_CState149, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState149, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState149, CShiftExpressionType}:          _CGotoState122Action,
	{_CState149, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState149, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState149, CAndExpressionType}:            _CGotoState101Action,
	{_CState149, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState149, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState149, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState149, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState149, CConditionalExpressionType}:    _CGotoState132Action,
	{_CState149, CConstantExpressionType}:       _CGotoState243Action,
	{_CState150, CIdentifierToken}:              _CGotoState63Action,
	{_CState150, CEnumeratorType}:               _CGotoState244Action,
	{_CState152, CIdentifierToken}:              _CGotoState130Action,
	{_CState152, CConstantToken}:                _CGotoState77Action,
	{_CState152, CStringLiteralToken}:           _CGotoState96Action,
	{_CState152, CSizeofToken}:                  _CGotoState95Action,
	{_CState152, CIncOpToken}:                   _CGotoState87Action,
	{_CState152, CDecOpToken}:                   _CGotoState79Action,
	{_CState152, CLparamToken}:                  _CGotoState88Action,
	{_CState152, CLcurlToken}:                   _CGotoState152Action,
	{_CState152, CMulToken}:                     _CGotoState90Action,
	{_CState152, CMinusToken}:                   _CGotoState89Action,
	{_CState152, CPlusToken}:                    _CGotoState91Action,
	{_CState152, CAndToken}:                     _CGotoState74Action,
	{_CState152, CExclaimToken}:                 _CGotoState82Action,
	{_CState152, CTildaToken}:                   _CGotoState98Action,
	{_CState152, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState152, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState152, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState152, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState152, CCastExpressionType}:           _CGotoState103Action,
	{_CState152, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState152, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState152, CShiftExpressionType}:          _CGotoState122Action,
	{_CState152, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState152, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState152, CAndExpressionType}:            _CGotoState101Action,
	{_CState152, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState152, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState152, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState152, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState152, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState152, CAssignmentExpressionType}:     _CGotoState153Action,
	{_CState152, CInitializerType}:              _CGotoState245Action,
	{_CState152, CInitializerListType}:          _CGotoState246Action,
	{_CState158, CColonToken}:                   _CGotoState247Action,
	{_CState160, CIdentifierToken}:              _CGotoState130Action,
	{_CState160, CConstantToken}:                _CGotoState77Action,
	{_CState160, CStringLiteralToken}:           _CGotoState96Action,
	{_CState160, CSizeofToken}:                  _CGotoState95Action,
	{_CState160, CIncOpToken}:                   _CGotoState87Action,
	{_CState160, CDecOpToken}:                   _CGotoState79Action,
	{_CState160, CLparamToken}:                  _CGotoState88Action,
	{_CState160, CMulToken}:                     _CGotoState90Action,
	{_CState160, CMinusToken}:                   _CGotoState89Action,
	{_CState160, CPlusToken}:                    _CGotoState91Action,
	{_CState160, CAndToken}:                     _CGotoState74Action,
	{_CState160, CExclaimToken}:                 _CGotoState82Action,
	{_CState160, CTildaToken}:                   _CGotoState98Action,
	{_CState160, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState160, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState160, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState160, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState160, CCastExpressionType}:           _CGotoState103Action,
	{_CState160, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState160, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState160, CShiftExpressionType}:          _CGotoState122Action,
	{_CState160, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState160, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState160, CAndExpressionType}:            _CGotoState101Action,
	{_CState160, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState160, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState160, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState160, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState160, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState160, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState160, CExpressionType}:               _CGotoState169Action,
	{_CState162, CIdentifierToken}:              _CGotoState85Action,
	{_CState162, CConstantToken}:                _CGotoState77Action,
	{_CState162, CStringLiteralToken}:           _CGotoState96Action,
	{_CState162, CSizeofToken}:                  _CGotoState95Action,
	{_CState162, CIncOpToken}:                   _CGotoState87Action,
	{_CState162, CDecOpToken}:                   _CGotoState79Action,
	{_CState162, CCaseToken}:                    _CGotoState76Action,
	{_CState162, CDefaultToken}:                 _CGotoState80Action,
	{_CState162, CIfToken}:                      _CGotoState86Action,
	{_CState162, CSwitchToken}:                  _CGotoState97Action,
	{_CState162, CWhileToken}:                   _CGotoState99Action,
	{_CState162, CDoToken}:                      _CGotoState81Action,
	{_CState162, CForToken}:                     _CGotoState83Action,
	{_CState162, CGotoToken}:                    _CGotoState84Action,
	{_CState162, CContinueToken}:                _CGotoState78Action,
	{_CState162, CBreakToken}:                   _CGotoState75Action,
	{_CState162, CReturnToken}:                  _CGotoState93Action,
	{_CState162, CLparamToken}:                  _CGotoState88Action,
	{_CState162, CLcurlToken}:                   _CGotoState49Action,
	{_CState162, CSemicolonToken}:               _CGotoState94Action,
	{_CState162, CMulToken}:                     _CGotoState90Action,
	{_CState162, CMinusToken}:                   _CGotoState89Action,
	{_CState162, CPlusToken}:                    _CGotoState91Action,
	{_CState162, CAndToken}:                     _CGotoState74Action,
	{_CState162, CExclaimToken}:                 _CGotoState82Action,
	{_CState162, CTildaToken}:                   _CGotoState98Action,
	{_CState162, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState162, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState162, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState162, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState162, CCastExpressionType}:           _CGotoState103Action,
	{_CState162, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState162, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState162, CShiftExpressionType}:          _CGotoState122Action,
	{_CState162, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState162, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState162, CAndExpressionType}:            _CGotoState101Action,
	{_CState162, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState162, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState162, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState162, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState162, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState162, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState162, CExpressionType}:               _CGotoState109Action,
	{_CState162, CStatementType}:                _CGotoState248Action,
	{_CState162, CLabeledStatementType}:         _CGotoState114Action,
	{_CState162, CCompoundStatementType}:        _CGotoState104Action,
	{_CState162, CExpressionStatementType}:      _CGotoState110Action,
	{_CState162, CSelectionStatementType}:       _CGotoState121Action,
	{_CState162, CIterationStatementType}:       _CGotoState112Action,
	{_CState162, CJumpStatementType}:            _CGotoState113Action,
	{_CState163, CWhileToken}:                   _CGotoState249Action,
	{_CState164, CIdentifierToken}:              _CGotoState130Action,
	{_CState164, CConstantToken}:                _CGotoState77Action,
	{_CState164, CStringLiteralToken}:           _CGotoState96Action,
	{_CState164, CSizeofToken}:                  _CGotoState95Action,
	{_CState164, CIncOpToken}:                   _CGotoState87Action,
	{_CState164, CDecOpToken}:                   _CGotoState79Action,
	{_CState164, CLparamToken}:                  _CGotoState88Action,
	{_CState164, CSemicolonToken}:               _CGotoState94Action,
	{_CState164, CMulToken}:                     _CGotoState90Action,
	{_CState164, CMinusToken}:                   _CGotoState89Action,
	{_CState164, CPlusToken}:                    _CGotoState91Action,
	{_CState164, CAndToken}:                     _CGotoState74Action,
	{_CState164, CExclaimToken}:                 _CGotoState82Action,
	{_CState164, CTildaToken}:                   _CGotoState98Action,
	{_CState164, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState164, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState164, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState164, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState164, CCastExpressionType}:           _CGotoState103Action,
	{_CState164, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState164, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState164, CShiftExpressionType}:          _CGotoState122Action,
	{_CState164, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState164, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState164, CAndExpressionType}:            _CGotoState101Action,
	{_CState164, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState164, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState164, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState164, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState164, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState164, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState164, CExpressionType}:               _CGotoState109Action,
	{_CState164, CExpressionStatementType}:      _CGotoState250Action,
	{_CState165, CSemicolonToken}:               _CGotoState251Action,
	{_CState166, CIdentifierToken}:              _CGotoState85Action,
	{_CState166, CConstantToken}:                _CGotoState77Action,
	{_CState166, CStringLiteralToken}:           _CGotoState96Action,
	{_CState166, CSizeofToken}:                  _CGotoState95Action,
	{_CState166, CIncOpToken}:                   _CGotoState87Action,
	{_CState166, CDecOpToken}:                   _CGotoState79Action,
	{_CState166, CCaseToken}:                    _CGotoState76Action,
	{_CState166, CDefaultToken}:                 _CGotoState80Action,
	{_CState166, CIfToken}:                      _CGotoState86Action,
	{_CState166, CSwitchToken}:                  _CGotoState97Action,
	{_CState166, CWhileToken}:                   _CGotoState99Action,
	{_CState166, CDoToken}:                      _CGotoState81Action,
	{_CState166, CForToken}:                     _CGotoState83Action,
	{_CState166, CGotoToken}:                    _CGotoState84Action,
	{_CState166, CContinueToken}:                _CGotoState78Action,
	{_CState166, CBreakToken}:                   _CGotoState75Action,
	{_CState166, CReturnToken}:                  _CGotoState93Action,
	{_CState166, CLparamToken}:                  _CGotoState88Action,
	{_CState166, CLcurlToken}:                   _CGotoState49Action,
	{_CState166, CSemicolonToken}:               _CGotoState94Action,
	{_CState166, CMulToken}:                     _CGotoState90Action,
	{_CState166, CMinusToken}:                   _CGotoState89Action,
	{_CState166, CPlusToken}:                    _CGotoState91Action,
	{_CState166, CAndToken}:                     _CGotoState74Action,
	{_CState166, CExclaimToken}:                 _CGotoState82Action,
	{_CState166, CTildaToken}:                   _CGotoState98Action,
	{_CState166, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState166, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState166, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState166, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState166, CCastExpressionType}:           _CGotoState103Action,
	{_CState166, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState166, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState166, CShiftExpressionType}:          _CGotoState122Action,
	{_CState166, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState166, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState166, CAndExpressionType}:            _CGotoState101Action,
	{_CState166, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState166, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState166, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState166, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState166, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState166, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState166, CExpressionType}:               _CGotoState109Action,
	{_CState166, CStatementType}:                _CGotoState252Action,
	{_CState166, CLabeledStatementType}:         _CGotoState114Action,
	{_CState166, CCompoundStatementType}:        _CGotoState104Action,
	{_CState166, CExpressionStatementType}:      _CGotoState110Action,
	{_CState166, CSelectionStatementType}:       _CGotoState121Action,
	{_CState166, CIterationStatementType}:       _CGotoState112Action,
	{_CState166, CJumpStatementType}:            _CGotoState113Action,
	{_CState167, CIdentifierToken}:              _CGotoState130Action,
	{_CState167, CConstantToken}:                _CGotoState77Action,
	{_CState167, CStringLiteralToken}:           _CGotoState96Action,
	{_CState167, CSizeofToken}:                  _CGotoState95Action,
	{_CState167, CIncOpToken}:                   _CGotoState87Action,
	{_CState167, CDecOpToken}:                   _CGotoState79Action,
	{_CState167, CLparamToken}:                  _CGotoState88Action,
	{_CState167, CMulToken}:                     _CGotoState90Action,
	{_CState167, CMinusToken}:                   _CGotoState89Action,
	{_CState167, CPlusToken}:                    _CGotoState91Action,
	{_CState167, CAndToken}:                     _CGotoState74Action,
	{_CState167, CExclaimToken}:                 _CGotoState82Action,
	{_CState167, CTildaToken}:                   _CGotoState98Action,
	{_CState167, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState167, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState167, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState167, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState167, CCastExpressionType}:           _CGotoState103Action,
	{_CState167, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState167, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState167, CShiftExpressionType}:          _CGotoState122Action,
	{_CState167, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState167, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState167, CAndExpressionType}:            _CGotoState101Action,
	{_CState167, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState167, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState167, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState167, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState167, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState167, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState167, CExpressionType}:               _CGotoState253Action,
	{_CState169, CRparamToken}:                  _CGotoState254Action,
	{_CState169, CCommaToken}:                   _CGotoState186Action,
	{_CState170, CLparamToken}:                  _CGotoState255Action,
	{_CState170, CLbraceToken}:                  _CGotoState223Action,
	{_CState170, CMulToken}:                     _CGotoState13Action,
	{_CState170, CPointerType}:                  _CGotoState257Action,
	{_CState170, CAbstractDeclaratorType}:       _CGotoState256Action,
	{_CState170, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState171, CRparamToken}:                  _CGotoState258Action,
	{_CState173, CSemicolonToken}:               _CGotoState259Action,
	{_CState173, CCommaToken}:                   _CGotoState186Action,
	{_CState174, CIdentifierToken}:              _CGotoState130Action,
	{_CState174, CConstantToken}:                _CGotoState77Action,
	{_CState174, CStringLiteralToken}:           _CGotoState96Action,
	{_CState174, CSizeofToken}:                  _CGotoState95Action,
	{_CState174, CIncOpToken}:                   _CGotoState87Action,
	{_CState174, CDecOpToken}:                   _CGotoState79Action,
	{_CState174, CTypeNameToken}:                _CGotoState20Action,
	{_CState174, CCharToken}:                    _CGotoState3Action,
	{_CState174, CShortToken}:                   _CGotoState15Action,
	{_CState174, CIntToken}:                     _CGotoState10Action,
	{_CState174, CLongToken}:                    _CGotoState11Action,
	{_CState174, CSignedToken}:                  _CGotoState16Action,
	{_CState174, CUnsignedToken}:                _CGotoState22Action,
	{_CState174, CFloatToken}:                   _CGotoState8Action,
	{_CState174, CDoubleToken}:                  _CGotoState5Action,
	{_CState174, CConstToken}:                   _CGotoState4Action,
	{_CState174, CVolatileToken}:                _CGotoState24Action,
	{_CState174, CVoidToken}:                    _CGotoState23Action,
	{_CState174, CStructToken}:                  _CGotoState18Action,
	{_CState174, CUnionToken}:                   _CGotoState21Action,
	{_CState174, CEnumToken}:                    _CGotoState6Action,
	{_CState174, CLparamToken}:                  _CGotoState88Action,
	{_CState174, CMulToken}:                     _CGotoState90Action,
	{_CState174, CMinusToken}:                   _CGotoState89Action,
	{_CState174, CPlusToken}:                    _CGotoState91Action,
	{_CState174, CAndToken}:                     _CGotoState74Action,
	{_CState174, CExclaimToken}:                 _CGotoState82Action,
	{_CState174, CTildaToken}:                   _CGotoState98Action,
	{_CState174, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState174, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState174, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState174, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState174, CCastExpressionType}:           _CGotoState103Action,
	{_CState174, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState174, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState174, CShiftExpressionType}:          _CGotoState122Action,
	{_CState174, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState174, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState174, CAndExpressionType}:            _CGotoState101Action,
	{_CState174, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState174, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState174, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState174, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState174, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState174, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState174, CExpressionType}:               _CGotoState169Action,
	{_CState174, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState174, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState174, CStructOrUnionType}:            _CGotoState34Action,
	{_CState174, CSpecifierQualifierListType}:   _CGotoState170Action,
	{_CState174, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState174, CTypeQualifierType}:            _CGotoState146Action,
	{_CState174, CTypeNameType}:                 _CGotoState260Action,
	{_CState176, CIdentifierToken}:              _CGotoState130Action,
	{_CState176, CConstantToken}:                _CGotoState77Action,
	{_CState176, CStringLiteralToken}:           _CGotoState96Action,
	{_CState176, CSizeofToken}:                  _CGotoState95Action,
	{_CState176, CIncOpToken}:                   _CGotoState87Action,
	{_CState176, CDecOpToken}:                   _CGotoState79Action,
	{_CState176, CLparamToken}:                  _CGotoState88Action,
	{_CState176, CMulToken}:                     _CGotoState90Action,
	{_CState176, CMinusToken}:                   _CGotoState89Action,
	{_CState176, CPlusToken}:                    _CGotoState91Action,
	{_CState176, CAndToken}:                     _CGotoState74Action,
	{_CState176, CExclaimToken}:                 _CGotoState82Action,
	{_CState176, CTildaToken}:                   _CGotoState98Action,
	{_CState176, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState176, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState176, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState176, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState176, CCastExpressionType}:           _CGotoState103Action,
	{_CState176, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState176, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState176, CShiftExpressionType}:          _CGotoState122Action,
	{_CState176, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState176, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState176, CAndExpressionType}:            _CGotoState101Action,
	{_CState176, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState176, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState176, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState176, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState176, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState176, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState176, CExpressionType}:               _CGotoState261Action,
	{_CState177, CIdentifierToken}:              _CGotoState130Action,
	{_CState177, CConstantToken}:                _CGotoState77Action,
	{_CState177, CStringLiteralToken}:           _CGotoState96Action,
	{_CState177, CSizeofToken}:                  _CGotoState95Action,
	{_CState177, CIncOpToken}:                   _CGotoState87Action,
	{_CState177, CDecOpToken}:                   _CGotoState79Action,
	{_CState177, CLparamToken}:                  _CGotoState88Action,
	{_CState177, CMulToken}:                     _CGotoState90Action,
	{_CState177, CMinusToken}:                   _CGotoState89Action,
	{_CState177, CPlusToken}:                    _CGotoState91Action,
	{_CState177, CAndToken}:                     _CGotoState74Action,
	{_CState177, CExclaimToken}:                 _CGotoState82Action,
	{_CState177, CTildaToken}:                   _CGotoState98Action,
	{_CState177, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState177, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState177, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState177, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState177, CCastExpressionType}:           _CGotoState103Action,
	{_CState177, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState177, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState177, CShiftExpressionType}:          _CGotoState122Action,
	{_CState177, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState177, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState177, CAndExpressionType}:            _CGotoState101Action,
	{_CState177, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState177, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState177, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState177, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState177, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState177, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState177, CExpressionType}:               _CGotoState262Action,
	{_CState178, CIdentifierToken}:              _CGotoState130Action,
	{_CState178, CConstantToken}:                _CGotoState77Action,
	{_CState178, CStringLiteralToken}:           _CGotoState96Action,
	{_CState178, CSizeofToken}:                  _CGotoState95Action,
	{_CState178, CIncOpToken}:                   _CGotoState87Action,
	{_CState178, CDecOpToken}:                   _CGotoState79Action,
	{_CState178, CLparamToken}:                  _CGotoState88Action,
	{_CState178, CMulToken}:                     _CGotoState90Action,
	{_CState178, CMinusToken}:                   _CGotoState89Action,
	{_CState178, CPlusToken}:                    _CGotoState91Action,
	{_CState178, CAndToken}:                     _CGotoState74Action,
	{_CState178, CExclaimToken}:                 _CGotoState82Action,
	{_CState178, CTildaToken}:                   _CGotoState98Action,
	{_CState178, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState178, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState178, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState178, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState178, CCastExpressionType}:           _CGotoState103Action,
	{_CState178, CMultiplicativeExpressionType}: _CGotoState263Action,
	{_CState179, CIdentifierToken}:              _CGotoState130Action,
	{_CState179, CConstantToken}:                _CGotoState77Action,
	{_CState179, CStringLiteralToken}:           _CGotoState96Action,
	{_CState179, CSizeofToken}:                  _CGotoState95Action,
	{_CState179, CIncOpToken}:                   _CGotoState87Action,
	{_CState179, CDecOpToken}:                   _CGotoState79Action,
	{_CState179, CLparamToken}:                  _CGotoState88Action,
	{_CState179, CMulToken}:                     _CGotoState90Action,
	{_CState179, CMinusToken}:                   _CGotoState89Action,
	{_CState179, CPlusToken}:                    _CGotoState91Action,
	{_CState179, CAndToken}:                     _CGotoState74Action,
	{_CState179, CExclaimToken}:                 _CGotoState82Action,
	{_CState179, CTildaToken}:                   _CGotoState98Action,
	{_CState179, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState179, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState179, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState179, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState179, CCastExpressionType}:           _CGotoState103Action,
	{_CState179, CMultiplicativeExpressionType}: _CGotoState264Action,
	{_CState180, CIdentifierToken}:              _CGotoState130Action,
	{_CState180, CConstantToken}:                _CGotoState77Action,
	{_CState180, CStringLiteralToken}:           _CGotoState96Action,
	{_CState180, CSizeofToken}:                  _CGotoState95Action,
	{_CState180, CIncOpToken}:                   _CGotoState87Action,
	{_CState180, CDecOpToken}:                   _CGotoState79Action,
	{_CState180, CLparamToken}:                  _CGotoState88Action,
	{_CState180, CMulToken}:                     _CGotoState90Action,
	{_CState180, CMinusToken}:                   _CGotoState89Action,
	{_CState180, CPlusToken}:                    _CGotoState91Action,
	{_CState180, CAndToken}:                     _CGotoState74Action,
	{_CState180, CExclaimToken}:                 _CGotoState82Action,
	{_CState180, CTildaToken}:                   _CGotoState98Action,
	{_CState180, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState180, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState180, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState180, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState180, CCastExpressionType}:           _CGotoState103Action,
	{_CState180, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState180, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState180, CShiftExpressionType}:          _CGotoState122Action,
	{_CState180, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState180, CEqualityExpressionType}:       _CGotoState265Action,
	{_CState182, CIdentifierToken}:              _CGotoState85Action,
	{_CState182, CConstantToken}:                _CGotoState77Action,
	{_CState182, CStringLiteralToken}:           _CGotoState96Action,
	{_CState182, CSizeofToken}:                  _CGotoState95Action,
	{_CState182, CIncOpToken}:                   _CGotoState87Action,
	{_CState182, CDecOpToken}:                   _CGotoState79Action,
	{_CState182, CCaseToken}:                    _CGotoState76Action,
	{_CState182, CDefaultToken}:                 _CGotoState80Action,
	{_CState182, CIfToken}:                      _CGotoState86Action,
	{_CState182, CSwitchToken}:                  _CGotoState97Action,
	{_CState182, CWhileToken}:                   _CGotoState99Action,
	{_CState182, CDoToken}:                      _CGotoState81Action,
	{_CState182, CForToken}:                     _CGotoState83Action,
	{_CState182, CGotoToken}:                    _CGotoState84Action,
	{_CState182, CContinueToken}:                _CGotoState78Action,
	{_CState182, CBreakToken}:                   _CGotoState75Action,
	{_CState182, CReturnToken}:                  _CGotoState93Action,
	{_CState182, CLparamToken}:                  _CGotoState88Action,
	{_CState182, CLcurlToken}:                   _CGotoState49Action,
	{_CState182, CRcurlToken}:                   _CGotoState266Action,
	{_CState182, CSemicolonToken}:               _CGotoState94Action,
	{_CState182, CMulToken}:                     _CGotoState90Action,
	{_CState182, CMinusToken}:                   _CGotoState89Action,
	{_CState182, CPlusToken}:                    _CGotoState91Action,
	{_CState182, CAndToken}:                     _CGotoState74Action,
	{_CState182, CExclaimToken}:                 _CGotoState82Action,
	{_CState182, CTildaToken}:                   _CGotoState98Action,
	{_CState182, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState182, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState182, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState182, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState182, CCastExpressionType}:           _CGotoState103Action,
	{_CState182, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState182, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState182, CShiftExpressionType}:          _CGotoState122Action,
	{_CState182, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState182, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState182, CAndExpressionType}:            _CGotoState101Action,
	{_CState182, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState182, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState182, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState182, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState182, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState182, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState182, CExpressionType}:               _CGotoState109Action,
	{_CState182, CStatementType}:                _CGotoState208Action,
	{_CState182, CLabeledStatementType}:         _CGotoState114Action,
	{_CState182, CCompoundStatementType}:        _CGotoState104Action,
	{_CState182, CExpressionStatementType}:      _CGotoState110Action,
	{_CState182, CSelectionStatementType}:       _CGotoState121Action,
	{_CState182, CIterationStatementType}:       _CGotoState112Action,
	{_CState182, CJumpStatementType}:            _CGotoState113Action,
	{_CState183, CIdentifierToken}:              _CGotoState130Action,
	{_CState183, CConstantToken}:                _CGotoState77Action,
	{_CState183, CStringLiteralToken}:           _CGotoState96Action,
	{_CState183, CSizeofToken}:                  _CGotoState95Action,
	{_CState183, CIncOpToken}:                   _CGotoState87Action,
	{_CState183, CDecOpToken}:                   _CGotoState79Action,
	{_CState183, CLparamToken}:                  _CGotoState88Action,
	{_CState183, CMulToken}:                     _CGotoState90Action,
	{_CState183, CMinusToken}:                   _CGotoState89Action,
	{_CState183, CPlusToken}:                    _CGotoState91Action,
	{_CState183, CAndToken}:                     _CGotoState74Action,
	{_CState183, CExclaimToken}:                 _CGotoState82Action,
	{_CState183, CTildaToken}:                   _CGotoState98Action,
	{_CState183, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState183, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState183, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState183, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState183, CCastExpressionType}:           _CGotoState103Action,
	{_CState183, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState183, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState183, CShiftExpressionType}:          _CGotoState122Action,
	{_CState183, CRelationalExpressionType}:     _CGotoState267Action,
	{_CState184, CIdentifierToken}:              _CGotoState130Action,
	{_CState184, CConstantToken}:                _CGotoState77Action,
	{_CState184, CStringLiteralToken}:           _CGotoState96Action,
	{_CState184, CSizeofToken}:                  _CGotoState95Action,
	{_CState184, CIncOpToken}:                   _CGotoState87Action,
	{_CState184, CDecOpToken}:                   _CGotoState79Action,
	{_CState184, CLparamToken}:                  _CGotoState88Action,
	{_CState184, CMulToken}:                     _CGotoState90Action,
	{_CState184, CMinusToken}:                   _CGotoState89Action,
	{_CState184, CPlusToken}:                    _CGotoState91Action,
	{_CState184, CAndToken}:                     _CGotoState74Action,
	{_CState184, CExclaimToken}:                 _CGotoState82Action,
	{_CState184, CTildaToken}:                   _CGotoState98Action,
	{_CState184, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState184, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState184, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState184, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState184, CCastExpressionType}:           _CGotoState103Action,
	{_CState184, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState184, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState184, CShiftExpressionType}:          _CGotoState122Action,
	{_CState184, CRelationalExpressionType}:     _CGotoState268Action,
	{_CState185, CIdentifierToken}:              _CGotoState130Action,
	{_CState185, CConstantToken}:                _CGotoState77Action,
	{_CState185, CStringLiteralToken}:           _CGotoState96Action,
	{_CState185, CSizeofToken}:                  _CGotoState95Action,
	{_CState185, CIncOpToken}:                   _CGotoState87Action,
	{_CState185, CDecOpToken}:                   _CGotoState79Action,
	{_CState185, CLparamToken}:                  _CGotoState88Action,
	{_CState185, CMulToken}:                     _CGotoState90Action,
	{_CState185, CMinusToken}:                   _CGotoState89Action,
	{_CState185, CPlusToken}:                    _CGotoState91Action,
	{_CState185, CAndToken}:                     _CGotoState74Action,
	{_CState185, CExclaimToken}:                 _CGotoState82Action,
	{_CState185, CTildaToken}:                   _CGotoState98Action,
	{_CState185, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState185, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState185, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState185, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState185, CCastExpressionType}:           _CGotoState103Action,
	{_CState185, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState185, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState185, CShiftExpressionType}:          _CGotoState122Action,
	{_CState185, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState185, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState185, CAndExpressionType}:            _CGotoState269Action,
	{_CState186, CIdentifierToken}:              _CGotoState130Action,
	{_CState186, CConstantToken}:                _CGotoState77Action,
	{_CState186, CStringLiteralToken}:           _CGotoState96Action,
	{_CState186, CSizeofToken}:                  _CGotoState95Action,
	{_CState186, CIncOpToken}:                   _CGotoState87Action,
	{_CState186, CDecOpToken}:                   _CGotoState79Action,
	{_CState186, CLparamToken}:                  _CGotoState88Action,
	{_CState186, CMulToken}:                     _CGotoState90Action,
	{_CState186, CMinusToken}:                   _CGotoState89Action,
	{_CState186, CPlusToken}:                    _CGotoState91Action,
	{_CState186, CAndToken}:                     _CGotoState74Action,
	{_CState186, CExclaimToken}:                 _CGotoState82Action,
	{_CState186, CTildaToken}:                   _CGotoState98Action,
	{_CState186, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState186, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState186, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState186, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState186, CCastExpressionType}:           _CGotoState103Action,
	{_CState186, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState186, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState186, CShiftExpressionType}:          _CGotoState122Action,
	{_CState186, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState186, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState186, CAndExpressionType}:            _CGotoState101Action,
	{_CState186, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState186, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState186, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState186, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState186, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState186, CAssignmentExpressionType}:     _CGotoState270Action,
	{_CState188, CIdentifierToken}:              _CGotoState130Action,
	{_CState188, CConstantToken}:                _CGotoState77Action,
	{_CState188, CStringLiteralToken}:           _CGotoState96Action,
	{_CState188, CSizeofToken}:                  _CGotoState95Action,
	{_CState188, CIncOpToken}:                   _CGotoState87Action,
	{_CState188, CDecOpToken}:                   _CGotoState79Action,
	{_CState188, CLparamToken}:                  _CGotoState88Action,
	{_CState188, CMulToken}:                     _CGotoState90Action,
	{_CState188, CMinusToken}:                   _CGotoState89Action,
	{_CState188, CPlusToken}:                    _CGotoState91Action,
	{_CState188, CAndToken}:                     _CGotoState74Action,
	{_CState188, CExclaimToken}:                 _CGotoState82Action,
	{_CState188, CTildaToken}:                   _CGotoState98Action,
	{_CState188, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState188, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState188, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState188, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState188, CCastExpressionType}:           _CGotoState103Action,
	{_CState188, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState188, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState188, CShiftExpressionType}:          _CGotoState122Action,
	{_CState188, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState188, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState188, CAndExpressionType}:            _CGotoState101Action,
	{_CState188, CExclusiveOrExpressionType}:    _CGotoState271Action,
	{_CState189, CIdentifierToken}:              _CGotoState130Action,
	{_CState189, CConstantToken}:                _CGotoState77Action,
	{_CState189, CStringLiteralToken}:           _CGotoState96Action,
	{_CState189, CSizeofToken}:                  _CGotoState95Action,
	{_CState189, CIncOpToken}:                   _CGotoState87Action,
	{_CState189, CDecOpToken}:                   _CGotoState79Action,
	{_CState189, CLparamToken}:                  _CGotoState88Action,
	{_CState189, CMulToken}:                     _CGotoState90Action,
	{_CState189, CMinusToken}:                   _CGotoState89Action,
	{_CState189, CPlusToken}:                    _CGotoState91Action,
	{_CState189, CAndToken}:                     _CGotoState74Action,
	{_CState189, CExclaimToken}:                 _CGotoState82Action,
	{_CState189, CTildaToken}:                   _CGotoState98Action,
	{_CState189, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState189, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState189, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState189, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState189, CCastExpressionType}:           _CGotoState103Action,
	{_CState189, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState189, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState189, CShiftExpressionType}:          _CGotoState122Action,
	{_CState189, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState189, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState189, CAndExpressionType}:            _CGotoState101Action,
	{_CState189, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState189, CInclusiveOrExpressionType}:    _CGotoState272Action,
	{_CState190, CIdentifierToken}:              _CGotoState130Action,
	{_CState190, CConstantToken}:                _CGotoState77Action,
	{_CState190, CStringLiteralToken}:           _CGotoState96Action,
	{_CState190, CSizeofToken}:                  _CGotoState95Action,
	{_CState190, CIncOpToken}:                   _CGotoState87Action,
	{_CState190, CDecOpToken}:                   _CGotoState79Action,
	{_CState190, CLparamToken}:                  _CGotoState88Action,
	{_CState190, CMulToken}:                     _CGotoState90Action,
	{_CState190, CMinusToken}:                   _CGotoState89Action,
	{_CState190, CPlusToken}:                    _CGotoState91Action,
	{_CState190, CAndToken}:                     _CGotoState74Action,
	{_CState190, CExclaimToken}:                 _CGotoState82Action,
	{_CState190, CTildaToken}:                   _CGotoState98Action,
	{_CState190, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState190, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState190, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState190, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState190, CCastExpressionType}:           _CGotoState103Action,
	{_CState190, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState190, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState190, CShiftExpressionType}:          _CGotoState122Action,
	{_CState190, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState190, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState190, CAndExpressionType}:            _CGotoState101Action,
	{_CState190, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState190, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState190, CLogicalAndExpressionType}:     _CGotoState273Action,
	{_CState191, CIdentifierToken}:              _CGotoState130Action,
	{_CState191, CConstantToken}:                _CGotoState77Action,
	{_CState191, CStringLiteralToken}:           _CGotoState96Action,
	{_CState191, CSizeofToken}:                  _CGotoState95Action,
	{_CState191, CIncOpToken}:                   _CGotoState87Action,
	{_CState191, CDecOpToken}:                   _CGotoState79Action,
	{_CState191, CLparamToken}:                  _CGotoState88Action,
	{_CState191, CMulToken}:                     _CGotoState90Action,
	{_CState191, CMinusToken}:                   _CGotoState89Action,
	{_CState191, CPlusToken}:                    _CGotoState91Action,
	{_CState191, CAndToken}:                     _CGotoState74Action,
	{_CState191, CExclaimToken}:                 _CGotoState82Action,
	{_CState191, CTildaToken}:                   _CGotoState98Action,
	{_CState191, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState191, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState191, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState191, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState191, CCastExpressionType}:           _CGotoState103Action,
	{_CState191, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState191, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState191, CShiftExpressionType}:          _CGotoState122Action,
	{_CState191, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState191, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState191, CAndExpressionType}:            _CGotoState101Action,
	{_CState191, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState191, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState191, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState191, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState191, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState191, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState191, CExpressionType}:               _CGotoState274Action,
	{_CState192, CIdentifierToken}:              _CGotoState130Action,
	{_CState192, CConstantToken}:                _CGotoState77Action,
	{_CState192, CStringLiteralToken}:           _CGotoState96Action,
	{_CState192, CSizeofToken}:                  _CGotoState95Action,
	{_CState192, CIncOpToken}:                   _CGotoState87Action,
	{_CState192, CDecOpToken}:                   _CGotoState79Action,
	{_CState192, CLparamToken}:                  _CGotoState88Action,
	{_CState192, CMulToken}:                     _CGotoState90Action,
	{_CState192, CMinusToken}:                   _CGotoState89Action,
	{_CState192, CPlusToken}:                    _CGotoState91Action,
	{_CState192, CAndToken}:                     _CGotoState74Action,
	{_CState192, CExclaimToken}:                 _CGotoState82Action,
	{_CState192, CTildaToken}:                   _CGotoState98Action,
	{_CState192, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState192, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState192, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState192, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState192, CCastExpressionType}:           _CGotoState275Action,
	{_CState193, CIdentifierToken}:              _CGotoState130Action,
	{_CState193, CConstantToken}:                _CGotoState77Action,
	{_CState193, CStringLiteralToken}:           _CGotoState96Action,
	{_CState193, CSizeofToken}:                  _CGotoState95Action,
	{_CState193, CIncOpToken}:                   _CGotoState87Action,
	{_CState193, CDecOpToken}:                   _CGotoState79Action,
	{_CState193, CLparamToken}:                  _CGotoState88Action,
	{_CState193, CMulToken}:                     _CGotoState90Action,
	{_CState193, CMinusToken}:                   _CGotoState89Action,
	{_CState193, CPlusToken}:                    _CGotoState91Action,
	{_CState193, CAndToken}:                     _CGotoState74Action,
	{_CState193, CExclaimToken}:                 _CGotoState82Action,
	{_CState193, CTildaToken}:                   _CGotoState98Action,
	{_CState193, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState193, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState193, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState193, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState193, CCastExpressionType}:           _CGotoState276Action,
	{_CState194, CIdentifierToken}:              _CGotoState130Action,
	{_CState194, CConstantToken}:                _CGotoState77Action,
	{_CState194, CStringLiteralToken}:           _CGotoState96Action,
	{_CState194, CSizeofToken}:                  _CGotoState95Action,
	{_CState194, CIncOpToken}:                   _CGotoState87Action,
	{_CState194, CDecOpToken}:                   _CGotoState79Action,
	{_CState194, CLparamToken}:                  _CGotoState88Action,
	{_CState194, CMulToken}:                     _CGotoState90Action,
	{_CState194, CMinusToken}:                   _CGotoState89Action,
	{_CState194, CPlusToken}:                    _CGotoState91Action,
	{_CState194, CAndToken}:                     _CGotoState74Action,
	{_CState194, CExclaimToken}:                 _CGotoState82Action,
	{_CState194, CTildaToken}:                   _CGotoState98Action,
	{_CState194, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState194, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState194, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState194, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState194, CCastExpressionType}:           _CGotoState277Action,
	{_CState196, CIdentifierToken}:              _CGotoState278Action,
	{_CState198, CIdentifierToken}:              _CGotoState130Action,
	{_CState198, CConstantToken}:                _CGotoState77Action,
	{_CState198, CStringLiteralToken}:           _CGotoState96Action,
	{_CState198, CSizeofToken}:                  _CGotoState95Action,
	{_CState198, CIncOpToken}:                   _CGotoState87Action,
	{_CState198, CDecOpToken}:                   _CGotoState79Action,
	{_CState198, CLparamToken}:                  _CGotoState88Action,
	{_CState198, CMulToken}:                     _CGotoState90Action,
	{_CState198, CMinusToken}:                   _CGotoState89Action,
	{_CState198, CPlusToken}:                    _CGotoState91Action,
	{_CState198, CAndToken}:                     _CGotoState74Action,
	{_CState198, CExclaimToken}:                 _CGotoState82Action,
	{_CState198, CTildaToken}:                   _CGotoState98Action,
	{_CState198, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState198, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState198, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState198, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState198, CCastExpressionType}:           _CGotoState103Action,
	{_CState198, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState198, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState198, CShiftExpressionType}:          _CGotoState122Action,
	{_CState198, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState198, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState198, CAndExpressionType}:            _CGotoState101Action,
	{_CState198, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState198, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState198, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState198, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState198, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState198, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState198, CExpressionType}:               _CGotoState279Action,
	{_CState199, CIdentifierToken}:              _CGotoState130Action,
	{_CState199, CConstantToken}:                _CGotoState77Action,
	{_CState199, CStringLiteralToken}:           _CGotoState96Action,
	{_CState199, CSizeofToken}:                  _CGotoState95Action,
	{_CState199, CIncOpToken}:                   _CGotoState87Action,
	{_CState199, CDecOpToken}:                   _CGotoState79Action,
	{_CState199, CLparamToken}:                  _CGotoState88Action,
	{_CState199, CRparamToken}:                  _CGotoState280Action,
	{_CState199, CMulToken}:                     _CGotoState90Action,
	{_CState199, CMinusToken}:                   _CGotoState89Action,
	{_CState199, CPlusToken}:                    _CGotoState91Action,
	{_CState199, CAndToken}:                     _CGotoState74Action,
	{_CState199, CExclaimToken}:                 _CGotoState82Action,
	{_CState199, CTildaToken}:                   _CGotoState98Action,
	{_CState199, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState199, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState199, CArgumentExpressionListType}:   _CGotoState281Action,
	{_CState199, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState199, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState199, CCastExpressionType}:           _CGotoState103Action,
	{_CState199, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState199, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState199, CShiftExpressionType}:          _CGotoState122Action,
	{_CState199, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState199, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState199, CAndExpressionType}:            _CGotoState101Action,
	{_CState199, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState199, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState199, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState199, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState199, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState199, CAssignmentExpressionType}:     _CGotoState282Action,
	{_CState200, CIdentifierToken}:              _CGotoState283Action,
	{_CState201, CIdentifierToken}:              _CGotoState130Action,
	{_CState201, CConstantToken}:                _CGotoState77Action,
	{_CState201, CStringLiteralToken}:           _CGotoState96Action,
	{_CState201, CSizeofToken}:                  _CGotoState95Action,
	{_CState201, CIncOpToken}:                   _CGotoState87Action,
	{_CState201, CDecOpToken}:                   _CGotoState79Action,
	{_CState201, CLparamToken}:                  _CGotoState88Action,
	{_CState201, CMulToken}:                     _CGotoState90Action,
	{_CState201, CMinusToken}:                   _CGotoState89Action,
	{_CState201, CPlusToken}:                    _CGotoState91Action,
	{_CState201, CAndToken}:                     _CGotoState74Action,
	{_CState201, CExclaimToken}:                 _CGotoState82Action,
	{_CState201, CTildaToken}:                   _CGotoState98Action,
	{_CState201, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState201, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState201, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState201, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState201, CCastExpressionType}:           _CGotoState103Action,
	{_CState201, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState201, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState201, CShiftExpressionType}:          _CGotoState284Action,
	{_CState202, CIdentifierToken}:              _CGotoState130Action,
	{_CState202, CConstantToken}:                _CGotoState77Action,
	{_CState202, CStringLiteralToken}:           _CGotoState96Action,
	{_CState202, CSizeofToken}:                  _CGotoState95Action,
	{_CState202, CIncOpToken}:                   _CGotoState87Action,
	{_CState202, CDecOpToken}:                   _CGotoState79Action,
	{_CState202, CLparamToken}:                  _CGotoState88Action,
	{_CState202, CMulToken}:                     _CGotoState90Action,
	{_CState202, CMinusToken}:                   _CGotoState89Action,
	{_CState202, CPlusToken}:                    _CGotoState91Action,
	{_CState202, CAndToken}:                     _CGotoState74Action,
	{_CState202, CExclaimToken}:                 _CGotoState82Action,
	{_CState202, CTildaToken}:                   _CGotoState98Action,
	{_CState202, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState202, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState202, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState202, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState202, CCastExpressionType}:           _CGotoState103Action,
	{_CState202, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState202, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState202, CShiftExpressionType}:          _CGotoState285Action,
	{_CState203, CIdentifierToken}:              _CGotoState130Action,
	{_CState203, CConstantToken}:                _CGotoState77Action,
	{_CState203, CStringLiteralToken}:           _CGotoState96Action,
	{_CState203, CSizeofToken}:                  _CGotoState95Action,
	{_CState203, CIncOpToken}:                   _CGotoState87Action,
	{_CState203, CDecOpToken}:                   _CGotoState79Action,
	{_CState203, CLparamToken}:                  _CGotoState88Action,
	{_CState203, CMulToken}:                     _CGotoState90Action,
	{_CState203, CMinusToken}:                   _CGotoState89Action,
	{_CState203, CPlusToken}:                    _CGotoState91Action,
	{_CState203, CAndToken}:                     _CGotoState74Action,
	{_CState203, CExclaimToken}:                 _CGotoState82Action,
	{_CState203, CTildaToken}:                   _CGotoState98Action,
	{_CState203, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState203, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState203, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState203, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState203, CCastExpressionType}:           _CGotoState103Action,
	{_CState203, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState203, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState203, CShiftExpressionType}:          _CGotoState286Action,
	{_CState204, CIdentifierToken}:              _CGotoState130Action,
	{_CState204, CConstantToken}:                _CGotoState77Action,
	{_CState204, CStringLiteralToken}:           _CGotoState96Action,
	{_CState204, CSizeofToken}:                  _CGotoState95Action,
	{_CState204, CIncOpToken}:                   _CGotoState87Action,
	{_CState204, CDecOpToken}:                   _CGotoState79Action,
	{_CState204, CLparamToken}:                  _CGotoState88Action,
	{_CState204, CMulToken}:                     _CGotoState90Action,
	{_CState204, CMinusToken}:                   _CGotoState89Action,
	{_CState204, CPlusToken}:                    _CGotoState91Action,
	{_CState204, CAndToken}:                     _CGotoState74Action,
	{_CState204, CExclaimToken}:                 _CGotoState82Action,
	{_CState204, CTildaToken}:                   _CGotoState98Action,
	{_CState204, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState204, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState204, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState204, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState204, CCastExpressionType}:           _CGotoState103Action,
	{_CState204, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState204, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState204, CShiftExpressionType}:          _CGotoState287Action,
	{_CState205, CIdentifierToken}:              _CGotoState130Action,
	{_CState205, CConstantToken}:                _CGotoState77Action,
	{_CState205, CStringLiteralToken}:           _CGotoState96Action,
	{_CState205, CSizeofToken}:                  _CGotoState95Action,
	{_CState205, CIncOpToken}:                   _CGotoState87Action,
	{_CState205, CDecOpToken}:                   _CGotoState79Action,
	{_CState205, CLparamToken}:                  _CGotoState88Action,
	{_CState205, CMulToken}:                     _CGotoState90Action,
	{_CState205, CMinusToken}:                   _CGotoState89Action,
	{_CState205, CPlusToken}:                    _CGotoState91Action,
	{_CState205, CAndToken}:                     _CGotoState74Action,
	{_CState205, CExclaimToken}:                 _CGotoState82Action,
	{_CState205, CTildaToken}:                   _CGotoState98Action,
	{_CState205, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState205, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState205, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState205, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState205, CCastExpressionType}:           _CGotoState103Action,
	{_CState205, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState205, CAdditiveExpressionType}:       _CGotoState288Action,
	{_CState206, CIdentifierToken}:              _CGotoState130Action,
	{_CState206, CConstantToken}:                _CGotoState77Action,
	{_CState206, CStringLiteralToken}:           _CGotoState96Action,
	{_CState206, CSizeofToken}:                  _CGotoState95Action,
	{_CState206, CIncOpToken}:                   _CGotoState87Action,
	{_CState206, CDecOpToken}:                   _CGotoState79Action,
	{_CState206, CLparamToken}:                  _CGotoState88Action,
	{_CState206, CMulToken}:                     _CGotoState90Action,
	{_CState206, CMinusToken}:                   _CGotoState89Action,
	{_CState206, CPlusToken}:                    _CGotoState91Action,
	{_CState206, CAndToken}:                     _CGotoState74Action,
	{_CState206, CExclaimToken}:                 _CGotoState82Action,
	{_CState206, CTildaToken}:                   _CGotoState98Action,
	{_CState206, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState206, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState206, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState206, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState206, CCastExpressionType}:           _CGotoState103Action,
	{_CState206, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState206, CAdditiveExpressionType}:       _CGotoState289Action,
	{_CState220, CIdentifierToken}:              _CGotoState130Action,
	{_CState220, CConstantToken}:                _CGotoState77Action,
	{_CState220, CStringLiteralToken}:           _CGotoState96Action,
	{_CState220, CSizeofToken}:                  _CGotoState95Action,
	{_CState220, CIncOpToken}:                   _CGotoState87Action,
	{_CState220, CDecOpToken}:                   _CGotoState79Action,
	{_CState220, CLparamToken}:                  _CGotoState88Action,
	{_CState220, CMulToken}:                     _CGotoState90Action,
	{_CState220, CMinusToken}:                   _CGotoState89Action,
	{_CState220, CPlusToken}:                    _CGotoState91Action,
	{_CState220, CAndToken}:                     _CGotoState74Action,
	{_CState220, CExclaimToken}:                 _CGotoState82Action,
	{_CState220, CTildaToken}:                   _CGotoState98Action,
	{_CState220, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState220, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState220, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState220, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState220, CCastExpressionType}:           _CGotoState103Action,
	{_CState220, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState220, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState220, CShiftExpressionType}:          _CGotoState122Action,
	{_CState220, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState220, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState220, CAndExpressionType}:            _CGotoState101Action,
	{_CState220, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState220, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState220, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState220, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState220, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState220, CAssignmentExpressionType}:     _CGotoState290Action,
	{_CState223, CIdentifierToken}:              _CGotoState130Action,
	{_CState223, CConstantToken}:                _CGotoState77Action,
	{_CState223, CStringLiteralToken}:           _CGotoState96Action,
	{_CState223, CSizeofToken}:                  _CGotoState95Action,
	{_CState223, CIncOpToken}:                   _CGotoState87Action,
	{_CState223, CDecOpToken}:                   _CGotoState79Action,
	{_CState223, CLparamToken}:                  _CGotoState88Action,
	{_CState223, CRbraceToken}:                  _CGotoState291Action,
	{_CState223, CMulToken}:                     _CGotoState90Action,
	{_CState223, CMinusToken}:                   _CGotoState89Action,
	{_CState223, CPlusToken}:                    _CGotoState91Action,
	{_CState223, CAndToken}:                     _CGotoState74Action,
	{_CState223, CExclaimToken}:                 _CGotoState82Action,
	{_CState223, CTildaToken}:                   _CGotoState98Action,
	{_CState223, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState223, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState223, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState223, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState223, CCastExpressionType}:           _CGotoState103Action,
	{_CState223, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState223, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState223, CShiftExpressionType}:          _CGotoState122Action,
	{_CState223, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState223, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState223, CAndExpressionType}:            _CGotoState101Action,
	{_CState223, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState223, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState223, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState223, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState223, CConditionalExpressionType}:    _CGotoState132Action,
	{_CState223, CConstantExpressionType}:       _CGotoState292Action,
	{_CState224, CIdentifierToken}:              _CGotoState9Action,
	{_CState224, CTypeNameToken}:                _CGotoState20Action,
	{_CState224, CTypedefToken}:                 _CGotoState19Action,
	{_CState224, CExternToken}:                  _CGotoState7Action,
	{_CState224, CStaticToken}:                  _CGotoState17Action,
	{_CState224, CAutoToken}:                    _CGotoState2Action,
	{_CState224, CRegisterToken}:                _CGotoState14Action,
	{_CState224, CCharToken}:                    _CGotoState3Action,
	{_CState224, CShortToken}:                   _CGotoState15Action,
	{_CState224, CIntToken}:                     _CGotoState10Action,
	{_CState224, CLongToken}:                    _CGotoState11Action,
	{_CState224, CSignedToken}:                  _CGotoState16Action,
	{_CState224, CUnsignedToken}:                _CGotoState22Action,
	{_CState224, CFloatToken}:                   _CGotoState8Action,
	{_CState224, CDoubleToken}:                  _CGotoState5Action,
	{_CState224, CConstToken}:                   _CGotoState4Action,
	{_CState224, CVolatileToken}:                _CGotoState24Action,
	{_CState224, CVoidToken}:                    _CGotoState23Action,
	{_CState224, CStructToken}:                  _CGotoState18Action,
	{_CState224, CUnionToken}:                   _CGotoState21Action,
	{_CState224, CEnumToken}:                    _CGotoState6Action,
	{_CState224, CLparamToken}:                  _CGotoState224Action,
	{_CState224, CRparamToken}:                  _CGotoState293Action,
	{_CState224, CLbraceToken}:                  _CGotoState223Action,
	{_CState224, CMulToken}:                     _CGotoState13Action,
	{_CState224, CDeclarationSpecifiersType}:    _CGotoState137Action,
	{_CState224, CStorageClassSpecifierType}:    _CGotoState33Action,
	{_CState224, CTypeSpecifierType}:            _CGotoState37Action,
	{_CState224, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState224, CStructOrUnionType}:            _CGotoState34Action,
	{_CState224, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState224, CTypeQualifierType}:            _CGotoState36Action,
	{_CState224, CDeclaratorType}:               _CGotoState41Action,
	{_CState224, CDirectDeclaratorType}:         _CGotoState28Action,
	{_CState224, CPointerType}:                  _CGotoState228Action,
	{_CState224, CParameterTypeListType}:        _CGotoState295Action,
	{_CState224, CParameterListType}:            _CGotoState140Action,
	{_CState224, CParameterDeclarationType}:     _CGotoState139Action,
	{_CState224, CAbstractDeclaratorType}:       _CGotoState294Action,
	{_CState224, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState227, CLparamToken}:                  _CGotoState297Action,
	{_CState227, CLbraceToken}:                  _CGotoState296Action,
	{_CState228, CIdentifierToken}:              _CGotoState9Action,
	{_CState228, CLparamToken}:                  _CGotoState224Action,
	{_CState228, CLbraceToken}:                  _CGotoState223Action,
	{_CState228, CDirectDeclaratorType}:         _CGotoState56Action,
	{_CState228, CDirectAbstractDeclaratorType}: _CGotoState298Action,
	{_CState229, CIdentifierToken}:              _CGotoState299Action,
	{_CState231, CTypeNameToken}:                _CGotoState20Action,
	{_CState231, CTypedefToken}:                 _CGotoState19Action,
	{_CState231, CExternToken}:                  _CGotoState7Action,
	{_CState231, CStaticToken}:                  _CGotoState17Action,
	{_CState231, CAutoToken}:                    _CGotoState2Action,
	{_CState231, CRegisterToken}:                _CGotoState14Action,
	{_CState231, CCharToken}:                    _CGotoState3Action,
	{_CState231, CShortToken}:                   _CGotoState15Action,
	{_CState231, CIntToken}:                     _CGotoState10Action,
	{_CState231, CLongToken}:                    _CGotoState11Action,
	{_CState231, CSignedToken}:                  _CGotoState16Action,
	{_CState231, CUnsignedToken}:                _CGotoState22Action,
	{_CState231, CFloatToken}:                   _CGotoState8Action,
	{_CState231, CDoubleToken}:                  _CGotoState5Action,
	{_CState231, CConstToken}:                   _CGotoState4Action,
	{_CState231, CVolatileToken}:                _CGotoState24Action,
	{_CState231, CVoidToken}:                    _CGotoState23Action,
	{_CState231, CStructToken}:                  _CGotoState18Action,
	{_CState231, CUnionToken}:                   _CGotoState21Action,
	{_CState231, CEnumToken}:                    _CGotoState6Action,
	{_CState231, CEllipsisToken}:                _CGotoState300Action,
	{_CState231, CDeclarationSpecifiersType}:    _CGotoState137Action,
	{_CState231, CStorageClassSpecifierType}:    _CGotoState33Action,
	{_CState231, CTypeSpecifierType}:            _CGotoState37Action,
	{_CState231, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState231, CStructOrUnionType}:            _CGotoState34Action,
	{_CState231, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState231, CTypeQualifierType}:            _CGotoState36Action,
	{_CState231, CParameterDeclarationType}:     _CGotoState301Action,
	{_CState233, CTypeNameToken}:                _CGotoState20Action,
	{_CState233, CCharToken}:                    _CGotoState3Action,
	{_CState233, CShortToken}:                   _CGotoState15Action,
	{_CState233, CIntToken}:                     _CGotoState10Action,
	{_CState233, CLongToken}:                    _CGotoState11Action,
	{_CState233, CSignedToken}:                  _CGotoState16Action,
	{_CState233, CUnsignedToken}:                _CGotoState22Action,
	{_CState233, CFloatToken}:                   _CGotoState8Action,
	{_CState233, CDoubleToken}:                  _CGotoState5Action,
	{_CState233, CConstToken}:                   _CGotoState4Action,
	{_CState233, CVolatileToken}:                _CGotoState24Action,
	{_CState233, CVoidToken}:                    _CGotoState23Action,
	{_CState233, CStructToken}:                  _CGotoState18Action,
	{_CState233, CUnionToken}:                   _CGotoState21Action,
	{_CState233, CEnumToken}:                    _CGotoState6Action,
	{_CState233, CRcurlToken}:                   _CGotoState302Action,
	{_CState233, CTypeSpecifierType}:            _CGotoState147Action,
	{_CState233, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState233, CStructOrUnionType}:            _CGotoState34Action,
	{_CState233, CStructDeclarationType}:        _CGotoState239Action,
	{_CState233, CSpecifierQualifierListType}:   _CGotoState143Action,
	{_CState233, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState233, CTypeQualifierType}:            _CGotoState146Action,
	{_CState234, CIdentifierToken}:              _CGotoState130Action,
	{_CState234, CConstantToken}:                _CGotoState77Action,
	{_CState234, CStringLiteralToken}:           _CGotoState96Action,
	{_CState234, CSizeofToken}:                  _CGotoState95Action,
	{_CState234, CIncOpToken}:                   _CGotoState87Action,
	{_CState234, CDecOpToken}:                   _CGotoState79Action,
	{_CState234, CLparamToken}:                  _CGotoState88Action,
	{_CState234, CMulToken}:                     _CGotoState90Action,
	{_CState234, CMinusToken}:                   _CGotoState89Action,
	{_CState234, CPlusToken}:                    _CGotoState91Action,
	{_CState234, CAndToken}:                     _CGotoState74Action,
	{_CState234, CExclaimToken}:                 _CGotoState82Action,
	{_CState234, CTildaToken}:                   _CGotoState98Action,
	{_CState234, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState234, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState234, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState234, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState234, CCastExpressionType}:           _CGotoState103Action,
	{_CState234, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState234, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState234, CShiftExpressionType}:          _CGotoState122Action,
	{_CState234, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState234, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState234, CAndExpressionType}:            _CGotoState101Action,
	{_CState234, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState234, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState234, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState234, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState234, CConditionalExpressionType}:    _CGotoState132Action,
	{_CState234, CConstantExpressionType}:       _CGotoState303Action,
	{_CState235, CColonToken}:                   _CGotoState304Action,
	{_CState237, CSemicolonToken}:               _CGotoState306Action,
	{_CState237, CCommaToken}:                   _CGotoState305Action,
	{_CState246, CRcurlToken}:                   _CGotoState308Action,
	{_CState246, CCommaToken}:                   _CGotoState307Action,
	{_CState247, CIdentifierToken}:              _CGotoState85Action,
	{_CState247, CConstantToken}:                _CGotoState77Action,
	{_CState247, CStringLiteralToken}:           _CGotoState96Action,
	{_CState247, CSizeofToken}:                  _CGotoState95Action,
	{_CState247, CIncOpToken}:                   _CGotoState87Action,
	{_CState247, CDecOpToken}:                   _CGotoState79Action,
	{_CState247, CCaseToken}:                    _CGotoState76Action,
	{_CState247, CDefaultToken}:                 _CGotoState80Action,
	{_CState247, CIfToken}:                      _CGotoState86Action,
	{_CState247, CSwitchToken}:                  _CGotoState97Action,
	{_CState247, CWhileToken}:                   _CGotoState99Action,
	{_CState247, CDoToken}:                      _CGotoState81Action,
	{_CState247, CForToken}:                     _CGotoState83Action,
	{_CState247, CGotoToken}:                    _CGotoState84Action,
	{_CState247, CContinueToken}:                _CGotoState78Action,
	{_CState247, CBreakToken}:                   _CGotoState75Action,
	{_CState247, CReturnToken}:                  _CGotoState93Action,
	{_CState247, CLparamToken}:                  _CGotoState88Action,
	{_CState247, CLcurlToken}:                   _CGotoState49Action,
	{_CState247, CSemicolonToken}:               _CGotoState94Action,
	{_CState247, CMulToken}:                     _CGotoState90Action,
	{_CState247, CMinusToken}:                   _CGotoState89Action,
	{_CState247, CPlusToken}:                    _CGotoState91Action,
	{_CState247, CAndToken}:                     _CGotoState74Action,
	{_CState247, CExclaimToken}:                 _CGotoState82Action,
	{_CState247, CTildaToken}:                   _CGotoState98Action,
	{_CState247, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState247, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState247, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState247, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState247, CCastExpressionType}:           _CGotoState103Action,
	{_CState247, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState247, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState247, CShiftExpressionType}:          _CGotoState122Action,
	{_CState247, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState247, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState247, CAndExpressionType}:            _CGotoState101Action,
	{_CState247, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState247, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState247, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState247, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState247, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState247, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState247, CExpressionType}:               _CGotoState109Action,
	{_CState247, CStatementType}:                _CGotoState309Action,
	{_CState247, CLabeledStatementType}:         _CGotoState114Action,
	{_CState247, CCompoundStatementType}:        _CGotoState104Action,
	{_CState247, CExpressionStatementType}:      _CGotoState110Action,
	{_CState247, CSelectionStatementType}:       _CGotoState121Action,
	{_CState247, CIterationStatementType}:       _CGotoState112Action,
	{_CState247, CJumpStatementType}:            _CGotoState113Action,
	{_CState249, CLparamToken}:                  _CGotoState310Action,
	{_CState250, CIdentifierToken}:              _CGotoState130Action,
	{_CState250, CConstantToken}:                _CGotoState77Action,
	{_CState250, CStringLiteralToken}:           _CGotoState96Action,
	{_CState250, CSizeofToken}:                  _CGotoState95Action,
	{_CState250, CIncOpToken}:                   _CGotoState87Action,
	{_CState250, CDecOpToken}:                   _CGotoState79Action,
	{_CState250, CLparamToken}:                  _CGotoState88Action,
	{_CState250, CSemicolonToken}:               _CGotoState94Action,
	{_CState250, CMulToken}:                     _CGotoState90Action,
	{_CState250, CMinusToken}:                   _CGotoState89Action,
	{_CState250, CPlusToken}:                    _CGotoState91Action,
	{_CState250, CAndToken}:                     _CGotoState74Action,
	{_CState250, CExclaimToken}:                 _CGotoState82Action,
	{_CState250, CTildaToken}:                   _CGotoState98Action,
	{_CState250, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState250, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState250, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState250, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState250, CCastExpressionType}:           _CGotoState103Action,
	{_CState250, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState250, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState250, CShiftExpressionType}:          _CGotoState122Action,
	{_CState250, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState250, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState250, CAndExpressionType}:            _CGotoState101Action,
	{_CState250, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState250, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState250, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState250, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState250, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState250, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState250, CExpressionType}:               _CGotoState109Action,
	{_CState250, CExpressionStatementType}:      _CGotoState311Action,
	{_CState253, CRparamToken}:                  _CGotoState312Action,
	{_CState253, CCommaToken}:                   _CGotoState186Action,
	{_CState255, CTypeNameToken}:                _CGotoState20Action,
	{_CState255, CTypedefToken}:                 _CGotoState19Action,
	{_CState255, CExternToken}:                  _CGotoState7Action,
	{_CState255, CStaticToken}:                  _CGotoState17Action,
	{_CState255, CAutoToken}:                    _CGotoState2Action,
	{_CState255, CRegisterToken}:                _CGotoState14Action,
	{_CState255, CCharToken}:                    _CGotoState3Action,
	{_CState255, CShortToken}:                   _CGotoState15Action,
	{_CState255, CIntToken}:                     _CGotoState10Action,
	{_CState255, CLongToken}:                    _CGotoState11Action,
	{_CState255, CSignedToken}:                  _CGotoState16Action,
	{_CState255, CUnsignedToken}:                _CGotoState22Action,
	{_CState255, CFloatToken}:                   _CGotoState8Action,
	{_CState255, CDoubleToken}:                  _CGotoState5Action,
	{_CState255, CConstToken}:                   _CGotoState4Action,
	{_CState255, CVolatileToken}:                _CGotoState24Action,
	{_CState255, CVoidToken}:                    _CGotoState23Action,
	{_CState255, CStructToken}:                  _CGotoState18Action,
	{_CState255, CUnionToken}:                   _CGotoState21Action,
	{_CState255, CEnumToken}:                    _CGotoState6Action,
	{_CState255, CLparamToken}:                  _CGotoState255Action,
	{_CState255, CRparamToken}:                  _CGotoState293Action,
	{_CState255, CLbraceToken}:                  _CGotoState223Action,
	{_CState255, CMulToken}:                     _CGotoState13Action,
	{_CState255, CDeclarationSpecifiersType}:    _CGotoState137Action,
	{_CState255, CStorageClassSpecifierType}:    _CGotoState33Action,
	{_CState255, CTypeSpecifierType}:            _CGotoState37Action,
	{_CState255, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState255, CStructOrUnionType}:            _CGotoState34Action,
	{_CState255, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState255, CTypeQualifierType}:            _CGotoState36Action,
	{_CState255, CPointerType}:                  _CGotoState257Action,
	{_CState255, CParameterTypeListType}:        _CGotoState295Action,
	{_CState255, CParameterListType}:            _CGotoState140Action,
	{_CState255, CParameterDeclarationType}:     _CGotoState139Action,
	{_CState255, CAbstractDeclaratorType}:       _CGotoState294Action,
	{_CState255, CDirectAbstractDeclaratorType}: _CGotoState227Action,
	{_CState257, CLparamToken}:                  _CGotoState255Action,
	{_CState257, CLbraceToken}:                  _CGotoState223Action,
	{_CState257, CDirectAbstractDeclaratorType}: _CGotoState298Action,
	{_CState258, CIdentifierToken}:              _CGotoState130Action,
	{_CState258, CConstantToken}:                _CGotoState77Action,
	{_CState258, CStringLiteralToken}:           _CGotoState96Action,
	{_CState258, CSizeofToken}:                  _CGotoState95Action,
	{_CState258, CIncOpToken}:                   _CGotoState87Action,
	{_CState258, CDecOpToken}:                   _CGotoState79Action,
	{_CState258, CLparamToken}:                  _CGotoState88Action,
	{_CState258, CMulToken}:                     _CGotoState90Action,
	{_CState258, CMinusToken}:                   _CGotoState89Action,
	{_CState258, CPlusToken}:                    _CGotoState91Action,
	{_CState258, CAndToken}:                     _CGotoState74Action,
	{_CState258, CExclaimToken}:                 _CGotoState82Action,
	{_CState258, CTildaToken}:                   _CGotoState98Action,
	{_CState258, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState258, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState258, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState258, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState258, CCastExpressionType}:           _CGotoState313Action,
	{_CState260, CRparamToken}:                  _CGotoState314Action,
	{_CState261, CRparamToken}:                  _CGotoState315Action,
	{_CState261, CCommaToken}:                   _CGotoState186Action,
	{_CState262, CRparamToken}:                  _CGotoState316Action,
	{_CState262, CCommaToken}:                   _CGotoState186Action,
	{_CState263, CMulToken}:                     _CGotoState194Action,
	{_CState263, CDivToken}:                     _CGotoState192Action,
	{_CState263, CModToken}:                     _CGotoState193Action,
	{_CState264, CMulToken}:                     _CGotoState194Action,
	{_CState264, CDivToken}:                     _CGotoState192Action,
	{_CState264, CModToken}:                     _CGotoState193Action,
	{_CState265, CEqOpToken}:                    _CGotoState183Action,
	{_CState265, CNeOpToken}:                    _CGotoState184Action,
	{_CState267, CLeOpToken}:                    _CGotoState203Action,
	{_CState267, CGeOpToken}:                    _CGotoState201Action,
	{_CState267, CLtToken}:                      _CGotoState204Action,
	{_CState267, CGtToken}:                      _CGotoState202Action,
	{_CState268, CLeOpToken}:                    _CGotoState203Action,
	{_CState268, CGeOpToken}:                    _CGotoState201Action,
	{_CState268, CLtToken}:                      _CGotoState204Action,
	{_CState268, CGtToken}:                      _CGotoState202Action,
	{_CState269, CAndToken}:                     _CGotoState180Action,
	{_CState271, CHatToken}:                     _CGotoState185Action,
	{_CState272, COrToken}:                      _CGotoState188Action,
	{_CState273, CAndOpToken}:                   _CGotoState189Action,
	{_CState274, CColonToken}:                   _CGotoState317Action,
	{_CState274, CCommaToken}:                   _CGotoState186Action,
	{_CState279, CRbraceToken}:                  _CGotoState318Action,
	{_CState279, CCommaToken}:                   _CGotoState186Action,
	{_CState281, CRparamToken}:                  _CGotoState320Action,
	{_CState281, CCommaToken}:                   _CGotoState319Action,
	{_CState284, CLeftOpToken}:                  _CGotoState205Action,
	{_CState284, CRightOpToken}:                 _CGotoState206Action,
	{_CState285, CLeftOpToken}:                  _CGotoState205Action,
	{_CState285, CRightOpToken}:                 _CGotoState206Action,
	{_CState286, CLeftOpToken}:                  _CGotoState205Action,
	{_CState286, CRightOpToken}:                 _CGotoState206Action,
	{_CState287, CLeftOpToken}:                  _CGotoState205Action,
	{_CState287, CRightOpToken}:                 _CGotoState206Action,
	{_CState288, CMinusToken}:                   _CGotoState178Action,
	{_CState288, CPlusToken}:                    _CGotoState179Action,
	{_CState289, CMinusToken}:                   _CGotoState178Action,
	{_CState289, CPlusToken}:                    _CGotoState179Action,
	{_CState292, CRbraceToken}:                  _CGotoState321Action,
	{_CState294, CRparamToken}:                  _CGotoState322Action,
	{_CState295, CRparamToken}:                  _CGotoState323Action,
	{_CState296, CIdentifierToken}:              _CGotoState130Action,
	{_CState296, CConstantToken}:                _CGotoState77Action,
	{_CState296, CStringLiteralToken}:           _CGotoState96Action,
	{_CState296, CSizeofToken}:                  _CGotoState95Action,
	{_CState296, CIncOpToken}:                   _CGotoState87Action,
	{_CState296, CDecOpToken}:                   _CGotoState79Action,
	{_CState296, CLparamToken}:                  _CGotoState88Action,
	{_CState296, CRbraceToken}:                  _CGotoState324Action,
	{_CState296, CMulToken}:                     _CGotoState90Action,
	{_CState296, CMinusToken}:                   _CGotoState89Action,
	{_CState296, CPlusToken}:                    _CGotoState91Action,
	{_CState296, CAndToken}:                     _CGotoState74Action,
	{_CState296, CExclaimToken}:                 _CGotoState82Action,
	{_CState296, CTildaToken}:                   _CGotoState98Action,
	{_CState296, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState296, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState296, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState296, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState296, CCastExpressionType}:           _CGotoState103Action,
	{_CState296, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState296, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState296, CShiftExpressionType}:          _CGotoState122Action,
	{_CState296, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState296, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState296, CAndExpressionType}:            _CGotoState101Action,
	{_CState296, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState296, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState296, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState296, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState296, CConditionalExpressionType}:    _CGotoState132Action,
	{_CState296, CConstantExpressionType}:       _CGotoState325Action,
	{_CState297, CTypeNameToken}:                _CGotoState20Action,
	{_CState297, CTypedefToken}:                 _CGotoState19Action,
	{_CState297, CExternToken}:                  _CGotoState7Action,
	{_CState297, CStaticToken}:                  _CGotoState17Action,
	{_CState297, CAutoToken}:                    _CGotoState2Action,
	{_CState297, CRegisterToken}:                _CGotoState14Action,
	{_CState297, CCharToken}:                    _CGotoState3Action,
	{_CState297, CShortToken}:                   _CGotoState15Action,
	{_CState297, CIntToken}:                     _CGotoState10Action,
	{_CState297, CLongToken}:                    _CGotoState11Action,
	{_CState297, CSignedToken}:                  _CGotoState16Action,
	{_CState297, CUnsignedToken}:                _CGotoState22Action,
	{_CState297, CFloatToken}:                   _CGotoState8Action,
	{_CState297, CDoubleToken}:                  _CGotoState5Action,
	{_CState297, CConstToken}:                   _CGotoState4Action,
	{_CState297, CVolatileToken}:                _CGotoState24Action,
	{_CState297, CVoidToken}:                    _CGotoState23Action,
	{_CState297, CStructToken}:                  _CGotoState18Action,
	{_CState297, CUnionToken}:                   _CGotoState21Action,
	{_CState297, CEnumToken}:                    _CGotoState6Action,
	{_CState297, CRparamToken}:                  _CGotoState326Action,
	{_CState297, CDeclarationSpecifiersType}:    _CGotoState137Action,
	{_CState297, CStorageClassSpecifierType}:    _CGotoState33Action,
	{_CState297, CTypeSpecifierType}:            _CGotoState37Action,
	{_CState297, CStructOrUnionSpecifierType}:   _CGotoState35Action,
	{_CState297, CStructOrUnionType}:            _CGotoState34Action,
	{_CState297, CEnumSpecifierType}:            _CGotoState29Action,
	{_CState297, CTypeQualifierType}:            _CGotoState36Action,
	{_CState297, CParameterTypeListType}:        _CGotoState327Action,
	{_CState297, CParameterListType}:            _CGotoState140Action,
	{_CState297, CParameterDeclarationType}:     _CGotoState139Action,
	{_CState298, CLparamToken}:                  _CGotoState297Action,
	{_CState298, CLbraceToken}:                  _CGotoState296Action,
	{_CState304, CIdentifierToken}:              _CGotoState130Action,
	{_CState304, CConstantToken}:                _CGotoState77Action,
	{_CState304, CStringLiteralToken}:           _CGotoState96Action,
	{_CState304, CSizeofToken}:                  _CGotoState95Action,
	{_CState304, CIncOpToken}:                   _CGotoState87Action,
	{_CState304, CDecOpToken}:                   _CGotoState79Action,
	{_CState304, CLparamToken}:                  _CGotoState88Action,
	{_CState304, CMulToken}:                     _CGotoState90Action,
	{_CState304, CMinusToken}:                   _CGotoState89Action,
	{_CState304, CPlusToken}:                    _CGotoState91Action,
	{_CState304, CAndToken}:                     _CGotoState74Action,
	{_CState304, CExclaimToken}:                 _CGotoState82Action,
	{_CState304, CTildaToken}:                   _CGotoState98Action,
	{_CState304, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState304, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState304, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState304, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState304, CCastExpressionType}:           _CGotoState103Action,
	{_CState304, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState304, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState304, CShiftExpressionType}:          _CGotoState122Action,
	{_CState304, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState304, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState304, CAndExpressionType}:            _CGotoState101Action,
	{_CState304, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState304, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState304, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState304, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState304, CConditionalExpressionType}:    _CGotoState132Action,
	{_CState304, CConstantExpressionType}:       _CGotoState328Action,
	{_CState305, CIdentifierToken}:              _CGotoState9Action,
	{_CState305, CLparamToken}:                  _CGotoState12Action,
	{_CState305, CColonToken}:                   _CGotoState234Action,
	{_CState305, CMulToken}:                     _CGotoState13Action,
	{_CState305, CStructDeclaratorType}:         _CGotoState329Action,
	{_CState305, CDeclaratorType}:               _CGotoState235Action,
	{_CState305, CDirectDeclaratorType}:         _CGotoState28Action,
	{_CState305, CPointerType}:                  _CGotoState32Action,
	{_CState307, CIdentifierToken}:              _CGotoState130Action,
	{_CState307, CConstantToken}:                _CGotoState77Action,
	{_CState307, CStringLiteralToken}:           _CGotoState96Action,
	{_CState307, CSizeofToken}:                  _CGotoState95Action,
	{_CState307, CIncOpToken}:                   _CGotoState87Action,
	{_CState307, CDecOpToken}:                   _CGotoState79Action,
	{_CState307, CLparamToken}:                  _CGotoState88Action,
	{_CState307, CLcurlToken}:                   _CGotoState152Action,
	{_CState307, CRcurlToken}:                   _CGotoState330Action,
	{_CState307, CMulToken}:                     _CGotoState90Action,
	{_CState307, CMinusToken}:                   _CGotoState89Action,
	{_CState307, CPlusToken}:                    _CGotoState91Action,
	{_CState307, CAndToken}:                     _CGotoState74Action,
	{_CState307, CExclaimToken}:                 _CGotoState82Action,
	{_CState307, CTildaToken}:                   _CGotoState98Action,
	{_CState307, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState307, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState307, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState307, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState307, CCastExpressionType}:           _CGotoState103Action,
	{_CState307, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState307, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState307, CShiftExpressionType}:          _CGotoState122Action,
	{_CState307, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState307, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState307, CAndExpressionType}:            _CGotoState101Action,
	{_CState307, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState307, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState307, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState307, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState307, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState307, CAssignmentExpressionType}:     _CGotoState153Action,
	{_CState307, CInitializerType}:              _CGotoState331Action,
	{_CState310, CIdentifierToken}:              _CGotoState130Action,
	{_CState310, CConstantToken}:                _CGotoState77Action,
	{_CState310, CStringLiteralToken}:           _CGotoState96Action,
	{_CState310, CSizeofToken}:                  _CGotoState95Action,
	{_CState310, CIncOpToken}:                   _CGotoState87Action,
	{_CState310, CDecOpToken}:                   _CGotoState79Action,
	{_CState310, CLparamToken}:                  _CGotoState88Action,
	{_CState310, CMulToken}:                     _CGotoState90Action,
	{_CState310, CMinusToken}:                   _CGotoState89Action,
	{_CState310, CPlusToken}:                    _CGotoState91Action,
	{_CState310, CAndToken}:                     _CGotoState74Action,
	{_CState310, CExclaimToken}:                 _CGotoState82Action,
	{_CState310, CTildaToken}:                   _CGotoState98Action,
	{_CState310, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState310, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState310, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState310, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState310, CCastExpressionType}:           _CGotoState103Action,
	{_CState310, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState310, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState310, CShiftExpressionType}:          _CGotoState122Action,
	{_CState310, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState310, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState310, CAndExpressionType}:            _CGotoState101Action,
	{_CState310, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState310, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState310, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState310, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState310, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState310, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState310, CExpressionType}:               _CGotoState332Action,
	{_CState311, CIdentifierToken}:              _CGotoState130Action,
	{_CState311, CConstantToken}:                _CGotoState77Action,
	{_CState311, CStringLiteralToken}:           _CGotoState96Action,
	{_CState311, CSizeofToken}:                  _CGotoState95Action,
	{_CState311, CIncOpToken}:                   _CGotoState87Action,
	{_CState311, CDecOpToken}:                   _CGotoState79Action,
	{_CState311, CLparamToken}:                  _CGotoState88Action,
	{_CState311, CRparamToken}:                  _CGotoState333Action,
	{_CState311, CMulToken}:                     _CGotoState90Action,
	{_CState311, CMinusToken}:                   _CGotoState89Action,
	{_CState311, CPlusToken}:                    _CGotoState91Action,
	{_CState311, CAndToken}:                     _CGotoState74Action,
	{_CState311, CExclaimToken}:                 _CGotoState82Action,
	{_CState311, CTildaToken}:                   _CGotoState98Action,
	{_CState311, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState311, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState311, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState311, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState311, CCastExpressionType}:           _CGotoState103Action,
	{_CState311, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState311, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState311, CShiftExpressionType}:          _CGotoState122Action,
	{_CState311, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState311, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState311, CAndExpressionType}:            _CGotoState101Action,
	{_CState311, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState311, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState311, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState311, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState311, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState311, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState311, CExpressionType}:               _CGotoState334Action,
	{_CState312, CIdentifierToken}:              _CGotoState85Action,
	{_CState312, CConstantToken}:                _CGotoState77Action,
	{_CState312, CStringLiteralToken}:           _CGotoState96Action,
	{_CState312, CSizeofToken}:                  _CGotoState95Action,
	{_CState312, CIncOpToken}:                   _CGotoState87Action,
	{_CState312, CDecOpToken}:                   _CGotoState79Action,
	{_CState312, CCaseToken}:                    _CGotoState76Action,
	{_CState312, CDefaultToken}:                 _CGotoState80Action,
	{_CState312, CIfToken}:                      _CGotoState86Action,
	{_CState312, CSwitchToken}:                  _CGotoState97Action,
	{_CState312, CWhileToken}:                   _CGotoState99Action,
	{_CState312, CDoToken}:                      _CGotoState81Action,
	{_CState312, CForToken}:                     _CGotoState83Action,
	{_CState312, CGotoToken}:                    _CGotoState84Action,
	{_CState312, CContinueToken}:                _CGotoState78Action,
	{_CState312, CBreakToken}:                   _CGotoState75Action,
	{_CState312, CReturnToken}:                  _CGotoState93Action,
	{_CState312, CLparamToken}:                  _CGotoState88Action,
	{_CState312, CLcurlToken}:                   _CGotoState49Action,
	{_CState312, CSemicolonToken}:               _CGotoState94Action,
	{_CState312, CMulToken}:                     _CGotoState90Action,
	{_CState312, CMinusToken}:                   _CGotoState89Action,
	{_CState312, CPlusToken}:                    _CGotoState91Action,
	{_CState312, CAndToken}:                     _CGotoState74Action,
	{_CState312, CExclaimToken}:                 _CGotoState82Action,
	{_CState312, CTildaToken}:                   _CGotoState98Action,
	{_CState312, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState312, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState312, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState312, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState312, CCastExpressionType}:           _CGotoState103Action,
	{_CState312, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState312, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState312, CShiftExpressionType}:          _CGotoState122Action,
	{_CState312, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState312, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState312, CAndExpressionType}:            _CGotoState101Action,
	{_CState312, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState312, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState312, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState312, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState312, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState312, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState312, CExpressionType}:               _CGotoState109Action,
	{_CState312, CStatementType}:                _CGotoState337Action,
	{_CState312, CLabeledStatementType}:         _CGotoState114Action,
	{_CState312, CCompoundStatementType}:        _CGotoState104Action,
	{_CState312, CExpressionStatementType}:      _CGotoState110Action,
	{_CState312, CSelectionStatementType}:       _CGotoState121Action,
	{_CState312, CIterationStatementType}:       _CGotoState112Action,
	{_CState312, CJumpStatementType}:            _CGotoState113Action,
	{_CState315, CIdentifierToken}:              _CGotoState85Action,
	{_CState315, CConstantToken}:                _CGotoState77Action,
	{_CState315, CStringLiteralToken}:           _CGotoState96Action,
	{_CState315, CSizeofToken}:                  _CGotoState95Action,
	{_CState315, CIncOpToken}:                   _CGotoState87Action,
	{_CState315, CDecOpToken}:                   _CGotoState79Action,
	{_CState315, CCaseToken}:                    _CGotoState76Action,
	{_CState315, CDefaultToken}:                 _CGotoState80Action,
	{_CState315, CIfToken}:                      _CGotoState86Action,
	{_CState315, CSwitchToken}:                  _CGotoState97Action,
	{_CState315, CWhileToken}:                   _CGotoState99Action,
	{_CState315, CDoToken}:                      _CGotoState81Action,
	{_CState315, CForToken}:                     _CGotoState83Action,
	{_CState315, CGotoToken}:                    _CGotoState84Action,
	{_CState315, CContinueToken}:                _CGotoState78Action,
	{_CState315, CBreakToken}:                   _CGotoState75Action,
	{_CState315, CReturnToken}:                  _CGotoState93Action,
	{_CState315, CLparamToken}:                  _CGotoState88Action,
	{_CState315, CLcurlToken}:                   _CGotoState49Action,
	{_CState315, CSemicolonToken}:               _CGotoState94Action,
	{_CState315, CMulToken}:                     _CGotoState90Action,
	{_CState315, CMinusToken}:                   _CGotoState89Action,
	{_CState315, CPlusToken}:                    _CGotoState91Action,
	{_CState315, CAndToken}:                     _CGotoState74Action,
	{_CState315, CExclaimToken}:                 _CGotoState82Action,
	{_CState315, CTildaToken}:                   _CGotoState98Action,
	{_CState315, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState315, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState315, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState315, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState315, CCastExpressionType}:           _CGotoState103Action,
	{_CState315, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState315, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState315, CShiftExpressionType}:          _CGotoState122Action,
	{_CState315, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState315, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState315, CAndExpressionType}:            _CGotoState101Action,
	{_CState315, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState315, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState315, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState315, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState315, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState315, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState315, CExpressionType}:               _CGotoState109Action,
	{_CState315, CStatementType}:                _CGotoState338Action,
	{_CState315, CLabeledStatementType}:         _CGotoState114Action,
	{_CState315, CCompoundStatementType}:        _CGotoState104Action,
	{_CState315, CExpressionStatementType}:      _CGotoState110Action,
	{_CState315, CSelectionStatementType}:       _CGotoState121Action,
	{_CState315, CIterationStatementType}:       _CGotoState112Action,
	{_CState315, CJumpStatementType}:            _CGotoState113Action,
	{_CState316, CIdentifierToken}:              _CGotoState85Action,
	{_CState316, CConstantToken}:                _CGotoState77Action,
	{_CState316, CStringLiteralToken}:           _CGotoState96Action,
	{_CState316, CSizeofToken}:                  _CGotoState95Action,
	{_CState316, CIncOpToken}:                   _CGotoState87Action,
	{_CState316, CDecOpToken}:                   _CGotoState79Action,
	{_CState316, CCaseToken}:                    _CGotoState76Action,
	{_CState316, CDefaultToken}:                 _CGotoState80Action,
	{_CState316, CIfToken}:                      _CGotoState86Action,
	{_CState316, CSwitchToken}:                  _CGotoState97Action,
	{_CState316, CWhileToken}:                   _CGotoState99Action,
	{_CState316, CDoToken}:                      _CGotoState81Action,
	{_CState316, CForToken}:                     _CGotoState83Action,
	{_CState316, CGotoToken}:                    _CGotoState84Action,
	{_CState316, CContinueToken}:                _CGotoState78Action,
	{_CState316, CBreakToken}:                   _CGotoState75Action,
	{_CState316, CReturnToken}:                  _CGotoState93Action,
	{_CState316, CLparamToken}:                  _CGotoState88Action,
	{_CState316, CLcurlToken}:                   _CGotoState49Action,
	{_CState316, CSemicolonToken}:               _CGotoState94Action,
	{_CState316, CMulToken}:                     _CGotoState90Action,
	{_CState316, CMinusToken}:                   _CGotoState89Action,
	{_CState316, CPlusToken}:                    _CGotoState91Action,
	{_CState316, CAndToken}:                     _CGotoState74Action,
	{_CState316, CExclaimToken}:                 _CGotoState82Action,
	{_CState316, CTildaToken}:                   _CGotoState98Action,
	{_CState316, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState316, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState316, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState316, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState316, CCastExpressionType}:           _CGotoState103Action,
	{_CState316, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState316, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState316, CShiftExpressionType}:          _CGotoState122Action,
	{_CState316, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState316, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState316, CAndExpressionType}:            _CGotoState101Action,
	{_CState316, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState316, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState316, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState316, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState316, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState316, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState316, CExpressionType}:               _CGotoState109Action,
	{_CState316, CStatementType}:                _CGotoState339Action,
	{_CState316, CLabeledStatementType}:         _CGotoState114Action,
	{_CState316, CCompoundStatementType}:        _CGotoState104Action,
	{_CState316, CExpressionStatementType}:      _CGotoState110Action,
	{_CState316, CSelectionStatementType}:       _CGotoState121Action,
	{_CState316, CIterationStatementType}:       _CGotoState112Action,
	{_CState316, CJumpStatementType}:            _CGotoState113Action,
	{_CState317, CIdentifierToken}:              _CGotoState130Action,
	{_CState317, CConstantToken}:                _CGotoState77Action,
	{_CState317, CStringLiteralToken}:           _CGotoState96Action,
	{_CState317, CSizeofToken}:                  _CGotoState95Action,
	{_CState317, CIncOpToken}:                   _CGotoState87Action,
	{_CState317, CDecOpToken}:                   _CGotoState79Action,
	{_CState317, CLparamToken}:                  _CGotoState88Action,
	{_CState317, CMulToken}:                     _CGotoState90Action,
	{_CState317, CMinusToken}:                   _CGotoState89Action,
	{_CState317, CPlusToken}:                    _CGotoState91Action,
	{_CState317, CAndToken}:                     _CGotoState74Action,
	{_CState317, CExclaimToken}:                 _CGotoState82Action,
	{_CState317, CTildaToken}:                   _CGotoState98Action,
	{_CState317, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState317, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState317, CUnaryExpressionType}:          _CGotoState134Action,
	{_CState317, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState317, CCastExpressionType}:           _CGotoState103Action,
	{_CState317, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState317, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState317, CShiftExpressionType}:          _CGotoState122Action,
	{_CState317, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState317, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState317, CAndExpressionType}:            _CGotoState101Action,
	{_CState317, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState317, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState317, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState317, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState317, CConditionalExpressionType}:    _CGotoState340Action,
	{_CState319, CIdentifierToken}:              _CGotoState130Action,
	{_CState319, CConstantToken}:                _CGotoState77Action,
	{_CState319, CStringLiteralToken}:           _CGotoState96Action,
	{_CState319, CSizeofToken}:                  _CGotoState95Action,
	{_CState319, CIncOpToken}:                   _CGotoState87Action,
	{_CState319, CDecOpToken}:                   _CGotoState79Action,
	{_CState319, CLparamToken}:                  _CGotoState88Action,
	{_CState319, CMulToken}:                     _CGotoState90Action,
	{_CState319, CMinusToken}:                   _CGotoState89Action,
	{_CState319, CPlusToken}:                    _CGotoState91Action,
	{_CState319, CAndToken}:                     _CGotoState74Action,
	{_CState319, CExclaimToken}:                 _CGotoState82Action,
	{_CState319, CTildaToken}:                   _CGotoState98Action,
	{_CState319, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState319, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState319, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState319, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState319, CCastExpressionType}:           _CGotoState103Action,
	{_CState319, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState319, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState319, CShiftExpressionType}:          _CGotoState122Action,
	{_CState319, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState319, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState319, CAndExpressionType}:            _CGotoState101Action,
	{_CState319, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState319, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState319, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState319, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState319, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState319, CAssignmentExpressionType}:     _CGotoState341Action,
	{_CState325, CRbraceToken}:                  _CGotoState342Action,
	{_CState327, CRparamToken}:                  _CGotoState343Action,
	{_CState332, CRparamToken}:                  _CGotoState344Action,
	{_CState332, CCommaToken}:                   _CGotoState186Action,
	{_CState333, CIdentifierToken}:              _CGotoState85Action,
	{_CState333, CConstantToken}:                _CGotoState77Action,
	{_CState333, CStringLiteralToken}:           _CGotoState96Action,
	{_CState333, CSizeofToken}:                  _CGotoState95Action,
	{_CState333, CIncOpToken}:                   _CGotoState87Action,
	{_CState333, CDecOpToken}:                   _CGotoState79Action,
	{_CState333, CCaseToken}:                    _CGotoState76Action,
	{_CState333, CDefaultToken}:                 _CGotoState80Action,
	{_CState333, CIfToken}:                      _CGotoState86Action,
	{_CState333, CSwitchToken}:                  _CGotoState97Action,
	{_CState333, CWhileToken}:                   _CGotoState99Action,
	{_CState333, CDoToken}:                      _CGotoState81Action,
	{_CState333, CForToken}:                     _CGotoState83Action,
	{_CState333, CGotoToken}:                    _CGotoState84Action,
	{_CState333, CContinueToken}:                _CGotoState78Action,
	{_CState333, CBreakToken}:                   _CGotoState75Action,
	{_CState333, CReturnToken}:                  _CGotoState93Action,
	{_CState333, CLparamToken}:                  _CGotoState88Action,
	{_CState333, CLcurlToken}:                   _CGotoState49Action,
	{_CState333, CSemicolonToken}:               _CGotoState94Action,
	{_CState333, CMulToken}:                     _CGotoState90Action,
	{_CState333, CMinusToken}:                   _CGotoState89Action,
	{_CState333, CPlusToken}:                    _CGotoState91Action,
	{_CState333, CAndToken}:                     _CGotoState74Action,
	{_CState333, CExclaimToken}:                 _CGotoState82Action,
	{_CState333, CTildaToken}:                   _CGotoState98Action,
	{_CState333, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState333, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState333, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState333, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState333, CCastExpressionType}:           _CGotoState103Action,
	{_CState333, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState333, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState333, CShiftExpressionType}:          _CGotoState122Action,
	{_CState333, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState333, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState333, CAndExpressionType}:            _CGotoState101Action,
	{_CState333, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState333, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState333, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState333, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState333, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState333, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState333, CExpressionType}:               _CGotoState109Action,
	{_CState333, CStatementType}:                _CGotoState345Action,
	{_CState333, CLabeledStatementType}:         _CGotoState114Action,
	{_CState333, CCompoundStatementType}:        _CGotoState104Action,
	{_CState333, CExpressionStatementType}:      _CGotoState110Action,
	{_CState333, CSelectionStatementType}:       _CGotoState121Action,
	{_CState333, CIterationStatementType}:       _CGotoState112Action,
	{_CState333, CJumpStatementType}:            _CGotoState113Action,
	{_CState334, CRparamToken}:                  _CGotoState346Action,
	{_CState334, CCommaToken}:                   _CGotoState186Action,
	{_CState335, CElseToken}:                    _CGotoState347Action,
	{_CState336, CElseToken}:                    _CGotoState347Action,
	{_CState337, CElseToken}:                    _CGotoState347Action,
	{_CState344, CSemicolonToken}:               _CGotoState348Action,
	{_CState346, CIdentifierToken}:              _CGotoState85Action,
	{_CState346, CConstantToken}:                _CGotoState77Action,
	{_CState346, CStringLiteralToken}:           _CGotoState96Action,
	{_CState346, CSizeofToken}:                  _CGotoState95Action,
	{_CState346, CIncOpToken}:                   _CGotoState87Action,
	{_CState346, CDecOpToken}:                   _CGotoState79Action,
	{_CState346, CCaseToken}:                    _CGotoState76Action,
	{_CState346, CDefaultToken}:                 _CGotoState80Action,
	{_CState346, CIfToken}:                      _CGotoState86Action,
	{_CState346, CSwitchToken}:                  _CGotoState97Action,
	{_CState346, CWhileToken}:                   _CGotoState99Action,
	{_CState346, CDoToken}:                      _CGotoState81Action,
	{_CState346, CForToken}:                     _CGotoState83Action,
	{_CState346, CGotoToken}:                    _CGotoState84Action,
	{_CState346, CContinueToken}:                _CGotoState78Action,
	{_CState346, CBreakToken}:                   _CGotoState75Action,
	{_CState346, CReturnToken}:                  _CGotoState93Action,
	{_CState346, CLparamToken}:                  _CGotoState88Action,
	{_CState346, CLcurlToken}:                   _CGotoState49Action,
	{_CState346, CSemicolonToken}:               _CGotoState94Action,
	{_CState346, CMulToken}:                     _CGotoState90Action,
	{_CState346, CMinusToken}:                   _CGotoState89Action,
	{_CState346, CPlusToken}:                    _CGotoState91Action,
	{_CState346, CAndToken}:                     _CGotoState74Action,
	{_CState346, CExclaimToken}:                 _CGotoState82Action,
	{_CState346, CTildaToken}:                   _CGotoState98Action,
	{_CState346, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState346, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState346, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState346, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState346, CCastExpressionType}:           _CGotoState103Action,
	{_CState346, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState346, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState346, CShiftExpressionType}:          _CGotoState122Action,
	{_CState346, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState346, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState346, CAndExpressionType}:            _CGotoState101Action,
	{_CState346, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState346, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState346, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState346, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState346, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState346, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState346, CExpressionType}:               _CGotoState109Action,
	{_CState346, CStatementType}:                _CGotoState349Action,
	{_CState346, CLabeledStatementType}:         _CGotoState114Action,
	{_CState346, CCompoundStatementType}:        _CGotoState104Action,
	{_CState346, CExpressionStatementType}:      _CGotoState110Action,
	{_CState346, CSelectionStatementType}:       _CGotoState121Action,
	{_CState346, CIterationStatementType}:       _CGotoState112Action,
	{_CState346, CJumpStatementType}:            _CGotoState113Action,
	{_CState347, CIdentifierToken}:              _CGotoState85Action,
	{_CState347, CConstantToken}:                _CGotoState77Action,
	{_CState347, CStringLiteralToken}:           _CGotoState96Action,
	{_CState347, CSizeofToken}:                  _CGotoState95Action,
	{_CState347, CIncOpToken}:                   _CGotoState87Action,
	{_CState347, CDecOpToken}:                   _CGotoState79Action,
	{_CState347, CCaseToken}:                    _CGotoState76Action,
	{_CState347, CDefaultToken}:                 _CGotoState80Action,
	{_CState347, CIfToken}:                      _CGotoState86Action,
	{_CState347, CSwitchToken}:                  _CGotoState97Action,
	{_CState347, CWhileToken}:                   _CGotoState99Action,
	{_CState347, CDoToken}:                      _CGotoState81Action,
	{_CState347, CForToken}:                     _CGotoState83Action,
	{_CState347, CGotoToken}:                    _CGotoState84Action,
	{_CState347, CContinueToken}:                _CGotoState78Action,
	{_CState347, CBreakToken}:                   _CGotoState75Action,
	{_CState347, CReturnToken}:                  _CGotoState93Action,
	{_CState347, CLparamToken}:                  _CGotoState88Action,
	{_CState347, CLcurlToken}:                   _CGotoState49Action,
	{_CState347, CSemicolonToken}:               _CGotoState94Action,
	{_CState347, CMulToken}:                     _CGotoState90Action,
	{_CState347, CMinusToken}:                   _CGotoState89Action,
	{_CState347, CPlusToken}:                    _CGotoState91Action,
	{_CState347, CAndToken}:                     _CGotoState74Action,
	{_CState347, CExclaimToken}:                 _CGotoState82Action,
	{_CState347, CTildaToken}:                   _CGotoState98Action,
	{_CState347, CPrimaryExpressionType}:        _CGotoState119Action,
	{_CState347, CPostfixExpressionType}:        _CGotoState118Action,
	{_CState347, CUnaryExpressionType}:          _CGotoState125Action,
	{_CState347, CUnaryOperatorType}:            _CGotoState126Action,
	{_CState347, CCastExpressionType}:           _CGotoState103Action,
	{_CState347, CMultiplicativeExpressionType}: _CGotoState117Action,
	{_CState347, CAdditiveExpressionType}:       _CGotoState100Action,
	{_CState347, CShiftExpressionType}:          _CGotoState122Action,
	{_CState347, CRelationalExpressionType}:     _CGotoState120Action,
	{_CState347, CEqualityExpressionType}:       _CGotoState107Action,
	{_CState347, CAndExpressionType}:            _CGotoState101Action,
	{_CState347, CExclusiveOrExpressionType}:    _CGotoState108Action,
	{_CState347, CInclusiveOrExpressionType}:    _CGotoState111Action,
	{_CState347, CLogicalAndExpressionType}:     _CGotoState115Action,
	{_CState347, CLogicalOrExpressionType}:      _CGotoState116Action,
	{_CState347, CConditionalExpressionType}:    _CGotoState105Action,
	{_CState347, CAssignmentExpressionType}:     _CGotoState102Action,
	{_CState347, CExpressionType}:               _CGotoState109Action,
	{_CState347, CStatementType}:                _CGotoState350Action,
	{_CState347, CLabeledStatementType}:         _CGotoState114Action,
	{_CState347, CCompoundStatementType}:        _CGotoState104Action,
	{_CState347, CExpressionStatementType}:      _CGotoState110Action,
	{_CState347, CSelectionStatementType}:       _CGotoState121Action,
	{_CState347, CIterationStatementType}:       _CGotoState112Action,
	{_CState347, CJumpStatementType}:            _CGotoState113Action,
	{_CState2, _CWildcardMarker}:                _CReduceDToStorageClassSpecifierAction,
	{_CState3, _CWildcardMarker}:                _CReduceBToTypeSpecifierAction,
	{_CState4, _CWildcardMarker}:                _CReduceAToTypeQualifierAction,
	{_CState5, _CWildcardMarker}:                _CReduceGToTypeSpecifierAction,
	{_CState7, _CWildcardMarker}:                _CReduceBToStorageClassSpecifierAction,
	{_CState8, _CWildcardMarker}:                _CReduceFToTypeSpecifierAction,
	{_CState9, _CWildcardMarker}:                _CReduceAToDirectDeclaratorAction,
	{_CState10, _CWildcardMarker}:               _CReduceDToTypeSpecifierAction,
	{_CState11, _CWildcardMarker}:               _CReduceEToTypeSpecifierAction,
	{_CState13, _CWildcardMarker}:               _CReduceAToPointerAction,
	{_CState14, _CWildcardMarker}:               _CReduceEToStorageClassSpecifierAction,
	{_CState15, _CWildcardMarker}:               _CReduceCToTypeSpecifierAction,
	{_CState16, _CWildcardMarker}:               _CReduceHToTypeSpecifierAction,
	{_CState17, _CWildcardMarker}:               _CReduceCToStorageClassSpecifierAction,
	{_CState18, _CWildcardMarker}:               _CReduceAToStructOrUnionAction,
	{_CState19, _CWildcardMarker}:               _CReduceAToStorageClassSpecifierAction,
	{_CState20, _CWildcardMarker}:               _CReduceLToTypeSpecifierAction,
	{_CState21, _CWildcardMarker}:               _CReduceBToStructOrUnionAction,
	{_CState22, _CWildcardMarker}:               _CReduceIToTypeSpecifierAction,
	{_CState23, _CWildcardMarker}:               _CReduceAToTypeSpecifierAction,
	{_CState24, _CWildcardMarker}:               _CReduceBToTypeQualifierAction,
	{_CState25, _CWildcardMarker}:               _CReduceBToExternalDeclarationAction,
	{_CState28, _CWildcardMarker}:               _CReduceBToDeclaratorAction,
	{_CState29, _CWildcardMarker}:               _CReduceKToTypeSpecifierAction,
	{_CState30, _CWildcardMarker}:               _CReduceAToTranslationUnitAction,
	{_CState31, _CWildcardMarker}:               _CReduceAToExternalDeclarationAction,
	{_CState33, _CWildcardMarker}:               _CReduceAToDeclarationSpecifiersAction,
	{_CState35, _CWildcardMarker}:               _CReduceJToTypeSpecifierAction,
	{_CState36, _CWildcardMarker}:               _CReduceEToDeclarationSpecifiersAction,
	{_CState37, _CWildcardMarker}:               _CReduceCToDeclarationSpecifiersAction,
	{_CState38, _CWildcardMarker}:               _CReduceBToTranslationUnitAction,
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
	{_CState60, _CWildcardMarker}:               _CReduceFToDeclarationSpecifiersAction,
	{_CState61, _CWildcardMarker}:               _CReduceDToDeclarationSpecifiersAction,
	{_CState63, _CWildcardMarker}:               _CReduceAToEnumeratorAction,
	{_CState64, _CWildcardMarker}:               _CReduceAToEnumeratorListAction,
	{_CState66, _CWildcardMarker}:               _CReduceBToDirectDeclaratorAction,
	{_CState67, _CWildcardMarker}:               _CReduceDToPointerAction,
	{_CState68, _CWildcardMarker}:               _CReduceBToTypeQualifierListAction,
	{_CState70, _CWildcardMarker}:               _CReduceBToFunctionDefinitionAction,
	{_CState73, _CWildcardMarker}:               _CReduceBToDeclarationAction,
	{_CState74, _CWildcardMarker}:               _CReduceAToUnaryOperatorAction,
	{_CState77, _CWildcardMarker}:               _CReduceBToPrimaryExpressionAction,
	{_CState82, _CWildcardMarker}:               _CReduceFToUnaryOperatorAction,
	{_CState85, _CWildcardMarker}:               _CReduceAToPrimaryExpressionAction,
	{_CState89, _CWildcardMarker}:               _CReduceDToUnaryOperatorAction,
	{_CState90, _CWildcardMarker}:               _CReduceBToUnaryOperatorAction,
	{_CState91, _CWildcardMarker}:               _CReduceCToUnaryOperatorAction,
	{_CState92, _CWildcardMarker}:               _CReduceAToCompoundStatementAction,
	{_CState94, _CWildcardMarker}:               _CReduceAToExpressionStatementAction,
	{_CState96, _CWildcardMarker}:               _CReduceCToPrimaryExpressionAction,
	{_CState98, _CWildcardMarker}:               _CReduceEToUnaryOperatorAction,
	{_CState100, _CWildcardMarker}:              _CReduceAToShiftExpressionAction,
	{_CState101, _CWildcardMarker}:              _CReduceAToExclusiveOrExpressionAction,
	{_CState102, _CWildcardMarker}:              _CReduceAToExpressionAction,
	{_CState103, _CWildcardMarker}:              _CReduceAToMultiplicativeExpressionAction,
	{_CState104, _CWildcardMarker}:              _CReduceBToStatementAction,
	{_CState105, _CWildcardMarker}:              _CReduceAToAssignmentExpressionAction,
	{_CState107, _CWildcardMarker}:              _CReduceAToAndExpressionAction,
	{_CState108, _CWildcardMarker}:              _CReduceAToInclusiveOrExpressionAction,
	{_CState110, _CWildcardMarker}:              _CReduceCToStatementAction,
	{_CState111, _CWildcardMarker}:              _CReduceAToLogicalAndExpressionAction,
	{_CState112, _CWildcardMarker}:              _CReduceEToStatementAction,
	{_CState113, _CWildcardMarker}:              _CReduceFToStatementAction,
	{_CState114, _CWildcardMarker}:              _CReduceAToStatementAction,
	{_CState115, _CWildcardMarker}:              _CReduceAToLogicalOrExpressionAction,
	{_CState116, _CWildcardMarker}:              _CReduceAToConditionalExpressionAction,
	{_CState117, _CWildcardMarker}:              _CReduceAToAdditiveExpressionAction,
	{_CState118, _CWildcardMarker}:              _CReduceAToUnaryExpressionAction,
	{_CState119, _CWildcardMarker}:              _CReduceAToPostfixExpressionAction,
	{_CState120, _CWildcardMarker}:              _CReduceAToEqualityExpressionAction,
	{_CState121, _CWildcardMarker}:              _CReduceDToStatementAction,
	{_CState122, _CWildcardMarker}:              _CReduceAToRelationalExpressionAction,
	{_CState123, _CWildcardMarker}:              _CReduceAToStatementListAction,
	{_CState125, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState127, _CWildcardMarker}:              _CReduceCToFunctionDefinitionAction,
	{_CState128, _CWildcardMarker}:              _CReduceBToDeclarationListAction,
	{_CState129, _CWildcardMarker}:              _CReduceAToInitDeclaratorAction,
	{_CState130, _CWildcardMarker}:              _CReduceAToPrimaryExpressionAction,
	{_CState131, _CWildcardMarker}:              _CReduceDToDirectDeclaratorAction,
	{_CState132, _CWildcardMarker}:              _CReduceAToConstantExpressionAction,
	{_CState134, _CWildcardMarker}:              _CReduceAToCastExpressionAction,
	{_CState135, _CWildcardMarker}:              _CReduceAToIdentifierListAction,
	{_CState136, _CWildcardMarker}:              _CReduceGToDirectDeclaratorAction,
	{_CState137, _CWildcardMarker}:              _CReduceCToParameterDeclarationAction,
	{_CState139, _CWildcardMarker}:              _CReduceAToParameterListAction,
	{_CState140, CRparamToken}:                  _CReduceAToParameterTypeListAction,
	{_CState144, _CWildcardMarker}:              _CReduceAToStructDeclarationListAction,
	{_CState146, _CWildcardMarker}:              _CReduceDToSpecifierQualifierListAction,
	{_CState147, _CWildcardMarker}:              _CReduceBToSpecifierQualifierListAction,
	{_CState151, _CWildcardMarker}:              _CReduceAToEnumSpecifierAction,
	{_CState153, _CWildcardMarker}:              _CReduceAToInitializerAction,
	{_CState154, _CWildcardMarker}:              _CReduceBToInitDeclaratorAction,
	{_CState155, _CWildcardMarker}:              _CReduceAToFunctionDefinitionAction,
	{_CState156, _CWildcardMarker}:              _CReduceBToInitDeclaratorListAction,
	{_CState157, _CWildcardMarker}:              _CReduceCToJumpStatementAction,
	{_CState159, _CWildcardMarker}:              _CReduceBToJumpStatementAction,
	{_CState161, _CWildcardMarker}:              _CReduceCToUnaryExpressionAction,
	{_CState168, _CWildcardMarker}:              _CReduceBToUnaryExpressionAction,
	{_CState170, CRparamToken}:                  _CReduceAToTypeNameAction,
	{_CState172, _CWildcardMarker}:              _CReduceDToJumpStatementAction,
	{_CState175, _CWildcardMarker}:              _CReduceEToUnaryExpressionAction,
	{_CState181, _CWildcardMarker}:              _CReduceCToCompoundStatementAction,
	{_CState187, _CWildcardMarker}:              _CReduceBToExpressionStatementAction,
	{_CState195, _CWildcardMarker}:              _CReduceHToPostfixExpressionAction,
	{_CState197, _CWildcardMarker}:              _CReduceGToPostfixExpressionAction,
	{_CState207, _CWildcardMarker}:              _CReduceBToCompoundStatementAction,
	{_CState208, _CWildcardMarker}:              _CReduceBToStatementListAction,
	{_CState209, _CWildcardMarker}:              _CReduceEToAssignmentOperatorAction,
	{_CState210, _CWildcardMarker}:              _CReduceIToAssignmentOperatorAction,
	{_CState211, _CWildcardMarker}:              _CReduceCToAssignmentOperatorAction,
	{_CState212, _CWildcardMarker}:              _CReduceAToAssignmentOperatorAction,
	{_CState213, _CWildcardMarker}:              _CReduceGToAssignmentOperatorAction,
	{_CState214, _CWildcardMarker}:              _CReduceDToAssignmentOperatorAction,
	{_CState215, _CWildcardMarker}:              _CReduceBToAssignmentOperatorAction,
	{_CState216, _CWildcardMarker}:              _CReduceKToAssignmentOperatorAction,
	{_CState217, _CWildcardMarker}:              _CReduceHToAssignmentOperatorAction,
	{_CState218, _CWildcardMarker}:              _CReduceFToAssignmentOperatorAction,
	{_CState219, _CWildcardMarker}:              _CReduceJToAssignmentOperatorAction,
	{_CState221, _CWildcardMarker}:              _CReduceDToUnaryExpressionAction,
	{_CState222, _CWildcardMarker}:              _CReduceCToDirectDeclaratorAction,
	{_CState225, _CWildcardMarker}:              _CReduceBToParameterDeclarationAction,
	{_CState226, _CWildcardMarker}:              _CReduceAToParameterDeclarationAction,
	{_CState227, _CWildcardMarker}:              _CReduceBToAbstractDeclaratorAction,
	{_CState228, _CWildcardMarker}:              _CReduceAToAbstractDeclaratorAction,
	{_CState230, _CWildcardMarker}:              _CReduceFToDirectDeclaratorAction,
	{_CState232, _CWildcardMarker}:              _CReduceEToDirectDeclaratorAction,
	{_CState235, _CWildcardMarker}:              _CReduceAToStructDeclaratorAction,
	{_CState236, _CWildcardMarker}:              _CReduceAToStructDeclaratorListAction,
	{_CState238, _CWildcardMarker}:              _CReduceBToStructOrUnionSpecifierAction,
	{_CState239, _CWildcardMarker}:              _CReduceBToStructDeclarationListAction,
	{_CState240, _CWildcardMarker}:              _CReduceCToSpecifierQualifierListAction,
	{_CState241, _CWildcardMarker}:              _CReduceAToSpecifierQualifierListAction,
	{_CState242, _CWildcardMarker}:              _CReduceBToEnumSpecifierAction,
	{_CState243, _CWildcardMarker}:              _CReduceBToEnumeratorAction,
	{_CState244, _CWildcardMarker}:              _CReduceBToEnumeratorListAction,
	{_CState245, _CWildcardMarker}:              _CReduceAToInitializerListAction,
	{_CState248, _CWildcardMarker}:              _CReduceCToLabeledStatementAction,
	{_CState251, _CWildcardMarker}:              _CReduceAToJumpStatementAction,
	{_CState252, _CWildcardMarker}:              _CReduceAToLabeledStatementAction,
	{_CState254, _CWildcardMarker}:              _CReduceDToPrimaryExpressionAction,
	{_CState256, CRparamToken}:                  _CReduceBToTypeNameAction,
	{_CState257, CRparamToken}:                  _CReduceAToAbstractDeclaratorAction,
	{_CState259, _CWildcardMarker}:              _CReduceEToJumpStatementAction,
	{_CState263, _CWildcardMarker}:              _CReduceCToAdditiveExpressionAction,
	{_CState264, _CWildcardMarker}:              _CReduceBToAdditiveExpressionAction,
	{_CState265, _CWildcardMarker}:              _CReduceBToAndExpressionAction,
	{_CState266, _CWildcardMarker}:              _CReduceDToCompoundStatementAction,
	{_CState267, _CWildcardMarker}:              _CReduceBToEqualityExpressionAction,
	{_CState268, _CWildcardMarker}:              _CReduceCToEqualityExpressionAction,
	{_CState269, _CWildcardMarker}:              _CReduceBToExclusiveOrExpressionAction,
	{_CState270, _CWildcardMarker}:              _CReduceBToExpressionAction,
	{_CState271, _CWildcardMarker}:              _CReduceBToInclusiveOrExpressionAction,
	{_CState272, _CWildcardMarker}:              _CReduceBToLogicalAndExpressionAction,
	{_CState273, _CWildcardMarker}:              _CReduceBToLogicalOrExpressionAction,
	{_CState275, _CWildcardMarker}:              _CReduceCToMultiplicativeExpressionAction,
	{_CState276, _CWildcardMarker}:              _CReduceDToMultiplicativeExpressionAction,
	{_CState277, _CWildcardMarker}:              _CReduceBToMultiplicativeExpressionAction,
	{_CState278, _CWildcardMarker}:              _CReduceEToPostfixExpressionAction,
	{_CState280, _CWildcardMarker}:              _CReduceCToPostfixExpressionAction,
	{_CState282, _CWildcardMarker}:              _CReduceAToArgumentExpressionListAction,
	{_CState283, _CWildcardMarker}:              _CReduceFToPostfixExpressionAction,
	{_CState284, _CWildcardMarker}:              _CReduceEToRelationalExpressionAction,
	{_CState285, _CWildcardMarker}:              _CReduceCToRelationalExpressionAction,
	{_CState286, _CWildcardMarker}:              _CReduceDToRelationalExpressionAction,
	{_CState287, _CWildcardMarker}:              _CReduceBToRelationalExpressionAction,
	{_CState288, _CWildcardMarker}:              _CReduceBToShiftExpressionAction,
	{_CState289, _CWildcardMarker}:              _CReduceCToShiftExpressionAction,
	{_CState290, _CWildcardMarker}:              _CReduceBToAssignmentExpressionAction,
	{_CState291, _CWildcardMarker}:              _CReduceBToDirectAbstractDeclaratorAction,
	{_CState293, _CWildcardMarker}:              _CReduceFToDirectAbstractDeclaratorAction,
	{_CState298, _CWildcardMarker}:              _CReduceCToAbstractDeclaratorAction,
	{_CState299, _CWildcardMarker}:              _CReduceBToIdentifierListAction,
	{_CState300, CRparamToken}:                  _CReduceBToParameterTypeListAction,
	{_CState301, _CWildcardMarker}:              _CReduceBToParameterListAction,
	{_CState302, _CWildcardMarker}:              _CReduceAToStructOrUnionSpecifierAction,
	{_CState303, _CWildcardMarker}:              _CReduceBToStructDeclaratorAction,
	{_CState306, _CWildcardMarker}:              _CReduceAToStructDeclarationAction,
	{_CState308, _CWildcardMarker}:              _CReduceBToInitializerAction,
	{_CState309, _CWildcardMarker}:              _CReduceBToLabeledStatementAction,
	{_CState313, _CWildcardMarker}:              _CReduceBToCastExpressionAction,
	{_CState314, _CWildcardMarker}:              _CReduceFToUnaryExpressionAction,
	{_CState318, _CWildcardMarker}:              _CReduceBToPostfixExpressionAction,
	{_CState320, _CWildcardMarker}:              _CReduceDToPostfixExpressionAction,
	{_CState321, _CWildcardMarker}:              _CReduceCToDirectAbstractDeclaratorAction,
	{_CState322, _CWildcardMarker}:              _CReduceAToDirectAbstractDeclaratorAction,
	{_CState323, _CWildcardMarker}:              _CReduceGToDirectAbstractDeclaratorAction,
	{_CState324, _CWildcardMarker}:              _CReduceDToDirectAbstractDeclaratorAction,
	{_CState326, _CWildcardMarker}:              _CReduceHToDirectAbstractDeclaratorAction,
	{_CState328, _CWildcardMarker}:              _CReduceCToStructDeclaratorAction,
	{_CState329, _CWildcardMarker}:              _CReduceBToStructDeclaratorListAction,
	{_CState330, _CWildcardMarker}:              _CReduceCToInitializerAction,
	{_CState331, _CWildcardMarker}:              _CReduceBToInitializerListAction,
	{_CState335, _CWildcardMarker}:              _CReduceAToSelectionStatementAction,
	{_CState336, CAndToken}:                     _CReduceAToSelectionStatementAction,
	{_CState336, CBreakToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CCaseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState336, CConstantToken}:                _CReduceAToSelectionStatementAction,
	{_CState336, CContinueToken}:                _CReduceAToSelectionStatementAction,
	{_CState336, CDecOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CDefaultToken}:                 _CReduceAToSelectionStatementAction,
	{_CState336, CDoToken}:                      _CReduceAToSelectionStatementAction,
	{_CState336, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState336, CExclaimToken}:                 _CReduceAToSelectionStatementAction,
	{_CState336, CForToken}:                     _CReduceAToSelectionStatementAction,
	{_CState336, CGotoToken}:                    _CReduceAToSelectionStatementAction,
	{_CState336, CIdentifierToken}:              _CReduceAToSelectionStatementAction,
	{_CState336, CIfToken}:                      _CReduceAToSelectionStatementAction,
	{_CState336, CIncOpToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CLcurlToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CLparamToken}:                  _CReduceAToSelectionStatementAction,
	{_CState336, CMinusToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CMulToken}:                     _CReduceAToSelectionStatementAction,
	{_CState336, CPlusToken}:                    _CReduceAToSelectionStatementAction,
	{_CState336, CRcurlToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CReturnToken}:                  _CReduceAToSelectionStatementAction,
	{_CState336, CSemicolonToken}:               _CReduceAToSelectionStatementAction,
	{_CState336, CSizeofToken}:                  _CReduceAToSelectionStatementAction,
	{_CState336, CStringLiteralToken}:           _CReduceAToSelectionStatementAction,
	{_CState336, CSwitchToken}:                  _CReduceAToSelectionStatementAction,
	{_CState336, CTildaToken}:                   _CReduceAToSelectionStatementAction,
	{_CState336, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState337, CElseToken}:                    _CReduceAToSelectionStatementAction,
	{_CState337, CWhileToken}:                   _CReduceAToSelectionStatementAction,
	{_CState338, _CWildcardMarker}:              _CReduceCToSelectionStatementAction,
	{_CState339, _CWildcardMarker}:              _CReduceAToIterationStatementAction,
	{_CState340, _CWildcardMarker}:              _CReduceBToConditionalExpressionAction,
	{_CState341, _CWildcardMarker}:              _CReduceBToArgumentExpressionListAction,
	{_CState342, _CWildcardMarker}:              _CReduceEToDirectAbstractDeclaratorAction,
	{_CState343, _CWildcardMarker}:              _CReduceIToDirectAbstractDeclaratorAction,
	{_CState345, _CWildcardMarker}:              _CReduceCToIterationStatementAction,
	{_CState348, _CWildcardMarker}:              _CReduceBToIterationStatementAction,
	{_CState349, _CWildcardMarker}:              _CReduceDToIterationStatementAction,
	{_CState350, _CWildcardMarker}:              _CReduceBToSelectionStatementAction,
}

var _CExpectedTerminals = map[_CStateId][]CSymbolId{
	_CState0:   []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken},
	_CState1:   []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, _CEndMarker},
	_CState6:   []CSymbolId{CIdentifierToken, CLcurlToken},
	_CState12:  []CSymbolId{CIdentifierToken, CLparamToken, CMulToken},
	_CState13:  []CSymbolId{CConstToken, CVolatileToken, CMulToken},
	_CState26:  []CSymbolId{CIdentifierToken, CLparamToken, CSemicolonToken, CMulToken},
	_CState27:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken},
	_CState28:  []CSymbolId{CLparamToken, CLbraceToken},
	_CState32:  []CSymbolId{CIdentifierToken, CLparamToken},
	_CState33:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState34:  []CSymbolId{CIdentifierToken, CLcurlToken},
	_CState36:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState37:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
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
	_CState62:  []CSymbolId{CIdentifierToken},
	_CState63:  []CSymbolId{CEqToken},
	_CState65:  []CSymbolId{CRcurlToken, CCommaToken},
	_CState69:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState71:  []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLcurlToken},
	_CState72:  []CSymbolId{CIdentifierToken, CLparamToken, CMulToken},
	_CState75:  []CSymbolId{CSemicolonToken},
	_CState76:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState78:  []CSymbolId{CSemicolonToken},
	_CState79:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState80:  []CSymbolId{CColonToken},
	_CState81:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState83:  []CSymbolId{CLparamToken},
	_CState84:  []CSymbolId{CIdentifierToken},
	_CState85:  []CSymbolId{CColonToken},
	_CState86:  []CSymbolId{CLparamToken},
	_CState87:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState88:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState93:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState95:  []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState97:  []CSymbolId{CLparamToken},
	_CState99:  []CSymbolId{CLparamToken},
	_CState100: []CSymbolId{CMinusToken, CPlusToken},
	_CState101: []CSymbolId{CAndToken},
	_CState106: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState107: []CSymbolId{CEqOpToken, CNeOpToken},
	_CState108: []CSymbolId{CHatToken},
	_CState109: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState111: []CSymbolId{COrToken},
	_CState115: []CSymbolId{CAndOpToken},
	_CState116: []CSymbolId{COrOpToken, CQuestionToken},
	_CState117: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState118: []CSymbolId{CPtrOpToken, CIncOpToken, CDecOpToken, CLparamToken, CLbraceToken, CDotToken},
	_CState120: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState122: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState124: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState125: []CSymbolId{CMulAssignToken, CDivAssignToken, CModAssignToken, CAddAssignToken, CSubAssignToken, CLeftAssignToken, CRightAssignToken, CAndAssignToken, CXorAssignToken, COrAssignToken, CEqToken},
	_CState126: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState129: []CSymbolId{CEqToken},
	_CState133: []CSymbolId{CRbraceToken},
	_CState137: []CSymbolId{CIdentifierToken, CLparamToken, CLbraceToken, CMulToken},
	_CState138: []CSymbolId{CRparamToken, CCommaToken},
	_CState140: []CSymbolId{CCommaToken, CRparamToken},
	_CState141: []CSymbolId{CRparamToken},
	_CState142: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState143: []CSymbolId{CIdentifierToken, CLparamToken, CColonToken, CMulToken},
	_CState145: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRcurlToken},
	_CState146: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState147: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken},
	_CState148: []CSymbolId{CRcurlToken, CCommaToken},
	_CState149: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState150: []CSymbolId{CIdentifierToken},
	_CState152: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState158: []CSymbolId{CColonToken},
	_CState160: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState162: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState163: []CSymbolId{CWhileToken},
	_CState164: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState165: []CSymbolId{CSemicolonToken},
	_CState166: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState167: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState169: []CSymbolId{CRparamToken, CCommaToken},
	_CState170: []CSymbolId{CLparamToken, CLbraceToken, CMulToken, CRparamToken},
	_CState171: []CSymbolId{CRparamToken},
	_CState173: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState174: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState176: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState177: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState178: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState179: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState180: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState182: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CRcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState183: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState184: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState185: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState186: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState188: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState189: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState190: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState191: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState192: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState193: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState194: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState196: []CSymbolId{CIdentifierToken},
	_CState198: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState199: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState200: []CSymbolId{CIdentifierToken},
	_CState201: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState202: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState203: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState204: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState205: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState206: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState220: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState223: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRbraceToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState224: []CSymbolId{CIdentifierToken, CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CRparamToken, CLbraceToken, CMulToken},
	_CState227: []CSymbolId{CLparamToken, CLbraceToken},
	_CState228: []CSymbolId{CIdentifierToken, CLparamToken, CLbraceToken},
	_CState229: []CSymbolId{CIdentifierToken},
	_CState231: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CEllipsisToken},
	_CState233: []CSymbolId{CTypeNameToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRcurlToken},
	_CState234: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState235: []CSymbolId{CColonToken},
	_CState237: []CSymbolId{CSemicolonToken, CCommaToken},
	_CState246: []CSymbolId{CRcurlToken, CCommaToken},
	_CState247: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState249: []CSymbolId{CLparamToken},
	_CState250: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState253: []CSymbolId{CRparamToken, CCommaToken},
	_CState255: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CLparamToken, CRparamToken, CLbraceToken, CMulToken},
	_CState256: []CSymbolId{CRparamToken},
	_CState257: []CSymbolId{CLparamToken, CLbraceToken, CRparamToken},
	_CState258: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState260: []CSymbolId{CRparamToken},
	_CState261: []CSymbolId{CRparamToken, CCommaToken},
	_CState262: []CSymbolId{CRparamToken, CCommaToken},
	_CState263: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState264: []CSymbolId{CMulToken, CDivToken, CModToken},
	_CState265: []CSymbolId{CEqOpToken, CNeOpToken},
	_CState267: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState268: []CSymbolId{CLeOpToken, CGeOpToken, CLtToken, CGtToken},
	_CState269: []CSymbolId{CAndToken},
	_CState271: []CSymbolId{CHatToken},
	_CState272: []CSymbolId{COrToken},
	_CState273: []CSymbolId{CAndOpToken},
	_CState274: []CSymbolId{CColonToken, CCommaToken},
	_CState279: []CSymbolId{CRbraceToken, CCommaToken},
	_CState281: []CSymbolId{CRparamToken, CCommaToken},
	_CState284: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState285: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState286: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState287: []CSymbolId{CLeftOpToken, CRightOpToken},
	_CState288: []CSymbolId{CMinusToken, CPlusToken},
	_CState289: []CSymbolId{CMinusToken, CPlusToken},
	_CState292: []CSymbolId{CRbraceToken},
	_CState294: []CSymbolId{CRparamToken},
	_CState295: []CSymbolId{CRparamToken},
	_CState296: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRbraceToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState297: []CSymbolId{CTypeNameToken, CTypedefToken, CExternToken, CStaticToken, CAutoToken, CRegisterToken, CCharToken, CShortToken, CIntToken, CLongToken, CSignedToken, CUnsignedToken, CFloatToken, CDoubleToken, CConstToken, CVolatileToken, CVoidToken, CStructToken, CUnionToken, CEnumToken, CRparamToken},
	_CState298: []CSymbolId{CLparamToken, CLbraceToken},
	_CState300: []CSymbolId{CRparamToken},
	_CState304: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState305: []CSymbolId{CIdentifierToken, CLparamToken, CColonToken, CMulToken},
	_CState307: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CLcurlToken, CRcurlToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState310: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState311: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CRparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState312: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState315: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState316: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState317: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState319: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CLparamToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState325: []CSymbolId{CRbraceToken},
	_CState327: []CSymbolId{CRparamToken},
	_CState332: []CSymbolId{CRparamToken, CCommaToken},
	_CState333: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState334: []CSymbolId{CRparamToken, CCommaToken},
	_CState335: []CSymbolId{CElseToken},
	_CState336: []CSymbolId{CElseToken, CAndToken, CBreakToken, CCaseToken, CConstantToken, CContinueToken, CDecOpToken, CDefaultToken, CDoToken, CElseToken, CExclaimToken, CForToken, CGotoToken, CIdentifierToken, CIfToken, CIncOpToken, CLcurlToken, CLparamToken, CMinusToken, CMulToken, CPlusToken, CRcurlToken, CReturnToken, CSemicolonToken, CSizeofToken, CStringLiteralToken, CSwitchToken, CTildaToken, CWhileToken},
	_CState337: []CSymbolId{CElseToken, CElseToken, CWhileToken},
	_CState344: []CSymbolId{CSemicolonToken},
	_CState346: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
	_CState347: []CSymbolId{CIdentifierToken, CConstantToken, CStringLiteralToken, CSizeofToken, CIncOpToken, CDecOpToken, CCaseToken, CDefaultToken, CIfToken, CSwitchToken, CWhileToken, CDoToken, CForToken, CGotoToken, CContinueToken, CBreakToken, CReturnToken, CLparamToken, CLcurlToken, CSemicolonToken, CMulToken, CMinusToken, CPlusToken, CAndToken, CExclaimToken, CTildaToken},
}

/*
Parser Debug States:
  State 0:
    Kernel Items:
      #accept: ^.translation_unit
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 12
      MUL -> State 13
      declaration -> State 25
      declaration_specifiers -> State 26
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      declarator -> State 27
      direct_declarator -> State 28
      pointer -> State 32
      translation_unit -> State 1
      external_declaration -> State 30
      function_definition -> State 31

  State 1:
    Kernel Items:
      #accept: ^ translation_unit., $
      translation_unit: translation_unit.external_declaration
    Reduce:
      $ -> [#accept]
    Goto:
      IDENTIFIER -> State 9
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 12
      MUL -> State 13
      declaration -> State 25
      declaration_specifiers -> State 26
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      declarator -> State 27
      direct_declarator -> State 28
      pointer -> State 32
      external_declaration -> State 38
      function_definition -> State 31

  State 2:
    Kernel Items:
      storage_class_specifier: AUTO., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 3:
    Kernel Items:
      type_specifier: CHAR., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 4:
    Kernel Items:
      type_qualifier: CONST., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      type_specifier: DOUBLE., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      enum_specifier: ENUM.IDENTIFIER
      enum_specifier: ENUM.IDENTIFIER LCURL enumerator_list RCURL
      enum_specifier: ENUM.LCURL enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 39
      LCURL -> State 40

  State 7:
    Kernel Items:
      storage_class_specifier: EXTERN., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      type_specifier: FLOAT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      direct_declarator: IDENTIFIER., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      type_specifier: INT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      type_specifier: LONG., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      direct_declarator: LPARAM.declarator RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      MUL -> State 13
      declarator -> State 41
      direct_declarator -> State 28
      pointer -> State 32

  State 13:
    Kernel Items:
      pointer: MUL., *
      pointer: MUL.pointer
      pointer: MUL.type_qualifier_list
      pointer: MUL.type_qualifier_list pointer
    Reduce:
      * -> [pointer]
    Goto:
      CONST -> State 4
      VOLATILE -> State 24
      MUL -> State 13
      type_qualifier -> State 43
      pointer -> State 42
      type_qualifier_list -> State 44

  State 14:
    Kernel Items:
      storage_class_specifier: REGISTER., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      type_specifier: SHORT., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      type_specifier: SIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      storage_class_specifier: STATIC., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      struct_or_union: STRUCT., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      storage_class_specifier: TYPEDEF., *
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      type_specifier: TYPE_NAME., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      struct_or_union: UNION., *
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      type_specifier: UNSIGNED., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      type_specifier: VOID., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      type_qualifier: VOLATILE., *
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      external_declaration: declaration., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      declaration: declaration_specifiers.SEMICOLON
      declaration: declaration_specifiers.init_declarator_list SEMICOLON
      function_definition: declaration_specifiers.declarator compound_statement
      function_definition: declaration_specifiers.declarator declaration_list compound_statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      SEMICOLON -> State 45
      MUL -> State 13
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 46
      direct_declarator -> State 28
      pointer -> State 32

  State 27:
    Kernel Items:
      function_definition: declarator.compound_statement
      function_definition: declarator.declaration_list compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LCURL -> State 49
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      compound_statement -> State 50
      declaration_list -> State 52

  State 28:
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

  State 29:
    Kernel Items:
      type_specifier: enum_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      translation_unit: external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      external_declaration: function_definition., *
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      declarator: pointer.direct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      direct_declarator -> State 56

  State 33:
    Kernel Items:
      declaration_specifiers: storage_class_specifier., *
      declaration_specifiers: storage_class_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      declaration_specifiers -> State 57
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36

  State 34:
    Kernel Items:
      struct_or_union_specifier: struct_or_union.IDENTIFIER
      struct_or_union_specifier: struct_or_union.IDENTIFIER LCURL struct_declaration_list RCURL
      struct_or_union_specifier: struct_or_union.LCURL struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 58
      LCURL -> State 59

  State 35:
    Kernel Items:
      type_specifier: struct_or_union_specifier., *
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      declaration_specifiers: type_qualifier., *
      declaration_specifiers: type_qualifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      declaration_specifiers -> State 60
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36

  State 37:
    Kernel Items:
      declaration_specifiers: type_specifier., *
      declaration_specifiers: type_specifier.declaration_specifiers
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      declaration_specifiers -> State 61
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36

  State 38:
    Kernel Items:
      translation_unit: translation_unit external_declaration., *
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER., *
      enum_specifier: ENUM IDENTIFIER.LCURL enumerator_list RCURL
    Reduce:
      * -> [enum_specifier]
    Goto:
      LCURL -> State 62

  State 40:
    Kernel Items:
      enum_specifier: ENUM LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator_list -> State 65
      enumerator -> State 64

  State 41:
    Kernel Items:
      direct_declarator: LPARAM declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 66

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
      CONST -> State 4
      VOLATILE -> State 24
      MUL -> State 13
      type_qualifier -> State 68
      pointer -> State 67

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
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LCURL -> State 49
      EQ -> State 69
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      compound_statement -> State 70
      declaration_list -> State 71

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
      SEMICOLON -> State 73
      COMMA -> State 72

  State 49:
    Kernel Items:
      compound_statement: LCURL.RCURL
      compound_statement: LCURL.declaration_list RCURL
      compound_statement: LCURL.declaration_list statement_list RCURL
      compound_statement: LCURL.statement_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      RCURL -> State 92
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      declaration -> State 51
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      statement -> State 123
      labeled_statement -> State 114
      compound_statement -> State 104
      declaration_list -> State 106
      statement_list -> State 124
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

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
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LCURL -> State 49
      declaration -> State 128
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      compound_statement -> State 127

  State 53:
    Kernel Items:
      declaration: declaration_specifiers.SEMICOLON
      declaration: declaration_specifiers.init_declarator_list SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      SEMICOLON -> State 45
      MUL -> State 13
      init_declarator_list -> State 48
      init_declarator -> State 47
      declarator -> State 129
      direct_declarator -> State 28
      pointer -> State 32

  State 54:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE.RBRACE
      direct_declarator: direct_declarator LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      RBRACE -> State 131
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 133

  State 55:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM.RPARAM
      direct_declarator: direct_declarator LPARAM.identifier_list RPARAM
      direct_declarator: direct_declarator LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      RPARAM -> State 136
      declaration_specifiers -> State 137
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      parameter_type_list -> State 141
      parameter_list -> State 140
      parameter_declaration -> State 139
      identifier_list -> State 138

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
      LCURL -> State 142

  State 59:
    Kernel Items:
      struct_or_union_specifier: struct_or_union LCURL.struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      struct_declaration_list -> State 145
      struct_declaration -> State 144
      specifier_qualifier_list -> State 143
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 60:
    Kernel Items:
      declaration_specifiers: type_qualifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      declaration_specifiers: type_specifier declaration_specifiers., *
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator_list -> State 148
      enumerator -> State 64

  State 63:
    Kernel Items:
      enumerator: IDENTIFIER., *
      enumerator: IDENTIFIER.EQ constant_expression
    Reduce:
      * -> [enumerator]
    Goto:
      EQ -> State 149

  State 64:
    Kernel Items:
      enumerator_list: enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      enum_specifier: ENUM LCURL enumerator_list.RCURL
      enumerator_list: enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 151
      COMMA -> State 150

  State 66:
    Kernel Items:
      direct_declarator: LPARAM declarator RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      pointer: MUL type_qualifier_list pointer., *
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      type_qualifier_list: type_qualifier_list type_qualifier., *
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      init_declarator: declarator EQ.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      LCURL -> State 152
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 153
      initializer -> State 154

  State 70:
    Kernel Items:
      function_definition: declaration_specifiers declarator compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      declaration_list: declaration_list.declaration
      function_definition: declaration_specifiers declarator declaration_list.compound_statement
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LCURL -> State 49
      declaration -> State 128
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      compound_statement -> State 155

  State 72:
    Kernel Items:
      init_declarator_list: init_declarator_list COMMA.init_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      MUL -> State 13
      init_declarator -> State 156
      declarator -> State 129
      direct_declarator -> State 28
      pointer -> State 32

  State 73:
    Kernel Items:
      declaration: declaration_specifiers init_declarator_list SEMICOLON., *
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      unary_operator: AND., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      jump_statement: BREAK.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 157

  State 76:
    Kernel Items:
      labeled_statement: CASE.constant_expression COLON statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 158

  State 77:
    Kernel Items:
      primary_expression: CONSTANT., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      jump_statement: CONTINUE.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 159

  State 79:
    Kernel Items:
      unary_expression: DEC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 160
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 161
      unary_operator -> State 126

  State 80:
    Kernel Items:
      labeled_statement: DEFAULT.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 162

  State 81:
    Kernel Items:
      iteration_statement: DO.statement WHILE LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 163
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 82:
    Kernel Items:
      unary_operator: EXCLAIM., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      iteration_statement: FOR.LPARAM expression_statement expression_statement RPARAM statement
      iteration_statement: FOR.LPARAM expression_statement expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 164

  State 84:
    Kernel Items:
      jump_statement: GOTO.IDENTIFIER SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 165

  State 85:
    Kernel Items:
      labeled_statement: IDENTIFIER.COLON statement
      primary_expression: IDENTIFIER., *
    Reduce:
      * -> [primary_expression]
    Goto:
      COLON -> State 166

  State 86:
    Kernel Items:
      selection_statement: IF.LPARAM expression RPARAM statement
      selection_statement: IF.LPARAM expression RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 167

  State 87:
    Kernel Items:
      unary_expression: INC_OP.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 160
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 168
      unary_operator -> State 126

  State 88:
    Kernel Items:
      cast_expression: LPARAM.type_name RPARAM cast_expression
      primary_expression: LPARAM.expression RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 169
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      specifier_qualifier_list -> State 170
      enum_specifier -> State 29
      type_qualifier -> State 146
      type_name -> State 171

  State 89:
    Kernel Items:
      unary_operator: MINUS., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      unary_operator: MUL., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      unary_operator: PLUS., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      compound_statement: LCURL RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      jump_statement: RETURN.SEMICOLON
      jump_statement: RETURN.expression SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      SEMICOLON -> State 172
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 173

  State 94:
    Kernel Items:
      expression_statement: SEMICOLON., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      unary_expression: SIZEOF.LPARAM type_name RPARAM
      unary_expression: SIZEOF.unary_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 174
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 175
      unary_operator -> State 126

  State 96:
    Kernel Items:
      primary_expression: STRING_LITERAL., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      selection_statement: SWITCH.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 176

  State 98:
    Kernel Items:
      unary_operator: TILDA., *
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      iteration_statement: WHILE.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 177

  State 100:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 101:
    Kernel Items:
      and_expression: and_expression.AND equality_expression
      exclusive_or_expression: and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 180

  State 102:
    Kernel Items:
      expression: assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      multiplicative_expression: cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      statement: compound_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      assignment_expression: conditional_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      compound_statement: LCURL declaration_list.RCURL
      compound_statement: LCURL declaration_list.statement_list RCURL
      declaration_list: declaration_list.declaration
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      RCURL -> State 181
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      declaration -> State 128
      declaration_specifiers -> State 53
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      statement -> State 123
      labeled_statement -> State 114
      compound_statement -> State 104
      statement_list -> State 182
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 107:
    Kernel Items:
      and_expression: equality_expression., *
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 183
      NE_OP -> State 184

  State 108:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.HAT and_expression
      inclusive_or_expression: exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 185

  State 109:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      expression_statement: expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 187
      COMMA -> State 186

  State 110:
    Kernel Items:
      statement: expression_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression: inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 188

  State 112:
    Kernel Items:
      statement: iteration_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      statement: jump_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      statement: labeled_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 189

  State 116:
    Kernel Items:
      conditional_expression: logical_or_expression., *
      conditional_expression: logical_or_expression.QUESTION expression COLON conditional_expression
      logical_or_expression: logical_or_expression.OR_OP logical_and_expression
    Reduce:
      * -> [conditional_expression]
    Goto:
      OR_OP -> State 190
      QUESTION -> State 191

  State 117:
    Kernel Items:
      additive_expression: multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 118:
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
      PTR_OP -> State 200
      INC_OP -> State 197
      DEC_OP -> State 195
      LPARAM -> State 199
      LBRACE -> State 198
      DOT -> State 196

  State 119:
    Kernel Items:
      postfix_expression: primary_expression., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      equality_expression: relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 121:
    Kernel Items:
      statement: selection_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      relational_expression: shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 123:
    Kernel Items:
      statement_list: statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      compound_statement: LCURL statement_list.RCURL
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      RCURL -> State 207
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 208
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 125:
    Kernel Items:
      assignment_expression: unary_expression.assignment_operator assignment_expression
      cast_expression: unary_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      MUL_ASSIGN -> State 215
      DIV_ASSIGN -> State 211
      MOD_ASSIGN -> State 214
      ADD_ASSIGN -> State 209
      SUB_ASSIGN -> State 218
      LEFT_ASSIGN -> State 213
      RIGHT_ASSIGN -> State 217
      AND_ASSIGN -> State 210
      XOR_ASSIGN -> State 219
      OR_ASSIGN -> State 216
      EQ -> State 212
      assignment_operator -> State 220

  State 126:
    Kernel Items:
      unary_expression: unary_operator.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 221

  State 127:
    Kernel Items:
      function_definition: declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      declaration_list: declaration_list declaration., *
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      init_declarator: declarator., *
      init_declarator: declarator.EQ initializer
    Reduce:
      * -> [init_declarator]
    Goto:
      EQ -> State 69

  State 130:
    Kernel Items:
      primary_expression: IDENTIFIER., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE RBRACE., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      constant_expression: conditional_expression., *
    Reduce:
      * -> [constant_expression]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 222

  State 134:
    Kernel Items:
      cast_expression: unary_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      identifier_list: IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      parameter_declaration: declaration_specifiers., *
      parameter_declaration: declaration_specifiers.abstract_declarator
      parameter_declaration: declaration_specifiers.declarator
    Reduce:
      * -> [parameter_declaration]
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 224
      LBRACE -> State 223
      MUL -> State 13
      declarator -> State 226
      direct_declarator -> State 28
      pointer -> State 228
      abstract_declarator -> State 225
      direct_abstract_declarator -> State 227

  State 138:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM identifier_list.RPARAM
      identifier_list: identifier_list.COMMA IDENTIFIER
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 230
      COMMA -> State 229

  State 139:
    Kernel Items:
      parameter_list: parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      parameter_list: parameter_list.COMMA parameter_declaration
      parameter_type_list: parameter_list., RPARAM
      parameter_type_list: parameter_list.COMMA ELLIPSIS
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      COMMA -> State 231

  State 141:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 232

  State 142:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL.struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      struct_declaration_list -> State 233
      struct_declaration -> State 144
      specifier_qualifier_list -> State 143
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 143:
    Kernel Items:
      struct_declaration: specifier_qualifier_list.struct_declarator_list SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      COLON -> State 234
      MUL -> State 13
      struct_declarator_list -> State 237
      struct_declarator -> State 236
      declarator -> State 235
      direct_declarator -> State 28
      pointer -> State 32

  State 144:
    Kernel Items:
      struct_declaration_list: struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      struct_declaration_list: struct_declaration_list.struct_declaration
      struct_or_union_specifier: struct_or_union LCURL struct_declaration_list.RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      RCURL -> State 238
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      struct_declaration -> State 239
      specifier_qualifier_list -> State 143
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 146:
    Kernel Items:
      specifier_qualifier_list: type_qualifier., *
      specifier_qualifier_list: type_qualifier.specifier_qualifier_list
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      specifier_qualifier_list -> State 240
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 147:
    Kernel Items:
      specifier_qualifier_list: type_specifier., *
      specifier_qualifier_list: type_specifier.specifier_qualifier_list
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      specifier_qualifier_list -> State 241
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 148:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL enumerator_list.RCURL
      enumerator_list: enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 242
      COMMA -> State 150

  State 149:
    Kernel Items:
      enumerator: IDENTIFIER EQ.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 243

  State 150:
    Kernel Items:
      enumerator_list: enumerator_list COMMA.enumerator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator -> State 244

  State 151:
    Kernel Items:
      enum_specifier: ENUM LCURL enumerator_list RCURL., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      initializer: LCURL.initializer_list COMMA RCURL
      initializer: LCURL.initializer_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      LCURL -> State 152
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 153
      initializer -> State 245
      initializer_list -> State 246

  State 153:
    Kernel Items:
      initializer: assignment_expression., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      init_declarator: declarator EQ initializer., *
    Reduce:
      * -> [init_declarator]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      function_definition: declaration_specifiers declarator declaration_list compound_statement., *
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      init_declarator_list: init_declarator_list COMMA init_declarator., *
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      jump_statement: BREAK SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      labeled_statement: CASE constant_expression.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 247

  State 159:
    Kernel Items:
      jump_statement: CONTINUE SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      primary_expression: LPARAM.expression RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 169

  State 161:
    Kernel Items:
      unary_expression: DEC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      labeled_statement: DEFAULT COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 248
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 163:
    Kernel Items:
      iteration_statement: DO statement.WHILE LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      WHILE -> State 249

  State 164:
    Kernel Items:
      iteration_statement: FOR LPARAM.expression_statement expression_statement RPARAM statement
      iteration_statement: FOR LPARAM.expression_statement expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      expression_statement -> State 250

  State 165:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 251

  State 166:
    Kernel Items:
      labeled_statement: IDENTIFIER COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 252
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 167:
    Kernel Items:
      selection_statement: IF LPARAM.expression RPARAM statement
      selection_statement: IF LPARAM.expression RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 253

  State 168:
    Kernel Items:
      unary_expression: INC_OP unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      primary_expression: LPARAM expression.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 254
      COMMA -> State 186

  State 170:
    Kernel Items:
      type_name: specifier_qualifier_list., RPARAM
      type_name: specifier_qualifier_list.abstract_declarator
    Reduce:
      RPARAM -> [type_name]
    Goto:
      LPARAM -> State 255
      LBRACE -> State 223
      MUL -> State 13
      pointer -> State 257
      abstract_declarator -> State 256
      direct_abstract_declarator -> State 227

  State 171:
    Kernel Items:
      cast_expression: LPARAM type_name.RPARAM cast_expression
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 258

  State 172:
    Kernel Items:
      jump_statement: RETURN SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      jump_statement: RETURN expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 259
      COMMA -> State 186

  State 174:
    Kernel Items:
      primary_expression: LPARAM.expression RPARAM
      unary_expression: SIZEOF LPARAM.type_name RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 169
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      specifier_qualifier_list -> State 170
      enum_specifier -> State 29
      type_qualifier -> State 146
      type_name -> State 260

  State 175:
    Kernel Items:
      unary_expression: SIZEOF unary_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      selection_statement: SWITCH LPARAM.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 261

  State 177:
    Kernel Items:
      iteration_statement: WHILE LPARAM.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 262

  State 178:
    Kernel Items:
      additive_expression: additive_expression MINUS.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 263

  State 179:
    Kernel Items:
      additive_expression: additive_expression PLUS.multiplicative_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 264

  State 180:
    Kernel Items:
      and_expression: and_expression AND.equality_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 265

  State 181:
    Kernel Items:
      compound_statement: LCURL declaration_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      compound_statement: LCURL declaration_list statement_list.RCURL
      statement_list: statement_list.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      RCURL -> State 266
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 208
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 183:
    Kernel Items:
      equality_expression: equality_expression EQ_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 267

  State 184:
    Kernel Items:
      equality_expression: equality_expression NE_OP.relational_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 268

  State 185:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression HAT.and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 269

  State 186:
    Kernel Items:
      expression: expression COMMA.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 270

  State 187:
    Kernel Items:
      expression_statement: expression SEMICOLON., *
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression OR.exclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 271

  State 189:
    Kernel Items:
      logical_and_expression: logical_and_expression AND_OP.inclusive_or_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 272

  State 190:
    Kernel Items:
      logical_or_expression: logical_or_expression OR_OP.logical_and_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 273

  State 191:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION.expression COLON conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 274

  State 192:
    Kernel Items:
      multiplicative_expression: multiplicative_expression DIV.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 275

  State 193:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MOD.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 276

  State 194:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MUL.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 277

  State 195:
    Kernel Items:
      postfix_expression: postfix_expression DEC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      postfix_expression: postfix_expression DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 278

  State 197:
    Kernel Items:
      postfix_expression: postfix_expression INC_OP., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      postfix_expression: postfix_expression LBRACE.expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 279

  State 199:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM.RPARAM
      postfix_expression: postfix_expression LPARAM.argument_expression_list RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      RPARAM -> State 280
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      argument_expression_list -> State 281
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 282

  State 200:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 283

  State 201:
    Kernel Items:
      relational_expression: relational_expression GE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 284

  State 202:
    Kernel Items:
      relational_expression: relational_expression GT.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 285

  State 203:
    Kernel Items:
      relational_expression: relational_expression LE_OP.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 286

  State 204:
    Kernel Items:
      relational_expression: relational_expression LT.shift_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 287

  State 205:
    Kernel Items:
      shift_expression: shift_expression LEFT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 288

  State 206:
    Kernel Items:
      shift_expression: shift_expression RIGHT_OP.additive_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 289

  State 207:
    Kernel Items:
      compound_statement: LCURL statement_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      statement_list: statement_list statement., *
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      assignment_operator: ADD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      assignment_operator: AND_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      assignment_operator: DIV_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      assignment_operator: EQ., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      assignment_operator: LEFT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      assignment_operator: MOD_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      assignment_operator: MUL_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      assignment_operator: OR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      assignment_operator: RIGHT_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      assignment_operator: SUB_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      assignment_operator: XOR_ASSIGN., *
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 290

  State 221:
    Kernel Items:
      unary_expression: unary_operator cast_expression., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      direct_declarator: direct_declarator LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      direct_abstract_declarator: LBRACE.RBRACE
      direct_abstract_declarator: LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      RBRACE -> State 291
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 292

  State 224:
    Kernel Items:
      direct_abstract_declarator: LPARAM.RPARAM
      direct_abstract_declarator: LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator: LPARAM.parameter_type_list RPARAM
      direct_declarator: LPARAM.declarator RPARAM
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 224
      RPARAM -> State 293
      LBRACE -> State 223
      MUL -> State 13
      declaration_specifiers -> State 137
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      declarator -> State 41
      direct_declarator -> State 28
      pointer -> State 228
      parameter_type_list -> State 295
      parameter_list -> State 140
      parameter_declaration -> State 139
      abstract_declarator -> State 294
      direct_abstract_declarator -> State 227

  State 225:
    Kernel Items:
      parameter_declaration: declaration_specifiers abstract_declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      parameter_declaration: declaration_specifiers declarator., *
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      abstract_declarator: direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator: direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 297
      LBRACE -> State 296

  State 228:
    Kernel Items:
      abstract_declarator: pointer., *
      abstract_declarator: pointer.direct_abstract_declarator
      declarator: pointer.direct_declarator
    Reduce:
      * -> [abstract_declarator]
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 224
      LBRACE -> State 223
      direct_declarator -> State 56
      direct_abstract_declarator -> State 298

  State 229:
    Kernel Items:
      identifier_list: identifier_list COMMA.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 299

  State 230:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM identifier_list RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      parameter_list: parameter_list COMMA.parameter_declaration
      parameter_type_list: parameter_list COMMA.ELLIPSIS
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      ELLIPSIS -> State 300
      declaration_specifiers -> State 137
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      parameter_declaration -> State 301

  State 232:
    Kernel Items:
      direct_declarator: direct_declarator LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      struct_declaration_list: struct_declaration_list.struct_declaration
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL struct_declaration_list.RCURL
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      RCURL -> State 302
      type_specifier -> State 147
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      struct_declaration -> State 239
      specifier_qualifier_list -> State 143
      enum_specifier -> State 29
      type_qualifier -> State 146

  State 234:
    Kernel Items:
      struct_declarator: COLON.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 303

  State 235:
    Kernel Items:
      struct_declarator: declarator., *
      struct_declarator: declarator.COLON constant_expression
    Reduce:
      * -> [struct_declarator]
    Goto:
      COLON -> State 304

  State 236:
    Kernel Items:
      struct_declarator_list: struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list.SEMICOLON
      struct_declarator_list: struct_declarator_list.COMMA struct_declarator
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 306
      COMMA -> State 305

  State 238:
    Kernel Items:
      struct_or_union_specifier: struct_or_union LCURL struct_declaration_list RCURL., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      struct_declaration_list: struct_declaration_list struct_declaration., *
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      specifier_qualifier_list: type_qualifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      specifier_qualifier_list: type_specifier specifier_qualifier_list., *
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      enum_specifier: ENUM IDENTIFIER LCURL enumerator_list RCURL., *
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      enumerator: IDENTIFIER EQ constant_expression., *
    Reduce:
      * -> [enumerator]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      enumerator_list: enumerator_list COMMA enumerator., *
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      initializer_list: initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      initializer: LCURL initializer_list.COMMA RCURL
      initializer: LCURL initializer_list.RCURL
      initializer_list: initializer_list.COMMA initializer
    Reduce:
      (nil)
    Goto:
      RCURL -> State 308
      COMMA -> State 307

  State 247:
    Kernel Items:
      labeled_statement: CASE constant_expression COLON.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 309
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 248:
    Kernel Items:
      labeled_statement: DEFAULT COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      iteration_statement: DO statement WHILE.LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 310

  State 250:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement.expression_statement RPARAM statement
      iteration_statement: FOR LPARAM expression_statement.expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      expression_statement -> State 311

  State 251:
    Kernel Items:
      jump_statement: GOTO IDENTIFIER SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      labeled_statement: IDENTIFIER COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      selection_statement: IF LPARAM expression.RPARAM statement
      selection_statement: IF LPARAM expression.RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 312
      COMMA -> State 186

  State 254:
    Kernel Items:
      primary_expression: LPARAM expression RPARAM., *
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      direct_abstract_declarator: LPARAM.RPARAM
      direct_abstract_declarator: LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator: LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      LPARAM -> State 255
      RPARAM -> State 293
      LBRACE -> State 223
      MUL -> State 13
      declaration_specifiers -> State 137
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      pointer -> State 257
      parameter_type_list -> State 295
      parameter_list -> State 140
      parameter_declaration -> State 139
      abstract_declarator -> State 294
      direct_abstract_declarator -> State 227

  State 256:
    Kernel Items:
      type_name: specifier_qualifier_list abstract_declarator., RPARAM
    Reduce:
      RPARAM -> [type_name]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      abstract_declarator: pointer., RPARAM
      abstract_declarator: pointer.direct_abstract_declarator
    Reduce:
      RPARAM -> [abstract_declarator]
    Goto:
      LPARAM -> State 255
      LBRACE -> State 223
      direct_abstract_declarator -> State 298

  State 258:
    Kernel Items:
      cast_expression: LPARAM type_name RPARAM.cast_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 313

  State 259:
    Kernel Items:
      jump_statement: RETURN expression SEMICOLON., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      unary_expression: SIZEOF LPARAM type_name.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 314

  State 261:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      selection_statement: SWITCH LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 315
      COMMA -> State 186

  State 262:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: WHILE LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 316
      COMMA -> State 186

  State 263:
    Kernel Items:
      additive_expression: additive_expression MINUS multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 264:
    Kernel Items:
      additive_expression: additive_expression PLUS multiplicative_expression., *
      multiplicative_expression: multiplicative_expression.DIV cast_expression
      multiplicative_expression: multiplicative_expression.MOD cast_expression
      multiplicative_expression: multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 265:
    Kernel Items:
      and_expression: and_expression AND equality_expression., *
      equality_expression: equality_expression.EQ_OP relational_expression
      equality_expression: equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 183
      NE_OP -> State 184

  State 266:
    Kernel Items:
      compound_statement: LCURL declaration_list statement_list RCURL., *
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      equality_expression: equality_expression EQ_OP relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 268:
    Kernel Items:
      equality_expression: equality_expression NE_OP relational_expression., *
      relational_expression: relational_expression.GE_OP shift_expression
      relational_expression: relational_expression.GT shift_expression
      relational_expression: relational_expression.LE_OP shift_expression
      relational_expression: relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 269:
    Kernel Items:
      and_expression: and_expression.AND equality_expression
      exclusive_or_expression: exclusive_or_expression HAT and_expression., *
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 180

  State 270:
    Kernel Items:
      expression: expression COMMA assignment_expression., *
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      exclusive_or_expression: exclusive_or_expression.HAT and_expression
      inclusive_or_expression: inclusive_or_expression OR exclusive_or_expression., *
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 185

  State 272:
    Kernel Items:
      inclusive_or_expression: inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression: logical_and_expression AND_OP inclusive_or_expression., *
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 188

  State 273:
    Kernel Items:
      logical_and_expression: logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression: logical_or_expression OR_OP logical_and_expression., *
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 189

  State 274:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression.COLON conditional_expression
      expression: expression.COMMA assignment_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 317
      COMMA -> State 186

  State 275:
    Kernel Items:
      multiplicative_expression: multiplicative_expression DIV cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MOD cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      multiplicative_expression: multiplicative_expression MUL cast_expression., *
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      postfix_expression: postfix_expression DOT IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      postfix_expression: postfix_expression LBRACE expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 318
      COMMA -> State 186

  State 280:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM RPARAM., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      argument_expression_list: argument_expression_list.COMMA assignment_expression
      postfix_expression: postfix_expression LPARAM argument_expression_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 320
      COMMA -> State 319

  State 282:
    Kernel Items:
      argument_expression_list: assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      postfix_expression: postfix_expression PTR_OP IDENTIFIER., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      relational_expression: relational_expression GE_OP shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 285:
    Kernel Items:
      relational_expression: relational_expression GT shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 286:
    Kernel Items:
      relational_expression: relational_expression LE_OP shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 287:
    Kernel Items:
      relational_expression: relational_expression LT shift_expression., *
      shift_expression: shift_expression.LEFT_OP additive_expression
      shift_expression: shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 288:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: shift_expression LEFT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 289:
    Kernel Items:
      additive_expression: additive_expression.MINUS multiplicative_expression
      additive_expression: additive_expression.PLUS multiplicative_expression
      shift_expression: shift_expression RIGHT_OP additive_expression., *
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 290:
    Kernel Items:
      assignment_expression: unary_expression assignment_operator assignment_expression., *
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      direct_abstract_declarator: LBRACE RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      direct_abstract_declarator: LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 321

  State 293:
    Kernel Items:
      direct_abstract_declarator: LPARAM RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      direct_abstract_declarator: LPARAM abstract_declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 322

  State 295:
    Kernel Items:
      direct_abstract_declarator: LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 323

  State 296:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE.RBRACE
      direct_abstract_declarator: direct_abstract_declarator LBRACE.constant_expression RBRACE
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      RBRACE -> State 324
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 325

  State 297:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM.RPARAM
      direct_abstract_declarator: direct_abstract_declarator LPARAM.parameter_type_list RPARAM
    Reduce:
      (nil)
    Goto:
      TYPE_NAME -> State 20
      TYPEDEF -> State 19
      EXTERN -> State 7
      STATIC -> State 17
      AUTO -> State 2
      REGISTER -> State 14
      CHAR -> State 3
      SHORT -> State 15
      INT -> State 10
      LONG -> State 11
      SIGNED -> State 16
      UNSIGNED -> State 22
      FLOAT -> State 8
      DOUBLE -> State 5
      CONST -> State 4
      VOLATILE -> State 24
      VOID -> State 23
      STRUCT -> State 18
      UNION -> State 21
      ENUM -> State 6
      RPARAM -> State 326
      declaration_specifiers -> State 137
      storage_class_specifier -> State 33
      type_specifier -> State 37
      struct_or_union_specifier -> State 35
      struct_or_union -> State 34
      enum_specifier -> State 29
      type_qualifier -> State 36
      parameter_type_list -> State 327
      parameter_list -> State 140
      parameter_declaration -> State 139

  State 298:
    Kernel Items:
      abstract_declarator: pointer direct_abstract_declarator., *
      direct_abstract_declarator: direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator: direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator: direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 297
      LBRACE -> State 296

  State 299:
    Kernel Items:
      identifier_list: identifier_list COMMA IDENTIFIER., *
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      parameter_type_list: parameter_list COMMA ELLIPSIS., RPARAM
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      parameter_list: parameter_list COMMA parameter_declaration., *
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      struct_or_union_specifier: struct_or_union IDENTIFIER LCURL struct_declaration_list RCURL., *
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      struct_declarator: COLON constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      struct_declarator: declarator COLON.constant_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 132
      constant_expression -> State 328

  State 305:
    Kernel Items:
      struct_declarator_list: struct_declarator_list COMMA.struct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      COLON -> State 234
      MUL -> State 13
      struct_declarator -> State 329
      declarator -> State 235
      direct_declarator -> State 28
      pointer -> State 32

  State 306:
    Kernel Items:
      struct_declaration: specifier_qualifier_list struct_declarator_list SEMICOLON., *
    Reduce:
      * -> [struct_declaration]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      initializer: LCURL initializer_list COMMA.RCURL
      initializer_list: initializer_list COMMA.initializer
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      LCURL -> State 152
      RCURL -> State 330
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 153
      initializer -> State 331

  State 308:
    Kernel Items:
      initializer: LCURL initializer_list RCURL., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      labeled_statement: CASE constant_expression COLON statement., *
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM.expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 332

  State 311:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement.RPARAM statement
      iteration_statement: FOR LPARAM expression_statement expression_statement.expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      RPARAM -> State 333
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 334

  State 312:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM.statement
      selection_statement: IF LPARAM expression RPARAM.statement ELSE statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 337
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 313:
    Kernel Items:
      cast_expression: LPARAM type_name RPARAM cast_expression., *
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      unary_expression: SIZEOF LPARAM type_name RPARAM., *
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      selection_statement: SWITCH LPARAM expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 338
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 316:
    Kernel Items:
      iteration_statement: WHILE LPARAM expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 339
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 317:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression COLON.conditional_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 134
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 340

  State 318:
    Kernel Items:
      postfix_expression: postfix_expression LBRACE expression RBRACE., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      argument_expression_list: argument_expression_list COMMA.assignment_expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      LPARAM -> State 88
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 341

  State 320:
    Kernel Items:
      postfix_expression: postfix_expression LPARAM argument_expression_list RPARAM., *
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      direct_abstract_declarator: LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      direct_abstract_declarator: LPARAM abstract_declarator RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      direct_abstract_declarator: LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 342

  State 326:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 343

  State 328:
    Kernel Items:
      struct_declarator: declarator COLON constant_expression., *
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      struct_declarator_list: struct_declarator_list COMMA struct_declarator., *
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      initializer: LCURL initializer_list COMMA RCURL., *
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      initializer_list: initializer_list COMMA initializer., *
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: DO statement WHILE LPARAM expression.RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 344
      COMMA -> State 186

  State 333:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 345
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 334:
    Kernel Items:
      expression: expression.COMMA assignment_expression
      iteration_statement: FOR LPARAM expression_statement expression_statement expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 346
      COMMA -> State 186

  State 335:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement., *
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement
    Reduce:
      * -> [selection_statement]
    Goto:
      ELSE -> State 347

  State 336:
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
      ELSE -> State 347
    Shift/reduce conflict symbols:
      [ELSE]

  State 337:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement., ELSE
      selection_statement: IF LPARAM expression RPARAM statement., WHILE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , ELSE
      selection_statement: IF LPARAM expression RPARAM statement.ELSE statement , WHILE
    Reduce:
      ELSE -> [selection_statement]
      WHILE -> [selection_statement]
    Goto:
      ELSE -> State 347
    Shift/reduce conflict symbols:
      [ELSE]

  State 338:
    Kernel Items:
      selection_statement: SWITCH LPARAM expression RPARAM statement., *
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      iteration_statement: WHILE LPARAM expression RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      conditional_expression: logical_or_expression QUESTION expression COLON conditional_expression., *
    Reduce:
      * -> [conditional_expression]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      argument_expression_list: argument_expression_list COMMA assignment_expression., *
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LBRACE constant_expression RBRACE., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      direct_abstract_declarator: direct_abstract_declarator LPARAM parameter_type_list RPARAM., *
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM expression RPARAM.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 348

  State 345:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement expression RPARAM.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 349
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 347:
    Kernel Items:
      selection_statement: IF LPARAM expression RPARAM statement ELSE.statement
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 85
      CONSTANT -> State 77
      STRING_LITERAL -> State 96
      SIZEOF -> State 95
      INC_OP -> State 87
      DEC_OP -> State 79
      CASE -> State 76
      DEFAULT -> State 80
      IF -> State 86
      SWITCH -> State 97
      WHILE -> State 99
      DO -> State 81
      FOR -> State 83
      GOTO -> State 84
      CONTINUE -> State 78
      BREAK -> State 75
      RETURN -> State 93
      LPARAM -> State 88
      LCURL -> State 49
      SEMICOLON -> State 94
      MUL -> State 90
      MINUS -> State 89
      PLUS -> State 91
      AND -> State 74
      EXCLAIM -> State 82
      TILDA -> State 98
      primary_expression -> State 119
      postfix_expression -> State 118
      unary_expression -> State 125
      unary_operator -> State 126
      cast_expression -> State 103
      multiplicative_expression -> State 117
      additive_expression -> State 100
      shift_expression -> State 122
      relational_expression -> State 120
      equality_expression -> State 107
      and_expression -> State 101
      exclusive_or_expression -> State 108
      inclusive_or_expression -> State 111
      logical_and_expression -> State 115
      logical_or_expression -> State 116
      conditional_expression -> State 105
      assignment_expression -> State 102
      expression -> State 109
      statement -> State 350
      labeled_statement -> State 114
      compound_statement -> State 104
      expression_statement -> State 110
      selection_statement -> State 121
      iteration_statement -> State 112
      jump_statement -> State 113

  State 348:
    Kernel Items:
      iteration_statement: DO statement WHILE LPARAM expression RPARAM SEMICOLON., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      iteration_statement: FOR LPARAM expression_statement expression_statement expression RPARAM statement., *
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 350:
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
