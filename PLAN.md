# db-scale-journal — Plan

Roadmap por fases. As fases são sequenciais no início, mas a partir da Fase 5 o plano vira um ciclo — ele vai crescer conforme os gargalos forem aparecendo, então não há uma "fase final" fixa.

## Fase 0 — Setup

- Inicializar módulo Go e estrutura mínima do projeto.
- Subir MySQL local (docker compose é uma opção razoável).
- Desenhar o schema inicial (users, posts, follows).
- Escrever um seed inicial pequeno (milhares de linhas) só para validar que tudo funciona antes de ir para milhões.

## Fase 1 — API ingênua (v0)

- Implementar rotas básicas: criar usuário, criar post, seguir/deixar de seguir, buscar feed.
- De propósito, sem otimizar nada ainda: queries diretas, sem cache, sem pensar em índice.
- Critério de saída da fase: a API funciona corretamente com poucos dados — correção antes de performance.

## Fase 2 — Escala dos dados

- Seed de milhões de linhas (usuários, posts, relações de seguir) com distribuição realista — não uniforme (ex: alguns usuários com muitíssimos seguidores, a maioria com poucos).
- Estudar: geração de dados sintéticos em massa, distribuições tipo power-law/Zipfian (comuns em redes sociais), inserção em lote vs. individual.

## Fase 3 — Geração de carga e baseline

- Escolher/montar a ferramenta de geração de carga.
- Definir cenários de carga (leitura pesada tipo feed, escrita pesada tipo post, misto).
- Rodar baseline com a API ainda ingênua e registrar métricas: latência p50/p95/p99, throughput, taxa de erro.

## Fase 4 — Observabilidade mínima

- Antes de sair otimizando às cegas, conseguir ver o que está acontecendo: métricas da aplicação, slow query log do MySQL, `EXPLAIN` de queries.
- Estudar: o que observar num banco sob carga, como ler um plano de execução, monitoring básico.

## Fase 5+ — Ciclo iterativo de otimização

Repetir:

1. Rodar carga.
2. Identificar o gargalo (o que está lento/saturado e por quê).
3. Formular uma hipótese.
4. Estudar e aplicar uma mudança pontual.
5. Medir de novo.
6. Registrar a iteração.

Temas que tendem a aparecer organicamente nesse ciclo (não implementar antecipadamente — deixar o gargalo real apontar pra eles): indexação, connection pooling, paginação do feed, N+1, desnormalização, fan-out on write vs. fan-out on read, cache, réplicas de leitura, particionamento.

## Registro das iterações

Sugestão: uma pasta `journal/` com um arquivo por iteração — data, o que foi testado, métricas antes/depois, conclusão. Isso é o que dá valor de estudo ao projeto ao longo do tempo, além do código em si.
