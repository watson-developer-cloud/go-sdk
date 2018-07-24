package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	watson "golang-sdk"
	"golang-sdk/personalityinsightsv3"
)

func main() {
	// Instantiate the Watson Personality Insights service
	piV3, piV3Err := personalityinsightsv3.NewPersonalityInsightsV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2017-10-13",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if piV3Err != nil {
		fmt.Println(piV3Err)
		return
	}

	// Read file with example speech
	pwd, _ := os.Getwd()
	speech, speechErr := ioutil.ReadFile(pwd + "/resources/personality-v3.txt")

	// Check successful file read
	if speechErr != nil {
		fmt.Println(speechErr)
		return
	}

	// Create request for Profile method
	profileReq := personalityinsightsv3.Content{
		ContentItems: []personalityinsightsv3.ContentItem{
			{
				Content: string(speech),
			},
		},
	}

	// Call the personality insights Profile method
	prof, profErr := piV3.Profile(&profileReq, "text/plain;charset=utf-8", "en", "en", false, false, false)

	// Check successful call
	if profErr != nil {
		fmt.Println(profErr)
		return
	}

	// Cast response from call to the specific struct returned by GetProfileResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	profResult := personalityinsightsv3.GetProfileResult(prof)

	// Check successful casting
	if profResult != nil {
		// Print result with pretty indenting
		output, _ := json.MarshalIndent(profResult, "", "    ")
		fmt.Printf("PROFILE %+v\n", string(output))
	}
}
