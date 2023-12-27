# Manifesto do Desafio Golang

Esse repositório contém o manifesto do Desafio Golang para os candidatos à vaga de Dev Backend Pleno. Você deve fazer um fork desse repositório e desenvolver sua própria implementação do Desafio. Boa sorte.

## 1) Introdução

A Flex Consulta existe para inovar o setor logístico do agronegócio conectando quem conecta o Brasil de ponta a ponta. Nosso compromisso é de desenvolver soluções tecnológicas entregando segurança e agilidade na operação, reduzindo riscos e gastos desnecessários para as transportadoras de cargas. Visite nosso website em flexconsulta.com.br.

O Propósito desse Desafio é verificar sua capacidade de seguir um manifesto com instruções básicas sobre demandas corriqueiras em nosso dia a dia de trabalho. Somos constantemente confrontados com exigências novas e desafiadoras, e precisamos de um profissional que gosta de desafios e seja resoluto, sempre entregando o melhor possível. 

## 2) Objetivo do Desafio

Esse Desafio Golang, como o próprio nome já deixa evidente, é verificar sua capacidade de desenvolver em Golang, além de algumas outras capacidades que alvejamos, a saber:

- Habilidades de programação;
- Capacidade de resolver problemas; 
- Familiaridade com tecnologias específicas;
- Proficiência com versionamento de código com git;
- Conhecimento e prática com mensageiria;

## 3) Detalhes do Desafio

Para concluir esse Desafio você deverá entregar um sistema distribuído de envio de dados, aonde uma aplicação atuará como server destinatário dos dados, e outra aplicação atuará como um cliente emissor dos dados. O cliente deverá enviar ao server um pacote de dados a cada 5 segundos. A mensagem do pacote deve ter a seguinte estrutura:

```
{
	"id": "dce4cc68-848f-11ee-8542-f3104f1fd201",
	"cliente": "e87a9f80-848f-11ee-839a-2394cef1c867",
	"date": "2023-12-27 08:28:07",
	"payload": ""
}
```
Essa mensagem simples possui apenas quatro índices, sendo:

* id: Deve ser um UUID único para cada mensagem enviada;
* cliente: Deve ser o UUID que você desiginou para o cliente;
* date: Deve o time stamp no momento do envio da mensagem;
* payload: Deve conter os dados enviados pelo cliente;

Obs: O payload você pode enviar vazio, visto se tratar de uma simulação.

Você deverá compilar os binários executáveis das suas aplicações "server" e "cliente" para ambos os sistemas operacionais Linux (Ubuntu) e Windows. Nós testaremos seus executáveis para verificar se eles atendem às demandas desse manifesto. 

Basicamente, o que queremos, é rodar o server em um terminal, rodar o cliente em outro terminal, e ver o server receber as mensagens enviadas pelo cliente. Logo, o server deve escrever as mensagens recebidas no terminal.

## 4) Requisitos Técnicos

* Linguagens e Frameworks: A única linguagem utilizada deve ser o Golang. Gostaríamos que você desenvolvesse sem o uso de frameworks, utilizando apenas as livrarias padrão do Golang ou as de terceiros. 

* Configuração do ambiente: Como você desenvolverá as aplicações para uso no terminal do Linux ou Windows, assim como essas aplicações serão binários executáveis que rodam "stand alone", além de serem muito simples, não é necessário configuração alguma de ambiente. Obs: Em uma solução no mundo real, certamente haveria de se lidar com variáveis de ambiente específicas para cada aplicação.

* Padrões de codificação: Gostamos de código limpo e organizado, entendendo que escrevemos código para que seja facilmente lido e gerenciado por outros devs. Gostaríamos de ver seu código Golang bem organizado e seguindo os melhores padrões para codificação em Golang. 

* Message Broker: Assuma que suas aplicações terão uma instância do RabbitMQ rodando em localhost:5672

## Entregáveis Esperados

* Quando você achar que já pode entregar a sua solução, você deverá transferir a propriedade do seu repositório para a Flex Consulta. Transfira para o usuário "flexconsulta".

* Esperamos os executáveis no diretório "bin" do seu repositório.

* Não é necessário escrever a documentação do projeto, mas você impressionaria a gente se descrevesse os detalhes mais importantes do seu projeto no README.md na raiz do seu repositório.

* Não é necessário escrever testes para sua impementação, mas você impressionaria a gente se disponibizasse alguns testes.


