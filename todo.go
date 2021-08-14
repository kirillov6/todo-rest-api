package todo

type (
	TodoList struct {
		Id    int    `json:"-"`
		Title string `json:"title"`
	}

	TodoItem struct {
		Id    int    `json:"-"`
		Title string `json:"title"`
		Note  string `json:"note"`
		Done  bool   `json:"done"`
	}

	ListsItems struct {
		Id     int
		ListId int
		ItemId int
	}
)
