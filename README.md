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

We want the API to be the same when running the `dev.sh` script directly or by using the helper.
This is why we decided to NOT introduce a `run` argument for this helper.

When thinking about ...

`./dev.sh sometask` & `dev run sometask` vs.

`./dev.sh sometask` & `dev sometask` then

`./dev.sh sometask` & `dev sometask` **is the clear winner ;)**

Some functionality will only be provided by the helper, e.g. running an init. 
As we do not want to confuse this additional functionality with running a task, we use UPPERCASE arguments for the API.

Example: `dev INIT` to create the files needed in your project.

## Setup

run `brew install sandstorm/tap/dev-script-runner` to install
run `brew upgrade sandstorm/tap/dev-script-runner` to upgrade

Go to your project root and run `dev INIT`to create a `dev.js` with examples for different types of tasks.
The `$@` at the end of your `dev.sh` dispatches the script arguments to a function (so `dev sometask` calls `sometask`).

The script is only picked up by the helper if `DEV_SCRIPT_MARKER` is present in the file. 

## Usage

`dev <TASK_NAME>` to run a task
`dev INIT` to create the files needed in your project

## TODOs
* Autocompletion for tasks
