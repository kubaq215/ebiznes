import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
  return (
    <div className="container">
      <h1>Welcome to OAuth2 Demo</h1>
      <Link to="/login">
        <button className="btn">Login</button>
      </Link>
      <Link to="/register">
        <button className="btn">Register</button>
      </Link>
    </div>
  );
};

export default Home;
