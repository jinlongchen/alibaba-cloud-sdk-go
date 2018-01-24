package vpc

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

func (client *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
	response = CreateModifyNatGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyNatGatewayAttributeWithChan(request *ModifyNatGatewayAttributeRequest) (<-chan *ModifyNatGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyNatGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNatGatewayAttribute(request)
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

func (client *Client) ModifyNatGatewayAttributeWithCallback(request *ModifyNatGatewayAttributeRequest, callback func(response *ModifyNatGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNatGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyNatGatewayAttribute(request)
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

type ModifyNatGatewayAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	Name                 string           `position:"Query" name:"Name"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
}

type ModifyNatGatewayAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyNatGatewayAttributeRequest() (request *ModifyNatGatewayAttributeRequest) {
	request = &ModifyNatGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewayAttribute", "vpc", "openAPI")
	return
}

func CreateModifyNatGatewayAttributeResponse() (response *ModifyNatGatewayAttributeResponse) {
	response = &ModifyNatGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
