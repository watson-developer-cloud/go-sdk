package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/languageTranslatorV3"
)

func main() {
	// Instantiate the Watson Language Translator service
	languageTranslator, languageTranslatorErr := languageTranslatorV3.NewLanguageTranslatorV3(watson.Credentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-02-16",
		APIkey: "YOUR API KEY",
	})

	// Check successful instantiation
	if languageTranslatorErr != nil {
		fmt.Println(languageTranslatorErr)
		return
	}

	/* LIST IDENTIFIABLE LANGUAGES */

	// Call the languageTranslator List Identifiable Languages method
	list, listErr := languageTranslator.ListIdentifiableLanguages()

	// Check successful call
	if listErr != nil {
		fmt.Println(listErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListIdentifiableLanguagesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listResult := languageTranslatorV3.GetListIdentifiableLanguagesResult(list)

	// Check successful casting
	if listResult != nil {
		// Print result
		fmt.Println(listResult)
	}

	/* IDENTIFY */

	var entry string
	entry = "YOUR STRING TO IDENTIFY"

	// Call the languageTranslator Identify method
	identify, identifyErr := languageTranslator.Identify(&entry)

	// Check successful call
	if identifyErr != nil {
		fmt.Println(identifyErr)
		return
	}

	// Cast response from call to the specific struct returned by GetIdentifyResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	identifyResult := languageTranslatorV3.GetIdentifyResult(identify)

	// Check successful casting
	if identifyResult != nil {
		// Print result
		fmt.Println(identifyResult)
	}
}
