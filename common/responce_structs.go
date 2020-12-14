package common

type UnitForClient struct {
	Title       string `json: "title"`
	Description string `json: "description"`
}

type PostResult struct {
	Id int `json: "id"`
}
