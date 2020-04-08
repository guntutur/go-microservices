package github

import (
	"encoding/json"
	"fmt"
	"github.com/guntutur/go-microservices/src/api/clients/restclient"
	"github.com/guntutur/go-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, GetAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	// since we are using ioutil.ReadAll func to read all the response body
	// we need some way to properly close the reader
	// it would be nice if we put at the very end of the func method, just before the final return
	// but this is the perks of go, go compiler would actually evaluate this particular defer keyword here,
	// let it be as it is, continuing the syntax execution and so on
	// and just after this very function is executing its final return
	// the response.Body.Close() will be executed,
	// and that's my ni99a, happened with the help of de-fucking-fer keyword
	// heads to TestNotDefer and TestDefer at github_provider_test to see it in action
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errorResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errorResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo sucessfull response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	return &result, nil
}
