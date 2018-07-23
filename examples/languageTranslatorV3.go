package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/languageTranslatorV3"
	"golang-sdk/resources"
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


	/* TRANSLATE */

	entryTranslate := languageTranslatorV3.TranslateRequest{
		Text: []string {"Let's translate this message"},
		ModelId: "es-en",
	}

	// Call the languageTranslator Translate method
	translate, translateErr := languageTranslator.Translate(&entryTranslate)

	// Check successful call
	if translateErr != nil {
		fmt.Println(translateErr)
		return
	}

	// Cast response from call to the specific struct returned by GetTranslateResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	translateResult := languageTranslatorV3.GetTranslateResult(translate)

	// Check successful casting
	if translateResult != nil {
		// Print result
		fmt.Println(translateResult)
	}


	/* LIST IDENTIFIABLE LANGUAGES */

	// Call the languageTranslator List Identifiable Languages method
	listLanguage, listLanguageErr := languageTranslator.ListIdentifiableLanguages()

	// Check successful call
	if listLanguageErr != nil {
		fmt.Println(listLanguageErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListIdentifiableLanguagesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listLanguageResult := languageTranslatorV3.GetListIdentifiableLanguagesResult(listLanguage)

	// Check successful casting
	if listLanguageResult != nil {
		// Print result
		fmt.Println(listLanguageResult)
	}


	/* IDENTIFY */

	entryIdentify := "What language is this message in?"

	// Call the languageTranslator Identify method
	identify, identifyErr := languageTranslator.Identify(&entryIdentify)

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


	/* LIST MODELS */

	// Call the languageTranslator List Models method
	listModel, listModelErr := languageTranslator.ListModels("YOUR SOURCE LANGUAGE", "YOUR TARGET LANGUAGE", true)

	// Check successful call
	if listModelErr != nil {
		fmt.Println(listModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListModelsResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listModelResult := languageTranslatorV3.GetListModelsResult(listModel)

	// Check successful casting
	if listModelResult != nil {
		// Print result
		fmt.Println(listModelResult)
	}


	/* CREATE MODEL */

	// Call the languageTranslator Create Model method
	createModel, createModelErr := languageTranslator.CreateModel("YOUR MODEL ID", "YOUR MODEL NAME", "YOUR FORCED GLOSSARY", "YOUR PARALLEL CORPUS")

	// Check successful call
	if createModelErr != nil {
		fmt.Println(createModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetCreateModelsResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	createModelResult := languageTranslatorV3.GetIdentifyResult(createModel)

	// Check successful casting
	if createModelResult != nil {
		// Print result
		fmt.Println(createModelResult)
	}


	/* DELETE MODEL */

	// Call the languageTranslator Delete Model method
	deleteModel, deleteModelErr := languageTranslator.DeleteModel("9f8d9c6f-2123-462f-9793-f17fdcb77cd6")

	// Check successful call
	if deleteModelErr != nil {
		fmt.Println(deleteModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetDeleteModelResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	deleteModelResult := languageTranslatorV3.GetDeleteModelResult(deleteModel)

	// Check successful casting
	if deleteModelResult != nil {
		// Print result
		fmt.Println(deleteModelResult)
	}


	/* GET MODEL */

	// Call the languageTranslator Get Model method
	getModel, getModelErr := languageTranslator.GetModel("9f8d9c6f-2123-462f-9793-f17fdcb77cd6")

	// Check successful call
	if getModelErr != nil {
		fmt.Println(getModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetGetModelResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	getModelResult := languageTranslatorV3.GetGetModelResult(getModel)

	// Check successful casting
	if getModelResult != nil {
		// Print result
		fmt.Println(getModelResult)
	}
}
