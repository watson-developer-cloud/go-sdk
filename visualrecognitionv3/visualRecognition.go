// Package visualrecognitionv3 : Operations and models for the VisualRecognitionV3 service
package visualrecognitionv3
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
    "github.com/go-openapi/strfmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

// VisualRecognitionV3 : The VisualRecognitionV3 service
type VisualRecognitionV3 struct {
	client *watson.Client
}

// NewVisualRecognitionV3 : Instantiate VisualRecognitionV3
func NewVisualRecognitionV3(creds watson.Credentials) (*VisualRecognitionV3, error) {
    if creds.ServiceURL == "" {
        creds.ServiceURL = "https://gateway.watsonplatform.net/visual-recognition/api"
    }

	client, clientErr := watson.NewClient(creds, "watson_vision_combined")

	if clientErr != nil {
		return nil, clientErr
	}

	return &VisualRecognitionV3{ client: client }, nil
}

// Classify : Classify images
func (visualRecognition *VisualRecognitionV3) Classify(options *ClassifyOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classify"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    if options.IsAcceptLanguageSet {
        request.Set("Accept-Language", fmt.Sprint(options.AcceptLanguage))
    }
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    form := map[string]interface{}{}
    if options.IsImagesFileSet {
        request.SendFile(options.ImagesFile, "", "images_file")
    }
    if options.IsURLSet {
        form["url"] = options.URL
    }
    if options.IsThresholdSet {
        form["threshold"] = options.Threshold
    }
    if options.IsOwnersSet {
        form["owners"] = strings.Join(options.Owners, ",")
    }
    if options.IsClassifierIdsSet {
        form["classifier_ids"] = strings.Join(options.ClassifierIds, ",")
    }
    request.Send(form)

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

    response.Result = new(ClassifiedImages)
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

// GetClassifyResult : Cast result of Classify operation
func GetClassifyResult(response *watson.WatsonResponse) *ClassifiedImages {
    result, ok := response.Result.(*ClassifiedImages)

    if ok {
        return result
    }

    return nil
}

// DetectFaces : Detect faces in images
func (visualRecognition *VisualRecognitionV3) DetectFaces(options *DetectFacesOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/detect_faces"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    form := map[string]interface{}{}
    if options.IsImagesFileSet {
        request.SendFile(options.ImagesFile, "", "images_file")
    }
    if options.IsURLSet {
        form["url"] = options.URL
    }
    request.Send(form)

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

    response.Result = new(DetectedFaces)
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

// GetDetectFacesResult : Cast result of DetectFaces operation
func GetDetectFacesResult(response *watson.WatsonResponse) *DetectedFaces {
    result, ok := response.Result.(*DetectedFaces)

    if ok {
        return result
    }

    return nil
}

// CreateClassifier : Create a classifier
func (visualRecognition *VisualRecognitionV3) CreateClassifier(options *CreateClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    form := map[string]interface{}{}
    form["name"] = options.Name
    request.SendFile(options.ClassnamePositiveExamples, "", "classname_positive_examples")
    if options.IsNegativeExamplesSet {
        request.SendFile(options.NegativeExamples, "", "negative_examples")
    }
    request.Send(form)

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

    response.Result = new(Classifier)
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

// GetCreateClassifierResult : Cast result of CreateClassifier operation
func GetCreateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// DeleteClassifier : Delete a classifier
func (visualRecognition *VisualRecognitionV3) DeleteClassifier(options *DeleteClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)

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

    res, _, err := request.End()

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


// GetClassifier : Retrieve classifier details
func (visualRecognition *VisualRecognitionV3) GetClassifier(options *GetClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)

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

    response.Result = new(Classifier)
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

// GetGetClassifierResult : Cast result of GetClassifier operation
func GetGetClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// ListClassifiers : Retrieve a list of classifiers
func (visualRecognition *VisualRecognitionV3) ListClassifiers(options *ListClassifiersOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    if options.IsVerboseSet {
        request.Query("verbose=" + fmt.Sprint(options.Verbose))
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

    response.Result = new(Classifiers)
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

// GetListClassifiersResult : Cast result of ListClassifiers operation
func GetListClassifiersResult(response *watson.WatsonResponse) *Classifiers {
    result, ok := response.Result.(*Classifiers)

    if ok {
        return result
    }

    return nil
}

// UpdateClassifier : Update a classifier
func (visualRecognition *VisualRecognitionV3) UpdateClassifier(options *UpdateClassifierOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Post(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Query("version=" + creds.Version)
    request.Type("multipart")
    if options.IsClassnamePositiveExamplesSet {
        request.SendFile(options.ClassnamePositiveExamples, "", "classname_positive_examples")
    }
    if options.IsNegativeExamplesSet {
        request.SendFile(options.NegativeExamples, "", "negative_examples")
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

    response.Result = new(Classifier)
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

// GetUpdateClassifierResult : Cast result of UpdateClassifier operation
func GetUpdateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// GetCoreMlModel : Retrieve a Core ML model of a classifier
func (visualRecognition *VisualRecognitionV3) GetCoreMlModel(options *GetCoreMlModelOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}/core_ml_model"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", options.ClassifierID, 1)
    request := req.New().Get(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/octet-stream")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)

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

// GetGetCoreMlModelResult : Cast result of GetCoreMlModel operation
func GetGetCoreMlModelResult(response *watson.WatsonResponse) *os.File {
    result, ok := response.Result.(*os.File)

    if ok {
        return result
    }

    return nil
}

// DeleteUserData : Delete labeled data
func (visualRecognition *VisualRecognitionV3) DeleteUserData(options *DeleteUserDataOptions) (*watson.WatsonResponse, []error) {
    path := "/v3/user_data"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    request := req.New().Delete(creds.ServiceURL + path)

    for headerName, headerValue := range options.Headers {
        request.Set(headerName, headerValue)
    }

    request.Set("Accept", "application/json")
    request.Set("Content-Type", "application/json")
    request.Query("version=" + creds.Version)
    request.Query("customer_id=" + fmt.Sprint(options.CustomerID))

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

    res, _, err := request.End()

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



// Class : A category within a classifier.
type Class struct {

	// The name of the class.
	ClassName string `json:"class"`
}

// ClassResult : Result of a class within a classifier.
type ClassResult struct {

	// Name of the class.
	ClassName string `json:"class"`

	// Confidence score for the property in the range of 0 to 1. A higher score indicates greater likelihood that the class is depicted in the image. The default threshold for returning scores from a classifier is 0.5.
	Score float32 `json:"score,omitempty"`

	// Knowledge graph of the property. For example, `/fruit/pome/apple/eating apple/Granny Smith`. Included only if identified.
	TypeHierarchy string `json:"type_hierarchy,omitempty"`
}

// ClassifiedImage : Results for one image.
type ClassifiedImage struct {

	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceURL string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedURL string `json:"resolved_url,omitempty"`

	// Relative path of the image file if uploaded directly. Not returned when the image is passed by URL.
	Image string `json:"image,omitempty"`

	// Information about what might have caused a failure, such as an image that is too large. Not returned when there is no error.
	Error ErrorInfo `json:"error,omitempty"`

	// The classifiers.
	Classifiers []ClassifierResult `json:"classifiers"`
}

// ClassifiedImages : Results for all images.
type ClassifiedImages struct {

	// Number of custom classes identified in the images.
	CustomClasses int64 `json:"custom_classes,omitempty"`

	// Number of images processed for the API call.
	ImagesProcessed int64 `json:"images_processed,omitempty"`

	// Classified images.
	Images []ClassifiedImage `json:"images"`

	// Information about what might cause less than optimal output. For example, a request sent with a corrupt .zip file and a list of image URLs will still complete, but does not return the expected output. Not returned when there is no warning.
	Warnings []WarningInfo `json:"warnings,omitempty"`
}

// Classifier : Information about a classifier.
type Classifier struct {

	// ID of a classifier identified in the image.
	ClassifierID string `json:"classifier_id"`

	// Name of the classifier.
	Name string `json:"name"`

	// Unique ID of the account who owns the classifier. Returned when verbose=`true`. Might not be returned by some requests.
	Owner string `json:"owner,omitempty"`

	// Training status of classifier.
	Status string `json:"status,omitempty"`

	// Whether the classifier can be downloaded as a Core ML model after the training status is `ready`.
	CoreMlEnabled bool `json:"core_ml_enabled,omitempty"`

	// If classifier training has failed, this field may explain why.
	Explanation string `json:"explanation,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was created.
	Created strfmt.DateTime `json:"created,omitempty"`

	// Classes that define a classifier.
	Classes []Class `json:"classes,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was updated. Returned when verbose=`true`. Might not be returned by some requests. Identical to `updated` and retained for backward compatibility.
	Retrained strfmt.DateTime `json:"retrained,omitempty"`

	// Date and time in Coordinated Universal Time (UTC) that the classifier was most recently updated. The field matches either `retrained` or `created`.  Returned when verbose=`true`. Might not be returned by some requests.
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// ClassifierResult : Classifier and score combination.
type ClassifierResult struct {

	// Name of the classifier.
	Name string `json:"name"`

	// ID of a classifier identified in the image.
	ClassifierID string `json:"classifier_id"`

	// Classes within the classifier.
	Classes []ClassResult `json:"classes"`
}

// Classifiers : A container for the list of classifiers.
type Classifiers struct {

	// List of classifiers.
	Classifiers []Classifier `json:"classifiers"`
}

// ClassifyOptions : The classify options.
type ClassifyOptions struct {

	// An image file (.jpg, .png) or .zip file with images. Maximum image size is 10 MB. Include no more than 20 images and limit the .zip file to 100 MB. Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters. You can also include an image with the **url** parameter.
	ImagesFile os.File `json:"images_file,omitempty"`

    // Indicates whether user set optional parameter ImagesFile
    IsImagesFileSet bool

	// The language of the output class names. The full set of languages is supported for the built-in classifier IDs: `default`, `food`, and `explicit`. The class names of custom classifiers are not translated. The response might not be in the specified language when the requested language is not supported or when there is no translation for the class name.
	AcceptLanguage string `json:"Accept-Language,omitempty"`

    // Indicates whether user set optional parameter AcceptLanguage
    IsAcceptLanguageSet bool

	// The URL of an image to analyze. Must be in .jpg, or .png format. The minimum recommended pixel density is 32X32 pixels per inch, and the maximum image size is 10 MB. You can also include images with the **images_file** parameter.
	URL string `json:"url,omitempty"`

    // Indicates whether user set optional parameter URL
    IsURLSet bool

	// The minimum score a class must have to be displayed in the response. Set the threshold to `0.0` to ignore the classification score and return all values.
	Threshold float32 `json:"threshold,omitempty"`

    // Indicates whether user set optional parameter Threshold
    IsThresholdSet bool

	// The categories of classifiers to apply. Use `IBM` to classify against the `default` general classifier, and use `me` to classify against your custom classifiers. To analyze the image against both classifier categories, set the value to both `IBM` and `me`. The built-in `default` classifier is used if both **classifier_ids** and **owners** parameters are empty. The **classifier_ids** parameter overrides **owners**, so make sure that **classifier_ids** is empty.
	Owners []string `json:"owners,omitempty"`

    // Indicates whether user set optional parameter Owners
    IsOwnersSet bool

	// Which classifiers to apply. Overrides the **owners** parameter. You can specify both custom and built-in classifier IDs. The built-in `default` classifier is used if both **classifier_ids** and **owners** parameters are empty. The following built-in classifier IDs require no training: - `default`: Returns classes from thousands of general tags. - `food`: (Beta) Enhances specificity and accuracy for images of food items. - `explicit`: (Beta) Evaluates whether the image might be pornographic.
	ClassifierIds []string `json:"classifier_ids,omitempty"`

    // Indicates whether user set optional parameter ClassifierIds
    IsClassifierIdsSet bool

	// The content type of ImagesFile.
	ImagesFileContentType string `json:"images_file_content_type,omitempty"`

    // Indicates whether user set optional parameter ImagesFileContentType
    IsImagesFileContentTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewClassifyOptions : Instantiate ClassifyOptions
func NewClassifyOptions() *ClassifyOptions {
    return &ClassifyOptions{}
}

// SetImagesFile : Allow user to set ImagesFile
func (options *ClassifyOptions) SetImagesFile(param os.File) *ClassifyOptions {
    options.ImagesFile = param
    options.IsImagesFileSet = true
    return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ClassifyOptions) SetAcceptLanguage(param string) *ClassifyOptions {
    options.AcceptLanguage = param
    options.IsAcceptLanguageSet = true
    return options
}

// SetURL : Allow user to set URL
func (options *ClassifyOptions) SetURL(param string) *ClassifyOptions {
    options.URL = param
    options.IsURLSet = true
    return options
}

// SetThreshold : Allow user to set Threshold
func (options *ClassifyOptions) SetThreshold(param float32) *ClassifyOptions {
    options.Threshold = param
    options.IsThresholdSet = true
    return options
}

// SetOwners : Allow user to set Owners
func (options *ClassifyOptions) SetOwners(param []string) *ClassifyOptions {
    options.Owners = param
    options.IsOwnersSet = true
    return options
}

// SetClassifierIds : Allow user to set ClassifierIds
func (options *ClassifyOptions) SetClassifierIds(param []string) *ClassifyOptions {
    options.ClassifierIds = param
    options.IsClassifierIdsSet = true
    return options
}

// SetImagesFileContentType : Allow user to set ImagesFileContentType
func (options *ClassifyOptions) SetImagesFileContentType(param string) *ClassifyOptions {
    options.ImagesFileContentType = param
    options.IsImagesFileContentTypeSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ClassifyOptions) SetHeaders(param map[string]string) *ClassifyOptions {
    options.Headers = param
    return options
}

// CreateClassifierOptions : The createClassifier options.
type CreateClassifierOptions struct {

	// The name of the new classifier. Encode special characters in UTF-8.
	Name string `json:"name"`

	// A .zip file of images that depict the visual subject of a class in the new classifier. You can include more than one positive example file in a call. Specify the parameter name by appending `_positive_examples` to the class name. For example, `goldenretriever_positive_examples` creates the class **goldenretriever**. Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The maximum number of images is 10,000 images or 100 MB per .zip file. Encode special characters in the file name in UTF-8.
	ClassnamePositiveExamples os.File `json:"classname_positive_examples"`

	// A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must contain a minimum of 10 images. Encode special characters in the file name in UTF-8.
	NegativeExamples os.File `json:"negative_examples,omitempty"`

    // Indicates whether user set optional parameter NegativeExamples
    IsNegativeExamplesSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewCreateClassifierOptions : Instantiate CreateClassifierOptions
func NewCreateClassifierOptions(name string, classnamePositiveExamples os.File) *CreateClassifierOptions {
    return &CreateClassifierOptions{
        Name: name,
        ClassnamePositiveExamples: classnamePositiveExamples,
    }
}

// SetName : Allow user to set Name
func (options *CreateClassifierOptions) SetName(param string) *CreateClassifierOptions {
    options.Name = param
    return options
}

// SetClassnamePositiveExamples : Allow user to set ClassnamePositiveExamples
func (options *CreateClassifierOptions) SetClassnamePositiveExamples(param os.File) *CreateClassifierOptions {
    options.ClassnamePositiveExamples = param
    return options
}

// SetNegativeExamples : Allow user to set NegativeExamples
func (options *CreateClassifierOptions) SetNegativeExamples(param os.File) *CreateClassifierOptions {
    options.NegativeExamples = param
    options.IsNegativeExamplesSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateClassifierOptions) SetHeaders(param map[string]string) *CreateClassifierOptions {
    options.Headers = param
    return options
}

// DeleteClassifierOptions : The deleteClassifier options.
type DeleteClassifierOptions struct {

	// The ID of the classifier.
	ClassifierID string `json:"classifier_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteClassifierOptions : Instantiate DeleteClassifierOptions
func NewDeleteClassifierOptions(classifierID string) *DeleteClassifierOptions {
    return &DeleteClassifierOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *DeleteClassifierOptions) SetClassifierID(param string) *DeleteClassifierOptions {
    options.ClassifierID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteClassifierOptions) SetHeaders(param map[string]string) *DeleteClassifierOptions {
    options.Headers = param
    return options
}

// DeleteUserDataOptions : The deleteUserData options.
type DeleteUserDataOptions struct {

	// The customer ID for which all data is to be deleted.
	CustomerID string `json:"customer_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
    return &DeleteUserDataOptions{
        CustomerID: customerID,
    }
}

// SetCustomerID : Allow user to set CustomerID
func (options *DeleteUserDataOptions) SetCustomerID(param string) *DeleteUserDataOptions {
    options.CustomerID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
    options.Headers = param
    return options
}

// DetectFacesOptions : The detectFaces options.
type DetectFacesOptions struct {

	// An image file (gif, .jpg, .png, .tif.) or .zip file with images. Limit the .zip file to 100 MB. You can include a maximum of 15 images in a request. Encode the image and .zip file names in UTF-8 if they contain non-ASCII characters. The service assumes UTF-8 encoding if it encounters non-ASCII characters. You can also include an image with the **url** parameter.
	ImagesFile os.File `json:"images_file,omitempty"`

    // Indicates whether user set optional parameter ImagesFile
    IsImagesFileSet bool

	// The URL of an image to analyze. Must be in .gif, .jpg, .png, or .tif format. The minimum recommended pixel density is 32X32 pixels per inch, and the maximum image size is 10 MB. Redirects are followed, so you can use a shortened URL. You can also include images with the **images_file** parameter.
	URL string `json:"url,omitempty"`

    // Indicates whether user set optional parameter URL
    IsURLSet bool

	// The content type of ImagesFile.
	ImagesFileContentType string `json:"images_file_content_type,omitempty"`

    // Indicates whether user set optional parameter ImagesFileContentType
    IsImagesFileContentTypeSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewDetectFacesOptions : Instantiate DetectFacesOptions
func NewDetectFacesOptions() *DetectFacesOptions {
    return &DetectFacesOptions{}
}

// SetImagesFile : Allow user to set ImagesFile
func (options *DetectFacesOptions) SetImagesFile(param os.File) *DetectFacesOptions {
    options.ImagesFile = param
    options.IsImagesFileSet = true
    return options
}

// SetURL : Allow user to set URL
func (options *DetectFacesOptions) SetURL(param string) *DetectFacesOptions {
    options.URL = param
    options.IsURLSet = true
    return options
}

// SetImagesFileContentType : Allow user to set ImagesFileContentType
func (options *DetectFacesOptions) SetImagesFileContentType(param string) *DetectFacesOptions {
    options.ImagesFileContentType = param
    options.IsImagesFileContentTypeSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *DetectFacesOptions) SetHeaders(param map[string]string) *DetectFacesOptions {
    options.Headers = param
    return options
}

// DetectedFaces : Results for all faces.
type DetectedFaces struct {

	// Number of images processed for the API call.
	ImagesProcessed int64 `json:"images_processed,omitempty"`

	// The images.
	Images []ImageWithFaces `json:"images"`

	// Information about what might cause less than optimal output. For example, a request sent with a corrupt .zip file and a list of image URLs will still complete, but does not return the expected output. Not returned when there is no warning.
	Warnings []WarningInfo `json:"warnings,omitempty"`
}

// ErrorInfo : Information about what might have caused a failure, such as an image that is too large. Not returned when there is no error.
type ErrorInfo struct {

	// HTTP status code.
	Code int64 `json:"code"`

	// Human-readable error description. For example, `File size limit exceeded`.
	Description string `json:"description"`

	// Codified error string. For example, `limit_exceeded`.
	ErrorID string `json:"error_id"`
}

// Face : Information about the face.
type Face struct {

	// Age information about a face.
	Age FaceAge `json:"age,omitempty"`

	// Information about the gender of the face.
	Gender FaceGender `json:"gender,omitempty"`

	// The location of the bounding box around the face.
	FaceLocation FaceLocation `json:"face_location,omitempty"`
}

// FaceAge : Age information about a face.
type FaceAge struct {

	// Estimated minimum age.
	Min int64 `json:"min,omitempty"`

	// Estimated maximum age.
	Max int64 `json:"max,omitempty"`

	// Confidence score in the range of 0 to 1. A higher score indicates greater confidence in the estimated value for the property.
	Score float32 `json:"score,omitempty"`
}

// FaceGender : Information about the gender of the face.
type FaceGender struct {

	// Gender identified by the face. For example, `MALE` or `FEMALE`.
	Gender string `json:"gender"`

	// Confidence score in the range of 0 to 1. A higher score indicates greater confidence in the estimated value for the property.
	Score float32 `json:"score,omitempty"`
}

// FaceLocation : The location of the bounding box around the face.
type FaceLocation struct {

	// Width in pixels of face region.
	Width float64 `json:"width"`

	// Height in pixels of face region.
	Height float64 `json:"height"`

	// X-position of top-left pixel of face region.
	Left float64 `json:"left"`

	// Y-position of top-left pixel of face region.
	Top float64 `json:"top"`
}

// GetClassifierOptions : The getClassifier options.
type GetClassifierOptions struct {

	// The ID of the classifier.
	ClassifierID string `json:"classifier_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetClassifierOptions : Instantiate GetClassifierOptions
func NewGetClassifierOptions(classifierID string) *GetClassifierOptions {
    return &GetClassifierOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *GetClassifierOptions) SetClassifierID(param string) *GetClassifierOptions {
    options.ClassifierID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetClassifierOptions) SetHeaders(param map[string]string) *GetClassifierOptions {
    options.Headers = param
    return options
}

// GetCoreMlModelOptions : The getCoreMlModel options.
type GetCoreMlModelOptions struct {

	// The ID of the classifier.
	ClassifierID string `json:"classifier_id"`

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewGetCoreMlModelOptions : Instantiate GetCoreMlModelOptions
func NewGetCoreMlModelOptions(classifierID string) *GetCoreMlModelOptions {
    return &GetCoreMlModelOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *GetCoreMlModelOptions) SetClassifierID(param string) *GetCoreMlModelOptions {
    options.ClassifierID = param
    return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCoreMlModelOptions) SetHeaders(param map[string]string) *GetCoreMlModelOptions {
    options.Headers = param
    return options
}

// ImageWithFaces : Information about faces in the image.
type ImageWithFaces struct {

	// Faces detected in the images.
	Faces []Face `json:"faces"`

	// Relative path of the image file if uploaded directly. Not returned when the image is passed by URL.
	Image string `json:"image,omitempty"`

	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceURL string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedURL string `json:"resolved_url,omitempty"`

	// Information about what might have caused a failure, such as an image that is too large. Not returned when there is no error.
	Error ErrorInfo `json:"error,omitempty"`
}

// ListClassifiersOptions : The listClassifiers options.
type ListClassifiersOptions struct {

	// Specify `true` to return details about the classifiers. Omit this parameter to return a brief list of classifiers.
	Verbose bool `json:"verbose,omitempty"`

    // Indicates whether user set optional parameter Verbose
    IsVerboseSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewListClassifiersOptions : Instantiate ListClassifiersOptions
func NewListClassifiersOptions() *ListClassifiersOptions {
    return &ListClassifiersOptions{}
}

// SetVerbose : Allow user to set Verbose
func (options *ListClassifiersOptions) SetVerbose(param bool) *ListClassifiersOptions {
    options.Verbose = param
    options.IsVerboseSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *ListClassifiersOptions) SetHeaders(param map[string]string) *ListClassifiersOptions {
    options.Headers = param
    return options
}

// UpdateClassifierOptions : The updateClassifier options.
type UpdateClassifierOptions struct {

	// The ID of the classifier.
	ClassifierID string `json:"classifier_id"`

	// A .zip file of images that depict the visual subject of a class in the classifier. The positive examples create or update classes in the classifier. You can include more than one positive example file in a call. Specify the parameter name by appending `_positive_examples` to the class name. For example, `goldenretriever_positive_examples` creates the class `goldenretriever`. Include at least 10 images in .jpg or .png format. The minimum recommended image resolution is 32X32 pixels. The maximum number of images is 10,000 images or 100 MB per .zip file. Encode special characters in the file name in UTF-8.
	ClassnamePositiveExamples os.File `json:"classname_positive_examples,omitempty"`

    // Indicates whether user set optional parameter ClassnamePositiveExamples
    IsClassnamePositiveExamplesSet bool

	// A .zip file of images that do not depict the visual subject of any of the classes of the new classifier. Must contain a minimum of 10 images. Encode special characters in the file name in UTF-8.
	NegativeExamples os.File `json:"negative_examples,omitempty"`

    // Indicates whether user set optional parameter NegativeExamples
    IsNegativeExamplesSet bool

    // Allows users to set headers to be GDPR compliant
    Headers map[string]string
}

// NewUpdateClassifierOptions : Instantiate UpdateClassifierOptions
func NewUpdateClassifierOptions(classifierID string) *UpdateClassifierOptions {
    return &UpdateClassifierOptions{
        ClassifierID: classifierID,
    }
}

// SetClassifierID : Allow user to set ClassifierID
func (options *UpdateClassifierOptions) SetClassifierID(param string) *UpdateClassifierOptions {
    options.ClassifierID = param
    return options
}

// SetClassnamePositiveExamples : Allow user to set ClassnamePositiveExamples
func (options *UpdateClassifierOptions) SetClassnamePositiveExamples(param os.File) *UpdateClassifierOptions {
    options.ClassnamePositiveExamples = param
    options.IsClassnamePositiveExamplesSet = true
    return options
}

// SetNegativeExamples : Allow user to set NegativeExamples
func (options *UpdateClassifierOptions) SetNegativeExamples(param os.File) *UpdateClassifierOptions {
    options.NegativeExamples = param
    options.IsNegativeExamplesSet = true
    return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateClassifierOptions) SetHeaders(param map[string]string) *UpdateClassifierOptions {
    options.Headers = param
    return options
}

// WarningInfo : Information about something that went wrong.
type WarningInfo struct {

	// Codified warning string, such as `limit_reached`.
	WarningID string `json:"warning_id"`

	// Information about the error.
	Description string `json:"description"`
}
