<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md). Proposta; publicação pendente.

# SPEC-001 — Catálogo de artefatos por sessão

**Estado:** proposto. **Versão aprovada:** nenhuma.

## Objetivo e escopo

Permitir ao cliente autorizado acompanhar passos e recuperar arquivos antes e depois do encerramento do navegador, até a expiração. Inclui downloads, screenshots, HTML e texto; não inclui armazenamento permanente nem publicação HTTP.

## Pré-condições e decisões

Sessão criada e acesso autorizado. As definições de diretórios privados, catálogo SQLite, eventos incrementais e projeção JSON estão em [CONTR-004](../contratos/CONTR-004-artefatos-resultados.md). Essa é a fonte das regras existentes; a presente SPEC reúne o aceite sem criar uma decisão duplicada.

## Critérios de aceite

- **CA-01** — Isolamento de duas sessões: um cliente não acessa ou lista artefatos de outra sem autorização explícita.
- **CA-02** — `result.json` sempre é um JSON íntegro durante atualização; `events.jsonl` conserva ordem por sessão e suporta reconstrução.
- **CA-03** — Uma navegação, um clique e um download surgem como steps com `artifact_id` quando houver arquivo; falhas também ficam registradas.
- **CA-04** — HTML bruto, texto e screenshots são capturados com limites, checksum e metadados; não salvar HTML/valores confidenciais em SQLite ou logs.
- **CA-05** — Consulta por cursor entrega somente novos registros, sem duplicar steps; funciona antes e depois do fim.
- **CA-06** — Downloads e consultas por `artifact_id` funcionam após `end_session`, mas não após o expurgo.
- **CA-07** — Falhas, cancelamento e reinício do serviço não deixam diretórios ilimitados ou arquivos parcialmente publicados.
- **CA-08** — Testes cobrem permissões `0700/0600`, path traversal, symlink, limites, expiração, checksum e download parcial.

## Evidências e lacunas

Sem código do serviço ou testes executáveis do catálogo. [TEST-001](../../../guias/testes/TEST-001-navegacao-starlark.md) lista cenários previstos. Definir cursor para atualização de steps, atomicidade entre SQLite/eventos/JSON e recuperação após queda. Os testes com mocks não comprovam esses critérios.

<a id="sdd"></a>
## SDD — Spec-Driven Development

**Pendente:** o fluxo completo do catálogo ainda precisa ser decomposto em assinaturas e exceções tipadas; não recebe marca de validação 2.0. A validação lexical proposta em [FLUX-0003](../fluxos/FLUX-0003-validar-session-id.md) cobre apenas um pré-requisito e não comprova autorização nem segurança do filesystem.
