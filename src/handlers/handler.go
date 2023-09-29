package handlers

import (
	"log"
	"net/http"
	"tomovale_deleta_duplicado/src/database"
	"tomovale_deleta_duplicado/src/repository"
	"tomovale_deleta_duplicado/src/service"
)

func BuscaPacientesNoClinux(w http.ResponseWriter, r *http.Request) {
	db, erro := database.GetConnection()
	if erro != nil {
		log.Fatalf("erro ao tentar pegar a conexão com banco de dados. Erro :%v", erro)
	}

	rep := repository.NewServiceRepository(db)

	pacientes := rep.BuscaPacientesNoClinux()

	serv := service.NewService()

	serv.SearchPatientsInEndPointAndDeleteDuplicates(pacientes)

}

func DeletarAgendamentosQmatic(w http.ResponseWriter, r *http.Request) {
	db, erro := database.GetConnection()

	if erro != nil {
		log.Fatalf("erro ao tentar pegar a conexão com banco de dados. Erro :%v", erro)
	}

	rep := repository.NewServiceRepository(db)

	serv := service.NewService()

	agendamentosClinux, err := rep.BuscaAgendamentoNoClinux()

	if err != nil {
		log.Fatalf("Erro ao tentar buscar os agendamentos do clinux. Erro: %v", err)
	}

	pacientes := rep.BuscaPacientesNoClinux()

	serv.SearchSchedulesOnEndPointAndDelete(agendamentosClinux)

	serv.SearchPatientsInEndPointAndDeleteDuplicates(pacientes)
}
