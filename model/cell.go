package model

type Cell struct {
	Pos    Pos   `json:"pos"`
	Start  bool  `json:"start_pos"`
	Locked bool  `json:"locked"`
	Enemy  Enemy `json:"enemy"`
}

func (c Cell) IsFree() bool {
	enemy := Enemy{}
	if c.Start || c.Locked || c.Enemy == enemy {
		return true
	}
	return false
}
