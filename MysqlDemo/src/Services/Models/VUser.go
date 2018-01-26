package Models

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}
