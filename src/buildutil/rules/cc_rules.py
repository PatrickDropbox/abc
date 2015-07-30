import os.path

from buildutil.target_rule import TargetRule


HDR_EXTS = ('.h', '.hh')
SRC_EXTS = ('.c', '.cc', '.cpp')

CC_SECTION = 'cc_rules'

DEFAULT_CC = 'g++'
DEFAULT_CFLAGS = '-Wall'


class CcBaseTargetRule(TargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      hdrs=(),
      deps=(),
      visibility=None):

    sources = []
    artifacts = []

    for hdr in set(hdrs):
      _, ext = os.path.splitext(hdr)
      assert ext in HDR_EXTS, (
          'Unexpected header extension: %s (target: %s:%s)' % (
              hdr,
              pkg_path,
              name))

      sources.append(hdr)
      artifacts.append(hdr)

    for src in set(srcs):
      src_name, ext = os.path.splitext(src)
      assert ext in SRC_EXTS, (
          'Unexpected src extension: %s (target: %s:%s)' % (
              src,
              pkg_path,
              name))

      sources.append(src)
      artifacts.append(src_name + '.o')

    super(CcBaseTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        sources=sources,
        dependencies=deps,
        artifacts=artifacts,
        visibility_set=visibility)

  def build(self):
    cc = self.config.get(CC_SECTION, 'cc_location', DEFAULT_CC)
    cflags = self.config.get(CC_SECTION, 'cflags', DEFAULT_CFLAGS)
    hdr_dirs = (self.config.src_dir_abs_path, self.config.genfile_dir_abs_path)

    for src in self.sources:
      name, ext = os.path.splitext(src)
      if ext in HDR_EXTS:
        continue

      self.execute_cmd(
          '%s -c %s %s -o %s %s' % (
              cc,
              cflags,
              ' '.join('-I' + d for d in hdr_dirs),
              self.build_abs_path(name=name + '.o'),
              self.locate_file(src)))

    return True


class CcLibraryTargetRule(CcBaseTargetRule):
  def __init__(
      self,
      config,
      pkg_path,
      name,
      srcs=(),
      hdrs=(),
      deps=(),
      visibility=None):
    super(CcLibraryTargetRule, self).__init__(
        config,
        pkg_path,
        name,
        srcs=srcs,
        hdrs=hdrs,
        deps=deps,
        visibility=visibility)

  @classmethod
  def rule_name(cls):
    return "cc_library"

