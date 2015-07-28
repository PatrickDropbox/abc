import os
import os.path

from buildutil.target_rule import TargetRule
from buildutil.target_patterns import TargetPatterns


class PyInitTargetRule(TargetRule):
  def __init__(self, config, pkg_path):
    super(PyInitTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name='__init__.py',
        sources=[],
        dependencies=[],
        artifacts=['__init__.py'],
        visibility_set=['//...'])  # public to all

  @classmethod
  def is_unique_target(cls):
    return False

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    for pkg_path in PyInitTargetRule.list_dep_pkg_paths(current_pkg_path):
      targets_accumulator.append(
          PyInitTargetRule(config=config, pkg_path=pkg_path))

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
      # TODO: disallow __init__.py in src
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

  @staticmethod
  def list_dep_pkg_paths(pkg_path):
    result = []

    while True:
      assert pkg_path.startswith('//')
      result.append(pkg_path)

      if pkg_path == '//':
        return result

      pkg_path, _ = os.path.split(pkg_path)

  @staticmethod
  def list_dep_target_paths(pkg_path):
    result = []
    for path in PyInitTargetRule.list_dep_pkg_paths(pkg_path):
      result.append(path + ':__init__.py')

    return result


class PyLibraryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    deps = list(deps) + PyInitTargetRule.list_dep_target_paths(pkg_path)

    super(PyLibraryTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=srcs,
        visibility_set=visibility)

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    PyInitTargetRule.generate_targets(
        targets_accumulator=targets_accumulator,
        config=config,
        current_pkg_path=current_pkg_path)
    targets_accumulator.append(
        cls(config=config, pkg_path=current_pkg_path, **kwargs))

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


DEFAULT_BASH_ABS_PATH = '/bin/bash'
DEFAULT_PYTHON_ABS_PATH = '/usr/bin/python'

RUNNER_TEMPLATE = """#!%(bash)s
PYTHONPATH=%(runtime_dir)s %(python)s %(runtime_dir)s/%(main_py)s $@
"""


class PyBinaryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):

    assert len(srcs) == 1, 'Must have exactly one .py src file'

    deps = list(deps) + PyInitTargetRule.list_dep_target_paths(pkg_path)

    super(PyBinaryTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=[],  # compute dynamically
        visibility_set=visibility)

  @classmethod
  def generate_targets(
      cls,
      targets_accumulator,
      config,
      current_pkg_path,
      **kwargs):
    PyInitTargetRule.generate_targets(
        targets_accumulator,
        config,
        current_pkg_path)
    targets_accumulator.append(
        cls(config=config, pkg_path=current_pkg_path, **kwargs))

    name = kwargs['name']
    visibility = None
    if 'visibility' in kwargs:
      visibility = kwargs['visibility']

    targets_accumulator.append(
        PyParTargetRule(
            config=config,
            pkg_path=current_pkg_path,
            name=name,
            visibility=visibility))

  @classmethod
  def rule_name(cls):
    return "py_binary"

  def artifacts(self):
    assert self.deps_binded

    result = set()
    for artifact in self.list_dependencies_artifacts():
      result.add(
        os.path.join(self.name + '.runtime', artifact[2:]))

    result.add(self.name)
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

    return self.write_runner_script()

  def write_runner_script(self):
    script_abs_path = self.build_abs_path(name=self.name)

    section = 'py_binary'
    tmpl_vals ={
        'bash': self.config.get(
            section,
            'bash_location',
            DEFAULT_BASH_ABS_PATH),
        'runtime_dir': script_abs_path + '.runtime',
        'python': self.config.get(
            section,
            'python_location',
            DEFAULT_PYTHON_ABS_PATH),
        'main_py': self.pkg_path(name=self.sources[0])[2:],
        }

    print 'Writing:', script_abs_path, tmpl_vals
    with open(script_abs_path, 'w') as f:
      f.write(RUNNER_TEMPLATE % tmpl_vals)

    os.chmod(script_abs_path, 0755)
    return True


# TODO
class PyParTargetRule(TargetRule):
  def __init__(self, config, pkg_path, name, visibility=None):
    super(PyParTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name + '.par',
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


# TODO don't subclass ByBinary cuz it's limits to a single src file ...
class PyTestTargetRule(PyBinaryTargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):
    super(PyTestTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        srcs=srcs,
        deps=deps,
        visibility=visibility)

  @classmethod
  def rule_name(cls):
    return "py_test"

  @classmethod
  def is_test_rule(cls):
    return True

  def test(self):
    print 'TEST', self.name
    return True
