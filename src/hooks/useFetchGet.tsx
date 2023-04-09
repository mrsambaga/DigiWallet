import { useEffect, useState } from 'react';
import axios from 'axios';

export type FetchData = {
  data: any;
  loading: boolean;
  error: any;
};

const useFetchGet = (url: string, token?: string): FetchData => {
  const [data, setData] = useState<any | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<any | null>(null);

  useEffect(() => {
    setLoading(true);
    const fetchData = async () => {
      try {
        const headers = token ? { Authorization: `Bearer ${token}` } : {};
        const config = { headers };
        const response = await axios.get(url, config);
        setData(response?.data);
      } catch (error) {
        setError(error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [url]);

  return { data, loading, error };
};

export default useFetchGet;
