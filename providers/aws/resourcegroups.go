// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
)

var resourcegroupsAllowEmptyValues = []string{"tags."}

type ResourceGroupsGenerator struct {
	AWSService
}

func (g *ResourceGroupsGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := resourcegroups.New(config)
	p := resourcegroups.NewListGroupsPaginator(svc.ListGroupsRequest(&resourcegroups.ListGroupsInput{}))
	var resources []terraformutils.Resource
	for p.Next(context.Background()) {
		for _, group := range p.CurrentPage().Groups {
			groupName := aws.StringValue(group.Name)
			resources = append(resources, terraformutils.NewSimpleResource(
				groupName,
				groupName,
				"aws_resourcegroups_group",
				"aws",
				resourcegroupsAllowEmptyValues))
		}
	}
	g.Resources = resources
	return p.Err()
}
