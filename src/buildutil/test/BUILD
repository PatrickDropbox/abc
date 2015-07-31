cc_library(
    name='lib1',
    srcs=[
        'lib1a.cc',
        'lib1b.cc',
        ],
    hdrs=['lib1.h'])


cc_library(
    name='lib2',
    srcs=['lib2.cc'],
    hdrs=['lib2.h'],
    deps=[
        ':lib1',
        ':tmpl',
        ])

cc_library(
    name='tmpl',
    hdrs=['tmpl_inl.h'])

cc_binary(
    name='hello_world',
    srcs=[
        'hello.cc',
        'world.cc',
        ],
    hdrs=['hello_world.h'],
    deps=[':lib2'])

cc_test(
    name='foo_unittest',
    srcs=['foo_unittest.cc'],
    deps=['//gtest:gtest'])

gen_lex(
    name='scanner_lex',
    srcs=['scanner.l'])

cc_binary(
    name='scanner',
    srcs=['scanner.yy.c'],
    deps=[':scanner_lex'])
