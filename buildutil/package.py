import os
import os.path
import re

from functools import partial

from buildutil.util import (
    pkg_dir_to_pkg_name,
    pkg_name_to_pkg_dir,
    split_target_path,
    validate_pkg_full_path,
    )
from buildutil.target_patterns import TargetPatterns
from buildutil.rules.rules import RULES


CONFIG_FILE_NAME = 'BUILD'


class PackageSet(object):
  def __init__(self, root_dir_abs_path):
    root_dir_abs_path = os.path.abspath(root_dir_abs_path)
    assert os.path.isdir(root_dir_abs_path)

    self.root_dir = root_dir_abs_path


    self.pkgs = {}

  def get_or_load_package(self, pkg_full_name):
    if pkg_full_name in self.pkgs:
      return self.pkgs[pkg_full_name]

    pkg = Package(pkg_full_name)
    pkg.load(pkg_name_to_pkg_dir(pkg_full_name, self.root_dir))
    self.pkgs[pkg_full_name] = pkg
    return pkg

  def get_or_load_all_subpackages(self, pkg_full_name):
    pkgs = []

    pkg_dir = pkg_name_to_pkg_dir(pkg_full_name, self.root_dir)
    for dir_name, _, file_names in os.walk(pkg_dir):
      if CONFIG_FILE_NAME not in file_names:
        continue

      subpkg_name = pkg_dir_to_pkg_name(dir_name, self.root_dir)
      pkgs.append(self.get_or_load_package(subpkg_name))

    return pkgs

  def get_or_load_target(self, target_full_name):
    pkg_full_name, target_name = split_target_path(target_full_name)
    return self.get_or_load_package(pkg_full_name).get_target(target_name)


class Package(object):
  def __init__(self, pkg_full_path):
    assert validate_pkg_full_path(pkg_full_path), (
        'Invalid package full path: %s' % pkg_full_path)

    self.full_path = pkg_full_path
    self.targets = {}
    self.visibility_patterns = TargetPatterns(self.full_path)

  def set_visibility(self, visibilty_patterns):
    self.visibility_patterns.set_patterns(visibility_set)

  def get_target(self, target_name):
    return self.targets[target_name]

  def get_all_targets(self):
    return self.targets.values()

  def path():
    return self.full_path

  def load(self, pkg_dir_abs_path):
    # TODO add package func to globals
    globals_dict = {}
    for rule in RULES:
      def func(**kwargs):
        target = rule(pkg=self, **kwargs)
        assert target.name not in self.targets, (
            'Duplicate target name: %s (pkg: %s)' % (
                target.name,
                self.full_path))
        self.targets[target.name] = target
      func.__name__ = rule.rule_name()

      assert rule.rule_name() not in globals_dict
      globals_dict[rule.rule_name()] = func

    config_file = os.path.join(pkg_dir_abs_path, CONFIG_FILE_NAME)
    try:
      # XXX: probably should restrict python exec here, but o well ...
      execfile(config_file, globals_dict, {})
    except Exception as e:
      print 'Failed to load', config_file, '-', e
      raise

