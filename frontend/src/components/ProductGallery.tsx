

// src/components/ProductGallery.tsx
import React, { useState } from "react";
import { Product } from "../types";
import ProductCard from "./ProductCard";
import ProductModal from "./ProductModal";
import "./style/ProductGallery.css"

interface ProductGalleryProps {
  products: Product[];
  onAddToCart: (product: Product) => void;
  onRemoveFromCart: (productId: number) => void;
  isInCart: (productId: number) => boolean;
}

const ProductGallery: React.FC<ProductGalleryProps> = ({ products, onAddToCart, onRemoveFromCart, isInCart }) => {
  const [selected, setSelected] = useState<Product | null>(null);

  const closeModal = () => setSelected(null);

  return (
    <>
        <div className="product-gallery">
            {products.length > 0 ? (
                products.map((p, i) => (
                <ProductCard key={i} product={p} onClick={setSelected} />
                ))
            ) : (
                <div className="product-gallery-empty">No products available</div>
            )}
        </div>

        {selected && (
            <ProductModal 
                product={selected} 
                onClose={closeModal}
                onAddToCart={onAddToCart}
                onRemoveFromCart={onRemoveFromCart}
                inCart={isInCart(selected.id)}
            />
        )}
    </>
  );
};

export default ProductGallery;
