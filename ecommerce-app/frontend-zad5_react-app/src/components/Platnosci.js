import React, { useState } from 'react';
import PropTypes from 'prop-types'; // Import PropTypes for props validation
import axiosInstance from '../api/axios';
import { useNavigate } from 'react-router-dom'; // Import useNavigate from react-router-dom

function Platnosci({ koszykID }) {
  const navigate = useNavigate(); // Hook to access navigation function
  const [paymentDetails, setPaymentDetails] = useState({
    cardNumber: '',
    expiryDate: '',
    cvv: '',
    amount: '',
    cart_id: koszykID
  });

  const handleChange = (event) => {
    setPaymentDetails({
      ...paymentDetails,
      [event.target.name]: event.target.value
    });
  };

  const handlePaymentSubmission = () => {
    axiosInstance.post('/pay', paymentDetails)
      .then(response => {
        alert('Payment Successful!');
        return axiosInstance.delete(`/koszyk/${koszykID}`);
      })
      .then(() => {
        alert('Cart cleared');
        navigate('/'); // Redirect to the HomePage using navigate
      })
      .catch(error => {
        console.error('Error during payment or clearing cart:', error);
        alert('Payment or Cart clearing failed. Please try again later.');
      });
  };

  return (
    <div className="payment-form">
      <h2>Payment Details</h2>
      <input
        type="text"
        name="cardNumber"
        placeholder="Card Number"
        value={paymentDetails.cardNumber}
        onChange={handleChange}
      />
      <input
        type="text"
        name="expiryDate"
        placeholder="Expiry Date (MM/YY)"
        value={paymentDetails.expiryDate}
        onChange={handleChange}
      />
      <input
        type="text"
        name="cvv"
        placeholder="CVV"
        value={paymentDetails.cvv}
        onChange={handleChange}
      />
      <input
        type="number"
        name="amount"
        placeholder="Amount"
        value={paymentDetails.amount}
        onChange={handleChange}
      />
      <button onClick={handlePaymentSubmission}>Submit Payment</button>
    </div>
  );
}

// Define prop types
Platnosci.propTypes = {
  koszykID: PropTypes.number.isRequired,
};

export default Platnosci;
