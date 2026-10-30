# 🪔 AstroDiwali · AstroEngine & Prosperity Platform

[![CI - Lint & Build](https://github.com/your-username/astrodiwali/actions/workflows/ci.yml/badge.svg)](https://github.com/your-username/astrodiwali/actions/workflows/ci.yml)
[![Deploy to GitHub Pages](https://github.com/your-username/astrodiwali/actions/workflows/deploy-pages.yml/badge.svg)](https://github.com/your-username/astrodiwali/actions/workflows/deploy-pages.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-amber.svg)](https://opensource.org/licenses/MIT)
[![Node.js](https://img.shields.io/badge/Node.js-20.x-green.svg)](https://nodejs.org/)
[![React](https://img.shields.io/badge/React-19-blue.svg)](https://react.dev/)
[![Vite](https://img.shields.io/badge/Vite-6.x-purple.svg)](https://vitejs.dev/)
[![TailwindCSS](https://img.shields.io/badge/TailwindCSS-v4-06B6D4.svg)](https://tailwindcss.com/)

An enterprise-grade, high-performance Vedic AstroEngine and Festival Prosperity Platform. Features high-precision planetary calculations, symbolic contradiction resolution, Birth Time Rectification (BTR), Remedial Vector Matrix (RVM), and real-time Diwali Panchang countdown in an imperial dark maroon and antique gold aesthetic.

---

## 🌟 Highlights & Core Architecture

### 1. 🪔 Live Diwali Countdown & Vedic Panchang Cosmic Clock
- **Real-Time Astronomical Clock**: Live 1-second countdown to the next upcoming Diwali (Lakshmi Puja on Kartik Amavasya).
- **Vedic Panchang Metrics**: Displays Tithi, Paksha, Nakshatra, and Shubh Muhurta (*Pradosh Kaal & Vrishabha Sthira Lagna*).
- **5-Day Festival Sequence**: One-click countdown switcher across **Dhanteras**, **Narak Chaturdashi / Choti Diwali**, **Diwali (Lakshmi Puja)**, **Govardhan Puja**, and **Bhai Dooj**.
- **Natal Lagna Synchronization**: Real-time cross-referencing between user's natal lagna and the Puja Muhurta.
- **Instant Export**: Copy verified Muhurta details to clipboard.

### 2. ⚡ 5-Screen Conversion Life Cycle & Neuro-Design
- **Screen 1 — Social Acquisition Hook**: High-converting 9:16 kinetic Instagram-style ad card with animated typography tickers.
- **Screen 2 — Smart Onboarding**: 3-step progressive ephemeris input (Date/Time, Coordinates, True Lahiri / KP / Raman Ayanamsas).
- **Screen 3 — Kinetic Calculation Stream**: Real-time telemetry feed tracking planetary calculations, divisional harmonic matrix synthesis, and contradiction resolution.
- **Screen 4 — Curiosity-Gap Paywall**: Gaussian-blurred planetary vector preview with countdown timer and 1-Click checkout simulation.
- **Screen 5 — Production Astronomical Dashboard (C1–C6)**:
  - **C1: Native Identity Card**: Lagna, Moon Nakshatra, Ayanamsa, optimal Muhurta, and AI Desha-Kala-Patra verdict.
  - **C2: Chrono-Weather Dial**: 24-hour circular solar intersect mapping Rahu Kaal, Abhijit Muhurta, and Amrit Kaal with a live clock needle.
  - **C3: Event Matrix Strip**: Horizontal festival timeline with auspiciousness scores.
  - **C4: 365-Day High-Density Heatmap**: GitHub-style interactive yearly grid visualizing planetary transit auspiciousness scores.
  - **C5: Sensory Remedial Suite**: Acoustic 432 Hz frequencies, Kinetic yantras, Temporal timings, and Behavioral actions.
  - **C6: Calibration Metadata**: Julian Day calculations, execution benchmarks (<12ms), and cryptographic transaction verification.

### 3. 🪐 AstroEngine Computational Core
- **Divisional Charts**: High-precision calculations for D-1 (Rasi), D-2 (Hora), D-9 (Navamsa), and D-60 (Shashtiamsa).
- **Vimshottari Dasha Engine**: Maha-dasha and Antar-dasha progression calculations.
- **Symbolic Contradiction Resolver**: Resolves conflicting astrological yogas (e.g., Kemadruma vs. Gajakesari) via weighted graph synthesis.
- **Birth Time Rectification (BTR)**: Fine-grained ±15-minute life milestone calibration using Tata/BTR convergence algorithms.
- **Remedial Vector Matrix (RVM)**: Multi-modal remedies (Acoustic, Kinetic, Temporal, Behavioral).

### 4. 📺 Living Room & Ambient Display
- **TV Dashboard**: High-contrast, ambient 10-foot UI designed for living room screens.
- **Interactive Aarti**: Real-time bell, conch, and diya animations with synchronized audio.
- **Vedic Housie / Tambola**: Interactive multiplayer family game for festival gatherings.

### 5. 🔒 Transactional Backend & Webhooks
- `server.ts` Express tier with crypto HMAC SHA-256 webhook verification.
- Server-Sent Events (SSE) `/api/payment/sse/:token` for instant zero-latency dashboard paywall unlock.
- Ready-to-deploy SQL schemas (`schema.sql`) and Go microservice blueprints (`normalization_engine.go`, `btr_processor.go`, `graph_hydrator.go`, `conflict_resolver.go`, `rvm_engine.go`, `transactional_webhook.go`).

---

## 🛠️ Tech Stack

- **Frontend**: React 19, TypeScript, Tailwind CSS v4, Lucide React, Canvas Confetti, Motion
- **Backend**: Node.js, Express, TypeScript (`tsx`)
- **Build Tool**: Vite 6+
- **Containerization**: Multi-stage Dockerfile, Docker Compose
- **CI/CD**: GitHub Actions (Lint, Build, Automated GitHub Pages Deployment, Docker verification)

---

## 🚀 Quick Start (Local Development)

### Prerequisites
- [Node.js](https://nodejs.org/) (v20 or higher recommended)
- `npm` (v9 or higher)

### Installation

```bash
# 1. Clone repository
git clone https://github.com/your-username/astrodiwali.git
cd astrodiwali

# 2. Install dependencies
npm install

# 3. Setup environment variables (optional for AI features)
cp .env.example .env

# 4. Start the full-stack development server
npm run dev
```

Open your browser at **`http://localhost:3000`**.

---

## 🚢 Deployment Guides

### Option 1: GitHub Pages (Automated Zero-Config)
This repository includes a pre-configured GitHub Actions workflow (`.github/workflows/deploy-pages.yml`):
1. Push this repository to GitHub.
2. In your GitHub repository, navigate to **Settings** → **Pages**.
3. Under **Build and deployment** → **Source**, select **GitHub Actions**.
4. Every push to `main` will automatically build and deploy your live site!

### Option 2: Docker & Docker Compose
Run the entire production containerized application with a single command:

```bash
# Build and run container
docker compose up --build

# Or with Docker directly:
docker build -t astrodiwali .
docker run -p 3000:3000 -e PORT=3000 astrodiwali
```

### Option 3: Google Cloud Run / AWS ECS / VPS
Build the Docker image and deploy to any container platform:
```bash
docker build -t gcr.io/<PROJECT_ID>/astrodiwali:latest .
docker push gcr.io/<PROJECT_ID>/astrodiwali:latest
```

### Option 4: Vercel / Netlify / Render
- **Build Command**: `npm run build`
- **Output Directory**: `dist`
- **Install Command**: `npm install`

---

## 📁 Repository Structure

```
.
├── .github/
│   └── workflows/
│       ├── ci.yml                 # Automated CI: Lint & build test
│       ├── deploy-pages.yml       # Automated GitHub Pages deployment
│       └── docker-build.yml       # Docker build verification
├── src/
│   ├── components/
│   │   ├── DiwaliCountdown.tsx     # 🪔 Live Diwali Countdown & Panchang clock
│   │   ├── ConversionFunnelView.tsx# ⚡ 5-Screen Acquisition & Neuro-Design
│   │   ├── DashboardView.tsx       # 🪐 Astronomical Dashboard (C1-C6)
│   │   ├── Navigation.tsx          # Top navigation bar
│   │   ├── DiwaliProsperityKit.tsx # ₹99 Festive Prosperity kit
│   │   ├── RemediationView.tsx     # RVM Remedial suite
│   │   ├── LivingRoomTvDashboard.tsx# 10-foot TV ambient dashboard
│   │   ├── FamilyGameNight.tsx     # Aarti & Vedic Housie
│   │   ├── UatAndStressSuite.tsx   # DevOps benchmark suite
│   │   ├── ConflictResolverModal.tsx# Graph contradiction modal
│   │   ├── BtrModal.tsx            # Birth Time Rectification modal
│   │   └── CheckoutModal.tsx       # 1-Click payment modal
│   ├── engine/
│   │   ├── astroMath.ts            # Swiss Ephemeris & astronomical math
│   │   ├── conflictResolver.ts     # Graph conflict resolver
│   │   └── rvmEngine.ts            # Remedial Vector Matrix engine
│   ├── App.tsx                     # Root application container
│   ├── main.tsx                    # React DOM entry
│   └── index.css                   # Global styles & Dark Maroon theme
├── server.ts                       # Express backend & payment webhook SSE tier
├── schema.sql                      # PostgreSQL schema & payment transactions
├── Dockerfile                      # Production multi-stage Docker build
├── docker-compose.yml              # Local container orchestration
├── .dockerignore                   # Docker build exclusions
├── .gitignore                      # Git exclusions
├── vite.config.ts                  # Vite configuration with relative base support
├── package.json                    # Dependencies & scripts
└── LICENSE                         # MIT License
```

---

## 🔑 Environment Variables

Refer to `.env.example`:

| Variable | Description | Default / Example |
|---|---|---|
| `PORT` | Backend server port | `3000` |
| `GEMINI_API_KEY` | Optional Google Gemini API key for Desha-Kala-Patra AI | `""` |
| `APP_URL` | Base host URL | `http://localhost:3000` |

---

## 📜 Available Scripts

- `npm run dev`: Starts the full-stack server with live reloading via `tsx server.ts`
- `npm run build`: Builds the production Vite SPA bundle to `/dist`
- `npm run lint`: Runs TypeScript validation (`tsc --noEmit`)
- `npm run preview`: Previews the production build locally
- `npm run start`: Starts the production server

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
