package weather 

import (
	"fmt"
	 "math"
	"log"
	"errors"
	owm "github.com/briandowns/openweathermap"
	"github.com/bwmarrin/discordgo"
	"github.com/DMXMax/hugin/command"
	"github.com/DMXMax/hugin/util"
)

const weather_key = "projects/131909995516/secrets/weather_key/versions/1"

type WindVelocity struct{
	KpH float64
	MpH float64
}


var WeatherToken string

func Init() error {
	wt ,err := util.GetSecret(weather_key)
	if err != nil {
		return errors.New("No weather token")	
	} else {
		WeatherToken = wt
	}
	return nil //No error
}

var WeatherCommand command.Command = command.Command{
	Name : "GetWeather",
	Scope : "any",
	Op : HandleWeatherRequest,
}

func HandleWeatherRequest(s *discordgo.Session, m *discordgo.MessageCreate) (map[string]string, error) {

	type City struct {
		name  string
		locID int
	}

	var cities = []City{City{"Millennium City and Windsor, Ontario", 4990729}}

	for _,c := range cities {

		w, err := owm.NewCurrent("C", "EN", WeatherToken)

		if err != nil {
			fmt.Println(err)
			s.ChannelMessageSend(m.ChannelID,"I cannot get the weather right now.")
		} else {
			w.CurrentByID(c.locID)
		var wv = convertWindVelocity(w.Wind.Speed)	
		s.ChannelMessageSend(m.ChannelID, 
		fmt.Sprintf("WCOC Action Weather For %v: %v.\nThe current "+
		"temperature is %v degrees Celcius, %v degrees Farenheit.\n"+
		"Humidity is %v%%, winds are out of the %v at %.1f kph (%.1f mph)",
			c.name,
			w.Weather[0].Description,
			math.Round(w.Main.Temp),
			math.Round((w.Main.Temp * 1.8) + 32),
			w.Main.Humidity,
			getWindDirection(w.Wind.Deg),
			wv.KpH, wv.MpH))
	
	
	}
}
	return nil,nil
}

func getWindDirection(dir float64) string{
	var strDirection string
	
	if dir < 0 {
		dir = 360-dir
	}
	switch {
	case dir <27.5:
		strDirection =  "north"
	case dir < 72.5:
		strDirection = "north east"
	case dir <  117.5:
		strDirection = "east"
	case dir < 162.5:
		strDirection = "south east"
	case dir <  207.5:
		strDirection = "south"
	case dir <  252.5:
		strDirection = "south west"
	case dir < 297.5:
		strDirection = "west"
	case dir  < 342.5:
		strDirection = "north west"
	case dir <= 360:
		strDirection = "north"
	default:
		strDirection = "Some Direction"
		log.Printf("Direction = %v\n", dir)
	}

	return strDirection
}
/* 
   Wind velocity comes in Meters per Second. This converts to Kilometers
   per hour and Miles per hour
*/
func convertWindVelocity(MpS float64) (converted WindVelocity){
	converted.KpH = MpS*3600/1000
	converted.MpH = converted.KpH*.621371
	return
}
