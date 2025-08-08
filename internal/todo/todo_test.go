package todo

import (
	"os"
	"testing"
)

func TestCreatingNewItem(t *testing.T) {
	item := NewItem("Learn Go")

	if item.Text != "Learn Go" {
		t.Errorf("Expected text to be 'Learn Go', got '%s'", item.Text)
	}

	if item.Done {
		t.Error("New item should not be marked as done")
	}
}

func TestAddItem(t *testing.T) {
	list := NewList()

	list.Add("Go outside")

	if len(list.Items) != 1 {
		t.Errorf("Expected 1 item in list, got %d", len(list.Items))
	}

	if list.Items[0].Text != "Go outside" {
		t.Errorf("Expected text to be 'Go outside', got '%s'", list.Items[0].Text)
	}
}

func TestCompleteItem(t *testing.T) {
	list := NewList()
	list.Add("Write code")

	err := list.Complete(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !list.Items[0].Done {
		t.Error("Expected item to be marked as done")
	}

	// Test completing an item that doesn't exist
	err = list.Complete(1)
	if err == nil {
		t.Error("Expected error when completing non-existent item")
	}
}

func TestSaveAndLoad(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "todo-test")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}
	// Tear down
	defer os.Remove(tmpfile.Name())

	// Create and save a list
	list := NewList()
	list.Add("Task 1")
	list.Add("Task 2")
	list.Complete(0)

	if err := list.Save(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to save list: %v", err)
	}

	// Load the list from the file
	loadedList := NewList()
	if err := loadedList.Load(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to load list: %v", err)
	}

	// Verify the loaded list matches the original
	if len(loadedList.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(loadedList.Items))
	}

	if !loadedList.Items[0].Done {
		t.Error("Expected first item to be completed")
	}

	if loadedList.Items[0].Text != "Task 1" {
		t.Errorf("Expected text 'Task 1', got '%s'", loadedList.Items[0].Text)
	}
}
