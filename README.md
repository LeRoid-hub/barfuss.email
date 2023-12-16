# barfuss-website

## useful links
- CI https://www.programonaut.com/how-to-deploy-a-docker-image-to-a-server-using-github-actions/#action
  
## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Build Docker
```sh
docker build . -t barfuss-website
docker run -d -p 8080:80 barfuss-website
```
