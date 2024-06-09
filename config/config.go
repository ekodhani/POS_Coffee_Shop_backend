package config

func Port() string {
	return ":8081"
}

func AllowAkses() []string {
	return []string{
		"http://localhost:3002", "http://localhost", "http://192.168.1.47:3002", "http://192.168.1.47"}
}
