package ecs

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

func (client *Client) DescribeNetworkInterfaces(request *DescribeNetworkInterfacesRequest) (response *DescribeNetworkInterfacesResponse, err error) {
	response = CreateDescribeNetworkInterfacesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeNetworkInterfacesWithChan(request *DescribeNetworkInterfacesRequest) (<-chan *DescribeNetworkInterfacesResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkInterfacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkInterfaces(request)
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

func (client *Client) DescribeNetworkInterfacesWithCallback(request *DescribeNetworkInterfacesRequest, callback func(response *DescribeNetworkInterfacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkInterfacesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkInterfaces(request)
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

type DescribeNetworkInterfacesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupId      string           `position:"Query" name:"SecurityGroupId"`
	Type                 string           `position:"Query" name:"Type"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NetworkInterfaceName string           `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	PrimaryIpAddress     string           `position:"Query" name:"PrimaryIpAddress"`
	NetworkInterfaceId   *[]string        `position:"Query" name:"NetworkInterfaceId"  type:"Repeated"`
}

type DescribeNetworkInterfacesResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	TotalCount           int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber           int    `json:"PageNumber" xml:"PageNumber"`
	PageSize             int    `json:"PageSize" xml:"PageSize"`
	NetworkInterfaceSets struct {
		NetworkInterfaceSet []struct {
			NetworkInterfaceId   string `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
			Status               string `json:"Status" xml:"Status"`
			Type                 string `json:"Type" xml:"Type"`
			VpcId                string `json:"VpcId" xml:"VpcId"`
			VSwitchId            string `json:"VSwitchId" xml:"VSwitchId"`
			ZoneId               string `json:"ZoneId" xml:"ZoneId"`
			PrivateIpAddress     string `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
			MacAddress           string `json:"MacAddress" xml:"MacAddress"`
			NetworkInterfaceName string `json:"NetworkInterfaceName" xml:"NetworkInterfaceName"`
			Description          string `json:"Description" xml:"Description"`
			InstanceId           string `json:"InstanceId" xml:"InstanceId"`
			CreationTime         string `json:"CreationTime" xml:"CreationTime"`
			SecurityGroupIds     struct {
				SecurityGroupId []string `json:"SecurityGroupId" xml:"SecurityGroupId"`
			} `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
			AssociatedPublicIp struct {
				PublicIpAddress string `json:"PublicIpAddress" xml:"PublicIpAddress"`
				AllocationId    string `json:"AllocationId" xml:"AllocationId"`
			} `json:"AssociatedPublicIp" xml:"AssociatedPublicIp"`
			PrivateIpSets struct {
				PrivateIpSet []struct {
					PrivateIpAddress    string `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
					Primary             bool   `json:"Primary" xml:"Primary"`
					AssociatedPublicIp1 struct {
						PublicIpAddress string `json:"PublicIpAddress" xml:"PublicIpAddress"`
						AllocationId    string `json:"AllocationId" xml:"AllocationId"`
					} `json:"AssociatedPublicIp" xml:"AssociatedPublicIp"`
				} `json:"PrivateIpSet" xml:"PrivateIpSet"`
			} `json:"PrivateIpSets" xml:"PrivateIpSets"`
		} `json:"NetworkInterfaceSet" xml:"NetworkInterfaceSet"`
	} `json:"NetworkInterfaceSets" xml:"NetworkInterfaceSets"`
}

func CreateDescribeNetworkInterfacesRequest() (request *DescribeNetworkInterfacesRequest) {
	request = &DescribeNetworkInterfacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfaces", "ecs", "openAPI")
	return
}

func CreateDescribeNetworkInterfacesResponse() (response *DescribeNetworkInterfacesResponse) {
	response = &DescribeNetworkInterfacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
