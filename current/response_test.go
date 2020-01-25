package current

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResponseUnmarshalling(t *testing.T) {
	responseJson := `{
  "coord": {
    "lon": -122.08,
    "lat": 37.39
  },
  "weather": [
    {
      "id": 800,
      "main": "Clear",
      "description": "clear sky",
      "icon": "01d"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 282.55,
    "feels_like": 281.86,
    "temp_min": 280.37,
    "temp_max": 284.26,
    "pressure": 1023,
    "humidity": 100
  },
  "visibility": 16093,
  "wind": {
    "speed": 1.5,
    "deg": 350
  },
  "clouds": {
    "all": 1
  },
  "dt": 1560350645,
  "sys": {
    "type": 1,
    "id": 5122,
    "message": 0.0139,
    "country": "US",
    "sunrise": 1560343627,
    "sunset": 1560396563
  },
  "timezone": -25200,
  "id": 420006353,
  "name": "Mountain View",
  "cod": 200
}`

	expect := Response{
		Coord: Coord{
			Lon: -122.08,
			Lat: 37.39,
		},
		Weather: []Weather{
			{
				ID:          800,
				Main:        "Clear",
				Description: "clear sky",
				Icon:        "01d",
			},
		},
		Base: "stations",
		Main: Main{
			Temp:      282.55,
			FeelsLike: 281.86,
			TempMin:   280.37,
			TempMax:   284.26,
			Pressure:  1023,
			Humidity:  100,
		},
		Visibility: 16093,
		Wind: Wind{
			Speed: 1.5,
			Deg:   350,
		},
		Clouds: Clouds{
			All: 1,
		},
		Dt: 1560350645,
		Sys: Sys{
			Type:    1,
			ID:      5122,
			Message: 0.0139,
			Country: "US",
			Sunrise: 1560343627,
			Sunset:  1560396563,
		},
		Timezone: -25200,
		ID:       420006353,
		Name:     "Mountain View",
		Cod:      200,
	}

	var got Response
	if err := json.Unmarshal([]byte(responseJson), &got); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expect, got) {
		t.Error("un-marshalled Result does not match expected struct")
	}
}
