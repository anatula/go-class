## class 02
- In some languages the function main takes parameters that represent the command line arguments but Go doesn't do that, need to import the `os` package and run like `go run <file.go> <arg1> <arg2>`
- Create a package to execute a specific function, "Say" (note uppercase since its exported)

### [Package](https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules)
- named collection of 1 or more related `.go` files 
- to isolate and reuse code
- every `.go` file should begin with `package <name>` statement which indicates the name of the package that file is part of.
- **code in a package can access and use all types, constants, variables and functions whitin that package (even if they are declared in a != `.go` files)**
- if you import a package but don't use it in your code it will give you a compile-time error (same goes for referenced packages without importing)

### The main package
- `main` is a special package name which indicates that the package contains the code for an *executable application* (indicates that the package contains code that can be built into a binary and run)
- Any package with name `main` must also contain a `main()` function somewhere in the package which acts as the entry point for the program
- 👍 Convention: `main()` function lives in a file with name `main.go` (entry point easier to find)
- if you try to run a non-main package it will result in an error

### Importing and using standard library packages

- When importing a package from the standard library you need to use the *full path to the package* in the standard library tree, not just the name of the package (for example `math/rand`)

#### Standard Library Layout
```
math/                  # Directory for the main `math` package
├── abs.go             # File 1: `package math`
├── acosh.go           # File 2: `package math`
├── rand/              # Subdirectory for the `rand` subpackage
│   └── rand.go        # File: `package rand`
└── (20+ other files)  # All with `package math`
```

- Single Package, Multiple Files
- All .go files in the root math/ directory declare package math. Together, they form one logical package (the compiler merges them).No math.go is needed—the package can be split across many files (e.g., sin.go, log.go).

- Subpackages Require Subdirectories: The rand/ subdirectory is a separate package (package rand). This is why it needs its own directory and rand.go file.

The import paths become:
- "math" (for the math package)
- "math/rand" (for the rand package)
- The /rand suffix in the import path comes directly from the directory name, not from any special relationship between the packages.

However, `rand` is not a "child" of `math`:
- rand cannot access math’s unexported (lowercase) functions.
- math cannot access rand’s functions unless it explicitly imports "math/rand".

They are independent packages that happen to share a directory prefix.

So, why together? group related packages and to avoid name collision, for example math/rand crypto/rand.

- Not like Java/Python where nested packages imply inheritance/scoping.

- Not like C++ namespaces where math::rand would be a sub-namespace.

- ***Go is flat: Packages are equal citizens, regardless of directory depth.***

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
```
class02/
├── go.mod
├── hello.go      # package hello
├── hello2.go     # package hello (same package!)
├── hello3/       # DIFFERENT package!
│   └── hello3.go # package hello3 (must differ)
└── cmd/
    └── main.go   # imports both
```
- 1 folder = 1 package (all files must declare same package name and they share everything without exporting)
- Subfolders = separate packages (need to export)

To import:
```
import (
    "class02"        // root files (package hello)
    "class02/hello3" // subfolder (package hello3)
)
```

#### Formatter 

The way (import "class02" + hello.X()) works because:
- Root hello.go declares package hello
- Go automatically links the module (class02) to its root package (hello)

Formatter adds `hello "class02"` in the import because:
- Some tools prefer explicit aliases (safety for future refactors)

- Automatic Aliasing: Some formatters (like gofmt or IDE formatters) add the hello alias when it detects:
    - Multiple packages with the same name (e.g., hello from different paths)
    - Potential naming conflicts
- Explicit Clarity:
- It makes absolutely clear that hello.Say() comes from "class02" (not another package)

IDE/formatter likely has a setting for:
"go.importAlias": "explicit" // vs "auto"

- we'll create a test file `<name>_test.go` (comvention) a test for the whole package
- run `go run` inside the class02 directory

### :=

- := declares and assigns (infers type).
- = only assigns (must be declared first).

### Testing standard library
https://pkg.go.dev/testing

- provides support for automated testing of Go packages
- to be used in concert with the `go test` command, which automates execution of any function of the form
- give that file a name ending in `_test.go`
```
func TestXxx(*testing.T)
where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.
```
- `go test` creates a testing.T instance and calls TestXxx(&t)
- Pointer (*testing.T) → Allows modifying test state (failures, logs).
- Automatically called by go test → You never invoke TestXxx manually.
- testing.T is managed by the test runner → Ensures proper test tracking.
- If t were passed as a value (testing.T), changes would only apply to a copy, and the test runner wouldn’t know if the test failed
- You don’t manage t, just use it


Rules:
- := only inside functions.
- Can’t reuse := in same scope (no redeclaration).
- Works for multiple variables.
- Can reuse := if at least one variable is new.
- Scoped redeclaration is allowed (e.g., inside a new function).
- Works in if/for/switch blocks (scoped locally). 

More on this [StackOverlow anwser](https://stackoverflow.com/a/45654233/5425908)

#### String and string slice
- Quick way to create a slice with pre-filled values
- `[]string` (dynamic slice, not a fixed-size array)
- `[]string{"test"}` in Go is a string slice initialized with one element: "test"
Slice vs. Array:
- `[]string{"test"}` → Slice (dynamic size, backed by an array).
- `[1]string{"test"}` → Array (fixed size of 1).

#### Conventions 
- Generally, it is not recommended to use multiple Go modules in one repository
- Base your module path on a URL that you own or control. For example, a good module path would be github.com/user/project.
- The module path doesn't need to be a 'real' functioning URL with something hosted at it. It's really just an arbitrary string which acts as a unique identifier for your module unless you make an open source package then your module path must be the location that the code will be fetchable from (github.com/example/package).
- use camelCase for functions (snake_case is rare) and keep names short
