# Testes Starlark — InsulaTestBrowser

Os scripts em scripts/ cobrem 4 cenários: página sem lista, lista paginada, pesquisa Google e pesquisa YouTube com filtro de vídeos.

Para testar localmente, na raiz do repositório:

    bash tests/starlark/run.sh

O runner Go usa starlark.ExecFile e respostas simuladas de browser/job. Executa 7 cenários (4 principais e 3 de ausência de elementos). Não usa Chromium, Podman, rede, UDS ou CDP. A dependência go.starlark.net é baixada por go mod tidy.

Os scripts do YouTube e Google dependem da interface real de cada site; são smoke tests, não prova de seletores atuais. Quando os comandos run_script e métodos do navegador forem implementados, realizar E2E conforme docs/testes/TEST-001-navegacao-starlark.md.
