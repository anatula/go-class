# go-class
Repo for Go series from Matt Holiday
https://www.youtube.com/watch?v=iDQAZEJK8lI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=1

class 00
- simple to deploy, put a go program by itself in a container no need JVM or libc
- container is small and secure
- Easy, fast, safe

class 01
Every program has to have a main function, it tells go where does the program start
Modular language you can put the program in different files and compile it together
Put different parts of the program into packages. But the main function has to be in a package main
We need to import any package we use

Compiles (it sticks it in some secret temp directory) and run (and then get rid of what left over)
`go run hello-world.go`