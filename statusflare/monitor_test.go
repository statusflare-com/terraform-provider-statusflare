package statusflare

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_MonitorIntegration(t *testing.T) {

	var m *Monitor

	client, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	// scenario: create new monitor
	Convey("When we create a new monitor", t, func() {

		m = &Monitor{
			Name:         "test-1",
			URL:          "sme.sk",
			Scheme:       "https",
			Method:       "GET",
			ExpectStatus: 200,
			Worker:       "managed",
		}

		err := client.CreateMonitor(m)
		if err != nil {
			t.Fatalf("%v", err)
		}

		Convey("Then we cat get this monitor by his ID", func() {
			_, err := client.GetMonitor(m.ID)
			if err != nil {
				t.Fatalf("%v", err)
			}
		})
	})

	// scenario: change the monitor
	Convey("When we save the changed monitor", t, func() {
		m.Name = "test-2"
		m.URL = "hnonline.sk"
		err := client.SaveMonitor(m)
		if err != nil {
			t.Fatalf("cannot update monitor: %v", err)
		}

		Convey("Then the monitor's name is changed", func() {
			changedm, _ := client.GetMonitor(m.ID)
			if changedm.Name != "test-zvr-6" {
				t.Fatalf("The name of the monitor is unchanged")
			}
		})
	})

	// scenario: delete the monitor
	Convey("When we delete the monitor", t, func() {
		err = client.DeleteMonitor(m.ID)
		if err != nil {
			t.Fatalf("error in delete of monitor: %v", err)
		}

		Convey("Then monitor is no more available in Statusflare", func() {
			res, err := client.GetMonitor(m.ID)
			if err == nil && res.ID != "" {
				t.Fatalf("The monitor still exist, even we delete it (%s)", m.ID)
			}
		})
	})
}
