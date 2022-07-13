package go

type Todolist struct {
	Id 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id 			int    `json:"id"`
	I 		string `json:"title"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id 			int     `json:"id"`
	Title 		string  `json:"title"`
	Description string  `json:"description"`
	Done		bool	`json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int 
}