package data

var Users = []User{
	{Id: 1, FirstName: "dexter", LastName: "rain", Age: 23},
	{Id: 1, FirstName: "dennis", LastName: "rain", Age: 13},
}

type User struct {
	Id        int64    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
