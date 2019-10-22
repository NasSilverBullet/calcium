# calcium

## this is task runner inspired by Makefile

## Installation

```sh
# Please set $GOPATH
$ git clone https://github.com/NasSilverBullet/calcium
$ go install
```

## Usage

```sh
$ cat calcium.yml
version: 1

tasks:

  - task:
    use: test1
    run: |
      echo test

  - task:
    use: test2
    flags:
      - name: value
        short: v
        long: val
        description: for echo value

      - name: secondvalue
        short: sv
        long: secval
        description: for echo second value
    run: |
      echo {{value}}
      echo {{secondvalue}}

$ calcium run test1 # call task: test1
test

$ calcium run test2 -v foo -sv bar # call task: test2
foo
bar

$ calcium run test2 -v foo # call faild task: test2
Error:
Missing flags: [secondvalue]

Usage:
  calcium run test2 [flags]

Flags:
  -v,  --val      for echo value
  -sv, --secval   for echo second value
```

## License

MIT License. See LICENSE.txt for more information.
