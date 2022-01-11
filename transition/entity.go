// ini berisi struct yg akan mengenerate tabel didatabase berdasarkan perancangan disini
package transition

import "time"

// ketika digenerate ke db maka nama tabelnya ke db menjd plural(sama sprti laravel), misal struct Book akan menjd Books. untuk mengenerate sehingga akan dibuatkan
// tabelnya maka perlu memakai autoMigration, jika tdk maka kita hrs buat tabel manual didb
type Penyimpanan struct {
	ID        int
	Judul     string
	Rating    int
	SubTitle  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
