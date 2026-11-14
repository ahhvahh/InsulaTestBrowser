<!-- A-BIOMA-DOC-STANDARD -->
# EX-001 — Varredura paginada de PDFs com Starlark

**Situação:** Exemplo **não executável ainda**; depende da implementação de [CONTR-002](../contratos/CONTR-002-api-starlark.md) e da inspeção real do site. Não é comprovação dos seletores/HTML atuais do portal Tatuí.

## Finalidade
Varredura de uma lista, abertura de páginas de detalhe, identificação de links/controles de PDF e gravação de artefatos. A navegação principal permanece na aba original para não perder a paginação. O identificador `list_ref` será descoberto novamente a cada página.

## Script de referência (`.star`)

```python
URL = "https://portal.tatui.sp.gov.br/atos-oficiais/legislacao/decretos-municipais"
MAX_PAGES = 25

def escolher_lista(candidatas):
    for lista in candidatas:
        if "Decreto" in lista.get("text_sample", ""):
            return lista
    return None

def executar():
    sessao = browser.create_session(profile = "chromium-base")
    sid = sessao["session_id"]
    browser.navigate(session_id = sid, url = URL)

    checkpoint = job.get_checkpoint()
    ultima_pagina = 0
    if checkpoint != None:
        ultima_pagina = checkpoint["page"]

    for pagina in range(1, MAX_PAGES + 1):
        # Retomada real exige voltar à página correta a partir do início.
        # Não saltamos páginas sem navegação correspondente.
        candidatas = browser.inspect_lists(session_id = sid)
        lista = escolher_lista(candidatas["lists"])
        if lista == None:
            job.log(level = "warn", message = "Lista não identificada; coleta interrompida")
            break

        itens = browser.get_list_items(session_id = sid, list_ref = lista["ref"])
        if pagina > ultima_pagina:
            for item in itens["items"]:
                detalhe = item.get("url")
                if detalhe == None:
                    continue
                aba = browser.open_tab(session_id = sid, url = detalhe)
                tid = aba["tab_id"]
                encontrados = browser.find_downloads(
                    session_id = sid,
                    tab_id = tid,
                    file_type = "pdf"
                )
                for pdf in encontrados["downloads"]:
                    chave = detalhe + "#" + (pdf.get("url") or pdf["ref"])
                    if job.is_processed(key = chave):
                        continue
                    arquivo = browser.download(
                        session_id = sid,
                        tab_id = tid,
                        element_ref = pdf["ref"]
                    )
                    if arquivo["status"] == "completed":
                        job.record_document(
                            title = item.get("title") or "Documento",
                            source_url = detalhe,
                            artifact_id = arquivo["artifact_id"],
                            key = chave
                        )
                browser.close_tab(session_id = sid, tab_id = tid)

            job.checkpoint(page = pagina)

        pag = browser.inspect_pagination(session_id = sid)
        if not pag["has_next"]:
            break
        proxima = browser.next_page(
            session_id = sid,
            pagination_ref = pag["ref"]
        )
        if not proxima["advanced"]:
            break

    browser.end_session(session_id = sid)

executar()
```

## Premissas e cuidados
- Este script faz **descoberta heurística**; se os registros estiverem em modal, tabela ou links indiretos, adaptar a seleção usando dados reais de `get_dom`, `get_links` e `inspect_lists`.
- `MAX_PAGES` é limite de segurança, não número de páginas observado.
- `find_downloads` apenas lista candidatos; validar PDF com `artifacts.get_info` se exigido.
- Scripts Starlark não possuem `try/finally` como Python. O serviço Go **deve** encerrar sessões, abas e downloads remanescentes em erros/cancelamentos.
- A coleta respeita bloqueios, limites, redirecionamentos e políticas de acesso; não deve tentar contornar CAPTCHA, autenticação ou restrições.
- O checkpoint avança apenas depois de processar a página. Na retomada, o script visita as páginas anteriores sem coletá-las, então processa a próxima.
- Há risco de mudança de ordem da listagem entre execuções: deduplicação por chave e reconciliação dos itens são necessárias.

## Referências
- [Contrato de métodos](../contratos/CONTR-002-api-starlark.md)
- [Prompt de geração](../prompts/PROMPT-001-geracao-starlark.md)
