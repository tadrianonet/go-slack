# Slack Bot - Envio de Mensagens com Go

Este é um projeto Go que utiliza a API do Slack para enviar mensagens a um canal específico. 

## Configuração
Para enviar mensagens para o Slack voce precisa criar um token e ter o ID do seu canal. Para obter essas informacoes, basta seguir os passos abaixo.

### Instruções para Configurar o Bot no Slack

Este guia fornece um passo a passo para criar um app no Slack e configurar as credenciais necessárias (`SLACK_BOT_TOKEN` e `SLACK_CHANNEL_ID`) para que o bot possa enviar mensagens aos canais.

## 1. Criando um Slack App para Obter o `SLACK_BOT_TOKEN`

1. **Acesse o Slack API**: Visite [https://api.slack.com/](https://api.slack.com/).
   
2. **Crie um novo App**:
   - Clique em **Create an App**.
   - Selecione **From scratch**.
   - Escolha o workspace onde o app será instalado e dê um nome ao app.

3. **Configurar Permissões do Bot**:
   - No menu do lado esquerdo, vá para **OAuth & Permissions**.
   - Em **Scopes**, na seção **Bot Token Scopes**, adicione as permissões que seu bot precisará. As permissões comuns para envio de mensagens são:
     - `chat:write`: para enviar mensagens aos canais.
     - `channels:read`: para ler informações de canais públicos (opcional, mas útil para identificar canais).

4. **Instalar o App**:
   - No topo da página de **OAuth & Permissions**, clique em **Install to Workspace** e autorize o app.
   - Após a instalação, você verá um **Bot User OAuth Token**. Esse é o seu `SLACK_BOT_TOKEN`.


Agora basta pegar o ID da URL e enviar as suas mensagens.

```plaintext
    https://app.../client/TXXXXXXX/CYYYYYYYYY
 ```

Neste exemplo, o SLACK_CHANNEL_ID é CYYYYYYYYY

## Debug

Para testar o projeto, bastar abrir um terminal e executar o comando go run main.go para enviar a mensagem.