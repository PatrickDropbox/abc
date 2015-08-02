from buildutil.target_rule import TargetRule


# TODO
class GoLibraryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name='go',
      deps=()):

    sources = []
    artifacts = []

    super(GoPkgTargetRule, self).__init__(
        self,
        config,
        pkg_path,
        name,
        sources=sources,
        dependencies=deps,
        artifacts=artifacts,
        visibility_set=['//...'])

  @classmethod
  def rule_name(cls):
    return "go_library"

  def is_test_rule(self):
    return True

class GoBinaryTargetRule(GoLibraryTargetRule):
  @classmethod
  def rule_name(cls):
    return "go_binary"

