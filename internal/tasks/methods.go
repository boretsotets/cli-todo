package tasks

import (
	"time"
	"slices" // Requires Go 1.21+
	"strconv"
	"fmt"
	"sort"
)

type TaskList struct {
	Tasks []Task
}

type Task struct {
	Id int
	Description string
	Status string
	CreatedAt string
	UpdatedAt string
}

func NoSuchMethod(a string) {
	response := 
`No such method: %s
Please, enter one of available methods:
"add", "update", "delete", "mark-done",
"mark-in-progress", "list", "list done",
"list todo", "list in-progress"
`

	fmt.Printf(response, a)
}

func AddMethod(a []string) {
	all_data, _ := LoadFromFile()
	if len(a) < 2 {
		fmt.Println("Write description!")
		return
	} else if len(a) > 2 {
		fmt.Println("Write desctiption in brackets!")
		return
	}

	formattedTime := time.Now().Format("2006-01-02 15:04:05")

	var ids []int
	var id int
	for i := range all_data.Tasks {
		ids = append(ids, all_data.Tasks[i].Id)
	}
	sort.Ints(ids)
	
	for i := range ids {
		if i+1 != ids[i] {
			id = i+1
			break
		}
	}

	if id == 0 {
		id = len(ids)+1
	}

	task := Task{id, a[1], "todo", formattedTime, formattedTime}
	all_data.Tasks = append(all_data.Tasks, task)

	WriteToFile("./internal/database/tasks_data.json", all_data)

	return
}

func DeleteMethod(a []string) {
	all_data, _ := LoadFromFile()

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
			return
		}
	}

	WriteToFile("./internal/database/tasks_data.json", all_data)
	
}

func UpdateMethod(a []string) {
	all_data, _ := LoadFromFile()

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
			all_data.Tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		} else if i == len(all_data.Tasks)-1 {
			fmt.Println("There is no task with this id!")
			return
		}
	}

	WriteToFile("./internal/database/tasks_data.json", all_data)
}

const LISTALL int = 0
const LISTDONE int = 1
const LISTTODO int = 2
const LISTINPROGRESS int = 3

func ListMethod(a []string) {
	var flag int

	if len(a) > 2 {
		fmt.Println("write correct option for list method!")
		return
	} else if len(a) == 1 {
		flag = LISTALL
	} else {
		switch a[1] {
			case "done": flag = LISTDONE
			case "todo": flag = LISTTODO
			case "in-progress": flag = LISTINPROGRESS
			default: {
				NoSuchMethod(a[1])
				return
			    }
		}
	}
	PrintList(flag)
}

func PrintList(flag int) {

	all_data, _ := LoadFromFile()

	for i := range all_data.Tasks {
		if flag == LISTALL || (flag == LISTDONE && all_data.Tasks[i].Status == "done") ||
		(flag == LISTTODO && all_data.Tasks[i].Status == "todo") ||
		(flag == LISTINPROGRESS && all_data.Tasks[i].Status == "in-progress") {
			fmt.Println(all_data.Tasks[i].Id, all_data.Tasks[i].Description, 
				all_data.Tasks[i].Status, all_data.Tasks[i].CreatedAt,
			all_data.Tasks[i].UpdatedAt)
		}
	}
}

func MarkInProgress(a []string) {
	all_data, _ := LoadFromFile()

	if len(a) != 2 {
		fmt.Println("needed 2 parameters")
	} else if neededId, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("error parsing id")
	} else {
		for i := range all_data.Tasks {
			if all_data.Tasks[i].Id == neededId {
				all_data.Tasks[i].Status = "in-progress"
				all_data.Tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
				break
			} else if i == len(all_data.Tasks)-1 {
				fmt.Println("There is no task with this id!")
			}
		}
	}
	WriteToFile("./internal/database/tasks_data.json", all_data)
}

func MarkDone(a []string) {
	all_data, _ := LoadFromFile()
	if len(a) != 2 {
		fmt.Println("needed 2 parameters")
	} else if neededId, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("error parsing id")
	} else {
		for i := range all_data.Tasks {
			if all_data.Tasks[i].Id == neededId {
				all_data.Tasks[i].Status = "done"
				all_data.Tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
				break
			} else if i == len(all_data.Tasks)-1 {
				fmt.Println("There is no task with this id!")
			}
		}
	}
	WriteToFile("./internal/database/tasks_data.json", all_data)

}
