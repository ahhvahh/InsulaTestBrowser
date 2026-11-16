<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../mil-0001.md).

# Módulos — arquitetura hexagonal

- [sessoes](sessoes/README.md): Coordena casos de uso e compensação; define políticas de sessão e usa portas injetadas.
- [ambiente](ambiente/README.md): Implementa a porta de ambiente com Podman Rootless, sem conhecer tabelas de histórico.
- [navegacao](navegacao/README.md): Implementa a porta Chromium e devolve resultados tipados de páginas, cliques e downloads.
- [historico](historico/README.md): Controla passos, alvos sanitizados e catálogo de arquivos; solicita persistência por portas.
- [persistencia](persistencia/README.md): Implementa portas de repositório em SQLite, com SQL parametrizado e transações.

## Documentação em absorção

[Automação anterior](automacao/README.md) mantém os contratos agregados e as funcionalidades ainda não decompostas. A arquitetura vigente de trabalho está em [ARQ-005](../arquitetura/ARQ-005-arquitetura-hexagonal.md).
