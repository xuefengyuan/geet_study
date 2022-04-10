package user

type User struct {
    Name     string
    Password string
}

func NewUser()*User{
    return &User{}
}

func (u *User) Save() error {
    // 这里调用操作db保存用户信息
    return nil
}

func (u *User) GetUserInfo(id int64) error {
    // 这里调用操作db获取用户信息
    u.Name = "test name"
    return nil
}
