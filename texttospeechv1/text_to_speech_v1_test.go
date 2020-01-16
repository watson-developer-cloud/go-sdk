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

package texttospeechv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe(`TextToSpeechV1`, func() {
	Describe(`ListVoices(listVoicesOptions *ListVoicesOptions)`, func() {
		listVoicesPath := "/v1/voices"
		bearerToken := "0ui9876453"
		Context(`Successfully - List voices`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVoicesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"voices": []}`)
			}))
			It(`Succeed to call ListVoices`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListVoices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listVoicesOptions := testService.NewListVoicesOptions()
				result, response, operationErr = testService.ListVoices(listVoicesOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVoice(getVoiceOptions *GetVoiceOptions)`, func() {
		getVoicePath := "/v1/voices/{voice}"
		bearerToken := "0ui9876453"
		voice := "exampleString"
		getVoicePath = strings.Replace(getVoicePath, "{voice}", voice, 1)
		Context(`Successfully - Get a voice`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVoicePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "fake_URL", "gender": "fake_Gender", "name": "fake_Name", "language": "fake_Language", "description": "fake_Description", "customizable": true, "supported_features": {"custom_pronunciation": false, "voice_transformation": false}}`)
			}))
			It(`Succeed to call GetVoice`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetVoice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getVoiceOptions := testService.NewGetVoiceOptions(voice)
				result, response, operationErr = testService.GetVoice(getVoiceOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Synthesize(synthesizeOptions *SynthesizeOptions)`, func() {
		synthesizePath := "/v1/synthesize"
		bearerToken := "0ui9876453"
		text := "exampleString"
		Context(`Successfully - Synthesize audio`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(synthesizePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `Contents of response byte-stream...`)
			}))
			It(`Succeed to call Synthesize`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.Synthesize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				synthesizeOptions := testService.NewSynthesizeOptions(text)
				result, response, operationErr = testService.Synthesize(synthesizeOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetPronunciation(getPronunciationOptions *GetPronunciationOptions)`, func() {
		getPronunciationPath := "/v1/pronunciation"
		bearerToken := "0ui9876453"
		text := "exampleString"
		Context(`Successfully - Get pronunciation`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getPronunciationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["text"]).To(Equal([]string{text}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"pronunciation": "fake_Pronunciation"}`)
			}))
			It(`Succeed to call GetPronunciation`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetPronunciation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getPronunciationOptions := testService.NewGetPronunciationOptions(text)
				result, response, operationErr = testService.GetPronunciation(getPronunciationOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateVoiceModel(createVoiceModelOptions *CreateVoiceModelOptions)`, func() {
		createVoiceModelPath := "/v1/customizations"
		bearerToken := "0ui9876453"
		name := "exampleString"
		Context(`Successfully - Create a custom model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call CreateVoiceModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.CreateVoiceModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				createVoiceModelOptions := testService.NewCreateVoiceModelOptions(name)
				result, response, operationErr = testService.CreateVoiceModel(createVoiceModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListVoiceModels(listVoiceModelsOptions *ListVoiceModelsOptions)`, func() {
		listVoiceModelsPath := "/v1/customizations"
		bearerToken := "0ui9876453"
		Context(`Successfully - List custom models`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listVoiceModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customizations": []}`)
			}))
			It(`Succeed to call ListVoiceModels`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.ListVoiceModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				listVoiceModelsOptions := testService.NewListVoiceModelsOptions()
				result, response, operationErr = testService.ListVoiceModels(listVoiceModelsOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateVoiceModel(updateVoiceModelOptions *UpdateVoiceModelOptions)`, func() {
		updateVoiceModelPath := "/v1/customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		updateVoiceModelPath = strings.Replace(updateVoiceModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Update a custom model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call UpdateVoiceModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.UpdateVoiceModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				updateVoiceModelOptions := testService.NewUpdateVoiceModelOptions(customizationID)
				response, operationErr = testService.UpdateVoiceModel(updateVoiceModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetVoiceModel(getVoiceModelOptions *GetVoiceModelOptions)`, func() {
		getVoiceModelPath := "/v1/customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		getVoiceModelPath = strings.Replace(getVoiceModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Get a custom model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getVoiceModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"customization_id": "fake_CustomizationID"}`)
			}))
			It(`Succeed to call GetVoiceModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				result, response, operationErr := testService.GetVoiceModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				getVoiceModelOptions := testService.NewGetVoiceModelOptions(customizationID)
				result, response, operationErr = testService.GetVoiceModel(getVoiceModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteVoiceModel(deleteVoiceModelOptions *DeleteVoiceModelOptions)`, func() {
		deleteVoiceModelPath := "/v1/customizations/{customization_id}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		deleteVoiceModelPath = strings.Replace(deleteVoiceModelPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Delete a custom model`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteVoiceModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteVoiceModel`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				response, operationErr := testService.DeleteVoiceModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				deleteVoiceModelOptions := testService.NewDeleteVoiceModelOptions(customizationID)
				response, operationErr = testService.DeleteVoiceModel(deleteVoiceModelOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`AddWords(addWordsOptions *AddWordsOptions)`, func() {
		addWordsPath := "/v1/customizations/{customization_id}/words"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		words := []texttospeechv1.Word{}
		addWordsPath = strings.Replace(addWordsPath, "{customization_id}", customizationID, 1)
		Context(`Successfully - Add custom words`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addWordsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call AddWords`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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
	Describe(`AddWord(addWordOptions *AddWordOptions)`, func() {
		addWordPath := "/v1/customizations/{customization_id}/words/{word}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		word := "exampleString"
		translation := "exampleString"
		addWordPath = strings.Replace(addWordPath, "{customization_id}", customizationID, 1)
		addWordPath = strings.Replace(addWordPath, "{word}", word, 1)
		Context(`Successfully - Add a custom word`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addWordPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(200)
			}))
			It(`Succeed to call AddWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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

				addWordOptions := testService.NewAddWordOptions(customizationID, word, translation)
				response, operationErr = testService.AddWord(addWordOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetWord(getWordOptions *GetWordOptions)`, func() {
		getWordPath := "/v1/customizations/{customization_id}/words/{word}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		word := "exampleString"
		getWordPath = strings.Replace(getWordPath, "{customization_id}", customizationID, 1)
		getWordPath = strings.Replace(getWordPath, "{word}", word, 1)
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
				fmt.Fprintf(res, `{"translation": "fake_Translation"}`)
			}))
			It(`Succeed to call GetWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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

				getWordOptions := testService.NewGetWordOptions(customizationID, word)
				result, response, operationErr = testService.GetWord(getWordOptions)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteWord(deleteWordOptions *DeleteWordOptions)`, func() {
		deleteWordPath := "/v1/customizations/{customization_id}/words/{word}"
		bearerToken := "0ui9876453"
		customizationID := "exampleString"
		word := "exampleString"
		deleteWordPath = strings.Replace(deleteWordPath, "{customization_id}", customizationID, 1)
		deleteWordPath = strings.Replace(deleteWordPath, "{word}", word, 1)
		Context(`Successfully - Delete a custom word`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteWordPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Succeed to call DeleteWord`, func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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

				deleteWordOptions := testService.NewDeleteWordOptions(customizationID, word)
				response, operationErr = testService.DeleteWord(deleteWordOptions)
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

				testService, testServiceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
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
	Describe("Model constructor tests", func() {
		Context("with a sample service", func() {
			testService, _ := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL:           "http://texttospeechv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It("should call NewTranslation successfully", func() {
				translation := "exampleString"
				model, err := testService.NewTranslation(translation)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewWord successfully", func() {
				word := "exampleString"
				translation := "exampleString"
				model, err := testService.NewWord(word, translation)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It("should call NewWords successfully", func() {
				words := []texttospeechv1.Word{}
				model, err := testService.NewWords(words)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
