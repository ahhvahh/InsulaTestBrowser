<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../../mil-0001.md). Proposta; publicação pendente.

# ENT-002 — sessions

**Estado:** modelo lógico proposto; não existe migração SQL.

## Campos
| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária opaca |
| `image_id` | TEXT | Referência a `images.id` |
| `container_id` | TEXT NULL | ID real retornado pelo Podman, único quando presente |
| `status` | TEXT | Estado da sessão; ver [fluxo](../../../automacao/fluxos/FLUX-001-ciclo-de-vida.md) |
| `created_at` | TEXT | UTC, obrigatório |
| `started_at` | TEXT NULL | UTC |
| `ended_at` | TEXT NULL | UTC |
| `purge_after` | TEXT NULL | Prazo UTC de expurgo dos metadados da sessão |
| `failure_code` | TEXT NULL | Código técnico de erro, sem segredos |

## Origem e integridade

Extraído de [PERS-001](../PERS-001-sessoes-sqlite.md). Tipos e nulabilidade foram preservados; CHECK, FKs e regras de exclusão ainda precisam de definição.
