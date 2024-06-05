import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('http://localhost:5000/login', {
        email,
        password,
      });
      console.log(response);
      if (response.status === 200) {
        localStorage.setItem('token', response.data.token);
        setMessage('Login successful!');
        // Redirect to home page or another page
        // window.location.href = '/';
      } else {
        setMessage('Login failed.');
      }
      console.log(response);
    } catch (error) {
      setMessage('An error occurred during login.');
    }
  };

  const handleOAuthLogin = () => {
    window.location.href = 'http://localhost:5000/auth/google';
  };

  return (
    <div className="container">
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <div className="input-group">
          <label>Email</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div className="input-group">
          <label>Password</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit" className="btn">Login</button>
      </form>
      <button onClick={handleOAuthLogin} className="btn oauth-btn">Login with Google</button>
      {message && <p>{message}</p>}
    </div>
  );
};

export default Login;
