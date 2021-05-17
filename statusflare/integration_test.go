package statusflare

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Integrations(t *testing.T) {
	client, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	var i *Integration

	// scenario: create new integration
	Convey("When we create a new integration", t, func() {

		i = &Integration{
			Name:   "Go test integration",
			Type:   "webhook",
			Secret: "some-secret-webhook",
		}

		err := client.CreateIntegration(i)
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we can get this integration by its ID", func() {
			_, err := client.GetIntegration(i.Id)
			if err != nil {
				t.Fatalf("%v", err)
			}
		})
	})

	Convey("When we get all integrations", t, func() {
		integrations, err := client.AllIntegrations()
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we have at least one integration", func() {
			if len(integrations) <= 0 {
				t.Fatal("No integration retrieved")
			}
		})
	})

	// scenario: change the integration
	Convey("When we change the integration name", t, func() {
		i.Name = "Go changed test integration"
		err := client.SaveIntegration(i)
		if err != nil {
			t.Fatalf("cannot update integration: %v", err)
		}

		Convey("Then the integration's name is changed", func() {
			changedm, _ := client.GetIntegration(i.Id)
			if changedm.Name != "Go changed test integration" {
				t.Fatalf("The name of the integration is unchanged")
			}
		})
	})

	// scenario: delete the integration
	Convey("When we delete the integration", t, func() {
		err = client.DeleteIntegration(i.Id)
		if err != nil {
			t.Fatalf("error in delete of integration: %v", err)
		}

		Convey("Then integration is no more available in Statusflare", func() {
			res, err := client.GetIntegration(i.Id)
			if err == nil && res.Id != "" {
				t.Fatalf("The integration still exist (%s)", i.Id)
			}
		})
	})
}
