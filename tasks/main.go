package tasks

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/MahikaJaguste/todocli/db"
	"github.com/MahikaJaguste/todocli/schema"
)

var ColumnHeaders = []string{"Id", "Description"}

func CreateTask(taskDescription string) error {
	if len(taskDescription) < 3 {
		return errors.New("please enter valid task description after add, eg. add \"do math homework\"")
	}

	tasks, getErr := db.GetTasks()
	if getErr != nil {
		return getErr
	}

	var maxIndex int8 = 0
	for _, task := range *tasks {
		if task.Id > maxIndex {
			maxIndex = task.Id
		}
	}

	task := &schema.Task{
		Id:          maxIndex + 1,
		Description: taskDescription,
	}
	taskRow := GetTaskString(task)
	createErr := db.AppendTask(taskRow)
	return createErr
}

func GetTasks() error {
	tasks, getErr := db.GetTasks()
	if getErr != nil {
		return getErr
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 1, 8, ' ', tabwriter.TabIndent)

	s := ""
	for _, header := range ColumnHeaders {
		s += header + "\t"
	}
	fmt.Fprintln(w, s)

	// Print the CSV data
	for _, task := range *tasks {
		s = ""
		row := GetTaskString(&task)
		for _, col := range row {
			s += col + "\t"
		}
		fmt.Fprintln(w, s)
	}
	w.Flush()

	return nil
}

func GetTaskString(task *schema.Task) []string {
	taskRow := []string{
		strconv.Itoa(int(task.Id)),
		task.Description,
	}
	return taskRow
}
