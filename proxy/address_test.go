package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoService_AddressSearch(t *testing.T) {
	geoService := NewGeoService("c9d99e291a596d154d758473370cb61cd97b3c0f", "2e0cc6173ff587cb6753d5455691c6c53f16f53d")

	addresses, err := geoService.AddressSearch("г Москва, ул Сухонская, д 11")

	assert.NoError(t, err)
	assert.NotNil(t, addresses)
	assert.Equal(t, "Москва", addresses[0].City)
	assert.Equal(t, "Сухонская", addresses[0].Street)
	assert.Equal(t, "11", addresses[0].House)
	assert.Equal(t, "55.878315", addresses[0].Lat)
	assert.Equal(t, "37.65372", addresses[0].Lon)
}

func TestGeoService_GeoCode(t *testing.T) {
	geoService := NewGeoService("c9d99e291a596d154d758473370cb61cd97b3c0f", "2e0cc6173ff587cb6753d5455691c6c53f16f53d")

	addresses, err := geoService.GeoCode("56.3282034", "44.0032667")

	assert.NoError(t, err)
	assert.NotNil(t, addresses)
	assert.Equal(t, "Нижний Новгород", addresses[0].City)
	assert.Equal(t, "Кремль", addresses[0].Street)
	assert.Equal(t, "1", addresses[0].House)
	assert.Equal(t, "56.3282034", addresses[0].Lat)
	assert.Equal(t, "44.0032667", addresses[0].Lon)
}
