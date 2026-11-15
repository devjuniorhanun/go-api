package user // Define o pacote responsável pelos usuários.

// User representa um usuário cadastrado.
type User struct {
	ID           int    `json:"id"`    // Identificador único.
	Name         string `json:"name"`  // Nome do usuário.
	Email        string `json:"email"` // E-mail do usuário.
	PasswordHash string `json:"-"`     // Hash da senha, oculto no JSON.
}
