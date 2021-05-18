package statusflare

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_StatusPage(t *testing.T) {

	var s *StatusPage

	client, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	// scenario: create new status page
	Convey("When we create a new status page", t, func() {

		s = &StatusPage{
			Name:     "Go test status page",
			Monitors: []string{},
		}

		err := client.CreateStatusPage(s)
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we cat get this status page by its ID", func() {
			_, err := client.GetStatusPage(s.Id)
			if err != nil {
				t.Fatalf("%v", err)
			}
		})
	})

	// scenario: change the status page
	Convey("When we save the changed status page", t, func() {
		s.Name = "Go changed test status page"
		err := client.SaveStatusPage(s)
		if err != nil {
			t.Fatalf("cannot update status page: %v", err)
		}

		Convey("Then the status pages's name is changed", func() {
			changedm, _ := client.GetStatusPage(s.Id)
			if changedm.Name != "Go changed test status page" {
				t.Fatalf("The name of the status page is unchanged")
			}
		})
	})

	// scenario: delete the status page
	Convey("When we delete the status page", t, func() {
		err = client.DeleteStatusPage(s.Id)
		if err != nil {
			t.Fatalf("error in delete of status page: %v", err)
		}

		Convey("Then status page is no more available in Statusflare", func() {
			res, err := client.GetMonitor(s.Id)
			if err == nil && res.Id != "" {
				t.Fatalf("The status page still exist, even we delete it (%s)", s.Id)
			}
		})
	})
}
