<!-- A-BIOMA-DOC-STANDARD -->
# DEC-003 — Versionamento de Chromium e reutilização de imagens

**Decisão:** adotada em princípio.  
**Implementação:** não verificada.

## Contexto

O serviço deve oferecer opção de atualizar Chromium para a versão estável mais recente do canal configurado ou executar uma versão específica. É necessário reproduzir testes e evitar instalação direta no host.

## Decisão

Distribuir Chromium dentro de imagens OCI. Oferecer seleção `latest`, `pinned` e `cached`:
- `latest`: buscar versão estável mais atual publicada nas origens de pacotes Debian confiáveis configuradas;
- `pinned`: exigir versão/revisão fixa resolvida por origem confiável;
- `cached`: reutilizar imagem já validada pelo digest, inclusive sem acesso externo.

Uma imagem validada pode servir a muitas sessões, mas seus perfis e volumes mutáveis nunca serão compartilhados. A criação de imagens ocorre a partir de base/Containerfile controlados, não de sessões que navegaram na internet.

## Consequências

- Incluir a versão/revisão real do pacote Chromium e o digest OCI no inventário.
- A cada nova imagem, executar verificações de sandbox e isolamento antes de habilitar novas sessões.
- Versão mais recente é relativa ao canal configurado; não garantir que seja o commit mais novo do projeto Chromium.
- Não modificar imagem de sessão em andamento; novas imagens só atendem novas sessões.
- Definir futuramente política de descarte para imagens sem referências ativas.

## Referências

- [Imagens e atualização](../arquitetura/ARQ-003-imagens-chromium.md)
- [Persistência](../persistencia/PERS-001-sessoes-sqlite.md)
- [Snapshot Debian](https://snapshot.debian.org/)
