# RealQuick

RealQuick is a web-based scheduling platform that delivers an intuitive experience by quickly generating customizable schedules, seamlessly syncing with 3rd party Calendars, and providing accessibility and responsiveness for users.

## Overview

RealQuick aims to be the default personal scheduling tool for college and university students around the globe. It offers:

- Dynamic schedule generation based on user preferences and fixed timeblocks
- Seamless integration with third-party calendar tools
- Flexible schedule display formats (week and timesheet views)
- Analytics dashboard with interpreted user data
- Weather integration for enhanced user experience
- Secure and persistent data storage

## Tech Stack

RealQuick frontend is built with modern and robust technologies:

- **Framework**: [Nuxt.js 3](https://nuxt.com/) - Vue.js framework for building modern web applications
- **UI Library**: [Vuetify](https://vuetifyjs.com/) - Material Design component framework
- **State Management**: [Pinia](https://pinia.vuejs.org/) - Intuitive, type safe store for Vue
- **Styling**: [TailwindCSS](https://tailwindcss.com/) - Utility-first CSS framework
- **Calendar**: [Vue Calendar](https://vcalendar.io/) - Calendar component system
- **Charts**: [Chart.js](https://www.chartjs.org/) with Vue wrapper for analytics
- **HTTP Client**: [Axios](https://axios-http.com/) - Promise based HTTP client

## Prerequisites

Before you begin, ensure you have the following installed:
- [Node.js](https://nodejs.org/) (v16 or higher)
- Package manager of your choice ([npm](https://www.npmjs.com/), [pnpm](https://pnpm.io/), [yarn](https://yarnpkg.com/), or [bun](https://bun.sh/))
- Git for version control

this is a change to check if signed commits work

## Setup

1. Clone the repository:
```bash
git clone https://github.com/your-username/real-quick-frontend.git
cd real-quick-frontend
```

2. Create a `.env` file in the root directory:
```bash
cp .env.example .env
```

3. Install dependencies using your preferred package manager:

```bash
# npm
npm install

# pnpm
pnpm install

# yarn
yarn install

# bun
bun install
```

## Key Features

- Automatic schedule generation around fixed timeblocks
- Customizable preferences for scheduling (weekend hours, task duration limits)
- Task management with tags and timeblocks
- Weather API integration
- Analytics visualization
- Issue reporting with session logs
- Responsive and intuitive user interface

## Development Server

Start the development server on `http://localhost:3000`:

```bash
# npm
npm run dev

# pnpm
pnpm dev

# yarn
yarn dev

# bun
bun run dev
```

## Production

Build the application for production:

```bash
# npm
npm run build

# pnpm
pnpm build

# yarn
yarn build

# bun
bun run build
```

Locally preview production build:

```bash
# npm
npm run preview

# pnpm
pnpm preview

# yarn
yarn preview

# bun
bun run preview
```

## Development Approach

- Structured team with dedicated frontend/backend leads
- Agile development methodology
- Test-driven development
- Enforced branch rules and coding standards
- Docker containerization for deployment
- Robust CI/CD pipeline

## Goals

1. Create an advanced, dynamic scheduling platform
2. Provide responsive and intuitive user experience
3. Ensure secure and reliable data storage
4. Design clean and engaging dashboard
5. Enable cross-session data persistence
6. Integrate third-party calendar synchronization
7. Support multiple schedule view formats
8. Implement comprehensive task management
9. Provide user feedback system
10. Deliver weather-integrated dashboard
11. Offer detailed analytics visualization

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

For more information about deployment, check out the [deployment documentation](https://nuxt.com/docs/getting-started/deployment).
