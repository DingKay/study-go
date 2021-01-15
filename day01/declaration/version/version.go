package version

import "fmt"

func init() {
	fmt.Println("version/version.go ==> init")
}

var Version = nowVersion()

func nowVersion() string {
	fmt.Println("version/version.go nowVersion")
	return "v1.2.1"
}
