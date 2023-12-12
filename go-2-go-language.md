# The Go Language

## 1. The myth of creation 
There is a myth around the creation of Go. The language was born inside a Google office, and it happened during a very long build that took 45 minutes.

This story is told by Rob Pike in. It gives us a piece of precious information about the motivations behind the creation of Go. Build times were too long and painful... they had to find a way to avoid them; that was Go Genesis’ entry point.

Robert Griesemer, Ken Thompson, and Rob Pike are the developers that started working on Go back in 2007. Rob Pike claims that by mid-2008, the language was “mostly designed and the implementation (compiler, run-time) starting to work.”. After that, Ian Lance Taylor and Russ Cox joined the team in 2008.

Go is an open-source programming language maintained by its community and a core team of developers working at Google. The 16th March of 2011 is the date of the first Go release. (It was named “r56”). Go version 1 was released on the 28th March of 2012.

## 2. Motivations

Go (or Golang) was built by Google to solve the company’s problems. To better understands the reasons, it’s worth reading the keynote talk of Rob Pike.

What are the challenges of software at large worldwide companies?

- The codebase of Google services is massive. Google has millions of lines of code.

- Those lines are written in different languages: C, C++, Java, and others.

- The build time of those applications “have stretched to many minutes, even hours”.

- Updates of some application parts can be costly.

The objective of the first Gophers1 was to make the life of developers easier by :

- Reducing the build time of programs drastically.

- Designing a language that is easy to learn, read and debug for young developers that have been exposed to C, C++, or Java.

- Designing an effective dependency management system.

- Building a language that can produce software that scale well, on hardware.

### 2.1. Definition of some concepts

