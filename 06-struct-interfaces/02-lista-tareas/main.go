package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) print() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de go",
		desc:      "Completar curso de go este mes",
		completed: false,
	}
	t2 := Task{
		name:      "Revision de tools",
		desc:      "Revisar las herramientas este mes",
		completed: true,
	}

	t1.print()
	t2.print()

	tl := TaskList{}
	tl.appendTask(&t1)
	tl.appendTask(&t2)
	fmt.Println(tl)

	t3 := Task{
		name:      "Curso de terraform",
		desc:      "Completar curso de terraform este mes",
		completed: false,
	}
	tl.appendTask(&t3)

	for i, task := range tl.tasks {
		fmt.Println(i, task.name)
	}
}
