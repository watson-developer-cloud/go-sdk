package texttospeechv1_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
)

var _ = Describe("TextToSpeechV1", func() {
	Describe("GetVoice(getVoiceOptions *GetVoiceOptions)", func() {
		GetVoicePath := "/v1/voices/{voice}"
		Voice := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetVoicePath = strings.Replace(GetVoicePath, "{voice}", Voice, 1)
		Context("Successfully - Get a voice", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(GetVoicePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"url":"https://test.com", "gender":"female"}`)
			}))
			It("Succeed to call GetVoice", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetVoiceOptions := testService.NewGetVoiceOptions(Voice)
				returnValue, returnValueErr := testService.GetVoice(GetVoiceOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetVoiceResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListVoices(listVoicesOptions *ListVoicesOptions)", func() {
		ListVoicesPath := "/v1/voices"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List voices", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(ListVoicesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `["gender":"female"]`)
			}))
			It("Succeed to call ListVoices", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ListVoicesOptions := testService.NewListVoicesOptions()
				returnValue, returnValueErr := testService.ListVoices(ListVoicesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListVoicesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Synthesize(synthesizeOptions *SynthesizeOptions)", func() {
		SynthesizePath := "/v1/synthesize"
		Text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Synthesize audio", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(SynthesizePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))

				res.WriteHeader(http.StatusOK)
				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/output.wav")
				if err != nil {
					panic(err)
				}
				bytes, err := ioutil.ReadAll(file)
				if err != nil {
					panic(err)
				}
				res.Write(bytes)
			}))
			It("Succeed to call Synthesize", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				SynthesizeOptions := testService.NewSynthesizeOptions(Text)
				returnValue, returnValueErr := testService.Synthesize(SynthesizeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetSynthesizeResult(returnValue)
				Expect(result).ToNot(BeNil())
				result.Close()
			})
		})
	})
	Describe("GetPronunciation(getPronunciationOptions *GetPronunciationOptions)", func() {
		GetPronunciationPath := "/v1/pronunciation"
		Text := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Get pronunciation", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(GetPronunciationPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"pronunciation":"creepy"}`)
			}))
			It("Succeed to call GetPronunciation", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetPronunciationOptions := testService.NewGetPronunciationOptions(Text)
				returnValue, returnValueErr := testService.GetPronunciation(GetPronunciationOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetPronunciationResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateVoiceModel(createVoiceModelOptions *CreateVoiceModelOptions)", func() {
		CreateVoiceModelPath := "/v1/customizations"
		Name := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(CreateVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call CreateVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				CreateVoiceModelOptions := testService.NewCreateVoiceModelOptions(Name)
				returnValue, returnValueErr := testService.CreateVoiceModel(CreateVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateVoiceModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteVoiceModel(deleteVoiceModelOptions *DeleteVoiceModelOptions)", func() {
		DeleteVoiceModelPath := "/v1/customizations/{customization_id}"
		CustomizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		DeleteVoiceModelPath = strings.Replace(DeleteVoiceModelPath, "{customization_id}", CustomizationID, 1)
		Context("Successfully - Delete a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(DeleteVoiceModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteVoiceModelOptions := testService.NewDeleteVoiceModelOptions(CustomizationID)
				returnValue, returnValueErr := testService.DeleteVoiceModel(DeleteVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetVoiceModel(getVoiceModelOptions *GetVoiceModelOptions)", func() {
		GetVoiceModelPath := "/v1/customizations/{customization_id}"
		CustomizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetVoiceModelPath = strings.Replace(GetVoiceModelPath, "{customization_id}", CustomizationID, 1)
		Context("Successfully - Get a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(GetVoiceModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call GetVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetVoiceModelOptions := testService.NewGetVoiceModelOptions(CustomizationID)
				returnValue, returnValueErr := testService.GetVoiceModel(GetVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetVoiceModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListVoiceModels(listVoiceModelsOptions *ListVoiceModelsOptions)", func() {
		ListVoiceModelsPath := "/v1/customizations"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List custom models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(ListVoiceModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListVoiceModels", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ListVoiceModelsOptions := testService.NewListVoiceModelsOptions()
				returnValue, returnValueErr := testService.ListVoiceModels(ListVoiceModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListVoiceModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateVoiceModel(updateVoiceModelOptions *UpdateVoiceModelOptions)", func() {
		UpdateVoiceModelPath := "/v1/customizations/{customization_id}"
		CustomizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		UpdateVoiceModelPath = strings.Replace(UpdateVoiceModelPath, "{customization_id}", CustomizationID, 1)
		Context("Successfully - Update a custom model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(UpdateVoiceModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpdateVoiceModel", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				UpdateVoiceModelOptions := testService.NewUpdateVoiceModelOptions(CustomizationID)
				returnValue, returnValueErr := testService.UpdateVoiceModel(UpdateVoiceModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddWord(addWordOptions *AddWordOptions)", func() {
		AddWordPath := "/v1/customizations/{customization_id}/words/{word}"
		CustomizationID := "exampleString"
		Word := "exampleString"
		Translation := "ex ample string"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		AddWordPath = strings.Replace(AddWordPath, "{customization_id}", CustomizationID, 1)
		AddWordPath = strings.Replace(AddWordPath, "{word}", Word, 1)
		Context("Successfully - Add a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(AddWordPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddWord", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				AddWordOptions := testService.NewAddWordOptions(CustomizationID, Word, Translation)
				returnValue, returnValueErr := testService.AddWord(AddWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddWords(addWordsOptions *AddWordsOptions)", func() {
		AddWordsPath := "/v1/customizations/{customization_id}/words"
		CustomizationID := "exampleString"
		username := "user1"
		password := "pass1"
		words := []texttospeechv1.Word{}
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		AddWordsPath = strings.Replace(AddWordsPath, "{customization_id}", CustomizationID, 1)
		Context("Successfully - Add custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(AddWordsPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddWords", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				AddWordsOptions := testService.NewAddWordsOptions(CustomizationID, words)
				returnValue, returnValueErr := testService.AddWords(AddWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteWord(deleteWordOptions *DeleteWordOptions)", func() {
		DeleteWordPath := "/v1/customizations/{customization_id}/words/{word}"
		CustomizationID := "exampleString"
		Word := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		DeleteWordPath = strings.Replace(DeleteWordPath, "{customization_id}", CustomizationID, 1)
		DeleteWordPath = strings.Replace(DeleteWordPath, "{word}", Word, 1)
		Context("Successfully - Delete a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(DeleteWordPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteWord", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteWordOptions := testService.NewDeleteWordOptions(CustomizationID, Word)
				returnValue, returnValueErr := testService.DeleteWord(DeleteWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetWord(getWordOptions *GetWordOptions)", func() {
		GetWordPath := "/v1/customizations/{customization_id}/words/{word}"
		CustomizationID := "exampleString"
		Word := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		GetWordPath = strings.Replace(GetWordPath, "{customization_id}", CustomizationID, 1)
		GetWordPath = strings.Replace(GetWordPath, "{word}", Word, 1)
		Context("Successfully - Get a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(GetWordPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"translation":"hello"}`)
			}))
			It("Succeed to call GetWord", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				GetWordOptions := testService.NewGetWordOptions(CustomizationID, Word)
				returnValue, returnValueErr := testService.GetWord(GetWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetWordResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWords(listWordsOptions *ListWordsOptions)", func() {
		ListWordsPath := "/v1/customizations/{customization_id}/words"
		CustomizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		ListWordsPath = strings.Replace(ListWordsPath, "{customization_id}", CustomizationID, 1)
		Context("Successfully - List custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(ListWordsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListWords", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				ListWordsOptions := testService.NewListWordsOptions(CustomizationID)
				returnValue, returnValueErr := testService.ListWords(ListWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListWordsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		DeleteUserDataPath := "/v1/user_data"
		CustomerID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Delete labeled data", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(DeleteUserDataPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := texttospeechv1.
					NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
						URL:      testServer.URL,
						Username: username,
						Password: password,
					})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				DeleteUserDataOptions := testService.NewDeleteUserDataOptions(CustomerID)
				returnValue, returnValueErr := testService.DeleteUserData(DeleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
})
