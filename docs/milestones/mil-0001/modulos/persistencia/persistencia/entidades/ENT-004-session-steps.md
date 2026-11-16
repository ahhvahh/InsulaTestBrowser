<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../../mil-0001.md). Proposta; publicação pendente.

# ENT-004 — session_steps

**Estado:** modelo lógico proposto; não existe migração SQL.

## Campos
| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária do step |
| `session_id` | TEXT | Referência a `sessions.id` |
| `execution_id` | TEXT NULL | Execução associada |
| `sequence` | INTEGER | Ordem monotônica por sessão; `UNIQUE(session_id,sequence)` |
| `operation` | TEXT | Método/ação: `browser.navigate`, `browser.click`, `browser.download` etc. |
| `status` | TEXT | `started`, `completed`, `failed`, `cancelled` |
| `started_at` | TEXT | UTC |
| `finished_at` | TEXT NULL | UTC |
| `metadata` | TEXT NULL | JSON sanitizado sem valores de formulários, tokens ou conteúdo sensível |

## Origem e integridade

Extraído de [PERS-001](../PERS-001-sessoes-sqlite.md). Tipos e nulabilidade foram preservados; CHECK, FKs e regras de exclusão ainda precisam de definição.
