# Etapa 1: Build da aplicação Angular
FROM node:18 AS build

# Definir o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos de dependências
COPY package.json package-lock.json ./

# Instalar as dependências
RUN npm install

# Copiar o restante dos arquivos do projeto
COPY . .

# Construir a aplicação Angular para produção
RUN npm run build -- --configuration production

# Etapa 2: Configuração do Nginx para servir a aplicação
FROM nginx:alpine

# Remover a página padrão do Nginx
RUN rm -rf /usr/share/nginx/html/*

# Copiar os arquivos de build do Angular para o diretório padrão do Nginx
COPY --from=build /app/dist/frontend /usr/share/nginx/html

# Copiar o arquivo de configuração personalizado para o Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expor a porta 80
EXPOSE 80

# Comando para iniciar o Nginx
CMD ["nginx", "-g", "daemon off;"]