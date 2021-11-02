Naive / (mostly) canonical lr parser generator.

This deviates slightly from the dragon book's LR(1) presentation:

1. The start production is item is [#accept -> ^ . S, $] instead of
   [#accept -> . S, $].  This avoids the need to special case the start
   state's kernel.

2. The token stream act as pseudo symbol stack.  Instead of relying on
   a secondary GOTO table for reduction, the goto entries are encoded as
   shift actions in the ACTION table.  The reduction action will pop
   n symbols from state stack, and push the reduced symbol to the front
   of the pseudo symbol stack.  The parser then perform a second lookup
   to shift the reduced symbol.

3. In states with deterministic reduction (i.e., all items in the state of
   the form [A -> B . , x] have identical A and B), the reduction items
   are merged into a single entry [A -> B ., *] (and hence a single reduce
   action), where * indicates the look ahead symbol can be anything.
   Internally, an ACTION table lookup will first attempt to look up the look
   ahead specific action, and if that fails, fall back to looking up the
   wildcard action.

   In case of syntax error, ACTION lookup will fail on the next token, after
   the reduce symbol is pushed onto the state stack.

   Note that this potentially introduces new shift/reduce conflicit (in
   which case shifting is always favored), but unlike LALR, this does not
   introduce new reduce/reduce conflicit.

The generator's parser is self bootstrapped from ./internal/parser/grammar.lr
