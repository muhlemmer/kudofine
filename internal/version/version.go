package version

var (
	version     = "dev"
	environment = "dev"
)

func SetVersion(v string) {
	version = v
}
func SetEnvironment(e string) {
	environment = e
}
func Version() string {
	return version
}
func Environment() string {
	return environment
}
