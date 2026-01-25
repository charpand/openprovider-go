// Package provider implements the Terraform provider for OpenProvider.
package provider

import (
	"context"
	"fmt"

	"github.com/charpand/terraform-provider-openprovider/internal/client"
	"github.com/charpand/terraform-provider-openprovider/internal/client/customers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CustomerDataSource{}
	_ datasource.DataSourceWithConfigure = &CustomerDataSource{}
)

// CustomerDataSource is the data source implementation.
type CustomerDataSource struct {
	client *client.Client
}

// NewCustomerDataSource returns a new instance of the customer data source.
func NewCustomerDataSource() datasource.DataSource {
	return &CustomerDataSource{}
}

// Metadata returns the data source type name.
func (d *CustomerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer"
}

// Schema defines the schema for the data source.
func (d *CustomerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about an OpenProvider customer (contact handle).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The customer identifier (same as handle).",
				Computed:            true,
			},
			"handle": schema.StringAttribute{
				MarkdownDescription: "The customer handle to look up (e.g., XX123456-XX).",
				Required:            true,
			},
			"company_name": schema.StringAttribute{
				MarkdownDescription: "The company name.",
				Computed:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "The customer's email address.",
				Computed:            true,
			},
			"locale": schema.StringAttribute{
				MarkdownDescription: "The customer's language/locale.",
				Computed:            true,
			},
			"comments": schema.StringAttribute{
				MarkdownDescription: "Custom notes about this customer.",
				Computed:            true,
			},
		},
		Blocks: map[string]schema.Block{
			"phone": schema.SingleNestedBlock{
				MarkdownDescription: "The customer's phone number.",
				Attributes: map[string]schema.Attribute{
					"country_code": schema.StringAttribute{
						MarkdownDescription: "Country code.",
						Computed:            true,
					},
					"area_code": schema.StringAttribute{
						MarkdownDescription: "Area code.",
						Computed:            true,
					},
					"number": schema.StringAttribute{
						MarkdownDescription: "Phone number.",
						Computed:            true,
					},
				},
			},
			"address": schema.SingleNestedBlock{
				MarkdownDescription: "The customer's address.",
				Attributes: map[string]schema.Attribute{
					"street": schema.StringAttribute{
						MarkdownDescription: "Street name.",
						Computed:            true,
					},
					"number": schema.StringAttribute{
						MarkdownDescription: "Street number.",
						Computed:            true,
					},
					"suffix": schema.StringAttribute{
						MarkdownDescription: "Address suffix.",
						Computed:            true,
					},
					"city": schema.StringAttribute{
						MarkdownDescription: "City name.",
						Computed:            true,
					},
					"state": schema.StringAttribute{
						MarkdownDescription: "State or province.",
						Computed:            true,
					},
					"zipcode": schema.StringAttribute{
						MarkdownDescription: "Postal/ZIP code.",
						Computed:            true,
					},
					"country": schema.StringAttribute{
						MarkdownDescription: "Country code.",
						Computed:            true,
					},
				},
			},
			"name": schema.SingleNestedBlock{
				MarkdownDescription: "The customer's name.",
				Attributes: map[string]schema.Attribute{
					"first_name": schema.StringAttribute{
						MarkdownDescription: "First name.",
						Computed:            true,
					},
					"last_name": schema.StringAttribute{
						MarkdownDescription: "Last name.",
						Computed:            true,
					},
					"initials": schema.StringAttribute{
						MarkdownDescription: "Initials.",
						Computed:            true,
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: "Name prefix.",
						Computed:            true,
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *CustomerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

// Read retrieves the customer information.
func (d *CustomerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CustomerModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	handle := config.Handle.ValueString()

	// Get customer
	customer, err := customers.Get(d.client, handle)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Customer",
			fmt.Sprintf("Could not read customer %s: %s", handle, err.Error()),
		)
		return
	}

	if customer == nil {
		resp.Diagnostics.AddError(
			"Customer Not Found",
			fmt.Sprintf("Customer %s not found", handle),
		)
		return
	}

	// Map to state
	var state CustomerModel
	state.Handle = types.StringValue(customer.Handle)
	state.ID = types.StringValue(customer.Handle)
	state.Email = types.StringValue(customer.Email)
	
	if customer.CompanyName != "" {
		state.CompanyName = types.StringValue(customer.CompanyName)
	} else {
		state.CompanyName = types.StringNull()
	}
	
	if customer.Locale != "" {
		state.Locale = types.StringValue(customer.Locale)
	} else {
		state.Locale = types.StringNull()
	}
	
	if customer.Comments != "" {
		state.Comments = types.StringValue(customer.Comments)
	} else {
		state.Comments = types.StringNull()
	}

	// Map phone
	state.Phone = &PhoneModel{
		CountryCode: types.StringValue(customer.Phone.CountryCode),
		AreaCode:    types.StringValue(customer.Phone.AreaCode),
		Number:      types.StringValue(customer.Phone.Number),
	}

	// Map address
	state.Address = &AddressModel{
		Street:  types.StringValue(customer.Address.Street),
		City:    types.StringValue(customer.Address.City),
		Country: types.StringValue(customer.Address.Country),
	}
	if customer.Address.Number != "" {
		state.Address.Number = types.StringValue(customer.Address.Number)
	} else {
		state.Address.Number = types.StringNull()
	}
	if customer.Address.Suffix != "" {
		state.Address.Suffix = types.StringValue(customer.Address.Suffix)
	} else {
		state.Address.Suffix = types.StringNull()
	}
	if customer.Address.State != "" {
		state.Address.State = types.StringValue(customer.Address.State)
	} else {
		state.Address.State = types.StringNull()
	}
	if customer.Address.Zipcode != "" {
		state.Address.Zipcode = types.StringValue(customer.Address.Zipcode)
	} else {
		state.Address.Zipcode = types.StringNull()
	}

	// Map name
	state.Name = &NameModel{
		FirstName: types.StringValue(customer.Name.FirstName),
		LastName:  types.StringValue(customer.Name.LastName),
	}
	if customer.Name.Initials != "" {
		state.Name.Initials = types.StringValue(customer.Name.Initials)
	} else {
		state.Name.Initials = types.StringNull()
	}
	if customer.Name.Prefix != "" {
		state.Name.Prefix = types.StringValue(customer.Name.Prefix)
	} else {
		state.Name.Prefix = types.StringNull()
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
