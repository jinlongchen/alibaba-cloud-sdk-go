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

func (client *Client) BindInputBucket(request *BindInputBucketRequest) (response *BindInputBucketResponse, err error) {
	response = CreateBindInputBucketResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) BindInputBucketWithChan(request *BindInputBucketRequest) (<-chan *BindInputBucketResponse, <-chan error) {
	responseChan := make(chan *BindInputBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindInputBucket(request)
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

func (client *Client) BindInputBucketWithCallback(request *BindInputBucketRequest, callback func(response *BindInputBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindInputBucketResponse
		var err error
		defer close(result)
		response, err = client.BindInputBucket(request)
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

type BindInputBucketRequest struct {
	*requests.RpcRequest
	Bucket               string           `position:"Query" name:"Bucket"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	RoleArn              string           `position:"Query" name:"RoleArn"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type BindInputBucketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateBindInputBucketRequest() (request *BindInputBucketRequest) {
	request = &BindInputBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "BindInputBucket", "mts", "openAPI")
	return
}

func CreateBindInputBucketResponse() (response *BindInputBucketResponse) {
	response = &BindInputBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
