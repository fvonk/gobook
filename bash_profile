source ~/.profile
export GOPATH=~/go/src/gobook/

export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8

export PATH=$GOPATH/bin:$PATH
# Git Shortcuts
alias gsu='git submodule update'
alias g='git'
alias gs='git status'
alias gst='git status -sb'
alias ga='git add'
alias gau='git add -u' # Removes deleted files
alias gp='git pull'
alias gpu='git push'
alias gc='git commit -a -m'
alias gca='git commit -v -a' # Does both add and commit in same command, add -m 'blah' for comment
alias gco='git checkout'
alias gcoa='git checkout -- .'
alias gl='git log --oneline'
#alias ghist='hist = log --graph --all --pretty=format:'%Cred%h%Creset %ad %s %C\(yellow\)%d%Creset %C\(bold blue\)<%ad>%Creset' --date=shor'
alias gb='git branch -av'
alias gam='git commit -a --amend'
alias gla='git log --stat --no-merges --max-count=1'
alias gprd='git pull --rebase origin dev'
alias gprm='git pull --rebase origin master'
alias gf='git fetch -ap --tags'
alias ggc='git gc && git fsck'
alias gd='git diff'

[[ -s "$HOME/.rvm/scripts/rvm" ]] && source "$HOME/.rvm/scripts/rvm" # Load RVM into a shell session *as a function*
