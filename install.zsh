#!/usr/bin/zsh

if [ -f ~/.zshrc ]; then
    echo "alias zs='/Users/$USER/zs/dist/zs'" >> ~/.zshrc
else
    echo "No .zshrc file found"
fi