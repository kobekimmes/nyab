

import React, { useEffect, useState } from "react";
import { getProducts } from "../api";
import { Product, DataStatus } from "../types";
import ProductGallery from "./ProductGallery";
import "./style/Store.css"

interface StoreProps {
  onAddToCart: (product: Product) => void;
  onRemoveFromCart: (product_id: number) => void;
  isInCart: (product_id: number) => boolean;
}

const Store: React.FC<StoreProps> = ({onAddToCart, onRemoveFromCart, isInCart}) => {
  const [dataStatus, setDataStatus] = useState<DataStatus>("Loading");
  const [error, setError] = useState<Error | null>(null);
  const [products, setProducts] = useState<Product[]>([]);

  useEffect(() => {
    const fetchProducts = async () => {
      setDataStatus("Loading");
      try {
        const result = await getProducts();
        setProducts(result);
        setDataStatus("Success");
      } catch (error) {
        setDataStatus("Failure");
        if (error instanceof Error) {
            setError(error)
        } else {
            setError(new Error("Unexpected/unknown error occured"))
        }
      }
    };

    fetchProducts();
  }, []);


  if (dataStatus === "Loading") return <p>Loading products...</p>;
  if (dataStatus === "Failure") return <p>Failed to load products. Please try again. Error message: ${error?.message}</p>;

  return (
    <div className="store-container">
      <h2>Find something that's uniquely you</h2>
      <ProductGallery
        products={products}
        onAddToCart={onAddToCart}
        onRemoveFromCart={onRemoveFromCart}
        isInCart={isInCart}
      />
    </div>
  );
};

export default Store;
