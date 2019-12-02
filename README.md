# Red Coins API

---
Red Coins API é uma API de autenticação de usuário e operação de compra/venda de cripto moedas.

**Objetivo**: desafio técnico Red Ventures.


Ferramentas: Golang | Docker | Docker-compose

### Requerimentos ###

É necessário que a instalação prévia do Docker e do Docker-compose

* **[Docker 19.03.x](https://docs.docker.com)** :white_check_mark:
* **[Docker compose 1.24.x](https://docs.docker.com/compose/)** :white_check_mark:

### Instalação do Projeto ###

**Obs.: As seguintes instruções foram testadas na distribuição do Ubuntu 18.04 Bionic**

1 - Depois de clonar o repositório 'git clone' (comando), execute os seguintes comandos para criar as imagens docker "redcoins-api", "mysqldb" e "swagger-ui":
  - user@user:~/diretorio_projeto_clonado/$ **docker-compose up -d**
  - certifique-se se as portas :8000, :8001 e :3306 estão liberadas
  - acessa red-coins-api http://localhost:8000
  - acessa swagger interface http://localhost:8001

2 - Execução dos testes unitários
  - use@user:~/diretorio_projeto_clonado/ **go test ./... -v**

### Rotas da API ###
|   Ação                   | Requerido  | Role  |  Método  | URL
|   -----------------------|------------| ----- |----------|--------------
|   AUTENTICAR USUÁRIO     |            |       | `POST`   | /api/v1/users/login
|   CRIAR USUÁRIO          |            |       | `POST`   | /api/v1/users/signup
|   CRIAR OPERAÇÃO         | Autenticar | User  | `POST`   | /api/v1/operations
|   BUSCAR OPERAÇÕES       | Autenticar | User/Admin  | `GET  `  | /api/v1/operations
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `POST`   | /api/v1/operations/email
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `GET`    | /api/v1/operations/date/{date}
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `GET`    | /api/v1/operations/name/{name}

#### AUTENTICAR USUÁRIO ####
* REQUISIÇÃO
```
POST /api/v1/users/login
```
```json
{
  "email": "user@gmail.com",
  "password": "user123"
}
```
* RESPOSTA
```json
{
  "code": "200",
  "token": "eyJhbGciOiJI..."
}
```
#### CRIAR USUÁRIO ####
* REQUISIÇÃO
```
POST /api/v1/users/signup
```
```json
{
  "name": "User test",
  "email": "user@gmail.com",
  "password": "user123",
  "confirm_password": "user123",
  "secret": "" // Para criação de usuário Admin
}
```
* RESPOSTA
```json
{
  "code": "201",
  "message": "User registered with success."
}
```
#### CRIAR OPERAÇÔES ####
* REQUISIÇÃO
```
POST /api/v1/operations
```
```json
{
  "operation_type": "purchase", // purchase ou sale
  "amount": "150",
}
```
* RESPOSTA
```json
{
  "code": "200",
  "message": "Operation successfully performed."
}
```
#### RECUPERAR OPERAÇÔES ####
* REQUISIÇÃO
```
GET /api/v1/operations
```
* RESPOSTA
```json
{
  "code": "200",
  "operations": [
    {
      ...
    },
  ]
}
```
#### RECUPERAR OPERAÇÔES BY DATE ####
* REQUISIÇÃO
```
GET /api/v1/operations/date/{date}
```
* RESPOSTA
```json
{
  "code": "200",
  "operations": [
    {
      ...
    },
  ]
}
```
#### RECUPERAR OPERAÇÔES BY NAME ####
* REQUISIÇÃO
```
GET /api/v1/operations/name/{name}
```
* RESPOSTA
```json
{
  "code": "200",
  "operations": [
    {
      ...
    },
  ]
}
```
#### RECUPERAR OPERAÇÔES BY EMAIL ####
* REQUISIÇÃO
```
POST /api/v1/operations/email
```
```json
{
  "email": "admin@gmail.com",
}
```
* RESPOSTA
```json
{
  "code": "200",
  "operations": [
    {
      ...
    },
  ]
}
```

### Autor

* Thiago Luiz Pereira Nunes ([ThiagoLuizNunes](https://github.com/ThiagoLuizNunes)) thiagoluiz.dev@gmail.com

### Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo [LICENSE.md](LICENSE.md) para obter detalhes

>Criado por **[ThiagoLuizNunes](https://www.linkedin.com/in/thiago-luiz-507483112/)** 2019.

---
