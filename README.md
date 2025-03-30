# Desafio Client-Server API

Desafio para a pós graduação Go Expoert da Full Cycle.

## Detalhes

Projeto desenvolvido em Golang usando apenas uma lib externa (driver SQLite).

O projeto implementa tanto um servidor quanto um cliente em go. O servidor faz uma requisição para uma API externa, buscando a cotação atual do dólar, e persiste essa informação em um banco SQLite (`cotacao.db`), retornando o resultado na API.

O Servidor sobe na porta `8080` e expõe a rota `/cotacao` (`GET`, sem parâmetros).

O cliente faz uma requisição para o servidor e persiste o resultado em um arquivo local com nome `cotacao.txt`, onde cada linha é uma requisição feita com sucesso.

As chamadas a recursos bloqueantes possuem timeouts:
- chamada cliente -> servidor tem timeout de 300ms
- chamada servidor -> API tem timeout de 200ms
- chamada servidor -> DB tem timeout de 10ms
- escrita cliente -> arquivo não tem timeout

## Rodando a aplicação

Para rodar a aplicação, primeiro é necessário ter a versão do GoLang 1.18 ou superior.

Além disso, é preciso baixaar as dependências rodando o comando 

```
go mod tidy
```

### Rodando
Após baixar as dependências, precisamos rodar o servidor com o seguinte comando
```
go run server.go
```

Isso irá configurar o banco de dados, criando-o e criando a tabela se ainda não existir, e iniciar o servidor HTTP na porta `8080`.

### Rodando o client

Para rodar o client, primeiro precisamos rodar o servidor, caso contrário o caminho que o client chama não existirá.

Tendo o servidor rodando, rode
```
go run client.go
```

Isso fará uma requisição para o server e salvará a resposta em um arquivo. Qualquer erro será logado. A aplicação finaliza sozinha após executar.

## Bibliotecas
O projeto usa apenas uma biblioteca externa, o driver do SQLite. Todas outras bibliotecas utilizadas são nativas do GoLang 1.18.