# React PWA Boilerplate

Progressive Web Application (PWA) boilerplate yang dibangun dengan React, TypeScript, dan Vite. Template siap pakai untuk memulai project PWA dengan konfigurasi optimal.

## 📱 Spesifikasi PWA

### Fitur PWA
- ✅ **Installable** - Dapat diinstall di berbagai platform (Android, iOS, Windows, macOS)
- ✅ **Offline Support** - Bekerja tanpa koneksi internet dengan Service Worker
- ✅ **Background Sync** - Sinkronisasi data di background ketika koneksi kembali
- ✅ **Auto Update** - Otomatis update ke versi terbaru
- ✅ **Fast Loading** - Pre-caching assets untuk performa optimal
- ✅ **Responsive** - Mendukung berbagai ukuran layar

### Platform Support
- **Android** - 6 ukuran icon (48px - 512px)
- **iOS** - 21 ukuran icon (16px - 1024px) 
- **Windows 11** - 95 varian icon (SmallTile, Square, Wide, Large, StoreLogo, SplashScreen)
- **Desktop** - Chrome, Edge, Safari

### Manifest Configuration
```json
{
  "name": "Your App Name",
  "short_name": "App Name",
  "description": "Your App Description",
  "theme_color": "#E6E6FA",
  "background_color": "#E6E6FA",
  "start_url": "/",
  "display": "standalone",
  "orientation": "portrait"
}
```

## 🛠️ Tech Stack

### Core Dependencies

| Library | Version | Keterangan |
|---------|---------|------------|
| **React** | ^19.2.0 | UI Framework |
| **React DOM** | ^19.2.0 | React renderer untuk web |
| **TypeScript** | ~5.9.3 | Type-safe JavaScript |
| **Vite** | ^7.1.9 | Build tool & dev server |

### Routing
| Library | Version | Keterangan |
|---------|---------|------------|
| **React Router** | ^7.2.0 | Client-side routing |
| **React Router DOM** | ^7.2.0 | DOM bindings untuk React Router |

### Styling
| Library | Version | Keterangan |
|---------|---------|------------|
| **Tailwind CSS** | ^4.0.9 | Utility-first CSS framework |
| **@tailwindcss/vite** | ^4.0.9 | Vite plugin untuk Tailwind v4 |
| **tailwindcss-animate** | ^1.0.7 | Animation utilities |
| **clsx** | ^2.1.1 | Conditional className utility |
| **tailwind-merge** | ^3.1.0 | Merge Tailwind classes |

### PWA & Service Worker
| Library | Version | Keterangan |
|---------|---------|------------|
| **vite-plugin-pwa** | ^1.1.0 | Vite plugin untuk PWA |
| **workbox-core** | ^7.3.0 | Core Workbox library |
| **workbox-precaching** | ^7.3.0 | Pre-caching strategies |
| **workbox-routing** | ^7.3.0 | Request routing |
| **workbox-strategies** | ^7.3.0 | Caching strategies |
| **workbox-window** | ^7.3.0 | Window helper untuk SW |

### Development Tools
| Library | Version | Keterangan |
|---------|---------|------------|
| **ESLint** | ^9.37.0 | Linter |
| **TypeScript ESLint** | ^8.45.0 | TypeScript linting rules |
| **Prettier** | ^3.5.3 | Code formatter |
| **prettier-plugin-tailwindcss** | ^0.6.11 | Tailwind class sorting |
| **@vite-pwa/assets-generator** | ^1.0.2 | Generate PWA assets |

## 📁 Struktur Folder

```
react-pwa-boilerplate/
├── public/                      # Static assets
│   ├── android/                 # Android launcher icons
│   ├── ios/                     # iOS app icons
│   └── windows11/              # Windows 11 tiles
├── src/
│   ├── assets/                 # Dynamic assets
│   ├── lib/                    # Utility functions
│   │   └── utils.ts           # Helper utilities
│   ├── pages/                  # Page components
│   │   └── Home.tsx           # Home page
│   ├── App.tsx                 # Main app component
│   ├── PWABadge.tsx           # PWA update notification
│   ├── main.tsx               # Entry point
│   ├── sw.ts                  # Service Worker
│   └── index.css              # Global styles
├── vite.config.ts             # Vite configuration
├── tsconfig.json              # TypeScript config
└── package.json               # Dependencies
```

## 🚀 Instalasi & Development

### Prerequisites
- Node.js >= 18.x
- Bun >= 1.1.13 (atau npm/yarn/pnpm)

### Install Dependencies
```bash
bun install
# atau
npm install
```

### Development Server
```bash
bun run dev
# atau
npm run dev
```
Server akan berjalan di `http://localhost:5173`

### Build untuk Production
```bash
bun run build
# atau
npm run build
```

Output akan berada di folder `dist/`

### Preview Production Build
```bash
bun run preview
# atau
npm run preview
```

## 🔧 Konfigurasi PWA

### Service Worker Strategy
Menggunakan **InjectManifest** strategy dengan custom Service Worker (`src/sw.ts`):

- **Pre-caching**: Assets statis (JS, CSS, HTML, images)
- **Runtime Caching**:
  - Images: CacheFirst
  - Fonts: CacheFirst
  - API Calls: NetworkFirst
  - Navigation: NetworkFirst
- **Background Sync**: POST/PATCH requests dengan retry queue

### Workbox Configuration
```typescript
workbox: {
  globPatterns: ["**/*.{js,css,html,ico,png,svg,json,woff,woff2}"],
  runtimeCaching: [
    {
      urlPattern: /^https:\/\/fonts\.googleapis\.com\/.*/i,
      handler: "CacheFirst",
      options: {
        cacheName: "google-fonts-cache",
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 365, // 1 year
        },
      },
    },
    // ... more caching rules
  ],
}
```

## 🎨 Theme & Styling

### Color Palette (Default Theme)

#### Light Mode
- **Primary**: `#0077B6` - Biru utama
- **Secondary**: `#00B4D8` - Biru muda  
- **Accent**: `#0096C7` - Aksen gradasi tengah
- **Background**: `#E3F2FD` - Putih kebiruan

#### Dark Mode
- **Primary**: `#00B4D8` - Biru muda
- **Secondary**: `#023E8A` - Biru gelap
- **Accent**: `#0077B6` - Biru utama
- **Background**: `#0A0A0A` - Hitam

> **Note**: Customize color palette di `src/index.css` sesuai brand Anda

### Font
- **Manrope** - Google Fonts (200-800 weight)
- Ganti font di `index.html` dan `src/index.css` sesuai kebutuhan

## 🐳 Docker Deployment

### Build Docker Image
```bash
docker build -t my-pwa-app \
  --build-arg VITE_MODE=production \
  --build-arg VITE_API_URL=https://api.example.com .
```

### Run Container
```bash
docker run -d -p 80:80 my-pwa-app
```

### Docker Features
- **Multi-stage build**: Builder stage + nginx stage
- **Brotli compression**: Pre-compress assets dengan brotli
- **Nginx optimization**: Static file caching, gzip fallback
- **Small image size**: Alpine-based images

## 📦 Icon Assets

Total **122 icon variants** untuk berbagai platform:

### Android (6 icons)
- 512x512, 192x192, 144x144, 96x96, 72x72, 48x48

### iOS (21 icons)
- 1024x1024, 512x512, 256x256, 192x192, 180x180, 167x167, 152x152, 144x144, 128x128, 120x120, 114x114, 100x100, 87x87, 80x80, 76x76, 72x72, 64x64, 60x60, 58x58, 57x57, 50x50, 40x40, 32x32, 29x29, 20x20, 16x16

### Windows 11 (95 icons)
- SmallTile (5 scales)
- Square150x150Logo (5 scales)
- Wide310x150Logo (5 scales)
- LargeTile (5 scales)
- Square44x44Logo (5 scales)
- StoreLogo (5 scales)
- SplashScreen (5 scales)
- Target size variants (16-256px)
- Unplated variants
- Light unplated variants

## 🔐 Environment Variables

Buat file `.env` di root folder:

```env
VITE_MODE=development
VITE_API_URL=http://localhost:3000
```

## 📝 Scripts

| Command | Deskripsi |
|---------|-----------|
| `bun run dev` | Start development server |
| `bun run build` | Build untuk production |
| `bun run preview` | Preview production build |
| `bun run lint` | Run ESLint |

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open Pull Request

## 📄 License

MIT License - feel free to use this boilerplate for your projects

## 🌟 Features to Add

Boilerplate ini siap untuk dikembangkan dengan fitur-fitur seperti:
- Authentication (Firebase, Supabase, Auth0)
- State Management (Zustand, Redux, Jotai)
- Database (IndexedDB, Dexie.js)
- Push Notifications
- Biometric Authentication
- Camera & Media Access
- Geolocation
- Payment Integration

---

**Note**: Pastikan semua icon assets tersedia di folder `public/android/`, `public/ios/`, dan `public/windows11/` sebelum build production. Gunakan tool seperti [PWA Asset Generator](https://github.com/vite-pwa/assets-generator) untuk generate icons otomatis.