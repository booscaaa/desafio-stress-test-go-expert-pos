# Aplicação de Teste de Estresse

## Visão Geral
Esta aplicação em Go, intitulada "stress-test", é projetada para realizar testes de carga em serviços web. Especificando uma URL de destino, o número total de solicitações e o grau de concorrência, ela simula tráfego para o serviço, permitindo que você avalie seu desempenho sob estresse.

## Primeiros Passos

### Pré-requisitos
Certifique-se de ter o Go instalado em seu sistema.

### Instalação
Clone ou baixe o repositório para sua máquina local.

```bash
git clone https://github.com/booscaaa/desafio-stress-test-go-expert-pos.git
```

### Navegue até o diretório do projeto

```bash
cd desafio-stress-test-go-expert-pos
```

## Uso
Na pasta dist existe um executavel linux. Acesse a pasta

```bash
cd dist
```

### Comando Básico
O formato básico para executar o teste de estresse é:


```bash
docker run booscaaa/desafio-stress-test-go-expert-pos [flags]
#ou
./stress-test [flags]
```

### Flags
- `-u, --url`: URL do serviço a ser testado (obrigatório).
- `-r, --requests`: Número total de solicitações a serem enviadas (padrão 10).
- `-c, --concurrency`: Número de chamadas simultâneas (padrão 2).

### Exemplo
Para testar um serviço em `http://example.com` com 100 solicitações e um nível de concorrência de 5, use:

```bash
docker run booscaaa/desafio-stress-test-go-expert-pos -u http://example.com -r 100 -c 5
# ou
./stress-test -u http://example.com -r 100 -c 5

# ou também

docker run booscaaa/desafio-stress-test-go-expert-pos --url http://example.com --requests 100 --concurrency 5
# ou
./stress-test --url http://example.com --requests 100 --concurrency 5
```

## Mais detalhes
### Como buildar e rodar do zero localmente

Na pasta raíz do projeto execute:

```bash
go mod tidy
go build -o stress-test main.go
./stress-test [flags]

# ou

docker build -t <nome_imagem> .
docker run <nome_imagem> [flags]

```

## Retornos
Ao executar, casos de sucesso retornarão assim: 
```bash
Test completed
----------------------------------------------------------------------
Total requests: 1
----------------------------------------------------------------------
Successful requests: status code 200; total 1
----------------------------------------------------------------------
Total execution time: 448.047892ms
```
Ao executar, casos de erro retornarão assim: 
```bash
Test completed
----------------------------------------------------------------------
Total requests: 100000
----------------------------------------------------------------------
Requests with error: status code 520; total 24371
----------------------------------------------------------------------
Requests with error: status code 403; total 12799
----------------------------------------------------------------------
Requests with error: status code 0; total 16548
----------------------------------------------------------------------
Requests with error: status code 404; total 46282
----------------------------------------------------------------------
Total execution time: 1m21.21908682s
```

## Erros com status code 0
Quando houver esse tipo de erro significa que há falha na execução da request sem retorno de status code.
```bash
Requests with error: status code 0
```

## Estrutura do Código

### Funcionalidade Principal
- `rootCmd`: Define a interface de linha de comando, analisando flags e executando a lógica principal.
- `makeRequest`: Envia solicitações HTTP GET para a URL especificada e atualiza o relatório.
- `makeReport`: Gera um relatório resumido do teste, incluindo o total de solicitações e a distribuição do status de resposta.
- `addToReport`: Atualiza de forma segura o relatório com novas informações de solicitação, usando mutexes para controle de concorrência.

### Gerenciamento de Concorrência
A aplicação utiliza recursos de concorrência do Go (`sync.WaitGroup` e goroutines) para lidar com solicitações simultâneas de forma eficiente.

### Tratamento de Erros
Implementação adequada de tratamento de erros, particularmente para falhas de solicitação HTTP e cenários de entrada inválida.

### Registro (Logging)
A aplicação registra informações essenciais, como o número total de solicitações, solicitações bem-sucedidas e falhas, e o tempo total de execução.