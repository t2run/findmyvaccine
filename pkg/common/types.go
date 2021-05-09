package common

//StateList model
type StateList struct {
	States []State `json:"states"`
	TTL    int     `json:"ttl"`
}

//State model
type State struct {
	StateID   int    `json:"state_id"`
	StateName string `json:"state_name"`
}

//DistrictList model
type DistrictList struct {
	Districts []District `json:"districts"`
	TTL       int        `json:"ttl"`
}

//District model
type District struct {
	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

type SessionList struct {
	Sessions []struct {
		CenterID          int      `json:"center_id"`
		Name              string   `json:"name"`
		Address           string   `json:"address"`
		StateName         string   `json:"state_name"`
		DistrictName      string   `json:"district_name"`
		BlockName         string   `json:"block_name"`
		Pincode           int      `json:"pincode"`
		From              string   `json:"from"`
		To                string   `json:"to"`
		Lat               int      `json:"lat"`
		Long              int      `json:"long"`
		FeeType           string   `json:"fee_type"`
		SessionID         string   `json:"session_id"`
		Date              string   `json:"date"`
		AvailableCapacity int      `json:"available_capacity"`
		Fee               string   `json:"fee"`
		MinAgeLimit       int      `json:"min_age_limit"`
		Vaccine           string   `json:"vaccine"`
		Slots             []string `json:"slots"`
	} `json:"sessions"`
}
