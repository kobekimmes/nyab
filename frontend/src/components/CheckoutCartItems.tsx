
import React from "react";
import { Product } from "../types";
import "./style/CheckoutCartItem.css";

interface CheckoutCartItemProps {
    product: Product;
    onRemoveFromCart: (productId: number) => void;
}


export const CheckoutCartItem: React.FC<CheckoutCartItemProps> = ({product, onRemoveFromCart}) => {
    return (
        <div className="checkout-cart-item">
            <img src={product.images[0]} alt={product.name} />
            <div className="checkout-cart-item-content">
                <h1>{product.name}</h1>
                <h2>${(product.price * (1 - (product.discount / 100))).toFixed(2)}</h2>
            </div>
            <button
                title="Remove from cart"
                onClick={() => onRemoveFromCart(product.id)}
            >
                &times;
            </button>
        </div>
  );
}

export default CheckoutCartItem;