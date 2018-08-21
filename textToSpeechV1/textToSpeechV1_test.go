package textToSpeechV1_test

import (
	"github.ibm.com/arf/go-sdk/textToSpeechV1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/base64"
	"net/http/httptest"
	"net/http"
	"strings"
	"fmt"
	"os"
	"encoding/json"
	"github.com/cloudfoundry-community/go-cfenv"
)

var _ = Describe("TextToSpeechV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "text_to_speech",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://stream.watsonplatform.net/text-to-speech/api",
						"username": username,
						"password": password,
					},
				},
			},
		}
		VCAPbytes, _ := json.Marshal(cfenv.App{})
		os.Setenv("VCAP_APPLICATION", string(VCAPbytes))
		VCAPbytes, _ = json.Marshal(VCAPservices)
		os.Setenv("VCAP_SERVICES", string(VCAPbytes))
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create TextToSpeechV1 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create TextToSpeechV1", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListVoiceModels(textToSpeechV1.NewListVoiceModelsOptions())
			})
		})
	})
	Describe("GetVoice(options *GetVoiceOptions)", func() {
		getVoicePath := "/v1/voices/{voice}"
        version := "exampleString"
        voice := "exampleString"
        getVoiceOptions := textToSpeechV1.NewGetVoiceOptions(voice)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getVoicePath = strings.Replace(getVoicePath, "{voice}", voice, 1)
		Context("Successfully - Get a voice", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getVoicePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetVoice", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetVoice(getVoiceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetGetVoiceResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListVoices(options *ListVoicesOptions)", func() {
		listVoicesPath := "/v1/voices"
        version := "exampleString"
        listVoicesOptions := textToSpeechV1.NewListVoicesOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List voices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listVoicesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListVoices", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListVoices(listVoicesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetListVoicesResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Synthesize(options *SynthesizeOptions)", func() {
		synthesizePath := "/v1/synthesize"
        version := "exampleString"
        text := "exampleString"
        synthesizeOptions := textToSpeechV1.NewSynthesizeOptions(text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Synthesize audio", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(synthesizePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Synthesize", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Synthesize(synthesizeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetSynthesizeResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetPronunciation(options *GetPronunciationOptions)", func() {
		getPronunciationPath := "/v1/pronunciation"
        version := "exampleString"
        text := "exampleString"
        getPronunciationOptions := textToSpeechV1.NewGetPronunciationOptions(text)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Get pronunciation", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getPronunciationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetPronunciation", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetPronunciation(getPronunciationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetGetPronunciationResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateVoiceModel(options *CreateVoiceModelOptions)", func() {
		createVoiceModelPath := "/v1/customizations"
        version := "exampleString"
        name := "exampleString"
        createVoiceModelOptions := textToSpeechV1.NewCreateVoiceModelOptions(name)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(createVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateVoiceModel(createVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetCreateVoiceModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteVoiceModel(options *DeleteVoiceModelOptions)", func() {
		deleteVoiceModelPath := "/v1/customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        deleteVoiceModelOptions := textToSpeechV1.NewDeleteVoiceModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteVoiceModelPath = strings.Replace(deleteVoiceModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Delete a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteVoiceModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteVoiceModel(deleteVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetVoiceModel(options *GetVoiceModelOptions)", func() {
		getVoiceModelPath := "/v1/customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        getVoiceModelOptions := textToSpeechV1.NewGetVoiceModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getVoiceModelPath = strings.Replace(getVoiceModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Get a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getVoiceModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetVoiceModel(getVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetGetVoiceModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListVoiceModels(options *ListVoiceModelsOptions)", func() {
		listVoiceModelsPath := "/v1/customizations"
        version := "exampleString"
        listVoiceModelsOptions := textToSpeechV1.NewListVoiceModelsOptions()
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List custom models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listVoiceModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListVoiceModels", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListVoiceModels(listVoiceModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetListVoiceModelsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateVoiceModel(options *UpdateVoiceModelOptions)", func() {
		updateVoiceModelPath := "/v1/customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        updateVoiceModelOptions := textToSpeechV1.NewUpdateVoiceModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        updateVoiceModelPath = strings.Replace(updateVoiceModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Update a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(updateVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpdateVoiceModel(updateVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddWord(options *AddWordOptions)", func() {
		addWordPath := "/v1/customizations/{customization_id}/words/{word}"
        version := "exampleString"
        customizationID := "exampleString"
        word := "exampleString"
        addWordOptions := textToSpeechV1.NewAddWordOptions(customizationID, word)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        addWordPath = strings.Replace(addWordPath, "{customization_id}", customizationID, 1)
        addWordPath = strings.Replace(addWordPath, "{word}", word, 1)
		Context("Successfully - Add a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(addWordPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddWord", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.AddWord(addWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddWords(options *AddWordsOptions)", func() {
		addWordsPath := "/v1/customizations/{customization_id}/words"
        version := "exampleString"
        customizationID := "exampleString"
        addWordsOptions := textToSpeechV1.NewAddWordsOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        addWordsPath = strings.Replace(addWordsPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Add custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(addWordsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddWords", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.AddWords(addWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteWord(options *DeleteWordOptions)", func() {
		deleteWordPath := "/v1/customizations/{customization_id}/words/{word}"
        version := "exampleString"
        customizationID := "exampleString"
        word := "exampleString"
        deleteWordOptions := textToSpeechV1.NewDeleteWordOptions(customizationID, word)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteWordPath = strings.Replace(deleteWordPath, "{customization_id}", customizationID, 1)
        deleteWordPath = strings.Replace(deleteWordPath, "{word}", word, 1)
		Context("Successfully - Delete a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteWordPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteWord", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteWord(deleteWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetWord(options *GetWordOptions)", func() {
		getWordPath := "/v1/customizations/{customization_id}/words/{word}"
        version := "exampleString"
        customizationID := "exampleString"
        word := "exampleString"
        getWordOptions := textToSpeechV1.NewGetWordOptions(customizationID, word)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getWordPath = strings.Replace(getWordPath, "{customization_id}", customizationID, 1)
        getWordPath = strings.Replace(getWordPath, "{word}", word, 1)
		Context("Successfully - Get a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getWordPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetWord", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetWord(getWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetGetWordResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWords(options *ListWordsOptions)", func() {
		listWordsPath := "/v1/customizations/{customization_id}/words"
        version := "exampleString"
        customizationID := "exampleString"
        listWordsOptions := textToSpeechV1.NewListWordsOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listWordsPath = strings.Replace(listWordsPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listWordsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListWords", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListWords(listWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := textToSpeechV1.GetListWordsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(options *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
        version := "exampleString"
        customerID := "exampleString"
        deleteUserDataOptions := textToSpeechV1.NewDeleteUserDataOptions(customerID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := textToSpeechV1.NewTextToSpeechV1(&textToSpeechV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
