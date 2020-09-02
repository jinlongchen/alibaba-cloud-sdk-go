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

// GetSmsConfig invokes the cloudcallcenter.GetSmsConfig API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getsmsconfig.html
func (client *Client) GetSmsConfig(request *GetSmsConfigRequest) (response *GetSmsConfigResponse, err error) {
	response = CreateGetSmsConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetSmsConfigWithChan invokes the cloudcallcenter.GetSmsConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getsmsconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSmsConfigWithChan(request *GetSmsConfigRequest) (<-chan *GetSmsConfigResponse, <-chan error) {
	responseChan := make(chan *GetSmsConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSmsConfig(request)
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

// GetSmsConfigWithCallback invokes the cloudcallcenter.GetSmsConfig API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getsmsconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSmsConfigWithCallback(request *GetSmsConfigRequest, callback func(response *GetSmsConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSmsConfigResponse
		var err error
		defer close(result)
		response, err = client.GetSmsConfig(request)
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

// GetSmsConfigRequest is the request struct for api GetSmsConfig
type GetSmsConfigRequest struct {
	*requests.RpcRequest
	InstanceId string    `position:"Query" name:"InstanceId"`
	Scenario   *[]string `position:"Query" name:"Scenario"  type:"Repeated"`
}

// GetSmsConfigResponse is the response struct for api GetSmsConfig
type GetSmsConfigResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SmsConfigs     SmsConfigs `json:"SmsConfigs" xml:"SmsConfigs"`
}

// CreateGetSmsConfigRequest creates a request to invoke GetSmsConfig API
func CreateGetSmsConfigRequest() (request *GetSmsConfigRequest) {
	request = &GetSmsConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetSmsConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSmsConfigResponse creates a response to parse from GetSmsConfig response
func CreateGetSmsConfigResponse() (response *GetSmsConfigResponse) {
	response = &GetSmsConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}