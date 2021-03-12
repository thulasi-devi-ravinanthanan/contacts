package contact

type Contact struct {
	Id    uint `json:"id" gorm:"primary_key"`
	Name  string
	Phone string
	Email string
	City  string
	Age   int
}

type ContactRepository interface {
	New(*Contact) (*Contact, error)
	Update(*Contact) (*Contact, error)
	Delete(*Contact) error

	Get(uint) (*Contact, error)
	List() ([]*Contact, error)
}
