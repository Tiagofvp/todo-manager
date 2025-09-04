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

func (list *List) Add(text string) {
	item := NewItem(text)
	list.Items = append(list.Items, item)
}

func (list *List) Complete(index int) error {
	if index < 0 || index >= len(list.Items) {
		return errors.New("item index out of range")
	}
	list.Items[index].Done = true

	return nil
}

func (list *List) Delete(index int) error {
	if index < 0 || index >= len(list.Items) {
		return errors.New("item index out of range")
	}
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
	return nil
}

func (list *List) Save(filename string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (list *List) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, list)
}

// String returns a formatted string representation of the list
func (list *List) String() string {
	if len(list.Items) == 0 {
		return "No items in the todo list"
	}

	result := "Todo List:\n"
	for i, item := range list.Items {
		status := " "
		if item.Done {
			status = "✓"
		}
		result += fmt.Sprintf("%d. [%s] %s\n", i+1, status, item.Text)
	}

	return result
}
