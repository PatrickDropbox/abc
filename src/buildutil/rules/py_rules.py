import os
import os.path

from buildutil.target_rule import TargetRule
from buildutil.target_patterns import TargetPatterns


class PyInitTargetRule(TargetRule):
  def __init__(self, pkg):
    super(PyInitTargetRule, self).__init__(
        pkg,
        '__init__.py',
        sources=[],
        dependencies=[],
        artifacts=['__init__.py'],
        visibility_set=[])  # private

  @classmethod
  def rule_name(cls):
    assert False, 'Should never be called directly'

  def should_build(self):
    init = '__init__.py'
    src_init = self.src_file_path(init)
    genfile_init = self.genfile_file_path(init)
    build_init = self.build_file_path(init)

    assert not os.path.exists(build_init)

    if os.path.isfile(src_init):
      return False
    elif os.path.isfile(genfile_init):
      return False

    return True

  def build(self):
    init = '__init__.py'
    src_init = self.src_file_path(init)
    genfile_init = self.genfile_file_path(init)

    assert not os.path.isfile(src_init)

    r = self.execute_cmd('touch %s' % genfile_init)
    return r == 0


class PyLibraryTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    super(PyLibraryTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def register(cls, pkg, **kwargs):
    pkg.register(PyInitTargetRule(pkg), ignore_duplicate=True)
    pkg.register(cls(pkg=pkg, **kwargs))

  @classmethod
  def rule_name(cls):
    return "py_library"

  def should_build(self):
    # Never build since python "binary creation" is done via symlinks
    return False

  def build(self):
    # Do nothing
    assert False, 'This should never execute'


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
    pkg.register(PyInitTargetRule(pkg), ignore_duplicate=True)
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
  def register(cls, pkg, **kwargs):
    pkg.register(PyInitTargetRule(pkg), ignore_duplicate=True)
    pkg.register(cls(pkg=pkg, **kwargs))

  @classmethod
  def rule_name(cls):
    return "py_test"

  @classmethod
  def is_test_rule(cls):
    return True

  def test(self):
    print 'TEST', self.name
    return True
