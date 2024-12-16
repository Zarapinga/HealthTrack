# HealthTrack

**HealthTrack** é um sistema web desenvolvido para modernizar a gestão de clínicas médicas, centralizando o gerenciamento de prontuários, agendamentos e receitas. O objetivo do projeto é otimizar os processos internos da clínica, reduzir o uso de papel e melhorar a experiência de pacientes e profissionais da saúde.

---

## Funcionalidades

- **Pacientes**:
  - Visualização de consultas passadas e futuras.
  - Acesso a diagnósticos e receitas médicas.
  - Agendamento de consultas com visualização dos horários disponíveis.

- **Secretárias**:
  - Gerenciamento completo dos dados de pacientes e médicos.
  - Agendamento de consultas para qualquer paciente.

- **Médicos**:
  - Acesso aos prontuários de pacientes com quem têm ou tiveram consultas.
  - Emissão de diagnósticos e receitas médicas.

- **Super-Usuário**:
  - Controle total do sistema, incluindo o cadastro de médicos e secretárias.

---
## Boas Práticas de Desenvolvimento

**Comentários e Documentação:**

- Código Autoexplicativo: Escrever Código claro e compreensível. Usar comentários para explicar partes complexas quando necessário.
- Documentar Métodos e Funções: Descrever o propósito, os parâmetros e os valores de retorno de métodos e funções de forma clara.

**Código Limpo (Clean Code):**

- Usar Nomes Significativos: Usar nomes descritivos para variáveis, métodos e classes.
- Escrever Funções Curtas e Objetivas: Fazer com que cada Função deve realizar uma única tarefa, facilitando a legibilidade e manutenção.

**Responsabilidade Única:**

- Funções com Responsabilidade Única: A função deve focar em uma tarefa específica ou um grupo relacionado de tarefas.

**Formatação e Organização:**

- Consistência: Manter uma formatação consistente no código.
- Organização Lógica: Agrupar funções e variáveis relacionadas para melhorar a clareza.

**Tratamento de Erros:**

- Exceções: Usar exceções para erros excepcionais, não para controle de fluxo.

**Testabilidade:**

- Código Testável: Separar a lógica de negócios das dependências para facilitar o teste.

**Princípios SOLID:**

- SRP (Single Responsibility Principle): A classe deve ter uma única responsabilidade.
- OCP (Open/Closed Principle): O código deve ser extensível sem modificar o existente.
- LSP (Liskov Substitution Principle): Subtipos devem ser substituíveis por seus tipos base.
ISP (Interface Segregation Principle): Usar várias interfaces específicas em vez de uma única geral.
-DIP (Dependency Inversion Principle): Depender de abstrações, não de implementações concretas.

Refatoração:

- Melhoria Contínua: Refatorar regularmente para melhorar a estrutura sem mudar o comportamento.
- Evitar Código Duplicado (DRY): Não repitir códigos para reduzir erros e facilitar manutenção.

**Simplicidade:**

- KISS (Keep It Simple, Stupid): Manter o código simples e evite complexidade desnecessária.

**Controle de Versionamento:**

- Branch main: FAZER alterações na branch main e use commits com mensagens claras.

**Testes Automatizados:**

- Cobertura de Testes: Garantir que novas funcionalidades sejam cobertas por testes automatizados.
- Testes Independentes: Testes devem ser pequenos e verificar um comportamento específico por vez, usando frameworks apropriados.

---
## Regras e Padrões de uso do Git


- **Commits**

- Use mensagens de commit claras e diretas, preferencialmente no gerúndio
- Cada commit deve ser atômico, representando uma única alteração lógica.
- Sempre vincule os commits às issues correspondentes no backlog.


- **Branches**

- Utilize branches para os códigos.
- Crie branches dedicadas para correções críticas.
- Nomeie as branches com letras minúsculas apenas.

- **Organização**
  
- Mantenha a estrutura do projeto organizada, separando claramente a documentação do código.

---

## Tecnologias Utilizadas

### **Frontend**
- **Linguagem**: JavaScript (React.js)  
- **Versão**: React 18.2.0  
- **Gerenciamento de Pacotes**: NPM 9.x  

### **Backend**
- **Linguagem**: GO(Golang)  
- **Versão**: 2.1.x    

### **Banco de Dados**
- **Sistema Gerenciador**: MySQL  
- **Versão**: 8.0  

### **Servidor**
- **Servidor de Desenvolvimento**: Node.js (Express)  
- **Servidor de Produção**: [Especifique aqui o serviço, ex.: Apache/Nginx, ou serviço na nuvem, ex.: Heroku/AWS]  

### **Outras Tecnologias**
- **Autenticação**: JSON Web Tokens (JWT)  
- **Criptografia**: bcrypt 5.x  
- **Comunicação API**: RESTful  
- **Documentação da API**: Swagger (opcional)  

---

## Instalação e Execução

### **Requisitos**
1. **Node.js**: 18.x ou superior  
2. **MySQL**: 8.0 ou superior  
3. **NPM**: 9.x ou superior  

### **Passos para Configuração**
1. Clone este repositório:
   ```bash
   git clone https://github.com/Zarapinga/HealthTrack.git
   cd HealthTrack
