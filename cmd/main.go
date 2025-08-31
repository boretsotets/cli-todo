package main
import (
	"os"
	"github.com/boretsotets/task-tracker/internal/tasks"
)

func main() {
	cli_args := os.Args[1:]

	switch cli_args[0] {
	case "add": tasks.AddMethod(cli_args)
	case "update": tasks.UpdateMethod(cli_args[1:])
	case "delete": tasks.DeleteMethod(cli_args)
	case "mark-done":
	case "mark-in-progress":
	case "list":
	default: tasks.NoSuchMethod(cli_args[0])
	}

	return
}

