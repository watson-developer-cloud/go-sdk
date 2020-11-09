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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.17.0-8d569e8f-20201030-142059
 */

// Package personalityinsightsv3 : Operations and models for the PersonalityInsightsV3 service
package personalityinsightsv3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/watson-developer-cloud/go-sdk/common"
)

// PersonalityInsightsV3 : The IBM Watson&trade; Personality Insights service enables applications to derive insights
// from social media, enterprise data, or other digital communications. The service uses linguistic analytics to infer
// individuals' intrinsic personality characteristics, including Big Five, Needs, and Values, from digital
// communications such as email, text messages, tweets, and forum posts.
//
// The service can automatically infer, from potentially noisy social media, portraits of individuals that reflect their
// personality characteristics. The service can infer consumption preferences based on the results of its analysis and,
// for JSON content that is timestamped, can report temporal behavior.
// * For information about the meaning of the models that the service uses to describe personality characteristics, see
// [Personality models](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-models#models).
// * For information about the meaning of the consumption preferences, see [Consumption
// preferences](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-preferences#preferences).
//
// **Note:** Request logging is disabled for the Personality Insights service. Regardless of whether you set the
// `X-Watson-Learning-Opt-Out` request header, the service does not log or retain data from requests and responses.
//
// Version: 3.4.4
// See: https://cloud.ibm.com/docs/personality-insights
type PersonalityInsightsV3 struct {
	Service *core.BaseService

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2017-10-13`.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.personality-insights.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "personality_insights"

// PersonalityInsightsV3Options : Service options
type PersonalityInsightsV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format. The current version is
	// `2017-10-13`.
	Version *string `validate:"required"`
}

// NewPersonalityInsightsV3 : constructs an instance of PersonalityInsightsV3 with passed in options.
func NewPersonalityInsightsV3(options *PersonalityInsightsV3Options) (service *PersonalityInsightsV3, err error) {
	// Log deprecation warning
	core.GetLogger().Log(core.LevelWarn, "", "On 1 December 2021, Personality Insights will no longer be available. Consider migrating to Watson Natural Language Understanding. For more information, see Personality Insights Deprecation.")

	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	err = baseService.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &PersonalityInsightsV3{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (personalityInsights *PersonalityInsightsV3) SetServiceURL(url string) error {
	return personalityInsights.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (personalityInsights *PersonalityInsightsV3) GetServiceURL() string {
	return personalityInsights.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (personalityInsights *PersonalityInsightsV3) SetDefaultHeaders(headers http.Header) {
	personalityInsights.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (personalityInsights *PersonalityInsightsV3) SetEnableGzipCompression(enableGzip bool) {
	personalityInsights.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (personalityInsights *PersonalityInsightsV3) GetEnableGzipCompression() bool {
	return personalityInsights.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (personalityInsights *PersonalityInsightsV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	personalityInsights.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (personalityInsights *PersonalityInsightsV3) DisableRetries() {
	personalityInsights.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (personalityInsights *PersonalityInsightsV3) DisableSSLVerification() {
	personalityInsights.Service.DisableSSLVerification()
}

// Profile : Get profile
// Generates a personality profile for the author of the input text. The service accepts a maximum of 20 MB of input
// content, but it requires much less text to produce an accurate profile. The service can analyze text in Arabic,
// English, Japanese, Korean, or Spanish. It can return its results in a variety of languages.
//
// **See also:**
// * [Requesting a profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#input)
// * [Providing sufficient
// input](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#sufficient)
//
// ### Content types
//
//  You can provide input content as plain text (`text/plain`), HTML (`text/html`), or JSON (`application/json`) by
// specifying the **Content-Type** parameter. The default is `text/plain`.
// * Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.
// * Per the HTTP specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII
// character set).
//
// When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character
// encoding of the input text; for example, `Content-Type: text/plain;charset=utf-8`.
//
// **See also:** [Specifying request and response
// formats](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#formats)
//
// ### Accept types
//
//  You must request a response as JSON (`application/json`) or comma-separated values (`text/csv`) by specifying the
// **Accept** parameter. CSV output includes a fixed number of columns. Set the **csv_headers** parameter to `true` to
// request optional column headers for CSV output.
//
// **See also:**
// * [Understanding a JSON
// profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-output#output)
// * [Understanding a CSV
// profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-outputCSV#outputCSV).
func (personalityInsights *PersonalityInsightsV3) Profile(profileOptions *ProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	return personalityInsights.ProfileWithContext(context.Background(), profileOptions)
}

// ProfileWithContext is an alternate form of the Profile method which supports a Context parameter
func (personalityInsights *PersonalityInsightsV3) ProfileWithContext(ctx context.Context, profileOptions *ProfileOptions) (result *Profile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(profileOptions, "profileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(profileOptions, "profileOptions")
	if err != nil {
		return
	}

	if profileOptions.Content != nil && profileOptions.ContentType == nil {
		profileOptions.SetContentType("application/json")
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = personalityInsights.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(personalityInsights.Service.Options.URL, `/v3/profile`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range profileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("personality_insights", "V3", "Profile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if profileOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*profileOptions.ContentType))
	}
	if profileOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*profileOptions.ContentLanguage))
	}
	if profileOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*profileOptions.AcceptLanguage))
	}

	builder.AddQuery("version", fmt.Sprint(*personalityInsights.Version))
	if profileOptions.RawScores != nil {
		builder.AddQuery("raw_scores", fmt.Sprint(*profileOptions.RawScores))
	}
	if profileOptions.CsvHeaders != nil {
		builder.AddQuery("csv_headers", fmt.Sprint(*profileOptions.CsvHeaders))
	}
	if profileOptions.ConsumptionPreferences != nil {
		builder.AddQuery("consumption_preferences", fmt.Sprint(*profileOptions.ConsumptionPreferences))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(profileOptions.ContentType), profileOptions.Content, nil, profileOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = personalityInsights.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfile)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ProfileAsCsv : Get profile as csv
// Generates a personality profile for the author of the input text. The service accepts a maximum of 20 MB of input
// content, but it requires much less text to produce an accurate profile. The service can analyze text in Arabic,
// English, Japanese, Korean, or Spanish. It can return its results in a variety of languages.
//
// **See also:**
// * [Requesting a profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#input)
// * [Providing sufficient
// input](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#sufficient)
//
// ### Content types
//
//  You can provide input content as plain text (`text/plain`), HTML (`text/html`), or JSON (`application/json`) by
// specifying the **Content-Type** parameter. The default is `text/plain`.
// * Per the JSON specification, the default character encoding for JSON content is effectively always UTF-8.
// * Per the HTTP specification, the default encoding for plain text and HTML is ISO-8859-1 (effectively, the ASCII
// character set).
//
// When specifying a content type of plain text or HTML, include the `charset` parameter to indicate the character
// encoding of the input text; for example, `Content-Type: text/plain;charset=utf-8`.
//
// **See also:** [Specifying request and response
// formats](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#formats)
//
// ### Accept types
//
//  You must request a response as JSON (`application/json`) or comma-separated values (`text/csv`) by specifying the
// **Accept** parameter. CSV output includes a fixed number of columns. Set the **csv_headers** parameter to `true` to
// request optional column headers for CSV output.
//
// **See also:**
// * [Understanding a JSON
// profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-output#output)
// * [Understanding a CSV
// profile](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-outputCSV#outputCSV).
func (personalityInsights *PersonalityInsightsV3) ProfileAsCsv(profileOptions *ProfileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return personalityInsights.ProfileAsCsvWithContext(context.Background(), profileOptions)
}

// ProfileAsCsvWithContext is an alternate form of the ProfileAsCsv method which supports a Context parameter
func (personalityInsights *PersonalityInsightsV3) ProfileAsCsvWithContext(ctx context.Context, profileOptions *ProfileOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(profileOptions, "profileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(profileOptions, "profileOptions")
	if err != nil {
		return
	}

	if profileOptions.Content != nil && profileOptions.ContentType == nil {
		profileOptions.SetContentType("application/json")
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = personalityInsights.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(personalityInsights.Service.Options.URL, `/v3/profile`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range profileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("personality_insights", "V3", "ProfileAsCsv")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/csv")
	if profileOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*profileOptions.ContentType))
	}
	if profileOptions.ContentLanguage != nil {
		builder.AddHeader("Content-Language", fmt.Sprint(*profileOptions.ContentLanguage))
	}
	if profileOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*profileOptions.AcceptLanguage))
	}

	builder.AddQuery("version", fmt.Sprint(*personalityInsights.Version))
	if profileOptions.RawScores != nil {
		builder.AddQuery("raw_scores", fmt.Sprint(*profileOptions.RawScores))
	}
	if profileOptions.CsvHeaders != nil {
		builder.AddQuery("csv_headers", fmt.Sprint(*profileOptions.CsvHeaders))
	}
	if profileOptions.ConsumptionPreferences != nil {
		builder.AddQuery("consumption_preferences", fmt.Sprint(*profileOptions.ConsumptionPreferences))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(profileOptions.ContentType), profileOptions.Content, nil, profileOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = personalityInsights.Service.Request(request, &result)

	return
}

// Behavior : The temporal behavior for the input content.
type Behavior struct {
	// The unique, non-localized identifier of the characteristic to which the results pertain. IDs have the form
	// `behavior_{value}`.
	TraitID *string `json:"trait_id" validate:"required"`

	// The user-visible, localized name of the characteristic.
	Name *string `json:"name" validate:"required"`

	// The category of the characteristic: `behavior` for temporal data.
	Category *string `json:"category" validate:"required"`

	// For JSON content that is timestamped, the percentage of timestamped input data that occurred during that day of the
	// week or hour of the day. The range is 0 to 1.
	Percentage *float64 `json:"percentage" validate:"required"`
}

// UnmarshalBehavior unmarshals an instance of Behavior from the specified map of raw messages.
func UnmarshalBehavior(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Behavior)
	err = core.UnmarshalPrimitive(m, "trait_id", &obj.TraitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "percentage", &obj.Percentage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionPreferences : A consumption preference that the service inferred from the input content.
type ConsumptionPreferences struct {
	// The unique, non-localized identifier of the consumption preference to which the results pertain. IDs have the form
	// `consumption_preferences_{preference}`.
	ConsumptionPreferenceID *string `json:"consumption_preference_id" validate:"required"`

	// The user-visible, localized name of the consumption preference.
	Name *string `json:"name" validate:"required"`

	// The score for the consumption preference:
	// * `0.0`: Unlikely
	// * `0.5`: Neutral
	// * `1.0`: Likely
	//
	// The scores for some preferences are binary and do not allow a neutral value. The score is an indication of
	// preference based on the results inferred from the input text, not a normalized percentile.
	Score *float64 `json:"score" validate:"required"`
}

// UnmarshalConsumptionPreferences unmarshals an instance of ConsumptionPreferences from the specified map of raw messages.
func UnmarshalConsumptionPreferences(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionPreferences)
	err = core.UnmarshalPrimitive(m, "consumption_preference_id", &obj.ConsumptionPreferenceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConsumptionPreferencesCategory : The consumption preferences that the service inferred from the input content.
type ConsumptionPreferencesCategory struct {
	// The unique, non-localized identifier of the consumption preferences category to which the results pertain. IDs have
	// the form `consumption_preferences_{category}`.
	ConsumptionPreferenceCategoryID *string `json:"consumption_preference_category_id" validate:"required"`

	// The user-visible name of the consumption preferences category.
	Name *string `json:"name" validate:"required"`

	// Detailed results inferred from the input text for the individual preferences of the category.
	ConsumptionPreferences []ConsumptionPreferences `json:"consumption_preferences" validate:"required"`
}

// UnmarshalConsumptionPreferencesCategory unmarshals an instance of ConsumptionPreferencesCategory from the specified map of raw messages.
func UnmarshalConsumptionPreferencesCategory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConsumptionPreferencesCategory)
	err = core.UnmarshalPrimitive(m, "consumption_preference_category_id", &obj.ConsumptionPreferenceCategoryID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "consumption_preferences", &obj.ConsumptionPreferences, UnmarshalConsumptionPreferences)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Content : The full input content that the service is to analyze.
type Content struct {
	// An array of `ContentItem` objects that provides the text that is to be analyzed.
	ContentItems []ContentItem `json:"contentItems" validate:"required"`
}

// NewContent : Instantiate Content (Generic Model Constructor)
func (*PersonalityInsightsV3) NewContent(contentItems []ContentItem) (model *Content, err error) {
	model = &Content{
		ContentItems: contentItems,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalContent unmarshals an instance of Content from the specified map of raw messages.
func UnmarshalContent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Content)
	err = core.UnmarshalModel(m, "contentItems", &obj.ContentItems, UnmarshalContentItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ContentItem : An input content item that the service is to analyze.
type ContentItem struct {
	// The content that is to be analyzed. The service supports up to 20 MB of content for all `ContentItem` objects
	// combined.
	Content *string `json:"content" validate:"required"`

	// A unique identifier for this content item.
	ID *string `json:"id,omitempty"`

	// A timestamp that identifies when this content was created. Specify a value in milliseconds since the UNIX Epoch
	// (January 1, 1970, at 0:00 UTC). Required only for results that include temporal behavior data.
	Created *int64 `json:"created,omitempty"`

	// A timestamp that identifies when this content was last updated. Specify a value in milliseconds since the UNIX Epoch
	// (January 1, 1970, at 0:00 UTC). Required only for results that include temporal behavior data.
	Updated *int64 `json:"updated,omitempty"`

	// The MIME type of the content. The default is plain text. The tags are stripped from HTML content before it is
	// analyzed; plain text is processed as submitted.
	Contenttype *string `json:"contenttype,omitempty"`

	// The language identifier (two-letter ISO 639-1 identifier) for the language of the content item. The default is `en`
	// (English). Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. A
	// language specified with the **Content-Type** parameter overrides the value of this parameter; any content items that
	// specify a different language are ignored. Omit the **Content-Type** parameter to base the language on the most
	// prevalent specification among the content items; again, content items that specify a different language are ignored.
	// You can specify any combination of languages for the input and response content.
	Language *string `json:"language,omitempty"`

	// The unique ID of the parent content item for this item. Used to identify hierarchical relationships between
	// posts/replies, messages/replies, and so on.
	Parentid *string `json:"parentid,omitempty"`

	// Indicates whether this content item is a reply to another content item.
	Reply *bool `json:"reply,omitempty"`

	// Indicates whether this content item is a forwarded/copied version of another content item.
	Forward *bool `json:"forward,omitempty"`
}

// Constants associated with the ContentItem.Contenttype property.
// The MIME type of the content. The default is plain text. The tags are stripped from HTML content before it is
// analyzed; plain text is processed as submitted.
const (
	ContentItem_Contenttype_TextHTML  = "text/html"
	ContentItem_Contenttype_TextPlain = "text/plain"
)

// Constants associated with the ContentItem.Language property.
// The language identifier (two-letter ISO 639-1 identifier) for the language of the content item. The default is `en`
// (English). Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. A
// language specified with the **Content-Type** parameter overrides the value of this parameter; any content items that
// specify a different language are ignored. Omit the **Content-Type** parameter to base the language on the most
// prevalent specification among the content items; again, content items that specify a different language are ignored.
// You can specify any combination of languages for the input and response content.
const (
	ContentItem_Language_Ar = "ar"
	ContentItem_Language_En = "en"
	ContentItem_Language_Es = "es"
	ContentItem_Language_Ja = "ja"
	ContentItem_Language_Ko = "ko"
)

// NewContentItem : Instantiate ContentItem (Generic Model Constructor)
func (*PersonalityInsightsV3) NewContentItem(content string) (model *ContentItem, err error) {
	model = &ContentItem{
		Content: core.StringPtr(content),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalContentItem unmarshals an instance of ContentItem from the specified map of raw messages.
func UnmarshalContentItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContentItem)
	err = core.UnmarshalPrimitive(m, "content", &obj.Content)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "contenttype", &obj.Contenttype)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parentid", &obj.Parentid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reply", &obj.Reply)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "forward", &obj.Forward)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Profile : The personality profile that the service generated for the input content.
type Profile struct {
	// The language model that was used to process the input.
	ProcessedLanguage *string `json:"processed_language" validate:"required"`

	// The number of words from the input that were used to produce the profile.
	WordCount *int64 `json:"word_count" validate:"required"`

	// When guidance is appropriate, a string that provides a message that indicates the number of words found and where
	// that value falls in the range of required or suggested number of words.
	WordCountMessage *string `json:"word_count_message,omitempty"`

	// A recursive array of `Trait` objects that provides detailed results for the Big Five personality characteristics
	// (dimensions and facets) inferred from the input text.
	Personality []Trait `json:"personality" validate:"required"`

	// Detailed results for the Needs characteristics inferred from the input text.
	Needs []Trait `json:"needs" validate:"required"`

	// Detailed results for the Values characteristics inferred from the input text.
	Values []Trait `json:"values" validate:"required"`

	// For JSON content that is timestamped, detailed results about the social behavior disclosed by the input in terms of
	// temporal characteristics. The results include information about the distribution of the content over the days of the
	// week and the hours of the day.
	Behavior []Behavior `json:"behavior,omitempty"`

	// If the **consumption_preferences** parameter is `true`, detailed results for each category of consumption
	// preferences. Each element of the array provides information inferred from the input text for the individual
	// preferences of that category.
	ConsumptionPreferences []ConsumptionPreferencesCategory `json:"consumption_preferences,omitempty"`

	// An array of warning messages that are associated with the input text for the request. The array is empty if the
	// input generated no warnings.
	Warnings []Warning `json:"warnings" validate:"required"`
}

// Constants associated with the Profile.ProcessedLanguage property.
// The language model that was used to process the input.
const (
	Profile_ProcessedLanguage_Ar = "ar"
	Profile_ProcessedLanguage_En = "en"
	Profile_ProcessedLanguage_Es = "es"
	Profile_ProcessedLanguage_Ja = "ja"
	Profile_ProcessedLanguage_Ko = "ko"
)

// UnmarshalProfile unmarshals an instance of Profile from the specified map of raw messages.
func UnmarshalProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Profile)
	err = core.UnmarshalPrimitive(m, "processed_language", &obj.ProcessedLanguage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "word_count", &obj.WordCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "word_count_message", &obj.WordCountMessage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "personality", &obj.Personality, UnmarshalTrait)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "needs", &obj.Needs, UnmarshalTrait)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalTrait)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "behavior", &obj.Behavior, UnmarshalBehavior)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "consumption_preferences", &obj.ConsumptionPreferences, UnmarshalConsumptionPreferencesCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfileOptions : The Profile options.
type ProfileOptions struct {
	// A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see
	// [Providing sufficient
	// input](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#sufficient). For JSON input,
	// provide an object of type `Content`.
	Content *Content `json:"content,omitempty"`

	// A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see
	// [Providing sufficient
	// input](https://cloud.ibm.com/docs/personality-insights?topic=personality-insights-input#sufficient). For JSON input,
	// provide an object of type `Content`.
	Body *string `json:"body,omitempty"`

	// The type of the input. For more information, see **Content types** in the method description.
	ContentType *string `json:"Content-Type,omitempty"`

	// The language of the input text for the request: Arabic, English, Japanese, Korean, or Spanish. Regional variants are
	// treated as their parent language; for example, `en-US` is interpreted as `en`.
	//
	// The effect of the **Content-Language** parameter depends on the **Content-Type** parameter. When **Content-Type** is
	// `text/plain` or `text/html`, **Content-Language** is the only way to specify the language. When **Content-Type** is
	// `application/json`, **Content-Language** overrides a language specified with the `language` parameter of a
	// `ContentItem` object, and content items that specify a different language are ignored; omit this parameter to base
	// the language on the specification of the content items. You can specify any combination of languages for
	// **Content-Language** and **Accept-Language**.
	ContentLanguage *string `json:"Content-Language,omitempty"`

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent
	// language; for example, `en-US` is interpreted as `en`. You can specify any combination of languages for the input
	// and response content.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Indicates whether a raw score in addition to a normalized percentile is returned for each characteristic; raw scores
	// are not compared with a sample population. By default, only normalized percentiles are returned.
	RawScores *bool `json:"raw_scores,omitempty"`

	// Indicates whether column labels are returned with a CSV response. By default, no column labels are returned. Applies
	// only when the response type is CSV (`text/csv`).
	CsvHeaders *bool `json:"csv_headers,omitempty"`

	// Indicates whether consumption preferences are returned with the results. By default, no consumption preferences are
	// returned.
	ConsumptionPreferences *bool `json:"consumption_preferences,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ProfileOptions.ContentLanguage property.
// The language of the input text for the request: Arabic, English, Japanese, Korean, or Spanish. Regional variants are
// treated as their parent language; for example, `en-US` is interpreted as `en`.
//
// The effect of the **Content-Language** parameter depends on the **Content-Type** parameter. When **Content-Type** is
// `text/plain` or `text/html`, **Content-Language** is the only way to specify the language. When **Content-Type** is
// `application/json`, **Content-Language** overrides a language specified with the `language` parameter of a
// `ContentItem` object, and content items that specify a different language are ignored; omit this parameter to base
// the language on the specification of the content items. You can specify any combination of languages for
// **Content-Language** and **Accept-Language**.
const (
	ProfileOptions_ContentLanguage_Ar = "ar"
	ProfileOptions_ContentLanguage_En = "en"
	ProfileOptions_ContentLanguage_Es = "es"
	ProfileOptions_ContentLanguage_Ja = "ja"
	ProfileOptions_ContentLanguage_Ko = "ko"
)

// Constants associated with the ProfileOptions.AcceptLanguage property.
// The desired language of the response. For two-character arguments, regional variants are treated as their parent
// language; for example, `en-US` is interpreted as `en`. You can specify any combination of languages for the input and
// response content.
const (
	ProfileOptions_AcceptLanguage_Ar   = "ar"
	ProfileOptions_AcceptLanguage_De   = "de"
	ProfileOptions_AcceptLanguage_En   = "en"
	ProfileOptions_AcceptLanguage_Es   = "es"
	ProfileOptions_AcceptLanguage_Fr   = "fr"
	ProfileOptions_AcceptLanguage_It   = "it"
	ProfileOptions_AcceptLanguage_Ja   = "ja"
	ProfileOptions_AcceptLanguage_Ko   = "ko"
	ProfileOptions_AcceptLanguage_PtBr = "pt-br"
	ProfileOptions_AcceptLanguage_ZhCn = "zh-cn"
	ProfileOptions_AcceptLanguage_ZhTw = "zh-tw"
)

// NewProfileOptions : Instantiate ProfileOptions
func (*PersonalityInsightsV3) NewProfileOptions() *ProfileOptions {
	return &ProfileOptions{}
}

// SetContent : Allow user to set Content
func (options *ProfileOptions) SetContent(content *Content) *ProfileOptions {
	options.Content = content
	return options
}

// SetBody : Allow user to set Body
func (options *ProfileOptions) SetBody(body string) *ProfileOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentType : Allow user to set ContentType
func (options *ProfileOptions) SetContentType(contentType string) *ProfileOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ProfileOptions) SetContentLanguage(contentLanguage string) *ProfileOptions {
	options.ContentLanguage = core.StringPtr(contentLanguage)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ProfileOptions) SetAcceptLanguage(acceptLanguage string) *ProfileOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetRawScores : Allow user to set RawScores
func (options *ProfileOptions) SetRawScores(rawScores bool) *ProfileOptions {
	options.RawScores = core.BoolPtr(rawScores)
	return options
}

// SetCsvHeaders : Allow user to set CsvHeaders
func (options *ProfileOptions) SetCsvHeaders(csvHeaders bool) *ProfileOptions {
	options.CsvHeaders = core.BoolPtr(csvHeaders)
	return options
}

// SetConsumptionPreferences : Allow user to set ConsumptionPreferences
func (options *ProfileOptions) SetConsumptionPreferences(consumptionPreferences bool) *ProfileOptions {
	options.ConsumptionPreferences = core.BoolPtr(consumptionPreferences)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ProfileOptions) SetHeaders(param map[string]string) *ProfileOptions {
	options.Headers = param
	return options
}

// Trait : The characteristics that the service inferred from the input content.
type Trait struct {
	// The unique, non-localized identifier of the characteristic to which the results pertain. IDs have the form
	// * `big5_{characteristic}` for Big Five personality dimensions
	// * `facet_{characteristic}` for Big Five personality facets
	// * `need_{characteristic}` for Needs
	//  *`value_{characteristic}` for Values.
	TraitID *string `json:"trait_id" validate:"required"`

	// The user-visible, localized name of the characteristic.
	Name *string `json:"name" validate:"required"`

	// The category of the characteristic: `personality` for Big Five personality characteristics, `needs` for Needs, and
	// `values` for Values.
	Category *string `json:"category" validate:"required"`

	// The normalized percentile score for the characteristic. The range is 0 to 1. For example, if the percentage for
	// Openness is 0.60, the author scored in the 60th percentile; the author is more open than 59 percent of the
	// population and less open than 39 percent of the population.
	Percentile *float64 `json:"percentile" validate:"required"`

	// The raw score for the characteristic. The range is 0 to 1. A higher score generally indicates a greater likelihood
	// that the author has that characteristic, but raw scores must be considered in aggregate: The range of values in
	// practice might be much smaller than 0 to 1, so an individual score must be considered in the context of the overall
	// scores and their range.
	//
	// The raw score is computed based on the input and the service model; it is not normalized or compared with a sample
	// population. The raw score enables comparison of the results against a different sampling population and with a
	// custom normalization approach.
	RawScore *float64 `json:"raw_score,omitempty"`

	// **`2017-10-13`**: Indicates whether the characteristic is meaningful for the input language. The field is always
	// `true` for all characteristics of English, Spanish, and Japanese input. The field is `false` for the subset of
	// characteristics of Arabic and Korean input for which the service's models are unable to generate meaningful results.
	// **`2016-10-19`**: Not returned.
	Significant *bool `json:"significant,omitempty"`

	// For `personality` (Big Five) dimensions, more detailed results for the facets of each dimension as inferred from the
	// input text.
	Children []Trait `json:"children,omitempty"`
}

// Constants associated with the Trait.Category property.
// The category of the characteristic: `personality` for Big Five personality characteristics, `needs` for Needs, and
// `values` for Values.
const (
	Trait_Category_Needs       = "needs"
	Trait_Category_Personality = "personality"
	Trait_Category_Values      = "values"
)

// UnmarshalTrait unmarshals an instance of Trait from the specified map of raw messages.
func UnmarshalTrait(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Trait)
	err = core.UnmarshalPrimitive(m, "trait_id", &obj.TraitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "percentile", &obj.Percentile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "raw_score", &obj.RawScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "significant", &obj.Significant)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "children", &obj.Children, UnmarshalTrait)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Warning : A warning message that is associated with the input content.
type Warning struct {
	// The identifier of the warning message.
	WarningID *string `json:"warning_id" validate:"required"`

	// The message associated with the `warning_id`:
	// * `WORD_COUNT_MESSAGE`: "There were {number} words in the input. We need a minimum of 600, preferably 1,200 or more,
	// to compute statistically significant estimates."
	// * `JSON_AS_TEXT`: "Request input was processed as text/plain as indicated, however detected a JSON input. Did you
	// mean application/json?"
	// * `CONTENT_TRUNCATED`: "For maximum accuracy while also optimizing processing time, only the first 250KB of input
	// text (excluding markup) was analyzed. Accuracy levels off at approximately 3,000 words so this did not affect the
	// accuracy of the profile."
	// * `PARTIAL_TEXT_USED`, "The text provided to compute the profile was trimmed for performance reasons. This action
	// does not affect the accuracy of the output, as not all of the input text was required." Applies only when Arabic
	// input text exceeds a threshold at which additional words do not contribute to the accuracy of the profile.
	Message *string `json:"message" validate:"required"`
}

// Constants associated with the Warning.WarningID property.
// The identifier of the warning message.
const (
	Warning_WarningID_ContentTruncated = "CONTENT_TRUNCATED"
	Warning_WarningID_JSONAsText       = "JSON_AS_TEXT"
	Warning_WarningID_PartialTextUsed  = "PARTIAL_TEXT_USED"
	Warning_WarningID_WordCountMessage = "WORD_COUNT_MESSAGE"
)

// UnmarshalWarning unmarshals an instance of Warning from the specified map of raw messages.
func UnmarshalWarning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Warning)
	err = core.UnmarshalPrimitive(m, "warning_id", &obj.WarningID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
