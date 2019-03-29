package comparecomplyv1_test

import (
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/comparecomplyv1"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

var _ = Describe("CompareComplyV1", func() {
	Describe("ConvertToHTML(convertToHTMLOptions *ConvertToHTMLOptions)", func() {
		convertToHTMLPath := "/v1/html_conversion"
		version := "exampleString"
		pwd, _ := os.Getwd()
		testPDF, testPDFErr := os.Open(pwd + "/../resources/contract_A.pdf")
		if testPDFErr != nil {
			fmt.Println(testPDFErr)
		}
		filename := "exampleString"
		Context("Successfully - Convert file to HTML", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()
				if !strings.Contains(req.URL.String(), "/v1/html_conversion") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(convertToHTMLPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(convertToHTMLPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"num_pages":"20"}`)
				}
			}))
			It("Succeed to call ConvertToHTML", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "iamAPiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ConvertToHTML(nil)
				Expect(returnValueErr).NotTo(BeNil())

				convertToHTMLOptions := testService.NewConvertToHTMLOptions(testPDF, filename)
				returnValue, returnValueErr = testService.ConvertToHTML(convertToHTMLOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetConvertToHTMLResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ClassifyElements(classifyElementsOptions *ClassifyElementsOptions)", func() {
		classifyElementsPath := "/v1/element_classification"
		version := "exampleString"
		pwd, _ := os.Getwd()
		testPDF, testPDFErr := os.Open(pwd + "/../resources/contract_A.pdf")
		if testPDFErr != nil {
			fmt.Println(testPDFErr)
		}
		Context("Successfully - Classify the elements of a document", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/element_classification") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(classifyElementsPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(classifyElementsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"model_id":"xxx"}`)
				}
			}))
			It("Succeed to call ClassifyElements", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ClassifyElements(nil)
				Expect(returnValueErr).NotTo(BeNil())

				classifyElementsOptions := testService.NewClassifyElementsOptions(testPDF)
				returnValue, returnValueErr = testService.ClassifyElements(classifyElementsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetClassifyElementsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ExtractTables(extractTablesOptions *ExtractTablesOptions)", func() {
		extractTablesPath := "/v1/tables"
		version := "exampleString"
		pwd, _ := os.Getwd()
		testPDF, testPDFErr := os.Open(pwd + "/../resources/sample-tables.pdf")
		if testPDFErr != nil {
			fmt.Println(testPDFErr)
		}
		Context("Successfully - Extract a document's tables", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/tables") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(extractTablesPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(extractTablesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"model_id":"xxx"}`)
				}
			}))
			It("Succeed to call ExtractTables", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ExtractTables(nil)
				Expect(returnValueErr).NotTo(BeNil())

				extractTablesOptions := testService.NewExtractTablesOptions(testPDF)
				returnValue, returnValueErr = testService.ExtractTables(extractTablesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetExtractTablesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CompareDocuments(compareDocumentsOptions *CompareDocumentsOptions)", func() {
		compareDocumentsPath := "/v1/comparison"
		version := "exampleString"
		pwd, _ := os.Getwd()
		file1, file1Err := os.Open(pwd + "/../resources/contract_A.pdf")
		if file1Err != nil {
			fmt.Println(file1Err)
		}
		file2, file2Err := os.Open(pwd + "/../resources/contract_B.pdf")
		if file2Err != nil {
			fmt.Println(file2Err)
		}
		Context("Successfully - Compare two documents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/comparison") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(compareDocumentsPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(compareDocumentsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"documents":[]}`)
				}
			}))
			It("Succeed to call CompareDocuments", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CompareDocuments(nil)
				Expect(returnValueErr).NotTo(BeNil())

				compareDocumentsOptions := testService.NewCompareDocumentsOptions(file1, file2)
				returnValue, returnValueErr = testService.CompareDocuments(compareDocumentsOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCompareDocumentsResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("AddFeedback(addFeedbackOptions *AddFeedbackOptions)", func() {
		addFeedbackPath := "/v1/feedback"
		version := "exampleString"
		Context("Successfully - Add feedback", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/feedback") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(addFeedbackPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(addFeedbackPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"feedback_id":"xxx"}`)
				}
			}))
			It("Succeed to call AddFeedback", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.AddFeedback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				addFeedbackOptions := testService.NewAddFeedbackOptions(&comparecomplyv1.FeedbackDataInput{
					FeedbackType: core.StringPtr("element_classification"),
					Document: &comparecomplyv1.ShortDoc{
						Title: core.StringPtr("title"),
					},
					ModelID:      core.StringPtr("contracts"),
					ModelVersion: core.StringPtr("11.00"),
					Location: &comparecomplyv1.Location{
						Begin: core.Int64Ptr(214),
						End:   core.Int64Ptr(237),
					},
					Text: core.StringPtr("1. IBM will provide a Senior Managing Consultant / expert resource, for up to 80 hours, to assist Florida Power & Light (FPL) with the creation of an IT infrastructure unit cost model for existing infrastructure."),
					OriginalLabels: &comparecomplyv1.OriginalLabelsIn{
						Types: []comparecomplyv1.TypeLabel{
							comparecomplyv1.TypeLabel{
								Label: &comparecomplyv1.Label{
									Nature: core.StringPtr("Obligation"),
									Party:  core.StringPtr("IBM"),
								},
								ProvenanceIds: []string{"85f5981a-ba91-44f5-9efa-0bd22e64b7bc", "ce0480a1-5ef1-4c3e-9861-3743b5610795"},
							},
						},
						Categories: []comparecomplyv1.Category{
							comparecomplyv1.Category{
								Label: core.StringPtr(comparecomplyv1.Category_Label_Amendments),
							},
						},
					},
					UpdatedLabels: &comparecomplyv1.UpdatedLabelsIn{
						Types: []comparecomplyv1.TypeLabel{
							comparecomplyv1.TypeLabel{
								Label: &comparecomplyv1.Label{
									Nature: core.StringPtr("Obligation"),
									Party:  core.StringPtr("IBM"),
								},
							},
							comparecomplyv1.TypeLabel{
								Label: &comparecomplyv1.Label{
									Nature: core.StringPtr("Disclaimer"),
									Party:  core.StringPtr("Buyer"),
								},
							},
						},
						Categories: []comparecomplyv1.Category{
							comparecomplyv1.Category{
								Label: core.StringPtr("Responsibilities"),
							},
						},
					},
				})
				returnValue, returnValueErr = testService.AddFeedback(addFeedbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetAddFeedbackResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("DeleteFeedback(deleteFeedbackOptions *DeleteFeedbackOptions)", func() {
		deleteFeedbackPath := "/v1/feedback/{feedback_id}"
		version := "exampleString"
		feedbackID := "exampleString"
		Path := strings.Replace(deleteFeedbackPath, "{feedback_id}", feedbackID, 1)
		Context("Successfully - Deletes a specified feedback entry", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/feedback") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
					Expect(req.URL.Path).To(Equal(Path))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"status":"deleted"}`)
				}
			}))
			It("Succeed to call DeleteFeedback", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.DeleteFeedback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				deleteFeedbackOptions := testService.NewDeleteFeedbackOptions(feedbackID)
				returnValue, returnValueErr = testService.DeleteFeedback(deleteFeedbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetDeleteFeedbackResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetFeedback(getFeedbackOptions *GetFeedbackOptions)", func() {
		getFeedbackPath := "/v1/feedback/{feedback_id}"
		version := "exampleString"
		feedbackID := "exampleString"
		Path := strings.Replace(getFeedbackPath, "{feedback_id}", feedbackID, 1)
		Context("Successfully - List a specified feedback entry", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/feedback") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
					Expect(req.URL.Path).To(Equal(Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"feedback_id":"xxx"}`)
				}
			}))
			It("Succeed to call GetFeedback", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetFeedback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getFeedbackOptions := testService.NewGetFeedbackOptions(feedbackID)
				returnValue, returnValueErr = testService.GetFeedback(getFeedbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetFeedbackResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("ListFeedback(listFeedbackOptions *ListFeedbackOptions)", func() {
		listFeedbackPath := "/v1/feedback"
		version := "exampleString"
		Context("Successfully - List the feedback in documents", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/feedback") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(listFeedbackPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(listFeedbackPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"feedback_id":"xxx"}`)
				}
			}))
			It("Succeed to call ListFeedback", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListFeedback(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listFeedbackOptions := testService.NewListFeedbackOptions()
				returnValue, returnValueErr = testService.ListFeedback(listFeedbackOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListFeedbackResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("CreateBatch(createBatchOptions *CreateBatchOptions)", func() {
		createBatchPath := "/v1/batches"
		version := "exampleString"
		function := "exampleString"
		pwd, _ := os.Getwd()
		inputCredentialsFile, inputCredentialsFileErr := os.Open(pwd + "/../resources/dummy-storage-credentials.json")
		if inputCredentialsFileErr != nil {
			fmt.Println(inputCredentialsFileErr)
		}
		inputBucketLocation := "exampleString"
		inputBucketName := "exampleString"
		outputCredentialsFile, outputCredentialsFileErr := os.Open(pwd + "/../resources/dummy-storage-credentials.json")
		if outputCredentialsFileErr != nil {
			fmt.Println(outputCredentialsFileErr)
		}
		outputBucketLocation := "exampleString"
		outputBucketName := "exampleString"
		Context("Successfully - Submit a batch-processing request", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/batches") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(createBatchPath + "?function=" + function + "&version=" + version))
					Expect(req.URL.Path).To(Equal(createBatchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"function":"hello_world"}`)
				}
			}))
			It("Succeed to call CreateBatch", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.CreateBatch(nil)
				Expect(returnValueErr).NotTo(BeNil())

				createBatchOptions := testService.NewCreateBatchOptions(function, inputCredentialsFile, inputBucketLocation, inputBucketName, outputCredentialsFile, outputBucketLocation, outputBucketName)
				returnValue, returnValueErr = testService.CreateBatch(createBatchOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetCreateBatchResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetBatch(getBatchOptions *GetBatchOptions)", func() {
		getBatchPath := "/v1/batches/{batch_id}"
		version := "exampleString"
		batchID := "exampleString"
		Path := strings.Replace(getBatchPath, "{batch_id}", batchID, 1)
		Context("Successfully - Gets information about a specific batch-processing request", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/batches") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(Path + "?version=" + version))
					Expect(req.URL.Path).To(Equal(Path))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"function":"hello_world"}`)
				}
			}))
			It("Succeed to call GetBatch", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.GetBatch(nil)
				Expect(returnValueErr).NotTo(BeNil())

				getBatchOptions := testService.NewGetBatchOptions(batchID)
				returnValue, returnValueErr = testService.GetBatch(getBatchOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetGetBatchResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("GetBatches(getBatchesOptions *GetBatchesOptions)", func() {
		getBatchesPath := "/v1/batches"
		version := "exampleString"
		Context("Successfully - Gets the list of submitted batch-processing jobs", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/batches") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(getBatchesPath + "?version=" + version))
					Expect(req.URL.Path).To(Equal(getBatchesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"function":"hello_world"}`)
				}
			}))
			It("Succeed to call GetBatches", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.ListBatches(nil)
				Expect(returnValueErr).NotTo(BeNil())

				listBatchesOptions := testService.NewListBatchesOptions()
				returnValue, returnValueErr = testService.ListBatches(listBatchesOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetListBatchesResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe("UpdateBatch(updateBatchOptions *UpdateBatchOptions)", func() {
		updateBatchPath := "/v1/batches/{batch_id}"
		version := "exampleString"
		batchID := "exampleString"
		action := "exampleString"
		Path := strings.Replace(updateBatchPath, "{batch_id}", batchID, 1)
		Context("Successfully - Updates a pending or active batch-processing request", func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				if !strings.Contains(req.URL.String(), "/v1/batches") {
					res.WriteHeader(http.StatusOK)
					fmt.Fprintf(res, `{
							"access_token": "oAeisG8yqPY7sFR_x66Z15",
							"token_type": "Bearer",
							"expires_in": 3600,
							"expiration": 1524167011,
							"refresh_token": "jy4gl91BQ"
						}`)
				} else {
					Expect(req.URL.String()).To(Equal(Path + "?action=" + action + "&version=" + version))
					Expect(req.URL.Path).To(Equal(Path))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Authorization"]).ToNot(BeNil())
					Expect(req.Header["Authorization"][0]).To(Equal("Bearer oAeisG8yqPY7sFR_x66Z15"))
					res.Header().Set("Content-type", "application/json")
					fmt.Fprintf(res, `{"function":"hello_world"}`)
				}
			}))
			It("Succeed to call UpdateBatch", func() {
				defer testServer.Close()

				testService, testServiceErr := comparecomplyv1.NewCompareComplyV1(&comparecomplyv1.CompareComplyV1Options{
					URL:       testServer.URL,
					Version:   version,
					IAMURL:    testServer.URL,
					IAMApiKey: "IAMApiKey",
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Pass empty options
				returnValue, returnValueErr := testService.UpdateBatch(nil)
				Expect(returnValueErr).NotTo(BeNil())

				updateBatchOptions := testService.NewUpdateBatchOptions(batchID, action)
				returnValue, returnValueErr = testService.UpdateBatch(updateBatchOptions)
				Expect(returnValueErr).To(BeNil())
				Expect(returnValue).ToNot(BeNil())

				result := testService.GetUpdateBatchResult(returnValue)
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
