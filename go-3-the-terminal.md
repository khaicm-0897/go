# The terminal

## 1. Introduction

The day-to-day life of a Go programmer requires the usage of the terminal. This chapter will explain what the terminal is and how to use it. If you are already familiar with it, you can skip the chapter.

## 2. Graphical User Interface (GUI)

Operating systems like macOS, Windows, or Linux offer a rich graphical user interfaces (GUI). To launch a program installed on your computer, you generally double-click on an icon placed on your desktop. Programs will be displayed in windows with an interactive interface: menus, sidebars, buttons...

We say that those programs offer a graphical user interface (GUI). The vast majority of users will use programs that provide these interfaces. Graphical interfaces are easy to use and intuitive.

## 3. Command Line Interface (CLI)

Graphical user interfaces have not always existed. The first computers did not have such capability. But how do the users of those computers managed to launch and use programs? Computers were shipped with a command-line interface. This interface is also called a “shell”. A shell is a program that can pass commands to an operating system1. Shell is a generic term that designates such programs. The most famous shell is bash (Bourne Again Shell). Bash is shipped by default with macOS and a vast majority of Linux Distributions. Windows is also shipped by default with a shell (but it’s not bash).

## 4. How to interact with a shell: the terminal

The shell was presented to the user directly after startup on old computers. On modern computers, we have to launch a program to interact with a shell. This program is often called a terminal.

### 4.1.Window Cmder

The terminal and the Windows shell is not very practical on a day today basis. I advise you to install cmder To make your developer life easier on windows. Cmder is an emulator that will allow you to use commands available on Linux/MacOS. The installation process is easy (download the latest release on GitHub) then launch the installation wizard.

After installing cmder, launch the program “Cmder” to open your brand new terminal.

### 4.2 Bash for Windows

By default, you cannot use bash on a Windows computer. This is not a problem, but it means that you will have to find the Windows equivalent for each macOS/Linux command. It can be cumbersome at some point because a lot of examples and tutorials on the web do not always provide equivalence for Windows.

Microsoft announced that you can now install a “Windows Subsystem for Linux” (WSL) on your Windows computer. This is a piece of good news because you will use bash. You can find the installation instructions on the Microsoft website: https://docs.microsoft.com/en-us/windows/wsl/install-win10 (for Windows 10).

I strongly advise you to install this, because it will make your life easier even if I try to give Windows equivalent for basic commands in the next sections.

## 5. How to use the terminal

After you opened your terminal, you will see a black window. This is the interface where you can type commands. This is not intuitive because to type commands; you have to know them before! Each command has a name. To launch a command, you will type its name, eventually type some options and then press enter. Let’s take an example.

## 5.1. About the dollar symbol

In the examples, you will see a line that begins with a dollar symbol. This is a convention; it means “type it into your terminal”, when you want to reproduce the examples do not type the dollar, just everything after it.

## 6. Test yourself

1. What is the command to list the contents of a directory?
- ls (for UNIX sytems)
- dir for Windows
2. What is a terminal?
- A terminal is a program that offers an interface to use a shell
3. Give the name of a well-known shell.

- bash
 ## 7. Key takeaways

- Graphical User Interfaces have not always existed

- To interact with computers, we can use a Graphical Interface or a Command line Interface. (CLI)

- To use a CLI, we have to open a terminal application that offers an interface to interact with a shell

- Bash is a shell.

- You can launch commands by typing the command name and eventually options, then press enter.

- We will use the terminal to launch go-specific commands.
