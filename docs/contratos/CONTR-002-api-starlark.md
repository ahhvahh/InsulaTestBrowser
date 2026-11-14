<!-- A-BIOMA-DOC-STANDARD -->
# CONTR-002 — API pública para scripts Starlark

**Situação:** Proposto. **Versão de contrato:** `starlark_api_version: 1` (reservada; não implementada).
**Escopo:** chamadas disponíveis **dentro** dos arquivos `.star`. Para comandos externos, consultar [CONTR-003](CONTR-003-administracao-scripts.md).

## Convenções obrigatórias

- Um método é chamado por `modulo.metodo(parametro = valor)`; parâmetros com `? ` nas tabelas são opcionais, com o padrão indicado.
- Valores: `string`, `int`, `bool`, `dict`, `list` e `None` Starlark. Retornos são `dict`, `list`, `bool` ou `None`, conversíveis para JSON/MessagePack. Exemplo JSON `null` equivale a `None` em Starlark.
- Identificadores `session_id`, `tab_id`, `element_ref`, `list_ref`, `pagination_ref` e `artifact_id` são strings opacas, nunca caminhos ou permissões.
- `session_id` é obrigatório nas chamadas `browser` após criar a sessão; `tab_id` opcional seleciona uma aba, sendo a aba principal o padrão.
- `*_ref` vale somente para a aba e versão do DOM que o originou; após navegação, atualização dinâmica ou paginação, redescobrir os elementos.
- Os tempos `timeout_ms` são inteiros positivos em milissegundos, limitados por configuração. `limit` e `max_depth` são limitados pelo serviço, independentemente do script.
- Falhas de chamada geram erro estruturado na execução, com classe `invalid_request | forbidden | not_found | invalid_state | stale_reference | resource_exhausted | browser_failure | runtime_failure | timeout | cancelled`. A tradução do erro para exceção Starlark ainda será definida; **não presumir `try/except`**.
- `get_*` e `inspect_*` são leituras; `navigate`, `click`, `fill_*`, `submit_form`, `scroll`, `next_page` e `download` têm efeitos observáveis.
- Os métodos retornam **dados**, não strings JSON serializadas: `job.emit_result(data = ...)` recebe dict/list, não `json.dumps`.

## A. `browser` — sessões e navegação (8 métodos)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `create_session(profile=None, viewport=None, timeout_ms=30000)` | Seleciona perfil autorizado, cria contêiner e Chromium isolados | `{session_id, status, default_tab_id}` |
| `get_session(session_id)` | Consulta propriedade, estado, imagem e horários | `{session_id, status, image_digest, created_at}` |
| `end_session(session_id)` | Encerra e agenda limpeza; idempotente | `{session_id, status}` |
| `navigate(session_id, url, tab_id=None, wait_until="load", timeout_ms=30000)` | Valida destino e navega; espera `load`, `domcontentloaded` ou condição suportada | `{url, title, tab_id, document_revision}` |
| `open_tab(session_id, url=None)` | Cria aba adicional; URL opcional sujeita à política | `{tab_id, url}` |
| `close_tab(session_id, tab_id)` | Fecha aba pertencente à sessão | `{tab_id, closed}` |
| `get_page_info(session_id, tab_id=None)` | Consulta URL, título e revisão do DOM | `{url, title, tab_id, document_revision}` |
| `wait_for(session_id, selector, state="visible", tab_id=None, timeout_ms=10000)` | Aguarda seletor CSS em `attached`, `visible`, `hidden` ou `detached` | `{matched, element_ref, state}` |

`viewport`: `{width: int, height: int}` dentro das cotas. `create_session` só pode usar perfis previamente permitidos no `config.yaml`.

## B. `browser` — inspeção de página (9 métodos)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `get_elements(session_id, selector, tab_id=None, limit=100)` | Consulta elementos por seletor CSS | `{elements: [Element], truncated}` |
| `get_element(session_id, element_ref, tab_id=None)` | Retorna detalhes de um elemento identificado | `Element` |
| `get_forms(session_id, tab_id=None, limit=100)` | Identifica `form`, inputs, textarea, select e botões associados | `{forms: [Form], truncated}` |
| `get_buttons(session_id, tab_id=None, limit=100)` | Lista `button`, `input[type=button|submit]` e elementos com papel de botão | `{buttons: [Element], truncated}` |
| `get_links(session_id, selector="a[href]", tab_id=None, limit=100)` | Extrai texto, `href` e URL resolvida | `{links: [Link], truncated}` |
| `inspect_lists(session_id, tab_id=None, limit=20)` | Detecta candidatos a listas/tabelas/grupos repetidos | `{lists: [ListCandidate], truncated}` |
| `get_list_items(session_id, list_ref, tab_id=None, limit=100)` | Extrai linhas/itens da lista selecionada | `{items: [ListItem], truncated}` |
| `inspect_pagination(session_id, tab_id=None)` | Detecta controles de próxima página ou rolagem incremental | `{ref, mode, has_next, current_page}` |
| `get_dom(session_id, selector="html", tab_id=None, max_depth=4, max_nodes=500)` | Retorna DOM simplificado e limitado, não o HTML integral | `{root: Node, truncated, document_revision}` |

`Element` mínimo: `{ref, tag, id, name, type, text, label, attributes, visible, disabled, selector}`. Campos não existentes: `None`. `selector` é sugestão, não garantia de unicidade.

`Form` mínimo: `{ref, id, method, action, fields: [Element], buttons: [Element]}`.

`Link` mínimo: `{ref, text, href, url}`, onde `url` é a URL absoluta normalizada.

`ListCandidate` mínimo: `{ref, selector, item_count, text_sample, confidence}`; `confidence` é heurística, não confirmação. `ListItem` mínimo: `{ref, text, title, url, fields}`, sendo `url` opcional.

`inspect_pagination` pode retornar `{ref: None, mode: "unknown", has_next: False, current_page: None}` caso não haja paginação comprovada; **não inventar páginas**.

## C. `browser` — interação e captura (7 métodos)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `click(session_id, element_ref, tab_id=None, timeout_ms=10000)` | Aciona elemento clicável e aguarda confirmação da ação | `{clicked, navigated, document_revision}` |
| `fill_field(session_id, element_ref, value, tab_id=None)` | Preenche campo (string) com eventos de formulário | `{filled}` |
| `select_option(session_id, element_ref, value, tab_id=None)` | Seleciona opção por valor no `select` | `{selected}` |
| `submit_form(session_id, form_ref, tab_id=None, timeout_ms=30000)` | Submete formulário autorizado | `{submitted, navigated}` |
| `scroll(session_id, direction="down", pixels=700, tab_id=None)` | Executa rolagem `up/down` na aba | `{scrolled, document_revision}` |
| `next_page(session_id, pagination_ref, tab_id=None, timeout_ms=30000)` | Aciona controle validado e aguarda mudança de listagem | `{advanced, url, document_revision}` |
| `screenshot(session_id, tab_id=None, full_page=False, format="png")` | Captura imagem na área de artefatos; `png/jpeg` permitidos | `{artifact_id, mime_type}` |

A aplicação não interpreta diretamente JavaScript arbitrário passado por scripts. Interações são tipadas e validadas pelo adaptador Go.

## D. `browser` — detecção e obtenção de documentos (2 métodos)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `find_downloads(session_id, tab_id=None, file_type="pdf", limit=100)` | Identifica links/controles **candidatos** a download | `{downloads: [{ref, url, text, suspected_type}], truncated}` |
| `download(session_id, url=None, element_ref=None, tab_id=None, timeout_ms=60000)` | Abre/aciona download monitorado pelo Chromium, confere conclusão e registra artefato | `{artifact_id, status, mime_type, size_bytes, sha256}` |

Em `download`, informar **exatamente um** entre `url` e `element_ref`. O serviço valida redirects, tamanho e tipo real do arquivo (PDF com assinatura `%PDF-` e verificações adicionais); nome/extensão/MIME de resposta não bastam. `find_downloads` não garante que um link retornará um PDF; links de visualização, redirecionamentos, blobs e geradores JS exigem tratamento pelo adaptador. Se o navegador não identificar download como concluído, retornar erro, não inventar `artifact_id`.

## E. `artifacts` — consulta de artefatos (1 método)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `get_info(artifact_id)` | Verifica autorização e lê metadados de arquivo armazenado | `{artifact_id, kind, mime_type, size_bytes, sha256, created_at}` |

O script não recebe caminho absoluto de arquivo. Transferência dos bytes ao cliente externo exigirá contrato próprio; **não é função Starlark**.

## F. `job` — execução e checkpoints (6 métodos)

| Método e parâmetros | Ação | Retorno principal |
| --- | --- | --- |
| `checkpoint(page, metadata=None)` | Confirma conclusão de uma página, grava progresso transacional | `{page, saved_at}` |
| `get_checkpoint()` | Lê último checkpoint da execução ou do reinício autorizado | `{page, metadata}` ou `None` |
| `is_processed(key)` | Consulta chave idempotente da execução/coleta | `bool` |
| `record_document(title, source_url, artifact_id, key=None)` | Registra documento e marca chave deduplicável; `key` padrão é `source_url` | `{document_id, recorded}` |
| `emit_result(data)` | Publica dados primitivos estruturados em resultado paginável | `{accepted}` |
| `log(level, message)` | Registra mensagem sanitizada em `debug/info/warn/error` | `None` |

`job` já pertence à execução corrente: **não recebe `execution_id` do script**. A gravação de `record_document` deve concluir antes de `is_processed(key)` retornar verdadeiro; falha de persistência não pode marcar processamento como concluído. `checkpoint(page=N)` só após concluir todos os itens da página N; na retomada o serviço visita a página seguinte, ou reprocessa a atual com deduplicação quando necessário.

## Exemplos mínimos

### Criar sessão e inspecionar formulários

```python
def executar():
    sessao = browser.create_session(profile = "chromium-base")
    sid = sessao["session_id"]
    browser.navigate(session_id = sid, url = "https://www.google.com")
    dados = browser.get_forms(session_id = sid)
    job.emit_result(data = dados)
    browser.end_session(session_id = sid)

executar()
```

### Resposta representativa de `get_forms` (JSON ilustrativo)

```json
{
  "forms": [{
    "ref": "form_1", "id": null, "action": "/search", "method": "GET",
    "fields": [{"ref": "field_1", "tag": "textarea", "name": "q", "type": "text"}],
    "buttons": [{"ref": "button_1", "tag": "input", "type": "submit"}]
  }],
  "truncated": false
}
```

O exemplo é ilustrativo, não uma leitura verificada do DOM atual do Google. Não presumir IDs ou seletores de sites reais sem observação.

## Condições de aceite

1. Todos os 33 métodos são expostos por um catálogo versionado de assinaturas; parâmetro desconhecido gera erro determinístico.
2. Saídas de inspeção têm limites, sinalizador `truncated` e representação estável de ausências.
3. `element_ref` de outra sessão, aba ou revisão não pode ser acionado.
4. Testes cobrem formulário, lista, paginação por link/JS, download PDF e comportamento sem resultados.
5. Timeouts, rede e autorizações são verificados **no Go**, não delegados ao script.
6. Mesmo sem `browser.end_session` (erro/cancelamento), o supervisor encerra os recursos.

## Referências
- [Arquitetura do motor](../arquitetura/ARQ-004-motor-starlark.md)
- [Comandos administrativos](CONTR-003-administracao-scripts.md)
- [Prompt gerador](../prompts/PROMPT-001-geracao-starlark.md)
- [Exemplo de coleta](../exemplos/EX-001-coleta-pdfs-starlark.md)
- [Operações gerais](CONTR-001-operacoes-de-navegacao.md)
