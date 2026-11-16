<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../mil-0001.md).

# Módulo persistencia

Implementa portas de repositório em SQLite, com SQL parametrizado e transações.

[Arquitetura e chamadas permitidas](../../arquitetura/ARQ-005-arquitetura-hexagonal.md). Todos os contratos desta pasta são propostas de portas internas; implementação pendente.

## Documentos

- [CONTR-112-salvar-sessao](contratos/CONTR-112-salvar-sessao.md)
- [CONTR-115-gravar-passo](contratos/CONTR-115-gravar-passo.md)
- [CONTR-133-salvar-artefato](contratos/CONTR-133-salvar-artefato.md)
- [DEC-002-sqlite](decisoes/DEC-002-sqlite.md)
- [FLUX-0500-salvar-sessao](fluxos/FLUX-0500-salvar-sessao.md)
- [FLUX-0501-gravar-passo](fluxos/FLUX-0501-gravar-passo.md)
- [FLUX-0502-salvar-artefato](fluxos/FLUX-0502-salvar-artefato.md)
- [PERS-001-sessoes-sqlite](persistencia/PERS-001-sessoes-sqlite.md)
- [ENT-001-images](persistencia/entidades/ENT-001-images.md)
- [ENT-002-sessions](persistencia/entidades/ENT-002-sessions.md)
- [ENT-003-artifacts](persistencia/entidades/ENT-003-artifacts.md)
- [ENT-004-session-steps](persistencia/entidades/ENT-004-session-steps.md)
