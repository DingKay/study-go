package version

import "fmt"

var FerchVersion = fetchVersion()

func init() {
	fmt.Println("version/ferch.go ==> init")
}

func fetchVersion() string {
	fmt.Println("version/fetch.go fetchVersion")
	return Version
}
