import os

from buildutil.analysis_passes import (
    BindDependencies,
    CheckCycles,
    )
from buildutil.package import PackageSet
from buildutil.target_patterns import TargetPatterns
from buildutil.topo_sorter import TopoSorter


def main():
  pkgs = PackageSet(os.getcwd())
  pkgs.get_or_load_all_subpackages('//')

  for name, pkg in pkgs.pkgs.items():
    print name
    for target in pkg.targets.values():
      print ' ', target.full_name()

  print '-' * 80

  pkgs = PackageSet(os.getcwd())

  p = TargetPatterns('//buildutil')
  p.set_patterns(['//...'])
  #p.set_patterns(['//buildutil:buildutil'])

  passes = [
      BindDependencies(pkgs),
      CheckCycles(),
      ]

  sorter = TopoSorter()

  targets = p.get_matching_targets(pkgs)
  for target in targets:
    for p in passes:
      p.run(target)
    order = sorter.sort(target)
    print target.full_name()
    print '   ', [[t.full_name() for t in l] for l in order]

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
