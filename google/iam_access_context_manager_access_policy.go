// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

var AccessContextManagerAccessPolicyIamSchema = map[string]*schema.Schema{
	"name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type AccessContextManagerAccessPolicyIamUpdater struct {
	name   string
	d      TerraformResourceData
	Config *transport_tpg.Config
}

func AccessContextManagerAccessPolicyIamUpdaterProducer(d TerraformResourceData, config *transport_tpg.Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("name"); ok {
		values["name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"accessPolicies/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config, d.Get("name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &AccessContextManagerAccessPolicyIamUpdater{
		name:   values["name"],
		d:      d,
		Config: config,
	}

	if err := d.Set("name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting name: %s", err)
	}

	return u, nil
}

func AccessContextManagerAccessPolicyIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	m, err := getImportIdQualifiers([]string{"accessPolicies/(?P<name>[^/]+)", "(?P<name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &AccessContextManagerAccessPolicyIamUpdater{
		name:   values["name"],
		d:      d,
		Config: config,
	}
	if err := d.Set("name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *AccessContextManagerAccessPolicyIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAccessPolicyUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "POST", "", url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *AccessContextManagerAccessPolicyIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAccessPolicyUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", "", url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *AccessContextManagerAccessPolicyIamUpdater) qualifyAccessPolicyUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{AccessContextManagerBasePath}}%s:%s", fmt.Sprintf("accessPolicies/%s", u.name), methodIdentifier)
	url, err := ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *AccessContextManagerAccessPolicyIamUpdater) GetResourceId() string {
	return fmt.Sprintf("accessPolicies/%s", u.name)
}

func (u *AccessContextManagerAccessPolicyIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-accesscontextmanager-accesspolicy-%s", u.GetResourceId())
}

func (u *AccessContextManagerAccessPolicyIamUpdater) DescribeResource() string {
	return fmt.Sprintf("accesscontextmanager accesspolicy %q", u.GetResourceId())
}
