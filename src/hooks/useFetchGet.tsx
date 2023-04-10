import { useEffect, useState } from 'react';
import axios from 'axios';

export type FetchData = {
  out: any;
  loading: boolean;
  error: any;
};

const useFetchGet = (url: string, token?: string): FetchData => {
  const [out, setOut] = useState<any | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<any | null>(null);

  useEffect(() => {
    setLoading(true);
    const fetchData = async () => {
      try {
        const headers = token ? { Authorization: `Bearer ${token}` } : {};
        const config = { headers };
        const response = await axios.get(url, config);
        setOut(response?.data);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [url]);

  return { out, loading, error };
};

export default useFetchGet;
