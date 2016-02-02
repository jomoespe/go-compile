Compile passing variales
========================

Go have the capacity to set variables *inside* binary in compilation time. This is very usefull, for example, to set an application version, the SHA1 of the Git commit, etc.

To do it, simply set  **-ldflags** in build process, with a "-X <param_name=param_value>" for each param to export to binary.

    $ go build -ldflags "-X main.myVariable=value"

In the Go code this value should setted to a *public* variable.

    package main

    var myVariable string;


Use cases could be, for example, specify business data to binary that will run in an runtime context with an specific info (for exaple, to fix to a machine or build especific binaries for each client setting values in binary).


# The example

In this example we are going to ser a *build id*, a *version* (this two are artifact information) and a *Id* (this is business data) to the project main class. TO run the example:

    $ go run -ldflags "-X main.Build build-1 -X main.Version 1.0.0-SNAPSHOT -X main.Id M7160FC" example.go


# Another interesting usages

Another interesting use is to burn a *build id* in the binaries. You can get, for example, the SHA-1 of the Git commit and set to program.

Get the SHA-1 of the commit...

    $ git rev-parse HEAD
    $ git rev-parse --short HEAD


And then pass to program:

    $ go build -ldflags "-X main.Build [the_build_id]" example.go

Or *inlining* the command execution:

    $ go run -ldflags "-X main.Build `git rev-parse HEAD`" example.go

