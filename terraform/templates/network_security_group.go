package templates

import artifacts "github.com/kun-lun/artifacts/pkg/apis"

var networkSecurityGroupTF = []byte(`
  resource "azurerm_network_security_group" "{{.nsgName}}" {
	name                = "${var.env_name}-{{.nsgName}}"
	location            = "${var.location}"
	resource_group_name = "${var.resource_group_name}"
  }

`)

var networkSecurityRuleTF = []byte(`
  resource "azurerm_network_security_rule" "{{.nsrName}}" {
	name                        = "{{.nsrName}}"
	priority                    = "${var.{{.nsrName}}_ansr_priority}"
	direction                   = "${var.{{.nsrName}}_ansr_direction}"
	access                      = "${var.{{.nsrName}}_ansr_access}"
	protocol                    = "${var.{{.nsrName}}_ansr_protocol}"
	source_port_range           = "${var.{{.nsrName}}_ansr_source_port_range}"
	destination_port_range      = "${var.{{.nsrName}}_ansr_destination_port_range}"
	source_address_prefix       = "${var.{{.nsrName}}_ansr_source_address_prefix}"
	destination_address_prefix  = "${var.{{.nsrName}}_ansr_destination_address_prefix}"
	resource_group_name         = "${var.resource_group_name}"
	network_security_group_name = "${azurerm_network_security_group.{{.nsgName}}.name}"
  }

  variable "{{.nsrName}}_ansr_priority" {}

  variable "{{.nsrName}}_ansr_direction" {}

  variable "{{.nsrName}}_ansr_access" {}

  variable "{{.nsrName}}_ansr_protocol" {}

  variable "{{.nsrName}}_ansr_source_port_range" {}

  variable "{{.nsrName}}_ansr_destination_port_range" {}

  variable "{{.nsrName}}_ansr_source_address_prefix" {}

  variable "{{.nsrName}}_ansr_destination_address_prefix" {}

`)

var networkSecurityRuleTFVars = []byte(`
{{.nsrName}}_ansr_priority = "{{.ansr_priority}}"
{{.nsrName}}_ansr_direction = "{{.ansr_direction}}"
{{.nsrName}}_ansr_access = "{{.ansr_access}}"
{{.nsrName}}_ansr_protocol = "{{.ansr_protocol}}"
{{.nsrName}}_ansr_source_port_range = "{{.ansr_source_port_range}}"
{{.nsrName}}_ansr_destination_port_range = "{{.ansr_destination_port_range}}"
{{.nsrName}}_ansr_source_address_prefix = "{{.ansr_source_address_prefix}}"
{{.nsrName}}_ansr_destination_address_prefix = "{{.ansr_destination_address_prefix}}"
`)

func NewNSGTemplate(nsg artifacts.NetworkSecurityGroup) (string, error) {
	tf := ""
	nsgTF, err := render(networkSecurityGroupTF, getNSGTFParams(nsg))
	if err != nil {
		return "", err
	}
	tf += nsgTF

	for _, nsr := range nsg.NetworkSecurityRules {
		nsrTF, err := render(networkSecurityRuleTF, getNSRTFParams(nsr, nsg.Name))
		if err != nil {
			return "", err
		}
		tf += nsrTF
	}
	return tf, nil
}

func NewNSGInput(nsg artifacts.NetworkSecurityGroup) (string, error) {
	tfvars := ""
	for _, nsr := range nsg.NetworkSecurityRules {
		nsrVars, err := render(networkSecurityRuleTFVars, getNSRTFVarsParams(nsr))
		if err != nil {
			return "", err
		}
		tfvars += nsrVars
	}
	return tfvars, nil
}

func getNSGTFParams(nsg artifacts.NetworkSecurityGroup) map[string]interface{} {
	return map[string]interface{}{
		"nsgName": nsg.Name,
	}
}

func getNSRTFParams(nsr artifacts.NetworkSecurityRule, nsgName string) map[string]interface{} {
	return map[string]interface{}{
		"nsgName": nsgName,
		"nsrName": nsr.Name,
	}
}

func getNSRTFVarsParams(nsr artifacts.NetworkSecurityRule) map[string]interface{} {
	return map[string]interface{}{
		"nsrName":                         nsr.Name,
		"ansr_priority":                   nsr.Priority,
		"ansr_direction":                  nsr.Direction,
		"ansr_access":                     nsr.Access,
		"ansr_protocol":                   nsr.Protocol,
		"ansr_source_port_range":          nsr.SourcePortRange,
		"ansr_destination_port_range":     nsr.DestinationPortRange,
		"ansr_source_address_prefix":      nsr.SourceAddressPrefix,
		"ansr_destination_address_prefix": nsr.DestinationAddressPrefix,
	}
}
