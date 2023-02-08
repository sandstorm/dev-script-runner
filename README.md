# Dev Script Runner

Use a shell script to document recurring development tasks and run them from anywhere inside your projects
folder structure using `dev some-task` while still being able to run them without this helper calling the 
script directly `./dev.sh some-task`.


<!-- TOC -->
* [Dev Script Runner](#dev-script-runner)
  * [Features](#features)
  * [Motivation](#motivation)
  * [Setup](#setup)
    * [Writing tasks](#writing-tasks)
    * [Passing arguments](#passing-arguments)
    * [Documenting Tasks](#documenting-tasks)
    * [Structuring tasks into different files](#structuring-tasks-into-different-files)
  * [Usage](#usage)
  * [Roadmap](#roadmap)
<!-- TOC -->

## Features

* run your development tasks from within a nested folder structure
* easy initialization of your project
* autocompletion
* documentation of your tasks will be used to provide help
* structure your tasks in different files

## Motivation

In all of our projects we have recurring tasks that require us running multiple commands, e.g. to

* set up the project initially
* start and stop containers
* enter containers
* inspect logs
* run test
* ...

These commands need to be documented somewhere. So why not have them documented in a shell script.
This way you can run them without the need of copy and pasting them.

We tried Makefiles, but they are complicated to use because of their syntax. It is also not possible
to copy commands from the Makefile and run them in the shell.

**Always changing back to the directory containing the dev script to run a task is annoying.
This helper allows running tasks from inside a nested folder structure ;)**

On the way up, the first `dev.sh` marked with a comment containing `DEV_SCRIPT_MARKER` is used to
execute a task.

We want the API of the helper to be the same as running the `dev.sh` script directly.
This is why we decided to NOT introduce a `run` argument for this helper.

This is how we ended up with the following API to run your tasks.

`./dev.sh sometask` or `dev sometask`

Some functionality will only be provided by the helper, e.g. running an init. 
As we do not want to confuse this additional functionality with running a task, 
we use UPPERCASE arguments prefixed with `DSR` for utils only provided by the `dev` command.

EXAMPLE: `dev DSR_INIT` to create the files needed in your project.

## Setup

run `brew install sandstorm/tap/dev-script-runner` to install

run `brew upgrade sandstorm/tap/dev-script-runner` to upgrade

**Autocompletion**

run `dev completion [bash|zsh|fish|powershell] --help` and follow instructions on how to set up autocompletion

> For **MacOS on ARM** the zsh instructions will not work. You have to add the 
> 
> `$(brew --prefix)/share/zsh/site-functions` to your `FPATH` 
> 
> before calling `autoload -U compinit; compinit`
> ```bash
> FPATH=“$(brew --prefix)/share/zsh/site-functions:${FPATH}”
> autoload -U compinit; compinit
> ```

**Initialization**

In your project root run `dev DSR_INIT`

This will create the files needed to start writing your own tasks. Run `dev` and you will see
example tasks that have been created for you. You can now start editing your `dev.sh` ;)


### Writing tasks

```bash
function sometask() {
  echo "TODO: implement sometask()"
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
function sometask() {
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
function sometask() {
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


## Usage

run `dev` for more information.

## Roadmap

* parse from comments: usage, examples, flags, params, ...
* add ready made tasks or provide copy paste examples
* CLI test
