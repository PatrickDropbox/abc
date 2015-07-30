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

cc_test(
    name='hello_world',
    srcs=[
        'hello.cc',
        'world.cc',
        ],
    hdrs=['hello_world.h'],
    deps=[':lib2'])
