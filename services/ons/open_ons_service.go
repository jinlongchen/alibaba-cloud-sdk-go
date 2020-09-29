package ons

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

// OpenOnsService invokes the ons.OpenOnsService API synchronously
func (client *Client) OpenOnsService(request *OpenOnsServiceRequest) (response *OpenOnsServiceResponse, err error) {
	response = CreateOpenOnsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenOnsServiceWithChan invokes the ons.OpenOnsService API asynchronously
func (client *Client) OpenOnsServiceWithChan(request *OpenOnsServiceRequest) (<-chan *OpenOnsServiceResponse, <-chan error) {
	responseChan := make(chan *OpenOnsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenOnsService(request)
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

// OpenOnsServiceWithCallback invokes the ons.OpenOnsService API asynchronously
func (client *Client) OpenOnsServiceWithCallback(request *OpenOnsServiceRequest, callback func(response *OpenOnsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenOnsServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenOnsService(request)
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

// OpenOnsServiceRequest is the request struct for api OpenOnsService
type OpenOnsServiceRequest struct {
	*requests.RpcRequest
}

// OpenOnsServiceResponse is the response struct for api OpenOnsService
type OpenOnsServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateOpenOnsServiceRequest creates a request to invoke OpenOnsService API
func CreateOpenOnsServiceRequest() (request *OpenOnsServiceRequest) {
	request = &OpenOnsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OpenOnsService", "", "")
	request.Method = requests.POST
	return
}

// CreateOpenOnsServiceResponse creates a response to parse from OpenOnsService response
func CreateOpenOnsServiceResponse() (response *OpenOnsServiceResponse) {
	response = &OpenOnsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
