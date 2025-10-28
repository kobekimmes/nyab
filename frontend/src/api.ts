

import { Product, Order } from "./types";
import { CheckoutRequest, CheckoutResponse } from "./types";

const BACKEND_DOMAIN = import.meta.env.BE_DOMAIN;

export async function getProducts(): Promise<Product[]> {
    const res = await fetch(`${BACKEND_DOMAIN}/api/products`);
    if (!res.ok) throw new Error("Failed to fetch products");
    return res.json();
}

export async function checkout(order: CheckoutRequest): Promise<CheckoutResponse> {
    const res = await fetch(`${BACKEND_DOMAIN}/api/checkout`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(order)
    });

  if (!res.ok) {
    const errorText = await res.text();
    throw new Error(errorText || "Failed to process checkout");
  }

  return res.json();
}