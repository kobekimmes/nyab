

// src/components/ProductModal.tsx
import React, { useState } from "react";
import { Product } from "../types"
import "./style/ProductModal.css"

interface ProductModalProps {
  product: Product;
  onClose: () => void;
  inCart: boolean;
  onAddToCart: (product: Product) => void;
  onRemoveFromCart: (productId: number) => void
}

const ProductModal: React.FC<ProductModalProps> = ({ product, onClose, inCart, onAddToCart, onRemoveFromCart }) => {
  const [currentImage, setCurrentImage] = useState(0);

  const nextImage = () => setCurrentImage((currentImage + 1) % product.images.length);
  const prevImage = () => setCurrentImage((currentImage - 1 + product.images.length) % product.images.length);

  return (
    <div className="modal-overlay">
        <div className="modal-container">
            <button className="modal-close" onClick={onClose}>
                &times;
            </button>

            <div className="modal-content">
                <div className="modal-image-container">
                    <img src={product.images[currentImage]} alt={product.name} />
                    {product.images.length > 1 && (
                        <div className="image-nav">
                            <button onClick={prevImage}>&lt;</button>
                            <button onClick={nextImage}>&gt;</button>
                        </div>
                    )}
                </div>

                <div className="modal-details">
                    <h2>{product.name}</h2>
                    <p>{product.description}</p>
                    <p>${(product.price * (1 - product.discount / 100)).toFixed(2)}</p>

                    {inCart ? (
                        <button onClick={() => onRemoveFromCart(product.id)}>
                            Remove from Cart
                        </button>
                    ) : (
                        <button onClick={() => onAddToCart(product)}>
                            Add to Cart
                        </button>
                    )}
                </div>
            </div>
        </div>
    </div>
  );
};

export default ProductModal;
