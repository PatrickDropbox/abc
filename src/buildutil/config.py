import os.path


class Config(object):
  def __init__(
      self,
      project_root_dir_path,
      src_dir_name,
      genrule_dir_name,
      build_dir_name):
    self.project_dir = os.path.abspath(project_root_dir_path)
    self.src_dir = os.path.normpath(
        os.path.join(self.project_dir, src_dir_name))
    self.genrule_dir = os.path.normpath(
        os.path.join(self.project_dir, genrule_dir_name))
    self.build_dir = os.path.normpath(
        os.path.join(self.project_dir, build_dir_name))

    assert os.path.isdir(self.project_dir)
    assert os.path.isdir(self.src_dir)

    # TODO make genrule / build dir

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

  def pkg_name_to_pkg_genrule_dir(self, pkg_name):
    assert pkg_name.startswith('//'), pkg_name
    return os.path.join(self.genrule_dir, pkg_name[2:])

  def pkg_name_to_pkg_build_dir(self, pkg_name):
    assert pkg_name.startswith('//'), pkg_name
    return os.path.join(self.build_dir, pkg_name[2:])

  def locate_file(
      self,
      package_name,
      file_name,
      include_build_dir=True,
      include_genrule_dir=True,
      include_src_dir=True):
    assert not file_name.startswith('/')
    assert '../' not in file_name

    if include_build_dir:
      build_path = os.path.join(
          self.pkg_name_to_pkg_build_dir(package_name),
          file_name)
      if os.path.exists(build_path):
        return build_path

    if include_genrule_dir:
      genrule_path = os.path.join(
          self.pkg_name_to_pkg_genrule_dir(package_name),
          file_name)
      if os.path.exists(genrule_path):
        return genrule_path

    if include_src_dir:
      src_path = os.path.join(
          self.pkg_name_to_pkg_src_dir(package_name),
          file_name)
      if os.path.exists(src_path):
        return src_path

    return None

