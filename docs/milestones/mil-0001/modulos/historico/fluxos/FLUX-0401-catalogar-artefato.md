<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# FLUX-0401 — Catalogar artefato

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** historico.

## Objetivo e pré-condições

Validar metadados confirmados pelo produtor confiável e persistir localização do arquivo.

IDs opacos; storage_key relativa à raiz privada da sessão, sem traversal; tamanho não negativo; sha256 hexadecimal de 64 caracteres; MIME validado pelo produtor. Apenas artefato concluído pode ser catalogado como ready.

Contratos: [entrada, saída e falha](../contratos/CONTR-130-catalogar-artefato.md). Assinatura pública, consumível somente conforme [matriz de dependências](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

## Sequência

```mermaid
sequenceDiagram
    participant U as Chamador
    participant A as historico
    participant B as SQLite
    U->>A: Contrato de entrada
    A->>A: Validar campos e contexto
    alt Entrada inválida
        A-->>U: invalid_request
    else Entrada válida
    A->>B: Operação contratada
    alt Dependência confirma
        B-->>A: Resultado tipado
        A->>A: Validar saída
        A-->>U: Resultado ou falha de contrato
    else Dependência falha
        B-->>A: Falha tipada
        A-->>U: Falha sanitizada
    end
    end
```

## Falhas, efeitos e telemetria

Não há retry automático. Cada chamada trata todas as falhas declaradas. Timeouts podem deixar efeito externo incerto; reconciliar antes de repetir. Correlacionar execução, session_id, step_id quando presente, fluxo, PAS e resultado; nunca registrar argumentos completos, tokens ou conteúdo de formulários. Logs técnicos do adaptador não substituem o histórico funcional.

## Critérios documentais

Cobrir sucesso, rejeição antes de efeito, falha da dependência e resposta incompatível. Não existem testes executados nem implementação destas portas. Em provisionamento, testar falha de persistência e falha da compensação; na navegação, testar histórico indisponível após efeito confirmado. As portas de CDP, autorização e migração SQL ainda precisam de implementação.


<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: FLUX-0401
  modulo: historico
  estado: proposto
  tipos: []
  contratos:
  - id: CONTR-130
    documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
  - id: CONTR-131
    documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
  - id: CONTR-132
    documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
  excecoes:
  - id: EXC-0401
    nome: Falha0401
    condicao: Validação, autorização, timeout ou falha da dependência desta porta.
    resultado: Encerrar a operação com código sanitizado, preservando efeitos já confirmados.
    recuperavel: false
    contrato:
      id: CONTR-132
      documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
  fluxos:
  - id: FLUX-0401
    nome: Catalogar artefato
    categoria: metodo
    entradas:
    - nome: solicitacao
      contrato:
        id: CONTR-130
        documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
    saida:
      id: CONTR-131
      documento: ../contratos/CONTR-130-catalogar-artefato.md#sdd
    excecoes:
    - id: EXC-0401
    efeitos:
    - Persistir catálogo de artefato
    passos:
    - id: PAS-1096
      acao: validar
      descricao: Validar tipos e regras de entrada antes de acessar dependências.
      condicao:
        descricao: IDs opacos; storage_key relativa à raiz privada da sessão, sem
          traversal; tamanho não negativo; sha256 hexadecimal de 64 caracteres; MIME
          validado pelo produtor. Apenas artefato concluído pode ser catalogado como
          ready. Todos os campos presentes e com os tipos declarados; rejeitar antes
          de interpretar valores.
        dados:
        - caminho: entradas.solicitacao
      falha:
        id: EXC-0401
      dados_falha:
        literal:
          code: invalid_request
    - id: PAS-1097
      acao: atribuir
      descricao: Inicializar objeto local.
      destino: variaveis.registro
      valor:
        literal: {}
    - id: PAS-1098
      acao: atribuir
      descricao: Copiar campo tipado session_id.
      destino: variaveis.registro.session_id
      valor:
        caminho: entradas.solicitacao.session_id
    - id: PAS-1099
      acao: atribuir
      descricao: Copiar campo tipado artifact_id.
      destino: variaveis.registro.artifact_id
      valor:
        caminho: entradas.solicitacao.artifact_id
    - id: PAS-1100
      acao: atribuir
      descricao: Copiar campo tipado storage_key.
      destino: variaveis.registro.storage_key
      valor:
        caminho: entradas.solicitacao.storage_key
    - id: PAS-1101
      acao: atribuir
      descricao: Copiar campo tipado mime_type.
      destino: variaveis.registro.mime_type
      valor:
        caminho: entradas.solicitacao.mime_type
    - id: PAS-1102
      acao: atribuir
      descricao: Copiar campo tipado size_bytes.
      destino: variaveis.registro.size_bytes
      valor:
        caminho: entradas.solicitacao.size_bytes
    - id: PAS-1103
      acao: atribuir
      descricao: Copiar campo tipado sha256.
      destino: variaveis.registro.sha256
      valor:
        caminho: entradas.solicitacao.sha256
    - id: PAS-1105
      acao: chamar
      descricao: Chamar porta pública FLUX-0502.
      fluxo:
        id: FLUX-0502
        documento: ../../persistencia/fluxos/FLUX-0502-salvar-artefato.md#sdd
      argumentos:
      - nome: solicitacao
        valor:
          caminho: variaveis.registro
      guardar_em: variaveis.resultado
      erros:
      - origem:
          id: EXC-0502
          documento: ../../persistencia/fluxos/FLUX-0502-salvar-artefato.md#sdd
        guardar_em: falhas.dependencia
        passos:
        - id: PAS-1104
          acao: falhar
          descricao: Encerrar com código sanitizado; não repetir efeitos.
          excecao:
            id: EXC-0401
          dados:
            literal:
              code: dependency_failure
    - id: PAS-1106
      acao: validar
      descricao: Verificar contrato de saída antes de retornar.
      condicao:
        descricao: Resultado contém todos os campos obrigatórios nos tipos declarados;
          session_id corresponde à entrada; IDs de recurso não vazios quando presentes;
          revision e sequence não negativos; saved/removed verdadeiros quando presentes.
        dados:
        - caminho: variaveis.resultado
        - caminho: entradas.solicitacao
      falha:
        id: EXC-0401
      dados_falha:
        literal:
          code: invalid_dependency_result
    - id: PAS-1107
      acao: retornar
      descricao: Entregar a única saída normal desta porta.
      valor:
        caminho: variaveis.resultado
```
