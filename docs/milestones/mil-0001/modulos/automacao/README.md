<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->
<!-- A-BIOMA-DOC-STANDARD -->
<!-- A-BIOMA-DOC-REVISADO -->
<!-- A-BIOMA-DOC-VERSION {"plugin":"a-bioma-doc-standard","version":"2.0.0"} -->

[Milestone MIL-0001](../../../mil-0001.md). Proposta; publicação pendente.

# Módulo automacao

Sessões, imagens, scripts e artefatos pertencem ao mesmo módulo nesta revisão: os contratos e relações de persistência atuais são interdependentes. Não há isolamento de módulos implementado. Uma divisão futura exige extrair contratos compartilhados para `docs/global/`, antes de introduzir dependências entre módulos. Não criar global apenas para contornar a regra de isolamento.

- [CONTR-001-operacoes-de-navegacao](contratos/CONTR-001-operacoes-de-navegacao.md)
- [CONTR-002-api-starlark](contratos/CONTR-002-api-starlark.md)
- [CONTR-003-administracao-scripts](contratos/CONTR-003-administracao-scripts.md)
- [CONTR-004-artefatos-resultados](contratos/CONTR-004-artefatos-resultados.md)
- [DEC-001-isolamento-rootless](decisoes/DEC-001-isolamento-rootless.md)
- [DEC-002-sqlite](decisoes/DEC-002-sqlite.md)
- [DEC-003-versoes-chromium](decisoes/DEC-003-versoes-chromium.md)
- [DEC-004-configuracao-local](decisoes/DEC-004-configuracao-local.md)
- [SPEC-001-catalogo-artefatos](especificacoes/SPEC-001-catalogo-artefatos.md)
- [FLUX-0003-validar-session-id](fluxos/FLUX-0003-validar-session-id.md)
- [FLUX-001-ciclo-de-vida](fluxos/FLUX-001-ciclo-de-vida.md)
- [FLUX-002-preparacao-imagem](fluxos/FLUX-002-preparacao-imagem.md)
- [INTF-TERM-001-inicializacao](interfaces/terminal/INTF-TERM-001-inicializacao.md)
- [PERS-001-sessoes-sqlite](persistencia/PERS-001-sessoes-sqlite.md)
- [ENT-001-images](persistencia/entidades/ENT-001-images.md)
- [ENT-002-sessions](persistencia/entidades/ENT-002-sessions.md)
- [ENT-003-artifacts](persistencia/entidades/ENT-003-artifacts.md)
- [ENT-004-session-steps](persistencia/entidades/ENT-004-session-steps.md)
