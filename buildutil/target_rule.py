from buildutil.target_patterns import TargetPatterns
from buildutil.util import validate_target_name

class TargetRule(object):
  def __init__(
      self,
      pkg,
      name,
      sources=(),
      dependencies=(),
      artifacts=(),
      visibility_set=None):
    """
    sources: list of files (in relative path form) in the current directory
      (may also be in sub directories).
    dependencies: list of target path.
    artifacts: list of "output" files (in relative path form).
    visibility_set: list of visibility targets (None means use package default)
    """
    assert validate_target_name(name), (
        'Invalid target name: %s (pkg: %s)' % (name, pkg.full_path))

    self.package_path = pkg.full_path
    self.name = name
    self.sources = sources
    self.dependency_patterns = TargetPatterns(pkg.full_path)
    self.dependency_patterns.set_patterns(dependencies)
    self.artifacts = artifacts

    if visibility_set is not None:
      self.visibility_patterns = TargetPatterns(pkg.full_path)
      self.visibility_patterns.set_patterns(visibility_set)
    else:
      self.visibility_patterns = pkg.visibility_patterns

    # The following are initialized in later analysis passes.

    self.binded = False
    self.dependencies = {}

    # None for not checked, False for no cycle, True for cycle.
    self.in_cycle = None

  def full_name(self):
    return self.package_path + ':' + self.name

  def is_visible_to(self, target):
    if self.package_path == target.package_path:
      return True

    return self.visibility_patterns.matches(target)

  @classmethod
  def rule_name(cls):
    raise NotImplemented

  # When building libraries / binaries, test targets are ignored.
  @classmethod
  def is_test_rule(cls):
    return False
