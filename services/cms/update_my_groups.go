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

func (client *Client) UpdateMyGroups(request *UpdateMyGroupsRequest) (response *UpdateMyGroupsResponse, err error) {
	response = CreateUpdateMyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateMyGroupsWithChan(request *UpdateMyGroupsRequest) (<-chan *UpdateMyGroupsResponse, <-chan error) {
	responseChan := make(chan *UpdateMyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMyGroups(request)
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

func (client *Client) UpdateMyGroupsWithCallback(request *UpdateMyGroupsRequest, callback func(response *UpdateMyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMyGroupsResponse
		var err error
		defer close(result)
		response, err = client.UpdateMyGroups(request)
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

type UpdateMyGroupsRequest struct {
	*requests.RpcRequest
	GroupName     string           `position:"Query" name:"GroupName"`
	GroupId       string           `position:"Query" name:"GroupId"`
	Type          string           `position:"Query" name:"Type"`
	BindUrls      string           `position:"Query" name:"BindUrls"`
	ContactGroups string           `position:"Query" name:"ContactGroups"`
	ServiceId     requests.Integer `position:"Query" name:"ServiceId"`
}

type UpdateMyGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

func CreateUpdateMyGroupsRequest() (request *UpdateMyGroupsRequest) {
	request = &UpdateMyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "UpdateMyGroups", "cms", "openAPI")
	return
}

func CreateUpdateMyGroupsResponse() (response *UpdateMyGroupsResponse) {
	response = &UpdateMyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
