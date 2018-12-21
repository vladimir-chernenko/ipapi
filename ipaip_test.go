package ipapi

import "testing"

type IpDetailsTestCase struct {
	Ip          string
	City        string
	PhonePrefix string
	MustFail    bool
}

func TestGetIpDetails(t *testing.T) {
	testCases := []IpDetailsTestCase{
		IpDetailsTestCase{
			Ip:          "95.82.135.131",
			City:        "Brno (Brno-Bystrc)",
			PhonePrefix: "420",
			MustFail:    false,
		},
		IpDetailsTestCase{
			Ip:          "0.0.0.0",
			City:        "",
			PhonePrefix: "",
			MustFail:    true,
		},
	}

	for _, iptc := range testCases {
		ipd, err := GetIpDetails(iptc.Ip)

		if err != nil && !iptc.MustFail {
			t.Error(err)
		}

		if ipd.City != iptc.City {
			t.Errorf("%s: expected city is '%s', but got '%s'", iptc.Ip, iptc.City, ipd.City)
		}

		if ipd.PhonePrefix != iptc.PhonePrefix {
			t.Errorf("%s: expected phone prefix is '%s', but got '%s'", iptc.Ip, iptc.PhonePrefix, ipd.PhonePrefix)
		}
	}
}
