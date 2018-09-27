package main

import (
	"encoding/json"
	"fmt"
	"os"

	languagetranslator "github.com/ibm-watson/go-sdk/languagetranslatorv3"
)

// PrettyPrint print pretty
func PrettyPrint(result interface{}, resultName string) {
	output, err := json.MarshalIndent(result, "", "    ")

	if err == nil {
		fmt.Printf("%v:\n%+v\n\n", resultName, string(output))
	}
}

func main() {
	// Instantiate the Watson Language Translator service
	service, serviceErr := languagetranslator.
		NewLanguageTranslatorV3(&languagetranslator.LanguageTranslatorV3Options{
			URL:       "YOUR SERVICE URL",
			Version:   "2018-02-16",
			IAMApiKey: "YOUR IAM API KEY",
		})

	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println(serviceErr)
		return
	}

	/* TRANSLATE */

	textToTranslate := []string{
		"Let's translate this message",
		"And this one",
	}

	translateOptions := service.NewTranslateOptions(textToTranslate).
		SetModelID("en-es")

	// Call the languageTranslator Translate method
	response, responseErr := service.Translate(translateOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// Cast translate.Result to the specific dataType returned by Translate
	// NOTE: most methods have a corresponding Get<methodName>Result() function
	translateResult := service.GetTranslateResult(response)

	// Check successful casting
	if translateResult != nil {
		fmt.Println("The word count is ", *translateResult.WordCount)
	}

	/* LIST IDENTIFIABLE LANGUAGES */

	listIdentifiableLanguagesOptions := service.NewListIdentifiableLanguagesOptions()

	// Call the languageTranslator ListIdentifiableLanguages method
	response, responseErr = service.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast result
	listLanguageResult := service.GetListIdentifiableLanguagesResult(response)

	// Check successful casting
	if listLanguageResult != nil {
		PrettyPrint(listLanguageResult, "Identifiable Languages")
	}

	/* IDENTIFY */

	textToIdentify := "What language is this message in?"
	identifyOptions := service.NewIdentifyOptions(textToIdentify)

	// Call the languageTranslator Identify method
	response, responseErr = service.Identify(identifyOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast result
	identifyResult := service.GetIdentifyResult(response)

	// Check successful casting
	if identifyResult != nil {
		PrettyPrint(identifyResult, "Identify")
	}

	/* LIST MODELS */

	listModelsOptions := service.NewListModelsOptions().
		SetSource("es").
		SetTarget("en").
		SetDefaultModels(true)

	// Call the languageTranslator ListModels method
	response, responseErr = service.ListModels(listModelsOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast result
	listModelResult := service.GetListModelsResult(response)

	// Check successful casting
	if listModelResult != nil {
		PrettyPrint(listModelResult, "Models")
	}

	/* CREATE MODEL */

	pwd, _ := os.Getwd()

	glossary, glossaryErr := os.Open(pwd + "/../resources/glossary.tmx")
	fmt.Println(glossary)
	if glossaryErr != nil {
		fmt.Println(glossaryErr)
	}

	createModelOptions := service.NewCreateModelOptions("en-fr").
		SetName("custom-en-fr").
		SetForcedGlossary(glossary)

	// Call the languageTranslator CreateModel method
	response, responseErr = service.CreateModel(createModelOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast result
	createModelResult := service.GetCreateModelResult(response)

	// Check successful casting
	if createModelResult != nil {
		PrettyPrint(createModelResult, "Create Model")
	}

	/* GET MODEL */

	// Call the languageTranslator GetModel method
	getModelOptions := service.NewGetModelOptions(*createModelResult.ModelID)
	response, getModelErr := service.GetModel(getModelOptions)

	// Check successful call
	if getModelErr != nil {
		panic(getModelErr)
	}

	// Cast result
	getModelResult := service.GetGetModelResult(response)

	// Check successful casting
	if getModelResult != nil {
		PrettyPrint(getModelResult, "Get Model")
	}

	/* DELETE MODEL */

	// Call the languageTranslator DeleteModel method
	deleteModelOptions := service.NewDeleteModelOptions(*getModelResult.ModelID)
	response, responseErr = service.DeleteModel(deleteModelOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Cast result
	deleteModelResult := service.GetDeleteModelResult(response)

	// Check successful casting
	if deleteModelResult != nil {
		PrettyPrint(deleteModelResult, "Delete Model")
	}
}
