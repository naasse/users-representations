package representations

var GetUsersRel = "getUsers"

type User struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}
