// +build integration

package discoveryv2_test

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
	"net/http"
	"os"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/watson-developer-cloud/go-sdk/v2/discoveryv2"
)

const skipMessage = "External configuration could not be loaded, skipping..."

var configLoaded bool
var configFile = "../.env"

var service *discoveryv2.DiscoveryV2
var projectID *string
var collectionID *string
var serviceURL *string
var APIKey *string

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

	projectEnv := os.Getenv("DISCOVERY_V2_PROJECT_ID")
	if projectEnv != "" {
		projectID = &projectEnv
	}

	collectionEnv := os.Getenv("DISCOVERY_V2_COLLECTION_ID")
	if collectionEnv != "" {
		collectionID = &collectionEnv
	}

	serviceURLEnv := os.Getenv("DISCOVERY_V2_URL")
	if serviceURLEnv != "" {
		serviceURL = &serviceURLEnv
	}

	APIKeyEnv := os.Getenv("DISCOVERY_V2_APIKEY")
	if APIKeyEnv != "" {
		APIKey = &APIKeyEnv
	}

	if projectID != nil && collectionID != nil && serviceURL != nil && APIKey != nil {
		configLoaded = true
	}
}

func TestConstructService(t *testing.T) {
	shouldSkipTest(t)

	var err error

	service, err = discoveryv2.NewDiscoveryV2(&discoveryv2.DiscoveryV2Options{
		Version: core.StringPtr("2020-08-23"),
		URL:     *serviceURL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: *APIKey,
		},
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

func TestCollection(t *testing.T) {
	shouldSkipTest(t)

	// Create collection
	createCollection, _, responseErr := service.CreateCollection(
		&discoveryv2.CreateCollectionOptions{
			ProjectID: projectID,
			Name:      core.StringPtr("Name for GO collection"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, createCollection)

	// List collection
	listCollection, _, responseErr := service.ListCollections(
		&discoveryv2.ListCollectionsOptions{
			ProjectID: projectID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listCollection)

	// Get collection
	getCollection, _, responseErr := service.GetCollection(
		&discoveryv2.GetCollectionOptions{
			ProjectID:    projectID,
			CollectionID: createCollection.CollectionID,
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, getCollection)

	// Update collection
	updateCollection, _, responseErr := service.UpdateCollection(
		&discoveryv2.UpdateCollectionOptions{
			ProjectID:    projectID,
			CollectionID: createCollection.CollectionID,
			Name:         core.StringPtr("Update collection in GO"),
			Description:  core.StringPtr("Update description in GO"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateCollection)

	// List fields
	listFields, _, responseErr := service.ListFields(
		&discoveryv2.ListFieldsOptions{
			ProjectID:     projectID,
			CollectionIds: []string{*createCollection.CollectionID},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, listFields)

	_, deleteErr := service.DeleteCollection(
		&discoveryv2.DeleteCollectionOptions{
			ProjectID:    projectID,
			CollectionID: createCollection.CollectionID,
		},
	)

	assert.Nil(t, deleteErr)
}

func TestDocument(t *testing.T) {
	shouldSkipTest(t)

	// Add document
	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileErr)

	addDocument, _, responseErr := service.AddDocument(
		&discoveryv2.AddDocumentOptions{
			ProjectID:       projectID,
			CollectionID:    collectionID,
			File:            file,
			Filename:        core.StringPtr("test_file"),
			FileContentType: core.StringPtr("application/html"),
		},
	)

	assert.Nil(t, responseErr)
	assert.NotNil(t, addDocument)

	// Update document
	fileInfo, fileInfoErr := os.Open(pwd + "/../resources/example.html")
	assert.Nil(t, fileInfoErr)

	updateDocument, _, responseErr := service.UpdateDocument(
		&discoveryv2.UpdateDocumentOptions{
			ProjectID:       projectID,
			CollectionID:    collectionID,
			DocumentID:      addDocument.DocumentID,
			File:            fileInfo,
			Filename:        core.StringPtr("test_file"),
			FileContentType: core.StringPtr("application/html"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, updateDocument)

	// Delete document
	_, _, responseErr = service.DeleteDocument(
		&discoveryv2.DeleteDocumentOptions{
			ProjectID:    projectID,
			CollectionID: collectionID,
			DocumentID:   addDocument.DocumentID,
		},
	)
	assert.Nil(t, responseErr)
}

func TestQuery(t *testing.T) {
	shouldSkipTest(t)

	// Query
	query, _, responseErr := service.Query(
		&discoveryv2.QueryOptions{
			ProjectID:     projectID,
			CollectionIds: []string{*collectionID},
			Filter:        core.StringPtr("extracted_metadata.sha1::9181d244*"),
			Return:        []string{"extracted_metadata.sha1"},
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, query)
}

func TestQueryWithTimesliceAggregation(t *testing.T) {
	shouldSkipTest(t)

	// Query
	query, _, responseErr := service.Query(
		&discoveryv2.QueryOptions{
			ProjectID:     projectID,
			CollectionIds: []string{*collectionID},
			Query:         core.StringPtr("enriched_text.concepts.text:\"Cloud computing\""),
			Aggregation:   core.StringPtr("timeslice(publication_date,1day,anomaly:false)"),
		},
	)
	assert.Nil(t, responseErr)
	assert.NotNil(t, query)
}

func TestEnrichments(t *testing.T) {
	shouldSkipTest(t)

	pwd, _ := os.Getwd()
	file, fileErr := os.Open(pwd + "/../resources/TestEnrichments.csv")
	assert.Nil(t, fileErr)

	enrichment, _, createErr := service.CreateEnrichment(
		&discoveryv2.CreateEnrichmentOptions{
			ProjectID: projectID,
			File:      file,
			Enrichment: &discoveryv2.CreateEnrichment{
				Name:        core.StringPtr("test enrichment"),
				Description: core.StringPtr("enrichment used for go"),
				Type:        core.StringPtr("dictionary"),
				Options: &discoveryv2.EnrichmentOptions{
					Languages:  []string{"en"},
					EntityType: core.StringPtr("keyword"),
				},
			},
		},
	)

	assert.Nil(t, createErr)
	assert.NotNil(t, enrichment)

	enrichmentList, _, listErr := service.ListEnrichments(
		&discoveryv2.ListEnrichmentsOptions{
			ProjectID: projectID,
		},
	)

	assert.Nil(t, listErr)
	assert.NotNil(t, enrichmentList)

	var collectionFound = false
	for _, e := range enrichmentList.Enrichments {
		if *e.EnrichmentID == *enrichment.EnrichmentID {
			collectionFound = true
		}
	}

	assert.Equal(t, collectionFound, true)

	updatedEnrichment, _, updateErr := service.UpdateEnrichment(
		&discoveryv2.UpdateEnrichmentOptions{
			ProjectID:    projectID,
			EnrichmentID: enrichment.EnrichmentID,
			Name:         core.StringPtr("updated enrichment"),
			Description:  core.StringPtr("updated"),
		},
	)

	assert.Nil(t, updateErr)
	assert.NotNil(t, updatedEnrichment)
	assert.Equal(t, *updatedEnrichment.Name, "updated enrichment")

	_, deleteErr := service.DeleteEnrichment(
		&discoveryv2.DeleteEnrichmentOptions{
			ProjectID:    projectID,
			EnrichmentID: enrichment.EnrichmentID,
		},
	)

	assert.Nil(t, deleteErr)
}
