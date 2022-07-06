package zip

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestUnzipAndGetNameOfFilesInZip(t *testing.T) {
	out, err := UnzipAndGetNameOfFilesInZip("test.zip", "./in")

	if err != nil {
		fmt.Println("Error in unzip file process ", err)
	}

	fmt.Println(out)

	for _, s := range out {

		if strings.Contains(s, "MOBILE") {
			mobile, _ := ioutil.ReadFile(out[0])

			fmt.Println(string(mobile))
		}

	}

}
