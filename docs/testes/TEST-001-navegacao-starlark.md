<!-- A-BIOMA-DOC-STANDARD -->
# TEST-001 — Smoke tests Starlark: navegação e pesquisa

**Situação:** testes simulados preparados; execução real com Chromium/Podman/UDS ainda pendente.

## Casos

| Script | Referência | Verificação |
| --- | --- | --- |
| 01-sem-lista.star | https://example.com | Criação de sessão, navegação, título, links e ausência de lista obrigatória |
| 02-lista-paginada.star | https://portal.tatui.sp.gov.br/atos-oficiais/legislacao/decretos-municipais | Descoberta de lista, itens, avanço de página e checkpoints; máximo 2 páginas |
| 03-google-pesquisa.star | https://www.google.com | Inspeção de formulário, preenchimento do campo de busca e envio; fallback de URL fixa |
| 04-youtube-filtro.star | https://www.youtube.com/results?search_query=golang | Pesquisa de golang, descoberta do botão Filtros e tentativa de selecionar Vídeos |

## Executar

Na raiz do repositório: bash tests/starlark/run.sh

O runner Go usa Starlark-Go com objetos browser/job simulados, sem rede. Valida fluxos, retornos, avanços de página, ausência de elementos e encerramento de sessão. Baixa dependências com go mod tidy antes dos testes.

**Limites:** não acessa os sites. Para validação real faltam executor Starlark, Podman Rootless, Chromium/CDP e o contrato de run_script. Uma vez implementados, enviar a fonte completa autorizada ou script_id, não um caminho arbitrário do host, usando CONTR-003. Somente relatar filtro YouTube aplicado após verificar o estado real do site; um clique não prova a filtragem.

Se não encontrar lista ou botão, o script emite list_not_found ou filter_button_not_found; sem formulário do Google utiliza direct_url_fallback. Não contornar login, CAPTCHA, consentimento ou bloqueio.

## Referências
- [API Starlark](../contratos/CONTR-002-api-starlark.md)
- [Administração](../contratos/CONTR-003-administracao-scripts.md)
- [Motor](../arquitetura/ARQ-004-motor-starlark.md)
- [Scripts e runner](../../tests/starlark/README.md)
