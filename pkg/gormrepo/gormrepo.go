package gormrepo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajch/contacts/pkg/contact"
)

type Gormrepo struct {
	filename string
	db       *gorm.DB
}

func NewGormrepo(fn string) (*Gormrepo, error) {
	db, err := gorm.Open("sqlite3", fn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&contact.Contact{})

	return &Gormrepo{
		filename: fn,
		db:       db,
	}, nil

}

func (g *Gormrepo) New(c *contact.Contact) (*contact.Contact, error) {
	err := g.db.Create(c).Error
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (g *Gormrepo) Update(c *contact.Contact) (*contact.Contact, error) {
	_, err := g.Get(c.Id)
	if err != nil {
		return nil, err
	}

	err = g.db.Save(c).Error
	if err != nil {
		return nil, err
	}

	return c, nil

}

func (g *Gormrepo) Delete(c *contact.Contact) error {
	_, err := g.Get(c.Id)
	if err != nil {
		return err
	}

	return g.db.Delete(c).Error
}

func (g *Gormrepo) Get(id uint) (*contact.Contact, error) {
	var c contact.Contact
	err := g.db.Where("id = ?", id).First(&c).Error
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (g *Gormrepo) List() ([]*contact.Contact, error) {
	var allcontacts []*contact.Contact
	g.db.Find(&allcontacts)
	return allcontacts, nil
}

func (g *Gormrepo) Close() {
	g.db.Close()
}
