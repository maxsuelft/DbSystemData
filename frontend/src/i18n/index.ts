import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import ptBR from './locales/pt-BR.json';

const defaultNS = 'common';

i18n.use(initReactI18next).init({
  resources: {
    'pt-BR': { [defaultNS]: ptBR },
  },
  lng: 'pt-BR',
  fallbackLng: 'pt-BR',
  defaultNS,
  interpolation: {
    escapeValue: false,
  },
});

export default i18n;
