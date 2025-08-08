package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type List struct {
	Items []Item
}

type Item struct {
	Text string
	Done bool
}

func NewItem(text string) Item {
	return Item{
		Text: text,
		Done: false,
	}
}

func NewList() *List {
	return &List{
		Items: []Item{},
	}
}

func (l *List) Add(text string) {
	item := NewItem(text)
	l.Items = append(l.Items, item)
}

func (l *List) Complete(index int) error {
	if index < 0 || index >= len(l.Items) {
		return errors.New("item index out of range")
	}
	l.Items[index].Done = true

	return nil
}

func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (l *List) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, l)
}

// String returns a formatted string representation of the list
func (l *List) String() string {
	if len(l.Items) == 0 {
		return "No items in the todo list"
	}

	result := "Todo List:\n"
	for i, item := range l.Items {
		status := " "
		if item.Done {
			status = "✓"
		}
		result += fmt.Sprintf("%d. [%s] %s\n", i+1, status, item.Text)
	}

	return result
}
