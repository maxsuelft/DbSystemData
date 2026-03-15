import type { Storage } from '../../../../../entity/storages';

interface Props {
  storage: Storage;
}

export function ShowDropboxStorageComponent({ storage }: Props) {
  return (
    <div className="mb-1">
      <div className="min-w-[110px]">App Key</div>
      <div className="text-gray-600 dark:text-gray-400">
        {storage?.dropboxStorage?.appKey ? `${storage.dropboxStorage.appKey.slice(0, 8)}...` : '-'}
      </div>
    </div>
  );
}
