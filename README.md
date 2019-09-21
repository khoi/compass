[![Action](https://github.com/khoi/compass/workflows/test/badge.svg)](https://github.com/khoi/compass/actions)
[![@khoiracle](https://img.shields.io/badge/contact-@khoiracle-4bbee3.svg?style=flat)](https://twitter.com/khoiracle)
<img width="200" align="right" src="https://github.com/khoi/compass/blob/master/logo.svg">
# Compass
Compass learns your habit, and help navigate to your "frecently used" directory.

## Usage
By default, `s` is the key-binding wrapper around `compass`.

- Fuzzily navigate to directory contains `go` and `compass` :

```bash
s compass
# ~/Workspace/go/src/github.com/khoi/compass
```

- For more option refer to:

```bash
compass --help
```

## Install

Use Homebrew:

```bash
brew install khoi/tap/compass
```

For development build:

```bash
go get github.com/khoi/compass
```

Add this to the end of your `.zshrc` or `.bash_profile` 

```bash
eval "$(compass shell)"
```

For fish shell add the line below to your `~/.config/fish/config.fish`

```bash
if type -q compass
  status --is-interactive; and source (compass shell --type fish -|psub)
end
```

If you want to use different key binding pass `--bind-to` to the `compass shell` command:

For instance, if you want to use `z` instead of `s`

```bash
eval "$(compass shell --bind-to z)"
```

## Improvements

- [ ] Add support for child queries. For instance: `s go gallery` 

## References

- [rupa/z](https://github.com/rupa/z)
- [wting/autojump](https://github.com/wting/autojump)
- [gsamokovarov/jump](https://github.com/gsamokovarov/jump)
