package cloudcallcenter

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

// ListRolesOfUser invokes the cloudcallcenter.ListRolesOfUser API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrolesofuser.html
func (client *Client) ListRolesOfUser(request *ListRolesOfUserRequest) (response *ListRolesOfUserResponse, err error) {
	response = CreateListRolesOfUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListRolesOfUserWithChan invokes the cloudcallcenter.ListRolesOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrolesofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRolesOfUserWithChan(request *ListRolesOfUserRequest) (<-chan *ListRolesOfUserResponse, <-chan error) {
	responseChan := make(chan *ListRolesOfUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRolesOfUser(request)
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

// ListRolesOfUserWithCallback invokes the cloudcallcenter.ListRolesOfUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrolesofuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRolesOfUserWithCallback(request *ListRolesOfUserRequest, callback func(response *ListRolesOfUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRolesOfUserResponse
		var err error
		defer close(result)
		response, err = client.ListRolesOfUser(request)
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

// ListRolesOfUserRequest is the request struct for api ListRolesOfUser
type ListRolesOfUserRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

// ListRolesOfUserResponse is the response struct for api ListRolesOfUser
type ListRolesOfUserResponse struct {
	*responses.BaseResponse
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Success        bool                   `json:"Success" xml:"Success"`
	Code           string                 `json:"Code" xml:"Code"`
	Message        string                 `json:"Message" xml:"Message"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Roles          RolesInListRolesOfUser `json:"Roles" xml:"Roles"`
}

// CreateListRolesOfUserRequest creates a request to invoke ListRolesOfUser API
func CreateListRolesOfUserRequest() (request *ListRolesOfUserRequest) {
	request = &ListRolesOfUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListRolesOfUser", "", "")
	request.Method = requests.POST
	return
}

// CreateListRolesOfUserResponse creates a response to parse from ListRolesOfUser response
func CreateListRolesOfUserResponse() (response *ListRolesOfUserResponse) {
	response = &ListRolesOfUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}