from buildutil.rules.cc_rules import CcLibraryTargetRule
from buildutil.rules.py_rules import (
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    )

# list of TargetRule subclasses.
RULES = [
    CcLibraryTargetRule,
    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
