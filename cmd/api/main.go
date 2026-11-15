package main // Define o pacote executável.

import ( // Inicia as importações.
	"log"      // Permite registrar mensagens.
	"net/http" // Disponibiliza o servidor HTTP.
	"time"     // Disponibiliza configurações de tempo.

	"github.com/devjuniorhanun/go-api/internal/handler" // Importa os handlers.
	"github.com/devjuniorhanun/go-api/internal/store"   // Importa o armazenamento.
) // Finaliza as importações.

// main é o ponto de entrada da aplicação.
func main() {
	memory := store.NewMemory()          // Inicializa o armazenamento.
	auth := &handler.Auth{Store: memory} // Cria os handlers.
	mux := http.NewServeMux()            // Cria o roteador HTTP.

	mux.HandleFunc("POST /users", auth.Register) // Registra a rota de cadastro.
	mux.HandleFunc("POST /login", auth.Login)    // Registra a rota de login.

	server := &http.Server{ // Configura o servidor.
		Addr:              ":8081",         // Define a porta HTTP.
		Handler:           mux,             // Define o roteador.
		ReadHeaderTimeout: 5 * time.Second, // Limita o tempo de leitura dos cabeçalhos.
	} // Finaliza a configuração.

	log.Println("API iniciada em http://localhost:8081") // Exibe o endereço.
	log.Fatal(server.ListenAndServe())                   // Inicia o servidor HTTP.
} // Finaliza a função principal.
