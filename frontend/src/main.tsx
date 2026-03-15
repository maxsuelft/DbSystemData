import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import utc from 'dayjs/plugin/utc';
import { createRoot } from 'react-dom/client';

import './i18n';
import './index.css';

import App from './App.tsx';

dayjs.extend(utc);
dayjs.extend(relativeTime);

createRoot(document.getElementById('root')!).render(<App />);
