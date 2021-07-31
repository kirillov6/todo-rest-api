package todo

type (
	User struct {
		Id       int    `json:"-"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserLists struct {
		Id     int
		UserId int
		ListId int
	}
)
