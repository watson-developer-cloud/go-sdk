package main

import (
	"fmt"
	watson "golang-sdk"
	"golang-sdk/languagetranslatorv3"
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
	languageTranslator, languageTranslatorErr := languagetranslatorv3.NewLanguageTranslatorV3(watson.Credentials{
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

	translateOptions := languagetranslatorv3.NewTranslateOptions(textToTranslate).
		SetModelID("en-es")

	// Call the languageTranslator Translate method
	translate, translateErr := languageTranslator.Translate(translateOptions)

	// Check successful call
	if translateErr != nil {
		fmt.Println(translateErr)
		return
	}

	// Cast response from call to the specific struct returned by GetTranslateResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	translateResult := languagetranslatorv3.GetTranslateResult(translate)

	// Check successful casting
	if translateResult != nil {
		prettyPrint(translateResult, "Translation")
	}


	/* LIST IDENTIFIABLE LANGUAGES */

	listIdentifiableLanguagesOptions := languagetranslatorv3.NewListIdentifiableLanguagesOptions()

	// Call the languageTranslator ListIdentifiableLanguages method
	listLanguage, listLanguageErr := languageTranslator.ListIdentifiableLanguages(listIdentifiableLanguagesOptions)

	// Check successful call
	if listLanguageErr != nil {
		fmt.Println(listLanguageErr)
		return
	}

	// Cast response from call to the specific struct returned by GetListIdentifiableLanguagesResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listLanguageResult := languagetranslatorv3.GetListIdentifiableLanguagesResult(listLanguage)

	// Check successful casting
	if listLanguageResult != nil {
		prettyPrint(listLanguageResult, "Identifiable Languages")
	}


	/* IDENTIFY */

	textToIdentify := "What language is this message in?"
	identifyOptions := languagetranslatorv3.NewIdentifyOptions(textToIdentify)

	// Call the languageTranslator Identify method
	identify, identifyErr := languageTranslator.Identify(identifyOptions)

	// Check successful call
	if identifyErr != nil {
		fmt.Println(identifyErr)
		return
	}

	// Cast response from call to the specific struct returned by GetIdentifyResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	identifyResult := languagetranslatorv3.GetIdentifyResult(identify)

	// Check successful casting
	if identifyResult != nil {
		prettyPrint(identifyResult, "Identify")
	}


	/* LIST MODELS */

	listModelsOptions := languagetranslatorv3.NewListModelsOptions().
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

	// Cast response from call to the specific struct returned by GetListModelsResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	listModelResult := languagetranslatorv3.GetListModelsResult(listModel)

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

	createModelOptions := languagetranslatorv3.NewCreateModelOptions("en-fr").
		SetName("custom-en-fr").
		SetForcedGlossary(*glossary)

	// Call the languageTranslator CreateModel method
	createModel, createModelErr := languageTranslator.CreateModel(createModelOptions)

	// Check successful call
	if createModelErr != nil {
		fmt.Println(createModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetCreateModelResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	createModelResult := languagetranslatorv3.GetCreateModelResult(createModel)

	// Check successful casting
	if createModelResult != nil {
		prettyPrint(createModelResult, "Create Model")
	}


	/* GET MODEL */

	// Call the languageTranslator GetModel method
	getModelOptions := languagetranslatorv3.NewGetModelOptions(createModelResult.ModelID)
	getModel, getModelErr := languageTranslator.GetModel(getModelOptions)

	// Check successful call
	if getModelErr != nil {
		fmt.Println(getModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetGetModelResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	getModelResult := languagetranslatorv3.GetGetModelResult(getModel)

	// Check successful casting
	if getModelResult != nil {
		prettyPrint(getModelResult, "Get Model")
	}


	/* DELETE MODEL */

	// Call the languageTranslator DeleteModel method
	deleteModelOptions := languagetranslatorv3.NewDeleteModelOptions(getModelResult.ModelID)
	deleteModel, deleteModelErr := languageTranslator.DeleteModel(deleteModelOptions)

	// Check successful call
	if deleteModelErr != nil {
		fmt.Println(deleteModelErr)
		return
	}

	// Cast response from call to the specific struct returned by GetDeleteModelResult
	// NOTE: other than DELETE requests, every method has a corresponding Get<methodName>Result() function
	deleteModelResult := languagetranslatorv3.GetDeleteModelResult(deleteModel)

	// Check successful casting
	if deleteModelResult != nil {
		prettyPrint(deleteModelResult, "Delete Model")
	}
}
