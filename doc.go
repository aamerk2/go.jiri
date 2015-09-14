// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated via go generate.
// DO NOT UPDATE MANUALLY

/*
Command jiri is a multi-purpose tool for multi-repo development.

Usage:
   jiri [flags] <command>

The jiri commands are:
   cl           Manage project changelists
   contributors List project contributors
   project      Manage the jiri projects
   rebuild      Rebuild the jiri command line tool
   snapshot     Manage project snapshots
   update       Update all jiri tools and projects
   version      Print version
   api          Manage vanadium public API
   copyright    Manage vanadium copyright
   env          Print vanadium environment variables
   go           Execute the go tool using the vanadium environment
   goext        Vanadium extensions of the go tool
   oncall       Manage vanadium oncall schedule
   profile      Manage vanadium profiles
   run          Run an executable using the vanadium environment
   test         Manage vanadium tests
   help         Display help for commands or topics

The jiri flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

The global flags are:
 -metadata=<just specify -metadata to activate>
   Displays metadata for the program and exits.

Jiri cl - Manage project changelists

Manage project changelists.

Usage:
   jiri cl <command>

The jiri cl commands are:
   cleanup     Clean up changelists that have been merged
   mail        Mail a changelist for review
   new         Create a new local branch for a changelist
   sync        Bring a changelist up to date

Jiri cl cleanup - Clean up changelists that have been merged

Command "cleanup" checks that the given branches have been merged into the
corresponding remote branch. If a branch differs from the corresponding remote
branch, the command reports the difference and stops. Otherwise, it deletes the
given branches.

Usage:
   jiri cl cleanup [flags] <branches>

<branches> is a list of branches to cleanup.

The jiri cl cleanup flags are:
 -f=false
   Ignore unmerged changes.
 -remote-branch=master
   Name of the remote branch the CL pertains to.

Jiri cl mail - Mail a changelist for review

Command "mail" squashes all commits of a local branch into a single "changelist"
and mails this changelist to Gerrit as a single commit. First time the command
is invoked, it generates a Change-Id for the changelist, which is appended to
the commit message. Consecutive invocations of the command use the same
Change-Id by default, informing Gerrit that the incomming commit is an update of
an existing changelist.

Usage:
   jiri cl mail [flags]

The jiri cl mail flags are:
 -autosubmit=false
   Automatically submit the changelist when feasiable.
 -cc=
   Comma-seperated list of emails or LDAPs to cc.
 -check-uncommitted=true
   Check that no uncommitted changes exist.
 -d=false
   Send a draft changelist.
 -edit=true
   Open an editor to edit the CL description.
 -host=https://vanadium-review.googlesource.com/
   Gerrit host to use.
 -m=
   CL description.
 -presubmit=all
   The type of presubmit tests to run. Valid values: none,all.
 -r=
   Comma-seperated list of emails or LDAPs to request review.
 -remote-branch=master
   Name of the remote branch the CL pertains to.
 -set-topic=true
   Set Gerrit CL topic.
 -topic=
   CL topic, defaults to <username>-<branchname>.

Jiri cl new - Create a new local branch for a changelist

Command "new" creates a new local branch for a changelist. In particular, it
forks a new branch with the given name from the current branch and records the
relationship between the current branch and the new branch in the .v23 metadata
directory. The information recorded in the .v23 metadata directory tracks
dependencies between CLs and is used by the "jiri cl sync" and "jiri cl mail"
commands.

Usage:
   jiri cl new <name>

<name> is the changelist name.

Jiri cl sync - Bring a changelist up to date

Command "sync" brings the CL identified by the current branch up to date with
the branch tracking the remote branch this CL pertains to. To do that, the
command uses the information recorded in the .v23 metadata directory to identify
the sequence of dependent CLs leading to the current branch. The command then
iterates over this sequence bringing each of the CLs up to date with its
ancestor. The end result of this process is that all CLs in the sequence are up
to date with the branch that tracks the remote branch this CL pertains to.

NOTE: It is possible that the command cannot automatically merge changes in an
ancestor into its dependent. When that occurs, the command is aborted and prints
instructions that need to be followed before the command can be retried.

Usage:
   jiri cl sync [flags]

The jiri cl sync flags are:
 -remote-branch=master
   Name of the remote branch the CL pertains to.

Jiri contributors - List project contributors

Lists project contributors. Projects to consider can be specified as an
argument. If no projects are specified, all projects in the current manifest are
considered by default.

Usage:
   jiri contributors [flags] <projects>

<projects> is a list of projects to consider.

The jiri contributors flags are:
 -aliases=
   Path to the aliases file.
 -n=false
   Show number of contributions.

Jiri project - Manage the jiri projects

Manage the jiri projects.

Usage:
   jiri project <command>

The jiri project commands are:
   clean        Restore jiri projects to their pristine state
   list         List existing jiri projects and branches
   shell-prompt Print a succinct status of projects, suitable for shell prompts
   poll         Poll existing jiri projects

Jiri project clean - Restore jiri projects to their pristine state

Restore jiri projects back to their master branches and get rid of all the local
branches and changes.

Usage:
   jiri project clean [flags] <project ...>

<project ...> is a list of projects to clean up.

The jiri project clean flags are:
 -branches=false
   Delete all non-master branches.

Jiri project list - List existing jiri projects and branches

Inspect the local filesystem and list the existing projects and branches.

Usage:
   jiri project list [flags]

The jiri project list flags are:
 -branches=false
   Show project branches.
 -nopristine=false
   If true, omit pristine projects, i.e. projects with a clean master branch and
   no other branches.

Jiri project shell-prompt

Reports current branches of jiri projects (repositories) as well as an
indication of each project's status:
  *  indicates that a repository contains uncommitted changes
  %  indicates that a repository contains untracked files

Usage:
   jiri project shell-prompt [flags]

The jiri project shell-prompt flags are:
 -check-dirty=true
   If false, don't check for uncommitted changes or untracked files. Setting
   this option to false is dangerous: dirty master branches will not appear in
   the output.
 -show-name=false
   Show the name of the current repo.

Jiri project poll - Poll existing jiri projects

Poll jiri projects that can affect the outcome of the given tests and report
whether any new changes in these projects exist. If no tests are specified, all
projects are polled by default.

Usage:
   jiri project poll [flags] <test ...>

<test ...> is a list of tests that determine what projects to poll.

The jiri project poll flags are:
 -manifest=
   Name of the project manifest.

Jiri rebuild - Rebuild the jiri command line tool

Rebuild the jiri command line tool.

Usage:
   jiri rebuild

Jiri snapshot - Manage project snapshots

The "jiri snapshot" command can be used to manage project snapshots. In
particular, it can be used to create new snapshots and to list existing
snapshots.

The command-line flag "-remote" determines whether the command pertains to
"local" snapshots that are only stored locally or "remote" snapshots the are
revisioned in the manifest repository.

Usage:
   jiri snapshot [flags] <command>

The jiri snapshot commands are:
   create      Create a new project snapshot
   list        List existing project snapshots

The jiri snapshot flags are:
 -remote=false
   Manage remote snapshots.

Jiri snapshot create - Create a new project snapshot

The "jiri snapshot create <label>" command captures the current project state in
a manifest and, depending on the value of the -remote flag, the command either
stores the manifest in the local $JIRI_ROOT/.snapshots directory, or in the
manifest repository, pushing the change to the remote repository and thus making
it available globally.

Internally, snapshots are organized as follows:

 <snapshot-dir>/
   labels/
     <label1>/
       <label1-snapshot1>
       <label1-snapshot2>
       ...
     <label2>/
       <label2-snapshot1>
       <label2-snapshot2>
       ...
     <label3>/
     ...
   <label1> # a symlink to the latest <label1-snapshot*>
   <label2> # a symlink to the latest <label2-snapshot*>
   ...

NOTE: Unlike the jiri tool commands, the above internal organization is not an
API. It is an implementation and can change without notice.

Usage:
   jiri snapshot create [flags] <label>

<label> is the snapshot label.

The jiri snapshot create flags are:
 -time-format=2006-01-02T15:04:05Z07:00
   Time format for snapshot file name.

Jiri snapshot list - List existing project snapshots

The "snapshot list" command lists existing snapshots of the labels specified as
command-line arguments. If no arguments are provided, the command lists
snapshots for all known labels.

Usage:
   jiri snapshot list <label ...>

<label ...> is a list of snapshot labels.

Jiri update - Update all jiri tools and projects

Updates all jiri projects, builds the latest version of jiri tools, and installs
the resulting binaries into $JIRI_ROOT/devtools/bin. The sequence in which the
individual updates happen guarantees that we end up with a consistent set of
tools and source code.

The set of project and tools to update is describe by a manifest. Jiri manifests
are revisioned and stored in a "manifest" repository, that is available locally
in $JIRI_ROOT/.manifest. The manifest uses the following XML schema:

 <manifest>
   <imports>
     <import name="default"/>
     ...
   </imports>
   <projects>
     <project name="release.go.jiri"
              path="release/go/src/v.io/jiri"
              protocol="git"
              name="https://vanadium.googlesource.com/release.go.jiri"
              revision="HEAD"/>
     ...
   </projects>
   <tools>
     <tool name="jiri" package="v.io/jiri"/>
     ...
   </tools>
 </manifest>

The <import> element can be used to share settings across multiple manifests.
Import names are interpreted relative to the $JIRI_ROOT/.manifest/v2 directory.
Import cycles are not allowed and if a project or a tool is specified multiple
times, the last specification takes effect. In particular, the elements <project
name="foo" exclude="true"/> and <tool name="bar" exclude="true"/> can be used to
exclude previously included projects and tools.

The tool identifies which manifest to use using the following algorithm. If the
$JIRI_ROOT/.local_manifest file exists, then it is used. Otherwise, the
$JIRI_ROOT/.manifest/v2/<manifest>.xml file is used, which <manifest> is the
value of the -manifest command-line flag, which defaults to "default".

NOTE: Unlike the jiri tool commands, the above manifest file format is not an
API. It is an implementation and can change without notice.

Usage:
   jiri update [flags]

The jiri update flags are:
 -attempts=1
   Number of attempts before failing.
 -gc=false
   Garbage collect obsolete repositories.
 -manifest=
   Name of the project manifest.

Jiri version - Print version

Print version of the jiri tool.

Usage:
   jiri version

Jiri api - Manage vanadium public API

Use this command to ensure that no unintended changes are made to the vanadium
public API.

Usage:
   jiri api [flags] <command>

The jiri api commands are:
   check       Check if any changes have been made to the public API
   fix         Update .api files to reflect changes to the public API

The jiri api flags are:
 -color=true
   Use color to format output.
 -gotools-bin=
   The path to the gotools binary to use. If empty, gotools will be built if
   necessary.
 -manifest=
   Name of the project manifest.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri api check - Check if any changes have been made to the public API

Check if any changes have been made to the public API.

Usage:
   jiri api check [flags] <projects>

<projects> is a list of vanadium projects to check. If none are specified, all
projects that require a public API check upon presubmit are checked.

The jiri api check flags are:
 -detailed=true
   If true, shows each API change in an expanded form. Otherwise, only a summary
   is shown.

Jiri api fix

Update .api files to reflect changes to the public API.

Usage:
   jiri api fix <projects>

<projects> is a list of vanadium projects to update. If none are specified, all
project APIs are updated.

Jiri copyright - Manage vanadium copyright

This command can be used to check if all source code files of Vanadium projects
contain the appropriate copyright header and also if all projects contains the
appropriate licensing files. Optionally, the command can be used to fix the
appropriate copyright headers and licensing files.

In order to ignore checked in third-party assets which have their own copyright
and licensing headers a ".jiriignore" file can be added to a project. The
".jiriignore" file is expected to contain a single regular expression pattern
per line.

Usage:
   jiri copyright [flags] <command>

The jiri copyright commands are:
   check       Check copyright headers and licensing files
   fix         Fix copyright headers and licensing files

The jiri copyright flags are:
 -color=true
   Use color to format output.
 -manifest=
   Name of the project manifest.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri copyright check - Check copyright headers and licensing files

Check copyright headers and licensing files.

Usage:
   jiri copyright check <projects>

<projects> is a list of projects to check.

Jiri copyright fix - Fix copyright headers and licensing files

Fix copyright headers and licensing files.

Usage:
   jiri copyright fix <projects>

<projects> is a list of projects to fix.

Jiri env - Print vanadium environment variables

Print vanadium environment variables.

If no arguments are given, prints all variables in NAME="VALUE" format, each on
a separate line ordered by name.  This format makes it easy to set all vars by
running the following bash command (or similar for other shells):
   eval $(jiri env)

If arguments are given, prints only the value of each named variable, each on a
separate line in the same order as the arguments.

Usage:
   jiri env [flags] [name ...]

[name ...] is an optional list of variable names.

The jiri env flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri go - Execute the go tool using the vanadium environment

Wrapper around the 'go' tool that can be used for compilation of vanadium Go
sources. It takes care of vanadium-specific setup, such as setting up the Go
specific environment variables or making sure that VDL generated files are
regenerated before compilation.

In particular, the tool invokes the following command before invoking any go
tool commands that compile vanadium Go code:

vdl generate -lang=go all

Usage:
   jiri go [flags] <arg ...>

<arg ...> is a list of arguments for the go tool.

The jiri go flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri goext - Vanadium extensions of the go tool

Vanadium extension of the go tool.

Usage:
   jiri goext [flags] <command>

The jiri goext commands are:
   distclean   Restore the vanadium Go workspaces to their pristine state

The jiri goext flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri goext distclean - Restore the vanadium Go workspaces to their pristine
state

Unlike the 'go clean' command, which only removes object files for packages in
the source tree, the 'goext disclean' command removes all object files from
vanadium Go workspaces. This functionality is needed to avoid accidental use of
stale object files that correspond to packages that no longer exist in the
source tree.

Usage:
   jiri goext distclean

Jiri oncall - Manage vanadium oncall schedule

Manage vanadium oncall schedule. If no subcommand is given, it shows the LDAP of
the current oncall.

Usage:
   jiri oncall [flags]
   jiri oncall [flags] <command>

The jiri oncall commands are:
   list        List available oncall schedule

The jiri oncall flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri oncall list - List available oncall schedule

List available oncall schedule.

Usage:
   jiri oncall list

Jiri profile - Manage vanadium profiles

To facilitate development across different host platforms, vanadium defines
platform-independent "profiles" that map different platforms to a set of
libraries and tools that can be used for a facet of vanadium development.

Each profile can be in one of three states: absent, up-to-date, or out-of-date.
The subcommands of the profile command realize the following transitions:

  install:   absent => up-to-date
  update:    out-of-date => up-to-date
  uninstall: up-to-date or out-of-date => absent

In addition, a profile can transition from being up-to-date to out-of-date by
the virtue of a new version of the profile being released.

To enable cross-compilation, a profile can be installed for multiple targets. If
a profile supports multiple targets the above state transitions are applied on a
profile + target basis.

Usage:
   jiri profile [flags] <command>

The jiri profile commands are:
   install     Install the given vanadium profiles
   list        List known vanadium profiles
   setup       Set up the given vanadium profiles
   uninstall   Uninstall the given vanadium profiles
   update      Update the given vanadium profiles

The jiri profile flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri profile install - Install the given vanadium profiles

Install the given vanadium profiles.

Usage:
   jiri profile install <profiles>

<profiles> is a list of profiles to install.

Jiri profile list - List known vanadium profiles

List known vanadium profiles.

Usage:
   jiri profile list

Jiri profile setup - Set up the given vanadium profiles

Set up the given vanadium profiles. This command is identical to 'install' and
is provided for backwards compatibility.

Usage:
   jiri profile setup <profiles>

<profiles> is a list of profiles to set up.

Jiri profile uninstall - Uninstall the given vanadium profiles

Uninstall the given vanadium profiles.

Usage:
   jiri profile uninstall <profiles>

<profiles> is a list of profiles to uninstall.

Jiri profile update - Update the given vanadium profiles

Update the given vanadium profiles.

Usage:
   jiri profile update <profiles>

<profiles> is a list of profiles to update.

Jiri run - Run an executable using the vanadium environment

Run an executable using the vanadium environment.

Usage:
   jiri run [flags] <executable> [arg ...]

<executable> [arg ...] is the executable to run and any arguments to pass
verbatim to the executable.

The jiri run flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri test - Manage vanadium tests

Manage vanadium tests.

Usage:
   jiri test [flags] <command>

The jiri test commands are:
   generate    Generate supporting code for v23 integration tests
   project     Run tests for a vanadium project
   run         Run vanadium tests
   list        List vanadium tests

The jiri test flags are:
 -color=true
   Use color to format output.
 -n=false
   Show what commands will run but do not execute them.
 -v=false
   Print verbose output.

Jiri test generate - Generate supporting code for v23 integration tests

The generate command supports the vanadium integration test framework and unit
tests by generating go files that contain supporting code.  jiri test generate
is intended to be invoked via the 'go generate' mechanism and the resulting
files are to be checked in.

Integration tests are functions of the following form:

    func V23Test<x>(i *v23tests.T)

These functions are typically defined in 'external' *_test packages, to ensure
better isolation.  But they may also be defined directly in the 'internal' *
package.  The following helper functions will be generated:

    func TestV23<x>(t *testing.T) {
      v23tests.RunTest(t, V23Test<x>)
    }

In addition a TestMain function is generated, if it doesn't already exist.  Note
that Go requires that at most one TestMain function is defined across both the
internal and external test packages.

The generated TestMain performs common initialization, and also performs child
process dispatching for tests that use "v.io/veyron/test/modules".

Usage:
   jiri test generate [flags] [packages]

list of go packages

The jiri test generate flags are:
 -prefix=v23
   Specifies the prefix to use for generated files. Up to two files may
   generated, the defaults are v23_test.go and v23_internal_test.go, or
   <prefix>_test.go and <prefix>_internal_test.go.
 -progress=false
   Print verbose progress information.

Jiri test project - Run tests for a vanadium project

Runs tests for a vanadium project that is by the remote URL specified as the
command-line argument. Projects hosted on googlesource.com, can be specified
using the basename of the URL (e.g. "vanadium.go.core" implies
"https://vanadium.googlesource.com/vanadium.go.core").

Usage:
   jiri test project <project>

<project> identifies the project for which to run tests.

Jiri test run - Run vanadium tests

Run vanadium tests.

Usage:
   jiri test run [flags] <name...>

<name...> is a list names identifying the tests to run.

The jiri test run flags are:
 -blessings-root=dev.v.io
   The blessings root.
 -clean-go=true
   Specify whether to remove Go object files and binaries before running the
   tests. Setting this flag to 'false' may lead to faster Go builds, but it may
   also result in some source code changes not being reflected in the tests
   (e.g., if the change was made in a different Go workspace).
 -num-test-workers=<runtime.NumCPU()>
   Set the number of test workers to use; use 1 to serialize all tests.
 -output-dir=
   Directory to output test results into.
 -part=-1
   Specify which part of the test to run.
 -pkgs=
   Comma-separated list of Go package expressions that identify a subset of
   tests to run; only relevant for Go-based tests
 -v23.credentials.admin=
   Directory for vanadium credentials.
 -v23.credentials.publisher=
   Directory for vanadium credentials for publishing new binaries.
 -v23.namespace.root=/ns.dev.v.io:8101
   The namespace root.

Jiri test list - List vanadium tests

List vanadium tests.

Usage:
   jiri test list

Jiri help - Display help for commands or topics

Help with no args displays the usage of the parent command.

Help with args displays the usage of the specified sub-command or help topic.

"help ..." recursively displays help for all commands and topics.

Usage:
   jiri help [flags] [command/topic ...]

[command/topic ...] optionally identifies a specific sub-command or help topic.

The jiri help flags are:
 -style=compact
   The formatting style for help output:
      compact - Good for compact cmdline output.
      full    - Good for cmdline output, shows all global flags.
      godoc   - Good for godoc processing.
   Override the default by setting the CMDLINE_STYLE environment variable.
 -width=<terminal width>
   Format output to this target width in runes, or unlimited if width < 0.
   Defaults to the terminal width if available.  Override the default by setting
   the CMDLINE_WIDTH environment variable.
*/
package main