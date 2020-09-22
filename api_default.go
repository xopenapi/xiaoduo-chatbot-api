/*
 * Xiaoduo Chatbot open api
 *
 * xiaoduo chatbot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package xiaoduo

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// MatchQuestionOpts Optional parameters for the method 'MatchQuestion'
type MatchQuestionOpts struct {
    UserId optional.String
    Nick optional.String
}

/*
MatchQuestion Method for MatchQuestion
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param unitId 企业ID
 * @param channelId 渠道ID
 * @param salt 签名加盐,生成方式如下
 * @param sign 签名,生成方式如下
 * @param question 用户提问
 * @param optional nil or *MatchQuestionOpts - Optional Parameters:
 * @param "UserId" (optional.String) -  用户ID,要求全局唯一
 * @param "Nick" (optional.String) -  用户昵称
@return MatchQuestionRsp
*/
func (a *DefaultApiService) MatchQuestion(ctx _context.Context, unitId interface{}, channelId interface{}, salt string, sign string, question string, localVarOptionals *MatchQuestionOpts) (MatchQuestionRsp, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MatchQuestionRsp
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/match_question"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/formData"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("unit_id", parameterToString(unitId, ""))
	localVarFormParams.Add("channel_id", parameterToString(channelId, ""))
	localVarFormParams.Add("salt", parameterToString(salt, ""))
	localVarFormParams.Add("sign", parameterToString(sign, ""))
	if localVarOptionals != nil && localVarOptionals.UserId.IsSet() {
		localVarFormParams.Add("user_id", parameterToString(localVarOptionals.UserId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nick.IsSet() {
		localVarFormParams.Add("nick", parameterToString(localVarOptionals.Nick.Value(), ""))
	}
	localVarFormParams.Add("question", parameterToString(question, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
