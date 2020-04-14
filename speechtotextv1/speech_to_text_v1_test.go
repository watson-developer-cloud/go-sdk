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

package speechtotextv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe(`SpeechToTextV1`, func() {
	Describe(`ListModels(listModelsOptions *ListModelsOptions)`, func() {
		listModelsPath := "/v1/models"
		bearerToken := "0ui9876453"
		Context(`Successfully - List models`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"models": []}`)
			}))
			It(`Succeed to call ListModels`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listModelsOptions := testService.NewListModelsOptions()
				result, response, operationErr = testService.ListModels(listModelsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetModel(getModelOptions *GetModelOptions)`, func() {
		getModelPath := "/v1/models/{model_id}"
		bearerToken := "0ui9876453"
		modelID := "exampleString"
		getModelPath = strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context(`Successfully - Get a model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name", "language": "fake_Language", "rate": 4, "url": "fake_URL", "supported_features": {"custom_language_model": false, "speaker_labels": false}, "description": "fake_Description"}`)
			}))
			It(`Succeed to call GetModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getModelOptions := testService.NewGetModelOptions(modelID)
				result, response, operationErr = testService.GetModel(getModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Recognize(recognizeOptions *RecognizeOptions)`, func() {
		recognizePath := "/v1/recognize"
		bearerToken := "0ui9876453"
		Context(`Successfully - Recognize audio`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(recognizePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call Recognize`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Recognize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				file, err := os.Open("../resources/output.wav")
				if err != nil {
					panic(err)
				}
				recognizeOptions := testService.
					NewRecognizeOptions(file)
				result, response, operationErr = testService.Recognize(recognizeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RegisterCallback(registerCallbackOptions *RegisterCallbackOptions)`, func() {
		registerCallbackPath := "/v1/register_callback"
		bearerToken := "0ui9876453"
		callbackURL := "exampleString"
		Context(`Successfully - Register a callback`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(registerCallbackPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["callback_url"]).To(Equal([]string{callbackURL}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"status": "fake_Status", "url": "fake_URL"}`)
			}))
			It(`Succeed to call RegisterCallback`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.RegisterCallback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				registerCallbackOptions := testService.NewRegisterCallbackOptions(callbackURL)
				result, response, operationErr = testService.RegisterCallback(registerCallbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UnregisterCallback(unregisterCallbackOptions *UnregisterCallbackOptions)`, func() {
		unregisterCallbackPath := "/v1/unregister_callback"
		bearerToken := "0ui9876453"
		callbackURL := "exampleString"
		Context(`Successfully - Unregister a callback`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(unregisterCallbackPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["callback_url"]).To(Equal([]string{callbackURL}))

				res.WriteHeader(200)
			}))
			It(`Succeed to call UnregisterCallback`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.UnregisterCallback(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				unregisterCallbackOptions := testService.NewUnregisterCallbackOptions(callbackURL)
				response, operationErr = testService.UnregisterCallback(unregisterCallbackOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateJob(createJobOptions *CreateJobOptions)`, func() {
		createJobPath := "/v1/recognitions"
		bearerToken := "0ui9876453"
		Context(`Successfully - Create a job`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createJobPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "fake_ID", "status": "fake_Status", "created": "fake_Created"}`)
			}))
			It(`Succeed to call CreateJob`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				file, err := os.Open("../resources/output.wav")
				if err != nil {
					panic(err)
				}
				createJobOptions := testService.
					NewCreateJobOptions(file).
					SetContentType("audio/wav")
				result, response, operationErr = testService.CreateJob(createJobOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CheckJobs(checkJobsOptions *CheckJobsOptions)`, func() {
		checkJobsPath := "/v1/recognitions"
		bearerToken := "0ui9876453"
		Context(`Successfully - Check jobs`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(checkJobsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"recognitions": []}`)
			}))
			It(`Succeed to call CheckJobs`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CheckJobs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				checkJobsOptions := testService.NewCheckJobsOptions()
				result, response, operationErr = testService.CheckJobs(checkJobsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CheckJob(checkJobOptions *CheckJobOptions)`, func() {
		checkJobPath := "/v1/recognitions/{id}"
		bearerToken := "0ui9876453"
		id := "exampleString"
		checkJobPath = strings.Replace(checkJobPath, "{id}", id, 1)
		Context(`Successfully - Check a job`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(checkJobPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "fake_ID", "status": "fake_Status", "created": "fake_Created"}`)
			}))
			It(`Succeed to call CheckJob`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CheckJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				checkJobOptions := testService.NewCheckJobOptions(id)
				result, response, operationErr = testService.CheckJob(checkJobOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteJob(deleteJobOptions *DeleteJobOptions)`, func() {
		deleteJobPath := "/v1/recognitions/{id}"
		bearerToken := "0ui9876453"
		id := "exampleString"
		deleteJobPath = strings.Replace(deleteJobPath, "{id}", id, 1)
		Context(`Successfully - Delete a job`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteJobPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteJob`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteJob(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteJobOptions := testService.NewDeleteJobOptions(id)
				response, operationErr = testService.DeleteJob(deleteJobOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions)`, func() {
		createLanguageModelPath := "/v1/customizations"
		bearerToken := "0ui9876453"
		name := "exampleString"
		baseModelName := "exampleString"
		Context(`Successfully - Create a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call CreateLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createLanguageModelOptions := testService.NewCreateLanguageModelOptions(name, baseModelName)
				result, response, operationErr = testService.CreateLanguageModel(createLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions)`, func() {
		listLanguageModelsPath := "/v1/customizations"
		bearerToken := "0ui9876453"
		Context(`Successfully - List custom language models`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listLanguageModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customizations": []}`)
			}))
			It(`Succeed to call ListLanguageModels`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListLanguageModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listLanguageModelsOptions := testService.NewListLanguageModelsOptions()
				result, response, operationErr = testService.ListLanguageModels(listLanguageModelsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions)`, func() {
		getLanguageModelPath := "/v1/customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		getLanguageModelPath = strings.Replace(getLanguageModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Get a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getLanguageModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call GetLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getLanguageModelOptions := testService.NewGetLanguageModelOptions(customizationID)
				result, response, operationErr = testService.GetLanguageModel(getLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteLanguageModel(deleteLanguageModelOptions *DeleteLanguageModelOptions)`, func() {
		deleteLanguageModelPath := "/v1/customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		deleteLanguageModelPath = strings.Replace(deleteLanguageModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Delete a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteLanguageModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteLanguageModelOptions := testService.NewDeleteLanguageModelOptions(customizationID)
				response, operationErr = testService.DeleteLanguageModel(deleteLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions)`, func() {
		trainLanguageModelPath := "/v1/customizations/{customization_id}/train"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		trainLanguageModelPath = strings.Replace(trainLanguageModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Train a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(trainLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call TrainLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.TrainLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				trainLanguageModelOptions := testService.NewTrainLanguageModelOptions(customizationID)
				result, response, operationErr = testService.TrainLanguageModel(trainLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ResetLanguageModel(resetLanguageModelOptions *ResetLanguageModelOptions)`, func() {
		resetLanguageModelPath := "/v1/customizations/{customization_id}/reset"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		resetLanguageModelPath = strings.Replace(resetLanguageModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Reset a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(resetLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call ResetLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.ResetLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				resetLanguageModelOptions := testService.NewResetLanguageModelOptions(customizationID)
				response, operationErr = testService.ResetLanguageModel(resetLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpgradeLanguageModel(upgradeLanguageModelOptions *UpgradeLanguageModelOptions)`, func() {
		upgradeLanguageModelPath := "/v1/customizations/{customization_id}/upgrade_model"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		upgradeLanguageModelPath = strings.Replace(upgradeLanguageModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Upgrade a custom language model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(upgradeLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpgradeLanguageModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.UpgradeLanguageModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				upgradeLanguageModelOptions := testService.NewUpgradeLanguageModelOptions(customizationID)
				response, operationErr = testService.UpgradeLanguageModel(upgradeLanguageModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListCorpora(listCorporaOptions *ListCorporaOptions)`, func() {
		listCorporaPath := "/v1/customizations/{customization_id}/corpora"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		listCorporaPath = strings.Replace(listCorporaPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - List corpora`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listCorporaPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"corpora": []}`)
			}))
			It(`Succeed to call ListCorpora`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListCorpora(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listCorporaOptions := testService.NewListCorporaOptions(customizationID)
				result, response, operationErr = testService.ListCorpora(listCorporaOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddCorpus(addCorpusOptions *AddCorpusOptions)`, func() {
		addCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		corpusName := "exampleString"
		corpusFile, corpusErr := os.Open("../resources/corpus-short-1.txt")
		if corpusErr != nil {
			panic(corpusErr)
		}
		addCorpusPath = strings.Replace(addCorpusPath, "{customization_id}", customizationID, 1)
		addCorpusPath = strings.Replace(addCorpusPath, "{corpus_name}", corpusName, 1)
		Context(`Successfully - Add a corpus`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addCorpusPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(201)
			}))
			It(`Succeed to call AddCorpus`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.AddCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				addCorpusOptions := testService.NewAddCorpusOptions(customizationID, corpusName, corpusFile)
				response, operationErr = testService.AddCorpus(addCorpusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCorpus(getCorpusOptions *GetCorpusOptions)`, func() {
		getCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		corpusName := "exampleString"
		getCorpusPath = strings.Replace(getCorpusPath, "{customization_id}", customizationID, 1)
		getCorpusPath = strings.Replace(getCorpusPath, "{corpus_name}", corpusName, 1)
		Context(`Successfully - Get a corpus`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCorpusPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name", "total_words": 10, "out_of_vocabulary_words": 20, "status": "fake_Status"}`)
			}))
			It(`Succeed to call GetCorpus`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getCorpusOptions := testService.NewGetCorpusOptions(customizationID, corpusName)
				result, response, operationErr = testService.GetCorpus(getCorpusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteCorpus(deleteCorpusOptions *DeleteCorpusOptions)`, func() {
		deleteCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		corpusName := "exampleString"
		deleteCorpusPath = strings.Replace(deleteCorpusPath, "{customization_id}", customizationID, 1)
		deleteCorpusPath = strings.Replace(deleteCorpusPath, "{corpus_name}", corpusName, 1)
		Context(`Successfully - Delete a corpus`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteCorpusPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteCorpus`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteCorpusOptions := testService.NewDeleteCorpusOptions(customizationID, corpusName)
				response, operationErr = testService.DeleteCorpus(deleteCorpusOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListWords(listWordsOptions *ListWordsOptions)`, func() {
		listWordsPath := "/v1/customizations/{customization_id}/words"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		listWordsPath = strings.Replace(listWordsPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - List custom words`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listWordsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"words": []}`)
			}))
			It(`Succeed to call ListWords`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listWordsOptions := testService.NewListWordsOptions(customizationID)
				result, response, operationErr = testService.ListWords(listWordsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddWords(addWordsOptions *AddWordsOptions)`, func() {
		addWordsPath := "/v1/customizations/{customization_id}/words"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		words := []speechtotextv1.CustomWord{}
		addWordsPath = strings.Replace(addWordsPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Add custom words`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addWordsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(201)
			}))
			It(`Succeed to call AddWords`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.AddWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				addWordsOptions := testService.NewAddWordsOptions(customizationID, words)
				response, operationErr = testService.AddWords(addWordsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`AddWord(addWordOptions *AddWordOptions)`, func() {
		addWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		wordName := "exampleString"
		addWordPath = strings.Replace(addWordPath, "{customization_id}", customizationID, 1)
		addWordPath = strings.Replace(addWordPath, "{word_name}", wordName, 1)
		Context(`Successfully - Add a custom word`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addWordPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(201)
			}))
			It(`Succeed to call AddWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.AddWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				addWordOptions := testService.NewAddWordOptions(customizationID, wordName)
				response, operationErr = testService.AddWord(addWordOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetWord(getWordOptions *GetWordOptions)`, func() {
		getWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		wordName := "exampleString"
		getWordPath = strings.Replace(getWordPath, "{customization_id}", customizationID, 1)
		getWordPath = strings.Replace(getWordPath, "{word_name}", wordName, 1)
		Context(`Successfully - Get a custom word`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getWordPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"word": "fake_Word", "sounds_like": [], "display_as": "fake_DisplayAs", "count": 5, "source": []}`)
			}))
			It(`Succeed to call GetWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getWordOptions := testService.NewGetWordOptions(customizationID, wordName)
				result, response, operationErr = testService.GetWord(getWordOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteWord(deleteWordOptions *DeleteWordOptions)`, func() {
		deleteWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		wordName := "exampleString"
		deleteWordPath = strings.Replace(deleteWordPath, "{customization_id}", customizationID, 1)
		deleteWordPath = strings.Replace(deleteWordPath, "{word_name}", wordName, 1)
		Context(`Successfully - Delete a custom word`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteWordPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteWordOptions := testService.NewDeleteWordOptions(customizationID, wordName)
				response, operationErr = testService.DeleteWord(deleteWordOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListGrammars(listGrammarsOptions *ListGrammarsOptions)`, func() {
		listGrammarsPath := "/v1/customizations/{customization_id}/grammars"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		listGrammarsPath = strings.Replace(listGrammarsPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - List grammars`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listGrammarsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"grammars": []}`)
			}))
			It(`Succeed to call ListGrammars`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListGrammars(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listGrammarsOptions := testService.NewListGrammarsOptions(customizationID)
				result, response, operationErr = testService.ListGrammars(listGrammarsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddGrammar(addGrammarOptions *AddGrammarOptions)`, func() {
		addGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		grammarName := "exampleString"
		grammarFile, grammarFileErr := os.Open("../resources/confirm-grammar.xml")
		if grammarFileErr != nil {
			panic(grammarFileErr)
		}
		contentType := "exampleString"
		addGrammarPath = strings.Replace(addGrammarPath, "{customization_id}", customizationID, 1)
		addGrammarPath = strings.Replace(addGrammarPath, "{grammar_name}", grammarName, 1)
		Context(`Successfully - Add a grammar`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addGrammarPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(201)
			}))
			It(`Succeed to call AddGrammar`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.AddGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				addGrammarOptions := testService.NewAddGrammarOptions(customizationID, grammarName, grammarFile, contentType)
				response, operationErr = testService.AddGrammar(addGrammarOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetGrammar(getGrammarOptions *GetGrammarOptions)`, func() {
		getGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		grammarName := "exampleString"
		getGrammarPath = strings.Replace(getGrammarPath, "{customization_id}", customizationID, 1)
		getGrammarPath = strings.Replace(getGrammarPath, "{grammar_name}", grammarName, 1)
		Context(`Successfully - Get a grammar`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getGrammarPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"name": "fake_Name", "out_of_vocabulary_words": 20, "status": "fake_Status"}`)
			}))
			It(`Succeed to call GetGrammar`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getGrammarOptions := testService.NewGetGrammarOptions(customizationID, grammarName)
				result, response, operationErr = testService.GetGrammar(getGrammarOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteGrammar(deleteGrammarOptions *DeleteGrammarOptions)`, func() {
		deleteGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		grammarName := "exampleString"
		deleteGrammarPath = strings.Replace(deleteGrammarPath, "{customization_id}", customizationID, 1)
		deleteGrammarPath = strings.Replace(deleteGrammarPath, "{grammar_name}", grammarName, 1)
		Context(`Successfully - Delete a grammar`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteGrammarPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteGrammar`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteGrammar(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteGrammarOptions := testService.NewDeleteGrammarOptions(customizationID, grammarName)
				response, operationErr = testService.DeleteGrammar(deleteGrammarOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions)`, func() {
		createAcousticModelPath := "/v1/acoustic_customizations"
		bearerToken := "0ui9876453"
		name := "exampleString"
		baseModelName := "exampleString"
		Context(`Successfully - Create a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call CreateAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createAcousticModelOptions := testService.NewCreateAcousticModelOptions(name, baseModelName)
				result, response, operationErr = testService.CreateAcousticModel(createAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions)`, func() {
		listAcousticModelsPath := "/v1/acoustic_customizations"
		bearerToken := "0ui9876453"
		Context(`Successfully - List custom acoustic models`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAcousticModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customizations": []}`)
			}))
			It(`Succeed to call ListAcousticModels`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListAcousticModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listAcousticModelsOptions := testService.NewListAcousticModelsOptions()
				result, response, operationErr = testService.ListAcousticModels(listAcousticModelsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions)`, func() {
		getAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		getAcousticModelPath = strings.Replace(getAcousticModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Get a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAcousticModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call GetAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getAcousticModelOptions := testService.NewGetAcousticModelOptions(customizationID)
				result, response, operationErr = testService.GetAcousticModel(getAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteAcousticModel(deleteAcousticModelOptions *DeleteAcousticModelOptions)`, func() {
		deleteAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		deleteAcousticModelPath = strings.Replace(deleteAcousticModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Delete a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteAcousticModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteAcousticModelOptions := testService.NewDeleteAcousticModelOptions(customizationID)
				response, operationErr = testService.DeleteAcousticModel(deleteAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions)`, func() {
		trainAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/train"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		trainAcousticModelPath = strings.Replace(trainAcousticModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Train a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(trainAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call TrainAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.TrainAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				trainAcousticModelOptions := testService.NewTrainAcousticModelOptions(customizationID)
				result, response, operationErr = testService.TrainAcousticModel(trainAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ResetAcousticModel(resetAcousticModelOptions *ResetAcousticModelOptions)`, func() {
		resetAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/reset"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		resetAcousticModelPath = strings.Replace(resetAcousticModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Reset a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(resetAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call ResetAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.ResetAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				resetAcousticModelOptions := testService.NewResetAcousticModelOptions(customizationID)
				response, operationErr = testService.ResetAcousticModel(resetAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpgradeAcousticModel(upgradeAcousticModelOptions *UpgradeAcousticModelOptions)`, func() {
		upgradeAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/upgrade_model"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		upgradeAcousticModelPath = strings.Replace(upgradeAcousticModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Upgrade a custom acoustic model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(upgradeAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpgradeAcousticModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.UpgradeAcousticModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				upgradeAcousticModelOptions := testService.NewUpgradeAcousticModelOptions(customizationID)
				response, operationErr = testService.UpgradeAcousticModel(upgradeAcousticModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAudio(listAudioOptions *ListAudioOptions)`, func() {
		listAudioPath := "/v1/acoustic_customizations/{customization_id}/audio"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		listAudioPath = strings.Replace(listAudioPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - List audio resources`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAudioPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_minutes_of_audio": 19, "audio": []}`)
			}))
			It(`Succeed to call ListAudio`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listAudioOptions := testService.NewListAudioOptions(customizationID)
				result, response, operationErr = testService.ListAudio(listAudioOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddAudio(addAudioOptions *AddAudioOptions)`, func() {
		addAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		audioName := "exampleString"
		audioResource, audioResourceErr := os.Open("../resources/audio_example.mp3")
		if audioResourceErr != nil {
			panic(audioResourceErr)
		}
		addAudioPath = strings.Replace(addAudioPath, "{customization_id}", customizationID, 1)
		addAudioPath = strings.Replace(addAudioPath, "{audio_name}", audioName, 1)
		Context(`Successfully - Add an audio resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addAudioPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(201)
			}))
			It(`Succeed to call AddAudio`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.AddAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				addAudioOptions := testService.NewAddAudioOptions(customizationID, audioName, audioResource)
				response, operationErr = testService.AddAudio(addAudioOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAudio(getAudioOptions *GetAudioOptions)`, func() {
		getAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		audioName := "exampleString"
		getAudioPath = strings.Replace(getAudioPath, "{customization_id}", customizationID, 1)
		getAudioPath = strings.Replace(getAudioPath, "{audio_name}", audioName, 1)
		Context(`Successfully - Get an audio resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAudioPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{}`)
			}))
			It(`Succeed to call GetAudio`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getAudioOptions := testService.NewGetAudioOptions(customizationID, audioName)
				result, response, operationErr = testService.GetAudio(getAudioOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteAudio(deleteAudioOptions *DeleteAudioOptions)`, func() {
		deleteAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		audioName := "exampleString"
		deleteAudioPath = strings.Replace(deleteAudioPath, "{customization_id}", customizationID, 1)
		deleteAudioPath = strings.Replace(deleteAudioPath, "{audio_name}", audioName, 1)
		Context(`Successfully - Delete an audio resource`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteAudioPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteAudio`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteAudio(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteAudioOptions := testService.NewDeleteAudioOptions(customizationID, audioName)
				response, operationErr = testService.DeleteAudio(deleteAudioOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v1/user_data"
		bearerToken := "0ui9876453"
		customerID := "exampleString"
		Context(`Successfully - Delete labeled data`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["customer_id"]).To(Equal([]string{customerID}))

				res.WriteHeader(200)
			}))
			It(`Succeed to call DeleteUserData`, func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL: testServer.URL,
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
})
