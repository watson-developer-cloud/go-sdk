// +build integration

package discoveryv1_test

/**
 * (C) Copyright IBM Corp. 2018, 2020.
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

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *discoveryv1.DiscoveryV1
var environmentID *string
var configurationID *string
var collectionID *string

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load(configFile)
	if err != nil {
		t.Skip(skipMessage)
	}

	s := os.Getenv("DISCOVERY_ENVIRONMENT_ID")
	assert.NotEmpty(t, s)
	if s != "" {
		configLoaded = true
		environmentID = &s
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = discoveryv1.NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
		Version: core.StringPtr("2020-11-05"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, service)

	if err == nil {
		customHeaders := http.Header{}
		customHeaders.Add("X-Watson-Learning-Opt-Out", "1")
		customHeaders.Add("X-Watson-Test", "1")
		service.Service.SetDefaultHeaders(customHeaders)
	}
}

func TestEnvironment(t *testing.T) {
	shouldSkipTest(t)

	// List environment
	listEnvironments, _, responseErr := service.ListEnvironments(
		&discoveryv1.ListEnvironmentsOptions{},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listEnvironments)

	t.Skip("Skip rest of environment API tests")

	// Create environment
	createEnvironment, _, responseErr := service.CreateEnvironment(
		&discoveryv1.CreateEnvironmentOptions{
			Name:        core.StringPtr("test environment for SDKS DO NOT DELETE"),
			Description: core.StringPtr("My environment"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createEnvironment)

	var newsEnvironmentId *string
	for _, environment := range listEnvironments.Environments {
		if *environment.Name == "Watson Discovery News Environment" {
			newsEnvironmentId = environment.EnvironmentID
		}
	}
	if newsEnvironmentId != nil {
		_, _, responseErr = service.ListCollections(
			&discoveryv1.ListCollectionsOptions{
				EnvironmentID: newsEnvironmentId,
			},
		)
		assert.Nil(t, responseErr)
	}

	// Get environment
	getEnvironment, _, responseErr := service.GetEnvironment(
		&discoveryv1.GetEnvironmentOptions{
			EnvironmentID: createEnvironment.EnvironmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getEnvironment)

	// Update environment
	updateEnvironment, _, responseErr := service.UpdateEnvironment(
		&discoveryv1.UpdateEnvironmentOptions{
			EnvironmentID: createEnvironment.EnvironmentID,
			Name:          core.StringPtr("Updated name"),
			Description:   core.StringPtr("Updated description"),
		},
	)
	assert.Nil(t, responseErr)

	assert.NotNil(t, updateEnvironment)
}

func TestConfiguration(t *testing.T) {
	shouldSkipTest(t)

	// Create Configuraion
	createConfiguration, _, responseErr := service.CreateConfiguration(
		&discoveryv1.CreateConfigurationOptions{
			EnvironmentID: environmentID,
			Name:          core.StringPtr("Test configuration for GO"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createConfiguration)

	// List Configuraion
	listConfiguration, _, responseErr := service.ListConfigurations(
		&discoveryv1.ListConfigurationsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listConfiguration)

	// Get configuration
	getConfiguration, _, responseErr := service.GetConfiguration(
		&discoveryv1.GetConfigurationOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: createConfiguration.ConfigurationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getConfiguration)

	// Update configuration
	updateConfiguration, _, responseErr := service.UpdateConfiguration(
		&discoveryv1.UpdateConfigurationOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: createConfiguration.ConfigurationID,
			Name:            core.StringPtr("Test configuration for GO with name update"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateConfiguration)

	configurationID = createConfiguration.ConfigurationID
}

func TestCollection(t *testing.T) {
	shouldSkipTest(t)

	// Create collection
	createCollection, _, responseErr := service.CreateCollection(
		&discoveryv1.CreateCollectionOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: configurationID,
			Name:            core.StringPtr("Name for GO collection"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createCollection)

	// List collection
	listCollection, _, responseErr := service.ListCollections(
		&discoveryv1.ListCollectionsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCollection)

	// Get collection
	getCollection, _, responseErr := service.GetCollection(
		&discoveryv1.GetCollectionOptions{
			EnvironmentID: environmentID,
			CollectionID:  createCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getCollection)

	// Update collection
	updateCollection, _, responseErr := service.UpdateCollection(
		&discoveryv1.UpdateCollectionOptions{
			EnvironmentID:   environmentID,
			CollectionID:    createCollection.CollectionID,
			ConfigurationID: configurationID,
			Name:            core.StringPtr("Update collection in GO"),
			Description:     core.StringPtr("Update description in GO"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateCollection)

	// List collection fields
	listCollectionFields, _, responseErr := service.ListCollectionFields(
		&discoveryv1.ListCollectionFieldsOptions{
			EnvironmentID: environmentID,
			CollectionID:  createCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCollectionFields)

	// List fields
	listFields, _, responseErr := service.ListFields(
		&discoveryv1.ListFieldsOptions{
			EnvironmentID: environmentID,
			CollectionIds: []string{*createCollection.CollectionID},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listFields)

	collectionID = createCollection.CollectionID
}

func TestDocument(t *testing.T) {
	shouldSkipTest(t)

	// Add document
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileErr)

	addDocument, _, responseErr := service.AddDocument(
		&discoveryv1.AddDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			File:          file,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, addDocument)

	// Get document
	getDocument, _, responseErr := service.GetDocumentStatus(
		&discoveryv1.GetDocumentStatusOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getDocument)

	// Update document
	fileInfo, fileInfoErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileInfoErr)

	updateDocument, _, responseErr := service.UpdateDocument(
		&discoveryv1.UpdateDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
			File:          fileInfo,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateDocument)

	// Delete document
	_, _, responseErr = service.DeleteDocument(
		&discoveryv1.DeleteDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestQuery(t *testing.T) {
	shouldSkipTest(t)

	// Query
	query, _, responseErr := service.Query(
		&discoveryv1.QueryOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			Filter:        core.StringPtr("extracted_metadata.sha1::9181d244*"),
			Return:        core.StringPtr("extracted_metadata.sha1"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, query)
}

func TestQueryWithTimesliceAggregation(t *testing.T) {
	shouldSkipTest(t)

	// Query
	query, _, responseErr := service.Query(
		&discoveryv1.QueryOptions{
			EnvironmentID: core.StringPtr("system"),
			CollectionID:  core.StringPtr("news-en"),
			Query:         core.StringPtr("enriched_text.concepts.text:\"Cloud computing\""),
			Aggregation:   core.StringPtr("timeslice(publication_date,1day,anomaly:false)"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, query)
}

func TestTokenizationDictionary(t *testing.T) {
	shouldSkipTest(t)

	t.Skip("Disable temporarily")
	// Create collection in Japanese as create tokenization dictionary is only supported in JA
	testCollection, _, responseErr := service.CreateCollection(
		&discoveryv1.CreateCollectionOptions{
			EnvironmentID: environmentID,
			Name:          core.StringPtr("Test Tokenization Dictionary For Golang"),
			Language:      core.StringPtr(discoveryv1.CreateCollectionOptions_Language_Ja),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, testCollection)

	// create tokenization dictionary
	createTokenizationDictionary, _, responseErr := service.CreateTokenizationDictionary(
		&discoveryv1.CreateTokenizationDictionaryOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
			TokenizationRules: []discoveryv1.TokenDictRule{
				{
					Text:         core.StringPtr("token"),
					Tokens:       []string{"token 1", "token 2"},
					Readings:     []string{"reading 1", "reading 2"},
					PartOfSpeech: core.StringPtr("noun"),
				},
			},
		},
	)
	assert.NotNil(t, createTokenizationDictionary)

	// get tokenization dictionary status
	getTokenizationDictionaryStatus, _, responseErr := service.GetTokenizationDictionaryStatus(
		&discoveryv1.GetTokenizationDictionaryStatusOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
		},
	)
	assert.NotNil(t, getTokenizationDictionaryStatus)

	// delete tokenization dictionary
	_, _ = service.DeleteTokenizationDictionary(
		&discoveryv1.DeleteTokenizationDictionaryOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
		},
	)
}

func TestStopwordOperations(t *testing.T) {
	shouldSkipTest(t)

	t.Skip("Disable temporarily")
	// Create stopword list
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/stopwords.txt")
	assert.Nil(t, fileErr)

	stopwordsListStatus, _, responseErr := service.CreateStopwordList(
		&discoveryv1.CreateStopwordListOptions{
			EnvironmentID:    environmentID,
			CollectionID:     collectionID,
			StopwordFile:     file,
			StopwordFilename: core.StringPtr("stopwords.txt"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, stopwordsListStatus)

	// Delete stopword list
	_, responseErr = service.DeleteStopwordList(
		&discoveryv1.DeleteStopwordListOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestGatewayConfiguration(t *testing.T) {
	shouldSkipTest(t)

	// Create gateway
	createGateway, _, responseErr := service.CreateGateway(
		&discoveryv1.CreateGatewayOptions{
			EnvironmentID: environmentID,
			Name:          core.StringPtr("test-gateway-configuration-go"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createGateway)

	// Get gateway
	getGateway, _, responseErr := service.GetGateway(
		&discoveryv1.GetGatewayOptions{
			EnvironmentID: environmentID,
			GatewayID:     createGateway.GatewayID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getGateway)

	// List gateways
	listGateways, _, responseErr := service.ListGateways(
		&discoveryv1.ListGatewaysOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listGateways)

	// Delete gateway
	deleteGateway, _, responseErr := service.DeleteGateway(
		&discoveryv1.DeleteGatewayOptions{
			EnvironmentID: environmentID,
			GatewayID:     getGateway.GatewayID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, deleteGateway)
}

func TestDeleteOperations(t *testing.T) {
	shouldSkipTest(t)

	// Delete all collections
	listCollection, _, responseErr := service.ListCollections(
		&discoveryv1.ListCollectionsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCollection)

	for _, collection := range listCollection.Collections {
		fmt.Println("Deleting collection " + *collection.Name)
		// delete the collection
		service.DeleteCollection(
			&discoveryv1.DeleteCollectionOptions{
				EnvironmentID: environmentID,
				CollectionID:  collection.CollectionID,
			},
		)
	}

	// Delete all configurations
	listConfigurations, _, responseErr := service.ListConfigurations(
		&discoveryv1.ListConfigurationsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listConfigurations)

	for _, configuration := range listConfigurations.Configurations {
		// delete the configuration
		if *configuration.Name != "Default Configuration" {
			fmt.Println("Deleting configuration " + *configuration.Name)
			service.DeleteConfiguration(
				&discoveryv1.DeleteConfigurationOptions{
					EnvironmentID:   environmentID,
					ConfigurationID: configuration.ConfigurationID,
				},
			)
		}
	}
}
