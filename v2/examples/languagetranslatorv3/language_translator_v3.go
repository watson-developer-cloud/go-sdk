package main

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	languagetranslator "github.com/watson-developer-cloud/go-sdk/v2/languagetranslatorv3"
)

func main() {
	// Instantiate the Watson Language Translator service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("YOUR IAM API KEY"),
	}
	service, serviceErr := languagetranslator.
		NewLanguageTranslatorV3(&languagetranslator.LanguageTranslatorV3Options{
			URL:           "YOUR SERVICE URL",
			Version:       core.StringPtr("2018-02-16"),
			Authenticator: authenticator,
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
	translateResult, response, responseErr := service.Translate(translateOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	fmt.Println(response)

	// Check successful casting
	if translateResult != nil {
		fmt.Println("The word count is ", *translateResult.WordCount)
	}

	/* LIST IDENTIFIABLE LANGUAGES */

	listIdentifiableLanguagesOptions := service.NewListIdentifiableLanguagesOptions()

	// Call the languageTranslator ListIdentifiableLanguages method
	listLanguageResult, response, responseErr := service.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if listLanguageResult != nil {
		core.PrettyPrint(listLanguageResult, "Identifiable Languages")
	}

	/* IDENTIFY */

	textToIdentify := "What language is this message in?"
	identifyOptions := service.NewIdentifyOptions(textToIdentify)

	// Call the languageTranslator Identify method
	identifyResult, response, responseErr := service.Identify(identifyOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if identifyResult != nil {
		core.PrettyPrint(identifyResult, "Identify")
	}

	/* LIST MODELS */

	listModelsOptions := service.NewListModelsOptions().
		SetSource("es").
		SetTarget("en").
		SetDefault(true)

	// Call the languageTranslator ListModels method
	listModelResult, response, responseErr := service.ListModels(listModelsOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if listModelResult != nil {
		core.PrettyPrint(listModelResult, "Models")
	}

	/* CREATE MODEL */

	pwd, _ := os.Getwd()

	glossary, glossaryErr := os.Open(pwd + "/../../resources/glossary.tmx")
	fmt.Println(glossary)
	if glossaryErr != nil {
		fmt.Println(glossaryErr)
	}

	createModelOptions := service.NewCreateModelOptions("en-fr").
		SetName("custom-en-fr").
		SetForcedGlossary(glossary)

	// Call the languageTranslator CreateModel method
	createModelResult, response, responseErr := service.CreateModel(createModelOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if createModelResult != nil {
		core.PrettyPrint(createModelResult, "Create Model")
	}

	/* GET MODEL */

	// Call the languageTranslator GetModel method
	getModelOptions := service.NewGetModelOptions(*createModelResult.ModelID)
	getModelResult, response, getModelErr := service.GetModel(getModelOptions)

	// Check successful call
	if getModelErr != nil {
		panic(getModelErr)
	}

	// Check successful casting
	if getModelResult != nil {
		core.PrettyPrint(getModelResult, "Get Model")
	}

	/* DELETE MODEL */

	// Call the languageTranslator DeleteModel method
	deleteModelOptions := service.NewDeleteModelOptions(*getModelResult.ModelID)
	deleteModelResult, response, responseErr := service.DeleteModel(deleteModelOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}

	// Check successful casting
	if deleteModelResult != nil {
		core.PrettyPrint(deleteModelResult, "Delete Model")
	}
}
