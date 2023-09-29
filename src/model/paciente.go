package model

import "database/sql"

type Paciente struct {
	CdPaciente float64
	DsPaciente string
	DsCpf      sql.NullString
}
