# Manifesto do Desafio Golang

Esse repositório contém o manifesto do Desafio Golang para os candidatos à vaga de Dev Backend Pleno. Você deve dar um fork desse repositório e desenvolver sua própria implementação do Desafio. Boa sorte.

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


