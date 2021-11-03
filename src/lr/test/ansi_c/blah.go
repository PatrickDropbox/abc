package ansi_c

type Symbol string

func (Symbol) Id() CSymbolId { return "" }

func (Symbol) Location() CLocation { return CLocation{} }
