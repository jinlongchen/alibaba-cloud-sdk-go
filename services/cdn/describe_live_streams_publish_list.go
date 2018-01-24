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

func (client *Client) DescribeLiveStreamsPublishList(request *DescribeLiveStreamsPublishListRequest) (response *DescribeLiveStreamsPublishListResponse, err error) {
	response = CreateDescribeLiveStreamsPublishListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamsPublishListWithChan(request *DescribeLiveStreamsPublishListRequest) (<-chan *DescribeLiveStreamsPublishListResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamsPublishListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamsPublishList(request)
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

func (client *Client) DescribeLiveStreamsPublishListWithCallback(request *DescribeLiveStreamsPublishListRequest, callback func(response *DescribeLiveStreamsPublishListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamsPublishListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamsPublishList(request)
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

type DescribeLiveStreamsPublishListRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StreamName    string           `position:"Query" name:"StreamName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamsPublishListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PublishInfo struct {
		LiveStreamPublishInfo []struct {
			DomainName   string `json:"DomainName" xml:"DomainName"`
			AppName      string `json:"AppName" xml:"AppName"`
			StreamName   string `json:"StreamName" xml:"StreamName"`
			StreamUrl    string `json:"StreamUrl" xml:"StreamUrl"`
			PublishTime  string `json:"PublishTime" xml:"PublishTime"`
			StopTime     string `json:"StopTime" xml:"StopTime"`
			PublishUrl   string `json:"PublishUrl" xml:"PublishUrl"`
			ClientAddr   string `json:"ClientAddr" xml:"ClientAddr"`
			EdgeNodeAddr string `json:"EdgeNodeAddr" xml:"EdgeNodeAddr"`
		} `json:"LiveStreamPublishInfo" xml:"LiveStreamPublishInfo"`
	} `json:"PublishInfo" xml:"PublishInfo"`
}

func CreateDescribeLiveStreamsPublishListRequest() (request *DescribeLiveStreamsPublishListRequest) {
	request = &DescribeLiveStreamsPublishListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamsPublishList", "", "")
	return
}

func CreateDescribeLiveStreamsPublishListResponse() (response *DescribeLiveStreamsPublishListResponse) {
	response = &DescribeLiveStreamsPublishListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
