package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func setAPIKey() {
	file, err1 := os.Open("API_KEY.txt")
	if err1 != nil {
		log.Fatal(err1)
	}

	val, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Fatal(err2)
	}

	err3 := os.Setenv("API_KEY", string(val))
	if err3 != nil {
		log.Fatal(err3)
	}
}

func main() {
	setAPIKey()
	fetchGPSData()
	handleRequests()
}

func fetchGPSData() {
	uri := "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + os.Getenv("API_KEY")
	resp, err := http.Get(uri)

	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Print("resp", fromJSON(body))
}

func fromJSON(data []byte) []GpsDevice {
	var gpsDeviceList []GpsDevice

	err := json.Unmarshal(data, &gpsDeviceList)
	if err != nil {
		print(err)
	}

	return gpsDeviceList
}

func handleRequests() {
	http.HandleFunc("/", landingPage)
	http.HandleFunc("/echo", echo) //TODO remove
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//ConnectionData is subtype for GpsDevice
type ConnectionData struct {
	Calamp_next_lookup_time   string `json:"calamp_next_lookup_time"`
	Calamp_ipr_omega_fee_paid bool   `json:"calamp_ipr_omega_fee_paid"`
}

//Settings is subtype for GpsDevice
type Settings struct {
	Begin_moving_speed    SpeedAltitudeOrDistance `json:"begin_moving_speed"`
	Begin_stopped_speed   SpeedAltitudeOrDistance `json:"begin_stopped_speed"`
	Max_drift_distance    interface{}             `json:"max_drift_distance"`
	Min_num_satellites    int32                   `json:"min_num_satellites"`
	Max_hdop              float32                 `json:"max_hdop"`
	Drive_timeout         interface{}             `json:"drive_timeout"`
	Stop_timeout          interface{}             `json:"stop_timeout"`
	Offline_timeout       interface{}             `json:"offline_timeout"`
	History_calc_duration interface{}             `json:"history_calc_duration"`
}

//Params is subtype for LatestDevicePoint
type Params struct {
	Acc            string `json:"acc"`
	Gpslev         string `json:"gpslev"`
	Hdop           string `json:"hdop"`
	Rssi           string `json:"rssi"`
	Evtc           string `json:"evtc"`
	Acm12          string `json:"acm12"`
	Acm13          string `json:"acm13"`
	Acm0           string `json:"acm0"`
	Acm1           string `json:"acm1"`
	Acm10          string `json:"acm10"`
	Acm11          string `json:"acm11"`
	Acm2           string `json:"acm2"`
	Acm3           string `json:"acm3"`
	Acm5           string `json:"acm5"`
	Acm6           string `json:"acm6"`
	Acm7           string `json:"acm7"`
	Acm8           string `json:"acm8"`
	Acm9           string `json:"acm9"`
	Evti           string `json:"evti"`
	Input1         string `json:"input1"`
	Input2         string `json:"input2"`
	Last_known_loc string `json:"last_known_loc"`
	Seq_num        string `json:"seq_num"`
	Service_type   string `json:"service_type"`
	Valid_fix      string `json:"valid_fix"`
}

//DevicePointExternal is subtype for LatestDevicePoint
type DevicePointExternal struct {
	Posted_speed_limit        interface{} `json:"posted_speed_limit"`
	Software_odometer_reading interface{} `json:"software_odometer_reading"`
}

//DevicePointDetail is subtype for LatestDevicePoint
type DevicePointDetail struct {
	FactoryID               string                  `json:"factory_id"`
	TransmitTime            string                  `json:"transmit_time"`
	GpsTime                 string                  `json:"gps_time"`
	Acc                     bool                    `json:"acc"`
	LatLng                  LatLong                 `json:"lat_lng"`
	Altitude                SpeedAltitudeOrDistance `json:"altitude"`
	Speed                   SpeedAltitudeOrDistance `json:"speed"`
	Heading                 int32                   `json:"heading"`
	Hdop                    float32                 `json:"hdop"`
	NumSatellites           int32                   `json:"num_satellites"`
	MotionLog               MotionLog               `json:"motion_log"`
	PacketSequenceID        string                  `json:"packet_sequence_id"`
	Rssi                    int32                   `json:"rssi"`
	TripDistance            interface{}             `json:"tripD_distance"`
	TravelDistance          interface{}             `json:"travel_distance"`
	ExternalVolt            float32                 `json:"external_volt"`
	BackupBatteryVolt       float32                 `json:"backup_battery_volt"`
	VideoOriginalEvent      VideoOriginalEvent      `json:"video_original_event"`
	CalampServiceType       string                  `json:"calamp_service_type"`
	RemainingBatteryPercent interface{}             `json:"remaining_battery_percent"`
	Historic                bool                    `json:"historic"`
	SessionConnected        bool                    `json:"session_connected"`
	DiffCorrected           bool                    `json:"diff_corrected"`
	Predicted               bool                    `json:"predicted"`
	Input01high             bool                    `json:"input_1_high"`
	Input02high             bool                    `json:"input_2_high"`
	Input03high             bool                    `json:"input_3_high"`
	Input04high             bool                    `json:"input_4_high"`
	VbusOdometer            interface{}             `json:"vbus_odometer"`
	VbusSpeed               interface{}             `json:"vbus_speed"`
	VbusEngineOn            interface{}             `json:"vbus_engine_on"`
	VbusInMotion            interface{}             `json:"vbus_in_motion"`
	FuelPercent             interface{}             `json:"fuel_percent"`
	FuelLevelHeight         interface{}             `json:"fuel_level_height"`
	FuelTemp                interface{}             `json:"fuel_temp"`
	EngOilTemp              interface{}             `json:"eng_oil_temp"`
	EngHoursTotal           interface{}             `json:"eng_hours_total"`
	EngHoursDeviceDerived   interface{}             `json:"eng_hours_device_derived"`
	TransFluidTemp          interface{}             `json:"trans_fluid_temp"`
	EngRpm                  interface{}             `json:"eng_rpm"`
	ThrottlePercent         interface{}             `json:"throttle_percent"`
	CoolantTemp             interface{}             `json:"coolant_temp"`
	AmbientAirTemp          interface{}             `json:"ambient_air_temp"`
	VbusAvgFuelEconomy      interface{}             `json:"vbus_avg_fuel_economy"`
	BarometricPressure      interface{}             `json:"barometric_pressure"`
}

//LatLong is subtype for LatestDevicePoint via DevicePointDetail
type LatLong struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

//SpeedAltitudeOrDistance is subtype for LatestDevicePoint via DevicePointDetail
type SpeedAltitudeOrDistance struct {
	Value   float32 `json:"value"`
	Unit    string  `json:"unit"`
	Display string  `json:"display"`
}

//MotionLog is subtype for LatestDevicePoint via DevicePointDetail
type MotionLog struct {
	Start_time             string      `json:"start_time"`
	End_time               string      `json:"end_time"`
	Start_heading          interface{} `json:"start_heading"`
	End_heading            interface{} `json:"end_heading"`
	Max_accelerating_force interface{} `json:"max_accelerating_force"`
	Max_decelerating_force interface{} `json:"max_decelerating_force"`
	Max_right_turn_force   interface{} `json:"max_right_turn_force"`
	Max_left_turn_force    interface{} `json:"max_left_turn_force"`
}

//VideoOriginalEvent is subtype for LatestDevicePoint via DevicePointDetail
type VideoOriginalEvent struct {
	Hevent_type string `json:"hevent_type"`
}

//DeviceState is subtype for LatestDevicePoint via DevicePointDetail
type DeviceState struct {
	Drive_status                        string                  `json:"drive_status"`
	Drive_status_id                     string                  `json:"drive_status_id"`
	Drive_status_duration               SpeedAltitudeOrDistance `json:"drive_status_id"`
	Drive_status_distance               SpeedAltitudeOrDistance `json:"drive_status_duration"`
	Drive_status_lat_lng_distance       SpeedAltitudeOrDistance `json:"drive_status_lat_lng_distance"`
	Drive_status_begin_time             string                  `json:"drive_status_begin_time"`
	Adjusted_lat_lng                    LatLong                 `json:"adjusted_lat_lng"`
	Prev_drive_status_duration          interface{}             `json:"prev_drive_status_duration"`
	Prev_drive_status_distance          SpeedAltitudeOrDistance `json:"prev_drive_status_distance"`
	Prev_drive_status_lat_lng_distance  interface{}             `json:"prev_drive_status_lat_lng_distance"`
	Prev_drive_status_begin_time        interface{}             `json:"prev_drive_status_begin_time"`
	Prev_adjusted_lat_lng               interface{}             `json:"prev_adjusted_lat_lng"`
	Inaccurate_per_device_settings      bool                    `json:"inaccurate_per_device_settings"`
	Software_odometer                   interface{}             `json:"software_odometer"`
	Last_software_odometer_reading_time interface{}             `json:"last_software_odometer_reading_time"`
	Hardware_odometer                   interface{}             `json:"hardware_odometer"`
	odometer                            interface{}             `json:"odometer"`
}

//LatestDevicePoint is subtype for GpsDevice
type LatestDevicePoint struct {
	devicePointID         string              `json:"device_point_id"`
	dtServer              string              `json:"dt_server"`
	dtTracker             string              `json:"dt_tracker"`
	lat                   float32             `json:"lat"`
	lng                   float32             `json:"lat"`
	altitude              float32             `json:"altitude"`
	angle                 int16               `json:"angle"`
	speed                 int16               `json:"speed"`
	params                Params              `json:"params"`
	device_point_external DevicePointExternal `json:"device_point_external"`
	device_point_detail   DevicePointDetail   `json:"device_point_detail"`
	device_state          DeviceState         `json:"device_point_detail"`
	device_state_stale    bool                `json:"device_state_stale"`
	sequence              string              `json:"sequence"`
}

//GPS struct
type GpsDevice struct {
	device_id                    string            `json:"device_id"`
	created_at                   string            `json:"device_id"`
	updated_at                   string            `json:"updated_at"`
	activated_at                 string            `json:"activated_at"`
	delivered_at                 string            `json:"delivered_at"`
	factory_id                   string            `json:"factory_id"`
	active_state                 string            `json:"active_state"`
	display_name                 string            `json:"display_name"`
	bcc_id                       string            `json:"device_id"`
	make                         string            `json:"make"`
	model                        string            `json:"model"`
	conn_type                    string            `json:"conn_type"`
	conn_data                    ConnectionData    `json:"conn_data"`
	data_node                    string            `json:"data_node"`
	settings                     Settings          `json:"settings"`
	secondary_id                 string            `json:"secondary_id"`
	user_id_list                 []string          `json:"user_id_list"`
	online                       bool              `json:"online"`
	latest_device_point          LatestDevicePoint `json:"latest_device_point"`
	latest_accurate_device_point LatestDevicePoint `json:"latest_accurate_device_point"`
	device_groups_id_list        string            `json:"device_groups_id_list"`
	device_ui_settings           interface{}       `json:"device_ui_settings"`
}

// Article struct
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles struct array
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "ComputerRAM", Desc: "New RAM technologies", Content: "Yooooooooooo this ram is fast"},
	}
	fmt.Println("Articles endpoint hit!")
	json.NewEncoder(w).Encode(articles)
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Landing Page Hit")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	fmt.Println(message)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
