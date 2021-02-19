package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/domain/repo"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/services"
	"github.com/Kungfucoding23/microservices-go/consuming_external_api/src/api/utils/errors"
)

var (
	success map[string]string
	failed  map[string]errors.APIError
)

type createRepoResult struct {
	Request repo.CreateRepoRequest
	Result  *repo.CreateRepoResponse
	Error   errors.APIError
}

func getRequests() []repo.CreateRepoRequest {
	result := make([]repo.CreateRepoRequest, 0)

	file, err := os.Open("requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//scan line by line until we reach the end of the file
		line := scanner.Text()
		// fmt.Println(line)
		request := repo.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}

	return result
}

func createRepo(buffer chan bool, output chan createRepoResult, request repo.CreateRepoRequest) {
	result, err := services.RepoService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer // take from the channel to release more space
}

func handleResults(wg *sync.WaitGroup, input chan createRepoResult) {
	//if we create a repo and we have an error, we need to review it
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func main() {
	requests := getRequests()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10) //this way we limit to 10 go routines at the same time
	var wg sync.WaitGroup

	go handleResults(&wg, input)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}
	wg.Wait()
	close(input)
	//Now you can write success and failed maps to disks or notify them via email or anything you need to do

}
