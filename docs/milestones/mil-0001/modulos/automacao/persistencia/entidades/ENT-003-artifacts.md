<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../../mil-0001.md). Proposta; publicação pendente.

# ENT-003 — artifacts

**Estado:** modelo lógico proposto; não existe migração SQL.

## Campos
| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária |
| `session_id` | TEXT | Referência a `sessions.id` |
| `execution_id` | TEXT NULL | Execução Starlark que originou o artefato |
| `step_id` | TEXT NULL | Step que criou o artefato |
| `kind` | TEXT | `screenshot`, `download`, `html`, `text`, `log` |
| `storage_key` | TEXT | Caminho relativo a `/tmp/itb/{session_id}` (não aceitar caminho fornecido pelo cliente) |
| `mime_type` | TEXT | Tipo validado do conteúdo |
| `size_bytes` | INTEGER | Tamanho físico, não negativo |
| `sha256` | TEXT | Hash do arquivo pronto, verificado antes de publicar |
| `status` | TEXT | `writing`, `ready`, `failed`, `expired` |
| `created_at` | TEXT | UTC |
| `purge_after` | TEXT | Instante de expurgo do arquivo |

## Origem e integridade

Extraído de [PERS-001](../PERS-001-sessoes-sqlite.md). Tipos e nulabilidade foram preservados; CHECK, FKs e regras de exclusão ainda precisam de definição.
