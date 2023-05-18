package configs

type ServiceInfoStruct struct {
	Host string
}

type Cors struct {
	AllowedHeaders     []string
	AllowedMethods     []string
	AllowedCredentials bool
	AllowedOrigin      []string
}
