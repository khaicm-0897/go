# Setup your dev environment

## 1. Computer architecture

In the following sections, you will need to know your computer architecture. This technical term sometimes brings confusion to newcomers. What is it exactly?

Computers come in different shapes, prices and forms. A lot of consumers focus on the brand on the computer and their high level specifications (like RAM, amount of storage on the hard drive). The question of the processor is sometimes eluded by marketing, but it’s a fundamental part of the computer. The processor is also called the CPU (Central Processing Unit).

The CPU is responsible for running system instructions but also programs. CPU is a generic term. There are different types of CPUs available.

Programs use a set of predefined instructions to speak to the machine. We call this list of orders the instruction set. Some processors share the same instruction sets. Others have entirely different ones. The architecture notion covers the instruction set and the physical organization, and computer implementation.

You understand now why this notion is essential at the machine level. The Go toolchain is a collection of programs that we will use to build our programs.

- Those programs are written in Go and compiled.

- The Go team provide different versions of the toolchain for each OS and architecture supported

    - Go has a minimum system and OS requirements.

    - You can check them on this webpage:

        - https://github.com/golang/go/wiki/MinimumRequirements.

        - The page is updated frequently.

## 2. Go toolchain installation

To start developing a Go program, you have first to install Go on your computer. To install it, you will have to download a file. The installation procedure is slightly different depending on your OS (Operating System). Just follow the steps corresponding to your OS.

### 2.1. Linux

#### 2.2.1. Determine your architecture

You have to determine the architecture of your computer first. It will allow you to select the right file to download.

Open a terminal and type the following command:

```bash
$ uname -p
```

Uname is a program that displays characteristics of your system. The flag -p will display “print the machine processor architecture name”.

Golang Linux versions are available for the following architectures :

- x86

- x86-64

- ARMv6

- ARMv8

- ppc64le

- s390x

#### 2.2.2. Check the archive hash

This step is strongly advised. The objective is to check that you do not have a corrupted version of the archive. To do that, open a terminal and type :

```bash
$ cd /where/you/have/downloaded/the/file
$ sha256sum go1.12.8.linux-amd64.tar.gz
```

The first command use cd (change directory) to go to the directory where the file you have downloaded has been put. For instance, if the file is in /home/max/ you have to execute :

```bash
$ cd /home/max
```

The next command will compute the SHA256 hash (which is a cryptographic hash function) of the file go1.12.8.linux-amd64.tar.gz.

**Cryptographic hash function**: A function that will take as input variable size data (a file, a word, a sentence) and output a fixed size data blob. Having the output, it’s almost impossible to retrieve the function’s input. The result is called the “hash” or the “message-digest”. Here we use it to ensure that the file has not been altered in its way to our computer. To do that, we compare the hash that we computed on our computer with the hash provided on the Go website. If both are equal, there has been no alteration.

When the **sha256sum** has been executed, you will see on your screen a set of characters that you have to compare with the hash displayed on the website. If there are not identical, something went wrong, do not use the file downloaded. If the hash you have computed and the hash displayed on the Go website is equal, your good to go.


#### 2.2.3. Decompress the archive

The archive is gzipped (compressed with gzip, it’s often called a tarball). To deflate it, you have two options :

1. Use the graphical interface (if you have one): double click on the archive and a window will open. Follow the instructions. You have to extract and deflate the files in /usr/local

2. Use the terminal.

We will take option 2.

Open your terminal and type the command :

```bash
$ sudo tar -C /usr/local -xzf go1.12.8.linux-amd64.tar.gz
```
We will use the application tar (in sudo mode)

- **-C /usr/local** : means that we will change the exection directory to /usr/local

- **-xzf go1.12.8.linux-amd64.tar.gz**: means that we want to extract (x) the archive go1.12.8.linux-amd64.tar.gz that is compressed with gzip.

#### 2.1.4. Set your PATH

We have eight folders: API, bin, doc, lib, misc, pkg, src, and tests.

The folder bin contains two executables :

- **go** : this is the main executable

- **godoc** : this is a program used to generate documentation1

- **gofmt** : this program will format your source files according to the language conventions

If you are into the bin directory, you can launch go. Try this command :

```go
$ go version
```

### 3. A tour of Go environment variables

Go use environment variables for its configuration. In this section, we will detail some of them :

- **GOBIN**: By default, Go will place compiled programs into $GOPATH/bin. If you want to override this behavior, you can set this variable.

- **GOROOT**: The absolute path where the Go distribution is installed (for Linux and macOS user, it’s by default /usr/local/go).

- **GOHOSTOS**: This is the operating system of the Go toolchain

- **GOHOSTARCH**: This is the system architecture of the Go toolchain binaries

To print all the Go environment variables, you can use this command :

```go
$ go env
```
Use this command when you have trouble with your Go settings. Here is the output of the command :

```go
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/maximilienandile/Library/Caches/go-build"
GOENV="/Users/maximilienandile/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/maximilienandile/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/maximilienandile/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.16"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/lm/9l2tk4811x32rmw4407f9h3m0000gn/T/go-build744222834=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

### 4. Key takeaways

- The Go toolchain is available on the Go official website

- The Go toolchain is available on Windows, macOS, and Linux Operating systems

- You should download the version that matches your OS and your Architecture.

- After installing the Go toolchain, you can launch the go command with your favorite terminal.

    - You might need to add the go binary to your PATH environment variable
