import os
import os.path
import re

from functools import partial

from buildutil.util import (
    split_target_path,
    validate_pkg_path,
    )
from buildutil.target_patterns import TargetPatterns
from buildutil.rules.rules import RULES


CONFIG_FILE_NAME = 'BUILD'


class PackageSet(object):
  def __init__(self, config):
    self.config = config
    self.pkgs = {}

  def get_or_load_package(self, pkg_path):
    if pkg_path in self.pkgs:
      return self.pkgs[pkg_path]

    pkg = Package(pkg_path, self.config)
    pkg.load(self.config.pkg_path_to_src_abs_path(pkg_path))
    self.pkgs[pkg_path] = pkg
    return pkg

  def get_or_load_all_subpackages(self, pkg_path):
    pkgs = []

    pkg_dir = self.config.pkg_path_to_src_abs_path(pkg_path)
    for dir_name, _, file_names in os.walk(pkg_dir):
      if CONFIG_FILE_NAME not in file_names:
        continue

      sub_pkg_path = self.config.src_abs_path_to_pkg_path(dir_name)
      pkgs.append(self.get_or_load_package(sub_pkg_path))

    return pkgs

  def get_or_load_target(self, target_path):
    pkg_path, target_name = split_target_path(target_path)
    return self.get_or_load_package(pkg_path).get_target(target_name)


class Package(object):
  def __init__(self, pkg_path, config):
    assert validate_pkg_path(pkg_path), (
        'Invalid package full path: %s' % pkg_path)

    self.pkg_path = pkg_path
    self.config = config
    self.targets = {}
    self.visibility_patterns = TargetPatterns(self.pkg_path)

  def set_visibility(self, visibilty_patterns):
    self.visibility_patterns.set_patterns(visibility_set)

  def get_target(self, target_name):
    return self.targets[target_name]

  def get_all_targets(self):
    return self.targets.values()

  def register(self, target, ignore_duplicate=False):
    if not ignore_duplicate:
      assert target.name not in self.targets, (
          'Duplicate target name: %s (pkg: %s)' % (
              target.name,
              self.pkg_path))
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

