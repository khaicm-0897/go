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



