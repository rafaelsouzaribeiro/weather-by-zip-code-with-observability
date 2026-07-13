package dto

type TempResponseOutput struct {
	Currents Current `json:"current"`
}

type Current struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
}

type TempResponseInput struct {
	Currents Current `json:"current"`
}
