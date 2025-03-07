package application

import (
	"context"
	//"fmt"
	"encoding/json"

	"github.com/andriykutsevol/WeatherServer/internal/domain/weather"
)


type WeatherResponseDTO struct{
	MessageJson map[string]interface{}
}

type Weather interface {
	HandleGet(ctx context.Context, id string) (error, WeatherResponseDTO)
}

type weatherApp struct {
	weatherRepo weather.Repository
}

func NewWeather(weatherRepo weather.Repository) Weather {
	return &weatherApp{
		weatherRepo: weatherRepo,
	}
}


func (a *weatherApp) HandleGet(ctx context.Context, id string) (error, WeatherResponseDTO){

	response := WeatherResponseDTO{}

	app_response, _ := a.weatherRepo.Retrieve(ctx, id)

	body := []byte(app_response)

	// We're doing this because the "c.JSON" will double escape nested jsob string.
	// So we first decode it into a Go map or struct, then pass it to c.JSON.
	var messageMap map[string]interface{}
	json.Unmarshal(body, &messageMap)	
	response.MessageJson = messageMap



	return nil, response
}