package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	nlu "github.com/watson-developer-cloud/go-sdk/v2/naturallanguageunderstandingv1"
)

func main() {
	// Instantiate the Watson Natural Language Understanding service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("YOUR IAM API KEY"),
	}
	service, serviceErr := nlu.
		NewNaturalLanguageUnderstandingV1(&nlu.NaturalLanguageUnderstandingV1Options{
			URL:           "YOUR SERVICE URL",
			Version:       core.StringPtr("2017-02-27"),
			Authenticator: authenticator,
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* ANALYZE */

	pwd, _ := os.Getwd()
	file, fileErr := ioutil.ReadFile(pwd + "/../../resources/energy-policy.html")

	// Check successful file read
	if fileErr != nil {
		panic(fileErr)
	}

	analyzeOptions := service.NewAnalyzeOptions(&nlu.Features{
		Entities: &nlu.EntitiesOptions{},
		Keywords: &nlu.KeywordsOptions{},
	}).
		SetHTML(string(file))

	// Call the naturalLanguageUnderstanding Analyze method
	analyzeResult, response, responseErr := service.Analyze(analyzeOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Print the entire detailed response
	fmt.Println(response)

	// Check successful casting
	if analyzeResult != nil {
		core.PrettyPrint(analyzeResult, "Analyze")
	}
}
