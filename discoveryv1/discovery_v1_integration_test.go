// +build integration

package discoveryv1_test

/**
 * Copyright 2018 IBM All Rights Reserved.
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
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/watson-developer-cloud/go-sdk/core"
	"github.com/watson-developer-cloud/go-sdk/discoveryv1"
	"os"
	"testing"
)

var service *discoveryv1.DiscoveryV1
var serviceErr error
var environmentID *string
var configurationID *string
var collectionID *string

func TestInitialization(t *testing.T) {
	err := godotenv.Load("../.env")
	require.Nil(t, err)

	service, serviceErr = discoveryv1.
		NewDiscoveryV1(&discoveryv1.DiscoveryV1Options{
			URL:       os.Getenv("DISCOVERY_URL"),
			Version:   "2018-03-05",
			IAMApiKey: os.Getenv("DISCOVERY_IAMAPIKEY"),
		})
	environmentID = core.StringPtr(os.Getenv("DISCOVERY_ENVIRONMENT_ID"))
	require.Nil(t, serviceErr)
}

func TestEnvironment(t *testing.T) {
	// List environment
	response, responseErr := service.ListEnvironments(
		&discoveryv1.ListEnvironmentsOptions{},
	)
	assert.Nil(t, responseErr)

	listEnvironments := service.GetListEnvironmentsResult(response)
	assert.NotNil(t, listEnvironments)

	t.Skip("Skip rest of environment API tests")
	// Create environment
	response, responseErr = service.CreateEnvironment(
		&discoveryv1.CreateEnvironmentOptions{
			Name:        core.StringPtr("test environment for SDKS DO NOT DELETE"),
			Description: core.StringPtr("My environment"),
		},
	)

	assert.Nil(t, responseErr)

	createEnvironment := service.GetCreateEnvironmentResult(response)
	assert.NotNil(t, createEnvironment)

	var newsEnvironmentId *string
	for _, environment := range listEnvironments.Environments {
		if *environment.Name == "Watson Discovery News Environment" {
			newsEnvironmentId = environment.EnvironmentID
		}
	}
	if newsEnvironmentId != nil {
		response, responseErr = service.ListCollections(
			&discoveryv1.ListCollectionsOptions{
				EnvironmentID: newsEnvironmentId,
			},
		)
		assert.Nil(t, responseErr)
	}

	// Get environment
	response, responseErr = service.GetEnvironment(
		&discoveryv1.GetEnvironmentOptions{
			EnvironmentID: createEnvironment.EnvironmentID,
		},
	)
	assert.Nil(t, responseErr)

	getEnvironment := service.GetGetEnvironmentResult(response)
	assert.NotNil(t, getEnvironment)

	// Update environment
	response, responseErr = service.UpdateEnvironment(
		&discoveryv1.UpdateEnvironmentOptions{
			EnvironmentID: createEnvironment.EnvironmentID,
			Name:          core.StringPtr("Updated name"),
			Description:   core.StringPtr("Updated description"),
		},
	)
	assert.Nil(t, responseErr)

	updateEnvironment := service.GetUpdateEnvironmentResult(response)
	assert.NotNil(t, updateEnvironment)
}

func TestConfiguration(t *testing.T) {
	// Create Configuraion
	response, responseErr := service.CreateConfiguration(
		&discoveryv1.CreateConfigurationOptions{
			EnvironmentID: environmentID,
			Name:          core.StringPtr("Test configuration for GO"),
		},
	)
	assert.Nil(t, responseErr)

	createConfiguration := service.GetCreateConfigurationResult(response)
	assert.NotNil(t, createConfiguration)

	// List Configuraion
	response, responseErr = service.ListConfigurations(
		&discoveryv1.ListConfigurationsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)

	listConfiguration := service.GetListConfigurationsResult(response)
	assert.NotNil(t, listConfiguration)

	// Test configuration
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/simple.html")
	assert.Nil(t, fileErr)

	response, responseErr = service.TestConfigurationInEnvironment(
		&discoveryv1.TestConfigurationInEnvironmentOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: createConfiguration.ConfigurationID,
			File:            file,
		},
	)
	assert.Nil(t, responseErr)

	testConfiguration := service.GetTestConfigurationInEnvironmentResult(response)
	assert.NotNil(t, testConfiguration)

	// Get configuration
	response, responseErr = service.GetConfiguration(
		&discoveryv1.GetConfigurationOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: createConfiguration.ConfigurationID,
		},
	)
	assert.Nil(t, responseErr)

	getConfiguration := service.GetGetConfigurationResult(response)
	assert.NotNil(t, getConfiguration)

	// Update configuration
	response, responseErr = service.UpdateConfiguration(
		&discoveryv1.UpdateConfigurationOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: createConfiguration.ConfigurationID,
			Name:            core.StringPtr("Test configuration for GO with name update"),
		},
	)
	assert.Nil(t, responseErr)

	updateConfiguration := service.GetUpdateConfigurationResult(response)
	assert.NotNil(t, updateConfiguration)

	configurationID = createConfiguration.ConfigurationID
}

func TestCollection(t *testing.T) {
	// Create collection
	response, responseErr := service.CreateCollection(
		&discoveryv1.CreateCollectionOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: configurationID,
			Name:            core.StringPtr("Name for GO collection"),
		},
	)
	assert.Nil(t, responseErr)

	createCollection := service.GetCreateCollectionResult(response)
	assert.NotNil(t, createCollection)

	// List collection
	response, responseErr = service.ListCollections(
		&discoveryv1.ListCollectionsOptions{
			EnvironmentID: environmentID,
		},
	)
	assert.Nil(t, responseErr)

	listCollection := service.GetListCollectionsResult(response)
	assert.NotNil(t, listCollection)

	// Get collection
	response, responseErr = service.GetCollection(
		&discoveryv1.GetCollectionOptions{
			EnvironmentID: environmentID,
			CollectionID:  createCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)

	getCollection := service.GetGetCollectionResult(response)
	assert.NotNil(t, getCollection)

	// Update collection
	response, responseErr = service.UpdateCollection(
		&discoveryv1.UpdateCollectionOptions{
			EnvironmentID:   environmentID,
			CollectionID:    createCollection.CollectionID,
			ConfigurationID: configurationID,
			Name:            core.StringPtr("Update collection in GO"),
			Description:     core.StringPtr("Update description in GO"),
		},
	)
	assert.Nil(t, responseErr)

	updateCollection := service.GetUpdateCollectionResult(response)
	assert.NotNil(t, updateCollection)

	// List collection fields
	response, responseErr = service.ListCollectionFields(
		&discoveryv1.ListCollectionFieldsOptions{
			EnvironmentID: environmentID,
			CollectionID:  createCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)

	listCollectionFields := service.GetListCollectionFieldsResult(response)
	assert.NotNil(t, listCollectionFields)

	// List fields
	response, responseErr = service.ListFields(
		&discoveryv1.ListFieldsOptions{
			EnvironmentID: environmentID,
			CollectionIds: []string{*createCollection.CollectionID},
		},
	)
	assert.Nil(t, responseErr)

	listFields := service.GetListFieldsResult(response)
	assert.NotNil(t, listFields)

	collectionID = createCollection.CollectionID
}

func TestDocument(t *testing.T) {
	// Add document
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileErr)

	response, responseErr := service.AddDocument(
		&discoveryv1.AddDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			File:          file,
		},
	)
	assert.Nil(t, responseErr)

	addDocument := service.GetAddDocumentResult(response)
	assert.NotNil(t, addDocument)

	// Get document
	response, responseErr = service.GetDocumentStatus(
		&discoveryv1.GetDocumentStatusOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)

	getDocument := service.GetGetDocumentStatusResult(response)
	assert.NotNil(t, getDocument)

	// Update document
	fileInfo, fileInfoErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileInfoErr)

	response, responseErr = service.UpdateDocument(
		&discoveryv1.UpdateDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
			File:          fileInfo,
		},
	)
	assert.Nil(t, responseErr)

	updateDocument := service.GetUpdateDocumentResult(response)
	assert.NotNil(t, updateDocument)

	// Delete document
	response, responseErr = service.DeleteDocument(
		&discoveryv1.DeleteDocumentOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			DocumentID:    addDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestQuery(t *testing.T) {
	// Query
	response, responseErr := service.Query(
		&discoveryv1.QueryOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
			Filter:        core.StringPtr("extracted_metadata.sha1::9181d244*"),
			ReturnFields:  core.StringPtr("extracted_metadata.sha1"),
		},
	)
	assert.Nil(t, responseErr)

	query := service.GetQueryResult(response)
	assert.NotNil(t, query)
}

func TestTokenizationDictionary(t *testing.T) {
	// Create collection in Japanese as create tokenization dictionary is only supported in JA
	response, responseErr := service.CreateCollection(
		&discoveryv1.CreateCollectionOptions{
			EnvironmentID: environmentID,
			Name:          core.StringPtr("Test Tokenization Dictionary For Golang"),
			Language:      core.StringPtr(discoveryv1.CreateCollectionOptions_Language_Ja),
		},
	)
	assert.Nil(t, responseErr)

	testCollection := service.GetCreateCollectionResult(response)
	assert.NotNil(t, testCollection)

	// create tokenization dictionary
	response, responseErr = service.CreateTokenizationDictionary(
		&discoveryv1.CreateTokenizationDictionaryOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
			TokenizationRules: []discoveryv1.TokenDictRule{
				discoveryv1.TokenDictRule{
					Text:         core.StringPtr("token"),
					Tokens:       []string{"token 1", "token 2"},
					Readings:     []string{"reading 1", "reading 2"},
					PartOfSpeech: core.StringPtr("noun"),
				},
			},
		},
	)

	createTokenizationDictionary := service.GetCreateTokenizationDictionaryResult(response)
	assert.NotNil(t, createTokenizationDictionary)

	// get tokenization dictionary status
	response, responseErr = service.GetTokenizationDictionaryStatus(
		&discoveryv1.GetTokenizationDictionaryStatusOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
		},
	)

	getTokenizationDictionaryStatus := service.GetGetTokenizationDictionaryStatusResult(response)
	assert.NotNil(t, getTokenizationDictionaryStatus)

	// delete tokenization dictionary
	response, responseErr = service.DeleteTokenizationDictionary(
		&discoveryv1.DeleteTokenizationDictionaryOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
		},
	)

	// Delete collection
	response, responseErr = service.DeleteCollection(
		&discoveryv1.DeleteCollectionOptions{
			EnvironmentID: environmentID,
			CollectionID:  testCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
}

func TestDeleteOperations(t *testing.T) {
	// Delete collection
	response, responseErr := service.DeleteCollection(
		&discoveryv1.DeleteCollectionOptions{
			EnvironmentID: environmentID,
			CollectionID:  collectionID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)

	// Delete configuration
	response, responseErr = service.DeleteConfiguration(
		&discoveryv1.DeleteConfigurationOptions{
			EnvironmentID:   environmentID,
			ConfigurationID: configurationID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, response)
}
