package config

func GetUserInfo() (string, string) {
	return GetConfig("user", "admin"), GetConfig("passwd", "admin")
}
