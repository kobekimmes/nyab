import React, { useEffect } from "react";
import { Product } from "../types";
import CheckoutCartItem from "./CheckoutCartItems";
import "./style/CheckoutCart.css"

interface CheckoutCartProps {
  isOpen: boolean;
  onClose: () => void;
  cartItems: Product[];
  onRemoveFromCart: (productId: number) => void;
}

const CheckoutCart: React.FC<CheckoutCartProps> = ({ isOpen, onClose, cartItems, onRemoveFromCart }) => {


  // useEffect(() => {

  // }, )

  const subtotal = cartItems.reduce((sum, product) => sum + (product.price * (1 - (product.discount / 100))), 0)

  return (
  <>
    <div className={`cart-overlay ${isOpen ? "opacity-100 visible" : "opacity-0 invisible"}`} onClick={onClose}/>

    <div className={`cart-drawer ${isOpen ? "open" : ""}`}>
    
      <div className="cart-header">
        <h2>Your Cart</h2>
        <button onClick={onClose}>&times;</button>
      </div>

      <div className="cart-items">
        {cartItems.length < 1 ? (
          <p className="cart-empty">Your cart is empty (for now 😅)</p>
        ) : (
        <>
          {cartItems.map((product, index) => (
          <CheckoutCartItem
          key={index}
          product={product}
          onRemoveFromCart={onRemoveFromCart}
          />
          ))}
        </>
        )}
      </div>

      {cartItems.length > 0 && (
        <div className="cart-subtotal">
          <strong>Subtotal:</strong> ${subtotal.toFixed(2)}
        </div>
      )}

      <div className="cart-footer">
        <button>Checkout</button>
      </div>
    </div>
  </>

  );
};

export default CheckoutCart;
