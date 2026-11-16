<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../mil-0001.md).

# REF-003 — Spec Rules e referências abreviadas

**Estado:** em refinamento. **Formato piloto:** a-bioma-spec-rules 0.2.0.

## Decisão e escopo

Em 2026-11-16 o usuário solicitou substituir o tópico SDD por Spec Rules, iniciar o YAML pelo ID do documento e representar nele a estrutura descrita em texto: contratos, campos, atores, passos, decisões, chamadas e exceções. Aplicação piloto em CONTR-103 e FLUX-0201; o restante da milestone conserva o formato anterior até migração explícita. O plugin não foi atualizado nesta rodada.

## Sintaxe de referências

| Forma | Resolução |
| --- | --- |
| `"CONTR-103"` | Definição do mesmo módulo, com preferência pela proposta desta milestone |
| `"@TIP-0100"` | Definição global |
| `"@[ambiente](CONTR-104)"` | Definição pública no módulo ambiente |
| `"@[ambiente](FLUX-0201)"` | Fluxo encapsulado público |

O caractere @ é seguido do ID no caso global. As referências são strings entre aspas no YAML; não usar barra invertida antes de parênteses. Elas não são links Markdown nem comandos executáveis. Não ampliar permissões: uma referência entre módulos continua sujeita à matriz ARQ-005.

Resolver por (escopo, ID), não presumir que cada ID tenha arquivo próprio. CONTR-104/105 são definidos dentro de CONTR-103. Priorizar proposta da própria milestone quando declarada, senão documento vigente; nunca consumir silenciosamente proposta de outra milestone. ID ausente ou ambíguo é erro documental.

## Catálogo de resolução do piloto

| Escopo | IDs definidos | Documento |
| --- | --- | --- |
| global | TIP-0100, TIP-0101, TIP-0102 | [CONTR-090](../global/contratos/CONTR-090-tipos-primitivos.md#sdd) |
| ambiente | CONTR-103, CONTR-104, CONTR-105; CMP-1007 a CMP-1011 | [CONTR-103](../modulos/ambiente/contratos/CONTR-103-remover-container.md#spec-rules) |
| ambiente | FLUX-0201, EXC-0201; PAS-1006 a PAS-1010 | [FLUX-0201](../modulos/ambiente/fluxos/FLUX-0201-remover-container.md#spec-rules) |

CHAMADA-103-01 a CHAMADA-103-03 são identificadores locais do encadeamento contratual, qualificados por CONTR-103. Não substituem os PAS do fluxo encapsulado.

## Estrutura obrigatória para comportamento

O último tópico é **Spec Rules**, contendo um único bloco YAML. Sua única chave raiz é o ID do documento. Abaixo dela: formato, versao, tipo_documento, modulo e estado. Contratos definem dados; fluxos definem execução.

- Estado documental: `proposta`, `em refinamento` ou `refinado`. Implementação é uma dimensão separada.
- Entrada é um objeto. Para uma entrada, usar nome e doc. Para várias, usar nomes como chaves e doc dentro de cada nome; não misturar as duas formas. Nunca usar lista em entrada.
- Saída é um objeto com doc para exatamente um contrato normal. Nunca usar lista em saida. Falhas ficam em excecoes.
- Contratos definidos no documento incluem campos, tipos, obrigatoriedade, IDs e regras. Referenciar outra definição não cria uma cópia.
- Fluxo contém atores e root. Root é lista ordenada de passos com id, acao e descricao quando aplicável; os nomes humanos são descrições, não chaves dinâmicas de execução.
- Cada chamada encapsulada informa doc, argumentos, guardar_em e tratamento das exceções expostas.
- Validar informa condicao, valido e invalido. O ramo inválido identifica contrato, exceção e dados de falha e encerra a execução. O válido continua no próximo passo.
- Operar identifica uma operação primitiva externa, seus contratos e erros; não é uma referência documental.
- Retornar e falhar encerram o caminho. Caminhos de dados usam objetos `{caminho: ...}`; constantes usam `{literal: ...}`.
- Exceções são declaradas com contrato de dados. Lista vazia somente quando ausência conhecida for justificada.
- Objetivo, acesso, regras, efeitos, telemetria, recuperação, critérios e pendências devem constar no YAML quando apresentados no documento e pertinentes ao comportamento.
- Regras por campo são declarativas; a execução é explícita nos passos, evitando recursão automática.
- Não usar aliases, anchors, tags ou chaves duplicadas. A ordem da lista é a ordem de execução.

O contrato apenas referencia o fluxo responsável pelas regras. Atores, sequência, chamadas e efeitos pertencem ao documento FLUX/ALG.

## Coexistência e validação

Versão 0.2.0 é a primeira proposta deste novo formato. Ela não deve usar a marca A-BIOMA-SPEC:2.0 nem alegar aprovação pelo schema anterior. O marcador piloto é A-BIOMA-SPEC-RULES:0.2.0.

O alias HTML sdd foi preservado para links já existentes; isso preserva navegação, não compatibilidade de parsers. Consumidores documentais antigos, inclusive as referências de FLUX-0100 à remoção, precisam reconhecer o novo formato ou permanecer explicitamente pendentes de migração. Chamadas existentes preservam a assinatura e os IDs.

Nesta rodada houve conferência estrutural dos objetos usados para gerar o YAML e resolução das referências do piloto. Não houve validação por parser YAML nem por um JSON Schema novo. O gerador usa strings JSON entre aspas e indentação de mapas/listas, sem anchors. Schema, parser e migração dos demais documentos permanecem pendentes. Não há alteração no código do serviço.

## Objetos de entrada e saída

Uma entrada:

```yaml
entrada:
  nome: "solicitacao"
  doc: "@[ambiente](CONTR-103)"
saida:
  doc: "@[ambiente](CONTR-104)"
```

Várias entradas usam a mesma ordem lógica dos argumentos nomeados, sem depender da ordem das propriedades:

```yaml
entrada:
  pedido:
    doc: "CONTR-EXEMPLO-ENTRADA"
  contexto:
    doc: "CONTR-EXEMPLO-CONTEXTO"
```

O segundo exemplo é fictício. Listas continuam adequadas a passos ordenados, critérios e coleções; seu marcador deve acompanhar a primeira propriedade, por exemplo `- id: PAS-1006`, evitando linhas com um traço isolado.

## Atualização do plugin

A proposta de atualização está em [plugin-update.json](plugin-update.json), com os arquivos alterados, identidade e release-base verificadas. Publicação do plugin pendente: requer empacotamento local indisponível nesta sessão. Não apresentar a versão preparada como instalada.

## Separação por tipo documental

CONTR define campos, tipos, obrigatoriedade, restrições, estruturas e referências declarativas às regras. Não exige atores, root, chamadas, passos, efeitos ou retorno de execução. FLUX/ALG define entrada e saída como objetos, atores, sequência, decisões, chamadas, efeitos e tratamento das exceções. Referencia os contratos sem copiar seus campos.

O schema deve selecionar as exigências conforme tipo_documento. O pacote de atualização do plugin anteriormente preparado contém um schema inicial de assinatura de fluxo; sua adequação a contratos permanece pendente antes de publicar o plugin.
