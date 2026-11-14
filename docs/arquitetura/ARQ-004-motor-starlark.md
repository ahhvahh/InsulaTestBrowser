<!-- A-BIOMA-DOC-STANDARD -->
# ARQ-004 — Motor de scripts Starlark

**Situação:** Proposto; nenhum motor ou método implementado neste repositório.

## Objetivo
Executar algoritmos de coleta produzidos por humanos ou LLMs sem recompilar o serviço Go. A LLM **gera** o arquivo `.star`; o InsulaTestBrowser **valida e executa** o arquivo mediante política autorizada. A LLM não participa do laço de navegação em tempo de execução.

## Componentes e fronteiras

| Elemento | Responsabilidade |
| --- | --- |
| Cliente local (UDS) | Autorizar e solicitar validação/execução; consultar progresso |
| Serviço Go | Aplicar permissões e políticas; criar e supervisionar execuções e sessões |
| Executor Starlark em processo auxiliar | Interpretar `.star`, com limites de CPU, memória, duração e passos |
| API exposta ao script | Somente `browser`, `artifacts` e `job`; objetos são wrappers restritos |
| Adaptador Chromium/CDP | Inspecionar DOM, navegar, interagir e acompanhar downloads |
| Podman Rootless | Isolar o Chromium em contêiner descartável por sessão |
| SQLite | Registrar metadados, checkpoint e estado da execução, sem conteúdo sensível |
| Armazenamento privado | Guardar PDFs, capturas e resultados, acessíveis por `artifact_id` |

O executor **não recebe** acesso irrestrito ao socket Podman, filesystem do host, segredos, rede ou funções Go internas. Chamadas à API passam pelo serviço e por suas políticas. O transporte interno do executor e o canal concreto Go→CDP ainda precisam de protótipos. Interface externa: UDS; proposta de serialização: MessagePack, com exemplos legíveis em JSON.

## Modelo de execução

1. Cliente envia fonte Starlark ou identificador de script previamente cadastrado, com parâmetros e política autorizada.
2. Serviço valida sintaxe, nomes de funções, parâmetros e política; prepara `execution_id` e supervisão.
3. Executor recebe exclusivamente os objetos `browser`, `artifacts`, `job` e dados primitivos autorizados.
4. O script cria uma sessão, navega, inspeciona a página e identifica candidatos a links, listas e paginação.
5. Cada chamada de browser é mediada pelo serviço, sujeita a autorização, quotas e allowlist de destinos.
6. Downloads completos viram artefatos; documentos e checkpoints são registrados de modo idempotente.
7. Ao terminar, falhar, ser cancelado ou exceder prazo, o **serviço Go** encerra todas as sessões/abas pertencentes à execução e reconcilia os recursos.

Os métodos são especificados em [CONTR-002](../contratos/CONTR-002-api-starlark.md) e os comandos de controle em [CONTR-003](../contratos/CONTR-003-administracao-scripts.md).

## Segurança e limites obrigatórios

- Starlark **não é Python**: sem imports arbitrários, sistema de arquivos, subprocessos, sockets ou bibliotecas externas. O executor expõe apenas as funções documentadas.
- Usar `starlark.Thread.SetMaxExecutionSteps` e cancelamento, **mas esses recursos não limitam memória nem necessariamente interrompem imediatamente uma função Go bloqueada**. Exigir processo auxiliar com limites de recursos e timeouts por chamada.
- Rejeitar execução sem teto de páginas, downloads, tempo, bytes por arquivo, resposta DOM e sessões concorrentes.
- URLs, redirecionamentos e requisições do navegador devem respeitar política de saída de rede, inclusive destinos internos/SSRF.
- Atributos, texto e HTML de sites são **dados não confiáveis**. Nunca interpretar conteúdo de páginas como instruções de política, comandos internos ou novo script autorizado.
- `element_ref` / `list_ref` / `pagination_ref` são opacos, específicos da aba e revisão do DOM; após mudança de página precisam ser redescobertos.
- Não persistir cookies, credenciais ou HTML sensível em SQLite/logs; somente metadados e identificadores de artefatos.
- Sem `try/finally` garantido nos scripts: a limpeza é **responsabilidade do supervisor Go**, inclusive após encerramento abrupto.

## Condições para implementação
Validar o sandbox Chromium/Podman e o canal CDP antes de considerar a API operacional. Fixar uma versão testada de `go.starlark.net` e executar testes com scripts válidos, inválidos, maliciosos e cancelados.

## Referências
- [Arquitetura geral](ARQ-001-visao-geral.md)
- [Segurança](ARQ-002-seguranca.md)
- [Contrato Starlark](../contratos/CONTR-002-api-starlark.md)
- [Administração](../contratos/CONTR-003-administracao-scripts.md)
- [API oficial de Starlark-Go](https://pkg.go.dev/go.starlark.net/starlark)
- [Especificação da linguagem](https://chromium.googlesource.com/external/github.com/google/starlark-go/+/HEAD/doc/spec.md)
- [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
