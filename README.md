Compiling with Go
=================

This project shows some tips and trick about compilation with Go. It's going to apply the tips and techniques step by step with a very [simple example code (m ain.go)](./main.go). Is tested with **go version go1.6 linux/amd64** over Ubuntu 15.10 (4.2.0-30-generic #36-Ubuntu SMP).

The example code is:

    package main
    
    import "fmt"
    
    var (
        version, build, buildDate string
    )
    
    func main() {
        fmt.Printf("version:   [%s]\n", version)
        fmt.Printf("build:     [%s]\n", build)
        fmt.Printf("buildDate: [%s]\n", buildDate)
    }


Basic compilation
-----------------

The basic compilation is with **go build** command. You cax use **-o** argument to set the output filename.

    $ go build -o example main.go

    $ ./example



Passing variables
-----------------

Go have the capacity to set variables *inside* binary in compilation time. This is very usefull, for example, to set an application version, the SHA1 of the Git commit, etc.

To do it, simply set  **-ldflags** in build process, with a "-X <variable_name=param_value>" for each param to export to binary. 

Example: 

    // In the go source file 
    package main
    
    var myVariable string;


    $ go build -ldflags "-X main.myVariable=MyValue" main.go


Use cases could be, for example, specify business data to binary that will run in an runtime context with an specific info (for exaple, to fix to a machine or build especific binaries for each client setting values in binary).


### Interesting uses

A tipical usage is add release information to binary file. For example, version, build identifier or compilation date. In our example:


    $ go build -o example -ldflags "-X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" main.go
    
    $ ./example

The binary file size, in our example, is * **2287896** * bytes


Binary file size optimiztions
-----------------------------

There are some compiling params than you can apply to optimize final binary.


**Disable DWARF generated debugging information**

    -ldflags "-w"


    $ go build -o example -ldflags "-w -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" main.go


After this optimization the file size is * **1722974** * bytes.


**Omit symbol table and debugging information**

    -ldflags "-s"

    $ go build -o example -ldflags "-s -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" main.go

After this optimization the file size is * **1589888** * bytes.


You can merge both params, but not reduces size

    -ldflags "-s -w"

    $ go build -o example -ldflags "-s -w -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" main.go

After this optimization the file size is the same, * **1589888** * bytes.



Also, to reduce the file size  you can use [The Ultimate Packer for eXecutables (UPX)](http://upx.sourceforge.net/) ([wikipedia](https://en.wikipedia.org/wiki/UPX)). But running UPX repacked program may require a bit more memory than the original. Also, it adds to the startup time, it might be noticeable. So... up to you.

To repackage with UPX, after Go compilation:

    $ upx <binary_file>


Generate full independent binary
--------------------------------

A go binary will need some runtime dependencies (shared libraries) in target executable machine. To create a full independent binary (specific for each platform) you can use a few environment variables and compilation flags.

    $ CGO_ENABLED=0 GOOS=<target_platform> go build [arguments] -a -installsuffix cgo <source_files>


To build the current package/project for Linux, for example:

    $ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo


In our example, applying all previous optimizations:

    # for Linux
    
    $ CGO_ENABLED=0 GOOS=linux go build -o example -ldflags "-s -w -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" -a -installsuffix cgo main.go


    # for Windows
    
    $ CGO_ENABLED=0 GOOS=windows go build -o example.exe -ldflags "-s -w -X 'main.version=$(cat VERSION)' -X 'main.build=$(git rev-parse --short HEAD)' -X 'main.buildDate=$(date --rfc-3339=seconds)'" -a -installsuffix cgo main.go

