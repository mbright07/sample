package handler

func CheckEmptyEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	return true
}

func CheckPassword(password string) bool {
	if len(password) > 3 {
		return true
	}
	return false
}

func CheckExistedUser(email string, password string) bool {
	user := GetUserByEmail(email)
	if (user == User{}) {
		return false
	}

	if !CheckPasswordHash(password, user.Password) {
		return false
	}

	return true
}

func CheckNotExistedUser(email string, password string) bool {
	user := GetUserByEmail(email)
	if (user == User{}) {
		return true
	}

	return false
}

func CheckLoginUser(email string, password string) bool {
	if !CheckEmptyEmail(email) || !CheckPassword(password) || !CheckExistedUser(email, password) {
		return false
	}
	return true
}

func CheckSignupUser(email string, password string) bool {
	if !CheckEmptyEmail(email) || !CheckPassword(password) || !CheckNotExistedUser(email, password) {
		return false
	}
	return true
}
