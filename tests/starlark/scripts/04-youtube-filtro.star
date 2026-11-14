# Pesquisa YouTube: tenta escolher Filtros > Vídeos por rótulos observados.
SEARCH_URL = "https://www.youtube.com/results?search_query=golang"

def choose_visible(items, names):
    for item in items:
        label = (item.get("text") or item.get("label") or "").strip().lower()
        if item.get("visible", True) and label in names:
            return item
    return None

def executar():
    s = browser.create_session(profile = "chromium-base")
    sid = s["session_id"]
    browser.navigate(session_id = sid, url = SEARCH_URL)
    buttons = browser.get_buttons(session_id = sid, limit = 100)
    filters = choose_visible(buttons["buttons"], ["filters", "filtros"])
    status = "filter_button_not_found"
    if filters != None:
        browser.click(session_id = sid, element_ref = filters["ref"])
        options = browser.get_elements(session_id = sid, selector = "button, a[href], tp-yt-paper-item, ytd-search-filter-renderer", limit = 100)
        videos = choose_visible(options["elements"], ["video", "videos", "vídeo", "vídeos"])
        if videos != None:
            browser.click(session_id = sid, element_ref = videos["ref"])
            status = "video_filter_clicked"
        else:
            status = "video_filter_not_found"
    links = browser.get_links(session_id = sid, limit = 10)
    job.emit_result(data = {"scenario": "youtube_filter", "status": status, "url": browser.get_page_info(session_id = sid)["url"], "links": links["links"]})
    browser.end_session(session_id = sid)

executar()
