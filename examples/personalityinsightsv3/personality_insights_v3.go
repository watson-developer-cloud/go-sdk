package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/personalityinsightsv3"
)

func main() {
	// Instantiate the Watson Personality Insights service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("YOUR API KEY"),
	}
	service, serviceErr := personalityinsightsv3.
		NewPersonalityInsightsV3(&personalityinsightsv3.PersonalityInsightsV3Options{
			URL:           "YOUR SERVICE URL",
			Version:       "2017-10-13",
			Authenticator: authenticator,
		})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	/* PROFILE */

	// Read txt file with example speech
	pwd, _ := os.Getwd()
	fileName := "personality-v3.txt"
	file, fileErr := ioutil.ReadFile(pwd + "/../../resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		panic(fileErr)
	}

	// Create a new ProfileOptions for ContentType "text/plain"
	profileOptions := service.
		NewProfileOptions().
		SetContentType("text/plain")
	profileOptions.SetBody(string(file))
	profileOptions.ContentLanguage = core.StringPtr("en")
	profileOptions.AcceptLanguage = core.StringPtr("en")

	// Call the personality insights Profile method
	profResult, _, responseErr := service.Profile(profileOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if profResult != nil {
		core.PrettyPrint(profResult, "Profile for "+fileName)
	}

	// Read JSON file with example tweets
	fileName = "personality-v3.json"
	file, fileErr = ioutil.ReadFile(pwd + "/../../resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		panic(fileErr)
	}

	// Unmarshal JSON into Content struct
	content := new(personalityinsightsv3.Content)
	json.Unmarshal(file, content)

	// Set Content of profileOptions
	profileOptions.Content = content

	// Call the personality insights Profile method now with JSON Content
	profResult, _, responseErr = service.Profile(profileOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if profResult != nil {
		core.PrettyPrint(profResult, "Profile for "+fileName)
	}

	/* PROFILE AS CSV */

	// Read txt file with example speech
	fileName = "personality-v3.txt"
	file, fileErr = ioutil.ReadFile(pwd + "/../../resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		panic(fileErr)
	}

	// Set text/plain of profileOptions
	profileOptions.SetBody(string(file))

	// Call the personality insights ProfileAsCsv method
	profCsvResult, _, responseErr := service.ProfileAsCsv(profileOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if profCsvResult != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(profCsvResult)
		fmt.Printf("Profile as CSV for %v\n%v\n", fileName, buff.String())

		file, _ := os.Create("profile_example.csv")
		file.Write(buff.Bytes())
		file.Close()
	}
	profCsvResult.Close()
}
