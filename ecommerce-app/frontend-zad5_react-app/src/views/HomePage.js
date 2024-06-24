import React from 'react';
import { Link } from 'react-router-dom';
import Produkty from '../components/Produkty.js';

function HomePage() {
  return (
    <div className="home">
      <h1>Welcome to Our Store from Github Actions</h1>
      <Produkty />
      <Link to="/checkout" className="checkout-button">Go to Checkout</Link>
    </div>
  );
}

export default HomePage;
