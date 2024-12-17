package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalGeoCode(t *testing.T) {
	validJSON := `{
		"suggestions": [
			{
				"value": "Some Value",
				"unrestricted_value": "Some Unrestricted Value",
				"data": {
					"postal_code": "190000",
					"country": "Россия",
					"country_iso_code": "RU",
					"federal_district": "Северо-Западный",
					"region_fias_id": "1",
					"region_kladr_id": "78",
					"region_iso_code": "RU-SPE",
					"region_with_type": "г Санкт-Петербург",
					"region_type": "г",
					"region_type_full": "город",
					"region": "Санкт-Петербург",
					"city_fias_id": "2",
					"city_kladr_id": "78",
					"city_with_type": "г Санкт-Петербург",
					"city_type": "г",
					"city_type_full": "город",
					"city": "Санкт-Петербург",
					"street_fias_id": "3",
					"street_kladr_id": "4",
					"street_with_type": "ул Казанская",
					"street_type": "ул",
					"street_type_full": "улица",
					"street": "Казанская",
					"house": "10"
				}
			}
		]
	}`

	var expectedGeoCode GeoCode
	_ = json.Unmarshal([]byte(validJSON), &expectedGeoCode)

	geoCode, err := UnmarshalGeoCode([]byte(validJSON))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(geoCode, expectedGeoCode) {
		t.Errorf("expected %+v, got %+v", expectedGeoCode, geoCode)
	}

	invalidJSON := `{ invalid json }`
	_, err = UnmarshalGeoCode([]byte(invalidJSON))
	if err == nil {
		t.Error("expected error for invalid JSON, got nil")
	}
}

func TestMarshalGeoCode(t *testing.T) {
	geoCode := GeoCode{
		Suggestions: []Suggestion{
			{
				Value:             "Some Value",
				UnrestrictedValue: "Some Unrestricted Value",
				Data: Data{
					PostalCode:     "190000",
					Country:        "Россия",
					CountryISOCode: "RU",
					FederalDistrict: "Северо-Западный",
					RegionFiasID:    "1",
					RegionKladrID:   "78",
					RegionISOCode:   "RU-SPE",
					RegionWithType:  "г Санкт-Петербург",
					RegionType:      "г",
					RegionTypeFull:  "город",
					Region:          "Санкт-Петербург",
					CityFiasID:      "2",
					CityKladrID:     "78",
					CityWithType:    "г Санкт-Петербург",
					CityType:        "г",
					CityTypeFull:    "город",
					City:            "Санкт-Петербург",
					StreetFiasID:    "3",
					StreetKladrID:   "4",
					StreetWithType:  "ул Казанская",
					StreetType:      "ул",
					StreetTypeFull:  "улица",
					Street:          "Казанская",
					House:           "10",
				},
			},
		},
	}

	data, err := geoCode.Marshal()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var unmarshalledGeoCode GeoCode
	if err := json.Unmarshal(data, &unmarshalledGeoCode); err != nil {
		t.Fatalf("failed to unmarshal data: %v", err)
	}

	if !reflect.DeepEqual(geoCode, unmarshalledGeoCode) {
		t.Errorf("expected %+v, got %+v", geoCode, unmarshalledGeoCode)
	}
}
