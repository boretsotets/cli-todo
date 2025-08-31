package main
import (
	"os"
	"github.com/boretsotets/task-tracker/internal/tasks"
)

func main() {
	cli_args := os.Args[1:]

	switch cli_args[0] {
	case "add": tasks.Add_method(cli_args)
	case "update": tasks.Update_method(cli_args[1:])
	case "delete": tasks.Delete_method(cli_args)
	case "mark-done":
	case "mark-in-progress":
	case "list":
	default: tasks.No_such_method(cli_args[0])
	}

	return
}

