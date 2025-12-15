package gitlab

import (
	"context"
	"log"
	"net/http"
	"testing"
)

func Test_Client(t *testing.T) {
	c, err := NewGitlabClient("https://gitlab.tha.kz/api/v4/projects/321/variables/COUNTER", "", "LTxRf3GYQdC-ASfjzhzb")
	if err != nil {
		log.Fatal(err)
	}

	req, err := c.Request(context.Background(), http.MethodPut, "v1.2.2-dev+6f69d972")
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Send(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(res))
}
