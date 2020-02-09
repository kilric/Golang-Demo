package util

func (user User) GetAge() string {
	return user.Name
}

func (user User) GetId() int {
	return user.Id
}
func (user User) GetName() string {
	return user.Name
}
func (user *User) SetId(id int) bool {
	if id > 0 {
		user.Id = id
		return true
	} else {
		return false
	}
}

func (user *User) SetName(name string) bool {
	if name != "" {
		user.Name = name
		return true
	}
	return false
}

func (admin Admin) GetName() string {
	return admin.GetName() //直接利用User的方法
}

func (admin Admin) GetId() int {
	return admin.Id
}

func (admin Admin) GetLevel() string {
	return admin.level
}
func (admin *Admin) SetId(id int) bool {
	if id > 0 {
		admin.User.Id = id
		return true
	}
	return false
}

func (admin *Admin) SetName(name string) bool {
	if name != "" {
		admin.Name = name
		return true
	}
	return false
}

func (admin *Admin) SetLevel(level string) bool {
	if level != "" {
		admin.level = level
		return true
	}
	return false
}
