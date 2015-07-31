py_library(
    name='cc_rules',
    srcs=['cc_rules.py'],
    deps=['//buildutil:target_rule'],
    visibility=[])

py_library(
    name='lex_yacc_rules',
    srcs=['lex_yacc_rules.py'],
    deps=[
        '//buildutil:target_rule',
        ],
    visibility=[])

py_library(
    name='py_rules',
    srcs=['py_rules.py'],
    deps=[
        '//buildutil:target_patterns',
        '//buildutil:target_rule',
        ],
    visibility=[])

py_library(
    name='rules',
    srcs=['rules.py'],
    deps=[
        ':cc_rules',
        ':lex_yacc_rules',
        ':py_rules',
        ],
    visibility=['//buildutil:package'])

