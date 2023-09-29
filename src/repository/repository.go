package repository

import (
	"database/sql"
	"fmt"
	"log"
	"tomovale_deleta_duplicado/src/model"
)

type InterfaceRepository interface {
	BuscaPacientesNoClinux() []model.Paciente
	BuscaAgendamentoNoClinux() ([]model.Agendamento, error)
	BuscarPacientesNoClinuxHoje() []model.Paciente
}

type Repository struct {
	DB *sql.DB
}

func NewServiceRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (repository *Repository) BuscaPacientesNoClinux() (pacientes []model.Paciente) {
	begin, erro := repository.DB.Begin()

	if erro != nil {
		log.Fatalf("Erro ao tentar começar uma transação no banco de dados. Erro: %v", erro)
	}

	rows, erro := begin.Query(`select 
									ds_paciente, pa.ds_cpf, cd_paciente
									from atendimentos ae
									join pacientes pa using (cd_paciente)
									where dt_envio_qmatic is not null
									ORDER BY ds_paciente`)

	if erro != nil {
		log.Fatalf("Erro ao tentar executar a query de consulta de pacientes")
	}

	for rows.Next() {
		var paciente model.Paciente

		if erro = rows.Scan(&paciente.DsPaciente, &paciente.DsCpf, &paciente.CdPaciente); erro != nil {
			log.Fatalf("Erro ao tentar executar o scan. Erro : %v", erro)
			return
		}
		pacientes = append(pacientes, paciente)
	}

	return pacientes
}

func (rep *Repository) BuscaAgendamentoNoClinux() ([]model.Agendamento, error) {
	rows, err := rep.DB.Query(`SELECT cd_atendimento, 
cd_paciente FROM atendimentos WHERE dt_data >= '2023-03-01' and cd_paciente is not null order by dt_data`)

	if err != nil {
		log.Fatalf("Erro ao tentar executar a query. Erro: %v", err)
	}

	defer rows.Close()

	var agendamentos []model.Agendamento

	for rows.Next() {

		var agendamento model.Agendamento

		if err := rows.Scan(&agendamento.CdAgendamento, &agendamento.CdPaciente); err != nil {
			log.Fatalf("Erro ao tentar escanear. Erro: %v", err)
		}
		agendamentos = append(agendamentos, agendamento)

	}

	return agendamentos, nil
}

func (rep *Repository) BuscarPacientesNoClinuxHoje() []model.Paciente {
	rows, erro := rep.DB.Query(`select ds_paciente, pa.ds_cpf, 
									cd_paciente from atendimentos ae join pacientes pa using (cd_paciente) where dt_data = '2023-09-29' ORDER BY ds_paciente`)

	if erro != nil {
		log.Fatalf("Erro ao tentar executar a query. Erro: %v", erro)
	}

	defer rows.Close()

	var pacientes []model.Paciente

	for rows.Next() {
		var paciente model.Paciente

		if erro = rows.Scan(
			&paciente.CdPaciente,
			&paciente.DsPaciente,
			&paciente.DsCpf,
		); erro != nil {
			log.Fatalf("Erro ao tentar escanear o resultado da query. Erro: %v", erro)
		}
		pacientes = append(pacientes, paciente)
	}

	fmt.Println(pacientes)

	return pacientes

}
