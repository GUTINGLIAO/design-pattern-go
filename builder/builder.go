package builder

type User struct {
	Name   string
	Age    int
	Sex    int
	Height int
	Weight int
	Father string
	Mother string
	Son    string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: User{},
	}
}

type UserBuilder struct {
	user User
}

func (u *UserBuilder) Build() *User {

	if u.user.Name == "" {
		print("warning, name parameter is required")
	}

	return &u.user
}

func (u *UserBuilder) Name(name string) *UserBuilder {
	u.user.Name = name
	return u
}

func (u *UserBuilder) Age(age int) *UserBuilder {
	u.user.Age = age
	return u
}

func (u *UserBuilder) Sex(sex int) *UserBuilder {
	u.user.Sex = sex
	return u
}

func (u *UserBuilder) Height(height int) *UserBuilder {
	u.user.Height = height
	return u
}

func (u *UserBuilder) Weight(weight int) *UserBuilder {
	u.user.Weight = weight
	return u
}

func (u *UserBuilder) Father(father string) *UserBuilder {
	u.user.Father = father
	return u
}

func (u *UserBuilder) Mother(mother string) *UserBuilder {
	u.user.Mother = mother
	return u
}

func (u *UserBuilder) Son(son string) *UserBuilder {
	u.user.Son = son
	return u
}
