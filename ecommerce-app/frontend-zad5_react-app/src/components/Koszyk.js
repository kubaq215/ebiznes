import React, { useEffect, useState } from 'react';
import axiosInstance from '../api/axios'; // Import the customized axios instance

function Koszyk({koszykID}) {
  const [cartItems, setCartItems] = useState([]);
  const [totalPrice, setTotalPrice] = useState(0);

  // Fetch cart items from the server
  useEffect(() => {
    const fetchCartItems = async () => {
      try {
        const response = await axiosInstance.get(`/koszyk/${koszykID}`); // Replace '12' with dynamic cart ID if necessary
        setCartItems(response.data);
        calculateTotal(response.data);
      } catch (error) {
        console.error('Error fetching cart items:', error);
      }
    };

    if (koszykID) {
      fetchCartItems();
    }
  }, [koszykID]); // This effect runs whenever koszykID changes

  // Function to calculate total price
  const calculateTotal = (items) => {
    const total = items.reduce((acc, item) => acc + item.Cena, 0);
    setTotalPrice(total);
  };

  return (
    <div>
      <h2>Your Cart</h2>
      {cartItems.length > 0 ? (
        <ul>
          {cartItems.map(item => (
            <li key={item.ID}>
              {item.Nazwa} - ${item.Cena.toFixed(2)} <br />
              Description: {item.Opis}
            </li>
          ))}
        </ul>
      ) : (
        <p>Your cart is empty.</p>
      )}
      <h3>Total: ${totalPrice.toFixed(2)}</h3>
    </div>
  );
}

export default Koszyk;
