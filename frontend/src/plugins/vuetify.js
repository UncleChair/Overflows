/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import 'vuetify/styles'
import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'
import MDIAliases from './MDI'

// Composable
import { createVuetify } from 'vuetify'
import { createVueI18nAdapter } from 'vuetify/locale/adapters/vue-i18n'
import { useI18n } from 'vue-i18n'
import i18n from './i18n'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    defaultTheme: localStorage.getItem('theme') ?? 'light',
  },
  locale: {
    adapter: createVueI18nAdapter({ i18n, useI18n })
  },
  icons: {
    defaultSet: 'mdi',
    aliases: {
      ...aliases,
      ...MDIAliases,
    },
    sets: {
      mdi,
    },
  },
})
