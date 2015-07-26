import os

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
  config = Config(os.getcwd() + '/..', 'src', 'genfile', 'build')

  pkgs = PackageSet(config)
  pkgs.get_or_load_all_subpackages('//')

  for name, pkg in pkgs.pkgs.items():
    print name
    for target in pkg.targets.values():
      print ' ', target.full_name()

  print '-' * 80

  pkgs = PackageSet(config)

  p = TargetPatterns('//buildutil')
  p.set_patterns(['//...'])
  #p.set_patterns(['//buildutil:buildutil'])

  passes = [
      BindDependencies(pkgs),
      CheckCycles(),
      BuildTargets(),
      ]

  targets = p.get_matching_targets(pkgs)
  for p in passes:
    p.run(targets)

  """
  pkg = pkgs.get_or_load_package('//buildutil/rules')
  for target in pkg.targets.values():
    print target.full_name()
  """

  """
  target = pkgs.get_or_load_target('//buildutil/rules:py_rules')
  print target.full_name()
  """


if __name__ == '__main__':
  main()
