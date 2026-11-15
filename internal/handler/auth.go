package handler // Define o pacote dos handlers HTTP.

import ( // Inicia as importações.
	"encoding/json" // Permite converter JSON.
	"net/http"      // Disponibiliza recursos HTTP.
	"strings"       // Permite manipular textos.

	"github.com/devjuniorhanun/go-api/internal/store" // Importa o armazenamento.
	"github.com/devjuniorhanun/go-api/internal/user"  // Importa a entidade.
	"golang.org/x/crypto/bcrypt"                      // Disponibiliza hash de senhas.
) // Finaliza as importações.

// Auth recebe as requisições de cadastro e login.
type Auth struct {
	Store *store.Memory // Referência ao armazenamento.
}

// Credentials representa os dados recebidos do cliente.
type Credentials struct {
	Name     string `json:"name"`     // Nome informado.
	Email    string `json:"email"`    // E-mail informado.
	Password string `json:"password"` // Senha informada.
}

// Register realiza o cadastro de usuários.
func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var input Credentials                         // Declara os dados recebidos.
	err := json.NewDecoder(r.Body).Decode(&input) // Lê o JSON.
	if err != nil {                               // Verifica erros de leitura.
		http.Error(w, "JSON inválido", http.StatusBadRequest) // Retorna erro.
		return                                                // Interrompe a execução.
	} // Finaliza a verificação.

	input.Email = strings.TrimSpace(input.Email)                                            // Remove espaços externos.
	if input.Name == "" || !strings.Contains(input.Email, "@") || len(input.Password) < 8 { // Realiza validações iniciais.
		http.Error(w, "Dados inválidos", http.StatusBadRequest) // Retorna erro.
		return                                                  // Interrompe a execução.
	} // Finaliza a validação.

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost) // Gera o hash.
	if err != nil {                                                                      // Verifica falha na geração.
		http.Error(w, "Erro interno", http.StatusInternalServerError) // Retorna erro.
		return                                                        // Interrompe a execução.
	} // Finaliza a verificação.

	newUser := user.User{ // Prepara o novo usuário.
		Name:         input.Name,   // Define o nome.
		Email:        input.Email,  // Define o e-mail.
		PasswordHash: string(hash), // Armazena o hash.
	} // Finaliza a estrutura.

	saved, ok := a.Store.Save(newUser) // Cadastra o usuário.
	if !ok {                           // Verifica duplicidade.
		http.Error(w, "E-mail já cadastrado", http.StatusConflict) // Retorna HTTP 409.
		return                                                     // Interrompe a execução.
	} // Finaliza a verificação.

	w.Header().Set("Content-Type", "application/json") // Define a resposta JSON.
	w.WriteHeader(http.StatusCreated)                  // Define HTTP 201.
	json.NewEncoder(w).Encode(saved)                   // Retorna o usuário cadastrado.
} // Finaliza o cadastro.

// Login verifica as credenciais informadas.
func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var input Credentials                         // Declara os dados recebidos.
	err := json.NewDecoder(r.Body).Decode(&input) // Lê o JSON.
	if err != nil {                               // Verifica JSON inválido.
		http.Error(w, "JSON inválido", http.StatusBadRequest) // Retorna erro.
		return                                                // Interrompe a execução.
	} // Finaliza a verificação.

	saved, found := a.Store.FindByEmail(strings.TrimSpace(input.Email)) // Procura o usuário.
	if !found {                                                         // Verifica se foi encontrado.
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized) // Retorna HTTP 401.
		return                                                          // Interrompe a execução.
	} // Finaliza a verificação.

	err = bcrypt.CompareHashAndPassword([]byte(saved.PasswordHash), []byte(input.Password)) // Compara a senha.
	if err != nil {                                                                         // Verifica se a senha está incorreta.
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized) // Retorna HTTP 401.
		return                                                          // Interrompe a execução.
	} // Finaliza a verificação.

	w.Header().Set("Content-Type", "application/json")                      // Define JSON.
	json.NewEncoder(w).Encode(map[string]string{"message": "Login válido"}) // Confirma as credenciais.
} // Finaliza o login.
