class BindDependencies(object):
  def __init__(self, pkgs):
    self.pkgs = pkgs

  def run(self, seed_target):
    frontier = [seed_target]
    while frontier:
      next_frontier = []
      for target in frontier:
        if target.binded:
          continue

        deps = target.dependency_patterns.get_matching_targets(self.pkgs)

        # filter test rules from library / binary targets.
        if not target.is_test_rule():
          deps = [t for t in deps if not t.is_test_rule()]

        for t in deps:
          assert t.is_visible_to(target), (
              '%s is not visible to %s' % (t.full_name(), target.full_name()))
        target.dependencies = {t.full_name(): t for t in deps}
        target.binded = True

        next_frontier.extend(deps)
      frontier = next_frontier


class CheckCycles(object):
  def __init__(self):
    pass

  def run(self, seed_target):
    stack = [seed_target]

    # NOTE: use DFS instead of topo sort for checking since I want to know the
    # exact cycle.
    while stack:
      for d in stack[-1].dependencies.values():
        assert d.in_cycle is not True
        if d.in_cycle is False:
          continue
        if d in stack:
          cycle = stack[stack.index(d):]
          for t in cycle:
            t.in_cycle = True
          assert False, ('Cycle detected: %s -> ...' %
              ' -> '.join([t.full_name() for t in cycle]))
        stack.append(d)
        break
      else:
        stack.pop().in_cycle = False
