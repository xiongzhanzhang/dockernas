package config

func GetGateWayInstanceName() string {
	return "http_gateway"
}

func GetHostNameInStats() string {
	return "dockernas"
}

func IsInstanceNameConflict(name string) bool {
	if GetGateWayInstanceName() == name {
		return true
	}

	if GetHostNameInStats() == name {
		return true
	}

	return false
}
