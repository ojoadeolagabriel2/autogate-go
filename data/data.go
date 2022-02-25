package data

var Users = []User{
	{Id: 1, FirstName: "dexter", LastName: "rain", Age: 23, Address: AddressLine{PostCode: "N165UA", Street: "94 Lordship Park, Flat 2. London"}},
	{Id: 2, FirstName: "dennis", LastName: "bain", Age: 13},
	{Id: 3, FirstName: "marcus", LastName: "gawain", Age: 13},
}

type AddressLine struct {
	PostCode string `json:"post_code"`
	Street string `json:"street"`
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address AddressLine `json:"address"`
}