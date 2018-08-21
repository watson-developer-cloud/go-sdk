package speechToTextV1_test

import (
	"github.ibm.com/arf/go-sdk/speechToTextV1"
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
	"io/ioutil"
	"bytes"
)

var _ = Describe("SpeechToTextV1", func() {
	Describe("Get credentials from VCAP", func() {
		version := "exampleString"
		username := "hyphenated-user"
		password := "hyphenated-pass"
		VCAPservices := cfenv.Services{
			"conversation" : {
				{
					Name: "speech_to_text",
					Tags: []string{},
					Credentials: map[string]interface{}{
						"url": "https://stream.watsonplatform.net/speech-to-text/api",
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
		Context("Successfully - Create SpeechToTextV1 with VCAP credentials", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
			}))
			It("Succeed to create SpeechToTextV1", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				testService.ListAcousticModels(speechToTextV1.NewListAcousticModelsOptions())
			})
		})
	})
	Describe("GetModel(options *GetModelOptions)", func() {
		getModelPath := "/v1/models/{model_id}"
        version := "exampleString"
        modelID := "exampleString"
        getModelOptions := speechToTextV1.NewGetModelOptions(modelID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getModelPath = strings.Replace(getModelPath, "{model_id}", modelID, 1)
		Context("Successfully - Get a model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetModel(getModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetGetModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListModels(options *ListModelsOptions)", func() {
		listModelsPath := "/v1/models"
        version := "exampleString"
        listModelsOptions := speechToTextV1.NewListModelsOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListModels(listModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetListModelsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("Recognize(options *RecognizeOptions)", func() {
		recognizePath := "/v1/recognize"
        version := "exampleString"
        audio := ioutil.NopCloser(bytes.NewReader([]byte("exampleString")))
        recognizeOptions := speechToTextV1.NewRecognizeOptionsForMp3(audio)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call Recognize", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.Recognize(recognizeOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetRecognizeResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CheckJob(options *CheckJobOptions)", func() {
		checkJobPath := "/v1/recognitions/{id}"
        version := "exampleString"
        id := "exampleString"
        checkJobOptions := speechToTextV1.NewCheckJobOptions(id)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        checkJobPath = strings.Replace(checkJobPath, "{id}", id, 1)
		Context("Successfully - Check a job", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(checkJobPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CheckJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CheckJob(checkJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetCheckJobResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CheckJobs(options *CheckJobsOptions)", func() {
		checkJobsPath := "/v1/recognitions"
        version := "exampleString"
        checkJobsOptions := speechToTextV1.NewCheckJobsOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CheckJobs", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CheckJobs(checkJobsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetCheckJobsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateJob(options *CreateJobOptions)", func() {
		createJobPath := "/v1/recognitions"
        version := "exampleString"
        audio := ioutil.NopCloser(bytes.NewReader([]byte("exampleString")))
        createJobOptions := speechToTextV1.NewCreateJobOptionsForBasic(audio)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateJob(createJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetCreateJobResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteJob(options *DeleteJobOptions)", func() {
		deleteJobPath := "/v1/recognitions/{id}"
        version := "exampleString"
        id := "exampleString"
        deleteJobOptions := speechToTextV1.NewDeleteJobOptions(id)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteJobPath = strings.Replace(deleteJobPath, "{id}", id, 1)
		Context("Successfully - Delete a job", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteJobPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteJob", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteJob(deleteJobOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("RegisterCallback(options *RegisterCallbackOptions)", func() {
		registerCallbackPath := "/v1/register_callback"
        version := "exampleString"
        callbackURL := "exampleString"
        registerCallbackOptions := speechToTextV1.NewRegisterCallbackOptions(callbackURL)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call RegisterCallback", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.RegisterCallback(registerCallbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetRegisterCallbackResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UnregisterCallback(options *UnregisterCallbackOptions)", func() {
		unregisterCallbackPath := "/v1/unregister_callback"
        version := "exampleString"
        callbackURL := "exampleString"
        unregisterCallbackOptions := speechToTextV1.NewUnregisterCallbackOptions(callbackURL)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UnregisterCallback", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UnregisterCallback(unregisterCallbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("CreateLanguageModel(options *CreateLanguageModelOptions)", func() {
		createLanguageModelPath := "/v1/customizations"
        version := "exampleString"
        name := "exampleString"
        baseModelName := "exampleString"
        createLanguageModelOptions := speechToTextV1.NewCreateLanguageModelOptions(name, baseModelName)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateLanguageModel(createLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetCreateLanguageModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteLanguageModel(options *DeleteLanguageModelOptions)", func() {
		deleteLanguageModelPath := "/v1/customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        deleteLanguageModelOptions := speechToTextV1.NewDeleteLanguageModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteLanguageModelPath = strings.Replace(deleteLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Delete a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteLanguageModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteLanguageModel(deleteLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetLanguageModel(options *GetLanguageModelOptions)", func() {
		getLanguageModelPath := "/v1/customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        getLanguageModelOptions := speechToTextV1.NewGetLanguageModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getLanguageModelPath = strings.Replace(getLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Get a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getLanguageModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetLanguageModel(getLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetGetLanguageModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListLanguageModels(options *ListLanguageModelsOptions)", func() {
		listLanguageModelsPath := "/v1/customizations"
        version := "exampleString"
        listLanguageModelsOptions := speechToTextV1.NewListLanguageModelsOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListLanguageModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListLanguageModels(listLanguageModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetListLanguageModelsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ResetLanguageModel(options *ResetLanguageModelOptions)", func() {
		resetLanguageModelPath := "/v1/customizations/{customization_id}/reset"
        version := "exampleString"
        customizationID := "exampleString"
        resetLanguageModelOptions := speechToTextV1.NewResetLanguageModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        resetLanguageModelPath = strings.Replace(resetLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Reset a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(resetLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ResetLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ResetLanguageModel(resetLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("TrainLanguageModel(options *TrainLanguageModelOptions)", func() {
		trainLanguageModelPath := "/v1/customizations/{customization_id}/train"
        version := "exampleString"
        customizationID := "exampleString"
        trainLanguageModelOptions := speechToTextV1.NewTrainLanguageModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        trainLanguageModelPath = strings.Replace(trainLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Train a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(trainLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call TrainLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.TrainLanguageModel(trainLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("UpgradeLanguageModel(options *UpgradeLanguageModelOptions)", func() {
		upgradeLanguageModelPath := "/v1/customizations/{customization_id}/upgrade_model"
        version := "exampleString"
        customizationID := "exampleString"
        upgradeLanguageModelOptions := speechToTextV1.NewUpgradeLanguageModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        upgradeLanguageModelPath = strings.Replace(upgradeLanguageModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Upgrade a custom language model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(upgradeLanguageModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpgradeLanguageModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpgradeLanguageModel(upgradeLanguageModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteCorpus(options *DeleteCorpusOptions)", func() {
		deleteCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
        version := "exampleString"
        customizationID := "exampleString"
        corpusName := "exampleString"
        deleteCorpusOptions := speechToTextV1.NewDeleteCorpusOptions(customizationID, corpusName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteCorpusPath = strings.Replace(deleteCorpusPath, "{customization_id}", customizationID, 1)
        deleteCorpusPath = strings.Replace(deleteCorpusPath, "{corpus_name}", corpusName, 1)
		Context("Successfully - Delete a corpus", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteCorpusPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteCorpus", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteCorpus(deleteCorpusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetCorpus(options *GetCorpusOptions)", func() {
		getCorpusPath := "/v1/customizations/{customization_id}/corpora/{corpus_name}"
        version := "exampleString"
        customizationID := "exampleString"
        corpusName := "exampleString"
        getCorpusOptions := speechToTextV1.NewGetCorpusOptions(customizationID, corpusName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getCorpusPath = strings.Replace(getCorpusPath, "{customization_id}", customizationID, 1)
        getCorpusPath = strings.Replace(getCorpusPath, "{corpus_name}", corpusName, 1)
		Context("Successfully - Get a corpus", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getCorpusPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetCorpus", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetCorpus(getCorpusOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetGetCorpusResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListCorpora(options *ListCorporaOptions)", func() {
		listCorporaPath := "/v1/customizations/{customization_id}/corpora"
        version := "exampleString"
        customizationID := "exampleString"
        listCorporaOptions := speechToTextV1.NewListCorporaOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listCorporaPath = strings.Replace(listCorporaPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List corpora", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listCorporaPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListCorpora", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListCorpora(listCorporaOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetListCorporaResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddWord(options *AddWordOptions)", func() {
		addWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
        version := "exampleString"
        customizationID := "exampleString"
        wordName := "exampleString"
        addWordOptions := speechToTextV1.NewAddWordOptions(customizationID, wordName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        addWordPath = strings.Replace(addWordPath, "{customization_id}", customizationID, 1)
        addWordPath = strings.Replace(addWordPath, "{word_name}", wordName, 1)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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
        words := []speechToTextV1.CustomWord{}
        addWordsOptions := speechToTextV1.NewAddWordsOptions(customizationID, words)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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
		deleteWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
        version := "exampleString"
        customizationID := "exampleString"
        wordName := "exampleString"
        deleteWordOptions := speechToTextV1.NewDeleteWordOptions(customizationID, wordName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteWordPath = strings.Replace(deleteWordPath, "{customization_id}", customizationID, 1)
        deleteWordPath = strings.Replace(deleteWordPath, "{word_name}", wordName, 1)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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
		getWordPath := "/v1/customizations/{customization_id}/words/{word_name}"
        version := "exampleString"
        customizationID := "exampleString"
        wordName := "exampleString"
        getWordOptions := speechToTextV1.NewGetWordOptions(customizationID, wordName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getWordPath = strings.Replace(getWordPath, "{customization_id}", customizationID, 1)
        getWordPath = strings.Replace(getWordPath, "{word_name}", wordName, 1)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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

                result := speechToTextV1.GetGetWordResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListWords(options *ListWordsOptions)", func() {
		listWordsPath := "/v1/customizations/{customization_id}/words"
        version := "exampleString"
        customizationID := "exampleString"
        listWordsOptions := speechToTextV1.NewListWordsOptions(customizationID)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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

                result := speechToTextV1.GetListWordsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateAcousticModel(options *CreateAcousticModelOptions)", func() {
		createAcousticModelPath := "/v1/acoustic_customizations"
        version := "exampleString"
        name := "exampleString"
        baseModelName := "exampleString"
        createAcousticModelOptions := speechToTextV1.NewCreateAcousticModelOptions(name, baseModelName)
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call CreateAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.CreateAcousticModel(createAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetCreateAcousticModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAcousticModel(options *DeleteAcousticModelOptions)", func() {
		deleteAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        deleteAcousticModelOptions := speechToTextV1.NewDeleteAcousticModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteAcousticModelPath = strings.Replace(deleteAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Delete a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteAcousticModelPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteAcousticModel(deleteAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetAcousticModel(options *GetAcousticModelOptions)", func() {
		getAcousticModelPath := "/v1/acoustic_customizations/{customization_id}"
        version := "exampleString"
        customizationID := "exampleString"
        getAcousticModelOptions := speechToTextV1.NewGetAcousticModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getAcousticModelPath = strings.Replace(getAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Get a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getAcousticModelPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetAcousticModel(getAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetGetAcousticModelResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAcousticModels(options *ListAcousticModelsOptions)", func() {
		listAcousticModelsPath := "/v1/acoustic_customizations"
        version := "exampleString"
        listAcousticModelsOptions := speechToTextV1.NewListAcousticModelsOptions()
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
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListAcousticModels", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListAcousticModels(listAcousticModelsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetListAcousticModelsResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ResetAcousticModel(options *ResetAcousticModelOptions)", func() {
		resetAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/reset"
        version := "exampleString"
        customizationID := "exampleString"
        resetAcousticModelOptions := speechToTextV1.NewResetAcousticModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        resetAcousticModelPath = strings.Replace(resetAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Reset a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(resetAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ResetAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ResetAcousticModel(resetAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("TrainAcousticModel(options *TrainAcousticModelOptions)", func() {
		trainAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/train"
        version := "exampleString"
        customizationID := "exampleString"
        trainAcousticModelOptions := speechToTextV1.NewTrainAcousticModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        trainAcousticModelPath = strings.Replace(trainAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Train a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(trainAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call TrainAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.TrainAcousticModel(trainAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("UpgradeAcousticModel(options *UpgradeAcousticModelOptions)", func() {
		upgradeAcousticModelPath := "/v1/acoustic_customizations/{customization_id}/upgrade_model"
        version := "exampleString"
        customizationID := "exampleString"
        upgradeAcousticModelOptions := speechToTextV1.NewUpgradeAcousticModelOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        upgradeAcousticModelPath = strings.Replace(upgradeAcousticModelPath, "{customization_id}", customizationID, 1)
		Context("Successfully - Upgrade a custom acoustic model", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(upgradeAcousticModelPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call UpgradeAcousticModel", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.UpgradeAcousticModel(upgradeAcousticModelOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("AddAudio(options *AddAudioOptions)", func() {
		addAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
        version := "exampleString"
        customizationID := "exampleString"
        audioName := "exampleString"
		audioResource := ioutil.NopCloser(bytes.NewReader([]byte("exampleString")))
        addAudioOptions := speechToTextV1.NewAddAudioOptionsForBasic(audioResource).
        	SetCustomizationID(customizationID).
        	SetAudioName(audioName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        addAudioPath = strings.Replace(addAudioPath, "{customization_id}", customizationID, 1)
        addAudioPath = strings.Replace(addAudioPath, "{audio_name}", audioName, 1)
		Context("Successfully - Add an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(addAudioPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call AddAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.AddAudio(addAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteAudio(options *DeleteAudioOptions)", func() {
		deleteAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
        version := "exampleString"
        customizationID := "exampleString"
        audioName := "exampleString"
        deleteAudioOptions := speechToTextV1.NewDeleteAudioOptions(customizationID, audioName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        deleteAudioPath = strings.Replace(deleteAudioPath, "{customization_id}", customizationID, 1)
        deleteAudioPath = strings.Replace(deleteAudioPath, "{audio_name}", audioName, 1)
		Context("Successfully - Delete an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(deleteAudioPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call DeleteAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.DeleteAudio(deleteAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())
			})
		})
	})
	Describe("GetAudio(options *GetAudioOptions)", func() {
		getAudioPath := "/v1/acoustic_customizations/{customization_id}/audio/{audio_name}"
        version := "exampleString"
        customizationID := "exampleString"
        audioName := "exampleString"
        getAudioOptions := speechToTextV1.NewGetAudioOptions(customizationID, audioName)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        getAudioPath = strings.Replace(getAudioPath, "{customization_id}", customizationID, 1)
        getAudioPath = strings.Replace(getAudioPath, "{audio_name}", audioName, 1)
		Context("Successfully - Get an audio resource", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(getAudioPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call GetAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.GetAudio(getAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetGetAudioResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListAudio(options *ListAudioOptions)", func() {
		listAudioPath := "/v1/acoustic_customizations/{customization_id}/audio"
        version := "exampleString"
        customizationID := "exampleString"
        listAudioOptions := speechToTextV1.NewListAudioOptions(customizationID)
		username := "user1"
		password := "pass1"
		encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
        listAudioPath = strings.Replace(listAudioPath, "{customization_id}", customizationID, 1)
		Context("Successfully - List audio resources", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				Expect(req.URL.Path).To(Equal(listAudioPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Basic " + encodedBasicAuth))
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"hi":"there"}`)
			}))
			It("Succeed to call ListAudio", func() {
				defer testServer.Close()

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
					ServiceURL: testServer.URL,
					Version: version,
					Username: username,
					Password: password,
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				returnValue, returnValueErr := testService.ListAudio(listAudioOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

                result := speechToTextV1.GetListAudioResult(returnValue)
                Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteUserData(options *DeleteUserDataOptions)", func() {
		deleteUserDataPath := "/v1/user_data"
        version := "exampleString"
        customerID := "exampleString"
        deleteUserDataOptions := speechToTextV1.NewDeleteUserDataOptions(customerID)
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

				testService, testServiceErr := speechToTextV1.NewSpeechToTextV1(&speechToTextV1.ServiceCredentials{
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
