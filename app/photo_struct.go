package app

type Photos struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Caption  string `db:"caption"`
	PhotoUrl string `db:"photo_url"`
	UserId   int    `db:"user_id"`
}

type PhotoRequest struct {
	Caption string      `db:"caption" binding:"required"`
	Photo   interface{} `json:"photo" binding:"required,file"`
	UserId  int         `db:"user_id" binding:"required"`
}

type PhotoResponseSuccess struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserId   int    `json:"user_id"`
	} `json:"data"`
}
