from buildutil.rules.cc_rules import (
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,
    )
from buildutil.rules.go_rules import (
    GoBinaryTargetRule,
    GoLibraryTargetRule,
    )
from buildutil.rules.lex_yacc_rules import (
    LexTargetRule,
    YaccTargetRule,
    )
from buildutil.rules.py_rules import (
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    )

# list of TargetRule subclasses.
RULES = [
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,

    GoBinaryTargetRule,
    GoLibraryTargetRule,

    LexTargetRule,
    YaccTargetRule,

    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
