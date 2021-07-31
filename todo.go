package todo

type (
	TodoList struct {
		Id     int    `json:"-"`
		Titles string `json:"title"`
	}

	TodoItem struct {
		Id    int    `json:"-"`
		Title string `json:"title"`
		Note  string `json:"note"`
		Done  bool   `json:"done"`
	}

	ListItems struct {
		Id     int
		ListId int
		ItemId int
	}
)
