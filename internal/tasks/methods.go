package tasks

import (
	"time"
	"slices" // Requires Go 1.21+
	"strconv"
	"fmt"
)

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

func No_such_method(a string) {
	response := 
`No such method: %s
Please, enter one of available methods:
"add", "update", "delete", "mark-done",
"mark-in-progress", "list", "list done",
"list todo", "list in-progress"
`

	fmt.Printf(response, a)
}

func Add_method(a []string) {
	all_data, _ := Load_from_file()
	if len(a) < 2 {
		fmt.Println("Write description!")
		return
	} else if len(a) > 2 {
		fmt.Println("Write desctiption in brackets!")
		return
	}

	task := Task{len(all_data.Tasks)+1, a[1], "todo", time.Now(), time.Now()}
	all_data.Tasks = append(all_data.Tasks, task)

	Write_to_file("./internal/database/tasks_data.json", all_data)

	return
}

func Delete_method(a []string) {
	all_data, _ := Load_from_file()

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

	Write_to_file("./internal/database/tasks_data.json", all_data)
	
}

func Update_method(a []string) {
	all_data, _ := Load_from_file()

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

	Write_to_file("./internal/database/tasks_data.json", all_data)
}



