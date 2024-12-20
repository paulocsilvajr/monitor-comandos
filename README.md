# Monitor de comandos remotos
### Desenvolvido no Ubuntu 24.04, Golang 1.21.9, Gin Web Framework 1.10

Este repositório contém um API desenvolvida em Golang para executar comandos previamente cadastrados em um servidor remoto e retornar seu resultado. Os comandos são executados via funções do pacote "os/exec" em uma goroutine para evitar ficar esperando pela resposta na requisição. Após finalizar o comando, pode-se obter o resultado via consulta em rota baseada em UUID retornado na solicitação de execução do comando.

### Pré-requisitos
Para compilar o projeto, deve-se instalar a linguagem de programação Go(golang) na versão 1.21.9 ou superior. No Ubuntu 24.04, execute:
```
sudo apt install golang-go
```
Se necessário, adicione ao PATH o local do executável do go, como no exemplo abaixo do arquivo .zshrc e reinicie a sessão:
```
export PATH=/usr/lib/go-1.21/bin/:$PATH
```
Instale as dependências via `go get` na pasta do projeto. Somente será possível instalar dependências e compilar se a versão do Go instalada seja igual ou superior a informada no arquivo 'go.mod'

O script *build.sh* compila o código na pasta bin, arquivo *monitor-comandos*. Pode-se executar também o projeto via comando:
```
go run main.go
```

Na primeira execução do programa é criado o arquivo *comandos* no mesmo diretório do executável da API. Se executado a API via 'go run', o arquivo 'comandos' vai ficar na pasta temporária do SO, portanto para facilitar testes, compile via 'build.sh'. 

Deve-se adicionar no arquivo 'comandos' os comandos que se deseja executar no servidor remoto. A convenção para esse arquivo deve ser a seguinte:
```
nomeRotaSemEspacos comando1 -p1 --parametro2 "parametro em string"
outraRota comando2 -p1
```
A rota informada como primeiro campo de cada linha do arquivo deve ser concisa e sem espaços. Ela será usada na rota para invocar o comando associado. Abaixo exemplo de listagem de palavras de dicionário em inglês presente no sistema, respondendo a rota lsdict:
```
lsdict cat /usr/share/dict/american-english
```
O comando pode conter quantos parâmetros forem necessários, mas se um deles precisar ser um texto com espaços, obrigatóriamente deve-se informar esse parâmetro entre aspas(").
Cada linha do arquivo deve conter somente um comando. Para comandos mais complexos, crie um script.sh com a lógica necessária e informe esse script no arquivo, de preferência com o diretório absoluto do script, caso contrário, ele deve estar na pasta do executável desta API. Lembre-se da permissão de eXecução do script.

Por padrão, a API executa na porta 8080, padrão do [Gin Framework](https://gin-gonic.com/), mas pode-se alterar para a porta desejada ao informar a variável PORT na execução, como no exemplo abaixo com a porta 9000:
```
PORT=9000 ./monitor-comandos  // ou
PORT=9000 go run main.go
```

O script *install_service.sh* adiciona a API da pasta 'bin' como um serviço do Systemd. O serviço que será adicionado terá a permissão do usuário que executar o script citado, portanto, se for necessário executar comandos como root, execute o script 'install_service.sh' como root. Para detalhes de uso consulte a ajuda do script com os parametros -h ou --help.

### Rotas:
- **[GET]/rota-comando-cadastrado-em-arquivo**: Retorna o *status-code* e a rota para consultar o resultado do comando no formato da rota abaixo
- **[GET]/resultados/UUID**: Retorna o *status-code* e o resultado da execução do comando associado. O resultado da execução do comando é separado por *stdout*, *err* e *exit-code*.

Após obter o resultado via consulta na rota associada, ele é removido e não pode ser mais obtido. Deve-se reexecutar o comando na rota associada para obter novamente a resposta.

Se a API for finalizada, resultados que não foram solicitados via rota associada serão perdidos, pois são armazenados em memória.

Para execução de comandos como administrador(root), execute a API como administrador.

### Arquivos

```
bin/: Pasta com o executável da API gerada via script "build.sh".
controller/: Pasta dos controllers cadastrados da API.
helper/: Pasta com funções e estruturas auxiliares, para tarefas diversas, usadas na API.
model/: Pasta de modelos(structs e métodos) da API.
router/: Pasta com o roteador baseado no Gin Web Framework e funções associadas.
view/: Pasta que agrupa as visualizações JSON da API.
.gitignore: arquivo do git para ignorar arquivos desnecessários ou que não devem ser compartilhados.
build.sh: Script para compilar a API na pasta bin.
comandos: Arquivo com a listagem de rotas e comandos associados usados pela API para criar as rotas e definir o comando que será executado. É criado vazio ao executar a API. Deve-se adicionar entradas para a API funcionar.
executa_remoto.sh: Script para testes local de API - consulta a API no endereço "localhost:8080" e formata a resposta JSON via jq.
go_teste_all.sh: Script para executar os testes unitários que foram implementados. Testes em construção constante.
go.mod e go.sum: Arquivos do golang que definem os pacotes usados nesse projeto go. "go get" consulta esse arquivo para instalar as dependências.
install_service.sh: Script para criar e registrar a API como um serviço do systemd. Execute com a opção -h ou --help para ajuda.
license_mit.txt: Licença MIT associada ao projeto.
main.go: Arquivo principal do projeto go.
monitorar_requisicoes.sh: Script para monitorar a API executada como um serviço via journalctl.
README.md: Este arquivo de ajuda.
```

### Licença

[Licença MIT](https://github.com/paulocsilvajr/monitor-comandos/blob/master/license_mit.txt), arquivo em anexo no repositório.

### Contato

Paulo Carvalho da Silva Junior - pauluscave@gmail.com
