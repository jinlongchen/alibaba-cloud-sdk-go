package mts

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

func (client *Client) AddPornPipeline(request *AddPornPipelineRequest) (response *AddPornPipelineResponse, err error) {
	response = CreateAddPornPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddPornPipelineWithChan(request *AddPornPipelineRequest) (<-chan *AddPornPipelineResponse, <-chan error) {
	responseChan := make(chan *AddPornPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPornPipeline(request)
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

func (client *Client) AddPornPipelineWithCallback(request *AddPornPipelineRequest, callback func(response *AddPornPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPornPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddPornPipeline(request)
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

type AddPornPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
}

type AddPornPipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		Priority     int    `json:"Priority" xml:"Priority"`
		State        string `json:"State" xml:"State"`
		NotifyConfig struct {
			Topic string `json:"Topic" xml:"Topic"`
			Queue string `json:"Queue" xml:"Queue"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"Pipeline" xml:"Pipeline"`
}

func CreateAddPornPipelineRequest() (request *AddPornPipelineRequest) {
	request = &AddPornPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddPornPipeline", "mts", "openAPI")
	return
}

func CreateAddPornPipelineResponse() (response *AddPornPipelineResponse) {
	response = &AddPornPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
