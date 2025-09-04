package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	// "slices"
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
	fmt.Printf("%d\t|%s\t|%s\t|%s\t|\n", task.Id, task.Name, task.Description, task.Status)
}

func (tasks *TaskList) deleteTask(task int) {
	var tmp []Task
	count := 0
	for t := range tasks.Tasks {
		if tasks.Tasks[t].Id != task {
			tmp = append(tmp, tasks.Tasks[t])
		} else {
			count++
		}
	}
	tasks.Tasks = tmp
	if count == 0 {
		fmt.Println("task not found")
	}
}

func (tasks *TaskList) doubleCheck(taskname string) bool {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].Name == taskname {
			return true
		}
	}
	return false
}

func statusChanger(tasks TaskList, task int, status string) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].Id == task {
			tasks.Tasks[i].Status = status
			fmt.Printf("%d\t|%s\t|%s\t|%s\t|\n", tasks.Tasks[i].Id, tasks.Tasks[i].Name, tasks.Tasks[i].Description, tasks.Tasks[i].Status)
			break
		}
	}
}

func (tasks *TaskList) startTask(task int) {
	statusChanger(*tasks, task, "ongoing")
}

func (tasks *TaskList) finishTask(task int) {
	statusChanger(*tasks, task, "finished")
}

func (tasks *TaskList) blockTask(task int) {
	statusChanger(*tasks, task, "blocked")
}

func (tasks *TaskList) awaitTask(task int) {
	statusChanger(*tasks, task, "waiting")
}

func (tasks *TaskList) listtask() {
	for i := range tasks.Tasks {
		fmt.Printf("%d\t|%s\t|%s\t|%s\t|\n", tasks.Tasks[i].Id, tasks.Tasks[i].Name, tasks.Tasks[i].Description, tasks.Tasks[i].Status)
	}
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
}

func (tasks *TaskList) loadFromFile() {
	jsonData, err := os.ReadFile("tasks.json")
	if err != nil {
		panic("teste")
	}
	json.Unmarshal(jsonData, &tasks)
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
	listtask := flag.Bool("listtask", false, "List all task")
	flag.Parse()

	switch {
	case *listtask:
		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.listtask()
	case *inittasklist:
		tasklist := TaskList{}
		tasklist.writeToFile()

	case *addtask:
		tasklist := TaskList{}
		tasklist.loadFromFile()
		if tasklist.doubleCheck(*taskname) {
			fmt.Println("Task already exists")
			break
		}
		fmt.Println(*taskname)

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
		tasklist.loadFromFile()
		tasklist.startTask(*taskid)
		tasklist.writeToFile()

	case *finishtask:
		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.finishTask(*taskid)
		tasklist.writeToFile()
	case *blocktask:
		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.blockTask(*taskid)
		tasklist.writeToFile()
	case *awaittask:
		tasklist := TaskList{}
		tasklist.loadFromFile()
		tasklist.awaitTask(*taskid)
		tasklist.writeToFile()
	default:
		fmt.Println(`
	-addtask
			add a new task
			use this command with the option -taskname & -taskdescription
			Ex: -addtask -taskname="" -taskdescription=""
			
	-awaittask
			Declare a task as waiting
	-blocktask
			declare a task as blocked
	-deletetask
			delete a task
	-finishtask
			finish a task
	-inittasklist
			Init the file storage of the task
	-listtask
			List all task
	-starttask
			start a task 
	-taskdescription string
			the description of the task (default "description")
	-taskid int
			The id a task (default 1)
	-taskname string
			this is the name of the string (default "task")
		`)
	}
}
