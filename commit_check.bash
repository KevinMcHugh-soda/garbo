#!/usr/bin/env bash

if git log --pretty=format:"%H" -- "README.md" | xargs git show | grep -q "Please do not commit this line"; then
  say -v Jester "Hey"
fi

if git log --pretty=format: --name-only | grep -q additional_file.txt; then
  say -v "Bad News" "Hey I think you included a file you shouldn't have"
fi
