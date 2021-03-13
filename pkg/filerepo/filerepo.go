package filerepo

import (
	"os"

	"github.com/thulasi-devi-ravinanthanan/contacts/pkg/contact"
)

type Filerepo struct {
	filepath string
	file     *os.File
	LastId   uint
	Contacts []*contact.Contact
}

func (f *Filerepo) New(c *contact.Contact) (*contact.Contact, error) {
	f.LastId += 1
	c.Id = f.LastId
	f.Contacts = append(f.Contacts, c)
	return c, nil
}
