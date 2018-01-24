package ccc

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

func (client *Client) ListUsers(request *ListUsersRequest) (response *ListUsersResponse, err error) {
	response = CreateListUsersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListUsersWithChan(request *ListUsersRequest) (<-chan *ListUsersResponse, <-chan error) {
	responseChan := make(chan *ListUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsers(request)
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

func (client *Client) ListUsersWithCallback(request *ListUsersRequest, callback func(response *ListUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersResponse
		var err error
		defer close(result)
		response, err = client.ListUsers(request)
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

type ListUsersRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

type ListUsersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Users          struct {
		TotalCount int `json:"TotalCount" xml:"TotalCount"`
		PageNumber int `json:"PageNumber" xml:"PageNumber"`
		PageSize   int `json:"PageSize" xml:"PageSize"`
		List       struct {
			User []struct {
				UserId     string `json:"UserId" xml:"UserId"`
				RamId      string `json:"RamId" xml:"RamId"`
				InstanceId string `json:"InstanceId" xml:"InstanceId"`
				Primary    bool   `json:"Primary" xml:"Primary"`
				Detail     struct {
					LoginName   string `json:"LoginName" xml:"LoginName"`
					DisplayName string `json:"DisplayName" xml:"DisplayName"`
					Phone       string `json:"Phone" xml:"Phone"`
					Email       string `json:"Email" xml:"Email"`
					Department  string `json:"Department" xml:"Department"`
				} `json:"Detail" xml:"Detail"`
				Roles struct {
					Role []struct {
						RoleId          string `json:"RoleId" xml:"RoleId"`
						InstanceId      string `json:"InstanceId" xml:"InstanceId"`
						RoleName        string `json:"RoleName" xml:"RoleName"`
						RoleDescription string `json:"RoleDescription" xml:"RoleDescription"`
					} `json:"Role" xml:"Role"`
				} `json:"Roles" xml:"Roles"`
				SkillLevels struct {
					SkillLevel []struct {
						SkillLevelId string `json:"SkillLevelId" xml:"SkillLevelId"`
						Level        int    `json:"Level" xml:"Level"`
						Skill        struct {
							SkillGroupId          string `json:"SkillGroupId" xml:"SkillGroupId"`
							InstanceId            string `json:"InstanceId" xml:"InstanceId"`
							SkillGroupName        string `json:"SkillGroupName" xml:"SkillGroupName"`
							SkillGroupDescription string `json:"SkillGroupDescription" xml:"SkillGroupDescription"`
						} `json:"Skill" xml:"Skill"`
					} `json:"SkillLevel" xml:"SkillLevel"`
				} `json:"SkillLevels" xml:"SkillLevels"`
			} `json:"User" xml:"User"`
		} `json:"List" xml:"List"`
	} `json:"Users" xml:"Users"`
}

func CreateListUsersRequest() (request *ListUsersRequest) {
	request = &ListUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListUsers", "", "")
	return
}

func CreateListUsersResponse() (response *ListUsersResponse) {
	response = &ListUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
