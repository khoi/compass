[![CircleCI](https://circleci.com/gh/khoiln/sextant.svg?style=shield)](https://circleci.com/gh/khoiln/sextant) [![Go Report Card](https://goreportcard.com/badge/github.com/khoiln/sextant)](https://goreportcard.com/report/github.com/khoiln/sextant)
<img width="200" align="right" src="https://github.com/khoiln/sextant/blob/master/logo.svg">
# Sextant 
Sextant learns your habit, and help navigate to your "frecently used" directory.

## Usage
By default, `s` is the key-binding wrapper around `sextant`. 

- Fuzzily navigate to directory contains `go` and `sextant` :

```bash
s sext
# ~/Workspace/go/src/github.com/khoiln/sextant
```

- For more option refer to:

```bash
sextant --help
```

## Install

To install, use `go get`:

```bash
$ go get github.com/khoiln/sextant
```

Add this to the end of your `.zshrc` or `.bash_profile` 

```bash
eval "$(sextant shell)"
```

## Improvements

- [ ] Add support for child queries. For instance: `s go gallery` 
- [ ] Add shells auto completion
- [ ] Add a cleanup command to remove non-exist folder

