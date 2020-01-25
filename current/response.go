package current

type Response struct {
	Base       string
	Coord      Coord
	Weather    []Weather
	Main       Main
	Visibility int64
	Wind       Wind
	Clouds     Clouds
	Dt         int64 // unix time
	Sys        Sys
	Timezone   int64
	ID         int64
	Name       string
	Cod        int64
}

type Coord struct {
	Lon float64
	Lat float64
}

type Weather struct {
	ID          int64
	Main        string
	Description string
	Icon        string // this is kind of an enum
}

type Main struct {
	Temp      float64
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64
	Humidity  int64
}

type Wind struct {
	Speed float64
	Deg   int64
}

type Clouds struct {
	All int64
}

type Sys struct {
	Type    int64
	ID      int64
	Message float64
	Country string
	Sunrise int64
	Sunset  int64
}
