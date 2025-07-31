Okay, auqi vou ter algumas noções sobre como deve funcionar a API e como ela se comunicará com o front
Modelo provavelmente vai ser MVM pq é mais fácil e simples.

1. Receber Dados WRSAT
2. Formatar Dados WRSAT (labeling)
3. Armazenamento de Dados -> Preciso ver a posição, labeling e velocidade
4. Relatório baseado nos dados que foram salvos
5. Visualização em tempo real baseado nesses dados também


Proximos passos:
A ideia é, ao invés de gerar os relatórios e salvar em um DB, eu gero um Excel do dia especifico, assim então, encaminho via email e faço uma query para deletar os dados enviados.

1. Verificar o posicionamento dos motoristas
2. Validar o tipo do posicionamento (Carregamento, etc)
3. Modificar o State do motorista (Carregado, Vazio, etc)
4. Verificar o tempo que ficou parado em certos locais 

Acho que seria interessante dar prioridade no quarto tópico, e já fazer um acumulo de histórico com o Redis
Gerar relatórios em Excel tbm, mas ele não vai ficar mt tempo na memória uma vez que vai ser excluído uma vez que enviado por email.

