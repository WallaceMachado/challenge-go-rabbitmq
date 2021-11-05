# challenge-go-rabbitmq

> status:	🚧 api-bank-transfers 🚀 em construção..  🚧


API Restful de controle de pessoas: 
* Implementa um CRUD (create, read,update,delete) básico
* Os registros são salvos em um banco de dados mongoDB. 
* Cada operação executada na API gera uma mensagem de LOG informativa,
  que é publicada em uma FILA RabbitMQ.
* Cada registro de uma pessoa contém as informações de NOME, SEXO, PESO, ALTURA e IMC.
    * Nome (Name):
        * Tipo: string
        * Regras: deve possuir entre 3 e 100 caracteres
    * Sexo (Gender):
        * Tipo: string
        * Regras: deve possuir o valor de "male" ou "female"
    * Peso (Weight):
        * Tipo: float32
        * Regras: deve ser informado em kg, sem a unidade e somente com valor entre 10 e 900
    * Altura (Height):
        * Tipo: float32
        * Regras: deve ser informado em cm, sem a unidade e somente com valor entre 30 e 250
    * IMC:
        * Tipo: float32
        * Regras: deve possuir valor entre 10 e 100
          
 

Um segundo serviço, consome os eventos gerados em fila, e atualiza
um websocket para exibição dos mesmos. Link do repositório: https://github.com/WallaceMachado/websocket-consumer-rabbitmq




## Indice

* <p><a href="#pré-requisitos">Pré Requisitos</a> </p>
* <p><a href="#iniciando-projeto">Iniciando Projeto</a></p>
* <p><a href="#variáveis-de-ambiente">Variáveis de Ambiente</a></p>
* <p><a href="#rotas">Rotas</a></p>
* <p><a href="#controle-de-versão">Controle de versão</a></p>
* <p><a href="#testes">Testes</a></p>
* <p><a href="#autor">Autor</a></p>




## Pré Requisitos

Antes de começar, você precisará ter as seguintes ferramentas instaladas em sua máquina:
* [Git](https://git-scm.com)

Para rodar via docker
* [Docker](https://docs.docker.com/)

Para rodar Local
* [Go](https://golang.org/) versão 1.16.7
* [MongoDB](https://docs.mongodb.com/)
* [RabbitMQ](https://www.rabbitmq.com/)

Além disso, é bom ter um editor para trabalhar com o código como: [VSCode](https://code.visualstudio.com/)



## Iniciando Projeto 

### Local

Deverá criar/rodar banco de dados e rabbitmq com as configurações informadas no arquivo ``` .env ``` 
Altere as varáveis de ambiente DB_HOST e RABBITMQ_DEFAULT_HOST  para localhost
 

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/challenge-go-rabbitmq.git

# Acesse a pasta do projeto no terminal / cmd
$ cd challenge-go-rabbitmq

# Instale as dependências e rode o projeto
$ go run main.go

# Server is running
```
#### Configuração RabbitMQ
* Acessar via navegador http://localhost:15672, 
* Efetuar login com usuário e senha do arquivo ``` .env ```
* Criar nova exchange com o mesmo nome da variável de ambiente RABBITMQ_EXCHANGE_PERSON
* Criar nova queue com o mesmo nome da variável de ambiente RABBITMQ_QUEUE_PERSON
* Fazer o bind da queue com a exchange

### Docker

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/challenge-go-rabbitmq.git

# Acesse a pasta do projeto no terminal / cmd
$ cdchallenge-go-rabbitmq

# Instale as dependências e rode o projeto
$ docker-compose up -d --build

```
#### Configuração RabbitMQ
* Acessar via navegador http://localhost:15672, 
* Efetuar login com usuário e senha do arquivo ``` .env ```
* Criar nova exchange com o mesmo nome da variável de ambiente RABBITMQ_EXCHANGE_PERSON
* Criar nova queue com o mesmo nome da variável de ambiente RABBITMQ_QUEUE_PERSON
* Fazer o bind da queue com a exchange

## Variáveis de Ambiente

Após clonar o repositório, renomeie o ``` .env.example ``` no diretório raiz para ``` .env ``` e atualize com suas configurações.


| Chave  |  Descrição  | Predefinição  |
| :---: | :---: | :---: | 
|  PORT |  Número da porta em que o aplicativo será executado. | 5000  |
|  DB_PORT |  Porta Mongo.  |  27017  |
|  DB_USERNAME |  Usuário Mongo. |  -  |
|  DB_PASSWORD |  Senha do Mongo.  |  -   |
|  DB_NAME |  Nome do banco de dados do aplicativo. |  -  |
|  DB_COLLECTION |  Nome da coleção do banco de dados Mongo. |  -  |
|  DB_HOST |  host do Mongo. |  mongo  |
|  RABBITMQ_DEFAULT_USER | Usuário RabbitMQ.  |  -   |
|  RABBITMQ_DEFAULT_PASS | Senha RabbitMQ.  |  -   |
|  RABBITMQ_DEFAULT_VHOST | virtual host do RabbitMQ.  |  /   |
|  RABBITMQ_DEFAULT_HOST | host do RabbitMQ.  |  rabbit   |
|  RABBITMQ_DEFAULT_PORT | Porta em que o RabbitMQ está sendo excecutado  |  5672   |
|  RABBITMQ_EXCHANGE_PERSON | Exchanche do RabbitMQ  |  person_ex   |
|  RABBITMQ_QUEUE_PERSON  | Fila do RabbitMQ  |  person_queue   |

## Controle de versão
Para contrele de versão, foi inserida a versão ``` v1 ``` após o  ``` host ```

```
GET http://localhost:5000/api/v1/person

```

 
## Rotas

```
http://localhost:5000/api/v1/

```

| Rotas  |  HTTP Method  | Params  |  Descrição  | 
| :---: | :---: | :---: | :---: |
|  /person |  POST |  Body: ``` name ```, ``` gender ```, ``` weigth ```, ``` height ``` e ``` imc ``` |  Cadastre uma nova pessoa |
|  /person |  GET |  -  | Recupere uma lista com todas as pessoas cadastradas |
|  /person/:id |  GET |  Params: ``` id ``` |  Consulte o cadastro de uma pessoa |
|  /person/:id |  PUT |   Params: ``` id ``` Body: ``` name ```, ``` gender ```, ``` weigth ```, ``` height ``` e ``` imc ``` |  Edite o cadastro de uma pessoa |
|  /person/:id |  DELETE |  Params: ``` id ``` |  Exclua o cadastro de uma pessoa|



### Requisições
* ``` POST /person ```

Corpo da requisição:
  
```
{
    "name":  "teste",
    "gender": "male",
    "weight": 80,
    "height": 174,
    "imc":    20
}

```

* ``` PUT /person ```

Corpo da requisição:
  
```
{
    "name":  "teste atualizado",
    "gender": "male",
    "weight": 80,
    "height": 174,
    "imc":    20
}

```
Parâmetro da requisição: 

```  http://localhost:5000/api/v1/person/4b808c0b-5822-410d-a67d-d87241f03e9d ```



## Testes
Para executar os testes :

### Local

Altere as varáveis de ambiente DB_HOST e RABBITMQ_DEFAULT_HOST  para localhost

```bash

  # Acesse a pasta do projeto no terminal / cmd
  $ cd challenge-go-rabbitmq
  
  # Rode os testes
  $ go test ./...
  
```


### Docker


```bash

    # Acesse a pasta do projeto no terminal / cmd
    $ cd challenge-go-rabbitmq

    # Instale as dependências e rode o projeto
    $ docker-compose up --build

    # Pegar o CONTNAINER ID  da imagem github.com/wallacemachado/challenge-go-rabbitmq
    $ docker ps
    # Exemplo de retorno:
    # CONTAINER ID      IMAGE                                             COMMAND                  
    # xxxxxxxxxxxxxxx   github.com/wallacemachado/challenge-go-rabbitmq      "go run main.go"         0.0.0.0:5003->5000/tcp, :::5003->5003 ....  
    # yyyyyyyyyyyyyyy   github.com/wallacemachado/challenge-go-rabbitmq   "docker-entrypoint.s…"   0.0.0.0:5432->5432/tcp, :::5432->5432 ....

    # acesse o contaniner
    $ docker exec -it <CONTAINER ID> /bin/bash
    # Exemplo:
    # docker exec -it xxxxxxxxxxxxxxx /bin/bash

    #root@xxxxxxxxxxxxxxx:/go/src/app# 
    #rode os testes
    $ go test ./...

```



## Autor


Feito com ❤️ por [Wallace Machado](https://github.com/WallaceMachado) 🚀🏽 Entre em contato!

[<img src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/in/wallace-machado-b2054246/)
