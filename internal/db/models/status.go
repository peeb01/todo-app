package models

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "inprogress"
	StatusTest       Status = "test"
	StatusBlock      Status = "block"
	StatusDone       Status = "done"
)
