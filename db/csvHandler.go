package db

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MahikaJaguste/todocli/schema"
)

const FilePath = "./tasks.csv"

func GetFilePath() (string, error) {
	filePath, filePathErr := filepath.Abs(FilePath)
	return filePath, filePathErr
}

func InitFile() error {
	filePath, filePathErr := GetFilePath()
	if filePathErr != nil {
		return filePathErr
	}

	file, fileErr := os.OpenFile(filePath, os.O_CREATE, 0644)
	if fileErr != nil {
		return fileErr
	}

	defer file.Close()
	return nil
}

func GetFile(isAppend bool) (*os.File, error) {
	filePath, filePathErr := GetFilePath()
	if filePathErr != nil {
		return nil, filePathErr
	}

	flag := os.O_RDONLY
	if isAppend {
		flag = os.O_WRONLY | os.O_APPEND
	}

	file, fileErr := os.OpenFile(filePath, flag, 0644)
	return file, fileErr
}

// returns any error, else the ID of this new task
func AppendTask(taskRow []string) error {
	file, fileErr := GetFile(true)
	if fileErr != nil {
		return fileErr
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	appendErr := writer.Write(taskRow)
	return appendErr
}

func GetTasks() (*[]schema.Task, error) {

	file, fileErr := GetFile(false)
	if fileErr != nil {
		return nil, fileErr
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	tasks, readErr := reader.ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	tasksResult := make([]schema.Task, len(tasks))
	for i, row := range tasks {

		id, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}

		tasksResult[i] = schema.Task{
			Id:          int8(id),
			Description: row[1],
		}
	}

	return &tasksResult, nil

}
