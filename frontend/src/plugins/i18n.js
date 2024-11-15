import { createI18n } from 'vue-i18n'
import zhHans from '@/locales/zhHans'
import en from '@/locales/en'

const messages = {
  en: { ...en, $vuetify: 'en' },
  zhHans: { ...zhHans, $vuetify: 'zhHans' },
}

export default createI18n({
  legacy: false, // Vuetify does not support the legacy mode of vue-i18n
  locale: localStorage.getItem('locale') ?? 'zhHans',
  fallbackLocale: 'zhHans',
  messages,
  missingWarn: false
})
