package structures

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
	Role       string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Get_user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Id       int    `json:"id"`
}

type Error struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type Customer_id struct {
	Id int `json:"id"`
}
