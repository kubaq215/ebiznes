import React, { useEffect, useState } from 'react';
import axiosInstance from '../api/axios'; // Import the customized axios instance

function Produkty() {
  const [products, setProducts] = useState([]);
  const [cartId, setCartId] = useState(null); // State to store the current cart ID

  useEffect(() => {
    axiosInstance.get('/produkty')
      .then(response => {
        setProducts(response.data);
      })
      .catch(error => console.error('Error fetching products:', error));
  }, []);

  const addToCart = (produktId) => {
    if (!cartId) {
      // If it's the first product, create a new cart
      axiosInstance.post('/newkoszyk')
        .then(response => {
          const newCartId = response.data.ID; // Assuming the response includes the new cart ID
          setCartId(newCartId);
          document.cookie = `koszykID=${newCartId}; path=/; max-age=86400`; // 86400 seconds = 1 day
          return axiosInstance.post(`/koszyk/${newCartId}/${produktId}`);
        })
        .then(() => alert('Product added to new cart'))
        .catch(error => console.error('Error adding product to new cart:', error));
    } else {
      // If it's not the first product, use the existing cart ID
      axiosInstance.post(`/koszyk/${cartId}/${produktId}`)
        .then(() => alert('Product added to existing cart'))
        .catch(error => console.error('Error adding product to cart:', error));
    }
  };

  return (
    <div>
      {products.map(product => (
        <div key={product.ID} className="product-item">
          <h3>{product.Nazwa}</h3>
          <p><strong>Price:</strong> ${product.Cena.toFixed(2)}</p>
          <p><strong>Description:</strong> {product.Opis}</p>
          <p><strong>Category ID:</strong> {product.KategoriaID}</p>
          <button onClick={() => addToCart(product.ID)}>Add to Cart</button>
        </div>
      ))}
    </div>
  );
}

export default Produkty;
