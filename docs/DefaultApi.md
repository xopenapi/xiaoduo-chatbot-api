# \DefaultApi

All URIs are relative to *https://cvd.xiaoduoai.com/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MatchQuestion**](DefaultApi.md#MatchQuestion) | **Post** /match_question | 



## MatchQuestion

> MatchQuestionRsp MatchQuestion(ctx, unitId, channelId, salt, sign, question, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unitId** | [**interface{}**](interface{}.md)| 企业ID | 
**channelId** | [**interface{}**](interface{}.md)| 渠道ID | 
**salt** | **string**| 签名加盐,生成方式如下 | 
**sign** | **string**| 签名,生成方式如下 | 
**question** | **string**| 用户提问 | 
 **optional** | ***MatchQuestionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MatchQuestionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| 用户ID,要求全局唯一 | 
 **nick** | **optional.String**| 用户昵称 | 

### Return type

[**MatchQuestionRsp**](MatchQuestionRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/formData
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

