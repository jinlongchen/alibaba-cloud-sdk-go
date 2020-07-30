package iot

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

// QueryClientIds invokes the iot.QueryClientIds API synchronously
// api document: https://help.aliyun.com/api/iot/queryclientids.html
func (client *Client) QueryClientIds(request *QueryClientIdsRequest) (response *QueryClientIdsResponse, err error) {
	response = CreateQueryClientIdsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryClientIdsWithChan invokes the iot.QueryClientIds API asynchronously
// api document: https://help.aliyun.com/api/iot/queryclientids.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryClientIdsWithChan(request *QueryClientIdsRequest) (<-chan *QueryClientIdsResponse, <-chan error) {
	responseChan := make(chan *QueryClientIdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryClientIds(request)
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

// QueryClientIdsWithCallback invokes the iot.QueryClientIds API asynchronously
// api document: https://help.aliyun.com/api/iot/queryclientids.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryClientIdsWithCallback(request *QueryClientIdsRequest, callback func(response *QueryClientIdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryClientIdsResponse
		var err error
		defer close(result)
		response, err = client.QueryClientIds(request)
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

// QueryClientIdsRequest is the request struct for api QueryClientIds
type QueryClientIdsRequest struct {
	*requests.RpcRequest
	AuthConfig    string `position:"Query" name:"AuthConfig"`
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryClientIdsResponse is the response struct for api QueryClientIds
type QueryClientIdsResponse struct {
	*responses.BaseResponse
	RequestId    string               `json:"RequestId" xml:"RequestId"`
	Success      bool                 `json:"Success" xml:"Success"`
	Code         string               `json:"Code" xml:"Code"`
	ErrorMessage string               `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryClientIds `json:"Data" xml:"Data"`
}

// CreateQueryClientIdsRequest creates a request to invoke QueryClientIds API
func CreateQueryClientIdsRequest() (request *QueryClientIdsRequest) {
	request = &QueryClientIdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryClientIds", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryClientIdsResponse creates a response to parse from QueryClientIds response
func CreateQueryClientIdsResponse() (response *QueryClientIdsResponse) {
	response = &QueryClientIdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}