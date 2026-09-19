# db-scale-journal — Spec

## Objetivo

Projeto pessoal de estudo prático de performance de banco de dados. Construir uma aplicação em Go, sobre MySQL, simulando o domínio de uma rede social simplificada (posts, seguidores, feed), e evoluí-la deliberadamente a partir de uma versão ingênua — usando testes de carga com milhões de requisições sobre milhões de linhas para descobrir gargalos reais e entender, na prática, como resolvê-los.

## Escopo do domínio

Entidades esperadas (o desenho exato do schema — colunas, tipos, relações, índices — é decisão de quem implementa, não está fixado aqui):

- Usuários
- Posts
- Relação de "seguir" (followers/following)
- Feed / timeline — provavelmente o ponto mais interessante do ponto de vista de performance, já que envolve agregar dados de múltiplas fontes

## Stack

- Linguagem: Go
- Banco: MySQL (ponto de partida — pode ser comparado ou trocado mais adiante)
- Geração de carga: a definir

## Não-objetivos (por agora)

- Não é objetivo ter a versão "certa"/otimizada desde o início.
- Não é objetivo cobrir autenticação, autorização, UI — o foco é API + banco + carga.
- Não é objetivo, nesta fase, cobrir sharding, multi-região ou cache distribuído — isso pode surgir como evolução futura, não como ponto de partida.

## Princípio de trabalho

Começar simples e ingênuo de propósito (queries diretas, sem índices pensados, sem pool de conexões ajustado, possíveis N+1 etc.), rodar carga, observar o que quebra ou degrada, entender a causa raiz, ajustar uma coisa por vez, medir de novo. O registro dessas iterações é parte do valor do projeto — daí o nome "journal".

## Critérios de sucesso

Não é "chegar no mais rápido possível". É:

- Conseguir simular volume real (milhões de linhas / milhões de requisições).
- Para cada gargalo encontrado, conseguir explicar por que ele existe e por que a correção aplicada resolveu (ou não resolveu).
- Ter um histórico registrado da evolução, com métricas de antes/depois de cada mudança.

## Papel do Claude neste projeto

Orientação: sugestão de tópicos de estudo, revisão crítica do que for escrito/decidido, ajuda com boilerplate e tooling (ex: libs de geração de carga, scaffolding, seed de dados em massa). Não entrega soluções prontas para as decisões centrais de performance (schema, índices, queries, arquitetura) — essas são para quem está estudando decidir e testar.
