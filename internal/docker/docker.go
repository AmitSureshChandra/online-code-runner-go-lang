package docker

var containerMap = map[string]string{
	"jdk8":  "online-compiler-jdk8",
	"jdk20": "online-compiler-jdk8",
}

func GetContainerName(compiler string) string {
	return containerMap[compiler]
}
