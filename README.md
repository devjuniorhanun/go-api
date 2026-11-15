# Go API Auth — Cadastro de usuários e login

API REST básica desenvolvida em **Go**, para estudar a criação de um servidor HTTP, cadastro de usuários e validação de login com e-mail e senha.

O projeto utiliza a biblioteca padrão `net/http`, `encoding/json` e a dependência `golang.org/x/crypto/bcrypt` para armazenar **hashes de senhas**, nunca senhas em texto puro. Os usuários são armazenados **em memória** e são perdidos ao reiniciar a aplicação.

## Requisitos

- Go **1.27 ou superior** (necessário para os padrões de rota `POST /users` e `POST /login`).
- Visual Studio Code com a extensão **REST Client**, para executar as requisições do arquivo `api.http`.

## Estrutura do projeto

```text
go-api-auth/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── user/
│   │   └── user.go
│   ├── store/
│   │   └── memory.go
│   └── handler/
│       └── auth.go
├── api.http              # Requisições de teste com REST Client
├── README.md             # Documentação do projeto
├── go.mod
└── go.sum
```

## Instalação

No diretório raiz do projeto:

```bash
go get golang.org/x/crypto/bcrypt
go mod tidy
```

> Se o arquivo `go.mod` já existir, **não execute novamente** `go mod init`. O caminho `example.com/go-api-auth` deve coincidir com os imports usados no código.

## Executando a API

```bash
go run ./cmd/api
```

O servidor estará disponível em `http://localhost:8081`.

## Endpoints

| Método | Rota | Descrição | Resposta de sucesso |
| --- | --- | --- | --- |
| `POST` | `/users` | Cadastra usuário com nome, e-mail e senha | `201 Created` |
| `POST` | `/login` | Verifica e-mail e senha | `200 OK` |

### O Cadastro de usuário e o Login estão no arquivo
```
api.http
```

## Testes pelo arquivo `api.http` usando REST Client

Foi criado o arquivo **`api.http`** na raiz do projeto, com requisições HTTP prontas para testar o cadastro, o login correto, a senha incorreta e a tentativa de cadastro duplicado.

1. Abra o projeto no **Visual Studio Code**.
2. Instale a extensão **REST Client** (autor: **Huachao Mao**, identificador `humao.rest-client`).
3. Inicie a API com `go run ./cmd/api` em um terminal.
4. Abra o arquivo **`api.http`** no editor.
5. Clique em **Send Request** acima da requisição que deseja executar.
6. Confira o código HTTP e o corpo da resposta na aba aberta pela extensão.

**Ordem recomendada:** execute primeiro **Cadastrar usuário**, depois **Login válido**. As requisições usam o mesmo usuário de exemplo. Para testar **E-mail duplicado**, execute o cadastro novamente.

## Códigos HTTP esperados

| Código | Situação |
| --- | --- |
| `200 OK` | Login válido |
| `201 Created` | Usuário cadastrado |
| `400 Bad Request` | JSON ou dados inválidos |
| `401 Unauthorized` | E-mail ou senha incorretos |
| `409 Conflict` | E-mail já cadastrado |
| `500 Internal Server Error` | Falha interna |

## Limitações e próximos passos

Este é um projeto **didático**, não preparado para produção. Ainda faltam validação robusta dos dados, limitação do tamanho das requisições, proteção contra tentativas repetidas de login, HTTPS, persistência em banco de dados, testes automatizados e emissão de tokens de acesso. As próximas etapas poderão introduzir TDD, contratos, casos de uso, SOLID e Clean Architecture.
