//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package discoveryv2_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/discoveryv2"
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
const externalConfigFile = "../discovery_v2.env"

var (
	discoveryService *discoveryv2.DiscoveryV2
	config           map[string]string
	configLoaded     bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`DiscoveryV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(discoveryv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			discoveryServiceOptions := &discoveryv2.DiscoveryV2Options{
				Version: core.StringPtr("testString"),
			}

			discoveryService, err = discoveryv2.NewDiscoveryV2(discoveryServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(discoveryService).ToNot(BeNil())
		})
	})

	Describe(`DiscoveryV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
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
		It(`CreateCollection request example`, func() {
			fmt.Println("\nCreateCollection() result:")
			// begin-createCollection

			createCollectionOptions := discoveryService.NewCreateCollectionOptions(
				"testString",
				"testString",
			)

			collectionDetails, response, err := discoveryService.CreateCollection(createCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionDetails, "", "  ")
			fmt.Println(string(b))

			// end-createCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionDetails).ToNot(BeNil())

		})
		It(`GetCollection request example`, func() {
			fmt.Println("\nGetCollection() result:")
			// begin-getCollection

			getCollectionOptions := discoveryService.NewGetCollectionOptions(
				"testString",
				"testString",
			)

			collectionDetails, response, err := discoveryService.GetCollection(getCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionDetails, "", "  ")
			fmt.Println(string(b))

			// end-getCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionDetails).ToNot(BeNil())

		})
		It(`UpdateCollection request example`, func() {
			fmt.Println("\nUpdateCollection() result:")
			// begin-updateCollection

			updateCollectionOptions := discoveryService.NewUpdateCollectionOptions(
				"testString",
				"testString",
			)

			collectionDetails, response, err := discoveryService.UpdateCollection(updateCollectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(collectionDetails, "", "  ")
			fmt.Println(string(b))

			// end-updateCollection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(collectionDetails).ToNot(BeNil())

		})
		It(`Query request example`, func() {
			fmt.Println("\nQuery() result:")
			// begin-query

			queryOptions := discoveryService.NewQueryOptions(
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
		It(`GetAutocompletion request example`, func() {
			fmt.Println("\nGetAutocompletion() result:")
			// begin-getAutocompletion

			getAutocompletionOptions := discoveryService.NewGetAutocompletionOptions(
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
		It(`QueryCollectionNotices request example`, func() {
			fmt.Println("\nQueryCollectionNotices() result:")
			// begin-queryCollectionNotices

			queryCollectionNoticesOptions := discoveryService.NewQueryCollectionNoticesOptions(
				"testString",
				"testString",
			)

			queryNoticesResponse, response, err := discoveryService.QueryCollectionNotices(queryCollectionNoticesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(queryNoticesResponse, "", "  ")
			fmt.Println(string(b))

			// end-queryCollectionNotices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(queryNoticesResponse).ToNot(BeNil())

		})
		It(`QueryNotices request example`, func() {
			fmt.Println("\nQueryNotices() result:")
			// begin-queryNotices

			queryNoticesOptions := discoveryService.NewQueryNoticesOptions(
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
		It(`ListFields request example`, func() {
			fmt.Println("\nListFields() result:")
			// begin-listFields

			listFieldsOptions := discoveryService.NewListFieldsOptions(
				"testString",
			)

			listFieldsResponse, response, err := discoveryService.ListFields(listFieldsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listFieldsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listFields

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listFieldsResponse).ToNot(BeNil())

		})
		It(`GetComponentSettings request example`, func() {
			fmt.Println("\nGetComponentSettings() result:")
			// begin-getComponentSettings

			getComponentSettingsOptions := discoveryService.NewGetComponentSettingsOptions(
				"testString",
			)

			componentSettingsResponse, response, err := discoveryService.GetComponentSettings(getComponentSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(componentSettingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-getComponentSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(componentSettingsResponse).ToNot(BeNil())

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
		It(`ListTrainingQueries request example`, func() {
			fmt.Println("\nListTrainingQueries() result:")
			// begin-listTrainingQueries

			listTrainingQueriesOptions := discoveryService.NewListTrainingQueriesOptions(
				"testString",
			)

			trainingQuerySet, response, err := discoveryService.ListTrainingQueries(listTrainingQueriesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuerySet, "", "  ")
			fmt.Println(string(b))

			// end-listTrainingQueries

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingQuerySet).ToNot(BeNil())

		})
		It(`CreateTrainingQuery request example`, func() {
			fmt.Println("\nCreateTrainingQuery() result:")
			// begin-createTrainingQuery

			trainingExampleModel := &discoveryv2.TrainingExample{
				DocumentID:   core.StringPtr("testString"),
				CollectionID: core.StringPtr("testString"),
				Relevance:    core.Int64Ptr(int64(38)),
			}

			createTrainingQueryOptions := discoveryService.NewCreateTrainingQueryOptions(
				"testString",
				"testString",
				[]discoveryv2.TrainingExample{*trainingExampleModel},
			)

			trainingQuery, response, err := discoveryService.CreateTrainingQuery(createTrainingQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuery, "", "  ")
			fmt.Println(string(b))

			// end-createTrainingQuery

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trainingQuery).ToNot(BeNil())

		})
		It(`GetTrainingQuery request example`, func() {
			fmt.Println("\nGetTrainingQuery() result:")
			// begin-getTrainingQuery

			getTrainingQueryOptions := discoveryService.NewGetTrainingQueryOptions(
				"testString",
				"testString",
			)

			trainingQuery, response, err := discoveryService.GetTrainingQuery(getTrainingQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuery, "", "  ")
			fmt.Println(string(b))

			// end-getTrainingQuery

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trainingQuery).ToNot(BeNil())

		})
		It(`UpdateTrainingQuery request example`, func() {
			fmt.Println("\nUpdateTrainingQuery() result:")
			// begin-updateTrainingQuery

			trainingExampleModel := &discoveryv2.TrainingExample{
				DocumentID:   core.StringPtr("testString"),
				CollectionID: core.StringPtr("testString"),
				Relevance:    core.Int64Ptr(int64(38)),
			}

			updateTrainingQueryOptions := discoveryService.NewUpdateTrainingQueryOptions(
				"testString",
				"testString",
				"testString",
				[]discoveryv2.TrainingExample{*trainingExampleModel},
			)

			trainingQuery, response, err := discoveryService.UpdateTrainingQuery(updateTrainingQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trainingQuery, "", "  ")
			fmt.Println(string(b))

			// end-updateTrainingQuery

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trainingQuery).ToNot(BeNil())

		})
		It(`AnalyzeDocument request example`, func() {
			fmt.Println("\nAnalyzeDocument() result:")
			// begin-analyzeDocument

			analyzeDocumentOptions := discoveryService.NewAnalyzeDocumentOptions(
				"testString",
				"testString",
			)

			analyzedDocument, response, err := discoveryService.AnalyzeDocument(analyzeDocumentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(analyzedDocument, "", "  ")
			fmt.Println(string(b))

			// end-analyzeDocument

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(analyzedDocument).ToNot(BeNil())

		})
		It(`ListEnrichments request example`, func() {
			fmt.Println("\nListEnrichments() result:")
			// begin-listEnrichments

			listEnrichmentsOptions := discoveryService.NewListEnrichmentsOptions(
				"testString",
			)

			enrichments, response, err := discoveryService.ListEnrichments(listEnrichmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enrichments, "", "  ")
			fmt.Println(string(b))

			// end-listEnrichments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enrichments).ToNot(BeNil())

		})
		It(`CreateEnrichment request example`, func() {
			fmt.Println("\nCreateEnrichment() result:")
			// begin-createEnrichment

			createEnrichmentModel := &discoveryv2.CreateEnrichment{}

			createEnrichmentOptions := discoveryService.NewCreateEnrichmentOptions(
				"testString",
				createEnrichmentModel,
			)

			enrichment, response, err := discoveryService.CreateEnrichment(createEnrichmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enrichment, "", "  ")
			fmt.Println(string(b))

			// end-createEnrichment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(enrichment).ToNot(BeNil())

		})
		It(`GetEnrichment request example`, func() {
			fmt.Println("\nGetEnrichment() result:")
			// begin-getEnrichment

			getEnrichmentOptions := discoveryService.NewGetEnrichmentOptions(
				"testString",
				"testString",
			)

			enrichment, response, err := discoveryService.GetEnrichment(getEnrichmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enrichment, "", "  ")
			fmt.Println(string(b))

			// end-getEnrichment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enrichment).ToNot(BeNil())

		})
		It(`UpdateEnrichment request example`, func() {
			fmt.Println("\nUpdateEnrichment() result:")
			// begin-updateEnrichment

			updateEnrichmentOptions := discoveryService.NewUpdateEnrichmentOptions(
				"testString",
				"testString",
				"testString",
			)

			enrichment, response, err := discoveryService.UpdateEnrichment(updateEnrichmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enrichment, "", "  ")
			fmt.Println(string(b))

			// end-updateEnrichment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enrichment).ToNot(BeNil())

		})
		It(`ListProjects request example`, func() {
			fmt.Println("\nListProjects() result:")
			// begin-listProjects

			listProjectsOptions := discoveryService.NewListProjectsOptions()

			listProjectsResponse, response, err := discoveryService.ListProjects(listProjectsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listProjectsResponse, "", "  ")
			fmt.Println(string(b))

			// end-listProjects

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listProjectsResponse).ToNot(BeNil())

		})
		It(`CreateProject request example`, func() {
			fmt.Println("\nCreateProject() result:")
			// begin-createProject

			createProjectOptions := discoveryService.NewCreateProjectOptions(
				"testString",
				"document_retrieval",
			)

			projectDetails, response, err := discoveryService.CreateProject(createProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectDetails, "", "  ")
			fmt.Println(string(b))

			// end-createProject

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectDetails).ToNot(BeNil())

		})
		It(`GetProject request example`, func() {
			fmt.Println("\nGetProject() result:")
			// begin-getProject

			getProjectOptions := discoveryService.NewGetProjectOptions(
				"testString",
			)

			projectDetails, response, err := discoveryService.GetProject(getProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectDetails, "", "  ")
			fmt.Println(string(b))

			// end-getProject

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectDetails).ToNot(BeNil())

		})
		It(`UpdateProject request example`, func() {
			fmt.Println("\nUpdateProject() result:")
			// begin-updateProject

			updateProjectOptions := discoveryService.NewUpdateProjectOptions(
				"testString",
			)

			projectDetails, response, err := discoveryService.UpdateProject(updateProjectOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(projectDetails, "", "  ")
			fmt.Println(string(b))

			// end-updateProject

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(projectDetails).ToNot(BeNil())

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

			// end-deleteUserData
			fmt.Printf("\nDeleteUserData() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`DeleteTrainingQuery request example`, func() {
			// begin-deleteTrainingQuery

			deleteTrainingQueryOptions := discoveryService.NewDeleteTrainingQueryOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteTrainingQuery(deleteTrainingQueryOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteTrainingQuery
			fmt.Printf("\nDeleteTrainingQuery() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTrainingQueries request example`, func() {
			// begin-deleteTrainingQueries

			deleteTrainingQueriesOptions := discoveryService.NewDeleteTrainingQueriesOptions(
				"testString",
			)

			response, err := discoveryService.DeleteTrainingQueries(deleteTrainingQueriesOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteTrainingQueries
			fmt.Printf("\nDeleteTrainingQueries() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteProject request example`, func() {
			// begin-deleteProject

			deleteProjectOptions := discoveryService.NewDeleteProjectOptions(
				"testString",
			)

			response, err := discoveryService.DeleteProject(deleteProjectOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteProject
			fmt.Printf("\nDeleteProject() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteEnrichment request example`, func() {
			// begin-deleteEnrichment

			deleteEnrichmentOptions := discoveryService.NewDeleteEnrichmentOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteEnrichment(deleteEnrichmentOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteEnrichment
			fmt.Printf("\nDeleteEnrichment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

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
		It(`DeleteCollection request example`, func() {
			// begin-deleteCollection

			deleteCollectionOptions := discoveryService.NewDeleteCollectionOptions(
				"testString",
				"testString",
			)

			response, err := discoveryService.DeleteCollection(deleteCollectionOptions)
			if err != nil {
				panic(err)
			}

			// end-deleteCollection
			fmt.Printf("\nDeleteCollection() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
