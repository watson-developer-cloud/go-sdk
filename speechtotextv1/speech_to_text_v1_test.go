package speechtotextv1_test

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

var _ = Describe("SpeechToTextV1", func() {
	Describe("GetModel(getModelOptions *GetModelOptions)", func() {
		getModelPath := "/v1/models/{model_id}"
		modelID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Get a model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call GetModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getModelOptions := testService.NewGetModelOptions(modelID)
				returnValue, returnValueErr = testService.GetModel(getModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(listModelsOptions *ListModelsOptions)", func() {
		listModelsPath := "/v1/models"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[{"name":"xxx"}]`)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListModels(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listModelsOptions := testService.NewListModelsOptions()
				returnValue, returnValueErr = testService.ListModels(listModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Recognize(recognizeOptions *RecognizeOptions)", func() {
		recognizePath := "/v1/recognize"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Recognize audio", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(recognizePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"results":[{"final": true}]}`)
			}))
			It("Succeed to call Recognize", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.Recognize(nil)
				Expect(returnValueErr).NotTo(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/output.wav")
				if err != nil {
					panic(err)
				}
				recognizeOptions := testService.
					NewRecognizeOptions(file)
				returnValue, returnValueErr = testService.Recognize(recognizeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetRecognizeResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CheckJob(checkJobOptions *CheckJobOptions)", func() {
		checkJobPath := "/v1/recognitions/{id}"
		ID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(checkJobPath, "{id}", ID, 1)
		Context("Successfully - Check a job", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"id":"xxx", "status": "active"}`)
			}))
			It("Succeed to call CheckJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CheckJob(nil)
				Expect(returnValueErr).NotTo(BeNil())

				checkJobOptions := testService.NewCheckJobOptions(ID)
				returnValue, returnValueErr = testService.CheckJob(checkJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCheckJobResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CheckJobs(checkJobsOptions *CheckJobsOptions)", func() {
		checkJobsPath := "/v1/recognitions"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Check jobs", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(checkJobsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call CheckJobs", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CheckJobs(nil)
				Expect(returnValueErr).NotTo(BeNil())

				checkJobsOptions := testService.NewCheckJobsOptions()
				returnValue, returnValueErr = testService.CheckJobs(checkJobsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCheckJobsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateJob(createJobOptions *CreateJobOptions)", func() {
		createJobPath := "/v1/recognitions"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a job", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(createJobPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"id":"xxx", "status":"active"}`)
			}))
			It("Succeed to call CreateJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateJob(nil)
				Expect(returnValueErr).NotTo(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/output.wav")
				if err != nil {
					panic(err)
				}
				createJobOptions := testService.
					NewCreateJobOptions(file).
					SetContentType(speechtotextv1.CreateJobOptions_ContentType_AudioWav)
				returnValue, returnValueErr = testService.CreateJob(createJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateJobResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteJob(deleteJobOptions *DeleteJobOptions)", func() {
		deleteJobPath := "/v1/recognitions/{id}"
		ID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteJobPath, "{id}", ID, 1)
		Context("Successfully - Delete a job", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call DeleteJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteJob(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteJobOptions := testService.NewDeleteJobOptions(ID)
				returnValue, returnValueErr = testService.DeleteJob(deleteJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("RegisterCallback(registerCallbackOptions *RegisterCallbackOptions)", func() {
		registerCallbackPath := "/v1/register_callback"
		callbackURL := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Register a callback", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(registerCallbackPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"status":"active", "url":"www.test.com"}`)
			}))
			It("Succeed to call RegisterCallback", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.RegisterCallback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				registerCallbackOptions := testService.NewRegisterCallbackOptions(callbackURL)
				returnValue, returnValueErr = testService.RegisterCallback(registerCallbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetRegisterCallbackResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UnregisterCallback(unregisterCallbackOptions *UnregisterCallbackOptions)", func() {
		unregisterCallbackPath := "/v1/unregister_callback"
		callbackURL := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Unregister a callback", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(unregisterCallbackPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{}`)
			}))
			It("Succeed to call UnregisterCallback", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UnregisterCallback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				unregisterCallbackOptions := testService.NewUnregisterCallbackOptions(callbackURL)
				returnValue, returnValueErr = testService.UnregisterCallback(unregisterCallbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("CreateLanguageModel(createLanguageModelOptions *CreateLanguageModelOptions)", func() {
		createLanguageModelPath := "/v1/customizations"
		name := "exampleString"
		baseModelName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(createLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call CreateLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createLanguageModelOptions := testService.NewCreateLanguageModelOptions(name, baseModelName)
				returnValue, returnValueErr = testService.CreateLanguageModel(createLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateLanguageModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteLanguageModel(deleteLanguageModelOptions *DeleteLanguageModelOptions)", func() {
		deleteLanguageModelPath := "/v1/customizations/{customization_id}"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Delete a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteLanguageModelOptions := testService.NewDeleteLanguageModelOptions(customizationID)
				returnValue, returnValueErr = testService.DeleteLanguageModel(deleteLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetLanguageModel(getLanguageModelOptions *GetLanguageModelOptions)", func() {
		getLanguageModelPath := "/v1/customizations/{customization_id}"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Get a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call GetLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getLanguageModelOptions := testService.NewGetLanguageModelOptions(customizationID)
				returnValue, returnValueErr = testService.GetLanguageModel(getLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetLanguageModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListLanguageModels(listLanguageModelsOptions *ListLanguageModelsOptions)", func() {
		listLanguageModelsPath := "/v1/customizations"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List custom language models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listLanguageModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListLanguageModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListLanguageModels(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listLanguageModelsOptions := testService.NewListLanguageModelsOptions()
				returnValue, returnValueErr = testService.ListLanguageModels(listLanguageModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListLanguageModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ResetLanguageModel(resetLanguageModelOptions *ResetLanguageModelOptions)", func() {
		resetLanguageModelPath := "/v1/customizations/{customization_id}/reset"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(resetLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Reset a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ResetLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ResetLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				resetLanguageModelOptions := testService.NewResetLanguageModelOptions(customizationID)
				returnValue, returnValueErr = testService.ResetLanguageModel(resetLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("TrainLanguageModel(trainLanguageModelOptions *TrainLanguageModelOptions)", func() {
		trainLanguageModelPath := "/v1/customizations/{customization_id}/train"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(trainLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Train a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call TrainLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.TrainLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				trainLanguageModelOptions := testService.NewTrainLanguageModelOptions(customizationID)
				returnValue, returnValueErr = testService.TrainLanguageModel(trainLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("UpgradeLanguageModel(upgradeLanguageModelOptions *UpgradeLanguageModelOptions)", func() {
		upgradeLanguageModelPath := "/v1/customizations/{customization_id}/upgrade_model"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(upgradeLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Upgrade a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call UpgradeLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpgradeLanguageModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				upgradeLanguageModelOptions := testService.NewUpgradeLanguageModelOptions(customizationID)
				returnValue, returnValueErr = testService.UpgradeLanguageModel(upgradeLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddCorpus(addCorpusOptions *AddCorpusOptions)", func() {
		addCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		customizationID := "exampleString"
		corpusName := "exampleString"
		pwd, _ := os.Getwd()
		corpus, corpusErr := os.Open(pwd + "/../resources/corpus-short-1.txt")
		if corpusErr != nil {
			panic(corpusErr)
		}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(addCorpusPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{corpus_name}", corpusName, 1)
		Context("Successfully - Add a corpus", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call AddCorpus", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddCorpus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addCorpusOptions := testService.NewAddCorpusOptions(customizationID, corpusName, corpus)
				returnValue, returnValueErr = testService.AddCorpus(addCorpusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCorpus(deleteCorpusOptions *DeleteCorpusOptions)", func() {
		deleteCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		customizationID := "exampleString"
		corpusName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteCorpusPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{corpus_name}", corpusName, 1)
		Context("Successfully - Delete a corpus", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteCorpus", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteCorpus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteCorpusOptions := testService.NewDeleteCorpusOptions(customizationID, corpusName)
				returnValue, returnValueErr = testService.DeleteCorpus(deleteCorpusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetCorpus(getCorpusOptions *GetCorpusOptions)", func() {
		getCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
		customizationID := "exampleString"
		corpusName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getCorpusPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{corpus_name}", corpusName, 1)
		Context("Successfully - Get a corpus", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx", "total_words: 22"}`)
			}))
			It("Succeed to call GetCorpus", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetCorpus(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getCorpusOptions := testService.NewGetCorpusOptions(customizationID, corpusName)
				returnValue, returnValueErr = testService.GetCorpus(getCorpusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetCorpusResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCorpora(listCorporaOptions *ListCorporaOptions)", func() {
		listCorporaPath := "/v1/customizations/{customization_id}/corpora"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listCorporaPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List corpora", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListCorpora", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListCorpora(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listCorporaOptions := testService.NewListCorporaOptions(customizationID)
				returnValue, returnValueErr = testService.ListCorpora(listCorporaOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListCorporaResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddWord(addWordOptions *AddWordOptions)", func() {
		addWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		customizationID := "exampleString"
		wordName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(addWordPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{word_name}", wordName, 1)
		Context("Successfully - Add a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call AddWord", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddWord(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addWordOptions := testService.NewAddWordOptions(customizationID, wordName)
				returnValue, returnValueErr = testService.AddWord(addWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddWords(addWordsOptions *AddWordsOptions)", func() {
		addWordsPath := "/v1/customizations/{customization_id}/words"
		customizationID := "exampleString"
		words := []speechtotextv1.CustomWord{}
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(addWordsPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Add custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call AddWords", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddWords(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addWordsOptions := testService.NewAddWordsOptions(customizationID, words)
				returnValue, returnValueErr = testService.AddWords(addWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteWord(deleteWordOptions *DeleteWordOptions)", func() {
		deleteWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		customizationID := "exampleString"
		wordName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteWordPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{word_name}", wordName, 1)
		Context("Successfully - Delete a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteWord", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteWord(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteWordOptions := testService.NewDeleteWordOptions(customizationID, wordName)
				returnValue, returnValueErr = testService.DeleteWord(deleteWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetWord(getWordOptions *GetWordOptions)", func() {
		getWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
		customizationID := "exampleString"
		wordName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getWordPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{word_name}", wordName, 1)
		Context("Successfully - Get a custom word", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"word":"test"}`)
			}))
			It("Succeed to call GetWord", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetWord(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getWordOptions := testService.NewGetWordOptions(customizationID, wordName)
				returnValue, returnValueErr = testService.GetWord(getWordOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetWordResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWords(listWordsOptions *ListWordsOptions)", func() {
		listWordsPath := "/v1/customizations/{customization_id}/words"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listWordsPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List custom words", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListWords", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListWords(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listWordsOptions := testService.NewListWordsOptions(customizationID)
				returnValue, returnValueErr = testService.ListWords(listWordsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListWordsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateAcousticModel(createAcousticModelOptions *CreateAcousticModelOptions)", func() {
		createAcousticModelPath := "/v1/acoustic_customizations"
		name := "exampleString"
		baseModelName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - Create a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(createAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call CreateAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createAcousticModelOptions := testService.NewCreateAcousticModelOptions(name, baseModelName)
				returnValue, returnValueErr = testService.CreateAcousticModel(createAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateAcousticModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAcousticModel(deleteAcousticModelOptions *DeleteAcousticModelOptions)", func() {
		deleteAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Delete a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteAcousticModelOptions := testService.NewDeleteAcousticModelOptions(customizationID)
				returnValue, returnValueErr = testService.DeleteAcousticModel(deleteAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetAcousticModel(getAcousticModelOptions *GetAcousticModelOptions)", func() {
		getAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Get a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"customization_id":"xxx"}`)
			}))
			It("Succeed to call GetAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getAcousticModelOptions := testService.NewGetAcousticModelOptions(customizationID)
				returnValue, returnValueErr = testService.GetAcousticModel(getAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetAcousticModelResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAcousticModels(listAcousticModelsOptions *ListAcousticModelsOptions)", func() {
		listAcousticModelsPath := "/v1/acoustic_customizations"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Context("Successfully - List custom acoustic models", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listAcousticModelsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `[]`)
			}))
			It("Succeed to call ListAcousticModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListAcousticModels(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listAcousticModelsOptions := testService.NewListAcousticModelsOptions()
				returnValue, returnValueErr = testService.ListAcousticModels(listAcousticModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListAcousticModelsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ResetAcousticModel(resetAcousticModelOptions *ResetAcousticModelOptions)", func() {
		resetAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/reset"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(resetAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Reset a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call ResetAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ResetAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				resetAcousticModelOptions := testService.NewResetAcousticModelOptions(customizationID)
				returnValue, returnValueErr = testService.ResetAcousticModel(resetAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("TrainAcousticModel(trainAcousticModelOptions *TrainAcousticModelOptions)", func() {
		trainAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/train"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(trainAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Train a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call TrainAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.TrainAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				trainAcousticModelOptions := testService.NewTrainAcousticModelOptions(customizationID)
				returnValue, returnValueErr = testService.TrainAcousticModel(trainAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("UpgradeAcousticModel(upgradeAcousticModelOptions *UpgradeAcousticModelOptions)", func() {
		upgradeAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/upgrade_model"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(upgradeAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Upgrade a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call UpgradeAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpgradeAcousticModel(nil)
				Expect(returnValueErr).NotTo(BeNil())

				upgradeAcousticModelOptions := testService.NewUpgradeAcousticModelOptions(customizationID)
				returnValue, returnValueErr = testService.UpgradeAcousticModel(upgradeAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddAudio(addAudioOptions *AddAudioOptions)", func() {
		addAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		customizationID := "exampleString"
		audioName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(addAudioPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{audio_name}", audioName, 1)
		Context("Successfully - Add an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call AddAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddAudio(nil)
				Expect(returnValueErr).NotTo(BeNil())

				pwd, _ := os.Getwd()
				file, err := os.Open(pwd + "/../resources/output.wav")
				if err != nil {
					panic(err)
				}
				addAudioOptions := testService.
					NewAddAudioOptions(customizationID,
						audioName,
						file).
					SetContentType(speechtotextv1.AddAudioOptions_ContentType_AudioWav)
				returnValue, returnValueErr = testService.AddAudio(addAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAudio(deleteAudioOptions *DeleteAudioOptions)", func() {
		deleteAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		customizationID := "exampleString"
		audioName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteAudioPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{audio_name}", audioName, 1)
		Context("Successfully - Delete an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
			}))
			It("Succeed to call DeleteAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteAudio(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteAudioOptions := testService.NewDeleteAudioOptions(customizationID, audioName)
				returnValue, returnValueErr = testService.DeleteAudio(deleteAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetAudio(getAudioOptions *GetAudioOptions)", func() {
		getAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
		customizationID := "exampleString"
		audioName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getAudioPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{audio_name}", audioName, 1)
		Context("Successfully - Get an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"duration":2}`)
			}))
			It("Succeed to call GetAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetAudio(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getAudioOptions := testService.NewGetAudioOptions(customizationID, audioName)
				returnValue, returnValueErr = testService.GetAudio(getAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetAudioResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAudio(listAudioOptions *ListAudioOptions)", func() {
		listAudioPath := "/v1/acoustic_customizations/{customization_id}/audio"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listAudioPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List audio resources", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"total_minutes_of_audio":20}`)
			}))
			It("Succeed to call ListAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListAudio(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listAudioOptions := testService.NewListAudioOptions(customizationID)
				returnValue, returnValueErr = testService.ListAudio(listAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListAudioResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
		customerID := "exampleString"
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
			}))
			It("Succeed to call DeleteUserData", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteUserData(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteUserDataOptions := testService.NewDeleteUserDataOptions(customerID)
				returnValue, returnValueErr = testService.DeleteUserData(deleteUserDataOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddGrammar(addGrammarOptions *AddGrammarOptions)", func() {
		addGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		customizationID := "exampleString"
		grammarName := "exampleString"
		pwd, _ := os.Getwd()
		grammarFile, err := os.Open(pwd + "/../resources/confirm-grammar.xml")
		if err != nil {
			panic(err)
		}
		contentType := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(addGrammarPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{grammar_name}", grammarName, 1)
		Context("Successfully - Add a grammar", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddGrammar", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddGrammar(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addGrammarOptions := testService.NewAddGrammarOptions(customizationID, grammarName, grammarFile, contentType)
				returnValue, returnValueErr = testService.AddGrammar(addGrammarOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteGrammar(deleteGrammarOptions *DeleteGrammarOptions)", func() {
		deleteGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		customizationID := "exampleString"
		grammarName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(deleteGrammarPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{grammar_name}", grammarName, 1)
		Context("Successfully - Delete a grammar", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteGrammar", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteGrammar(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteGrammarOptions := testService.NewDeleteGrammarOptions(customizationID, grammarName)
				returnValue, returnValueErr = testService.DeleteGrammar(deleteGrammarOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetGrammar(getGrammarOptions *GetGrammarOptions)", func() {
		getGrammarPath := "/v1/customizations/{customization_id}/grammars/{grammar_name}"
		customizationID := "exampleString"
		grammarName := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(getGrammarPath, "{customization_id}", customizationID, 1)
		Path = strings.Replace(Path, "{grammar_name}", grammarName, 1)
		Context("Successfully - Get a grammar", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `{"name":"xxx"}`)
			}))
			It("Succeed to call GetGrammar", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetGrammar(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getGrammarOptions := testService.NewGetGrammarOptions(customizationID, grammarName)
				returnValue, returnValueErr = testService.GetGrammar(getGrammarOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetGrammarResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListGrammars(listGrammarsOptions *ListGrammarsOptions)", func() {
		listGrammarsPath := "/v1/customizations/{customization_id}/grammars"
		customizationID := "exampleString"
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		Path := strings.Replace(listGrammarsPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List grammars", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(Path))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.Header().Set("Content-type", "application/json")
				fmt.Fprintf(res, `["grammars": {"name":"xxx"}]`)
			}))
			It("Succeed to call ListGrammars", func() {
				defer testServer.Close()

				testService, testServiceErr := speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
					URL:      testServer.URL,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListGrammars(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listGrammarsOptions := testService.NewListGrammarsOptions(customizationID)
				returnValue, returnValueErr = testService.ListGrammars(listGrammarsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListGrammarsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
