/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package controllers

import (
	"net/http"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsVersionController struct {
	client *nitro.Client
}

func NewNsVersionController(client *nitro.Client) *NsVersionController {
	c := NsVersionController{
		client: client,
	}
	return &c
}

func (c *NsVersionController) Get(attributes []string) (*nitro.Response[config.NsVersion], error) {
	r := nitro.Request[config.NsVersion]{
		Method:     http.MethodGet,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *NsVersionController) GetVersion() (string, error) {
	r := nitro.Request[config.NsVersion]{
		Method:     http.MethodGet,
		Attributes: []string{"version"},
	}
	res, err := nitro.ExecuteNitroRequest(c.client, &r)
	if err != nil {
		return "", err
	}

	return res.Data[0].Version, nil
}
