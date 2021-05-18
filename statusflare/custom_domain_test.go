package statusflare

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_CustomDomains(t *testing.T) {
	client, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	var d *CustomDomain

	// scenario: create new custom domain
	Convey("When we create a new custom domain", t, func() {

		d = &CustomDomain{
			Domain: "acc-test.statusflare.app",
			Type:   "full_domain",
		}

		err := client.CreateCustomDomain(d)
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we can get this custom domain by its ID", func() {
			_, err := client.GetCustomDomain(d.Id)
			if err != nil {
				t.Fatalf("%v", err)
			}
		})
	})

	Convey("When we get all custom domains", t, func() {
		customDomains, err := client.AllCustomDomains()
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we have at least one custom domain", func() {
			if len(customDomains) <= 0 {
				t.Fatal("No custom domains retrieved")
			}
		})
	})

	// scenario: delete the custom domain
	Convey("When we delete the custom domain", t, func() {
		err = client.DeleteCustomDomain(d.Id)
		if err != nil {
			t.Fatalf("error in delete of custom domain: %v", err)
		}

		Convey("Then custom domain is no more available in Statusflare", func() {
			res, err := client.GetCustomDomain(d.Id)
			if err == nil && res.Id != "" {
				t.Fatalf("The custom domain still exist (%s)", d.Id)
			}
		})
	})
}
