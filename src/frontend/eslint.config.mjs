import antfu from '@antfu/eslint-config'

// Formats html, css, detects vue and typescript and uses the default rules from the library
// Indent is 2 and quotes are single (although we could change this)
export default antfu({
  typescript: {
    tsconfigPath: 'tsconfig.json',
  },
  formatters: {
    css: true,
    html: true,
  },
  stylistic: {
    indent: 2,
    quotes: 'single',
  },
  vue: true,

  ignores: [ // We might don't want it to check docker
    'docker-compose.yml',
    'Dockerfile',
  ],
})
