package alidns

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

// CheckDomainRecord invokes the alidns.CheckDomainRecord API synchronously
// api document: https://help.aliyun.com/api/alidns/checkdomainrecord.html
func (client *Client) CheckDomainRecord(request *CheckDomainRecordRequest) (response *CheckDomainRecordResponse, err error) {
	response = CreateCheckDomainRecordResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDomainRecordWithChan invokes the alidns.CheckDomainRecord API asynchronously
// api document: https://help.aliyun.com/api/alidns/checkdomainrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainRecordWithChan(request *CheckDomainRecordRequest) (<-chan *CheckDomainRecordResponse, <-chan error) {
	responseChan := make(chan *CheckDomainRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDomainRecord(request)
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

// CheckDomainRecordWithCallback invokes the alidns.CheckDomainRecord API asynchronously
// api document: https://help.aliyun.com/api/alidns/checkdomainrecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckDomainRecordWithCallback(request *CheckDomainRecordRequest, callback func(response *CheckDomainRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDomainRecordResponse
		var err error
		defer close(result)
		response, err = client.CheckDomainRecord(request)
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

// CheckDomainRecordRequest is the request struct for api CheckDomainRecord
type CheckDomainRecordRequest struct {
	*requests.RpcRequest
	RR           string `position:"Query" name:"RR"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	Lang         string `position:"Query" name:"Lang"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
}

// CheckDomainRecordResponse is the response struct for api CheckDomainRecord
type CheckDomainRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IsExist   bool   `json:"IsExist" xml:"IsExist"`
}

// CreateCheckDomainRecordRequest creates a request to invoke CheckDomainRecord API
func CreateCheckDomainRecordRequest() (request *CheckDomainRecordRequest) {
	request = &CheckDomainRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "CheckDomainRecord", "alidns", "openAPI")
	return
}

// CreateCheckDomainRecordResponse creates a response to parse from CheckDomainRecord response
func CreateCheckDomainRecordResponse() (response *CheckDomainRecordResponse) {
	response = &CheckDomainRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
