import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';

const App = () => {
  const history = useHistory();
  const [token, setToken] = useState(null);

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    if (token) {
      localStorage.setItem('token', token);
      setToken(token);
      history.push('/');
    }
  }, [history]);

  const handleLogin = () => {
    window.location.href = 'http://localhost:5000/auth/google';
  };

  return (
    <div>
      <h1>OAuth2 Login</h1>
      {!token ? (
        <button onClick={handleLogin}>Login with Google</button>
      ) : (
        <p>Logged in with token: {token}</p>
      )}
    </div>
  );
};

export default App;
