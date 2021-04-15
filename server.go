package doorbell

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", EventHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Hello, World!")
}
