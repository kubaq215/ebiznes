import React, { useState } from 'react';
import Koszyk from '../components/Koszyk';
import Platnosci from '../components/Platnosci';

function CheckoutPage() {
  // Example hardcoded koszykID, replace with dynamic data as necessary
  const koszykCookie = document.cookie
    .split('; ')
    .find(row => row.startsWith('koszykID='))
    ?.split('=')[1];

  const [koszykID, setKoszykID] = useState(koszykCookie || 1);

  return (
    <div className="checkout">
      <h1>Checkout</h1>
      <Koszyk koszykID={koszykID} />  {/* Passing koszykID as a prop */}
      <Platnosci koszykID={koszykID}/>
    </div>
  );
}

export default CheckoutPage;
