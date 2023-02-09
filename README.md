<h1 align="center">Dev Script Runner</h1>
<p align="center">Write and document recurring development tasks inside a shell script and run them from anywhere inside your projects folder structure.</p>

<p align="center">

<a style="text-decoration: none" href="https://github.com/sandstorm/dev-script-runner/releases">
<img src="https://img.shields.io/github/v/release/sandstorm/dev-script-runner?style=flat-square" alt="Latest Release">
</a>

<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

</p>

----

* [Features](#features)
* [Motivation](#motivation)
* [Setup](#setup)
  * [Writing tasks](#writing-tasks)
  * [Passing arguments](#passing-arguments)
  * [Documenting Tasks](#documenting-tasks)
  * [Structuring tasks into different files](#structuring-tasks-into-different-files)
* [Usage](#usage)
* [Roadmap](#roadmap)

----

## Features

* run your development tasks from within a nested folder structure
* easy initialization of your project
* autocompletion
* documentation of your tasks will be used to provide help
* structure your tasks in different files

## Motivation

We have recurring tasks in all our projects. Often they require multiple shell commands, e.g. for ...

* setting up a project
* starting and stopping containers
* running test
* ...

These tasks need to be documented, so why not have them documented in a shell script?
This way you can run them without the need of copy and pasting them from a Readme.

We tried Makefiles, but they are complicated to use. It is also not possible to copy commands 
from the Makefile and run them in the shell.

Always changing back to the root of your project to run a script can be annoying. This is why we want 
to run tasks from inside a nested folder structure.

Running tasks using the helper should be the same as running tasks directly using the script.

This is how we ended up with the following API.

`./dev.sh sometask` vs. `dev sometask`

Functionality only provided by the helper will be uppercase and prefixed with `DSR`.

Example: `dev DSR_INIT` to create the files needed in your project.

## Setup

run `brew install sandstorm/tap/dev-script-runner` to install

run `brew upgrade sandstorm/tap/dev-script-runner` to upgrade

**Autocompletion**

run `dev completion [bash|zsh|fish|powershell] --help` and follow instructions on how to set up autocompletion

> For **MacOS on ARM** the zsh instructions will not work. You have to change your `FPATH` first.
>
> ```bash
> # .zshrc
> 
> FPATH="$(brew --prefix)/share/zsh/site-functions:${FPATH}"
> # compinit MUST be called afterwards!
> autoload -U compinit; compinit
> ```
> Now run 
> 
> ```
> dev completion zsh > $(brew --prefix)/share/zsh/site-functions/_dev
> ```
> 
> and restart your terminal.

**Initialization**

In your project root run `dev DSR_INIT`

This will create the files needed to start writing your own tasks. Run `dev` and you will see
example tasks that have been created for you. You can now start editing your `dev.sh` ;)

### Writing tasks

```bash
function sometask {
  echo "TODO: implement sometask"
}
```
**Tasks starting with `_` are expected to be private and will be ignored**

**You should not use UPPERCASE tasks in your `dev.sh`** as they might be used to provide
utilities. e.g. `dev DSR_INIT` to init your folder with all the files needed by the 
Dev Script Runner

### Passing arguments

You can run a task providing additional arguments that will be passed to your `dev.sh`
transparently. 

The only exception are the `-h` and `--help` flags that are handled by the DevScriptRunner.
If you implement them in your `dev.sh` script they will only work if you call your script
directly. `./dev.sh --help` will run your own implementation.

```bash
dev sometask arg1 arg2 agr3
```

```bash
function sometask {
  echo "$@" # -> "arg1 arg2 agr3"
  echo "$1" # -> "arg1"
  echo "$2" # -> "arg2"
  echo "$3" # -> "arg3"
}
```

`$@` can be used for passing all arguments to other tasks in your `dev.sh`.

`$@` MUST always be present at the end of our your `dev.sh` script.

### Documenting Tasks

```bash
# Short description in first line
#
# Some more comments giving a detailed description about your task.
# Your description can span multiple lines. There MUST NOT be any empty lines 
# between the comment block and your task.
function sometask {
  echo "TODO: implement"
}
```

### Structuring tasks into different files

```bash
#!/bin/bash
############################## DEV_SCRIPT_MARKER ##############################

# You can structure your tasks into different files using `source`. We will also
# parse these files. The order will be: 
#   1. tasks of the dev.sh
#   2. tasks of sourced files
source ./dev_utilities.sh
source ./dev_tasks_testing.sh
source ./dev_tasks_release.sh
```

## Roadmap

* parse from comments: usage, examples, flags, params, ...
* add ready made tasks or provide copy paste examples
* CLI test
