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

func (client *Client) DescribeLiveStreamTranscodeStreamNum(request *DescribeLiveStreamTranscodeStreamNumRequest) (response *DescribeLiveStreamTranscodeStreamNumResponse, err error) {
	response = CreateDescribeLiveStreamTranscodeStreamNumResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamTranscodeStreamNumWithChan(request *DescribeLiveStreamTranscodeStreamNumRequest) (<-chan *DescribeLiveStreamTranscodeStreamNumResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamTranscodeStreamNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamTranscodeStreamNum(request)
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

func (client *Client) DescribeLiveStreamTranscodeStreamNumWithCallback(request *DescribeLiveStreamTranscodeStreamNumRequest, callback func(response *DescribeLiveStreamTranscodeStreamNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamTranscodeStreamNumResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamTranscodeStreamNum(request)
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

type DescribeLiveStreamTranscodeStreamNumRequest struct {
	*requests.RpcRequest
	PullDomain    string           `position:"Query" name:"PullDomain"`
	PushDomain    string           `position:"Query" name:"PushDomain"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Total             int    `json:"Total" xml:"Total"`
	TranscodedNumber  int    `json:"TranscodedNumber" xml:"TranscodedNumber"`
	UntranscodeNumber int    `json:"UntranscodeNumber" xml:"UntranscodeNumber"`
}

func CreateDescribeLiveStreamTranscodeStreamNumRequest() (request *DescribeLiveStreamTranscodeStreamNumRequest) {
	request = &DescribeLiveStreamTranscodeStreamNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamTranscodeStreamNum", "", "")
	return
}

func CreateDescribeLiveStreamTranscodeStreamNumResponse() (response *DescribeLiveStreamTranscodeStreamNumResponse) {
	response = &DescribeLiveStreamTranscodeStreamNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
