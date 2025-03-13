# Go CRUD Pet API

Este repositório contém uma aplicação CRUD (Create, Read, Update, Delete) simples para gerenciar informações sobre pets. A aplicação foi desenvolvida em Go (Golang) e utiliza o framework **Gin** para gerenciamento de rotas e o **GORM** para interação com o banco de dados **PostgreSQL**. Além disso, a aplicação é configurada para ser executada em um ambiente Docker utilizando o `docker-compose`.

## Pré-requisitos

Antes de executar a aplicação, certifique-se de que você possui os seguintes requisitos instalados em sua máquina:

1. **Docker**: A aplicação utiliza Docker para containerização. Você pode baixar e instalar o Docker a partir do [site oficial](https://www.docker.com/).

2. **Docker Compose**: O `docker-compose` é utilizado para orquestrar os containers da aplicação e do banco de dados. Ele geralmente vem incluído com a instalação do Docker, mas você pode verificar a instalação executando `docker-compose --version`.

3. **Git**: Para clonar o repositório, você precisa ter o Git instalado. Você pode baixar e instalar o Git a partir do [site oficial](https://git-scm.com/).

## Passo a Passo para Executar a Aplicação

Siga os passos abaixo para clonar, configurar e executar a aplicação em sua máquina:

### 1. Clonar o Repositório

Primeiro, clone o repositório para o seu ambiente local:

```bash
git clone https://github.com/Eduardo-Rosa/go-crud-pet-api.git
cd go-crud-pet-api
```

### 2. Configurar Variáveis de Ambiente

A aplicação utiliza variáveis de ambiente para configurar a conexão com o banco de dados. Crie um arquivo `.env` na raiz do projeto (caso ainda não exista) e adicione as seguintes variáveis:

```env
POSTGRES_USER=root
POSTGRES_PASSWORD=pass
POSTGRES_DB=pet_db
POSTGRES_HOST=db
POSTGRES_PORT=5432
```

### 3. Executar a Aplicação com Docker Compose

Para subir a aplicação e o banco de dados PostgreSQL, utilize o seguinte comando:

```bash
docker-compose up --build
```

Isso irá:

1. Construir a imagem Docker da aplicação Go.
2. Iniciar o container do PostgreSQL.
3. Executar as migrações para criar a tabela `pets`.
4. Iniciar a aplicação Go na porta `8080`.

A aplicação estará rodando na porta `8080`. Você pode acessar a aplicação através do endereço `http://localhost:8080`.

### 4. Testar a Aplicação

Você pode testar a aplicação utilizando ferramentas como o [Postman](https://www.postman.com/) ou [insomnia](https://insomnia.rest/download). Abaixo estão alguns exemplos de requisições que você pode fazer:

#### Criar um novo pet (POST)

```bash
curl -X POST http://localhost:8080/pets \
-H "Content-Type: application/json" \
-d '{"name": "Rex", "species": "Dog", "age": 3}'
```

#### Listar todos os pets (GET)

```bash
curl -X GET http://localhost:8080/pets
```

#### Obter um pet pelo ID (GET)

```bash
curl -X GET http://localhost:8080/pets/1
```

#### Atualizar um pet (PUT)

```bash
curl -X PUT http://localhost:8080/pets/1 \
-H "Content-Type: application/json" \
-d '{"name": "Rex", "species": "Dog", "age": 4}'
```

#### Deletar um pet (DELETE)

```bash
curl -X DELETE http://localhost:8080/pets/1
```

## Estrutura do Projeto

A estrutura do projeto está organizada da seguinte forma:

```
go-crud-pet-api/
├── backend/
│   ├── Dockerfile          # Dockerfile para construir a imagem da aplicação
│   ├── go.mod             # Arquivo de dependências do Go Modules
│   ├── go.sum             # Arquivo de checksum do Go Modules
│   ├── main.go            # Ponto de entrada da aplicação
│   ├── handlers/
│   │   └── pet_handler.go # Handlers para as rotas da API
│   ├── models/
│   │   └── pet.go         # Modelo de dados para Pet
│   ├── repositories/
│   │   └── pet_repository.go # Lógica de acesso ao banco de dados
│   └── migrations/
│       └── 001_create_pets_table.up.sql # Migração para criar a tabela `pets`
├── docker-compose.yml      # Arquivo de configuração do Docker Compose
├── .env                    # Variáveis de ambiente
├── .gitignore              # Arquivo para ignorar arquivos no Git
└── README.md               # Documentação do projeto
```

## Contribuindo

Se você deseja contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar um pull request. Todas as contribuições são bem-vindas!

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Se você tiver alguma dúvida ou sugestão, entre em contato através do [GitHub Issues](https://github.com/Eduardo-Rosa/go-crud-pet-api/issues) ou linkedin [Eduardo Rosa](https://www.linkedin.com/in/eduardobetimrosa/).
