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

For fish shell add the line below to your `~/.config/fish/config.fish`

```bash
if type -q sextant
  status --is-interactive; and source (sextant shell --type fish -|psub)
end
```

If you want to use different key binding pass `--bind-to` to the `sextant shell` command:

For instance, if you want to use `z` instead of `s`

```bash
eval "$(sextant shell --bind-to z)"
```

## Improvements

- [ ] Add support for child queries. For instance: `s go gallery` 

## References

- [rupa/z](https://github.com/rupa/z)
- [wting/autojump](https://github.com/wting/autojump)
- [gsamokovarov/jump](https://github.com/gsamokovarov/jump)
