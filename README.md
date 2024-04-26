# WELCOME TO GARBO SEARCH
It's a garbage lil search tool

## Install Steps:
* Open VS Code
    * CMD+Shift+P: Install 'Code' command in path
* Set your editor:
    * Open a terminal
    * `code ~/.zprofile`
    * Add `export EDITOR='code'` to `~/.zprofile`
    * save `.zprofile`
    * `source ~/.zprofile`
* Copy my favorite gitconfig options:
```
[alias]
	ls = log --pretty=format:"%C(yellow)%h%Cred%d\\ %Creset%s%Cblue\\ [%cn]" --decorate
	g = grep
	st = status
	ap = add -p
	co = checkout
	df = diff
	lg = "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
	find = ls-files
	wc = whatchanged
	sw = switch
	pdiff = diff --color=always | less -R
	fza = "!git ls-files -m -o --exclude-standard | fzf -m --print0 | xargs -0 git add"
	undo-commit = "reset --soft 'HEAD^'"
```
    * git config --global -e


## TODO:
* ~~Build the basic search tool~~
* ~~Parameterize the file to search~~
* ~~Parameterize the search value~~
* ~~Parameterize the search field~~
* ~~Support dynamic search fields~~
* Support searching multiple fields
