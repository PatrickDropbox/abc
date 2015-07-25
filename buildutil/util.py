import os.path
import re


# XXX: maybe allow % and $
TARGET_PATTERN_REGEX = re.compile(
    '^(?:(?:(?://)?(?:\w+/)*\.\.\.)|(?:(?://)?(?:(?:\w+/)*\w+)?:\w+(\.\w+)*))$')
FULL_PKG_NAME_REGEX = re.compile('^//((?:\w+/)*\w+)?$')
TARGE_NAME_REGEX = re.compile('^\w+(\.\w+)*$')

def validate_target_pattern(pattern):
  return TARGET_PATTERN_REGEX.match(pattern) is not None


def validate_pkg_full_path(path):
  return FULL_PKG_NAME_REGEX.match(path) is not None


def validate_target_name(name):
  if name == 'all':
    return False

  return TARGE_NAME_REGEX.match(name) is not None


def split_target_path(target_path):
  pkg_name, _, target_name = target_path.partition(':')
  return pkg_name, target_name


def pkg_dir_to_pkg_name(pkg_dir, root_dir):
  if pkg_dir == root_dir:
    return '//'

  if not root_dir.endswith('/'):
    root_dir += '/'

  assert pkg_dir.startswith(root_dir)
  return '//' + pkg_dir[len(root_dir):]


def pkg_name_to_pkg_dir(pkg_name, root_dir):
  assert pkg_name.startswith('//'), pkg_name
  return os.path.join(root_dir, pkg_name[2:])


def pkg_name_join(curr_pkg, rel_pkg):
  if rel_pkg.startswith('//'):
    return rel_pkg
  if rel_pkg.startswith(':'):
    return curr_pkg + rel_pkg
  return curr_pkg + '/' + rel_pkg
