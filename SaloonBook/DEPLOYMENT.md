# Quick Deployment Guide

This guide shows how to deploy your frontend-only SaloonBook app to various platforms.

## Prerequisites

Build your app first:
```bash
cd frontend/web
npm install
npm run build
```

The output will be in `frontend/web/dist/`

---

## Vercel (Recommended - Easiest)

### One-Command Deploy
```bash
cd frontend/web
npx vercel --prod
```

### Or with Vercel CLI
```bash
npm install -g vercel
cd frontend/web
vercel --prod
```

### Auto-Deploy from Git
1. Push your code to GitHub
2. Go to [vercel.com](https://vercel.com)
3. Click "New Project"
4. Import your repository
5. Set:
   - **Framework Preset**: Vite
   - **Root Directory**: `frontend/web`
   - **Build Command**: `npm run build`
   - **Output Directory**: `dist`
6. Click "Deploy"

**Done!** Auto-deploys on every push.

---

## Netlify

### Drag & Drop (Easiest)
1. Go to [netlify.com](https://netlify.com)
2. Drag `frontend/web/dist` folder onto the page
3. Done!

### CLI Deploy
```bash
npm install -g netlify-cli
cd frontend/web
npm run build
netlify deploy --prod --dir=dist
```

### Auto-Deploy from Git
1. Push code to GitHub
2. Go to [netlify.com](https://netlify.com)
3. Click "Add new site" → "Import an existing project"
4. Connect GitHub repository
5. Set:
   - **Base directory**: `frontend/web`
   - **Build command**: `npm run build`
   - **Publish directory**: `frontend/web/dist`
6. Click "Deploy"

---

## GitHub Pages

### Using gh-pages package
```bash
cd frontend/web
npm install -g gh-pages

# Build and deploy
npm run build
npx gh-pages -d dist
```

### Manual
```bash
cd frontend/web
npm run build

# Copy dist to root or gh-pages branch
git checkout -b gh-pages
cp -r dist/* .
git add .
git commit -m "Deploy"
git push origin gh-pages
```

Enable in repo settings: Settings → Pages → Source: gh-pages branch

---

## Cloudflare Pages

### From Git (Recommended)
1. Push code to GitHub
2. Go to [pages.cloudflare.com](https://pages.cloudflare.com)
3. Click "Create a project"
4. Connect GitHub repository
5. Set:
   - **Framework preset**: Vite
   - **Build command**: `cd frontend/web && npm run build`
   - **Build output directory**: `frontend/web/dist`
6. Click "Save and Deploy"

### Direct Upload
```bash
cd frontend/web
npm run build

# Upload dist folder via Cloudflare dashboard
```

---

## Firebase Hosting

### Setup
```bash
npm install -g firebase-tools
firebase login
firebase init hosting
```

When prompted:
- **Public directory**: `frontend/web/dist`
- **Single-page app**: Yes
- **Overwrite index.html**: No

### Deploy
```bash
cd frontend/web
npm run build
firebase deploy --only hosting
```

---

## Render

1. Push code to GitHub
2. Go to [render.com](https://render.com)
3. Click "New" → "Static Site"
4. Connect repository
5. Set:
   - **Build Command**: `cd frontend/web && npm install && npm run build`
   - **Publish Directory**: `frontend/web/dist`
6. Click "Create Static Site"

---

## Railway

1. Push code to GitHub
2. Go to [railway.app](https://railway.app)
3. Click "New Project" → "Deploy from GitHub repo"
4. Select repository
5. Add build settings:
   - **Build Command**: `cd frontend/web && npm install && npm run build`
   - **Start Command**: `npx serve frontend/web/dist`
6. Deploy!

---

## Surge.sh

### Quick Deploy
```bash
npm install -g surge
cd frontend/web
npm run build
cd dist
surge
```

Follow prompts to create account and deploy.

---

## Custom Server (VPS/Dedicated)

### Using Nginx

1. Build locally:
```bash
cd frontend/web
npm run build
```

2. Copy `dist/` to server:
```bash
scp -r dist/* user@server:/var/www/saloonbook/
```

3. Nginx config:
```nginx
server {
    listen 80;
    server_name your-domain.com;
    root /var/www/saloonbook;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

4. Restart Nginx:
```bash
sudo systemctl restart nginx
```

### Using Docker

The project includes a `docker-compose.yml`:

```bash
cd frontend/web
npm run build
cd ../..
docker compose up -d
```

Access at: http://localhost:8080

---

## Apache Server

1. Copy `dist/` files to web root:
```bash
cp -r dist/* /var/www/html/saloonbook/
```

2. Create `.htaccess` in the directory:
```apache
<IfModule mod_rewrite.c>
  RewriteEngine On
  RewriteBase /saloonbook/
  RewriteRule ^index\.html$ - [L]
  RewriteCond %{REQUEST_FILENAME} !-f
  RewriteCond %{REQUEST_FILENAME} !-d
  RewriteRule . /saloonbook/index.html [L]
</IfModule>
```

---

## AWS S3 + CloudFront

### S3 Bucket
```bash
cd frontend/web
npm run build

# Upload to S3
aws s3 sync dist/ s3://your-bucket-name/ --delete
```

### Enable Static Website Hosting
1. Go to S3 bucket properties
2. Enable "Static website hosting"
3. Index document: `index.html`
4. Error document: `index.html` (for SPA routing)

### CloudFront (Optional, for CDN)
1. Create CloudFront distribution
2. Origin: Your S3 bucket
3. Default root object: `index.html`
4. Error pages: Redirect 404 to `/index.html`

---

## Environment-Specific Builds

### Add environment variables (optional)

Create `.env.production`:
```bash
VITE_APP_NAME=SaloonBook
VITE_API_URL=https://api.yourdomain.com
```

Build:
```bash
npm run build
```

---

## Custom Domain

Most platforms support custom domains:

1. **Vercel/Netlify/Cloudflare**: 
   - Go to project settings
   - Add custom domain
   - Update DNS records

2. **DNS Settings**:
   ```
   Type: CNAME
   Name: @ or www
   Value: your-project.vercel.app
   ```

---

## Post-Deployment Checklist

✅ Test all pages load correctly  
✅ Test service selection flow  
✅ Create a test booking  
✅ Verify LocalStorage persists data  
✅ Test on mobile devices  
✅ Check browser console for errors  
✅ Test QR code generation  
✅ Verify responsive design  

---

## Performance Tips

1. **Compression**: Most hosts enable gzip/brotli automatically
2. **Caching**: Set cache headers for static assets
3. **CDN**: Use Cloudflare or similar for global distribution
4. **Lighthouse**: Run audit to check performance

---

## Troubleshooting

### Issue: 404 on Refresh
**Solution**: Configure server to serve `index.html` for all routes (SPA routing)

### Issue: Assets Not Loading
**Solution**: Check `base` in `vite.config.ts` if not deploying to root

### Issue: Build Fails
**Solution**: 
```bash
rm -rf node_modules package-lock.json
npm install
npm run build
```

---

## Quick Reference

| Platform | Command | Time |
|----------|---------|------|
| Vercel | `npx vercel --prod` | ~2 min |
| Netlify | `netlify deploy --prod` | ~2 min |
| Surge | `cd dist && surge` | ~1 min |
| Firebase | `firebase deploy` | ~3 min |
| GitHub Pages | `npx gh-pages -d dist` | ~3 min |

---

**Need help?** Check platform-specific docs or open an issue!

