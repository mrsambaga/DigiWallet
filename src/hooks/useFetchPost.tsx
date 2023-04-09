import { useEffect, useState } from 'react';
import axios, { AxiosError } from 'axios';

export type FetchData = {
  data: any | null;
  loading: boolean;
  error: any | null;
};

const useFetchPost = <T,>(
  url: string,
  body: T,
  submit: boolean,
  toggleSubmit: () => void,
): FetchData => {
  const [data, setData] = useState<any | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<AxiosError | null>(null);

  useEffect(() => {
    if (submit) {
      setLoading(true);
      const fetchData = async () => {
        try {
          const response = await axios.post(url, body);
          setData(response?.data);
          setError(null);
        } catch (error) {
          setError(error as AxiosError);
        } finally {
          setLoading(false);
          toggleSubmit();
        }
      };
      fetchData();
    }
  }, [submit]);

  return { data, loading, error };
};

export default useFetchPost;
