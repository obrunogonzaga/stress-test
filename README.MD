# Sistema de Stress Test

Este é um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deve fornecer a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas.

## Funcionalidades

- Realizar requisições HTTP para uma URL especificada.
- Distribuir as requisições de acordo com o nível de concorrência definido.
- Gerar um relatório com informações sobre a execução do teste.

## Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de chamadas simultâneas.

## Relatório

O relatório final contém:

- Tempo total gasto na execução.
- Quantidade total de requisições realizadas.
- Quantidade de requisições com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Como Executar

### Usando Docker

1. Construa a imagem Docker:
    ```sh
    docker build -t loadtester .
    ```

2. Execute a aplicação:
    ```sh
    docker run loadtester load --url=http://google.com --requests=100 --concurrency=10
    ```

### Localmente

1. Instale as dependências:
    ```sh
    go mod download
    ```

2. Compile o projeto:
    ```sh
    go build -o loadtester main.go
    ```

3. Execute o programa:
    ```sh
    ./loadtester load --url=http://google.com --requests=100 --concurrency=10
    ```

## Estrutura do Projeto

```plaintext
stress-test/
│
├── cmd/
│   └── main.go         # Ponto de entrada da aplicação
├── internal/
│   ├── load.go         # Lógica de carregamento
│   └── stresstest.go   # Lógica do teste de stress
├── Dockerfile          # Dockerfile para a aplicação
├── docker-compose.yml  # Arquivo Docker Compose (se necessário)
├── go.mod              # Módulo Go
├── go.sum              # Dependências Go
├── LICENSE             # Licença do projeto
└── README.md           # Documentação do projeto
```

