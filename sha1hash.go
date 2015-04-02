package main // import "github.com/cloudspace/microservice-sha1hasher"

import  (
	"fmt"
	"os"
	"crypto/sha1"
	"encoding/json"
)

func main() {
	unEncryptedInput := []byte(os.Args[1])
	result := sha1.Sum(unEncryptedInput)
	var strResult = fmt.Sprintf("%x", result)
	var wrappedResults = map[string]string{
		"result": strResult,
	}

	encodedResults, _ := json.Marshal(wrappedResults)
	fmt.Println(string(encodedResults))
}