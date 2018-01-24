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

func (client *Client) DescribeRouterInterfacesForGlobal(request *DescribeRouterInterfacesForGlobalRequest) (response *DescribeRouterInterfacesForGlobalResponse, err error) {
	response = CreateDescribeRouterInterfacesForGlobalResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRouterInterfacesForGlobalWithChan(request *DescribeRouterInterfacesForGlobalRequest) (<-chan *DescribeRouterInterfacesForGlobalResponse, <-chan error) {
	responseChan := make(chan *DescribeRouterInterfacesForGlobalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouterInterfacesForGlobal(request)
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

func (client *Client) DescribeRouterInterfacesForGlobalWithCallback(request *DescribeRouterInterfacesForGlobalRequest, callback func(response *DescribeRouterInterfacesForGlobalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouterInterfacesForGlobalResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouterInterfacesForGlobal(request)
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

type DescribeRouterInterfacesForGlobalRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Status               string           `position:"Query" name:"Status"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeRouterInterfacesForGlobalResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Code               string `json:"Code" xml:"Code"`
	Message            string `json:"Message" xml:"Message"`
	Desc               string `json:"desc" xml:"desc"`
	Success            bool   `json:"Success" xml:"Success"`
	PageSize           int    `json:"PageSize" xml:"PageSize"`
	PageNumber         int    `json:"PageNumber" xml:"PageNumber"`
	TotalCount         int    `json:"TotalCount" xml:"TotalCount"`
	RouterInterfaceSet struct {
		RouterInterfaceType []struct {
			BusinessStatus                  string `json:"BusinessStatus" xml:"BusinessStatus"`
			AccessPointId                   string `json:"AccessPointId" xml:"AccessPointId"`
			ChargeType                      string `json:"ChargeType" xml:"ChargeType"`
			ConnectedTime                   string `json:"ConnectedTime" xml:"ConnectedTime"`
			CreationTime                    string `json:"CreationTime" xml:"CreationTime"`
			RouterInterfaceId               string `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
			OppositeInterfaceBusinessStatus string `json:"OppositeInterfaceBusinessStatus" xml:"OppositeInterfaceBusinessStatus"`
			OppositeInterfaceId             string `json:"OppositeInterfaceId" xml:"OppositeInterfaceId"`
			OppositeInterfaceOwnerId        int    `json:"OppositeInterfaceOwnerId" xml:"OppositeInterfaceOwnerId"`
			OppositeInterfaceSpec           string `json:"OppositeInterfaceSpec" xml:"OppositeInterfaceSpec"`
			OppositeInterfaceStatus         string `json:"OppositeInterfaceStatus" xml:"OppositeInterfaceStatus"`
			OppositeRegionId                string `json:"OppositeRegionId" xml:"OppositeRegionId"`
			OppositeAccessPointId           string `json:"OppositeAccessPointId" xml:"OppositeAccessPointId"`
			OppositeRouterId                string `json:"OppositeRouterId" xml:"OppositeRouterId"`
			OppositeRouterType              string `json:"OppositeRouterType" xml:"OppositeRouterType"`
			OppositeVpcInstanceId           string `json:"OppositeVpcInstanceId" xml:"OppositeVpcInstanceId"`
			RegionId                        string `json:"RegionId" xml:"RegionId"`
			Role                            string `json:"Role" xml:"Role"`
			RouterId                        string `json:"RouterId" xml:"RouterId"`
			RouterType                      string `json:"RouterType" xml:"RouterType"`
			Spec                            string `json:"Spec" xml:"Spec"`
			Status                          string `json:"Status" xml:"Status"`
			VpcInstanceId                   string `json:"VpcInstanceId" xml:"VpcInstanceId"`
			Name                            string `json:"Name" xml:"Name"`
			Description                     string `json:"Description" xml:"Description"`
			HealthCheckSourceIp             string `json:"HealthCheckSourceIp" xml:"HealthCheckSourceIp"`
			HealthCheckTargetIp             string `json:"HealthCheckTargetIp" xml:"HealthCheckTargetIp"`
		} `json:"RouterInterfaceType" xml:"RouterInterfaceType"`
	} `json:"RouterInterfaceSet" xml:"RouterInterfaceSet"`
}

func CreateDescribeRouterInterfacesForGlobalRequest() (request *DescribeRouterInterfacesForGlobalRequest) {
	request = &DescribeRouterInterfacesForGlobalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfacesForGlobal", "vpc", "openAPI")
	return
}

func CreateDescribeRouterInterfacesForGlobalResponse() (response *DescribeRouterInterfacesForGlobalResponse) {
	response = &DescribeRouterInterfacesForGlobalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
