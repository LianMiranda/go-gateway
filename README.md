# Gateway de Pagamento - API Gateway (Go)

> Microsserviço da API Gateway desenvolvido em Go, parte do projeto Gateway de Pagamento criado durante a Imersão Full Stack & Full Cycle.

---

## ⚠️ Aviso Importante

Este projeto foi desenvolvido exclusivamente para fins didáticos como parte da Imersão Full Stack & Full Cycle.

---

## 📖 Sobre o Projeto

O **Gateway de Pagamento** é um sistema distribuído composto por:

- **Frontend** em Next.js  (Link)[https://github.com/LianMiranda/FrontEnd-Gateway]
- **API Gateway** em Go (este repositório)  
- **Sistema de Antifraude** em Nest.js  
- **Apache Kafka** para comunicação assíncrona  

---

## 🚀 Status Atual

### Implementado

- Setup e estrutura base do projeto  
- Endpoints de gerenciamento de **accounts** (criação e consulta)  
- Sistema completo de **faturas (invoices)** com:  
  - Criação e processamento automático de pagamentos  
  - Validação de limites (faturas > R$ 10.000 ficam pendentes)  
  - Consulta individual e listagem de faturas  
  - Atualização automática de saldo da conta  

### Pendências

- Integração com **Apache Kafka** para:  
  - Envio de transações para o microsserviço de antifraude  
  - Consumo de respostas do serviço de antifraude  
- Processamento de pagamentos baseado na análise de fraude  

---

## 🏗️ Arquitetura da Aplicação

> Visualize a arquitetura completa [neste diagrama](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy).

---

## 🎯 Pré-requisitos

- **Go** 1.24 ou superior  
- **Docker**  
  - Para Windows: WSL2 é necessário  
- **golang-migrate**  
  ```bash
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```  
- **Extensão REST Client** (opcional, para testes)

---

## ⚙️ Setup do Projeto

1. **Clone o repositório**  
   ```bash
   git clone https://github.com/devfullcycle/imersao22.git
   cd imersao22/go-gateway
   ```

2. **Configure as variáveis de ambiente**  
   ```bash
   cp .env.example .env
   ```

3. **Inicie o banco de dados**  
   ```bash
   docker compose up -d
   ```

4. **Execute as migrations**  
   ```bash
   migrate -path db/migrations \
     -database "postgresql://postgres:postgres@localhost:5432/gateway?sslmode=disable" up
   ```

5. **Execute a aplicação**  
   ```bash
   go run cmd/app/main.go
   ```

---

## 📡 API Endpoints

### Criar Conta

- **POST** `/accounts`  
- **Headers**:  
  ```http
  Content-Type: application/json
  ```
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@doe.com"
  }
  ```
- **Resposta**: dados da conta criada, incluindo o `api_key`.

---

### Consultar Conta

- **GET** `/accounts`  
- **Headers**:  
  ```http
  X-API-Key: {api_key}
  ```
- **Resposta**: dados da conta associada ao `api_key`.

---

### Criar Fatura

- **POST** `/invoice`  
- **Headers**:  
  ```http
  Content-Type: application/json
  X-API-Key: {api_key}
  ```
- **Body**:
  ```json
  {
    "amount": 100.50,
    "description": "Compra de produto",
    "payment_type": "credit_card",
    "card_number": "4111111111111111",
    "cvv": "123",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cardholder_name": "John Doe"
  }
  ```
- **Observação**: faturas acima de R$ 10.000 ficam pendentes para análise manual.

---

### Consultar Fatura

- **GET** `/invoice/{id}`  
- **Headers**:  
  ```http
  X-API-Key: {api_key}
  ```
- **Resposta**: dados da fatura específica.

---

### Listar Faturas

- **GET** `/invoice`  
- **Headers**:  
  ```http
  X-API-Key: {api_key}
  ```
- **Resposta**: lista de todas as faturas da conta.

---

## 🧪 Testando a API

O projeto inclui um arquivo **`test.http`** para uso com a extensão REST Client do VS Code. Ele contém:

- Variáveis globais pré-configuradas  
- Exemplos de todas as requisições  
- Captura automática do `api_key` após criação da conta  

### Para usar

1. Instale a extensão **REST Client** no VS Code  
2. Abra o arquivo `test.http`  
3. Clique em **Send Request** acima de cada requisição  

---
```
