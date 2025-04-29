API Rest para administrar clientes, quartos e reservas de projeto acadêmico da UMC (Universidade de Mogi das Cruzes), ele foi feito seguindo princípios REST e MVC

## 🛠️ Tecnologias Utilizadas
- Go
- MySQL
- Docker

## 🚀 Como Executar o Projeto
Antes de executar, você deve configurar o projeto, renomeie o arquivo `config.toml.example` para `config.toml`
```toml
mysql_dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
# Mude para "root:password@tcp(db)/hotel?charset=utf8mb4&parseTime=True&loc=Local" caso vá rodar com docker

[server]
port=8080
# Caso altere a porta, deve alterar também nos arquivos Dockerfile e docker-compose.yml
```
### ✅ Rodando com Go
```bash
# Clone o repositório
git clone https://github.com/avrahambenaram/hotel-backend.git

# Acesse a pasta do projeto
cd hotel-backend

# Configure o config.toml conforme explicado anteriormente

# Execute a aplicação
go run .
```

### 🐳 Rodando com Docker
```bash
# Clone o repositório
git clone https://github.com/avrahambenaram/hotel-backend.git

# Acesse a pasta do projeto
cd hotel-backend

# Configure o config.toml conforme explicado anteriormente

docker-compose up -d
```


## 📌 Funcionalidades
- 👤 Clientes: Listagem, criação, atualização e exclusão
- 🛏️ Quartos: Listagem, criação, atualização e exclusão
- 🟢 Reservas: Listagem, criação e exclusão de reservas
