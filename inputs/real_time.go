package inputs

type RealTime struct {
	Meta []interface{} `json:"meta"`
	Data Data          `json:"data"`
}

type Data struct {
	Temperature  float64     `json:"temperature"`
	Setpoint     float64     `json:"setpoint"`
	Consumption  float64     `json:"consumption"`
	Dissipation  float64     `json:"dissipation"`
	DissipationC float64     `json:"dissipationC"`
	DissipationW float64     `json:"dissipationW"`
	Mpue         float64     `json:"mpue"`
	Pump1Status  float64     `json:"pump1status"`
	Pump1RPM     float64     `json:"pump1rpm"`
	Pump2Status  float64     `json:"pump2status"`
	Pump2RPM     float64     `json:"pump2rpm"`
	CTI          float64     `json:"cti"`
	Cto          float64     `json:"cto"`
	CF           float64     `json:"cf"`
	Wti          float64     `json:"wti"`
	Wto          float64     `json:"wto"`
	Wf           float64     `json:"wf"`
	Alarm        float64     `json:"alarm"`
	Cpu0Temp     float64     `json:"cpu0temp"`
	Cpu1Temp     float64     `json:"cpu1temp"`
	Errors       []Error     `json:"errors"`
	Warnings     []Warning   `json:"warnings"`
	Mode         string      `json:"mode"`
	Test         interface{} `json:"test"`
	Maintenance  interface{} `json:"maintenance"`
	Demo         bool        `json:"demo"`
	Factory      bool        `json:"factory"`
}

type Error struct {
	IDFailure   string `json:"idFailure"`
	StartTime   string `json:"startTime"`
	FailureType string `json:"failureType"`
	Description string `json:"description"`
}

type Warning struct {
	IDWarning   string `json:"idWarning"`
	StartTime   string `json:"startTime"`
	WarningType string `json:"warningType"`
	Description string `json:"description"`
}
