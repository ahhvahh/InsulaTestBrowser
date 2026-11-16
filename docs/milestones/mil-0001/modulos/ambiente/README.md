<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../mil-0001.md).

# Módulo ambiente

Implementa a porta de ambiente com Podman Rootless, sem conhecer tabelas de histórico.

[Arquitetura e chamadas permitidas](../../arquitetura/ARQ-005-arquitetura-hexagonal.md). Todos os contratos desta pasta são propostas de portas internas; implementação pendente.

## Documentos

- [CONTR-100-criar-container](contratos/CONTR-100-criar-container.md)
- [CONTR-103-remover-container](contratos/CONTR-103-remover-container.md)
- [DEC-001-isolamento-rootless](decisoes/DEC-001-isolamento-rootless.md)
- [DEC-003-versoes-chromium](decisoes/DEC-003-versoes-chromium.md)
- [FLUX-0200-criar-container](fluxos/FLUX-0200-criar-container.md)
- [FLUX-0201-remover-container](fluxos/FLUX-0201-remover-container.md)
