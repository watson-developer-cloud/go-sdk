package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IBM/go-sdk-core/core"
	nlu "github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

func main() {
	// Instantiate the Watson Natural Language Understanding service
	service, serviceErr := nlu.
		NewNaturalLanguageUnderstandingV1(&nlu.NaturalLanguageUnderstandingV1Options{
			URL:       "YOUR SERVICE URL",
			Version:   "2017-02-27",
			IAMApiKey: "YOUR API KEY",
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
	response, responseErr := service.Analyze(analyzeOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Print the entire detailed response
	fmt.Println(response)

	// Cast analyze.Result to the specific dataType returned by Analyze
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	analyzeResult := service.GetAnalyzeResult(response)

	// Check successful casting
	if analyzeResult != nil {
		core.PrettyPrint(analyzeResult, "Analyze")
	}
}
