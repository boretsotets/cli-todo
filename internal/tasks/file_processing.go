package tasks

import (
	"os"
	"encoding/json"
	"fmt"
)

func LoadFromFile() (TaskList, error) {
	all_data := TaskList{}
	file, err := os.Open("./internal/database/tasks_data.json")
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

func WriteToFile(filename string, all_data TaskList) {
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
