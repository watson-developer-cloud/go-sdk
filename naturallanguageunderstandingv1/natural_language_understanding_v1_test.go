/**
 * (C) Copyright IBM Corp. 2020.
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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`NaturalLanguageUnderstandingV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeFalse())
			naturalLanguageUnderstandingService.DisableSSLVerification()
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "https://naturallanguageunderstandingv1/api",
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
				"NATURAL_LANGUAGE_UNDERSTANDING_URL": "https://naturallanguageunderstandingv1/api",
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
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
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
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_UNDERSTANDING_URL": "https://naturallanguageunderstandingv1/api",
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
				"NATURAL_LANGUAGE_UNDERSTANDING_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageUnderstandingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Analyze(analyzeOptions *AnalyzeOptions) - Operation response error`, func() {
		version := "testString"
		analyzePath := "/v1/analyze"
		Context(`Using mock server endpoint`, func() {
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
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

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
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "analyzed_text": "AnalyzedText", "retrieved_url": "RetrievedURL", "usage": {"features": 8, "text_characters": 14, "text_units": 9}, "concepts": [{"text": "Text", "relevance": 9, "dbpedia_resource": "DbpediaResource"}], "entities": [{"type": "Type", "text": "Text", "relevance": 9, "confidence": 10, "mentions": [{"text": "Text", "location": [8], "confidence": 10}], "count": 5, "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}, "disambiguation": {"name": "Name", "dbpedia_resource": "DbpediaResource", "subtype": ["Subtype"]}}], "keywords": [{"count": 5, "relevance": 9, "text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}, "sentiment": {"score": 5}}], "categories": [{"label": "Label", "score": 5, "explanation": {"relevant_text": [{"text": "Text"}]}}], "emotion": {"document": {"emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}, "targets": [{"text": "Text", "emotion": {"anger": 5, "disgust": 7, "fear": 4, "joy": 3, "sadness": 7}}]}, "metadata": {"authors": [{"name": "Name"}], "publication_date": "PublicationDate", "title": "Title", "image": "Image", "feeds": [{"link": "Link"}]}, "relations": [{"score": 5, "sentence": "Sentence", "type": "Type", "arguments": [{"entities": [{"text": "Text", "type": "Type"}], "location": [8], "text": "Text"}]}], "semantic_roles": [{"sentence": "Sentence", "subject": {"text": "Text", "entities": [{"type": "Type", "text": "Text"}], "keywords": [{"text": "Text"}]}, "action": {"text": "Text", "normalized": "Normalized", "verb": {"text": "Text", "tense": "Tense"}}, "object": {"text": "Text", "keywords": [{"text": "Text"}]}}], "sentiment": {"document": {"label": "Label", "score": 5}, "targets": [{"text": "Text", "score": 5}]}, "syntax": {"tokens": [{"text": "Text", "part_of_speech": "ADJ", "location": [8], "lemma": "Lemma"}], "sentences": [{"text": "Text", "location": [8]}]}}`)
				}))
			})
			It(`Invoke Analyze successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := naturalLanguageUnderstandingService.Analyze(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

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
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr = naturalLanguageUnderstandingService.Analyze(analyzeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.AnalyzeWithContext(ctx, analyzeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke Analyze with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())

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
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
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
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeFalse())
			naturalLanguageUnderstandingService.DisableSSLVerification()
			Expect(naturalLanguageUnderstandingService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(naturalLanguageUnderstandingService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "https://naturallanguageunderstandingv1/api",
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
				"NATURAL_LANGUAGE_UNDERSTANDING_URL": "https://naturallanguageunderstandingv1/api",
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
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL: "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
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
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"NATURAL_LANGUAGE_UNDERSTANDING_URL": "https://naturallanguageunderstandingv1/api",
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
				"NATURAL_LANGUAGE_UNDERSTANDING_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL: "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(naturalLanguageUnderstandingService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListModels(listModelsOptions *ListModelsOptions) - Operation response error`, func() {
		version := "testString"
		listModelsPath := "/v1/models"
		Context(`Using mock server endpoint`, func() {
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
					Version: core.StringPtr(version),
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"models": [{"status": "starting", "model_id": "ModelID", "language": "Language", "description": "Description", "workspace_id": "WorkspaceID", "model_version": "ModelVersion", "version": "Version", "version_description": "VersionDescription", "created": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListModels successfully`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr = naturalLanguageUnderstandingService.ListModels(listModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.ListModelsWithContext(ctx, listModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListModels with error: Operation request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
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
	})
	Describe(`DeleteModel(deleteModelOptions *DeleteModelOptions) - Operation response error`, func() {
		version := "testString"
		deleteModelPath := "/v1/models/testString"
		Context(`Using mock server endpoint`, func() {
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
					Version: core.StringPtr(version),
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
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

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
					Version: core.StringPtr(version),
				})
				Expect(serviceErr).To(BeNil())
				Expect(naturalLanguageUnderstandingService).ToNot(BeNil())
				naturalLanguageUnderstandingService.EnableRetries(0, 0)

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

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				naturalLanguageUnderstandingService.DisableRetries()
				result, response, operationErr = naturalLanguageUnderstandingService.DeleteModel(deleteModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = naturalLanguageUnderstandingService.DeleteModelWithContext(ctx, deleteModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteModel with error: Operation validation and request error`, func() {
				naturalLanguageUnderstandingService, serviceErr := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version: core.StringPtr(version),
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
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			naturalLanguageUnderstandingService, _ := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(&naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options{
				URL:           "http://naturallanguageunderstandingv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version: core.StringPtr(version),
			})
			It(`Invoke NewAnalyzeOptions successfully`, func() {
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
				Expect(sentimentOptionsModel.Document).To(Equal(core.BoolPtr(true)))
				Expect(sentimentOptionsModel.Targets).To(Equal([]string{"testString"}))

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
				featuresModel.Concepts = conceptsOptionsModel
				featuresModel.Emotion = emotionOptionsModel
				featuresModel.Entities = entitiesOptionsModel
				featuresModel.Keywords = keywordsOptionsModel
				featuresModel.Metadata = map[string]interface{}{"anyKey": "anyValue"}
				featuresModel.Relations = relationsOptionsModel
				featuresModel.SemanticRoles = semanticRolesOptionsModel
				featuresModel.Sentiment = sentimentOptionsModel
				featuresModel.Categories = categoriesOptionsModel
				featuresModel.Syntax = syntaxOptionsModel
				Expect(featuresModel.Concepts).To(Equal(conceptsOptionsModel))
				Expect(featuresModel.Emotion).To(Equal(emotionOptionsModel))
				Expect(featuresModel.Entities).To(Equal(entitiesOptionsModel))
				Expect(featuresModel.Keywords).To(Equal(keywordsOptionsModel))
				Expect(featuresModel.Metadata).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(featuresModel.Relations).To(Equal(relationsOptionsModel))
				Expect(featuresModel.SemanticRoles).To(Equal(semanticRolesOptionsModel))
				Expect(featuresModel.Sentiment).To(Equal(sentimentOptionsModel))
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
			It(`Invoke NewListModelsOptions successfully`, func() {
				// Construct an instance of the ListModelsOptions model
				listModelsOptionsModel := naturalLanguageUnderstandingService.NewListModelsOptions()
				listModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listModelsOptionsModel).ToNot(BeNil())
				Expect(listModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
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

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
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
