package model

type Cell struct {
	Pos    Pos   `json:"pos"`
	Start  Pos   `json:"start_pos"`
	Locked bool  `json:"locked"`
	IsOpen bool  `json:"is_open"`
	Enemy  Enemy `json:"enemy"`
}
