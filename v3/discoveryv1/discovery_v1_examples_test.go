//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package discoveryv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/discoveryv1"
)

//
// This file provides an example of how to use the Discovery service.
//
// The following configuration properties are assumed to be defined:
// DISCOVERY_URL=<service base url>
// DISCOVERY_AUTH_TYPE=iam
// DISCOVERY_APIKEY=<IAM apikey>
// DISCOVERY_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`DiscoveryV1 Examples Tests`, func() {

	const externalConfigFile = "../discovery_v1.env"

	var (
		discoveryService *discoveryv1.DiscoveryV1
		config           map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(discoveryv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			discoveryServiceOptions := &discoveryv1.DiscoveryV1Options{
				Version: core.StringPtr("testString"),
			}

			discoveryService, err = discoveryv1.NewDiscoveryV1(discoveryServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(discoveryService).ToNot(BeNil())
		})
	})

	Describe(`DiscoveryV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEnvironment request example`, func() {
			fmt.Println("\nCreateEnvironment() result:")
			// begin-createEnvironment

			createEnvironmentOptions := discoveryService.NewCreateEnvironmentOptions(
				"testString",
			)

			environment, response, err := discoveryService.CreateEnvironment(createEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environment, "", "  ")
			fmt.Println(string(b))

			// end-createEnvironment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(environment).ToNot(BeNil())

		})
		It(`ListEnvironments request example`, func() {
			fmt.Println("\nListEnvironments() result:")
			// begin-listEnvironments

			listEnvironmentsOptions := discoveryService.NewListEnvironmentsOptions()

			listEnvironmentsResponse, response, err := discoveryService.ListEnvironments(listEnvironmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listEnvironmentsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listEnvironments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listEnvironmentsResponse).ToNot(BeNil())

		})
		It(`GetEnvironment request example`, func() {
			fmt.Println("\nGetEnvironment() result:")
			// begin-getEnvironment

			getEnvironmentOptions := discoveryService.NewGetEnvironmentOptions(
				"testString",
			)

			environment, response, err := discoveryService.GetEnvironment(getEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environment, "", "  ")
			fmt.Println(string(b))

			// end-getEnvironment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())

		})
		It(`UpdateEnvironment request example`, func() {
			fmt.Println("\nUpdateEnvironment() result:")
			// begin-updateEnvironment

			updateEnvironmentOptions := discoveryService.NewUpdateEnvironmentOptions(
				"testString",
			)

			environment, response, err := discoveryService.UpdateEnvironment(updateEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(environment, "", "  ")
			fmt.Println(string(b))

			// end-updateEnvironment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(environment).ToNot(BeNil())

		})
		It(`ListFields request example`, func() {
			fmt.Println("\nListFields() result:")
			// begin-listFields

			listFieldsOptions := discoveryService.NewListFieldsOptions(
				"testString",
				[]string{"testString"},
			)

			listCollectionFieldsResponse, response, err := discoveryService.ListFields(listFieldsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listCollectionFieldsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listFields

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listCollectionFieldsResponse).ToNot(BeNil())

		})
		It(`CreateConfiguration request example`, func() {
			fmt.Println("\nCreateConfiguration() result:")
			// begin-createConfiguration

			createConfigurationOptions := discoveryService.NewCreateConfigurationOptions(
				"testString",
				"testString",
			)

			configuration, response, err := discoveryService.CreateConfiguration(createConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configuration, "", "  ")
			fmt.Println(string(b))

			// end-createConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(configuration).ToNot(BeNil())

		})
		It(`ListConfigurations request example`, func() {
			fmt.Println("\nListConfigurations() result:")
			// begin-listConfigurations

			listConfigurationsOptions := discoveryService.NewListConfigurationsOptions(
				"testString",
			)

			listConfigurationsResponse, response, err := discoveryService.ListConfigurations(listConfigurationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listConfigurationsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listConfigurations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listConfigurationsResponse).ToNot(BeNil())

		})
		It(`GetConfiguration request example`, func() {
			fmt.Println("\nGetConfiguration() result:")
			// begin-getConfiguration

			getConfigurationOptions := discoveryService.NewGetConfigurationOptions(
				"testString",
				"testString",
			)

			configuration, response, err := discoveryService.GetConfiguration(getConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configuration, "", "  ")
			fmt.Println(string(b))

			// end-getConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configuration).ToNot(BeNil())

		})
		It(`UpdateConfiguration request example`, func() {
			fmt.Println("\nUpdateConfiguration() result:")
			// begin-updateConfiguration

			updateConfigurationOptions := discoveryService.NewUpdateConfigurationOptions(
				"testString",
				"testString",
				"testString",
			)

			configuration, response, err := discoveryService.UpdateConfiguration(updateConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configuration, "", "  ")
			fmt.Println(string(b))

			// end-updateConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configuration).ToNot(BeNil())

		})
		It(`CreateCollection request example`, func() {
			fmt.Println("\nCreateCollection() result:")
			// begin-createCollection

			createCollectionOptions := discoveryService.NewCreateCollectionOptions(
				"testString",
				"testString",
			)

			collection, response, err := discoveryService.CreateCollection(createCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-createCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collection).ToNot(BeNil())

		})
		It(`ListCollections request example`, func() {
			fmt.Println("\nListCollections() result:")
			// begin-listCollections

			listCollectionsOptions := discoveryService.NewListCollectionsOptions(
				"testString",
			)

			listCollectionsResponse, response, err := discoveryService.ListCollections(listCollectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listCollectionsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listCollections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listCollectionsResponse).ToNot(BeNil())

		})
		It(`GetCollection request example`, func() {
			fmt.Println("\nGetCollection() result:")
			// begin-getCollection

			getCollectionOptions := discoveryService.NewGetCollectionOptions(
				"testString",
				"testString",
			)

			collection, response, err := discoveryService.GetCollection(getCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-getCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collection).ToNot(BeNil())

		})
		It(`UpdateCollection request example`, func() {
			fmt.Println("\nUpdateCollection() result:")
			// begin-updateCollection

			updateCollectionOptions := discoveryService.NewUpdateCollectionOptions(
				"testString",
				"testString",
				"testString",
			)

			collection, response, err := discoveryService.UpdateCollection(updateCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collection, "", "  ")
			fmt.Println(string(b))

			// end-updateCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(collection).ToNot(BeNil())

		})
		It(`ListCollectionFields request example`, func() {
			fmt.Println("\nListCollectionFields() result:")
			// begin-listCollectionFields

			listCollectionFieldsOptions := discoveryService.NewListCollectionFieldsOptions(
				"testString",
				"testString",
			)

			listCollectionFieldsResponse, response, err := discoveryService.ListCollectionFields(listCollectionFieldsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listCollectionFieldsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listCollectionFields

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listCollectionFieldsResponse).ToNot(BeNil())

		})
		It(`ListExpansions request example`, func() {
			fmt.Println("\nListExpansions() result:")
			// begin-listExpansions

			listExpansionsOptions := discoveryService.NewListExpansionsOptions(
				"testString",
				"testString",
			)

			expansions, response, err := discoveryService.ListExpansions(listExpansionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(expansions, "", "  ")
			fmt.Println(string(b))

			// end-listExpansions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(expansions).ToNot(BeNil())

		})
		It(`CreateExpansions request example`, func() {
			fmt.Println("\nCreateExpansions() result:")
			// begin-createExpansions

			expansionModel := &discoveryv1.Expansion{
				ExpandedTerms: []string{"testString"},
			}

			createExpansionsOptions := discoveryService.NewCreateExpansionsOptions(
				"testString",
				"testString",
				[]discoveryv1.Expansion{*expansionModel},
			)

			expansions, response, err := discoveryService.CreateExpansions(createExpansionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(expansions, "", "  ")
			fmt.Println(string(b))

			// end-createExpansions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(expansions).ToNot(BeNil())

		})
		It(`GetTokenizationDictionaryStatus request example`, func() {
			fmt.Println("\nGetTokenizationDictionaryStatus() result:")
			// begin-getTokenizationDictionaryStatus

			getTokenizationDictionaryStatusOptions := discoveryService.NewGetTokenizationDictionaryStatusOptions(
				"testString",
				"testString",
			)

			tokenDictStatusResponse, response, err := discoveryService.GetTokenizationDictionaryStatus(getTokenizationDictionaryStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tokenDictStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-getTokenizationDictionaryStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tokenDictStatusResponse).ToNot(BeNil())

		})
		It(`CreateTokenizationDictionary request example`, func() {
			fmt.Println("\nCreateTokenizationDictionary() result:")
			// begin-createTokenizationDictionary

			createTokenizationDictionaryOptions := discoveryService.NewCreateTokenizationDictionaryOptions(
				"testString",
				"testString",
			)

			tokenDictStatusResponse, response, err := discoveryService.CreateTokenizationDictionary(createTokenizationDictionaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tokenDictStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-createTokenizationDictionary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(tokenDictStatusResponse).ToNot(BeNil())

		})
		It(`GetStopwordListStatus request example`, func() {
			fmt.Println("\nGetStopwordListStatus() result:")
			// begin-getStopwordListStatus

			getStopwordListStatusOptions := discoveryService.NewGetStopwordListStatusOptions(
				"testString",
				"testString",
			)

			tokenDictStatusResponse, response, err := discoveryService.GetStopwordListStatus(getStopwordListStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tokenDictStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-getStopwordListStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tokenDictStatusResponse).ToNot(BeNil())

		})
		It(`CreateStopwordList request example`, func() {
			fmt.Println("\nCreateStopwordList() result:")
			// begin-createStopwordList

			createStopwordListOptions := discoveryService.NewCreateStopwordListOptions(
				"testString",
				"testString",
				CreateMockReader("This is a mock file."),
				"testString",
			)

			tokenDictStatusResponse, response, err := discoveryService.CreateStopwordList(createStopwordListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tokenDictStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-createStopwordList

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tokenDictStatusResponse).ToNot(BeNil())

		})
		It(`AddDocument request example`, func() {
			fmt.Println("\nAddDocument() result:")
			// begin-addDocument

			addDocumentOptions := discoveryService.NewAddDocumentOptions(
				"testString",
				"testString",
			)

			documentAccepted, response, err := discoveryService.AddDocument(addDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentAccepted, "", "  ")
			fmt.Println(string(b))

			// end-addDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(documentAccepted).ToNot(BeNil())

		})
		It(`GetDocumentStatus request example`, func() {
			fmt.Println("\nGetDocumentStatus() result:")
			// begin-getDocumentStatus

			getDocumentStatusOptions := discoveryService.NewGetDocumentStatusOptions(
				"testString",
				"testString",
				"testString",
			)

			documentStatus, response, err := discoveryService.GetDocumentStatus(getDocumentStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentStatus, "", "  ")
			fmt.Println(string(b))

			// end-getDocumentStatus

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(documentStatus).ToNot(BeNil())

		})
		It(`UpdateDocument request example`, func() {
			fmt.Println("\nUpdateDocument() result:")
			// begin-updateDocument

			updateDocumentOptions := discoveryService.NewUpdateDocumentOptions(
				"testString",
				"testString",
				"testString",
			)

			documentAccepted, response, err := discoveryService.UpdateDocument(updateDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(documentAccepted, "", "  ")
			fmt.Println(string(b))

			// end-updateDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(documentAccepted).ToNot(BeNil())

		})
		It(`Query request example`, func() {
			fmt.Println("\nQuery() result:")
			// begin-query

			queryOptions := discoveryService.NewQueryOptions(
				"testString",
				"testString",
			)

			queryResponse, response, err := discoveryService.Query(queryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(queryResponse, "", "  ")
			fmt.Println(string(b))

			// end-query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(queryResponse).ToNot(BeNil())

		})
		It(`QueryNotices request example`, func() {
			fmt.Println("\nQueryNotices() result:")
			// begin-queryNotices

			queryNoticesOptions := discoveryService.NewQueryNoticesOptions(
				"testString",
				"testString",
			)

			queryNoticesResponse, response, err := discoveryService.QueryNotices(queryNoticesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(queryNoticesResponse, "", "  ")
			fmt.Println(string(b))

			// end-queryNotices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(queryNoticesResponse).ToNot(BeNil())

		})
		It(`FederatedQuery request example`, func() {
			fmt.Println("\nFederatedQuery() result:")
			// begin-federatedQuery

			federatedQueryOptions := discoveryService.NewFederatedQueryOptions(
				"testString",
				"testString",
			)

			queryResponse, response, err := discoveryService.FederatedQuery(federatedQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(queryResponse, "", "  ")
			fmt.Println(string(b))

			// end-federatedQuery

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(queryResponse).ToNot(BeNil())

		})
		It(`FederatedQueryNotices request example`, func() {
			fmt.Println("\nFederatedQueryNotices() result:")
			// begin-federatedQueryNotices

			federatedQueryNoticesOptions := discoveryService.NewFederatedQueryNoticesOptions(
				"testString",
				[]string{"testString"},
			)

			queryNoticesResponse, response, err := discoveryService.FederatedQueryNotices(federatedQueryNoticesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(queryNoticesResponse, "", "  ")
			fmt.Println(string(b))

			// end-federatedQueryNotices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(queryNoticesResponse).ToNot(BeNil())

		})
		It(`GetAutocompletion request example`, func() {
			fmt.Println("\nGetAutocompletion() result:")
			// begin-getAutocompletion

			getAutocompletionOptions := discoveryService.NewGetAutocompletionOptions(
				"testString",
				"testString",
				"testString",
			)

			completions, response, err := discoveryService.GetAutocompletion(getAutocompletionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(completions, "", "  ")
			fmt.Println(string(b))

			// end-getAutocompletion

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(completions).ToNot(BeNil())

		})
		It(`ListTrainingData request example`, func() {
			fmt.Println("\nListTrainingData() result:")
			// begin-listTrainingData

			listTrainingDataOptions := discoveryService.NewListTrainingDataOptions(
				"testString",
				"testString",
			)

			trainingDataSet, response, err := discoveryService.ListTrainingData(listTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingDataSet, "", "  ")
			fmt.Println(string(b))

			// end-listTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingDataSet).ToNot(BeNil())

		})
		It(`AddTrainingData request example`, func() {
			fmt.Println("\nAddTrainingData() result:")
			// begin-addTrainingData

			addTrainingDataOptions := discoveryService.NewAddTrainingDataOptions(
				"testString",
				"testString",
			)

			trainingQuery, response, err := discoveryService.AddTrainingData(addTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuery, "", "  ")
			fmt.Println(string(b))

			// end-addTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingQuery).ToNot(BeNil())

		})
		It(`GetTrainingData request example`, func() {
			fmt.Println("\nGetTrainingData() result:")
			// begin-getTrainingData

			getTrainingDataOptions := discoveryService.NewGetTrainingDataOptions(
				"testString",
				"testString",
				"testString",
			)

			trainingQuery, response, err := discoveryService.GetTrainingData(getTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuery, "", "  ")
			fmt.Println(string(b))

			// end-getTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingQuery).ToNot(BeNil())

		})
		It(`ListTrainingExamples request example`, func() {
			fmt.Println("\nListTrainingExamples() result:")
			// begin-listTrainingExamples

			listTrainingExamplesOptions := discoveryService.NewListTrainingExamplesOptions(
				"testString",
				"testString",
				"testString",
			)

			trainingExampleList, response, err := discoveryService.ListTrainingExamples(listTrainingExamplesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingExampleList, "", "  ")
			fmt.Println(string(b))

			// end-listTrainingExamples

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingExampleList).ToNot(BeNil())

		})
		It(`CreateTrainingExample request example`, func() {
			fmt.Println("\nCreateTrainingExample() result:")
			// begin-createTrainingExample

			createTrainingExampleOptions := discoveryService.NewCreateTrainingExampleOptions(
				"testString",
				"testString",
				"testString",
			)

			trainingExample, response, err := discoveryService.CreateTrainingExample(createTrainingExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingExample, "", "  ")
			fmt.Println(string(b))

			// end-createTrainingExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trainingExample).ToNot(BeNil())

		})
		It(`UpdateTrainingExample request example`, func() {
			fmt.Println("\nUpdateTrainingExample() result:")
			// begin-updateTrainingExample

			updateTrainingExampleOptions := discoveryService.NewUpdateTrainingExampleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			trainingExample, response, err := discoveryService.UpdateTrainingExample(updateTrainingExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingExample, "", "  ")
			fmt.Println(string(b))

			// end-updateTrainingExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingExample).ToNot(BeNil())

		})
		It(`GetTrainingExample request example`, func() {
			fmt.Println("\nGetTrainingExample() result:")
			// begin-getTrainingExample

			getTrainingExampleOptions := discoveryService.NewGetTrainingExampleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			trainingExample, response, err := discoveryService.GetTrainingExample(getTrainingExampleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingExample, "", "  ")
			fmt.Println(string(b))

			// end-getTrainingExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingExample).ToNot(BeNil())

		})
		It(`CreateEvent request example`, func() {
			fmt.Println("\nCreateEvent() result:")
			// begin-createEvent

			eventDataModel := &discoveryv1.EventData{
				EnvironmentID: core.StringPtr("testString"),
				SessionToken:  core.StringPtr("testString"),
				CollectionID:  core.StringPtr("testString"),
				DocumentID:    core.StringPtr("testString"),
			}

			createEventOptions := discoveryService.NewCreateEventOptions(
				"click",
				eventDataModel,
			)

			createEventResponse, response, err := discoveryService.CreateEvent(createEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createEventResponse, "", "  ")
			fmt.Println(string(b))

			// end-createEvent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createEventResponse).ToNot(BeNil())

		})
		It(`QueryLog request example`, func() {
			fmt.Println("\nQueryLog() result:")
			// begin-queryLog

			queryLogOptions := discoveryService.NewQueryLogOptions()

			logQueryResponse, response, err := discoveryService.QueryLog(queryLogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logQueryResponse, "", "  ")
			fmt.Println(string(b))

			// end-queryLog

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logQueryResponse).ToNot(BeNil())

		})
		It(`GetMetricsQuery request example`, func() {
			fmt.Println("\nGetMetricsQuery() result:")
			// begin-getMetricsQuery

			getMetricsQueryOptions := discoveryService.NewGetMetricsQueryOptions()

			metricResponse, response, err := discoveryService.GetMetricsQuery(getMetricsQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(metricResponse, "", "  ")
			fmt.Println(string(b))

			// end-getMetricsQuery

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metricResponse).ToNot(BeNil())

		})
		It(`GetMetricsQueryEvent request example`, func() {
			fmt.Println("\nGetMetricsQueryEvent() result:")
			// begin-getMetricsQueryEvent

			getMetricsQueryEventOptions := discoveryService.NewGetMetricsQueryEventOptions()

			metricResponse, response, err := discoveryService.GetMetricsQueryEvent(getMetricsQueryEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(metricResponse, "", "  ")
			fmt.Println(string(b))

			// end-getMetricsQueryEvent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metricResponse).ToNot(BeNil())

		})
		It(`GetMetricsQueryNoResults request example`, func() {
			fmt.Println("\nGetMetricsQueryNoResults() result:")
			// begin-getMetricsQueryNoResults

			getMetricsQueryNoResultsOptions := discoveryService.NewGetMetricsQueryNoResultsOptions()

			metricResponse, response, err := discoveryService.GetMetricsQueryNoResults(getMetricsQueryNoResultsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(metricResponse, "", "  ")
			fmt.Println(string(b))

			// end-getMetricsQueryNoResults

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metricResponse).ToNot(BeNil())

		})
		It(`GetMetricsEventRate request example`, func() {
			fmt.Println("\nGetMetricsEventRate() result:")
			// begin-getMetricsEventRate

			getMetricsEventRateOptions := discoveryService.NewGetMetricsEventRateOptions()

			metricResponse, response, err := discoveryService.GetMetricsEventRate(getMetricsEventRateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(metricResponse, "", "  ")
			fmt.Println(string(b))

			// end-getMetricsEventRate

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metricResponse).ToNot(BeNil())

		})
		It(`GetMetricsQueryTokenEvent request example`, func() {
			fmt.Println("\nGetMetricsQueryTokenEvent() result:")
			// begin-getMetricsQueryTokenEvent

			getMetricsQueryTokenEventOptions := discoveryService.NewGetMetricsQueryTokenEventOptions()

			metricTokenResponse, response, err := discoveryService.GetMetricsQueryTokenEvent(getMetricsQueryTokenEventOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(metricTokenResponse, "", "  ")
			fmt.Println(string(b))

			// end-getMetricsQueryTokenEvent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(metricTokenResponse).ToNot(BeNil())

		})
		It(`ListCredentials request example`, func() {
			fmt.Println("\nListCredentials() result:")
			// begin-listCredentials

			listCredentialsOptions := discoveryService.NewListCredentialsOptions(
				"testString",
			)

			credentialsList, response, err := discoveryService.ListCredentials(listCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentialsList, "", "  ")
			fmt.Println(string(b))

			// end-listCredentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentialsList).ToNot(BeNil())

		})
		It(`CreateCredentials request example`, func() {
			fmt.Println("\nCreateCredentials() result:")
			// begin-createCredentials

			createCredentialsOptions := discoveryService.NewCreateCredentialsOptions(
				"testString",
			)

			credentials, response, err := discoveryService.CreateCredentials(createCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentials, "", "  ")
			fmt.Println(string(b))

			// end-createCredentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentials).ToNot(BeNil())

		})
		It(`GetCredentials request example`, func() {
			fmt.Println("\nGetCredentials() result:")
			// begin-getCredentials

			getCredentialsOptions := discoveryService.NewGetCredentialsOptions(
				"testString",
				"testString",
			)

			credentials, response, err := discoveryService.GetCredentials(getCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentials, "", "  ")
			fmt.Println(string(b))

			// end-getCredentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentials).ToNot(BeNil())

		})
		It(`UpdateCredentials request example`, func() {
			fmt.Println("\nUpdateCredentials() result:")
			// begin-updateCredentials

			updateCredentialsOptions := discoveryService.NewUpdateCredentialsOptions(
				"testString",
				"testString",
			)

			credentials, response, err := discoveryService.UpdateCredentials(updateCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(credentials, "", "  ")
			fmt.Println(string(b))

			// end-updateCredentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(credentials).ToNot(BeNil())

		})
		It(`ListGateways request example`, func() {
			fmt.Println("\nListGateways() result:")
			// begin-listGateways

			listGatewaysOptions := discoveryService.NewListGatewaysOptions(
				"testString",
			)

			gatewayList, response, err := discoveryService.ListGateways(listGatewaysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gatewayList, "", "  ")
			fmt.Println(string(b))

			// end-listGateways

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gatewayList).ToNot(BeNil())

		})
		It(`CreateGateway request example`, func() {
			fmt.Println("\nCreateGateway() result:")
			// begin-createGateway

			createGatewayOptions := discoveryService.NewCreateGatewayOptions(
				"testString",
			)

			gateway, response, err := discoveryService.CreateGateway(createGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gateway, "", "  ")
			fmt.Println(string(b))

			// end-createGateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gateway).ToNot(BeNil())

		})
		It(`GetGateway request example`, func() {
			fmt.Println("\nGetGateway() result:")
			// begin-getGateway

			getGatewayOptions := discoveryService.NewGetGatewayOptions(
				"testString",
				"testString",
			)

			gateway, response, err := discoveryService.GetGateway(getGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gateway, "", "  ")
			fmt.Println(string(b))

			// end-getGateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gateway).ToNot(BeNil())

		})
		It(`DeleteUserData request example`, func() {
			// begin-deleteUserData

			deleteUserDataOptions := discoveryService.NewDeleteUserDataOptions(
				"testString",
			)

			response, err := discoveryService.DeleteUserData(deleteUserDataOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteUserData(): %d\n", response.StatusCode)
			}

			// end-deleteUserData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteTrainingExample request example`, func() {
			// begin-deleteTrainingExample

			deleteTrainingExampleOptions := discoveryService.NewDeleteTrainingExampleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteTrainingExample(deleteTrainingExampleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTrainingExample(): %d\n", response.StatusCode)
			}

			// end-deleteTrainingExample

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTrainingData request example`, func() {
			// begin-deleteTrainingData

			deleteTrainingDataOptions := discoveryService.NewDeleteTrainingDataOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteTrainingData(deleteTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTrainingData(): %d\n", response.StatusCode)
			}

			// end-deleteTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTokenizationDictionary request example`, func() {
			// begin-deleteTokenizationDictionary

			deleteTokenizationDictionaryOptions := discoveryService.NewDeleteTokenizationDictionaryOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteTokenizationDictionary(deleteTokenizationDictionaryOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteTokenizationDictionary(): %d\n", response.StatusCode)
			}

			// end-deleteTokenizationDictionary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteStopwordList request example`, func() {
			// begin-deleteStopwordList

			deleteStopwordListOptions := discoveryService.NewDeleteStopwordListOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteStopwordList(deleteStopwordListOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DeleteStopwordList(): %d\n", response.StatusCode)
			}

			// end-deleteStopwordList

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteGateway request example`, func() {
			fmt.Println("\nDeleteGateway() result:")
			// begin-deleteGateway

			deleteGatewayOptions := discoveryService.NewDeleteGatewayOptions(
				"testString",
				"testString",
			)

			gatewayDelete, response, err := discoveryService.DeleteGateway(deleteGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(gatewayDelete, "", "  ")
			fmt.Println(string(b))

			// end-deleteGateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(gatewayDelete).ToNot(BeNil())

		})
		It(`DeleteExpansions request example`, func() {
			// begin-deleteExpansions

			deleteExpansionsOptions := discoveryService.NewDeleteExpansionsOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteExpansions(deleteExpansionsOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteExpansions(): %d\n", response.StatusCode)
			}

			// end-deleteExpansions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteEnvironment request example`, func() {
			fmt.Println("\nDeleteEnvironment() result:")
			// begin-deleteEnvironment

			deleteEnvironmentOptions := discoveryService.NewDeleteEnvironmentOptions(
				"testString",
			)

			deleteEnvironmentResponse, response, err := discoveryService.DeleteEnvironment(deleteEnvironmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteEnvironmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteEnvironment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteEnvironmentResponse).ToNot(BeNil())

		})
		It(`DeleteDocument request example`, func() {
			fmt.Println("\nDeleteDocument() result:")
			// begin-deleteDocument

			deleteDocumentOptions := discoveryService.NewDeleteDocumentOptions(
				"testString",
				"testString",
				"testString",
			)

			deleteDocumentResponse, response, err := discoveryService.DeleteDocument(deleteDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteDocumentResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteDocumentResponse).ToNot(BeNil())

		})
		It(`DeleteCredentials request example`, func() {
			fmt.Println("\nDeleteCredentials() result:")
			// begin-deleteCredentials

			deleteCredentialsOptions := discoveryService.NewDeleteCredentialsOptions(
				"testString",
				"testString",
			)

			deleteCredentials, response, err := discoveryService.DeleteCredentials(deleteCredentialsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteCredentials, "", "  ")
			fmt.Println(string(b))

			// end-deleteCredentials

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteCredentials).ToNot(BeNil())

		})
		It(`DeleteConfiguration request example`, func() {
			fmt.Println("\nDeleteConfiguration() result:")
			// begin-deleteConfiguration

			deleteConfigurationOptions := discoveryService.NewDeleteConfigurationOptions(
				"testString",
				"testString",
			)

			deleteConfigurationResponse, response, err := discoveryService.DeleteConfiguration(deleteConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteConfigurationResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteConfiguration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteConfigurationResponse).ToNot(BeNil())

		})
		It(`DeleteCollection request example`, func() {
			fmt.Println("\nDeleteCollection() result:")
			// begin-deleteCollection

			deleteCollectionOptions := discoveryService.NewDeleteCollectionOptions(
				"testString",
				"testString",
			)

			deleteCollectionResponse, response, err := discoveryService.DeleteCollection(deleteCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteCollectionResponse, "", "  ")
			fmt.Println(string(b))

			// end-deleteCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(deleteCollectionResponse).ToNot(BeNil())

		})
		It(`DeleteAllTrainingData request example`, func() {
			// begin-deleteAllTrainingData

			deleteAllTrainingDataOptions := discoveryService.NewDeleteAllTrainingDataOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteAllTrainingData(deleteAllTrainingDataOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAllTrainingData(): %d\n", response.StatusCode)
			}

			// end-deleteAllTrainingData

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
