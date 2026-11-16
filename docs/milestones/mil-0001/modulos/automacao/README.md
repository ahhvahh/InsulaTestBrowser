<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../mil-0001.md).

# Automação — documentação em absorção

A organização em módulo único foi substituída pela [arquitetura hexagonal](../../arquitetura/ARQ-005-arquitetura-hexagonal.md). Esta pasta preserva contratos agregados e fluxos anteriores até a absorção integral. Não define um sexto módulo executável nem autoriza chamadas aos seus detalhes internos.

- [CONTR-001-operacoes-de-navegacao](contratos/CONTR-001-operacoes-de-navegacao.md)
- [CONTR-002-api-starlark](contratos/CONTR-002-api-starlark.md)
- [CONTR-003-administracao-scripts](contratos/CONTR-003-administracao-scripts.md)
- [CONTR-004-artefatos-resultados](contratos/CONTR-004-artefatos-resultados.md)
- [SPEC-001-catalogo-artefatos](especificacoes/SPEC-001-catalogo-artefatos.md)
- [FLUX-0003-validar-session-id](fluxos/FLUX-0003-validar-session-id.md)
- [FLUX-001-ciclo-de-vida](fluxos/FLUX-001-ciclo-de-vida.md)
- [FLUX-002-preparacao-imagem](fluxos/FLUX-002-preparacao-imagem.md)
- [INTF-TERM-001-inicializacao](interfaces/terminal/INTF-TERM-001-inicializacao.md)

Fluxos antigos FLUX-001/002 descrevem ciclos completos; os novos fluxos são recortes com assinaturas tipadas. Persistência física foi transferida para persistencia; decisões técnicas foram distribuídas por responsabilidade.
