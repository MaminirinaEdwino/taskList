package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
)

type Task struct {
	Id          int
	Name        string
	Description string
	Status      string
}
type TaskList struct {
	Tasks []Task
}

func (tasks *TaskList) addTask(task Task) {
	if len(tasks.Tasks) == 0 {
		task.Id = 1
	} else {
		task.Id = tasks.Tasks[len(tasks.Tasks)-1].Id + 1
	}
	task.Status = "en attente"
	tasks.Tasks = append(tasks.Tasks, task)
}

func (tasks *TaskList) deleteTask(task int) {
	for t := range tasks.Tasks {
		if tasks.Tasks[t].Id == task {
			tasks.Tasks = slices.Delete(tasks.Tasks, t, t+1)
		}
	}
}

func (tasks *TaskList) startTask(task int) {

}
func (tasks *TaskList) finishTask(task int) {

}
func (tasks *TaskList) blockTask(task int) {

}
func (tasks *TaskList) awaitTask(task int) {

}
func (tasks *TaskList) writeToFile() {
	taks, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic("error")
	}
	if len(tasks.Tasks) == 0 {
		os.WriteFile("tasks.json", taks, 0644)
	}
	os.WriteFile("tasks.json", taks, 0644)
	// file, err := os.Open("tasks.json")
	var tak TaskList
	jsonData, err := os.ReadFile("tasks.json")
	if err!= nil {
		panic("teste")
	}
	json.Unmarshal(jsonData, &tak)
	fmt.Println(tak.Tasks)
	
}
func (tasks *TaskList) loadFromFile() {

}

func main() {
	fmt.Println("Cli todoList vaovao")
	taskname := flag.String("taskname", "task", "this is the name of the string")
	taskdescription := flag.String("taskdescription", "description", "the description of the task")
	taskid := flag.Int("taskid", 1, "THe id a task")
	addtask := flag.Bool("addtask", false, "add a new task")
	deletetask := flag.Bool("deletetask", false, "delete a task")
	starttask := flag.Bool("starttask", false, "start a task ")
	finishtask := flag.Bool("finishtask", false, "finish a task")
	blocktask := flag.Bool("blocktask", false, "declare a task as blocked")
	awaittask := flag.Bool("awaittask", false, "Declare a task as waiting")
	inittasklist := flag.Bool("inittasklist", false, "Init the file storage of the task")
	flag.Parse()

	switch {
	case *inittasklist:
		tasklist := TaskList{}
		tasklist.writeToFile()
	case *addtask:
		fmt.Println(*taskname)

		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.addTask(Task{Name: *taskname, Description: *taskdescription})
		tasklist.writeToFile()

	case *deletetask:
		fmt.Println("delete task")
		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.deleteTask(*taskid)
		tasklist.writeToFile()

	case *starttask:
		tasklist := TaskList{}
		tasklist.startTask(0)
		tasklist.writeToFile()

	case *finishtask:
		tasklist := TaskList{}
		tasklist.finishTask(0)
		tasklist.writeToFile()
	case *blocktask:
		tasklist := TaskList{}
		tasklist.blockTask(0)
		tasklist.writeToFile()
	case *awaittask:
		tasklist := TaskList{}
		tasklist.awaitTask(0)
		tasklist.writeToFile()
	default:
		fmt.Println("ouin")
	}
}
