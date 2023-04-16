import { useEffect, useState } from 'react';
import axios, { AxiosError } from 'axios';

export type FetchData = {
  out: any;
  loading: boolean;
  error: AxiosError | null;
};

const useFetchGet = (
  url: string,
  token?: string,
  propsChange?: boolean,
): FetchData => {
  const [out, setOut] = useState<any | null>(null);
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
