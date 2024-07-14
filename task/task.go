package task

// type TaskStatus int
//
// // Enum for tasks status
// const (
// 	Todo TaskStatus = iota
// 	InProgress
// 	Done
// )

type Task struct {
	title       string
	description string
	// status      TaskStatus
}

// func (t Task) Description() string {
// 	return t.description
// }
//
// func (t Task) Title() string {
// 	return t.title
// }

// Function used by bubble tea List to search element on.
func (t *Task) FilterValue() string {
	return t.title
}
