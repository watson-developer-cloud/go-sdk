package main

import (
	"fmt"
	. "go-sdk/languageTranslatorV3"
	"os"
	"encoding/json"
)

func prettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Language Translator service
	languageTranslator, languageTranslatorErr := NewLanguageTranslatorV3(&ServiceCredentials{
		ServiceURL: "YOUR SERVICE URL",
		Version: "2018-02-16",
		APIkey: "YOUR SERVICE API KEY",
	})

	// Check successful instantiation
	if languageTranslatorErr != nil {
		fmt.Println(languageTranslatorErr)
		return
	}


	/* TRANSLATE */

	textToTranslate := []string{
		"Let's translate this message",
		"And this one",
	}

	translateOptions := NewTranslateOptions(textToTranslate).
		SetModelID("en-es")

	// Call the languageTranslator Translate method
	translate, translateErr := languageTranslator.Translate(translateOptions)

	// Check successful call
	if translateErr != nil {
		fmt.Println(translateErr)
		return
	}

	// Cast translate.Result to the specific dataType returned by Translate
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	translateResult := GetTranslateResult(translate)

	// Check successful casting
	if translateResult != nil {
		prettyPrint(translateResult, "Translation")
	}


	/* LIST IDENTIFIABLE LANGUAGES */

	listIdentifiableLanguagesOptions := NewListIdentifiableLanguagesOptions()

	// Call the languageTranslator ListIdentifiableLanguages method
	listLanguage, listLanguageErr := languageTranslator.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)

	// Check successful call
	if listLanguageErr != nil {
		fmt.Println(listLanguageErr)
		return
	}

	// Cast result
	listLanguageResult := GetListIdentifiableLanguagesResult(listLanguage)

	// Check successful casting
	if listLanguageResult != nil {
		prettyPrint(listLanguageResult, "Identifiable Languages")
	}


	/* IDENTIFY */

	textToIdentify := "What language is this message in?"
	identifyOptions := NewIdentifyOptions(textToIdentify)

	// Call the languageTranslator Identify method
	identify, identifyErr := languageTranslator.Identify(identifyOptions)

	// Check successful call
	if identifyErr != nil {
		fmt.Println(identifyErr)
		return
	}

	// Cast result
	identifyResult := GetIdentifyResult(identify)

	// Check successful casting
	if identifyResult != nil {
		prettyPrint(identifyResult, "Identify")
	}


	/* LIST MODELS */

	listModelsOptions := NewListModelsOptions().
		SetSource("es").
		SetTarget("en").
		SetDefaultModels(true)

	// Call the languageTranslator ListModels method
	listModel, listModelErr := languageTranslator.ListModels(listModelsOptions)

	// Check successful call
	if listModelErr != nil {
		fmt.Println(listModelErr)
		return
	}

	// Cast result
	listModelResult := GetListModelsResult(listModel)

	// Check successful casting
	if listModelResult != nil {
		prettyPrint(listModelResult, "Models")
	}


	/* CREATE MODEL */

	pwd, _ := os.Getwd()

	glossary, glossaryErr := os.Open(pwd + "/resources/glossary.tmx")
	if glossaryErr != nil {
		fmt.Println(glossaryErr)
	}

	createModelOptions := NewCreateModelOptions("en-fr").
		SetName("custom-en-fr").
		SetForcedGlossary(*glossary)

	// Call the languageTranslator CreateModel method
	createModel, createModelErr := languageTranslator.CreateModel(createModelOptions)

	// Check successful call
	if createModelErr != nil {
		fmt.Println(createModelErr)
		return
	}

	// Cast result
	createModelResult := GetCreateModelResult(createModel)

	// Check successful casting
	if createModelResult != nil {
		prettyPrint(createModelResult, "Create Model")
	}


	/* GET MODEL */

	// Call the languageTranslator GetModel method
	getModelOptions := NewGetModelOptions(createModelResult.ModelID)
	getModel, getModelErr := languageTranslator.GetModel(getModelOptions)

	// Check successful call
	if getModelErr != nil {
		fmt.Println(getModelErr)
		return
	}

	// Cast result
	getModelResult := GetGetModelResult(getModel)

	// Check successful casting
	if getModelResult != nil {
		prettyPrint(getModelResult, "Get Model")
	}


	/* DELETE MODEL */

	// Call the languageTranslator DeleteModel method
	deleteModelOptions := NewDeleteModelOptions(getModelResult.ModelID)
	deleteModel, deleteModelErr := languageTranslator.DeleteModel(deleteModelOptions)

	// Check successful call
	if deleteModelErr != nil {
		fmt.Println(deleteModelErr)
		return
	}

	// Cast result
	deleteModelResult := GetDeleteModelResult(deleteModel)

	// Check successful casting
	if deleteModelResult != nil {
		prettyPrint(deleteModelResult, "Delete Model")
	}
}
