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
    src_init = self.src_abs_path(name=init)
    genfile_init = self.genfile_abs_path(name=init)
    build_init = self.build_abs_path(name=init)

    assert not os.path.exists(build_init)

    if os.path.isfile(src_init):
      return False
    elif os.path.isfile(genfile_init):
      return False

    return True

  def build(self):
    init = '__init__.py'
    src_init = self.src_abs_path(name=init)
    genfile_init = self.genfile_abs_path(name=init)

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
        dependencies=list(deps) + [':__init__.py'],
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def register(cls, pkg, **kwargs):
    pkg.register(PyInitTargetRule(pkg), ignore_duplicate=True)
    pkg.register(cls(pkg=pkg, **kwargs))

  @classmethod
  def rule_name(cls):
    return "py_library"

  def build(self):
    for src_file_name in self.sources:
      abs_path = self.locate_file(src_file_name)
      assert abs_path

      r = self.execute_cmd('touch %s' % abs_path)
      if r != 0:
        return False

    return True


class PyBinaryTargetRule(TargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    super(PyBinaryTargetRule, self).__init__(
        pkg,
        name,
        sources=srcs,
        dependencies=list(deps) + [':__init__.py'],
        artifacts=[],  # compute dynamically
        visibility_set=visibility)

  @classmethod
  def register(cls, pkg, **kwargs):
    pkg.register(PyInitTargetRule(pkg), ignore_duplicate=True)
    pkg.register(cls(pkg=pkg, **kwargs))

    name = kwargs['name']
    visibility = None
    if 'visibility' in kwargs:
      visibility = kwargs['visibility']
    pkg.register(PyParTargetRule(pkg=pkg, name=name, visibility=visibility))

  @classmethod
  def rule_name(cls):
    return "py_binary"

  def artifacts(self):
    assert self.deps_binded

    result = set()
    for artifact in self.list_dependencies_artifacts():
      result.add(
        os.path.join(self.name + '.runtime', artifact[2:]))

    return result

  @classmethod
  def include_dependencies_artifacts(cls):
    return False

  def build(self):
    src_pkg_paths = self.list_dependencies_artifacts()
    for name in self.sources:
      src_pkg_paths.add(self.pkg_path(name=name))

    runtime_abs_path = self.build_abs_path(name=self.name + '.runtime')
    r = self.execute_cmd('rm -rf %s' % runtime_abs_path)
    if r != 0:
      return False

    dir_abs_paths = set()
    for pkg_path in src_pkg_paths:
      pkg_path, _ = os.path.split(pkg_path)
      assert pkg_path.startswith('//')
      dir_abs_paths.add(os.path.join(runtime_abs_path, pkg_path[2:]))

    for abs_path in dir_abs_paths:
      r = self.execute_cmd('mkdir -p %s' % abs_path)
      if r != 0:
        return False

    for pkg_path in src_pkg_paths:
      abs_path = self.config.locate_file(pkg_path, include_build=False)
      assert abs_path, 'Failed to locate: %s' % pkg_path

      r = self.execute_cmd('ln -s %s %s' % (
          abs_path,
          os.path.join(runtime_abs_path, pkg_path[2:])))
      if r != 0:
        return False
    return True

# TODO
class PyParTargetRule(TargetRule):
  def __init__(self, pkg, name, visibility=None):

    super(PyParTargetRule, self).__init__(
        pkg,
        name + '.par',
        sources=(),
        dependencies=[':%s' % name],
        artifacts=[name + '.par'],
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    assert False, 'Should never be called directly'

  @classmethod
  def include_dependencies_artifacts(cls):
    return False


class PyTestTargetRule(PyBinaryTargetRule):
  def __init__(
      self,
      pkg,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    super(PyTestTargetRule, self).__init__(
        pkg,
        name,
        srcs=srcs,
        deps=deps,
        visibility=visibility)

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
