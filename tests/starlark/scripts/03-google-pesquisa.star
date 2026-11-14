# Pesquisa no Google usando o formulário detectado; URL alternativa fixa.
URL = "https://www.google.com/"
QUERY = "Starlark"
FALLBACK_SEARCH_URL = "https://www.google.com/search?q=Starlark"

def find_search_field(forms):
    for form in forms:
        for field in form.get("fields", []):
            if field.get("name") == "q":
                return {"field_ref": field["ref"], "form_ref": form["ref"]}
    return None

def executar():
    s = browser.create_session(profile = "chromium-base")
    sid = s["session_id"]
    browser.navigate(session_id = sid, url = URL)
    forms = browser.get_forms(session_id = sid)
    search = find_search_field(forms["forms"])
    method = "form"
    if search != None:
        browser.fill_field(session_id = sid, element_ref = search["field_ref"], value = QUERY)
        browser.submit_form(session_id = sid, form_ref = search["form_ref"])
    else:
        method = "direct_url_fallback"
        browser.navigate(session_id = sid, url = FALLBACK_SEARCH_URL)
    links = browser.get_links(session_id = sid, limit = 10)
    page = browser.get_page_info(session_id = sid)
    job.emit_result(data = {"scenario": "google_search", "method": method, "query": QUERY, "url": page["url"], "links": links["links"]})
    browser.end_session(session_id = sid)

executar()
