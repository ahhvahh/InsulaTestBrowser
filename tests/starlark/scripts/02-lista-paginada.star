# Lista de decretos: descoberta e até duas páginas, sem IDs inventados.
URL = "https://portal.tatui.sp.gov.br/atos-oficiais/legislacao/decretos-municipais"
MAX_PAGES = 2

def escolher_lista(candidates):
    for item in candidates:
        if "decreto" in item.get("text_sample", "").lower():
            return item
    return None

def executar():
    s = browser.create_session(profile = "chromium-base")
    sid = s["session_id"]
    browser.navigate(session_id = sid, url = URL)
    for page in range(1, MAX_PAGES + 1):
        candidates = browser.inspect_lists(session_id = sid)
        target = escolher_lista(candidates["lists"])
        if target == None:
            job.emit_result(data = {"scenario": "list", "status": "list_not_found", "page": page})
            break
        found = browser.get_list_items(session_id = sid, list_ref = target["ref"], limit = 10)
        job.emit_result(data = {"scenario": "list", "status": "ok", "page": page, "item_count": len(found["items"]), "items": found["items"]})
        job.checkpoint(page = page)
        pag = browser.inspect_pagination(session_id = sid)
        if not pag["has_next"] or page == MAX_PAGES:
            break
        advanced = browser.next_page(session_id = sid, pagination_ref = pag["ref"])
        if not advanced["advanced"]:
            break
    browser.end_session(session_id = sid)

executar()
