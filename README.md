# redcoins-api

---
Redcoins-api é uma API de autenticação de usuário e operação de compra/venda de cripto moedas.

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

2 - Execução dos testes unitários
  - use@user:~/diretorio_projeto_clonado/ **go test ./... -v**

### Rotas da API ###
|   Ação                   | Requerido  | Role  |  Método  | URL
|   -----------------------|------------| ----- |----------|--------------
|   AUTENTICAR USUÁRIO     |            |       | `POST`   | /api/v1/users/login
|   CRIAR USUÁRIO          |            |       | `POST`   | /api/v1/users/signup
|   CRIAR OPERAÇÃO         | Autenticar | User  | `POST`   | /api/v1/operations
|   BUSCAR OPERAÇÕES       | Autenticar | User  | `GET  `  | /api/v1/operations
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `GET`    | /api/v1/operations/email
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `GET`    | /api/v1/operations/date
|   BUSCAR OPERAÇÕES       | Autenticar | Admin | `GET`    | /api/v1/operations/name

<!-- #### AUTENTICAR USUÁRIO ####
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
  "token": "eyJhbGciOiJI..."
}
```

#### CRIAR USUÁRIO ####
* REQUISIÇÃO
```
POST /api/v1/users
```
```json
{
  "name": "user",
  "email": "user@gmail.com",
  "password": "user123"
}
```
* RESPOSTA
```json
{
  "message": "User registered with success",
  "token": "eyJhbGciOiJI..."
}
```

#### OBTER USUÁRIO POR TOKEN ####
* REQUISIÇÃO
```
GET /api/v1/users
```
```javascript
const token = 'eyJhbGciOiJI...';
req.setRequestHeader('Authorization', token);
```
* RESPOSTA
```json
{
  "id": 46643,
  "name": "user",
  "email": "user@gmail.com",
  "password": "user123",
  "imageUrl": "https://almsaeedstudio.com/themes/AdminLTE/dist/img/user2-160x160.jpg"
}
```

#### ATUALIZAR USUÁRIO POR TOKEN ####
* REQUISIÇÃO
```
PATCH /api/v1/users
```
```javascript
const token = 'eyJhbGciOiJI...';
req.setRequestHeader('Authorization', token);
```
```json
{
  "name": "newName",
  "email": "new_email@gmail.com",
  "password": "newpassword123"
}
```
* RESPOSTA
```json
{
  "message": "User updated with success"
}
```

#### DELETAR USUÁRIO POR TOKEN ####
* REQUISIÇÃO
```
DELETE /api/v1/users
```
```javascript
const token = 'eyJhbGciOiJI...';
req.setRequestHeader('Authorization', token);
```
* RESPOSTA
```json
{
  "message": "User deleted with success"
}
```

#### DELETAR TODOS USUÁRIOS ####
* REQUISIÇÃO
```
DELETE /api/v1/all-users
```
```json
{
  "key_admin": "keyadmin123"
}
```
* RESPOSTA
```json
{
  "message": "Users deleted with success"
}
``` -->

### Autor

* Thiago Luiz Pereira Nunes ([ThiagoLuizNunes](https://github.com/ThiagoLuizNunes)) thiagoluiz.dev@gmail.com

### Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo [LICENSE.md](LICENSE.md) para obter detalhes

>Criado por **[ThiagoLuizNunes](https://www.linkedin.com/in/thiago-luiz-507483112/)** 2019.

---
