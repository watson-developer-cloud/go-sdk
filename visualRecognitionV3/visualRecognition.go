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
package visualRecognitionV3

import (
    "bytes"
    "fmt"
    "github.com/go-openapi/strfmt"
    "os"
    "strings"
    req "github.com/parnurzeal/gorequest"
    watson "golang-sdk"
)

type VisualRecognitionV3 struct {
	client *watson.Client
}

func NewVisualRecognitionV3(creds watson.Credentials) (*VisualRecognitionV3, error) {
	client, clientErr := watson.NewClient(creds, "watson_vision_combined")

	if clientErr != nil {
		return nil, clientErr
	}

	return &VisualRecognitionV3{ client: client }, nil
}

// Classify images
func (visualRecognition *VisualRecognitionV3) Classify(imagesFile os.File, acceptLanguage string, url string, threshold float32, owners []string, classifierIds []string, imagesFileContentType string) (*watson.WatsonResponse, []error) {
    path := "/v3/classify"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager


    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetClassifyResult(response *watson.WatsonResponse) *ClassifiedImages {
    result, ok := response.Result.(*ClassifiedImages)

    if ok {
        return result
    }

    return nil
}

// Detect faces in images
func (visualRecognition *VisualRecognitionV3) DetectFaces(imagesFile os.File, url string, imagesFileContentType string) (*watson.WatsonResponse, []error) {
    path := "/v3/detect_faces"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager


    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetDetectFacesResult(response *watson.WatsonResponse) *DetectedFaces {
    result, ok := response.Result.(*DetectedFaces)

    if ok {
        return result
    }

    return nil
}

// Create a classifier
func (visualRecognition *VisualRecognitionV3) CreateClassifier(name string, classnamePositiveExamples os.File, negativeExamples os.File) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager


    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetCreateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// Delete a classifier
func (visualRecognition *VisualRecognitionV3) DeleteClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Delete(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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


// Retrieve classifier details
func (visualRecognition *VisualRecognitionV3) GetClassifier(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Get(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetGetClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// Retrieve a list of classifiers
func (visualRecognition *VisualRecognitionV3) ListClassifiers(verbose bool) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager


    request := req.New().Get(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)

    request.Query("verbose=" + fmt.Sprint(verbose))

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

func GetListClassifiersResult(response *watson.WatsonResponse) *Classifiers {
    result, ok := response.Result.(*Classifiers)

    if ok {
        return result
    }

    return nil
}

// Update a classifier
func (visualRecognition *VisualRecognitionV3) UpdateClassifier(classifierID string, classnamePositiveExamples os.File, negativeExamples os.File) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Post(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetUpdateClassifierResult(response *watson.WatsonResponse) *Classifier {
    result, ok := response.Result.(*Classifier)

    if ok {
        return result
    }

    return nil
}

// Retrieve a Core ML model of a classifier
func (visualRecognition *VisualRecognitionV3) GetCoreMlModel(classifierID string) (*watson.WatsonResponse, []error) {
    path := "/v3/classifiers/{classifier_id}/core_ml_model"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager

    path = strings.Replace(path, "{classifier_id}", classifierID, 1)

    request := req.New().Get(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)


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

func GetGetCoreMlModelResult(response *watson.WatsonResponse) *os.File {
    result, ok := response.Result.(*os.File)

    if ok {
        return result
    }

    return nil
}

// Delete labeled data
func (visualRecognition *VisualRecognitionV3) DeleteUserData(customerID string) (*watson.WatsonResponse, []error) {
    path := "/v3/user_data"
    creds := visualRecognition.client.Creds
    useTM := visualRecognition.client.UseTM
    tokenManager := visualRecognition.client.TokenManager


    request := req.New().Delete(creds.ServiceURL + path).
        Set("Accept", "application/json").
        Query("version=" + creds.Version)

    request.Query("customer_id=" + fmt.Sprint(customerID))

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



type Class struct {

	// The name of the class.
	ClassName string `json:"class_name"`
}

type ClassResult struct {

	// Name of the class.
	ClassName string `json:"class_name"`

	// Confidence score for the property in the range of 0 to 1. A higher score indicates greater likelihood that the class is depicted in the image. The default threshold for returning scores from a classifier is 0.5.
	Score float32 `json:"score,omitempty"`

	// Knowledge graph of the property. For example, `/fruit/pome/apple/eating apple/Granny Smith`. Included only if identified.
	TypeHierarchy string `json:"type_hierarchy,omitempty"`
}

type ClassifiedImage struct {

	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceUrl string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedUrl string `json:"resolved_url,omitempty"`

	// Relative path of the image file if uploaded directly. Not returned when the image is passed by URL.
	Image string `json:"image,omitempty"`

	// Information about what might have caused a failure, such as an image that is too large. Not returned when there is no error.
	Error ErrorInfo `json:"error,omitempty"`

	// The classifiers.
	Classifiers []ClassifierResult `json:"classifiers"`
}

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

type Classifier struct {

	// ID of a classifier identified in the image.
	ClassifierId string `json:"classifier_id"`

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

type ClassifierResult struct {

	// Name of the classifier.
	Name string `json:"name"`

	// ID of a classifier identified in the image.
	ClassifierId string `json:"classifier_id"`

	// Classes within the classifier.
	Classes []ClassResult `json:"classes"`
}

type Classifiers struct {

	// List of classifiers.
	Classifiers []Classifier `json:"classifiers"`
}

type DetectedFaces struct {

	// Number of images processed for the API call.
	ImagesProcessed int64 `json:"images_processed,omitempty"`

	// The images.
	Images []ImageWithFaces `json:"images"`

	// Information about what might cause less than optimal output. For example, a request sent with a corrupt .zip file and a list of image URLs will still complete, but does not return the expected output. Not returned when there is no warning.
	Warnings []WarningInfo `json:"warnings,omitempty"`
}

type ErrorInfo struct {

	// HTTP status code.
	Code int64 `json:"code"`

	// Human-readable error description. For example, `File size limit exceeded`.
	Description string `json:"description"`

	// Codified error string. For example, `limit_exceeded`.
	ErrorId string `json:"error_id"`
}

type Face struct {

	// Age information about a face.
	Age FaceAge `json:"age,omitempty"`

	// Information about the gender of the face.
	Gender FaceGender `json:"gender,omitempty"`

	// The location of the bounding box around the face.
	FaceLocation FaceLocation `json:"face_location,omitempty"`
}

type FaceAge struct {

	// Estimated minimum age.
	Min int64 `json:"min,omitempty"`

	// Estimated maximum age.
	Max int64 `json:"max,omitempty"`

	// Confidence score in the range of 0 to 1. A higher score indicates greater confidence in the estimated value for the property.
	Score float32 `json:"score,omitempty"`
}

type FaceGender struct {

	// Gender identified by the face. For example, `MALE` or `FEMALE`.
	Gender string `json:"gender"`

	// Confidence score in the range of 0 to 1. A higher score indicates greater confidence in the estimated value for the property.
	Score float32 `json:"score,omitempty"`
}

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

type ImageWithFaces struct {

	// Faces detected in the images.
	Faces []Face `json:"faces"`

	// Relative path of the image file if uploaded directly. Not returned when the image is passed by URL.
	Image string `json:"image,omitempty"`

	// Source of the image before any redirects. Not returned when the image is uploaded.
	SourceUrl string `json:"source_url,omitempty"`

	// Fully resolved URL of the image after redirects are followed. Not returned when the image is uploaded.
	ResolvedUrl string `json:"resolved_url,omitempty"`

	// Information about what might have caused a failure, such as an image that is too large. Not returned when there is no error.
	Error ErrorInfo `json:"error,omitempty"`
}

type WarningInfo struct {

	// Codified warning string, such as `limit_reached`.
	WarningId string `json:"warning_id"`

	// Information about the error.
	Description string `json:"description"`
}
