import { Button, Input } from 'antd';
import { useTranslation } from 'react-i18next';

import type { Storage } from '../../../../../entity/storages';
import type { StorageOauthDto } from '../../../../../entity/storages/models/StorageOauthDto';

interface Props {
  storage: Storage;
  setStorage: (storage: Storage) => void;
  setUnsaved: () => void;
}

export function EditDropboxStorageComponent({ storage, setStorage, setUnsaved }: Props) {
  const { t } = useTranslation();

  const goToAuthUrl = () => {
    if (!storage?.dropboxStorage?.appKey || !storage?.dropboxStorage?.appSecret) {
      return;
    }

    const redirectUri = `${window.location.origin}/storages/dropbox-oauth`;
    const appKey = storage.dropboxStorage!.appKey;

    const oauthDto: StorageOauthDto = {
      storage: storage,
      authCode: '',
    };

    const authUrl = `https://www.dropbox.com/oauth2/authorize?client_id=${appKey}&redirect_uri=${encodeURIComponent(redirectUri)}&response_type=code&token_access_type=offline&state=${encodeURIComponent(JSON.stringify(oauthDto))}`;

    window.open(authUrl);
  };

  return (
    <>
      <div className="mb-2 flex items-center">
        <div className="hidden min-w-[110px] sm:block" />

        <div className="text-xs text-blue-600">
          <a
            href="https://github.com/dbsystemdata/DbSystemData#readme"
            target="_blank"
            rel="noreferrer"
          >
            Como conectar o Dropbox?
          </a>
        </div>
      </div>

      <div className="mb-1 flex w-full flex-col items-start sm:flex-row sm:items-center">
        <div className="mb-1 min-w-[110px] sm:mb-0">{t('storages.appKey')}</div>
        <Input
          value={storage?.dropboxStorage?.appKey || ''}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
            if (!storage?.dropboxStorage) return;

            setStorage({
              ...storage,
              dropboxStorage: {
                ...storage.dropboxStorage,
                appKey: e.target.value.trim(),
              },
            });
            setUnsaved();
          }}
          size="small"
          className="w-full max-w-[250px]"
          placeholder="app-key"
          disabled={!!storage?.dropboxStorage?.tokenJson}
        />
      </div>

      <div className="mb-1 flex w-full flex-col items-start sm:flex-row sm:items-center">
        <div className="mb-1 min-w-[110px] sm:mb-0">{t('storages.appSecret')}</div>
        <Input
          value={storage?.dropboxStorage?.appSecret || ''}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
            if (!storage?.dropboxStorage) return;

            setStorage({
              ...storage,
              dropboxStorage: {
                ...storage.dropboxStorage,
                appSecret: e.target.value.trim(),
              },
            });
            setUnsaved();
          }}
          size="small"
          className="w-full max-w-[250px]"
          placeholder="app-secret"
          disabled={!!storage?.dropboxStorage?.tokenJson}
        />
      </div>

      {storage?.dropboxStorage?.tokenJson && (
        <div className="mb-1 flex w-full flex-col items-start sm:flex-row sm:items-center">
          <div className="mb-1 min-w-[110px] sm:mb-0">{t('storages.token')}</div>
          <Input value="••••••••" disabled size="small" className="w-full max-w-[250px]" />
        </div>
      )}

      {!storage?.dropboxStorage?.tokenJson && (
        <Button
          type="primary"
          disabled={!storage?.dropboxStorage?.appKey || !storage?.dropboxStorage?.appSecret}
          onClick={goToAuthUrl}
        >
          {t('storages.authorize')}
        </Button>
      )}
    </>
  );
}
