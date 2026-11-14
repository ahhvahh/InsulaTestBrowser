<!-- A-BIOMA-DOC-STANDARD -->
# CONTR-003 — Administração de scripts e execuções

**Situação:** Proposto; comandos externos ainda não implementados.

## Objetivo e transporte
Operações solicitadas por cliente autorizado ao serviço Go, via socket Unix (UDS). **MessagePack é a serialização proposta**, ainda dependente de validação de compatibilidade; estruturas JSON abaixo são apenas forma legível do payload. Esses comandos **não** são funções chamadas dentro de scripts Starlark.

## Comandos externos (5)

| Comando | Entrada | Ação e retorno |
| --- | --- | --- |
| `validate_script` | `script_id` **ou** `source`, `api_version` | Analisa sintaxe, referências permitidas e política; `{valid, diagnostics, api_version}` |
| `run_script` | `script_id` **ou** `source`, `parameters`, `limits` opcionais | Autoriza, valida e inicia execução; `{execution_id, status}` |
| `get_script_status` | `execution_id` | Retorna `status`, `current_page`, `documents_count`, `last_error` e horários |
| `cancel_script` | `execution_id` | Solicita cancelamento e limpeza; `{execution_id, status}` |
| `get_script_results` | `execution_id`, `cursor` opcional, `limit` opcional | Consulta resultados persistidos; `{items, next_cursor}` |

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
{"execution_id": "exec_123", "status": "queued"}
```

Identificadores e valores são exemplos, não resultados reais.

## Estados da execução
`queued` → `running` → `finished` / `failed` / `cancelled` / `expired`.

Executar em processo auxiliar com limite de RAM/CPU; cada chamada de navegação tem timeout e cancelamento. Limites efetivos são a interseção mais restritiva entre configuração administrativa, perfil, permissão de usuário e pedido. Não aceitar que o script eleve seus próprios limites.

## Persistência e retomada
Persistir, em migração SQLite futura: `script_executions` (`id`, `script_hash`, `api_version`, `status`, `created_at`, `started_at`, `ended_at`, `last_error`, `limits` sem segredo), `script_checkpoints` (`execution_id`, `page`, `metadata`, `updated_at`) e `script_results` (`execution_id`, `ordinal`, `document_key`, `artifact_id`, `metadata`). Tabelas ainda **não existem**. `script_hash` vincula o checkpoint à versão exata do algoritmo e parâmetros relevantes.

Na recuperação: nova sessão Chromium, redescoberta do DOM e reaplicação da navegação até a última página confirmada; nunca presumir que abas ou referências do DOM persistam. Chaves idempotentes impedem rebaixar o mesmo PDF. Retomada de script alterado exige nova execução ou migração explícita.

`cancel_script` é idempotente e deve cancelar navegador, downloads e subprocesso supervisionado, mesmo que o script não consiga executar limpeza.

## Falhas e segurança
Classes: `script_invalid`, `api_incompatible`, `forbidden`, `not_found`, `timeout`, `cancelled`, `resource_exhausted`, `browser_failure`, `runtime_failure`. Logs nunca carregam conteúdo de formulários, tokens ou cookies.

### Aceite
Validar scripts com funções inexistentes, parâmetros errados, loop muito grande, resposta DOM excessiva, site bloqueado, cancelamento durante download, reinício do serviço e retomada com deduplicação.

## Referências
- [API Starlark](CONTR-002-api-starlark.md)
- [Arquitetura de execução](../arquitetura/ARQ-004-motor-starlark.md)
- [Persistência existente](../persistencia/PERS-001-sessoes-sqlite.md)
- [Segurança](../arquitetura/ARQ-002-seguranca.md)
