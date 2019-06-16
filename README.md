# vimalter
[![CircleCI](https://circleci.com/gh/tennashi/vimalter/tree/master.svg?style=shield)](https://circleci.com/gh/tennashi/vimalter/tree/master)
Switch commands between inside and outside of the vim/nvim.

# Screenshot
## In neovim
[![asciicast](https://asciinema.org/a/qmUVLV7e93kLWgBcI3ARt64vt.svg)](https://asciinema.org/a/qmUVLV7e93kLWgBcI3ARt64vt)

## In vim8 with `--remote`
[![asciicast](https://asciinema.org/a/045L59uL9XthDB6UBStzP1ojC.svg)](https://asciinema.org/a/045L59uL9XthDB6UBStzP1ojC)

## In vim8 with `terminal-api`
[![asciicast](https://asciinema.org/a/4UM372nJ5LY65SKWNLhVqvcls.svg)](https://asciinema.org/a/4UM372nJ5LY65SKWNLhVqvcls)

# Install
Get from the [release page](https://github.com/tennashi/vimalter/releases)(recomended)

or

```shell
$ go get -u github.com/tennashi/vimalter
```

## Neovim
* Please install [mhinz/neovim-remote](https://github.com/mhinz/neovim-remote) in the `$PATH`

## Vim8 with `--remote`
* Please install or build the vim with `+clientserver` support.

## Vim8 `terminal-api`
* Please install [tennashi/termopen.vim](https://github.com/tennashi/termopen.vim)

# Usage
```bash
$ vimalter [option] [file ...]
```

## Options
* `-tab`: Open specified file in a new tab when executed from terminal mode of vim/nvim.

## Neovim & vim8 with `--remote`
* If you start it as the default editor, such as `git commit`, you need to exit with `: w | bd`.
* If you don't like this, put the following in your vimrc.(see. [mhinz/neovim-remote](https://github.com/mhinz/neovim-remote#typical-use-cases))
```vim
autocmd FileType gitcommit set bufhidden=delete
```
