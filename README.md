<!-- A-BIOMA-DOC-STANDARD -->
# InsulaTestBrowser

**Serviço de automação e testes de navegação web com Chromium isolado por sessão em contêineres Podman Rootless.**

**Situação:** planejamento arquitetural. Nenhuma implementação de navegador, contêiner, API ou banco foi comprovada neste repositório.

## Objetivo

Permitir que uma aplicação Go crie, controle e encerre sessões independentes de Chromium para navegar em sites, interagir com elementos de páginas e produzir screenshots. Cada sessão usará um contêiner descartável, um perfil próprio e limites de recursos. A prioridade é segurança e isolamento, não a velocidade de inicialização.

## Tecnologias e responsabilidades

| Tecnologia | Papel |
| --- | --- |
| Go | Serviço orquestrador e interfaces de controle |
| Podman Rootless | Criação, inspeção e exclusão de contêineres independentes |
| Chromium | Navegação, automação, interação e screenshots dentro do contêiner |
| SQLite | Inventário de imagens, sessões, estados e retenção de artefatos |
| OCI | Imagens versionadas e reutilizáveis contendo Chromium |

O Chromium será distribuído na imagem OCI e **não precisa estar instalado no host**. O Podman, suas dependências e o suporte Linux a execução rootless continuam necessários.

## Recursos planejados

- Criar e encerrar sessões independentes, com viewport e resolução configuráveis.
- Navegar em URLs, interagir com elementos e capturar screenshots.
- Controlar Podman por socket Unix local, sem expor o socket administrativo ao navegador.
- Criar e reutilizar imagens OCI de Chromium com modos de versão `latest`, `pinned` e `cached`.
- Carregar `config.yaml` ao lado do executável ou na pasta indicada por `--config-dir`; centralizar banco SQLite, acesso, referências a credenciais e perfis de contêiner.
- Instanciar sessões a partir de imagem validada ou preparar um contêiner novo e salvar sua imagem para reutilização, sem capturar dados de navegação.
- Persistir estados mínimos em SQLite e reconciliar sessões com os contêineres reais após reinícios.
- Expirar perfis imediatamente após o encerramento e aplicar retenção configurável a screenshots, downloads, logs e metadados.
- Aplicar limites de recursos e políticas de acesso à rede; exigir sandbox Chromium funcional.

Contêineres compartilham o kernel Linux. Por isso, o nível de proteção **não equivale ao de uma VM**, e a compatibilidade do sandbox Chromium com Podman Rootless deverá ser testada.

## Documentação

Consulte [docs/README.md](docs/README.md) para o índice completo do padrão A-Bioma. A referência de [configuração centralizada](docs/referencia/REF-001-configuracao.md) e a [interface de inicialização](docs/interfaces/terminal/INTF-TERM-001-inicializacao.md) detalham a descoberta do `config.yaml`.

## Licença

[Apache License 2.0](LICENSE). Componentes de terceiros, incluindo Chromium e dependências presentes nas imagens, mantêm suas próprias condições de licença.
