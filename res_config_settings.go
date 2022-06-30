package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	return c.Create(ResConfigSettingsModel, rcs)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config.settings not found", id)
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.config.settings was not found")
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config.settings was not found")
}
