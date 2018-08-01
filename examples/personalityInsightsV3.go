package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	watson "golang-sdk"
	"golang-sdk/personalityinsightsv3"
	"bytes"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

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


	/* PROFILE */

	// Read txt file with example speech
	pwd, _ := os.Getwd()
	fileName := "personality-v3.txt"
	file, fileErr := ioutil.ReadFile(pwd + "/resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	// Create a new ProfileOptions for ContentType "text/plain"
	profileOptions := personalityinsightsv3.NewProfileOptionsForPlain(string(file)).
		SetContentLanguage("en").
		SetAcceptLanguage("en")

	// Call the personality insights Profile method
	prof, profErr := piV3.Profile(profileOptions)

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
		prettyPrint(profResult, "Profile for " + fileName)
	}

	// Read JSON file with example tweets
	fileName = "personality-v3.json"
	file, fileErr = ioutil.ReadFile(pwd + "/resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	// Unmarshal JSON into Content struct
	content := new(personalityinsightsv3.Content)
	json.Unmarshal(file, content)

	// Set Content of profileOptions
	profileOptions.SetContent(*content)

	// Call the personality insights Profile method now with JSON Content
	prof, profErr = piV3.Profile(profileOptions)

	// Check successful call
	if profErr != nil {
		fmt.Println(profErr)
		return
	}

	// Cast response again
	profResult = personalityinsightsv3.GetProfileResult(prof)

	// Check successful casting
	if profResult != nil {
		prettyPrint(profResult, "Profile for " + fileName)
	}


	/* PROFILE AS CSV */

	// Read txt file with example speech
	fileName = "personality-v3.txt"
	file, fileErr = ioutil.ReadFile(pwd + "/resources/" + fileName)

	// Check successful file read
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	// Set text/plain of profileOptions
	profileOptions.SetPlain(string(file))

	// Call the personality insights ProfileAsCsv method
	prof, profErr = piV3.ProfileAsCsv(profileOptions)

	// Check successful call
	if profErr != nil {
		fmt.Println(profErr)
		return
	}

	// Cast response
	profCsvResult := personalityinsightsv3.GetProfileAsCsvResult(prof)

	// Check successful casting
	if profCsvResult != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(profCsvResult)
		fmt.Printf("Profile as CSV for %v\n%v\n", fileName, buff.String())

		file, _ := os.Create("profile_example.csv")
		file.Write(buff.Bytes())
		file.Close()
	}
}
