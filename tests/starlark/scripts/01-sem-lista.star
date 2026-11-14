# Página sem lista: sessão, informações e links.
URL = "https://example.com/"

def executar():
    s = browser.create_session(profile = "chromium-base")
    sid = s["session_id"]
    browser.navigate(session_id = sid, url = URL)
    page = browser.get_page_info(session_id = sid)
    lists = browser.inspect_lists(session_id = sid)
    links = browser.get_links(session_id = sid, limit = 10)
    job.emit_result(data = {"scenario": "simple_page", "url": page["url"], "title": page["title"], "list_count": len(lists["lists"]), "links": links["links"]})
    browser.end_session(session_id = sid)

executar()
