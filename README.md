# calcium

## This is task runner inspired by Makefile

## Installation

```sh
# Please set $GOPATH
$ git clone https://github.com/NasSilverBullet/calcium
$ go install
```

## Usage

```sh
# Please create calcium.yml on your working directory
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

# call task: test1
$ calcium run test1
test # echo test

# call task: test2
$ calcium run test2 -v foo -sv bar
foo # echo {{value}} => echo foo
bar # echo {{secondvalue}} => echo bar

# call faild task: test2
$ calcium run test2 -v foo
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
