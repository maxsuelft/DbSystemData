<div align="center">
  <img src="assets/logo.svg" alt="DbSystemData Logo" width="250"/>

  <h3>Ferramenta de backup para PostgreSQL, MySQL e MongoDB</h3>
  <p>DbSystemData é um fork focado em backend e frontend. Ferramenta gratuita, open source e self-hosted para backup de bancos de dados (com foco em PostgreSQL). Backups com múltiplos destinos (S3, Google Drive, FTP, etc.) e notificações (Slack, Discord, Telegram, etc.).</p>
  <p><em>Derivado do projeto <a href="https://github.com/databasus/databasus">Databasus</a>.</em></p>

  <!-- Badges -->
  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
  [![MySQL](https://img.shields.io/badge/MySQL-4479A1?logo=mysql&logoColor=white)](https://www.mysql.com/)
  [![MariaDB](https://img.shields.io/badge/MariaDB-003545?logo=mariadb&logoColor=white)](https://mariadb.org/)
  [![MongoDB](https://img.shields.io/badge/MongoDB-47A248?logo=mongodb&logoColor=white)](https://www.mongodb.com/)
  <br />
  [![Apache 2.0 License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
  [![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-lightgrey)](https://github.com/databasus/DbSystemData)
  [![Self Hosted](https://img.shields.io/badge/self--hosted-yes-brightgreen)](https://github.com/databasus/DbSystemData)
  [![Open Source](https://img.shields.io/badge/open%20source-❤️-red)](https://github.com/databasus/DbSystemData)

  <p>
    <a href="#-features">Features</a> •
    <a href="#-installation">Instalação</a> •
    <a href="#-usage">Uso</a> •
    <a href="#-license">Licença</a> •
    <a href="#-contributing">Contribuindo</a>
  </p>

  <img src="assets/dashboard-dark.svg" alt="DbSystemData Dark Dashboard" width="800" style="margin-bottom: 10px;"/>

  <img src="assets/dashboard.svg" alt="DbSystemData Dashboard" width="800"/>
</div>

---

## ✨ Features

### 💾 **Bancos suportados**

- **PostgreSQL**: 12, 13, 14, 15, 16, 17 e 18
- **MySQL**: 5.7, 8 e 9
- **MariaDB**: 10, 11 e 12
- **MongoDB**: 4, 5, 6, 7 e 8

### 🔄 **Backups agendados**

- **Agendamento flexível**: horário, diário, semanal, mensal ou cron
- **Horários precisos**: execute em horários específicos (ex.: 4h da manhã)
- **Compressão**: economia de espaço com compressão balanceada (~20% de overhead)

### 🗑️ **Políticas de retenção**

- **Período**: Manter backups por duração fixa (ex.: 7 dias, 3 meses, 1 ano)
- **Quantidade**: Manter um número fixo dos backups mais recentes (ex.: últimos 30)
- **GFS (Grandfather-Father-Son)**: Retenção em camadas — horário, diário, semanal, mensal e anual
- **Limites de tamanho**: Limites por backup e total de armazenamento

### 🗄️ **Múltiplos destinos de armazenamento**

- **Armazenamento local**: Manter backups no seu VPS/servidor
- **Nuvem**: S3, Cloudflare R2, Google Drive, NAS, Dropbox, SFTP, Rclone e outros
- **Seguro**: Seus dados permanecem sob seu controle

### 📱 **Notificações**

- **Múltiplos canais**: Email, Telegram, Slack, Discord, webhooks
- **Atualizações em tempo real**: Sucesso e falha
- **Integração**: Ideal para fluxos DevOps

### 🔒 **Segurança**

- **Criptografia AES-256-GCM**: Proteção para arquivos de backup
- **Armazenamento zero-trust**: Backups criptografados
- **Criptografia de segredos**: Dados sensíveis nunca expostos em logs
- **Usuário read-only**: Usa usuário somente leitura por padrão para backups

É importante que você consiga descriptografar e restaurar backups a partir dos storages (local, S3, etc.) sem o DbSystemData. Evitamos vendor lock-in.

### 👥 **Para times**

- **Workspaces**: Agrupe bancos, notifiers e storages por projeto ou time
- **Controle de acesso**: Permissões por função
- **Audit logs**: Rastreie atividades e alterações
- **Roles**: Viewer, member, admin ou owner

### 🎨 **Interface**

- **UI limpa**: Interface intuitiva
- **Temas claro e escuro**
- **Adaptável a mobile**

### ☁️ **Self-hosted e nuvem**

- **Nuvem**: AWS RDS, Google Cloud SQL, Azure Database for PostgreSQL
- **Self-hosted**: Qualquer instância PostgreSQL que você gerencie
- **Granularidade**: Backups horários e diários cobrem a maioria dos casos

### 🐳 **Self-hosted e seguro**

- **Docker**: Deploy e gerenciamento simples
- **Privacidade**: Dados na sua infraestrutura
- **Open source**: Licença Apache 2.0

### 📦 Instalação

Quatro formas de instalar: script automatizado (recomendado), Docker run, Docker Compose ou Kubernetes com Helm.

<img src="assets/healthchecks.svg" alt="Dashboard" width="800"/>

---

## 📦 Instalação

### Opção 1: Script de instalação (recomendado, Linux)

O script irá instalar Docker/Docker Compose (se necessário), configurar o DbSystemData e inicialização no boot.

```bash
sudo apt-get install -y curl && \
sudo curl -sSL https://raw.githubusercontent.com/databasus/DbSystemData/refs/heads/main/install-databasus.sh \
| sudo bash
```

*(Substitua a URL pelo repositório do seu fork se necessário.)*

### Opção 2: Docker run

Use a imagem publicada no GitHub Container Registry (após o workflow de build no seu repositório):

```bash
docker run -d \
  --name dbsystemdata \
  -p 4005:4005 \
  -v ./dbsystemdata-data:/dbsystemdata-data \
  --restart unless-stopped \
  ghcr.io/SEU_USUARIO/dbsystemdata:latest
```

- Substitua `SEU_USUARIO` pelo seu usuário do GitHub.
- Acesse em `http://localhost:4005`
- Dados em `./dbsystemdata-data`
- Reinício automático habilitado

### Opção 3: Docker Compose

1. Copie os arquivos de exemplo e defina seu usuário GitHub:

```bash
cp docker-compose.yml.example docker-compose.yml
cp .env.example .env
# Edite .env e defina: GITHUB_USER=seu_usuario_github
```

2. Suba o serviço:

```bash
docker compose up -d
```

A imagem usada será `ghcr.io/SEU_USUARIO/dbsystemdata:latest` (gerada pelo workflow em `.github/workflows/docker-build-push.yml` ao dar push em `main`/`master`).

**Build da imagem no GitHub:** o workflow `.github/workflows/docker-build-push.yml` faz build e push da imagem no push para `main` ou `master`. Para o build concluir, o repositório precisa ter os binários em `assets/tools/x64` e `assets/tools/arm` (PostgreSQL, MySQL, MariaDB); veja `assets/tools/README.md`. Se esses diretórios não existirem, o build falhará até você adicioná-los.

### Persistência de dados (Docker)

Para **não perder dados** ao apagar o container e subir de novo, o DbSystemData precisa gravar em um **volume**:

- **Docker Compose:** o `docker-compose.yml.example` já usa o volume nomeado `dbsystemdata-data`. Ao rodar `docker compose up -d`, os dados (banco interno PostgreSQL, Valkey, backups locais) ficam nesse volume e **persistem** quando você remove o container e sobe de novo. Não use `docker compose down -v` se quiser manter os dados (o `-v` apaga volumes nomeados).
- **Docker run:** use sempre `-v ./dbsystemdata-data:/dbsystemdata-data` (ou um volume nomeado). Sem o `-v`, tudo fica só dentro do container e é **perdido** ao remover o container.

Resumo: **com o volume configurado**, você pode apagar o container e subir de novo que usuários, bancos, storages e agendamentos continuam lá. Para usar uma pasta no host (ex. `./dbsystemdata-data`) em vez do volume nomeado, no `docker-compose.yml` troque `dbsystemdata-data:/dbsystemdata-data` por `./dbsystemdata-data:/dbsystemdata-data` e remova a seção `volumes:` no final do arquivo.

**Se a pasta do volume ficar vazia** (container sobe mas não aparecem arquivos em `pgdata`, `backups`, etc.):

1. **Crie a pasta no host antes de subir** (e garanta permissão de escrita):  
   `mkdir -p /caminho/para/data && chmod 755 /caminho/para/data`
2. **SELinux (RHEL/CentOS/Rocky)**  
   Adicione `:z` ao volume para o container poder escrever:  
   `- /www/wwwroot/docker/dbsystemdata/data:/dbsystemdata-data:z`
3. **Logs do container**  
   Verifique se o script de inicialização falhou:  
   `docker logs dbsystemdata`  
   Erros de "Permission denied" indicam problema de permissão ou SELinux.

### Opção 4: Kubernetes com Helm

Consulte o [README do chart Helm](deploy/helm/README.md) para instalação via OCI registry e opções (ClusterIP, LoadBalancer, Ingress).

---

## 🚀 Uso

1. **Acesse o dashboard**: `http://localhost:4005`
2. **Adicione seu primeiro banco**: Clique em "New Database" e siga o assistente
3. **Configure o agendamento**: Horário, diário, semanal, mensal ou cron
4. **Conexão**: Informe credenciais e detalhes do banco
5. **Escolha o storage**: Local, S3, Google Drive, etc.
6. **Política de retenção**: Período, quantidade ou GFS
7. **Notificações** (opcional): Email, Telegram, Slack, webhook
8. **Salve e inicie**: O DbSystemData validará e iniciará o agendamento

### 🔑 Redefinir senha

```bash
docker exec -it dbsystemdata ./main --new-password="SuaNovaSenhaSegura123" --email="admin"
```

Substitua `admin` pelo email do usuário.

### 💾 Backup do próprio DbSystemData

Após a instalação, recomenda-se fazer backup do próprio DbSystemData ou copiar a chave secreta de criptografia. Consulte a documentação no repositório.

---

## 📝 Licença

Este projeto está sob a licença Apache 2.0 — veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🤝 Contribuindo

Contribuições são bem-vindas. Abra issues e pull requests neste repositório.
