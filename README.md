# vimalter
[![CircleCI](https://circleci.com/gh/tennashi/vimalter/tree/master.svg?style=shield)](https://circleci.com/gh/tennashi/vimalter/tree/master)  
Switch commands between inside and outside of the vim/nvim.

# Screenshot
## In neovim
[![asciicast](https://asciinema.org/a/qmUVLV7e93kLWgBcI3ARt64vt.svg)](https://asciinema.org/a/qmUVLV7e93kLWgBcI3ARt64vt)

## In vim8 with `--remote`
[![asciicast](https://asciinema.org/a/045L59uL9XthDB6UBStzP1ojC.svg)](https://asciinema.org/a/045L59uL9XthDB6UBStzP1ojC)

## In vim8 with `terminal-api`(partial support)
[![asciicast](https://asciinema.org/a/4UM372nJ5LY65SKWNLhVqvcls.svg)](https://asciinema.org/a/4UM372nJ5LY65SKWNLhVqvcls)

# Usage
```bash
$ vimalter [option] [file ...]
```

## Options
* `--tab`: Open specified file in a new tab when executed from terminal mode of vim/nvim.

## Neovim
* Please install [mhinz/neovim-remote](https://github.com/mhinz/neovim-remote) in the `$PATH`

## Vim8 with `--remote`
* Please install or build the vim with `+serverclient` support.

## Vim8 `terminal-api` support
Vim `terminal-api` support is partial.
* `--tab` is not supported.
* No argument(i.e. you want to start editing an empty file in new buffer) is not supported.
