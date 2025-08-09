# go-class
Repo for [Go class](https://www.youtube.com/watch?v=iDQAZEJK8lI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=1
) from [Matt Holiday](https://github.com/matt4biz)


## class 00
- simple to deploy, put a go program by itself in a container no need JVM or libc
- container is small and secure
- Easy, fast, safe

## class 01
- Every program has to have a main function, it tells go where does the program start
- Modular language you can put the program in different files and compile it together
- Put different parts of the program into packages. But the main function has to be in a package main
- We need to import any package we use
- To compiles (it sticks it in some secret temp directory) and run (and then get rid of what left over) 

    `go run hello-world.go`

## class 02
- In some languages the function main takes parameters that represent the command line arguments but Go doesn't do that, need to import the `os` package and run like `go run <file.go> <arg1> <arg2>`
- Create a package to execute a specific function, "say"

### [Package](https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules)
- named collection of 1 or more related `.go` files 
- to isolate and reuse code
- every `.go` file should begin with `package <name>` statement which indicates the name of the package that file is part of.
- code in a package can access and use all types, constants, variables and functions whitin that package (even if they are declared in a != `.go` files)

### The main package
- `main` is a special package name which indicates that the package contains the code for an *executable application* (indicates that the package contains code that can be built into a binary and run)
- Any package with name `main` must also contain a `main()` function somewhere in the package which acts as the entry point for the program
- convention: `main()` function lives in a file with name `main.go` (entry point easier to find)
- if you try to run a non-main package it will result in an error

### Importing and using standard library packages

- When importing a package from the standard library you need to use the *full path to the package* in the standard library tree, not jus the name of the package (`math/rand`)
- if you import a package but don't use it in your code it will give you a compile-time error (same goes for referenced packages without importing)

### exported vs unexported

- Something in Go is exported if *its name starts with a capital letter*. For example: `func FooBar() {...}`
- unexported things are 'private' to the package that they are declared in
- exported things in a package are 'public' and visible to any code that imports the package. When you **import a package you get to use its exported things**
- a `main` package should never normally be imported
- package names are short and clear, lower case, with no under_score or mixedCaps, often simple nouns (more on [package names](https://go.dev/blog/package-names))

### Modules

- what if you need to: 
    - import and use 3rd party packages (not from the standard library)
    - structure your code so it's split into **multiple packages**

A [module](https://go.dev/blog/using-go-modules) is a tree of Go source files with a `go.mod` file in the tree's root directory.

- A module is a group of packages (or a single package) that are(is) versioned. 

-  A module is a collection of Go packages stored in a directory with a `go.mod` file at its root. 
-  The `go.mod` file contains the module name and the versions of the dependencies the module uses.

The `go.mod` file has the following structure

```
module gitlab.com/userA/myModule

go 1.15

require (
    github.com/PuerkitoBio/goquery v1.6.1
    github.com/aws/aws-sdk-go v1.36.31
)
```

The first line gives the `module path` (it can be either a local path or a URI to a repository). It should be any string unique and unlikely to be used by anyone else. 

The second line will give the `version of Go` used by the developer

Then another section define the `dependencies` that are used by the module:

```
require (
    DEPENDENCY_1_PATH VERSION_OF_DEPENDENCY_1
    DEPENDENCY_2_PATH VERSION_OF_DEPENDENCY_2
   //...
)
```

First, the module path (Github URL of repo) and then the desired version.

The `go.mod` file can be created automatically with the go command line `go mod init my/module/path`
This will generate the following file :
```
module my/module/path

go 1.15
```

- Run commands from the directory containing go.mod
- When you place .go files directly in the root directory (where go.mod is), their package names (package hello inside hello.go) become directly accessible under the module name. Files in subdirectories (e.g., `hello/hello.go`) add a subpath to imports: `import "class02/hello"`

- Root files are treated as part of the module’s "core" packages

- in the class02 module, we'll create a new package test 

#### Conventions 
- Generally, it is not recommended to use multiple Go modules in one repository
- Base your module path on a URL that you own or control. For example, a good module path would be github.com/user/project.
- The module path doesn't need to be a 'real' functioning URL with something hosted at it. It's really just an arbitrary string which acts as a unique identifier for your module unless you make an open source package then your module path must be the location that the code will be fetchable from (github.com/example/package).

[Continue with class 02](https://youtu.be/-EYNVEv-snE?feature=shared&t=298)