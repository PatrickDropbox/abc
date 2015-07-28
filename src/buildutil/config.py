# Convention ...
# *pkg_path = relative to project root
# *target_path = relative to project root
# *abs_path = relative to filesystem's root
# *name = relative to local context
import ConfigParser
import os
import os.path


class Config(object):
  def __init__(self, ini_config):
    self.ini_config = ini_config

    section = 'buildutil'

    project_dir_name = self.get(section, 'project_dir_name')
    src_dir_name = self.get(section, 'src_dir_name')
    genfile_dir_name = self.get(section, 'genfile_dir_name')
    build_dir_name = self.get(section, 'build_dir_name')
    """
    project_dir_name = self.get(section, 'project_dir_name', 'abc')
    src_dir_name = self.get(section, 'src_dir_name', 'src')
    genfile_dir_name = self.get(section, 'genfile_dir_name', 'genfile')
    build_dir_name = self.get(section, 'build_dir_name', 'build')
    """

    assert '/' not in project_dir_name
    assert '/' not in src_dir_name
    assert '/' not in build_dir_name
    assert '/' not in genfile_dir_name
    assert src_dir_name != genfile_dir_name
    assert src_dir_name != build_dir_name
    assert genfile_dir_name != build_dir_name

    cwd = os.getcwd() + '/'

    sep = '/%s/' % project_dir_name
    assert sep in cwd
    base_abs_path, _, _ = cwd.partition(sep)

    self.project_dir_abs_path = os.path.normpath(
        base_abs_path + '/' + project_dir_name)
    self.src_dir_abs_path = os.path.normpath(
        os.path.join(self.project_dir_abs_path, src_dir_name))
    self.genfile_dir_abs_path = os.path.normpath(
        os.path.join(self.project_dir_abs_path, genfile_dir_name))
    self.build_dir_abs_path = os.path.normpath(
        os.path.join(self.project_dir_abs_path, build_dir_name))

    assert os.path.isdir(self.project_dir_abs_path)
    assert os.path.isdir(self.src_dir_abs_path)

  def get(self, section, key, default_value=None):
    if self.ini_config is None:
      return default_value

    if self.ini_config.has_option(section, key):
      return self.ini_config.get(section, key)

    return default_value

  def src_abs_path_to_pkg_path(self, abs_path):
    abs_path = os.path.normpath(abs_path)
    if abs_path == self.src_dir_abs_path:
      return '//'

    src_dir_abs_path = self.src_dir_abs_path + '/'

    assert abs_path.startswith(src_dir_abs_path)
    return '//' + abs_path[len(src_dir_abs_path):]

  def pkg_path_to_src_abs_path(self, pkg_path):
    assert pkg_path.startswith('//'), pkg_path
    assert '../' not in pkg_path
    return os.path.normpath(os.path.join(self.src_dir_abs_path, pkg_path[2:]))

  def pkg_path_to_genfile_abs_path(self, pkg_path):
    assert pkg_path.startswith('//'), pkg_path
    assert '../' not in pkg_path
    return os.path.normpath(
        os.path.join(self.genfile_dir_abs_path, pkg_path[2:]))

  def pkg_path_to_build_abs_path(self, pkg_path):
    assert pkg_path.startswith('//'), pkg_path
    assert '../' not in pkg_path
    return os.path.normpath(
        os.path.join(self.build_dir_abs_path, pkg_path[2:]))

  def locate_file(
      self,
      file_pkg_path,
      include_src=True,
      include_genfile=True,
      include_build=True):
    if include_src:
      src_path = self.pkg_path_to_src_abs_path(file_pkg_path)
      if os.path.isfile(src_path):
        return src_path

    if include_genfile:
      genfile_path = self.pkg_path_to_genfile_abs_path(file_pkg_path)
      if os.path.isfile(genfile_path):
        return genfile_path

    if include_build:
      build_path = self.pkg_path_to_build_abs_path(file_pkg_path)
      if os.path.isfile(build_path):
        return build_path

    return None

