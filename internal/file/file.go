package file

import (
	"io/ioutil"
	"os"
)

func WriteFile(filePath string, content string) {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func CreateFolderIfNotExists(folder string) {
	err := os.Mkdir(folder, 0755)
	if err != nil {
		panic(err)
	} // 0755 is the permission mode (read-write-execute for owner, read-execute for group and others)
}
