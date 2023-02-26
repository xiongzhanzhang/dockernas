package config

func GetBindAddr() string {
	return GetConfig("bindAddr", "127.0.0.1:8080")
}

func GetDomain() string {
	return GetConfig("domain", "")
}

func SetDomain(domain string) {
	SetConfig("domain", domain)
	SaveConfig()
}

func GetIsHttpsEnabled() bool {
	return GetConfig("https", "") == "true"
}

func EnableHttps() {
	SetConfig("https", "true")
	SaveConfig()
}

func DisableHttps() {
	SetConfig("https", "false")
	SaveConfig()
}

func EnableHttpGateway() {
	SetConfig("httpGateWay", "true")
	SaveConfig()
}

func DisableHttpGateway() {
	SetConfig("httpGateWay", "false")
	SaveConfig()
}
