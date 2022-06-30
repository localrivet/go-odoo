package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	AccountNo                  *String    `xmlrpc:"account_no,omptempty"`
	BankIds                    *Relation  `xmlrpc:"bank_ids,omptempty"`
	BaseOnboardingCompanyState *Selection `xmlrpc:"base_onboarding_company_state,omptempty"`
	ChildIds                   *Relation  `xmlrpc:"child_ids,omptempty"`
	City                       *String    `xmlrpc:"city,omptempty"`
	CompanyRegistry            *String    `xmlrpc:"company_registry,omptempty"`
	CountryId                  *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                 *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	Email                      *String    `xmlrpc:"email,omptempty"`
	ExternalReportLayoutId     *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	Favicon                    *String    `xmlrpc:"favicon,omptempty"`
	Font                       *Selection `xmlrpc:"font,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Logo                       *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                    *String    `xmlrpc:"logo_web,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	PaperformatId              *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                   *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerId                  *Many2One  `xmlrpc:"partner_id,omptempty"`
	Phone                      *String    `xmlrpc:"phone,omptempty"`
	PrimaryColor               *String    `xmlrpc:"primary_color,omptempty"`
	ReportFooter               *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader               *String    `xmlrpc:"report_header,omptempty"`
	SecondaryColor             *String    `xmlrpc:"secondary_color,omptempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omptempty"`
	StateId                    *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                     *String    `xmlrpc:"street,omptempty"`
	Street2                    *String    `xmlrpc:"street2,omptempty"`
	UserIds                    *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                        *String    `xmlrpc:"vat,omptempty"`
	Website                    *String    `xmlrpc:"website,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                        *String    `xmlrpc:"zip,omptempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	return c.Create(ResCompanyModel, rc)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.company was not found")
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.company was not found")
}
