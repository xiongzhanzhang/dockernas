package config

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

func GetCaFileDir() string {
	return GetConfig("ca_dir", "")
}

func SetCaFileDir(path string) {
	SetConfig("ca_dir", path)
	SaveConfig()
}
