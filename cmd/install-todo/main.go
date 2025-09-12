package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Pfad zum Go-Bin-Verzeichnis
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = filepath.Join(os.Getenv("HOME"), "go", "bin")
	}

	// Prüfen, ob todo existiert
	todoPath := filepath.Join(gobin, "todo")
	if _, err := os.Stat(todoPath); err == nil {
		fmt.Printf("✅ Todo-Manager CLI gefunden: %s\n", todoPath)
	} else {
		fmt.Printf("⚠️ Todo-Manager CLI nicht gefunden im Go-Bin-Verzeichnis (%s)\n", gobin)
		fmt.Println("Bitte zuerst installieren: go install github.com/Tiagofvp/todo-manager/cmd/todo@latest")
		return
	}

	// Prüfen, ob gobin im PATH ist
	path := os.Getenv("PATH")
	if !containsPath(path, gobin) {
		fmt.Printf("⚠️ Achtung: %s ist nicht im PATH!\n", gobin)
		fmt.Printf("Füge es hinzu, z.B.:\n\n  export PATH=$PATH:%s\n\n", gobin)
		fmt.Println("Danach kannst du 'todo' von überall ausführen.")
	} else {
		fmt.Println("✅ PATH korrekt gesetzt, du kannst 'todo' direkt nutzen.")
	}

	// Optional: Testlauf von 'todo list'
	fmt.Println("\n==> Testlauf: todo list")
	cmd := exec.Command("todo", "list")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func containsPath(paths string, dir string) bool {
	for _, p := range filepath.SplitList(paths) {
		if p == dir {
			return true
		}
	}
	return false
}
