# First Go Application

## 1. Application specification
Before writing any code, we need to decide what our application will do. Most projects begin with this phase. This is called the specification phase. It aims to produce precise requirements that the application should fulfill. Those requirements are the specifications (or the specs).

Our application specification is simple: when launched, the application should display the date and time, and exit.

## 2. Project directory

A Go application is composed of files. On those files, we will write Go code. We call those file “source files”. An application is stored in the main directory. This main directory can contain just one source file. Most of the time, it’s composed of several sub-directories.

We will create this main directory for our application. You can do it using the command line (or with the graphical interface of your system) :

```go
$ cd Documents/code
$ mkdir dateAndTime
```

## 3. IDE

We are just minutes away from writing code. Do you need special software to write Go code? In theory, you can write code with a standard text editor. There is dedicated software on the market specially developed for developers. They are called IDE: Integrated Development Environment.

-IDE provides functionalities like :

- Automatic coloration for reserved words (syntax highlighting)

- Autocompletion

- Refactoring capabilities

They are many IDE on the market. You can search google to find the best fit for you. I use Goland, which IntelliJ develops. This software is not free (it’s subscription-based), but I find it easy to use.

## 4. Source file

Let’s create our source file. We will name it main.go.

```go
// first-go-application/first/main.go
package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)
}
```

### 4.1. Code explanation

The first line is mandatory in Go. In all files, you must add the package declaration. Such declaration is composed of the keyword package then the name of the package.

On the second line, you can read another keyword: import. It is usually followed by an open parenthesis and a list of the program’s imported packages. Every package is written on a new line. Each package has a name delimited by double-quotes. Here our application depends on two packages :

- fmt
- time

Those packages are part of the standard library.

Then you find the declaration of a function named main. We will go deeper into the syntax of functions later.

- The function declaration is enclosed with curly brackets : { and }.

-  Inside the function declaration, we have two statements :

    - The first instruction is an affectation. We initialize the variable “now” and we give it the value returned by the function call Now()from the package time.

    - The second instruction is a call of the function Print from the package fmt

WARNING! be sure that you only use packages that you actually import. Otherwise your program will not compile... When you use a package not imported, your Go program will not compile

Make sure that you use imported packages. In the following code sample I import the package fmt and time but I do not use them in my **main** function :

```go
// DO NOT COMPILE
// first-go-application/import-issue/main.go
package main

import (
    "fmt"
    "time"
)

func main(){

}
```
### 4.2. About the main function

The main function is the **entry point** of the program. In every application, you have at least one main function. The program will start with the first statement of this function. (Note that in C, C++, Java, the main function’s concept exists).


## 5. Compilation

The source file is ready to be transformed into a binary (or executable). To do so, we will use the Go toolchain. Open a terminal :

```go
$ cd Documents/code/dateAndTime
$ go build main.go
```

The first instruction (cd) instructs the shell to change the current directory to Documents/code/dateAndTime. This command will compile the program into an executable. The executable is named main (the same name as the source file, without the .go extension). Let’s see the files that are now into our dateAndTime directory :

```bash
$ ls -lh
total 4160
-rwxr-xr-x  1 maximilienandile  staff   2.0M Aug 16 11:27 main
-rw-r--r--  1 maximilienandile  staff    94B Aug 16 11:00 main.go
```

We use the command ls. (for windows user, you can use the command dir). You can see that we have two files :

main that weights 2.0M (MegaBytes), and that is executable.

main.go that weights only 94 Bytes. (this is our source file).

Now it’s time to launch our application :

```bash
$ ./main
2019-08-16 11:45:44.435637 +0200 CEST m=+0.000263533
```
## 6. Test yourself

1. How to compile a Go application?

- Open a terminal

- Go to the directory of your application. Providing that there is a file named main.go that contains a main function issue those commands :

    - $ cd /code/myApp

    - $ go build main.go

2. How is the result of the compilation called?

- It’s called an executable or binary
3. What is the name of the entry point function of a Go application?

- main
4. What is the usage of the import statement?

- It’s used to import packages (from the standard library or other sources) into the program.

- An imported package can then be used inside the code.

## 7. Key takeaways

- To create a simple program

    - Create a file

    - Name it main.go

    - Here is the basic skeleton of a program (that does nothing)

```go
// first-go-application/skeleton/main.go
package main

func main() {
    
}
```

- This file is denoted “source file”

- From this source file, we can create an executable program (that can be launched).

- The creation of the executable is called “compilation”.

- To compile a program, type the following command into the terminal

    - go build main.go
- To launch the compiled program, type the following into your terminal :

    - ./main
