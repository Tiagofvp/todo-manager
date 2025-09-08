package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"to-do-app/internal/todo"
)

const (
	todoFile = "todos.json"
)

func main() {
	// Define flags
	interactiveFlag := flag.Bool("i", false, "Run in interactive mode")
	helpFlag := flag.Bool("h", false, "Show help information")

	// Parse flags but keep access to non-flag arguments
	flag.Parse()
	args := flag.Args()

	// Load existing todos
	todoList := todo.NewList()
	if _, err := os.Stat(todoFile); err == nil {
		if err := todoList.Load(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error loading todos:", err)
			os.Exit(1)
		}
	}
	// Handle interactive mode flag
	if *interactiveFlag {
		runInteractive(todoList)
		return
	}

	// Show help if the -h flag is provided
	if *helpFlag {
		printHelp()
		return
	}

	if len(args) == 0 {
		fmt.Println(todoList)
		return
	}
	// Get the command
	command := args[0]

	switch command {
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: missing todo text")
			os.Exit(1)
		}

		text := strings.Join(args[1:], " ")

		todoList.Add(text)

		saveTodos(todoList)

		fmt.Println("Task added:", text)
	case "list":
		fmt.Println(todoList)
	case "complete":
		if len(args) < 2 {
			fmt.Println("Error: missing item number")
			os.Exit(1)
		}

		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid item number:", args[1])
			os.Exit(1)
		}

		if err := todoList.Complete(num - 1); err != nil {
			fmt.Fprintln(os.Stderr, "Error completing todo:", err)
			os.Exit(1)
		}

		saveTodos(todoList)

		fmt.Println("Marked item as completed")

	case "unmark":
		if len(args) < 2 {
			fmt.Println("Error: missing item number")
			os.Exit(1)
		}

		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid item number:", args[1])
			os.Exit(1)
		}

		if err := todoList.Unmark(num - 1); err != nil {
			fmt.Fprintln(os.Stderr, "Error completing todo:", err)
			os.Exit(1)
		}

		saveTodos(todoList)

		fmt.Println("Marked item as uncompleted")

	case "delete":
		if len(args) < 2 {
			fmt.Println("Error: missing item number")
		}
		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: invalid item number:", args[1])
			os.Exit(1)
		}

		if err := todoList.Delete(num - 1); err != nil {
			fmt.Println(os.Stderr, "Error deleting todo:", err)
			os.Exit(1)
		}
		saveTodos(todoList)
		fmt.Println("Deleted To-Do", num)
	case "help":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'todo help' or 'todo -h' for usage information")
		os.Exit(1)
	}
}

func saveTodos(list *todo.List) {
	if err := list.Save(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving todos:", err)
		os.Exit(1)
	}
}

func printHelp() {
	helpText := `
Todo - A simple command line todo manager

Usage:
  todo [command] [arguments]
  todo [flags]

Commands:
  add <text>     	Add a new todo item
  list           	List all todo items
  complete <number>   	Mark item n as completed
  unmark <number>    	Unmark item n as uncompleted
  delete <number>	Delete an item 
  help           	Show this help message

Flags:
  -h             	Show this help message
  -i             	Run in interactive mode
`
	fmt.Println(helpText)
}

func runInteractive(list *todo.List) {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("\n ________         _______                     __        __              __     \n" +
			"/        |       /       \\                   /  |      /  |            /  |    \n" +
			"$$$$$$$$/______  $$$$$$$  |  ______          $$ |      $$/   _______  _$$ |_   \n" +
			"   $$ | /      \\ $$ |  $$ | /      \\  ______ $$ |      /  | /       |/ $$   |  \n" +
			"   $$ |/$$$$$$  |$$ |  $$ |/$$$$$$  |/      |$$ |      $$ |/$$$$$$$/ $$$$$$/   \n" +
			"   $$ |$$ |  $$ |$$ |  $$ |$$ |  $$ |$$$$$$/ $$ |      $$ |$$      \\   $$ | __ \n" +
			"   $$ |$$ \\__$$ |$$ |__$$ |$$ \\__$$ |        $$ |_____ $$ | $$$$$$  |  $$ |/  |\n" +
			"   $$ |$$    $$/ $$    $$/ $$    $$/         $$       |$$ |/     $$/   $$  $$/ \n" +
			"   $$/  $$$$$$/  $$$$$$$/   $$$$$$/          $$$$$$$$/ $$/ $$$$$$$/     $$$$/  \n" +
			"                                                                               \n")
		fmt.Println("\nThis is the interactive mode of the ToDo-List-Manager")
		fmt.Println("\nUse following commands:")
		fmt.Println("  list 				- List all items")
		fmt.Println("  add [text]			- Add an item")
		fmt.Println("  complete [number]		- Set an item to completed")
		fmt.Println("  unmark [number]		- Set an item to uncompleted")
		fmt.Println("  delete [number]		- Delete an item")
		fmt.Println("  help				- Show this help message")
		fmt.Println("  quit/exit			- Exit interactive mode")
		fmt.Println("\n\n", list)

		for {
			fmt.Print("\n> ")

			if !scanner.Scan() {
				break
			}

			input := scanner.Text()
			parts := strings.SplitN(input, " ", 2)
			cmd := parts[0]

			switch cmd {
			case "add":
				if len(parts) < 2 {
					fmt.Println("Error: missing todo text")
					continue
				}
				list.Add(parts[1])
				saveTodos(list)
				fmt.Println("Added:", parts[1])
			case "list":
				fmt.Println(list)
			case "complete":
				if len(parts) < 2 {
					fmt.Println("Error: missing item number")
					continue
				}
				num, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Error: invalid item number")
					continue
				}
				if err := list.Complete(num - 1); err != nil {
					fmt.Println("Error:", err)
					continue
				}
				saveTodos(list)
				fmt.Println("Marked item as completed")
			case "unmark":
				if len(parts) < 2 {
					fmt.Println("Error: missing item number")
					continue
				}
				num, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Error: invalid item number")
					continue
				}
				if err := list.Unmark(num - 1); err != nil {
					fmt.Println("Error:", err)
					continue
				}
				saveTodos(list)
				fmt.Println("Marked item as uncompleted")
			case "delete":
				if len(parts) < 2 {
					fmt.Println("Error: missing item number")
					continue
				}
				num, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Error: invalid item number")
					continue
				}
				if err := list.Delete(num - 1); err != nil {
					fmt.Println("Error:", err)
					continue
				}
				saveTodos(list)
				fmt.Println("Deleted item", num)
			case "quit", "exit":
				return
			case "help":
				fmt.Println("\nAvailable commands:")
				fmt.Println("  add <text>			- Add a new todo")
				fmt.Println("  list				- List all todos")
				fmt.Println("  complete <number>		- Mark item <n> as completed")
				fmt.Println("  delete <number>		- Delete item <n> from ToDo-List")
				fmt.Println("  help				- Show this help message")
				fmt.Println("  quit/exit			- Exit the program")
			default:
				fmt.Println("Unknown command:", cmd)
				fmt.Println("Type 'help' for available commands")
			}
		}
	}
}
