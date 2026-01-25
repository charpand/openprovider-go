// Package provider implements the Terraform provider for OpenProvider.
package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AddressModel represents a customer address in Terraform state.
type AddressModel struct {
	City    types.String `tfsdk:"city"`
	Country types.String `tfsdk:"country"`
	Number  types.String `tfsdk:"number"`
	State   types.String `tfsdk:"state"`
	Street  types.String `tfsdk:"street"`
	Suffix  types.String `tfsdk:"suffix"`
	Zipcode types.String `tfsdk:"zipcode"`
}

// NameModel represents a customer name in Terraform state.
type NameModel struct {
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
	Initials  types.String `tfsdk:"initials"`
	Prefix    types.String `tfsdk:"prefix"`
}

// PhoneModel represents a customer phone in Terraform state.
type PhoneModel struct {
	AreaCode    types.String `tfsdk:"area_code"`
	CountryCode types.String `tfsdk:"country_code"`
	Number      types.String `tfsdk:"number"`
}

// CustomerModel represents the Terraform state model for a customer.
type CustomerModel struct {
	ID          types.String `tfsdk:"id"`
	Handle      types.String `tfsdk:"handle"`
	CompanyName types.String `tfsdk:"company_name"`
	Email       types.String `tfsdk:"email"`
	Locale      types.String `tfsdk:"locale"`
	Comments    types.String `tfsdk:"comments"`
	Phone       *PhoneModel  `tfsdk:"phone"`
	Address     *AddressModel `tfsdk:"address"`
	Name        *NameModel    `tfsdk:"name"`
}
