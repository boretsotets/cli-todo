package main
import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	"slices" // Requires Go 1.21+
	"strconv"
)

func main() {
	cli_args := os.Args[1:]

	switch cli_args[0] {
	case "add": add_method(cli_args)
	case "update": update_method(cli_args[1:])
	case "delete": delete_method(cli_args)
	case "mark-done":
	case "mark-in-progress":
	case "list":
	default: no_such_method(cli_args[0])
	}

	return
}

func no_such_method(a string) {
	response := 
`No such method: %s
Please, enter one of available methods:
"add", "update", "delete", "mark-done",
"mark-in-progress", "list", "list done",
"list todo", "list in-progress"
`

	fmt.Printf(response, a)
}

type TaskList struct {
	Tasks []Task
}

type Task struct {
	Id int
	Description string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func load_from_file() (TaskList, error) {
	all_data := TaskList{}
	file, err := os.Open("./database/tasks_data.json")
	if err == nil || os.IsNotExist(err) {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&all_data)
	} else {
		fmt.Println(err)
		return all_data, err
	}
	return all_data, nil
}

func write_to_file(filename string, all_data TaskList) {
	writeFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer writeFile.Close()

	encoder := json.NewEncoder(writeFile)
	encoder.SetIndent("", "    ")
	encoder.Encode(all_data)

}

func add_method(a []string) {
	all_data, _ := load_from_file()
	if len(a) < 2 {
		fmt.Println("Write description!")
		return
	} else if len(a) > 2 {
		fmt.Println("Write desctiption in brackets!")
		return
	}

	task := Task{len(all_data.Tasks)+1, a[1], "todo", time.Now(), time.Now()}
	all_data.Tasks = append(all_data.Tasks, task)

	write_to_file("./database/tasks_data.json", all_data)

	return
}

func delete_method(a []string) {
	all_data, _ := load_from_file()

	index, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range all_data.Tasks {
		if all_data.Tasks[i].Id == index {
			all_data.Tasks = slices.Delete(all_data.Tasks, i, i+1)
			break
		} else if i == len(all_data.Tasks)-1 {
			fmt.Println("There is no task with this id!")
		}
	}

	write_to_file("./database/tasks_data.json", all_data)
	
}

func update_method(a []string) {
	all_data, _ := load_from_file()

	if len(a) != 2 {
		fmt.Println("Write task id and new descrition in brackets!")
		return
	}

	index, err := strconv.Atoi(a[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range all_data.Tasks {
		if all_data.Tasks[i].Id == index {
			all_data.Tasks[i].Description = a[1]
			break
		} else if i == len(all_data.Tasks)-1 {
			fmt.Println("There is no task with this id!")
		}
	}

	write_to_file("./database/tasks_data.json", all_data)
}
