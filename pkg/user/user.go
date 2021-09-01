package user

// cущность JSON для запроса клиента
type User struct {
	Login     string `json:"login"`
	Passworld string `json:""`
	LastName  string `json:"lastname"`
}

// cущность для получения данных из БД
type UserDB struct {
	Status   int    `json:"status"`
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

