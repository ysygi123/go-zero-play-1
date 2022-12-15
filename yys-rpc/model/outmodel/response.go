package outmodel

type YysAccountInfoResponse struct {
	Equip struct {
		EquipDesc string `json:"equip_desc"`
		EquipName string `json:"equip_name"`
	} `json:"equip"`
}

type DetailNeedUn struct {
	Inventory map[string]struct {
		ItemID      int             `json:"itemId"`
		Rattr       [][]interface{} `json:"rattr"`
		Name        string          `json:"name"`
		DiscardTime int             `json:"discard_time"`
		Level       int             `json:"level"`
		NewGet      int             `json:"newGet"`
		Lock        bool            `json:"lock"`
		BaseRindex  int             `json:"base_rindex"`
		SingleAttr  []string        `json:"single_attr"`
		Pos         int             `json:"pos"`
		Qua         int             `json:"qua"`
		Suitid      int             `json:"suitid"`
		Isuseless   bool            `json:"isuseless"`
		Attrs       [][]string      `json:"attrs"`
		Exp         int             `json:"exp"`
		BaseR       float64         `json:"base_r"`
		UUID        string          `json:"uuid"`
	} `json:"inventory"`
}
