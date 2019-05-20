package todo

import (
	"fmt"
	"log"
	"strings"
)

var Todos []Todo
var nextID int

// Todo denotes a single todo item.
type Todo struct {
	id    int
	Title string
	Done  bool
}

func init() {
	// Simulating a previously created task loaded from disk
	log.Print("Loading existing todo items from disk...")
	// TODO: The item's id should match the Todo slice's index, to be able to
	// retrieve todo items by id in O(1). Also, the slice allocation size
	// should be determined by the number of existing todo items.
	Todos = make([]Todo, 0)
	Todos = append(Todos, Todo{
		4,
		"An older todo task created in an earlier session",
		false,
	})
	// The next ID should be the highest ID of all existing todo items, plus 1
	for _, todo := range Todos {
		if nextID <= todo.id {
			nextID = todo.id + 1
		}
	}
}

func (t Todo) String() string {
	var checkbox rune
	if t.Done {
		checkbox = '☑'
	} else {
		checkbox = '☐'
	}
	return fmt.Sprintf("%c  %s (id: %d)", checkbox, t.Title, t.id)
}

// AddTodo adds a new todo item with the given title.
func AddTodo(title string) {
	newTodo := Todo{Title: title}
	newTodo.id = nextID
	nextID++
	Todos = append(Todos, newTodo)
}

// FindByTitle returns the first todo item that contains the given search term.
// Returns nil if no matching todo item was found.
func FindByTitle(searchTerm string) *Todo {
	for i, todo := range Todos {
		if strings.Contains(todo.Title, searchTerm) {
			return &Todos[i]
		}
	}
	return nil
}

// Save writes the in-memory todo slice persistently to disk.
func Save() {
	// TODO: Implementation.
}
