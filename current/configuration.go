package current

import "net/url"

type Configuration func(*WeatherService)

func SetLanguage(code string) Configuration {
	return func(w *WeatherService) {
		w.Language = code
	}
}

func SetUnits(units string) Configuration {
	return func(w *WeatherService) {
		w.Units = units
	}
}

var SetUnitsImperial Configuration = func(service *WeatherService) {
	service.Units = "imperial"
}

var SetUnitsKelvin Configuration = func(service *WeatherService) {
	service.Units = "kelvin"
}

var SetUnitsMetric Configuration = func(service *WeatherService) {
	service.Units = "metric"
}

func SetUrl(url url.URL) Configuration {
	return func(service *WeatherService) {
		service.URL = url
	}
}

func SetExecutor(executor Executor) Configuration {
	return func(service *WeatherService) {
		service.execute = executor
	}
}
