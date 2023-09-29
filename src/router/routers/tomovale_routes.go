package routers

import (
	"net/http"
	"tomovale_deleta_duplicado/src/handlers"
)

var tomovaleRoutes = []Route{
	{
		URI:    "/buscar",
		Method: http.MethodGet,
		Func:   handlers.BuscaPacientesNoClinux,
	},
	{
		URI:    "/deletar_agendamentos",
		Method: http.MethodGet,
		Func:   handlers.DeletarAgendamentosQmatic,
	},
}
