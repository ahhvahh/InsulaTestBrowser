<!-- A-BIOMA-DOC-STANDARD -->
# PERS-001 — Modelo de persistência SQLite

**Situação:** Proposto. O esquema abaixo é um **modelo lógico inicial**, não uma migração existente.

## Objetivo

Persistir o estado operacional mínimo para gerenciar contêineres, rastrear imagens, aplicar retenção e reconciliar sessões após falha do serviço. O SQLite permanece no host, fora do alcance dos contêineres.

## Entidades propostas

### `images`

| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária local, identificador opaco |
| `digest` | TEXT | Digest OCI imutável, único |
| `chromium_version` | TEXT | Versão/revisão efetivamente incluída |
| `source` | TEXT | Origem de construção, sem credenciais |
| `created_at` | TEXT | Instante UTC em formato ISO 8601 |
| `verified_at` | TEXT NULL | Última validação de segurança concluída |

### `sessions`

| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária opaca |
| `image_id` | TEXT | Referência a `images.id` |
| `container_id` | TEXT NULL | ID real retornado pelo Podman, único quando presente |
| `status` | TEXT | Estado da sessão; ver [fluxo](../fluxos/FLUX-001-ciclo-de-vida.md) |
| `created_at` | TEXT | UTC, obrigatório |
| `started_at` | TEXT NULL | UTC |
| `ended_at` | TEXT NULL | UTC |
| `purge_after` | TEXT NULL | Prazo UTC de expurgo dos metadados da sessão |
| `failure_code` | TEXT NULL | Código técnico de erro, sem segredos |

### `artifacts`

| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária |
| `session_id` | TEXT | Referência a `sessions.id` |
| `kind` | TEXT | `screenshot`, `download`, `log` |
| `storage_key` | TEXT | Identificador relativo ao armazenamento privado, nunca caminho arbitrário |
| `created_at` | TEXT | UTC |
| `purge_after` | TEXT | Instante de expurgo do arquivo |

O schema definitivo, as restrições `CHECK`, índices, regras de exclusão por FK e migrações ficam para implementação e testes. Recomenda-se índices em `sessions(status)`, `sessions(purge_after)` e `artifacts(purge_after)`.

## Regras de consistência

- Habilitar `PRAGMA foreign_keys=ON` em cada conexão e configurar `busy_timeout`.
- Avaliar `journal_mode=WAL` no filesystem local: múltiplos leitores podem coexistir, mas SQLite mantém **um escritor por vez**.
- Usar transações curtas e idempotência em operações de criação/encerramento.
- Após reinício, consultar Podman e reconciliar `container_id` e estados. O SQLite não substitui a inspeção real do runtime.
- Não registrar cookies, tokens, senhas, HTML de páginas nem conteúdo de downloads no banco.
- Na retenção, remover primeiro dados sensíveis e artefatos físicos, depois os registros relacionados, sem perder rastreabilidade de falhas de exclusão.

## Retenção proposta

| Classe | Valor inicial proposto | Observação |
| --- | --- | --- |
| Perfil, cookies e cache | 0m após encerramento | Remover imediatamente |
| Screenshots | 1h | Parametrizável |
| Downloads | 1h | Parametrizável |
| Logs | 1h | Parametrizável; sem dados sensíveis |
| Metadados de sessão encerrada | 24h | Parametrizável |

Os valores são **padrões candidatos**, não configurações já implementadas. `0m` significa exclusão imediata após o encerramento. A limpeza também será acionada na inicialização do serviço e periodicamente.

## Referências

- [Ciclo de vida](../fluxos/FLUX-001-ciclo-de-vida.md)
- [Decisão SQLite](../decisoes/DEC-002-sqlite.md)
- [SQLite WAL](https://www.sqlite.org/wal.html)
- [SQLite — isolamento](https://www.sqlite.org/isolation.html)
