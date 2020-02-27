package main

import (
	"fmt"
	"os"
)

func updateChromePolicies() {

	sfi, _ := os.Stat("./URLBlacklist.json")

	fmt.Println(sfi)

	// dstLocation := "/etc/opt/chrome/policies/managed"
	// data, err := ioutil.ReadFile(dstLocation)

	// if err == nil {

	// }

}
