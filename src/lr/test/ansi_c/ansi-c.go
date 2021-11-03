// Auto-generated from source: ansi-c.lr

package ansi_c

import (
	fmt "fmt"
	io "io"
)

type CSymbolId string

const (
	CIdentifierSymbol    = CSymbolId("IDENTIFIER")
	CConstantSymbol      = CSymbolId("CONSTANT")
	CStringLiteralSymbol = CSymbolId("STRING_LITERAL")
	CSizeofSymbol        = CSymbolId("SIZEOF")
	CPtrOpSymbol         = CSymbolId("PTR_OP")
	CIncOpSymbol         = CSymbolId("INC_OP")
	CDecOpSymbol         = CSymbolId("DEC_OP")
	CLeftOpSymbol        = CSymbolId("LEFT_OP")
	CRightOpSymbol       = CSymbolId("RIGHT_OP")
	CLeOpSymbol          = CSymbolId("LE_OP")
	CGeOpSymbol          = CSymbolId("GE_OP")
	CEqOpSymbol          = CSymbolId("EQ_OP")
	CNeOpSymbol          = CSymbolId("NE_OP")
	CAndOpSymbol         = CSymbolId("AND_OP")
	COrOpSymbol          = CSymbolId("OR_OP")
	CMulAssignSymbol     = CSymbolId("MUL_ASSIGN")
	CDivAssignSymbol     = CSymbolId("DIV_ASSIGN")
	CModAssignSymbol     = CSymbolId("MOD_ASSIGN")
	CAddAssignSymbol     = CSymbolId("ADD_ASSIGN")
	CSubAssignSymbol     = CSymbolId("SUB_ASSIGN")
	CLeftAssignSymbol    = CSymbolId("LEFT_ASSIGN")
	CRightAssignSymbol   = CSymbolId("RIGHT_ASSIGN")
	CAndAssignSymbol     = CSymbolId("AND_ASSIGN")
	CXorAssignSymbol     = CSymbolId("XOR_ASSIGN")
	COrAssignSymbol      = CSymbolId("OR_ASSIGN")
	CTypeNameSymbol      = CSymbolId("TYPE_NAME")
	CTypedefSymbol       = CSymbolId("TYPEDEF")
	CExternSymbol        = CSymbolId("EXTERN")
	CStaticSymbol        = CSymbolId("STATIC")
	CAutoSymbol          = CSymbolId("AUTO")
	CRegisterSymbol      = CSymbolId("REGISTER")
	CCharSymbol          = CSymbolId("CHAR")
	CShortSymbol         = CSymbolId("SHORT")
	CIntSymbol           = CSymbolId("INT")
	CLongSymbol          = CSymbolId("LONG")
	CSignedSymbol        = CSymbolId("SIGNED")
	CUnsignedSymbol      = CSymbolId("UNSIGNED")
	CFloatSymbol         = CSymbolId("FLOAT")
	CDoubleSymbol        = CSymbolId("DOUBLE")
	CConstSymbol         = CSymbolId("CONST")
	CVolatileSymbol      = CSymbolId("VOLATILE")
	CVoidSymbol          = CSymbolId("VOID")
	CStructSymbol        = CSymbolId("STRUCT")
	CUnionSymbol         = CSymbolId("UNION")
	CEnumSymbol          = CSymbolId("ENUM")
	CEllipsisSymbol      = CSymbolId("ELLIPSIS")
	CCaseSymbol          = CSymbolId("CASE")
	CDefaultSymbol       = CSymbolId("DEFAULT")
	CIfSymbol            = CSymbolId("IF")
	CElseSymbol          = CSymbolId("ELSE")
	CSwitchSymbol        = CSymbolId("SWITCH")
	CWhileSymbol         = CSymbolId("WHILE")
	CDoSymbol            = CSymbolId("DO")
	CForSymbol           = CSymbolId("FOR")
	CGotoSymbol          = CSymbolId("GOTO")
	CContinueSymbol      = CSymbolId("CONTINUE")
	CBreakSymbol         = CSymbolId("BREAK")
	CReturnSymbol        = CSymbolId("RETURN")
	CLparamSymbol        = CSymbolId("LPARAM")
	CRparamSymbol        = CSymbolId("RPARAM")
	CLcurlSymbol         = CSymbolId("LCURL")
	CRcurlSymbol         = CSymbolId("RCURL")
	CLbraceSymbol        = CSymbolId("LBRACE")
	CRbraceSymbol        = CSymbolId("RBRACE")
	CSemicolonSymbol     = CSymbolId("SEMICOLON")
	CColonSymbol         = CSymbolId("COLON")
	CCommaSymbol         = CSymbolId("COMMA")
	CEqSymbol            = CSymbolId("EQ")
	CQuestionSymbol      = CSymbolId("QUESTION")
	CMulSymbol           = CSymbolId("MUL")
	CDivSymbol           = CSymbolId("DIV")
	CMinusSymbol         = CSymbolId("MINUS")
	CPlusSymbol          = CSymbolId("PLUS")
	CModSymbol           = CSymbolId("MOD")
	CAndSymbol           = CSymbolId("AND")
	COrSymbol            = CSymbolId("OR")
	CExclaimSymbol       = CSymbolId("EXCLAIM")
	CDotSymbol           = CSymbolId("DOT")
	CHatSymbol           = CSymbolId("HAT")
	CLtSymbol            = CSymbolId("LT")
	CGtSymbol            = CSymbolId("GT")
	CTildaSymbol         = CSymbolId("TILDA")
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

type CSymbol interface {
	Id() CSymbolId
	Location() CLocation
}

type CLexer interface {
	// Note: Return io.EOF to indicate end of stream
	Next() (CSymbol, error)
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
	Error(nextToken CSymbol, parseStack _CStack) error
}

type CDefaultParseErrorHandler struct{}

func (CDefaultParseErrorHandler) Error(nextToken CSymbol, stack _CStack) error {
	return fmt.Errorf("Syntax error: unexpected symbol %v. Expecting: %v (%v)", nextToken.Id(), _CExpectedTerminals[stack[len(stack)-1].StateId], nextToken.Location())
}

func CParse(lexer CLexer, reducer CReducer) (Symbol, error) {
	return CParseWithCustomErrorHandler(lexer, reducer, CDefaultParseErrorHandler{})
}

func CParseWithCustomErrorHandler(lexer CLexer, reducer CReducer, errHandler CParseErrorHandler) (Symbol, error) {
	var errRetVal Symbol
	stateStack := _CStack{
		&_CStackItem{_CState0, &_CSymbol{SymbolId_: _CStartMarker}},
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
			var reduceSymbol CSymbol
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
			panic("Unknown action type: " + action.ActionType)
		}
	}
}

// =======================================================
// Parser internal implementation
// User should avoid directly accessing the following code
// =======================================================

const (
	_CAcceptMarker   = CSymbolId("#accept")
	_CStartMarker    = CSymbolId("^")
	_CEndMarker      = CSymbolId("$")
	_CWildcardMarker = CSymbolId("*")

	_CPrimaryExpressionSymbol        = CSymbolId("primary_expression")
	_CPostfixExpressionSymbol        = CSymbolId("postfix_expression")
	_CArgumentExpressionListSymbol   = CSymbolId("argument_expression_list")
	_CUnaryExpressionSymbol          = CSymbolId("unary_expression")
	_CUnaryOperatorSymbol            = CSymbolId("unary_operator")
	_CCastExpressionSymbol           = CSymbolId("cast_expression")
	_CMultiplicativeExpressionSymbol = CSymbolId("multiplicative_expression")
	_CAdditiveExpressionSymbol       = CSymbolId("additive_expression")
	_CShiftExpressionSymbol          = CSymbolId("shift_expression")
	_CRelationalExpressionSymbol     = CSymbolId("relational_expression")
	_CEqualityExpressionSymbol       = CSymbolId("equality_expression")
	_CAndExpressionSymbol            = CSymbolId("and_expression")
	_CExclusiveOrExpressionSymbol    = CSymbolId("exclusive_or_expression")
	_CInclusiveOrExpressionSymbol    = CSymbolId("inclusive_or_expression")
	_CLogicalAndExpressionSymbol     = CSymbolId("logical_and_expression")
	_CLogicalOrExpressionSymbol      = CSymbolId("logical_or_expression")
	_CConditionalExpressionSymbol    = CSymbolId("conditional_expression")
	_CAssignmentExpressionSymbol     = CSymbolId("assignment_expression")
	_CAssignmentOperatorSymbol       = CSymbolId("assignment_operator")
	_CExpressionSymbol               = CSymbolId("expression")
	_CConstantExpressionSymbol       = CSymbolId("constant_expression")
	_CDeclarationSymbol              = CSymbolId("declaration")
	_CDeclarationSpecifiersSymbol    = CSymbolId("declaration_specifiers")
	_CInitDeclaratorListSymbol       = CSymbolId("init_declarator_list")
	_CInitDeclaratorSymbol           = CSymbolId("init_declarator")
	_CStorageClassSpecifierSymbol    = CSymbolId("storage_class_specifier")
	_CTypeSpecifierSymbol            = CSymbolId("type_specifier")
	_CStructOrUnionSpecifierSymbol   = CSymbolId("struct_or_union_specifier")
	_CStructOrUnionSymbol            = CSymbolId("struct_or_union")
	_CStructDeclarationListSymbol    = CSymbolId("struct_declaration_list")
	_CStructDeclarationSymbol        = CSymbolId("struct_declaration")
	_CSpecifierQualifierListSymbol   = CSymbolId("specifier_qualifier_list")
	_CStructDeclaratorListSymbol     = CSymbolId("struct_declarator_list")
	_CStructDeclaratorSymbol         = CSymbolId("struct_declarator")
	_CEnumSpecifierSymbol            = CSymbolId("enum_specifier")
	_CEnumeratorListSymbol           = CSymbolId("enumerator_list")
	_CEnumeratorSymbol               = CSymbolId("enumerator")
	_CTypeQualifierSymbol            = CSymbolId("type_qualifier")
	_CDeclaratorSymbol               = CSymbolId("declarator")
	_CDirectDeclaratorSymbol         = CSymbolId("direct_declarator")
	_CPointerSymbol                  = CSymbolId("pointer")
	_CTypeQualifierListSymbol        = CSymbolId("type_qualifier_list")
	_CParameterTypeListSymbol        = CSymbolId("parameter_type_list")
	_CParameterListSymbol            = CSymbolId("parameter_list")
	_CParameterDeclarationSymbol     = CSymbolId("parameter_declaration")
	_CIdentifierListSymbol           = CSymbolId("identifier_list")
	_CTypeNameSymbol                 = CSymbolId("type_name")
	_CAbstractDeclaratorSymbol       = CSymbolId("abstract_declarator")
	_CDirectAbstractDeclaratorSymbol = CSymbolId("direct_abstract_declarator")
	_CInitializerSymbol              = CSymbolId("initializer")
	_CInitializerListSymbol          = CSymbolId("initializer_list")
	_CStatementSymbol                = CSymbolId("statement")
	_CLabeledStatementSymbol         = CSymbolId("labeled_statement")
	_CCompoundStatementSymbol        = CSymbolId("compound_statement")
	_CDeclarationListSymbol          = CSymbolId("declaration_list")
	_CStatementListSymbol            = CSymbolId("statement_list")
	_CExpressionStatementSymbol      = CSymbolId("expression_statement")
	_CSelectionStatementSymbol       = CSymbolId("selection_statement")
	_CIterationStatementSymbol       = CSymbolId("iteration_statement")
	_CJumpStatementSymbol            = CSymbolId("jump_statement")
	_CTranslationUnitSymbol          = CSymbolId("translation_unit")
	_CExternalDeclarationSymbol      = CSymbolId("external_declaration")
	_CFunctionDefinitionSymbol       = CSymbolId("function_definition")
)

type _CActionType string

const (
	// NOTE: error action is implicit
	_CShiftAction  = _CActionType("shift")
	_CReduceAction = _CActionType("reduce")
	_CAcceptAction = _CActionType("accept")
)

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

type _CStateId string

const (
	_CState0   = _CStateId("State 0")
	_CState1   = _CStateId("State 1")
	_CState2   = _CStateId("State 2")
	_CState3   = _CStateId("State 3")
	_CState4   = _CStateId("State 4")
	_CState5   = _CStateId("State 5")
	_CState6   = _CStateId("State 6")
	_CState7   = _CStateId("State 7")
	_CState8   = _CStateId("State 8")
	_CState9   = _CStateId("State 9")
	_CState10  = _CStateId("State 10")
	_CState11  = _CStateId("State 11")
	_CState12  = _CStateId("State 12")
	_CState13  = _CStateId("State 13")
	_CState14  = _CStateId("State 14")
	_CState15  = _CStateId("State 15")
	_CState16  = _CStateId("State 16")
	_CState17  = _CStateId("State 17")
	_CState18  = _CStateId("State 18")
	_CState19  = _CStateId("State 19")
	_CState20  = _CStateId("State 20")
	_CState21  = _CStateId("State 21")
	_CState22  = _CStateId("State 22")
	_CState23  = _CStateId("State 23")
	_CState24  = _CStateId("State 24")
	_CState25  = _CStateId("State 25")
	_CState26  = _CStateId("State 26")
	_CState27  = _CStateId("State 27")
	_CState28  = _CStateId("State 28")
	_CState29  = _CStateId("State 29")
	_CState30  = _CStateId("State 30")
	_CState31  = _CStateId("State 31")
	_CState32  = _CStateId("State 32")
	_CState33  = _CStateId("State 33")
	_CState34  = _CStateId("State 34")
	_CState35  = _CStateId("State 35")
	_CState36  = _CStateId("State 36")
	_CState37  = _CStateId("State 37")
	_CState38  = _CStateId("State 38")
	_CState39  = _CStateId("State 39")
	_CState40  = _CStateId("State 40")
	_CState41  = _CStateId("State 41")
	_CState42  = _CStateId("State 42")
	_CState43  = _CStateId("State 43")
	_CState44  = _CStateId("State 44")
	_CState45  = _CStateId("State 45")
	_CState46  = _CStateId("State 46")
	_CState47  = _CStateId("State 47")
	_CState48  = _CStateId("State 48")
	_CState49  = _CStateId("State 49")
	_CState50  = _CStateId("State 50")
	_CState51  = _CStateId("State 51")
	_CState52  = _CStateId("State 52")
	_CState53  = _CStateId("State 53")
	_CState54  = _CStateId("State 54")
	_CState55  = _CStateId("State 55")
	_CState56  = _CStateId("State 56")
	_CState57  = _CStateId("State 57")
	_CState58  = _CStateId("State 58")
	_CState59  = _CStateId("State 59")
	_CState60  = _CStateId("State 60")
	_CState61  = _CStateId("State 61")
	_CState62  = _CStateId("State 62")
	_CState63  = _CStateId("State 63")
	_CState64  = _CStateId("State 64")
	_CState65  = _CStateId("State 65")
	_CState66  = _CStateId("State 66")
	_CState67  = _CStateId("State 67")
	_CState68  = _CStateId("State 68")
	_CState69  = _CStateId("State 69")
	_CState70  = _CStateId("State 70")
	_CState71  = _CStateId("State 71")
	_CState72  = _CStateId("State 72")
	_CState73  = _CStateId("State 73")
	_CState74  = _CStateId("State 74")
	_CState75  = _CStateId("State 75")
	_CState76  = _CStateId("State 76")
	_CState77  = _CStateId("State 77")
	_CState78  = _CStateId("State 78")
	_CState79  = _CStateId("State 79")
	_CState80  = _CStateId("State 80")
	_CState81  = _CStateId("State 81")
	_CState82  = _CStateId("State 82")
	_CState83  = _CStateId("State 83")
	_CState84  = _CStateId("State 84")
	_CState85  = _CStateId("State 85")
	_CState86  = _CStateId("State 86")
	_CState87  = _CStateId("State 87")
	_CState88  = _CStateId("State 88")
	_CState89  = _CStateId("State 89")
	_CState90  = _CStateId("State 90")
	_CState91  = _CStateId("State 91")
	_CState92  = _CStateId("State 92")
	_CState93  = _CStateId("State 93")
	_CState94  = _CStateId("State 94")
	_CState95  = _CStateId("State 95")
	_CState96  = _CStateId("State 96")
	_CState97  = _CStateId("State 97")
	_CState98  = _CStateId("State 98")
	_CState99  = _CStateId("State 99")
	_CState100 = _CStateId("State 100")
	_CState101 = _CStateId("State 101")
	_CState102 = _CStateId("State 102")
	_CState103 = _CStateId("State 103")
	_CState104 = _CStateId("State 104")
	_CState105 = _CStateId("State 105")
	_CState106 = _CStateId("State 106")
	_CState107 = _CStateId("State 107")
	_CState108 = _CStateId("State 108")
	_CState109 = _CStateId("State 109")
	_CState110 = _CStateId("State 110")
	_CState111 = _CStateId("State 111")
	_CState112 = _CStateId("State 112")
	_CState113 = _CStateId("State 113")
	_CState114 = _CStateId("State 114")
	_CState115 = _CStateId("State 115")
	_CState116 = _CStateId("State 116")
	_CState117 = _CStateId("State 117")
	_CState118 = _CStateId("State 118")
	_CState119 = _CStateId("State 119")
	_CState120 = _CStateId("State 120")
	_CState121 = _CStateId("State 121")
	_CState122 = _CStateId("State 122")
	_CState123 = _CStateId("State 123")
	_CState124 = _CStateId("State 124")
	_CState125 = _CStateId("State 125")
	_CState126 = _CStateId("State 126")
	_CState127 = _CStateId("State 127")
	_CState128 = _CStateId("State 128")
	_CState129 = _CStateId("State 129")
	_CState130 = _CStateId("State 130")
	_CState131 = _CStateId("State 131")
	_CState132 = _CStateId("State 132")
	_CState133 = _CStateId("State 133")
	_CState134 = _CStateId("State 134")
	_CState135 = _CStateId("State 135")
	_CState136 = _CStateId("State 136")
	_CState137 = _CStateId("State 137")
	_CState138 = _CStateId("State 138")
	_CState139 = _CStateId("State 139")
	_CState140 = _CStateId("State 140")
	_CState141 = _CStateId("State 141")
	_CState142 = _CStateId("State 142")
	_CState143 = _CStateId("State 143")
	_CState144 = _CStateId("State 144")
	_CState145 = _CStateId("State 145")
	_CState146 = _CStateId("State 146")
	_CState147 = _CStateId("State 147")
	_CState148 = _CStateId("State 148")
	_CState149 = _CStateId("State 149")
	_CState150 = _CStateId("State 150")
	_CState151 = _CStateId("State 151")
	_CState152 = _CStateId("State 152")
	_CState153 = _CStateId("State 153")
	_CState154 = _CStateId("State 154")
	_CState155 = _CStateId("State 155")
	_CState156 = _CStateId("State 156")
	_CState157 = _CStateId("State 157")
	_CState158 = _CStateId("State 158")
	_CState159 = _CStateId("State 159")
	_CState160 = _CStateId("State 160")
	_CState161 = _CStateId("State 161")
	_CState162 = _CStateId("State 162")
	_CState163 = _CStateId("State 163")
	_CState164 = _CStateId("State 164")
	_CState165 = _CStateId("State 165")
	_CState166 = _CStateId("State 166")
	_CState167 = _CStateId("State 167")
	_CState168 = _CStateId("State 168")
	_CState169 = _CStateId("State 169")
	_CState170 = _CStateId("State 170")
	_CState171 = _CStateId("State 171")
	_CState172 = _CStateId("State 172")
	_CState173 = _CStateId("State 173")
	_CState174 = _CStateId("State 174")
	_CState175 = _CStateId("State 175")
	_CState176 = _CStateId("State 176")
	_CState177 = _CStateId("State 177")
	_CState178 = _CStateId("State 178")
	_CState179 = _CStateId("State 179")
	_CState180 = _CStateId("State 180")
	_CState181 = _CStateId("State 181")
	_CState182 = _CStateId("State 182")
	_CState183 = _CStateId("State 183")
	_CState184 = _CStateId("State 184")
	_CState185 = _CStateId("State 185")
	_CState186 = _CStateId("State 186")
	_CState187 = _CStateId("State 187")
	_CState188 = _CStateId("State 188")
	_CState189 = _CStateId("State 189")
	_CState190 = _CStateId("State 190")
	_CState191 = _CStateId("State 191")
	_CState192 = _CStateId("State 192")
	_CState193 = _CStateId("State 193")
	_CState194 = _CStateId("State 194")
	_CState195 = _CStateId("State 195")
	_CState196 = _CStateId("State 196")
	_CState197 = _CStateId("State 197")
	_CState198 = _CStateId("State 198")
	_CState199 = _CStateId("State 199")
	_CState200 = _CStateId("State 200")
	_CState201 = _CStateId("State 201")
	_CState202 = _CStateId("State 202")
	_CState203 = _CStateId("State 203")
	_CState204 = _CStateId("State 204")
	_CState205 = _CStateId("State 205")
	_CState206 = _CStateId("State 206")
	_CState207 = _CStateId("State 207")
	_CState208 = _CStateId("State 208")
	_CState209 = _CStateId("State 209")
	_CState210 = _CStateId("State 210")
	_CState211 = _CStateId("State 211")
	_CState212 = _CStateId("State 212")
	_CState213 = _CStateId("State 213")
	_CState214 = _CStateId("State 214")
	_CState215 = _CStateId("State 215")
	_CState216 = _CStateId("State 216")
	_CState217 = _CStateId("State 217")
	_CState218 = _CStateId("State 218")
	_CState219 = _CStateId("State 219")
	_CState220 = _CStateId("State 220")
	_CState221 = _CStateId("State 221")
	_CState222 = _CStateId("State 222")
	_CState223 = _CStateId("State 223")
	_CState224 = _CStateId("State 224")
	_CState225 = _CStateId("State 225")
	_CState226 = _CStateId("State 226")
	_CState227 = _CStateId("State 227")
	_CState228 = _CStateId("State 228")
	_CState229 = _CStateId("State 229")
	_CState230 = _CStateId("State 230")
	_CState231 = _CStateId("State 231")
	_CState232 = _CStateId("State 232")
	_CState233 = _CStateId("State 233")
	_CState234 = _CStateId("State 234")
	_CState235 = _CStateId("State 235")
	_CState236 = _CStateId("State 236")
	_CState237 = _CStateId("State 237")
	_CState238 = _CStateId("State 238")
	_CState239 = _CStateId("State 239")
	_CState240 = _CStateId("State 240")
	_CState241 = _CStateId("State 241")
	_CState242 = _CStateId("State 242")
	_CState243 = _CStateId("State 243")
	_CState244 = _CStateId("State 244")
	_CState245 = _CStateId("State 245")
	_CState246 = _CStateId("State 246")
	_CState247 = _CStateId("State 247")
	_CState248 = _CStateId("State 248")
	_CState249 = _CStateId("State 249")
	_CState250 = _CStateId("State 250")
	_CState251 = _CStateId("State 251")
	_CState252 = _CStateId("State 252")
	_CState253 = _CStateId("State 253")
	_CState254 = _CStateId("State 254")
	_CState255 = _CStateId("State 255")
	_CState256 = _CStateId("State 256")
	_CState257 = _CStateId("State 257")
	_CState258 = _CStateId("State 258")
	_CState259 = _CStateId("State 259")
	_CState260 = _CStateId("State 260")
	_CState261 = _CStateId("State 261")
	_CState262 = _CStateId("State 262")
	_CState263 = _CStateId("State 263")
	_CState264 = _CStateId("State 264")
	_CState265 = _CStateId("State 265")
	_CState266 = _CStateId("State 266")
	_CState267 = _CStateId("State 267")
	_CState268 = _CStateId("State 268")
	_CState269 = _CStateId("State 269")
	_CState270 = _CStateId("State 270")
	_CState271 = _CStateId("State 271")
	_CState272 = _CStateId("State 272")
	_CState273 = _CStateId("State 273")
	_CState274 = _CStateId("State 274")
	_CState275 = _CStateId("State 275")
	_CState276 = _CStateId("State 276")
	_CState277 = _CStateId("State 277")
	_CState278 = _CStateId("State 278")
	_CState279 = _CStateId("State 279")
	_CState280 = _CStateId("State 280")
	_CState281 = _CStateId("State 281")
	_CState282 = _CStateId("State 282")
	_CState283 = _CStateId("State 283")
	_CState284 = _CStateId("State 284")
	_CState285 = _CStateId("State 285")
	_CState286 = _CStateId("State 286")
	_CState287 = _CStateId("State 287")
	_CState288 = _CStateId("State 288")
	_CState289 = _CStateId("State 289")
	_CState290 = _CStateId("State 290")
	_CState291 = _CStateId("State 291")
	_CState292 = _CStateId("State 292")
	_CState293 = _CStateId("State 293")
	_CState294 = _CStateId("State 294")
	_CState295 = _CStateId("State 295")
	_CState296 = _CStateId("State 296")
	_CState297 = _CStateId("State 297")
	_CState298 = _CStateId("State 298")
	_CState299 = _CStateId("State 299")
	_CState300 = _CStateId("State 300")
	_CState301 = _CStateId("State 301")
	_CState302 = _CStateId("State 302")
	_CState303 = _CStateId("State 303")
	_CState304 = _CStateId("State 304")
	_CState305 = _CStateId("State 305")
	_CState306 = _CStateId("State 306")
	_CState307 = _CStateId("State 307")
	_CState308 = _CStateId("State 308")
	_CState309 = _CStateId("State 309")
	_CState310 = _CStateId("State 310")
	_CState311 = _CStateId("State 311")
	_CState312 = _CStateId("State 312")
	_CState313 = _CStateId("State 313")
	_CState314 = _CStateId("State 314")
	_CState315 = _CStateId("State 315")
	_CState316 = _CStateId("State 316")
	_CState317 = _CStateId("State 317")
	_CState318 = _CStateId("State 318")
	_CState319 = _CStateId("State 319")
	_CState320 = _CStateId("State 320")
	_CState321 = _CStateId("State 321")
	_CState322 = _CStateId("State 322")
	_CState323 = _CStateId("State 323")
	_CState324 = _CStateId("State 324")
	_CState325 = _CStateId("State 325")
	_CState326 = _CStateId("State 326")
	_CState327 = _CStateId("State 327")
	_CState328 = _CStateId("State 328")
	_CState329 = _CStateId("State 329")
	_CState330 = _CStateId("State 330")
	_CState331 = _CStateId("State 331")
	_CState332 = _CStateId("State 332")
	_CState333 = _CStateId("State 333")
	_CState334 = _CStateId("State 334")
	_CState335 = _CStateId("State 335")
	_CState336 = _CStateId("State 336")
	_CState337 = _CStateId("State 337")
	_CState338 = _CStateId("State 338")
	_CState339 = _CStateId("State 339")
	_CState340 = _CStateId("State 340")
	_CState341 = _CStateId("State 341")
	_CState342 = _CStateId("State 342")
	_CState343 = _CStateId("State 343")
	_CState344 = _CStateId("State 344")
	_CState345 = _CStateId("State 345")
	_CState346 = _CStateId("State 346")
	_CState347 = _CStateId("State 347")
	_CState348 = _CStateId("State 348")
	_CState349 = _CStateId("State 349")
	_CState350 = _CStateId("State 350")
)

type _CSymbol struct {
	SymbolId_ CSymbolId

	T Symbol
}

func (s *_CSymbol) Id() CSymbolId {
	return s.SymbolId_
}

func (s *_CSymbol) Location() CLocation {
	type locator interface{ Location() CLocation }
	switch s.SymbolId_ {
	case CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CPtrOpSymbol, CIncOpSymbol, CDecOpSymbol, CLeftOpSymbol, CRightOpSymbol, CLeOpSymbol, CGeOpSymbol, CEqOpSymbol, CNeOpSymbol, CAndOpSymbol, COrOpSymbol, CMulAssignSymbol, CDivAssignSymbol, CModAssignSymbol, CAddAssignSymbol, CSubAssignSymbol, CLeftAssignSymbol, CRightAssignSymbol, CAndAssignSymbol, CXorAssignSymbol, COrAssignSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CEllipsisSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CElseSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CRparamSymbol, CLcurlSymbol, CRcurlSymbol, CLbraceSymbol, CRbraceSymbol, CSemicolonSymbol, CColonSymbol, CCommaSymbol, CEqSymbol, CQuestionSymbol, CMulSymbol, CDivSymbol, CMinusSymbol, CPlusSymbol, CModSymbol, CAndSymbol, COrSymbol, CExclaimSymbol, CDotSymbol, CHatSymbol, CLtSymbol, CGtSymbol, CTildaSymbol, _CPrimaryExpressionSymbol, _CPostfixExpressionSymbol, _CArgumentExpressionListSymbol, _CUnaryExpressionSymbol, _CUnaryOperatorSymbol, _CCastExpressionSymbol, _CMultiplicativeExpressionSymbol, _CAdditiveExpressionSymbol, _CShiftExpressionSymbol, _CRelationalExpressionSymbol, _CEqualityExpressionSymbol, _CAndExpressionSymbol, _CExclusiveOrExpressionSymbol, _CInclusiveOrExpressionSymbol, _CLogicalAndExpressionSymbol, _CLogicalOrExpressionSymbol, _CConditionalExpressionSymbol, _CAssignmentExpressionSymbol, _CAssignmentOperatorSymbol, _CExpressionSymbol, _CConstantExpressionSymbol, _CDeclarationSymbol, _CDeclarationSpecifiersSymbol, _CInitDeclaratorListSymbol, _CInitDeclaratorSymbol, _CStorageClassSpecifierSymbol, _CTypeSpecifierSymbol, _CStructOrUnionSpecifierSymbol, _CStructOrUnionSymbol, _CStructDeclarationListSymbol, _CStructDeclarationSymbol, _CSpecifierQualifierListSymbol, _CStructDeclaratorListSymbol, _CStructDeclaratorSymbol, _CEnumSpecifierSymbol, _CEnumeratorListSymbol, _CEnumeratorSymbol, _CTypeQualifierSymbol, _CDeclaratorSymbol, _CDirectDeclaratorSymbol, _CPointerSymbol, _CTypeQualifierListSymbol, _CParameterTypeListSymbol, _CParameterListSymbol, _CParameterDeclarationSymbol, _CIdentifierListSymbol, _CTypeNameSymbol, _CAbstractDeclaratorSymbol, _CDirectAbstractDeclaratorSymbol, _CInitializerSymbol, _CInitializerListSymbol, _CStatementSymbol, _CLabeledStatementSymbol, _CCompoundStatementSymbol, _CDeclarationListSymbol, _CStatementListSymbol, _CExpressionStatementSymbol, _CSelectionStatementSymbol, _CIterationStatementSymbol, _CJumpStatementSymbol, _CTranslationUnitSymbol, _CExternalDeclarationSymbol, _CFunctionDefinitionSymbol:
		loc, ok := interface{}(s.T).(locator)
		if ok {
			return loc.Location()
		}
	}
	return CLocation{}
}

type _CPseudoToken CSymbolId

func (t _CPseudoToken) Id() CSymbolId { return CSymbolId(t) }

func (_CPseudoToken) Location() CLocation { return CLocation{} }

type _CPseudoSymbolStack struct {
	lexer CLexer
	top   []CSymbol
}

func (stack *_CPseudoSymbolStack) Top() (CSymbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = _CPseudoToken(_CEndMarker)
		}
		stack.top = append(stack.top, token)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_CPseudoSymbolStack) Push(symbol CSymbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_CPseudoSymbolStack) Pop() (CSymbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _CStackItem struct {
	StateId _CStateId

	*_CSymbol
}

type _CStack []*_CStackItem

type _CAction struct {
	ActionType _CActionType

	ShiftStateId _CStateId
	ReduceType   _CReduceType
}

func (act *_CAction) ShiftItem(symbol CSymbol) *_CStackItem {
	item := &_CStackItem{StateId: act.ShiftStateId}

	reducedSymbol, ok := symbol.(*_CSymbol)
	if ok {
		item._CSymbol = reducedSymbol
		return item
	}

	item._CSymbol = &_CSymbol{SymbolId_: symbol.Id()}

	switch symbol.Id() {
	case CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CPtrOpSymbol, CIncOpSymbol, CDecOpSymbol, CLeftOpSymbol, CRightOpSymbol, CLeOpSymbol, CGeOpSymbol, CEqOpSymbol, CNeOpSymbol, CAndOpSymbol, COrOpSymbol, CMulAssignSymbol, CDivAssignSymbol, CModAssignSymbol, CAddAssignSymbol, CSubAssignSymbol, CLeftAssignSymbol, CRightAssignSymbol, CAndAssignSymbol, CXorAssignSymbol, COrAssignSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CEllipsisSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CElseSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CRparamSymbol, CLcurlSymbol, CRcurlSymbol, CLbraceSymbol, CRbraceSymbol, CSemicolonSymbol, CColonSymbol, CCommaSymbol, CEqSymbol, CQuestionSymbol, CMulSymbol, CDivSymbol, CMinusSymbol, CPlusSymbol, CModSymbol, CAndSymbol, COrSymbol, CExclaimSymbol, CDotSymbol, CHatSymbol, CLtSymbol, CGtSymbol, CTildaSymbol:
		item.T = symbol.(Symbol)
	case _CEndMarker: // EOF has no value
	default:
		panic("Unexpected symbol type: " + symbol.Id())
	}
	return item
}

func (act *_CAction) ReduceSymbol(reducer CReducer, stack _CStack) (_CStack, *_CSymbol, error) {
	var err error
	symbol := &_CSymbol{}
	switch act.ReduceType {
	case _CReduceAToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CPrimaryExpressionSymbol
		symbol.T, err = reducer.AToPrimaryExpression(args[0].T)
	case _CReduceBToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CPrimaryExpressionSymbol
		symbol.T, err = reducer.BToPrimaryExpression(args[0].T)
	case _CReduceCToPrimaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CPrimaryExpressionSymbol
		symbol.T, err = reducer.CToPrimaryExpression(args[0].T)
	case _CReduceDToPrimaryExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CPrimaryExpressionSymbol
		symbol.T, err = reducer.DToPrimaryExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToPostfixExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.AToPostfixExpression(args[0].T)
	case _CReduceBToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.BToPostfixExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.CToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToPostfixExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.DToPostfixExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceEToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.EToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceFToPostfixExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.FToPostfixExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceGToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.GToPostfixExpression(args[0].T, args[1].T)
	case _CReduceHToPostfixExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CPostfixExpressionSymbol
		symbol.T, err = reducer.HToPostfixExpression(args[0].T, args[1].T)
	case _CReduceAToArgumentExpressionList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CArgumentExpressionListSymbol
		symbol.T, err = reducer.AToArgumentExpressionList(args[0].T)
	case _CReduceBToArgumentExpressionList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CArgumentExpressionListSymbol
		symbol.T, err = reducer.BToArgumentExpressionList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToUnaryExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.AToUnaryExpression(args[0].T)
	case _CReduceBToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.BToUnaryExpression(args[0].T, args[1].T)
	case _CReduceCToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.CToUnaryExpression(args[0].T, args[1].T)
	case _CReduceDToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.DToUnaryExpression(args[0].T, args[1].T)
	case _CReduceEToUnaryExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.EToUnaryExpression(args[0].T, args[1].T)
	case _CReduceFToUnaryExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CUnaryExpressionSymbol
		symbol.T, err = reducer.FToUnaryExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.AToUnaryOperator(args[0].T)
	case _CReduceBToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.BToUnaryOperator(args[0].T)
	case _CReduceCToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.CToUnaryOperator(args[0].T)
	case _CReduceDToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.DToUnaryOperator(args[0].T)
	case _CReduceEToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.EToUnaryOperator(args[0].T)
	case _CReduceFToUnaryOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CUnaryOperatorSymbol
		symbol.T, err = reducer.FToUnaryOperator(args[0].T)
	case _CReduceAToCastExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CCastExpressionSymbol
		symbol.T, err = reducer.AToCastExpression(args[0].T)
	case _CReduceBToCastExpression:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CCastExpressionSymbol
		symbol.T, err = reducer.BToCastExpression(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToMultiplicativeExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CMultiplicativeExpressionSymbol
		symbol.T, err = reducer.AToMultiplicativeExpression(args[0].T)
	case _CReduceBToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CMultiplicativeExpressionSymbol
		symbol.T, err = reducer.BToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CMultiplicativeExpressionSymbol
		symbol.T, err = reducer.CToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToMultiplicativeExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CMultiplicativeExpressionSymbol
		symbol.T, err = reducer.DToMultiplicativeExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAdditiveExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAdditiveExpressionSymbol
		symbol.T, err = reducer.AToAdditiveExpression(args[0].T)
	case _CReduceBToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CAdditiveExpressionSymbol
		symbol.T, err = reducer.BToAdditiveExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToAdditiveExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CAdditiveExpressionSymbol
		symbol.T, err = reducer.CToAdditiveExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToShiftExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CShiftExpressionSymbol
		symbol.T, err = reducer.AToShiftExpression(args[0].T)
	case _CReduceBToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CShiftExpressionSymbol
		symbol.T, err = reducer.BToShiftExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToShiftExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CShiftExpressionSymbol
		symbol.T, err = reducer.CToShiftExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToRelationalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CRelationalExpressionSymbol
		symbol.T, err = reducer.AToRelationalExpression(args[0].T)
	case _CReduceBToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CRelationalExpressionSymbol
		symbol.T, err = reducer.BToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CRelationalExpressionSymbol
		symbol.T, err = reducer.CToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceDToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CRelationalExpressionSymbol
		symbol.T, err = reducer.DToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceEToRelationalExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CRelationalExpressionSymbol
		symbol.T, err = reducer.EToRelationalExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEqualityExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CEqualityExpressionSymbol
		symbol.T, err = reducer.AToEqualityExpression(args[0].T)
	case _CReduceBToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CEqualityExpressionSymbol
		symbol.T, err = reducer.BToEqualityExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceCToEqualityExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CEqualityExpressionSymbol
		symbol.T, err = reducer.CToEqualityExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAndExpressionSymbol
		symbol.T, err = reducer.AToAndExpression(args[0].T)
	case _CReduceBToAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CAndExpressionSymbol
		symbol.T, err = reducer.BToAndExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToExclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CExclusiveOrExpressionSymbol
		symbol.T, err = reducer.AToExclusiveOrExpression(args[0].T)
	case _CReduceBToExclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CExclusiveOrExpressionSymbol
		symbol.T, err = reducer.BToExclusiveOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToInclusiveOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CInclusiveOrExpressionSymbol
		symbol.T, err = reducer.AToInclusiveOrExpression(args[0].T)
	case _CReduceBToInclusiveOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CInclusiveOrExpressionSymbol
		symbol.T, err = reducer.BToInclusiveOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToLogicalAndExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CLogicalAndExpressionSymbol
		symbol.T, err = reducer.AToLogicalAndExpression(args[0].T)
	case _CReduceBToLogicalAndExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CLogicalAndExpressionSymbol
		symbol.T, err = reducer.BToLogicalAndExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToLogicalOrExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CLogicalOrExpressionSymbol
		symbol.T, err = reducer.AToLogicalOrExpression(args[0].T)
	case _CReduceBToLogicalOrExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CLogicalOrExpressionSymbol
		symbol.T, err = reducer.BToLogicalOrExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToConditionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CConditionalExpressionSymbol
		symbol.T, err = reducer.AToConditionalExpression(args[0].T)
	case _CReduceBToConditionalExpression:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CConditionalExpressionSymbol
		symbol.T, err = reducer.BToConditionalExpression(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceAToAssignmentExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentExpressionSymbol
		symbol.T, err = reducer.AToAssignmentExpression(args[0].T)
	case _CReduceBToAssignmentExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CAssignmentExpressionSymbol
		symbol.T, err = reducer.BToAssignmentExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.AToAssignmentOperator(args[0].T)
	case _CReduceBToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.BToAssignmentOperator(args[0].T)
	case _CReduceCToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.CToAssignmentOperator(args[0].T)
	case _CReduceDToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.DToAssignmentOperator(args[0].T)
	case _CReduceEToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.EToAssignmentOperator(args[0].T)
	case _CReduceFToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.FToAssignmentOperator(args[0].T)
	case _CReduceGToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.GToAssignmentOperator(args[0].T)
	case _CReduceHToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.HToAssignmentOperator(args[0].T)
	case _CReduceIToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.IToAssignmentOperator(args[0].T)
	case _CReduceJToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.JToAssignmentOperator(args[0].T)
	case _CReduceKToAssignmentOperator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAssignmentOperatorSymbol
		symbol.T, err = reducer.KToAssignmentOperator(args[0].T)
	case _CReduceAToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CExpressionSymbol
		symbol.T, err = reducer.AToExpression(args[0].T)
	case _CReduceBToExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CExpressionSymbol
		symbol.T, err = reducer.BToExpression(args[0].T, args[1].T, args[2].T)
	case _CReduceAToConstantExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CConstantExpressionSymbol
		symbol.T, err = reducer.AToConstantExpression(args[0].T)
	case _CReduceAToDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclarationSymbol
		symbol.T, err = reducer.AToDeclaration(args[0].T, args[1].T)
	case _CReduceBToDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDeclarationSymbol
		symbol.T, err = reducer.BToDeclaration(args[0].T, args[1].T, args[2].T)
	case _CReduceAToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.AToDeclarationSpecifiers(args[0].T)
	case _CReduceBToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.BToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceCToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.CToDeclarationSpecifiers(args[0].T)
	case _CReduceDToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.DToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceEToDeclarationSpecifiers:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.EToDeclarationSpecifiers(args[0].T)
	case _CReduceFToDeclarationSpecifiers:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclarationSpecifiersSymbol
		symbol.T, err = reducer.FToDeclarationSpecifiers(args[0].T, args[1].T)
	case _CReduceAToInitDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CInitDeclaratorListSymbol
		symbol.T, err = reducer.AToInitDeclaratorList(args[0].T)
	case _CReduceBToInitDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CInitDeclaratorListSymbol
		symbol.T, err = reducer.BToInitDeclaratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToInitDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CInitDeclaratorSymbol
		symbol.T, err = reducer.AToInitDeclarator(args[0].T)
	case _CReduceBToInitDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CInitDeclaratorSymbol
		symbol.T, err = reducer.BToInitDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStorageClassSpecifierSymbol
		symbol.T, err = reducer.AToStorageClassSpecifier(args[0].T)
	case _CReduceBToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStorageClassSpecifierSymbol
		symbol.T, err = reducer.BToStorageClassSpecifier(args[0].T)
	case _CReduceCToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStorageClassSpecifierSymbol
		symbol.T, err = reducer.CToStorageClassSpecifier(args[0].T)
	case _CReduceDToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStorageClassSpecifierSymbol
		symbol.T, err = reducer.DToStorageClassSpecifier(args[0].T)
	case _CReduceEToStorageClassSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStorageClassSpecifierSymbol
		symbol.T, err = reducer.EToStorageClassSpecifier(args[0].T)
	case _CReduceAToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.AToTypeSpecifier(args[0].T)
	case _CReduceBToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.BToTypeSpecifier(args[0].T)
	case _CReduceCToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.CToTypeSpecifier(args[0].T)
	case _CReduceDToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.DToTypeSpecifier(args[0].T)
	case _CReduceEToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.EToTypeSpecifier(args[0].T)
	case _CReduceFToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.FToTypeSpecifier(args[0].T)
	case _CReduceGToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.GToTypeSpecifier(args[0].T)
	case _CReduceHToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.HToTypeSpecifier(args[0].T)
	case _CReduceIToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.IToTypeSpecifier(args[0].T)
	case _CReduceJToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.JToTypeSpecifier(args[0].T)
	case _CReduceKToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.KToTypeSpecifier(args[0].T)
	case _CReduceLToTypeSpecifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeSpecifierSymbol
		symbol.T, err = reducer.LToTypeSpecifier(args[0].T)
	case _CReduceAToStructOrUnionSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CStructOrUnionSpecifierSymbol
		symbol.T, err = reducer.AToStructOrUnionSpecifier(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToStructOrUnionSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CStructOrUnionSpecifierSymbol
		symbol.T, err = reducer.BToStructOrUnionSpecifier(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToStructOrUnionSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CStructOrUnionSpecifierSymbol
		symbol.T, err = reducer.CToStructOrUnionSpecifier(args[0].T, args[1].T)
	case _CReduceAToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStructOrUnionSymbol
		symbol.T, err = reducer.AToStructOrUnion(args[0].T)
	case _CReduceBToStructOrUnion:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStructOrUnionSymbol
		symbol.T, err = reducer.BToStructOrUnion(args[0].T)
	case _CReduceAToStructDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStructDeclarationListSymbol
		symbol.T, err = reducer.AToStructDeclarationList(args[0].T)
	case _CReduceBToStructDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CStructDeclarationListSymbol
		symbol.T, err = reducer.BToStructDeclarationList(args[0].T, args[1].T)
	case _CReduceAToStructDeclaration:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CStructDeclarationSymbol
		symbol.T, err = reducer.AToStructDeclaration(args[0].T, args[1].T, args[2].T)
	case _CReduceAToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CSpecifierQualifierListSymbol
		symbol.T, err = reducer.AToSpecifierQualifierList(args[0].T, args[1].T)
	case _CReduceBToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CSpecifierQualifierListSymbol
		symbol.T, err = reducer.BToSpecifierQualifierList(args[0].T)
	case _CReduceCToSpecifierQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CSpecifierQualifierListSymbol
		symbol.T, err = reducer.CToSpecifierQualifierList(args[0].T, args[1].T)
	case _CReduceDToSpecifierQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CSpecifierQualifierListSymbol
		symbol.T, err = reducer.DToSpecifierQualifierList(args[0].T)
	case _CReduceAToStructDeclaratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStructDeclaratorListSymbol
		symbol.T, err = reducer.AToStructDeclaratorList(args[0].T)
	case _CReduceBToStructDeclaratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CStructDeclaratorListSymbol
		symbol.T, err = reducer.BToStructDeclaratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStructDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStructDeclaratorSymbol
		symbol.T, err = reducer.AToStructDeclarator(args[0].T)
	case _CReduceBToStructDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CStructDeclaratorSymbol
		symbol.T, err = reducer.BToStructDeclarator(args[0].T, args[1].T)
	case _CReduceCToStructDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CStructDeclaratorSymbol
		symbol.T, err = reducer.CToStructDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEnumSpecifier:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CEnumSpecifierSymbol
		symbol.T, err = reducer.AToEnumSpecifier(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceBToEnumSpecifier:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CEnumSpecifierSymbol
		symbol.T, err = reducer.BToEnumSpecifier(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceCToEnumSpecifier:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CEnumSpecifierSymbol
		symbol.T, err = reducer.CToEnumSpecifier(args[0].T, args[1].T)
	case _CReduceAToEnumeratorList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CEnumeratorListSymbol
		symbol.T, err = reducer.AToEnumeratorList(args[0].T)
	case _CReduceBToEnumeratorList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CEnumeratorListSymbol
		symbol.T, err = reducer.BToEnumeratorList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToEnumerator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CEnumeratorSymbol
		symbol.T, err = reducer.AToEnumerator(args[0].T)
	case _CReduceBToEnumerator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CEnumeratorSymbol
		symbol.T, err = reducer.BToEnumerator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeQualifierSymbol
		symbol.T, err = reducer.AToTypeQualifier(args[0].T)
	case _CReduceBToTypeQualifier:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeQualifierSymbol
		symbol.T, err = reducer.BToTypeQualifier(args[0].T)
	case _CReduceAToDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclaratorSymbol
		symbol.T, err = reducer.AToDeclarator(args[0].T, args[1].T)
	case _CReduceBToDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDeclaratorSymbol
		symbol.T, err = reducer.BToDeclarator(args[0].T)
	case _CReduceAToDirectDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.AToDirectDeclarator(args[0].T)
	case _CReduceBToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.BToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceCToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.CToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceDToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.DToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceEToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.EToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceFToDirectDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.FToDirectDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceGToDirectDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectDeclaratorSymbol
		symbol.T, err = reducer.GToDirectDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceAToPointer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CPointerSymbol
		symbol.T, err = reducer.AToPointer(args[0].T)
	case _CReduceBToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CPointerSymbol
		symbol.T, err = reducer.BToPointer(args[0].T, args[1].T)
	case _CReduceCToPointer:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CPointerSymbol
		symbol.T, err = reducer.CToPointer(args[0].T, args[1].T)
	case _CReduceDToPointer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CPointerSymbol
		symbol.T, err = reducer.DToPointer(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeQualifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeQualifierListSymbol
		symbol.T, err = reducer.AToTypeQualifierList(args[0].T)
	case _CReduceBToTypeQualifierList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CTypeQualifierListSymbol
		symbol.T, err = reducer.BToTypeQualifierList(args[0].T, args[1].T)
	case _CReduceAToParameterTypeList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CParameterTypeListSymbol
		symbol.T, err = reducer.AToParameterTypeList(args[0].T)
	case _CReduceBToParameterTypeList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CParameterTypeListSymbol
		symbol.T, err = reducer.BToParameterTypeList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToParameterList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CParameterListSymbol
		symbol.T, err = reducer.AToParameterList(args[0].T)
	case _CReduceBToParameterList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CParameterListSymbol
		symbol.T, err = reducer.BToParameterList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CParameterDeclarationSymbol
		symbol.T, err = reducer.AToParameterDeclaration(args[0].T, args[1].T)
	case _CReduceBToParameterDeclaration:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CParameterDeclarationSymbol
		symbol.T, err = reducer.BToParameterDeclaration(args[0].T, args[1].T)
	case _CReduceCToParameterDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CParameterDeclarationSymbol
		symbol.T, err = reducer.CToParameterDeclaration(args[0].T)
	case _CReduceAToIdentifierList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CIdentifierListSymbol
		symbol.T, err = reducer.AToIdentifierList(args[0].T)
	case _CReduceBToIdentifierList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CIdentifierListSymbol
		symbol.T, err = reducer.BToIdentifierList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTypeName:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTypeNameSymbol
		symbol.T, err = reducer.AToTypeName(args[0].T)
	case _CReduceBToTypeName:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CTypeNameSymbol
		symbol.T, err = reducer.BToTypeName(args[0].T, args[1].T)
	case _CReduceAToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAbstractDeclaratorSymbol
		symbol.T, err = reducer.AToAbstractDeclarator(args[0].T)
	case _CReduceBToAbstractDeclarator:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CAbstractDeclaratorSymbol
		symbol.T, err = reducer.BToAbstractDeclarator(args[0].T)
	case _CReduceCToAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CAbstractDeclaratorSymbol
		symbol.T, err = reducer.CToAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceAToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.AToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceBToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.BToDirectAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceCToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.CToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceDToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.DToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceEToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.EToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceFToDirectAbstractDeclarator:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.FToDirectAbstractDeclarator(args[0].T, args[1].T)
	case _CReduceGToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.GToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceHToDirectAbstractDeclarator:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.HToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T)
	case _CReduceIToDirectAbstractDeclarator:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CDirectAbstractDeclaratorSymbol
		symbol.T, err = reducer.IToDirectAbstractDeclarator(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToInitializer:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CInitializerSymbol
		symbol.T, err = reducer.AToInitializer(args[0].T)
	case _CReduceBToInitializer:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CInitializerSymbol
		symbol.T, err = reducer.BToInitializer(args[0].T, args[1].T, args[2].T)
	case _CReduceCToInitializer:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CInitializerSymbol
		symbol.T, err = reducer.CToInitializer(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToInitializerList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CInitializerListSymbol
		symbol.T, err = reducer.AToInitializerList(args[0].T)
	case _CReduceBToInitializerList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CInitializerListSymbol
		symbol.T, err = reducer.BToInitializerList(args[0].T, args[1].T, args[2].T)
	case _CReduceAToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.AToStatement(args[0].T)
	case _CReduceBToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.BToStatement(args[0].T)
	case _CReduceCToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.CToStatement(args[0].T)
	case _CReduceDToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.DToStatement(args[0].T)
	case _CReduceEToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.EToStatement(args[0].T)
	case _CReduceFToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementSymbol
		symbol.T, err = reducer.FToStatement(args[0].T)
	case _CReduceAToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CLabeledStatementSymbol
		symbol.T, err = reducer.AToLabeledStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceBToLabeledStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CLabeledStatementSymbol
		symbol.T, err = reducer.BToLabeledStatement(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceCToLabeledStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CLabeledStatementSymbol
		symbol.T, err = reducer.CToLabeledStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceAToCompoundStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CCompoundStatementSymbol
		symbol.T, err = reducer.AToCompoundStatement(args[0].T, args[1].T)
	case _CReduceBToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CCompoundStatementSymbol
		symbol.T, err = reducer.BToCompoundStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceCToCompoundStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CCompoundStatementSymbol
		symbol.T, err = reducer.CToCompoundStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceDToCompoundStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CCompoundStatementSymbol
		symbol.T, err = reducer.DToCompoundStatement(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceAToDeclarationList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CDeclarationListSymbol
		symbol.T, err = reducer.AToDeclarationList(args[0].T)
	case _CReduceBToDeclarationList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CDeclarationListSymbol
		symbol.T, err = reducer.BToDeclarationList(args[0].T, args[1].T)
	case _CReduceAToStatementList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CStatementListSymbol
		symbol.T, err = reducer.AToStatementList(args[0].T)
	case _CReduceBToStatementList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CStatementListSymbol
		symbol.T, err = reducer.BToStatementList(args[0].T, args[1].T)
	case _CReduceAToExpressionStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CExpressionStatementSymbol
		symbol.T, err = reducer.AToExpressionStatement(args[0].T)
	case _CReduceBToExpressionStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CExpressionStatementSymbol
		symbol.T, err = reducer.BToExpressionStatement(args[0].T, args[1].T)
	case _CReduceAToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CSelectionStatementSymbol
		symbol.T, err = reducer.AToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToSelectionStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = _CSelectionStatementSymbol
		symbol.T, err = reducer.BToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceCToSelectionStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CSelectionStatementSymbol
		symbol.T, err = reducer.CToSelectionStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceAToIterationStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = _CIterationStatementSymbol
		symbol.T, err = reducer.AToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T)
	case _CReduceBToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = _CIterationStatementSymbol
		symbol.T, err = reducer.BToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceCToIterationStatement:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = _CIterationStatementSymbol
		symbol.T, err = reducer.CToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T)
	case _CReduceDToIterationStatement:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = _CIterationStatementSymbol
		symbol.T, err = reducer.DToIterationStatement(args[0].T, args[1].T, args[2].T, args[3].T, args[4].T, args[5].T, args[6].T)
	case _CReduceAToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CJumpStatementSymbol
		symbol.T, err = reducer.AToJumpStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceBToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CJumpStatementSymbol
		symbol.T, err = reducer.BToJumpStatement(args[0].T, args[1].T)
	case _CReduceCToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CJumpStatementSymbol
		symbol.T, err = reducer.CToJumpStatement(args[0].T, args[1].T)
	case _CReduceDToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CJumpStatementSymbol
		symbol.T, err = reducer.DToJumpStatement(args[0].T, args[1].T)
	case _CReduceEToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CJumpStatementSymbol
		symbol.T, err = reducer.EToJumpStatement(args[0].T, args[1].T, args[2].T)
	case _CReduceAToTranslationUnit:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CTranslationUnitSymbol
		symbol.T, err = reducer.AToTranslationUnit(args[0].T)
	case _CReduceBToTranslationUnit:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CTranslationUnitSymbol
		symbol.T, err = reducer.BToTranslationUnit(args[0].T, args[1].T)
	case _CReduceAToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CExternalDeclarationSymbol
		symbol.T, err = reducer.AToExternalDeclaration(args[0].T)
	case _CReduceBToExternalDeclaration:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = _CExternalDeclarationSymbol
		symbol.T, err = reducer.BToExternalDeclaration(args[0].T)
	case _CReduceAToFunctionDefinition:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = _CFunctionDefinitionSymbol
		symbol.T, err = reducer.AToFunctionDefinition(args[0].T, args[1].T, args[2].T, args[3].T)
	case _CReduceBToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CFunctionDefinitionSymbol
		symbol.T, err = reducer.BToFunctionDefinition(args[0].T, args[1].T, args[2].T)
	case _CReduceCToFunctionDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = _CFunctionDefinitionSymbol
		symbol.T, err = reducer.CToFunctionDefinition(args[0].T, args[1].T, args[2].T)
	case _CReduceDToFunctionDefinition:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = _CFunctionDefinitionSymbol
		symbol.T, err = reducer.DToFunctionDefinition(args[0].T, args[1].T)
	default:
		panic("Unknown reduce type: " + act.ReduceType)
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _CActionTableKey struct {
	_CStateId
	CSymbolId
}

type _CActionTableType map[_CActionTableKey]_CAction

func (table _CActionTableType) Get(stateId _CStateId, symbol CSymbolId) (_CAction, bool) {
	action, ok := table[_CActionTableKey{stateId, symbol}]
	if ok {
		return action, ok
	}
	action, ok = table[_CActionTableKey{stateId, _CWildcardMarker}]
	return action, ok
}

var _CActionTable = _CActionTableType{
	{_CState1, _CEndMarker}:                        {_CAcceptAction, "", ""},
	{_CState0, CIdentifierSymbol}:                  {_CShiftAction, _CState9, ""},
	{_CState0, CTypeNameSymbol}:                    {_CShiftAction, _CState20, ""},
	{_CState0, CTypedefSymbol}:                     {_CShiftAction, _CState19, ""},
	{_CState0, CExternSymbol}:                      {_CShiftAction, _CState7, ""},
	{_CState0, CStaticSymbol}:                      {_CShiftAction, _CState17, ""},
	{_CState0, CAutoSymbol}:                        {_CShiftAction, _CState2, ""},
	{_CState0, CRegisterSymbol}:                    {_CShiftAction, _CState14, ""},
	{_CState0, CCharSymbol}:                        {_CShiftAction, _CState3, ""},
	{_CState0, CShortSymbol}:                       {_CShiftAction, _CState15, ""},
	{_CState0, CIntSymbol}:                         {_CShiftAction, _CState10, ""},
	{_CState0, CLongSymbol}:                        {_CShiftAction, _CState11, ""},
	{_CState0, CSignedSymbol}:                      {_CShiftAction, _CState16, ""},
	{_CState0, CUnsignedSymbol}:                    {_CShiftAction, _CState22, ""},
	{_CState0, CFloatSymbol}:                       {_CShiftAction, _CState8, ""},
	{_CState0, CDoubleSymbol}:                      {_CShiftAction, _CState5, ""},
	{_CState0, CConstSymbol}:                       {_CShiftAction, _CState4, ""},
	{_CState0, CVolatileSymbol}:                    {_CShiftAction, _CState24, ""},
	{_CState0, CVoidSymbol}:                        {_CShiftAction, _CState23, ""},
	{_CState0, CStructSymbol}:                      {_CShiftAction, _CState18, ""},
	{_CState0, CUnionSymbol}:                       {_CShiftAction, _CState21, ""},
	{_CState0, CEnumSymbol}:                        {_CShiftAction, _CState6, ""},
	{_CState0, CLparamSymbol}:                      {_CShiftAction, _CState12, ""},
	{_CState0, CMulSymbol}:                         {_CShiftAction, _CState13, ""},
	{_CState0, _CDeclarationSymbol}:                {_CShiftAction, _CState25, ""},
	{_CState0, _CDeclarationSpecifiersSymbol}:      {_CShiftAction, _CState26, ""},
	{_CState0, _CStorageClassSpecifierSymbol}:      {_CShiftAction, _CState33, ""},
	{_CState0, _CTypeSpecifierSymbol}:              {_CShiftAction, _CState37, ""},
	{_CState0, _CStructOrUnionSpecifierSymbol}:     {_CShiftAction, _CState35, ""},
	{_CState0, _CStructOrUnionSymbol}:              {_CShiftAction, _CState34, ""},
	{_CState0, _CEnumSpecifierSymbol}:              {_CShiftAction, _CState29, ""},
	{_CState0, _CTypeQualifierSymbol}:              {_CShiftAction, _CState36, ""},
	{_CState0, _CDeclaratorSymbol}:                 {_CShiftAction, _CState27, ""},
	{_CState0, _CDirectDeclaratorSymbol}:           {_CShiftAction, _CState28, ""},
	{_CState0, _CPointerSymbol}:                    {_CShiftAction, _CState32, ""},
	{_CState0, _CTranslationUnitSymbol}:            {_CShiftAction, _CState1, ""},
	{_CState0, _CExternalDeclarationSymbol}:        {_CShiftAction, _CState30, ""},
	{_CState0, _CFunctionDefinitionSymbol}:         {_CShiftAction, _CState31, ""},
	{_CState1, CIdentifierSymbol}:                  {_CShiftAction, _CState9, ""},
	{_CState1, CTypeNameSymbol}:                    {_CShiftAction, _CState20, ""},
	{_CState1, CTypedefSymbol}:                     {_CShiftAction, _CState19, ""},
	{_CState1, CExternSymbol}:                      {_CShiftAction, _CState7, ""},
	{_CState1, CStaticSymbol}:                      {_CShiftAction, _CState17, ""},
	{_CState1, CAutoSymbol}:                        {_CShiftAction, _CState2, ""},
	{_CState1, CRegisterSymbol}:                    {_CShiftAction, _CState14, ""},
	{_CState1, CCharSymbol}:                        {_CShiftAction, _CState3, ""},
	{_CState1, CShortSymbol}:                       {_CShiftAction, _CState15, ""},
	{_CState1, CIntSymbol}:                         {_CShiftAction, _CState10, ""},
	{_CState1, CLongSymbol}:                        {_CShiftAction, _CState11, ""},
	{_CState1, CSignedSymbol}:                      {_CShiftAction, _CState16, ""},
	{_CState1, CUnsignedSymbol}:                    {_CShiftAction, _CState22, ""},
	{_CState1, CFloatSymbol}:                       {_CShiftAction, _CState8, ""},
	{_CState1, CDoubleSymbol}:                      {_CShiftAction, _CState5, ""},
	{_CState1, CConstSymbol}:                       {_CShiftAction, _CState4, ""},
	{_CState1, CVolatileSymbol}:                    {_CShiftAction, _CState24, ""},
	{_CState1, CVoidSymbol}:                        {_CShiftAction, _CState23, ""},
	{_CState1, CStructSymbol}:                      {_CShiftAction, _CState18, ""},
	{_CState1, CUnionSymbol}:                       {_CShiftAction, _CState21, ""},
	{_CState1, CEnumSymbol}:                        {_CShiftAction, _CState6, ""},
	{_CState1, CLparamSymbol}:                      {_CShiftAction, _CState12, ""},
	{_CState1, CMulSymbol}:                         {_CShiftAction, _CState13, ""},
	{_CState1, _CDeclarationSymbol}:                {_CShiftAction, _CState25, ""},
	{_CState1, _CDeclarationSpecifiersSymbol}:      {_CShiftAction, _CState26, ""},
	{_CState1, _CStorageClassSpecifierSymbol}:      {_CShiftAction, _CState33, ""},
	{_CState1, _CTypeSpecifierSymbol}:              {_CShiftAction, _CState37, ""},
	{_CState1, _CStructOrUnionSpecifierSymbol}:     {_CShiftAction, _CState35, ""},
	{_CState1, _CStructOrUnionSymbol}:              {_CShiftAction, _CState34, ""},
	{_CState1, _CEnumSpecifierSymbol}:              {_CShiftAction, _CState29, ""},
	{_CState1, _CTypeQualifierSymbol}:              {_CShiftAction, _CState36, ""},
	{_CState1, _CDeclaratorSymbol}:                 {_CShiftAction, _CState27, ""},
	{_CState1, _CDirectDeclaratorSymbol}:           {_CShiftAction, _CState28, ""},
	{_CState1, _CPointerSymbol}:                    {_CShiftAction, _CState32, ""},
	{_CState1, _CExternalDeclarationSymbol}:        {_CShiftAction, _CState38, ""},
	{_CState1, _CFunctionDefinitionSymbol}:         {_CShiftAction, _CState31, ""},
	{_CState6, CIdentifierSymbol}:                  {_CShiftAction, _CState39, ""},
	{_CState6, CLcurlSymbol}:                       {_CShiftAction, _CState40, ""},
	{_CState12, CIdentifierSymbol}:                 {_CShiftAction, _CState9, ""},
	{_CState12, CLparamSymbol}:                     {_CShiftAction, _CState12, ""},
	{_CState12, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState12, _CDeclaratorSymbol}:                {_CShiftAction, _CState41, ""},
	{_CState12, _CDirectDeclaratorSymbol}:          {_CShiftAction, _CState28, ""},
	{_CState12, _CPointerSymbol}:                   {_CShiftAction, _CState32, ""},
	{_CState13, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState13, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState13, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState13, _CTypeQualifierSymbol}:             {_CShiftAction, _CState43, ""},
	{_CState13, _CPointerSymbol}:                   {_CShiftAction, _CState42, ""},
	{_CState13, _CTypeQualifierListSymbol}:         {_CShiftAction, _CState44, ""},
	{_CState26, CIdentifierSymbol}:                 {_CShiftAction, _CState9, ""},
	{_CState26, CLparamSymbol}:                     {_CShiftAction, _CState12, ""},
	{_CState26, CSemicolonSymbol}:                  {_CShiftAction, _CState45, ""},
	{_CState26, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState26, _CInitDeclaratorListSymbol}:        {_CShiftAction, _CState48, ""},
	{_CState26, _CInitDeclaratorSymbol}:            {_CShiftAction, _CState47, ""},
	{_CState26, _CDeclaratorSymbol}:                {_CShiftAction, _CState46, ""},
	{_CState26, _CDirectDeclaratorSymbol}:          {_CShiftAction, _CState28, ""},
	{_CState26, _CPointerSymbol}:                   {_CShiftAction, _CState32, ""},
	{_CState27, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState27, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState27, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState27, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState27, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState27, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState27, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState27, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState27, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState27, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState27, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState27, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState27, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState27, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState27, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState27, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState27, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState27, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState27, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState27, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState27, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState27, _CDeclarationSymbol}:               {_CShiftAction, _CState51, ""},
	{_CState27, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState53, ""},
	{_CState27, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState27, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState27, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState27, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState27, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState27, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState27, _CCompoundStatementSymbol}:         {_CShiftAction, _CState50, ""},
	{_CState27, _CDeclarationListSymbol}:           {_CShiftAction, _CState52, ""},
	{_CState28, CLparamSymbol}:                     {_CShiftAction, _CState55, ""},
	{_CState28, CLbraceSymbol}:                     {_CShiftAction, _CState54, ""},
	{_CState32, CIdentifierSymbol}:                 {_CShiftAction, _CState9, ""},
	{_CState32, CLparamSymbol}:                     {_CShiftAction, _CState12, ""},
	{_CState32, _CDirectDeclaratorSymbol}:          {_CShiftAction, _CState56, ""},
	{_CState33, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState33, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState33, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState33, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState33, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState33, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState33, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState33, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState33, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState33, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState33, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState33, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState33, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState33, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState33, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState33, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState33, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState33, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState33, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState33, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState33, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState57, ""},
	{_CState33, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState33, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState33, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState33, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState33, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState33, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState34, CIdentifierSymbol}:                 {_CShiftAction, _CState58, ""},
	{_CState34, CLcurlSymbol}:                      {_CShiftAction, _CState59, ""},
	{_CState36, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState36, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState36, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState36, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState36, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState36, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState36, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState36, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState36, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState36, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState36, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState36, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState36, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState36, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState36, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState36, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState36, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState36, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState36, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState36, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState36, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState60, ""},
	{_CState36, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState36, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState36, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState36, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState36, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState36, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState37, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState37, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState37, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState37, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState37, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState37, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState37, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState37, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState37, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState37, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState37, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState37, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState37, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState37, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState37, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState37, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState37, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState37, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState37, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState37, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState37, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState61, ""},
	{_CState37, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState37, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState37, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState37, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState37, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState37, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState39, CLcurlSymbol}:                      {_CShiftAction, _CState62, ""},
	{_CState40, CIdentifierSymbol}:                 {_CShiftAction, _CState63, ""},
	{_CState40, _CEnumeratorListSymbol}:            {_CShiftAction, _CState65, ""},
	{_CState40, _CEnumeratorSymbol}:                {_CShiftAction, _CState64, ""},
	{_CState41, CRparamSymbol}:                     {_CShiftAction, _CState66, ""},
	{_CState44, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState44, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState44, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState44, _CTypeQualifierSymbol}:             {_CShiftAction, _CState68, ""},
	{_CState44, _CPointerSymbol}:                   {_CShiftAction, _CState67, ""},
	{_CState46, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState46, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState46, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState46, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState46, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState46, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState46, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState46, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState46, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState46, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState46, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState46, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState46, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState46, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState46, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState46, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState46, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState46, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState46, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState46, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState46, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState46, CEqSymbol}:                         {_CShiftAction, _CState69, ""},
	{_CState46, _CDeclarationSymbol}:               {_CShiftAction, _CState51, ""},
	{_CState46, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState53, ""},
	{_CState46, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState46, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState46, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState46, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState46, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState46, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState46, _CCompoundStatementSymbol}:         {_CShiftAction, _CState70, ""},
	{_CState46, _CDeclarationListSymbol}:           {_CShiftAction, _CState71, ""},
	{_CState48, CSemicolonSymbol}:                  {_CShiftAction, _CState73, ""},
	{_CState48, CCommaSymbol}:                      {_CShiftAction, _CState72, ""},
	{_CState49, CIdentifierSymbol}:                 {_CShiftAction, _CState85, ""},
	{_CState49, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState49, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState49, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState49, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState49, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState49, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState49, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState49, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState49, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState49, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState49, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState49, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState49, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState49, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState49, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState49, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState49, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState49, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState49, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState49, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState49, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState49, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState49, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState49, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState49, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState49, CCaseSymbol}:                       {_CShiftAction, _CState76, ""},
	{_CState49, CDefaultSymbol}:                    {_CShiftAction, _CState80, ""},
	{_CState49, CIfSymbol}:                         {_CShiftAction, _CState86, ""},
	{_CState49, CSwitchSymbol}:                     {_CShiftAction, _CState97, ""},
	{_CState49, CWhileSymbol}:                      {_CShiftAction, _CState99, ""},
	{_CState49, CDoSymbol}:                         {_CShiftAction, _CState81, ""},
	{_CState49, CForSymbol}:                        {_CShiftAction, _CState83, ""},
	{_CState49, CGotoSymbol}:                       {_CShiftAction, _CState84, ""},
	{_CState49, CContinueSymbol}:                   {_CShiftAction, _CState78, ""},
	{_CState49, CBreakSymbol}:                      {_CShiftAction, _CState75, ""},
	{_CState49, CReturnSymbol}:                     {_CShiftAction, _CState93, ""},
	{_CState49, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState49, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState49, CRcurlSymbol}:                      {_CShiftAction, _CState92, ""},
	{_CState49, CSemicolonSymbol}:                  {_CShiftAction, _CState94, ""},
	{_CState49, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState49, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState49, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState49, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState49, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState49, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState49, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState49, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState49, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState125, ""},
	{_CState49, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState49, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState49, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState49, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState49, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState49, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState49, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState49, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState49, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState49, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState49, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState49, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState49, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState105, ""},
	{_CState49, _CAssignmentExpressionSymbol}:      {_CShiftAction, _CState102, ""},
	{_CState49, _CExpressionSymbol}:                {_CShiftAction, _CState109, ""},
	{_CState49, _CDeclarationSymbol}:               {_CShiftAction, _CState51, ""},
	{_CState49, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState53, ""},
	{_CState49, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState49, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState49, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState49, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState49, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState49, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState49, _CStatementSymbol}:                 {_CShiftAction, _CState123, ""},
	{_CState49, _CLabeledStatementSymbol}:          {_CShiftAction, _CState114, ""},
	{_CState49, _CCompoundStatementSymbol}:         {_CShiftAction, _CState104, ""},
	{_CState49, _CDeclarationListSymbol}:           {_CShiftAction, _CState106, ""},
	{_CState49, _CStatementListSymbol}:             {_CShiftAction, _CState124, ""},
	{_CState49, _CExpressionStatementSymbol}:       {_CShiftAction, _CState110, ""},
	{_CState49, _CSelectionStatementSymbol}:        {_CShiftAction, _CState121, ""},
	{_CState49, _CIterationStatementSymbol}:        {_CShiftAction, _CState112, ""},
	{_CState49, _CJumpStatementSymbol}:             {_CShiftAction, _CState113, ""},
	{_CState52, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState52, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState52, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState52, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState52, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState52, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState52, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState52, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState52, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState52, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState52, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState52, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState52, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState52, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState52, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState52, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState52, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState52, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState52, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState52, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState52, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState52, _CDeclarationSymbol}:               {_CShiftAction, _CState128, ""},
	{_CState52, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState53, ""},
	{_CState52, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState52, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState52, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState52, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState52, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState52, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState52, _CCompoundStatementSymbol}:         {_CShiftAction, _CState127, ""},
	{_CState53, CIdentifierSymbol}:                 {_CShiftAction, _CState9, ""},
	{_CState53, CLparamSymbol}:                     {_CShiftAction, _CState12, ""},
	{_CState53, CSemicolonSymbol}:                  {_CShiftAction, _CState45, ""},
	{_CState53, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState53, _CInitDeclaratorListSymbol}:        {_CShiftAction, _CState48, ""},
	{_CState53, _CInitDeclaratorSymbol}:            {_CShiftAction, _CState47, ""},
	{_CState53, _CDeclaratorSymbol}:                {_CShiftAction, _CState129, ""},
	{_CState53, _CDirectDeclaratorSymbol}:          {_CShiftAction, _CState28, ""},
	{_CState53, _CPointerSymbol}:                   {_CShiftAction, _CState32, ""},
	{_CState54, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState54, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState54, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState54, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState54, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState54, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState54, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState54, CRbraceSymbol}:                     {_CShiftAction, _CState131, ""},
	{_CState54, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState54, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState54, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState54, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState54, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState54, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState54, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState54, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState54, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState134, ""},
	{_CState54, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState54, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState54, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState54, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState54, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState54, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState54, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState54, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState54, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState54, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState54, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState54, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState54, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState132, ""},
	{_CState54, _CConstantExpressionSymbol}:        {_CShiftAction, _CState133, ""},
	{_CState55, CIdentifierSymbol}:                 {_CShiftAction, _CState135, ""},
	{_CState55, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState55, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState55, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState55, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState55, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState55, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState55, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState55, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState55, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState55, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState55, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState55, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState55, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState55, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState55, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState55, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState55, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState55, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState55, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState55, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState55, CRparamSymbol}:                     {_CShiftAction, _CState136, ""},
	{_CState55, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState137, ""},
	{_CState55, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState55, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState55, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState55, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState55, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState55, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState55, _CParameterTypeListSymbol}:         {_CShiftAction, _CState141, ""},
	{_CState55, _CParameterListSymbol}:             {_CShiftAction, _CState140, ""},
	{_CState55, _CParameterDeclarationSymbol}:      {_CShiftAction, _CState139, ""},
	{_CState55, _CIdentifierListSymbol}:            {_CShiftAction, _CState138, ""},
	{_CState56, CLparamSymbol}:                     {_CShiftAction, _CState55, ""},
	{_CState56, CLbraceSymbol}:                     {_CShiftAction, _CState54, ""},
	{_CState58, CLcurlSymbol}:                      {_CShiftAction, _CState142, ""},
	{_CState59, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState59, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState59, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState59, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState59, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState59, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState59, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState59, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState59, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState59, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState59, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState59, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState59, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState59, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState59, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState59, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState147, ""},
	{_CState59, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState59, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState59, _CStructDeclarationListSymbol}:     {_CShiftAction, _CState145, ""},
	{_CState59, _CStructDeclarationSymbol}:         {_CShiftAction, _CState144, ""},
	{_CState59, _CSpecifierQualifierListSymbol}:    {_CShiftAction, _CState143, ""},
	{_CState59, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState59, _CTypeQualifierSymbol}:             {_CShiftAction, _CState146, ""},
	{_CState62, CIdentifierSymbol}:                 {_CShiftAction, _CState63, ""},
	{_CState62, _CEnumeratorListSymbol}:            {_CShiftAction, _CState148, ""},
	{_CState62, _CEnumeratorSymbol}:                {_CShiftAction, _CState64, ""},
	{_CState63, CEqSymbol}:                         {_CShiftAction, _CState149, ""},
	{_CState65, CRcurlSymbol}:                      {_CShiftAction, _CState151, ""},
	{_CState65, CCommaSymbol}:                      {_CShiftAction, _CState150, ""},
	{_CState69, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState69, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState69, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState69, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState69, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState69, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState69, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState69, CLcurlSymbol}:                      {_CShiftAction, _CState152, ""},
	{_CState69, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState69, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState69, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState69, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState69, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState69, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState69, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState69, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState69, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState125, ""},
	{_CState69, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState69, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState69, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState69, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState69, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState69, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState69, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState69, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState69, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState69, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState69, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState69, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState69, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState105, ""},
	{_CState69, _CAssignmentExpressionSymbol}:      {_CShiftAction, _CState153, ""},
	{_CState69, _CInitializerSymbol}:               {_CShiftAction, _CState154, ""},
	{_CState71, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState71, CTypedefSymbol}:                    {_CShiftAction, _CState19, ""},
	{_CState71, CExternSymbol}:                     {_CShiftAction, _CState7, ""},
	{_CState71, CStaticSymbol}:                     {_CShiftAction, _CState17, ""},
	{_CState71, CAutoSymbol}:                       {_CShiftAction, _CState2, ""},
	{_CState71, CRegisterSymbol}:                   {_CShiftAction, _CState14, ""},
	{_CState71, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState71, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState71, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState71, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState71, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState71, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState71, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState71, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState71, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState71, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState71, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState71, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState71, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState71, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState71, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState71, _CDeclarationSymbol}:               {_CShiftAction, _CState128, ""},
	{_CState71, _CDeclarationSpecifiersSymbol}:     {_CShiftAction, _CState53, ""},
	{_CState71, _CStorageClassSpecifierSymbol}:     {_CShiftAction, _CState33, ""},
	{_CState71, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState37, ""},
	{_CState71, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState71, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState71, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState71, _CTypeQualifierSymbol}:             {_CShiftAction, _CState36, ""},
	{_CState71, _CCompoundStatementSymbol}:         {_CShiftAction, _CState155, ""},
	{_CState72, CIdentifierSymbol}:                 {_CShiftAction, _CState9, ""},
	{_CState72, CLparamSymbol}:                     {_CShiftAction, _CState12, ""},
	{_CState72, CMulSymbol}:                        {_CShiftAction, _CState13, ""},
	{_CState72, _CInitDeclaratorSymbol}:            {_CShiftAction, _CState156, ""},
	{_CState72, _CDeclaratorSymbol}:                {_CShiftAction, _CState129, ""},
	{_CState72, _CDirectDeclaratorSymbol}:          {_CShiftAction, _CState28, ""},
	{_CState72, _CPointerSymbol}:                   {_CShiftAction, _CState32, ""},
	{_CState75, CSemicolonSymbol}:                  {_CShiftAction, _CState157, ""},
	{_CState76, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState76, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState76, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState76, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState76, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState76, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState76, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState76, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState76, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState76, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState76, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState76, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState76, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState76, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState76, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState76, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState134, ""},
	{_CState76, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState76, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState76, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState76, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState76, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState76, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState76, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState76, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState76, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState76, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState76, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState76, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState76, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState132, ""},
	{_CState76, _CConstantExpressionSymbol}:        {_CShiftAction, _CState158, ""},
	{_CState78, CSemicolonSymbol}:                  {_CShiftAction, _CState159, ""},
	{_CState79, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState79, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState79, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState79, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState79, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState79, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState79, CLparamSymbol}:                     {_CShiftAction, _CState160, ""},
	{_CState79, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState79, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState79, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState79, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState79, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState79, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState79, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState79, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState79, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState161, ""},
	{_CState79, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState80, CColonSymbol}:                      {_CShiftAction, _CState162, ""},
	{_CState81, CIdentifierSymbol}:                 {_CShiftAction, _CState85, ""},
	{_CState81, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState81, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState81, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState81, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState81, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState81, CCaseSymbol}:                       {_CShiftAction, _CState76, ""},
	{_CState81, CDefaultSymbol}:                    {_CShiftAction, _CState80, ""},
	{_CState81, CIfSymbol}:                         {_CShiftAction, _CState86, ""},
	{_CState81, CSwitchSymbol}:                     {_CShiftAction, _CState97, ""},
	{_CState81, CWhileSymbol}:                      {_CShiftAction, _CState99, ""},
	{_CState81, CDoSymbol}:                         {_CShiftAction, _CState81, ""},
	{_CState81, CForSymbol}:                        {_CShiftAction, _CState83, ""},
	{_CState81, CGotoSymbol}:                       {_CShiftAction, _CState84, ""},
	{_CState81, CContinueSymbol}:                   {_CShiftAction, _CState78, ""},
	{_CState81, CBreakSymbol}:                      {_CShiftAction, _CState75, ""},
	{_CState81, CReturnSymbol}:                     {_CShiftAction, _CState93, ""},
	{_CState81, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState81, CLcurlSymbol}:                      {_CShiftAction, _CState49, ""},
	{_CState81, CSemicolonSymbol}:                  {_CShiftAction, _CState94, ""},
	{_CState81, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState81, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState81, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState81, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState81, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState81, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState81, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState81, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState81, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState125, ""},
	{_CState81, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState81, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState81, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState81, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState81, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState81, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState81, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState81, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState81, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState81, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState81, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState81, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState81, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState105, ""},
	{_CState81, _CAssignmentExpressionSymbol}:      {_CShiftAction, _CState102, ""},
	{_CState81, _CExpressionSymbol}:                {_CShiftAction, _CState109, ""},
	{_CState81, _CStatementSymbol}:                 {_CShiftAction, _CState163, ""},
	{_CState81, _CLabeledStatementSymbol}:          {_CShiftAction, _CState114, ""},
	{_CState81, _CCompoundStatementSymbol}:         {_CShiftAction, _CState104, ""},
	{_CState81, _CExpressionStatementSymbol}:       {_CShiftAction, _CState110, ""},
	{_CState81, _CSelectionStatementSymbol}:        {_CShiftAction, _CState121, ""},
	{_CState81, _CIterationStatementSymbol}:        {_CShiftAction, _CState112, ""},
	{_CState81, _CJumpStatementSymbol}:             {_CShiftAction, _CState113, ""},
	{_CState83, CLparamSymbol}:                     {_CShiftAction, _CState164, ""},
	{_CState84, CIdentifierSymbol}:                 {_CShiftAction, _CState165, ""},
	{_CState85, CColonSymbol}:                      {_CShiftAction, _CState166, ""},
	{_CState86, CLparamSymbol}:                     {_CShiftAction, _CState167, ""},
	{_CState87, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState87, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState87, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState87, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState87, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState87, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState87, CLparamSymbol}:                     {_CShiftAction, _CState160, ""},
	{_CState87, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState87, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState87, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState87, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState87, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState87, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState87, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState87, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState87, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState168, ""},
	{_CState87, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState88, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState88, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState88, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState88, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState88, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState88, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState88, CTypeNameSymbol}:                   {_CShiftAction, _CState20, ""},
	{_CState88, CCharSymbol}:                       {_CShiftAction, _CState3, ""},
	{_CState88, CShortSymbol}:                      {_CShiftAction, _CState15, ""},
	{_CState88, CIntSymbol}:                        {_CShiftAction, _CState10, ""},
	{_CState88, CLongSymbol}:                       {_CShiftAction, _CState11, ""},
	{_CState88, CSignedSymbol}:                     {_CShiftAction, _CState16, ""},
	{_CState88, CUnsignedSymbol}:                   {_CShiftAction, _CState22, ""},
	{_CState88, CFloatSymbol}:                      {_CShiftAction, _CState8, ""},
	{_CState88, CDoubleSymbol}:                     {_CShiftAction, _CState5, ""},
	{_CState88, CConstSymbol}:                      {_CShiftAction, _CState4, ""},
	{_CState88, CVolatileSymbol}:                   {_CShiftAction, _CState24, ""},
	{_CState88, CVoidSymbol}:                       {_CShiftAction, _CState23, ""},
	{_CState88, CStructSymbol}:                     {_CShiftAction, _CState18, ""},
	{_CState88, CUnionSymbol}:                      {_CShiftAction, _CState21, ""},
	{_CState88, CEnumSymbol}:                       {_CShiftAction, _CState6, ""},
	{_CState88, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState88, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState88, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState88, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState88, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState88, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState88, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState88, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState88, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState88, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState125, ""},
	{_CState88, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState88, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState88, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState88, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState88, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState88, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState88, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState88, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState88, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState88, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState88, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState88, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState88, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState105, ""},
	{_CState88, _CAssignmentExpressionSymbol}:      {_CShiftAction, _CState102, ""},
	{_CState88, _CExpressionSymbol}:                {_CShiftAction, _CState169, ""},
	{_CState88, _CTypeSpecifierSymbol}:             {_CShiftAction, _CState147, ""},
	{_CState88, _CStructOrUnionSpecifierSymbol}:    {_CShiftAction, _CState35, ""},
	{_CState88, _CStructOrUnionSymbol}:             {_CShiftAction, _CState34, ""},
	{_CState88, _CSpecifierQualifierListSymbol}:    {_CShiftAction, _CState170, ""},
	{_CState88, _CEnumSpecifierSymbol}:             {_CShiftAction, _CState29, ""},
	{_CState88, _CTypeQualifierSymbol}:             {_CShiftAction, _CState146, ""},
	{_CState88, _CTypeNameSymbol}:                  {_CShiftAction, _CState171, ""},
	{_CState93, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState93, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState93, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState93, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState93, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState93, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState93, CLparamSymbol}:                     {_CShiftAction, _CState88, ""},
	{_CState93, CSemicolonSymbol}:                  {_CShiftAction, _CState172, ""},
	{_CState93, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState93, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState93, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState93, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState93, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState93, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState93, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState93, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState93, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState125, ""},
	{_CState93, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState93, _CCastExpressionSymbol}:            {_CShiftAction, _CState103, ""},
	{_CState93, _CMultiplicativeExpressionSymbol}:  {_CShiftAction, _CState117, ""},
	{_CState93, _CAdditiveExpressionSymbol}:        {_CShiftAction, _CState100, ""},
	{_CState93, _CShiftExpressionSymbol}:           {_CShiftAction, _CState122, ""},
	{_CState93, _CRelationalExpressionSymbol}:      {_CShiftAction, _CState120, ""},
	{_CState93, _CEqualityExpressionSymbol}:        {_CShiftAction, _CState107, ""},
	{_CState93, _CAndExpressionSymbol}:             {_CShiftAction, _CState101, ""},
	{_CState93, _CExclusiveOrExpressionSymbol}:     {_CShiftAction, _CState108, ""},
	{_CState93, _CInclusiveOrExpressionSymbol}:     {_CShiftAction, _CState111, ""},
	{_CState93, _CLogicalAndExpressionSymbol}:      {_CShiftAction, _CState115, ""},
	{_CState93, _CLogicalOrExpressionSymbol}:       {_CShiftAction, _CState116, ""},
	{_CState93, _CConditionalExpressionSymbol}:     {_CShiftAction, _CState105, ""},
	{_CState93, _CAssignmentExpressionSymbol}:      {_CShiftAction, _CState102, ""},
	{_CState93, _CExpressionSymbol}:                {_CShiftAction, _CState173, ""},
	{_CState95, CIdentifierSymbol}:                 {_CShiftAction, _CState130, ""},
	{_CState95, CConstantSymbol}:                   {_CShiftAction, _CState77, ""},
	{_CState95, CStringLiteralSymbol}:              {_CShiftAction, _CState96, ""},
	{_CState95, CSizeofSymbol}:                     {_CShiftAction, _CState95, ""},
	{_CState95, CIncOpSymbol}:                      {_CShiftAction, _CState87, ""},
	{_CState95, CDecOpSymbol}:                      {_CShiftAction, _CState79, ""},
	{_CState95, CLparamSymbol}:                     {_CShiftAction, _CState174, ""},
	{_CState95, CMulSymbol}:                        {_CShiftAction, _CState90, ""},
	{_CState95, CMinusSymbol}:                      {_CShiftAction, _CState89, ""},
	{_CState95, CPlusSymbol}:                       {_CShiftAction, _CState91, ""},
	{_CState95, CAndSymbol}:                        {_CShiftAction, _CState74, ""},
	{_CState95, CExclaimSymbol}:                    {_CShiftAction, _CState82, ""},
	{_CState95, CTildaSymbol}:                      {_CShiftAction, _CState98, ""},
	{_CState95, _CPrimaryExpressionSymbol}:         {_CShiftAction, _CState119, ""},
	{_CState95, _CPostfixExpressionSymbol}:         {_CShiftAction, _CState118, ""},
	{_CState95, _CUnaryExpressionSymbol}:           {_CShiftAction, _CState175, ""},
	{_CState95, _CUnaryOperatorSymbol}:             {_CShiftAction, _CState126, ""},
	{_CState97, CLparamSymbol}:                     {_CShiftAction, _CState176, ""},
	{_CState99, CLparamSymbol}:                     {_CShiftAction, _CState177, ""},
	{_CState100, CMinusSymbol}:                     {_CShiftAction, _CState178, ""},
	{_CState100, CPlusSymbol}:                      {_CShiftAction, _CState179, ""},
	{_CState101, CAndSymbol}:                       {_CShiftAction, _CState180, ""},
	{_CState106, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState106, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState106, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState106, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState106, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState106, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState106, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState106, CTypedefSymbol}:                   {_CShiftAction, _CState19, ""},
	{_CState106, CExternSymbol}:                    {_CShiftAction, _CState7, ""},
	{_CState106, CStaticSymbol}:                    {_CShiftAction, _CState17, ""},
	{_CState106, CAutoSymbol}:                      {_CShiftAction, _CState2, ""},
	{_CState106, CRegisterSymbol}:                  {_CShiftAction, _CState14, ""},
	{_CState106, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState106, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState106, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState106, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState106, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState106, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState106, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState106, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState106, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState106, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState106, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState106, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState106, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState106, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState106, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState106, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState106, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState106, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState106, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState106, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState106, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState106, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState106, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState106, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState106, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState106, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState106, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState106, CRcurlSymbol}:                     {_CShiftAction, _CState181, ""},
	{_CState106, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState106, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState106, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState106, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState106, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState106, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState106, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState106, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState106, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState106, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState106, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState106, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState106, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState106, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState106, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState106, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState106, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState106, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState106, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState106, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState106, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState106, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState106, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState106, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState106, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState106, _CDeclarationSymbol}:              {_CShiftAction, _CState128, ""},
	{_CState106, _CDeclarationSpecifiersSymbol}:    {_CShiftAction, _CState53, ""},
	{_CState106, _CStorageClassSpecifierSymbol}:    {_CShiftAction, _CState33, ""},
	{_CState106, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState37, ""},
	{_CState106, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState106, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState106, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState106, _CTypeQualifierSymbol}:            {_CShiftAction, _CState36, ""},
	{_CState106, _CStatementSymbol}:                {_CShiftAction, _CState123, ""},
	{_CState106, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState106, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState106, _CStatementListSymbol}:            {_CShiftAction, _CState182, ""},
	{_CState106, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState106, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState106, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState106, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState107, CEqOpSymbol}:                      {_CShiftAction, _CState183, ""},
	{_CState107, CNeOpSymbol}:                      {_CShiftAction, _CState184, ""},
	{_CState108, CHatSymbol}:                       {_CShiftAction, _CState185, ""},
	{_CState109, CSemicolonSymbol}:                 {_CShiftAction, _CState187, ""},
	{_CState109, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState111, COrSymbol}:                        {_CShiftAction, _CState188, ""},
	{_CState115, CAndOpSymbol}:                     {_CShiftAction, _CState189, ""},
	{_CState116, COrOpSymbol}:                      {_CShiftAction, _CState190, ""},
	{_CState116, CQuestionSymbol}:                  {_CShiftAction, _CState191, ""},
	{_CState117, CMulSymbol}:                       {_CShiftAction, _CState194, ""},
	{_CState117, CDivSymbol}:                       {_CShiftAction, _CState192, ""},
	{_CState117, CModSymbol}:                       {_CShiftAction, _CState193, ""},
	{_CState118, CPtrOpSymbol}:                     {_CShiftAction, _CState200, ""},
	{_CState118, CIncOpSymbol}:                     {_CShiftAction, _CState197, ""},
	{_CState118, CDecOpSymbol}:                     {_CShiftAction, _CState195, ""},
	{_CState118, CLparamSymbol}:                    {_CShiftAction, _CState199, ""},
	{_CState118, CLbraceSymbol}:                    {_CShiftAction, _CState198, ""},
	{_CState118, CDotSymbol}:                       {_CShiftAction, _CState196, ""},
	{_CState120, CLeOpSymbol}:                      {_CShiftAction, _CState203, ""},
	{_CState120, CGeOpSymbol}:                      {_CShiftAction, _CState201, ""},
	{_CState120, CLtSymbol}:                        {_CShiftAction, _CState204, ""},
	{_CState120, CGtSymbol}:                        {_CShiftAction, _CState202, ""},
	{_CState122, CLeftOpSymbol}:                    {_CShiftAction, _CState205, ""},
	{_CState122, CRightOpSymbol}:                   {_CShiftAction, _CState206, ""},
	{_CState124, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState124, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState124, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState124, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState124, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState124, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState124, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState124, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState124, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState124, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState124, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState124, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState124, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState124, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState124, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState124, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState124, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState124, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState124, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState124, CRcurlSymbol}:                     {_CShiftAction, _CState207, ""},
	{_CState124, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState124, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState124, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState124, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState124, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState124, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState124, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState124, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState124, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState124, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState124, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState124, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState124, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState124, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState124, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState124, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState124, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState124, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState124, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState124, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState124, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState124, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState124, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState124, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState124, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState124, _CStatementSymbol}:                {_CShiftAction, _CState208, ""},
	{_CState124, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState124, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState124, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState124, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState124, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState124, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState125, CMulAssignSymbol}:                 {_CShiftAction, _CState215, ""},
	{_CState125, CDivAssignSymbol}:                 {_CShiftAction, _CState211, ""},
	{_CState125, CModAssignSymbol}:                 {_CShiftAction, _CState214, ""},
	{_CState125, CAddAssignSymbol}:                 {_CShiftAction, _CState209, ""},
	{_CState125, CSubAssignSymbol}:                 {_CShiftAction, _CState218, ""},
	{_CState125, CLeftAssignSymbol}:                {_CShiftAction, _CState213, ""},
	{_CState125, CRightAssignSymbol}:               {_CShiftAction, _CState217, ""},
	{_CState125, CAndAssignSymbol}:                 {_CShiftAction, _CState210, ""},
	{_CState125, CXorAssignSymbol}:                 {_CShiftAction, _CState219, ""},
	{_CState125, COrAssignSymbol}:                  {_CShiftAction, _CState216, ""},
	{_CState125, CEqSymbol}:                        {_CShiftAction, _CState212, ""},
	{_CState125, _CAssignmentOperatorSymbol}:       {_CShiftAction, _CState220, ""},
	{_CState126, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState126, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState126, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState126, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState126, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState126, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState126, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState126, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState126, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState126, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState126, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState126, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState126, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState126, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState126, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState126, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState126, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState126, _CCastExpressionSymbol}:           {_CShiftAction, _CState221, ""},
	{_CState129, CEqSymbol}:                        {_CShiftAction, _CState69, ""},
	{_CState133, CRbraceSymbol}:                    {_CShiftAction, _CState222, ""},
	{_CState137, CIdentifierSymbol}:                {_CShiftAction, _CState9, ""},
	{_CState137, CLparamSymbol}:                    {_CShiftAction, _CState224, ""},
	{_CState137, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState137, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState137, _CDeclaratorSymbol}:               {_CShiftAction, _CState226, ""},
	{_CState137, _CDirectDeclaratorSymbol}:         {_CShiftAction, _CState28, ""},
	{_CState137, _CPointerSymbol}:                  {_CShiftAction, _CState228, ""},
	{_CState137, _CAbstractDeclaratorSymbol}:       {_CShiftAction, _CState225, ""},
	{_CState137, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState227, ""},
	{_CState138, CRparamSymbol}:                    {_CShiftAction, _CState230, ""},
	{_CState138, CCommaSymbol}:                     {_CShiftAction, _CState229, ""},
	{_CState140, CCommaSymbol}:                     {_CShiftAction, _CState231, ""},
	{_CState141, CRparamSymbol}:                    {_CShiftAction, _CState232, ""},
	{_CState142, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState142, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState142, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState142, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState142, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState142, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState142, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState142, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState142, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState142, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState142, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState142, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState142, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState142, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState142, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState142, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState142, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState142, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState142, _CStructDeclarationListSymbol}:    {_CShiftAction, _CState233, ""},
	{_CState142, _CStructDeclarationSymbol}:        {_CShiftAction, _CState144, ""},
	{_CState142, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState143, ""},
	{_CState142, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState142, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState143, CIdentifierSymbol}:                {_CShiftAction, _CState9, ""},
	{_CState143, CLparamSymbol}:                    {_CShiftAction, _CState12, ""},
	{_CState143, CColonSymbol}:                     {_CShiftAction, _CState234, ""},
	{_CState143, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState143, _CStructDeclaratorListSymbol}:     {_CShiftAction, _CState237, ""},
	{_CState143, _CStructDeclaratorSymbol}:         {_CShiftAction, _CState236, ""},
	{_CState143, _CDeclaratorSymbol}:               {_CShiftAction, _CState235, ""},
	{_CState143, _CDirectDeclaratorSymbol}:         {_CShiftAction, _CState28, ""},
	{_CState143, _CPointerSymbol}:                  {_CShiftAction, _CState32, ""},
	{_CState145, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState145, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState145, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState145, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState145, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState145, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState145, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState145, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState145, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState145, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState145, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState145, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState145, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState145, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState145, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState145, CRcurlSymbol}:                     {_CShiftAction, _CState238, ""},
	{_CState145, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState145, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState145, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState145, _CStructDeclarationSymbol}:        {_CShiftAction, _CState239, ""},
	{_CState145, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState143, ""},
	{_CState145, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState145, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState146, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState146, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState146, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState146, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState146, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState146, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState146, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState146, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState146, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState146, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState146, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState146, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState146, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState146, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState146, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState146, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState146, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState146, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState146, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState240, ""},
	{_CState146, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState146, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState147, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState147, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState147, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState147, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState147, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState147, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState147, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState147, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState147, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState147, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState147, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState147, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState147, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState147, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState147, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState147, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState147, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState147, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState147, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState241, ""},
	{_CState147, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState147, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState148, CRcurlSymbol}:                     {_CShiftAction, _CState242, ""},
	{_CState148, CCommaSymbol}:                     {_CShiftAction, _CState150, ""},
	{_CState149, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState149, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState149, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState149, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState149, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState149, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState149, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState149, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState149, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState149, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState149, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState149, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState149, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState149, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState149, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState149, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState149, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState149, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState149, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState149, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState149, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState149, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState149, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState149, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState149, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState149, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState149, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState149, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState149, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState132, ""},
	{_CState149, _CConstantExpressionSymbol}:       {_CShiftAction, _CState243, ""},
	{_CState150, CIdentifierSymbol}:                {_CShiftAction, _CState63, ""},
	{_CState150, _CEnumeratorSymbol}:               {_CShiftAction, _CState244, ""},
	{_CState152, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState152, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState152, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState152, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState152, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState152, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState152, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState152, CLcurlSymbol}:                     {_CShiftAction, _CState152, ""},
	{_CState152, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState152, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState152, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState152, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState152, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState152, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState152, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState152, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState152, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState152, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState152, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState152, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState152, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState152, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState152, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState152, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState152, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState152, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState152, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState152, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState152, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState152, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState152, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState153, ""},
	{_CState152, _CInitializerSymbol}:              {_CShiftAction, _CState245, ""},
	{_CState152, _CInitializerListSymbol}:          {_CShiftAction, _CState246, ""},
	{_CState158, CColonSymbol}:                     {_CShiftAction, _CState247, ""},
	{_CState160, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState160, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState160, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState160, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState160, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState160, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState160, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState160, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState160, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState160, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState160, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState160, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState160, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState160, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState160, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState160, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState160, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState160, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState160, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState160, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState160, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState160, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState160, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState160, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState160, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState160, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState160, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState160, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState160, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState160, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState160, _CExpressionSymbol}:               {_CShiftAction, _CState169, ""},
	{_CState162, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState162, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState162, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState162, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState162, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState162, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState162, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState162, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState162, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState162, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState162, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState162, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState162, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState162, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState162, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState162, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState162, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState162, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState162, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState162, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState162, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState162, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState162, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState162, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState162, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState162, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState162, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState162, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState162, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState162, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState162, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState162, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState162, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState162, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState162, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState162, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState162, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState162, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState162, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState162, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState162, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState162, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState162, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState162, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState162, _CStatementSymbol}:                {_CShiftAction, _CState248, ""},
	{_CState162, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState162, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState162, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState162, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState162, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState162, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState163, CWhileSymbol}:                     {_CShiftAction, _CState249, ""},
	{_CState164, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState164, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState164, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState164, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState164, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState164, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState164, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState164, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState164, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState164, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState164, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState164, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState164, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState164, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState164, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState164, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState164, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState164, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState164, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState164, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState164, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState164, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState164, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState164, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState164, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState164, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState164, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState164, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState164, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState164, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState164, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState164, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState164, _CExpressionStatementSymbol}:      {_CShiftAction, _CState250, ""},
	{_CState165, CSemicolonSymbol}:                 {_CShiftAction, _CState251, ""},
	{_CState166, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState166, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState166, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState166, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState166, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState166, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState166, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState166, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState166, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState166, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState166, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState166, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState166, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState166, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState166, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState166, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState166, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState166, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState166, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState166, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState166, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState166, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState166, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState166, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState166, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState166, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState166, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState166, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState166, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState166, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState166, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState166, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState166, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState166, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState166, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState166, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState166, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState166, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState166, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState166, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState166, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState166, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState166, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState166, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState166, _CStatementSymbol}:                {_CShiftAction, _CState252, ""},
	{_CState166, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState166, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState166, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState166, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState166, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState166, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState167, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState167, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState167, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState167, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState167, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState167, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState167, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState167, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState167, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState167, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState167, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState167, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState167, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState167, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState167, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState167, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState167, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState167, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState167, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState167, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState167, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState167, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState167, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState167, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState167, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState167, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState167, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState167, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState167, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState167, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState167, _CExpressionSymbol}:               {_CShiftAction, _CState253, ""},
	{_CState169, CRparamSymbol}:                    {_CShiftAction, _CState254, ""},
	{_CState169, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState170, CLparamSymbol}:                    {_CShiftAction, _CState255, ""},
	{_CState170, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState170, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState170, _CPointerSymbol}:                  {_CShiftAction, _CState257, ""},
	{_CState170, _CAbstractDeclaratorSymbol}:       {_CShiftAction, _CState256, ""},
	{_CState170, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState227, ""},
	{_CState171, CRparamSymbol}:                    {_CShiftAction, _CState258, ""},
	{_CState173, CSemicolonSymbol}:                 {_CShiftAction, _CState259, ""},
	{_CState173, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState174, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState174, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState174, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState174, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState174, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState174, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState174, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState174, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState174, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState174, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState174, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState174, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState174, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState174, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState174, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState174, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState174, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState174, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState174, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState174, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState174, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState174, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState174, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState174, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState174, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState174, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState174, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState174, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState174, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState174, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState174, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState174, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState174, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState174, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState174, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState174, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState174, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState174, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState174, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState174, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState174, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState174, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState174, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState174, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState174, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState174, _CExpressionSymbol}:               {_CShiftAction, _CState169, ""},
	{_CState174, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState174, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState174, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState174, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState170, ""},
	{_CState174, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState174, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState174, _CTypeNameSymbol}:                 {_CShiftAction, _CState260, ""},
	{_CState176, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState176, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState176, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState176, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState176, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState176, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState176, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState176, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState176, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState176, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState176, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState176, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState176, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState176, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState176, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState176, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState176, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState176, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState176, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState176, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState176, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState176, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState176, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState176, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState176, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState176, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState176, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState176, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState176, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState176, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState176, _CExpressionSymbol}:               {_CShiftAction, _CState261, ""},
	{_CState177, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState177, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState177, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState177, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState177, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState177, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState177, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState177, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState177, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState177, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState177, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState177, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState177, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState177, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState177, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState177, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState177, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState177, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState177, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState177, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState177, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState177, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState177, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState177, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState177, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState177, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState177, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState177, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState177, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState177, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState177, _CExpressionSymbol}:               {_CShiftAction, _CState262, ""},
	{_CState178, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState178, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState178, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState178, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState178, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState178, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState178, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState178, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState178, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState178, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState178, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState178, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState178, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState178, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState178, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState178, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState178, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState178, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState178, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState263, ""},
	{_CState179, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState179, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState179, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState179, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState179, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState179, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState179, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState179, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState179, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState179, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState179, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState179, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState179, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState179, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState179, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState179, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState179, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState179, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState179, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState264, ""},
	{_CState180, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState180, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState180, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState180, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState180, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState180, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState180, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState180, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState180, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState180, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState180, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState180, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState180, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState180, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState180, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState180, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState180, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState180, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState180, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState180, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState180, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState180, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState180, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState265, ""},
	{_CState182, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState182, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState182, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState182, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState182, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState182, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState182, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState182, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState182, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState182, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState182, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState182, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState182, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState182, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState182, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState182, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState182, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState182, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState182, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState182, CRcurlSymbol}:                     {_CShiftAction, _CState266, ""},
	{_CState182, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState182, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState182, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState182, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState182, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState182, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState182, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState182, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState182, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState182, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState182, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState182, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState182, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState182, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState182, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState182, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState182, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState182, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState182, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState182, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState182, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState182, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState182, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState182, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState182, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState182, _CStatementSymbol}:                {_CShiftAction, _CState208, ""},
	{_CState182, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState182, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState182, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState182, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState182, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState182, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState183, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState183, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState183, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState183, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState183, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState183, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState183, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState183, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState183, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState183, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState183, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState183, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState183, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState183, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState183, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState183, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState183, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState183, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState183, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState183, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState183, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState183, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState267, ""},
	{_CState184, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState184, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState184, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState184, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState184, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState184, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState184, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState184, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState184, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState184, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState184, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState184, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState184, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState184, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState184, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState184, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState184, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState184, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState184, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState184, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState184, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState184, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState268, ""},
	{_CState185, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState185, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState185, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState185, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState185, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState185, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState185, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState185, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState185, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState185, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState185, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState185, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState185, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState185, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState185, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState185, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState185, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState185, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState185, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState185, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState185, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState185, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState185, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState185, _CAndExpressionSymbol}:            {_CShiftAction, _CState269, ""},
	{_CState186, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState186, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState186, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState186, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState186, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState186, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState186, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState186, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState186, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState186, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState186, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState186, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState186, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState186, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState186, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState186, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState186, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState186, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState186, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState186, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState186, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState186, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState186, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState186, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState186, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState186, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState186, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState186, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState186, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState186, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState270, ""},
	{_CState188, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState188, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState188, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState188, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState188, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState188, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState188, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState188, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState188, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState188, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState188, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState188, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState188, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState188, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState188, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState188, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState188, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState188, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState188, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState188, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState188, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState188, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState188, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState188, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState188, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState271, ""},
	{_CState189, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState189, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState189, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState189, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState189, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState189, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState189, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState189, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState189, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState189, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState189, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState189, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState189, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState189, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState189, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState189, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState189, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState189, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState189, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState189, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState189, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState189, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState189, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState189, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState189, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState189, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState272, ""},
	{_CState190, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState190, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState190, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState190, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState190, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState190, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState190, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState190, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState190, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState190, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState190, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState190, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState190, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState190, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState190, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState190, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState190, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState190, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState190, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState190, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState190, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState190, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState190, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState190, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState190, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState190, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState190, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState273, ""},
	{_CState191, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState191, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState191, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState191, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState191, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState191, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState191, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState191, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState191, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState191, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState191, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState191, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState191, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState191, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState191, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState191, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState191, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState191, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState191, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState191, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState191, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState191, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState191, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState191, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState191, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState191, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState191, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState191, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState191, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState191, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState191, _CExpressionSymbol}:               {_CShiftAction, _CState274, ""},
	{_CState192, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState192, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState192, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState192, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState192, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState192, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState192, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState192, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState192, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState192, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState192, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState192, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState192, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState192, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState192, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState192, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState192, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState192, _CCastExpressionSymbol}:           {_CShiftAction, _CState275, ""},
	{_CState193, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState193, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState193, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState193, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState193, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState193, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState193, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState193, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState193, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState193, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState193, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState193, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState193, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState193, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState193, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState193, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState193, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState193, _CCastExpressionSymbol}:           {_CShiftAction, _CState276, ""},
	{_CState194, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState194, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState194, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState194, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState194, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState194, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState194, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState194, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState194, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState194, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState194, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState194, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState194, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState194, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState194, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState194, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState194, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState194, _CCastExpressionSymbol}:           {_CShiftAction, _CState277, ""},
	{_CState196, CIdentifierSymbol}:                {_CShiftAction, _CState278, ""},
	{_CState198, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState198, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState198, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState198, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState198, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState198, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState198, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState198, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState198, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState198, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState198, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState198, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState198, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState198, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState198, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState198, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState198, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState198, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState198, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState198, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState198, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState198, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState198, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState198, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState198, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState198, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState198, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState198, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState198, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState198, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState198, _CExpressionSymbol}:               {_CShiftAction, _CState279, ""},
	{_CState199, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState199, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState199, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState199, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState199, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState199, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState199, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState199, CRparamSymbol}:                    {_CShiftAction, _CState280, ""},
	{_CState199, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState199, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState199, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState199, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState199, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState199, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState199, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState199, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState199, _CArgumentExpressionListSymbol}:   {_CShiftAction, _CState281, ""},
	{_CState199, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState199, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState199, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState199, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState199, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState199, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState199, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState199, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState199, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState199, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState199, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState199, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState199, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState199, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState199, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState282, ""},
	{_CState200, CIdentifierSymbol}:                {_CShiftAction, _CState283, ""},
	{_CState201, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState201, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState201, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState201, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState201, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState201, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState201, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState201, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState201, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState201, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState201, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState201, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState201, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState201, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState201, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState201, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState201, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState201, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState201, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState201, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState201, _CShiftExpressionSymbol}:          {_CShiftAction, _CState284, ""},
	{_CState202, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState202, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState202, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState202, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState202, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState202, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState202, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState202, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState202, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState202, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState202, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState202, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState202, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState202, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState202, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState202, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState202, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState202, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState202, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState202, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState202, _CShiftExpressionSymbol}:          {_CShiftAction, _CState285, ""},
	{_CState203, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState203, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState203, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState203, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState203, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState203, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState203, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState203, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState203, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState203, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState203, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState203, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState203, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState203, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState203, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState203, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState203, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState203, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState203, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState203, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState203, _CShiftExpressionSymbol}:          {_CShiftAction, _CState286, ""},
	{_CState204, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState204, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState204, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState204, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState204, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState204, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState204, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState204, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState204, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState204, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState204, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState204, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState204, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState204, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState204, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState204, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState204, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState204, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState204, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState204, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState204, _CShiftExpressionSymbol}:          {_CShiftAction, _CState287, ""},
	{_CState205, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState205, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState205, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState205, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState205, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState205, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState205, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState205, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState205, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState205, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState205, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState205, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState205, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState205, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState205, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState205, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState205, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState205, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState205, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState205, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState288, ""},
	{_CState206, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState206, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState206, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState206, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState206, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState206, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState206, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState206, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState206, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState206, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState206, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState206, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState206, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState206, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState206, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState206, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState206, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState206, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState206, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState206, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState289, ""},
	{_CState220, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState220, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState220, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState220, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState220, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState220, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState220, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState220, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState220, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState220, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState220, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState220, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState220, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState220, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState220, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState220, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState220, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState220, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState220, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState220, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState220, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState220, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState220, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState220, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState220, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState220, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState220, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState220, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState220, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState220, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState290, ""},
	{_CState223, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState223, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState223, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState223, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState223, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState223, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState223, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState223, CRbraceSymbol}:                    {_CShiftAction, _CState291, ""},
	{_CState223, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState223, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState223, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState223, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState223, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState223, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState223, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState223, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState223, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState223, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState223, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState223, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState223, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState223, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState223, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState223, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState223, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState223, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState223, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState223, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState223, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState223, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState132, ""},
	{_CState223, _CConstantExpressionSymbol}:       {_CShiftAction, _CState292, ""},
	{_CState224, CIdentifierSymbol}:                {_CShiftAction, _CState9, ""},
	{_CState224, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState224, CTypedefSymbol}:                   {_CShiftAction, _CState19, ""},
	{_CState224, CExternSymbol}:                    {_CShiftAction, _CState7, ""},
	{_CState224, CStaticSymbol}:                    {_CShiftAction, _CState17, ""},
	{_CState224, CAutoSymbol}:                      {_CShiftAction, _CState2, ""},
	{_CState224, CRegisterSymbol}:                  {_CShiftAction, _CState14, ""},
	{_CState224, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState224, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState224, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState224, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState224, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState224, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState224, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState224, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState224, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState224, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState224, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState224, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState224, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState224, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState224, CLparamSymbol}:                    {_CShiftAction, _CState224, ""},
	{_CState224, CRparamSymbol}:                    {_CShiftAction, _CState293, ""},
	{_CState224, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState224, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState224, _CDeclarationSpecifiersSymbol}:    {_CShiftAction, _CState137, ""},
	{_CState224, _CStorageClassSpecifierSymbol}:    {_CShiftAction, _CState33, ""},
	{_CState224, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState37, ""},
	{_CState224, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState224, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState224, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState224, _CTypeQualifierSymbol}:            {_CShiftAction, _CState36, ""},
	{_CState224, _CDeclaratorSymbol}:               {_CShiftAction, _CState41, ""},
	{_CState224, _CDirectDeclaratorSymbol}:         {_CShiftAction, _CState28, ""},
	{_CState224, _CPointerSymbol}:                  {_CShiftAction, _CState228, ""},
	{_CState224, _CParameterTypeListSymbol}:        {_CShiftAction, _CState295, ""},
	{_CState224, _CParameterListSymbol}:            {_CShiftAction, _CState140, ""},
	{_CState224, _CParameterDeclarationSymbol}:     {_CShiftAction, _CState139, ""},
	{_CState224, _CAbstractDeclaratorSymbol}:       {_CShiftAction, _CState294, ""},
	{_CState224, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState227, ""},
	{_CState227, CLparamSymbol}:                    {_CShiftAction, _CState297, ""},
	{_CState227, CLbraceSymbol}:                    {_CShiftAction, _CState296, ""},
	{_CState228, CIdentifierSymbol}:                {_CShiftAction, _CState9, ""},
	{_CState228, CLparamSymbol}:                    {_CShiftAction, _CState224, ""},
	{_CState228, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState228, _CDirectDeclaratorSymbol}:         {_CShiftAction, _CState56, ""},
	{_CState228, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState298, ""},
	{_CState229, CIdentifierSymbol}:                {_CShiftAction, _CState299, ""},
	{_CState231, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState231, CTypedefSymbol}:                   {_CShiftAction, _CState19, ""},
	{_CState231, CExternSymbol}:                    {_CShiftAction, _CState7, ""},
	{_CState231, CStaticSymbol}:                    {_CShiftAction, _CState17, ""},
	{_CState231, CAutoSymbol}:                      {_CShiftAction, _CState2, ""},
	{_CState231, CRegisterSymbol}:                  {_CShiftAction, _CState14, ""},
	{_CState231, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState231, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState231, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState231, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState231, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState231, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState231, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState231, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState231, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState231, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState231, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState231, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState231, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState231, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState231, CEllipsisSymbol}:                  {_CShiftAction, _CState300, ""},
	{_CState231, _CDeclarationSpecifiersSymbol}:    {_CShiftAction, _CState137, ""},
	{_CState231, _CStorageClassSpecifierSymbol}:    {_CShiftAction, _CState33, ""},
	{_CState231, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState37, ""},
	{_CState231, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState231, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState231, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState231, _CTypeQualifierSymbol}:            {_CShiftAction, _CState36, ""},
	{_CState231, _CParameterDeclarationSymbol}:     {_CShiftAction, _CState301, ""},
	{_CState233, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState233, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState233, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState233, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState233, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState233, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState233, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState233, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState233, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState233, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState233, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState233, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState233, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState233, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState233, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState233, CRcurlSymbol}:                     {_CShiftAction, _CState302, ""},
	{_CState233, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState147, ""},
	{_CState233, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState233, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState233, _CStructDeclarationSymbol}:        {_CShiftAction, _CState239, ""},
	{_CState233, _CSpecifierQualifierListSymbol}:   {_CShiftAction, _CState143, ""},
	{_CState233, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState233, _CTypeQualifierSymbol}:            {_CShiftAction, _CState146, ""},
	{_CState234, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState234, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState234, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState234, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState234, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState234, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState234, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState234, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState234, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState234, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState234, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState234, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState234, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState234, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState234, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState234, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState234, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState234, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState234, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState234, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState234, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState234, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState234, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState234, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState234, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState234, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState234, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState234, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState234, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState132, ""},
	{_CState234, _CConstantExpressionSymbol}:       {_CShiftAction, _CState303, ""},
	{_CState235, CColonSymbol}:                     {_CShiftAction, _CState304, ""},
	{_CState237, CSemicolonSymbol}:                 {_CShiftAction, _CState306, ""},
	{_CState237, CCommaSymbol}:                     {_CShiftAction, _CState305, ""},
	{_CState246, CRcurlSymbol}:                     {_CShiftAction, _CState308, ""},
	{_CState246, CCommaSymbol}:                     {_CShiftAction, _CState307, ""},
	{_CState247, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState247, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState247, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState247, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState247, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState247, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState247, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState247, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState247, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState247, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState247, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState247, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState247, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState247, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState247, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState247, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState247, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState247, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState247, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState247, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState247, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState247, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState247, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState247, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState247, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState247, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState247, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState247, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState247, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState247, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState247, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState247, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState247, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState247, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState247, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState247, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState247, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState247, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState247, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState247, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState247, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState247, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState247, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState247, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState247, _CStatementSymbol}:                {_CShiftAction, _CState309, ""},
	{_CState247, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState247, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState247, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState247, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState247, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState247, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState249, CLparamSymbol}:                    {_CShiftAction, _CState310, ""},
	{_CState250, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState250, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState250, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState250, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState250, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState250, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState250, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState250, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState250, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState250, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState250, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState250, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState250, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState250, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState250, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState250, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState250, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState250, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState250, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState250, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState250, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState250, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState250, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState250, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState250, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState250, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState250, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState250, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState250, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState250, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState250, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState250, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState250, _CExpressionStatementSymbol}:      {_CShiftAction, _CState311, ""},
	{_CState253, CRparamSymbol}:                    {_CShiftAction, _CState312, ""},
	{_CState253, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState255, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState255, CTypedefSymbol}:                   {_CShiftAction, _CState19, ""},
	{_CState255, CExternSymbol}:                    {_CShiftAction, _CState7, ""},
	{_CState255, CStaticSymbol}:                    {_CShiftAction, _CState17, ""},
	{_CState255, CAutoSymbol}:                      {_CShiftAction, _CState2, ""},
	{_CState255, CRegisterSymbol}:                  {_CShiftAction, _CState14, ""},
	{_CState255, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState255, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState255, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState255, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState255, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState255, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState255, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState255, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState255, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState255, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState255, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState255, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState255, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState255, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState255, CLparamSymbol}:                    {_CShiftAction, _CState255, ""},
	{_CState255, CRparamSymbol}:                    {_CShiftAction, _CState293, ""},
	{_CState255, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState255, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState255, _CDeclarationSpecifiersSymbol}:    {_CShiftAction, _CState137, ""},
	{_CState255, _CStorageClassSpecifierSymbol}:    {_CShiftAction, _CState33, ""},
	{_CState255, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState37, ""},
	{_CState255, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState255, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState255, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState255, _CTypeQualifierSymbol}:            {_CShiftAction, _CState36, ""},
	{_CState255, _CPointerSymbol}:                  {_CShiftAction, _CState257, ""},
	{_CState255, _CParameterTypeListSymbol}:        {_CShiftAction, _CState295, ""},
	{_CState255, _CParameterListSymbol}:            {_CShiftAction, _CState140, ""},
	{_CState255, _CParameterDeclarationSymbol}:     {_CShiftAction, _CState139, ""},
	{_CState255, _CAbstractDeclaratorSymbol}:       {_CShiftAction, _CState294, ""},
	{_CState255, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState227, ""},
	{_CState257, CLparamSymbol}:                    {_CShiftAction, _CState255, ""},
	{_CState257, CLbraceSymbol}:                    {_CShiftAction, _CState223, ""},
	{_CState257, _CDirectAbstractDeclaratorSymbol}: {_CShiftAction, _CState298, ""},
	{_CState258, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState258, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState258, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState258, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState258, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState258, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState258, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState258, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState258, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState258, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState258, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState258, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState258, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState258, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState258, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState258, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState258, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState258, _CCastExpressionSymbol}:           {_CShiftAction, _CState313, ""},
	{_CState260, CRparamSymbol}:                    {_CShiftAction, _CState314, ""},
	{_CState261, CRparamSymbol}:                    {_CShiftAction, _CState315, ""},
	{_CState261, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState262, CRparamSymbol}:                    {_CShiftAction, _CState316, ""},
	{_CState262, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState263, CMulSymbol}:                       {_CShiftAction, _CState194, ""},
	{_CState263, CDivSymbol}:                       {_CShiftAction, _CState192, ""},
	{_CState263, CModSymbol}:                       {_CShiftAction, _CState193, ""},
	{_CState264, CMulSymbol}:                       {_CShiftAction, _CState194, ""},
	{_CState264, CDivSymbol}:                       {_CShiftAction, _CState192, ""},
	{_CState264, CModSymbol}:                       {_CShiftAction, _CState193, ""},
	{_CState265, CEqOpSymbol}:                      {_CShiftAction, _CState183, ""},
	{_CState265, CNeOpSymbol}:                      {_CShiftAction, _CState184, ""},
	{_CState267, CLeOpSymbol}:                      {_CShiftAction, _CState203, ""},
	{_CState267, CGeOpSymbol}:                      {_CShiftAction, _CState201, ""},
	{_CState267, CLtSymbol}:                        {_CShiftAction, _CState204, ""},
	{_CState267, CGtSymbol}:                        {_CShiftAction, _CState202, ""},
	{_CState268, CLeOpSymbol}:                      {_CShiftAction, _CState203, ""},
	{_CState268, CGeOpSymbol}:                      {_CShiftAction, _CState201, ""},
	{_CState268, CLtSymbol}:                        {_CShiftAction, _CState204, ""},
	{_CState268, CGtSymbol}:                        {_CShiftAction, _CState202, ""},
	{_CState269, CAndSymbol}:                       {_CShiftAction, _CState180, ""},
	{_CState271, CHatSymbol}:                       {_CShiftAction, _CState185, ""},
	{_CState272, COrSymbol}:                        {_CShiftAction, _CState188, ""},
	{_CState273, CAndOpSymbol}:                     {_CShiftAction, _CState189, ""},
	{_CState274, CColonSymbol}:                     {_CShiftAction, _CState317, ""},
	{_CState274, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState279, CRbraceSymbol}:                    {_CShiftAction, _CState318, ""},
	{_CState279, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState281, CRparamSymbol}:                    {_CShiftAction, _CState320, ""},
	{_CState281, CCommaSymbol}:                     {_CShiftAction, _CState319, ""},
	{_CState284, CLeftOpSymbol}:                    {_CShiftAction, _CState205, ""},
	{_CState284, CRightOpSymbol}:                   {_CShiftAction, _CState206, ""},
	{_CState285, CLeftOpSymbol}:                    {_CShiftAction, _CState205, ""},
	{_CState285, CRightOpSymbol}:                   {_CShiftAction, _CState206, ""},
	{_CState286, CLeftOpSymbol}:                    {_CShiftAction, _CState205, ""},
	{_CState286, CRightOpSymbol}:                   {_CShiftAction, _CState206, ""},
	{_CState287, CLeftOpSymbol}:                    {_CShiftAction, _CState205, ""},
	{_CState287, CRightOpSymbol}:                   {_CShiftAction, _CState206, ""},
	{_CState288, CMinusSymbol}:                     {_CShiftAction, _CState178, ""},
	{_CState288, CPlusSymbol}:                      {_CShiftAction, _CState179, ""},
	{_CState289, CMinusSymbol}:                     {_CShiftAction, _CState178, ""},
	{_CState289, CPlusSymbol}:                      {_CShiftAction, _CState179, ""},
	{_CState292, CRbraceSymbol}:                    {_CShiftAction, _CState321, ""},
	{_CState294, CRparamSymbol}:                    {_CShiftAction, _CState322, ""},
	{_CState295, CRparamSymbol}:                    {_CShiftAction, _CState323, ""},
	{_CState296, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState296, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState296, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState296, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState296, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState296, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState296, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState296, CRbraceSymbol}:                    {_CShiftAction, _CState324, ""},
	{_CState296, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState296, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState296, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState296, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState296, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState296, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState296, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState296, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState296, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState296, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState296, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState296, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState296, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState296, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState296, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState296, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState296, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState296, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState296, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState296, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState296, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState296, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState132, ""},
	{_CState296, _CConstantExpressionSymbol}:       {_CShiftAction, _CState325, ""},
	{_CState297, CTypeNameSymbol}:                  {_CShiftAction, _CState20, ""},
	{_CState297, CTypedefSymbol}:                   {_CShiftAction, _CState19, ""},
	{_CState297, CExternSymbol}:                    {_CShiftAction, _CState7, ""},
	{_CState297, CStaticSymbol}:                    {_CShiftAction, _CState17, ""},
	{_CState297, CAutoSymbol}:                      {_CShiftAction, _CState2, ""},
	{_CState297, CRegisterSymbol}:                  {_CShiftAction, _CState14, ""},
	{_CState297, CCharSymbol}:                      {_CShiftAction, _CState3, ""},
	{_CState297, CShortSymbol}:                     {_CShiftAction, _CState15, ""},
	{_CState297, CIntSymbol}:                       {_CShiftAction, _CState10, ""},
	{_CState297, CLongSymbol}:                      {_CShiftAction, _CState11, ""},
	{_CState297, CSignedSymbol}:                    {_CShiftAction, _CState16, ""},
	{_CState297, CUnsignedSymbol}:                  {_CShiftAction, _CState22, ""},
	{_CState297, CFloatSymbol}:                     {_CShiftAction, _CState8, ""},
	{_CState297, CDoubleSymbol}:                    {_CShiftAction, _CState5, ""},
	{_CState297, CConstSymbol}:                     {_CShiftAction, _CState4, ""},
	{_CState297, CVolatileSymbol}:                  {_CShiftAction, _CState24, ""},
	{_CState297, CVoidSymbol}:                      {_CShiftAction, _CState23, ""},
	{_CState297, CStructSymbol}:                    {_CShiftAction, _CState18, ""},
	{_CState297, CUnionSymbol}:                     {_CShiftAction, _CState21, ""},
	{_CState297, CEnumSymbol}:                      {_CShiftAction, _CState6, ""},
	{_CState297, CRparamSymbol}:                    {_CShiftAction, _CState326, ""},
	{_CState297, _CDeclarationSpecifiersSymbol}:    {_CShiftAction, _CState137, ""},
	{_CState297, _CStorageClassSpecifierSymbol}:    {_CShiftAction, _CState33, ""},
	{_CState297, _CTypeSpecifierSymbol}:            {_CShiftAction, _CState37, ""},
	{_CState297, _CStructOrUnionSpecifierSymbol}:   {_CShiftAction, _CState35, ""},
	{_CState297, _CStructOrUnionSymbol}:            {_CShiftAction, _CState34, ""},
	{_CState297, _CEnumSpecifierSymbol}:            {_CShiftAction, _CState29, ""},
	{_CState297, _CTypeQualifierSymbol}:            {_CShiftAction, _CState36, ""},
	{_CState297, _CParameterTypeListSymbol}:        {_CShiftAction, _CState327, ""},
	{_CState297, _CParameterListSymbol}:            {_CShiftAction, _CState140, ""},
	{_CState297, _CParameterDeclarationSymbol}:     {_CShiftAction, _CState139, ""},
	{_CState298, CLparamSymbol}:                    {_CShiftAction, _CState297, ""},
	{_CState298, CLbraceSymbol}:                    {_CShiftAction, _CState296, ""},
	{_CState304, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState304, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState304, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState304, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState304, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState304, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState304, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState304, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState304, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState304, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState304, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState304, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState304, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState304, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState304, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState304, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState304, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState304, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState304, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState304, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState304, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState304, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState304, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState304, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState304, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState304, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState304, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState304, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState304, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState132, ""},
	{_CState304, _CConstantExpressionSymbol}:       {_CShiftAction, _CState328, ""},
	{_CState305, CIdentifierSymbol}:                {_CShiftAction, _CState9, ""},
	{_CState305, CLparamSymbol}:                    {_CShiftAction, _CState12, ""},
	{_CState305, CColonSymbol}:                     {_CShiftAction, _CState234, ""},
	{_CState305, CMulSymbol}:                       {_CShiftAction, _CState13, ""},
	{_CState305, _CStructDeclaratorSymbol}:         {_CShiftAction, _CState329, ""},
	{_CState305, _CDeclaratorSymbol}:               {_CShiftAction, _CState235, ""},
	{_CState305, _CDirectDeclaratorSymbol}:         {_CShiftAction, _CState28, ""},
	{_CState305, _CPointerSymbol}:                  {_CShiftAction, _CState32, ""},
	{_CState307, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState307, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState307, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState307, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState307, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState307, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState307, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState307, CLcurlSymbol}:                     {_CShiftAction, _CState152, ""},
	{_CState307, CRcurlSymbol}:                     {_CShiftAction, _CState330, ""},
	{_CState307, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState307, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState307, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState307, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState307, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState307, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState307, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState307, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState307, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState307, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState307, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState307, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState307, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState307, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState307, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState307, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState307, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState307, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState307, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState307, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState307, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState307, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState307, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState153, ""},
	{_CState307, _CInitializerSymbol}:              {_CShiftAction, _CState331, ""},
	{_CState310, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState310, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState310, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState310, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState310, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState310, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState310, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState310, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState310, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState310, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState310, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState310, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState310, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState310, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState310, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState310, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState310, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState310, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState310, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState310, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState310, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState310, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState310, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState310, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState310, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState310, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState310, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState310, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState310, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState310, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState310, _CExpressionSymbol}:               {_CShiftAction, _CState332, ""},
	{_CState311, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState311, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState311, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState311, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState311, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState311, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState311, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState311, CRparamSymbol}:                    {_CShiftAction, _CState333, ""},
	{_CState311, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState311, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState311, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState311, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState311, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState311, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState311, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState311, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState311, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState311, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState311, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState311, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState311, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState311, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState311, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState311, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState311, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState311, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState311, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState311, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState311, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState311, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState311, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState311, _CExpressionSymbol}:               {_CShiftAction, _CState334, ""},
	{_CState312, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState312, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState312, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState312, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState312, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState312, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState312, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState312, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState312, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState312, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState312, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState312, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState312, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState312, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState312, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState312, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState312, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState312, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState312, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState312, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState312, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState312, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState312, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState312, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState312, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState312, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState312, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState312, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState312, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState312, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState312, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState312, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState312, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState312, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState312, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState312, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState312, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState312, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState312, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState312, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState312, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState312, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState312, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState312, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState312, _CStatementSymbol}:                {_CShiftAction, _CState337, ""},
	{_CState312, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState312, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState312, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState312, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState312, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState312, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState315, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState315, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState315, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState315, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState315, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState315, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState315, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState315, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState315, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState315, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState315, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState315, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState315, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState315, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState315, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState315, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState315, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState315, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState315, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState315, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState315, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState315, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState315, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState315, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState315, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState315, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState315, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState315, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState315, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState315, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState315, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState315, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState315, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState315, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState315, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState315, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState315, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState315, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState315, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState315, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState315, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState315, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState315, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState315, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState315, _CStatementSymbol}:                {_CShiftAction, _CState338, ""},
	{_CState315, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState315, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState315, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState315, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState315, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState315, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState316, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState316, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState316, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState316, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState316, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState316, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState316, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState316, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState316, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState316, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState316, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState316, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState316, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState316, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState316, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState316, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState316, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState316, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState316, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState316, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState316, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState316, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState316, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState316, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState316, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState316, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState316, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState316, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState316, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState316, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState316, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState316, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState316, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState316, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState316, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState316, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState316, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState316, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState316, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState316, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState316, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState316, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState316, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState316, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState316, _CStatementSymbol}:                {_CShiftAction, _CState339, ""},
	{_CState316, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState316, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState316, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState316, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState316, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState316, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState317, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState317, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState317, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState317, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState317, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState317, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState317, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState317, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState317, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState317, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState317, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState317, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState317, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState317, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState317, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState317, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState134, ""},
	{_CState317, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState317, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState317, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState317, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState317, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState317, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState317, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState317, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState317, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState317, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState317, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState317, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState317, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState340, ""},
	{_CState319, CIdentifierSymbol}:                {_CShiftAction, _CState130, ""},
	{_CState319, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState319, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState319, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState319, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState319, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState319, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState319, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState319, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState319, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState319, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState319, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState319, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState319, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState319, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState319, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState319, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState319, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState319, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState319, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState319, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState319, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState319, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState319, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState319, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState319, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState319, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState319, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState319, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState319, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState341, ""},
	{_CState325, CRbraceSymbol}:                    {_CShiftAction, _CState342, ""},
	{_CState327, CRparamSymbol}:                    {_CShiftAction, _CState343, ""},
	{_CState332, CRparamSymbol}:                    {_CShiftAction, _CState344, ""},
	{_CState332, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState333, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState333, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState333, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState333, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState333, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState333, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState333, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState333, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState333, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState333, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState333, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState333, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState333, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState333, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState333, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState333, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState333, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState333, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState333, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState333, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState333, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState333, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState333, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState333, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState333, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState333, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState333, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState333, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState333, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState333, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState333, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState333, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState333, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState333, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState333, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState333, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState333, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState333, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState333, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState333, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState333, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState333, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState333, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState333, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState333, _CStatementSymbol}:                {_CShiftAction, _CState345, ""},
	{_CState333, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState333, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState333, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState333, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState333, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState333, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState334, CRparamSymbol}:                    {_CShiftAction, _CState346, ""},
	{_CState334, CCommaSymbol}:                     {_CShiftAction, _CState186, ""},
	{_CState335, CElseSymbol}:                      {_CShiftAction, _CState347, ""},
	{_CState336, CElseSymbol}:                      {_CShiftAction, _CState347, ""},
	{_CState337, CElseSymbol}:                      {_CShiftAction, _CState347, ""},
	{_CState344, CSemicolonSymbol}:                 {_CShiftAction, _CState348, ""},
	{_CState346, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState346, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState346, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState346, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState346, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState346, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState346, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState346, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState346, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState346, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState346, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState346, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState346, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState346, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState346, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState346, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState346, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState346, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState346, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState346, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState346, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState346, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState346, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState346, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState346, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState346, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState346, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState346, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState346, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState346, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState346, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState346, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState346, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState346, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState346, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState346, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState346, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState346, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState346, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState346, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState346, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState346, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState346, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState346, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState346, _CStatementSymbol}:                {_CShiftAction, _CState349, ""},
	{_CState346, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState346, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState346, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState346, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState346, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState346, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState347, CIdentifierSymbol}:                {_CShiftAction, _CState85, ""},
	{_CState347, CConstantSymbol}:                  {_CShiftAction, _CState77, ""},
	{_CState347, CStringLiteralSymbol}:             {_CShiftAction, _CState96, ""},
	{_CState347, CSizeofSymbol}:                    {_CShiftAction, _CState95, ""},
	{_CState347, CIncOpSymbol}:                     {_CShiftAction, _CState87, ""},
	{_CState347, CDecOpSymbol}:                     {_CShiftAction, _CState79, ""},
	{_CState347, CCaseSymbol}:                      {_CShiftAction, _CState76, ""},
	{_CState347, CDefaultSymbol}:                   {_CShiftAction, _CState80, ""},
	{_CState347, CIfSymbol}:                        {_CShiftAction, _CState86, ""},
	{_CState347, CSwitchSymbol}:                    {_CShiftAction, _CState97, ""},
	{_CState347, CWhileSymbol}:                     {_CShiftAction, _CState99, ""},
	{_CState347, CDoSymbol}:                        {_CShiftAction, _CState81, ""},
	{_CState347, CForSymbol}:                       {_CShiftAction, _CState83, ""},
	{_CState347, CGotoSymbol}:                      {_CShiftAction, _CState84, ""},
	{_CState347, CContinueSymbol}:                  {_CShiftAction, _CState78, ""},
	{_CState347, CBreakSymbol}:                     {_CShiftAction, _CState75, ""},
	{_CState347, CReturnSymbol}:                    {_CShiftAction, _CState93, ""},
	{_CState347, CLparamSymbol}:                    {_CShiftAction, _CState88, ""},
	{_CState347, CLcurlSymbol}:                     {_CShiftAction, _CState49, ""},
	{_CState347, CSemicolonSymbol}:                 {_CShiftAction, _CState94, ""},
	{_CState347, CMulSymbol}:                       {_CShiftAction, _CState90, ""},
	{_CState347, CMinusSymbol}:                     {_CShiftAction, _CState89, ""},
	{_CState347, CPlusSymbol}:                      {_CShiftAction, _CState91, ""},
	{_CState347, CAndSymbol}:                       {_CShiftAction, _CState74, ""},
	{_CState347, CExclaimSymbol}:                   {_CShiftAction, _CState82, ""},
	{_CState347, CTildaSymbol}:                     {_CShiftAction, _CState98, ""},
	{_CState347, _CPrimaryExpressionSymbol}:        {_CShiftAction, _CState119, ""},
	{_CState347, _CPostfixExpressionSymbol}:        {_CShiftAction, _CState118, ""},
	{_CState347, _CUnaryExpressionSymbol}:          {_CShiftAction, _CState125, ""},
	{_CState347, _CUnaryOperatorSymbol}:            {_CShiftAction, _CState126, ""},
	{_CState347, _CCastExpressionSymbol}:           {_CShiftAction, _CState103, ""},
	{_CState347, _CMultiplicativeExpressionSymbol}: {_CShiftAction, _CState117, ""},
	{_CState347, _CAdditiveExpressionSymbol}:       {_CShiftAction, _CState100, ""},
	{_CState347, _CShiftExpressionSymbol}:          {_CShiftAction, _CState122, ""},
	{_CState347, _CRelationalExpressionSymbol}:     {_CShiftAction, _CState120, ""},
	{_CState347, _CEqualityExpressionSymbol}:       {_CShiftAction, _CState107, ""},
	{_CState347, _CAndExpressionSymbol}:            {_CShiftAction, _CState101, ""},
	{_CState347, _CExclusiveOrExpressionSymbol}:    {_CShiftAction, _CState108, ""},
	{_CState347, _CInclusiveOrExpressionSymbol}:    {_CShiftAction, _CState111, ""},
	{_CState347, _CLogicalAndExpressionSymbol}:     {_CShiftAction, _CState115, ""},
	{_CState347, _CLogicalOrExpressionSymbol}:      {_CShiftAction, _CState116, ""},
	{_CState347, _CConditionalExpressionSymbol}:    {_CShiftAction, _CState105, ""},
	{_CState347, _CAssignmentExpressionSymbol}:     {_CShiftAction, _CState102, ""},
	{_CState347, _CExpressionSymbol}:               {_CShiftAction, _CState109, ""},
	{_CState347, _CStatementSymbol}:                {_CShiftAction, _CState350, ""},
	{_CState347, _CLabeledStatementSymbol}:         {_CShiftAction, _CState114, ""},
	{_CState347, _CCompoundStatementSymbol}:        {_CShiftAction, _CState104, ""},
	{_CState347, _CExpressionStatementSymbol}:      {_CShiftAction, _CState110, ""},
	{_CState347, _CSelectionStatementSymbol}:       {_CShiftAction, _CState121, ""},
	{_CState347, _CIterationStatementSymbol}:       {_CShiftAction, _CState112, ""},
	{_CState347, _CJumpStatementSymbol}:            {_CShiftAction, _CState113, ""},
	{_CState2, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceDToStorageClassSpecifier},
	{_CState3, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceBToTypeSpecifier},
	{_CState4, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceAToTypeQualifier},
	{_CState5, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceGToTypeSpecifier},
	{_CState7, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceBToStorageClassSpecifier},
	{_CState8, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceFToTypeSpecifier},
	{_CState9, _CWildcardMarker}:                   {_CReduceAction, "", _CReduceAToDirectDeclarator},
	{_CState10, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceDToTypeSpecifier},
	{_CState11, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceEToTypeSpecifier},
	{_CState13, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToPointer},
	{_CState14, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceEToStorageClassSpecifier},
	{_CState15, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToTypeSpecifier},
	{_CState16, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceHToTypeSpecifier},
	{_CState17, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToStorageClassSpecifier},
	{_CState18, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToStructOrUnion},
	{_CState19, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToStorageClassSpecifier},
	{_CState20, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceLToTypeSpecifier},
	{_CState21, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToStructOrUnion},
	{_CState22, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceIToTypeSpecifier},
	{_CState23, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToTypeSpecifier},
	{_CState24, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToTypeQualifier},
	{_CState25, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToExternalDeclaration},
	{_CState28, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToDeclarator},
	{_CState29, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceKToTypeSpecifier},
	{_CState30, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToTranslationUnit},
	{_CState31, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToExternalDeclaration},
	{_CState33, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToDeclarationSpecifiers},
	{_CState35, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceJToTypeSpecifier},
	{_CState36, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceEToDeclarationSpecifiers},
	{_CState37, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToDeclarationSpecifiers},
	{_CState38, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToTranslationUnit},
	{_CState39, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToEnumSpecifier},
	{_CState42, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToPointer},
	{_CState43, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToTypeQualifierList},
	{_CState44, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToPointer},
	{_CState45, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToDeclaration},
	{_CState46, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToInitDeclarator},
	{_CState47, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToInitDeclaratorList},
	{_CState50, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceDToFunctionDefinition},
	{_CState51, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToDeclarationList},
	{_CState56, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToDeclarator},
	{_CState57, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToDeclarationSpecifiers},
	{_CState58, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToStructOrUnionSpecifier},
	{_CState60, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceFToDeclarationSpecifiers},
	{_CState61, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceDToDeclarationSpecifiers},
	{_CState63, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToEnumerator},
	{_CState64, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToEnumeratorList},
	{_CState66, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToDirectDeclarator},
	{_CState67, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceDToPointer},
	{_CState68, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToTypeQualifierList},
	{_CState70, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToFunctionDefinition},
	{_CState73, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToDeclaration},
	{_CState74, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToUnaryOperator},
	{_CState77, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToPrimaryExpression},
	{_CState82, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceFToUnaryOperator},
	{_CState85, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToPrimaryExpression},
	{_CState89, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceDToUnaryOperator},
	{_CState90, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceBToUnaryOperator},
	{_CState91, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToUnaryOperator},
	{_CState92, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToCompoundStatement},
	{_CState94, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceAToExpressionStatement},
	{_CState96, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceCToPrimaryExpression},
	{_CState98, _CWildcardMarker}:                  {_CReduceAction, "", _CReduceEToUnaryOperator},
	{_CState100, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToShiftExpression},
	{_CState101, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToExclusiveOrExpression},
	{_CState102, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToExpression},
	{_CState103, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToMultiplicativeExpression},
	{_CState104, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStatement},
	{_CState105, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToAssignmentExpression},
	{_CState107, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToAndExpression},
	{_CState108, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToInclusiveOrExpression},
	{_CState110, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToStatement},
	{_CState111, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToLogicalAndExpression},
	{_CState112, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToStatement},
	{_CState113, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToStatement},
	{_CState114, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStatement},
	{_CState115, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToLogicalOrExpression},
	{_CState116, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToConditionalExpression},
	{_CState117, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToAdditiveExpression},
	{_CState118, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToUnaryExpression},
	{_CState119, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToPostfixExpression},
	{_CState120, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToEqualityExpression},
	{_CState121, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToStatement},
	{_CState122, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToRelationalExpression},
	{_CState123, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStatementList},
	{_CState125, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToCastExpression},
	{_CState127, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToFunctionDefinition},
	{_CState128, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToDeclarationList},
	{_CState129, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToInitDeclarator},
	{_CState130, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToPrimaryExpression},
	{_CState131, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToDirectDeclarator},
	{_CState132, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToConstantExpression},
	{_CState134, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToCastExpression},
	{_CState135, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToIdentifierList},
	{_CState136, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceGToDirectDeclarator},
	{_CState137, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToParameterDeclaration},
	{_CState139, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToParameterList},
	{_CState140, CRparamSymbol}:                    {_CReduceAction, "", _CReduceAToParameterTypeList},
	{_CState144, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStructDeclarationList},
	{_CState146, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToSpecifierQualifierList},
	{_CState147, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToSpecifierQualifierList},
	{_CState151, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToEnumSpecifier},
	{_CState153, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToInitializer},
	{_CState154, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToInitDeclarator},
	{_CState155, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToFunctionDefinition},
	{_CState156, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToInitDeclaratorList},
	{_CState157, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToJumpStatement},
	{_CState159, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToJumpStatement},
	{_CState161, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToUnaryExpression},
	{_CState168, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToUnaryExpression},
	{_CState170, CRparamSymbol}:                    {_CReduceAction, "", _CReduceAToTypeName},
	{_CState172, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToJumpStatement},
	{_CState175, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToUnaryExpression},
	{_CState181, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToCompoundStatement},
	{_CState187, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToExpressionStatement},
	{_CState195, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceHToPostfixExpression},
	{_CState197, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceGToPostfixExpression},
	{_CState207, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToCompoundStatement},
	{_CState208, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStatementList},
	{_CState209, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToAssignmentOperator},
	{_CState210, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceIToAssignmentOperator},
	{_CState211, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToAssignmentOperator},
	{_CState212, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToAssignmentOperator},
	{_CState213, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceGToAssignmentOperator},
	{_CState214, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToAssignmentOperator},
	{_CState215, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToAssignmentOperator},
	{_CState216, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceKToAssignmentOperator},
	{_CState217, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceHToAssignmentOperator},
	{_CState218, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToAssignmentOperator},
	{_CState219, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceJToAssignmentOperator},
	{_CState221, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToUnaryExpression},
	{_CState222, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToDirectDeclarator},
	{_CState225, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToParameterDeclaration},
	{_CState226, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToParameterDeclaration},
	{_CState227, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToAbstractDeclarator},
	{_CState228, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToAbstractDeclarator},
	{_CState230, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToDirectDeclarator},
	{_CState232, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToDirectDeclarator},
	{_CState235, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStructDeclarator},
	{_CState236, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStructDeclaratorList},
	{_CState238, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStructOrUnionSpecifier},
	{_CState239, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStructDeclarationList},
	{_CState240, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToSpecifierQualifierList},
	{_CState241, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToSpecifierQualifierList},
	{_CState242, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToEnumSpecifier},
	{_CState243, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToEnumerator},
	{_CState244, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToEnumeratorList},
	{_CState245, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToInitializerList},
	{_CState248, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToLabeledStatement},
	{_CState251, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToJumpStatement},
	{_CState252, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToLabeledStatement},
	{_CState254, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToPrimaryExpression},
	{_CState256, CRparamSymbol}:                    {_CReduceAction, "", _CReduceBToTypeName},
	{_CState257, CRparamSymbol}:                    {_CReduceAction, "", _CReduceAToAbstractDeclarator},
	{_CState259, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToJumpStatement},
	{_CState263, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToAdditiveExpression},
	{_CState264, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToAdditiveExpression},
	{_CState265, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToAndExpression},
	{_CState266, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToCompoundStatement},
	{_CState267, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToEqualityExpression},
	{_CState268, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToEqualityExpression},
	{_CState269, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToExclusiveOrExpression},
	{_CState270, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToExpression},
	{_CState271, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToInclusiveOrExpression},
	{_CState272, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToLogicalAndExpression},
	{_CState273, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToLogicalOrExpression},
	{_CState275, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToMultiplicativeExpression},
	{_CState276, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToMultiplicativeExpression},
	{_CState277, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToMultiplicativeExpression},
	{_CState278, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToPostfixExpression},
	{_CState280, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToPostfixExpression},
	{_CState282, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToArgumentExpressionList},
	{_CState283, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToPostfixExpression},
	{_CState284, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToRelationalExpression},
	{_CState285, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToRelationalExpression},
	{_CState286, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToRelationalExpression},
	{_CState287, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToRelationalExpression},
	{_CState288, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToShiftExpression},
	{_CState289, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToShiftExpression},
	{_CState290, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToAssignmentExpression},
	{_CState291, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToDirectAbstractDeclarator},
	{_CState293, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToDirectAbstractDeclarator},
	{_CState298, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToAbstractDeclarator},
	{_CState299, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToIdentifierList},
	{_CState300, CRparamSymbol}:                    {_CReduceAction, "", _CReduceBToParameterTypeList},
	{_CState301, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToParameterList},
	{_CState302, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStructOrUnionSpecifier},
	{_CState303, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStructDeclarator},
	{_CState306, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToStructDeclaration},
	{_CState308, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToInitializer},
	{_CState309, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToLabeledStatement},
	{_CState313, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToCastExpression},
	{_CState314, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceFToUnaryExpression},
	{_CState318, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToPostfixExpression},
	{_CState320, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToPostfixExpression},
	{_CState321, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToDirectAbstractDeclarator},
	{_CState322, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToDirectAbstractDeclarator},
	{_CState323, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceGToDirectAbstractDeclarator},
	{_CState324, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToDirectAbstractDeclarator},
	{_CState326, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceHToDirectAbstractDeclarator},
	{_CState328, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToStructDeclarator},
	{_CState329, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToStructDeclaratorList},
	{_CState330, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToInitializer},
	{_CState331, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToInitializerList},
	{_CState335, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CAndSymbol}:                       {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CBreakSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CCaseSymbol}:                      {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CConstantSymbol}:                  {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CContinueSymbol}:                  {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CDecOpSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CDefaultSymbol}:                   {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CDoSymbol}:                        {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CElseSymbol}:                      {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CExclaimSymbol}:                   {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CForSymbol}:                       {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CGotoSymbol}:                      {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CIdentifierSymbol}:                {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CIfSymbol}:                        {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CIncOpSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CLcurlSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CLparamSymbol}:                    {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CMinusSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CMulSymbol}:                       {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CPlusSymbol}:                      {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CRcurlSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CReturnSymbol}:                    {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CSemicolonSymbol}:                 {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CSizeofSymbol}:                    {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CStringLiteralSymbol}:             {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CSwitchSymbol}:                    {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CTildaSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState336, CWhileSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState337, CElseSymbol}:                      {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState337, CWhileSymbol}:                     {_CReduceAction, "", _CReduceAToSelectionStatement},
	{_CState338, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToSelectionStatement},
	{_CState339, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceAToIterationStatement},
	{_CState340, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToConditionalExpression},
	{_CState341, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToArgumentExpressionList},
	{_CState342, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceEToDirectAbstractDeclarator},
	{_CState343, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceIToDirectAbstractDeclarator},
	{_CState345, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceCToIterationStatement},
	{_CState348, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToIterationStatement},
	{_CState349, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceDToIterationStatement},
	{_CState350, _CWildcardMarker}:                 {_CReduceAction, "", _CReduceBToSelectionStatement},
}

var _CExpectedTerminals = map[_CStateId][]CSymbolId{
	_CState0:   []CSymbolId{CIdentifierSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CMulSymbol},
	_CState1:   []CSymbolId{CIdentifierSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CMulSymbol, _CEndMarker},
	_CState6:   []CSymbolId{CIdentifierSymbol, CLcurlSymbol},
	_CState12:  []CSymbolId{CIdentifierSymbol, CLparamSymbol, CMulSymbol},
	_CState13:  []CSymbolId{CConstSymbol, CVolatileSymbol, CMulSymbol},
	_CState26:  []CSymbolId{CIdentifierSymbol, CLparamSymbol, CSemicolonSymbol, CMulSymbol},
	_CState27:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLcurlSymbol},
	_CState28:  []CSymbolId{CLparamSymbol, CLbraceSymbol},
	_CState32:  []CSymbolId{CIdentifierSymbol, CLparamSymbol},
	_CState33:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState34:  []CSymbolId{CIdentifierSymbol, CLcurlSymbol},
	_CState36:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState37:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState39:  []CSymbolId{CLcurlSymbol},
	_CState40:  []CSymbolId{CIdentifierSymbol},
	_CState41:  []CSymbolId{CRparamSymbol},
	_CState44:  []CSymbolId{CConstSymbol, CVolatileSymbol, CMulSymbol},
	_CState46:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLcurlSymbol, CEqSymbol},
	_CState48:  []CSymbolId{CSemicolonSymbol, CCommaSymbol},
	_CState49:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CRcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState52:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLcurlSymbol},
	_CState53:  []CSymbolId{CIdentifierSymbol, CLparamSymbol, CSemicolonSymbol, CMulSymbol},
	_CState54:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CRbraceSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState55:  []CSymbolId{CIdentifierSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CRparamSymbol},
	_CState56:  []CSymbolId{CLparamSymbol, CLbraceSymbol},
	_CState58:  []CSymbolId{CLcurlSymbol},
	_CState59:  []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState62:  []CSymbolId{CIdentifierSymbol},
	_CState63:  []CSymbolId{CEqSymbol},
	_CState65:  []CSymbolId{CRcurlSymbol, CCommaSymbol},
	_CState69:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CLcurlSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState71:  []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLcurlSymbol},
	_CState72:  []CSymbolId{CIdentifierSymbol, CLparamSymbol, CMulSymbol},
	_CState75:  []CSymbolId{CSemicolonSymbol},
	_CState76:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState78:  []CSymbolId{CSemicolonSymbol},
	_CState79:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState80:  []CSymbolId{CColonSymbol},
	_CState81:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState83:  []CSymbolId{CLparamSymbol},
	_CState84:  []CSymbolId{CIdentifierSymbol},
	_CState85:  []CSymbolId{CColonSymbol},
	_CState86:  []CSymbolId{CLparamSymbol},
	_CState87:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState88:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState93:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState95:  []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState97:  []CSymbolId{CLparamSymbol},
	_CState99:  []CSymbolId{CLparamSymbol},
	_CState100: []CSymbolId{CMinusSymbol, CPlusSymbol},
	_CState101: []CSymbolId{CAndSymbol},
	_CState106: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CRcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState107: []CSymbolId{CEqOpSymbol, CNeOpSymbol},
	_CState108: []CSymbolId{CHatSymbol},
	_CState109: []CSymbolId{CSemicolonSymbol, CCommaSymbol},
	_CState111: []CSymbolId{COrSymbol},
	_CState115: []CSymbolId{CAndOpSymbol},
	_CState116: []CSymbolId{COrOpSymbol, CQuestionSymbol},
	_CState117: []CSymbolId{CMulSymbol, CDivSymbol, CModSymbol},
	_CState118: []CSymbolId{CPtrOpSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CLbraceSymbol, CDotSymbol},
	_CState120: []CSymbolId{CLeOpSymbol, CGeOpSymbol, CLtSymbol, CGtSymbol},
	_CState122: []CSymbolId{CLeftOpSymbol, CRightOpSymbol},
	_CState124: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CRcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState125: []CSymbolId{CMulAssignSymbol, CDivAssignSymbol, CModAssignSymbol, CAddAssignSymbol, CSubAssignSymbol, CLeftAssignSymbol, CRightAssignSymbol, CAndAssignSymbol, CXorAssignSymbol, COrAssignSymbol, CEqSymbol},
	_CState126: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState129: []CSymbolId{CEqSymbol},
	_CState133: []CSymbolId{CRbraceSymbol},
	_CState137: []CSymbolId{CIdentifierSymbol, CLparamSymbol, CLbraceSymbol, CMulSymbol},
	_CState138: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState140: []CSymbolId{CCommaSymbol, CRparamSymbol},
	_CState141: []CSymbolId{CRparamSymbol},
	_CState142: []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState143: []CSymbolId{CIdentifierSymbol, CLparamSymbol, CColonSymbol, CMulSymbol},
	_CState145: []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CRcurlSymbol},
	_CState146: []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState147: []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol},
	_CState148: []CSymbolId{CRcurlSymbol, CCommaSymbol},
	_CState149: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState150: []CSymbolId{CIdentifierSymbol},
	_CState152: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CLcurlSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState158: []CSymbolId{CColonSymbol},
	_CState160: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState162: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState163: []CSymbolId{CWhileSymbol},
	_CState164: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState165: []CSymbolId{CSemicolonSymbol},
	_CState166: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState167: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState169: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState170: []CSymbolId{CLparamSymbol, CLbraceSymbol, CMulSymbol, CRparamSymbol},
	_CState171: []CSymbolId{CRparamSymbol},
	_CState173: []CSymbolId{CSemicolonSymbol, CCommaSymbol},
	_CState174: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState176: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState177: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState178: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState179: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState180: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState182: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CRcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState183: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState184: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState185: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState186: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState188: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState189: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState190: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState191: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState192: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState193: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState194: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState196: []CSymbolId{CIdentifierSymbol},
	_CState198: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState199: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CRparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState200: []CSymbolId{CIdentifierSymbol},
	_CState201: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState202: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState203: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState204: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState205: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState206: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState220: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState223: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CRbraceSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState224: []CSymbolId{CIdentifierSymbol, CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CRparamSymbol, CLbraceSymbol, CMulSymbol},
	_CState227: []CSymbolId{CLparamSymbol, CLbraceSymbol},
	_CState228: []CSymbolId{CIdentifierSymbol, CLparamSymbol, CLbraceSymbol},
	_CState229: []CSymbolId{CIdentifierSymbol},
	_CState231: []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CEllipsisSymbol},
	_CState233: []CSymbolId{CTypeNameSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CRcurlSymbol},
	_CState234: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState235: []CSymbolId{CColonSymbol},
	_CState237: []CSymbolId{CSemicolonSymbol, CCommaSymbol},
	_CState246: []CSymbolId{CRcurlSymbol, CCommaSymbol},
	_CState247: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState249: []CSymbolId{CLparamSymbol},
	_CState250: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState253: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState255: []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CLparamSymbol, CRparamSymbol, CLbraceSymbol, CMulSymbol},
	_CState256: []CSymbolId{CRparamSymbol},
	_CState257: []CSymbolId{CLparamSymbol, CLbraceSymbol, CRparamSymbol},
	_CState258: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState260: []CSymbolId{CRparamSymbol},
	_CState261: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState262: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState263: []CSymbolId{CMulSymbol, CDivSymbol, CModSymbol},
	_CState264: []CSymbolId{CMulSymbol, CDivSymbol, CModSymbol},
	_CState265: []CSymbolId{CEqOpSymbol, CNeOpSymbol},
	_CState267: []CSymbolId{CLeOpSymbol, CGeOpSymbol, CLtSymbol, CGtSymbol},
	_CState268: []CSymbolId{CLeOpSymbol, CGeOpSymbol, CLtSymbol, CGtSymbol},
	_CState269: []CSymbolId{CAndSymbol},
	_CState271: []CSymbolId{CHatSymbol},
	_CState272: []CSymbolId{COrSymbol},
	_CState273: []CSymbolId{CAndOpSymbol},
	_CState274: []CSymbolId{CColonSymbol, CCommaSymbol},
	_CState279: []CSymbolId{CRbraceSymbol, CCommaSymbol},
	_CState281: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState284: []CSymbolId{CLeftOpSymbol, CRightOpSymbol},
	_CState285: []CSymbolId{CLeftOpSymbol, CRightOpSymbol},
	_CState286: []CSymbolId{CLeftOpSymbol, CRightOpSymbol},
	_CState287: []CSymbolId{CLeftOpSymbol, CRightOpSymbol},
	_CState288: []CSymbolId{CMinusSymbol, CPlusSymbol},
	_CState289: []CSymbolId{CMinusSymbol, CPlusSymbol},
	_CState292: []CSymbolId{CRbraceSymbol},
	_CState294: []CSymbolId{CRparamSymbol},
	_CState295: []CSymbolId{CRparamSymbol},
	_CState296: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CRbraceSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState297: []CSymbolId{CTypeNameSymbol, CTypedefSymbol, CExternSymbol, CStaticSymbol, CAutoSymbol, CRegisterSymbol, CCharSymbol, CShortSymbol, CIntSymbol, CLongSymbol, CSignedSymbol, CUnsignedSymbol, CFloatSymbol, CDoubleSymbol, CConstSymbol, CVolatileSymbol, CVoidSymbol, CStructSymbol, CUnionSymbol, CEnumSymbol, CRparamSymbol},
	_CState298: []CSymbolId{CLparamSymbol, CLbraceSymbol},
	_CState300: []CSymbolId{CRparamSymbol},
	_CState304: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState305: []CSymbolId{CIdentifierSymbol, CLparamSymbol, CColonSymbol, CMulSymbol},
	_CState307: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CLcurlSymbol, CRcurlSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState310: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState311: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CRparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState312: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState315: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState316: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState317: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState319: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CLparamSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState325: []CSymbolId{CRbraceSymbol},
	_CState327: []CSymbolId{CRparamSymbol},
	_CState332: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState333: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState334: []CSymbolId{CRparamSymbol, CCommaSymbol},
	_CState335: []CSymbolId{CElseSymbol},
	_CState336: []CSymbolId{CElseSymbol, CAndSymbol, CBreakSymbol, CCaseSymbol, CConstantSymbol, CContinueSymbol, CDecOpSymbol, CDefaultSymbol, CDoSymbol, CElseSymbol, CExclaimSymbol, CForSymbol, CGotoSymbol, CIdentifierSymbol, CIfSymbol, CIncOpSymbol, CLcurlSymbol, CLparamSymbol, CMinusSymbol, CMulSymbol, CPlusSymbol, CRcurlSymbol, CReturnSymbol, CSemicolonSymbol, CSizeofSymbol, CStringLiteralSymbol, CSwitchSymbol, CTildaSymbol, CWhileSymbol},
	_CState337: []CSymbolId{CElseSymbol, CElseSymbol, CWhileSymbol},
	_CState344: []CSymbolId{CSemicolonSymbol},
	_CState346: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
	_CState347: []CSymbolId{CIdentifierSymbol, CConstantSymbol, CStringLiteralSymbol, CSizeofSymbol, CIncOpSymbol, CDecOpSymbol, CCaseSymbol, CDefaultSymbol, CIfSymbol, CSwitchSymbol, CWhileSymbol, CDoSymbol, CForSymbol, CGotoSymbol, CContinueSymbol, CBreakSymbol, CReturnSymbol, CLparamSymbol, CLcurlSymbol, CSemicolonSymbol, CMulSymbol, CMinusSymbol, CPlusSymbol, CAndSymbol, CExclaimSymbol, CTildaSymbol},
}

/*
Parser Debug States:
  State 0:
    Kernel Items:
      #accept:^.translation_unit
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
      #accept:^ translation_unit.,$
      translation_unit:translation_unit.external_declaration
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
      storage_class_specifier:AUTO.,*
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 3:
    Kernel Items:
      type_specifier:CHAR.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 4:
    Kernel Items:
      type_qualifier:CONST.,*
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 5:
    Kernel Items:
      type_specifier:DOUBLE.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      enum_specifier:ENUM.IDENTIFIER
      enum_specifier:ENUM.IDENTIFIER LCURL enumerator_list RCURL
      enum_specifier:ENUM.LCURL enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 39
      LCURL -> State 40

  State 7:
    Kernel Items:
      storage_class_specifier:EXTERN.,*
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      type_specifier:FLOAT.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      direct_declarator:IDENTIFIER.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      type_specifier:INT.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      type_specifier:LONG.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      direct_declarator:LPARAM.declarator RPARAM
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
      pointer:MUL.,*
      pointer:MUL.pointer
      pointer:MUL.type_qualifier_list
      pointer:MUL.type_qualifier_list pointer
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
      storage_class_specifier:REGISTER.,*
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      type_specifier:SHORT.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      type_specifier:SIGNED.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      storage_class_specifier:STATIC.,*
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      struct_or_union:STRUCT.,*
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      storage_class_specifier:TYPEDEF.,*
    Reduce:
      * -> [storage_class_specifier]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      type_specifier:TYPE_NAME.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      struct_or_union:UNION.,*
    Reduce:
      * -> [struct_or_union]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      type_specifier:UNSIGNED.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      type_specifier:VOID.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      type_qualifier:VOLATILE.,*
    Reduce:
      * -> [type_qualifier]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      external_declaration:declaration.,*
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      declaration:declaration_specifiers.SEMICOLON
      declaration:declaration_specifiers.init_declarator_list SEMICOLON
      function_definition:declaration_specifiers.declarator compound_statement
      function_definition:declaration_specifiers.declarator declaration_list compound_statement
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
      function_definition:declarator.compound_statement
      function_definition:declarator.declaration_list compound_statement
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
      declarator:direct_declarator.,*
      direct_declarator:direct_declarator.LBRACE RBRACE
      direct_declarator:direct_declarator.LBRACE constant_expression RBRACE
      direct_declarator:direct_declarator.LPARAM RPARAM
      direct_declarator:direct_declarator.LPARAM identifier_list RPARAM
      direct_declarator:direct_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [declarator]
    Goto:
      LPARAM -> State 55
      LBRACE -> State 54

  State 29:
    Kernel Items:
      type_specifier:enum_specifier.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      translation_unit:external_declaration.,*
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      external_declaration:function_definition.,*
    Reduce:
      * -> [external_declaration]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      declarator:pointer.direct_declarator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 9
      LPARAM -> State 12
      direct_declarator -> State 56

  State 33:
    Kernel Items:
      declaration_specifiers:storage_class_specifier.,*
      declaration_specifiers:storage_class_specifier.declaration_specifiers
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
      struct_or_union_specifier:struct_or_union.IDENTIFIER
      struct_or_union_specifier:struct_or_union.IDENTIFIER LCURL struct_declaration_list RCURL
      struct_or_union_specifier:struct_or_union.LCURL struct_declaration_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 58
      LCURL -> State 59

  State 35:
    Kernel Items:
      type_specifier:struct_or_union_specifier.,*
    Reduce:
      * -> [type_specifier]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      declaration_specifiers:type_qualifier.,*
      declaration_specifiers:type_qualifier.declaration_specifiers
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
      declaration_specifiers:type_specifier.,*
      declaration_specifiers:type_specifier.declaration_specifiers
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
      translation_unit:translation_unit external_declaration.,*
    Reduce:
      * -> [translation_unit]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      enum_specifier:ENUM IDENTIFIER.,*
      enum_specifier:ENUM IDENTIFIER.LCURL enumerator_list RCURL
    Reduce:
      * -> [enum_specifier]
    Goto:
      LCURL -> State 62

  State 40:
    Kernel Items:
      enum_specifier:ENUM LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator_list -> State 65
      enumerator -> State 64

  State 41:
    Kernel Items:
      direct_declarator:LPARAM declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 66

  State 42:
    Kernel Items:
      pointer:MUL pointer.,*
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      type_qualifier_list:type_qualifier.,*
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      pointer:MUL type_qualifier_list.,*
      pointer:MUL type_qualifier_list.pointer
      type_qualifier_list:type_qualifier_list.type_qualifier
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
      declaration:declaration_specifiers SEMICOLON.,*
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      function_definition:declaration_specifiers declarator.compound_statement
      function_definition:declaration_specifiers declarator.declaration_list compound_statement
      init_declarator:declarator.,*
      init_declarator:declarator.EQ initializer
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
      init_declarator_list:init_declarator.,*
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      declaration:declaration_specifiers init_declarator_list.SEMICOLON
      init_declarator_list:init_declarator_list.COMMA init_declarator
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 73
      COMMA -> State 72

  State 49:
    Kernel Items:
      compound_statement:LCURL.RCURL
      compound_statement:LCURL.declaration_list RCURL
      compound_statement:LCURL.declaration_list statement_list RCURL
      compound_statement:LCURL.statement_list RCURL
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
      function_definition:declarator compound_statement.,*
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      declaration_list:declaration.,*
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      declaration_list:declaration_list.declaration
      function_definition:declarator declaration_list.compound_statement
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
      declaration:declaration_specifiers.SEMICOLON
      declaration:declaration_specifiers.init_declarator_list SEMICOLON
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
      direct_declarator:direct_declarator LBRACE.RBRACE
      direct_declarator:direct_declarator LBRACE.constant_expression RBRACE
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
      direct_declarator:direct_declarator LPARAM.RPARAM
      direct_declarator:direct_declarator LPARAM.identifier_list RPARAM
      direct_declarator:direct_declarator LPARAM.parameter_type_list RPARAM
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
      declarator:pointer direct_declarator.,*
      direct_declarator:direct_declarator.LBRACE RBRACE
      direct_declarator:direct_declarator.LBRACE constant_expression RBRACE
      direct_declarator:direct_declarator.LPARAM RPARAM
      direct_declarator:direct_declarator.LPARAM identifier_list RPARAM
      direct_declarator:direct_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [declarator]
    Goto:
      LPARAM -> State 55
      LBRACE -> State 54

  State 57:
    Kernel Items:
      declaration_specifiers:storage_class_specifier declaration_specifiers.,*
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      struct_or_union_specifier:struct_or_union IDENTIFIER.,*
      struct_or_union_specifier:struct_or_union IDENTIFIER.LCURL struct_declaration_list RCURL
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      LCURL -> State 142

  State 59:
    Kernel Items:
      struct_or_union_specifier:struct_or_union LCURL.struct_declaration_list RCURL
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
      declaration_specifiers:type_qualifier declaration_specifiers.,*
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      declaration_specifiers:type_specifier declaration_specifiers.,*
    Reduce:
      * -> [declaration_specifiers]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      enum_specifier:ENUM IDENTIFIER LCURL.enumerator_list RCURL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator_list -> State 148
      enumerator -> State 64

  State 63:
    Kernel Items:
      enumerator:IDENTIFIER.,*
      enumerator:IDENTIFIER.EQ constant_expression
    Reduce:
      * -> [enumerator]
    Goto:
      EQ -> State 149

  State 64:
    Kernel Items:
      enumerator_list:enumerator.,*
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      enum_specifier:ENUM LCURL enumerator_list.RCURL
      enumerator_list:enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 151
      COMMA -> State 150

  State 66:
    Kernel Items:
      direct_declarator:LPARAM declarator RPARAM.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      pointer:MUL type_qualifier_list pointer.,*
    Reduce:
      * -> [pointer]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      type_qualifier_list:type_qualifier_list type_qualifier.,*
    Reduce:
      * -> [type_qualifier_list]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      init_declarator:declarator EQ.initializer
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
      function_definition:declaration_specifiers declarator compound_statement.,*
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      declaration_list:declaration_list.declaration
      function_definition:declaration_specifiers declarator declaration_list.compound_statement
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
      init_declarator_list:init_declarator_list COMMA.init_declarator
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
      declaration:declaration_specifiers init_declarator_list SEMICOLON.,*
    Reduce:
      * -> [declaration]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      unary_operator:AND.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      jump_statement:BREAK.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 157

  State 76:
    Kernel Items:
      labeled_statement:CASE.constant_expression COLON statement
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
      primary_expression:CONSTANT.,*
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      jump_statement:CONTINUE.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 159

  State 79:
    Kernel Items:
      unary_expression:DEC_OP.unary_expression
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
      labeled_statement:DEFAULT.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 162

  State 81:
    Kernel Items:
      iteration_statement:DO.statement WHILE LPARAM expression RPARAM SEMICOLON
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
      unary_operator:EXCLAIM.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      iteration_statement:FOR.LPARAM expression_statement expression_statement RPARAM statement
      iteration_statement:FOR.LPARAM expression_statement expression_statement expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 164

  State 84:
    Kernel Items:
      jump_statement:GOTO.IDENTIFIER SEMICOLON
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 165

  State 85:
    Kernel Items:
      labeled_statement:IDENTIFIER.COLON statement
      primary_expression:IDENTIFIER.,*
    Reduce:
      * -> [primary_expression]
    Goto:
      COLON -> State 166

  State 86:
    Kernel Items:
      selection_statement:IF.LPARAM expression RPARAM statement
      selection_statement:IF.LPARAM expression RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 167

  State 87:
    Kernel Items:
      unary_expression:INC_OP.unary_expression
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
      cast_expression:LPARAM.type_name RPARAM cast_expression
      primary_expression:LPARAM.expression RPARAM
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
      unary_operator:MINUS.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      unary_operator:MUL.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      unary_operator:PLUS.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      compound_statement:LCURL RCURL.,*
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      jump_statement:RETURN.SEMICOLON
      jump_statement:RETURN.expression SEMICOLON
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
      expression_statement:SEMICOLON.,*
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      unary_expression:SIZEOF.LPARAM type_name RPARAM
      unary_expression:SIZEOF.unary_expression
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
      primary_expression:STRING_LITERAL.,*
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      selection_statement:SWITCH.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 176

  State 98:
    Kernel Items:
      unary_operator:TILDA.,*
    Reduce:
      * -> [unary_operator]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      iteration_statement:WHILE.LPARAM expression RPARAM statement
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 177

  State 100:
    Kernel Items:
      additive_expression:additive_expression.MINUS multiplicative_expression
      additive_expression:additive_expression.PLUS multiplicative_expression
      shift_expression:additive_expression.,*
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 101:
    Kernel Items:
      and_expression:and_expression.AND equality_expression
      exclusive_or_expression:and_expression.,*
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 180

  State 102:
    Kernel Items:
      expression:assignment_expression.,*
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      multiplicative_expression:cast_expression.,*
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      statement:compound_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      assignment_expression:conditional_expression.,*
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      compound_statement:LCURL declaration_list.RCURL
      compound_statement:LCURL declaration_list.statement_list RCURL
      declaration_list:declaration_list.declaration
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
      and_expression:equality_expression.,*
      equality_expression:equality_expression.EQ_OP relational_expression
      equality_expression:equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 183
      NE_OP -> State 184

  State 108:
    Kernel Items:
      exclusive_or_expression:exclusive_or_expression.HAT and_expression
      inclusive_or_expression:exclusive_or_expression.,*
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 185

  State 109:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      expression_statement:expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 187
      COMMA -> State 186

  State 110:
    Kernel Items:
      statement:expression_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      inclusive_or_expression:inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression:inclusive_or_expression.,*
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 188

  State 112:
    Kernel Items:
      statement:iteration_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      statement:jump_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      statement:labeled_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      logical_and_expression:logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression:logical_and_expression.,*
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 189

  State 116:
    Kernel Items:
      conditional_expression:logical_or_expression.,*
      conditional_expression:logical_or_expression.QUESTION expression COLON conditional_expression
      logical_or_expression:logical_or_expression.OR_OP logical_and_expression
    Reduce:
      * -> [conditional_expression]
    Goto:
      OR_OP -> State 190
      QUESTION -> State 191

  State 117:
    Kernel Items:
      additive_expression:multiplicative_expression.,*
      multiplicative_expression:multiplicative_expression.DIV cast_expression
      multiplicative_expression:multiplicative_expression.MOD cast_expression
      multiplicative_expression:multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 118:
    Kernel Items:
      postfix_expression:postfix_expression.DEC_OP
      postfix_expression:postfix_expression.DOT IDENTIFIER
      postfix_expression:postfix_expression.INC_OP
      postfix_expression:postfix_expression.LBRACE expression RBRACE
      postfix_expression:postfix_expression.LPARAM RPARAM
      postfix_expression:postfix_expression.LPARAM argument_expression_list RPARAM
      postfix_expression:postfix_expression.PTR_OP IDENTIFIER
      unary_expression:postfix_expression.,*
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
      postfix_expression:primary_expression.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      equality_expression:relational_expression.,*
      relational_expression:relational_expression.GE_OP shift_expression
      relational_expression:relational_expression.GT shift_expression
      relational_expression:relational_expression.LE_OP shift_expression
      relational_expression:relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 121:
    Kernel Items:
      statement:selection_statement.,*
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      relational_expression:shift_expression.,*
      shift_expression:shift_expression.LEFT_OP additive_expression
      shift_expression:shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 123:
    Kernel Items:
      statement_list:statement.,*
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      compound_statement:LCURL statement_list.RCURL
      statement_list:statement_list.statement
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
      assignment_expression:unary_expression.assignment_operator assignment_expression
      cast_expression:unary_expression.,*
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
      unary_expression:unary_operator.cast_expression
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
      function_definition:declarator declaration_list compound_statement.,*
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      declaration_list:declaration_list declaration.,*
    Reduce:
      * -> [declaration_list]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      init_declarator:declarator.,*
      init_declarator:declarator.EQ initializer
    Reduce:
      * -> [init_declarator]
    Goto:
      EQ -> State 69

  State 130:
    Kernel Items:
      primary_expression:IDENTIFIER.,*
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      direct_declarator:direct_declarator LBRACE RBRACE.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      constant_expression:conditional_expression.,*
    Reduce:
      * -> [constant_expression]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      direct_declarator:direct_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 222

  State 134:
    Kernel Items:
      cast_expression:unary_expression.,*
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      identifier_list:IDENTIFIER.,*
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      direct_declarator:direct_declarator LPARAM RPARAM.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      parameter_declaration:declaration_specifiers.,*
      parameter_declaration:declaration_specifiers.abstract_declarator
      parameter_declaration:declaration_specifiers.declarator
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
      direct_declarator:direct_declarator LPARAM identifier_list.RPARAM
      identifier_list:identifier_list.COMMA IDENTIFIER
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 230
      COMMA -> State 229

  State 139:
    Kernel Items:
      parameter_list:parameter_declaration.,*
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      parameter_list:parameter_list.COMMA parameter_declaration
      parameter_type_list:parameter_list.,RPARAM
      parameter_type_list:parameter_list.COMMA ELLIPSIS
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      COMMA -> State 231

  State 141:
    Kernel Items:
      direct_declarator:direct_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 232

  State 142:
    Kernel Items:
      struct_or_union_specifier:struct_or_union IDENTIFIER LCURL.struct_declaration_list RCURL
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
      struct_declaration:specifier_qualifier_list.struct_declarator_list SEMICOLON
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
      struct_declaration_list:struct_declaration.,*
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      struct_declaration_list:struct_declaration_list.struct_declaration
      struct_or_union_specifier:struct_or_union LCURL struct_declaration_list.RCURL
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
      specifier_qualifier_list:type_qualifier.,*
      specifier_qualifier_list:type_qualifier.specifier_qualifier_list
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
      specifier_qualifier_list:type_specifier.,*
      specifier_qualifier_list:type_specifier.specifier_qualifier_list
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
      enum_specifier:ENUM IDENTIFIER LCURL enumerator_list.RCURL
      enumerator_list:enumerator_list.COMMA enumerator
    Reduce:
      (nil)
    Goto:
      RCURL -> State 242
      COMMA -> State 150

  State 149:
    Kernel Items:
      enumerator:IDENTIFIER EQ.constant_expression
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
      enumerator_list:enumerator_list COMMA.enumerator
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63
      enumerator -> State 244

  State 151:
    Kernel Items:
      enum_specifier:ENUM LCURL enumerator_list RCURL.,*
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      initializer:LCURL.initializer_list COMMA RCURL
      initializer:LCURL.initializer_list RCURL
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
      initializer:assignment_expression.,*
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      init_declarator:declarator EQ initializer.,*
    Reduce:
      * -> [init_declarator]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      function_definition:declaration_specifiers declarator declaration_list compound_statement.,*
    Reduce:
      * -> [function_definition]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      init_declarator_list:init_declarator_list COMMA init_declarator.,*
    Reduce:
      * -> [init_declarator_list]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      jump_statement:BREAK SEMICOLON.,*
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      labeled_statement:CASE constant_expression.COLON statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 247

  State 159:
    Kernel Items:
      jump_statement:CONTINUE SEMICOLON.,*
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      primary_expression:LPARAM.expression RPARAM
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
      unary_expression:DEC_OP unary_expression.,*
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      labeled_statement:DEFAULT COLON.statement
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
      iteration_statement:DO statement.WHILE LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      WHILE -> State 249

  State 164:
    Kernel Items:
      iteration_statement:FOR LPARAM.expression_statement expression_statement RPARAM statement
      iteration_statement:FOR LPARAM.expression_statement expression_statement expression RPARAM statement
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
      jump_statement:GOTO IDENTIFIER.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 251

  State 166:
    Kernel Items:
      labeled_statement:IDENTIFIER COLON.statement
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
      selection_statement:IF LPARAM.expression RPARAM statement
      selection_statement:IF LPARAM.expression RPARAM statement ELSE statement
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
      unary_expression:INC_OP unary_expression.,*
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      primary_expression:LPARAM expression.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 254
      COMMA -> State 186

  State 170:
    Kernel Items:
      type_name:specifier_qualifier_list.,RPARAM
      type_name:specifier_qualifier_list.abstract_declarator
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
      cast_expression:LPARAM type_name.RPARAM cast_expression
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 258

  State 172:
    Kernel Items:
      jump_statement:RETURN SEMICOLON.,*
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      jump_statement:RETURN expression.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 259
      COMMA -> State 186

  State 174:
    Kernel Items:
      primary_expression:LPARAM.expression RPARAM
      unary_expression:SIZEOF LPARAM.type_name RPARAM
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
      unary_expression:SIZEOF unary_expression.,*
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      selection_statement:SWITCH LPARAM.expression RPARAM statement
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
      iteration_statement:WHILE LPARAM.expression RPARAM statement
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
      additive_expression:additive_expression MINUS.multiplicative_expression
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
      additive_expression:additive_expression PLUS.multiplicative_expression
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
      and_expression:and_expression AND.equality_expression
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
      compound_statement:LCURL declaration_list RCURL.,*
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      compound_statement:LCURL declaration_list statement_list.RCURL
      statement_list:statement_list.statement
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
      equality_expression:equality_expression EQ_OP.relational_expression
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
      equality_expression:equality_expression NE_OP.relational_expression
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
      exclusive_or_expression:exclusive_or_expression HAT.and_expression
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
      expression:expression COMMA.assignment_expression
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
      expression_statement:expression SEMICOLON.,*
    Reduce:
      * -> [expression_statement]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      inclusive_or_expression:inclusive_or_expression OR.exclusive_or_expression
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
      logical_and_expression:logical_and_expression AND_OP.inclusive_or_expression
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
      logical_or_expression:logical_or_expression OR_OP.logical_and_expression
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
      conditional_expression:logical_or_expression QUESTION.expression COLON conditional_expression
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
      multiplicative_expression:multiplicative_expression DIV.cast_expression
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
      multiplicative_expression:multiplicative_expression MOD.cast_expression
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
      multiplicative_expression:multiplicative_expression MUL.cast_expression
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
      postfix_expression:postfix_expression DEC_OP.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      postfix_expression:postfix_expression DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 278

  State 197:
    Kernel Items:
      postfix_expression:postfix_expression INC_OP.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      postfix_expression:postfix_expression LBRACE.expression RBRACE
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
      postfix_expression:postfix_expression LPARAM.RPARAM
      postfix_expression:postfix_expression LPARAM.argument_expression_list RPARAM
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
      postfix_expression:postfix_expression PTR_OP.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 283

  State 201:
    Kernel Items:
      relational_expression:relational_expression GE_OP.shift_expression
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
      relational_expression:relational_expression GT.shift_expression
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
      relational_expression:relational_expression LE_OP.shift_expression
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
      relational_expression:relational_expression LT.shift_expression
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
      shift_expression:shift_expression LEFT_OP.additive_expression
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
      shift_expression:shift_expression RIGHT_OP.additive_expression
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
      compound_statement:LCURL statement_list RCURL.,*
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      statement_list:statement_list statement.,*
    Reduce:
      * -> [statement_list]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      assignment_operator:ADD_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      assignment_operator:AND_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      assignment_operator:DIV_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      assignment_operator:EQ.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      assignment_operator:LEFT_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      assignment_operator:MOD_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      assignment_operator:MUL_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      assignment_operator:OR_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      assignment_operator:RIGHT_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      assignment_operator:SUB_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      assignment_operator:XOR_ASSIGN.,*
    Reduce:
      * -> [assignment_operator]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      assignment_expression:unary_expression assignment_operator.assignment_expression
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
      unary_expression:unary_operator cast_expression.,*
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      direct_declarator:direct_declarator LBRACE constant_expression RBRACE.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      direct_abstract_declarator:LBRACE.RBRACE
      direct_abstract_declarator:LBRACE.constant_expression RBRACE
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
      direct_abstract_declarator:LPARAM.RPARAM
      direct_abstract_declarator:LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator:LPARAM.parameter_type_list RPARAM
      direct_declarator:LPARAM.declarator RPARAM
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
      parameter_declaration:declaration_specifiers abstract_declarator.,*
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      parameter_declaration:declaration_specifiers declarator.,*
    Reduce:
      * -> [parameter_declaration]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      abstract_declarator:direct_abstract_declarator.,*
      direct_abstract_declarator:direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator:direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator:direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator:direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 297
      LBRACE -> State 296

  State 228:
    Kernel Items:
      abstract_declarator:pointer.,*
      abstract_declarator:pointer.direct_abstract_declarator
      declarator:pointer.direct_declarator
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
      identifier_list:identifier_list COMMA.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 299

  State 230:
    Kernel Items:
      direct_declarator:direct_declarator LPARAM identifier_list RPARAM.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      parameter_list:parameter_list COMMA.parameter_declaration
      parameter_type_list:parameter_list COMMA.ELLIPSIS
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
      direct_declarator:direct_declarator LPARAM parameter_type_list RPARAM.,*
    Reduce:
      * -> [direct_declarator]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      struct_declaration_list:struct_declaration_list.struct_declaration
      struct_or_union_specifier:struct_or_union IDENTIFIER LCURL struct_declaration_list.RCURL
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
      struct_declarator:COLON.constant_expression
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
      struct_declarator:declarator.,*
      struct_declarator:declarator.COLON constant_expression
    Reduce:
      * -> [struct_declarator]
    Goto:
      COLON -> State 304

  State 236:
    Kernel Items:
      struct_declarator_list:struct_declarator.,*
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      struct_declaration:specifier_qualifier_list struct_declarator_list.SEMICOLON
      struct_declarator_list:struct_declarator_list.COMMA struct_declarator
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 306
      COMMA -> State 305

  State 238:
    Kernel Items:
      struct_or_union_specifier:struct_or_union LCURL struct_declaration_list RCURL.,*
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      struct_declaration_list:struct_declaration_list struct_declaration.,*
    Reduce:
      * -> [struct_declaration_list]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      specifier_qualifier_list:type_qualifier specifier_qualifier_list.,*
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      specifier_qualifier_list:type_specifier specifier_qualifier_list.,*
    Reduce:
      * -> [specifier_qualifier_list]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      enum_specifier:ENUM IDENTIFIER LCURL enumerator_list RCURL.,*
    Reduce:
      * -> [enum_specifier]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      enumerator:IDENTIFIER EQ constant_expression.,*
    Reduce:
      * -> [enumerator]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      enumerator_list:enumerator_list COMMA enumerator.,*
    Reduce:
      * -> [enumerator_list]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      initializer_list:initializer.,*
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      initializer:LCURL initializer_list.COMMA RCURL
      initializer:LCURL initializer_list.RCURL
      initializer_list:initializer_list.COMMA initializer
    Reduce:
      (nil)
    Goto:
      RCURL -> State 308
      COMMA -> State 307

  State 247:
    Kernel Items:
      labeled_statement:CASE constant_expression COLON.statement
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
      labeled_statement:DEFAULT COLON statement.,*
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      iteration_statement:DO statement WHILE.LPARAM expression RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      LPARAM -> State 310

  State 250:
    Kernel Items:
      iteration_statement:FOR LPARAM expression_statement.expression_statement RPARAM statement
      iteration_statement:FOR LPARAM expression_statement.expression_statement expression RPARAM statement
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
      jump_statement:GOTO IDENTIFIER SEMICOLON.,*
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      labeled_statement:IDENTIFIER COLON statement.,*
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      selection_statement:IF LPARAM expression.RPARAM statement
      selection_statement:IF LPARAM expression.RPARAM statement ELSE statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 312
      COMMA -> State 186

  State 254:
    Kernel Items:
      primary_expression:LPARAM expression RPARAM.,*
    Reduce:
      * -> [primary_expression]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      direct_abstract_declarator:LPARAM.RPARAM
      direct_abstract_declarator:LPARAM.abstract_declarator RPARAM
      direct_abstract_declarator:LPARAM.parameter_type_list RPARAM
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
      type_name:specifier_qualifier_list abstract_declarator.,RPARAM
    Reduce:
      RPARAM -> [type_name]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      abstract_declarator:pointer.,RPARAM
      abstract_declarator:pointer.direct_abstract_declarator
    Reduce:
      RPARAM -> [abstract_declarator]
    Goto:
      LPARAM -> State 255
      LBRACE -> State 223
      direct_abstract_declarator -> State 298

  State 258:
    Kernel Items:
      cast_expression:LPARAM type_name RPARAM.cast_expression
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
      jump_statement:RETURN expression SEMICOLON.,*
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      unary_expression:SIZEOF LPARAM type_name.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 314

  State 261:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      selection_statement:SWITCH LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 315
      COMMA -> State 186

  State 262:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      iteration_statement:WHILE LPARAM expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 316
      COMMA -> State 186

  State 263:
    Kernel Items:
      additive_expression:additive_expression MINUS multiplicative_expression.,*
      multiplicative_expression:multiplicative_expression.DIV cast_expression
      multiplicative_expression:multiplicative_expression.MOD cast_expression
      multiplicative_expression:multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 264:
    Kernel Items:
      additive_expression:additive_expression PLUS multiplicative_expression.,*
      multiplicative_expression:multiplicative_expression.DIV cast_expression
      multiplicative_expression:multiplicative_expression.MOD cast_expression
      multiplicative_expression:multiplicative_expression.MUL cast_expression
    Reduce:
      * -> [additive_expression]
    Goto:
      MUL -> State 194
      DIV -> State 192
      MOD -> State 193

  State 265:
    Kernel Items:
      and_expression:and_expression AND equality_expression.,*
      equality_expression:equality_expression.EQ_OP relational_expression
      equality_expression:equality_expression.NE_OP relational_expression
    Reduce:
      * -> [and_expression]
    Goto:
      EQ_OP -> State 183
      NE_OP -> State 184

  State 266:
    Kernel Items:
      compound_statement:LCURL declaration_list statement_list RCURL.,*
    Reduce:
      * -> [compound_statement]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      equality_expression:equality_expression EQ_OP relational_expression.,*
      relational_expression:relational_expression.GE_OP shift_expression
      relational_expression:relational_expression.GT shift_expression
      relational_expression:relational_expression.LE_OP shift_expression
      relational_expression:relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 268:
    Kernel Items:
      equality_expression:equality_expression NE_OP relational_expression.,*
      relational_expression:relational_expression.GE_OP shift_expression
      relational_expression:relational_expression.GT shift_expression
      relational_expression:relational_expression.LE_OP shift_expression
      relational_expression:relational_expression.LT shift_expression
    Reduce:
      * -> [equality_expression]
    Goto:
      LE_OP -> State 203
      GE_OP -> State 201
      LT -> State 204
      GT -> State 202

  State 269:
    Kernel Items:
      and_expression:and_expression.AND equality_expression
      exclusive_or_expression:exclusive_or_expression HAT and_expression.,*
    Reduce:
      * -> [exclusive_or_expression]
    Goto:
      AND -> State 180

  State 270:
    Kernel Items:
      expression:expression COMMA assignment_expression.,*
    Reduce:
      * -> [expression]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      exclusive_or_expression:exclusive_or_expression.HAT and_expression
      inclusive_or_expression:inclusive_or_expression OR exclusive_or_expression.,*
    Reduce:
      * -> [inclusive_or_expression]
    Goto:
      HAT -> State 185

  State 272:
    Kernel Items:
      inclusive_or_expression:inclusive_or_expression.OR exclusive_or_expression
      logical_and_expression:logical_and_expression AND_OP inclusive_or_expression.,*
    Reduce:
      * -> [logical_and_expression]
    Goto:
      OR -> State 188

  State 273:
    Kernel Items:
      logical_and_expression:logical_and_expression.AND_OP inclusive_or_expression
      logical_or_expression:logical_or_expression OR_OP logical_and_expression.,*
    Reduce:
      * -> [logical_or_expression]
    Goto:
      AND_OP -> State 189

  State 274:
    Kernel Items:
      conditional_expression:logical_or_expression QUESTION expression.COLON conditional_expression
      expression:expression.COMMA assignment_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 317
      COMMA -> State 186

  State 275:
    Kernel Items:
      multiplicative_expression:multiplicative_expression DIV cast_expression.,*
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      multiplicative_expression:multiplicative_expression MOD cast_expression.,*
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      multiplicative_expression:multiplicative_expression MUL cast_expression.,*
    Reduce:
      * -> [multiplicative_expression]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      postfix_expression:postfix_expression DOT IDENTIFIER.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      postfix_expression:postfix_expression LBRACE expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 318
      COMMA -> State 186

  State 280:
    Kernel Items:
      postfix_expression:postfix_expression LPARAM RPARAM.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      argument_expression_list:argument_expression_list.COMMA assignment_expression
      postfix_expression:postfix_expression LPARAM argument_expression_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 320
      COMMA -> State 319

  State 282:
    Kernel Items:
      argument_expression_list:assignment_expression.,*
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      postfix_expression:postfix_expression PTR_OP IDENTIFIER.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      relational_expression:relational_expression GE_OP shift_expression.,*
      shift_expression:shift_expression.LEFT_OP additive_expression
      shift_expression:shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 285:
    Kernel Items:
      relational_expression:relational_expression GT shift_expression.,*
      shift_expression:shift_expression.LEFT_OP additive_expression
      shift_expression:shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 286:
    Kernel Items:
      relational_expression:relational_expression LE_OP shift_expression.,*
      shift_expression:shift_expression.LEFT_OP additive_expression
      shift_expression:shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 287:
    Kernel Items:
      relational_expression:relational_expression LT shift_expression.,*
      shift_expression:shift_expression.LEFT_OP additive_expression
      shift_expression:shift_expression.RIGHT_OP additive_expression
    Reduce:
      * -> [relational_expression]
    Goto:
      LEFT_OP -> State 205
      RIGHT_OP -> State 206

  State 288:
    Kernel Items:
      additive_expression:additive_expression.MINUS multiplicative_expression
      additive_expression:additive_expression.PLUS multiplicative_expression
      shift_expression:shift_expression LEFT_OP additive_expression.,*
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 289:
    Kernel Items:
      additive_expression:additive_expression.MINUS multiplicative_expression
      additive_expression:additive_expression.PLUS multiplicative_expression
      shift_expression:shift_expression RIGHT_OP additive_expression.,*
    Reduce:
      * -> [shift_expression]
    Goto:
      MINUS -> State 178
      PLUS -> State 179

  State 290:
    Kernel Items:
      assignment_expression:unary_expression assignment_operator assignment_expression.,*
    Reduce:
      * -> [assignment_expression]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      direct_abstract_declarator:LBRACE RBRACE.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      direct_abstract_declarator:LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 321

  State 293:
    Kernel Items:
      direct_abstract_declarator:LPARAM RPARAM.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      direct_abstract_declarator:LPARAM abstract_declarator.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 322

  State 295:
    Kernel Items:
      direct_abstract_declarator:LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 323

  State 296:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LBRACE.RBRACE
      direct_abstract_declarator:direct_abstract_declarator LBRACE.constant_expression RBRACE
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
      direct_abstract_declarator:direct_abstract_declarator LPARAM.RPARAM
      direct_abstract_declarator:direct_abstract_declarator LPARAM.parameter_type_list RPARAM
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
      abstract_declarator:pointer direct_abstract_declarator.,*
      direct_abstract_declarator:direct_abstract_declarator.LBRACE RBRACE
      direct_abstract_declarator:direct_abstract_declarator.LBRACE constant_expression RBRACE
      direct_abstract_declarator:direct_abstract_declarator.LPARAM RPARAM
      direct_abstract_declarator:direct_abstract_declarator.LPARAM parameter_type_list RPARAM
    Reduce:
      * -> [abstract_declarator]
    Goto:
      LPARAM -> State 297
      LBRACE -> State 296

  State 299:
    Kernel Items:
      identifier_list:identifier_list COMMA IDENTIFIER.,*
    Reduce:
      * -> [identifier_list]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      parameter_type_list:parameter_list COMMA ELLIPSIS.,RPARAM
    Reduce:
      RPARAM -> [parameter_type_list]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      parameter_list:parameter_list COMMA parameter_declaration.,*
    Reduce:
      * -> [parameter_list]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      struct_or_union_specifier:struct_or_union IDENTIFIER LCURL struct_declaration_list RCURL.,*
    Reduce:
      * -> [struct_or_union_specifier]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      struct_declarator:COLON constant_expression.,*
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      struct_declarator:declarator COLON.constant_expression
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
      struct_declarator_list:struct_declarator_list COMMA.struct_declarator
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
      struct_declaration:specifier_qualifier_list struct_declarator_list SEMICOLON.,*
    Reduce:
      * -> [struct_declaration]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      initializer:LCURL initializer_list COMMA.RCURL
      initializer_list:initializer_list COMMA.initializer
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
      initializer:LCURL initializer_list RCURL.,*
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      labeled_statement:CASE constant_expression COLON statement.,*
    Reduce:
      * -> [labeled_statement]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      iteration_statement:DO statement WHILE LPARAM.expression RPARAM SEMICOLON
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
      iteration_statement:FOR LPARAM expression_statement expression_statement.RPARAM statement
      iteration_statement:FOR LPARAM expression_statement expression_statement.expression RPARAM statement
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
      selection_statement:IF LPARAM expression RPARAM.statement
      selection_statement:IF LPARAM expression RPARAM.statement ELSE statement
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
      cast_expression:LPARAM type_name RPARAM cast_expression.,*
    Reduce:
      * -> [cast_expression]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      unary_expression:SIZEOF LPARAM type_name RPARAM.,*
    Reduce:
      * -> [unary_expression]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      selection_statement:SWITCH LPARAM expression RPARAM.statement
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
      iteration_statement:WHILE LPARAM expression RPARAM.statement
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
      conditional_expression:logical_or_expression QUESTION expression COLON.conditional_expression
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
      postfix_expression:postfix_expression LBRACE expression RBRACE.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      argument_expression_list:argument_expression_list COMMA.assignment_expression
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
      postfix_expression:postfix_expression LPARAM argument_expression_list RPARAM.,*
    Reduce:
      * -> [postfix_expression]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      direct_abstract_declarator:LBRACE constant_expression RBRACE.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      direct_abstract_declarator:LPARAM abstract_declarator RPARAM.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      direct_abstract_declarator:LPARAM parameter_type_list RPARAM.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LBRACE RBRACE.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LBRACE constant_expression.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 342

  State 326:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LPARAM RPARAM.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LPARAM parameter_type_list.RPARAM
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 343

  State 328:
    Kernel Items:
      struct_declarator:declarator COLON constant_expression.,*
    Reduce:
      * -> [struct_declarator]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      struct_declarator_list:struct_declarator_list COMMA struct_declarator.,*
    Reduce:
      * -> [struct_declarator_list]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      initializer:LCURL initializer_list COMMA RCURL.,*
    Reduce:
      * -> [initializer]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      initializer_list:initializer_list COMMA initializer.,*
    Reduce:
      * -> [initializer_list]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      expression:expression.COMMA assignment_expression
      iteration_statement:DO statement WHILE LPARAM expression.RPARAM SEMICOLON
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 344
      COMMA -> State 186

  State 333:
    Kernel Items:
      iteration_statement:FOR LPARAM expression_statement expression_statement RPARAM.statement
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
      expression:expression.COMMA assignment_expression
      iteration_statement:FOR LPARAM expression_statement expression_statement expression.RPARAM statement
    Reduce:
      (nil)
    Goto:
      RPARAM -> State 346
      COMMA -> State 186

  State 335:
    Kernel Items:
      selection_statement:IF LPARAM expression RPARAM statement.,*
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement
    Reduce:
      * -> [selection_statement]
    Goto:
      ELSE -> State 347

  State 336:
    Kernel Items:
      selection_statement:IF LPARAM expression RPARAM statement.,AND
      selection_statement:IF LPARAM expression RPARAM statement.,BREAK
      selection_statement:IF LPARAM expression RPARAM statement.,CASE
      selection_statement:IF LPARAM expression RPARAM statement.,CONSTANT
      selection_statement:IF LPARAM expression RPARAM statement.,CONTINUE
      selection_statement:IF LPARAM expression RPARAM statement.,DEC_OP
      selection_statement:IF LPARAM expression RPARAM statement.,DEFAULT
      selection_statement:IF LPARAM expression RPARAM statement.,DO
      selection_statement:IF LPARAM expression RPARAM statement.,ELSE
      selection_statement:IF LPARAM expression RPARAM statement.,EXCLAIM
      selection_statement:IF LPARAM expression RPARAM statement.,FOR
      selection_statement:IF LPARAM expression RPARAM statement.,GOTO
      selection_statement:IF LPARAM expression RPARAM statement.,IDENTIFIER
      selection_statement:IF LPARAM expression RPARAM statement.,IF
      selection_statement:IF LPARAM expression RPARAM statement.,INC_OP
      selection_statement:IF LPARAM expression RPARAM statement.,LCURL
      selection_statement:IF LPARAM expression RPARAM statement.,LPARAM
      selection_statement:IF LPARAM expression RPARAM statement.,MINUS
      selection_statement:IF LPARAM expression RPARAM statement.,MUL
      selection_statement:IF LPARAM expression RPARAM statement.,PLUS
      selection_statement:IF LPARAM expression RPARAM statement.,RCURL
      selection_statement:IF LPARAM expression RPARAM statement.,RETURN
      selection_statement:IF LPARAM expression RPARAM statement.,SEMICOLON
      selection_statement:IF LPARAM expression RPARAM statement.,SIZEOF
      selection_statement:IF LPARAM expression RPARAM statement.,STRING_LITERAL
      selection_statement:IF LPARAM expression RPARAM statement.,SWITCH
      selection_statement:IF LPARAM expression RPARAM statement.,TILDA
      selection_statement:IF LPARAM expression RPARAM statement.,WHILE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,AND
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,BREAK
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,CASE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,CONSTANT
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,CONTINUE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,DEC_OP
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,DEFAULT
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,DO
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,ELSE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,EXCLAIM
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,FOR
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,GOTO
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,IDENTIFIER
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,IF
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,INC_OP
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,LCURL
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,LPARAM
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,MINUS
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,MUL
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,PLUS
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,RCURL
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,RETURN
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,SEMICOLON
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,SIZEOF
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,STRING_LITERAL
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,SWITCH
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,TILDA
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,WHILE
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
      selection_statement:IF LPARAM expression RPARAM statement.,ELSE
      selection_statement:IF LPARAM expression RPARAM statement.,WHILE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,ELSE
      selection_statement:IF LPARAM expression RPARAM statement.ELSE statement,WHILE
    Reduce:
      ELSE -> [selection_statement]
      WHILE -> [selection_statement]
    Goto:
      ELSE -> State 347
    Shift/reduce conflict symbols:
      [ELSE]

  State 338:
    Kernel Items:
      selection_statement:SWITCH LPARAM expression RPARAM statement.,*
    Reduce:
      * -> [selection_statement]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      iteration_statement:WHILE LPARAM expression RPARAM statement.,*
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      conditional_expression:logical_or_expression QUESTION expression COLON conditional_expression.,*
    Reduce:
      * -> [conditional_expression]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      argument_expression_list:argument_expression_list COMMA assignment_expression.,*
    Reduce:
      * -> [argument_expression_list]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LBRACE constant_expression RBRACE.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      direct_abstract_declarator:direct_abstract_declarator LPARAM parameter_type_list RPARAM.,*
    Reduce:
      * -> [direct_abstract_declarator]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      iteration_statement:DO statement WHILE LPARAM expression RPARAM.SEMICOLON
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 348

  State 345:
    Kernel Items:
      iteration_statement:FOR LPARAM expression_statement expression_statement RPARAM statement.,*
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      iteration_statement:FOR LPARAM expression_statement expression_statement expression RPARAM.statement
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
      selection_statement:IF LPARAM expression RPARAM statement ELSE.statement
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
      iteration_statement:DO statement WHILE LPARAM expression RPARAM SEMICOLON.,*
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      iteration_statement:FOR LPARAM expression_statement expression_statement expression RPARAM statement.,*
    Reduce:
      * -> [iteration_statement]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      selection_statement:IF LPARAM expression RPARAM statement ELSE statement.,*
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
