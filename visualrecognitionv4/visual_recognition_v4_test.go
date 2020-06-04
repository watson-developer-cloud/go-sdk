/**
 * (C) Copyright IBM Corp. 2019, 2020.
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

package visualrecognitionv4_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv4"
)

var _ = Describe(`VisualRecognitionV4`, func() {
	Describe(`Analyze(analyzeOptions *AnalyzeOptions)`, func() {
		analyzePath := "/v4/analyze"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionIds := []string{}
		features := []string{}
		Context(`Successfully - Analyze images`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(analyzePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"images": []}`)
			}))
			It(`Succeed to call Analyze`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Analyze(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				analyzeOptions := testService.NewAnalyzeOptions(collectionIds, features)
				result, response, operationErr = testService.Analyze(analyzeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateCollection(createCollectionOptions *CreateCollectionOptions)`, func() {
		createCollectionPath := "/v4/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Create a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"collection_id": "fake_CollectionID", "name": "fake_Name", "description": "fake_Description", "created": "2017-05-16T13:56:54.957Z", "updated": "2017-05-16T13:56:54.957Z", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "description": "fake_Description"}}}`)
			}))
			It(`Succeed to call CreateCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createCollectionOptions := testService.NewCreateCollectionOptions()
				result, response, operationErr = testService.CreateCollection(createCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCollections(listCollectionsOptions *ListCollectionsOptions)`, func() {
		listCollectionsPath := "/v4/collections"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - List collections`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCollectionsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"collections": []}`)
			}))
			It(`Succeed to call ListCollections`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCollections(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCollectionsOptions := testService.NewListCollectionsOptions()
				result, response, operationErr = testService.ListCollections(listCollectionsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCollection(getCollectionOptions *GetCollectionOptions)`, func() {
		getCollectionPath := "/v4/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		getCollectionPath = strings.Replace(getCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get collection details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"collection_id": "fake_CollectionID", "name": "fake_Name", "description": "fake_Description", "created": "2017-05-16T13:56:54.957Z", "updated": "2017-05-16T13:56:54.957Z", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "description": "fake_Description"}}}`)
			}))
			It(`Succeed to call GetCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCollectionOptions := testService.NewGetCollectionOptions(collectionID)
				result, response, operationErr = testService.GetCollection(getCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateCollection(updateCollectionOptions *UpdateCollectionOptions)`, func() {
		updateCollectionPath := "/v4/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		updateCollectionPath = strings.Replace(updateCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Update a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"collection_id": "fake_CollectionID", "name": "fake_Name", "description": "fake_Description", "created": "2017-05-16T13:56:54.957Z", "updated": "2017-05-16T13:56:54.957Z", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "description": "fake_Description"}}}`)
			}))
			It(`Succeed to call UpdateCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateCollectionOptions := testService.NewUpdateCollectionOptions(collectionID)
				result, response, operationErr = testService.UpdateCollection(updateCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCollection(deleteCollectionOptions *DeleteCollectionOptions)`, func() {
		deleteCollectionPath := "/v4/collections/{collection_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		deleteCollectionPath = strings.Replace(deleteCollectionPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Delete a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCollectionPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteCollection`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteCollection(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteCollectionOptions := testService.NewDeleteCollectionOptions(collectionID)
				response, operationErr = testService.DeleteCollection(deleteCollectionOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`AddImages(addImagesOptions *AddImagesOptions)`, func() {
		addImagesPath := "/v4/collections/{collection_id}/images"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		addImagesPath = strings.Replace(addImagesPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Add images`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addImagesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call AddImages`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.AddImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				addImagesOptions := testService.NewAddImagesOptions(collectionID).SetImageURL([]string{"bogus"})
				result, response, operationErr = testService.AddImages(addImagesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListImages(listImagesOptions *ListImagesOptions)`, func() {
		listImagesPath := "/v4/collections/{collection_id}/images"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		listImagesPath = strings.Replace(listImagesPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - List images`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listImagesPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"images": []}`)
			}))
			It(`Succeed to call ListImages`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListImages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listImagesOptions := testService.NewListImagesOptions(collectionID)
				result, response, operationErr = testService.ListImages(listImagesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetImageDetails(getImageDetailsOptions *GetImageDetailsOptions)`, func() {
		getImageDetailsPath := "/v4/collections/{collection_id}/images/{image_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		imageID := "exampleString"
		getImageDetailsPath = strings.Replace(getImageDetailsPath, "{collection_id}", collectionID, 1)
		getImageDetailsPath = strings.Replace(getImageDetailsPath, "{image_id}", imageID, 1)
		Context(`Successfully - Get image details`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getImageDetailsPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"source": {"type": "fake_Type"}}`)
			}))
			It(`Succeed to call GetImageDetails`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetImageDetails(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getImageDetailsOptions := testService.NewGetImageDetailsOptions(collectionID, imageID)
				result, response, operationErr = testService.GetImageDetails(getImageDetailsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {
		deleteImagePath := "/v4/collections/{collection_id}/images/{image_id}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		imageID := "exampleString"
		deleteImagePath = strings.Replace(deleteImagePath, "{collection_id}", collectionID, 1)
		deleteImagePath = strings.Replace(deleteImagePath, "{image_id}", imageID, 1)
		Context(`Successfully - Delete an image`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteImagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteImage`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteImageOptions := testService.NewDeleteImageOptions(collectionID, imageID)
				response, operationErr = testService.DeleteImage(deleteImageOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetJpegImage(getJpegImageOptions *GetJpegImageOptions)`, func() {
		getJpegImagePath := "/v4/collections/{collection_id}/images/{image_id}/jpeg"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		imageID := "exampleString"
		getJpegImagePath = strings.Replace(getJpegImagePath, "{collection_id}", collectionID, 1)
		getJpegImagePath = strings.Replace(getJpegImagePath, "{image_id}", imageID, 1)
		Context(`Successfully - Get a JPEG file of an image`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getJpegImagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
				myImage, err := os.Open("../resources/my-giraffe.jpeg")
				if err != nil {
					panic(err)
				}
				bytes, err := ioutil.ReadAll(myImage)
				if err != nil {
					panic(err)
				}
				_, _ = res.Write(bytes)
			}))
			It(`Succeed to call GetJpegImage`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetJpegImage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getJpegImageOptions := testService.NewGetJpegImageOptions(collectionID, imageID)
				result, response, operationErr = testService.GetJpegImage(getJpegImageOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListObjectMetadata(listObjectMetadataOptions *ListObjectMetadataOptions)`, func() {
		listObjectMetadataPath := "/v4/collections/{collection_id}/objects"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		listObjectMetadataPath = strings.Replace(listObjectMetadataPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - List object metadata`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listObjectMetadataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"object_count": 11}`)
			}))
			It(`Succeed to call ListObjectMetadata`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listObjectMetadataOptions := testService.NewListObjectMetadataOptions(collectionID)
				result, response, operationErr = testService.ListObjectMetadata(listObjectMetadataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateObjectMetadata(updateObjectMetadataOptions *UpdateObjectMetadataOptions)`, func() {
		updateObjectMetadataPath := "/v4/collections/{collection_id}/objects/{object}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		object := "exampleString"
		newObject := "exampleString"
		updateObjectMetadataPath = strings.Replace(updateObjectMetadataPath, "{collection_id}", collectionID, 1)
		updateObjectMetadataPath = strings.Replace(updateObjectMetadataPath, "{object}", object, 1)
		Context(`Successfully - Update an object name`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateObjectMetadataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"object": "fake_Object", "count": 5}`)
			}))
			It(`Succeed to call UpdateObjectMetadata`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.UpdateObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				updateObjectMetadataOptions := testService.NewUpdateObjectMetadataOptions(collectionID, object, newObject)
				result, response, operationErr = testService.UpdateObjectMetadata(updateObjectMetadataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetObjectMetadata(getObjectMetadataOptions *GetObjectMetadataOptions)`, func() {
		getObjectMetadataPath := "/v4/collections/{collection_id}/objects/{object}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		object := "exampleString"
		getObjectMetadataPath = strings.Replace(getObjectMetadataPath, "{collection_id}", collectionID, 1)
		getObjectMetadataPath = strings.Replace(getObjectMetadataPath, "{object}", object, 1)
		Context(`Successfully - Get object metadata`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getObjectMetadataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetObjectMetadata`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetObjectMetadata(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getObjectMetadataOptions := testService.NewGetObjectMetadataOptions(collectionID, object)
				result, response, operationErr = testService.GetObjectMetadata(getObjectMetadataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteObject(deleteObjectOptions *DeleteObjectOptions)`, func() {
		deleteObjectPath := "/v4/collections/{collection_id}/objects/{object}"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		object := "exampleString"
		deleteObjectPath = strings.Replace(deleteObjectPath, "{collection_id}", collectionID, 1)
		deleteObjectPath = strings.Replace(deleteObjectPath, "{object}", object, 1)
		Context(`Successfully - Delete an object`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteObjectPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteObject`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteObject(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteObjectOptions := testService.NewDeleteObjectOptions(collectionID, object)
				response, operationErr = testService.DeleteObject(deleteObjectOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`Train(trainOptions *TrainOptions)`, func() {
		trainPath := "/v4/collections/{collection_id}/train"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		trainPath = strings.Replace(trainPath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Train a collection`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(trainPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"collection_id": "fake_CollectionID", "name": "fake_Name", "description": "fake_Description", "created": "2017-05-16T13:56:54.957Z", "updated": "2017-05-16T13:56:54.957Z", "image_count": 10, "training_status": {"objects": {"ready": false, "in_progress": true, "data_changed": false, "latest_failed": true, "description": "fake_Description"}}}`)
			}))
			It(`Succeed to call Train`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Train(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				trainOptions := testService.NewTrainOptions(collectionID)
				result, response, operationErr = testService.Train(trainOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddImageTrainingData(addImageTrainingDataOptions *AddImageTrainingDataOptions)`, func() {
		addImageTrainingDataPath := "/v4/collections/{collection_id}/images/{image_id}/training_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		imageID := "exampleString"
		addImageTrainingDataPath = strings.Replace(addImageTrainingDataPath, "{collection_id}", collectionID, 1)
		addImageTrainingDataPath = strings.Replace(addImageTrainingDataPath, "{image_id}", imageID, 1)
		Context(`Successfully - Add training data to an image`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addImageTrainingDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call AddImageTrainingData`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.AddImageTrainingData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				addImageTrainingDataOptions := testService.NewAddImageTrainingDataOptions(collectionID, imageID)
				result, response, operationErr = testService.AddImageTrainingData(addImageTrainingDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetModelFile(getModelFileOptions *GetModelFileOptions)`, func() {
		getModelFilePath := "/v4/collections/{collection_id}/model"
		version := "exampleString"
		bearerToken := "0ui9876453"
		collectionID := "exampleString"
		feature := "exampleString"
		modelFormat := "exampleString"
		getModelFilePath = strings.Replace(getModelFilePath, "{collection_id}", collectionID, 1)
		Context(`Successfully - Get a model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getModelFilePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["feature"]).To(Equal([]string{feature}))

				Expect(req.URL.Query()["model_format"]).To(Equal([]string{modelFormat}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Succeed to call GetModelFile`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetModelFile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getModelFileOptions := testService.NewGetModelFileOptions(collectionID, feature, modelFormat)
				result, response, operationErr = testService.GetModelFile(getModelFileOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetTrainingUsage(getTrainingUsageOptions *GetTrainingUsageOptions)`, func() {
		getTrainingUsagePath := "/v4/training_usage"
		version := "exampleString"
		bearerToken := "0ui9876453"
		Context(`Successfully - Get training usage`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getTrainingUsagePath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetTrainingUsage`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetTrainingUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getTrainingUsageOptions := testService.NewGetTrainingUsageOptions()
				result, response, operationErr = testService.GetTrainingUsage(getTrainingUsageOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v4/user_data"
		version := "exampleString"
		bearerToken := "0ui9876453"
		customerID := "exampleString"
		Context(`Successfully - Delete labeled data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.URL.Query()["version"]).To(Equal([]string{version}))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(202)
			}))
			It(`Succeed to call DeleteUserData`, func() {
				defer testServer.Close()

				testService, testServiceErr := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
					URL:     testServer.URL,
					Version: version,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				response, operationErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			version := "1970-01-01"
			testService, _ := visualrecognitionv4.NewVisualRecognitionV4(&visualrecognitionv4.VisualRecognitionV4Options{
				URL:           "http://visualrecognitionv4modelgenerator.com",
				Version:       version,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewFileWithMetadata successfully", func() {
				data := new(os.File)
				model, err := testService.NewFileWithMetadata(data)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewLocation successfully", func() {
				top := int64(1234)
				left := int64(1234)
				width := int64(1234)
				height := int64(1234)
				model, err := testService.NewLocation(top, left, width, height)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewObjectTrainingStatus successfully", func() {
				ready := true
				inProgress := true
				dataChanged := true
				latestFailed := true
				rsCnnReady := false
				description := "exampleString"
				model, err := testService.NewObjectTrainingStatus(ready, inProgress, dataChanged, latestFailed, rsCnnReady, description)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewTrainingStatus successfully", func() {
				objects := &visualrecognitionv4.ObjectTrainingStatus{Ready: core.BoolPtr(true), InProgress: core.BoolPtr(true), DataChanged: core.BoolPtr(true), LatestFailed: core.BoolPtr(true), Description: core.StringPtr("exampleString"), RscnnReady: core.BoolPtr(false)}
				model, err := testService.NewTrainingStatus(objects)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewUpdateObjectMetadata successfully", func() {
				object := "exampleString"
				count := int64(1234)
				model, err := testService.NewUpdateObjectMetadata(object, count)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
