package models

type Smart_Data struct {
	Meter_id    string `binding:"required" json:"meter_id"`
    Consumption string `binding:"required" json:"consumption"`
}

