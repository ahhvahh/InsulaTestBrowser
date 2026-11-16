<!-- A-BIOMA-DOC-LEGADO -->

> Base preservada. [Proposta reorganizada na MIL-0001](../milestones/mil-0001/modulos/persistencia/persistencia/PERS-001-sessoes-sqlite.md); ainda não publicada.

# PERS-001 — Modelo de persistência SQLite

**Situação:** Proposto. O esquema abaixo é um **modelo lógico inicial**, não uma migração existente.

## Objetivo

Persistir o estado operacional mínimo para gerenciar contêineres, rastrear imagens, aplicar retenção e reconciliar sessões após falha do serviço. O SQLite permanece no host, fora do alcance dos contêineres. **O caminho físico do banco é definido por `database.path` no [config.yaml](../referencia/REF-001-configuracao.md).**

## Entidades propostas

### `images`

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

### `session_steps`

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

A lista de arquivos e passos alimenta `/tmp/itb/{session_id}/result.json`; os eventos incrementais também são registrados em `events.jsonl`. O banco mantém somente metadados, não bytes de screenshots, PDFs ou HTML bruto. Consulte [CONTR-004](../contratos/CONTR-004-artefatos-resultados.md).

O schema definitivo, as restrições `CHECK`, índices, regras de exclusão por FK e migrações ficam para implementação e testes. Recomenda-se índices em `sessions(status)`, `sessions(purge_after)`, `artifacts(session_id,status)`, `artifacts(purge_after)` e `session_steps(session_id,sequence)`.

## Regras de consistência

- Habilitar `PRAGMA foreign_keys=ON` em cada conexão e configurar `busy_timeout`.
- Avaliar `journal_mode=WAL` no filesystem local: múltiplos leitores podem coexistir, mas SQLite mantém **um escritor por vez**.
- Usar transações curtas e idempotência em operações de criação/encerramento.
- Após reinício, consultar Podman e reconciliar `container_id` e estados. O SQLite não substitui a inspeção real do runtime.
- Não registrar cookies, tokens, senhas, HTML de páginas nem conteúdo de downloads no banco.
- Na retenção, remover primeiro dados sensíveis e artefatos físicos, depois os registros relacionados, sem perder rastreabilidade de falhas de exclusão. Encerrar Chromium não exclui automaticamente artefatos ainda retidos; `/tmp` é volátil e não oferece garantia de durabilidade.

## Retenção proposta

| Classe | Valor inicial proposto | Observação |
| --- | --- | --- |
| Perfil, cookies e cache | 0m após encerramento | Remover imediatamente |
| Screenshots | 1h | Parametrizável |
| Downloads | 1h | Parametrizável |
| HTML bruto/textos | 1h | Parametrizável; conteúdo potencialmente sensível |
| `result.json`/`events.jsonl` | 1h | Retenção alinhada aos demais artefatos da sessão |
| Logs | 1h | Parametrizável; sem dados sensíveis |
| Metadados de sessão encerrada | 24h | Parametrizável |

Os valores são **padrões candidatos**, não configurações já implementadas. `0m` significa exclusão imediata após o encerramento. A limpeza também será acionada na inicialização do serviço e periodicamente.

## Referências

- [Ciclo de vida](../fluxos/FLUX-001-ciclo-de-vida.md)
- [Decisão SQLite](../decisoes/DEC-002-sqlite.md)
- [Configuração local](../referencia/REF-001-configuracao.md)
- [Fluxo de preparação](../fluxos/FLUX-002-preparacao-imagem.md)
- [SQLite WAL](https://www.sqlite.org/wal.html)
- [SQLite — isolamento](https://www.sqlite.org/isolation.html)
