import unittest

from buildutil.target_patterns import validate_target


class TargetPatternTestCase(unittest.TestCase):
  def test_target_regex(self):
    self.assertTrue(validate_target('...'))
    self.assertTrue(validate_target('bar/...'))
    self.assertTrue(validate_target('foo/bar/...'))
    self.assertTrue(validate_target('//...'))
    self.assertTrue(validate_target('//foo/...'))
    self.assertTrue(validate_target('//foo/bar/...'))

    self.assertTrue(validate_target(':all'))
    self.assertTrue(validate_target('foo:all'))
    self.assertTrue(validate_target('foo/bar:all'))
    self.assertTrue(validate_target('//:all'))
    self.assertTrue(validate_target('//foo:all'))
    self.assertTrue(validate_target('//foo/bar:zzz'))

    self.assertFalse(validate_target('//asdf/../foo'))
    self.assertFalse(validate_target('/...'))
    self.assertFalse(validate_target('/bar/...'))
    self.assertFalse(validate_target('/bar:all'))

if __name__ == '__main__':
    unittest.main()
