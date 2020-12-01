package evaluate

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func GetProcCapabilities() {
	data, err := ioutil.ReadFile("/proc/self/status")
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(data))

	pattern := regexp.MustCompile("(?i)capeff:\\s*?[a-z0-9]+\\s")
	params := pattern.FindStringSubmatch(string(data))
	for _, matched := range params {
		log.Println("Capabilities:")
		fmt.Printf("\t%s\n", matched)
		if strings.Contains(matched, "3fffffffff") {
			fmt.Println("Critical - Possible Privileged Container Found.")
		}
	}
}
