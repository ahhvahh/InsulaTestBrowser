<!-- A-BIOMA-DOC-STANDARD -->
# Documentação — InsulaTestBrowser

**Situação:** projeto em definição. As interfaces, o banco de dados e os fluxos abaixo são **propostos**, não implementados.

O InsulaTestBrowser é um serviço Go que controla o Podman Rootless para iniciar uma instância Chromium por sessão, executar operações de navegação e descartar os recursos ao término.

## Índice

| Tema | Documento |
| --- | --- |
| Contexto, componentes e implantação | [ARQ-001 — Arquitetura](arquitetura/ARQ-001-visao-geral.md) |
| Ameaças, limites e isolamento | [ARQ-002 — Segurança](arquitetura/ARQ-002-seguranca.md) |
| Versões e imagens OCI | [ARQ-003 — Imagens](arquitetura/ARQ-003-imagens-chromium.md) |
| Modelo de dados SQLite | [PERS-001 — Persistência](persistencia/PERS-001-sessoes-sqlite.md) |
| Criação e encerramento da sessão | [FLUX-001 — Ciclo de vida](fluxos/FLUX-001-ciclo-de-vida.md) |
| Operações da aplicação | [CONTR-001 — Contrato conceitual](contratos/CONTR-001-operacoes-de-navegacao.md) |
| Escolha de Podman Rootless | [DEC-001 — Isolamento](decisoes/DEC-001-isolamento-rootless.md) |
| Escolha de SQLite | [DEC-002 — Persistência](decisoes/DEC-002-sqlite.md) |
| Política de versão do Chromium | [DEC-003 — Versões](decisoes/DEC-003-versoes-chromium.md) |

## Convenções

- **Proposto:** requisito ou desenho ainda sem implementação comprovada.
- **Implementado:** exige evidência no código e testes; nenhum componente foi confirmado neste levantamento.
- **Não verificado:** depende de ensaio técnico, especialmente o sandbox do Chromium dentro do contêiner.
- `DEVE`, `NÃO DEVE`, `RECOMENDA-SE` e `PODE` expressam força de requisito como convenção interna inspirada nas RFC 2119/8174.

Prefixos locais: `ARQ` arquitetura, `PERS` persistência, `FLUX` fluxo, `CONTR` contrato e `DEC` decisão.

## Escopo documental

Este conjunto documenta decisões e propostas. Não pressupõe que já exista serviço, schema SQLite, imagem OCI, socket, comando CLI ou implementação do protocolo de navegação.
