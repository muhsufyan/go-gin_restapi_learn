package transition

import "time"

type Penyimpanan struct {
	ID        int
	Judul     string
	Rating    int
	SubTitle  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
