# Usar uma imagem base do Node.js
FROM node:18 as build

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos de dependência
COPY package.json package-lock.json ./

# Instalar as dependências
RUN npm install

# Copiar o restante do código
COPY . .

# Construir o projeto Angular
RUN npm run build -- --configuration production

# Usar uma imagem leve do Nginx para servir o frontend
FROM nginx:alpine

# Copiar os arquivos de build do Angular para o Nginx
COPY --from=build /app/dist/frontend-pet-app /usr/share/nginx/html

# Expor a porta 80
EXPOSE 80

# Iniciar o Nginx
CMD ["nginx", "-g", "daemon off;"]