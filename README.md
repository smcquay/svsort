# svsort

Sort lines of text that begin with semver.

## Example

```bash

$ git tag -n
0.0.1           Merge pull request #55 from prometheus/update-makefile
0.0.2           Merge pull request #73 from prometheus/version-0.0.2
0.0.3           Merge pull request #76 from prometheus/v0.0.3
0.0.4           Cut version 0.0.4.
0.1.0           Bump version to 0.1.0
0.1.0-alpha     Enable YAML highlighting in example config
0.1.0-beta0     Bump version to 0.1.0-beta0
0.1.0-beta1     Bump beta version
0.1.0-beta2     Add 0.1.0beta2 changelog
0.1.0beta2      Add 0.1.0beta2 changelog
0.1.1           Bump version to 0.1.1
0.2.0           Merge pull request #386 from prometheus/fabxc-0.2.0
0.2.1           Merge pull request #405 from prometheus/fabxc-0.2.1
v0.10.0         v0.10.0
v0.11.0         0.11.0 (#1109)
v0.3.0          Merge pull request #422 from prometheus/fabxc-0.3.0
v0.4.0          Bump version to 0.4 (#469)
v0.4.1          v0.4.1 / 2016-08-31
v0.4.2          v0.4.2 / 2016-09-02
v0.5.0          Cut 0.5.0 (#543)
v0.5.0-alpha.0  *: fix version
v0.5.0-beta.0   Merge pull request #527 from brancz/master
v0.5.1          *: cut v0.5.1 (#569)
v0.6.0          Cut v0.6.0
v0.6.1          v0.6.1
v0.6.2          *: cut v0.6.2 (#777)
v0.7.0          v0.7.0
v0.7.0-rc.0     v0.7.0-rc.0
v0.7.1          v0.7.1
v0.8.0          v0.8.0
v0.9.0          Add release-0.9 (#1004)
v0.9.1          v0.9.1 (#1010)

$ git tag -n | svsort
0.0.1           Merge pull request #55 from prometheus/update-makefile
0.0.2           Merge pull request #73 from prometheus/version-0.0.2
0.0.3           Merge pull request #76 from prometheus/v0.0.3
0.0.4           Cut version 0.0.4.
0.1.0-alpha     Enable YAML highlighting in example config
0.1.0-beta0     Bump version to 0.1.0-beta0
0.1.0-beta1     Bump beta version
0.1.0-beta2     Add 0.1.0beta2 changelog
0.1.0           Bump version to 0.1.0
0.1.1           Bump version to 0.1.1
0.2.0           Merge pull request #386 from prometheus/fabxc-0.2.0
0.2.1           Merge pull request #405 from prometheus/fabxc-0.2.1
v0.3.0          Merge pull request #422 from prometheus/fabxc-0.3.0
v0.4.0          Bump version to 0.4 (#469)
v0.4.1          v0.4.1 / 2016-08-31
v0.4.2          v0.4.2 / 2016-09-02
v0.5.0-alpha.0  *: fix version
v0.5.0-beta.0   Merge pull request #527 from brancz/master
v0.5.0          Cut 0.5.0 (#543)
v0.5.1          *: cut v0.5.1 (#569)
v0.6.0          Cut v0.6.0
v0.6.1          v0.6.1
v0.6.2          *: cut v0.6.2 (#777)
v0.7.0-rc.0     v0.7.0-rc.0
v0.7.0          v0.7.0
v0.7.1          v0.7.1
v0.8.0          v0.8.0
v0.9.0          Add release-0.9 (#1004)
v0.9.1          v0.9.1 (#1010)
v0.10.0         v0.10.0
v0.11.0         0.11.0 (#1109)
```

