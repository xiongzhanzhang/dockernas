package config

func GetDomain() string {
	return GetConfig("domain", "")
}

func SetDomain(domain string) {
	SetConfig("domain", domain)
	SaveConfig()
}

func GetGateWayInstanceName() string{
	return "http_gateway"
}
