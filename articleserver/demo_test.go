package articleserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/login", "http://localhost:8080/api"), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	fmt.Println("calling...")
	r, err := client.Do(req)
	if err != nil {
		fmt.Println("error occured here")
		return
	}
	fmt.Println(r)
	var tkn string
	// tkn = r.Header.Get("Token")
	var demo string
	err = json.NewDecoder(r.Body).Decode(&demo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(demo)
	fmt.Println(tkn)
	fmt.Println("login")
	r.Body.Close()
}