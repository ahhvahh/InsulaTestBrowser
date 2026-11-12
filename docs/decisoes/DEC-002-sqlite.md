<!-- A-BIOMA-DOC-STANDARD -->
# DEC-002 — SQLite para gerenciamento de sessões

**Decisão:** adotada para o projeto.  
**Implementação:** não verificada.

## Contexto

O serviço precisa manter um inventário pequeno de sessões, imagens OCI, identificadores de contêineres, estados e vencimento de artefatos. Não há requisito para banco de dados compartilhado em rede ou múltiplos nós gravando simultaneamente.

## Decisão

Usar SQLite local ao host, acessado somente pelo serviço Go. O banco persistirá metadados; perfis do navegador, cookies, capturas e downloads permanecerão fora dele.

## Alternativas consideradas

- **Somente arquivos JSON:** baixa complexidade inicial, mas concorrência, recuperação e consultas são mais difíceis.
- **PostgreSQL:** robusto, porém adiciona serviço e administração desnecessários para o escopo inicial.

## Consequências

- Usar transações curtas, chaves estrangeiras, `busy_timeout` e migrações versionadas.
- Avaliar WAL para concorrência de leitura; SQLite permite somente um escritor de cada vez.
- Manter o banco em armazenamento local, com permissões restritas; não colocar o arquivo WAL em filesystem de rede.
- Reconciliar estado com Podman após falha/reinício; registro em SQLite não prova que o contêiner existe.
- Implementar expurgo de metadados e retenção de artefatos de modo idempotente.

## Referências

- [Modelo relacional proposto](../persistencia/PERS-001-sessoes-sqlite.md)
- [SQLite WAL](https://www.sqlite.org/wal.html)
- [SQLite — isolamento](https://www.sqlite.org/isolation.html)
