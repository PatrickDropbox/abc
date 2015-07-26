import os
import os.path
import re

from functools import partial

from buildutil.util import (
    split_target_path,
    validate_pkg_full_path,
    )
from buildutil.target_patterns import TargetPatterns
from buildutil.rules.rules import RULES


CONFIG_FILE_NAME = 'BUILD'


class PackageSet(object):
  def __init__(self, config):
    self.config = config
    self.pkgs = {}

  def get_or_load_package(self, pkg_full_name):
    if pkg_full_name in self.pkgs:
      return self.pkgs[pkg_full_name]

    pkg = Package(pkg_full_name, self.config)
    pkg.load(self.config.pkg_name_to_pkg_src_dir(pkg_full_name))
    self.pkgs[pkg_full_name] = pkg
    return pkg

  def get_or_load_all_subpackages(self, pkg_full_name):
    pkgs = []

    pkg_dir = self.config.pkg_name_to_pkg_src_dir(pkg_full_name)
    for dir_name, _, file_names in os.walk(pkg_dir):
      if CONFIG_FILE_NAME not in file_names:
        continue

      subpkg_name = self.config.pkg_src_dir_to_pkg_name(dir_name)
      pkgs.append(self.get_or_load_package(subpkg_name))

    return pkgs

  def get_or_load_target(self, target_full_name):
    pkg_full_name, target_name = split_target_path(target_full_name)
    return self.get_or_load_package(pkg_full_name).get_target(target_name)


class Package(object):
  def __init__(self, pkg_full_path, config):
    assert validate_pkg_full_path(pkg_full_path), (
        'Invalid package full path: %s' % pkg_full_path)

    self.full_path = pkg_full_path
    self.config = config
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

  def register(self, target, ignore_duplicate=False):
    if not ignore_duplicate:
      assert target.name not in self.targets, (
          'Duplicate target name: %s (pkg: %s)' % (
              target.name,
              self.full_path))
    self.targets[target.name] = target

  def load(self, pkg_dir_abs_path):
    # TODO add package func to globals
    globals_dict = {}
    for rule in RULES:
      assert rule.rule_name() not in globals_dict
      globals_dict[rule.rule_name()] = partial(rule.register, pkg=self)

    config_file = os.path.join(pkg_dir_abs_path, CONFIG_FILE_NAME)
    try:
      # XXX: probably should restrict python exec here, but o well ...
      execfile(config_file, globals_dict, {})
    except Exception as e:
      print 'Failed to load', config_file, '-', e
      raise

