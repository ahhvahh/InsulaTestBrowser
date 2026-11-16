<!-- A-BIOMA-DOC-STANDARD -->
# CONTR-004 — Catálogo temporário de artefatos e resultado de sessão

**Situação:** Proposto; sem implementação no serviço Go. **Versão do resultado:** `result_version: 1`. A API externa continua planejada para UDS e MessagePack.

## Objetivo
Capturar os passos executados pelo Starlark e disponibilizar, **durante a execução e após o processamento**, os PDFs, downloads, screenshots, HTML bruto e texto extraído. O consumidor identifica arquivos por `artifact_id` e a navegação por `session_id`; nunca precisa conhecer caminhos físicos.

## Organização no host
Raiz inicial: `/tmp/itb`; configurável em `artifacts.root_dir`. Cada sessão possui diretório exclusivo:

```text
/tmp/itb/
  {session_id}/
    downloads/
    screenshots/
    pages/
      html/
      text/
    events.jsonl
    result.json
```

- Criar o diretório na abertura da sessão; `session_id` é opaco, gerado pelo serviço e validado como componente único de caminho (sem `/`, `..` ou caracteres de controle).
- Arquivos usam `artifact_id` no nome físico e extensão compatível com o tipo validado. O nome fornecido pelo site é só metadado sanitizado.
- `events.jsonl` recebe eventos de maneira incremental; `result.json` é uma projeção consolidada, regravada por troca atômica de arquivo temporário (sem edição parcial visível).
- `SQLite` é o catálogo consultável e a fonte operacional de metadados; o JSON é uma representação para consumo/diagnóstico. Reconciliar após quedas usando eventos e metadados confirmados.
- O encerramento do Chromium **não** exclui imediatamente os artefatos. Expiração e remoção são controladas por `purge_after`/política configurada, inclusive depois de reinicialização.
- `/tmp` é volátil e pode ser limpo pelo sistema/reinício: retenção é **melhor esforço**, não armazenamento permanente. Quem precisar preservar os bytes deve copiá-los antes da expiração.

### Acesso
O termo “público” significa **disponível para clientes autorizados do InsulaTestBrowser**, não `chmod 777`, servidor HTTP aberto ou leitura por qualquer usuário do host. Usar diretórios `0700`, arquivos `0600` e usuário de serviço dedicado; se outro processo precisar de acesso direto, definir grupo/ACL explicitamente. Validar propriedade, caminhos, symlinks, hardlinks e limites de tamanho antes de aceitar operações. Nunca servir `/tmp/itb` como diretório estático público. A API verifica posse/autorização por artefato e sessão.

## Esquema de `result.json`
Um resultado consolidado por sessão, inclusive enquanto `status=running`. Na versão inicial, **uma execução de script utiliza uma sessão principal**; se for habilitada múltipla sessão por execução, cada `result.json` permanece por sessão e `get_script_results` agrega por `execution_id`.

```json
{
  "result_version": 1,
  "execution_id": "exec_exemplo",
  "session_id": "sess_exemplo",
  "status": "running",
  "revision": 5,
  "updated_at": "2026-11-16T14:00:00Z",
  "steps": [
    {
      "step_id": "step_001",
      "sequence": 1,
      "operation": "browser.navigate",
      "status": "completed",
      "tab_id": "main",
      "url": "https://example.org/atos",
      "started_at": "2026-11-16T14:00:00Z",
      "finished_at": "2026-11-16T14:00:01Z",
      "artifact_ids": ["art_html001", "art_text001"]
    },
    {
      "step_id": "step_002",
      "sequence": 2,
      "operation": "browser.download",
      "status": "completed",
      "tab_id": "main",
      "started_at": "2026-11-16T14:00:02Z",
      "finished_at": "2026-11-16T14:00:03Z",
      "artifact_ids": ["art_pdf001"]
    }
  ],
  "data": {
    "artifacts": [
      {
        "artifact_id": "art_html001",
        "step_id": "step_001",
        "kind": "html",
        "storage_key": "pages/html/art_html001.html",
        "mime_type": "text/html",
        "size_bytes": 1200,
        "sha256": "<sha256_verificado>",
        "status": "ready"
      },
      {
        "artifact_id": "art_text001",
        "step_id": "step_001",
        "kind": "text",
        "storage_key": "pages/text/art_text001.txt",
        "mime_type": "text/plain; charset=utf-8",
        "size_bytes": 350,
        "sha256": "<sha256_verificado>",
        "status": "ready"
      },
      {
        "artifact_id": "art_pdf001",
        "step_id": "step_002",
        "kind": "download",
        "storage_key": "downloads/art_pdf001.pdf",
        "mime_type": "application/pdf",
        "size_bytes": 24000,
        "sha256": "<sha256_verificado>",
        "status": "ready"
      }
    ],
    "results": []
  }
}
```

Exemplo **ilustrativo**, sem arquivos efetivamente capturados. `data.results` recebe os dados estruturados publicados por `job.emit_result`; `data.artifacts` contém os arquivos físicos. O `artifact_id` é a chave canônica; `storage_key` é relativo à pasta de sessão e informativo, **não** um argumento de leitura/escrita.

### Catálogo de eventos e steps
- Uma chamada Starlark relevante gera um `step` registrado **pelo serviço Go**, não manualmente pelo script: criação de sessão, navegação, abertura/fechamento de aba, inspeção de listas/formulários, seleção de campos, clique, paginação, screenshot, captura de página e download.
- `sequence` é monotônico por sessão. Cada step possui `step_id`, `operation`, `status` (`started|completed|failed|cancelled`), horários, `tab_id` opcional, referência de destino sanitizada, `artifact_ids` e código de erro opcional. Uma transição de status não cria uma ação duplicada.
- Nunca armazenar valores de campos de formulário, cookies, tokens, cabeçalhos de autenticação ou URLs com segredos. HTML bruto e capturas são **conteúdo potencialmente sensível**; restringir leitura, aplicar retenção e evitar incluir seus bytes em logs/resultado.
- `events.jsonl`: um JSON por evento, por exemplo `{"sequence":3,"event":"step.completed","step_id":"step_002","artifact_ids":["art_pdf001"]}`. Incluir também `artifact.ready`, `result.emitted` e `session.finished`. Persistir apenas após confirmação da operação; eventos de falha registram código sanitizado.
- `revision` permite observar alterações em `result.json`, independentemente de paginação/cursor do cliente.

## Captura de HTML e texto
Além de `get_dom` (árvore limitada), propor `browser.capture_page(session_id, tab_id=None, formats=["html","text"])`, com retorno `{artifacts:[{artifact_id,kind,mime_type}]}`. Essa chamada gera arquivos e steps como screenshot/download. Alternativamente, `artifacts.capture_on_navigation: true` habilita captura automática após navegação bem-sucedida (incluindo navegação por `next_page` e cliques com mudança de documento); a captura automática usa os mesmos limites de bytes, quantidade e retenção e não armazena cookies. **A API não está implementada**.

## Interface externa futura: UDS / MessagePack

| Comando | Entrada | Retorno |
| --- | --- | --- |
| `get_script_status` | `execution_id` | `{execution_id,session_id,status,last_step,artifacts_count}` |
| `get_script_results` | `execution_id,cursor?,limit?` | `{execution_id,session_id,status,steps,data,next_cursor}` — apenas itens posteriores ao cursor, ou resultado completo na ausência dele, sujeito a paginação |
| `list_artifacts` | `session_id,cursor?,limit?,kind?` | `{artifacts,next_cursor}` |
| `get_artifact_info` | `artifact_id` | Metadados autorizados, `size_bytes`, `sha256`, `purge_after` |
| `read_artifact` | `artifact_id,offset=0,limit?` | `{artifact_id,offset,bytes,eof}`; `bytes` é **binário MessagePack**, com tamanho máximo de chunk |
 
`run_script` devolve `execution_id` de imediato; o `session_id` pode ainda ser `null` até que o Starlark chame `browser.create_session`. Consultas seguintes retornam o ID assim que existir. Após `end_session`, leitura de artefatos ainda funciona até a expiração. A ausência de arquivo, expiração, corrupção, erro de checksum e falta de permissão são erros distintos; `artifact_id` não concede permissão.

## Critérios de aceite
1. Isolamento de duas sessões: um cliente não acessa ou lista artefatos de outra sem autorização explícita.
2. `result.json` sempre é um JSON íntegro durante atualização; `events.jsonl` conserva ordem por sessão e suporta reconstrução.
3. Uma navegação, um clique e um download surgem como steps com `artifact_id` quando houver arquivo; falhas também ficam registradas.
4. HTML bruto, texto e screenshots são capturados com limites, checksum e metadados; não salvar HTML/valores confidenciais em SQLite ou logs.
5. Consulta por cursor entrega somente novos registros, sem duplicar steps; funciona antes e depois do fim.
6. Downloads e consultas por `artifact_id` funcionam após `end_session`, mas não após o expurgo.
7. Falhas, cancelamento e reinício do serviço não deixam diretórios ilimitados ou arquivos parcialmente publicados.
8. Testes cobrem permissões `0700/0600`, path traversal, symlink, limites, expiração, checksum e download parcial.

## Referências
- [ARQ-004 — Motor Starlark](../arquitetura/ARQ-004-motor-starlark.md)
- [CONTR-002 — API Starlark](CONTR-002-api-starlark.md)
- [CONTR-003 — Administração](CONTR-003-administracao-scripts.md)
- [PERS-001 — Persistência](../persistencia/PERS-001-sessoes-sqlite.md)
- [ARQ-002 — Segurança](../arquitetura/ARQ-002-seguranca.md)
