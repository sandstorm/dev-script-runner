# Sandstorm Dev Script Runner

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
As we do not want to confuse this additional functionality with running a task, we use UPPERCASE 
arguments prefixed with `DSR` for utils only provided by the `dev` command.

Example: `dev DSR_INIT` to create the files needed in your project.

## Setup

run `brew install sandstorm/tap/dev-script-runner` to install

run `brew upgrade sandstorm/tap/dev-script-runner` to upgrade

run `dev completion [bash|zsh|fish|powershell] --help` and follow instructions on how to setup autocompletion

Go to your project root and run `dev DSR_INIT` to create a `dev.sh` and a `dev_setup.sh` with examples for different types of tasks.
The `$@` at the end of your `dev.sh` dispatches the script arguments to a function (so `dev sometask` calls `sometask`).

The script is only picked up by the helper if `DEV_SCRIPT_MARKER` is present in the file. 

### Writing tasks

```bash
function sometask() {
  echo "TODO: implement sometask()"
}
```
**Tasks starting with `_` are expected to be private and will be ignored**

**You should not use UPPERCASE tasks in your `dev.sh`**

### Passing arguments

You can run a task providing additional arguments that will be passed to your dev.sh
transparently. 

The only exception are the `-h` and `--help` flags that are handled by the DevScriptRunner.
You should not use them in your `dev.sh` script.

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

`$@` can be pretty useful for passing all arguments to other tasks in your `dev.sh`.

`$@` should always be present at the end of our your `dev.sh` script.

### Documenting Tasks

```bash
# Short description in first line
#
# Some more comments giving a detailed description about your task.
# Your description can span multiple lines. There MUST not be any
# empty lines for the comments to be associated with your task.
function sometask() {
  echo "TODO: implement"
}
```

This currently is WIP and will be improved in the future ;)

## Usage

run `dev` for more information.

## TODOs

* more testing
  * dev script parsing
  * e2e tests actually testing the CLI
* more features for documenting tasks -> e.g. support usage, examples, params, ...
