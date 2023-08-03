package docker

var containerMap = map[string]string{
	"jdk8":     "online-compiler-jdk8",
	"jdk20":    "online-compiler-jdk8",
	"gcc11":    "online-compiler-gcc11",
	"golang12": "online-compiler-golang12",
	"python3":  "online-compiler-python3",
	"node20":   "online-compiler-node20",
}

var compilerMap = map[string]string{
	"jdk8":     "Java 8",
	"jdk20":    "Java 20",
	"gcc11":    "c/c++",
	"golang12": "Golang",
	"python3":  "Python 3",
	"node20":   "Node 20",
}

var fileNameMap = map[string]string{
	"jdk8":     "Solution.java",
	"jdk20":    "Solution.java",
	"gcc11":    "main.cpp",
	"golang12": "main.go",
	"python3":  "solution.py",
	"node20":   "solution.js",
}

func GetContainerName(compiler string) string {
	return containerMap[compiler]
}

func GetCompilers() map[string]string {
	return compilerMap
}

func GetFileName(compiler string) string {
	return fileNameMap[compiler]
}
