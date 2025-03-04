package iam

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/response"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"reflect"
)

const (
	DefaultIamEndpoint         = "https://iam.myhuaweicloud.com"
	KeystoneListProjectsUri    = "/v3/projects"
	KeystoneListAuthDomainsUri = "/v3/auth/domains"
)

type KeystoneListProjectsResponse struct {
	Projects *[]ProjectResult `json:"projects,omitempty"`
}

type ProjectResult struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetKeystoneListProjectsRequest(iamEndpoint string, regionId string) *request.DefaultHttpRequest {
	return request.NewHttpRequestBuilder().
		WithEndpoint(iamEndpoint).
		WithPath(KeystoneListProjectsUri).
		WithMethod("GET").
		AddQueryParam("name", reflect.ValueOf(regionId)).
		Build()
}

func KeystoneListProjects(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (string, error) {
	resp, err := client.SyncInvokeHttp(req)
	if err != nil {
		return "", err
	}

	data, err := GetResponseBody(resp)
	if err != nil {
		return "", err
	}

	keystoneListProjectResponse := new(KeystoneListProjectsResponse)
	err = jsoniter.Unmarshal(data, keystoneListProjectResponse)
	if err != nil {
		return "", err
	}

	if len(*keystoneListProjectResponse.Projects) > 0 {

		return (*keystoneListProjectResponse.Projects)[0].Id, nil
	}

	return "", errors.New("no project found")
}

type KeystoneListAuthDomainsResponse struct {
	Domains *[]Domains `json:"domains,omitempty"`
}

type Domains struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetKeystoneListAuthDomainsRequest(iamEndpoint string) *request.DefaultHttpRequest {
	return request.NewHttpRequestBuilder().
		WithEndpoint(fmt.Sprintf("%s%s", iamEndpoint, KeystoneListAuthDomainsUri)).
		WithMethod("GET").
		Build()
}

func KeystoneListAuthDomains(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (string, error) {
	resp, err := client.SyncInvokeHttp(req)
	if err != nil {
		return "", err
	}

	data, err := GetResponseBody(resp)
	if err != nil {
		return "", err
	}

	keystoneListAuthDomainsResponse := new(KeystoneListAuthDomainsResponse)
	err = jsoniter.Unmarshal(data, keystoneListAuthDomainsResponse)
	if err != nil {
		return "", err
	}

	if len(*keystoneListAuthDomainsResponse.Domains) > 0 {

		return (*keystoneListAuthDomainsResponse.Domains)[0].Id, nil
	}

	return "", errors.New("no domain found")
}

func GetResponseBody(resp *response.DefaultHttpResponse) ([]byte, error) {
	if resp.GetStatusCode() >= 400 {
		return nil, sdkerr.NewServiceResponseError(resp.Response)
	}

	data, err := ioutil.ReadAll(resp.Response.Body)

	if err != nil {
		if closeErr := resp.Response.Body.Close(); closeErr != nil {
			return nil, err
		}
		return nil, err
	}

	if err := resp.Response.Body.Close(); err != nil {
		return nil, err
	} else {
		resp.Response.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	return data, nil
}
