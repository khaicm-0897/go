# Learning  Go Lang

## How to Write Hello, World! in Go

In there, create a new folder, for example call it hello.

In there, create a hello.go file (you can name it as you want).

Add this content:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Let’s analyze this line by line.

```go
package main
```

We organize Go programs in packages.

Each .go file first declares which package it is part of.

A package can be composed by multiple files, or just one file.

A program can contain multiple packages.

The main package is the entry point of the program and identifies an executable program.

```go
import "fmt"
```

We use the import keyword to import a package.

fmt is a built-in package provided by Go that provides input/output utility functions.

We have a [large standard library](https://pkg.go.dev/std) ready to use that we can use for anything from network connectivity to math, crypto, image processing, filesystem access, and more.

```go
func main() {
	
}
```

Here we declare the main() function.

What’s a function? We’ll see more about them later, but in the meantime let’s say a function is a block of code that’s assigned a name, and contains some instructions.

The main function is special because what’s where the program starts.

In this simple case we just have one function – the program starts with that and then ends.


## Chapter 1: Programming A Computer

### 1. Technical concepts covered

- Memory Unit, Arithmetic and Logic Unit, Input/Output, Control Unit

- Central memory, Auxiliary memory

- Volatile and non-volatile memory

- RAM/ROM

- CPU

- High and low-level languages

- Assembly language, assembler

- Compiled and interpreted language

### 2. The four hardware components

A computer is composed of four main parts :

The memory unit (MU) where data and programs are stored.

For instance, we can store into the memory unit the grades of a college class. We can also hold a program that will compute the class’s average grade.
The arithmetic and logic unit (ALU). Its role is to perform arithmetic and logical operations on data stored into the memory unit. This unit can perform, for instance, additions, incrementations, decrementations, which are called operations. In general, each operation requires two operands and output a result.

Let’s say we want to add 5 and 4. Those two numbers are the operands. The result of this operation is 9. Operands are loaded from the memory unit. The ALU is an electrical circuitry that is designed to execute operations.
The input and output unit (I/OU) will be in charge of loading data into the memory unit from an input device. This unit also sends data from the memory unit to an output device.

An input device is, for example, the touchpad of your computer, your keyboard, your mouse.

An output device is, for instance, your monitor.

The control unit (CU) will receive instructions from programs and will control the activity of the other units.

The four hardware components represent a schematic view of the computer’s components.

### 3. Memory

A computer is composed of two types of memory :

- The central memory

- The auxiliary memory

Two categories of memory exist:

- Volatile

- Non Volatile

#### 3.1. The central memory

This type of memory is necessary to run the operating systems and all the other programs your computer will run. The central memory contains two types of storage:

- **RAM** (Random Access Memory). This type of storage requires electric power to persist data. When you turn your computer off, the memory contained in this type of storage will be erased. The operating system and the programs you use will be loaded into this memory. This type of memory is volatile.

- **ROM** (Read-Only Memory). This is a memory that contains data necessary for the computer to run correctly. This kind of memory is not volatile (when you turn the computer off, it will not be erased). It’s designed to be only readable and not updated by the system.

#### 3.2. The auxiliary memory

This type of memory is not volatile. When the power is going off, the data stored will not be erased. Here are some examples of auxiliary memory: Hard drives, USB keys, CD-ROM, DVD ...etc.

Read and writes to this type of memory is **slow** compared to the **RAM**.

Some hard drives sequentially access memory. The system should respect a particular sequence. Respecting this access sequence takes a longer time than a random access mode. Note that some hard drives allow random memory access.

##### 3.2.1 SSD hard drive

Hard drives, also denoted Hard Disk Drive (HDD), are composed of magnetic disks that are rotating. Data are read and write thanks to a moving magnetic head. Reads and writes will generate a rotation and a magnetic head movement, which consumes time.

SSD (Solid-State Drives) are not constructed like that. There is no magnetic head neither magnetic disks. Instead, data is stored in flash memory cells. Data access is quicker on that kind of disk. Note that SSD also costs more than traditional electromagnetic hard drives.

### 4. CPU

CPU is the initials of Central Processing Unit. The CPU is also denoted processor. The CPU contains :

- The arithmetic and logic unit (ALU)

- The control unit (CU)

The CPU will be responsible for executing the instructions given by a program. For instance, the program can instruct to perform an addition between two numbers. Those numbers will be retrieved from the memory unit and passed to the ALU. The program might also require performing an I/O operation like reading data from the hard drive and loading it into the RAM for further processing. The CPU will execute those instructions.

The CPU is the central component of a computer.

### 5. What is a program

To make computers do something, we have to feed them with precise instructions. This set of instructions is called “program”.

Following a more official definition, a program is “a combination of computer instructions and data definitions that enable computer hardware to compute”

Let’s take an example. Imagine a program that asks the user to type two numbers. The program adds those numbers, and the result is then displayed on the monitor. The instructions that have to be written are :

1. Output “Please type your first number and press enter” on the monitor.

2. When a number is typed and the “Enter” key is pressed on the keyboard, store the number into memory. Let’s denote A this number.

3. Output “Please type your second number and press enter” on the monitor.

4. When a number is typed and the “Enter” key is pressed on the keyboard, store the number into memory. Let’s denote B this number.

5. Send to the ALU the two numbers (A and B) and the addition opcode and store the result into memory.

6. Output the result on the monitor

Two types of instructions are performed :

**I/O operations**: store numbers into memory from an the input device (the keyboard), load data from memory, retrieve numbers stored into memory, and display it to the user.

**An arithmetic operation**: add two numbers.

### 6. How to speak to the machine

#### 6.1. Programming languages are formal languages

Instructions that are given to the machine are written with programming languages. Programming languages are formal languages.They are composed of words that are constructed from an alphabet (a set of distinct characters). Those words are organized following specific rules. Go is a programming language, like x86 Assembly, Java, C, C++, Javascript...

They are two types of programming languages :

- Low level

- High level

Some high-level languages are compiled, others are interpreted, and some are in between. We will see in the next sections what those two terms mean.

![image](./images/program-language-classification.png)

