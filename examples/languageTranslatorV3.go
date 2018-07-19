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
	language, languageErr := languageTranslator.ListIdentifiableLanguages()

	// Check successful call
	if languageErr != nil {
		fmt.Println(languageErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListIdentifiableLanguagesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	resultLanguage := languageTranslatorV3.GetListIdentifiableLanguagesResult(language)

	// Check successful casting
	if resultLanguage != nil {
		// Print result
		fmt.Println(resultLanguage)
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
	resultIdentify := languageTranslatorV3.GetIdentifyResult(identify)

	// Check successful casting
	if resultIdentify != nil {
		// Print result
		fmt.Println(resultIdentify)
	}
}
