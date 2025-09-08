# To-Do Manager - CLI Task Management

![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue.svg)
![Version](https://img.shields.io/badge/CLI&nbsp;Version-1.0-blue.svg)
---

A lightweight, fast command-line task manager written in Go for efficient to-do list management right from your terminal.

---
## ✨ Features

- ✅ Create, edit, and delete tasks
- 📝 Mark tasks as complete/incomplete
- 🔍 Filter and search tasks
- 📁 Save tasks to local file
- 🚀 Fast execution with minimal overhead
- 📦 Zero external dependencies

---

## 📦 Installation

### Prerequisites
- [Go](https://golang.org/dl/) 1.16 or newer
- Terminal/Command Prompt access

---

###  Method 1: Install via Go (Recommended)
```bash
go install github.com/Tiagofvp/todo-manager@latest
```

Build an executable or use the already generated one
```bash
go build -o todo ./cmd/todo
```

#### Verify Go Binary Path
```bash
echo $PATH | grep -q "$(go env GOPATH)/bin" && echo "Go bin in PATH" || echo "Go bin NOT in PATH"
```
If not in PATH, add this to your shell config (.bashrc, .zshrc, etc.):
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```
---

###  Method 2: Manual Installation
```bash
go install github.com/Tiagofvp/todo-manager@latest
```

Go into the Project Folder
```bash
cd path/to/project/to-do-app
```

Build an executable or use the already generated one
```bash
go build -o todo ./cmd/todo
```

Move to a directory in your PATH:
#### Verify Go Binary Path
```bash
sudo mv todo-manager /usr/local/bin/<
```
