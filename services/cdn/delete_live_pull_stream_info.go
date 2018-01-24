package cdn

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

func (client *Client) DeleteLivePullStreamInfo(request *DeleteLivePullStreamInfoRequest) (response *DeleteLivePullStreamInfoResponse, err error) {
	response = CreateDeleteLivePullStreamInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteLivePullStreamInfoWithChan(request *DeleteLivePullStreamInfoRequest) (<-chan *DeleteLivePullStreamInfoResponse, <-chan error) {
	responseChan := make(chan *DeleteLivePullStreamInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLivePullStreamInfo(request)
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

func (client *Client) DeleteLivePullStreamInfoWithCallback(request *DeleteLivePullStreamInfoRequest, callback func(response *DeleteLivePullStreamInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLivePullStreamInfoResponse
		var err error
		defer close(result)
		response, err = client.DeleteLivePullStreamInfo(request)
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

type DeleteLivePullStreamInfoRequest struct {
	*requests.RpcRequest
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DeleteLivePullStreamInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteLivePullStreamInfoRequest() (request *DeleteLivePullStreamInfoRequest) {
	request = &DeleteLivePullStreamInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLivePullStreamInfo", "", "")
	return
}

func CreateDeleteLivePullStreamInfoResponse() (response *DeleteLivePullStreamInfoResponse) {
	response = &DeleteLivePullStreamInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
