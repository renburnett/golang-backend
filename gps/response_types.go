package gps

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
	TripDistance            interface{}             `json:"trip_distance"`
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
	HeventType string `json:"hevent_type"`
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
	Odometer                            interface{}             `json:"odometer"`
}

//LatestDevicePoint is subtype for GpsDevice
type LatestDevicePoint struct {
	DevicePointID         string              `json:"device_point_id"`
	DtServer              string              `json:"dt_server"`
	DtTracker             string              `json:"dt_tracker"`
	Lat                   float32             `json:"lat"`
	Lng                   float32             `json:"lat"`
	Altitude              float32             `json:"altitude"`
	Angle                 int16               `json:"angle"`
	Speed                 int16               `json:"speed"`
	Params                Params              `json:"params"`
	Device_point_external DevicePointExternal `json:"device_point_external"`
	Device_point_detail   DevicePointDetail   `json:"device_point_detail"`
	Device_state          DeviceState         `json:"device_point_detail"`
	Device_state_stale    bool                `json:"device_state_stale"`
	Sequence              string              `json:"sequence"`
}

//Device models GPS device
type Device struct {
	DeviceID                  string            `json:"device_id"`
	CreatedAt                 string            `json:"device_id"`
	UpdatedAt                 string            `json:"updated_at"`
	ActivatedAt               string            `json:"activated_at"`
	DeliveredAt               string            `json:"delivered_at"`
	FactoryID                 string            `json:"factory_id"`
	ActiveState               string            `json:"active_state"`
	DisplayName               string            `json:"display_name"`
	BccID                     string            `json:"device_id"`
	Make                      string            `json:"make"`
	Model                     string            `json:"model"`
	ConnType                  string            `json:"conn_type"`
	ConnData                  ConnectionData    `json:"conn_data"`
	DataNode                  string            `json:"data_node"`
	Settings                  Settings          `json:"settings"`
	SecondaryID               string            `json:"secondary_id"`
	UserIDList                []string          `json:"user_id_list"`
	Online                    bool              `json:"online"`
	LatestDevicePoint         LatestDevicePoint `json:"latest_device_point"`
	LatestAccurateDevicePoint LatestDevicePoint `json:"latest_accurate_device_point"`
	DeviceGroupsIDList        string            `json:"device_groups_id_list"`
	DeviceUISettings          interface{}       `json:"device_ui_settings"`
}

// Article struct
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// APIResponse is array of gps devices
type APIResponse struct {
	ResultList []GpsDevice `json:"result_list"`
}
