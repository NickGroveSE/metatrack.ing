# Frontend Dockerfile for SSR
# Place this file as: frontend/Dockerfile

FROM node:20-alpine

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Expose port
EXPOSE 4321

# Set environment variables
ENV HOST=0.0.0.0
ENV PORT=4321

# Start the Node.js server with sh to ensure env vars are used
CMD ["sh", "-c", "HOST=0.0.0.0 PORT=4321 node ./dist/server/entry.mjs"]