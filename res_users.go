package odoo

import (
	"fmt"
)

// ResUsers represents res.users model.
type ResUsers struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	AccessesCount         *Int       `xmlrpc:"accesses_count,omptempty"`
	ActionId              *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount       *Int       `xmlrpc:"active_lang_count,omptempty"`
	ActivePartner         *Bool      `xmlrpc:"active_partner,omptempty"`
	BankIds               *Relation  `xmlrpc:"bank_ids,omptempty"`
	CategoryId            *Relation  `xmlrpc:"category_id,omptempty"`
	ChildIds              *Relation  `xmlrpc:"child_ids,omptempty"`
	City                  *String    `xmlrpc:"city,omptempty"`
	Color                 *Int       `xmlrpc:"color,omptempty"`
	Comment               *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId   *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompaniesCount        *Int       `xmlrpc:"companies_count,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyIds            *Relation  `xmlrpc:"company_ids,omptempty"`
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
	GroupsCount           *Int       `xmlrpc:"groups_count,omptempty"`
	GroupsId              *Relation  `xmlrpc:"groups_id,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Image1024             *String    `xmlrpc:"image_1024,omptempty"`
	Image128              *String    `xmlrpc:"image_128,omptempty"`
	Image1920             *String    `xmlrpc:"image_1920,omptempty"`
	Image256              *String    `xmlrpc:"image_256,omptempty"`
	Image512              *String    `xmlrpc:"image_512,omptempty"`
	IndustryId            *Many2One  `xmlrpc:"industry_id,omptempty"`
	IsCompany             *Bool      `xmlrpc:"is_company,omptempty"`
	Lang                  *Selection `xmlrpc:"lang,omptempty"`
	LogIds                *Relation  `xmlrpc:"log_ids,omptempty"`
	Login                 *String    `xmlrpc:"login,omptempty"`
	LoginDate             *Time      `xmlrpc:"login_date,omptempty"`
	Mobile                *String    `xmlrpc:"mobile,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	NewPassword           *String    `xmlrpc:"new_password,omptempty"`
	ParentId              *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName            *String    `xmlrpc:"parent_name,omptempty"`
	PartnerId             *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude       *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude      *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare          *Bool      `xmlrpc:"partner_share,omptempty"`
	Password              *String    `xmlrpc:"password,omptempty"`
	Phone                 *String    `xmlrpc:"phone,omptempty"`
	Ref                   *String    `xmlrpc:"ref,omptempty"`
	RulesCount            *Int       `xmlrpc:"rules_count,omptempty"`
	SameVatPartnerId      *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	Self                  *Many2One  `xmlrpc:"self,omptempty"`
	Share                 *Bool      `xmlrpc:"share,omptempty"`
	Signature             *String    `xmlrpc:"signature,omptempty"`
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

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	return c.Create(ResUsersModel, ru)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users not found", id)
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("res.users was not found")
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users was not found")
}
