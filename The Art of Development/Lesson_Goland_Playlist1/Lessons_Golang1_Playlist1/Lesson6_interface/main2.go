package main

<<<<<<< HEAD
var _ User = superUser{}

type superUser struct {
	Name string
	Age  int
}

func (s *superUser) ChangeFIO(newFIO string) {
	//TODO implement me
	panic("implement me")
}

func (s *superUser) ChangeAddress(newAddress string) {
	//TODO implement me
	panic("implement me")
=======
var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	IsBlocked bool
}

func (s *superUser) Block() {
	u.IsBlocked = true
 origin/ main
}

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
<<<<<<< HEAD
}

func (u *user) ChangeFIO(newFIO string) {
	u.FIO = newFIO
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func NewUser(fio, address, phone string) {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
}
func main() {

=======
	IsBlocked           bool
}

func (u *user) Block() {
	u.IsBlocked = true
}

type User interface {
	Block()
}

func NewUser(fio, address, phone string) User {
	u := user{}
	return &u
}
func main() {
	u := NewUser("", "", "")
	u.Block()
>>>>>>> origin/ main
}
