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
}
