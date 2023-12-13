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

- **Build time** : the amount of time necessary for a compiler to generate a machine-readable executable.

- **Statically Typed Language** : Giving a precise definition of this concept is now premature. We will cover this term in the next chapters.

- **Dependency** : a piece of software that is used by another software.

- **Scalability** : the ability for a program to handle an increasing amount of tasks to be performed. For instance, a website is said to be scalable if it can accept an increasing number of requests without downtimes or increasing loading latency.

## 3. Go key features

Go creators have concentrated their efforts on several critical design choices.:

- A compiled language
- With a semantic easy to understand and learn
- Statically typed
- With a built-in **concurrency** a system that is easy for developers to work on
- With a robust dependency management
- With a garbage collector

The main objective, as stated by Rob Pike, was to give developers an easy to learn, a language for “engineering large software projects”

### 3.1. Some Concepts

- **Concurrency** : A program is concurrent when tasks can be executed out-of-order or in a partial order

- **Garbage collector** (often called GC): When we build programs, we need to store data and fetch data from memory. Memory is not an infinite resource. Therefore the developer must ensure that unused elements stored in memory are destroyed from time to time. Putting some data into memory is called allocation; the inverse action, which consists of removing data from memory, is called deallocation. The garbage collector’s role is to deallocate memory when it is not used anymore. When the language does not have any garbage collector, the developer has to collect his garbage and free memory that is no longer used... By chance, Go has a garbage collector.

## 4. The State of Go

- The project has grown very fast, and count now more than one thousand contributors
- At the time of writing (8th January 2020), the latest version of Go is 1.15.6
- A lot of meetups and conferences are organized to federate the community4
    - In 2018 there was 19 conferences organized : 3 in the US and 16 in other countries5
    - In 2017 there were 13 conferences organized

- Go is wanted by developers :
    - in 2018, 2019, and 2020 in the Stackoverflow Developer Survey Go is the top Three most wanted programming language

## 5. Test yourself

1. What is the date of birth of the Go language?
    - 2007
2. What does concurrency mean?
    - A program is concurrent when tasks might be executed at the same time.
3. On average, Go show very long build time? True or False?
    - False. The language was created to tackle this precise issue.

## 6. Key takeaways

- Go was born in 2007

- Go version 1 was released in 2012

- The language is easy to understand. Its semantics remain simple.

- It is statically typed

- It is compiled

- You can write concurrent programs with Go.
