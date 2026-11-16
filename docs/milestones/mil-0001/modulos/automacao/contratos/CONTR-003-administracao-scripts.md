<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md). Proposta; publicação pendente.

# CONTR-003 — Administração de scripts e execuções

**Base documental:** [CONTR-003-administracao-scripts.md](../../../../../contratos/CONTR-003-administracao-scripts.md) no commit `7a13bc974cabeffe89a469b1cf106103f834f02e`. Esta proposta preserva o ID da base.

**Situação:** Proposto; comandos externos ainda não implementados.

## Objetivo e transporte
Operações solicitadas por cliente autorizado ao serviço Go, via socket Unix (UDS). **MessagePack é a serialização proposta**, ainda dependente de validação de compatibilidade; estruturas JSON abaixo são apenas forma legível do payload. Esses comandos **não** são funções chamadas dentro de scripts Starlark.

## Comandos externos (8)

| Comando | Entrada | Ação e retorno |
| --- | --- | --- |
| `validate_script` | `script_id` **ou** `source`, `api_version` | Analisa sintaxe, referências permitidas e política; `{valid, diagnostics, api_version}` |
| `run_script` | `script_id` **ou** `source`, `parameters`, `limits` opcionais | Autoriza, valida e inicia execução; `{execution_id, session_id: null, status}`. O ID da sessão fica disponível após `browser.create_session` |
| `get_script_status` | `execution_id` | Retorna `session_id`, `status`, `current_page`, `documents_count`, `last_step`, `artifacts_count`, `last_error` e horários |
| `cancel_script` | `execution_id` | Solicita cancelamento e limpeza; `{execution_id, status}` |
| `get_script_results` | `execution_id`, `cursor` opcional, `limit` opcional | Consulta steps e artefatos produzidos; `{execution_id,session_id,status,steps,data,next_cursor}` | 
| `list_artifacts` | `session_id`, `cursor`/`limit`/`kind` opcionais | Catálogo paginável `{artifacts,next_cursor}` |
| `get_artifact_info` | `artifact_id` | Metadados e prazo de expiração do arquivo, sem conteúdo |
| `read_artifact` | `artifact_id`, `offset=0`, `limit` opcional | Trecho binário MessagePack `{artifact_id,offset,bytes,eof}` sujeito a limite |

`source` contém texto Starlark, limitado em tamanho e autorizado **somente** para clientes com permissão de administrar scripts. `script_id` aponta para script cadastrado pelo mecanismo confiável de instalação/catalogação, ainda a definir. Enviar fonte diretamente **não** confere permissão para executá-la.

## Exemplo conceitual de chamada

```json
{
  "command": "run_script",
  "api_version": 1,
  "script_id": "coleta_decretos_tatui_v1",
  "parameters": {"url": "https://portal.tatui.sp.gov.br/atos-oficiais/legislacao/decretos-municipais", "max_pages": 25},
  "limits": {"max_downloads": 250, "max_duration_seconds": 1800}
}
```

```json
{"execution_id": "exec_123", "session_id": null, "status": "queued"}
```

Identificadores e valores são exemplos, não resultados reais.

## Estados da execução
`queued` → `running` → `finished` / `failed` / `cancelled` / `expired`.

Executar em processo auxiliar com limite de RAM/CPU; cada chamada de navegação tem timeout e cancelamento. Limites efetivos são a interseção mais restritiva entre configuração administrativa, perfil, permissão de usuário e pedido. Não aceitar que o script eleve seus próprios limites.

## Persistência e retomada
Persistir, em migração SQLite futura: `session_steps` e `artifacts` conforme [PERS-001](../../persistencia/persistencia/PERS-001-sessoes-sqlite.md), além de `script_executions` (`id`, `script_hash`, `api_version`, `status`, `created_at`, `started_at`, `ended_at`, `last_error`, `limits` sem segredo), `script_checkpoints` (`execution_id`, `page`, `metadata`, `updated_at`) e `script_results` (`execution_id`, `ordinal`, `document_key`, `artifact_id`, `metadata`). Tabelas ainda **não existem**. `script_hash` vincula o checkpoint à versão exata do algoritmo e parâmetros relevantes.

Os arquivos temporários ficam em `/tmp/itb/{session_id}` e incluem `events.jsonl` e `result.json`. O `result.json` tem `session_id`, `steps` e `data.artifacts`/`data.results`; os resultados da API são incrementais e consultáveis inclusive depois do encerramento da sessão, até a expiração. `artifact_id` é opaco; a leitura deve validar autorização e integridade (checksum), sem acesso direto aos caminhos. [CONTR-004](CONTR-004-artefatos-resultados.md) define o esquema completo.

Na recuperação: nova sessão Chromium, redescoberta do DOM e reaplicação da navegação até a última página confirmada; nunca presumir que abas ou referências do DOM persistam. Chaves idempotentes impedem rebaixar o mesmo PDF. Retomada de script alterado exige nova execução ou migração explícita.

`cancel_script` é idempotente e deve cancelar navegador, downloads e subprocesso supervisionado, mesmo que o script não consiga executar limpeza.

## Falhas e segurança
Classes: `script_invalid`, `api_incompatible`, `forbidden`, `not_found`, `timeout`, `cancelled`, `resource_exhausted`, `browser_failure`, `runtime_failure`. Logs nunca carregam conteúdo de formulários, tokens ou cookies.

### Aceite
Validar scripts com funções inexistentes, parâmetros errados, loop muito grande, resposta DOM excessiva, site bloqueado, cancelamento durante download, reinício do serviço e retomada com deduplicação.

## Referências
- [API Starlark](CONTR-002-api-starlark.md)
- [Arquitetura de execução](../../../arquitetura/ARQ-004-motor-starlark.md)
- [Modelo de persistência proposto](../../persistencia/persistencia/PERS-001-sessoes-sqlite.md)
- [Catálogo e resultado incremental](CONTR-004-artefatos-resultados.md)
- [Segurança](../../../arquitetura/ARQ-002-seguranca.md)

## Pendências da migração comportamental

O conteúdo legado foi absorvido para refinamento; ainda não é uma especificação SDD 2.0 completa. Faltam decompor operações em contratos de entrada/saída, tipar cada campo e falha, vincular regras por fluxo e explicitar tratamento de falhas por chamada. As listas desconhecidas não são representadas como vazias.

<a id="sdd"></a>
## SDD — Spec-Driven Development

**Migração pendente:** formato-alvo 2.0. Não há YAML validado deste comportamento; preservar a descrição acima até fechar as assinaturas. Ver o exemplo tipado [FLUX-0003](../fluxos/FLUX-0003-validar-session-id.md).
