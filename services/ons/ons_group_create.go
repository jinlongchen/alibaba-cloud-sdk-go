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

// OnsGroupCreate invokes the ons.OnsGroupCreate API synchronously
func (client *Client) OnsGroupCreate(request *OnsGroupCreateRequest) (response *OnsGroupCreateResponse, err error) {
	response = CreateOnsGroupCreateResponse()
	err = client.DoAction(request, response)
	return
}

// OnsGroupCreateWithChan invokes the ons.OnsGroupCreate API asynchronously
func (client *Client) OnsGroupCreateWithChan(request *OnsGroupCreateRequest) (<-chan *OnsGroupCreateResponse, <-chan error) {
	responseChan := make(chan *OnsGroupCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsGroupCreate(request)
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

// OnsGroupCreateWithCallback invokes the ons.OnsGroupCreate API asynchronously
func (client *Client) OnsGroupCreateWithCallback(request *OnsGroupCreateRequest, callback func(response *OnsGroupCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsGroupCreateResponse
		var err error
		defer close(result)
		response, err = client.OnsGroupCreate(request)
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

// OnsGroupCreateRequest is the request struct for api OnsGroupCreate
type OnsGroupCreateRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	Remark     string `position:"Query" name:"Remark"`
	InstanceId string `position:"Query" name:"InstanceId"`
	GroupType  string `position:"Query" name:"GroupType"`
}

// OnsGroupCreateResponse is the response struct for api OnsGroupCreate
type OnsGroupCreateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsGroupCreateRequest creates a request to invoke OnsGroupCreate API
func CreateOnsGroupCreateRequest() (request *OnsGroupCreateRequest) {
	request = &OnsGroupCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsGroupCreate", "", "")
	request.Method = requests.POST
	return
}

// CreateOnsGroupCreateResponse creates a response to parse from OnsGroupCreate response
func CreateOnsGroupCreateResponse() (response *OnsGroupCreateResponse) {
	response = &OnsGroupCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
