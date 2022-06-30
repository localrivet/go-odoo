package odoo

import (
	"fmt"
)

// ResPartner represents res.partner model.
type ResPartner struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount       *Int       `xmlrpc:"active_lang_count,omptempty"`
	BankIds               *Relation  `xmlrpc:"bank_ids,omptempty"`
	CategoryId            *Relation  `xmlrpc:"category_id,omptempty"`
	ChildIds              *Relation  `xmlrpc:"child_ids,omptempty"`
	City                  *String    `xmlrpc:"city,omptempty"`
	Color                 *Int       `xmlrpc:"color,omptempty"`
	Comment               *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId   *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyName           *String    `xmlrpc:"company_name,omptempty"`
	CompanyType           *Selection `xmlrpc:"company_type,omptempty"`
	ContactAddress        *String    `xmlrpc:"contact_address,omptempty"`
	CountryId             *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CreditLimit           *Float     `xmlrpc:"credit_limit,omptempty"`
	Date                  *Time      `xmlrpc:"date,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Email                 *String    `xmlrpc:"email,omptempty"`
	EmailFormatted        *String    `xmlrpc:"email_formatted,omptempty"`
	Employee              *Bool      `xmlrpc:"employee,omptempty"`
	Function              *String    `xmlrpc:"function,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Image1024             *String    `xmlrpc:"image_1024,omptempty"`
	Image128              *String    `xmlrpc:"image_128,omptempty"`
	Image1920             *String    `xmlrpc:"image_1920,omptempty"`
	Image256              *String    `xmlrpc:"image_256,omptempty"`
	Image512              *String    `xmlrpc:"image_512,omptempty"`
	IndustryId            *Many2One  `xmlrpc:"industry_id,omptempty"`
	IsCompany             *Bool      `xmlrpc:"is_company,omptempty"`
	Lang                  *Selection `xmlrpc:"lang,omptempty"`
	Mobile                *String    `xmlrpc:"mobile,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	ParentId              *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName            *String    `xmlrpc:"parent_name,omptempty"`
	PartnerLatitude       *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude      *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare          *Bool      `xmlrpc:"partner_share,omptempty"`
	Phone                 *String    `xmlrpc:"phone,omptempty"`
	Ref                   *String    `xmlrpc:"ref,omptempty"`
	SameVatPartnerId      *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	Self                  *Many2One  `xmlrpc:"self,omptempty"`
	StateId               *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                *String    `xmlrpc:"street,omptempty"`
	Street2               *String    `xmlrpc:"street2,omptempty"`
	Title                 *Many2One  `xmlrpc:"title,omptempty"`
	Type                  *Selection `xmlrpc:"type,omptempty"`
	Tz                    *Selection `xmlrpc:"tz,omptempty"`
	TzOffset              *String    `xmlrpc:"tz_offset,omptempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds               *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                   *String    `xmlrpc:"vat,omptempty"`
	Website               *String    `xmlrpc:"website,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                   *String    `xmlrpc:"zip,omptempty"`
}

// ResPartners represents array of res.partner model.
type ResPartners []ResPartner

// ResPartnerModel is the odoo model name.
const ResPartnerModel = "res.partner"

// Many2One convert ResPartner to *Many2One.
func (rp *ResPartner) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartner(rp *ResPartner) (int64, error) {
	return c.Create(ResPartnerModel, rp)
}

// UpdateResPartner updates an existing res.partner record.
func (c *Client) UpdateResPartner(rp *ResPartner) error {
	return c.UpdateResPartners([]int64{rp.Id.Get()}, rp)
}

// UpdateResPartners updates existing res.partner records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateResPartners(ids []int64, rp *ResPartner) error {
	return c.Update(ResPartnerModel, ids, rp)
}

// DeleteResPartner deletes an existing res.partner record.
func (c *Client) DeleteResPartner(id int64) error {
	return c.DeleteResPartners([]int64{id})
}

// DeleteResPartners deletes existing res.partner records.
func (c *Client) DeleteResPartners(ids []int64) error {
	return c.Delete(ResPartnerModel, ids)
}

// GetResPartner gets res.partner existing record.
func (c *Client) GetResPartner(id int64) (*ResPartner, error) {
	rps, err := c.GetResPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner not found", id)
}

// GetResPartners gets res.partner existing records.
func (c *Client) GetResPartners(ids []int64) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.Read(ResPartnerModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartner finds res.partner record by querying it with criteria.
func (c *Client) FindResPartner(criteria *Criteria) (*ResPartner, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("res.partner was not found")
}

// FindResPartners finds res.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartners(criteria *Criteria, options *Options) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerId finds record id by querying it with criteria.
func (c *Client) FindResPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner was not found")
}
