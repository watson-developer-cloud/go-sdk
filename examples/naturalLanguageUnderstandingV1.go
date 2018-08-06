package main

import (
	"fmt"
	. "go-sdk/naturalLanguageUnderstandingV1"
	"encoding/json"
	"os"
	"io/ioutil"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Natural Language Understanding service
	nlu, nluErr := NewNaturalLanguageUnderstandingV1(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-02-27",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if nluErr != nil {
		fmt.Println(nluErr)
		return
	}


	/* ANALYZE */

	pwd, _ := os.Getwd()
	file, fileErr := ioutil.ReadFile(pwd + "/resources/energy-policy.html")

	// Check successful file read
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	analyzeOptions := NewAnalyzeOptions(Features{}).
		SetHTML(string(file))

	// Call the naturalLanguageUnderstanding Analyze method
	analyze, analyzeErr := nlu.Analyze(analyzeOptions)

	// Check successful call
	if analyzeErr != nil {
		fmt.Println(analyzeErr)
		return
	}

	// Cast analyze.Result to the specific dataType returned by Analyze
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	analyzeResult := GetAnalyzeResult(analyze)

	// Check successful casting
	if analyzeResult != nil {
		prettyPrint(analyzeResult, "Analyze")
	}
}
