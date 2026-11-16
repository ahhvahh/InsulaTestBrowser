<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md). Proposta; publicação pendente.

# PERS-001 — Modelo de persistência SQLite

**Base documental:** [PERS-001-sessoes-sqlite.md](../../../../../persistencia/PERS-001-sessoes-sqlite.md) no commit `7a13bc974cabeffe89a469b1cf106103f834f02e`. Esta proposta preserva o ID da base.

**Situação:** Proposto. O esquema abaixo é um **modelo lógico inicial**, não uma migração existente.

## Objetivo

Persistir o estado operacional mínimo para gerenciar contêineres, rastrear imagens, aplicar retenção e reconciliar sessões após falha do serviço. O SQLite permanece no host, fora do alcance dos contêineres. **O caminho físico do banco é definido por `database.path` no [config.yaml](../../../referencia/REF-001-configuracao.md).**

## Entidades propostas

### `images`

Campos definidos em [ENT-001](entidades/ENT-001-images.md).

### `sessions`

Campos definidos em [ENT-002](entidades/ENT-002-sessions.md).

### `artifacts`

Campos definidos em [ENT-003](entidades/ENT-003-artifacts.md).

### `session_steps`

Campos definidos em [ENT-004](entidades/ENT-004-session-steps.md).

A lista de arquivos e passos alimenta `/tmp/itb/{session_id}/result.json`; os eventos incrementais também são registrados em `events.jsonl`. O banco mantém somente metadados, não bytes de screenshots, PDFs ou HTML bruto. Consulte [CONTR-004](../../automacao/contratos/CONTR-004-artefatos-resultados.md).

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

- [Ciclo de vida](../../automacao/fluxos/FLUX-001-ciclo-de-vida.md)
- [Decisão SQLite](../decisoes/DEC-002-sqlite.md)
- [Configuração local](../../../referencia/REF-001-configuracao.md)
- [Fluxo de preparação](../../automacao/fluxos/FLUX-002-preparacao-imagem.md)
- [SQLite WAL](https://www.sqlite.org/wal.html)
- [SQLite — isolamento](https://www.sqlite.org/isolation.html)
