package store // Define o pacote de armazenamento.

import ( // Inicia as importações.
	"strings" // Disponibiliza operações com textos.
	"sync"    // Disponibiliza mecanismos de sincronização.

	"github.com/devjuniorhanun/go-api/internal/user" // Importa a entidade.
) // Finaliza as importações.

// Memory representa nosso armazenamento temporário.
type Memory struct {
	mu     sync.Mutex           // Protege operações simultâneas.
	users  map[string]user.User // Armazena usuários pelo e-mail.
	nextID int                  // Controla os identificadores.
}

// NewMemory cria um armazenamento vazio.
func NewMemory() *Memory {
	return &Memory{ // Retorna uma nova estrutura.
		users:  make(map[string]user.User), // Inicializa o mapa.
		nextID: 1,                          // Define o primeiro identificador.
	} // Finaliza a estrutura.
} // Finaliza a função.

// Save cadastra um usuário caso o e-mail esteja disponível.
func (m *Memory) Save(u user.User) (user.User, bool) {
	m.mu.Lock()                              // Bloqueia acessos simultâneos.
	defer m.mu.Unlock()                      // Libera o bloqueio ao terminar.
	email := strings.ToLower(u.Email)        // Normaliza o e-mail.
	if _, exists := m.users[email]; exists { // Verifica duplicidade.
		return user.User{}, false // Informa que já existe.
	} // Finaliza a verificação.
	u.ID = m.nextID    // Atribui o identificador.
	m.nextID++         // Incrementa o contador.
	m.users[email] = u // Armazena o usuário.
	return u, true     // Retorna o usuário cadastrado.
} // Finaliza a função.

// FindByEmail procura um usuário pelo e-mail.
func (m *Memory) FindByEmail(email string) (user.User, bool) {
	m.mu.Lock()                                 // Protege o acesso ao mapa.
	defer m.mu.Unlock()                         // Libera o bloqueio.
	u, found := m.users[strings.ToLower(email)] // Busca o usuário.
	return u, found                             // Retorna o resultado.
} // Finaliza a função.
