<div align="center">
<h1>zs - A better <code>ls</code> command</h1>
<img src="https://img.shields.io/badge/Language-Golang-informational?logo=go&style=plastic&logoColor=333333&color=666666&labelColor=ffff">
<img src="https://img.shields.io/badge/Made%20For-ZSH-informational?logo=go&style=plastic&logoColor=333333&color=666666&labelColor=ffff">
</div>

# ZS

**zs** is a replacement to the `ls` command, it lists directories and files in a nicely formatted and color coded way.

![](https://github.com/myferr/zs/blob/main/img.png?raw=true)

# Documentation

## Using arguments

To use a configuration argument you pass your directory argument and then pass your configuration argument.

**Example: `zs <directory> <argument>`**

You have to pass your directory argument before you pass your configuration argument.

## Arguments

A list of configuration arguments.

- `--show-hidden` When this argument is passed it will list the **.DS_Store** file.
- `--show-git` When this argument is passed it will list the **.git** directory and all subdirectories inside as well as files.
- `--show-node-modules` When this argument is passed it will list the **node_modules** directory and all contents found in the **node_modules** directory.

## Installation

Requirements:

- ZSH
- A `.zshrc` file in your `~` directory
- [Golang](https://go.dev)
- Git

Instructions:

1. Navigate to `/Users/$USERNAME`
2. Run the command `git clone https://github.com/myferr/zs`
3. Run the command `bash /Users/$HOME/zs/install.zsh`
