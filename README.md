# calcium

## this is task runner inspired by Makefile

## Installation

```sh
$ git clone https://github.com/NasSilverBullet/calcium
$ go install
```

## Usage

```sh
$ cat calcium.yml
version: 1

tasks:

  - task:
    use: "test1"
    run: |
      echo test

  - task:
    use: "test2"
    flags:
      - name: value
        short: v
        long: val
    run: |
      echo {{value}}

$ calcium run test1 # call task test1
test

$ calcium run test2 -v hoge # call task test2
hoge
```

## License

MIT License. See LICENSE.txt for more information.
