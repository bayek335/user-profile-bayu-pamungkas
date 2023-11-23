package app

type Photos struct {
	ID       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	Caption  string `db:"caption" json:"caption"`
	PhotoUrl string `db:"photo_url" json:"photo_url"`
	UserId   int    `db:"user_id" json:"user_id"`
}

type PhotoRequest struct {
	Title   string `form:"title" binding:"required"`
	Caption string `form:"caption" binding:"required"`
}

type PhotoResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserId   int    `json:"user_id"`
	} `json:"data"`
}
type PhotoResponseAll struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []*Photos `json:"data"`
}
