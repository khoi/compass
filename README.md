[![CircleCI](https://circleci.com/gh/khoiracle/sextant.svg?style=shield)](https://circleci.com/gh/khoiracle/sextant) [![Go Report Card](https://goreportcard.com/badge/github.com/khoiracle/sextant)](https://goreportcard.com/report/github.com/khoiracle/sextant)
<img width="200" align="right" src="https://github.com/khoiracle/sextant/blob/master/logo.svg">
# Sextant 
Sextant learns your habit, and help navigate to your "frecently used" directory.

## Usage
By default, `s` is the key-binding wrapper around `sextant`. 

- Fuzzily navigate to directory contains `go` and `sextant` :

```bash
s sext
# ~/Workspace/go/src/github.com/khoiracle/sextant
```

- For more option refer to:

```bash
sextant --help
```

## Install

Use Homebrew:

```bash
$ brew install khoiracle/tap/sextant
```

For development build:

```bash
$ go get github.com/khoiracle/sextant
```

Add this to the end of your `.zshrc` or `.bash_profile` 

```bash
eval "$(sextant shell)"
```

## Improvements

- [ ] Add support for child queries. For instance: `s go gallery` 
- [ ] Allow custom key binding

## References

- [rupa/z](https://github.com/rupa/z)
- [wting/autojump](https://github.com/wting/autojump)
- [gsamokovarov/jump](https://github.com/gsamokovarov/jump)
