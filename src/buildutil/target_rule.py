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

    assert artifacts, 'Target must have at least one artifact: %s (pkg: %s)' % (
        name,
        pkg.full_path)

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

    self.deps_binded = False
    self.dependencies = {}

    # None for not checked, False for no cycle, True for cycle.
    self.in_cycle = None

    # NOTE: If this was implemented as a server (like blaze), then we should
    # check source content hashes/size instead of sources/artifacts mtime
    self.soruces_max_mtime = None
    self.artifacts_max_mtime = None
    # None for not checked, True if has build, False if build was unnecessary.
    self.has_modified = None

  def full_name(self):
    """DO NOT OVERRIDE"""
    return self.package_path + ':' + self.name

  def is_visible_to(self, target):
    """DO NOT OVERRIDE"""
    if self.package_path == target.package_path:
      return True

    return self.visibility_patterns.matches(target)

  @classmethod
  def register(cls, pkg, **kwargs):
    """Override to customize target registration (see PyBinaryTargetRule for
    example)."""
    pkg.register(cls(pkg=pkg, **kwargs))

  @classmethod
  def rule_name(cls):
    """The function name used in config file, e.g., cc_library"""
    raise NotImplemented

  @classmethod
  def is_test_rule(cls):
    """When building libraries / binaries, test targets are ignored.  When
    true, must implement test"""
    return False

  def build(self, config):
    """How the target should be build"""
    print 'BUILD', self.name
    #raise NotImplemented

  def test(self, config):
    """How the target should be tested"""
    print 'TEST', self.name
    #raise NotImplemented

