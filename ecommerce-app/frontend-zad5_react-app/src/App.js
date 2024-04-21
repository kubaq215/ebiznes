import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './views/HomePage';
import CheckoutPage from './views/CheckoutPage';
import './App.css';

function App() {
  return (
    <Router>
      <div className="app">
        <Routes>
          {/* Home Route */}
          <Route path="/" element={<HomePage />} />
          {/* Checkout Route */}
          <Route path="/checkout" element={<CheckoutPage />} />
          {/* You can add more routes here */}
        </Routes>
      </div>
    </Router>
  );
}

export default App;
