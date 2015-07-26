import os

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
    self.config = pkg.config
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

  def _get_max_mtime(self, files, verify_existence=False):
    """DO NOT OVERRIDE"""
    max_mtime = None
    for f in files:
      abs_path = self.locate_file(f)
      if abs_path is None:
        assert not verify_existence, (
            'Failed to locate: %s (target: %s)' % (f, self.full_path()))
        return None

      mtime = os.lstat(abs_path).st_mtime
      if max_mtime is None or max_mtime < mtime:
        max_mtime = mtime

  def update_sources_max_mtime(self):
    """DO NOT OVERRIDE"""
    self.sources_max_mtime = self._get_max_mtime(
        self.sources,
        verify_existence=True)

  def update_artifacts_max_mtime(self, verify_existence=True):
    """DO NOT OVERRIDE"""
    self.artifacts_max_mtime = self._get_max_mtime(
        self.artifacts,
        verify_existence=verify_existence)

  def pkg_src_dir(self):
    """DO NOT OVERRIDE"""
    return self.config.pkg_name_to_pkg_src_dir(self.package_path)

  def pkg_genfile_dir(self):
    """DO NOT OVERRIDE"""
    return self.config.pkg_name_to_pkg_genfile_dir(self.package_path)

  def pkg_build_dir(self):
    """DO NOT OVERRIDE"""
    return self.config.pkg_name_to_pkg_build_dir(self.package_path)

  def src_file_path(self, file_name):
    """DO NOT OVERRIDE"""
    return self.config.src_file_path(self.package_path, file_name)

  def genfile_file_path(self, file_name):
    """DO NOT OVERRIDE"""
    return self.config.genfile_file_path(self.package_path, file_name)

  def build_file_path(self, file_name):
    """DO NOT OVERRIDE"""
    return self.config.build_file_path(self.package_path, file_name)

  def locate_file(
      self,
      file_name,
      include_src_dir=True,
      include_genfile_dir=True,
      include_build_dir=True):
    """DO NOT OVERRIDE"""
    return self.config.locate_file(
        self.package_path,
        file_name,
        include_src_dir=include_src_dir,
        include_genfile_dir=include_genfile_dir,
        include_build_dir=include_build_dir)

  @classmethod
  def register(cls, pkg, **kwargs):
    """Override to customize target registration (see PyBinaryTargetRule for
    example)."""
    pkg.register(cls(pkg=pkg, **kwargs))

  @classmethod
  def rule_name(cls):
    """The function name used in BUILD file, e.g., cc_library"""
    raise NotImplemented

  @classmethod
  def is_test_rule(cls):
    """When building libraries / binaries, test targets are ignored.  When
    true, must implement test"""
    return False

  def should_build(self):
    # Artifacts are created without source and dependencies.
    if not self.sources and not self.dependencies:
      return True

    # First time building the artifacts.
    if self.artifacts_max_mtime is None:
      return True

    if self.sources:
      assert self.sources_max_mtime

      # Sources are newer than the artifacts.
      if self.artifacts_max_mtime < self.sources_max_mtime:
        return True

    for dep in self.dependencies.values():
      assert dep.has_modified is not None
      # A dependency changed within the same session (i.e., when multiple
      # targets are specified in the same build command).  This is more
      # accurate than the mtime check.
      if dep.has_modified:
        return True

      assert dep.artifacts_max_mtime is not None
      # The dependency changed from a previous session.
      if self.artifacts_max_mtime < dep.artifacts_max_mtime:
        return True

    return False

  def build(self):
    """How the target should be build.  Returns true if build succeeded, false
    otherwise."""
    print 'BUILD', self.name
    #raise NotImplemented
    return True

  def test(self):
    """How the target should be tested.  Returns true if test succeeded, false
    otherwise."""
    print 'TEST', self.name
    #raise NotImplemented
    return True

