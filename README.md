# Microservice: Insurance

## Proteção para celular: 

*Endpoints (Oferta):*
| Endpoint | Http Method | Descrição |
| -------- | --------- | --------- |
| insurance/cellphone/v1/simulation | GET | Receber a marca e modelo e retornar o preço do valor e da franquia |
| insurance/cellphone/v1/buy | POST | Contratar o seguro, gerando um recibo |
| insurance/cellphone/v1/phone-brand-list | GET | Retorna uma lista com todas as marcas válidas |
| insurance/cellphone/v1/phone-model-list | GET | Retorna uma lista com todas os modelos válidos |

*Endpoints (Gestão):*
| Endpoint | Http Method | Descrição |
| -------- | --------- | ----------- |
| insurance/cellphone/v1/cancel | PUT | Retornar um feedback dizendo que o seguro foi cancelado |
| insurance/cellphone/v1/claim | POST | Acionar o sinistro passando o dados necessários |
| insurance/cellphone/v1/user-insurance | GET | Retorna os dados do seguro dele, se ele não tiver seguro retorna um 404

### Casos de uso

- Ao acionarem o sinistro é necessário enviar uma mensagem para a equipe de logistíca para que mandem um novo celular para o endereço;

- Calcular o preço e valor da franquia;
  
- O usuário pode cancelar o seguro quando ele quiser;

- O usário pode acionar o sinistro quando ele quiser;

- Quando o usuário clicar em contratar o seguro, deve-se chamar o sistema de pagamento, cobrando a primeira mensalidade;
  
- Necessário criar um motor para cobrar o usuário todo mês;
  
- Ao acionarem o sinistro é necessário chamar o sistema de pagamento cobrando o valor da franquia;

- Tem que receber o modelo e a marca do celular;

- É necesário receber informações do usuário para gerar o contrato.

