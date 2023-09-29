package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"tomovale_deleta_duplicado/src/model"
	"tomovale_deleta_duplicado/src/response"
)

type InterfaceService interface {
	SearchPatientsInEndPointAndDeleteDuplicates(pacientes []model.Paciente)
	SearchSchedulesOnEndPointAndDelete(atendimentos []model.Agendamento)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

var apiUrl = "https://tomovale.qmatic.cloud/qsystem/rest/appointment/customers;cardNumber="
var endPointDeletePacienteQmatic = "https://tomovale.qmatic.cloud/rest/entrypoint/customers/%d/"
var apiBuscarAgendamentoUrl = "https://tomovale.qmatic.cloud/qsystem/rest/appointment/appointments/external/%d"
var apiDeleteAgendamentoUrl = "https://tomovale.qmatic.cloud/qsystem/rest/appointment/appointments/%d"

var headers = map[string]string{
	"auth-token":    "1f6331a7-6457-46c5-a616-05a9c7cfd139",
	"Authorization": "Basic aW50ZWdyYXRpb25hcHBvaW50OlVsYW5AMTIzNA==",
	"Referer":       "https://tomovale.qmatic.cloud/",
	"Cookie":        "SSOcookie=98ffcd09-4012-4e5b-900c-16ac8d53e17e; SSOcookie=d746245e-b5b9-48a2-840f-2bbbbeba4db6",
	"Content-Type":  "application/json",
}

var retornoEndPointPacientes []response.RetornoEndPointPacientesQmatic
var retornoEndPointAgendamentosQmatic response.RetornoAgendamentosQmatic

func (service *Service) SearchPatientsInEndPointAndDeleteDuplicates(pacientes []model.Paciente) {
	for _, paciente := range pacientes {

		if paciente.DsCpf.Valid {

			urlWithCPF := apiUrl + paciente.DsCpf.String

			req, err := http.NewRequest("GET", urlWithCPF, nil)

			if err != nil {
				fmt.Printf("Erro ao criar a solicitação HTTP para CPF %s: %v\n", paciente.DsCpf.String, err)
				continue
			}

			for key, value := range headers {
				req.Header.Add(key, value)
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				fmt.Printf("Erro ao fazer a solicitação GET para CPF %s: %v\n", paciente.DsCpf.String, err)
				continue
			}

			defer resp.Body.Close()

			if err = json.NewDecoder(resp.Body).Decode(&retornoEndPointPacientes); err != nil {
				log.Fatalf("Erro ao tentar realizar o unmarshal do input. Erro: %v", err)
			}

			switch {

			case len(retornoEndPointPacientes) > 1:

				for _, pacienteQmatic := range retornoEndPointPacientes[1:] {

					url := fmt.Sprintf(endPointDeletePacienteQmatic, pacienteQmatic.ID)

					request, err := http.NewRequest(http.MethodDelete, url, nil)

					if err != nil {
						log.Fatalf("Erro: %v", err)
					}

					for key, value := range headers {
						request.Header.Add(key, value)
					}

					cliente := &http.Client{}

					do, err := cliente.Do(request)

					if err != nil {
						fmt.Printf("Erro ao fazer a solicitação DELETE para ID %d: %v\n", pacienteQmatic.ID, err)
						continue
					}

					log.Println(do.StatusCode)

					var res string

					json.NewDecoder(resp.Body).Decode(&res)

					log.Printf("Paciente duplicado deletado: . %v, id: %v, cpf: %v, ", pacienteQmatic.FirstName,
						pacienteQmatic.ID, pacienteQmatic.CardNumber)

					time.Sleep(5 * time.Millisecond)
				}
			}

			if err != nil {
				fmt.Printf("Erro ao ler o corpo da resposta para CPF %s: %v\n", paciente.DsCpf.String, err)
				continue
			}
		}
	}
}

func (service *Service) SearchSchedulesOnEndPointAndDelete(agendamentos []model.Agendamento) {
	for _, agendamento := range agendamentos {

		switch {

		case agendamento.CdPaciente != 0:

			url := fmt.Sprintf(apiBuscarAgendamentoUrl, agendamento.CdPaciente)

			req, err := http.NewRequest(http.MethodGet, url, nil)

			if err != nil {
				fmt.Printf("Erro ao criar a solicitação GET para o cd_paciente %d: %v\n", agendamento.CdPaciente, err)
				continue
			}

			for key, value := range headers {
				req.Header.Add(key, value)
			}

			if err != nil {
				fmt.Printf("Erro ao fazer a solicitação GET para o id %v: %v\n", agendamento.CdPaciente, err)
				continue
			}

			client := &http.Client{}

			resp, err := client.Do(req)

			if err != nil {
				log.Fatalf("Erro: %v", err)
			}

			bytes, err := io.ReadAll(resp.Body)

			if err != nil {
				log.Fatalf("Erro ao tentar converter para bytes. Erro : %v", err)
			}

			if len(bytes) == 0 {
				fmt.Println("O JSON está vazio.")
				continue
			}

			if err = json.Unmarshal(bytes, &retornoEndPointAgendamentosQmatic); err != nil {
				log.Fatalf("Erro ao tentar deserializar a resposta. Erro: %v", err)
			}

			urlDelete := fmt.Sprintf(apiDeleteAgendamentoUrl, retornoEndPointAgendamentosQmatic.Id)

			req, err = http.NewRequest(http.MethodDelete, urlDelete, nil)

			if err != nil {
				fmt.Printf("Erro ao criar a solicitação GET para o cd_paciente %d: %v\n", agendamento.CdPaciente, err)
				continue
			}

			for key, value := range headers {
				req.Header.Add(key, value)
			}

			if err != nil {
				fmt.Printf("Erro ao fazer a solicitação GET para o id %v: %v\n", agendamento.CdPaciente, err)
				continue
			}

			client = &http.Client{}

			resp, err = client.Do(req)

			if err != nil {
				log.Fatalf("Erro: %v", err)
			}

			if resp.StatusCode == 204 {

				for _, costumer := range retornoEndPointAgendamentosQmatic.Customers {
					log.Printf("Id %v deletado. data do agendamento: %v. id: %v", retornoEndPointAgendamentosQmatic.Id,
						retornoEndPointAgendamentosQmatic.StartTime, costumer.Id)
				}

			}

		case agendamento.CdAgendamento != 0:

			url := fmt.Sprintf(apiBuscarAgendamentoUrl, agendamento.CdAgendamento)

			req, erro := http.NewRequest(http.MethodGet, url, nil)

			if erro != nil {
				fmt.Printf("Erro ao criar a solicitação GET para o cd_paciente %d: %v\n", agendamento.CdPaciente, erro)
				continue
			}

			for key, value := range headers {
				req.Header.Add(key, value)
			}

			if erro != nil {
				fmt.Printf("Erro ao fazer a solicitação GET para o id %v: %v\n", agendamento.CdPaciente, erro)
				continue
			}

			client := &http.Client{}

			resp, erro := client.Do(req)

			if erro != nil {
				log.Fatalf("Erro: %v", erro)
			}

			bytes, erro := io.ReadAll(resp.Body)

			if erro != nil {
				log.Fatalf("Erro ao tentar converter para bytes. Erro : %v", erro)
			}

			if len(bytes) == 0 {
				fmt.Println("Nenhum resultado encontrado para esse cd_paciente.")
				continue
			}

			if erro = json.Unmarshal(bytes, &retornoEndPointAgendamentosQmatic); erro != nil {
				log.Fatalf("Erro ao tentar deserializar a resposta. Erro: %v", erro)
			}

			if retornoEndPointAgendamentosQmatic.ExternalId == "" && retornoEndPointAgendamentosQmatic.Id == 0 {
				fmt.Println("tá vazio.")
				return
			}

			urlDelete := fmt.Sprintf(apiDeleteAgendamentoUrl, retornoEndPointAgendamentosQmatic.Id)

			req, erro = http.NewRequest(http.MethodDelete, urlDelete, nil)

			if erro != nil {

				fmt.Printf("Erro ao criar a solicitação DELETE para o cd_atendimento %d: %v\n", agendamento.CdAgendamento, erro)
				continue

			}

			for key, value := range headers {

				req.Header.Add(key, value)

			}

			if erro != nil {

				fmt.Printf("Erro ao fazer a solicitação GET para o id %v: %v\n", agendamento.CdPaciente, erro)
				continue

			}

			client = &http.Client{}

			resp, erro = client.Do(req)

			if erro != nil {
				log.Fatalf("Erro ao tentar enviar a requisição. Erro: %v", erro)
			}

			if resp.StatusCode == 204 {

				for _, costumer := range retornoEndPointAgendamentosQmatic.Customers {
					log.Printf("Id %v deletado. data do agendamento: %v. id: %v", retornoEndPointAgendamentosQmatic.Id,
						retornoEndPointAgendamentosQmatic.StartTime, costumer.Id)
				}

			}

		}
	}
}
