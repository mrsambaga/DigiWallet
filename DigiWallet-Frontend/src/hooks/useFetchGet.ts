import { useEffect, useState } from 'react';
import axios, { AxiosError } from 'axios';

export type FetchData<T> = {
  out: T;
  loading: boolean;
  error: AxiosError | null;
};

const useFetchGet = <T>(
  url: string,
  token?: string,
  propsChange?: boolean,
): FetchData<T> => {
  const [out, setOut] = useState<T>({} as T);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<AxiosError | null>(null);

  useEffect(() => {
    setLoading(true);
    const fetchData = async () => {
      try {
        const headers = token ? { Authorization: `Bearer ${token}` } : {};
        const config = { headers };
        const response = await axios.get(url, config);
        setOut(response?.data);
      } catch (error) {
        const err = error as AxiosError | null;
        setError(err);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [url, propsChange]);

  return { out, loading, error };
};

export default useFetchGet;
