package current

import (
	"net/url"
	"strconv"
)

type WeatherService struct {
	AppID    string
	Language string
	Units    string
	Mode     string
	URL      url.URL
	execute  Executor
	decode   Decoder
}

func NewWeatherService(appID string, configurations ...Configuration) *WeatherService {
	w := newWeatherService(appID)
	for _, configure := range configurations {
		configure(w)
	}
	return w
}

func newWeatherService(appID string) *WeatherService {
	return &WeatherService{
		AppID: appID,
		Mode:  "json",
		URL: url.URL{
			Scheme: "https",
			Host:   "api.openweathermap.org",
			Path:   "data/2.5/weather",
		},
		execute: DefaultExecutor,
		decode:  DefaultDecoder,
	}
}

func (w *WeatherService) getWeather(query url.URL) (*Response, error) {
	responseReader, err := w.execute(query)
	if err != nil {
		return nil, err
	}

	var response = new(Response)
	err = w.decode(responseReader, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (w *WeatherService) setParams(v *url.Values) {
	v.Add("appid", w.AppID)
	if w.Language != "" {
		v.Add("lang", w.Language)
	}
	if w.Units != "" {
		v.Add("units", w.Units)
	}
	if w.Mode != "" {
		v.Add("mode", w.Mode)
	}
}

func (w *WeatherService) GetByCityName(name string) (*Response, error) {
	query := url.Values{}
	query.Add("q", name)
	w.setParams(&query)

	requestUrl := w.URL
	requestUrl.RawQuery = query.Encode()

	return w.getWeather(requestUrl)
}

func (w *WeatherService) GetByZipCode(zip string) (*Response, error) {
	query := url.Values{}
	query.Add("zip", zip)
	w.setParams(&query)

	requestUrl := w.URL
	requestUrl.RawQuery = query.Encode()

	return w.getWeather(requestUrl)
}

func (w *WeatherService) GetByCityID(id int) (*Response, error) {
	query := url.Values{}
	query.Add("id", strconv.Itoa(id))
	w.setParams(&query)

	requestUrl := w.URL
	requestUrl.RawQuery = query.Encode()

	return w.getWeather(requestUrl)
}

func (w *WeatherService) GetByGeographicCoordinates(lat, lon float64) (*Response, error) {
	query := url.Values{}
	query.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	query.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))
	w.setParams(&query)

	requestUrl := w.URL
	requestUrl.RawQuery = query.Encode()

	return w.getWeather(requestUrl)
}
