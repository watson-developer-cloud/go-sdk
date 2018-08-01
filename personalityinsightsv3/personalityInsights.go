// Package personalityinsightsv3 : Operations and models for the PersonalityInsightsV3 service
package personalityinsightsv3
/**
 * Copyright 2018 IBM All Rights Reserved.
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

import (
    "bytes"
    "fmt"
    "os"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// PersonalityInsightsV3 : The PersonalityInsightsV3 service
type PersonalityInsightsV3 struct {
	client *watson.Client
}

// NewPersonalityInsightsV3 : Instantiate PersonalityInsightsV3
func NewPersonalityInsightsV3(creds watson.Credentials) (*PersonalityInsightsV3, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/personality-insights/api"
    }

	client, clientErr := watson.NewClient(creds, "personality_insights")

	if clientErr != nil {
		return nil, clientErr
	}

	return &PersonalityInsightsV3{ client: client }, nil
}

// Profile : Get profile
func (personalityInsights *PersonalityInsightsV3) Profile(options *ProfileOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/profile"
    creds := personalityInsights.client.Creds
    useTM := personalityInsights.client.UseTM
    tokenManager := personalityInsights.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsContentLanguageSet {
        request.Set("Content-Language", fmt.Sprint(options.ContentLanguage))
    }
    if options.IsAcceptLanguageSet {
        request.Set("Accept-Language", fmt.Sprint(options.AcceptLanguage))
    }
    request.Query("version=" + creds.Version)
    if options.IsRawScoresSet {
        request.Query("raw_scores=" + fmt.Sprint(options.RawScores))
    }
    if options.IsCsvHeadersSet {
        request.Query("csv_headers=" + fmt.Sprint(options.CsvHeaders))
    }
    if options.IsConsumptionPreferencesSet {
        request.Query("consumption_preferences=" + fmt.Sprint(options.ConsumptionPreferences))
    }
    if options.ContentType == "application/json" {
        request.Send(options.Content)
    } else {
        request.Send(options.Body)
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(Profile)
    res, _, err := request.EndStruct(&response.Result)

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetProfileResult : Cast result of Profile operation
func GetProfileResult(response *watson.WatsonResponse) *Profile {
    result, ok := response.Result.(*Profile)

    if ok {
        return result
    }

    return nil
}

// ProfileAsCsv : Get profile as csv
func (personalityInsights *PersonalityInsightsV3) ProfileAsCsv(options *ProfileOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/profile"
    creds := personalityInsights.client.Creds
    useTM := personalityInsights.client.UseTM
    tokenManager := personalityInsights.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "text/csv")
    request.Set("Content-Type", fmt.Sprint(options.ContentType))
    if options.IsContentLanguageSet {
        request.Set("Content-Language", fmt.Sprint(options.ContentLanguage))
    }
    if options.IsAcceptLanguageSet {
        request.Set("Accept-Language", fmt.Sprint(options.AcceptLanguage))
    }
    request.Query("version=" + creds.Version)
    if options.IsRawScoresSet {
        request.Query("raw_scores=" + fmt.Sprint(options.RawScores))
    }
    if options.IsCsvHeadersSet {
        request.Query("csv_headers=" + fmt.Sprint(options.CsvHeaders))
    }
    if options.IsConsumptionPreferencesSet {
        request.Query("consumption_preferences=" + fmt.Sprint(options.ConsumptionPreferences))
    }
    if options.ContentType == "application/json" {
        request.Send(options.Content)
    } else {
        request.Send(options.Body)
    }

    if useTM {
        token, tokenErr := tokenManager.GetToken()

        if tokenErr != nil {
            return nil, tokenErr
        }

        request.Set("Authorization", "Bearer " + token)
    } else {
        request.SetBasicAuth(creds.Username, creds.Password)
    }

    response := new(watson.WatsonResponse)

    response.Result = new(os.File)
    res, _, err := request.EndStruct(&response.Result)

    response.Headers = res.Header
    response.StatusCode = res.StatusCode

    if err != nil {
        return nil, err
    }

    if res.StatusCode < 200 || res.StatusCode >= 300 {
        buff := new(bytes.Buffer)
        buff.ReadFrom(res.Body)
        errStr := buff.String()
        err = append(err, fmt.Errorf(errStr))
        return response, err
    }

    return response, nil
}

// GetProfileAsCsvResult : Cast result of ProfileAsCsv operation
func GetProfileAsCsvResult(response *watson.WatsonResponse) *os.File {
    result, ok := response.Result.(*os.File)

    if ok {
        return result
    }

    return nil
}


// Behavior : Behavior struct
type Behavior struct {

	// The unique, non-localized identifier of the characteristic to which the results pertain. IDs have the form `behavior_{value}`.
	TraitID string `json:"trait_id"`

	// The user-visible, localized name of the characteristic.
	Name string `json:"name"`

	// The category of the characteristic: `behavior` for temporal data.
	Category string `json:"category"`

	// For JSON content that is timestamped, the percentage of timestamped input data that occurred during that day of the week or hour of the day. The range is 0 to 1.
	Percentage float64 `json:"percentage"`
}

// ConsumptionPreferences : ConsumptionPreferences struct
type ConsumptionPreferences struct {

	// The unique, non-localized identifier of the consumption preference to which the results pertain. IDs have the form `consumption_preferences_{preference}`.
	ConsumptionPreferenceID string `json:"consumption_preference_id"`

	// The user-visible, localized name of the consumption preference.
	Name string `json:"name"`

	// The score for the consumption preference: * `0.0`: Unlikely * `0.5`: Neutral * `1.0`: Likely The scores for some preferences are binary and do not allow a neutral value. The score is an indication of preference based on the results inferred from the input text, not a normalized percentile.
	Score float64 `json:"score"`
}

// ConsumptionPreferencesCategory : ConsumptionPreferencesCategory struct
type ConsumptionPreferencesCategory struct {

	// The unique, non-localized identifier of the consumption preferences category to which the results pertain. IDs have the form `consumption_preferences_{category}`.
	ConsumptionPreferenceCategoryID string `json:"consumption_preference_category_id"`

	// The user-visible name of the consumption preferences category.
	Name string `json:"name"`

	// Detailed results inferred from the input text for the individual preferences of the category.
	ConsumptionPreferences []ConsumptionPreferences `json:"consumption_preferences"`
}

// Content : Content struct
type Content struct {

	// An array of `ContentItem` objects that provides the text that is to be analyzed.
	ContentItems []ContentItem `json:"contentItems"`
}

// ContentItem : ContentItem struct
type ContentItem struct {

	// The content that is to be analyzed. The service supports up to 20 MB of content for all `ContentItem` objects combined.
	Content string `json:"content"`

	// A unique identifier for this content item.
	ID string `json:"id,omitempty"`

	// A timestamp that identifies when this content was created. Specify a value in milliseconds since the UNIX Epoch (January 1, 1970, at 0:00 UTC). Required only for results that include temporal behavior data.
	Created int64 `json:"created,omitempty"`

	// A timestamp that identifies when this content was last updated. Specify a value in milliseconds since the UNIX Epoch (January 1, 1970, at 0:00 UTC). Required only for results that include temporal behavior data.
	Updated int64 `json:"updated,omitempty"`

	// The MIME type of the content. The default is plain text. The tags are stripped from HTML content before it is analyzed; plain text is processed as submitted.
	Contenttype string `json:"contenttype,omitempty"`

	// The language identifier (two-letter ISO 639-1 identifier) for the language of the content item. The default is `en` (English). Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. A language specified with the **Content-Type** parameter overrides the value of this parameter; any content items that specify a different language are ignored. Omit the **Content-Type** parameter to base the language on the most prevalent specification among the content items; again, content items that specify a different language are ignored. You can specify any combination of languages for the input and response content.
	Language string `json:"language,omitempty"`

	// The unique ID of the parent content item for this item. Used to identify hierarchical relationships between posts/replies, messages/replies, and so on.
	Parentid string `json:"parentid,omitempty"`

	// Indicates whether this content item is a reply to another content item.
	Reply bool `json:"reply,omitempty"`

	// Indicates whether this content item is a forwarded/copied version of another content item.
	Forward bool `json:"forward,omitempty"`
}

// Profile : Profile struct
type Profile struct {

	// The language model that was used to process the input.
	ProcessedLanguage string `json:"processed_language"`

	// The number of words from the input that were used to produce the profile.
	WordCount int64 `json:"word_count"`

	// When guidance is appropriate, a string that provides a message that indicates the number of words found and where that value falls in the range of required or suggested number of words.
	WordCountMessage string `json:"word_count_message,omitempty"`

	// A recursive array of `Trait` objects that provides detailed results for the Big Five personality characteristics (dimensions and facets) inferred from the input text.
	Personality []Trait `json:"personality"`

	// Detailed results for the Needs characteristics inferred from the input text.
	Needs []Trait `json:"needs"`

	// Detailed results for the Values characteristics inferred from the input text.
	Values []Trait `json:"values"`

	// For JSON content that is timestamped, detailed results about the social behavior disclosed by the input in terms of temporal characteristics. The results include information about the distribution of the content over the days of the week and the hours of the day.
	Behavior []Behavior `json:"behavior,omitempty"`

	// If the **consumption_preferences** parameter is `true`, detailed results for each category of consumption preferences. Each element of the array provides information inferred from the input text for the individual preferences of that category.
	ConsumptionPreferences []ConsumptionPreferencesCategory `json:"consumption_preferences,omitempty"`

	// Warning messages associated with the input text submitted with the request. The array is empty if the input generated no warnings.
	Warnings []Warning `json:"warnings"`
}

// ProfileOptions : The profile options.
type ProfileOptions struct {

	// A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://console.bluemix.net/docs/services/personality-insights/input.html#sufficient). For JSON input, provide an object of type `Content`.
	Content Content `json:"content,omitempty"`

    // Indicates whether user set optional parameter Content
    IsContentSet bool

	// A maximum of 20 MB of content to analyze, though the service requires much less text; for more information, see [Providing sufficient input](https://console.bluemix.net/docs/services/personality-insights/input.html#sufficient). For JSON input, provide an object of type `Content`.
	Body string `json:"body,omitempty"`

    // Indicates whether user set optional parameter Body
    IsBodySet bool

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example, 'text/html;charset=utf-8'.
	ContentType string `json:"Content-Type"`

	// The language of the input text for the request: Arabic, English, Japanese, Korean, or Spanish. Regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. The effect of the **Content-Language** parameter depends on the **Content-Type** parameter. When **Content-Type** is `text/plain` or `text/html`, **Content-Language** is the only way to specify the language. When **Content-Type** is `application/json`, **Content-Language** overrides a language specified with the `language` parameter of a `ContentItem` object, and content items that specify a different language are ignored; omit this parameter to base the language on the specification of the content items. You can specify any combination of languages for **Content-Language** and **Accept-Language**.
	ContentLanguage string `json:"Content-Language,omitempty"`

    // Indicates whether user set optional parameter ContentLanguage
    IsContentLanguageSet bool

	// The desired language of the response. For two-character arguments, regional variants are treated as their parent language; for example, `en-US` is interpreted as `en`. You can specify any combination of languages for the input and response content.
	AcceptLanguage string `json:"Accept-Language,omitempty"`

    // Indicates whether user set optional parameter AcceptLanguage
    IsAcceptLanguageSet bool

	// Indicates whether a raw score in addition to a normalized percentile is returned for each characteristic; raw scores are not compared with a sample population. By default, only normalized percentiles are returned.
	RawScores bool `json:"raw_scores,omitempty"`

    // Indicates whether user set optional parameter RawScores
    IsRawScoresSet bool

	// Indicates whether column labels are returned with a CSV response. By default, no column labels are returned. Applies only when the **Accept** parameter is set to `text/csv`.
	CsvHeaders bool `json:"csv_headers,omitempty"`

    // Indicates whether user set optional parameter CsvHeaders
    IsCsvHeadersSet bool

	// Indicates whether consumption preferences are returned with the results. By default, no consumption preferences are returned.
	ConsumptionPreferences bool `json:"consumption_preferences,omitempty"`

    // Indicates whether user set optional parameter ConsumptionPreferences
    IsConsumptionPreferencesSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewProfileOptionsForContent : Instantiate ProfileOptionsForContent
func NewProfileOptionsForContent(content Content) *ProfileOptions {
    return &ProfileOptions{
        Content: content,
        ContentType: "application/json",
    }
}

// SetContent : Allow user to set Content
func (options *ProfileOptions) SetContent(content Content) *ProfileOptions {
    options.Content = content
    options.ContentType = "application/json"
    return options
}

// NewProfileOptionsForHTML : Instantiate ProfileOptionsForHTML
func NewProfileOptionsForHTML(content string) *ProfileOptions {
    return &ProfileOptions{
        Body: content,
        ContentType: "text/html",
    }
}

// SetHTML : Allow user to set HTML
func (options *ProfileOptions) SetHTML(content string) *ProfileOptions {
    options.Body = content
    options.ContentType = "text/html"
    return options
}

// NewProfileOptionsForPlain : Instantiate ProfileOptionsForPlain
func NewProfileOptionsForPlain(content string) *ProfileOptions {
    return &ProfileOptions{
        Body: content,
        ContentType: "text/plain",
    }
}

// SetPlain : Allow user to set Plain
func (options *ProfileOptions) SetPlain(content string) *ProfileOptions {
    options.Body = content
    options.ContentType = "text/plain"
    return options
}

// SetContentLanguage : Allow user to set ContentLanguage
func (options *ProfileOptions) SetContentLanguage(param string) *ProfileOptions {
    options.ContentLanguage = param
    options.IsContentLanguageSet = true
    return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ProfileOptions) SetAcceptLanguage(param string) *ProfileOptions {
    options.AcceptLanguage = param
    options.IsAcceptLanguageSet = true
    return options
}

// SetRawScores : Allow user to set RawScores
func (options *ProfileOptions) SetRawScores(param bool) *ProfileOptions {
    options.RawScores = param
    options.IsRawScoresSet = true
    return options
}

// SetCsvHeaders : Allow user to set CsvHeaders
func (options *ProfileOptions) SetCsvHeaders(param bool) *ProfileOptions {
    options.CsvHeaders = param
    options.IsCsvHeadersSet = true
    return options
}

// SetConsumptionPreferences : Allow user to set ConsumptionPreferences
func (options *ProfileOptions) SetConsumptionPreferences(param bool) *ProfileOptions {
    options.ConsumptionPreferences = param
    options.IsConsumptionPreferencesSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ProfileOptions) SetHeaders(param map[string]string) *ProfileOptions {
    options.Headers = param
    return options
}

// Trait : Trait struct
type Trait struct {

	// The unique, non-localized identifier of the characteristic to which the results pertain. IDs have the form * `big5_{characteristic}` for Big Five personality dimensions * `facet_{characteristic}` for Big Five personality facets * `need_{characteristic}` for Needs *`value_{characteristic}` for Values.
	TraitID string `json:"trait_id"`

	// The user-visible, localized name of the characteristic.
	Name string `json:"name"`

	// The category of the characteristic: `personality` for Big Five personality characteristics, `needs` for Needs, and `values` for Values.
	Category string `json:"category"`

	// The normalized percentile score for the characteristic. The range is 0 to 1. For example, if the percentage for Openness is 0.60, the author scored in the 60th percentile; the author is more open than 59 percent of the population and less open than 39 percent of the population.
	Percentile float64 `json:"percentile"`

	// The raw score for the characteristic. The range is 0 to 1. A higher score generally indicates a greater likelihood that the author has that characteristic, but raw scores must be considered in aggregate: The range of values in practice might be much smaller than 0 to 1, so an individual score must be considered in the context of the overall scores and their range. The raw score is computed based on the input and the service model; it is not normalized or compared with a sample population. The raw score enables comparison of the results against a different sampling population and with a custom normalization approach.
	RawScore float64 `json:"raw_score,omitempty"`

	// **`2017-10-13`**: Indicates whether the characteristic is meaningful for the input language. The field is always `true` for all characteristics of English, Spanish, and Japanese input. The field is `false` for the subset of characteristics of Arabic and Korean input for which the service's models are unable to generate meaningful results. **`2016-10-19`**: Not returned.
	Significant bool `json:"significant,omitempty"`

	// For `personality` (Big Five) dimensions, more detailed results for the facets of each dimension as inferred from the input text.
	Children []Trait `json:"children,omitempty"`
}

// Warning : Warning struct
type Warning struct {

	// The identifier of the warning message.
	WarningID string `json:"warning_id"`

	// The message associated with the `warning_id`: * `WORD_COUNT_MESSAGE`: "There were {number} words in the input. We need a minimum of 600, preferably 1,200 or more, to compute statistically significant estimates." * `JSON_AS_TEXT`: "Request input was processed as text/plain as indicated, however detected a JSON input. Did you mean application/json?" * `CONTENT_TRUNCATED`: "For maximum accuracy while also optimizing processing time, only the first 250KB of input text (excluding markup) was analyzed. Accuracy levels off at approximately 3,000 words so this did not affect the accuracy of the profile." * `PARTIAL_TEXT_USED`, "The text provided to compute the profile was trimmed for performance reasons. This action does not affect the accuracy of the output, as not all of the input text was required." Applies only when Arabic input text exceeds a threshold at which additional words do not contribute to the accuracy of the profile.
	Message string `json:"message"`
}
