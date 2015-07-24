from buildutil.rules.cc_rules import CcLibraryTargetRule
from buildutil.rules.py_rules import (
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    )

# list of TargetRule subclasses.  Each subclass must take pkg as argument.
RULES = [
    CcLibraryTargetRule,
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
