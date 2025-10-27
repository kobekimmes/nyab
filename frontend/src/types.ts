

export interface Product {
    id: number;
    name: string;
    description: string;
    price: number;
    discount: number;
    images: string[];
    sold: boolean;
}

export interface Order {
    id: number;
    totalCost: number;
    productIds: number[];
    firstName: string;
    lastName: string;
    email: string;
    
}

export interface CheckoutRequest {
    productIds: number[];
    firstName: string;
    lastName: string;
    email: string;
}

export interface CheckoutResponse {
    orderId: number;
    totalCost: number;
    message: string;
}

export type DataStatus = "Loading" |"Success" | "Failure"

