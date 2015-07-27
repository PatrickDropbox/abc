from buildutil.target_rule import TargetRule


class CcLibraryTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      hdrs=(),
      deps=(),
      visibilty=None):

    # TODO
    artifiacts = () + hdrs
    super(CcLibraryTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        sources=srcs+hdrs,
        dependencies=deps,
        artifacts=artifacts,
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "cc_library"
