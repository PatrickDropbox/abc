import os.path


class Config(object):
  def __init__(
      self,
      project_root_dir_path,
      src_dir_name,
      genfile_dir_name,
      build_dir_name):
    assert '/' not in src_dir_name
    assert '/' not in build_dir_name
    assert '/' not in genfile_dir_name
    assert src_dir_name != genfile_dir_name
    assert src_dir_name != build_dir_name
    assert genfile_dir_name != build_dir_name

    self.project_dir = os.path.abspath(project_root_dir_path)
    self.src_dir = os.path.normpath(
        os.path.join(self.project_dir, src_dir_name))
    self.genfile_dir = os.path.normpath(
        os.path.join(self.project_dir, genfile_dir_name))
    self.build_dir = os.path.normpath(
        os.path.join(self.project_dir, build_dir_name))

    assert os.path.isdir(self.project_dir)
    assert os.path.isdir(self.src_dir)

  def pkg_src_dir_to_pkg_name(self, pkg_dir):
    pkg_dir = os.path.normpath(pkg_dir)
    if pkg_dir == self.src_dir:
      return '//'

    src_dir = self.src_dir + '/'

    assert pkg_dir.startswith(src_dir)
    return '//' + pkg_dir[len(src_dir):]

  def pkg_name_to_pkg_src_dir(self, pkg_name):
    assert pkg_name.startswith('//'), pkg_name
    return os.path.join(self.src_dir, pkg_name[2:])

  def pkg_name_to_pkg_genfile_dir(self, pkg_name):
    assert pkg_name.startswith('//'), pkg_name
    return os.path.join(self.genfile_dir, pkg_name[2:])

  def pkg_name_to_pkg_build_dir(self, pkg_name):
    assert pkg_name.startswith('//'), pkg_name
    return os.path.join(self.build_dir, pkg_name[2:])

  def locate_file(
      self,
      package_name,
      file_name,
      include_src_dir=True,
      include_genfile_dir=True,
      include_build_dir=True):
    assert not file_name.startswith('/')
    assert '../' not in file_name

    if include_src_dir:
      src_path = os.path.join(
          self.pkg_name_to_pkg_src_dir(package_name),
          file_name)
      if os.path.exists(src_path):
        return src_path

    if include_genfile_dir:
      genfile_path = os.path.join(
          self.pkg_name_to_pkg_genfile_dir(package_name),
          file_name)
      if os.path.exists(genfile_path):
        return genfile_path

    if include_build_dir:
      build_path = os.path.join(
          self.pkg_name_to_pkg_build_dir(package_name),
          file_name)
      if os.path.exists(build_path):
        return build_path

    return None

