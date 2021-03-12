package gormrepo_test

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajch/contacts/pkg/contact"
	"github.com/rajch/contacts/pkg/gormrepo"
)

func TestGormrepo(t *testing.T) {
	g, err := gormrepo.NewGormrepo("testdb.db")
	if err != nil {
		t.Fatalf("Could not open database: %v.", err)
	}
	defer g.Close()

	c1 := contact.Contact{
		Name: "First Contact",
		Age:  50,
		City: "Bangalore",
	}

	c2 := contact.Contact{
		Name: "Second Contact",
		Age:  25,
		City: "Mumbai",
	}

	t.Log("Testing New...")

	t.Logf("Before insert: %v\n", c1)
	pc1, err := g.New(&c1)
	if err != nil {
		t.Fatalf("Could not write first record: %v", err)
	}

	t.Logf("After insert: %v\n", *pc1)

	pc2, err := g.New(&c2)
	if err != nil {
		t.Fatalf("Could not write second record: %v", err)
	}

	t.Log("Testing List...")

	allcontacts, err := g.List()
	if len(allcontacts) < 2 {
		t.Fatal("Number of inserted records do not match.")
	}

	for _, c := range allcontacts {
		t.Log(*c)
	}

	if allcontacts[0].Name != pc1.Name ||
		allcontacts[1].Name != pc2.Name {

		t.Fatal("Inserted records do not match.")
	}

	t.Log("Testing Get...")
	gpc1, err := g.Get(pc1.Id)
	if err != nil {
		t.Fatalf("Could not read written record: %v", err)
	}

	if gpc1.Name != c1.Name || gpc1.Age != c1.Age {
		t.Fatal("Returned record does not match original.")
	}

	t.Logf("Returned record: %v\n", *gpc1)

	t.Log("Testing Update...")
	gpc1.Name = "Changed name"
	gpc3, err := g.Update(gpc1)
	if err != nil {
		t.Fatalf("Could not update: %v", err)
	}

	t.Logf("Updated record: %v\n", gpc3)

	t.Log("Testing invalid update")
	var c3 contact.Contact
	c3 = *gpc1
	c3.Id = 5678
	gpc3, err = g.Update(&c3)
	if err == nil {
		t.Fatalf("Invalid update was allowed.")
	}
	t.Logf("Update was prevented with an error: %v", err)

	t.Log("Testing Delete...")
	err = g.Delete(&c1)
	if err != nil {
		t.Fatalf("Could not delete: %v", err)
	}

	t.Log("Testing non-existant get...")
	_, err = g.Get(c1.Id)
	if err == nil {
		t.Fatal("Get succeded for a non-existant id")
	}
	t.Logf("Error returned was: %v", err)

	err = g.Delete(&c2)
	if err != nil {
		t.Fatalf("Could not delete: %v", err)
	}
}
