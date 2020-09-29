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

// OnsGroupList invokes the ons.OnsGroupList API synchronously
func (client *Client) OnsGroupList(request *OnsGroupListRequest) (response *OnsGroupListResponse, err error) {
	response = CreateOnsGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// OnsGroupListWithChan invokes the ons.OnsGroupList API asynchronously
func (client *Client) OnsGroupListWithChan(request *OnsGroupListRequest) (<-chan *OnsGroupListResponse, <-chan error) {
	responseChan := make(chan *OnsGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsGroupList(request)
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

// OnsGroupListWithCallback invokes the ons.OnsGroupList API asynchronously
func (client *Client) OnsGroupListWithCallback(request *OnsGroupListRequest, callback func(response *OnsGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsGroupListResponse
		var err error
		defer close(result)
		response, err = client.OnsGroupList(request)
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

// OnsGroupListRequest is the request struct for api OnsGroupList
type OnsGroupListRequest struct {
	*requests.RpcRequest
	GroupId    string             `position:"Query" name:"GroupId"`
	InstanceId string             `position:"Query" name:"InstanceId"`
	GroupType  string             `position:"Query" name:"GroupType"`
	Tag        *[]OnsGroupListTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// OnsGroupListTag is a repeated param struct in OnsGroupListRequest
type OnsGroupListTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// OnsGroupListResponse is the response struct for api OnsGroupList
type OnsGroupListResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	HelpUrl   string             `json:"HelpUrl" xml:"HelpUrl"`
	Data      DataInOnsGroupList `json:"Data" xml:"Data"`
}

// CreateOnsGroupListRequest creates a request to invoke OnsGroupList API
func CreateOnsGroupListRequest() (request *OnsGroupListRequest) {
	request = &OnsGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsGroupList", "", "")
	request.Method = requests.POST
	return
}

// CreateOnsGroupListResponse creates a response to parse from OnsGroupList response
func CreateOnsGroupListResponse() (response *OnsGroupListResponse) {
	response = &OnsGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
