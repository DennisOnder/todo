# todo
A simple todo app, created for learning purposes.

# Usage

~~~go
package main

import (
    "fmt"

    "github.com/adem/todo"
)

func main() {
    todo.NewTodo("Buy milk")
    todo.NewTodo("Fill out important documents")

    printTodos()
    // ☐  An older todo task created in an earlier session (id: 4)
    // ☐  Buy milk (id: 5)
    // ☐  Fill out important documents (id: 6)

    if todoItem := todo.FindByTitle("Buy milk"); todoItem != nil {
        todoItem.Done = true
    }

    printTodos()
    // ☐  An older todo task created in an earlier session (id: 4)
    // ☑  Buy milk (id: 5)
    // ☐  Fill out important documents (id: 6)
}

func printTodos() {
    for _, todo := range todo.Todos {
        fmt.Println(todo)
    }
}
~~~

# License
MIT
