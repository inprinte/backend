package structures

type User struct {
	Id_user   int    `json:"id_user"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Password  string `json:"password "`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Adress    Adress `json:"adress"`
}

type JsonResponseUsers struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}
