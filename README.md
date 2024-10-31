# Let's create a README.md file with the provided content.

# Frate-Go

`Frate-Go` is a command-line tool designed for efficient C/C++ project management. It simplifies building, managing dependencies, handling CMake configurations, and working with templates and package repositories.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Available Commands](#available-commands)
  - [Global Flags](#global-flags)
- [Shell Completion](#shell-completion)
  - [Zsh Completion](#zsh-completion)
- [Examples](#examples)
- [Debugging Completion](#debugging-completion)
- [Contributing](#contributing)
- [License](#license)

---

## Installation

1. **Download and Install**: Make sure you have Go installed on your system. Then, install `frate-go` using:

   ```bash
   go install github.com/frate-dev/frate-go@latest
   ```

### Linux and Mac
   ```bash

   ```

## Usage 


### Available Command

- build - Build the project using CMake.
- completion - Generate the autocompletion script for the specified shell.
- dependency - Manage your project dependencies.
- generate - Generate or regenerate your project's configuration using CMake.
- help - Display help information about any command.
- init (alias: i) - Initialize a new Frate-Go project.
- packages - Search and push your project to a package repository.
- repo - Add, remove, or list package repositories.
- run - Build and run your project.
- template (alias: t) - Manage templates, create, list, delete templates, and manage template repositories.


## Shell Completion 


### Zsh Completion


```bash
frate-go completion zsh > ~/.zsh-completions/_frate-go
echo "fpath=(~/.zsh-completions $fpath)" >> ~/.zshrc
```


## Examples


```bash
frate-go init project
```


```bash
cd project
frate-go run 
```


```bash
frate-go template init new_template 
cd new_template
frate-go template push .
```

```bash
frate-go packages list 
frate-go packages push . 
```



