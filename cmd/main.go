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
	case "mark-done": tasks.MarkDone(cli_args)
	case "mark-in-progress": tasks.MarkInProgress(cli_args)
	case "list": tasks.ListMethod(cli_args)
	default: tasks.NoSuchMethod(cli_args[0])
	}

	return
}
