package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// PutGroupDynamicRule invokes the cms.PutGroupDynamicRule API synchronously
// api document: https://help.aliyun.com/api/cms/putgroupdynamicrule.html
func (client *Client) PutGroupDynamicRule(request *PutGroupDynamicRuleRequest) (response *PutGroupDynamicRuleResponse, err error) {
	response = CreatePutGroupDynamicRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PutGroupDynamicRuleWithChan invokes the cms.PutGroupDynamicRule API asynchronously
// api document: https://help.aliyun.com/api/cms/putgroupdynamicrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutGroupDynamicRuleWithChan(request *PutGroupDynamicRuleRequest) (<-chan *PutGroupDynamicRuleResponse, <-chan error) {
	responseChan := make(chan *PutGroupDynamicRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutGroupDynamicRule(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// PutGroupDynamicRuleWithCallback invokes the cms.PutGroupDynamicRule API asynchronously
// api document: https://help.aliyun.com/api/cms/putgroupdynamicrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutGroupDynamicRuleWithCallback(request *PutGroupDynamicRuleRequest, callback func(response *PutGroupDynamicRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutGroupDynamicRuleResponse
		var err error
		defer close(result)
		response, err = client.PutGroupDynamicRule(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// PutGroupDynamicRuleRequest is the request struct for api PutGroupDynamicRule
type PutGroupDynamicRuleRequest struct {
	*requests.RpcRequest
	GroupId            requests.Integer `position:"Query" name:"GroupId"`
	GroupRuleArrayJson string           `position:"Query" name:"GroupRuleArrayJson"`
}

// PutGroupDynamicRuleResponse is the response struct for api PutGroupDynamicRule
type PutGroupDynamicRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreatePutGroupDynamicRuleRequest creates a request to invoke PutGroupDynamicRule API
func CreatePutGroupDynamicRuleRequest() (request *PutGroupDynamicRuleRequest) {
	request = &PutGroupDynamicRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "PutGroupDynamicRule", "cms", "openAPI")
	return
}

// CreatePutGroupDynamicRuleResponse creates a response to parse from PutGroupDynamicRule response
func CreatePutGroupDynamicRuleResponse() (response *PutGroupDynamicRuleResponse) {
	response = &PutGroupDynamicRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}