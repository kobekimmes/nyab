

import React from "react";
import { Product } from "../types";
import "./style/ProductCard.css"

interface ProductCardProps {
  product: Product;
  onClick: (product: Product) => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onClick }) => {
    return (
        <div className="product-card" onClick={() => onClick(product)} role="button" tabIndex={0}>
            {product.sold && <div className="sold-badge">Sold</div>}
            <img src={product.images[0]} alt={product.name}/>
            <h3>{product.name}</h3>
            <p>
                <span className="original-price">${product.price.toFixed(2)}</span>
                <span className="sale-price">${(product.price * (1 - (product.discount / 100.0))).toFixed(2)}</span>
            </p>
        </div>
    );
};

export default ProductCard;