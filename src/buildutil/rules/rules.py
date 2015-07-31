from buildutil.rules.cc_rules import (
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,
    )
from buildutil.rules.lex_yacc_rules import (
    LexTargetRule,
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

    LexTargetRule,

    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
