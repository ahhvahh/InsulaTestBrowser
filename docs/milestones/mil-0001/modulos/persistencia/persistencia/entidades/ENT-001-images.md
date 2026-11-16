<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../../mil-0001.md). Proposta; publicação pendente.

# ENT-001 — images

**Estado:** modelo lógico proposto; não existe migração SQL.

## Campos
| Campo | Tipo SQLite | Regra |
| --- | --- | --- |
| `id` | TEXT | Chave primária local, identificador opaco |
| `digest` | TEXT | Digest OCI imutável, único |
| `chromium_version` | TEXT | Versão/revisão efetivamente incluída |
| `source` | TEXT | Origem de construção, sem credenciais |
| `profile_name` | TEXT NULL | Perfil de configuração utilizado para preparar/selecionar a imagem |
| `creation_method` | TEXT NULL | Origem da imagem: construção ou snapshot de preparação; definição final sujeita a migração |
| `created_at` | TEXT | Instante UTC em formato ISO 8601 |
| `verified_at` | TEXT NULL | Última validação de segurança concluída |

## Origem e integridade

Extraído de [PERS-001](../PERS-001-sessoes-sqlite.md). Tipos e nulabilidade foram preservados; CHECK, FKs e regras de exclusão ainda precisam de definição.
