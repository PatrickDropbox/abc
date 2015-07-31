import os.path

from buildutil.target_rule import TargetRule


LEX_YACC_SECTION = 'lex_yacc_rules'
DEFAULT_LEX = '/usr/bin/flex'

class LexTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      deps=(),
      visibility=None):
    assert len(srcs) == 1
    file_name, ext = os.path.splitext(srcs[0])
    assert ext == '.l'

    super(LexTargetRule, self).__init__(
        config=config,
        pkg_path=pkg_path,
        name=name,
        sources=srcs,
        dependencies=deps,
        artifacts=[file_name + '.yy.c'],
        visibility_set=visibility)

  @classmethod
  def rule_name(cls):
    return "gen_lex"

  def build(self):
    lex = self.config.get(LEX_YACC_SECTION, 'lex_location', DEFAULT_LEX)

    src = self.sources[0]
    name, _ = os.path.splitext(src)

    self.execute_cmd('%s -o %s %s' % (
        lex,
        self.genfile_abs_path(name + '.yy.c'),
        self.locate_file(src)))

    return True

  def list_artifacts(self):
    # Don't propagate artifacts.
    return []
