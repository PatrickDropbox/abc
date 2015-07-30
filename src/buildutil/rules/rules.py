from buildutil.rules.cc_rules import (
    CcBinaryTargetRule,
    CcLibraryTargetRule,
    CcTestTargetRule,
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

    PyBinaryTargetRule,
    PyLibraryTargetRule,
    PyTestTargetRule,
    ]
