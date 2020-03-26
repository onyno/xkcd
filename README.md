# xkcd password generator

xkcd is a command line tool to generate passwords in the style of [XKCD Password Strength](https://xkcd.com/936/).

## Usage

Run with the system dictionary words, `/usr/share/dict/words`

```
Usage of xkcd:
  -common
    	use common words (default false)
  -max int
    	max length of words (default 8)
  -min int
    	min length of words (default 3)
  -n int
    	number of words (default 4)
```

## Examples

Default using the system dictionary of words

```
$ xkcd
string-economy-debt-floor
```

Use the common words

```
$ xkcd -common -n 3
breath-operate-mount
```

Use smaller common words

```
$ xkcd -common -n 3 -max 5
fish-seed-nice
```

Use longer common words

```
$ xkcd -common -n 3 -min 5 -max 15
agricultural-following-extension
```

## Installation

Using xkcd is easy. First, use `go get` to install the latest version of the binary.

```
go get -u github.com/onyno/xkcd
```
