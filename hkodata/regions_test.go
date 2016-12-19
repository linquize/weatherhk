package hkodata_test

import (
	"os"
	"testing"
	"time"

	"github.com/yookoala/weatherhk/hkodata"
)

func TestDecodeRegionJSON(t *testing.T) {
	file, err := os.Open("./test/region_json.201612191037.xml")
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}
	defer file.Close()

	regions, _ := hkodata.DecodeRegionJSON(file)
	if want, have := time.Date(2016, time.December, 19, 10, 20, 0, 0, hkodata.HKT), regions.PubDate; !want.Equal(have) {
		t.Errorf("expected %s, got %s", want, have)
	}

	expected := []hkodata.Region{
		hkodata.Region{
			Name:             hkodata.RegionName("hka"),
			ShortName:        "hka",
			CurrentTemp:      24.1,
			RelativeHumidity: .54,
			WindDirection:    "E",
			WindSpeed:        24,
			MaxTemp:          24.6,
			MinTemp:          19.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("cch"),
			ShortName:        "cch",
			CurrentTemp:      22.0,
			RelativeHumidity: .68,
			WindDirection:    "SE",
			WindSpeed:        30,
			MaxTemp:          22.0,
			MinTemp:          18.4,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("hpv"),
			ShortName:        "hpv",
			CurrentTemp:      23.1,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          23.1,
			MinTemp:          17.1,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("hko"),
			ShortName:        "hko",
			CurrentTemp:      20.8,
			RelativeHumidity: .74,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          20.8,
			MinTemp:          18.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("hkp"),
			ShortName:        "hkp",
			CurrentTemp:      21.6,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          21.6,
			MinTemp:          17.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("se1"),
			ShortName:        "se1",
			CurrentTemp:      21.8,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.5,
			MinTemp:          19.0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ksc"),
			ShortName:        "ksc",
			CurrentTemp:      22.2,
			RelativeHumidity: .66,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.2,
			MinTemp:          15.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("kp"),
			ShortName:        "kp",
			CurrentTemp:      21.7,
			RelativeHumidity: .66,
			WindDirection:    "SE",
			WindSpeed:        9,
			MaxTemp:          21.8,
			MinTemp:          17.7,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("klt"),
			ShortName:        "klt",
			CurrentTemp:      22.5,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          23.2,
			MinTemp:          17.8,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ktg"),
			ShortName:        "ktg",
			CurrentTemp:      21.9,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.0,
			MinTemp:          18.4,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("lfs"),
			ShortName:        "lfs",
			CurrentTemp:      23.8,
			RelativeHumidity: .62,
			WindDirection:    "SE",
			WindSpeed:        8,
			MaxTemp:          23.8,
			MinTemp:          15.9,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ngp"),
			ShortName:        "ngp",
			CurrentTemp:      19.3,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        39,
			MaxTemp:          20.1,
			MinTemp:          17.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tyw"),
			ShortName:        "tyw",
			CurrentTemp:      22.4,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.4,
			MinTemp:          12.4,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("pen"),
			ShortName:        "pen",
			CurrentTemp:      21.8,
			RelativeHumidity: .72,
			WindDirection:    "NE",
			WindSpeed:        18,
			MaxTemp:          21.8,
			MinTemp:          18.7,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("skg"),
			ShortName:        "skg",
			CurrentTemp:      20.6,
			RelativeHumidity: .72,
			WindDirection:    "SE",
			WindSpeed:        8,
			MaxTemp:          20.6,
			MinTemp:          16.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ssp"),
			ShortName:        "ssp",
			CurrentTemp:      24.1,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          24.3,
			MinTemp:          18.0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("sha"),
			ShortName:        "sha",
			CurrentTemp:      22.9,
			RelativeHumidity: .59,
			WindDirection:    "SE",
			WindSpeed:        5,
			MaxTemp:          23.0,
			MinTemp:          15.8,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("skw"),
			ShortName:        "skw",
			CurrentTemp:      21.2,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          21.7,
			MinTemp:          17.9,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("sek"),
			ShortName:        "sek",
			CurrentTemp:      22.9,
			RelativeHumidity: .63,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.9,
			MinTemp:          15.4,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ssh"),
			ShortName:        "ssh",
			CurrentTemp:      21.1,
			RelativeHumidity: .70,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          21.1,
			MinTemp:          16.1,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("sty"),
			ShortName:        "sty",
			CurrentTemp:      20.6,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        20,
			MaxTemp:          20.7,
			MinTemp:          18.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tkl"),
			ShortName:        "tkl",
			CurrentTemp:      21.9,
			RelativeHumidity: .65,
			WindDirection:    "E",
			WindSpeed:        10,
			MaxTemp:          21.9,
			MinTemp:          15.0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tms"),
			ShortName:        "tms",
			CurrentTemp:      16.4,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          17.0,
			MinTemp:          12.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tpo"),
			ShortName:        "tpo",
			CurrentTemp:      22.9,
			RelativeHumidity: .68,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.9,
			MinTemp:          16.7,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("vp1"),
			ShortName:        "vp1",
			CurrentTemp:      20.4,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          20.4,
			MinTemp:          15.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("jkb"),
			ShortName:        "jkb",
			CurrentTemp:      22.3,
			RelativeHumidity: .67,
			WindDirection:    "E",
			WindSpeed:        8,
			MaxTemp:          22.6,
			MinTemp:          17.0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ty1"),
			ShortName:        "ty1",
			CurrentTemp:      22.1,
			RelativeHumidity: .61,
			WindDirection:    "E",
			WindSpeed:        8,
			MaxTemp:          22.1,
			MinTemp:          16.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("twn"),
			ShortName:        "twn",
			CurrentTemp:      21.9,
			RelativeHumidity: .64,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          21.9,
			MinTemp:          16.1,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tw"),
			ShortName:        "tw",
			CurrentTemp:      22.8,
			RelativeHumidity: .63,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          22.8,
			MinTemp:          17.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tun"),
			ShortName:        "tun",
			CurrentTemp:      23.8,
			RelativeHumidity: .58,
			WindDirection:    "S",
			WindSpeed:        8,
			MaxTemp:          24.3,
			MinTemp:          17.9,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("wgl"),
			ShortName:        "wgl",
			CurrentTemp:      22.4,
			RelativeHumidity: .72,
			WindDirection:    "E",
			WindSpeed:        23,
			MaxTemp:          22.7,
			MinTemp:          18.6,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("wlp"),
			ShortName:        "wlp",
			CurrentTemp:      23.6,
			RelativeHumidity: .59,
			WindDirection:    "NE",
			WindSpeed:        6,
			MaxTemp:          24.0,
			MinTemp:          15.3,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("hks"),
			ShortName:        "hks",
			CurrentTemp:      22.3,
			RelativeHumidity: .62,
			WindDirection:    "NE",
			WindSpeed:        14,
			MaxTemp:          22.3,
			MinTemp:          17.2,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("wts"),
			ShortName:        "wts",
			CurrentTemp:      23.6,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          23.6,
			MinTemp:          17.5,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("ylp"),
			ShortName:        "ylp",
			CurrentTemp:      23.8,
			RelativeHumidity: 0,
			WindDirection:    "",
			WindSpeed:        0,
			MaxTemp:          23.8,
			MinTemp:          15.1,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tc"),
			ShortName:        "tc",
			CurrentTemp:      17.5,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        28,
			MaxTemp:          17.6,
			MinTemp:          14.0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("gi"),
			ShortName:        "gi",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "NE",
			WindSpeed:        29,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("se"),
			ShortName:        "se",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        17,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("sc"),
			ShortName:        "sc",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        18,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("sf"),
			ShortName:        "sf",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        16,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("plc"),
			ShortName:        "plc",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        11,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tpk"),
			ShortName:        "tpk",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "E",
			WindSpeed:        11,
			MaxTemp:          0,
			MinTemp:          0,
		},
		hkodata.Region{
			Name:             hkodata.RegionName("tap"),
			ShortName:        "tap",
			CurrentTemp:      0,
			RelativeHumidity: 0,
			WindDirection:    "SE",
			WindSpeed:        20,
			MaxTemp:          0,
			MinTemp:          0,
		},
	}

	for i := 0; i < len(expected); i++ {
		if want, have := expected[i], regions.Regions[i]; want != have {
			t.Errorf("\nexpected: %#v\n     got: %#v", want, have)
		}
		// special case: se1 has no name
		if regions.Regions[i].ShortName != "se1" {
			if regions.Regions[i].Name.En == "" {
				t.Errorf("The English name of %#v is empty string", regions.Regions[i].ShortName)
			}
			if regions.Regions[i].Name.Zh == "" {
				t.Errorf("The Chinese name of %#v is empty string", regions.Regions[i].ShortName)
			}
		}
	}

	/*
		// TODO: make sure empty windspeed / wind direction / relative humidity are
		// omitted in JSON
		data, _ := json.Marshal(regions.Regions)
		t.Logf("json: %s", data)
	*/
}