# parrotsay

Inspired by [parrotsay](https://github.com/matheuss/parrotsay) and slightly overengineered. Definitely a very serious project.
Most of the code here is taken by [Flavio's great work](https://flaviocopes.com/go-tutorial-cowsay/).

## Build, Usage...

### Requirements
- Go version 1.21

### Build

Clone the repository
```
git clone https://github.com/Rocche/parrotsay.git
```

Then move to the cloned folder and build the project
```
go build .
```

### Install

To install, move to the project folder and install it with go
```
go install
```

### Usage

```
# without input at all: print a random party phrase
parrotsay

# with input as argument
parrotsay "OMG YES LET'S PARTY!"

# from stdin
fortune | parrotsay
echo "OMG YES LET'S PARTY!" | parrotsay
```

![image](https://github.com/Rocche/parrotsay/assets/37312278/48ebf21b-353e-440c-8476-994377f09da0)

## Architecture

### Use Case Diagram
![p](https://github.com/Rocche/parrotsay/assets/37312278/93405658-5bfe-4e53-b572-a4129876cf59)

### Package Diagram
![package](https://github.com/Rocche/parrotsay/assets/37312278/c21866ff-c96c-45e4-9d24-a302c720b2c7)

## Related

- [Cult of the Party Parrot](https://cultofthepartyparrot.com/)
- [Building a CLI command with Go: cowsay](https://flaviocopes.com/go-tutorial-cowsay/) (huge thanks!)

## Contribute

PRs are very welcome! Let's contribute to take this **very serious** project to the next level!
