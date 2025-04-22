package internal

type CreateTask struct {
	Name string `json:name`
	Text string `json:text`
}

type TaskById struct {
	Id int `json:id`
}
