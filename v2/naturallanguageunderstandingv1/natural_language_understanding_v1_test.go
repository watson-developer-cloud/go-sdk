/**
 * (C) Copyright IBM Corp. 2018, 2021.
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

package naturallanguageunderstandingv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v2/naturallanguageunderstandingv1"
)

var _ = Describe(`NaturalLanguageUnderstandingV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeFalse())
			naturalLanguageUnderstandingService.DisableSSLVerification()
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL:     "https://naturallanguageunderstandingv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(naturalLanguageUnderstandingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{})
			Expect(naturalLanguageUnderstandingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_UNDERSTANDING_URL":       "https://naturallanguageunderstandingv1/api",
				"NATURAL_LANGUAGE_UNDERSTANDING_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					Version: core.StringPtr(version),
				})
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := naturalLanguageUnderstandingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != naturalLanguageUnderstandingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(naturalLanguageUnderstandingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(naturalLanguageUnderstandingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := naturalLanguageUnderstandingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != naturalLanguageUnderstandingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(naturalLanguageUnderstandingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(naturalLanguageUnderstandingService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					Version: core.StringPtr(version),
				})
				err := naturalLanguageUnderstandingService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := naturalLanguageUnderstandingService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != naturalLanguageUnderstandingService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(naturalLanguageUnderstandingService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(naturalLanguageUnderstandingService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_UNDERSTANDING_URL":       "https://naturallanguageunderstandingv1/api",
				"NATURAL_LANGUAGE_UNDERSTANDING_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageUnderstandingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_UNDERSTANDING_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageUnderstandingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = naturallanguageunderstandingv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Analyze(analyzeOptions *AnalyzeOptions) - Operation response error`, func() {
		version := "testString"
		analyzePath := "/v1/analyze"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Analyze with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				classificationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				relationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(naturallanguageunderstandingv1.AnalyzeOptions)
				analyzeOptionsModel.Features = featuresModel
				analyzeOptionsModel.Text = core.StringPtr("testString")
				analyzeOptionsModel.HTML = core.StringPtr("testString")
				analyzeOptionsModel.URL = core.StringPtr("testString")
				analyzeOptionsModel.Clean = core.BoolPtr(true)
				analyzeOptionsModel.Xpath = core.StringPtr("testString")
				analyzeOptionsModel.FallbackToRaw = core.BoolPtr(true)
				analyzeOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				analyzeOptionsModel.Language = core.StringPtr("testString")
				analyzeOptionsModel.LimitTextCharacters = core.Int64Ptr(int64(38))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Analyze(analyzeOptions *AnalyzeOptions)`, func() {
		version := "testString"
		analyzePath := "/v1/analyze"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "analyzed_text": "AnalyzedText", "retrieved_url": "RetrievedURL", "usage": {"features": 8, "text_characters": 14, "text_units": 9}, "concepts": [{"text": "Text", "relevance": 9, "dbpedia_resource": "DbpediaResource"}], "entities": [{"type": "Type", "text": "Text", "relevance": 9, "confidence": 10, "mentions": [{"text": "Text", "location": [8], "confidence": 10}], "count": 5, "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}, "disambiguation": {"name": "Name", "dbpedia_resource": "DbpediaResource", "subtype": ["Subtype"]}}], "keywords": [{"count": 5, "relevance": 9, "text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}}], "categories": [{"label": "Label", "score": 5, "explanation": {"relevant_text": [{"text": "Text"}]}}], "classifications": [{"class_name": "ClassName", "confidence": 10}], "emotion": {"document": {"emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}, "targets": [{"text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}]}, "metadata": {"authors": [{"name": "Name"}], "publication_date": "PublicationDate", "title": "Title", "image": "Image", "feeds": [{"link": "Link"}]}, "relations": [{"score": 5, "sentence": "Sentence", "type": "Type", "arguments": [{"entities": [{"text": "Text", "type": "Type"}], "location": [8], "text": "Text"}]}], "semantic_roles": [{"sentence": "Sentence", "subject": {"text": "Text", "entities": [{"type": "Type", "text": "Text"}], "keywords": [{"text": "Text"}]}, "action": {"text": "Text", "normalized": "Normalized", "verb": {"text": "Text", "tense": "Tense"}}, "object": {"text": "Text", "keywords": [{"text": "Text"}]}}], "sentiment": {"document": {"label": "Label", "score": 5}, "targets": [{"text": "Text", "score": 5}]}, "syntax": {"tokens": [{"text": "Text", "part_of_speech": "ADJ", "location": [8], "lemma": "Lemma"}], "sentences": [{"text": "Text", "location": [8]}]}}`)
				}))
			})
			It(`Invoke Analyze successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				classificationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				relationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(naturallanguageunderstandingv1.AnalyzeOptions)
				analyzeOptionsModel.Features = featuresModel
				analyzeOptionsModel.Text = core.StringPtr("testString")
				analyzeOptionsModel.HTML = core.StringPtr("testString")
				analyzeOptionsModel.URL = core.StringPtr("testString")
				analyzeOptionsModel.Clean = core.BoolPtr(true)
				analyzeOptionsModel.Xpath = core.StringPtr("testString")
				analyzeOptionsModel.FallbackToRaw = core.BoolPtr(true)
				analyzeOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				analyzeOptionsModel.Language = core.StringPtr("testString")
				analyzeOptionsModel.LimitTextCharacters = core.Int64Ptr(int64(38))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(analyzePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "analyzed_text": "AnalyzedText", "retrieved_url": "RetrievedURL", "usage": {"features": 8, "text_characters": 14, "text_units": 9}, "concepts": [{"text": "Text", "relevance": 9, "dbpedia_resource": "DbpediaResource"}], "entities": [{"type": "Type", "text": "Text", "relevance": 9, "confidence": 10, "mentions": [{"text": "Text", "location": [8], "confidence": 10}], "count": 5, "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}, "disambiguation": {"name": "Name", "dbpedia_resource": "DbpediaResource", "subtype": ["Subtype"]}}], "keywords": [{"count": 5, "relevance": 9, "text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}}], "categories": [{"label": "Label", "score": 5, "explanation": {"relevant_text": [{"text": "Text"}]}}], "classifications": [{"class_name": "ClassName", "confidence": 10}], "emotion": {"document": {"emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}, "targets": [{"text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}]}, "metadata": {"authors": [{"name": "Name"}], "publication_date": "PublicationDate", "title": "Title", "image": "Image", "feeds": [{"link": "Link"}]}, "relations": [{"score": 5, "sentence": "Sentence", "type": "Type", "arguments": [{"entities": [{"text": "Text", "type": "Type"}], "location": [8], "text": "Text"}]}], "semantic_roles": [{"sentence": "Sentence", "subject": {"text": "Text", "entities": [{"type": "Type", "text": "Text"}], "keywords": [{"text": "Text"}]}, "action": {"text": "Text", "normalized": "Normalized", "verb": {"text": "Text", "tense": "Tense"}}, "object": {"text": "Text", "keywords": [{"text": "Text"}]}}], "sentiment": {"document": {"label": "Label", "score": 5}, "targets": [{"text": "Text", "score": 5}]}, "syntax": {"tokens": [{"text": "Text", "part_of_speech": "ADJ", "location": [8], "lemma": "Lemma"}], "sentences": [{"text": "Text", "location": [8]}]}}`)
				}))
			})
			It(`Invoke Analyze successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				classificationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				relationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(naturallanguageunderstandingv1.AnalyzeOptions)
				analyzeOptionsModel.Features = featuresModel
				analyzeOptionsModel.Text = core.StringPtr("testString")
				analyzeOptionsModel.HTML = core.StringPtr("testString")
				analyzeOptionsModel.URL = core.StringPtr("testString")
				analyzeOptionsModel.Clean = core.BoolPtr(true)
				analyzeOptionsModel.Xpath = core.StringPtr("testString")
				analyzeOptionsModel.FallbackToRaw = core.BoolPtr(true)
				analyzeOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				analyzeOptionsModel.Language = core.StringPtr("testString")
				analyzeOptionsModel.LimitTextCharacters = core.Int64Ptr(int64(38))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Analyze with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				classificationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				relationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(naturallanguageunderstandingv1.AnalyzeOptions)
				analyzeOptionsModel.Features = featuresModel
				analyzeOptionsModel.Text = core.StringPtr("testString")
				analyzeOptionsModel.HTML = core.StringPtr("testString")
				analyzeOptionsModel.URL = core.StringPtr("testString")
				analyzeOptionsModel.Clean = core.BoolPtr(true)
				analyzeOptionsModel.Xpath = core.StringPtr("testString")
				analyzeOptionsModel.FallbackToRaw = core.BoolPtr(true)
				analyzeOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				analyzeOptionsModel.Language = core.StringPtr("testString")
				analyzeOptionsModel.LimitTextCharacters = core.Int64Ptr(int64(38))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AnalyzeOptions model with no property values
				analyzeOptionsModelNew := new(naturallanguageunderstandingv1.AnalyzeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.Analyze(analyzeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Analyze successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				classificationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				relationsOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel

				// Construct an instance of the AnalyzeOptions model
				analyzeOptionsModel := new(naturallanguageunderstandingv1.AnalyzeOptions)
				analyzeOptionsModel.Features = featuresModel
				analyzeOptionsModel.Text = core.StringPtr("testString")
				analyzeOptionsModel.HTML = core.StringPtr("testString")
				analyzeOptionsModel.URL = core.StringPtr("testString")
				analyzeOptionsModel.Clean = core.BoolPtr(true)
				analyzeOptionsModel.Xpath = core.StringPtr("testString")
				analyzeOptionsModel.FallbackToRaw = core.BoolPtr(true)
				analyzeOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				analyzeOptionsModel.Language = core.StringPtr("testString")
				analyzeOptionsModel.LimitTextCharacters = core.Int64Ptr(int64(38))
				analyzeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions) - Operation response error`, func() {
		version := "testString"
		listModelsPath := "/v1/models"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListModels with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(naturallanguageunderstandingv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions)`, func() {
		version := "testString"
		listModelsPath := "/v1/models"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"status": "starting", "model_id": "ModelID", "language": "Language", "description": "Description", "workspace_id": "WorkspaceID", "model_version": "ModelVersion", "version": "Version", "version_description": "VersionDescription", "created": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListModels successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(naturallanguageunderstandingv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"status": "starting", "model_id": "ModelID", "language": "Language", "description": "Description", "workspace_id": "WorkspaceID", "model_version": "ModelVersion", "version": "Version", "version_description": "VersionDescription", "created": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.ListModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(naturallanguageunderstandingv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListModels with error: Operation request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(naturallanguageunderstandingv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := new(naturallanguageunderstandingv1.ListModelsOptions)
				listModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteModelPath := "/v1/models/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(naturallanguageunderstandingv1.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions)`, func() {
		version := "testString"
		deleteModelPath := "/v1/models/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(naturallanguageunderstandingv1.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(naturallanguageunderstandingv1.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(naturallanguageunderstandingv1.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteModelOptions model with no property values
				deleteModelOptionsModelNew := new(naturallanguageunderstandingv1.DeleteModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteModelOptions model
				deleteModelOptionsModel := new(naturallanguageunderstandingv1.DeleteModelOptions)
				deleteModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSentimentModel(createSentimentModelOptions *CreateSentimentModelOptions) - Operation response error`, func() {
		version := "testString"
		createSentimentModelPath := "/v1/models/sentiment"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSentimentModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSentimentModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateSentimentModelOptions model
				createSentimentModelOptionsModel := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				createSentimentModelOptionsModel.Language = core.StringPtr("testString")
				createSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel.Name = core.StringPtr("testString")
				createSentimentModelOptionsModel.Description = core.StringPtr("testString")
				createSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSentimentModel(createSentimentModelOptions *CreateSentimentModelOptions)`, func() {
		version := "testString"
		createSentimentModelPath := "/v1/models/sentiment"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSentimentModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke CreateSentimentModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the CreateSentimentModelOptions model
				createSentimentModelOptionsModel := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				createSentimentModelOptionsModel.Language = core.StringPtr("testString")
				createSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel.Name = core.StringPtr("testString")
				createSentimentModelOptionsModel.Description = core.StringPtr("testString")
				createSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.CreateSentimentModelWithContext(ctx, createSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.CreateSentimentModelWithContext(ctx, createSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSentimentModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke CreateSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.CreateSentimentModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSentimentModelOptions model
				createSentimentModelOptionsModel := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				createSentimentModelOptionsModel.Language = core.StringPtr("testString")
				createSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel.Name = core.StringPtr("testString")
				createSentimentModelOptionsModel.Description = core.StringPtr("testString")
				createSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSentimentModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateSentimentModelOptions model
				createSentimentModelOptionsModel := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				createSentimentModelOptionsModel.Language = core.StringPtr("testString")
				createSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel.Name = core.StringPtr("testString")
				createSentimentModelOptionsModel.Description = core.StringPtr("testString")
				createSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSentimentModelOptions model with no property values
				createSentimentModelOptionsModelNew := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateSentimentModelOptions model
				createSentimentModelOptionsModel := new(naturallanguageunderstandingv1.CreateSentimentModelOptions)
				createSentimentModelOptionsModel.Language = core.StringPtr("testString")
				createSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel.Name = core.StringPtr("testString")
				createSentimentModelOptionsModel.Description = core.StringPtr("testString")
				createSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.CreateSentimentModel(createSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSentimentModels(listSentimentModelsOptions *ListSentimentModelsOptions) - Operation response error`, func() {
		version := "testString"
		listSentimentModelsPath := "/v1/models/sentiment"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSentimentModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSentimentModels with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := new(naturallanguageunderstandingv1.ListSentimentModelsOptions)
				listSentimentModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSentimentModels(listSentimentModelsOptions *ListSentimentModelsOptions)`, func() {
		version := "testString"
		listSentimentModelsPath := "/v1/models/sentiment"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSentimentModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}]}`)
				}))
			})
			It(`Invoke ListSentimentModels successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := new(naturallanguageunderstandingv1.ListSentimentModelsOptions)
				listSentimentModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.ListSentimentModelsWithContext(ctx, listSentimentModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.ListSentimentModelsWithContext(ctx, listSentimentModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSentimentModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}]}`)
				}))
			})
			It(`Invoke ListSentimentModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.ListSentimentModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := new(naturallanguageunderstandingv1.ListSentimentModelsOptions)
				listSentimentModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSentimentModels with error: Operation request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := new(naturallanguageunderstandingv1.ListSentimentModelsOptions)
				listSentimentModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSentimentModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := new(naturallanguageunderstandingv1.ListSentimentModelsOptions)
				listSentimentModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.ListSentimentModels(listSentimentModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSentimentModel(getSentimentModelOptions *GetSentimentModelOptions) - Operation response error`, func() {
		version := "testString"
		getSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSentimentModelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSentimentModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetSentimentModelOptions model
				getSentimentModelOptionsModel := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				getSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				getSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSentimentModel(getSentimentModelOptions *GetSentimentModelOptions)`, func() {
		version := "testString"
		getSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSentimentModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke GetSentimentModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the GetSentimentModelOptions model
				getSentimentModelOptionsModel := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				getSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				getSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.GetSentimentModelWithContext(ctx, getSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.GetSentimentModelWithContext(ctx, getSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSentimentModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke GetSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.GetSentimentModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSentimentModelOptions model
				getSentimentModelOptionsModel := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				getSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				getSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSentimentModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetSentimentModelOptions model
				getSentimentModelOptionsModel := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				getSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				getSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSentimentModelOptions model with no property values
				getSentimentModelOptionsModelNew := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetSentimentModelOptions model
				getSentimentModelOptionsModel := new(naturallanguageunderstandingv1.GetSentimentModelOptions)
				getSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				getSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.GetSentimentModel(getSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSentimentModel(updateSentimentModelOptions *UpdateSentimentModelOptions) - Operation response error`, func() {
		version := "testString"
		updateSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSentimentModelPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSentimentModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateSentimentModelOptions model
				updateSentimentModelOptionsModel := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				updateSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Language = core.StringPtr("testString")
				updateSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel.Name = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Description = core.StringPtr("testString")
				updateSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSentimentModel(updateSentimentModelOptions *UpdateSentimentModelOptions)`, func() {
		version := "testString"
		updateSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSentimentModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke UpdateSentimentModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the UpdateSentimentModelOptions model
				updateSentimentModelOptionsModel := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				updateSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Language = core.StringPtr("testString")
				updateSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel.Name = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Description = core.StringPtr("testString")
				updateSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModelWithContext(ctx, updateSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.UpdateSentimentModelWithContext(ctx, updateSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSentimentModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z", "name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "notices": [{"message": "Message"}], "workspace_id": "WorkspaceID", "version_description": "VersionDescription"}`)
				}))
			})
			It(`Invoke UpdateSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateSentimentModelOptions model
				updateSentimentModelOptionsModel := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				updateSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Language = core.StringPtr("testString")
				updateSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel.Name = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Description = core.StringPtr("testString")
				updateSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSentimentModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateSentimentModelOptions model
				updateSentimentModelOptionsModel := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				updateSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Language = core.StringPtr("testString")
				updateSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel.Name = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Description = core.StringPtr("testString")
				updateSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSentimentModelOptions model with no property values
				updateSentimentModelOptionsModelNew := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateSentimentModelOptions model
				updateSentimentModelOptionsModel := new(naturallanguageunderstandingv1.UpdateSentimentModelOptions)
				updateSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Language = core.StringPtr("testString")
				updateSentimentModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel.Name = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Description = core.StringPtr("testString")
				updateSentimentModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateSentimentModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateSentimentModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateSentimentModel(updateSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSentimentModel(deleteSentimentModelOptions *DeleteSentimentModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSentimentModelPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteSentimentModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteSentimentModelOptions model
				deleteSentimentModelOptionsModel := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				deleteSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSentimentModel(deleteSentimentModelOptions *DeleteSentimentModelOptions)`, func() {
		version := "testString"
		deleteSentimentModelPath := "/v1/models/sentiment/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSentimentModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteSentimentModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteSentimentModelOptions model
				deleteSentimentModelOptionsModel := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				deleteSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModelWithContext(ctx, deleteSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteSentimentModelWithContext(ctx, deleteSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSentimentModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteSentimentModelOptions model
				deleteSentimentModelOptionsModel := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				deleteSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteSentimentModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteSentimentModelOptions model
				deleteSentimentModelOptionsModel := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				deleteSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteSentimentModelOptions model with no property values
				deleteSentimentModelOptionsModelNew := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteSentimentModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteSentimentModelOptions model
				deleteSentimentModelOptionsModel := new(naturallanguageunderstandingv1.DeleteSentimentModelOptions)
				deleteSentimentModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteSentimentModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteSentimentModel(deleteSentimentModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCategoriesModel(createCategoriesModelOptions *CreateCategoriesModelOptions) - Operation response error`, func() {
		version := "testString"
		createCategoriesModelPath := "/v1/models/categories"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCategoriesModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCategoriesModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateCategoriesModelOptions model
				createCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				createCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				createCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				createCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCategoriesModel(createCategoriesModelOptions *CreateCategoriesModelOptions)`, func() {
		version := "testString"
		createCategoriesModelPath := "/v1/models/categories"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCategoriesModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCategoriesModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the CreateCategoriesModelOptions model
				createCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				createCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				createCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				createCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModelWithContext(ctx, createCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.CreateCategoriesModelWithContext(ctx, createCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCategoriesModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCategoriesModelOptions model
				createCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				createCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				createCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				createCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCategoriesModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateCategoriesModelOptions model
				createCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				createCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				createCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				createCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCategoriesModelOptions model with no property values
				createCategoriesModelOptionsModelNew := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateCategoriesModelOptions model
				createCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.CreateCategoriesModelOptions)
				createCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				createCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				createCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.CreateCategoriesModel(createCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCategoriesModels(listCategoriesModelsOptions *ListCategoriesModelsOptions) - Operation response error`, func() {
		version := "testString"
		listCategoriesModelsPath := "/v1/models/categories"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCategoriesModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCategoriesModels with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := new(naturallanguageunderstandingv1.ListCategoriesModelsOptions)
				listCategoriesModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCategoriesModels(listCategoriesModelsOptions *ListCategoriesModelsOptions)`, func() {
		version := "testString"
		listCategoriesModelsPath := "/v1/models/categories"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCategoriesModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListCategoriesModels successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := new(naturallanguageunderstandingv1.ListCategoriesModelsOptions)
				listCategoriesModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.ListCategoriesModelsWithContext(ctx, listCategoriesModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.ListCategoriesModelsWithContext(ctx, listCategoriesModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCategoriesModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListCategoriesModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.ListCategoriesModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := new(naturallanguageunderstandingv1.ListCategoriesModelsOptions)
				listCategoriesModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCategoriesModels with error: Operation request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := new(naturallanguageunderstandingv1.ListCategoriesModelsOptions)
				listCategoriesModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCategoriesModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := new(naturallanguageunderstandingv1.ListCategoriesModelsOptions)
				listCategoriesModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.ListCategoriesModels(listCategoriesModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCategoriesModel(getCategoriesModelOptions *GetCategoriesModelOptions) - Operation response error`, func() {
		version := "testString"
		getCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCategoriesModelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCategoriesModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetCategoriesModelOptions model
				getCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				getCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				getCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCategoriesModel(getCategoriesModelOptions *GetCategoriesModelOptions)`, func() {
		version := "testString"
		getCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCategoriesModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCategoriesModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the GetCategoriesModelOptions model
				getCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				getCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				getCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.GetCategoriesModelWithContext(ctx, getCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.GetCategoriesModelWithContext(ctx, getCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCategoriesModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.GetCategoriesModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCategoriesModelOptions model
				getCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				getCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				getCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCategoriesModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetCategoriesModelOptions model
				getCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				getCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				getCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCategoriesModelOptions model with no property values
				getCategoriesModelOptionsModelNew := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetCategoriesModelOptions model
				getCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.GetCategoriesModelOptions)
				getCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				getCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.GetCategoriesModel(getCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCategoriesModel(updateCategoriesModelOptions *UpdateCategoriesModelOptions) - Operation response error`, func() {
		version := "testString"
		updateCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCategoriesModelPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCategoriesModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateCategoriesModelOptions model
				updateCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				updateCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCategoriesModel(updateCategoriesModelOptions *UpdateCategoriesModelOptions)`, func() {
		version := "testString"
		updateCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCategoriesModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCategoriesModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the UpdateCategoriesModelOptions model
				updateCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				updateCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModelWithContext(ctx, updateCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.UpdateCategoriesModelWithContext(ctx, updateCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCategoriesModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateCategoriesModelOptions model
				updateCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				updateCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCategoriesModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateCategoriesModelOptions model
				updateCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				updateCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCategoriesModelOptions model with no property values
				updateCategoriesModelOptionsModelNew := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateCategoriesModelOptions model
				updateCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.UpdateCategoriesModelOptions)
				updateCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Language = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateCategoriesModelOptionsModel.Name = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Description = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateCategoriesModel(updateCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCategoriesModel(deleteCategoriesModelOptions *DeleteCategoriesModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCategoriesModelPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCategoriesModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteCategoriesModelOptions model
				deleteCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				deleteCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCategoriesModel(deleteCategoriesModelOptions *DeleteCategoriesModelOptions)`, func() {
		version := "testString"
		deleteCategoriesModelPath := "/v1/models/categories/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCategoriesModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteCategoriesModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteCategoriesModelOptions model
				deleteCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				deleteCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModelWithContext(ctx, deleteCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteCategoriesModelWithContext(ctx, deleteCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCategoriesModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCategoriesModelOptions model
				deleteCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				deleteCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteCategoriesModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteCategoriesModelOptions model
				deleteCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				deleteCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCategoriesModelOptions model with no property values
				deleteCategoriesModelOptionsModelNew := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteCategoriesModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteCategoriesModelOptions model
				deleteCategoriesModelOptionsModel := new(naturallanguageunderstandingv1.DeleteCategoriesModelOptions)
				deleteCategoriesModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteCategoriesModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteCategoriesModel(deleteCategoriesModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateClassificationsModel(createClassificationsModelOptions *CreateClassificationsModelOptions) - Operation response error`, func() {
		version := "testString"
		createClassificationsModelPath := "/v1/models/classifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassificationsModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateClassificationsModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateClassificationsModelOptions model
				createClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				createClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				createClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				createClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateClassificationsModel(createClassificationsModelOptions *CreateClassificationsModelOptions)`, func() {
		version := "testString"
		createClassificationsModelPath := "/v1/models/classifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassificationsModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateClassificationsModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the CreateClassificationsModelOptions model
				createClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				createClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				createClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				createClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModelWithContext(ctx, createClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.CreateClassificationsModelWithContext(ctx, createClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createClassificationsModelPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke CreateClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateClassificationsModelOptions model
				createClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				createClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				createClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				createClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateClassificationsModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateClassificationsModelOptions model
				createClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				createClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				createClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				createClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateClassificationsModelOptions model with no property values
				createClassificationsModelOptionsModelNew := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the CreateClassificationsModelOptions model
				createClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.CreateClassificationsModelOptions)
				createClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				createClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				createClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				createClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				createClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				createClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				createClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.CreateClassificationsModel(createClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClassificationsModels(listClassificationsModelsOptions *ListClassificationsModelsOptions) - Operation response error`, func() {
		version := "testString"
		listClassificationsModelsPath := "/v1/models/classifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassificationsModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListClassificationsModels with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := new(naturallanguageunderstandingv1.ListClassificationsModelsOptions)
				listClassificationsModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListClassificationsModels(listClassificationsModelsOptions *ListClassificationsModelsOptions)`, func() {
		version := "testString"
		listClassificationsModelsPath := "/v1/models/classifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassificationsModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"models": [{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke ListClassificationsModels successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := new(naturallanguageunderstandingv1.ListClassificationsModelsOptions)
				listClassificationsModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.ListClassificationsModelsWithContext(ctx, listClassificationsModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.ListClassificationsModelsWithContext(ctx, listClassificationsModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listClassificationsModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"models": [{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}]}]}`)
				}))
			})
			It(`Invoke ListClassificationsModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.ListClassificationsModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := new(naturallanguageunderstandingv1.ListClassificationsModelsOptions)
				listClassificationsModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListClassificationsModels with error: Operation request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := new(naturallanguageunderstandingv1.ListClassificationsModelsOptions)
				listClassificationsModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListClassificationsModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := new(naturallanguageunderstandingv1.ListClassificationsModelsOptions)
				listClassificationsModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.ListClassificationsModels(listClassificationsModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetClassificationsModel(getClassificationsModelOptions *GetClassificationsModelOptions) - Operation response error`, func() {
		version := "testString"
		getClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassificationsModelPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetClassificationsModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetClassificationsModelOptions model
				getClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				getClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				getClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetClassificationsModel(getClassificationsModelOptions *GetClassificationsModelOptions)`, func() {
		version := "testString"
		getClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassificationsModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetClassificationsModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the GetClassificationsModelOptions model
				getClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				getClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				getClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.GetClassificationsModelWithContext(ctx, getClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.GetClassificationsModelWithContext(ctx, getClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClassificationsModelPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke GetClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.GetClassificationsModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClassificationsModelOptions model
				getClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				getClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				getClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetClassificationsModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetClassificationsModelOptions model
				getClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				getClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				getClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClassificationsModelOptions model with no property values
				getClassificationsModelOptionsModelNew := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the GetClassificationsModelOptions model
				getClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.GetClassificationsModelOptions)
				getClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				getClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.GetClassificationsModel(getClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateClassificationsModel(updateClassificationsModelOptions *UpdateClassificationsModelOptions) - Operation response error`, func() {
		version := "testString"
		updateClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClassificationsModelPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateClassificationsModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateClassificationsModelOptions model
				updateClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				updateClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateClassificationsModel(updateClassificationsModelOptions *UpdateClassificationsModelOptions)`, func() {
		version := "testString"
		updateClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClassificationsModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateClassificationsModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the UpdateClassificationsModelOptions model
				updateClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				updateClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModelWithContext(ctx, updateClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.UpdateClassificationsModelWithContext(ctx, updateClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateClassificationsModelPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "user_metadata": {"mapKey": {"anyKey": "anyValue"}}, "language": "Language", "description": "Description", "model_version": "ModelVersion", "workspace_id": "WorkspaceID", "version_description": "VersionDescription", "features": ["Features"], "status": "starting", "model_id": "ModelID", "created": "2019-01-01T12:00:00.000Z", "notices": [{"message": "Message"}], "last_trained": "2019-01-01T12:00:00.000Z", "last_deployed": "2019-01-01T12:00:00.000Z"}`)
				}))
			})
			It(`Invoke UpdateClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateClassificationsModelOptions model
				updateClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				updateClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateClassificationsModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateClassificationsModelOptions model
				updateClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				updateClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateClassificationsModelOptions model with no property values
				updateClassificationsModelOptionsModelNew := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the UpdateClassificationsModelOptions model
				updateClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.UpdateClassificationsModelOptions)
				updateClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Language = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.TrainingData = CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel.TrainingDataContentType = core.StringPtr("json")
				updateClassificationsModelOptionsModel.Name = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Description = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.ModelVersion = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.WorkspaceID = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.VersionDescription = core.StringPtr("testString")
				updateClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.UpdateClassificationsModel(updateClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteClassificationsModel(deleteClassificationsModelOptions *DeleteClassificationsModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClassificationsModelPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteClassificationsModel with error: Operation response processing error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteClassificationsModelOptions model
				deleteClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				deleteClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				naturalLanguageUnderstandingService.EnableRetries(0, 0)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteClassificationsModel(deleteClassificationsModelOptions *DeleteClassificationsModelOptions)`, func() {
		version := "testString"
		deleteClassificationsModelPath := "/v1/models/classifications/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClassificationsModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteClassificationsModel successfully with retries`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Construct an instance of the DeleteClassificationsModelOptions model
				deleteClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				deleteClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModelWithContext(ctx, deleteClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteClassificationsModelWithContext(ctx, deleteClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteClassificationsModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"deleted": "Deleted"}`)
				}))
			})
			It(`Invoke DeleteClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteClassificationsModelOptions model
				deleteClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				deleteClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteClassificationsModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteClassificationsModelOptions model
				deleteClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				deleteClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := naturalLanguageUnderstandingService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteClassificationsModelOptions model with no property values
				deleteClassificationsModelOptionsModelNew := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteClassificationsModel successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

				// Construct an instance of the DeleteClassificationsModelOptions model
				deleteClassificationsModelOptionsModel := new(naturallanguageunderstandingv1.DeleteClassificationsModelOptions)
				deleteClassificationsModelOptionsModel.ModelID = core.StringPtr("testString")
				deleteClassificationsModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := naturalLanguageUnderstandingService.DeleteClassificationsModel(deleteClassificationsModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			naturalLanguageUnderstandingService, _ := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL:           "http://naturallanguageunderstandingv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAnalyzeOptions successfully`, func() {
				// Construct an instance of the ClassificationsOptions model
				classificationsOptionsModel := new(naturallanguageunderstandingv1.ClassificationsOptions)
				Expect(classificationsOptionsModel).ToNot(BeNil())
				classificationsOptionsModel.Model = core.StringPtr("testString")
				Expect(classificationsOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ConceptsOptions model
				conceptsOptionsModel := new(naturallanguageunderstandingv1.ConceptsOptions)
				Expect(conceptsOptionsModel).ToNot(BeNil())
				conceptsOptionsModel.Limit = core.Int64Ptr(int64(50))
				Expect(conceptsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))

				// Construct an instance of the EmotionOptions model
				emotionOptionsModel := new(naturallanguageunderstandingv1.EmotionOptions)
				Expect(emotionOptionsModel).ToNot(BeNil())
				emotionOptionsModel.Document = core.BoolPtr(true)
				emotionOptionsModel.Targets = []string{"testString"}
				Expect(emotionOptionsModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(emotionOptionsModel.Targets).To(Equal([]string{"testString"}))

				// Construct an instance of the EntitiesOptions model
				entitiesOptionsModel := new(naturallanguageunderstandingv1.EntitiesOptions)
				Expect(entitiesOptionsModel).ToNot(BeNil())
				entitiesOptionsModel.Limit = core.Int64Ptr(int64(250))
				entitiesOptionsModel.Mentions = core.BoolPtr(true)
				entitiesOptionsModel.Model = core.StringPtr("testString")
				entitiesOptionsModel.Sentiment = core.BoolPtr(true)
				entitiesOptionsModel.Emotion = core.BoolPtr(true)
				Expect(entitiesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(250))))
				Expect(entitiesOptionsModel.Mentions).To(Equal(core.BoolPtr(true)))
				Expect(entitiesOptionsModel.Model).To(Equal(core.StringPtr("testString")))
				Expect(entitiesOptionsModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(entitiesOptionsModel.Emotion).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the KeywordsOptions model
				keywordsOptionsModel := new(naturallanguageunderstandingv1.KeywordsOptions)
				Expect(keywordsOptionsModel).ToNot(BeNil())
				keywordsOptionsModel.Limit = core.Int64Ptr(int64(250))
				keywordsOptionsModel.Sentiment = core.BoolPtr(true)
				keywordsOptionsModel.Emotion = core.BoolPtr(true)
				Expect(keywordsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(250))))
				Expect(keywordsOptionsModel.Sentiment).To(Equal(core.BoolPtr(true)))
				Expect(keywordsOptionsModel.Emotion).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the MetadataOptions model
				metadataOptionsModel := new(naturallanguageunderstandingv1.MetadataOptions)
				Expect(metadataOptionsModel).ToNot(BeNil())

				// Construct an instance of the RelationsOptions model
				relationsOptionsModel := new(naturallanguageunderstandingv1.RelationsOptions)
				Expect(relationsOptionsModel).ToNot(BeNil())
				relationsOptionsModel.Model = core.StringPtr("testString")
				Expect(relationsOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SemanticRolesOptions model
				semanticRolesOptionsModel := new(naturallanguageunderstandingv1.SemanticRolesOptions)
				Expect(semanticRolesOptionsModel).ToNot(BeNil())
				semanticRolesOptionsModel.Limit = core.Int64Ptr(int64(38))
				semanticRolesOptionsModel.Keywords = core.BoolPtr(true)
				semanticRolesOptionsModel.Entities = core.BoolPtr(true)
				Expect(semanticRolesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(semanticRolesOptionsModel.Keywords).To(Equal(core.BoolPtr(true)))
				Expect(semanticRolesOptionsModel.Entities).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SentimentOptions model
				sentimentOptionsModel := new(naturallanguageunderstandingv1.SentimentOptions)
				Expect(sentimentOptionsModel).ToNot(BeNil())
				sentimentOptionsModel.Document = core.BoolPtr(true)
				sentimentOptionsModel.Targets = []string{"testString"}
				sentimentOptionsModel.Model = core.StringPtr("testString")
				Expect(sentimentOptionsModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(sentimentOptionsModel.Targets).To(Equal([]string{"testString"}))
				Expect(sentimentOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SummarizationOptions model
				summarizationOptionsModel := new(naturallanguageunderstandingv1.SummarizationOptions)
				Expect(summarizationOptionsModel).ToNot(BeNil())
				summarizationOptionsModel.Limit = core.Int64Ptr(int64(10))
				Expect(summarizationOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))

				// Construct an instance of the CategoriesOptions model
				categoriesOptionsModel := new(naturallanguageunderstandingv1.CategoriesOptions)
				Expect(categoriesOptionsModel).ToNot(BeNil())
				categoriesOptionsModel.Explanation = core.BoolPtr(true)
				categoriesOptionsModel.Limit = core.Int64Ptr(int64(10))
				categoriesOptionsModel.Model = core.StringPtr("testString")
				Expect(categoriesOptionsModel.Explanation).To(Equal(core.BoolPtr(true)))
				Expect(categoriesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(categoriesOptionsModel.Model).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SyntaxOptionsTokens model
				syntaxOptionsTokensModel := new(naturallanguageunderstandingv1.SyntaxOptionsTokens)
				Expect(syntaxOptionsTokensModel).ToNot(BeNil())
				syntaxOptionsTokensModel.Lemma = core.BoolPtr(true)
				syntaxOptionsTokensModel.PartOfSpeech = core.BoolPtr(true)
				Expect(syntaxOptionsTokensModel.Lemma).To(Equal(core.BoolPtr(true)))
				Expect(syntaxOptionsTokensModel.PartOfSpeech).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the SyntaxOptions model
				syntaxOptionsModel := new(naturallanguageunderstandingv1.SyntaxOptions)
				Expect(syntaxOptionsModel).ToNot(BeNil())
				syntaxOptionsModel.Tokens = syntaxOptionsTokensModel
				syntaxOptionsModel.Sentences = core.BoolPtr(true)
				Expect(syntaxOptionsModel.Tokens).To(Equal(syntaxOptionsTokensModel))
				Expect(syntaxOptionsModel.Sentences).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Features model
				featuresModel := new(naturallanguageunderstandingv1.Features)
				Expect(featuresModel).ToNot(BeNil())
				featuresModel.Classifications = classificationsOptionsModel
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = metadataOptionsModel
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Summarization = summarizationOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel
				Expect(featuresModel.Classifications).To(Equal(classificationsOptionsModel))
				Expect(featuresModel.Concepts).To(Equal(conceptsOptionsModel))
				Expect(featuresModel.Emotion).To(Equal(emotionOptionsModel))
				Expect(featuresModel.Entities).To(Equal(entitiesOptionsModel))
				Expect(featuresModel.Keywords).To(Equal(keywordsOptionsModel))
				Expect(featuresModel.Metadata).To(Equal(metadataOptionsModel))
				Expect(featuresModel.Relations).To(Equal(relationsOptionsModel))
				Expect(featuresModel.SemanticRoles).To(Equal(semanticRolesOptionsModel))
				Expect(featuresModel.Sentiment).To(Equal(sentimentOptionsModel))
				Expect(featuresModel.Summarization).To(Equal(summarizationOptionsModel))
				Expect(featuresModel.Categories).To(Equal(categoriesOptionsModel))
				Expect(featuresModel.Syntax).To(Equal(syntaxOptionsModel))

				// Construct an instance of the AnalyzeOptions model
				var analyzeOptionsFeatures *naturallanguageunderstandingv1.Features = nil
				analyzeOptionsModel := naturalLanguageUnderstandingService.NewAnalyzeOptions(analyzeOptionsFeatures)
				analyzeOptionsModel.SetFeatures(featuresModel)
				analyzeOptionsModel.SetText("testString")
				analyzeOptionsModel.SetHTML("testString")
				analyzeOptionsModel.SetURL("testString")
				analyzeOptionsModel.SetClean(true)
				analyzeOptionsModel.SetXpath("testString")
				analyzeOptionsModel.SetFallbackToRaw(true)
				analyzeOptionsModel.SetReturnAnalyzedText(true)
				analyzeOptionsModel.SetLanguage("testString")
				analyzeOptionsModel.SetLimitTextCharacters(int64(38))
				analyzeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(analyzeOptionsModel).ToNot(BeNil())
				Expect(analyzeOptionsModel.Features).To(Equal(featuresModel))
				Expect(analyzeOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(analyzeOptionsModel.HTML).To(Equal(core.StringPtr("testString")))
				Expect(analyzeOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(analyzeOptionsModel.Clean).To(Equal(core.BoolPtr(true)))
				Expect(analyzeOptionsModel.Xpath).To(Equal(core.StringPtr("testString")))
				Expect(analyzeOptionsModel.FallbackToRaw).To(Equal(core.BoolPtr(true)))
				Expect(analyzeOptionsModel.ReturnAnalyzedText).To(Equal(core.BoolPtr(true)))
				Expect(analyzeOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(analyzeOptionsModel.LimitTextCharacters).To(Equal(core.Int64Ptr(int64(38))))
				Expect(analyzeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCategoriesModelOptions successfully`, func() {
				// Construct an instance of the CreateCategoriesModelOptions model
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				createCategoriesModelOptionsModel := naturalLanguageUnderstandingService.NewCreateCategoriesModelOptions(language, trainingData)
				createCategoriesModelOptionsModel.SetLanguage("testString")
				createCategoriesModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				createCategoriesModelOptionsModel.SetTrainingDataContentType("json")
				createCategoriesModelOptionsModel.SetName("testString")
				createCategoriesModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				createCategoriesModelOptionsModel.SetDescription("testString")
				createCategoriesModelOptionsModel.SetModelVersion("testString")
				createCategoriesModelOptionsModel.SetWorkspaceID("testString")
				createCategoriesModelOptionsModel.SetVersionDescription("testString")
				createCategoriesModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCategoriesModelOptionsModel).ToNot(BeNil())
				Expect(createCategoriesModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createCategoriesModelOptionsModel.TrainingDataContentType).To(Equal(core.StringPtr("json")))
				Expect(createCategoriesModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(createCategoriesModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(createCategoriesModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateClassificationsModelOptions successfully`, func() {
				// Construct an instance of the CreateClassificationsModelOptions model
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				createClassificationsModelOptionsModel := naturalLanguageUnderstandingService.NewCreateClassificationsModelOptions(language, trainingData)
				createClassificationsModelOptionsModel.SetLanguage("testString")
				createClassificationsModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				createClassificationsModelOptionsModel.SetTrainingDataContentType("json")
				createClassificationsModelOptionsModel.SetName("testString")
				createClassificationsModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				createClassificationsModelOptionsModel.SetDescription("testString")
				createClassificationsModelOptionsModel.SetModelVersion("testString")
				createClassificationsModelOptionsModel.SetWorkspaceID("testString")
				createClassificationsModelOptionsModel.SetVersionDescription("testString")
				createClassificationsModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createClassificationsModelOptionsModel).ToNot(BeNil())
				Expect(createClassificationsModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createClassificationsModelOptionsModel.TrainingDataContentType).To(Equal(core.StringPtr("json")))
				Expect(createClassificationsModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(createClassificationsModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(createClassificationsModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSentimentModelOptions successfully`, func() {
				// Construct an instance of the CreateSentimentModelOptions model
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				createSentimentModelOptionsModel := naturalLanguageUnderstandingService.NewCreateSentimentModelOptions(language, trainingData)
				createSentimentModelOptionsModel.SetLanguage("testString")
				createSentimentModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				createSentimentModelOptionsModel.SetName("testString")
				createSentimentModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				createSentimentModelOptionsModel.SetDescription("testString")
				createSentimentModelOptionsModel.SetModelVersion("testString")
				createSentimentModelOptionsModel.SetWorkspaceID("testString")
				createSentimentModelOptionsModel.SetVersionDescription("testString")
				createSentimentModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSentimentModelOptionsModel).ToNot(BeNil())
				Expect(createSentimentModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createSentimentModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(createSentimentModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(createSentimentModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCategoriesModelOptions successfully`, func() {
				// Construct an instance of the DeleteCategoriesModelOptions model
				modelID := "testString"
				deleteCategoriesModelOptionsModel := naturalLanguageUnderstandingService.NewDeleteCategoriesModelOptions(modelID)
				deleteCategoriesModelOptionsModel.SetModelID("testString")
				deleteCategoriesModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCategoriesModelOptionsModel).ToNot(BeNil())
				Expect(deleteCategoriesModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCategoriesModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteClassificationsModelOptions successfully`, func() {
				// Construct an instance of the DeleteClassificationsModelOptions model
				modelID := "testString"
				deleteClassificationsModelOptionsModel := naturalLanguageUnderstandingService.NewDeleteClassificationsModelOptions(modelID)
				deleteClassificationsModelOptionsModel.SetModelID("testString")
				deleteClassificationsModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteClassificationsModelOptionsModel).ToNot(BeNil())
				Expect(deleteClassificationsModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteClassificationsModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteModelOptions successfully`, func() {
				// Construct an instance of the DeleteModelOptions model
				modelID := "testString"
				deleteModelOptionsModel := naturalLanguageUnderstandingService.NewDeleteModelOptions(modelID)
				deleteModelOptionsModel.SetModelID("testString")
				deleteModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteModelOptionsModel).ToNot(BeNil())
				Expect(deleteModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSentimentModelOptions successfully`, func() {
				// Construct an instance of the DeleteSentimentModelOptions model
				modelID := "testString"
				deleteSentimentModelOptionsModel := naturalLanguageUnderstandingService.NewDeleteSentimentModelOptions(modelID)
				deleteSentimentModelOptionsModel.SetModelID("testString")
				deleteSentimentModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSentimentModelOptionsModel).ToNot(BeNil())
				Expect(deleteSentimentModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSentimentModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCategoriesModelOptions successfully`, func() {
				// Construct an instance of the GetCategoriesModelOptions model
				modelID := "testString"
				getCategoriesModelOptionsModel := naturalLanguageUnderstandingService.NewGetCategoriesModelOptions(modelID)
				getCategoriesModelOptionsModel.SetModelID("testString")
				getCategoriesModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCategoriesModelOptionsModel).ToNot(BeNil())
				Expect(getCategoriesModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(getCategoriesModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClassificationsModelOptions successfully`, func() {
				// Construct an instance of the GetClassificationsModelOptions model
				modelID := "testString"
				getClassificationsModelOptionsModel := naturalLanguageUnderstandingService.NewGetClassificationsModelOptions(modelID)
				getClassificationsModelOptionsModel.SetModelID("testString")
				getClassificationsModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClassificationsModelOptionsModel).ToNot(BeNil())
				Expect(getClassificationsModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(getClassificationsModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSentimentModelOptions successfully`, func() {
				// Construct an instance of the GetSentimentModelOptions model
				modelID := "testString"
				getSentimentModelOptionsModel := naturalLanguageUnderstandingService.NewGetSentimentModelOptions(modelID)
				getSentimentModelOptionsModel.SetModelID("testString")
				getSentimentModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSentimentModelOptionsModel).ToNot(BeNil())
				Expect(getSentimentModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(getSentimentModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCategoriesModelsOptions successfully`, func() {
				// Construct an instance of the ListCategoriesModelsOptions model
				listCategoriesModelsOptionsModel := naturalLanguageUnderstandingService.NewListCategoriesModelsOptions()
				listCategoriesModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCategoriesModelsOptionsModel).ToNot(BeNil())
				Expect(listCategoriesModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListClassificationsModelsOptions successfully`, func() {
				// Construct an instance of the ListClassificationsModelsOptions model
				listClassificationsModelsOptionsModel := naturalLanguageUnderstandingService.NewListClassificationsModelsOptions()
				listClassificationsModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listClassificationsModelsOptionsModel).ToNot(BeNil())
				Expect(listClassificationsModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListModelsOptions successfully`, func() {
				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := naturalLanguageUnderstandingService.NewListModelsOptions()
				listModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listModelsOptionsModel).ToNot(BeNil())
				Expect(listModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSentimentModelsOptions successfully`, func() {
				// Construct an instance of the ListSentimentModelsOptions model
				listSentimentModelsOptionsModel := naturalLanguageUnderstandingService.NewListSentimentModelsOptions()
				listSentimentModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSentimentModelsOptionsModel).ToNot(BeNil())
				Expect(listSentimentModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCategoriesModelOptions successfully`, func() {
				// Construct an instance of the UpdateCategoriesModelOptions model
				modelID := "testString"
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				updateCategoriesModelOptionsModel := naturalLanguageUnderstandingService.NewUpdateCategoriesModelOptions(modelID, language, trainingData)
				updateCategoriesModelOptionsModel.SetModelID("testString")
				updateCategoriesModelOptionsModel.SetLanguage("testString")
				updateCategoriesModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				updateCategoriesModelOptionsModel.SetTrainingDataContentType("json")
				updateCategoriesModelOptionsModel.SetName("testString")
				updateCategoriesModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				updateCategoriesModelOptionsModel.SetDescription("testString")
				updateCategoriesModelOptionsModel.SetModelVersion("testString")
				updateCategoriesModelOptionsModel.SetWorkspaceID("testString")
				updateCategoriesModelOptionsModel.SetVersionDescription("testString")
				updateCategoriesModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCategoriesModelOptionsModel).ToNot(BeNil())
				Expect(updateCategoriesModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateCategoriesModelOptionsModel.TrainingDataContentType).To(Equal(core.StringPtr("json")))
				Expect(updateCategoriesModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateCategoriesModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateCategoriesModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateClassificationsModelOptions successfully`, func() {
				// Construct an instance of the UpdateClassificationsModelOptions model
				modelID := "testString"
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				updateClassificationsModelOptionsModel := naturalLanguageUnderstandingService.NewUpdateClassificationsModelOptions(modelID, language, trainingData)
				updateClassificationsModelOptionsModel.SetModelID("testString")
				updateClassificationsModelOptionsModel.SetLanguage("testString")
				updateClassificationsModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				updateClassificationsModelOptionsModel.SetTrainingDataContentType("json")
				updateClassificationsModelOptionsModel.SetName("testString")
				updateClassificationsModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				updateClassificationsModelOptionsModel.SetDescription("testString")
				updateClassificationsModelOptionsModel.SetModelVersion("testString")
				updateClassificationsModelOptionsModel.SetWorkspaceID("testString")
				updateClassificationsModelOptionsModel.SetVersionDescription("testString")
				updateClassificationsModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateClassificationsModelOptionsModel).ToNot(BeNil())
				Expect(updateClassificationsModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateClassificationsModelOptionsModel.TrainingDataContentType).To(Equal(core.StringPtr("json")))
				Expect(updateClassificationsModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateClassificationsModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateClassificationsModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSentimentModelOptions successfully`, func() {
				// Construct an instance of the UpdateSentimentModelOptions model
				modelID := "testString"
				language := "testString"
				trainingData := CreateMockReader("This is a mock file.")
				updateSentimentModelOptionsModel := naturalLanguageUnderstandingService.NewUpdateSentimentModelOptions(modelID, language, trainingData)
				updateSentimentModelOptionsModel.SetModelID("testString")
				updateSentimentModelOptionsModel.SetLanguage("testString")
				updateSentimentModelOptionsModel.SetTrainingData(CreateMockReader("This is a mock file."))
				updateSentimentModelOptionsModel.SetName("testString")
				updateSentimentModelOptionsModel.SetUserMetadata(make(map[string]interface{}))
				updateSentimentModelOptionsModel.SetDescription("testString")
				updateSentimentModelOptionsModel.SetModelVersion("testString")
				updateSentimentModelOptionsModel.SetWorkspaceID("testString")
				updateSentimentModelOptionsModel.SetVersionDescription("testString")
				updateSentimentModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSentimentModelOptionsModel).ToNot(BeNil())
				Expect(updateSentimentModelOptionsModel.ModelID).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.TrainingData).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(updateSentimentModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.UserMetadata).To(Equal(make(map[string]interface{})))
				Expect(updateSentimentModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.ModelVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.WorkspaceID).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.VersionDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateSentimentModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
