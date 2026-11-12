<!-- A-BIOMA-DOC-STANDARD -->
# ARQ-002 — Modelo de segurança e isolamento

**Situação:** Proposto. **Sandbox do Chromium:** Não verificado no ambiente-alvo.

## Objetivo e modelo de ameaça

Reduzir acesso de páginas hostis aos recursos do host e impedir que sessões distintas compartilhem perfis, cookies, arquivos e sockets de administração. Um contêiner compartilha o kernel com o host e **não oferece equivalência ao isolamento de uma VM**.

## Política obrigatória

| Área | Requisito |
| --- | --- |
| Identidade | Serviço Podman Rootless por usuário de sistema dedicado, sem privilégios administrativos |
| Execução | Um contêiner por sessão; não usar `--privileged` nem `--no-sandbox` |
| Navegador | Processo Chromium não privilegiado, sandbox interno funcional e comprovado antes de aceitar navegação |
| Filesystem | Imagem base imutável na sessão; volumes graváveis/temporários exclusivos, somente onde indispensáveis |
| Credenciais | Não disponibilizar credenciais do host, SSH agent nem socket Podman ao contêiner; arquivo `config.yaml` só aponta para segredos protegidos |
| Kernel | Limites de processos, CPU/memória, capacidades mínimas e seccomp; evitar `seccomp=unconfined` |
| Rede | Política de saída definida, negar redes de host/loopback/privadas e destinos de metadados por padrão |
| Depuração | Nunca publicar porta CDP em interfaces externas; acesso de comando estritamente autenticado/autorizado |
| Retenção | Remover perfis e segredos ao encerrar; artefatos podem permanecer apenas pelo prazo configurado |
| Logs | Não registrar tokens, cookies, conteúdo de formulários ou URLs sensíveis em texto aberto |

**Falha segura:** se sandbox, perfil privado, permissões ou política de rede não puderem ser estabelecidos, a criação da sessão DEVE falhar, sem afrouxamento automático.

## Cuidados específicos

- Rootless reduz privilégios, mas não impede todos os ataques ao kernel. Para ameaça elevada, avaliar VM ou gVisor como evolução.
- Políticas DNS isoladas não bastam contra SSRF; controlar destinos e redirecionamentos, inclusive alterações de resolução.
- A separação das sessões exige diretórios e volumes distintos, não apenas abas ou contextos do mesmo Chromium.
- O acesso ao socket Podman concede controle relevante sobre recursos do usuário executor; restringir permissões e não montar em contêiner.
- Compatibilidade entre sandbox Chromium, user namespaces e seccomp precisa ser comprovada no Debian-alvo. Não mascarar falhas com flags inseguras.
- Os artefatos retidos exigem permissões restritas e limpeza resistente a falhas/reinícios.
- A configuração e os arquivos de credenciais têm acesso restrito ao usuário de serviço; os caminhos referenciados são validados e nunca montados nas sessões.
- `podman commit` só é permitido a contêineres destinados à **preparação de imagens**, sem histórico de navegação de usuário. Antes de capturar, verificar ausência de segredos e resíduos; não considerar exclusão de arquivos após uso de segredos garantia de remoção de camadas anteriores.

## Critérios de aceite de segurança

1. Duas sessões simultâneas não leem cookies, cache, downloads, arquivos temporários ou sockets uma da outra.
2. O Chromium não inicia quando o sandbox interno é desabilitado ou indisponível.
3. Página controlada por atacante não alcança serviços locais do host nem o socket de Podman.
4. Desligamento abrupto seguido de reinício não deixa contêineres ou artefatos indefinidamente órfãos.
5. Contêiner não consegue ler o arquivo SQLite da aplicação.

## Referências

- [Arquitetura](ARQ-001-visao-geral.md)
- [Arquivo de configuração](../referencia/REF-001-configuracao.md)
- [Preparação de imagens](../fluxos/FLUX-002-preparacao-imagem.md)
- [Fluxo e limpeza](../fluxos/FLUX-001-ciclo-de-vida.md)
- [Podman Rootless](https://github.com/containers/podman/blob/main/docs/tutorials/rootless_tutorial.md)
- [Ativação por socket do Podman](https://github.com/containers/podman/blob/main/docs/tutorials/socket_activation.md)
