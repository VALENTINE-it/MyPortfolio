const API_BASE = import.meta.env.VITE_API_URL || '/api';

export async function request(endpoint, options = {}) {
  const url = `${API_BASE}${endpoint.startsWith('/') ? endpoint : `/${endpoint}`}`;
  
  const headers = {
    'Content-Type': 'application/json',
    ...(options.headers || {}),
  };

  const config = {
    credentials: options.credentials || 'include',
    ...options,
    headers,
  };

  try {
    const response = await fetch(url, config);
    const data = await response.json().catch(() => null);

    if (!response.ok) {
      const message = data?.message || data?.error?.message || `Request failed with status ${response.status}`;
      const error = new Error(message);
      error.status = response.status;
      error.data = data;
      throw error;
    }

    return data;
  } catch (err) {
    if (!err.status) {
      err.message = 'Unable to reach the server. Please check your network connection.';
    }
    throw err;
  }
}

export default request;
