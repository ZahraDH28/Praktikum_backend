package config

var AllowedOrigins = []string{
	"http://localhost:5000",
	"http://localhost:8080",
	"https://indrariksa.github.io",
	"https://frontend-coba-dudm6gdfx-zahradh28s-projects.vercel.app/",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}

