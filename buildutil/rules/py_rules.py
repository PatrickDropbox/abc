from buildutil.target_rule import TargetRule


class PyLibraryTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    # TODO
    # TODO include __init__.py implicitly
    super(PyLibraryTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "py_library"


class PyBinaryTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    # TODO
    # TODO include __init__.py implicitly
    super(PyBinaryTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "py_binary"

  @classmethod
  def register(cls, pkg, **kwargs):
    pkg.register(cls(pkg=pkg, **kwargs))
    kwargs['name'] = kwargs['name'] + '.par'
    pkg.register(PyParTargetRule(pkg=pkg, **kwargs))


# TODO
class PyParTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    super(PyParTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    assert False, 'Should never be called directly'


class PyTestTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    # TODO
    # TODO include __init__.py implicitly
    super(PyTestTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "py_test"

  @classmethod
  def is_test_rule(cls):
    return True
