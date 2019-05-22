# todo-app

A simple todo app, created for learning purposes.

# Usage

> Create your main/entry file in the src/directory, and paste in the code below:

```go
package main

import (
    "fmt"

    "github.com/<YOUR_USERNAME>/todo/src/todos"
)

func main() {
    todos.AddTodo("Buy milk")
    todos.AddTodo("Fill out important documents")

    printTodos()
    // ☐  An older todo task created in an earlier session (id: 4)
    // ☐  Buy milk (id: 5)
    // ☐  Fill out important documents (id: 6)

    if todoItem := todos.FindByTitle("Buy milk"); todoItem != nil {
        todoItem.Done = true
    }

    printTodos()
    // ☐  An older todo task created in an earlier session (id: 4)
    // ☑  Buy milk (id: 5)
    // ☐  Fill out important documents (id: 6)
}

func printTodos() {
    for _, todo := range todos.Todos {
        fmt.Println(todo)
    }
}
```

# License

MIT
