import ConfigParser
import os
import os.path

from buildutil.analysis_passes import (
    BindDependencies,
    BuildTargets,
    CheckCycles,
    TestTargets,
    )
from buildutil.config import Config
from buildutil.package import PackageSet
from buildutil.target_patterns import TargetPatterns


def main():
  ini_config = ConfigParser.ConfigParser()

  cfg_abs_path = os.path.expanduser('~/.buildutil_config.ini')
  if os.path.isfile(cfg_abs_path):
    ini_config.read(cfg_abs_path)

  config = Config(ini_config)

  pkgs = PackageSet(config)
  pkgs.get_or_load_all_subpackages('//')

  for name, pkg in pkgs.pkgs.items():
    print name
    for target in pkg.targets.values():
      print ' ', target.target_path()

  print '-' * 80

  pkgs = PackageSet(config)

  p = TargetPatterns('//buildutil')
  p.set_patterns(['//...'])
  #p.set_patterns(['//buildutil:buildutil'])

  passes = [
      BindDependencies(pkgs),
      CheckCycles(),
      BuildTargets(),
      TestTargets(),
      ]

  targets = p.get_matching_targets(pkgs)
  for p in passes:
    p.run(targets)

  #for t in targets:
  #  print t.target_path(), t.list_artifacts()

  """
  pkg = pkgs.get_or_load_package('//buildutil/rules')
  for target in pkg.targets.values():
    print target.target_path()
  """

  """
  target = pkgs.get_or_load_target('//buildutil/rules:py_rules')
  print target.target_path()
  """


if __name__ == '__main__':
  main()
