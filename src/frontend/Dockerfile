FROM node:20-slim as build

# Update and install the latest dependencies
RUN apt-get update && apt-get upgrade -y

# Set work dir as app
WORKDIR /app
# Copy the nuxt project package json and package json lock if available 
COPY package* ./
# Install pnpm
RUN npm install -g pnpm
# Copy all other project files to working directory
COPY . ./
# Build the nuxt project to generate the artifacts in .output directory
RUN pnpm install
RUN pnpm run build

# Second stage build
FROM node:20-slim

# Update packages and create non-root user
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y dumb-init && \
    groupadd -r nuxtuser && useradd -r -g nuxtuser nuxtuser

# Set work dir as app
WORKDIR /app

# Copy the output directory from build stage
COPY --chown=nuxtuser:nuxtuser --from=build /app/.output ./

# Set non root user
USER nuxtuser

# Expose 8080 on container
EXPOSE 3000

# Set app host and port
ENV HOST=0.0.0.0 PORT=3000 NODE_ENV=production

# Start the app with dumb-init
CMD ["dumb-init", "node", "/app/server/index.mjs"]